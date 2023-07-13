// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2023-07-13 16:10:33.8303184 +0200 CEST m=+0.008901701
package option

// optionize converts a nullary function to an option
func optionize[R any](f func() (R, bool)) Option[R] {
	if r, ok := f(); ok {
		return Some(r)
	}
	return None[R]()
}

// Optionize0 converts a function with 0 parameters returning a tuple of a return value R and a boolean into a function with 0 parameters returning an Option[R]
func Optionize0[F ~func() (R, bool), R any](f F) func() Option[R] {
	return func() Option[R] {
		return optionize(func() (R, bool) {
			return f()
		})
	}
}

// Unoptionize0 converts a function with 0 parameters returning a tuple of a return value R and a boolean into a function with 0 parameters returning an Option[R]
func Unoptionize0[F ~func() Option[R], R any](f F) func() (R, bool) {
	return func() (R, bool) {
		return Unwrap(f())
	}
}

// Optionize1 converts a function with 1 parameters returning a tuple of a return value R and a boolean into a function with 1 parameters returning an Option[R]
func Optionize1[F ~func(T0) (R, bool), T0, R any](f F) func(T0) Option[R] {
	return func(t0 T0) Option[R] {
		return optionize(func() (R, bool) {
			return f(t0)
		})
	}
}

// Unoptionize1 converts a function with 1 parameters returning a tuple of a return value R and a boolean into a function with 1 parameters returning an Option[R]
func Unoptionize1[F ~func(T0) Option[R], T0, R any](f F) func(T0) (R, bool) {
	return func(t0 T0) (R, bool) {
		return Unwrap(f(t0))
	}
}

// Optionize2 converts a function with 2 parameters returning a tuple of a return value R and a boolean into a function with 2 parameters returning an Option[R]
func Optionize2[F ~func(T0, T1) (R, bool), T0, T1, R any](f F) func(T0, T1) Option[R] {
	return func(t0 T0, t1 T1) Option[R] {
		return optionize(func() (R, bool) {
			return f(t0, t1)
		})
	}
}

// Unoptionize2 converts a function with 2 parameters returning a tuple of a return value R and a boolean into a function with 2 parameters returning an Option[R]
func Unoptionize2[F ~func(T0, T1) Option[R], T0, T1, R any](f F) func(T0, T1) (R, bool) {
	return func(t0 T0, t1 T1) (R, bool) {
		return Unwrap(f(t0, t1))
	}
}

// Optionize3 converts a function with 3 parameters returning a tuple of a return value R and a boolean into a function with 3 parameters returning an Option[R]
func Optionize3[F ~func(T0, T1, T2) (R, bool), T0, T1, T2, R any](f F) func(T0, T1, T2) Option[R] {
	return func(t0 T0, t1 T1, t2 T2) Option[R] {
		return optionize(func() (R, bool) {
			return f(t0, t1, t2)
		})
	}
}

// Unoptionize3 converts a function with 3 parameters returning a tuple of a return value R and a boolean into a function with 3 parameters returning an Option[R]
func Unoptionize3[F ~func(T0, T1, T2) Option[R], T0, T1, T2, R any](f F) func(T0, T1, T2) (R, bool) {
	return func(t0 T0, t1 T1, t2 T2) (R, bool) {
		return Unwrap(f(t0, t1, t2))
	}
}

// Optionize4 converts a function with 4 parameters returning a tuple of a return value R and a boolean into a function with 4 parameters returning an Option[R]
func Optionize4[F ~func(T0, T1, T2, T3) (R, bool), T0, T1, T2, T3, R any](f F) func(T0, T1, T2, T3) Option[R] {
	return func(t0 T0, t1 T1, t2 T2, t3 T3) Option[R] {
		return optionize(func() (R, bool) {
			return f(t0, t1, t2, t3)
		})
	}
}

// Unoptionize4 converts a function with 4 parameters returning a tuple of a return value R and a boolean into a function with 4 parameters returning an Option[R]
func Unoptionize4[F ~func(T0, T1, T2, T3) Option[R], T0, T1, T2, T3, R any](f F) func(T0, T1, T2, T3) (R, bool) {
	return func(t0 T0, t1 T1, t2 T2, t3 T3) (R, bool) {
		return Unwrap(f(t0, t1, t2, t3))
	}
}

// Optionize5 converts a function with 5 parameters returning a tuple of a return value R and a boolean into a function with 5 parameters returning an Option[R]
func Optionize5[F ~func(T0, T1, T2, T3, T4) (R, bool), T0, T1, T2, T3, T4, R any](f F) func(T0, T1, T2, T3, T4) Option[R] {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4) Option[R] {
		return optionize(func() (R, bool) {
			return f(t0, t1, t2, t3, t4)
		})
	}
}

// Unoptionize5 converts a function with 5 parameters returning a tuple of a return value R and a boolean into a function with 5 parameters returning an Option[R]
func Unoptionize5[F ~func(T0, T1, T2, T3, T4) Option[R], T0, T1, T2, T3, T4, R any](f F) func(T0, T1, T2, T3, T4) (R, bool) {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4) (R, bool) {
		return Unwrap(f(t0, t1, t2, t3, t4))
	}
}

