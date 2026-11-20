(ns data-pipeline.database
  (:require [clojure.java.jdbc :as jdbc]))

;; Database configuration
(def db-spec
  {:classname   "org.sqlite.JDBC"
   :subprotocol "sqlite"
   :subname     "students.db"})

;; Initialize database with table
(defn init-db!
  "Create students table if it doesn't exist"
  []
  (try
    (jdbc/execute! db-spec
      ["CREATE TABLE IF NOT EXISTS students (
          id INTEGER PRIMARY KEY,
          name TEXT NOT NULL,
          age INTEGER NOT NULL,
          grade REAL NOT NULL,
          subject TEXT NOT NULL)"])
    (println "✅ Database initialized successfully")
    {:success true :message "Database initialized"}
    (catch Exception e
      (println "❌ Database initialization failed:" (.getMessage e))
      {:success false :error (.getMessage e)})))

;; CREATE - Add student
(defn add-student!
  "Add a new student to the database"
  [student]
  (try
    (jdbc/insert! db-spec :students student)
    (println "✅ Student added:" (:name student))
    {:success true :student student}
    (catch Exception e
      (println "❌ Failed to add student:" (.getMessage e))
      {:success false :error (.getMessage e)})))

;; READ - Get all students
(defn get-all-students
  "Retrieve all students from database"
  []
  (try
    (jdbc/query db-spec ["SELECT * FROM students"])
    (catch Exception e
      (println "❌ Failed to retrieve students:" (.getMessage e))
      [])))

;; READ - Get student by ID
(defn get-student-by-id
  "Retrieve a student by ID"
  [id]
  (try
    (first (jdbc/query db-spec ["SELECT * FROM students WHERE id = ?" id]))
    (catch Exception e
      (println "❌ Failed to retrieve student:" (.getMessage e))
      nil)))

;; UPDATE - Update student
(defn update-student!
  "Update an existing student"
  [id student-data]
  (try
    (jdbc/update! db-spec :students student-data ["id = ?" id])
    (println "✅ Student updated, ID:" id)
    {:success true :id id}
    (catch Exception e
      (println "❌ Failed to update student:" (.getMessage e))
      {:success false :error (.getMessage e)})))

;; DELETE - Remove student
(defn delete-student!
  "Delete a student by ID"
  [id]
  (try
    (jdbc/delete! db-spec :students ["id = ?" id])
    (println "✅ Student deleted, ID:" id)
    {:success true :id id}
    (catch Exception e
      (println "❌ Failed to delete student:" (.getMessage e))
      {:success false :error (.getMessage e)})))

;; Batch insert for sample data
(defn add-sample-students!
  "Add sample student data"
  []
  (try
    (let [students [{:id 1 :name "Alice Johnson" :age 20 :grade 85.5 :subject "Mathematics"}
                    {:id 2 :name "Bob Smith" :age 22 :grade 92.0 :subject "Computer Science"}
                    {:id 3 :name "Charlie Brown" :age 19 :grade 78.5 :subject "Physics"}
                    {:id 4 :name "Diana Prince" :age 21 :grade 88.0 :subject "Chemistry"}
                    {:id 5 :name "Eve Wilson" :age 20 :grade 95.5 :subject "Computer Science"}
                    {:id 6 :name "Frank Miller" :age 23 :grade 72.0 :subject "Mathematics"}
                    {:id 7 :name "Grace Lee" :age 19 :grade 91.5 :subject "Biology"}
                    {:id 8 :name "Henry Davis" :age 22 :grade 84.0 :subject "Physics"}]]
      (doseq [student students]
        (jdbc/insert! db-spec :students student))
      (println "✅ Sample students added successfully")
      {:success true :count (count students)})
    (catch Exception e
      (println "❌ Failed to add sample students:" (.getMessage e))
      {:success false :error (.getMessage e)})))

;; Clear all students
(defn clear-all-students!
  "Remove all students from database"
  []
  (try
    (jdbc/execute! db-spec ["DELETE FROM students"])
    (println "✅ All students cleared")
    {:success true}
    (catch Exception e
      (println "❌ Failed to clear students:" (.getMessage e))
      {:success false :error (.getMessage e)})))
