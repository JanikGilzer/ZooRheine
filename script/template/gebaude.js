export async function setup_gebaude_banner(gebaude_id, tiere, fuetterungszeiten) {
    document.getElementById("tier-list").id = "tier-list" + gebaude_id;
    for(var t in tiere)
    {
        console.log("tier", tiere[t]);
        if(tiere[t].Gebaude.ID == gebaude_id)
        {
            const tierList = document.querySelector('#tier-list' + gebaude_id);
            const listItem = document.createElement('li');
            listItem.textContent = tiere[t].Name;
            tierList.appendChild(listItem);
        }
    }
    setup_zeiten(gebaude_id, fuetterungszeiten);

}

async function setup_zeiten(gebaude_id, fuetterungszeiten) {
    document.getElementById("fuetterungszeit-list").id = "fuetterungszeit-list" + gebaude_id;
    for (var f in fuetterungszeiten) {
        console.log(fuetterungszeiten[f])
        if (fuetterungszeiten[f].Gebaude.ID == gebaude_id) {
            const zeitenList = document.querySelector('#fuetterungszeit-list' + gebaude_id);
            const listItem = document.createElement('li');
            listItem.textContent = fuetterungszeiten[f].Zeit.Uhrzeit;
            zeitenList.appendChild(listItem);
        }
    }
}
    