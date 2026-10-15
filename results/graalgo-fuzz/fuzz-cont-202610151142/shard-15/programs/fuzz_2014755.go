package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint8 }
func genFn0(p0 int8) uint64 {
    v5 := []int{128, 4294967295}
    var v6 uint16 = 2
    v7 := GenS1{F0: (^(v6)), F1: 42, F2: 42}
    v8 := GenS1{F0: (v6 % (((v6 / ((42) | 1))) | 1)), F1: 0, F2: 127}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 uint64 = 0
    var v2 float32 = 1e10
    v3 := make(map[uint32]string)
    var v4 int32 = -1
    var v9 int32 = 0
    var v10 uint64 = 939
    switch -311 {
    case 504:
        fmt.Println(v10, v4)
    }
    _ = len(v3)
    if (((false || (true || true)) || true) && (true && ((v10 ^ v10) != v10))) {
        switch -1 {
        case -2147483648:
            var v11 uint8 = 0
            _ = v11
        case 32767:
            if (!("foo bar" >= ("\n\t" + ("\n\t" + ("" + "\n\t"))))) {
                switch "foo bar" {
                case "a":
                    if (0 > 1) {
                        var v12 uint = 0
                        _ = v12
                    } else {
                        _ = len(v3)
                        v2 = v2
                    }
                case "foo bar":
                    for fi13 := 0; fi13 < 12; fi13++ {
                        v9 = 0
                        v9 = -189
                        _ = fi13
                    }
                default:
                    delete(v3, 42)
                }
            } else {
                v2 -= (0.1 + (v2 / (((0.0 * 1e10) + 0.1) / (((v2 / v2) / 0.1) / 3.14159))))
                var v14 float64 = -2.25
                v15 := []int{127, 9223372036854775807}
                _ = v14
                _ = v15
            }
        }
    }
    switch "" {
    case "\n\t":
        v4 += 762
    case "foo bar":
        v1 = genFn0(-128)
    }
    v4 += int32(v1)
    var v16 int8 = -103
    v1 = 950
    v17 := v3[209]
    fmt.Println(v17, v16, v2)
    var v18 int64 = -9223372036854775808
    v19 := []int{}
    for fi20 := 0; fi20 < 18; fi20++ {
        switch 9223372036854775807 {
        case 127:
            v19 = append(v19, fi20)
        case 754:
            if (true && (!(((fi20 ^ (fi20 + fi20)) == fi20) || (false && ((!true) && (!(4294967295 < v10))))))) {
                var v21 uint32 = 2147483648
                _ = v21
            } else {
                v2 += v2
            }
        case 9223372036854775807:
            for fi22 := 0; fi22 < 7; fi22++ {
                if ((!(!false)) && false) {
                    var v23 int8 = v16
                    v16 = v23
                    _ = v23
                } else {
                    v4 -= (v9 >> ((2) & 63))
                    v4 -= (v4 - (v4 + -1))
                    delete(v3, 0)
                }
                _ = fi22
            }
        }
        v24 := GenS1{F0: 42, F1: (fi20 / (((fi20 - -235)) | 1)), F2: 1}
        _ = fi20
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
