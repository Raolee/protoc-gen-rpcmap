# protoc-gen-rpcmap

## 작동 방식
protoc의 plugin으로 작동하는데, `protoc`가 위치해있는 경로에 `protoc-gen-[plugin이름]` 으로 실행파일이 존재해야 한다.
이러면 protoc가 args로 들어온 내용을 보고, 알아서 plugin을 찾는다.

ex) 내 plugin 이름이 raol일 때
`실행파일 이름 : protoc-gen-raol`
`protoc에 옵션 주는 방식 : protoc ...중략... --raol_out=./ --raol_opt=path=SOURCE_RELATIVE my.proto`

## 작업 방법
1. cmd/main.go 를 빌드한다. (내 pc의 protoc가 위치한 곳에 실행파일이 떨어지도록)
2. test/ 디렉토리 안에 있는 proto를 컴파일하면서, 내가 만든 플러그인이 잘 작동하는지 테스트 한다.