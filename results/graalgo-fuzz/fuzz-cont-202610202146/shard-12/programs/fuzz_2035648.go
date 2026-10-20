package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int32 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint16; F1 string; F2 int64 }
type GenS4 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint32, p1 uint) int16 {
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint8 = 1
        v8 += (v8 / (((^(v8))) | 1))
        v9 := []int{fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 42, F1: -1, F2: 265}
    v11 := GenS2{F0: 42, F1: ("\n\t" > "")}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return int16(p1)
}
func genFn1(p0 int, p1 int32, p2 int8) int16 {
    var v12 float64 = (1e10 - (0.0 - 0.0))
    var v13 uint64 = 42
    for fi14 := 0; fi14 < 17; fi14++ {
        v13 += (v13 - (v13 + v13))
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return -349
}

func main() {
    var v1 int64 = -1
    v2 := make(map[string]int)
    var v3 uint64 = 1
    v4 := []string{"foo bar"}
    v5 := make(map[uint32]bool)
    var v6 int32 = -687
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 uint = 65535
        _ = fi15
        _ = v16
    }
    if (!false) {
        switch "" {
        case "\n\t":
            fmt.Println(v6)
        }
    }
    v4 = append(v4, "hello")
    v17 := GenS3{F0: 32767, F1: (("a" + "") + ""), F2: v1}
    var v18 uint64 = v3
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 uint32 = 912
        fmt.Println(v1, v6)
        if true {
            for fi21 := 0; fi21 < 11; fi21++ {
                if true {
                    fmt.Println(v6, v18)
                } else {
                    var v22 uint32 = 2147483648
                    switch 73 {
                    case -368:
                        var v23 int8 = 127
                        _ = v23
                    default:
                        if false {
                            v22 = (v20 / ((v20) | 1))
                            var v24 int64 = 206
                            _ = v24
                        }
                    }
                    _ = v22
                }
                v25 := GenS3{F0: 2, F1: "\n\t", F2: v1}
                v26 := v5[v20]
                _ = fi21
                _ = v25
                _ = v26
            }
            v27 := GenS1{F0: 495, F1: 0, F2: v6}
            v18 = v18
            _ = v27
        } else {
            _ = len(v2)
        }
        _ = fi19
        _ = v20
    }
    fmt.Println(v6, v6)
    v3 -= (v18 >> ((42) & 63))
    v3 = 524
    if false {
        fmt.Println(v18, v3, v18)
        if true {
            fmt.Println(v3, v3)
        }
        var v28 uint = 42
        _ = v28
    }
    v18 = 2
    for fi29 := 0; fi29 < 5; fi29++ {
        v1 = (v1 >> ((42) & 63))
        _ = fi29
    }
    switch -2147483648 {
    case 416:
        v30 := GenS4{F0: 1, F1: v3}
        _ = v30
    case -9223372036854775808:
        var v31 float32 = (float32(1e10) + 0.0)
        _ = v31
    case 2147483647:
        if len(v4) > 0 {
            v4[0] = ("" + "hello")
        }
    default:
        if false {
            var v32 float32 = 1.5
            _ = v32
        } else {
            var v33 string = ""
            _ = v33
        }
    }
    v5[631] = (179 > -472)
    v6 += v6
    for fi34 := 0; fi34 < 4; fi34++ {
        v3 = (v18 | v3)
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
