### Timed Quiz Game

### Installation
- Clone the repository
- Add csv file with the quiz of choice in to src/quizGame/internal
- Run the command in the command line: 
```
go run src/quiz/main.go -time=<seconds>[optional] [CSV file]
```
- Enjoy:)

### CVS file format
Starting from the first line of the file.
Questions must be on the left and the answer on the right. 
Eg: 
    5+5,10
    7+3,10
    1+1,2
    8+3,11
Question/Answers like these are also acceptable: 
Eg:
    "what 2+2, sir?",4
