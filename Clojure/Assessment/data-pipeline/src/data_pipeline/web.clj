(ns data-pipeline.web
  (:require [ring.adapter.jetty :refer [run-jetty]]
            [ring.middleware.json :refer [wrap-json-response wrap-json-body]]
            [ring.util.response :refer [response status]]
            [compojure.core :refer [defroutes GET POST PUT DELETE]]
            [compojure.route :as route]
            [data-pipeline.database :as db]
            [data-pipeline.processor :as proc]))

;; Helper function to create JSON response
(defn json-response
  ([data] (json-response data 200))
  ([data status-code]
   (-> (response data)
       (status status-code))))

;; Routes
(defroutes app-routes
  ;; Home endpoint
  (GET "/" []
    (json-response
     {:message "Student Data Pipeline API"
      :version "1.0.0"
      :endpoints ["/students"
                  "/students/:id"
                  "/students/high-performers"
                  "/students/statistics"
                  "/students/by-subject/:subject"
                  "/students/top/:n"
                  "/process/all"]}))

  ;; Get all students
  (GET "/students" []
    (try
      (let [students (db/get-all-students)]
        (json-response {:success true :data students :count (count students)}))
      (catch Exception e
        (json-response {:success false :error (.getMessage e)} 500))))

  ;; Get student by ID
  (GET "/students/:id" [id]
    (try
      (if-let [student (db/get-student-by-id (Integer/parseInt id))]
        (json-response {:success true :data student})
        (json-response {:success false :error "Student not found"} 404))
      (catch Exception e
        (json-response {:success false :error (.getMessage e)} 500))))

  ;; Add new student
  (POST "/students" request
    (try
      (let [student (get-in request [:body])]
        (db/add-student! student)
        (json-response {:success true :message "Student added" :data student} 201))
      (catch Exception e
        (json-response {:success false :error (.getMessage e)} 500))))

  ;; Update student
  (PUT "/students/:id" [id :as request]
    (try
      (let [student-data (get-in request [:body])
            result (db/update-student! (Integer/parseInt id) student-data)]
        (json-response result))
      (catch Exception e
        (json-response {:success false :error (.getMessage e)} 500))))

  ;; Delete student
  (DELETE "/students/:id" [id]
    (try
      (let [result (db/delete-student! (Integer/parseInt id))]
        (json-response result))
      (catch Exception e
        (json-response {:success false :error (.getMessage e)} 500))))

  ;; Get high performers
  (GET "/students/high-performers" [threshold]
    (try
      (let [students (db/get-all-students)
            thresh (if threshold (Double/parseDouble threshold) 85)
            high-perf (proc/high-performers students thresh)]
        (json-response {:success true
                        :threshold thresh
                        :data high-perf
                        :count (count high-perf)}))
      (catch Exception e
        (json-response {:success false :error (.getMessage e)} 500))))

  ;; Get statistics
  (GET "/students/statistics" []
    (try
      (let [students (db/get-all-students)
            stats (proc/calculate-statistics students)]
        (json-response {:success true :statistics stats}))
      (catch Exception e
        (json-response {:success false :error (.getMessage e)} 500))))

  ;; Get students by subject
  (GET "/students/by-subject/:subject" [subject]
    (try
      (let [students (db/get-all-students)
            filtered (proc/students-by-subject students subject)]
        (json-response {:success true
                        :subject subject
                        :data filtered
                        :count (count filtered)}))
      (catch Exception e
        (json-response {:success false :error (.getMessage e)} 500))))

  ;; Get top N performers
  (GET "/students/top/:n" [n]
    (try
      (let [students (db/get-all-students)
            num (Integer/parseInt n)
            top (proc/get-top-performers students num)]
        (json-response {:success true
                        :top num
                        :data top}))
      (catch Exception e
        (json-response {:success false :error (.getMessage e)} 500))))

  ;; Process all data (comprehensive endpoint)
  (GET "/process/all" []
    (try
      (let [result (proc/process-all-students)]
        (json-response {:success true :processed result}))
      (catch Exception e
        (json-response {:success false :error (.getMessage e)} 500))))

  ;; Subject-wise statistics
  (GET "/statistics/by-subject" []
    (try
      (let [students (db/get-all-students)
            stats (proc/subject-wise-statistics students)]
        (json-response {:success true :statistics stats}))
      (catch Exception e
        (json-response {:success false :error (.getMessage e)} 500))))

  ;; Default route
  (route/not-found
   (json-response {:success false :error "Route not found"} 404)))

;; Middleware wrapper
(def app
  (-> app-routes
      (wrap-json-body {:keywords? true})
      wrap-json-response))

;; Start server
(defn start-server!
  [port]
  (try
    (println (str "🚀 Starting server on port " port "..."))
    (run-jetty app {:port port :join? false})
    (println (str "✅ Server running at http://localhost:" port))
    (catch Exception e
      (println "❌ Failed to start server:" (.getMessage e)))))
