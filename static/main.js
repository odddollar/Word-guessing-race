// set initial word on page load
let currentWord = document.getElementById("word")
currentWord.innerText = words[wordIndex][0] + " _ " + words[wordIndex][2] + " _ " + words[wordIndex][4];

// check guess on input change
const guessBox = document.getElementById("guess");
guessBox.addEventListener("keypress", e => {
    checkGuess(guessBox.value);
})

// update scoreboard every X period of time
let scoreRefresh = setInterval(() => {
    getScoreboardJSON().then(res => {
        // remove old scoreboard and replace with updated one
        removeChildren(document.getElementById("scoreboard"), 1);
        addScoreboardRows(res);

        // check if there's a winner
        checkWinner(res);
    })
}, 2000);

// return promise that will resolve to json data
function getScoreboardJSON() {
    return fetch("http://" + url + "/score", {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8",
        },
    }).then(res => res.json())
    .catch(e => console.log("Error: ", e));
}

// take json data and add rows to scoreboard table
function addScoreboardRows(json) {
    const scores = JSON.parse(json);
    for (score of scores) {
        addScoreboardRow(score);
    }
}

function addScoreboardRow(row) {
    // get constants to append together
    const sb = document.getElementById("scoreboard");
    const newRow = document.createElement("tr");

    // create new table row and set data
    const name = document.createElement("td")
    name.innerText = row.name;
    newRow.appendChild(name);
    const score = document.createElement("td");
    score.innerText = row.score;
    newRow.appendChild(name);
    newRow.appendChild(score);

    // append new row to parent table element
    sb.appendChild(newRow);
}

// iterate through child elements of node and remove all but first
function removeChildren(element, remaining) {
    let child = element.lastElementChild;
    while (element.childElementCount > remaining) {
        element.removeChild(child);
        child = element.lastElementChild;
    }
}

function checkWinner(json) {
    // iterate through all users and check for winner
    for (user of JSON.parse(json)) {
        if (user.score == totalWords) {
            // stop pinging server for score data
            clearInterval(scoreRefresh);
            
            // display winner
            displayWinner(user.name);
        }
    }
}

function displayWinner(winnerName) {
    const game = document.getElementById("game");

    // remove game elements
    removeChildren(game, 0);

    // add winner elements
    const banner = document.createElement("p");
    banner.innerText = winnerName + " won!";

    game.appendChild(banner);
}

function checkGuess(guess) {
    if (guess == words[wordIndex]) {
        // update to new word
        wordIndex++;
        currentWord.innerText = words[wordIndex] != null ? words[wordIndex][0] + " _ " + words[wordIndex][2] + " _ " + words[wordIndex][4] : "No more words";
        guessBox.value = "";

        // post completion status back to server to update score
        updateScore();
    }
}

// post username as json to api endpoint to register completed word
function updateScore() {
    fetch("http://" + url + "/score", {
        method: "POST",
        headers: {
            "Content-type": "application/json; charset=UTF-8",
        },
        body: JSON.stringify({
            name: username,
        })
    }).catch(e => console.error("Error: ", e));
}
