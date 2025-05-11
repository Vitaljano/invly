import { Router } from '@lit-labs/router';
import { LitElement, html, css } from 'lit';
import { customElement } from 'lit/decorators.js';

import './components/navigation';
import './pages/login';

import 'urlpattern-polyfill';
import '@shoelace-style/shoelace/dist/components/icon/icon.js';
import '@shoelace-style/shoelace/dist/shoelace.js';
import '@shoelace-style/shoelace/dist/themes/dark.css';

import { setBasePath } from '@shoelace-style/shoelace/dist/utilities/base-path.js';

setBasePath('/node_modules/@shoelace-style/shoelace/dist');
// import '@shoelace-style/shoelace/dist/themes/light.css';

@customElement('app-root')
export class App extends LitElement {
	static styles = css``;
	private _routes = new Router(this, [
		{
			path: '/',
			render: () => html`<sl-button variant="primary">Home</sl-button>`,
		},
		{ path: '/login', render: () => html`<app-login></app-login>` },
		{ path: '/about', render: () => html`<h1>About</h1>` },
	]);

	render() {
		return html`<app-nav></app-nav>
			<main class="main">${this._routes.outlet()}</main> `;
	}
}
