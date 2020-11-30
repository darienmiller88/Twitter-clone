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
    form.reset()
    form.style.display = ''
    getAllTweeds()
    console.log(result)
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
        const date = document.createElement('p')

        header.textContent = tweed.name
        content.textContent = tweed.content
        date.textContent = new Date(tweed.created_at)

        div.append(header)
        div.append(content)
        div.append(date)

        tweedsElement.appendChild(div)
    });
    loadingElement.style.display = 'none'
    console.log(tweeds)
}