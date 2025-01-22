const express = require('express')
const app = express()
const port = 3000
const axios = require('axios')

app.get('/', async (req, res) => {
    const response = await axios.get("http://localhost:8080/info")

    console.log(`received response from server: ${JSON.stringify(response.data)}`)
    res.send(response.data)
})

app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})