import {getElements, getElement} from "../../read.js";
import {updateValue} from "../../update.js";
import {update_and_send_tier} from "./tier.js";

export async function setup_update_pfleger() {

    const pfleger = {
        id: parseInt(document.getElementById("pfleger_id").innerText.split(': ')[1]),
        name: document.getElementById("pfleger_name").innerText.split(': ')[1],
        telefonnummer: document.getElementById("pfleger_telefonnummer").innerText.split(': ')[1],
        adresse: document.getElementById("pfleger_adresse").innerText.split(': ')[1],
        ort: document.getElementById('pfleger_ort').innerText.split(': ')[1],
        revier: document.getElementById('pfleger_revier').innerText.split(': ')[1],
    }

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

    document.getElementById('updatePflegerForm').addEventListener('submit', async function (event) {
        event.preventDefault();
        await update_and_send_pfleger(pfleger);
    });
}

export async function update_and_send_pfleger(oldPfleger) {
    const form = document.getElementById('updatePflegerForm');
    const formData = new FormData(form);

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