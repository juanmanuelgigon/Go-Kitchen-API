document.addEventListener("DOMContentLoaded", async function () {
    function getParameterByName(name) {
        const urlParams = new URLSearchParams(window.location.search);
        return urlParams.get(name);
    }
    
    if (!isUserLogged()) {
        window.location = 'login.html?reason=login_required';
    }

    const alimentoID = getParameterByName('id');

    if (!alimentoID) {
        window.location = 'indiceAlimento.html';
    }

    await cargarDatosAlimento(alimentoID);

    document.getElementById("btnCargar").addEventListener("click", async function (event) {
        event.preventDefault();
        await guardarCambiosAlimento(alimentoID);
    });
});

async function cargarDatosAlimento(id) {
    try {
        await makeRequest(`http://localhost:8080/alimentos/${id}`, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successFn, errorFn);

    } catch (error) {
        console.error("Error al cargar los datos del alimento: ", error?.message || error);
        alert("Hubo un problema al cargar los datos del alimento.");
    }
}

function obtenerMomentosSeleccionados() {
    const momentosSeleccionados = [];
    const momentos = ["Almuerzo", "Merienda", "Cena", "Desayuno"];

    momentos.forEach(momento => {
        const checkbox = document.getElementById(momento);
        if (checkbox.checked) {
            momentosSeleccionados.push(checkbox.value);
        }
    });
    return momentosSeleccionados;
}

async function guardarCambiosAlimento(id) {
    try {
        const momentos = obtenerMomentosSeleccionados();

        const data = {
            NombreAlimento: document.getElementById("nombreAlimento").value,
            TipoAlimento: document.getElementById("tipoAlimento").value,
            PrecioUnitario: parseFloat(document.getElementById("precioUnitario").value),
            CantidadActual: parseInt(document.getElementById("cantidadActual").value),
            CantidadMinima: parseInt(document.getElementById("cantidadMinima").value),
            MomentoAlimento: momentos,
        };

        console.log("Datos a enviar:", data);

        await makeRequest(`http://localhost:8080/alimentos/${id}`, Method.PUT, data, ContentType.JSON, CallType.PRIVATE, funcionOk, funcionError);

    } catch (error) {        
        alert("Hubo un problema con la api.");
    }
}

function successFn(response) {
    document.getElementById("nombreAlimento").value = response.NombreAlimento;
    document.getElementById("tipoAlimento").value = response.TipoAlimento;
    document.getElementById("precioUnitario").value = response.PrecioUnitario;
    document.getElementById("cantidadActual").value = response.CantidadActual;
    document.getElementById("cantidadMinima").value = response.CantidadMinima;
    const momentos = response.MomentoAlimento || []; 

    ["Almuerzo", "Merienda", "Cena", "Desayuno"].forEach(momento => {
        const checkbox = document.getElementById(momento);
        checkbox.checked = momentos.includes(momento);
    });
}

function errorFn() {
    console.error("Error al cargar los datos del alimento");
    alert("Hubo un problema al cargar los datos del alimento.");
}

async function funcionOk(){
    alert("Alimento actualizado exitosamente");
    window.location = 'indiceAlimento.html';
}

async function funcionError(){
    console.error("Error al guardar los cambios del alimento");
    alert("Hubo un problema al guardar los cambios.");
}
