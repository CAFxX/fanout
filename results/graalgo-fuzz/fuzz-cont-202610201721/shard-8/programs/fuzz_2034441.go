package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 uint32; F1 int16; F2 int64 }
type GenS7 struct { F0 int; F1 int; F2 int16 }
type GenS8 struct { F0 bool; F1 int; F2 string }
func genFn0() uint8 {
    var v6 int16 = 32767
    switch 0 {
    case -705:
        v6 -= (v6 / (((v6 % ((v6) | 1))) | 1))
    case 0:
        v7 := []int{32767, 128, 9223372036854775807}
        _ = v7
    case 127:
        var v8 uint8 = 128
        _ = v8
    default:
        v9 := GenS1{F0: -834}
        _ = v9
    }
    _ = v6
    return 1
}
func genFn1(p0 uint) string {
    v10 := []int{2147483647, 667, 2147483647}
    v11 := GenS2{F0: 1, F1: 100}
    v12 := GenS3{F0: 1, F1: ("\n\t" + "a")}
    var v13 uint16 = 30
    if ((v13 * v13) > (v13 & v13)) {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        p0 += (p0 % ((18446744073709551615) | 1))
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return "hello"
}

func main() {
    var v1 float64 = 0.1
    var v2 int32 = -133
    var v3 bool = (v2 > 1)
    v4 := []uint32{0, 1, 2}
    var v5 int16 = -1
    var v14 int32 = -2147483648
    if false {
        var v15 int8 = 92
        _ = v15
    } else {
        var v16 int32 = 1
        _ = v16
    }
    fmt.Println(v3, v5)
    var v17 int64 = -1
    var v18 uint64 = 4294967295
    for fi19 := 0; fi19 < 4; fi19++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi19
    }
    v4 = append(v4, 720)
    v17 -= (v17 % (((-(v17))) | 1))
    fmt.Println(v5)
    v20 := GenS4{F0: (v18 % ((42) | 1)), F1: ""}
    if v3 {
        switch 4294967295 {
        case 344:
            switch "" {
            case "foo bar":
                var v21 uint32 = 65535
                _ = v21
            }
        case 9223372036854775807:
            var v22 int32 = int32(v18)
            _ = v22
        }
    }
    if ((629 >= 42) && v3) {
        fmt.Println(v14)
        v4 = append(v4, 0)
        v23 := GenS5{F0: 32767, F1: (78 == 42)}
        _ = v23
    } else {
        v24 := GenS6{F0: 1, F1: (v5 * (v5 & 42)), F2: 42}
        var v25 uint16 = 255
        _ = v24
        _ = v25
    }
    fmt.Println(v17)
    if (v3 || (!v3)) {
        var v26 bool = v3
        var v27 float64 = (((v1 + (v1 / -2.25)) - v1) * v1)
        _ = v26
        _ = v27
    } else {
        if v3 {
            v18 -= 4294967295
        } else {
            if (255 <= 17) {
                var v28 float32 = (float32((3.14159 - float64(0.1))) / (1e10 * ((0.0 + float32(((v1 * 1.5) - float64(float32(v1))))) - (0.1 / (0.0 + ((1.5 + 1.5) * 0.1))))))
                fmt.Println(v1)
                _ = len(v4)
                _ = v28
            } else {
                v18 += v18
            }
        }
        switch 817 {
        case 0:
            v14 = -1
        case -968:
            switch 127 {
            case 32767:
                v14 = int32(v18)
            case 4294967295:
                v29 := GenS7{F0: 127, F1: -671, F2: 0}
                _ = v29
            case 128:
                if true {
                    for fi30 := 0; fi30 < 1; fi30++ {
                        var v31 uint8 = 2
                        _ = fi30
                        _ = v31
                    }
                    if (((!(32768 >= -380)) && true) && v3) {
                        v32 := GenS8{F0: v3, F1: 1, F2: "\n\t"}
                        _ = v32
                    }
                } else {
                    v33 := GenS7{F0: 32767, F1: 42, F2: 1}
                    if ((!v3) && (18446744073709551615 >= 18446744073709551615)) {
                        _ = len(v4)
                        var v34 int = -9223372036854775808
                        _ = v34
                    }
                    v1 -= v1
                    _ = v33
                }
            default:
                for fi35 := 0; fi35 < 10; fi35++ {
                    v4 = append(v4, 2)
                    var v36 int32 = v2
                    v1 += 0.1
                    _ = fi35
                    _ = v36
                }
            }
        case -574:
            v3 = v3
        default:
            v14 = v2
        }
    }
    fmt.Println(v2, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
