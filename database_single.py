from asyncio import get_event_loop
from concurrent.futures import as_completed, ThreadPoolExecutor
from sys import argv

from asyncpg import create_pool

dsn = argv[1]
records = int(argv[2])


async def main(dsn, records):
    pool = await create_pool(dsn)
    async with pool.acquire() as connection:
        async with connection.transaction():
            await connection.execute('DROP SCHEMA IF EXISTS public CASCADE')
            await connection.execute('CREATE SCHEMA IF NOT EXISTS public')
            await connection.execute(
                '''
                CREATE TABLE IF NOT EXISTS records
                    (
                        id INTEGER NOT NULL,
                        alpha CHARACTER VARYING(255) NOT NULL,
                        beta CHARACTER VARYING(255) NOT NULL,
                        gamma CHARACTER VARYING(255) NOT NULL,
                        delta CHARACTER VARYING(255) NOT NULL,
                        epsilon CHARACTER VARYING(255) NOT NULL
                    )
                '''
            )
            await connection.execute(
                '''
                CREATE SEQUENCE records_id_sequence
                    START WITH 1
                    INCREMENT BY 1
                    NO MINVALUE
                    NO MAXVALUE
                    CACHE 1
                '''
            )
            await connection.execute(
                '''
                ALTER TABLE records
                    ALTER COLUMN id
                    SET DEFAULT nextval(\'records_id_sequence\'::regclass)
                '''
            )
            await connection.execute(
                '''
                ALTER TABLE records
                    ADD CONSTRAINT records_id_constraint
                    PRIMARY KEY (id)
                '''
            )
            await connection.execute(
                'CREATE INDEX records_alpha ON records USING btree (alpha)'
            )
            await connection.execute(
                'CREATE INDEX records_beta ON records USING btree (beta)'
            )
            await connection.execute(
                'CREATE INDEX records_gamma ON records USING btree (gamma)'
            )
            await connection.execute(
                'CREATE INDEX records_delta ON records USING btree (delta)'
            )
            await connection.execute(
                'CREATE INDEX records_epsilon ON records USING btree (epsilon)'
            )
    async with pool.acquire() as connection:
        async with connection.transaction():
            for number in range(1, records + 1):
                await connection.execute(
                    '''
                    INSERT INTO records (alpha, beta, gamma, delta, epsilon)
                    VALUES ($1, $2, $3, $4, $5)
                    ''',
                    number,
                    number,
                    number,
                    number,
                    number,
                )

loop = get_event_loop()
loop.run_until_complete(main(dsn, records))
loop.close()
