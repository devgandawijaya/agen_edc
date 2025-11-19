# TODO: Implement User API

## Steps to Complete

- [x] Add JWT and bcrypt dependencies to go.mod
- [x] Create User model in internal/models/user.go (with roles: user and admin)
- [x] Create migration for users table in migrations/002_create_users.sql
- [x] Create User repository in internal/repositories/user_repo.go
- [x] Create User service in internal/services/user_service.go (include CRUD, login, register, forgot password)
- [x] Create User controller in internal/controllers/user_controller.go
- [x] Add User routes in internal/routes/routes.go (CRUD with paging, login, register, forgot password)
- [x] Update Postman collection in postman_collection.json with User API endpoints
- [ ] Test the API endpoints
