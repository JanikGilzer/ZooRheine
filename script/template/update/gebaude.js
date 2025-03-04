import {getElement} from "../../read.js";
import {updateValue} from "../../update.js";


export async function update_and_send_gebaude() {
    const form = document.getElementById('updateGebaudeForm');
    const formData = new FormData(form);
    let newRevier = await getElement(formData.get('Revier'), "revier");

    const updatedGebaude = {
        id: parseInt(formData.get('ID')),
        name: formData.get('Name'),
        revier: newRevier
    };
    console.log(updatedGebaude)
    await updateValue(updatedGebaude, "gebaude");
}