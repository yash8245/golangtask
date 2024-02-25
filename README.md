# SECTION A 

The SQLite database schema for storing tasks will include the following fields:
• ID: The unique identifier for each task.
• Title: The title or name of the task.
• Description: A brief description of the task.
• Due Date: The deadline or due date for the task.
• Status: The status of the task (e.g., "pending," "completed," "in progress," etc.).


SOLUTION FOR THE ABOVE TASK 

# Schema created for the task 
<img width="960" alt="Screenshot 2024-02-24 185818" src="https://github.com/yash8245/golangtask/assets/135498759/1a52d670-ec82-4c8b-96d5-0c09819fcfb0">

# Table created 
<img width="960" alt="image" src="https://github.com/yash8245/golangtask/assets/135498759/1b3cb1a0-9aa1-4b5a-9bb5-605e19cfc246">

# SECTION B

Part 1: Create a new task
Endpoint: POST /tasks Accepts a JSON payload containing the task details (title,
description, due date). Generates a unique ID for the task and stores it in the database.
Returns the created task with the assigned ID.

# Solution for Part 1
<img width="960" alt="Screenshot 2024-02-25 180853" src="https://github.com/yash8245/golangtask/assets/135498759/a9d57f92-f472-4c27-93fc-e06683e4ffc5">
