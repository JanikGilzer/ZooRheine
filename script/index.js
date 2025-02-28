import { setup_revier_icon } from './template/revier.js';

export async function setup_index(id, tiere, gebaude, fuetterungszeiten) {
    const revierResponse = await fetch(`/server/template/read/revier?id=${id}`);
    const revierTemplate = await revierResponse.text();
    const revierWithLink = revierTemplate.replace("http://linkEinfügen", `/revier?id=${id}`);
    const container = document.querySelector("#template-container");

    // Create revier container
    const revier_container = document.createElement('div');
    revier_container.id = `revier-container-${id}`;
    revier_container.innerHTML = revierWithLink;

    // Add animation class after appending to the DOM
    revier_container.classList.add('fade-in'); // Add CSS animation class
    container.appendChild(revier_container);

    // Force reflow to trigger the animation
    void revier_container.offsetWidth;

    // Add animation for each gebaude banner
    for (const g in gebaude) {
        if (gebaude[g].Revier.ID == id) {
            await setup_revier_icon(id, gebaude[g].ID, tiere, fuetterungszeiten);
        }
    }
}
