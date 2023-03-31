<script>
    import Card from "./Card.svelte"
    import Modal from "./Modal.svelte"
    import {Modals, closeModal, openModal} from "svelte-modals"

    export let goly

    async function update(data) {
        const json = {
            redirect: data.redirect,
            goly: data.goly,
            random: data.random,
            id: goly.id,
        }

        await fetch("http://localhost:8888/redirect", {
            method: "PATCH",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(json)
        }).then(res => {
            console.log(res);
        })
    }

    function handleOpen(goly) {
        openModal(Modal,{
            title: 'Update Goly',
            send: update,
            goly: goly.go,
            redirect: goly.redirect,
            random: goly.random,
    })
    }
</script>

<Card>
    <p>Goly: http://localhost:8888/r/{goly.goly}</p>
    <p>Redirect: {goly.redirect}</p>
    <p>Clicked: </p>
    <button class="update" on:click="{handleOpen(goly)}">Update</button>
    <button class="delete">Delete</button>
</Card>
<Modals>
    <div
        slot="backdrop"
        class="backdrop"
        on:click="{closeModal}"
    />
</Modals>

<style>
    button {
        color: white;
        font-weight: bolder;
        border:none;
        padding: .75rem;
        border-radius: 4px ;
    }
    .update {
        background-color: yellowgreen;
    }
    .delete {
        background-color: red;
    }

    .backdrop {
        position: fixed;
        top: 0;
        bottom: 0;
        right: 0;
        left: 0;
        background: rgb(255, 255, 255)
    }
</style>