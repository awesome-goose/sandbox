import { Component, OnInit, signal } from '@angular/core';

interface VersionInfo {
  name: string;
  version: string;
}

interface User {
  id: number;
  name: string;
  email: string;
}

@Component({
  selector: 'app-root',
  standalone: true,
  template: `
    <main class="app">
      <h1>Goose + Angular</h1>
      <p>
        This page is served by the Goose SPA platform. The backend answers
        JSON under <code>/api</code>.
      </p>
      @if (info(); as i) {
        <p class="status">
          Backend says: <strong>{{ i.name }}</strong> v{{ i.version }}
        </p>
      }
      @if (users().length > 0) {
        <h2>Users from /api/users</h2>
        <ul class="users">
          @for (user of users(); track user.id) {
            <li>
              <strong>{{ user.name }}</strong> — {{ user.email }}
            </li>
          }
        </ul>
      }
      @if (error(); as e) {
        <p class="error">Backend unreachable: {{ e }}</p>
      }
    </main>
  `,
})
export class AppComponent implements OnInit {
  readonly info = signal<VersionInfo | null>(null);
  readonly users = signal<User[]>([]);
  readonly error = signal<string | null>(null);

  ngOnInit(): void {
    fetch('/api/version')
      .then((res) => res.json())
      .then((data: VersionInfo) => this.info.set(data))
      .catch((err) => this.error.set(String(err)));

    fetch('/api/users')
      .then((res) => res.json())
      .then((data: User[]) => this.users.set(data))
      .catch((err) => this.error.set(String(err)));
  }
}
