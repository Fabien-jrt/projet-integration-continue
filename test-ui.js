const puppeteer = require('puppeteer');

(async () => {
    const browser = await puppeteer.launch();
    const page = await browser.newPage();

    // Test for index.html
    await page.goto('http://localhost:1323');  // Change the URL to match your local server

    // Check for the presence of "Hello, World!" heading
    const heading = await page.$eval('h1', el => el.textContent);
    if (heading !== 'Hello, World!') {
        console.error('Test failed: Heading is not "Hello, World!"');
        process.exit(1);
    }

    // Check for the presence of the link to create a new user
    const linkExists = await page.$('a[href="/user"]') !== null;
    if (!linkExists) {
        console.error('Test failed: Link to create a new user not found.');
        process.exit(1);
    }

    console.log('index.html tests passed!');

    // Test for user.html
    await page.goto('http://localhost:1323/user');  // Change the URL to match your local server

    // Check for the presence of "Email validation" heading
    const headingUser = await page.$eval('h1', el => el.textContent);
    if (headingUser !== 'Email validation') {
        console.error('Test failed: Heading is not "Email validation"');
        process.exit(1);
    }

    // Check for the presence of the email input field
    const emailInputExists = await page.$('input[type="email"]#email') !== null;
    if (!emailInputExists) {
        console.error('Test failed: Email input field not found.');
        process.exit(1);
    }

    // Check for the presence of the submit button
    const submitButtonExists = await page.$('input[type="submit"]') !== null;
    if (!submitButtonExists) {
        console.error('Test failed: Submit button not found.');
        process.exit(1);
    }

    console.log('user.html tests passed!');
    await browser.close();
})();
