package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int32; F2 bool }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int; F1 int }
func genFn0(p0 int64, p1 string, p2 uint16) int64 {
    fmt.Println(p0, p0)
    if (p2 == 255) {
        var v7 string = ("a" + p1)
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return (^(p0))
}
func genFn1(p0 bool) uint16 {
    fmt.Println(p0)
    p0 = (!p0)
    if (!false) {
        fmt.Println(p0)
    } else {
        v8 := GenS1{F0: 18446744073709551615}
        _ = v8
    }
    _ = p0
    return 1
}

func main() {
    var v1 int32 = 0
    var v2 string = ("hello" + "\n\t")
    v3 := make(map[int]bool)
    v4 := []uint32{}
    v5 := make(map[uint32]int)
    v6 := make(map[int]uint32)
    v2 = "hello"
    if len(v4) > 0 {
        v4[0] = 0
    }
    v9 := GenS1{F0: 255}
    _ = len(v3)
    if (!false) {
        var v10 int = 1
        _ = v10
    } else {
        switch -666 {
        case 128:
            switch 0 {
            case -589:
                var v11 bool = ((!false) && true)
                _ = v11
            case 640:
                _ = len(v4)
            default:
                v2 = ""
            }
        }
        switch ("a" + v2) {
        case "foo bar":
            var v12 float64 = 3.14159
            _ = v12
        case "a":
            fmt.Println(v1)
        case "hello":
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 int32 = 507
                _ = fi13
                _ = v14
            }
        }
    }
    var v15 float32 = float32((-2.25 * 1e10))
    for fi16 := 0; fi16 < 13; fi16++ {
        for fi17 := 0; fi17 < 3; fi17++ {
            switch fi17 {
            case 908:
                var v18 uint64 = 4294967295
                _ = v18
            default:
                v1 += v1
            }
            _ = fi17
        }
        if false {
            v15 += v15
            if (0 == 2) {
                for fi19 := 0; fi19 < 6; fi19++ {
                    fmt.Println(v1, v2)
                    switch (fi19 - fi19) {
                    case 127:
                        var v20 int16 = 1000
                        _ = v20
                    case 554:
                        switch "" {
                        case "\n\t":
                            _ = len(v4)
                        case "a":
                            var v21 int64 = -1
                            _ = v21
                        }
                    case -9223372036854775808:
                        fmt.Println(fi16, v1)
                    default:
                        var v22 uint64 = 0
                        _ = v22
                    }
                    _ = fi19
                }
                if len(v4) > 0 {
                    v4[0] = 244
                }
            } else {
                for fi23 := 0; fi23 < 12; fi23++ {
                    switch ("foo bar" + v2) {
                    case "\n\t":
                        switch v2 {
                        case "\n\t":
                            v24 := GenS1{F0: 255}
                            _ = v24
                        default:
                            v15 -= (3.14159 - (((((v15 * (1e10 / v15)) * (v15 - v15)) - 0.1) / float32(1e10)) - 1e10))
                        }
                    case "hello":
                        var v25 float32 = ((3.14159 - v15) / (float32(((3.14159 / -2.25) * 3.14159)) * float32((-2.25 - float64(1.5)))))
                        _ = v25
                    default:
                        v26 := GenS2{F0: 0, F1: (-(v1)), F2: (true && true)}
                        _ = v26
                    }
                    v27 := GenS3{F0: 890, F1: 4294967295}
                    _ = fi23
                    _ = v27
                }
            }
        } else {
            fmt.Println(v15, v2)
            v28 := GenS4{F0: (^(fi16)), F1: 127}
            fmt.Println(fi16, fi16, v1)
            _ = v28
        }
        var v29 uint64 = 4294967295
        _ = fi16
        _ = v29
    }
    var v30 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
