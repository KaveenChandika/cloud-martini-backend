#!/bin/bash

# Define the log file
LOG_FILE="smoke-test.log"

# Run the smoke-test.sh script and write the output to the log file
if ./smoke-test.sh > "$LOG_FILE" 2>&1; then
  # Print a success message if the script succeeds
  echo "Smoke test completed successfully."
  echo "result=success"
else
  # Print a failure message if the script fails
  echo "Smoke test failed. Check the log file: $LOG_FILE"
  echo "result=failure"
fi