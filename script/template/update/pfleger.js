import {getElement} from "../../read.js";
import {updateValue} from "../../update.js";


export async function update_and_send_pfleger() {
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
    await updateValue(newPfleger, 'pfleger')
}