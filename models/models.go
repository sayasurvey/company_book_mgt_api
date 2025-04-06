package models

import "time"

type User struct {
    ID                  int64     `json:"id"`
    Name                string    `json:"name"`
    Email               string    `json:"email"`
    Password            string    `json:"-"`
    Role                string    `json:"role"`
    Books               []Book    `json:"books,omitempty"`
    BorrowedBooks       []BorrowedBook `json:"borrowed_books,omitempty"`
    BorrowingWishLists  []BorrowingWishList `json:"borrowing_wish_lists,omitempty"`
    CreatedAt           time.Time `json:"created_at"`
    UpdatedAt           time.Time `json:"updated_at"`
}

type Book struct {
    ID        int64     `json:"id"`
    UserID    int64     `json:"user_id"`
    Title     string    `json:"title"`
    ImageURL  string    `json:"image_url"`
    Loanable  bool      `json:"loanable"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type BorrowedBook struct {
    ID            int64     `json:"id"`
    UserID        int64     `json:"user_id"`
    BookID        int64     `json:"book_id"`
    CheckoutDate  time.Time `json:"checkout_date"`
    ReturnDueDate time.Time `json:"return_due_date"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
}

type BorrowingWishList struct {
    ID        int64     `json:"id"`
    UserID    int64     `json:"user_id"`
    BookID    int64     `json:"book_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type InvalidatedToken struct {
    ID        int64     `json:"id"`
    Token     string    `json:"token"`
    ExpiresAt time.Time `json:"expires_at"`
    CreatedAt time.Time `json:"created_at"`
}

type LoginRequest struct {
    ID        int64     `json:"id"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    CreatedAt time.Time `json:"created_at"`
}
