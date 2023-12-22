package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<style>
					body {
						font-family: 'Comic Sans MS', cursive, sans-serif;
						display: flex;
						flex-direction: column;
						align-items: center;
						justify-content: center;
						height: 100vh;
						margin: 0;
						background-color: #ffffcc;
					}
			
					#header {
						display: flex;
						justify-content: space-between;
						width: 100%;
						padding: 0 20px;
						box-sizing: border-box;
						background-color: #ffcc66;
						margin-bottom: 20px;
						border-bottom: 2px solid #cc9900;
					}
			
					#clock {
						font-size: 24px;
						color: #cc0000;
						text-shadow: 2px 2px 4px #ff6666;
					}
			
					#nameInput {
						margin-right: 10px;
						padding: 5px;
						border: 1px solid #666666;
						background-color: #ffffe6;
					}
			
					#nameForm {
						display: flex;
					}
			
					.nameEntry {
						display: flex;
						flex-direction: column;
						align-items: center;
						margin: 10px;
						border: 2px solid #cc9900;
						padding: 10px;
						max-width: 300px;
						background-color: #ffcc99;
					}
			
					.removeButton {
						cursor: pointer;
						color: #cc0000;
						margin-top: 5px;
						background-color: #ffe6e6;
						border: 1px solid #cc0000;
						padding: 3px 8px;
						border-radius: 5px;
					}
				</style>
				<script>
					function updateClock() {
						var now = new Date();
						var hours = now.getHours().toString().padStart(2, '0');
						var minutes = now.getMinutes().toString().padStart(2, '0');
						var seconds = now.getSeconds().toString().padStart(2, '0');
						document.getElementById('clock').innerText = hours + ':' + minutes + ':' + seconds;
					}
			
					function addNameEntry(name, hoursToAdd) {
						var entryTime = new Date();
						var expirationTime = new Date(entryTime.getTime() + hoursToAdd * 60 * 60 * 1000);
			
						var nameEntryContainer = document.createElement('div');
						nameEntryContainer.className = 'nameEntry';
			
						var nameElement = document.createElement('p');
						nameElement.innerText = 'Name: ' + name;
						nameEntryContainer.appendChild(nameElement);
			
						var entryTimeElement = document.createElement('p');
						entryTimeElement.innerText = 'Entry Time: ' + entryTime.toLocaleTimeString();
						nameEntryContainer.appendChild(entryTimeElement);
			
						var remainingTimeElement = document.createElement('p');
						nameEntryContainer.appendChild(remainingTimeElement);
			
						var removeButton = document.createElement('button');
						removeButton.className = 'removeButton';
						removeButton.innerText = 'Remove';
						removeButton.addEventListener('click', function () {
							document.body.removeChild(nameEntryContainer);
							clearInterval(intervalId);
						});
						nameEntryContainer.appendChild(removeButton);
			
						document.body.appendChild(nameEntryContainer);
			
						function updateRemainingTime() {
							var remainingTime = Math.ceil((expirationTime - new Date()) / 1000 / 60);
							remainingTimeElement.innerText = 'Remaining Time: ' + remainingTime + ' minutes';
			
							if (remainingTime <= -5) {
								document.body.removeChild(nameEntryContainer);
								clearInterval(intervalId);
							}
						}
			
						updateRemainingTime();
						var intervalId = setInterval(updateRemainingTime, 1000);
					}
			
					document.addEventListener('DOMContentLoaded', function () {
						setInterval(updateClock, 1000);
			
						document.getElementById('nameInput').addEventListener('input', function () {
							this.value = this.value.trim();
						});
			
						document.getElementById('nameForm').addEventListener('submit', function (event) {
							event.preventDefault();
							var name = document.getElementById('nameInput').value;
							if (name) {
								var hoursToAdd = parseInt(event.submitter.value, 10);
								addNameEntry(name, hoursToAdd);
							}
							document.getElementById('nameInput').value = '';
						});
					});
				</script>
			</head>
			<body>
			
			<div id="header">
				<div>
					<div id="clock"></div>
					<form id="nameForm">
						<label for="nameInput">Name:</label>
						<input type="text" id="nameInput" required>
						<button type="submit" name="submit" value="1">1 Hour</button>
						<button type="submit" name="submit" value="2">2 Hours</button>
					</form>
				</div>
			</div>
			
			</body>
			</html>


		`
		fmt.Fprint(w, html)
	})

	port := 8080
	fmt.Printf("Starting server on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
