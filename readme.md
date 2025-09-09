# Go Top 3 Words Finder 

A simple and efficient **Go program** that reads a text file, counts **word frequencies**, and prints the **top 3 most common words**.  
Uses a **min-heap** for better performance and handles complex text with punctuation, numbers, and mixed cases.

---

## Features 
- Reads files **line by line** using `bufio.Scanner`
- Counts **word frequency** efficiently using a map
- Uses a **min-heap** for faster top 3 extraction
- Handles **punctuation, numbers, and mixed cases**
- Includes a **sample file** (`sample.txt`) for testing

---
## Usage
**1. Clone the repository**
```bash
git clone git@github.com:abhishek5463/test.git
```
---
## Run Program
- go run main.go
---

## Sample Output
```bash
Top 3 most common words:
go: 10
fun: 5
testing: 3
```
---