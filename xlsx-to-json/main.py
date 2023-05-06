import pandas as pd
import os

filename = os.environ['FILENAME']

# Read Excel file
df = pd.read_excel(filename)


# Convert dataframe to JSON
json_data = df.to_json(orient='records')

print(json_data)

# Write JSON to file
with open('/app/data/output.json', 'w') as f:
    f.write(json_data)
