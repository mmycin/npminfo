from lib.imports import pl
from lib.data import data
from lib.printer import printer
from lib.types import Data
import os

def export(data: Data, filename: str) -> None:
    flattened_data = []
    for item in data:
        for name, details in item.items():
            description, link = details
            flattened_data.append([name, description, link])
    
    # Create a DataFrame
    df = pl.DataFrame(flattened_data, schema=["Name", "Description", "Link"])
    
    # Ensure the 'data' folder exists
    os.makedirs("data", exist_ok=True)
    
    # Write the DataFrame to a CSV file in the 'data' folder
    csv_file_path = f"data/{filename}"
    df.write_csv(csv_file_path)
    
    print(f"CSV file created successfully at {csv_file_path}.")