Object.defineProperty(Object.getPrototypeOf(navigator), 'webdriver', {
    get: () => undefined
});
