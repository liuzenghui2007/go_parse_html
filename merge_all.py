import pandas as pd
import os

# Directory containing the CSV files
csv_directory = './results'

# List all CSV files in the directory
csv_files = [f for f in os.listdir(csv_directory) if f.endswith('.csv')]

# Initialize an empty DataFrame
combined_df = pd.DataFrame()

# Loop through each CSV file
for csv_file in csv_files:
    # Read the CSV file
    df = pd.read_csv(os.path.join(csv_directory, csv_file))
    
    # Add a new column with the file name without the extension
    df['Source File'] = os.path.splitext(csv_file)[0]
    
    # Append the DataFrame to the combined DataFrame
    combined_df = pd.concat([combined_df, df], ignore_index=True)

# Save the combined DataFrame to an Excel file
combined_df.to_excel('./result.xlsx', index=False)