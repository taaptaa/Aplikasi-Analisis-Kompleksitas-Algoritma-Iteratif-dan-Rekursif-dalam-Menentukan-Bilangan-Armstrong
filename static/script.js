document.getElementById("form").addEventListener("submit", async (e) => {
    e.preventDefault();

    const numberInput = document.getElementById("number");
    const methodInput = document.getElementById("method");
    const btn = document.getElementById("check-btn");

    const number = numberInput.value;
    const method = methodInput.value;

    const resultContainer = document.getElementById("result-container");
    const resultText = document.getElementById("result-text");
    const timeOutput = document.getElementById("time-output");
    const memoryOutput = document.getElementById("memory-output");
    const errorMsg = document.getElementById("error-message");

    // Reset UI
    resultContainer.classList.add("hidden");
    errorMsg.classList.add("hidden");
    resultText.textContent = "Mengecek...";
    resultText.classList.remove("is-armstrong", "not-armstrong");

    btn.textContent = "Mengecek...";
    btn.disabled = true;

    // Validasi input
    if (number === "" || parseInt(number) < 0) {
        errorMsg.textContent = "Masukkan bilangan bulat positif (≥ 0).";
        errorMsg.classList.remove("hidden");
        btn.textContent = "Cek Sekarang";
        btn.disabled = false;
        return;
    }

    try {
        const res = await fetch("/check", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                number: parseInt(number),
                method: method
            })
        });

        if (!res.ok) {
            throw new Error(`Server error (${res.status})`);
        }

        const data = await res.json();

        // Hasil utama
        if (data.armstrong) {
            resultText.textContent = `🥳 ${number} ADALAH Bilangan Armstrong!`;
            resultText.classList.add("is-armstrong");
        } else {
            resultText.textContent = `😞 ${number} BUKAN Bilangan Armstrong.`;
            resultText.classList.add("not-armstrong");
        }

        // Metrik kinerja
        timeOutput.textContent = `${data.time.toLocaleString("id-ID")} µs`;
        memoryOutput.textContent = `${data.memory.toLocaleString("id-ID")} bytes`;

        resultContainer.classList.remove("hidden");

    } catch (err) {
        console.error(err);
        errorMsg.textContent = "Gagal menghubungi server. Pastikan backend Go berjalan.";
        errorMsg.classList.remove("hidden");
    } finally {
        btn.textContent = "Cek Sekarang";
        btn.disabled = false;
    }
});
