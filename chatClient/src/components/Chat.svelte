<script>
    import Message from "./Message.svelte"
    import { messages, username } from "../store"
    import {checkForId} from "../utils"
    export let server;
    
    let messageInput = ""
    let socket = new WebSocket(server)

    

    socket.onopen = () => {
        socket.send(JSON.stringify({action: "initial_connection", username: $username}))
    }

    socket.onmessage = (event) => {
        
        let data = JSON.parse(event.data)
        
        if (checkForId(data)) return
        
        messages.update( msgs => {
            if(Array.isArray(data)) return [...msgs, ...data]
            return [...msgs, data]
        })
    }

    const sendMessage = () => {
        if(messageInput.length){
            socket.send(JSON.stringify({action: "post_message", message: messageInput}))
        }
        messageInput = ""
    }

</script>

<div id="container">
    <div id="chatBox">
        {#if $messages}
        {#each $messages as message, i (i)}
            <Message message={message}/>
        {/each}
        {/if}
    </div>
    <div id="sendBox">
    <input bind:value={messageInput} type="text"><button on:click={sendMessage}>Send</button>
    </div>
</div>

<style>
    #container {
        width: 100%;
        height: 100%;
        max-width: 900px;
        margin: 0 auto;
        display: grid;
        grid-template-rows: auto 50px;
        grid-template-columns: 100%;
        grid-gap: 10px;
    }

    #chatBox {
        overflow-y: hidden;
        align-self: end;
        display: grid;
        grid-gap: 5px;
    }

    #sendBox {
        width: 100%;
        display: grid;
        grid-template-columns: auto 100px;
        grid-gap: 10px;
    }
</style>