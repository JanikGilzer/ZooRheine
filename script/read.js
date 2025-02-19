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
    await fetch(url)
        .then(response => response.text())
        .then(footer => {
            document.getElementById('Header').innerHTML = footer;
        });
}   
