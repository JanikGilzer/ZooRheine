import {getElements} from "../../read.js";

export async function setup_gebaude_modal()
{
    const modal = document.getElementById('myModal');
    const searchInput = document.getElementById('searchInput');

    // Initial load
    const gebaude = await getElements("gebaude");
    await fetchGebaudeBanners();
    modal.style.display = "block";
    document.body.classList.add('modal-open');

    // Search functionality
    searchInput.addEventListener('input', filterGebaude);

    async function fetchGebaudeBanners() {
        for (const g of gebaude ) {
            const gebaudeDiv = document.getElementById("alleGebaude");
            const response = await fetch('/server/template/read/gebaude-banner?id=' + g.ID);
            const gebaudeBanner = await response.text();
            gebaudeDiv.innerHTML += gebaudeBanner;
        }
    }

    function filterGebaude() {
        const searchTerm = searchInput.value.toLowerCase();
        const gebaudeItems = document.querySelectorAll('#alleGebaude > div');

        gebaudeItems.forEach(item => {
            const text = item.textContent.toLowerCase();
            item.style.display = text.includes(searchTerm) ? 'block' : 'none';
        });
    }



    document.getElementById('closeModal').onclick = () => {
        modal.style.display = "none";
        document.body.classList.remove('modal-open');
        window.location.reload()
    };

    window.onclick = (event) => {
        if (event.target === modal) {
            modal.style.display = "none";
            document.body.classList.remove('modal-open');
            window.location.reload()
        }
    };
}