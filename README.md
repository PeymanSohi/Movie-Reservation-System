# Movie Reservation Service

A backend service for managing movie reservations, built with Go and PostgreSQL.

## Features

- User authentication and authorization
- Movie management
- Showtime scheduling
- Seat reservation
- Reservation management
- Admin dashboard

## Prerequisites

- Go 1.21 or later
- PostgreSQL
- Make (optional)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/peyman/movie-res.git
cd movie-res
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the root directory with the following variables:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=movie_res
JWT_SECRET=your-secret-key-change-this-in-production
SERVER_PORT=8080
ADMIN_EMAIL=admin@example.com
ADMIN_PASSWORD=admin123
```

4. Create the database:
```bash
createdb movie_res
```

5. Run the application:
```bash
go run cmd/api/main.go
```

## API Endpoints

### Public Endpoints

- `POST /api/register` - Register a new user
- `POST /api/login` - Login and get JWT token
- `GET /api/movies` - Get all movies
- `GET /api/movies/date` - Get movies by date
- `GET /api/movies/:id` - Get movie details
- `GET /api/showtimes/:showtime_id/seats` - Get available seats for a showtime

### Protected Endpoints (Requires Authentication)

- `GET /api/profile` - Get user profile
- `PUT /api/profile` - Update user profile
- `POST /api/reservations` - Create a reservation
- `GET /api/reservations` - Get user's reservations
- `GET /api/reservations/:id` - Get reservation details
- `DELETE /api/reservations/:id` - Cancel a reservation

### Admin Endpoints (Requires Admin Role)

- `GET /api/admin/users` - List all users
- `POST /api/admin/users/:id/promote` - Promote user to admin
- `POST /api/admin/movies` - Create a movie
- `PUT /api/admin/movies/:id` - Update a movie
- `DELETE /api/admin/movies/:id` - Delete a movie
- `POST /api/admin/movies/:id/showtimes` - Add a showtime
- `GET /api/admin/reservations` - List all reservations
- `GET /api/admin/revenue` - Get revenue report

## Authentication

The API uses JWT (JSON Web Tokens) for authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your_token>
```

## Database Schema

The application uses the following main entities:

- Users
- Movies
- Showtimes
- Seats
- Reservations

## Error Handling

The API returns appropriate HTTP status codes and error messages in the following format:

```json
{
  "error": "Error message"
}
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 