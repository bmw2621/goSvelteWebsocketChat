<script>
    import {userId} from "../store"
    export let message;

    const getMessageType = () => {
        switch(message.senderId){
            case "System":
                return "system-message"
            case $userId :
                return "my-message"
            default:
                return null
        }
    }
</script>

<div id="message" class={`message ${getMessageType()}`}>
    <div class="messageSender">{message.senderId === "System" ? message.message : message.username}</div>
    {#if message.senderId !== "System"}
    <div class="messageText">{message.message}</div>
    {/if}
</div>

<style>
    .message {
        text-align: unset;
        justify-self: start;
        display: flex;
        flex-direction: column;
        text-align: left;
        max-width: 100%;
    }

    .message .messageText {
        background: blue;
        color: white;
        border-radius: 20px;
        padding: 5px 10px;
    }

    .my-message {
        justify-self: end;
        text-align: right;
    }

    .my-message .messageText {
        background: green;
        color: white;
    }

    .messageSender {
        font-style: italic;
        color: gray;
        font-size: 0.8rem;
    }

    .message.system-message {
        width: 100%;
        text-align: center;
    }
</style>