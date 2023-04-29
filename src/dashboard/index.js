cluster_urls = [
    'http://localhost:8000/metadata/',
    'http://localhost:8001/metadata/',
    'http://localhost:8002/metadata/',
]

function update_node_status(data) {
    if (data['error']) {
    }

    console.log(data)
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
    console.log('*')
}

// fetch_cluster_state()
setInterval(fetch_cluster_state, 1000)