import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators.js';

@customElement('user-profile')
export class Profile extends LitElement {
	render() {
		return html`<sl-avatar label="User avatar"></sl-avatar>`;
	}
}
