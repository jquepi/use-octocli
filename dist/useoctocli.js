const core = require('@actions/core');
const github = require('@actions/github');

async function run() {
    try {
        const project = core.getInput('project')
        const octopus_server = core.getInput('octopus-server')
        const api_key = core.getInput('api-key')
        const space = core.getInput('space')

    } catch (error) {
        core.setFailed(error.message);
    }
}

run()