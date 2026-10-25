package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS3 struct { F0 int }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = 42
    v3 := make(map[int]bool)
    v4 := []uint32{0}
    v5 := GenS1{F0: -237, F1: 42}
    if (2147483648 >= 329) {
        if ((3.14159 * ((1.5 + (3.14159 - 1e10)) * ((-2.25 + (3.14159 / (3.14159 / 1.5))) + 0.0))) > 0.0) {
            v6 := GenS2{F0: 127, F1: 102, F2: 73}
            v2 -= v2
            for fi7 := 0; fi7 < 5; fi7++ {
                if (!true) {
                    _ = len(v3)
                }
                _ = fi7
            }
            _ = v6
        }
        v8 := GenS2{F0: 242, F1: 1, F2: 925}
        _ = v8
    } else {
        if (true || ((950 > 1) || (1 >= 0))) {
            if (v2 >= (v2 >> ((2) & 63))) {
                var v9 int8 = 100
                delete(v1, 528)
                for fi10 := 0; fi10 < 12; fi10++ {
                    v2 -= v2
                    _ = fi10
                }
                _ = v9
            } else {
                v2 = -1
                var v11 float32 = float32((0.1 + (3.14159 / 1.0)))
                _ = v11
            }
            v2 += 0
        } else {
            var v12 uint16 = 255
            v1[42] = 2147483648
            fmt.Println(v2, v12)
            _ = v12
        }
        v2 = v2
        if true {
            if ((1 >= -100) || false) {
                switch "" {
                case "\n\t":
                    var v13 uint16 = 645
                    _ = v13
                case "":
                    for fi14 := 0; fi14 < 11; fi14++ {
                        var v15 int32 = 1
                        delete(v1, 524)
                        fmt.Println(v2, fi14, v2)
                        _ = fi14
                        _ = v15
                    }
                }
            } else {
                v2 -= 42
                for fi16 := 0; fi16 < 19; fi16++ {
                    for fi17 := 0; fi17 < 15; fi17++ {
                        if len(v4) > 0 {
                            v4[0] = 2147483647
                        }
                        var v18 bool = (!true)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi17
                        _ = v18
                    }
                    _ = fi16
                }
                var v19 uint8 = 255
                _ = v19
            }
            v20 := GenS2{F0: 255, F1: 42, F2: v2}
            _ = v20
        } else {
            if len(v4) > 0 {
                v4[0] = 65535
            }
        }
    }
    fmt.Println(v2)
    fmt.Println(v2)
    v2 = (v2 & -1)
    _ = len(v4)
    v21 := v1[4294967295]
    switch 127 {
    case 727:
        for fi22 := 0; fi22 < 9; fi22++ {
            v23 := GenS3{F0: (fi22 / ((127) | 1))}
            v24 := GenS2{F0: 128, F1: 255, F2: 32767}
            _ = len(v4)
            _ = fi22
            _ = v23
            _ = v24
        }
    default:
        for fi25 := 0; fi25 < 15; fi25++ {
            v26 := GenS1{F0: 42, F1: -312}
            var v27 int32 = -2147483648
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    v21 = (v21 + v21)
    delete(v3, 127)
    var v28 uint16 = 32767
    if true {
        fmt.Println(v2, v2)
        switch "foo bar" {
        case "a":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            var v29 int16 = (v2 % ((337) | 1))
            _ = v29
        }
        var v30 uint = 4294967295
        _ = v30
    }
    fmt.Println(v2, v21)
    if true {
        fmt.Println(v28, v21, v2)
    }
    if (0.1 >= 1.5) {
        var v31 int16 = (v2 << ((0) & 63))
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
