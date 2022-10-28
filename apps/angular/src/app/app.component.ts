import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { environment } from '../environments/environment';

@Component({
  selector: 'developer-first-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  goStatus$ = this._http.get(environment.goEndpoint);
  nestjsStatus$ = this._http.get(environment.nestjsEndpoint);
  dotnetStatus$ = this._http.get(environment.dotnetEndpoint);

  constructor(private _http: HttpClient) {}
}
