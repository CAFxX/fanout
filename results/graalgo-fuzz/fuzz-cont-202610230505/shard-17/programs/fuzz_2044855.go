package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 string; F1 uint16; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint64 }
func genFn0(p0 int64, p1 uint16) uint32 {
    var v4 int8 = 23
    v5 := []int{}
    for fi6 := 0; fi6 < 16; fi6++ {
        v5 = append(v5, fi6)
        fmt.Println(v4, p1, p1)
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 226
}
func genFn1() uint32 {
    var v7 uint8 = 0
    var v8 int32 = -682
    v9 := GenS1{F0: 266, F1: uint32(v7)}
    if true {
        var v10 float32 = 1e10
        for fi11 := 0; fi11 < 16; fi11++ {
            v7 -= 255
            _ = fi11
        }
        _ = v10
    }
    if (true && false) {
        switch 32767 {
        case 720:
            for fi12 := 0; fi12 < 7; fi12++ {
                switch "\n\t" {
                case "foo bar":
                    v13 := GenS1{F0: 640, F1: 754}
                    _ = v13
                }
                for fi14 := 0; fi14 < 17; fi14++ {
                    if (!true) {
                        v15 := GenS1{F0: 0, F1: 935}
                        _ = v15
                    }
                    if ((((-2.25 + (0.1 * -2.25)) - (((1e10 + 1.5) * -2.25) / float32(float64(((float32((3.14159 * 1.5)) / 3.14159) / 0.1))))) <= -2.25) && (true && false)) {
                        var v16 float64 = float64(1e10)
                        v8 -= v8
                        v17 := []int{(fi12 & (fi12 % (((fi14 + (fi14 / ((fi12) | 1)))) | 1))), fi14}
                        _ = v16
                        _ = v17
                    } else {
                        v7 = v7
                    }
                    _ = fi14
                }
                fmt.Println(v8, v8)
                _ = fi12
            }
        case 4294967295:
            v7 -= (v7 >> ((4294967295) & 63))
        }
        switch "\n\t" {
        case "":
            if false {
                var v18 int = 470
                _ = v18
            } else {
                for fi19 := 0; fi19 < 15; fi19++ {
                    fmt.Println(fi19)
                    v20 := []int{(fi19 >> ((uint(v7)) & 63)), fi19, (-(fi19))}
                    switch 2147483647 {
                    case -1:
                        fmt.Println(fi19, v8)
                    case -2147483648:
                        v8 = (v8 / ((v8) | 1))
                    case 32767:
                        var v21 uint16 = 255
                        _ = v21
                    }
                    _ = fi19
                    _ = v20
                }
                v22 := []int{-9223372036854775808, -165, 2147483647, int(v7)}
                _ = len(v22)
                _ = v22
            }
        default:
            v7 = (v7 ^ (v7 >> ((0) & 63)))
        }
    }
    _ = v7
    _ = v8
    _ = v9
    return 2147483648
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 127
    var v3 int32 = 42
    for fi23 := 0; fi23 < 19; fi23++ {
        v24 := GenS2{F0: ("\n\t" + "\n\t"), F1: 2, F2: 979}
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 bool = false
        _ = fi25
        _ = v26
    }
    var v27 int32 = v3
    v28 := []int{32768, v2, -2147483648}
    v29 := GenS3{F0: (!true), F1: 746, F2: 1}
    v30 := GenS2{F0: "hello", F1: 0, F2: 2}
    var v31 uint8 = 0
    var v32 int8 = -1
    v33 := GenS3{F0: false, F1: 42, F2: 0}
    v3 = (v27 & v3)
    var v34 float32 = float32(0.0)
    v34 = float32(1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}
