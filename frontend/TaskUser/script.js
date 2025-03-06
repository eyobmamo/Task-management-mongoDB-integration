document.addEventListener('DOMContentLoaded', async () => {
    const token = localStorage.getItem('token');

    // Check if the user is logged in
    if (!token) {
        alert('You are not logged in. Redirecting to login page...');
        window.location.href = '../user_authentication/login.html'; // Redirect to login page
        return;
    }

    // Logout button
    const logoutBtn = document.getElementById('logout-btn');
    logoutBtn.addEventListener('click', () => {
        localStorage.removeItem('token');
        window.location.href = './user_authentication/login.html';
    });

    try {
        // Fetch tasks from the backend
        const response = await fetch('http://localhost:8081/tasks', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`,
            },
        });

        const result = await response.json();
        if (response.ok) {
            renderTasks(result.tasks); // Render tasks
        } else {
            alert(`Error: ${result.message || 'Failed to fetch tasks'}`);
        }
    } catch (error) {
        console.error('Error:', error);
        alert('An error occurred. Please try again.');
    }
});

function renderTasks(tasks) {
    const tasksContainer = document.getElementById('tasks-container');
    tasksContainer.innerHTML = ''; // Clear previous content

    if (tasks.length === 0) {
        tasksContainer.innerHTML = '<p>No tasks found.</p>';
        return;
    }

    // Render each task
    tasks.forEach(task => {
        const taskElement = document.createElement('div');
        taskElement.className = 'task-card';
        taskElement.innerHTML = `
            <h3>${task.title}</h3>
            <p>${task.description}</p>
            <p><span class="status ${task.status.toLowerCase().replace(' ', '-')}">${task.status}</span></p>
        `;
        tasksContainer.appendChild(taskElement);
    });
}