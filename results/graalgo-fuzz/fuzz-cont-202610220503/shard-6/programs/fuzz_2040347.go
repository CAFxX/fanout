package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
func genFn0(p0 uint16, p1 uint32, p2 uint) bool {
    var v4 bool = (-1 != -9223372036854775808)
    p0 = p0
    switch 1 {
    case 1:
        var v5 float64 = -2.25
        _ = v5
    default:
        var v6 int8 = 17
        _ = v6
    }
    if false {
        var v7 int64 = -731
        var v8 float32 = 0.0
        _ = v7
        _ = v8
    } else {
        var v9 int8 = -126
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return ((^(p2)) == (p2 ^ p2))
}

func main() {
    var v1 uint64 = 42
    v2 := []string{}
    v3 := []string{"a"}
    switch 0 {
    case 4294967295:
        var v10 int16 = 32767
        _ = v10
    case 2147483647:
        if ((-802 == 0) || (!("foo bar" < ("" + "")))) {
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 float32 = 3.14159
                v12 -= 1.5
                v13 := GenS1{F0: 1000, F1: 18446744073709551615}
                _ = fi11
                _ = v12
                _ = v13
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi14 := 0; fi14 < 15; fi14++ {
                fmt.Println(v1, fi14)
                if len(v3) > 0 {
                    v3[0] = ""
                }
                if (255 == 255) {
                    if false {
                        var v15 bool = true
                        fmt.Println(fi14)
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 8; fi16++ {
                        var v17 uint32 = 682
                        _ = fi16
                        _ = v17
                    }
                } else {
                    var v18 uint16 = 255
                    _ = v18
                }
                _ = fi14
            }
        }
    default:
        var v19 float64 = 1.5
        _ = v19
    }
    var v20 int = -2147483648
    var v21 float32 = (1e10 * 3.14159)
    var v22 string = ("" + "\n\t")
    var v23 uint8 = 2
    fmt.Println(v22)
    v3 = append(v3, "a")
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
