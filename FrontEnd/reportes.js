let graficoBarras = null;
let graficoTorta = null;
document.addEventListener("DOMContentLoaded", function (eventDOM) {
  if (!isUserLogged()) {
    window.location = 'login.html?reason=login_required';
  }

  document.getElementById("btnMomento").addEventListener("click", function (eventClick) {
    if (graficoBarras) {
        graficoBarras.destroy();
    }
    if (graficoTorta) {
        graficoTorta.destroy();
    }
   eventClick.preventDefault();
    obtenerGraficoMomento();
  })
  document.getElementById("btnAlimento").addEventListener("click", function (eventClick) {
    if (graficoBarras) {
        graficoBarras.destroy();
    }
    if (graficoTorta) {
        graficoTorta.destroy();
    }
    eventClick.preventDefault();
     obtenerGraficoAlimento();
   })
   document.getElementById("btnCosto").addEventListener("click", function (eventClick) {
    if (graficoBarras) {
        graficoBarras.destroy();
    }
    if (graficoTorta) {
        graficoTorta.destroy();
    }
    eventClick.preventDefault();
     obtenerGraficoCostos();
   })
}) 
async function obtenerGraficoAlimento() {
    try {
        await makeRequest(`http://localhost:8080/reportes/tipo`, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successFnMomentoyAlimento, errorFn);
     } catch (error) {
      console.error("Error [obtenerAlimentos]: ", error?.message || error)
    }
}
async function obtenerGraficoMomento() {
    try {
        await makeRequest(`http://localhost:8080/reportes/momento`, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successFnMomentoyAlimento, errorFn);
     } catch (error) {
      console.error("Error [obtenerAlimentos]: ", error?.message || error)
    }
}  
async function obtenerGraficoCostos() {
    try {
        await makeRequest(`http://localhost:8080/reportes/costo`, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successFnCosto, errorFn);
     } catch (error) {
      console.error("Error [obtenerAlimentos]: ", error?.message || error)
    }
}   
function successFnMomentoyAlimento(data) {
    let info1 = [];
    let info2 = [];
    data.forEach(elemento => {
        info1.push(elemento.Tipo);
        info2.push(elemento.Cantidad);
    });

    const ctx = document.getElementById('GraficoBarras').getContext('2d');
    graficoBarras = new Chart(ctx, {
        type: 'bar',
        data: {
            labels: info1,
            datasets: [{
                data: info2,
                backgroundColor: Object.values(CHART_COLORS),
                borderColor: Object.values(CHART_COLORS),
                borderWidth: 1
            }]
        },
        options: {
            plugins: {
                legend: {
                    display: false
                }
            },
            scales: {
                y: {
                    beginAtZero: true
                }
            }
        }
    });

    const ctx2 = document.getElementById('GraficoTorta').getContext('2d');
    graficoTorta = new Chart(ctx2, {
        type: 'pie',
        data: {
            labels: info1,
            datasets: [{
                data: info2,
                backgroundColor: Object.values(CHART_COLORS),
                borderWidth: 1
            }]
        },
        options: {
            responsive: true,
            plugins: {
                legend: {
                    position: 'bottom'
                }
            }
        }
    });
}

function successFnCosto(data) {
    let info1 = [];
    let info2 = [];
    data.forEach(elemento => {
        info1.push(elemento.Tipo);    
        info2.push(elemento.Cantidad);
    })

    const ctx = document.getElementById('GraficoBarras').getContext('2d');
        graficoBarras = new Chart(ctx, {
            type: 'bar',
            data: {
                labels: info1,
                datasets: [{
                    data: info2,
                    backgroundColor: Object.values(CHART_COLORS),
                    borderColor: Object.values(CHART_COLORS),
                    borderWidth: 1
                }]
            },
            options: {
                plugins: {
                    legend: {
                        display: false
                    }
                },
                scales: {
                    y: {
                        beginAtZero: true
                    }
                }
            }
        });
}

function errorFn(response){
    console.log("Falla:", response);
}

const CHART_COLORS = {
    darkBrown: 'rgb(101, 67, 33)',       // Marrón oscuro
    saddleBrown: 'rgb(139, 69, 19)',     // Marrón silla de montar
    peru: 'rgb(205, 133, 63)',           // Marrón claro
    sienna: 'rgb(160, 82, 45)',          // Siena
    chocolate: 'rgb(210, 105, 30)',      // Chocolate
    tan: 'rgb(210, 180, 140)',           // Marrón claro
    wheat: 'rgb(245, 222, 179)',         // Trigo
    burlywood: 'rgb(222, 184, 135)',     // Madera
    lightGoldenrodYellow: 'rgb(250, 250, 210)', // Amarillo dorado claro
    lightYellow: 'rgb(255, 255, 224)',   // Amarillo muy claro
    gold: 'rgb(255, 215, 0)',            // Oro
    khaki: 'rgb(189, 183, 107)',         // Caqui
    olive: 'rgb(128, 128, 0)',           // Oliva
    darkOliveGreen: 'rgb(85, 107, 47)',  // Verde oliva oscuro
    darkGoldenrod: 'rgb(184, 134, 11)',  // Oro oscuro
    moccasin: 'rgb(255, 228, 181)',      // Mocasín
    papayaWhip: 'rgb(255, 239, 184)',    // Papaya
    lightSalmon: 'rgb(255, 160, 122)',   // Salmón claro
    coral: 'rgb(255, 127, 80)',          // Coral
    darkKhaki: 'rgb(189, 183, 107)'      // Caqui oscuro
};

