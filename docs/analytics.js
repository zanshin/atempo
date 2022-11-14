console.log("Cookies: " + navigator.cookieEnabled);
console.log("Browser language: " + navigator.browserLanguage);
console.log("Language: " + navigator.language);
console.log("Platform: " + navigator.platform);
console.log("Connection Speed: " + navigator.connectionSpeed);
console.log("User agent: " + navigator.userAgent);
console.log("Webdriver: " + navigator.webdriver);
console.log("Geolocation: " + navigator.geolocation);

console.log("Referrer: " + document.referrer)
console.log("Location: " + location.href)

console.log((window.decodeURI)?window.decodeURI(document.referrer):document.referrer)
console.log((window.decodeURI)?window.decodeURI(document.URL):document.URL)

//regular expressions to extract IP and country values
const countryCodeExpression = /loc=([\w]{2})/;
// const userIPExpression = /ip=([\w\.]+)/;
const userIPExpression = /ip=((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))/;

//automatic country determination.
function initCountry() {
    return new Promise((resolve, reject) => {
        var xhr = new XMLHttpRequest();
        xhr.timeout = 3000;
        xhr.onreadystatechange = function () {
            if (this.readyState == 4) {
                if (this.status == 200) {
                    countryCode = countryCodeExpression.exec(this.responseText)
                    ip = userIPExpression.exec(this.responseText)
                    if (countryCode === null || countryCode[1] === '' ||
                        ip === null || ip[1] === '') {
                        reject('IP/Country code detection failed');
                    }
                    let result = {
                        "countryCode": countryCode[1],
                        "IP": ip[1]
                    };
                    resolve(result)
                } else {
                    reject(xhr.status)
                }
            }
        }
        xhr.ontimeout = function () {
            reject('timeout')
        }
        xhr.open('GET', 'https://www.cloudflare.com/cdn-cgi/trace', true);
        xhr.send();
    });
}

// Call `initCountry` function
initCountry().then(result => console.log(JSON.stringify(result))).catch(e => console.log(e))
