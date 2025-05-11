import { html, css, LitElement } from 'lit';
import { customElement } from 'lit/decorators.js';

@customElement('app-login')
export class Login extends LitElement {
	static styles = css`
		.container {
			min-height: calc(100dvh - 4rem);
			display: grid;
			place-items: center;
		}
		.title {
			text-align: center;
		}
		.wrapper {
			display: flex;
			flex-direction: column;
			gap: 0.5rem;
		}
		.wrapper > :last-child {
			margin-top: 1rem;
		}
	`;
	render() {
		return html` <div class="container">
			<sl-card>
				<h2 class="title"><sl-icon name="home"></sl-icon>Login</h2>

				<div class="wrapper">
					<sl-input placeholder="Email" size="medium">
						<sl-icon name="envelope-at" slot="prefix"></sl-icon>
					</sl-input>
					<sl-input type="password" placeholder="Password" password-toggle>
						<sl-icon name="key" slot="prefix"></sl-icon>
					</sl-input>
					<sl-button class="submit" variant="success">
						<sl-icon slot="prefix" name="box-arrow-in-right"></sl-icon>
						Login</sl-button
					>
				</div>
			</sl-card>
			<div></div>
		</div>`;
	}
}
