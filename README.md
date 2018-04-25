# students-csv

[![Build Status](https://api.travis-ci.org/crazcalm/students-csv.svg?branch=master)](https://travis-ci.org/crazcalm/flash-cards)[![Go Report Card](https://goreportcard.com/badge/github.com/crazcalm/students-csv)](https://goreportcard.com/report/github.com/crazcalm/students-csv) [![Coverage Status](https://coveralls.io/repos/github/crazcalm/students-csv/badge.svg?branch=master)](https://coveralls.io/github/crazcalm/students-csv?branch=master)

## Purpose

This application holds generic tools I used while teaching English in China. Tools include:

- Taking attendance
- Breaking class into groups
- Selecting a random student

## How it works
### Manage your students with a csv file.
CSV heads must include:

- chinese_name
- pinyin
- english_name
- student_id

## Interface
	Usage of studentsCsv:
      -a    Take attendance
      -f string
            file: path to csv file
      -g int
            Groups the students
      -o string
            The directory where the files will be created (default ".")
      -r    Prints random student
      -s    Shuffle the students