export async function updateValue(oldValues, newValues, element)  {
    for (const key in newValues) {
        oldValues[key] = newValues[key];
    }
    try {
        console.log(oldValues)
        const url = '/server/update/' + element;
        console.log(url)
        console.log(oldValues)
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(oldValues)
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
    } catch (error) {
        console.error('Failed to update element:', error);
    }
}
