package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int }
func genFn0(p0 uint32) uint {
    var v4 float32 = 0.0
    v5 := GenS1{F0: 415}
    fmt.Println(p0)
    _ = p0
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 uint8 = 197
    var v2 int = -801
    v3 := make(map[string]int)
    switch ("" + "foo bar") {
    case "a":
        v6 := GenS1{F0: -384}
        _ = v6
    case "":
        if true {
            if true {
                v3[("\n\t" + "hello")] = (v2 >> ((42) & 63))
                if (127 <= v2) {
                    v7 := []int{(v2 * v2)}
                    if len(v7) > 0 {
                        v7[0] = v2
                    }
                    _ = v7
                }
                v8 := []int{(v2 >> ((genFn0(2147483647)) & 63)), (v2 % ((v2) | 1)), v2, v2}
                _ = v8
            }
            for fi9 := 0; fi9 < 4; fi9++ {
                var v10 uint64 = 1
                _ = fi9
                _ = v10
            }
            v1 = 0
        }
    }
    v11 := GenS1{F0: -9223372036854775808}
    if true {
        var v12 int32 = 42
        _ = v12
    }
    v1 += (v1 & v1)
    v13 := GenS2{F0: 100}
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 uint8 = (v1 << ((2) & 63))
        if (-29 > 1) {
            v3[("" + "foo bar")] = (v2 & (-(fi14)))
        }
        v16 := GenS2{F0: 1}
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        v2 -= (fi17 * (fi17 | (^(fi17))))
        v18 := []int{}
        v2 += (fi17 >> ((264) & 63))
        _ = fi17
        _ = v18
    }
    var v19 uint16 = 217
    switch 32767 {
    case -2147483648:
        switch "a" {
        case "hello":
            if true {
                for fi20 := 0; fi20 < 3; fi20++ {
                    fmt.Println(v19, v19)
                    if (false && true) {
                        v1 = v1
                        v21 := GenS1{F0: 9223372036854775807}
                        _ = v21
                    } else {
                        var v22 float64 = ((((1.5 - 1.5) / (0.1 * 1e10)) - 1.5) / float64((1.5 - (1.5 / float32(((-2.25 - ((((0.1 * 1e10) - 3.14159) * 1.5) * 1.5)) + float64(1.5)))))))
                        delete(v3, "hello")
                        _ = v22
                    }
                    _ = fi20
                }
                v23 := GenS3{F0: "", F1: (v2 % ((v2) | 1))}
                _ = v23
            }
        case "a":
            fmt.Println(v1, v2, v19)
        case "\n\t":
            var v24 int32 = 1
            _ = v24
        }
    case -1:
        switch v2 {
        case 0:
            v1 += v1
        case 32768:
            v1 = v1
        default:
            v25 := GenS1{F0: int64(v1)}
            _ = v25
        }
    default:
        for fi26 := 0; fi26 < 14; fi26++ {
            for fi27 := 0; fi27 < 5; fi27++ {
                v2 += (fi26 >> ((4294967295) & 63))
                _ = fi27
            }
            v3[(("foo bar" + "") + "\n\t")] = fi26
            var v28 float64 = -2.25
            _ = fi26
            _ = v28
        }
    }
    var v29 uint8 = (^(v1))
    switch 128 {
    case 128:
        v30 := []int{v2}
        _ = v30
    case 1:
        var v31 uint = 23
        _ = v31
    default:
        v32 := []int{2147483647}
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v13
    _ = v19
    _ = v2
    _ = v29
    _ = v3
}
