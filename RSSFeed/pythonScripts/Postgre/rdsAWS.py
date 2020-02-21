import psycopg2 as ps
import json
import sys

#Overkill
conn = ps.connect(host=credentials['POSTGRES_ADDRESS'],
              database=credentials['POSTGRES_DBNAME'],
              user=credentials['POSTGRES_USERNAME'],
              password=credentials['POSTGRES_PASSWORD'],
              port=credentials['POSTGRES_PORT'])

cur = conn.cursor()

#Sudo code
with open(date) as f:
	jsonInfo = json.load(f)

columns = [list(x.keys()) for x in record_list][0]
tb_name 