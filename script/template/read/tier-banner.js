import { getElements } from "../../../script/read.js";

export async function setup_and_create_tier_banner(id) {
    const futter = await getElements('benoetigtesfutter');    
    futter.forEach(f => {
        if (f.Tier.ID == id) {
            // Create the list item element
            const futterItem = document.createElement('li');
            // Set its text content
            futterItem.textContent = f.Futter.Name;    
            // Find the corresponding food list
            const futterListe = document.getElementById("futter-liste" + id);
            // Append the created list item
            futterListe.appendChild(futterItem);
        }
    });
}
