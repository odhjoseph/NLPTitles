import psycopg2 as ps

try: 
    connection = ps.connect(
        database="jsondatabase",
        user="postgres",
        password="ufkS3GfMnclQUCyKVbAN",
        host="jsondatabase.cm1zpldcgrzz.us-east-1.rds.amazonaws.com",
        port='5432'
    )
    cursor = connection.cursor()
    # Print PostgreSQL Connection properties
    print ( connection.get_dsn_parameters(),"\n")

    # Print PostgreSQL version
    cursor.execute("SELECT version();")
    record = cursor.fetchone()
    print("You are connected to - ", record,"\n")

except (Exception, ps.Error) as error :
    print ("Error while connecting to PostgreSQL", error)
finally:
    #closing database connection.
        if(connection):
            cursor.close()
            connection.close()
            print("PostgreSQL connection is closed")




