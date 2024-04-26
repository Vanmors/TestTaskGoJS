fetch('https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1')
    .then(response => response.json())
    .then(data => {
        const tableBody = document.getElementById('crypto-body');
        for (let i = 0; i < data.length; i++) {
            const currency = data[i];
            const row = document.createElement('tr');
            row.innerHTML = `
        <td>${currency.id}</td>
        <td>${currency.symbol}</td>
        <td>${currency.name}</td>
      `;
            console.log(i)
            if (currency.symbol === 'usdt') {
                console.log("here")
                row.classList.add('green-bg');
            } else if (i < 5) {
                console.log(currency.symbol)
                row.classList.add('blue-bg');
            }
            tableBody.appendChild(row);
        }
    })
    .catch(error => console.error('Error fetching data:', error));
