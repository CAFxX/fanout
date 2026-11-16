package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 string; F1 uint32; F2 bool }
type GenS4 struct { F0 uint32; F1 int; F2 int64 }
func genFn0(p0 int16) string {
    v6 := GenS1{F0: 255}
    var v7 uint16 = 0
    v8 := GenS1{F0: 194}
    fmt.Println(v7, p0)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return "foo bar"
}
func genFn1(p0 uint64, p1 int32, p2 int32) int16 {
    p0 -= 962
    p1 = -2147483648
    _ = p0
    _ = p1
    _ = p2
    return 476
}

func main() {
    var v1 int16 = -1
    v2 := make(map[uint32]bool)
    v3 := []uint32{}
    v4 := []string{"a"}
    var v5 uint16 = 2
    switch 32767 {
    case -2147483648:
        v9 := GenS2{F0: (v1 << ((42) & 63)), F1: 2147483647}
        _ = v9
    case -842:
        for fi10 := 0; fi10 < 4; fi10++ {
            v1 -= 42
            var v11 int16 = (^(v1))
            _ = fi10
            _ = v11
        }
    default:
        var v12 float64 = 0.0
        _ = v12
    }
    switch 4294967295 {
    case 2147483647:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        switch -763 {
        case 4294967295:
            v1 += (v1 | (v1 * v1))
        case 454:
            if (!((!(1e10 < (((1.5 + ((3.14159 + 3.14159) * 3.14159)) - (-2.25 / ((-2.25 * (0.1 / 0.1)) * (0.0 + 1e10)))) - 0.0))) || true)) {
                v13 := GenS3{F0: "\n\t", F1: 226, F2: true}
                _ = v13
            } else {
                v1 = v1
                var v14 string = "hello"
                _ = v14
            }
        case 42:
            if false {
                v15 := GenS4{F0: 0, F1: 400, F2: -344}
                v16 := GenS1{F0: v5}
                for fi17 := 0; fi17 < 20; fi17++ {
                    v1 -= (^(v1))
                    if (((801 == 0) && (true && true)) && true) {
                        _ = len(v2)
                        var v18 int32 = -2147483648
                        var v19 int64 = 9223372036854775807
                        _ = v18
                        _ = v19
                    } else {
                        v1 += (v1 ^ v1)
                        _ = len(v4)
                    }
                    _ = fi17
                }
                _ = v15
                _ = v16
            }
        default:
            v5 = v5
        }
    }
    v1 = (v1 + 1)
    v5 -= (v5 * 837)
    switch ("foo bar" + "") {
    case "foo bar":
        var v20 int16 = 42
        _ = v20
    }
    if (!true) {
        delete(v2, 0)
    } else {
        fmt.Println(v5, v1)
        var v21 uint32 = 135
        _ = v21
    }
    var v22 uint64 = 1
    var v23 string = ("foo bar" + "a")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
