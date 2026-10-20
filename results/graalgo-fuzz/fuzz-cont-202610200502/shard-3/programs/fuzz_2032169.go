package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 int32; F1 int8 }

func main() {
    v1 := []uint32{}
    var v2 int = 2147483647
    var v3 string = ("\n\t" + "foo bar")
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (-9223372036854775808 >= 42) {
        var v4 int16 = 0
        _ = v4
    } else {
        v5 := GenS1{F0: 461, F1: 127, F2: 65535}
        _ = v5
    }
    if (0 >= 127) {
        v2 += v2
        if (1 != -469) {
            if len(v1) > 0 {
                v1[0] = 969
            }
            v3 = (((("foo bar" + "hello") + "foo bar") + (("" + v3) + v3)) + (v3 + "foo bar"))
        }
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS2{F0: 517, F1: 4294967295}
        v8 := GenS3{F0: 419, F1: 255}
        var v9 string = v3
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := GenS4{F0: true, F1: 32767}
    var v11 uint32 = 167
    v12 := GenS5{F0: 1, F1: -72}
    for fi13 := 0; fi13 < 6; fi13++ {
        if (!false) {
            _ = len(v1)
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi13
    }
    v1 = append(v1, (v11 % ((v11) | 1)))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (-291 <= -798) {
        for fi14 := 0; fi14 < 7; fi14++ {
            if true {
                v11 -= v11
            }
            var v15 uint32 = 42
            switch ((v3 + "") + v3) {
            case "\n\t":
                fmt.Println(v11, v3, v2)
            }
            _ = fi14
            _ = v15
        }
    } else {
        var v16 uint32 = 2147483647
        if false {
            if true {
                fmt.Println(v3, v2)
                fmt.Println(v11, v3)
                v1 = append(v1, v11)
            } else {
                v16 += 65535
                for fi17 := 0; fi17 < 15; fi17++ {
                    if true {
                        if (179 <= 1) {
                            var v18 int8 = 81
                            _ = v18
                        }
                        fmt.Println(v16, fi17, v16)
                        v3 = v3
                    }
                    fmt.Println(v16, v11, fi17)
                    _ = fi17
                }
            }
        }
        _ = v16
    }
    switch "\n\t" {
    case "":
        switch "hello" {
        case "a":
            var v19 float32 = float32(0.0)
            _ = v19
        case "foo bar":
            var v20 float64 = 1e10
            _ = v20
        default:
            _ = len(v1)
        }
    default:
        if false {
            v11 -= (v11 / (((v11 / (((v11 * 4294967295)) | 1))) | 1))
            v3 = v3
            v1 = append(v1, 2147483647)
        } else {
            var v21 uint32 = v11
            v11 = 884
            _ = len(v1)
            _ = v21
        }
    }
    v11 = (v11 * 42)
    v3 = ("\n\t" + v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
}
