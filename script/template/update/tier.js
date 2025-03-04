import { getElement} from "../../read.js";
import {updateValue} from "../../update.js";


// Function to handle updating a tier
export async function update_and_send_tier(oldTier) {
    const form = document.getElementById('updateTierForm');
    const formData = new FormData(form);


    const updatedTier = {
        id: parseInt(formData.get('id')),
        name: formData.get('name'),
        geburtsdatum: formData.get('geburtsdatum'),
        gebaude: await getElement(formData.get('gebaude'), "gebaude"),
    };

    await updateValue(updatedTier, "tier");
}
