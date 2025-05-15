import { LitElement, html } from 'lit';
import { customElement, state } from 'lit/decorators.js';

const themeClass = 'sl-theme-dark';

@customElement('app-theme-switch')
export class ThemeSwitcher extends LitElement {
	@state()
	theme: 'dark' | 'light' = 'light';

	private _observer?: MutationObserver;

	connectedCallback(): void {
		super.connectedCallback();
		const el = document.querySelector('html');
		if (!el) return;

		this.theme = el?.classList.contains(themeClass) ? 'dark' : 'light';

		this._observer = new MutationObserver(() => {
			this.theme = el?.classList.contains(themeClass) ? 'dark' : 'light';
		});

		this._observer.observe(el, { attributes: true, attributeFilter: ['class'] });
	}

	disconnectedCallback(): void {
		this._observer?.disconnect();
		super.disconnectedCallback();
	}

	handleSwitch() {
		const el = document.querySelector('html');
		if (!el) return;
		if (this.theme === 'dark') {
			el.classList.remove(themeClass);
			this.theme = 'light';
		} else {
			el?.classList.add(themeClass);
			this.theme = 'dark';
		}
	}

	render() {
		const isDark = this.theme === 'dark';
		const icon = isDark ? 'moon-stars' : 'sun';
		const label = isDark ? 'Switch to light theme' : 'Switch to dark theme';
		return html`
			<sl-icon-button
				style="font-size: 1.2rem"
				name=${icon}
				@click=${this.handleSwitch}
				label=${label}
			></sl-icon-button>
		`;
	}
}
