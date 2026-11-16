package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int64 }
type GenS2 struct { F0 uint; F1 int64; F2 int8 }
func genFn0(p0 uint, p1 int) int16 {
    var v6 uint32 = 65535
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{p1, (p1 / (((p1 >> ((p0) & 63))) | 1)), p1}
        _ = fi7
        _ = v8
    }
    p0 = (p0 / ((p0) | 1))
    v9 := []int{}
    var v10 int16 = -676
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v9
    return (v10 & (v10 ^ v10))
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[uint32]bool)
    var v3 bool = true
    v4 := make(map[string]bool)
    var v5 uint64 = 0
    switch ("hello" + "foo bar") {
    case "hello":
        switch 541 {
        case -2147483648:
            var v11 int = 2147483647
            _ = v11
        case 9223372036854775807:
            if false {
                for fi12 := 0; fi12 < 9; fi12++ {
                    if v3 {
                        for fi13 := 0; fi13 < 17; fi13++ {
                            var v14 int8 = 62
                            v2[4294967295] = ((18446744073709551615 != 255) || ((v1 / ((v1) | 1)) > (v1 - v1)))
                            _ = fi13
                            _ = v14
                        }
                        v15 := GenS1{F0: genFn0(795, (fi12 * fi12)), F1: fi12, F2: -9223372036854775808}
                        _ = v15
                    } else {
                        var v16 int8 = 0
                        _ = v16
                    }
                    v3 = v3
                    _ = fi12
                }
                v17 := GenS2{F0: 7, F1: 9223372036854775807, F2: -128}
                _ = v17
            }
        }
    default:
        if v3 {
            var v18 int32 = v1
            _ = v18
        } else {
            var v19 int16 = -32768
            v1 = v1
            v20 := v4[("hello" + "hello")]
            _ = v19
            _ = v20
        }
    }
    if v3 {
        if v3 {
            fmt.Println(v5, v5, v3)
            v3 = v3
            var v21 uint64 = (v5 + 1)
            _ = v21
        } else {
            v22 := GenS2{F0: 65535, F1: 42, F2: 1}
            _ = v22
        }
        var v23 uint32 = 4294967295
        _ = v23
    } else {
        for fi24 := 0; fi24 < 6; fi24++ {
            v4[((((("foo bar" + (("a" + "\n\t") + (("" + "\n\t") + ""))) + "\n\t") + "a") + "") + "hello")] = v3
            _ = fi24
        }
        var v25 bool = (!(v3 && (127 >= 31)))
        for fi26 := 0; fi26 < 2; fi26++ {
            v5 += (^(v5))
            var v27 int64 = 0
            v5 += (v5 * v5)
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    v1 = v1
    switch "" {
    case "hello":
        for fi28 := 0; fi28 < 20; fi28++ {
            var v29 int8 = 1
            var v30 uint = 255
            _ = fi28
            _ = v29
            _ = v30
        }
    default:
        switch -340 {
        case 32768:
            fmt.Println(v5)
        default:
            fmt.Println(v5, v1, v1)
        }
    }
    fmt.Println(v3, v1, v1)
    v31 := []int{127, 32767, 128}
    v1 -= (^(v1))
    if false {
        var v32 uint16 = 0
        _ = v32
    }
    var v33 bool = v3
    var v34 uint32 = 326
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
