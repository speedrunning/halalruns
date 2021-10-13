.POSIX:

MANDIR	= /usr/share/man/man3

all:

.PHONY: check docs format
check:
	revive -formatter stylish *.go

docs:
	sed '1,/^$$/d' man/macros.tmac >tmp.tmac
	trap 'rm -f tmp.tmac tmp.3go' EXIT; \
	>/dev/null command -v gzip && has_gzip=true || has_gzip=false; \
	for manpage in man/*.3go; do \
		>tmp.3go sed -n -e '/^\.\s*so\s\s*macros\.tmac$$/! {p; d}' \
				-e 'r tmp.tmac' <$$manpage; \
		$$has_gzip && gzip -c9 tmp.3go >${MANDIR}/$${manpage##*/}.gz || \
		cp tmp.3go ${MANDIR}/; \
	done
	rm tmp.3go tmp.tmac

format:
	gofmt -w *.go
