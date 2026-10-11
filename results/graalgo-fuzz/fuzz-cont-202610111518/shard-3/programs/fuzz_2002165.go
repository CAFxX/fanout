package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint16 }
func genFn0() int16 {
    var v5 float64 = 0.0
    for fi6 := 0; fi6 < 14; fi6++ {
        v5 = -2.25
        _ = fi6
    }
    v5 += (-2.25 - ((-2.25 - ((v5 / ((v5 + ((0.0 - 3.14159) / v5)) + float64(3.14159))) - -2.25)) / (((3.14159 + v5) - float64(float32(v5))) - -2.25)))
    fmt.Println(v5)
    _ = v5
    return -32768
}
func genFn1(p0 int8, p1 bool, p2 int64) uint64 {
    v7 := GenS1{F0: (true || true)}
    var v8 float64 = float64(1e10)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 483
}

func main() {
    var v1 uint = 165
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]uint32)
    switch ("foo bar" + "") {
    case "foo bar":
        v1 = v1
    case "a":
        var v9 uint = (v1 % ((65535) | 1))
        _ = v9
    case "hello":
        v1 = (v1 ^ (v1 / (((v1 & v1)) | 1)))
    default:
        if false {
            v10 := GenS2{F0: 2, F1: 2, F2: 1}
            _ = v10
        } else {
            v11 := GenS3{F0: (^(v1)), F1: uint16(v1)}
            _ = v11
        }
    }
    if false {
        if false {
            for fi12 := 0; fi12 < 20; fi12++ {
                if false {
                    v13 := GenS3{F0: v1, F1: 0}
                    for fi14 := 0; fi14 < 3; fi14++ {
                        fmt.Println(fi14, v1)
                        _ = fi14
                    }
                    _ = v13
                } else {
                    v1 -= (v1 >> ((v1) & 63))
                }
                v15 := v2[65535]
                _ = fi12
                _ = v15
            }
        }
    } else {
        var v16 int = 127
        v17 := GenS1{F0: (true || (4294967295 <= 4294967295))}
        if true {
            for fi18 := 0; fi18 < 16; fi18++ {
                v1 -= (v1 ^ v1)
                fmt.Println(fi18, fi18, v16)
                _ = fi18
            }
            v1 = v1
            var v19 int = (v16 | v16)
            _ = v19
        }
        _ = v16
        _ = v17
    }
    v1 = (v1 << ((v1) & 63))
    v1 += (v1 * v1)
    v20 := GenS1{F0: false}
    v1 -= v1
    var v21 uint = v1
    fmt.Println(v21, v21)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
