package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint16; F2 bool }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 int16 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[uint32]string)
    var v3 uint32 = 42
    var v4 float32 = float32(0.1)
    v5 := make(map[string]string)
    fmt.Println(v3, v3)
    delete(v5, "foo bar")
    v4 += 1.5
    if (!true) {
        v1 -= v1
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 int64 = (v1 % ((v1) | 1))
        _ = fi6
        _ = v7
    }
    if (!(!true)) {
        if true {
            if false {
                if true {
                    v4 = float32(3.14159)
                } else {
                    var v8 int = 4294967295
                    _ = v8
                }
                v1 = 420
            }
            v9 := GenS1{F0: 2}
            var v10 int = 4294967295
            _ = v10
            _ = v9
        } else {
            v3 = 2
            v11 := GenS1{F0: 18446744073709551615}
            v4 -= v4
            _ = v11
        }
    }
    v3 = (^(v3))
    switch "a" {
    case "hello":
        v3 = (v3 ^ 2147483648)
    case "foo bar":
        switch 127 {
        case 4294967295:
            if (false && true) {
                for fi12 := 0; fi12 < 11; fi12++ {
                    for fi13 := 0; fi13 < 5; fi13++ {
                        var v14 bool = true
                        _ = fi13
                        _ = v14
                    }
                    v15 := GenS2{F0: 0, F1: 255, F2: (!true)}
                    _ = fi12
                    _ = v15
                }
                v4 -= (float32(float64(((v4 + v4) / v4))) * (0.0 / 1.0))
            } else {
                switch 665 {
                case -784:
                    for fi16 := 0; fi16 < 15; fi16++ {
                        v17 := GenS1{F0: 18446744073709551615}
                        for fi18 := 0; fi18 < 18; fi18++ {
                            v3 += v3
                            v19 := GenS3{F0: (v3 + v3), F1: 0}
                            var v20 uint16 = 255
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        _ = fi16
                        _ = v17
                    }
                }
                v21 := GenS3{F0: (^(v3)), F1: 32767}
                var v22 int8 = 127
                _ = v21
                _ = v22
            }
        case -9223372036854775808:
            v23 := GenS4{F0: 1000, F1: true}
            _ = v23
        default:
            v24 := GenS5{F0: 173}
            _ = v24
        }
    case "":
        var v25 uint16 = 42
        _ = v25
    }
    var v26 bool = true
    for fi27 := 0; fi27 < 12; fi27++ {
        v2[(v3 * (^(v3)))] = ""
        var v28 int16 = -209
        fmt.Println(v28, v4, v4)
        _ = fi27
        _ = v28
    }
    v29 := []int{-926, 32767, 297}
    var v30 uint64 = 18446744073709551615
    var v31 uint16 = uint16(v30)
    v30 = (v30 >> ((0) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
