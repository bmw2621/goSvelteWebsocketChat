import {userId} from "./store"

export const checkForId = (msg)=> {
    const regex = /^yourId:([a-zA-Z0-9]{10})$/
    const match = msg.message.match(regex)
    if(match){
        userId.set(match[1])
        return true
    } 
    return false
}