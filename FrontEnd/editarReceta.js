let listita = [];
let recetaOriginal = {}; 

document.addEventListener("DOMContentLoaded", async function () {
    function getParameterByName(name) {
        const urlParams = new URLSearchParams(window.location.search);
        return urlParams.get(name);
    }
 
    if (!isUserLogged()) {
        window.location = 'login.html?reason=login_required';
    }
    
    const recetaID = getParameterByName('id');

    if (!recetaID) {
        window.location = 'indiceRecetas.html';
    }

    cargarDatosReceta(recetaID);

    cargarOpcionesMenu(); 

    document.getElementById("btnGuardar").addEventListener("click", async function (event) {
        event.preventDefault();
        agregarReceta(listita, recetaID);
    });

    document.getElementById("btnAgregar").addEventListener("click", async function (event) {
        event.preventDefault(); 
         agregarALista(); 
      });
});

async function cargarDatosReceta(recetaID) {
        try {
          await makeRequest(`http://localhost:8080/recetas/${recetaID}`, Method.GET, null, ContentType.JSON, CallType.PRIVATE, (response) => {
              recetaOriginal = response; 
              successFn(response); 
          }, errorFn);
      } catch (error) {
          console.error("Error al cargar los datos de la receta: ", error?.message || error);
          alert("Hubo un problema al cargar los datos de la receta.");
      }
}

function successFn(response) {
    receta = response;
    document.getElementById("nombreReceta").value = receta.NombreReceta;

    receta.AlimentosNecesarios.forEach(alimento => {
        const row = document.createElement("tr");
        row.innerHTML = `
            <td>${alimento.NombreAlimento}</td>
            <td>
                <input type="number" value="${alimento.CantidadNecesaria}" 
                       onchange="actualizarCantidad('${alimento.IDAlimento}', this)">
            </td>
            <td>
                <button onclick="eliminarAlimento('${alimento.IDAlimento}', this)">Eliminar</button>
            </td>
        `;
        elementosTable.appendChild(row);
    });

    const momento = response.MomentoReceta;
    ["Almuerzo", "Merienda", "Cena", "Desayuno"].forEach(m => {
        const checkbox = document.getElementById(m);
        checkbox.checked = (m === momento);
    });
}

function actualizarCantidad(alimentoID, inputElement) {
    const nuevaCantidad = parseInt(inputElement.value, 10);
    if (!isNaN(nuevaCantidad)) {
        const alimento = recetaOriginal.AlimentosNecesarios.find(a => a.IDAlimento === alimentoID);
        if (alimento) {
            alimento.CantidadNecesaria = nuevaCantidad;
        }
    } else {
        alert("Por favor, ingrese un número válido.");
    }
}

function eliminarAlimento(alimentoID, buttonElement) {
    const row = buttonElement.closest("tr");
    row.remove();
    recetaOriginal.AlimentosNecesarios = recetaOriginal.AlimentosNecesarios.filter(
        alimento => alimento.IDAlimento !== alimentoID
    );
}

function errorFn() {
    console.error("Error al cargar los datos del alimento");
    alert("Hubo un problema al cargar los datos del alimento.");
}

function obtenerMomentoSeleccionado() {
  const momentos = ["Almuerzo", "Merienda", "Cena", "Desayuno"];
  for (let momento of momentos) {
      const checkbox = document.getElementById(momento);
      if (checkbox && checkbox.checked) {
          return momento;  
      }
  }
  return null; 
}

async function agregarReceta(lista, id) {
    try {
      const nombre = document.getElementById("nombreReceta").value || recetaOriginal.NombreReceta;; 
      const momento = obtenerMomentoSeleccionado() || recetaOriginal.MomentoReceta;;

      recetaOriginal.AlimentosNecesarios.forEach(alimento1 => {
        let encontrado = false;
    
        lista.forEach(alimento2 => {
            if (alimento1.IDAlimento === alimento2.IDAlimento) {
                encontrado = true;
            }
        });
    
        if (!encontrado) {
            listita.push(alimento1);
        }
    });
    
      const data = {
        NombreReceta: nombre,
        MomentoReceta: momento,
        AlimentosNecesarios: lista
      };
  
      await makeRequest(`http://localhost:8080/recetas/${id}`, Method.PUT, data, ContentType.JSON, CallType.PRIVATE, successFn1, errorFn1);
    } catch (error) {
      console.error("Error : ", error?.message || error); 
      alert("Carga fallida. Intente nuevamente."); 
    }
  }
  
  function errorFn1(status, response) {
    console.error("Error al guardar los cambios");
    console.log("Falla:", response);
    alert("Hubo un problema al guardar los cambios."); 
  }
  
  function successFn1() {
    alert("Receta actualizada"); 
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
const momento = obtenerMomentoSeleccionado();

selectAlimentos.innerHTML = `<option selected disabled>Seleccione un alimento</option>`;

opciones.forEach(alimento => {
    const optionElement = document.createElement("option");
    optionElement.value = alimento.IDAlimento;
    optionElement.textContent = alimento.NombreAlimento;
    selectAlimentos.appendChild(optionElement);
});

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

    listita.push(alimentoAgg);
  
    document.getElementById("opcionSeleccionada").textContent = "Ninguna";
    document.getElementById("cantidadContainer").innerHTML = ""; 
    document.getElementById("btnAgregar").style.display = "none"; 
}
function toggleCheckbox(selectedCheckbox) {
    const checkboxes = document.querySelectorAll('.form-check-input');

    checkboxes.forEach(checkbox => {
        if (checkbox !== selectedCheckbox) {
            checkbox.checked = false;
        }
    });
}
  
window.onclick = function (event) {
    if (!event.target.matches('.menu-button')) {
        const menuContent = document.getElementById("menuContent");
        if (menuContent && menuContent.style.display === "block") {
            menuContent.style.display = "none"; 
        }
    }
};