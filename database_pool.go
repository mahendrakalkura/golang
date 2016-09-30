package main

import "database/sql"
import "os"
import "strconv"
import "sync"
import _ "github.com/lib/pq"

var wg sync.WaitGroup

func consumer(dsn string, numbers <-chan int) {
	defer wg.Done()

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	for number := range numbers {
		transaction, err := database.Begin()
		if err != nil {
			panic(err)
		}

		_, err = transaction.Exec(
			`INSERT INTO records (alpha, beta, gamma, delta, epsilon)
            VALUES ($1, $2, $3, $4, $5)`,
			number,
			number,
			number,
			number,
			number,
		)
		if err == nil {
			transaction.Commit()
		} else {
			transaction.Rollback()
		}
	}

	err = database.Close()
	if err != nil {
		panic(err)
	}
}

func main() {
	arguments := os.Args[1:]

	dsn := arguments[0]

	consumers, err := strconv.Atoi(arguments[1])
	if err != nil {
		panic(err)
	}

	records, err := strconv.Atoi(arguments[2])
	if err != nil {
		panic(err)
	}

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	_, err = database.Exec("DROP SCHEMA IF EXISTS public CASCADE")
	if err != nil {
		panic(err)
	}

	_, err = database.Exec("CREATE SCHEMA IF NOT EXISTS public")
	if err != nil {
		panic(err)
	}

	_, err = database.Exec(
		`
        CREATE TABLE IF NOT EXISTS records
            (
                id INTEGER NOT NULL,
                alpha CHARACTER VARYING(255) NOT NULL,
                beta CHARACTER VARYING(255) NOT NULL,
                gamma CHARACTER VARYING(255) NOT NULL,
                delta CHARACTER VARYING(255) NOT NULL,
                epsilon CHARACTER VARYING(255) NOT NULL
            )
        `,
	)
	if err != nil {
		panic(err)
	}

	_, err = database.Exec(
		`
        CREATE SEQUENCE records_id_sequence
            START WITH 1
            INCREMENT BY 1
            NO MINVALUE
            NO MAXVALUE
            CACHE 1
        `,
	)
	if err != nil {
		panic(err)
	}

	_, err = database.Exec(
		`
        ALTER TABLE records
            ALTER COLUMN id
            SET DEFAULT nextval('records_id_sequence'::regclass)
        `,
	)
	if err != nil {
		panic(err)
	}

	_, err = database.Exec(
		`
        ALTER TABLE records
            ADD CONSTRAINT records_id_constraint
            PRIMARY KEY (id)
        `,
	)
	if err != nil {
		panic(err)
	}

	_, err = database.Exec(
		"CREATE INDEX records_alpha ON records USING btree (alpha)",
	)
	if err != nil {
		panic(err)
	}

	_, err = database.Exec(
		"CREATE INDEX records_beta ON records USING btree (beta)",
	)
	if err != nil {
		panic(err)
	}

	_, err = database.Exec(
		"CREATE INDEX records_gamma ON records USING btree (gamma)",
	)
	if err != nil {
		panic(err)
	}

	_, err = database.Exec(
		"CREATE INDEX records_delta ON records USING btree (delta)",
	)
	if err != nil {
		panic(err)
	}

	_, err = database.Exec(
		"CREATE INDEX records_epsilon ON records USING btree (epsilon)",
	)
	if err != nil {
		panic(err)
	}

	numbers := make(chan int, records)

	for index := 1; index <= consumers; index++ {
		wg.Add(1)
		go consumer(dsn, numbers)
	}

	for number := 1; number <= records; number++ {
		numbers <- number
	}
	close(numbers)

	wg.Wait()
}
