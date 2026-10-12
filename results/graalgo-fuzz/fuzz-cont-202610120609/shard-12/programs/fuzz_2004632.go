package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 bool; F2 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint }

func main() {
    var v1 int8 = 93
    v2 := make(map[uint32]uint32)
    v3 := []uint32{410, 1}
    v4 := GenS1{F0: 0}
    var v5 uint = 143
    var v6 uint16 = 2
    var v7 bool = false
    v1 -= (v1 + (v1 - v1))
    if v7 {
        v8 := GenS2{F0: int16(v5), F1: (-32768 != -385), F2: -1}
        _ = v8
    }
    v9 := GenS3{F0: 0}
    fmt.Println(v1, v7, v1)
    v5 = v5
    v10 := GenS1{F0: 65535}
    v1 += (v1 & v1)
    switch 910 {
    case 32767:
        v5 += 65535
    case -1:
        switch 2147483647 {
        case 636:
            var v11 int16 = -409
            _ = v11
        }
    default:
        v5 += 1
    }
    if false {
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 uint64 = 129
            _ = fi12
            _ = v13
        }
        if false {
            if (0.1 != -2.25) {
                v14 := GenS2{F0: 0, F1: false, F2: 1}
                if ("a" == (((((("foo bar" + "a") + "foo bar") + "hello") + "hello") + "hello") + ("foo bar" + "hello"))) {
                    v15 := GenS1{F0: uint32(v5)}
                    _ = v15
                }
                _ = v14
            }
            var v16 int8 = (v1 & -1)
            v1 = v16
            _ = v16
        } else {
            for fi17 := 0; fi17 < 11; fi17++ {
                v5 -= (^(v5))
                v5 += v5
                _ = fi17
            }
            if true {
                for fi18 := 0; fi18 < 3; fi18++ {
                    for fi19 := 0; fi19 < 18; fi19++ {
                        var v20 float32 = (-2.25 * (1.5 + (1.5 - (0.0 * (1.5 - 1.5)))))
                        v2[65535] = 337
                        v21 := GenS4{F0: 171, F1: (v5 / ((25) | 1)), F2: v5}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    _ = fi18
                }
                v6 -= 1
            } else {
                v5 += (v5 - v5)
                v22 := v2[2147483648]
                if ((int32(v22) > -44) && v7) {
                    v5 = v5
                }
                _ = v22
            }
            _ = len(v3)
        }
        switch 32767 {
        case 108:
            var v23 uint32 = 1
            _ = v23
        case 0:
            var v24 uint64 = 870
            _ = v24
        }
    }
    switch "a" {
    case "a":
        if (!((false || v7) || v7)) {
            var v25 int32 = 1
            _ = v25
        }
    default:
        fmt.Println(v6, v7, v7)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
