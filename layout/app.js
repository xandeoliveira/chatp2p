const inputElement = document.querySelector(".input-box input")
const chatContainer = document.querySelector("main.chat-box")
console.log(chatContainer)

const createMessage = (text, userName) => {
  const strong = document.createElement("strong")
  const head = document.createTextNode(userName)
  strong.appendChild(head)

  const br = document.createElement("br")

  const message = document.createTextNode(text)

  const p = document.createElement("p")
  p.appendChild(strong)
  p.appendChild(br)
  p.appendChild(message)

  const div = document.createElement("div")
  div.classList.add("message")
  div.appendChild(p)

  return div
}

const clearInput = () => inputElement.value = ""

const insertIntoChat = (element, userClass) => {
  element.classList.add(userClass)

  chatContainer.appendChild(element)

  clearInput()
}

const postUserMessage = text => {
  const message = createMessage(text, "User")
  insertIntoChat(message, "user")
}

const postMyMessage = text => {
  const message = createMessage(text, "Eu")
  insertIntoChat(message, "me")
}

const scrollChat = () => chatContainer.scrollTop = chatContainer.scrollHeight

const sendMessage = () => {
  message = inputElement.value.trim()

  if(!message.length) return
  // postMyMessage(message)

  send(message).then(myMessage => {
    postMyMessage(myMessage)
    scrollChat()
  })
}