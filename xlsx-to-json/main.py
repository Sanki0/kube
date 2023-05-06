import pandas as pd
import os
import sys

print("sys.path", sys.path[0])

filename = os.environ['FILENAME']

filePath = "/app/data/IN/"+filename

print("filePath", filePath)

# Read Excel file
df = pd.read_excel(filePath)


# Convert dataframe to JSON
json_data = df.to_json(orient='records')

print(json_data)

# Write JSON to file
with open('/app/data/OUT/output.json', 'w') as f:
    f.write(json_data)
