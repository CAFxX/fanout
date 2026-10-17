package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 string; F1 bool }
func genFn0(p0 int16) uint8 {
    fmt.Println(p0)
    if (1e10 != ((3.14159 + 1.5) * 1e10)) {
        if true {
            var v6 uint = 4294967295
            v7 := GenS1{F0: -405, F1: 100}
            _ = v6
            _ = v7
        }
    }
    _ = p0
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 4294967295
    v3 := make(map[string]bool)
    var v4 string = "foo bar"
    var v5 uint32 = (v2 & (v2 * (^(v2))))
    v2 -= 185
    v5 -= v5
    fmt.Println(v2)
    switch 293 {
    case 0:
        if false {
            v4 = "\n\t"
            v4 = v4
            if true {
                if (("a" + v4) < v4) {
                    var v8 uint32 = (v5 >> ((4294967295) & 63))
                    if false {
                        v9 := GenS2{F0: 480, F1: 1}
                        var v10 bool = (!true)
                        _ = v10
                        _ = v9
                    } else {
                        v2 += 611
                        var v11 int8 = 9
                        _ = v11
                    }
                    v12 := []int{2147483647, 0, 2147483647}
                    _ = v12
                    _ = v8
                }
                fmt.Println(v4, v2, v4)
            }
        }
    case -1:
        v13 := GenS2{F0: 0, F1: 0}
        _ = v13
    case 9223372036854775807:
        if (93 == -9223372036854775808) {
            v4 = ""
        } else {
            if false {
                switch 2147483647 {
                case 32767:
                    if false {
                        v5 = (v2 | (v5 - 42))
                        v4 = (v4 + v4)
                    }
                default:
                    var v14 int32 = 382
                    _ = v14
                }
                v2 = (v5 % (((v5 >> ((1) & 63))) | 1))
            } else {
                if ((true || ((true || false) && false)) && ((3.14159 == float64((-2.25 / 1.0))) && false)) {
                    v15 := GenS2{F0: 396, F1: -9223372036854775808}
                    _ = v15
                }
                var v16 int = -2147483648
                v17 := GenS2{F0: 1, F1: -9223372036854775808}
                _ = v16
                _ = v17
            }
            v18 := []int{99, 128, -2147483648, 128}
            _ = v18
        }
    default:
        v19 := GenS3{F0: -32768, F1: 215}
        _ = v19
    }
    var v20 uint16 = 0
    var v21 uint16 = (v20 / ((v20) | 1))
    delete(v1, v2)
    v22 := []int{9223372036854775807}
    v22 = append(v22, -455)
    v21 -= v20
    for fi23 := 0; fi23 < 9; fi23++ {
        v21 = (v21 >> ((595) & 63))
        _ = fi23
    }
    var v24 uint32 = v2
    v24 += 1
    for fi25 := 0; fi25 < 5; fi25++ {
        v5 = 4294967295
        v26 := GenS4{F0: (v4 + "a"), F1: false}
        var v27 uint16 = (v20 % (((v21 + (v21 | 2))) | 1))
        _ = fi25
        _ = v26
        _ = v27
    }
    v24 -= 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
