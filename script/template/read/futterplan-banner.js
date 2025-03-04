import { getElements } from "../../read.js";

export async function setup_fuetterrung_modal()
{
    const modal = document.getElementById('myModal');
    const searchInput = document.getElementById('searchInput');

    // Initial load
    const gebaude = await getElements("gebaude");
    await fetchfuetterrungBanners();
    modal.style.display = "block";
    document.body.classList.add('modal-open');

    // Search functionality
    searchInput.addEventListener('input', filterfuetterrunge);

    async function fetchfuetterrungBanners() {
        for (const g of gebaude) {
            const fuetterrungDiv = document.getElementById("alleFuetterungszeiten");
            const response = await fetch('/server/template/read/futterplan-banner?id=' + g.ID);
            const fuetterrungBanner = await response.text();
            fuetterrungDiv.innerHTML += fuetterrungBanner;
        }
    }

    function filterfuetterrunge() {
        const searchTerm = searchInput.value.toLowerCase();
        const fuetterrungItems = document.querySelectorAll('#alleFuetterungszeiten > div');

        fuetterrungItems.forEach(item => {
            const text = item.textContent.toLowerCase();
            item.style.display = text.includes(searchTerm) ? 'block' : 'none';
        });
    }
}