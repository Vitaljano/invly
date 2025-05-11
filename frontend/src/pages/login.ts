import { html, css, LitElement } from 'lit';
import { consume } from '@lit/context';
import { authContext, type AuthContextType } from '../context/auth';

import { customElement, state } from 'lit/decorators.js';
import z from 'zod';

const loginSchema = z.object({
	email: z.string().nonempty({ message: 'Email is required' }).email({ message: 'Invalid email address' }),
	password: z
		.string()
		.nonempty({ message: 'Password is required' })
		.min(6, { message: 'Password must be at least 6 characters' }),
});

@customElement('app-login')
export class Login extends LitElement {
	static styles = css`
		:host {
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
		.validity-styles sl-input,
		.validity-styles sl-select,
		.validity-styles sl-checkbox {
			display: block;
			margin-bottom: var(--sl-spacing-medium);
		}

		/* user invalid styles */
		.validity-styles sl-input[data-user-invalid]::part(base) {
			border-color: var(--sl-color-danger-600);
		}

		.validity-styles [data-user-invalid]::part(form-control-label),
		.validity-styles [data-user-invalid]::part(form-control-help-text),
		.validity-styles sl-checkbox[data-user-invalid]::part(label) {
			color: var(--sl-color-danger-700);
		}

		.validity-styles sl-checkbox[data-user-invalid]::part(control) {
			outline: none;
		}

		.validity-styles sl-input:focus-within[data-user-invalid]::part(base) {
			border-color: var(--sl-color-danger-600);
			box-shadow: 0 0 0 var(--sl-focus-ring-width) var(--sl-color-danger-300);
		}

		/* User valid styles */
		.validity-styles sl-input[data-user-valid]::part(base) {
			border-color: var(--sl-color-success-600);
		}

		.validity-styles [data-user-valid]::part(form-control-label),
		.validity-styles [data-user-valid]::part(form-control-help-text),
		.validity-styles sl-checkbox[data-user-valid]::part(label) {
			color: var(--sl-color-success-700);
		}

		.validity-styles sl-checkbox[data-user-valid]::part(control) {
			background-color: var(--sl-color-success-600);
			outline: none;
		}

		.validity-styles sl-input:focus-within[data-user-valid]::part(base) {
			border-color: var(--sl-color-success-600);
			box-shadow: 0 0 0 var(--sl-focus-ring-width) var(--sl-color-success-300);
		}
	`;

	@state() private _formData = {
		email: '',
		password: '',
	};
	@state() private _errors: Partial<Record<keyof typeof loginSchema.shape, string>> = {};
	@consume({ context: authContext })
	private _auth!: AuthContextType;

	handleInput(e: Event) {
		const target = e.target as HTMLInputElement;
		const { name, value } = target;

		if (name in this._formData) {
			this._formData = { ...this._formData, [name]: value.trim() };

			if (this._errors[name as keyof typeof this._errors]) {
				const fieldSchema = loginSchema.shape[name as keyof typeof loginSchema.shape];
				const result = fieldSchema.safeParse(value.trim());

				if (!result.success) {
					this._errors = {
						...this._errors,
						[name]: result.error.errors[0].message,
					};
					target.setCustomValidity(result.error.errors[0].message);
				} else {
					this._errors = { ...this._errors, [name]: undefined };
					target.setCustomValidity('');
				}

				target.reportValidity();
			}
		}
	}

	handleSubmit(e: Event) {
		e.preventDefault();
		const result = loginSchema.safeParse(this._formData);
		this._errors = {};

		if (!result.success) {
			const fieldErrors = result.error.flatten().fieldErrors;

			Object.entries(fieldErrors).forEach(([field, errors]) => {
				const input = this.renderRoot?.querySelector(`sl-input[name=${field}]`) as HTMLInputElement;
				if (errors?.[0]) {
					this._errors = { ...this._errors, [field]: errors[0] };
					input?.setCustomValidity(errors[0]);
				} else {
					input?.setCustomValidity('');
				}
				input?.reportValidity();
			});
			return;
		}

		// Clear custom validity for all fields
		Object.keys(this._formData).forEach((field) => {
			const input = this.renderRoot?.querySelector(`sl-input[name="${field}"]`) as HTMLInputElement;
			input?.setCustomValidity('');
			input?.reportValidity();
		});

		this._auth.login(this._formData.email, this._formData.password);
	}

	render() {
		return html`
			<sl-card>
				<form @submit=${this.handleSubmit} class="validity-styles">
					<h2 class="title"><sl-icon name="home"></sl-icon>Login</h2>

					<div class="wrapper">
						<sl-input
							id="email-input"
							name="email"
							placeholder="Email"
							size="medium"
							.value=${this._formData.email}
							@input=${this.handleInput}
							help-text=${this._errors.email ?? ''}
						>
							<sl-icon name="envelope-at" slot="prefix"></sl-icon>
						</sl-input>
						<sl-input
							id="password-input"
							name="password"
							type="password"
							placeholder="Password"
							password-toggle
							.value=${this._formData.password}
							@input=${this.handleInput}
							help-text=${this._errors.password ?? ''}
						>
							<sl-icon name="key" slot="prefix"></sl-icon>
						</sl-input>
						<sl-button class="submit" variant="success" type="submit">
							<sl-icon slot="prefix" name="box-arrow-in-right" style="font-size: 20px;"></sl-icon>
							Login</sl-button
						>
					</div>
				</form>
			</sl-card>
		`;
	}
}
