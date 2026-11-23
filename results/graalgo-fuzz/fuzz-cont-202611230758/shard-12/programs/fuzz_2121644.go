package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
func genFn0() int {
    var v7 float64 = 3.14159
    for fi8 := 0; fi8 < 19; fi8++ {
        v7 += 1e10
        switch "foo bar" {
        case "":
            var v9 string = (((((("foo bar" + ((("a" + "") + ("hello" + ("\n\t" + (("\n\t" + "a") + "foo bar")))) + ("foo bar" + ("hello" + ("" + "hello"))))) + "foo bar") + "a") + ("" + ("foo bar" + ""))) + (("hello" + "foo bar") + ("\n\t" + ("" + (("" + "foo bar") + "\n\t"))))) + "\n\t")
            _ = v9
        case "\n\t":
            if (((42 <= 820) || (!false)) && false) {
                if (0 == 2) {
                    var v10 int64 = -857
                    v10 -= -1
                    _ = v10
                }
                fmt.Println(fi8, fi8)
                if false {
                    v7 += v7
                }
            } else {
                var v11 int8 = 127
                if false {
                    v12 := []int{}
                    var v13 uint16 = 0
                    _ = v12
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 10; fi14++ {
                        v15 := GenS1{F0: 0, F1: 65535}
                        v16 := []int{}
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                }
                _ = v11
            }
        case "foo bar":
            var v17 string = "hello"
            _ = v17
        default:
            var v18 uint64 = 4294967295
            _ = v18
        }
        _ = fi8
    }
    _ = v7
    return -599
}
func genFn1(p0 int64) uint {
    v19 := []int{}
    for fi20 := 0; fi20 < 16; fi20++ {
        var v21 bool = true
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = v19
    return 4294967295
}

func main() {
    var v1 int32 = -1
    v2 := []int{0, -2147483648}
    var v3 float32 = 1.5
    v4 := make(map[string]bool)
    var v5 int8 = 111
    v6 := make(map[string]string)
    v4["hello"] = (!true)
    v5 = (v5 + (v5 * (v5 % ((v5) | 1))))
    fmt.Println(v3)
    for fi22 := 0; fi22 < 16; fi22++ {
        v5 -= v5
        switch "\n\t" {
        case "":
            var v23 int32 = v1
            _ = v23
        case "a":
            fmt.Println(v5, v5)
        case "\n\t":
            var v24 uint8 = 42
            _ = v24
        }
        fmt.Println(v5, v1)
        _ = fi22
    }
    var v25 int64 = -1
    v26 := GenS1{F0: 2, F1: 65535}
    if (367 < 1) {
        if true {
            if (false || (1 <= 874)) {
                v4["a"] = (((-2.25 > (1.5 - (3.14159 - (0.1 * 1.5)))) || (4294967295 > 112)) || ((0.1 == ((-2.25 - -2.25) + 0.1)) && true))
                for fi27 := 0; fi27 < 5; fi27++ {
                    v28 := GenS1{F0: 110, F1: 2}
                    _ = fi27
                    _ = v28
                }
                var v29 uint16 = 543
                _ = v29
            }
        }
    }
    var v30 int = 127
    v25 = v25
    v25 -= 0
    v3 += (((float32((3.14159 - (0.0 * 1e10))) * v3) - (0.1 * 0.0)) + ((v3 / -2.25) / v3))
    v25 = v25
    var v31 int32 = v1
    v30 = (v30 - (v30 ^ -274))
    for fi32 := 0; fi32 < 3; fi32++ {
        _ = len(v2)
        _ = fi32
    }
    switch v30 {
    case 4294967295:
        v25 = v25
    case 32767:
        fmt.Println(v25, v31)
    case -708:
        if (!false) {
            fmt.Println(v3, v25, v1)
        }
    default:
        v33 := GenS1{F0: 42, F1: 255}
        _ = v33
    }
    var v34 uint64 = 744
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
