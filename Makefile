build:
	GOOS=darwin GOARCH=amd64 go build -o sysinfo_macos_amd64-v0_0_1
	GOOS=linux GOARCH=amd64 go build -o sysinfo_linux_amd64-v0_0_1
	GOOS=windows GOARCH=amd64 go build -o sysinfo_windows_amd64-v0_0_1.exe
deploy:
	aws s3 cp ./sysinfo_macos_amd64-v0_0_1 s3://bss-sharedsoftware/util/sysinfo/sysinfo_macos_amd64-v0_0_1 --profile shared
	aws s3 cp ./sysinfo_linux_amd64-v0_0_1 s3://bss-sharedsoftware/util/sysinfo/sysinfo_linux_amd64-v0_0_1  --profile shared
	aws s3 cp ./sysinfo_windows_amd64-v0_0_1.exe s3://bss-sharedsoftware/util/sysinfo/sysinfo_windows_amd64-v0_0_1.exe --profile shared
replace:
	sysinfo replace --file ./test/config.groovy.template