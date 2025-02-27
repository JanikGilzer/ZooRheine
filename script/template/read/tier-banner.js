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
            await setup_and_create_tier_banner(t.ID);
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