import {getElements} from "../../read.js";


export async function setup_create_gebaude_form()
{
    const rData = await getElements("revier")
    const rSelect = document.getElementById("Revier")
    rData.forEach(revier => {
        const option = document.createElement('option');
        option.value = revier.ID
        option.text = revier.Name
        option.setAttribute('data-revier', JSON.stringify(revier));
        rSelect.appendChild(option);
    })

    const zData = await getElements("zeit")
    const zSelect = document.getElementById('Fuetterungszeit');
    zData.forEach(zeiten => {
        console.log(zeiten);
        const label = document.createElement('label');1
        const checkbox = document.createElement('input');
        checkbox.type = 'checkbox';
        checkbox.value = zeiten.Uhrzeit;
        label.appendChild(checkbox);
        label.appendChild(document.createTextNode(zeiten.Uhrzeit));
        zSelect.appendChild(label);
        zSelect.appendChild(document.createElement('br'));
    });


}

export async function create_and_send_gebaude() {
    document.getElementById('create-gebaude-form').addEventListener('submit', async function (event) {
        event.preventDefault(); // Prevent the default form submission
        const formData = new FormData(event.target);
        const jsonData = {
            'Gebaude': {
                'Name': formData.get('Name'),
                'Revier': {},
            },
            'Zeit': []
        }

        const revierSelect = document.getElementById('Revier');
        const selectedOrt = revierSelect.options[revierSelect.selectedIndex];
        jsonData['Gebaude']['Revier'] = JSON.parse(selectedOrt.getAttribute('data-revier'));

        const zeitSelect = document.getElementById('Fuetterungszeit');
        const checkboxes = zeitSelect.querySelectorAll('input[type="checkbox"]');
        checkboxes.forEach(checkbox => {
            if (checkbox.checked) {
                jsonData['Zeit'].push(checkbox.value);
            }
        });


        console.log(jsonData)

        const response = await fetch('/server/create/gebaude', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(jsonData) // Convert the object to JSON
        });
        window.location.reload();
    });
}


