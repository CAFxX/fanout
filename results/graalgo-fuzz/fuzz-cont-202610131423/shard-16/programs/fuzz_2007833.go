package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint8; F1 int16; F2 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 bool; F2 int32 }

func main() {
    var v1 int16 = 1000
    var v2 uint32 = 874
    var v3 int32 = -248
    if true {
        var v4 int8 = -111
        v5 := GenS1{F0: v4, F1: 130, F2: 275}
        _ = v4
        _ = v5
    } else {
        var v6 int64 = 42
        _ = v6
    }
    v7 := GenS1{F0: 100, F1: 68, F2: 943}
    v1 = v1
    var v8 int32 = -205
    switch 437 {
    case 4294967295:
        switch 1 {
        case 32768:
            switch 267 {
            case -1:
                if (-1 != 2147483647) {
                    for fi9 := 0; fi9 < 6; fi9++ {
                        fmt.Println(v1, v1, v8)
                        for fi10 := 0; fi10 < 10; fi10++ {
                            var v11 float64 = 1.5
                            v12 := []int{}
                            var v13 uint32 = (v2 ^ 65535)
                            _ = fi10
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                        _ = fi9
                    }
                    v14 := []int{42, -305}
                    _ = v14
                } else {
                    if false {
                        v15 := GenS2{F0: 79, F1: 127}
                        v3 = (^(v8))
                        v3 = (v3 >> ((1) & 63))
                        _ = v15
                    }
                    var v16 int8 = -110
                    _ = v16
                }
            case 9223372036854775807:
                for fi17 := 0; fi17 < 17; fi17++ {
                    v8 = (v8 & v8)
                    v1 = (^(v1))
                    v18 := GenS1{F0: 127, F1: 0, F2: v1}
                    _ = fi17
                    _ = v18
                }
            default:
                if true {
                    for fi19 := 0; fi19 < 12; fi19++ {
                        v3 += (-(v8))
                        switch "foo bar" {
                        case "foo bar":
                            v8 = v8
                        }
                        v20 := GenS1{F0: 100, F1: 18446744073709551615, F2: v1}
                        _ = fi19
                        _ = v20
                    }
                } else {
                    v2 += v2
                    fmt.Println(v2, v1)
                    v21 := GenS3{F0: 1, F1: (v1 % (((v1 + 118)) | 1)), F2: 1}
                    _ = v21
                }
            }
        case 42:
            var v22 float32 = 0.1
            _ = v22
        case 127:
            for fi23 := 0; fi23 < 2; fi23++ {
                v3 = v3
                switch fi23 {
                case 571:
                    for fi24 := 0; fi24 < 18; fi24++ {
                        v25 := GenS3{F0: 2, F1: v1, F2: -1}
                        _ = fi24
                        _ = v25
                    }
                case 375:
                    var v26 int32 = v3
                    _ = v26
                default:
                    v8 = -1
                }
                _ = fi23
            }
        }
    case 2147483647:
        v27 := GenS3{F0: 0, F1: 0, F2: 7}
        _ = v27
    case 9223372036854775807:
        fmt.Println(v8, v1, v3)
    }
    v28 := []int{42, -176, 4294967295}
    for fi29 := 0; fi29 < 7; fi29++ {
        v8 = (-(v8))
        switch 32768 {
        case 766:
            v30 := GenS4{F0: (v8 & (v3 & (v3 ^ (v8 - v3))))}
            _ = v30
        }
        fmt.Println(v1, v8)
        _ = fi29
    }
    var v31 int16 = 0
    var v32 uint16 = 42
    if len(v28) > 0 {
        v28[0] = 42
    }
    fmt.Println(v32)
    for fi33 := 0; fi33 < 4; fi33++ {
        for fi34 := 0; fi34 < 2; fi34++ {
            if false {
                if ((4294967295 <= 942) && true) {
                    var v35 float64 = float64((3.14159 + 0.0))
                    _ = v35
                }
                fmt.Println(v1, v31)
            } else {
                v2 -= 0
                fmt.Println(fi33, v2)
            }
            _ = fi34
        }
        _ = fi33
    }
    v36 := GenS5{F0: 9223372036854775807, F1: (int(v2) > 9223372036854775807), F2: v3}
    v37 := GenS3{F0: 238, F1: (v1 << ((781) & 63)), F2: int64(v3)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v36
    _ = v37
    _ = v7
    _ = v8
}
