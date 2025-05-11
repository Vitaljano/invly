import { LitElement, html, css } from 'lit';
import { customElement } from 'lit/decorators.js';

@customElement('app-logo')
export class Logo extends LitElement {
	static styles = css`
		div {
			display: flex;
			gap: 0.5rem;
		}

		.logo {
			background: linear-gradient(#bd317d, #f49f4c);
			border-radius: 50%;
			display: grid;
			place-items: center;
			width: 3rem;
			height: 3rem;
			color: #fff;
			rotate: 45deg;
			font-size: 2rem;
			font-weight: 700;
		}

		.text {
			font-size: 2rem;
		}
	`;
	render() {
		return html`
			<div>
				<span class="logo">U</span>
				<span class="text">invly.</span>
			</div>
		`;
	}
}
