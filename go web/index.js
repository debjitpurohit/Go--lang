
const express = require('express')
const app = express()
const port = process.env.PORT || 5000;
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

//routing
app.get('/', (req, res) => {
    res.status(200).send('Welcome to Go programming')
})
app.get('/get', (req, res) => {
    res.status(200).json({message: 'Welcome to Go programming json'})
})
app.post('/post', (req, res) => {
    let myData = req.body;

    res.status(200).send(myData);
})
app.post('/postform', (req, res) => {
    res.status(200).send(JSON.stringify(req.body));
})



app.listen(port, () => {
    console.log(`listening on port at http://localhost:${port}`)
})