// Optionize6 converts a function with 6 parameters returning a tuple of a return value R and a boolean into a function with 6 parameters returning an Option[R]
func Optionize6[F ~func(T0, T1, T2, T3, T4, T5) (R, bool), T0, T1, T2, T3, T4, T5, R any](f F) func(T0, T1, T2, T3, T4, T5) Option[R] {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5) Option[R] {
		return optionize(func() (R, bool) {
			return f(t0, t1, t2, t3, t4, t5)
		})
	}
}

// Unoptionize6 converts a function with 6 parameters returning a tuple of a return value R and a boolean into a function with 6 parameters returning an Option[R]
func Unoptionize6[F ~func(T0, T1, T2, T3, T4, T5) Option[R], T0, T1, T2, T3, T4, T5, R any](f F) func(T0, T1, T2, T3, T4, T5) (R, bool) {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5) (R, bool) {
		return Unwrap(f(t0, t1, t2, t3, t4, t5))
	}
}

// Optionize7 converts a function with 7 parameters returning a tuple of a return value R and a boolean into a function with 7 parameters returning an Option[R]
func Optionize7[F ~func(T0, T1, T2, T3, T4, T5, T6) (R, bool), T0, T1, T2, T3, T4, T5, T6, R any](f F) func(T0, T1, T2, T3, T4, T5, T6) Option[R] {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6) Option[R] {
		return optionize(func() (R, bool) {
			return f(t0, t1, t2, t3, t4, t5, t6)
		})
	}
}

// Unoptionize7 converts a function with 7 parameters returning a tuple of a return value R and a boolean into a function with 7 parameters returning an Option[R]
func Unoptionize7[F ~func(T0, T1, T2, T3, T4, T5, T6) Option[R], T0, T1, T2, T3, T4, T5, T6, R any](f F) func(T0, T1, T2, T3, T4, T5, T6) (R, bool) {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6) (R, bool) {
		return Unwrap(f(t0, t1, t2, t3, t4, t5, t6))
	}
}

// Optionize8 converts a function with 8 parameters returning a tuple of a return value R and a boolean into a function with 8 parameters returning an Option[R]
func Optionize8[F ~func(T0, T1, T2, T3, T4, T5, T6, T7) (R, bool), T0, T1, T2, T3, T4, T5, T6, T7, R any](f F) func(T0, T1, T2, T3, T4, T5, T6, T7) Option[R] {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7) Option[R] {
		return optionize(func() (R, bool) {
			return f(t0, t1, t2, t3, t4, t5, t6, t7)
		})
	}
}

// Unoptionize8 converts a function with 8 parameters returning a tuple of a return value R and a boolean into a function with 8 parameters returning an Option[R]
func Unoptionize8[F ~func(T0, T1, T2, T3, T4, T5, T6, T7) Option[R], T0, T1, T2, T3, T4, T5, T6, T7, R any](f F) func(T0, T1, T2, T3, T4, T5, T6, T7) (R, bool) {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7) (R, bool) {
		return Unwrap(f(t0, t1, t2, t3, t4, t5, t6, t7))
	}
}

// Optionize9 converts a function with 9 parameters returning a tuple of a return value R and a boolean into a function with 9 parameters returning an Option[R]
func Optionize9[F ~func(T0, T1, T2, T3, T4, T5, T6, T7, T8) (R, bool), T0, T1, T2, T3, T4, T5, T6, T7, T8, R any](f F) func(T0, T1, T2, T3, T4, T5, T6, T7, T8) Option[R] {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8) Option[R] {
		return optionize(func() (R, bool) {
			return f(t0, t1, t2, t3, t4, t5, t6, t7, t8)
		})
	}
}

// Unoptionize9 converts a function with 9 parameters returning a tuple of a return value R and a boolean into a function with 9 parameters returning an Option[R]
func Unoptionize9[F ~func(T0, T1, T2, T3, T4, T5, T6, T7, T8) Option[R], T0, T1, T2, T3, T4, T5, T6, T7, T8, R any](f F) func(T0, T1, T2, T3, T4, T5, T6, T7, T8) (R, bool) {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8) (R, bool) {
		return Unwrap(f(t0, t1, t2, t3, t4, t5, t6, t7, t8))
	}
}

// Optionize10 converts a function with 10 parameters returning a tuple of a return value R and a boolean into a function with 10 parameters returning an Option[R]
func Optionize10[F ~func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) (R, bool), T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, R any](f F) func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) Option[R] {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9) Option[R] {
		return optionize(func() (R, bool) {
			return f(t0, t1, t2, t3, t4, t5, t6, t7, t8, t9)
		})
	}
}

// Unoptionize10 converts a function with 10 parameters returning a tuple of a return value R and a boolean into a function with 10 parameters returning an Option[R]
func Unoptionize10[F ~func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) Option[R], T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, R any](f F) func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) (R, bool) {
	return func(t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9) (R, bool) {
		return Unwrap(f(t0, t1, t2, t3, t4, t5, t6, t7, t8, t9))
	}
}