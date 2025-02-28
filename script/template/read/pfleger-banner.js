import {getElements} from "../../read.js";

export async function setup_pfleger_modal()
{
    const modal = document.getElementById('myModal');
    const searchInput = document.getElementById('searchInput');

    // Initial load
    const pfleger = await getElements("pfleger");
    await fetchpflegerBanners();
    modal.style.display = "block";
    document.body.classList.add('modal-open');

    // Search functionality
    searchInput.addEventListener('input', filterpfleger);

    async function fetchpflegerBanners() {
        for (const p of pfleger ) {
            const pflegerDiv = document.getElementById("allePfleger");
            const response = await fetch('/server/template/read/pfleger-banner?id=' + p.ID);
            const pflegerBanner = await response.text();
            pflegerDiv.innerHTML += pflegerBanner;
        }
    }

    function filterpfleger() {
        const searchTerm = searchInput.value.toLowerCase();
        const pflegerItems = document.querySelectorAll('#allepfleger > div');

        pflegerItems.forEach(item => {
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