import { getElement } from "../../read.js";

export async function loadAnimals(id) {
    const tierart = await getElement(id, "tierart");
    if (!tierart) {
        console.error("No animal data found for ID:", id);
        return ""; // Return empty string if no data
    }

    // Return the HTML string instead of appending directly
    return `
        <div class="animal-card">
            <img src="../../../html/bilder/tierart/${tierart.ID}.jpg" alt="${tierart.ID}">
            <h3>${tierart.Name}</h3>
        </div>
    `;
}