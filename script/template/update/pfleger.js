import {getElements, getElement} from "../../read.js";
import {updateValue} from "../../update.js";

export async function setup_update_pfleger(pfleger) {
    const orte = await getElements("ort")
    const reviere = await getElements("revier")
    document.getElementById("updateId").value = pfleger.id
    document.getElementById("updateName").value = pfleger.name
    document.getElementById("updateTelefonnummer").value = pfleger.telefonnummer
    document.getElementById("updateAdresse").value = pfleger.adresse
    for (let ort in orte) {
        var option = document.createElement('option');
        option.value = orte[ort].ID;
        option.text = orte[ort].Stadt;
        document.getElementById('updateOrt').appendChild(option);
    }
    for (let rev in reviere) {
        var option = document.createElement('option');
        option.value = reviere[rev].ID;
        option.text = reviere[rev].Name;
        document.getElementById('updateRevier').appendChild(option);
    }
}

export async function update_and_send_pfleger(oldPfleger, form) {
    const formData = new FormData(form)
    const newPfleger = {
        id: parseInt(formData.get('id')),
        name: formData.get('name'),
        telefonnummer: formData.get('telefonnummer'),
        adresse: formData.get('adresse'),
        ort: await getElement(parseInt(formData.get('ort')), 'ort'),
        revier: await getElement(parseInt(formData.get('revier')), 'revier')
    }
    await updateValue(oldPfleger, newPfleger, 'pfleger')
    window.location.reload()
}