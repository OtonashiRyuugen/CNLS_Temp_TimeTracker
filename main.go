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
						font-family: Arial, sans-serif;
						display: flex;
						flex-direction: column;
						align-items: center;
						justify-content: center;
						height: 100vh;
						margin: 0;
					}
			
					#clock {
						font-size: 24px;
						margin-bottom: 20px;
					}
			
					#nameInput {
						margin-bottom: 10px;
					}
			
					.nameEntry {
						display: flex;
						flex-direction: column;
						align-items: center;
						margin-top: 10px;
						border: 1px solid #ccc;
						padding: 10px;
						max-width: 300px;
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

			<div id="clock"></div>
			
			<form id="nameForm">
				<label for="nameInput">Name:</label>
				<input type="text" id="nameInput" required>
				<button type="submit" name="submit" value="1">1 Hour</button>
				<button type="submit" name="submit" value="2">2 Hours</button>
			</form>
			
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
