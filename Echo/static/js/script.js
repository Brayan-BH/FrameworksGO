function changeMessage() {
    const messageElement = document.getElementById('message');
    const currentMessage = messageElement.textContent
    
    const alternateMessages = [
        'I am a JavaScript Developer',
        'I am a Python Developer',
        'I am a C++ Developer',
    ]

    let newMessage;

    do {
        newMessage = alternateMessages[Math.floor(Math.random() * alternateMessages.length)];
    } while(newMessage === currentMessage);

    messageElement.textContent = newMessage;
}