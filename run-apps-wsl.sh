#!/bin/bash

cd /mnt/c/MAMP/htdocs/Developer/go-htmx-mySQL

# Check if we have gnome-terminal
if command -v gnome-terminal &> /dev/null; then
    gnome-terminal -- bash -c "air -c .air.wsl.toml; exec bash"
    gnome-terminal -- bash -c "air -c .air.admin.wsl.toml; exec bash"
else
    # Alternative approach for terminals without gnome-terminal
    # Start the first app in the background
    echo "Starting main app..."
    air -c .air.wsl.toml &
    APP_PID=$!
    
    # Wait a bit for the first app to start
    sleep 2
    
    # Start the second app in the background
    echo "Starting admin app..."
    air -c .air.admin.wsl.toml &
    ADMIN_PID=$!
    
    echo "Both apps are running. Press Ctrl+C to stop both."
    
    # Wait for user to press Ctrl+C
    trap "kill $APP_PID $ADMIN_PID; exit" INT
    wait
fi
