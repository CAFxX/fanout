package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int; F2 uint32 }
func genFn0(p0 int16) uint8 {
    v5 := GenS1{F0: 257, F1: 0, F2: -36}
    var v6 string = "\n\t"
    if true {
        v7 := []int{42, -985, 2147483647}
        if true {
            if (-8 < -74) {
                v6 = v6
                for fi8 := 0; fi8 < 9; fi8++ {
                    fmt.Println(v6, v6, v6)
                    _ = fi8
                }
                p0 -= p0
            }
            if true {
                switch 0 {
                case -677:
                    p0 -= (p0 % (((p0 ^ p0)) | 1))
                case -9223372036854775808:
                    for fi9 := 0; fi9 < 4; fi9++ {
                        if (32767 < 344) {
                            for fi10 := 0; fi10 < 20; fi10++ {
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                _ = fi10
                            }
                        } else {
                            fmt.Println(fi9, p0, fi9)
                        }
                        _ = fi9
                    }
                default:
                    p0 = p0
                }
                var v11 int16 = -32768
                _ = v11
            }
            for fi12 := 0; fi12 < 15; fi12++ {
                switch "" {
                case "":
                    p0 = (p0 >> ((65535) & 63))
                case "hello":
                    v6 = v6
                case "\n\t":
                    p0 = -114
                default:
                    p0 = p0
                }
                switch "hello" {
                case "":
                    p0 = (p0 | (-(p0)))
                default:
                    v6 = "\n\t"
                }
                var v13 int8 = 3
                _ = fi12
                _ = v13
            }
        }
        _ = v7
    }
    var v14 int = 515
    var v15 string = ((v6 + (v6 + v6)) + (v6 + (v6 + v6)))
    _ = p0
    _ = v14
    _ = v15
    _ = v5
    _ = v6
    return 165
}

func main() {
    var v1 int64 = 0
    var v2 int64 = (v1 - (v1 % (((v1 | 42)) | 1)))
    var v3 float64 = (0.1 * (0.1 - -2.25))
    var v4 int8 = 100
    v1 -= v2
    v16 := []int{42, -804, 946, 1}
    v17 := GenS1{F0: 32767, F1: 849, F2: v4}
    fmt.Println(v1, v1, v2)
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 uint16 = 42
        _ = fi18
        _ = v19
    }
    _ = len(v16)
    for fi20 := 0; fi20 < 13; fi20++ {
        v2 = (v2 * 1)
        _ = fi20
    }
    if true {
        v3 = float64((1.5 / 3.14159))
        if false {
            v4 -= (v4 | (v4 % (((-(v4))) | 1)))
            v21 := GenS1{F0: 255, F1: 443, F2: 79}
            _ = v21
        } else {
            v2 -= (v1 << ((0) & 63))
        }
        var v22 float64 = v3
        _ = v22
    }
    v4 = (v4 | -75)
    var v23 uint64 = 290
    v24 := GenS2{F0: 602}
    v25 := GenS3{F0: "a", F1: int(v1), F2: 4294967295}
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    v26 := GenS2{F0: 65535}
    var v27 float64 = -2.25
    v16 = append(v16, 869)
    if true {
        var v28 float32 = (3.14159 * -2.25)
        var v29 uint = 601
        _ = v28
        _ = v29
    }
    if false {
        var v30 float64 = 0.0
        _ = v30
    } else {
        if false {
            if false {
                var v31 int = 622
                var v32 float64 = 3.14159
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                _ = v31
                _ = v32
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
