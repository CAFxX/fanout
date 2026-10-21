package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint32; F1 uint }

func main() {
    var v1 string = ("" + "a")
    v2 := make(map[int]string)
    v3 := []int{692, 4294967295, 814, 717}
    var v4 uint16 = 2
    for fi5 := 0; fi5 < 17; fi5++ {
        delete(v2, fi5)
        v6 := GenS1{F0: 883, F1: ("a" + ("" + ((v1 + "hello") + "foo bar")))}
        switch "\n\t" {
        case "foo bar":
            v1 = v1
        }
        _ = fi5
        _ = v6
    }
    if (false || false) {
        var v7 int32 = 42
        v8 := GenS2{F0: 0, F1: v4}
        _ = v7
        _ = v8
    } else {
        v1 = v1
        v9 := GenS1{F0: 42, F1: ("a" + ("a" + ""))}
        _ = v9
    }
    fmt.Println(v1)
    v1 = v1
    var v10 int8 = -1
    v4 = (v4 >> ((2) & 63))
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 uint64 = 42
        v10 -= v10
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 float32 = (0.0 * 0.0)
        v10 += (v10 + int8(fi13))
        switch 550 {
        case 0:
            _ = len(v2)
        case 2147483647:
            v10 = v10
        case 4294967295:
            if (1e10 < v14) {
                v1 = v1
                for fi15 := 0; fi15 < 10; fi15++ {
                    if false {
                        v16 := GenS2{F0: 42, F1: 65535}
                        v14 += (v14 - float32(0.1))
                        switch fi15 {
                        case 4294967295:
                            switch fi15 {
                            case -9223372036854775808:
                                var v17 uint64 = 42
                                _ = v17
                            case 128:
                                v10 -= 11
                            case 0:
                                v18 := GenS3{F0: (fi15 ^ (fi13 - -2147483648)), F1: -122}
                                _ = v18
                            default:
                                switch fi15 {
                                case 230:
                                    v19 := GenS2{F0: (-(v10)), F1: v4}
                                    _ = v19
                                case 588:
                                    v10 += v10
                                }
                            }
                        }
                        _ = v16
                    }
                    _ = fi15
                }
            } else {
                v1 = v1
            }
        }
        _ = fi13
        _ = v14
    }
    v4 = (v4 / ((v4) | 1))
    if false {
        switch -9223372036854775808 {
        case 981:
            var v20 uint = uint(v4)
            _ = v20
        }
    }
    switch 128 {
    case 128:
        v4 -= (v4 % ((v4) | 1))
    case 1:
        v21 := GenS1{F0: 42, F1: (((v1 + v1) + (((v1 + ("\n\t" + v1)) + (v1 + (v1 + v1))) + (v1 + v1))) + v1)}
        _ = v21
    }
    var v22 int32 = 2147483647
    v23 := GenS4{F0: 2, F1: 255}
    switch 32768 {
    case -1:
        switch v1 {
        case "":
            for fi24 := 0; fi24 < 7; fi24++ {
                v25 := GenS3{F0: (-(fi24)), F1: (v10 | v10)}
                _ = fi24
                _ = v25
            }
        default:
            _ = len(v3)
        }
    default:
        if (!(!false)) {
            switch 18 {
            case 0:
                switch 128 {
                case -9223372036854775808:
                    var v26 float64 = 1e10
                    _ = v26
                }
            }
            for fi27 := 0; fi27 < 12; fi27++ {
                v28 := GenS3{F0: fi27, F1: 0}
                _ = fi27
                _ = v28
            }
            fmt.Println(v1)
        }
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        if len(v3) > 0 {
            v3[0] = 0
        }
        for fi30 := 0; fi30 < 13; fi30++ {
            fmt.Println(v22, v4, fi29)
            var v31 uint16 = v4
            _ = fi30
            _ = v31
        }
        var v32 string = ((v1 + (("foo bar" + "foo bar") + (v1 + "a"))) + ((v1 + v1) + "hello"))
        _ = fi29
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
