GO = $(shell which go)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
HASH = $(shell git rev-parse HEAD)
VERSION = $(shell TZ=UTC git show -s --abbrev=12 --date=format-local:%Y%m%d%H%M%S --pretty=format:%cd-%h HEAD)
LDFLAGS = -ldflags '-X main.Version=v0.0.0-$(BRANCH)-$(VERSION)'
# LDFLAGS = -ldflags="-X 'main.Version=v1.0.0'"
GOBUILD = $(GO) build $(LDFLAGS) -o

help : Makefile
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## cleanBin: Удалить директорию с бинарниками
cleanBin:
	@rm -rf _bin

## build: Сбилдить приложение (publisher, consumer and main)
build: cleanBin
	$(GOBUILD) _bin/pub ./cmd/publisher
	$(GOBUILD) _bin/cons ./cmd/consumer
	$(GOBUILD) _bin/main .

## publish: Опубликовать сообщение (сообщение зашито)
publish:
	@./_bin/pub

## consume: Получить сообщения из канала (имя канала зашито)
consume:
	@./_bin/cons

## main: Запустить программу, которая выводит версию приложения
main:
	@./_bin/main
