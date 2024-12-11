document.getElementById('fetch-btn').addEventListener('click', async () => {
    try {
      const response = await fetch('/api/hello');
      const data = await response.json();
      document.getElementById('api-message').textContent = data.message;
    } catch (error) {
      document.getElementById('api-message').textContent = 'Error fetching API message.';
    }
  });
  