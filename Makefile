# makefile 関数
# https://qiita.com/chibi929/items/b8c5f36434d5d3fbfa4a

REPO_FULL=$(shell git rev-parse --show-toplevel)
REPO_DIR=$(shell basename $(shell dirname $(shell dirname $(REPO_FULL))))
REPO_AUTHOR=$(shell basename $(shell dirname $(REPO_FULL)))
REPO_NAME=$(shell basename $(REPO_FULL))
CMD_NAME=$(shell basename $(shell pwd))

FIRST_GOPATH=$(shell echo $(GOPATH) | awk -F ":" '{ print $1 }')
APP_ROOT=$(abspath ./)
WORK_ROOT=$(FIRST_GOPATH)/src/$(REPO_DIR)/$(REPO_AUTHOR)
WORK_DIR=$(WORK_ROOT)/$(REPO_NAME)


# コマンド一覧
.PHONY: setup deps build clean debug

deps:
		cd $(WORK_DIR)
		dep ensure -v

build: deps
		$(WORK_DIR)/build/crosscompile_for_go.sh -o $(CMD_NAME)

setup:
ifndef APP_ROOT
		#未定義の場合
		$(info $(APP_ROOT)"が存在しないため、setup完了していません")
else
		ln -s $(APP_ROOT) $(WORK_ROOT)
		$(info "gopath配下にsetupしました at $(WORK_ROOT)  with GOPATH=$(GOPATH)")
		$(shell cd $(WORK_DIR); dep init)
endif

# TODO: linux/mac向けのバイナリ削除
clean:
	  rm -rf *.exe

debug:
		$(info repo:    $(REPO_DIR)/$(REPO_AUTHOR)/$(REPO_NAME))
		$(info gopath:  $(FIRST_GOPATH))
		$(info workdir: $(WORK_DIR))
		$(info appRoot:     $(APP_ROOT))
		$(info cmd: $(CMD_NAME))