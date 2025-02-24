import {getElements, getElement} from "../../read.js";
import {updateValue} from "../../update.js";

export async function setup_update_gebaude() {

    let gebaude = {
        id: parseInt(document.getElementById("gebaude_id").innerText.split(': ')[1]),
        name: document.getElementById("gebaude_name").innerText.split(': ')[1],
        revier: document.getElementById('gebaude_revier').innerText.split(': ')[1],
    }

    let reviere = await getElements("revier")
    document.getElementById("updateId").value = gebaude.id
    document.getElementById("updateName").value = gebaude.name

    for (let rev in reviere) {
        var option = document.createElement('option');
        option.value = reviere[rev].ID;
        option.text = reviere[rev].Name;
        option.setAttribute('data-revier', JSON.stringify(reviere[rev]));
        document.getElementById('updateRevier').appendChild(option);
    }

    document.getElementById('updateGebaudeForm').addEventListener('submit', async function (event) {
        event.preventDefault();
        await update_and_send_gebaude(gebaude);
    });
}

export async function update_and_send_gebaude(oldGebaude) {
    const form = document.getElementById('updateGebaudeForm');
    const formData = new FormData(form);
    console.log("revier id ->" + formData.get('revier'));
    let newRevier = await getElement(formData.get('revier'), "revier");

    const updatedGebaude = {
        id: parseInt(formData.get('id')),
        name: formData.get('name'),
        revier: newRevier
    };
    console.log(updatedGebaude)
    await updateValue(oldGebaude, updatedGebaude, "gebaude");
}