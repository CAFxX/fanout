package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS2 struct { F0 string; F1 bool; F2 int32 }
type GenS3 struct { F0 int64; F1 int64; F2 bool }
type GenS4 struct { F0 int8; F1 uint32 }

func main() {
    var v1 uint64 = 1
    v2 := []int{}
    v3 := make(map[uint32]bool)
    v4 := make(map[string]int)
    var v5 string = "hello"
    for fi6 := 0; fi6 < 14; fi6++ {
        v5 = v5
        v5 = (("\n\t" + "") + v5)
        _ = fi6
    }
    if false {
        fmt.Println(v5, v5)
        v7 := GenS1{F0: -755, F1: 15, F2: v1}
        _ = v7
    } else {
        fmt.Println(v5, v1)
    }
    fmt.Println(v1)
    var v8 uint16 = 110
    if false {
        _ = len(v3)
        fmt.Println(v1)
        switch 4294967295 {
        case 717:
            v2 = append(v2, 32768)
        case 0:
            delete(v3, 2)
        case 42:
            switch v5 {
            case "hello":
                v5 = v5
            case "foo bar":
                if true {
                    v8 -= (v8 + 0)
                    if len(v2) > 0 {
                        v2[0] = 32767
                    }
                }
            case "":
                v5 = "hello"
            default:
                switch v5 {
                case "hello":
                    if (0.0 != -2.25) {
                        fmt.Println(v5, v1, v8)
                    } else {
                        v4[v5] = 32767
                        var v9 uint32 = 4294967295
                        _ = v9
                    }
                case "a":
                    v8 -= (v8 + (^(v8)))
                case "foo bar":
                    if false {
                        var v10 uint64 = (v1 + (v1 & uint64(v8)))
                        v8 += (^(v8))
                        _ = v10
                    }
                default:
                    v11 := GenS1{F0: 1, F1: 42, F2: v1}
                    _ = v11
                }
            }
        default:
            if false {
                fmt.Println(v1, v8, v5)
            }
        }
    }
    var v12 uint32 = 828
    if true {
        fmt.Println(v12)
        if false {
            fmt.Println(v5, v1)
            var v13 float64 = 0.0
            _ = v13
        } else {
            if (-981 == -32768) {
                v14 := GenS2{F0: v5, F1: ((980 > 42) || false), F2: 1}
                for fi15 := 0; fi15 < 20; fi15++ {
                    v16 := GenS1{F0: -867, F1: 0, F2: (v1 * (^(v1)))}
                    for fi17 := 0; fi17 < 19; fi17++ {
                        v8 = 32767
                        var v18 uint64 = v1
                        _ = fi17
                        _ = v18
                    }
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
            v19 := GenS3{F0: 42, F1: int64(v12), F2: false}
            var v20 bool = (true && true)
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 17; fi21++ {
            var v22 int16 = -317
            _ = fi21
            _ = v22
        }
    }
    v23 := GenS1{F0: -746, F1: 128, F2: (v1 | 1)}
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 int8 = 0
        var v26 float32 = (-2.25 + 0.0)
        v27 := GenS3{F0: -9223372036854775808, F1: -9223372036854775808, F2: (976 <= 65535)}
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    switch -1 {
    case -9223372036854775808:
        _ = len(v2)
    default:
        switch 42 {
        case 1:
            v5 = (v5 + "foo bar")
        default:
            fmt.Println(v12, v1, v8)
        }
    }
    v28 := GenS1{F0: 498, F1: 212, F2: v1}
    for fi29 := 0; fi29 < 8; fi29++ {
        for fi30 := 0; fi30 < 6; fi30++ {
            v31 := GenS4{F0: -1, F1: 501}
            if true {
                v5 = "foo bar"
                var v32 uint32 = v12
                _ = v32
            } else {
                var v33 uint8 = 42
                fmt.Println(v12, v1)
                _ = v33
            }
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
