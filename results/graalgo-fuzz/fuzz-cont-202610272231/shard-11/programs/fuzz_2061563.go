package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := []uint32{98, 701, 42}
    v2 := make(map[int]bool)
    var v3 string = ("foo bar" + "hello")
    v4 := make(map[uint32]string)
    v5 := []string{v3}
    switch (v3 + v3) {
    case "\n\t":
        v3 = (v3 + v3)
    default:
        v2[127] = true
    }
    fmt.Println(v3)
    if true {
        switch 42 {
        case 128:
            v6 := v2[128]
            _ = v6
        case 0:
            for fi7 := 0; fi7 < 11; fi7++ {
                fmt.Println(v3)
                _ = fi7
            }
        case 42:
            v1 = append(v1, 277)
        default:
            for fi8 := 0; fi8 < 3; fi8++ {
                for fi9 := 0; fi9 < 20; fi9++ {
                    var v10 uint = 454
                    v11 := GenS1{F0: 2147483647, F1: -1}
                    if false {
                        v3 = v3
                        v4[65535] = v3
                    } else {
                        v10 -= v10
                        v10 = (^(v10))
                        var v12 int8 = 2
                        _ = v12
                    }
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                fmt.Println(v3, fi8)
                switch fi8 {
                case -245:
                    var v13 float32 = 3.14159
                    _ = v13
                case -9223372036854775808:
                    v3 = v3
                case 649:
                    for fi14 := 0; fi14 < 14; fi14++ {
                        v3 = "foo bar"
                        v15 := GenS2{F0: -2147483648}
                        v16 := GenS2{F0: 42}
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                default:
                    v17 := GenS2{F0: 437}
                    _ = v17
                }
                _ = fi8
            }
        }
        v3 = ("hello" + (v3 + v3))
    } else {
        v3 = v3
        v18 := GenS2{F0: -1}
        for fi19 := 0; fi19 < 17; fi19++ {
            if false {
                v20 := GenS2{F0: int32(fi19)}
                _ = v20
            } else {
                v21 := GenS1{F0: 42, F1: 1}
                var v22 string = "foo bar"
                v23 := GenS3{F0: false}
                _ = v21
                _ = v22
                _ = v23
            }
            v3 = v3
            v1 = append(v1, 0)
            _ = fi19
        }
        _ = v18
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        v25 := GenS3{F0: ((fi24 + -9223372036854775808) <= fi24)}
        var v26 int64 = 0
        _ = fi24
        _ = v25
        _ = v26
    }
    v3 = "foo bar"
    var v27 uint = 18446744073709551615
    var v28 int32 = 0
    if (0.0 >= -2.25) {
        var v29 uint8 = 2
        if true {
            v30 := GenS4{F0: uint16(v27)}
            _ = v30
        } else {
            var v31 uint = (v27 >> ((65535) & 63))
            switch ("a" + v3) {
            case "a":
                _ = len(v1)
            case "":
                v3 = v3
            default:
                fmt.Println(v29, v27)
            }
            v32 := v2[42]
            _ = v31
            _ = v32
        }
        v1 = append(v1, 0)
        _ = v29
    }
    v27 = v27
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
