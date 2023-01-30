"use strict";

function GithubAuth() {
    return (
        <button>
            Hello there
        </button>
    )
};

const rootNode = document.getElementById("auth-component")
const root = ReactDOM.createRoot(rootNode);

root.render(React.createElement(GithubAuth))