package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint32 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint8; F1 bool; F2 uint }
type GenS4 struct { F0 int32; F1 int64; F2 uint8 }
func genFn0() uint {
    v5 := GenS1{F0: 269, F1: 0, F2: 2147483647}
    v6 := GenS2{F0: 323, F1: 42}
    v7 := GenS1{F0: 656, F1: 128, F2: 4294967295}
    var v8 int = 32768
    var v9 uint64 = 530
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 718
}
func genFn1(p0 int32) uint16 {
    var v10 int = 42
    if true {
        p0 += 2147483647
    }
    v10 += (v10 ^ v10)
    _ = p0
    _ = v10
    return 805
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    var v3 uint32 = 419
    var v4 bool = false
    fmt.Println(v3, v3)
    v3 = v3
    switch 32768 {
    case 42:
        v3 = (v3 >> ((646) & 63))
    case 511:
        v4 = v4
    }
    var v11 float32 = 0.1
    v12 := []int{307, 2147483647, 128, 4294967295}
    fmt.Println(v4, v11)
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            fmt.Println(v3, fi13, fi13)
            v15 := v1[v3]
            _ = fi14
            _ = v15
        }
        switch "\n\t" {
        case "a":
            v3 = 42
        case "hello":
            v16 := GenS3{F0: 128, F1: ("foo bar" > ""), F2: 2}
            _ = v16
        default:
            switch fi13 {
            case -2147483648:
                v4 = true
            }
        }
        _ = fi13
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 4; fi18++ {
            if ((((0.1 * 1e10) - 1.5) == ((1e10 + (1e10 - 0.1)) * 0.0)) && v4) {
                v4 = v4
                var v19 uint = 4294967295
                if (v19 <= (v19 + v19)) {
                    v11 -= (v11 * v11)
                    v19 = v19
                    v20 := GenS4{F0: -651, F1: 42, F2: 128}
                    _ = v20
                } else {
                    delete(v1, 4294967295)
                    v11 = v11
                }
                _ = v19
            }
            _ = len(v12)
            var v21 uint = 867
            v22 := GenS3{F0: 128, F1: ((float32(1.5) / v11) == 3.14159), F2: (^(v21))}
            _ = fi18
            _ = v21
            _ = v22
        }
        v11 = 1e10
        if v4 {
            for fi23 := 0; fi23 < 5; fi23++ {
                v11 = (v11 + v11)
                if len(v12) > 0 {
                    v12[0] = (fi23 >> ((836) & 63))
                }
                _ = fi23
            }
            var v24 int8 = 42
            _ = v24
        }
        _ = fi17
    }
    switch 365 {
    case 4294967295:
        v4 = true
    default:
        switch 42 {
        case 454:
            fmt.Println(v3, v4)
        case 32767:
            var v25 int16 = -718
            _ = v25
        default:
            if len(v12) > 0 {
                v12[0] = 992
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
