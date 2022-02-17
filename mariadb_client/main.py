#!/opt/homebrew/bin/python3

# brew install mariadb-connector-c
# pip3 install mariadb

# Module Import
import mariadb
import sys


key_path = "/Users/stevehuang/.tsh/keys/teleport.dev.aws.stevexin.me/STeve" 
ca_path  ="/Users/stevehuang/.tsh/keys/teleport.dev.aws.stevexin.me/certs.pem"
cert_path= "/Users/stevehuang/.tsh/keys/teleport.dev.aws.stevexin.me/STeve-db/teleport.dev.aws.stevexin.me/self-hosted-mariadb-x509.pem"

try:
    conn = mariadb.connect(
            host="127.0.0.1",
            port=34567,
            user="alice",
            ssl_key=key_path,
            ssl_ca=ca_path,
            ssl_cert=cert_path,
            database="test")
    print(conn.get_server_version())
    print(conn.server_version_info)
    print(conn.server_name)
    print(conn.server_info)

    cur = conn.cursor()

    cur.execute("DELETE FROM users WHERE age < 5")

    cur.execute("INSERT INTO users VALUES (?, ?)", ("macy", 4))

    cur.executemany("INSERT INTO users VALUES (?, ?)", [("miki", 3), ('mini', 2)])

    cur.execute("UPDATE users SET age=? WHERE name=?",(1, "macy"))

    cur.execute("SELECT age from users WHERE name=? and age < ?", ("macy", 100))
    print(cur.fetchall())
    print(cur.fetchall())

    cur.execute("SELECT * from users WHERE age < 100")
    print("--- all users ---")
    print(cur.fetchall())

    cur.execute("CREATE INDEX by_age ON users (age)")
    cur.execute("DROP INDEX by_age ON users (age)")


    conn.close()
except mariadb.Error as e:
    print(f"Error connecting to the database: {e}")
    sys.exit(1)

