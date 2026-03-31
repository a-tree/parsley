package api

import (
	"backend/internal/domain/models"
	"backend/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler は ServerInterface を実装する構造体です
type Handler struct {
	repo repository.UserRepository
}

// NewHandler はハンドラーを初期化します
func NewHandler(ur repository.UserRepository) *Handler {
	return &Handler{repo: ur}
}

// インターフェースが正しく実装されているかコンパイル時にチェックします
var _ ServerInterface = (*Handler)(nil)

// 全ユーザーの取得
func (h *Handler) ListUsers(ctx echo.Context) error {
	users, err := h.repo.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, users)
	}
	return ctx.JSON(http.StatusOK, users)
}

// 新規ユーザーの登録
func (h *Handler) CreateUser(ctx echo.Context) error {
	user := new(models.User)
	if err := ctx.Bind(user); err != nil {
		return err
	}
	if err := h.repo.Create(user); err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, user)
}

// 特定ユーザーの取得
func (h *Handler) GetUserById(ctx echo.Context, id uint) error {
	user, err := h.repo.FindUser(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, user)
	}
	return ctx.JSON(http.StatusOK, user)
}

// 登録内容の変更
func (h *Handler) UpdateUser(ctx echo.Context, id uint) error {
	return ctx.JSON(http.StatusOK, models.User{})
}

// ユーザーの削除
func (h *Handler) DeleteUser(ctx echo.Context, id uint) error {
	return ctx.JSON(http.StatusOK, models.User{})
}
