How to install?
===============

```
$ mkdir golang_vs_python
$ cd golang_vs_python
$ git clone --recursive git@github.com:mahendrakalkura/golang_vs_python.git .
$ mkvirtualenv --python=python3 golang_vs_python
$ pip install -r requirements.txt
$ go get github.com/lib/pq"
$ go get gopkg.in/cheggaaa/pb.v1
$ psql -c 'CREATE DATABASE golang_vs_python' -d postgres
$ deactivate
```

How to run?
===========

```
$ cd golang_vs_python
$ workon golang_vs_python
$ time go run database_pool.go 'host=0.0.0.0 port=5432 user={username} password={password} dbname=golang_vs_python sslmode=disable' 9 99999
$ time python database_pool.py postgresql://{username}:{password}@0.0.0.0:5432/golang_vs_python 9 99999
$ time go run database_single.go 'host=0.0.0.0 port=5432 user={username} password={password} dbname=golang_vs_python sslmode=disable' 99999
$ time python database_single.py postgresql://{username}:{password}@0.0.0.0:5432/golang_vs_python 99999
$ time go run net_http.go 9 99
$ time python net_http.py 9 99
$ time python requirements.txt
$ time go run total_pool.go 99 99999
$ time python total_pool.py 99 99999
$ time go run total_single.go 9999999
$ time python total_single.py 9999999
$ deactivate
```
