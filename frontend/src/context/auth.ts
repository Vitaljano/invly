import { createContext, provide } from '@lit/context';
import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators.js';

export interface AuthContextType {
	user: { name: string } | null;
	login: (username: string, password: string) => Promise<void>;
	logout: () => void;
	isAuthenticated: boolean;
}

export const authContext = createContext<AuthContextType>(Symbol('auth'));

@customElement('auth-provider')
export class AuthProvider extends LitElement {
	@provide({ context: authContext })
	auth: AuthContextType;

	constructor() {
		super();

		this.auth = {
			user: { name: 'user' },
			login: this.login.bind(this),
			logout: this.logout.bind(this),
			isAuthenticated: false,
		};
	}

	login = async (username: string, password: string) => {
		console.log('Not implemented login', username, password);

		this.auth = {
			...this.auth,
			isAuthenticated: true,
		};
	};
	logout = () => {
		console.log('Not implemented');
	};

	render() {
		return html`<slot></slot>`;
	}
}
