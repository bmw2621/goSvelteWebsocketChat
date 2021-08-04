<script>
    import {username, server} from "../store"


    const handleConnect = () => {
        server.set(`${document.getElementById("server_input").value}`)
    }

    const handleUsernameChange = e => {
        username.set(e.target.value)
    }
</script>

<h1>Enter your server address</h1>
<input type="text" id="server_input" placeholder="Server Address"><br/>
<input type="text" placeholder="Username" on:input={handleUsernameChange}><br/>
<button disabled={$username === ""} on:click={handleConnect}>Connect</button>

<style>
    h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}
</style>