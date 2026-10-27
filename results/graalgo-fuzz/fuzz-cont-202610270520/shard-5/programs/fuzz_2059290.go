package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 bool; F1 int64 }
func genFn0(p0 uint64, p1 bool, p2 uint32) uint64 {
    p2 = (^(p2))
    var v6 uint32 = (p2 | (p2 | (p2 & (p2 - p2))))
    var v7 int32 = -860
    v7 += -360
    v8 := []int{2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (p0 % (((^(p0))) | 1))
}
func genFn1(p0 int, p1 uint16, p2 int32) int16 {
    v9 := []int{}
    if ((656 <= 4294967295) && (0 == 1)) {
        if (!(true && true)) {
            switch p0 {
            case -924:
                if ((1000 > 578) && ((1e10 / 3.14159) <= (0.0 - 0.1))) {
                    p2 -= (p2 | 0)
                    if (!true) {
                        v10 := GenS1{F0: (^(p2))}
                        _ = v10
                    } else {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    }
                    if false {
                        if true {
                            p1 = 156
                            var v11 uint16 = (^(p1))
                            _ = v11
                        }
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    }
                }
            case -715:
                var v12 uint32 = 2
                _ = v12
            case 32768:
                v13 := GenS2{F0: true, F1: p0}
                _ = v13
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 32767
}

func main() {
    var v1 uint16 = 473
    v2 := make(map[uint32]uint32)
    v3 := []uint32{4294967295, 1, 2147483647, 42}
    var v4 uint8 = 255
    var v5 string = "a"
    fmt.Println(v4)
    fmt.Println(v1, v1)
    v2[758] = 4294967295
    if false {
        var v14 uint64 = 1
        _ = v14
    } else {
        switch v5 {
        case "":
            v15 := GenS1{F0: 1}
            _ = v15
        default:
            fmt.Println(v5, v4)
        }
        v16 := v2[2]
        if (v16 != (v16 % (((v16 ^ (v16 + (v16 * v16)))) | 1))) {
            var v17 uint32 = v16
            _ = v17
        } else {
            v1 = 0
            _ = len(v2)
        }
        _ = v16
    }
    var v18 int32 = 2147483647
    v3 = append(v3, 1)
    var v19 int8 = 127
    v20 := GenS3{F0: (true || (true && (false || (671 == 1)))), F1: -9223372036854775808}
    v5 = ("\n\t" + (((v5 + v5) + "\n\t") + v5))
    var v21 uint32 = 2147483648
    if false {
        v5 = v5
        v18 += (v18 * (v18 + v18))
        switch -67 {
        case 832:
            fmt.Println(v1, v4)
        case -1:
            var v22 uint32 = v21
            _ = v22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
