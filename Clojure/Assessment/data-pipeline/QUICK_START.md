# Quick Start Guide

## Run the Application

```bash
cd Clojure/Assessment/data-pipeline
lein run
```

## Quick Demo (Recommended Flow)

### Step 1: Setup
```
Enter choice: 1  (Initialize database)
Enter choice: 2  (Add sample students)
```

### Step 2: View Data
```
Enter choice: 3  (View all students)
```

### Step 3: See Data Processing
```
Enter choice: 4  (Process data - shows map, filter, reduce, transducers)
```

### Step 4: Statistics
```
Enter choice: 5  (Show statistics)
```

### Step 5: Start Web Server
```
Enter choice: 6  (Start web server on port 3000)
```

## Test API Endpoints

After starting the web server (choice 6), open a new terminal:

```bash
# Get all students
curl http://localhost:3000/students

# Get high performers (grade >= 85)
curl http://localhost:3000/students/high-performers

# Get statistics
curl http://localhost:3000/students/statistics

# Get students by subject
curl http://localhost:3000/students/by-subject/Mathematics

# Get top 3 students
curl http://localhost:3000/students/top/3

# Process all data (comprehensive)
curl http://localhost:3000/process/all
```

## What You'll See

### Choice 4 Output (Data Processing Demo):
- **FILTER**: High performers with grades >= 85
- **MAP**: Adding letter grades (A, B, C, D, F)
- **REDUCE**: Calculating statistics (average, min, max)
- **GROUP-BY**: Students grouped by subject
- **TRANSDUCER**: Efficient processing pipeline
- **JAVA INTEROP**: Formatted grades using Java DecimalFormat

### Sample Statistics Output:
```
Total Students: 8
Average Grade: 86.00
Highest Grade: 95.5
Lowest Grade: 72.0

Subject-wise Statistics:
  Mathematics:
    Students: 2
    Average: 78.75
  Computer Science:
    Students: 2
    Average: 93.75
```

## Requirements Demonstrated

✅ **Sequence Functions**: map, filter, reduce  
✅ **Transducers**: Efficient data processing  
✅ **Ring HTTP Endpoints**: REST API  
✅ **Database CRUD**: SQLite with error handling  
✅ **Java Interop**: DecimalFormat, ArrayList  

Enjoy exploring the data pipeline! 🚀
