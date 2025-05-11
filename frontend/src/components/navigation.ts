import { LitElement, css, html } from 'lit';
import { customElement } from 'lit/decorators.js';
import { consume } from '@lit/context';
import { authContext, type AuthContextType } from '../context/auth';

import './logo';
import './profile';

@customElement('app-nav')
export class Navigation extends LitElement {
	static styles = css`
		header {
			padding-inline: 2rem;
			height: 4rem;
			display: flex;
			place-items: center;
			justify-content: space-between;
			background: #1e1e1e;
		}

		nav {
			display: flex;
		}
	`;
	@consume({ context: authContext, subscribe: true })
	private _auth!: AuthContextType;

	render() {
		console.log('rerender', this._auth);

		return html` <header>
			<app-logo></app-logo>
			${this._auth.isAuthenticated
				? html`<user-profile></user-profile>`
				: html`
						<div>
							<sl-button href="/login" variant="text" size="medium">Sign in</sl-button>
							<sl-button href="/" variant="default" size="medium">Sign Up</sl-button>
						</div>
					`}
		</header>`;
	}
}
