package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int }
func genFn0(p0 string, p1 uint8) uint64 {
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    var v8 float64 = 0.1
    switch 42 {
    case 1:
        if ((42 <= 0) && false) {
            v9 := GenS1{F0: 2, F1: 0}
            _ = v9
        }
    case -813:
        var v10 string = ("a" + p0)
        _ = v10
    case 42:
        if (1 == 1) {
            p1 = p1
            fmt.Println(p0)
        }
    }
    _ = p0
    _ = p1
    _ = v8
    return 386
}
func genFn1(p0 bool, p1 int16, p2 string) uint16 {
    v11 := []int{-1, 42, 32767}
    p0 = p0
    v12 := GenS2{F0: 2147483648, F1: 742, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 65535
}

func main() {
    var v1 uint8 = 78
    v2 := []string{(("" + ("a" + "")) + "\n\t"), ("hello" + "a"), "\n\t"}
    var v3 uint64 = 4294967295
    var v4 uint8 = v1
    v5 := make(map[uint32]uint32)
    v5[65535] = 0
    var v13 uint8 = 128
    v14 := GenS1{F0: (^(v4)), F1: 1}
    var v15 float32 = float32(1.5)
    v16 := GenS2{F0: 404, F1: genFn0("\n\t", (v13 ^ (v1 | v1))), F2: 2147483647}
    v13 += v1
    switch 539 {
    case -249:
        var v17 float64 = 0.1
        _ = v17
    default:
        if (uint(v13) >= 184) {
            v18 := GenS1{F0: 201, F1: 2147483648}
            for fi19 := 0; fi19 < 19; fi19++ {
                for fi20 := 0; fi20 < 17; fi20++ {
                    fmt.Println(v1, v3)
                    switch "\n\t" {
                    case "a":
                        v21 := GenS2{F0: 750, F1: (v3 | genFn0(("foo bar" + ("a" + "hello")), v13)), F2: (fi19 - (fi20 * -1))}
                        _ = v21
                    }
                    _ = fi20
                }
                _ = fi19
            }
            v22 := GenS2{F0: 1, F1: v3, F2: 4294967295}
            _ = v18
            _ = v22
        }
    }
    if false {
        if (-2.25 < ((v15 + 1.5) * ((-2.25 * (1.5 + v15)) * ((v15 - v15) + ((-2.25 + v15) / v15))))) {
            v23 := GenS2{F0: 2147483648, F1: 18446744073709551615, F2: 593}
            var v24 uint = 42
            _ = v23
            _ = v24
        } else {
            for fi25 := 0; fi25 < 9; fi25++ {
                var v26 uint16 = 731
                var v27 uint8 = (v13 % ((v1) | 1))
                _ = fi25
                _ = v26
                _ = v27
            }
            var v28 int8 = 1
            _ = v28
        }
    }
    fmt.Println(v15, v15)
    v29 := GenS2{F0: 2, F1: (v3 << ((699) & 63)), F2: 9223372036854775807}
    var v30 bool = ("foo bar" < "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
