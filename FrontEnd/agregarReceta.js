let lista = []; 

document.addEventListener("DOMContentLoaded", function () {

  if (!isUserLogged()) {
    window.location = 'login.html?reason=login_required';
  }

  cargarOpcionesMenu(); 
  document.getElementById("btnAgregarReceta").addEventListener("click", async function (event) {
    event.preventDefault(); 
    await agregarReceta(lista); 
  });

  document.getElementById("btnAgregar").addEventListener("click", async function (event) {
    event.preventDefault(); 
    await agregarALista(); 
  });
});

async function agregarReceta(lista) {
  try {
    const nombre = document.getElementById("NombreReceta").value; 
    const momento = document.getElementById("MomentoReceta").value; 

    const data = {
      NombreReceta: nombre,
      MomentoReceta: momento,
      AlimentosNecesarios: lista 
    };

    await makeRequest(
      "http://localhost:8080/recetas", 
      Method.POST, 
      data,
      ContentType.JSON, 
      CallType.PRIVATE, 
      successFn, 
      errorFn 
    );
  } catch (error) {
    console.error("Error : ", error?.message || error); 
    alert("Carga fallida. Intente nuevamente.");
  }
}

function errorFn(status, response) {
  console.error("Error al guardar los cambios del alimento");
  console.log("Falla:", response);
  alert("Hubo un problema al guardar los cambios."); 
}

function successFn() {
  alert("Receta Cargada"); 
  window.location = 'indiceRecetas.html'; 
}

let opciones = []; 
async function cargarOpcionesMenu() {
  try {
    await makeRequest(`http://localhost:8080/alimentos`, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successFnOpciones, errorFnOpciones);
  } catch (error) {
    console.error("Error [cargarOpcionesMenu]: ", error); 
  }
}

function successFnOpciones(data) {
  opciones = data; 
  const selectAlimentos = document.getElementById("selectAlimentos");

  selectAlimentos.innerHTML = `<option selected disabled>Seleccione un alimento</option>`;
  opciones.sort((a, b) => a.NombreAlimento.localeCompare(b.NombreAlimento));

  opciones.forEach(opcion => {
      const optionElement = document.createElement("option");
      optionElement.value = opcion.IDAlimento; 
      optionElement.textContent = opcion.NombreAlimento; 
      selectAlimentos.appendChild(optionElement); 
  });
}

document.getElementById("selectAlimentos").addEventListener("change", function() {
    const selectedOption = this.options[this.selectedIndex];
    seleccionarOpcion(selectedOption.textContent, selectedOption.value); 
});


function errorFnOpciones() {
  alert("Error al cargar opciones del menú"); 
}

function seleccionarOpcion(opcionNombre, opcionID) {
  document.getElementById("opcionSeleccionada").textContent = opcionNombre; 

  const cantidadContainer = document.getElementById("cantidadContainer");
  cantidadContainer.innerHTML = ""; 

  const label = document.createElement("label");
  label.textContent = "Cantidad: ";
  label.setAttribute("for", "cantidadInput");

  const input = document.createElement("input");
  input.type = "number";
  input.id = "cantidadInput";
  input.placeholder = "Ingrese cantidad";

  cantidadContainer.appendChild(label);
  cantidadContainer.appendChild(input);
 
  const dataArreglo = JSON.stringify({
      NombreAlimento: opcionNombre,
      IDAlimento: opcionID,
  });

  document.getElementById("btnAgregar").style.display = "inline-block"; 
  document.getElementById("btnAgregar").dataset.opcion = dataArreglo; 
}

function agregarALista() {
  const opcion = JSON.parse(document.getElementById("btnAgregar").dataset.opcion);
  const cantidadInput = document.getElementById("cantidadInput");

  if (!cantidadInput) {
    return;
}

const cantidad = cantidadInput.value; 

if (cantidad === "") {
    alert("Por favor, ingrese una cantidad."); 
    return;
}

const cantidadNumerica = parseInt(cantidad);
if (isNaN(cantidadNumerica)) {
    alert("Por favor, ingrese un número válido para la cantidad."); 
    return;
}

  const listaSeleccion = document.getElementById("listaSeleccion");
  const item = document.createElement("li");
  item.textContent = `${opcion.NombreAlimento} - Cantidad: ${cantidadNumerica}`; 
  listaSeleccion.appendChild(item); 

  const alimentoAgg = {
      IDAlimento: opcion.IDAlimento,
      NombreAlimento: opcion.NombreAlimento,
      CantidadNecesaria: cantidadNumerica, 
  };

  lista.push(alimentoAgg);

  document.getElementById("opcionSeleccionada").textContent = "Ninguna";
  document.getElementById("cantidadContainer").innerHTML = ""; 
  document.getElementById("btnAgregar").style.display = "none"; 
}

window.onclick = function (event) {
  if (!event.target.matches('.menu-button')) {
      const menuContent = document.getElementById("menuContent");
      if (menuContent && menuContent.style.display === "block") {
          menuContent.style.display = "none"; 
      }
  }
};