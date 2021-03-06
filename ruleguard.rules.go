// +build ignore

package gorules

import "github.com/quasilyte/go-ruleguard/dsl/fluent"

// This is a collection of rules for ruleguard: https://github.com/quasilyte/go-ruleguard

// Remove extra conversions: mdempsky/unconvert
func unconvert(m fluent.Matcher) {
	m.Match("int($x)").Where(m["x"].Type.Is("int") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")

	m.Match("float32($x)").Where(m["x"].Type.Is("float32") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")
	m.Match("float64($x)").Where(m["x"].Type.Is("float64") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")

	// m.Match("byte($x)").Where(m["x"].Type.Is("byte")).Report("unnecessary conversion").Suggest("$x")
	// m.Match("rune($x)").Where(m["x"].Type.Is("rune")).Report("unnecessary conversion").Suggest("$x")
	m.Match("bool($x)").Where(m["x"].Type.Is("bool") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")

	m.Match("int8($x)").Where(m["x"].Type.Is("int8") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")
	m.Match("int16($x)").Where(m["x"].Type.Is("int16") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")
	m.Match("int32($x)").Where(m["x"].Type.Is("int32") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")
	m.Match("int64($x)").Where(m["x"].Type.Is("int64") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")

	m.Match("uint8($x)").Where(m["x"].Type.Is("uint8") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")
	m.Match("uint16($x)").Where(m["x"].Type.Is("uint16") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")
	m.Match("uint32($x)").Where(m["x"].Type.Is("uint32") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")
	m.Match("uint64($x)").Where(m["x"].Type.Is("uint64") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")
}

// Don't use == or != with time.Time
// https://github.com/dominikh/go-tools/issues/47 : Wontfix
func timeeq(m fluent.Matcher) {
	m.Match("$t0 == $t1").Where(m["t0"].Type.Is("time.Time")).Report("using == with time.Time")
	m.Match("$t0 != $t1").Where(m["t0"].Type.Is("time.Time")).Report("using != with time.Time")
}

// Wrong err in error check
func wrongerr(m fluent.Matcher) {
	m.Match("if $*_, $err0 := $*_; $err1 != nil { $*_ }").
		Where(m["err0"].Text == "err" && m["err0"].Type.Is("error") && m["err1"].Text != "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("if $*_, $err0 := $*_; $err1 != nil { $*_ }").
		Where(m["err0"].Text != "err" && m["err0"].Type.Is("error") && m["err1"].Text == "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("if $*_, $err0 = $*_; $err1 != nil { $*_ }").
		Where(m["err0"].Text == "err" && m["err0"].Type.Is("error") && m["err1"].Text != "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("if $*_, $err0 = $*_; $err1 != nil { $*_ }").
		Where(m["err0"].Text != "err" && m["err0"].Type.Is("error") && m["err1"].Text == "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("if $*_, $err0 := $*_; $err1 == nil { $*_ }").
		Where(m["err0"].Text == "err" && m["err0"].Type.Is("error") && m["err1"].Text != "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("if $*_, $err0 := $*_; $err1 == nil { $*_ }").
		Where(m["err0"].Text != "err" && m["err0"].Type.Is("error") && m["err1"].Text == "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("if $*_, $err0 = $*_; $err1 == nil { $*_ }").
		Where(m["err0"].Text == "err" && m["err0"].Type.Is("error") && m["err1"].Text != "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("if $*_, $err0 = $*_; $err1 == nil { $*_ }").
		Where(m["err0"].Text != "err" && m["err0"].Type.Is("error") && m["err1"].Text == "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("$*_, $err0 := $*_; if $err1 != nil { $*_ }").
		Where(m["err0"].Text == "err" && m["err0"].Type.Is("error") && m["err1"].Text != "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("$*_, $err0 := $*_; if $err1 != nil { $*_ }").
		Where(m["err0"].Text != "err" && m["err0"].Type.Is("error") && m["err1"].Text == "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("$*_, $err0 := $*_; if $err1 == nil { $*_ }").
		Where(m["err0"].Text == "err" && m["err0"].Type.Is("error") && m["err1"].Text != "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("$*_, $err0 := $*_; if $err1 == nil { $*_ }").
		Where(m["err0"].Text != "err" && m["err0"].Type.Is("error") && m["err1"].Text == "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("$*_, $err0 = $*_; if $err1 != nil { $*_ }").
		Where(m["err0"].Text == "err" && m["err0"].Type.Is("error") && m["err1"].Text != "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("$*_, $err0 = $*_; if $err1 != nil { $*_ }").
		Where(m["err0"].Text != "err" && m["err0"].Type.Is("error") && m["err1"].Text == "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("$*_, $err0 = $*_; if $err1 == nil { $*_ }").
		Where(m["err0"].Text == "err" && m["err0"].Type.Is("error") && m["err1"].Text != "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("$*_, $err0 = $*_; if $err1 == nil { $*_ }").
		Where(m["err0"].Text != "err" && m["err0"].Type.Is("error") && m["err1"].Text == "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")
}

// err but no an error
// False positive: if err := recover(); err != nil { ... }
func errnoterror(m fluent.Matcher) {
	m.Match("if $*_, err := $x; $err != nil { $*_ }").
		Where(m["err"].Text == "err" && !m["err"].Type.Is("error")).
		Report("err variable not error type")

	m.Match("if $*_, err = $x; $err != nil { $*_ }").
		Where(m["err"].Text == "err" && !m["err"].Type.Is("error")).
		Report("err variable not error type")

	m.Match("if $err != nil { $*_ }").
		Where(m["err"].Text == "err" && !m["err"].Type.Is("error")).
		Report("err variable not error type")
}

// Identical if and else bodies
func ifbodythenbody(m fluent.Matcher) {
	m.Match("if $*_ { $body } else { $body }").
		Report("identical if and else bodies")
	m.Match("if $*_ { $body } else if $*_ { $body }").
		Report("identical if and else bodies")
}
