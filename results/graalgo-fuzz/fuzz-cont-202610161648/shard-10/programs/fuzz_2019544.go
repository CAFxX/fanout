package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint16, p1 uint, p2 string) uint32 {
    for fi4 := 0; fi4 < 6; fi4++ {
        fmt.Println(p0, p0)
        var v5 float32 = -2.25
        for fi6 := 0; fi6 < 2; fi6++ {
            switch fi6 {
            case 9223372036854775807:
                v7 := GenS1{F0: 250, F1: (p0 >> ((p1) & 63))}
                _ = v7
            case 2147483647:
                v5 += 1e10
            case 0:
                var v8 float32 = (1e10 / (((0.1 - (0.0 * 3.14159)) + v5) - float32((1e10 - 1e10))))
                _ = v8
            default:
                if (4294967295 > 1) {
                    if (false || (42 == int16(p1))) {
                        v9 := []int{9223372036854775807, 217, -301}
                        if (((((("\n\t" + (p2 + p2)) + p2) + (p2 + p2)) + (p2 + p2)) + p2) > (p2 + p2)) {
                            v10 := GenS1{F0: 18446744073709551615, F1: 476}
                            _ = v10
                        }
                        v9 = append(v9, (fi4 - -9223372036854775808))
                        _ = v9
                    }
                } else {
                    var v11 int16 = 0
                    v12 := []int{(fi6 % (((fi4 & (fi6 & 127))) | 1)), (^(fi4)), -307}
                    v12 = append(v12, (fi6 - (fi6 << ((p1) & 63))))
                    var v13 int16 = int16(fi6)
                    _ = v11
                    _ = v12
                    _ = v13
                }
            }
            fmt.Println(p1)
            var v14 float32 = (0.1 + (1e10 - (v5 / (v5 - 1.5))))
            _ = fi6
            _ = v14
        }
        _ = fi4
        _ = v5
    }
    var v15 uint8 = 57
    var v16 uint64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 42
}

func main() {
    var v1 int16 = 953
    var v2 int16 = (v1 % ((32767) | 1))
    v3 := make(map[uint32]bool)
    var v17 int8 = 127
    v1 += v1
    fmt.Println(v2, v1, v2)
    v18 := GenS1{F0: 4294967295, F1: 310}
    fmt.Println(v1, v17, v17)
    if ((839 >= 637) || false) {
        switch ("hello" + "\n\t") {
        case "a":
            var v19 uint32 = 2147483647
            _ = v19
        case "":
            switch -953 {
            case 127:
                v2 += 1000
            }
        case "hello":
            v20 := []int{-9223372036854775808, -681}
            _ = v20
        default:
            switch -2147483648 {
            case -888:
                var v21 float64 = 0.0
                _ = v21
            case 32767:
                var v22 int32 = 0
                _ = v22
            }
        }
        var v23 int16 = (v1 | (-(v2)))
        _ = v23
    }
    v2 -= (v1 - v2)
    var v24 bool = (!true)
    v17 = v17
    v25 := []int{2147483647, 127, -918}
    var v26 int8 = (v17 * v17)
    v3[genFn0(255, 42, "\n\t")] = v24
    _ = len(v3)
    _ = len(v3)
    var v27 uint64 = 1
    if v24 {
        v26 = (v26 / (((v17 << ((4294967295) & 63))) | 1))
        if (!(1 <= 2147483647)) {
            var v28 uint32 = 1
            _ = v28
        } else {
            for fi29 := 0; fi29 < 4; fi29++ {
                if (!((v17 << ((573) & 63)) <= 0)) {
                    v25 = append(v25, fi29)
                }
                _ = fi29
            }
            if v24 {
                if (-2.25 >= 1e10) {
                    v17 -= 88
                }
                var v30 int64 = -184
                var v31 uint64 = (v27 << ((255) & 63))
                _ = v30
                _ = v31
            }
            var v32 uint32 = 2147483647
            _ = v32
        }
        fmt.Println(v2)
    }
    var v33 string = (("a" + ("hello" + "\n\t")) + ("" + (("foo bar" + ("a" + "hello")) + "\n\t")))
    v33 = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v33
}
