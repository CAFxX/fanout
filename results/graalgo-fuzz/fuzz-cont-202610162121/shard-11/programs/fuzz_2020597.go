package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int32 }
func genFn0(p0 string) int8 {
    var v6 float32 = (1e10 + float32(0.0))
    var v7 bool = false
    if (false || (1 >= 127)) {
        v8 := GenS1{F0: 2147483647}
        for fi9 := 0; fi9 < 3; fi9++ {
            p0 = p0
            _ = fi9
        }
        _ = v8
    }
    _ = p0
    _ = v6
    _ = v7
    return 95
}
func genFn1(p0 int8) int16 {
    var v10 float32 = float32(0.1)
    if false {
        p0 += (p0 << ((0) & 63))
        p0 -= genFn0(("" + ""))
        for fi11 := 0; fi11 < 17; fi11++ {
            switch "\n\t" {
            case "a":
                if (!("\n\t" < "a")) {
                    var v12 float64 = 1.5
                    _ = v12
                } else {
                    v13 := []int{(fi11 ^ fi11), (fi11 << ((42) & 63))}
                    _ = v13
                }
            default:
                for fi14 := 0; fi14 < 15; fi14++ {
                    v15 := GenS1{F0: (^(fi14))}
                    if false {
                        v16 := GenS1{F0: fi14}
                        _ = v16
                    } else {
                        v17 := GenS1{F0: 42}
                        _ = v17
                    }
                    v18 := GenS1{F0: 32768}
                    _ = fi14
                    _ = v15
                    _ = v18
                }
            }
            v19 := GenS1{F0: fi11}
            p0 = (p0 % ((p0) | 1))
            _ = fi11
            _ = v19
        }
    }
    v20 := GenS2{F0: (2 > 811)}
    switch 236 {
    case 42:
        v21 := []int{464, 2147483647, 2147483647}
        _ = v21
    case 128:
        v22 := []int{-737, 2147483647}
        _ = v22
    case -325:
        var v23 float32 = 0.1
        _ = v23
    default:
        p0 = p0
    }
    p0 += (p0 + (p0 | p0))
    _ = p0
    _ = v10
    _ = v20
    return 1000
}

func main() {
    var v1 int8 = 1
    var v2 uint32 = 65535
    var v3 uint32 = (v2 * (^(v2)))
    var v4 uint32 = (v2 & v3)
    v5 := make(map[int]uint32)
    var v24 string = (("\n\t" + "a") + ("" + ""))
    var v25 float32 = (float32(1.5) - ((1.5 / -2.25) - 0.1))
    fmt.Println(v1, v2)
    var v26 string = ((v24 + "hello") + v24)
    var v27 uint32 = (v3 >> ((2) & 63))
    if false {
        switch "a" {
        case "":
            switch 816 {
            case -717:
                var v28 uint8 = 127
                _ = v28
            default:
                v25 += ((((v25 * v25) * -2.25) - v25) / v25)
            }
        case "a":
            if false {
                v29 := GenS3{F0: 65535, F1: 770}
                for fi30 := 0; fi30 < 18; fi30++ {
                    _ = len(v5)
                    _ = fi30
                }
                _ = v29
            } else {
                v25 = 1.5
            }
        default:
            v31 := []int{32768, 128, 2147483647}
            _ = v31
        }
    } else {
        for fi32 := 0; fi32 < 11; fi32++ {
            var v33 int64 = 0
            var v34 int32 = -1
            _ = fi32
            _ = v33
            _ = v34
        }
        for fi35 := 0; fi35 < 4; fi35++ {
            v2 = (v27 - (v2 ^ v4))
            _ = fi35
        }
    }
    for fi36 := 0; fi36 < 15; fi36++ {
        v37 := GenS3{F0: 1, F1: 968}
        _ = len(v5)
        var v38 float64 = 1e10
        _ = fi36
        _ = v37
        _ = v38
    }
    v39 := GenS2{F0: false}
    var v40 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
