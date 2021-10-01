.POSIX:

MANDIR	= /usr/share/man/man3

all:

.PHONY: check format
check:
	revive -formatter stylish *.go

docs:
	>/dev/null command -v gzip && has_gzip=true || has_gzip=false; \
	for manpage in man/*; do \
		$$has_gzip && gzip -c9 $$manpage >${MANDIR}/$${manpage##*/}.gz || \
		cp $$manpage ${MANDIR}/; \
	done

format:
	gofmt -w *.go
