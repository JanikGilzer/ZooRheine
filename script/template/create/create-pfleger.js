import {getElements} from "./../../read.js";

export async function setup_create_pfleger_form() {
    // fetch revier
    var rData = await getElements('revier');
    const rSelect = document.getElementById('PRevier');
    rData.forEach(revier => {
        const option = document.createElement('option');
        option.value = revier.ID; // Store the id in the value attribute
        option.text = revier.Name; // Display the name as the text
        option.setAttribute('data-revier', JSON.stringify(revier)); // Store the whole object as a JSON string
        rSelect.appendChild(option);
    });

    // Fetch ort
    var oData = await getElements('ort');
    console.log(oData)
    const oSelect = document.getElementById('Ort');
    oData.forEach(ort => {
        const option = document.createElement('option');
        option.value = ort.ID; // Store the id in the value attribute
        option.text = ort.Stadt; // Display the name as the text
        option.setAttribute('data-ort', JSON.stringify(ort)); // Store the whole object as a JSON string
        oSelect.appendChild(option);
    });
}

export async function create_and_send_pfleger() { // Remove the parameter
    document.getElementById('create-pfleger-form').addEventListener('submit', async function(event) {
        event.preventDefault();
        const form = event.currentTarget;
        const formData = new FormData(form);

        const jsonData = {
            'Pfleger': {
                'ID' : 1,
                'Name': formData.get('Name'),
                'Telefonnummer': formData.get('Telefonnummer'),
                'Adresse': formData.get('Adresse'),
                'Ort': {},
                'Revier': {}
            }
        };

        // Add the selected Gebaude object to the jsonData
        const revierSelect = document.getElementById('PRevier');
        const selectedOption = revierSelect.options[revierSelect.selectedIndex];
        jsonData['Pfleger']['Revier'] = JSON.parse(selectedOption.getAttribute('data-revier'));

        // Add the selected Futter to the jsonData
        const ortSelect = document.getElementById('Ort');
        const ortSOptions = ortSelect.options[ortSelect.selectedIndex]
        jsonData['Pfleger']['Ort'] = JSON.parse(ortSOptions.getAttribute('data-ort'))

        console.log(jsonData);

        // Send the data as JSON in the POST request
        const response = await fetch('/server/create/pfleger', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(jsonData) // Convert the object to JSON
        });

        if (response.ok) {
            alert('Tier erfolgreich erstellt!');
        } else {
            alert('Fehler beim Erstellen des Tieres.');
        }
    });
}