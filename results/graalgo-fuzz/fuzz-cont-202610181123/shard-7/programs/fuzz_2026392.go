package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint32 }
func genFn0(p0 int8, p1 string, p2 int64) bool {
    if (0 != 18446744073709551615) {
        v4 := GenS1{F0: 0, F1: 833, F2: 65535}
        if (!true) {
            if (!(false && true)) {
                v5 := []int{723, 42, 4294967295, -9223372036854775808}
                var v6 float64 = 3.14159
                _ = v5
                _ = v6
            }
            v7 := []int{32768}
            v8 := GenS1{F0: 2, F1: 345, F2: 2147483647}
            _ = v7
            _ = v8
        }
        p2 -= p2
        _ = v4
    }
    p1 = "hello"
    switch p1 {
    case "a":
        var v9 int = 128
        _ = v9
    case "\n\t":
        var v10 float64 = ((((0.1 / -2.25) + float64(-2.25)) - (1.5 / -2.25)) * (3.14159 - 1e10))
        _ = v10
    default:
        v11 := GenS1{F0: 42, F1: 0, F2: 4294967295}
        _ = v11
    }
    p0 += (-(p0))
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    var v1 bool = false
    var v2 float32 = 1.5
    var v3 uint32 = 617
    if true {
        if (4294967295 == 731) {
            v3 = v3
        } else {
            switch 32768 {
            case 510:
                switch 4294967295 {
                case -1:
                    var v12 int = 32768
                    _ = v12
                case 626:
                    for fi13 := 0; fi13 < 20; fi13++ {
                        var v14 float32 = float32((((3.14159 * 0.0) / (-2.25 - (1.5 / ((0.0 - -2.25) + float64(((-2.25 / float32(1e10)) - v2)))))) + (0.0 * -2.25)))
                        _ = fi13
                        _ = v14
                    }
                }
            default:
                var v15 float32 = -2.25
                _ = v15
            }
            var v16 int64 = -619
            _ = v16
        }
    }
    var v17 uint = 65535
    v17 -= v17
    v18 := []int{32768}
    for fi19 := 0; fi19 < 16; fi19++ {
        _ = len(v18)
        v17 = 0
        _ = fi19
    }
    fmt.Println(v2, v17)
    var v20 int64 = 904
    var v21 int = 42
    switch v21 {
    case 42:
        fmt.Println(v1, v17, v2)
    default:
        if v1 {
            _ = len(v18)
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 int8 = 127
                var v24 uint = 1
                _ = fi22
                _ = v23
                _ = v24
            }
            var v25 uint = (v17 | 666)
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        fmt.Println(v17, v3, v17)
        v20 = v20
        var v27 uint16 = 1
        _ = fi26
        _ = v27
    }
    v28 := GenS1{F0: 44, F1: v17, F2: 918}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v3
}
