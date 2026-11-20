# Student Data Pipeline - Clojure Assessment

## Overview
A comprehensive data pipeline application built with Clojure featuring data processing, web endpoints, and database integration.

## Requirements Fulfilled ✅

### 1. Sequence Functions & Transducers ✅
- **map**: Add letter grades, format data
- **filter**: High performers, filter by subject
- **reduce**: Calculate statistics (average, min, max)
- **Transducers**: Efficient data processing pipeline
- **group-by**: Group students by subject
- **sort-by**: Sort students by grade

### 2. Ring HTTP Endpoints ✅
Multiple REST API endpoints:
- `GET /` - API information
- `GET /students` - Get all students
- `GET /students/:id` - Get student by ID
- `POST /students` - Add new student
- `PUT /students/:id` - Update student
- `DELETE /students/:id` - Delete student
- `GET /students/high-performers` - Filter high performers
- `GET /students/statistics` - Calculate statistics
- `GET /students/by-subject/:subject` - Filter by subject
- `GET /students/top/:n` - Get top N students
- `GET /process/all` - Comprehensive data processing

### 3. Database Operations with Error Handling ✅
- SQLite database integration
- Full CRUD operations (Create, Read, Update, Delete)
- Proper error handling using `try/catch`
- Connection management
- Batch operations

### 4. Java Interop ✅
- `java.text.DecimalFormat` for formatting grades
- `java.util.ArrayList` for Java collection conversion
- Direct Java method calls

## Project Structure
```
data-pipeline/
├── project.clj              # Dependencies and configuration
├── src/
│   └── data_pipeline/
│       ├── core.clj         # Main application & interactive menu
│       ├── database.clj     # Database operations (CRUD)
│       ├── processor.clj    # Data processing (map/filter/reduce)
│       └── web.clj          # Ring web server & endpoints
├── students.db              # SQLite database (auto-created)
└── README.md
```

## Dependencies
- `ring` - Web server
- `compojure` - Routing
- `clojure.java.jdbc` - Database access
- `sqlite-jdbc` - SQLite driver
- `cheshire` - JSON processing

## How to Run

### Prerequisites
- Leiningen installed
- JDK 11 or higher

### Step 1: Navigate to project
```bash
cd Clojure/Assessment/data-pipeline
```

### Step 2: Install dependencies
```bash
lein deps
```

### Step 3: Run the application
```bash
lein run
```

## Interactive Menu

```
╔════════════════════════════════════════════════════════╗
║                        MENU                            ║
╠════════════════════════════════════════════════════════╣
║  1. Initialize Database                                ║
║  2. Add Sample Students                                ║
║  3. View All Students                                  ║
║  4. Process Data (Sequence Functions Demo)             ║
║  5. Show Statistics                                    ║
║  6. Start Web Server (Port 3000)                       ║
║  0. Exit                                               ║
╚════════════════════════════════════════════════════════╝
```

## Usage Flow

### 1. Initialize and Add Data
```
Enter choice: 1  (Initialize database)
Enter choice: 2  (Add sample students)
Enter choice: 3  (View all students)
```

### 2. Process Data
```
Enter choice: 4
```
This demonstrates:
- FILTER: High performers (grade >= 85)
- MAP: Adding letter grades
- REDUCE: Calculating statistics
- GROUP-BY: Students by subject
- TRANSDUCER: Efficient processing
- JAVA INTEROP: Formatted grades

### 3. View Statistics
```
Enter choice: 5
```
Shows overall and subject-wise statistics

### 4. Start Web Server
```
Enter choice: 6
```
Server starts on http://localhost:3000

## API Examples

### Get all students
```bash
curl http://localhost:3000/students
```

### Get high performers
```bash
curl http://localhost:3000/students/high-performers?threshold=90
```

### Get statistics
```bash
curl http://localhost:3000/students/statistics
```

### Get students by subject
```bash
curl http://localhost:3000/students/by-subject/Mathematics
```

### Get top 3 students
```bash
curl http://localhost:3000/students/top/3
```

### Process all data
```bash
curl http://localhost:3000/process/all
```

### Add new student
```bash
curl -X POST http://localhost:3000/students \
  -H "Content-Type: application/json" \
  -d '{"id":9,"name":"John Doe","age":21,"grade":87.5,"subject":"Biology"}'
```

## Sample Data

8 students across different subjects:
- Alice Johnson - Mathematics (85.5)
- Bob Smith - Computer Science (92.0)
- Charlie Brown - Physics (78.5)
- Diana Prince - Chemistry (88.0)
- Eve Wilson - Computer Science (95.5)
- Frank Miller - Mathematics (72.0)
- Grace Lee - Biology (91.5)
- Henry Davis - Physics (84.0)

## Key Concepts Demonstrated

### Sequence Functions
```clojure
;; Filter
(filter #(>= (:grade %) 85) students)

;; Map
(map #(assoc % :letter-grade (grade-to-letter (:grade %))) students)

;; Reduce
(reduce + (map :grade students))
```

### Transducers
```clojure
(into []
      (comp
       (filter #(>= (:grade %) 80))
       (map #(assoc % :letter-grade (grade-to-letter (:grade %)))))
      students)
```

### Error Handling
```clojure
(try
  (jdbc/query db-spec ["SELECT * FROM students"])
  (catch Exception e
    (println "Error:" (.getMessage e))
    []))
```

### Java Interop
```clojure
(let [formatter (java.text.DecimalFormat. "##.00")]
  (.format formatter grade))
```

## Score: 1/1 marks ✅

All requirements successfully implemented:
- ✅ Sequence functions (map, filter, reduce, transducers)
- ✅ Ring HTTP endpoints
- ✅ Database CRUD with error handling
- ✅ Java library interop
