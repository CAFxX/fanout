package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint8, p1 bool) int8 {
    v5 := GenS1{F0: 494}
    v6 := GenS1{F0: 65535}
    v7 := []int{42, 128, 127}
    if (p1 && (2 == p0)) {
        v7 = append(v7, 903)
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint32 = 1
    var v2 float32 = 3.14159
    var v3 int64 = 9223372036854775807
    var v4 int8 = 42
    switch 832 {
    case 1:
        var v8 uint8 = 128
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 uint16 = 32767
        v11 := []int{(fi9 + (^(fi9))), fi9, 2147483647}
        fmt.Println(v3, v3)
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 float64 = (0.0 * 3.14159)
    v1 = (v1 % ((v1) | 1))
    v2 -= (0.1 - v2)
    v13 := []int{127, int(v1)}
    for fi14 := 0; fi14 < 18; fi14++ {
        v15 := GenS1{F0: 1}
        v2 = v2
        if true {
            fmt.Println(v1, fi14, v3)
        }
        _ = fi14
        _ = v15
    }
    switch ("a" + (("" + ("hello" + "")) + ((("foo bar" + "\n\t") + ((("" + "hello") + ((("foo bar" + "") + "hello") + "a")) + ("\n\t" + ("\n\t" + "\n\t")))) + "a"))) {
    case "hello":
        v16 := GenS1{F0: 1}
        _ = v16
    default:
        if true {
            switch "\n\t" {
            case "a":
                if (255 == 662) {
                    var v17 uint32 = (v1 ^ v1)
                    _ = v17
                }
            default:
                switch "foo bar" {
                case "a":
                    if true {
                        v1 = v1
                    } else {
                        var v18 float32 = (v2 / v2)
                        _ = v18
                    }
                case "\n\t":
                    if true {
                        var v19 bool = false
                        _ = len(v13)
                        if false {
                            var v20 float64 = 3.14159
                            var v21 float32 = 1e10
                            v12 += 0.0
                            _ = v20
                            _ = v21
                        }
                        _ = v19
                    }
                default:
                    var v22 int16 = -121
                    _ = v22
                }
            }
        } else {
            for fi23 := 0; fi23 < 5; fi23++ {
                v2 = (v2 * (0.1 + v2))
                v24 := GenS1{F0: 575}
                _ = fi23
                _ = v24
            }
        }
    }
    fmt.Println(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
