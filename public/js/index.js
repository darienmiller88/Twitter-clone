const form = document.querySelector('form')
const loadingElement = document.querySelector('.loading')
const API_URL = "http://localhost:8080/api/tweeds"

loadingElement.style.display = 'none'

form.addEventListener("submit", (event) => {
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
    fetch(API_URL, {
        method: "POST",
        body: JSON.stringify(tweed),
        headers: {
            "Content-type": "application/json",
        }
    })
    console.log(tweed)
})