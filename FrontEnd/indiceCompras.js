const url = "http://localhost:8080/compras";

document.addEventListener("DOMContentLoaded", function () {
  if (!isUserLogged()) {
    window.location = 'login.html?reason=login_required';
  }

  compraDto = obtenerCompras();

  document.getElementById("btnFiltrar").addEventListener("click", function (eventClick) {
    eventClick.preventDefault();
    obtenerCompras();
  })

  document.getElementById("botonGenerarCompra").addEventListener("click", function () {
    generarCompra();
  });

  document.getElementById("btnLimpiarFiltros").addEventListener("click", function () {
    document.getElementById('FiltroTipoAlimento').value = '';
    document.getElementById('FiltroNombreAlimento').value = '';

    obtenerCompras();
  });

  document.getElementById("botonGenerarCompraPer").addEventListener("click", function () {
    generarCompraPersonalizada();
  });
  
});

async function obtenerCompras() {
  try {
    const tipoAlimento = document.getElementById('FiltroTipoAlimento').value;
    const nombreAlimento = document.getElementById('FiltroNombreAlimento').value;

    const compraResumen = document.getElementById("compraResumen");
    if (nombreAlimento === "" && tipoAlimento === "") {
      compraResumen.style.display = "block";
      await makeRequest(`${url}`, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successFn, errorFn);
    } else {
      compraResumen.style.display = "none";
      await makeRequest(`${url}?nombre=${nombreAlimento}&tipo=${tipoAlimento}`, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successFn, errorFn);
    }
  } catch (error) {
    console.error("Error [obtenerCompras]: ", error?.message || error);
  }
}
let seleccionadosGlobal = new Set();
function successFn(data) {
  const elementosTable = document.querySelector("#elementosTable tbody");
  let totalCompra = 0;

  while (elementosTable.firstChild) {
    elementosTable.removeChild(elementosTable.firstChild);
  }

  if (data.AlimentosAComprar != null && data.AlimentosAComprar.length > 0) {
    data.AlimentosAComprar.forEach((elemento, index) => {
      const row = document.createElement("tr");
      row.id = elemento.IDAlimento;

      const nombreCell = document.createElement("td");
      nombreCell.textContent = elemento.NombreAlimento;

      const cantidadCell = document.createElement("td");
      cantidadCell.textContent = elemento.CantidadAComprar;

      const costoCell = document.createElement("td");
      costoCell.textContent = elemento.Costo.toLocaleString("es-AR", { style: "currency", currency: "ARS" });

      const checkboxCell = document.createElement("td");
      const checkbox = document.createElement("input");
      checkbox.type = "checkbox";
      checkbox.name = `alimento_${index}`;
      checkbox.classList.add("select-alimento");

      if (seleccionadosGlobal.has(elemento.IDAlimento.toString())) {
        checkbox.checked = true;
        totalCompra += elemento.Costo; 
      }

      checkbox.addEventListener("change", (event) => {
        if (event.target.checked) {
          seleccionadosGlobal.add(elemento.IDAlimento.toString());
          totalCompra += elemento.Costo;
        } else {
          seleccionadosGlobal.delete(elemento.IDAlimento.toString());
          totalCompra -= elemento.Costo;
        }

        document.getElementById("costoTotal2").textContent = totalCompra.toLocaleString("es-AR", {
          style: "currency",
          currency: "ARS"
        });
      });

      checkboxCell.appendChild(checkbox);
      row.appendChild(nombreCell);
      row.appendChild(cantidadCell);
      row.appendChild(costoCell);
      row.appendChild(checkboxCell);

      elementosTable.appendChild(row);
    });

    document.getElementById("costoTotal").textContent = data.CostoTotal.toLocaleString("es-AR", {
      style: "currency",
      currency: "ARS"
    });
    document.getElementById("costoTotal2").textContent = totalCompra.toLocaleString("es-AR", {
      style: "currency",
      currency: "ARS"
    });
  } else {
    document.getElementById("costoTotal").textContent = "$0,00";
    document.getElementById("costoTotal2").textContent = "$0,00";
  }
}

function obtenerAlimentosSeleccionados() {
  const checkboxes = document.querySelectorAll(".select-alimento:checked");
  const alimentosSeleccionados = [];

  checkboxes.forEach((checkbox) => {
    const row = checkbox.closest("tr"); 
    const id = row.id;
    const nombre = row.children[1].textContent; 
    const cantidad = parseFloat(row.children[2].textContent.replace(/[^0-9,.-]+/g, "").replace(",", "."));

    alimentosSeleccionados.push({
      IDAlimento: id, 
      NombreAlimento: nombre,
      CantidadAComprar: parseInt(cantidad, 10),
    });
  });

  return alimentosSeleccionados;
}
    
async function generarCompra() {
  const productos = {}
  try {
    await makeRequest(`${url}/`, Method.POST, compraDto, ContentType.JSON, CallType.PRIVATE, successFn1, errorFn);
  } catch (error) {
    console.error("Error [generarCompra]: ", error?.message || error);
  }
}

async function generarCompraPersonalizada() {
  const ids = Array.from(seleccionadosGlobal); 

  try {
    console.log(ids);
    await makeRequest(`${url}/`, Method.POST, { productos: ids }, ContentType.JSON, CallType.PRIVATE, successFn1, errorFn);
  } catch (error) {
    console.error("Error [generarCompraPersonalizada]: ", error?.message || error);
  }
}

function successFn1() {
  alert("Compra realizada con éxito");
  location.reload();
  obtenerCompras(); 
}

function errorFn(status, response) {
  console.log("Falla:", response);
}


