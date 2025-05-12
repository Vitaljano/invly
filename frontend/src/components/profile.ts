import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators.js';
import { authContext, type AuthContextType } from '../context/auth';
import { consume } from '@lit/context';

@customElement('user-profile')
export class Profile extends LitElement {
	@consume({ context: authContext })
	private _auth!: AuthContextType;
	handleLogout() {
		if (this._auth.isAuthenticated) {
			this._auth.logout();
		}
	}

	render() {
		return html`
			<sl-dropdown distance="10">
				<sl-button slot="trigger" variant="default" size="medium" circle>
					<sl-icon name="person" label="profile" style="font-size:18px"></sl-icon>
				</sl-button>
				<sl-menu>
					<sl-menu-item value="paste" @click=${this.handleLogout}>Logout</sl-menu-item>
				</sl-menu>
			</sl-dropdown>
		`;
	}
}
