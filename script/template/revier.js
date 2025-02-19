import { setup_gebaude_banner } from './gebaude.js';

export async function setup_revier_banner(rev_id, gebaude_id, tiere, fuetterungszeiten) {
    
    console.log("fill_tables", rev_id, gebaude_id);
    try {
        const response = await fetch(`/server/template/read/gebaude-banner?id=${gebaude_id}`);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const template = await response.text();
        // Select the gebaude-container within the specific revier container
        const revierContainer = document.querySelector(`#revier-container-${rev_id} .gebaude-container`);
        if (revierContainer) {
            revierContainer.innerHTML += template;
            setup_gebaude_banner(gebaude_id, tiere, fuetterungszeiten);
        }
    } catch (error) {
        console.error('Error fetching gebaude banner:', error);
    }
}