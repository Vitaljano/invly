import { LitElement, css, html } from 'lit';
import { customElement } from 'lit/decorators.js';

import './logo';

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
	render() {
		return html` <header>
			<app-logo></app-logo>
			<div>
				<sl-button href="/login" variant="text" size="medium"
					>Sing in</sl-button
				>
				<sl-button href="/" variant="default" size="medium">Sing Up</sl-button>
			</div>
		</header>`;
	}
}
