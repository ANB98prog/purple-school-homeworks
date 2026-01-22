package service

import (
	helper "github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/helpers/auth"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/helpers/file"
)

const codesFileName = "authCodes.json"

type AuthCodeService interface {
	GenerateAuthCode(phone string) (AuthCode, error)
	GetAuthCode(sessionId string) (AuthCode, bool)
	DeleteAuthCode(sessionId string) error
}

type authCodeService struct{}

func NewAuthCodeService() AuthCodeService {
	return &authCodeService{}
}

var _ AuthCodeService = (*authCodeService)(nil) // assertion

func (service *authCodeService) GenerateAuthCode(phone string) (AuthCode, error) {

	// Генерируем код и сессию
	code := helper.GenerateAuthCode()
	sessionId := helper.GenerateSessionId()

	// Читаем из файла коды
	authCodes, err := getCodesFromFile()
	if err != nil {
		return AuthCode{}, err
	}

	// Добавляем или обновляем код для пользователя
	authCode := authCodes.Upsert(sessionId, code, phone)

	// Сохраняем коды
	err = file.WriteFile(codesFileName, &authCodes)
	if err != nil {
		return AuthCode{}, err
	}

	// Возвращаем код
	return authCode, nil
}

func (service *authCodeService) GetAuthCode(sessionId string) (AuthCode, bool) {
	// Достаем коды из файла
	authCodes, err := getCodesFromFile()
	if err != nil {
		return AuthCode{}, false
	}

	// Ищем по идентификатору сессии
	authSession, ok := authCodes.GetBySessionId(sessionId)
	if !ok {
		return AuthCode{}, false
	}

	return authSession, true
}

func (service *authCodeService) DeleteAuthCode(sessionId string) error {
	// Достаем коды из файла
	authCodes, err := getCodesFromFile()
	if err != nil {
		return err
	}

	authCodes.Delete(sessionId)

	// Сохраняем коды
	err = file.WriteFile(codesFileName, &authCodes)
	if err != nil {
		return err
	}

	return nil
}

func getCodesFromFile() (*AuthCodes, error) {
	authCodes, err := file.ReadFile[AuthCodes](codesFileName)

	if err != nil {
		return nil, err
	}

	if authCodes == nil {
		return &AuthCodes{}, nil
	}

	return authCodes, nil
}
