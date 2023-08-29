package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score float64
}

type Students []Student

func (s Students) CalculateAverageScore() float64 {
	totalScore := 0.0
	for _, student := range s {
		totalScore += student.Score
	}
	averageScore := totalScore / float64(len(s))
	return averageScore
}

func (s Students) FindMinScoreStudent() Student {
	minScore := math.MaxFloat64
	var minScoreStudent Student
	for _, student := range s {
		if student.Score < minScore {
			minScore = student.Score
			minScoreStudent = student
		}
	}
	return minScoreStudent
}

func (s Students) FindMaxScoreStudent() Student {
	maxScore := -math.MaxFloat64
	var maxScoreStudent Student
	for _, student := range s {
		if student.Score > maxScore {
			maxScore = student.Score
			maxScoreStudent = student
		}
	}
	return maxScoreStudent
}

func main() {
	var students Students

	for i := 1; i <= 5; i++ {
		var name string
		var score float64

		fmt.Printf("Input %d Student's Name: ", i)
		fmt.Scan(&name)
		fmt.Printf("Input %d Student's Score: ", i)
		fmt.Scan(&score)

		students = append(students, Student{Name: name, Score: score})
	}

	averageScore := students.CalculateAverageScore()
	minScoreStudent := students.FindMinScoreStudent()
	maxScoreStudent := students.FindMaxScoreStudent()

	fmt.Printf("Average Score: %.2f\n", averageScore)
	fmt.Printf("Min Score of Students: %s (%.0f)\n", minScoreStudent.Name, minScoreStudent.Score)
	fmt.Printf("Max Score of Students: %s (%.0f)\n", maxScoreStudent.Name, maxScoreStudent.Score)
}
