(ns data-pipeline.processor
  (:require [data-pipeline.database :as db]))

;; Sequence functions for data processing

;; FILTER - Get students with grade above threshold
(defn high-performers
  "Filter students with grade above threshold (default 85)"
  ([students] (high-performers students 85))
  ([students threshold]
   (filter #(>= (:grade %) threshold) students)))

;; FILTER - Get students by subject
(defn students-by-subject
  "Filter students by subject"
  [students subject]
  (filter #(= (:subject %) subject) students))

;; MAP - Calculate letter grades
(defn grade-to-letter
  "Convert numeric grade to letter grade"
  [grade]
  (cond
    (>= grade 90) "A"
    (>= grade 80) "B"
    (>= grade 70) "C"
    (>= grade 60) "D"
    :else "F"))

(defn add-letter-grades
  "Add letter grade to each student"
  [students]
  (map #(assoc % :letter-grade (grade-to-letter (:grade %))) students))

;; REDUCE - Calculate statistics
(defn calculate-average-grade
  "Calculate average grade using reduce"
  [students]
  (if (empty? students)
    0
    (/ (reduce + (map :grade students))
       (count students))))

(defn calculate-statistics
  "Calculate comprehensive statistics using reduce"
  [students]
  (if (empty? students)
    {:count 0 :average 0 :total 0 :min 0 :max 0}
    (let [grades (map :grade students)
          total (reduce + grades)
          avg (/ total (count students))]
      {:count (count students)
       :average (double avg)
       :total total
       :min (apply min grades)
       :max (apply max grades)})))

;; TRANSDUCER - Efficient data processing
(defn process-with-transducer
  "Process students efficiently using transducers"
  [students]
  (into []
        (comp
         (filter #(>= (:grade %) 80))
         (map #(assoc % :letter-grade (grade-to-letter (:grade %))))
         (map #(select-keys % [:id :name :grade :letter-grade])))
        students))

;; GROUP-BY - Group students by subject
(defn group-by-subject
  "Group students by their subject"
  [students]
  (group-by :subject students))

;; SORT - Sort students by grade
(defn sort-by-grade
  "Sort students by grade (descending)"
  [students]
  (sort-by :grade > students))

;; Complex processing pipeline
(defn process-students-pipeline
  "Apply a complete processing pipeline"
  [students]
  (->> students
       (filter #(>= (:age %) 20))
       (map #(assoc % :letter-grade (grade-to-letter (:grade %))))
       (sort-by :grade >)))

;; Subject-wise statistics
(defn subject-wise-statistics
  "Calculate statistics grouped by subject"
  [students]
  (let [grouped (group-by-subject students)]
    (into {}
          (map (fn [[subject subj-students]]
                 [subject (calculate-statistics subj-students)])
               grouped))))

;; Top performers
(defn get-top-performers
  "Get top N students by grade"
  [students n]
  (->> students
       (sort-by :grade >)
       (take n)))

;; Java Interop - Using Java's DecimalFormat
(defn format-grade
  "Format grade using Java DecimalFormat"
  [grade]
  (let [formatter (java.text.DecimalFormat. "##.00")]
    (.format formatter grade)))

;; Using Java ArrayList (Java Interop example)
(defn students-to-java-list
  "Convert students to Java ArrayList"
  [students]
  (let [java-list (java.util.ArrayList.)]
    (doseq [student students]
      (.add java-list student))
    java-list))

;; Process all data from database
(defn process-all-students
  "Fetch and process all students from database"
  []
  (let [students (db/get-all-students)]
    {:raw-data students
     :high-performers (high-performers students)
     :with-letter-grades (add-letter-grades students)
     :statistics (calculate-statistics students)
     :by-subject (subject-wise-statistics students)
     :top-3 (get-top-performers students 3)
     :transducer-result (process-with-transducer students)}))
