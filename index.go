package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<!DOCTYPE html>
<html>
<head>
    <title>Pack Sizes Calculator</title>
    <script>
        function getPackSizes() {
            fetch("/sizes")
                .then(response => response.json())
                .then(data => {
                    const packSizesDiv = document.getElementById("packSizes");
                    packSizesDiv.innerHTML = JSON.stringify(data);
                })
                .catch(error => console.error("Error:", error));
        }

        function setPackSizes() {
            const input = document.getElementById("newPackSizes").value;
            const packSizes = input;
			const packSizesStr = packSizes.join(",");
            fetch("/sizes/set?sizes="+packSizesStr, { method: "POST" })
                .then(response => response.json())
                .then(data => {
                    const packSizesDiv = document.getElementById("packSizes");
                    packSizesDiv.innerHTML = JSON.stringify(data);
                })
                .catch(error => console.error("Error:", error));
        }

        function calculatePacks() {
            const items = parseInt(document.getElementById("items").value); 
            fetch("/calculate?items="+items)
                .then(response => response.json())
                .then(data => {
                    const resultDiv = document.getElementById("result");
                    resultDiv.innerHTML = JSON.stringify(data, null, 2);
                })
                .catch(error => console.error("Error:", error));
        }
    </script>
</head>
<body>
    <h1>Pack Sizes Calculator</h1>

    <h2>Available Pack Sizes</h2>
    <div id="packSizes"></div>
    <button onclick="getPackSizes()">Get Pack Sizes</button>

    <h2>Set Pack Sizes (coma separated list)</h2>
    <textarea id="newPackSizes" rows="4" cols="50" placeholder="1,2,3"></textarea>
    <button onclick="setPackSizes()">Set Pack Sizes</button>

    <h2>Calculate Required Packs</h2>
    <label for="items">Enter number of items:</label>
    <input type="number" id="items" min="1" value="1">
    <button onclick="calculatePacks()">Calculate Packs</button>
    <pre id="result"></pre>
    <script>
        getPackSizes();
    </script>
</body>
</html>
`)
}
