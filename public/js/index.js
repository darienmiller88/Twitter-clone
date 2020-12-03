const form = document.querySelector('form')
const loadingElement = document.querySelector('.loading')
const tweedsElement = document.querySelector('.tweeds')
const API_URL = "http://localhost:8080/api/tweeds"

loadingElement.style.display = ''
getAllTweeds()

form.addEventListener("submit", async (event) => {
    event.preventDefault()
    const formData = new FormData(form)
    const name = formData.get("name")
    const content = formData.get("content")
    const tweed = {
        name, 
        content,
    }

    form.style.display = 'none'
    loadingElement.style.display = ''
    const response = await fetch(API_URL, {
        method: "POST",
        body: JSON.stringify(tweed),
        headers: {
            "Content-type": "application/json",
        }
    })

    const result = await response.json()

    //The back end rate limits to one request per second, so hide the form for exactly one second on the front end
    setTimeout(() => {
        form.style.display = ''
    }, 1000)
    form.reset()
    getAllTweeds()
})

async function getAllTweeds(){
    const response = await fetch(API_URL)
    const tweeds = await response.json()

    tweedsElement.innerHTML = ''
    tweeds.reverse()
    tweeds.forEach(tweed => {
        const div = document.createElement('div')
        const header = document.createElement('h3')
        const content = document.createElement('p')
        const date = document.createElement('small')
        //const thumbsUp = document.createElement('i')
        
        header.textContent = tweed.name
        content.textContent = tweed.content
        date.textContent = new Date(tweed.created_at)
        // thumbsUp.className = "fa fa-thumbs-up"
        // thumbsUp.style = "font-size:36px"
        // thumbsUp.onclick = () => {
        //     console.log("icon clicked")
        // }

        div.append(header)
        div.append(content)
        div.append(date)
      //  div.append(thumbsUp)

        tweedsElement.appendChild(div)
    });
    loadingElement.style.display = 'none'
    console.log(tweeds)
}