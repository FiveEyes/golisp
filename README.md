# xlisp
Lisp interpreter in Go

## The idea

Lisp has two interesting features:
  * Everything is function call.
  * Quote is an unique operator in Lisp.

When I was learning Lisp, I was wondering if there exists a mini interpreter, such that all operators are normal native functions binding to keywords, instead of ad-hoc build-in functions.

I never implement this idea, because it would be another lisp interpreter, not very interesting to me.

During the spring break, I learnt a little bit of Golang. It also has some unique features, build-in multithreading and channels. Unlike coroution(continuation) in Lua, goroution is real thread and there is a scheduler hiding in runtime of Golang. It will be fun if we introduce goroution into lisp.

It's just a small imterpreter now, no marco, no coroution, and on goroution.

## Todo list
  * operator lazy
  * goroution
  * operator future and promise
  * coroution
  * marco
  
