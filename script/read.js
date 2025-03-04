export async function getElement(id, element){
    try {
        const response = await fetch('/server/json/' + element + "?id=" + id);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const values = await response.json();
        return values;

    } catch (error) {
        console.error('Failed to read element:', error);
    }
}

export async function getElements(element){
    try {
        const response = await fetch('/server/json/' + element);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const values = await response.json();
        return values;

    } catch (error) {
        console.error('Failed to read element:', error);
    }
}

export async function countElement(element){
    try {
        const response = await fetch('/server/count/' + element);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const value = await response.json();
        return value;

    } catch (error) {
        console.error('Failed to count element:', error);
    }
}

export async function getFooter() {
    const url = '/server/template/footer';
    await fetch(url)
        .then(response => response.text())
        .then(footer => {
            document.getElementById('Footer').innerHTML = footer;
        });
}

export async function getHeader() {
    const url = '/server/template/header';
    await fetch(url, {
        credentials: 'include'
    })
        .then(response => response.text())
        .then(header => {
            document.getElementById('Header').innerHTML = header;
            initializeHeaderFeatures();
            updateHeader();
        });
}

// Separate initialization logic
function initializeHeaderFeatures() {
    // Mobile menu toggle
    const toggle = document.querySelector('.mobile-menu-toggle');
    const nav = document.querySelector('.main-nav');

    if (toggle && nav) {
        toggle.addEventListener('click', function() {
            nav.classList.toggle('active');
            toggle.classList.toggle('active');
        });

        // Close menu when clicking outside
        document.addEventListener('click', function(event) {
            if (!event.target.closest('.header-container')) {
                nav.classList.remove('active');
                toggle.classList.remove('active');
            }
        });

        // Close menu when clicking a link
        document.querySelectorAll('.nav-link').forEach(link => {
            link.addEventListener('click', () => {
                nav.classList.remove('active');
                toggle.classList.remove('active');
            });
        });
    }

    // Desktop hover functionality
    const desktopNavItems = document.querySelectorAll('.main-nav > ul > li');
    desktopNavItems.forEach(item => {
        item.addEventListener('mouseenter', () => {
            if (window.innerWidth > 768) {
                item.classList.add('hover-active');
            }
        });
        item.addEventListener('mouseleave', () => {
            if (window.innerWidth > 768) {
                item.classList.remove('hover-active');
            }
        });
    });
}

function updateHeader() {
    if (isAuthenticated()) {
        updateHeaderForAuthenticatedUser();
    } else {
        updateHeaderForGuest();
    }
}

function updateHeaderForAuthenticatedUser() {
    const loginLi = document.getElementById('login');
    if (loginLi) {
        loginLi.innerHTML = `
            <a href="/admin-panel" class="nav-link">Verwaltung</a>
        `;
    }
}

function updateHeaderForGuest() {
    const loginLi = document.getElementById('login');
    if (loginLi) {
        // Replace the user profile with a login button
        loginLi.innerHTML = `
            <a href="/login" class="nav-link">Login</a>
        `;
    }
}




function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}

export function isAuthenticated() {
    const token = getCookie('token');
    if (!token) return false;

    try {
        const payload = JSON.parse(atob(token.split('.')[1]));
        return payload && payload.exp > Date.now() / 1000;
    } catch (error) {
        console.error('Error decoding token:', error);
        return false;
    }
}