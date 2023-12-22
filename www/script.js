document.addEventListener('DOMContentLoaded', function() {
    updateTime();
    setInterval(updateTime, 1000);

    document.getElementById('one-hour').addEventListener('click', () => addEntry(60));
    document.getElementById('two-hour').addEventListener('click', () => addEntry(120));
});

function updateTime() {
    const now = new Date();
    document.getElementById('time-display').innerText = now.toLocaleTimeString('en-US', { hour12: true });
}

function addEntry(minutes) {
    const name = document.getElementById('name-input').value;
    if (!name) return;

    const endTime = new Date(new Date().getTime() + minutes * 60000);
    const box = document.createElement('div');
    box.className = 'flex-box';
    box.innerHTML = `
        <span>${name}</span>
        <span>Time: ${endTime.toLocaleTimeString('en-US', { hour12: true })}</span>
        <span class="time-remaining"></span>
        <button onclick="this.parentNode.remove()">Remove</button>
    `;

    const interval = setInterval(() => {
        const remaining = Math.round((endTime - new Date()) / 60000);
        box.querySelector('.time-remaining').innerText = `Remaining: ${remaining} min`;
        if (remaining <= -5) {
            box.remove();
            clearInterval(interval);
        }
    }, 1000);

    document.getElementById('flex-container').appendChild(box);
}
