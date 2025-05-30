let ultimoTimestamp = null;

async function verificarServidor() {
  try {
    const resposta = await fetch("/status");
    const timestamp = await resposta.text();

    if (ultimoTimestamp && ultimoTimestamp !== timestamp) {
      console.log("Servidor reiniciado, recarregando página...");
      location.reload();
    }

    ultimoTimestamp = timestamp;
  } catch (error) {
    console.error("Erro ao verificar servidor:", error);
  }
}

// Checa a cada 3 segundos
setInterval(verificarServidor, 3000);
