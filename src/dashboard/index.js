cluster_urls = [
    'http://localhost:8000/metadata/',
    'http://localhost:8001/metadata/',
    'http://localhost:8002/metadata/',
]


servers = [
    'http://localhost:8000/upsert/',
    'http://localhost:8001/upsert/',
    'http://localhost:8002/upsert/',
]

var currentleader = null
const form = document.getElementById('key-value-form')
console.log(form)
if (form) {
    form.addEventListener('submit', async (event) => {
    event.preventDefault();
    const formData = new FormData(form);
    console.log(formData.get('key'))

    if (currentleader != null) {
        leaderurl = servers[currentleader]
        console.log(leaderurl)
        console.log(
            formData.get('key'),
            formData.get('value')
        )
        const response = await fetch(leaderurl, {
            method: 'POST',
            body: JSON.stringify({
                "key": formData.get('key'),
                "value": formData.get('value')
            })
        });
        console.log(await response.text());
        };
    }
    )
}


function update_node_status(data) {
    node_div_id_map = {
        0: 'node0',
        1: 'node1',
        2: 'node2'
    }
    if (data['IsLeader'] === true) {
        currentleader = data['Id']
    }
    id = data['Id']
    node_name = node_div_id_map[id]

    if (data['error']) {
        state_div = document.getElementById(node_name + '-' + 'state')
        state_div.innerHTML = 'Unreachable 🚫'
        return
    }

    state_div = document.getElementById(node_name + '-' + 'state')
    logsize_div = document.getElementById(node_name + '-' + 'logsize')
    term_div = document.getElementById(node_name + '-' + 'term')
    store_div = document.getElementById(node_name + '-' + 'store')

    state_div.innerHTML = 'is leader :- ' + data['IsLeader']
    logsize_div.innerHTML = 'log size :- ' + data['LogLength']
    term_div.innerHTML = 'term :- ' + data['Term']
    store_div.innerHTML = 'store :- ' + JSON.stringify(data['Store'])
}

function fetch_cluster_state() {
    cluster_urls.forEach((url, i) => {
        fetch(url, {
            method: "GET"
        }).then(
            (res) => res.json()
        ).then(
            (res) => update_node_status(res)
        ).catch(
            (e) => {
                update_node_status({'error': true, 'Id': i})
            }
        )
    })
}

// fetch_cluster_state()
setInterval(fetch_cluster_state, 1000)