package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint32 }
func genFn0() uint32 {
    var v6 int32 = 1
    v6 -= v6
    fmt.Println(v6)
    _ = v6
    return 65535
}

func main() {
    var v1 int16 = 0
    v2 := []string{"\n\t"}
    var v3 uint64 = 1
    v4 := make(map[int]uint32)
    var v5 bool = true
    switch 2147483647 {
    case 127:
        v7 := GenS1{F0: 112, F1: 128, F2: (v3 >> ((223) & 63))}
        _ = v7
    case 9223372036854775807:
        v1 -= v1
    default:
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := GenS1{F0: 1, F1: 220, F2: v3}
            if v5 {
                v5 = v5
                var v10 int = (fi8 >> ((18446744073709551615) & 63))
                if v5 {
                    v10 = (fi8 - 307)
                    for fi11 := 0; fi11 < 13; fi11++ {
                        var v12 float32 = (1e10 + -2.25)
                        var v13 int = fi11
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                } else {
                    v14 := GenS1{F0: -25, F1: 199, F2: 2}
                    var v15 int8 = -88
                    _ = v14
                    _ = v15
                }
                _ = v10
            }
            var v16 float32 = 1e10
            _ = fi8
            _ = v16
            _ = v9
        }
    }
    var v17 string = "foo bar"
    v18 := GenS2{F0: 255, F1: 835, F2: 18446744073709551615}
    if v5 {
        if (0.1 != 0.1) {
            v19 := GenS3{F0: false}
            switch v17 {
            case "":
                if ("hello" == ("hello" + "\n\t")) {
                    v5 = false
                    var v20 int32 = 286
                    _ = v20
                } else {
                    var v21 uint32 = 42
                    var v22 int8 = -1
                    fmt.Println(v21, v22)
                    _ = v21
                    _ = v22
                }
            case "foo bar":
                v23 := GenS4{F0: 950, F1: 32767, F2: 2147483647}
                _ = v23
            }
            v24 := GenS3{F0: v5}
            _ = v19
            _ = v24
        }
    }
    fmt.Println(v5, v17)
    var v25 float64 = 0.0
    var v26 uint64 = (v3 % ((0) | 1))
    if len(v2) > 0 {
        v2[0] = (v17 + (v17 + v17))
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        fmt.Println(v25, v5, v17)
        _ = fi27
    }
    v1 = v1
    switch 352 {
    case 9223372036854775807:
        var v28 uint8 = 42
        _ = v28
    case -2147483648:
        if (((1e10 * 0.0) - (float32(float64((3.14159 - (0.1 * 1e10)))) * -2.25)) < -2.25) {
            _ = len(v4)
            fmt.Println(v5, v17)
            var v29 uint = 0
            _ = v29
        } else {
            switch 32768 {
            case 655:
                var v30 int16 = v1
                _ = v30
            case -466:
                v31 := GenS2{F0: 65535, F1: 2, F2: 2}
                _ = v31
            case -1:
                switch "" {
                case "":
                    var v32 int32 = -1
                    _ = v32
                default:
                    v1 = v1
                }
            default:
                var v33 bool = (!(!(false && (!v5))))
                _ = v33
            }
            for fi34 := 0; fi34 < 8; fi34++ {
                v5 = v5
                _ = fi34
            }
            _ = len(v4)
        }
    default:
        v26 = (^(v3))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
