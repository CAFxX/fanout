package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int; F1 bool; F2 string }
type GenS4 struct { F0 int; F1 uint; F2 int16 }
func genFn0(p0 uint32) int8 {
    var v7 int = 32768
    v8 := GenS1{F0: 127, F1: 1, F2: 128}
    _ = p0
    _ = v7
    _ = v8
    return 83
}

func main() {
    v1 := []uint32{2, 2, 55, 0}
    var v2 int32 = 0
    var v3 uint8 = 255
    var v4 string = "foo bar"
    var v5 uint64 = 2
    var v6 uint8 = 1
    fmt.Println(v2, v6)
    if true {
        var v9 uint16 = 0
        if (255 < uint(v5)) {
            v10 := GenS2{F0: 830, F1: (v2 | (v2 - (v2 * (v2 | (^(v2))))))}
            _ = v10
        } else {
            v11 := GenS2{F0: 2147483647, F1: int32(v9)}
            _ = v11
        }
        _ = v9
    } else {
        fmt.Println(v6, v4, v4)
        var v12 uint = 2
        var v13 uint8 = (^(v3))
        _ = v12
        _ = v13
    }
    var v14 int32 = (-(v2))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v6, v14)
    v15 := GenS3{F0: 32767, F1: true, F2: ((v4 + (v4 + v4)) + v4)}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v16 := GenS4{F0: -718, F1: 2, F2: -139}
    if false {
        v5 -= (v5 >> ((65535) & 63))
    } else {
        if (42 <= 272) {
            if true {
                v17 := GenS4{F0: 795, F1: 42, F2: 42}
                _ = v17
            }
            if true {
                for fi18 := 0; fi18 < 7; fi18++ {
                    v1 = append(v1, 2147483647)
                    for fi19 := 0; fi19 < 12; fi19++ {
                        v1 = append(v1, 2147483647)
                        var v20 uint32 = 800
                        _ = fi19
                        _ = v20
                    }
                    _ = fi18
                }
            }
        }
        v1 = append(v1, 65535)
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi21
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        if (!true) {
            for fi23 := 0; fi23 < 6; fi23++ {
                v6 = 240
                v4 = ("foo bar" + v4)
                _ = fi23
            }
            fmt.Println(v3, v3, v4)
            var v24 bool = true
            _ = v24
        } else {
            v3 += 127
        }
        var v25 float32 = 0.1
        _ = fi22
        _ = v25
    }
    switch "a" {
    case "a":
        v1 = append(v1, 16)
    case "":
        v6 = v6
    default:
        var v26 bool = (3.14159 >= (0.0 / 1e10))
        _ = v26
    }
    if (v14 == v2) {
        if false {
            for fi27 := 0; fi27 < 5; fi27++ {
                for fi28 := 0; fi28 < 8; fi28++ {
                    var v29 string = v4
                    fmt.Println(v14, v5, fi27)
                    v5 = (v5 / (((v5 / ((v5) | 1))) | 1))
                    _ = fi28
                    _ = v29
                }
                v3 -= (v3 ^ v3)
                _ = fi27
            }
        }
    }
    if len(v1) > 0 {
        v1[0] = 410
    }
    v2 = (^(v14))
    v1 = append(v1, 2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
