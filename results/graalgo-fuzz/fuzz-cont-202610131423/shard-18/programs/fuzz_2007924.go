package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int }
func genFn0(p0 uint, p1 uint64) int64 {
    switch ("a" + ("a" + "foo bar")) {
    case "hello":
        for fi5 := 0; fi5 < 17; fi5++ {
            var v6 float32 = 1e10
            if (-35 < 127) {
                for fi7 := 0; fi7 < 12; fi7++ {
                    var v8 int64 = -1
                    _ = fi7
                    _ = v8
                }
                p0 += 18446744073709551615
                p1 = p1
            } else {
                for fi9 := 0; fi9 < 11; fi9++ {
                    var v10 int64 = 810
                    switch (fi9 >> ((408) & 63)) {
                    case -746:
                        v11 := []int{}
                        _ = v11
                    }
                    _ = fi9
                    _ = v10
                }
                v12 := []int{}
                for fi13 := 0; fi13 < 1; fi13++ {
                    if (false || (false && (!false))) {
                        if false {
                            var v14 int = int(p0)
                            var v15 uint8 = 2
                            _ = v14
                            _ = v15
                        }
                    }
                    if (false && false) {
                        p1 -= p1
                    }
                    _ = fi13
                }
                _ = v12
            }
            if true {
                var v16 uint64 = (^(p1))
                v16 = (v16 % ((0) | 1))
                v17 := []int{(fi5 << ((2) & 63)), -9223372036854775808, 9223372036854775807, fi5}
                _ = v16
                _ = v17
            }
            _ = fi5
            _ = v6
        }
    case "\n\t":
        var v18 bool = false
        _ = v18
    case "":
        if (1 >= 1) {
            v19 := GenS1{F0: p0}
            p0 -= p0
            _ = v19
        }
    default:
        v20 := []int{-959, 32767, -688}
        _ = v20
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 42
}
func genFn1() string {
    if true {
        var v21 int32 = -575
        _ = v21
    }
    var v22 uint8 = 255
    _ = v22
    return ""
}

func main() {
    var v1 int64 = 556
    v2 := make(map[uint32]bool)
    v3 := []string{(("foo bar" + (("hello" + "") + (("" + "a") + ""))) + ""), ("\n\t" + ""), ("hello" + ("hello" + "a"))}
    var v4 int32 = 42
    v23 := GenS2{F0: 2, F1: 325}
    v4 += v4
    v24 := v2[65535]
    v24 = false
    v4 += (v4 | 2147483647)
    var v25 int = 2147483647
    if v24 {
        if (1e10 != -2.25) {
            switch v25 {
            case -124:
                v26 := GenS1{F0: 0}
                _ = v26
            case 951:
                for fi27 := 0; fi27 < 14; fi27++ {
                    v25 += v25
                    var v28 int64 = v1
                    _ = fi27
                    _ = v28
                }
            default:
                for fi29 := 0; fi29 < 17; fi29++ {
                    for fi30 := 0; fi30 < 12; fi30++ {
                        v31 := GenS2{F0: 127, F1: (fi29 >> ((0) & 63))}
                        _ = fi30
                        _ = v31
                    }
                    _ = fi29
                }
            }
            var v32 float32 = 1.5
            _ = v32
        }
        if (42 > genFn0(1, 1)) {
            if len(v3) > 0 {
                v3[0] = (("a" + "foo bar") + "\n\t")
            }
            v2[2147483647] = (!true)
        }
        v24 = ((3.14159 + float64(0.1)) > ((3.14159 - (1.5 - float64(-2.25))) * float64(3.14159)))
    }
    fmt.Println(v25, v25, v4)
    _ = len(v3)
    v24 = true
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
