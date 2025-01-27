# Use an appropriate base image
FROM alpine:latest

# Install curl
RUN apk add --no-cache curl bash

# Set the working directory inside the container
WORKDIR /app

# Copy the agent.sh and smoke-test.sh scripts to the container
COPY agent.sh smoke-test.sh ./

# Make the scripts executable
RUN chmod +x agent.sh smoke-test.sh

# Run the agent.sh script
CMD ["./agent.sh"]