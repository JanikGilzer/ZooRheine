import { getElements } from '../../read.js';

export async function setup_create_tier_form() {
    // Fetch Gebäude data
    var gData = await getElements('gebaude');
    const gSelect = document.getElementById('Gebaude');
    gData.forEach(gebaude => {
        const option = document.createElement('option');
        option.value = gebaude.id; // Store the id in the value attribute
        option.text = gebaude.Name; // Display the name as the text
        option.setAttribute('data-gebaude', JSON.stringify(gebaude)); // Store the whole object as a JSON string
        gSelect.appendChild(option);
    });

    // tierarten data
    var taDaten = await getElements('tierart')
    const taSelect = document.getElementById('Tierart')
    taDaten.forEach(tierart => {
        const option = document.createElement('option')
        option.value = tierart.ID
        option.text = tierart.Name
        option.setAttribute('data-gebaude', JSON.stringify(tierart));
        taSelect.appendChild(option);
    })

    // Fetch Futter data
    var fData = await getElements('futter');
    const fSelect = document.getElementById('Futter');
    fData.forEach(futter => {
        console.log(futter);
        const label = document.createElement('label');
        const checkbox = document.createElement('input');
        checkbox.type = 'checkbox';
        checkbox.value = futter.Name;
        label.appendChild(checkbox);
        label.appendChild(document.createTextNode(futter.Name));
        fSelect.appendChild(label);
        fSelect.appendChild(document.createElement('br'));
    });
}


export async function create_and_send_tier() {
            // Handle form submission
            document.getElementById('create-tier-form').addEventListener('submit', async function(event) {
                event.preventDefault(); // Prevent the default form submission
    
                const form = event.target;
    
                // Collect form data and convert it to a JSON object
                const formData = new FormData(form);
                const jsonData = {
                    'Tier': {
                        'Name': formData.get('Name'),
                        'Geburtsdatum': formData.get('Geburtsdatum'),
                        'Gebaude': {}
                    },
                    'Futter': []
                };
    
                // Add the selected Gebaude object to the jsonData
                const gebaudeSelect = document.getElementById('Gebaude');
                const selectedOption = gebaudeSelect.options[gebaudeSelect.selectedIndex];
                jsonData['Tier']['Gebaude'] = JSON.parse(selectedOption.getAttribute('data-gebaude'));
    
                // Add the selected Futter to the jsonData
                const futterSelect = document.getElementById('Futter');
                const checkboxes = futterSelect.querySelectorAll('input[type="checkbox"]');
                checkboxes.forEach(checkbox => {
                    if (checkbox.checked) {
                        jsonData['Futter'].push(checkbox.value);
                    }
                });
    
                console.log(jsonData);
    
                // Send the data as JSON in the POST request
                const response = await fetch('/server/create/tier', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(jsonData)
                });
    
                if (response.ok) {
                    alert('Tier erfolgreich erstellt!');
                } else {
                    alert('Fehler beim Erstellen des Tieres.');
                }
            });
}