import { html, LitElement } from 'lit';
import { customElement } from 'lit/decorators.js';
import { Router } from '@lit-labs/router';

import "./components/navigation"

import "urlpattern-polyfill";
import '@shoelace-style/shoelace/dist/shoelace.js';
import '@shoelace-style/shoelace/dist/themes/dark.css';
// import '@shoelace-style/shoelace/dist/themes/light.css';



@customElement('app-root')
export class App extends LitElement {
private _routes = new Router(this,[
  {path: '/', render: () => html`<sl-button variant="primary" >Home</sl-button>`},
  {path: '/projects', render: () => html`<h1>Projects</h1>`},
  {path: '/about', render: () => html`<h1>About</h1>`},
])

  render() {
    return html`<app-nav><app-nav><main>${this._routes.outlet()}</main>
`;
  }
}
