_NAME=agent

compile: 
	for os in darwin linux; do \
		for arch in amd64 arm64; do \
			GOOS=$$os GOARCH=$$arch go build -o bin/$(_NAME) . ; \
			cd bin ; \
			tar czf $(_NAME)-$$os-$$arch.tar.gz $(_NAME) ; \
			sha256sum $(_NAME)-$$os-$$arch.tar.gz > $(_NAME)-$$os-$$arch.tar.gz.sha256 ; \
			cd - ; \
		done ; \
	done ; \
	for arm in 5 6 7; do \
		GOOS=linux GOARCH=arm GOARM=$$arm go build -o bin/$(_NAME) . ; \
		cd bin ; \
		tar czf $(_NAME)-linux-armv$$arm.tar.gz $(_NAME) ; \
		sha256sum $(_NAME)-linux-armv$$arm.tar.gz > $(_NAME)-linux-armv$$arm.tar.gz.sha256 ; \
		cd - ; \
	done

clean:
	rm bin/$(_NAME)*
