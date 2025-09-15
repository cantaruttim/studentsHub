const formulario = document.querySelector('form');

// Selecionando os elementos corretos do formulário
const nome = formulario.querySelector('#name'); // corrigido
const email = formulario.querySelector('#email');
const modulo = formulario.querySelector('#module');
const matricula = formulario.querySelector('#registrationNumber'); // corrigido
const q1 = formulario.querySelector('#questionOne'); // corrigido
const q2 = formulario.querySelector('#questionTwo'); // corrigido

function cadastrar() {
    // Validar comprimento mínimo das respostas
    if (q1.value.length < 300) {
        alert('A resposta 1 precisa ter pelo menos 300 caracteres.');
        return;
    }
    
    if (q2.value.length < 150) {
        alert('A resposta 2 precisa ter pelo menos 150 caracteres.');
        return;
    }

    fetch('http://localhost:8000/api/v1/forms-activities/response', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json'
        },
        body: JSON.stringify({
            name: nome.value,
            email: email.value,
            registrationNumber: matricula.value,
            module: modulo.value,
            questionOne: q1.value,
            questionTwo: q2.value
        })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Erro na resposta do servidor');
        }
        return response.json();
    })
    .then(data => {
        console.log('Success:', data);
        alert('Respostas enviadas com sucesso!');
        resetForm(); // só limpa se deu certo
    })
    .catch((error) => {
        console.error('Error:', error);
        alert('Erro ao enviar respostas. Tente novamente.');
    });
}

function resetForm() {
    nome.value = '';
    email.value = '';
    matricula.value = '';
    modulo.value = '';
    q1.value = '';
    q2.value = '';

    // Focar no primeiro campo para facilitar novo preenchimento
    setTimeout(() => {
        email.focus();
    }, 100);
    
    // Resetar mensagens de validação do navegador
    formulario.reset();
    
    // Remover mensagens de erro personalizadas se existirem
    const errorElements = formulario.querySelectorAll('.error-message');
    errorElements.forEach(element => element.remove());
    
    // Remover classes de validação se existirem
    const inputs = formulario.querySelectorAll('input, textarea');
    inputs.forEach(input => {
        input.classList.remove('is-invalid');
        input.classList.remove('is-valid');
    });

    console.log('Formulário resetado com sucesso');
}

formulario.addEventListener('submit', function(event) {
    event.preventDefault(); // impede o envio padrão
    cadastrar();
});
