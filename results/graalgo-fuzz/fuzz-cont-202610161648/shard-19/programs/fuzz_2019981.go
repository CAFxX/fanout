package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 int16 }
func genFn0(p0 int8, p1 bool, p2 uint8) int {
    fmt.Println(p1, p1, p0)
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 int16 = 32767
        var v8 float64 = 1e10
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 526
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 42
    v3 := []uint32{0, 4294967295, 65, 594}
    var v4 int8 = 127
    var v5 uint8 = 0
    v4 = (v4 * v4)
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 float64 = (0.0 + 3.14159)
        v12 := GenS1{F0: 2147483647, F1: 65, F2: int(v4)}
        _ = fi10
        _ = v11
        _ = v12
    }
    if ("foo bar" >= "a") {
        var v13 uint16 = 2
        _ = v13
    } else {
        v3 = append(v3, 2147483648)
        fmt.Println(v4, v5)
        if false {
            if len(v3) > 0 {
                v3[0] = 42
            }
            v4 += 100
        }
    }
    if true {
        var v14 uint16 = 548
        _ = v14
    } else {
        switch 128 {
        case 32767:
            var v15 int8 = v4
            _ = v15
        }
        _ = len(v1)
        var v16 uint8 = v5
        _ = v16
    }
    fmt.Println(v2)
    fmt.Println(v2)
    var v17 float32 = (0.0 * 1.5)
    if (4294967295 != uint(v5)) {
        v18 := GenS2{F0: (^(v5))}
        _ = v18
    } else {
        v17 -= float32(float64(v17))
        if true {
            if true {
                if (830 >= 1) {
                    if (42 != 2147483647) {
                        v17 -= 0.1
                        v5 += (^(v5))
                    } else {
                        var v19 bool = true
                        for fi20 := 0; fi20 < 9; fi20++ {
                            v5 = (^(v5))
                            v3 = append(v3, 42)
                            v4 -= v4
                            _ = fi20
                        }
                        _ = v19
                    }
                    var v21 int16 = int16(v5)
                    _ = v21
                } else {
                    var v22 int32 = 356
                    _ = v22
                }
                v5 = 128
            } else {
                switch "\n\t" {
                case "":
                    v5 = v5
                case "a":
                    var v23 int16 = (v2 & (-(v2)))
                    _ = v23
                default:
                    v2 = -661
                }
            }
            _ = len(v3)
        } else {
            var v24 uint32 = 442
            v25 := GenS3{F0: true}
            var v26 int16 = v2
            _ = v24
            _ = v25
            _ = v26
        }
        fmt.Println(v5, v5, v5)
    }
    v27 := GenS4{F0: 4294967295, F1: v2}
    var v28 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
