package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int16 }
func genFn0(p0 uint32) bool {
    p0 += 768
    p0 -= p0
    p0 = 118
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 int16 = 1
        var v6 float64 = 0.0
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := []int{-779}
    _ = p0
    _ = v7
    return (false && ("foo bar" >= ""))
}
func genFn1(p0 int, p1 bool, p2 uint64) uint {
    p1 = (!(!p1))
    var v8 int16 = 0
    fmt.Println(p1, p0, v8)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}

func main() {
    var v1 int64 = 909
    var v2 float64 = 1.5
    var v3 uint = 470
    if ("hello" != ("a" + "foo bar")) {
        v2 += float64(0.0)
    }
    var v9 int32 = 42
    var v10 float64 = v2
    v10 -= ((v10 + 1.5) / (((v2 + v2) * (float64(0.0) / v10)) * 0.1))
    var v11 int32 = (v9 * (-(v9)))
    v12 := []int{-9223372036854775808, -186, 127}
    var v13 int = 42
    var v14 int8 = 75
    v15 := GenS1{F0: v3, F1: 255, F2: -32768}
    for fi16 := 0; fi16 < 10; fi16++ {
        switch ("hello" + ("a" + ("\n\t" + ("hello" + "foo bar")))) {
        case "foo bar":
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 uint = 1
                _ = fi17
                _ = v18
            }
        case "":
            v19 := GenS1{F0: (v3 << ((v3) & 63)), F1: 42, F2: -32768}
            _ = v19
        }
        var v20 uint = 941
        if ("foo bar" < ("hello" + (("foo bar" + "foo bar") + "hello"))) {
            v21 := GenS1{F0: (v3 ^ (v20 >> ((v3) & 63))), F1: 2, F2: 1000}
            var v22 float64 = float64(((1e10 / 1.5) + (1e10 + 1e10)))
            var v23 int64 = v1
            _ = v21
            _ = v22
            _ = v23
        }
        _ = fi16
        _ = v20
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 float64 = (v2 * v2)
        _ = fi24
        _ = v25
    }
    v12 = append(v12, (v13 + 414))
    v14 = (v14 + (v14 >> ((genFn1(v13, (!(-1 == v11)), 18446744073709551615)) & 63)))
    var v26 int64 = v1
    v26 = (v1 | (v26 / ((v26) | 1)))
    var v27 int16 = -494
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v9
}
