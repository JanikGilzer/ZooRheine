export async function deleteValue(valueId, element) {
    const value = {
        id: parseInt(valueId),
    };

    try {
        const response = await fetch('/delete' + element, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(value)
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
    } catch (error) {
        console.error('Failed to delete user:', error);
    }
}