import { LitElement, html } from 'lit';
import { customElement, property } from 'lit/decorators.js';
import { consume } from '@lit/context';
import { authContext, type AuthContextType } from '../context/auth';

@customElement('app-protected')
export class Protect extends LitElement {
	@consume({ context: authContext })
	private _auth!: AuthContextType;

	@property({ type: Boolean }) reverse = false;

	render() {
		if (!this.reverse && !this._auth?.isAuthenticated) {
			globalThis.router.goto('login');
			return null;
		}

		if (this.reverse && this._auth?.isAuthenticated) {
			globalThis.router.goto('dashboard');
			return null;
		}

		return html`<slot></slot>`;
	}
}
