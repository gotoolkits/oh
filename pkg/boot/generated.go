// Released under an MIT-style license. See LICENSE.

// Generated by generate.oh

package boot

var Script string = `
define caar: method (l) =: car: car l
define cadr: method (l) =: car: cdr l
define cdar: method (l) =: cdr: car l
define cddr: method (l) =: cdr: cdr l
define caaar: method (l) =: car: caar l
define caadr: method (l) =: car: cadr l
define cadar: method (l) =: car: cdar l
define caddr: method (l) =: car: cddr l
define cdaar: method (l) =: cdr: caar l
define cdadr: method (l) =: cdr: cadr l
define cddar: method (l) =: cdr: cdar l
define cdddr: method (l) =: cdr: cddr l
define caaaar: method (l) =: caar: caar l
define caaadr: method (l) =: caar: cadr l
define caadar: method (l) =: caar: cdar l
define caaddr: method (l) =: caar: cddr l
define cadaar: method (l) =: cadr: caar l
define cadadr: method (l) =: cadr: cadr l
define caddar: method (l) =: cadr: cdar l
define cadddr: method (l) =: cadr: cddr l
define cdaaar: method (l) =: cdar: caar l
define cdaadr: method (l) =: cdar: cadr l
define cdadar: method (l) =: cdar: cdar l
define cdaddr: method (l) =: cdar: cddr l
define cddaar: method (l) =: cddr: caar l
define cddadr: method (l) =: cddr: cadr l
define cdddar: method (l) =: cddr: cdar l
define cddddr: method (l) =: cddr: cddr l
define $connect: syntax (conduit name) = {
	set conduit: eval conduit
	syntax e (left right) = {
		define p: conduit
		spawn {
			e::eval: quasiquote: block {
				public (unquote name) = (unquote p)
				eval (unquote left)
			}
			p::writer-close
		}
		block {
			e::eval: quasiquote: block {
				public $stdin = (unquote p)
				eval (unquote right)
			}
			p::reader-close
		}
	}
}
define $redirect: syntax (name mode closer) = {
	syntax e (c cmd) = {
		define c: e::eval c
		define f = ()
		if (not: or (is-channel c) (is-pipe c)) {
			set f: open mode c
			set c = f
		}
		e::eval: quasiquote: block {
			public (unquote name) (unquote c)
			eval (unquote cmd)
		}
		if (not: is-null f): eval: quasiquote: f::(unquote closer)
	}
}
define ...: method (: args) = {
	cd $origin
	define path: car args
	if (eq 2: length args) {
		cd: car args
		set path: cadr args
	}
	while true {
		define abs: symbol: "/"::join $cwd path
		if (exists abs): return abs
		if (eq $cwd /): return path
		cd ..
	}
}
define and: syntax e (: lst) = {
	define r = false
	while (not: is-null: car lst) {
		set r: e::eval: car lst
		if (not r): return r
		set lst: cdr lst
	}
	return r
}
define append-stderr: $redirect $stderr "a" writer-close
define append-stdout: $redirect $stdout "a" writer-close
define apply: method (f: args) =: f @args
define backtick: syntax e (cmd) = {
	define p: pipe
	spawn {
		e::eval: quasiquote: block {
			public $stdout = (unquote p)
			eval (unquote cmd)
		}
		p::writer-close
	}
	define r: cons () ()
	define c = r
	while (define l: p::readline) {
		set-cdr c: cons l ()
		set c: cdr c
	}
	p::reader-close
	return: cdr r
}
define catch: syntax e (name: clause) = {
	define args: list name (quote throw)
	define body: list (quote throw) name
	if (is-null clause) {
		set body: list body
	} else {
		set body: append clause body
	}
	define handler: e::eval {
		list (quote method) args (quote =) @body
	}
	define _return: e::eval (quote return)
	define _throw = throw
	e::public throw: method (condition) = {
		_return: handler condition _throw
	}
}
define channel-stderr: $connect channel $stderr
define channel-stdout: $connect channel $stdout
define echo: builtin (: args) = {
	if (is-null args) {
		$stdout::write: symbol ""
	} else {
		$stdout::write @(for args symbol)
	}
}
define error: builtin (: args) =: $stderr::write @args
public exception: method (_type _message _status _file _line) = {
	object {
		public type = _type
		public status = _status
		public message = _message
		public line = _line
		public file = _file
	}
}
define for: method (l m) = {
	define r: cons () ()
	define c = r
	while (not: is-null l) {
		set-cdr c: cons (m: car l) ()
		set c: cdr c
		set l: cdr l
	}
	return: cdr r
}
define glob: builtin (: args) =: return args
define import: syntax e (name) = {
	set name: e::eval name
	define m: module name
	if (or (is-null m) (is-object m)) {
		return m
	}

	e::eval: quasiquote: $root::define (unquote m): object {
		source (unquote name)
	}
}
define is-list: method (l) = {
	if (is-null l): return false
	if (not: is-cons l): return false
	if (is-null: cdr l): return true
	is-list: cdr l
}
define is-text: method (t) =: or (is-string t) (is-symbol t)
define list-ref: method (k x) =: car: list-tail k x
define list-tail: method (k x) = {
	if k {
		list-tail (sub k 1): cdr x
	} else {
		return x
	}
}
define object: syntax e (: body) = {
	e::eval: cons (quote block): append body (quote: context)
}
define or: syntax e (: lst) = {
	define r = false
	while (not: is-null: car lst) {
		set r: e::eval: car lst
		if r: return r
		set lst: cdr lst 
	}
	return r
}
define pipe-stderr: $connect pipe $stderr
define pipe-stdout: $connect pipe $stdout
define printf: method (f: args) =: echo: f::sprintf @args
define process-substitution: syntax e (:args) = {
	define fifos = ()
	define procs = ()
	define cmd: for args: method (arg) = {
		if (not: is-cons arg): return arg
		if (eq (quote substitute-stdin) (car arg)) {
			define fifo: temp-fifo
			define proc: spawn {
				e::eval: cdr arg < fifo
			}
			set fifos: cons fifo fifos
			set procs: cons proc procs
			return fifo
		}
		if (eq (quote substitute-stdout) (car arg)) {
			define fifo: temp-fifo
			define proc: spawn {
				e::eval: cdr arg > fifo
			}
			set fifos: cons fifo fifos
			set procs: cons proc procs
			return fifo
		}
		return arg
	}
	e::eval cmd
	wait @procs
	rm @fifos
}
public prompt: method (suffix) = {
	define folders: (string $cwd)::split "/"
	define last: sub (length folders) 1
	return (list-ref last folders) ^ suffix
}
define quasiquote: syntax e (cell) = {
	if (not: is-cons cell): return cell
	if (is-null cell): return cell
	if (eq (quote unquote): car cell): return: e::eval: cadr cell
	cons {
		e::eval: list (quote quasiquote): car cell
		e::eval: list (quote quasiquote): cdr cell
	}
}
define quote: syntax (cell) =: return cell
define read: builtin () =: $stdin::read
define readline: builtin () =: $stdin::readline
define redirect-stderr: $redirect $stderr "w" writer-close
define redirect-stdin: $redirect $stdin "r" reader-close
define redirect-stdout: $redirect $stdout "w" writer-close
define source: syntax e (name) = {
	define basename: e::eval name
	define paths = ()
	define name = basename

	if (has $OHPATH): set paths: (string $OHPATH)::split ":"
	while (and (not: is-null paths) (not: exists name)) {
		set name: "/"::join (car paths) basename
		set paths: cdr paths
	}

	if (not: exists name): set name = basename

	define f: open r- name

	define r: cons () ()
	define c = r
	while (define l: f::read) {
		set-cdr c: cons (cons (get-line-number) l) ()
		set c: cdr c
	}
	set c: cdr r
	f::close

	define rval: status 0
	define eval-list: syntax o (first rest) = {
		set first: o::eval first
		set rest: o::eval rest
		if (is-null first): return rval
		set-line-number: car first
		set rval: e::eval: cdr first
		eval-list (car rest) (cdr rest)
	}
	eval-list (car c) (cdr c)
	return rval
}
public throw: method (c) = {
	error: ": "::join c::file c::line c::type c::message
	fatal c::status
}
define write: method (: args) =: $stdout::write @args
define sys: object {
	public get-prompt: method (suffix) = {
		catch unused {
			return suffix
		}
		prompt suffix
	}
}

exists ("/"::join $HOME .ohrc) && source ("/"::join $HOME .ohrc)

`

//go:generate ./generate.oh
//go:generate go fmt generated.go
