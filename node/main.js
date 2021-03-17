async function job(n) {
    return new Promise(function(resolve, reject) {
        resolve(fibonacci(n))
    });
}

function fibonacci(n) {
    return n < 1 ? 0
         : n <= 2 ? 1
         : fibonacci(n - 1) + fibonacci(n - 2);
}

async function main() {
    const jobs = [];
    console.time('Runtime');
    for (let i = 0; i < 50; i++) {
        jobs.push(job(i));
    }
    await Promise.all(jobs);
    console.timeEnd('Runtime');
}

main();
