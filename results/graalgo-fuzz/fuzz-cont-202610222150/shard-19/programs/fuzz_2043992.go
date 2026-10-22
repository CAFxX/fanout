package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 string }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint64 = 4294967295
    v2 := []string{(("foo bar" + "\n\t") + "foo bar"), ("\n\t" + "foo bar"), "a", ("hello" + "")}
    v3 := make(map[int]bool)
    var v4 int = -2147483648
    var v5 bool = ((0.0 * 0.0) <= -2.25)
    delete(v3, (v4 + v4))
    v5 = (false || v5)
    var v6 uint = 555
    fmt.Println(v5, v4, v5)
    v5 = v5
    v7 := GenS1{F0: 538, F1: 251, F2: ("foo bar" + ("a" + "a"))}
    v6 = v6
    v1 -= 2
    v8 := GenS1{F0: 42, F1: 0, F2: (((("a" + ("a" + ("hello" + ("a" + "")))) + ("hello" + ("a" + "\n\t"))) + ("\n\t" + "hello")) + (("\n\t" + "hello") + "a"))}
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        fmt.Println(v1)
        _ = fi9
    }
    v4 = (v4 ^ 42)
    v1 -= v1
    if false {
        v6 = 4294967295
        v10 := GenS2{F0: uint16(v1)}
        _ = v10
    } else {
        v3[v4] = ((v6 | (v6 ^ (v6 * (v6 ^ 437)))) <= (v6 - v6))
    }
    if v5 {
        v11 := GenS1{F0: (^(v1)), F1: 148, F2: "hello"}
        _ = v11
    } else {
        fmt.Println(v4)
        var v12 bool = false
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
