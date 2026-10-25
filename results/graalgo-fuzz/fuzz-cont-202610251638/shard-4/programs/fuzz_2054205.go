package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 int32; F2 uint8 }
func genFn0(p0 uint64, p1 uint8, p2 int) int16 {
    if true {
        p1 = p1
        fmt.Println(p0, p1)
        fmt.Println(p1, p1, p0)
    }
    switch "foo bar" {
    case "\n\t":
        if (-5 <= -785) {
            fmt.Println(p2)
            var v5 string = ""
            _ = v5
        }
    default:
        v6 := []int{p2, p2, 32768}
        _ = v6
    }
    p1 -= (p1 >> ((599) & 63))
    fmt.Println(p1, p1, p1)
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 2}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 956
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 575
    var v3 uint = 757
    var v4 string = "a"
    if (0 == 864) {
        v2 = (v2 * (v2 - 4294967295))
        var v9 uint = (v3 & v3)
        switch "a" {
        case "foo bar":
            if ((v3 % ((65535) | 1)) != 42) {
                v2 += uint32(v9)
                v10 := []int{}
                _ = v10
            } else {
                switch 1 {
                case 0:
                    var v11 int64 = 9223372036854775807
                    _ = v11
                default:
                    v2 = v2
                }
            }
        case "hello":
            for fi12 := 0; fi12 < 19; fi12++ {
                v3 += (^(v9))
                _ = fi12
            }
        }
        _ = v9
    } else {
        v3 = (v3 >> (((v3 / (((^(v3))) | 1))) & 63))
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 int8 = 127
        _ = fi13
        _ = v14
    }
    var v15 uint = 279
    v2 -= 1
    if (!(0.0 > (1.5 + 1.5))) {
        fmt.Println(v15, v4, v4)
        var v16 uint32 = v2
        _ = v16
    }
    if false {
        switch "" {
        case "foo bar":
            if (((float32(0.1) + 0.0) + 0.0) >= 1.5) {
                switch -49 {
                case 0:
                    var v17 uint = (v15 * (v3 >> (((v15 & v15)) & 63)))
                    _ = v17
                default:
                    v2 = (v2 | (v2 / ((900) | 1)))
                }
                if false {
                    v18 := GenS1{F0: 42}
                    for fi19 := 0; fi19 < 19; fi19++ {
                        switch "foo bar" {
                        case "":
                            if false {
                                fmt.Println(v15)
                            }
                        case "hello":
                            v20 := []int{0, (fi19 | fi19)}
                            _ = v20
                        case "foo bar":
                            v21 := []int{(fi19 & (-(fi19))), fi19}
                            _ = v21
                        default:
                            v2 -= v2
                        }
                        _ = fi19
                    }
                    _ = v18
                }
                var v22 float32 = (1.5 * 0.1)
                _ = v22
            }
        case "\n\t":
            fmt.Println(v3)
        case "":
            var v23 int = 4294967295
            _ = v23
        }
        v2 -= 129
        v24 := GenS1{F0: v2}
        _ = v24
    }
    v25 := []int{42, 32767}
    v26 := GenS1{F0: (v2 | (v2 ^ (v2 | (v2 << ((uint(v2)) & 63)))))}
    var v27 float32 = 3.14159
    var v28 uint16 = 2
    v29 := GenS1{F0: uint32(v15)}
    for fi30 := 0; fi30 < 13; fi30++ {
        v15 -= v3
        _ = fi30
    }
    fmt.Println(v28, v15)
    v28 -= (v28 << ((v3) & 63))
    v31 := GenS2{F0: (v3 | 255)}
    var v32 uint8 = 189
    v33 := GenS3{F0: 4294967295, F1: 0, F2: 128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
