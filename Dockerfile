# Use Alpine as the base image
FROM alpine:latest
RUN apk add --no-cache libc6-compat
# Copy the binary into the container's /app directory
COPY ./build/go_build_ZooDaBa /app/go_build_ZooDaBa

# Set the working directory
WORKDIR /app

# Make the binary executable
RUN chmod +x /app/go_build_ZooDaBa

# Expose port 8090
EXPOSE 8090

# Run the binary
CMD ["/app/go_build_ZooDaBa"]