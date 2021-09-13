
const formLogin = document.getElementById("formLogin");
const formNewAccount = document.getElementById("formNewAccount");
const frmDeposito = document.getElementById("frmDeposito");
const frmTransferir = document.getElementById("frmTransferir");

document.getElementById("btn-new-account").addEventListener("click", () => {
    document.getElementById("dialog-cadastrar-nova-conta").showModal();
});

document.getElementById("cancel-new-account").addEventListener("click", () => {
    document.getElementById("dialog-cadastrar-nova-conta").close();
});

document.getElementById("btn-deposito").addEventListener("click", () => {
    document.getElementById("dialog-deposito").showModal();
});

document.getElementById("btn-transferencia").addEventListener("click", () => {
    getAccounts();
    document.getElementById("dialog-transferencia").showModal();
});


formNewAccount.onsubmit = async (e) => {
    e.preventDefault();
    createNewAccount(formNewAccount);
}

formLogin.onsubmit = async (e) => {
    e.preventDefault();
    getLogin(formLogin)
    getBalance()
}

frmDeposito.onsubmit = async (e) => {
    e.preventDefault();
    accountDeposit(frmDeposito)
}

frmTransferir.onsubmit = async (e) => {
    e.preventDefault();
    createTransfer(frmTransferir)
}

const getLogin = (form) => {
    const url = "http://localhost:5000/login";

    const options = {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        mode: 'no-cors',
        body: JSON.stringify(Object.fromEntries(new FormData(form)))
    }
    
    fetch(url, options)
    .then(response => {
        if (!response.ok) {
            throw new Error('Falha na requisição')
        }
        return response.json();
    })
    .then(data => {
        localStorage.setItem('user', JSON.stringify(data))
        const name = document.getElementById("user-name");
        name.innerHTML = `Seja bem vindo ${data.name}!`

        getTransfers();
    })
    .catch(err => {
        alert(err)
    })
}

const getBalance = () => {
    const user = JSON.parse(localStorage.getItem('user'));
    const url = `http://localhost:5000/accounts/${user.id}/balance`;

    const options = {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${user.token}`,
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
    }

    fetch(url, options)
    .then(response => {
        if (!response.ok) {
            throw new Error('Falha na requisição')
        }
        return response.json();
    })
    .then(data => {
        const saldo = document.getElementById("saldo");
        saldo.innerHTML = `O saldo atual da sua conta é: ${data.toLocaleString('pt-br',{style: 'currency', currency: 'BRL'})}`;
    })
    .catch(err => {
        alert(err)
    })
}

const createNewAccount = form => {
    const url = "http://localhost:5000/accounts";

    const options = {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(Object.fromEntries(new FormData(form)))
    }

    fetch(url, options)
    .then(response => {
        if (!response.ok) {
            throw new Error('Falha na requisição')
        }
        return response.json();
    })
    .then(() => {
        document.getElementById("dialog-cadastrar-nova-conta").close();
        alert("Conta criada com sucesso!")
    })
    .catch(err => {
        alert(err)
    })
}

const accountDeposit = form => {
    const amount = new FormData(form).get("deposit_amount");

    if (!amount) return alert("Informe o valor.")

    const token = JSON.parse(localStorage.getItem('user')).token;
    const url = `http://localhost:5000/accounts/${amount}/deposit`;

    const options = {
        method: 'POST',
        headers: {
            'Authorization': `Bearer ${token}`,
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
    }

    fetch(url, options)
    .then(response => {
        if (!response.ok) {
            if (response.status >= 400) {    
                if (response.status == 401) {
                    throw new Error('Para fazer um depósito, é preciso estar logado.');
                }
                throw new Error('Falha na requisição');
            }
            throw new Error('Falha na requisição')
        }
        return response.json();
    })
    .then(data => {
        getBalance();
        document.getElementById("dialog-deposito").close();
        alert(data)
    })
    .catch(err => {
        alert(err)
    })
}

const getAccounts = () => {
    const url = "http://localhost:5000/accounts";

    const options = {
        method: 'GET',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
    }

    fetch(url, options)
    .then(response => {
        if (!response.ok) {
            throw new Error('Falha na requisição')
        }
        return response.json();
    })
    .then(data => {
        const select = document.getElementById("account_destination_id");
        select.innerHTML = "";
        Object.entries(data).filter(i => i[1].id != parseInt(JSON.parse(localStorage.getItem('user')).id)).map(m => {
            select.options[select.options.length] = new Option(m[1].name, m[1].id);
        });

    })
    .catch(err => {
        alert(err)
    })
}

const createTransfer = form => {
    const amount = document.getElementById("transfer-amount").value;
    const account_destination_id = document.getElementById("account_destination_id").value;
    const url = `http://localhost:5000/transfers`;

    if (!amount) return alert("Informe o valor.")

    const options = {
        method: 'POST',
        headers: {
            'Authorization': `Bearer ${JSON.parse(localStorage.getItem('user')).token}`,
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "account_destination_id": account_destination_id,
            "amount" : amount
        })
    }

    fetch(url, options)
    .then(response => {
        if (!response.ok) {
            console.log(response.statusText);
            throw new Error('Falha na requisição')
        }
        return response.json();
    })
    .then(data => {
        getBalance();
        document.getElementById("dialog-transferencia").close();
        alert("Transferencia realizada com sucesso!")
        getTransfers();
        console.log(data)
    })
    .catch(err => {
        alert(err)
    })
}

const getTransfers = () => {
    const url = `http://localhost:5000/transfers`;
    const options = {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${JSON.parse(localStorage.getItem('user')).token}`,
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
    }

    fetch(url, options)
    .then(response => {
        if (!response.ok) {
            throw new Error('Falha na requisição')
        }
        return response.json();
    })
    .then(data => {
        mountTransfersTable(data)
    })
    .catch(err => {
        alert(err)
    })
}

const mountTransfersTable = data => {
    const container = document.getElementById("transfers-container");
    container.innerHTML = "";
    let table = "";
    let rows = "";

    Object.entries(data).map(m => {
        formatDate(m[1].created_at)
        rows += `<tr>
                    <td>${m[1].id}</td>
                    <td>${m[1].name}</td>
                    <td>${m[1].amount}</td>
                    <td>${formatDate(m[1].created_at)}</td>
                </tr>`
    })


    table = `<table>
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Nome</th>
                                <th>Valor</th>
                                <th>Data</th>
                            </tr>
                        </thead>
                        <tbody>
                            ${rows}
                        </tbody>
                    </table>`;

    container.insertAdjacentHTML('beforeend', table);
}

const formatDate = dt => {
    let data = new Date(dt);
    return `${String(data.getDay()).padStart(2, '0')}/${String(data.getMonth()+1).padStart(2,'0')}/${data.getFullYear()} ${data.getHours()}:${data.getMinutes()}`;
}