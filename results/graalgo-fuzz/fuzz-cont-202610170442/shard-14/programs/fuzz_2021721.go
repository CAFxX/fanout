package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := []string{(("\n\t" + ("a" + "\n\t")) + ("" + "foo bar")), ("hello" + ("a" + (("foo bar" + "foo bar") + ((((("\n\t" + (("foo bar" + "hello") + "a")) + ("foo bar" + "")) + "hello") + "") + (("" + "") + ""))))), "foo bar", "hello"}
    var v2 float64 = 1.5
    var v3 int8 = 111
    v4 := []string{("\n\t" + ("foo bar" + "\n\t")), "hello", "a", (("foo bar" + (("a" + "\n\t") + (("hello" + ("foo bar" + "")) + ((("hello" + "") + "foo bar") + ("" + "\n\t"))))) + "")}
    var v5 int32 = 42
    for fi6 := 0; fi6 < 12; fi6++ {
        v3 = v3
        var v7 bool = true
        _ = fi6
        _ = v7
    }
    v3 += (v3 ^ (v3 - v3))
    var v8 float64 = (v2 / (v2 * v2))
    v9 := GenS1{F0: 65535}
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 int = (fi10 - fi10)
        if len(v4) > 0 {
            v4[0] = "a"
        }
        var v12 int64 = 94
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint64 = 829
    v2 = v8
    v5 = 2147483647
    v5 -= (v5 << ((1) & 63))
    var v14 int64 = 299
    if (((v8 - float64(1.5)) >= (-2.25 / (((v2 / v8) + v2) * (float64(0.0) * float64((0.0 - (0.0 + 0.1))))))) && true) {
        var v15 uint = 1
        _ = v15
    }
    var v16 uint64 = v13
    v16 -= (v16 | 42)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
