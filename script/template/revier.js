import { setup_gebaude_icon } from './gebaude.js';

export async function setup_revier_icon(rev_id, gebaude_id) {
    
    console.log("fill_tables", rev_id, gebaude_id);
    try {
        const response = await fetch(`/server/template/read/gebaude-icon?id=${gebaude_id}`);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const template = await response.text();

        const revierContainer = document.querySelector(`#revier-container-${rev_id} .gebaude-container`);
        if (revierContainer) {
            revierContainer.innerHTML += template;

        }
    } catch (error) {
        console.error('Error fetching gebaude banner:', error);
    }
}