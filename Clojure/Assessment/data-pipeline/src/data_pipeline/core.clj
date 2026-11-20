(ns data-pipeline.core
  (:require [data-pipeline.database :as db]
            [data-pipeline.processor :as proc]
            [data-pipeline.web :as web])
  (:gen-class))

(defn display-banner []
  (println "==========================================================")
  (println "║     Student Data Pipeline - Clojure Assessment        ║")
  (println "==========================================================")
  (println))

(defn display-menu []
  (println "\n==========================================================")
  (println "║                        MENU                            ║")
  (println "╠════════════════════════════════════════════════════════╣")
  (println "║  1. Initialize Database                                ║")
  (println "║  2. Add Sample Students                                ║")
  (println "║  3. View All Students                                  ║")
  (println "║  4. Process Data (Sequence Functions Demo)             ║")
  (println "║  5. Show Statistics                                    ║")
  (println "║  6. Start Web Server (Port 3000)                       ║")
  (println "║  0. Exit                                               ║")
  (println "=========================================================="))

(defn demo-sequence-functions []
  (println "\n==========================================================")
  (println "DEMO: Sequence Functions (map, filter, reduce)")
  (println "==========================================================")
  
  (let [students (db/get-all-students)]
    (println "\n📊 Original Data Count:" (count students))
    
    ;; FILTER - High performers
    (println "\n🎯 FILTER - High Performers (grade >= 85):")
    (let [high-perf (proc/high-performers students)]
      (doseq [s high-perf]
        (println (format "  - %s: %.1f" (:name s) (:grade s))))
      (println "  Count:" (count high-perf)))
    
    ;; MAP - Add letter grades
    (println "\n📝 MAP - Adding Letter Grades:")
    (let [with-letters (proc/add-letter-grades students)]
      (doseq [s (take 3 with-letters)]
        (println (format "  - %s: %.1f (%s)"
                        (:name s) (:grade s) (:letter-grade s)))))
    
    ;; REDUCE - Calculate average
    (println "\n🔢 REDUCE - Calculate Statistics:")
    (let [stats (proc/calculate-statistics students)]
      (println (format "  Average Grade: %.2f" (:average stats)))
      (println (format "  Min Grade: %.1f" (:min stats)))
      (println (format "  Max Grade: %.1f" (:max stats)))
      (println (format "  Total Students: %d" (:count stats))))
    
    ;; GROUP-BY - By subject
    (println "\n📚 GROUP-BY - Students by Subject:")
    (let [by-subject (proc/group-by-subject students)]
      (doseq [[subject subj-students] by-subject]
        (println (format "  - %s: %d students" subject (count subj-students)))))
    
    ;; TRANSDUCER
    (println "\n⚡ TRANSDUCER - Efficient Processing:")
    (let [result (proc/process-with-transducer students)]
      (println "  Filtered (grade >= 80), mapped, and selected keys")
      (println "  Result count:" (count result))
      (doseq [s (take 3 result)]
        (println (format "  - %s: %.1f (%s)"
                        (:name s) (:grade s) (:letter-grade s)))))
    
    ;; Java Interop
    (println "\n☕ JAVA INTEROP - Using Java DecimalFormat:")
    (doseq [s (take 3 students)]
      (println (format "  - %s: %s"
                      (:name s) (proc/format-grade (:grade s)))))))

(defn display-statistics []
  (println "\==========================================================")
  (println "STATISTICS REPORT")
  (println "==========================================================")
  
  (let [students (db/get-all-students)
        stats (proc/calculate-statistics students)
        by-subject (proc/subject-wise-statistics students)]
    
    (println "\n📊 Overall Statistics:")
    (println (format "  Total Students: %d" (:count stats)))
    (println (format "  Average Grade: %.2f" (:average stats)))
    (println (format "  Highest Grade: %.1f" (:max stats)))
    (println (format "  Lowest Grade: %.1f" (:min stats)))
    
    (println "\n📚 Subject-wise Statistics:")
    (doseq [[subject subj-stats] by-subject]
      (println (format "  %s:" subject))
      (println (format "    Students: %d" (:count subj-stats)))
      (println (format "    Average: %.2f" (:average subj-stats))))))

(defn run-interactive []
  (display-banner)
  (loop []
    (display-menu)
    (print "\nEnter choice: ")
    (flush)
    (let [choice (read-line)]
      (case choice
        "1" (do (db/init-db!)
                (recur))
        
        "2" (do (db/clear-all-students!)
                (db/add-sample-students!)
                (recur))
        
        "3" (do (println "\n=== All Students ===")
                (let [students (db/get-all-students)]
                  (if (empty? students)
                    (println "No students found.")
                    (doseq [s students]
                      (println (format "ID: %d | %s | Age: %d | Grade: %.1f | Subject: %s"
                                      (:id s) (:name s) (:age s) (:grade s) (:subject s))))))
                (recur))
        
        "4" (do (demo-sequence-functions)
                (recur))
        
        "5" (do (display-statistics)
                (recur))
        
        "6" (do (println "\n🚀 Starting Web Server...")
                (web/start-server! 3000)
                (println "\n📡 API Endpoints available:")
                (println "  GET  http://localhost:3000/")
                (println "  GET  http://localhost:3000/students")
                (println "  GET  http://localhost:3000/students/high-performers")
                (println "  GET  http://localhost:3000/students/statistics")
                (println "  GET  http://localhost:3000/students/by-subject/:subject")
                (println "  GET  http://localhost:3000/students/top/:n")
                (println "  GET  http://localhost:3000/process/all")
                (println "\nPress Enter to continue (server runs in background)...")
                (read-line)
                (recur))
        
        "0" (println "\n👋 Goodbye!")
        
        (do (println "❌ Invalid choice. Try again.")
            (recur))))))

(defn -main
  [& args]
  (try
    (run-interactive)
    (catch Exception e
      (println "❌ Error:" (.getMessage e)))))
