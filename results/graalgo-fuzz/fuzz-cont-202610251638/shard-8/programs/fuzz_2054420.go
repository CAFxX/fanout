package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint64 }
func genFn0(p0 int64, p1 bool) uint32 {
    if (0 >= 32767) {
        var v5 int16 = 796
        v6 := GenS1{F0: 4294967295}
        _ = v5
        _ = v6
    } else {
        p1 = p1
        if p1 {
            if p1 {
                v7 := GenS2{F0: 2, F1: 382}
                p1 = (255 == 18446744073709551615)
                if true {
                    p0 = (p0 ^ (p0 + p0))
                    for fi8 := 0; fi8 < 18; fi8++ {
                        p1 = ((!(1e10 < (float64(0.1) + ((1e10 * (1e10 * 0.0)) / 1.0)))) || true)
                        var v9 int = (^(fi8))
                        _ = fi8
                        _ = v9
                    }
                    var v10 bool = (20 >= 0)
                    _ = v10
                }
                _ = v7
            } else {
                p0 = (p0 / ((42) | 1))
            }
        } else {
            v11 := GenS2{F0: 2, F1: 2}
            p0 += p0
            _ = v11
        }
    }
    v12 := []int{2147483647}
    p0 -= p0
    for fi13 := 0; fi13 < 19; fi13++ {
        p0 = p0
        for fi14 := 0; fi14 < 1; fi14++ {
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            _ = fi14
        }
        if (42 == 255) {
            if len(v12) > 0 {
                v12[0] = (fi13 & (fi13 >> ((4294967295) & 63)))
            }
            fmt.Println(p1, p1, p0)
            switch ("" + (("" + "") + "foo bar")) {
            case "hello":
                var v15 int = (^(fi13))
                _ = v15
            case "a":
                var v16 uint8 = 2
                _ = v16
            }
        }
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = v12
    return 515
}
func genFn1() string {
    var v17 uint = 4294967295
    v17 -= 286
    v17 = v17
    _ = v17
    return "hello"
}

func main() {
    var v1 float32 = 3.14159
    var v2 int16 = -1
    v3 := []string{"hello", ("\n\t" + (("" + "") + "hello")), ("hello" + ("foo bar" + ("hello" + ("hello" + "a"))))}
    var v4 uint64 = 58
    v18 := GenS2{F0: 2, F1: (v4 ^ 0)}
    fmt.Println(v1, v2, v2)
    v2 = (v2 & v2)
    var v19 uint64 = v4
    v4 -= v19
    var v20 uint64 = v4
    v21 := GenS2{F0: 255, F1: (^(v4))}
    var v22 float32 = ((0.1 - 0.0) / v1)
    var v23 string = "\n\t"
    v2 += -1
    fmt.Println(v2, v23)
    v24 := GenS1{F0: v20}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
