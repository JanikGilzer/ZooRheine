import { getElement, getElements } from "../../../script/read.js";
import {updateValue} from "../../update.js";
// Function to populate the update form with existing tier data
export async function setup_update_tier() {
    // Extract tier details from the pre-rendered HTML
    const tier = {
        id: parseInt(document.getElementById('tierId').innerText.split(': ')[1]),
        name: document.getElementById('tierName').innerText.split(': ')[1],
        geburtsdatum: document.getElementById('tierGeburtsdatum').innerText.split(': ')[1],
        gebaude_id: document.getElementById('tierGehege').innerText.split(': ')[1]
    };

    // Populate the update form fields
    document.getElementById('updateId').value = tier.id;
    document.getElementById('updateName').value = tier.name;
    document.getElementById('updateGeburtsdatum').value = tier.geburtsdatum;

    // Populate the Gebaude dropdown
    const gebaudeList = await getElements('gebaude');
    const gebaudeDropdown = document.getElementById('updateGebaude');
    gebaudeList.forEach(gebaude => {
        const option = document.createElement('option');
        option.value = gebaude.ID;
        option.text = gebaude.Name;
        gebaudeDropdown.appendChild(option);
    });

    // Set the selected Gebaude
    gebaudeDropdown.value = tier.gebaude_id;

    // Populate the delete form with the tier ID
    document.getElementById('deleteId').value = tier.id;

    // Add event listeners for form submissions
    document.getElementById('updateTierForm').addEventListener('submit', async function (event) {
        event.preventDefault();
        await update_and_send_tier(tier);
    });


}

// Function to handle updating a tier
export async function update_and_send_tier(oldTier) {
    const form = document.getElementById('updateTierForm');
    const formData = new FormData(form);

    const updatedTier = {
        id: parseInt(formData.get('id')),
        name: formData.get('name'),
        geburtsdatum: formData.get('geburtsdatum'),
        gebaude: await getElement(formData.get('gebaude'), "gebaude")
    };

    await updateValue(oldTier, updatedTier, "tier");
    alert("Tier updated successfully!");
    window.location.reload()
}
