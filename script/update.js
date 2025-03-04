export async function updateValue(newValues, element)  {
    try {
        console.log(newValues)
        const url = '/server/update/' + element;
        console.log(url)
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(newValues)
        });
        if (!response.ok) {
           console.log('Network response was not ok');
        }
    } catch (error) {
        console.error('Failed to update element:', error);
    }
}
