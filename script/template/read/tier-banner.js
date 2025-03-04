import { getElements } from "../../../script/read.js";



export async function setup_and_create_tier_modal()
{
    const modal = document.getElementById('myModal');
    const searchInput = document.getElementById('searchInput');

    // Initial load
    const tiere = await getElements("tier");
    await fetchTierBanners();
    modal.style.display = "block";
    document.body.classList.add('modal-open');

    // Search functionality
    searchInput.addEventListener('input', filterTiere);

    async function fetchTierBanners() {
        tiere.forEach(async t =>  {
            const tierDiv = document.getElementById("alleTiere");
            const response = await fetch('/server/template/read/tier-banner?id=' + t.ID);
            const tierBanner = await response.text();
            tierDiv.innerHTML += tierBanner;
        });
    }

    function filterTiere() {
        const searchTerm = searchInput.value.toLowerCase();
        const tierItems = document.querySelectorAll('#alleTiere > div');

        tierItems.forEach(item => {
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