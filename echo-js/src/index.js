// XXX even though ethers is not used in the code below, it's very likely
// it will be used by any DApp, so we are already including it here
const { ethers } = require("ethers");

const rollup_server = process.env.ROLLUP_HTTP_SERVER_URL;
console.log("HTTP rollup_server url is " + rollup_server);

async function handle_advance(data) {
    console.log("Received advance request data " + JSON.stringify(data));
    const payload = data["payload"];
    try {
        const payloadStr = ethers.utils.toUtf8String(payload);
        console.log(`Adding notice "${payloadStr}"`);
    } catch (e) {
        console.log(`Adding notice with binary value "${payload}"`);
    }
    const advance_req = await fetch(rollup_server + '/notice', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ payload })
    });
    const json = await advance_req.json();
    console.log("Received notice status " + advance_req.status + " with body " + JSON.stringify(json));
    return "accept";
}

async function handle_inspect(data) {
    console.log("Received inspect request data " + JSON.stringify(data));
    const payload = data["payload"];
    try {
        const payloadStr = ethers.utils.toUtf8String(payload);
        console.log(`Adding report "${payloadStr}"`);
    } catch (e) {
        console.log(`Adding report with binary value "${payload}"`);
    }
    const inspect_req = await fetch(rollup_server + '/report', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ payload })
    });
    console.log("Received report status " + inspect_req.status);
    return "accept";
}

var handlers = {
    advance_state: handle_advance,
    inspect_state: handle_inspect,
}

var finish = { status: "accept" };

(async () => {
    while (true) {
        console.log("Sending finish")

        const finish_req = await fetch(rollup_server + '/finish', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ status: 'accept' })
        });

        console.log("Received finish status " + finish_req.status);


        if (finish_req.status == 202) {
            console.log("No pending rollup request, trying again");
        } else {
            const rollup_req = await finish_req.json();
            var handler = handlers[rollup_req["request_type"]];
            finish["status"] = await handler(rollup_req["data"]);

        }
    }
})();