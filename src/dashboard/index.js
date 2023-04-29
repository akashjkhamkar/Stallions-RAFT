cluster_urls = [
    'http://localhost:8000/metadata/',
    'http://localhost:8001/metadata/',
    'http://localhost:8002/metadata/',
]

function update_node_status(data) {
    if (data['error']) {
        console.log('error')
        return
    }

    node_div_id_map = {
        0: 'node0',
        1: 'node1',
        2: 'node2'
    }

    id = data['Id']
    node_name = node_div_id_map[id]

    state_div = document.getElementById(node_name + '-' + 'state')
    logsize_div = document.getElementById(node_name + '-' + 'logsize')
    term_div = document.getElementById(node_name + '-' + 'term')
    store_div = document.getElementById(node_name + '-' + 'store')

    state_div.innerHTML = 'is leader :- ' + data['IsLeader']
    logsize_div.innerHTML = 'log size :- ' + data['LogLength']
    term_div.innerHTML = 'term :- ' + data['Term']
    store_div.innerHTML = 'store :- ' + data['Store']
}

function fetch_cluster_state() {
    cluster_urls.forEach(url => {
        fetch(url, {
            method: "GET"
        }).then(
            (res) => res.json()
        ).then(
            (res) => update_node_status(res)
        ).catch(
            (e) => {
                update_node_status({'error': true})
            }
        )
    })
}

// fetch_cluster_state()
setInterval(fetch_cluster_state, 1000)