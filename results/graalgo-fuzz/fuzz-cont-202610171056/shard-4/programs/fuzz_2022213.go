package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int64; F2 int32 }
type GenS4 struct { F0 uint; F1 int }
func genFn0(p0 int16, p1 int64, p2 bool) uint {
    var v7 uint8 = 42
    if (((!(p0 == p0)) && p2) && p2) {
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := GenS1{F0: -1, F1: 0}
            _ = fi8
            _ = v9
        }
        v7 += v7
        for fi10 := 0; fi10 < 6; fi10++ {
            if (v7 >= 61) {
                p0 = (p0 >> ((4294967295) & 63))
                var v11 float32 = (float32((1e10 * -2.25)) / 1e10)
                _ = v11
            }
            var v12 int32 = 42
            p2 = p2
            _ = fi10
            _ = v12
        }
    } else {
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 int64 = p1
            _ = fi13
            _ = v14
        }
    }
    p2 = ((float64(0.1) < 1e10) && p2)
    var v15 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v7
    return 2
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 213
    var v3 float64 = 3.14159
    v4 := []uint32{65535, 42}
    v5 := make(map[int]int)
    v6 := make(map[uint32]string)
    v16 := GenS1{F0: 1, F1: -1}
    for fi17 := 0; fi17 < 11; fi17++ {
        fmt.Println(v2, fi17)
        _ = fi17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 int32 = -641
        v19 += -2147483648
        _ = fi18
        _ = v19
    }
    fmt.Println(v2, v3)
    v2 = v2
    v3 -= 0.1
    if false {
        v3 = v3
    } else {
        v3 = ((v3 * float64(float32(((-2.25 / (v3 + float64(((((0.0 * (((3.14159 + -2.25) * 0.0) - 1e10)) / ((1e10 / (-2.25 * 3.14159)) - -2.25)) * -2.25) + float32(v3))))) * v3)))) / v3)
    }
    v20 := GenS1{F0: 1, F1: 42}
    v4 = append(v4, 978)
    var v21 float32 = 1e10
    if true {
        var v22 int32 = 1
        _ = len(v6)
        v23 := GenS1{F0: -1, F1: (v22 / (((v22 % ((v22) | 1))) | 1))}
        _ = v22
        _ = v23
    }
    v3 = 1.5
    var v24 bool = false
    v25 := GenS2{F0: 476}
    v5[1] = 287
    for fi26 := 0; fi26 < 9; fi26++ {
        for fi27 := 0; fi27 < 4; fi27++ {
            fmt.Println(v3, v2)
            fmt.Println(v2)
            _ = fi27
        }
        for fi28 := 0; fi28 < 17; fi28++ {
            var v29 int = fi26
            v30 := GenS1{F0: 0, F1: 466}
            _ = fi28
            _ = v29
            _ = v30
        }
        for fi31 := 0; fi31 < 14; fi31++ {
            if (false || v24) {
                if true {
                    var v32 float64 = (v3 * (v3 / ((((-2.25 / 1e10) / (v3 / (0.1 * v3))) * (3.14159 / v3)) + (v3 + v3))))
                    _ = v32
                }
                v33 := GenS3{F0: v2, F1: -772, F2: -1}
                var v34 int16 = 42
                v2 = 61
                _ = v33
                _ = v34
            } else {
                v35 := GenS4{F0: genFn0(1, 0, ((v2 / ((v2) | 1)) >= (^(v2)))), F1: fi31}
                v36 := GenS4{F0: 1, F1: fi31}
                var v37 float64 = float64(0.1)
                switch "" {
                case "\n\t":
                    var v38 float32 = float32(((v37 + 3.14159) - (float64(0.1) / v37)))
                    _ = v38
                case "a":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    if (true && (4294967295 < 0)) {
                        _ = len(v6)
                        v39 := GenS2{F0: 65535}
                        var v40 int8 = -1
                        _ = v39
                        _ = v40
                    }
                }
                _ = v35
                _ = v36
                _ = v37
            }
            if len(v4) > 0 {
                v4[0] = 42
            }
            v3 = 0.0
            _ = fi31
        }
        _ = fi26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
