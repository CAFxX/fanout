package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }

func main() {
    var v1 float32 = -2.25
    var v2 int16 = 32767
    v3 := []string{("" + ("\n\t" + "foo bar")), (("\n\t" + "") + "\n\t")}
    var v4 int16 = (v2 + (v2 >> ((18446744073709551615) & 63)))
    v5 := GenS1{F0: "a"}
    var v6 int16 = v2
    var v7 bool = false
    var v8 bool = (561 > 255)
    fmt.Println(v6, v1, v8)
    v2 = (v6 & v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v7 = (!(("a" + (("hello" + ("" + (((((("\n\t" + ("foo bar" + "")) + "a") + "a") + "foo bar") + ("" + ("\n\t" + ((((("" + "a") + "a") + ("a" + ("\n\t" + ""))) + "foo bar") + ("foo bar" + ("foo bar" + ("foo bar" + "foo bar"))))))) + ((((("\n\t" + "") + ("" + (("a" + "") + "hello"))) + ("" + "hello")) + "a") + "")))) + ("" + "\n\t"))) == ((((("foo bar" + (("" + (("" + ("foo bar" + "hello")) + (("a" + (("" + (("hello" + ("" + "hello")) + "")) + ("" + "a"))) + "hello"))) + "")) + "a") + "hello") + ("a" + (("hello" + ("\n\t" + ("a" + (("a" + "") + "foo bar")))) + "a"))) + "a")))
    v9 := GenS2{F0: 2147483647}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
