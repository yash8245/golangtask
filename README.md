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

Part 2: Retrieve a task
Endpoint: GET /tasks/{id} Accepts a task ID as a parameter. Retrieves the
corresponding task from the database. Returns the task details if found, or an
appropriate error message if not found.

# Solution for Part 2
<img width="960" alt="Screenshot 2024-02-25 181347" src="https://github.com/yash8245/golangtask/assets/135498759/b3657c8d-c3c5-492d-977f-e38ffd9117ed">

Part 3: Update a task
Endpoint: PUT /tasks/{id} Accepts a task ID as a parameter. Accepts a JSON payload
containing the updated task details (title, description, due date). Updates the
corresponding task in the database. Returns the updated task if successful, or an
appropriate error message if not found.

# Solution for Part 3
First status is setted as pending
<img width="960" alt="Screenshot 2024-02-25 181648" src="https://github.com/yash8245/golangtask/assets/135498759/b2c5122d-4dbd-422e-9723-063417786b33">

After using PUT Method status updated to complete
<img width="960" alt="Screenshot 2024-02-25 181700" src="https://github.com/yash8245/golangtask/assets/135498759/b2a385cb-5148-4434-994f-c3d467aca65b">

Part 4: Delete a task
Endpoint: DELETE /tasks/{id} Accepts a task ID as a parameter. Deletes the
corresponding task from the database. Returns a success message if the deletion is
successful, or an appropriate error message if not found.

# Solution for Part 4
<img width="960" alt="Screenshot 2024-02-25 182010" src="https://github.com/yash8245/golangtask/assets/135498759/ddb7df42-61af-4b90-a39d-04c91888f75d">

Part 5: List all tasks
Endpoint: GET /tasks Retrieves all tasks from the database. Returns a list of tasks,
including their details (title, description, due date).

# Solution for Part 5
<img width="960" alt="Screenshot 2024-02-25 182225" src="https://github.com/yash8245/golangtask/assets/135498759/92e25385-a650-4255-a58b-1b9b3f00ced9">

