#!/bin/bash

# Define the log file
LOG_FILE="smoke-test.log"

# Run the smoke-test.sh script and write the output to the log file
./smoke-test.sh > "$LOG_FILE" 2>&1

# Print a message indicating that the smoke test has completed
echo "Smoke test completed. Check the log file: $LOG_FILE"