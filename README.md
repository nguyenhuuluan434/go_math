# go_math

### This repo used for learning golang module

### 1. create module
go mod init github.com/nguyenhuuluan434/go_math

### 2. create version

#### 2.1  commit
git add .

git commit -m "init version 0.0.1"

####2.2 create tag what is version
git tag -a v0.0.1 -m "initial version"

git push original master --tags

### 3. patch version

#### 3.1  commit
git add .

git commit -m "patch the Div function to handle divide by zero"

####3.2 create tag what is version
git tag -a v0.0.2 -m "patch the Div function to handle divide by zero"

git push original master --tags
