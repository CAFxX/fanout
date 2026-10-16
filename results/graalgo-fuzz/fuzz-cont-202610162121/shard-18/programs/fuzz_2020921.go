package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int8; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int16, p1 uint, p2 int8) int64 {
    var v7 float64 = 1e10
    v8 := GenS1{F0: 42, F1: false}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int8, p1 int, p2 uint) uint64 {
    switch "hello" {
    case "hello":
        var v9 uint16 = 0
        _ = v9
    default:
        if true {
            var v10 uint64 = 378
            _ = v10
        }
    }
    p1 -= (p1 | -17)
    v11 := GenS1{F0: 281, F1: ((!(255 == p2)) && (true && false))}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 968
}

func main() {
    v1 := []uint32{}
    var v2 uint = 18446744073709551615
    var v3 int8 = 0
    var v4 uint = 1
    var v5 int16 = -32768
    var v6 uint8 = 128
    switch 1 {
    case 0:
        var v12 int8 = v3
        _ = v12
    }
    switch "hello" {
    case "hello":
        var v13 int64 = 184
        _ = v13
    case "":
        v14 := GenS1{F0: (v5 ^ 122), F1: (32767 > (v5 << (((v2 ^ 1)) & 63)))}
        _ = v14
    case "foo bar":
        var v15 int64 = -66
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 float32 = (1e10 * 1.5)
            if len(v1) > 0 {
                v1[0] = 0
            }
            fmt.Println(v6, v2)
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    if true {
        for fi19 := 0; fi19 < 9; fi19++ {
            _ = len(v1)
            _ = fi19
        }
    } else {
        v20 := GenS1{F0: v5, F1: false}
        _ = v20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        if len(v1) > 0 {
            v1[0] = uint32(v4)
        }
        fmt.Println(fi21, v6, fi21)
        var v22 uint16 = 1
        _ = fi21
        _ = v22
    }
    if true {
        v6 -= (v6 % (((v6 << ((v2) & 63))) | 1))
    }
    v4 = (^(v4))
    fmt.Println(v2, v3)
    v5 += v5
    if false {
        if (1 != 18446744073709551615) {
            if (!(2 != 0)) {
                if (!false) {
                    var v23 int = 1
                    var v24 uint16 = 390
                    switch v23 {
                    case 42:
                        v4 += 803
                    case 128:
                        if len(v1) > 0 {
                            v1[0] = 0
                        }
                    default:
                        fmt.Println(v6, v4)
                    }
                    _ = v23
                    _ = v24
                }
                for fi25 := 0; fi25 < 12; fi25++ {
                    var v26 uint8 = v6
                    _ = fi25
                    _ = v26
                }
            }
            var v27 int = 127
            _ = v27
        } else {
            if len(v1) > 0 {
                v1[0] = 0
            }
            if false {
                if ((^(v5)) <= -32768) {
                    v28 := GenS2{F0: (v3 * v3), F1: 42, F2: 32767}
                    v2 += 815
                    _ = len(v1)
                    _ = v28
                } else {
                    v1 = append(v1, 42)
                    var v29 uint64 = 863
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v29
                }
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        var v30 uint64 = 1
        _ = v30
    }
    var v31 int32 = -848
    switch 42 {
    case 931:
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
    case -9223372036854775808:
        for fi32 := 0; fi32 < 5; fi32++ {
            v2 -= v4
            switch fi32 {
            case 1:
                var v33 string = "foo bar"
                _ = v33
            case -576:
                for fi34 := 0; fi34 < 12; fi34++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(v4, v2)
                    _ = fi34
                }
            default:
                if (1.5 < ((0.0 - 0.1) / (0.0 + float64(float32(1.5))))) {
                    v3 = v3
                } else {
                    if len(v1) > 0 {
                        v1[0] = 2147483648
                    }
                    v4 = (v2 / (((v4 + 0)) | 1))
                    switch (-(fi32)) {
                    case -2147483648:
                        var v35 int16 = (v5 * -32768)
                        _ = v35
                    case 128:
                        v5 += v5
                    case 42:
                        if (true && ((^(fi32)) >= 0)) {
                            var v36 uint64 = 42
                            v31 = 974
                            v37 := GenS3{F0: 93}
                            _ = v36
                            _ = v37
                        }
                    default:
                        v38 := GenS3{F0: 2147483648}
                        _ = v38
                    }
                }
            }
            var v39 bool = true
            _ = fi32
            _ = v39
        }
    }
    v31 = v31
    if (!(v6 > v6)) {
        fmt.Println(v31, v3, v6)
    }
    v31 = v31
    var v40 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
