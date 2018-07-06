module.exports = {
    apps: [
        {
            name: "goapp",
            script: "./config.json",
            instances: 1,
            exec_mode: "fork",
            interpreter: "./main",
        }
    ]
}