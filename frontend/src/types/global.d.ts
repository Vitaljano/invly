export {};

interface Router {
	goto: (path: string) => void;
}

declare global {
	// eslint-disable-next-line no-var
	var router: Router;
}
