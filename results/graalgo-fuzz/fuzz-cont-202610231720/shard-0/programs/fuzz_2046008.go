package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint; F1 int64; F2 uint16 }
func genFn0(p0 bool) uint32 {
    switch "a" {
    case "hello":
        if p0 {
            v4 := GenS1{F0: 2147483647, F1: 255, F2: 128}
            _ = v4
        } else {
            var v5 uint64 = 28
            if false {
                for fi6 := 0; fi6 < 4; fi6++ {
                    fmt.Println(v5, p0)
                    v5 -= (^(v5))
                    v5 = v5
                    _ = fi6
                }
                v7 := GenS1{F0: 2147483648, F1: 65535, F2: 23}
                _ = v7
            }
            _ = v5
        }
    default:
        v8 := []int{}
        _ = v8
    }
    fmt.Println(p0)
    if false {
        if true {
            for fi9 := 0; fi9 < 16; fi9++ {
                p0 = ((1 != 9223372036854775807) && p0)
                var v10 string = "foo bar"
                var v11 int16 = 32767
                _ = fi9
                _ = v10
                _ = v11
            }
        } else {
            v12 := []int{-477, -2147483648, 9223372036854775807, 42}
            if (0 != 0) {
                for fi13 := 0; fi13 < 1; fi13++ {
                    p0 = true
                    p0 = ((!p0) || p0)
                    _ = fi13
                }
                for fi14 := 0; fi14 < 12; fi14++ {
                    if ("\n\t" <= ("\n\t" + ("hello" + "\n\t"))) {
                        var v15 uint32 = 2147483647
                        fmt.Println(fi14, v15)
                        _ = v15
                    } else {
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        if true {
                            switch -1 {
                            case 32768:
                                _ = len(v12)
                            case 42:
                                switch ("" + "hello") {
                                case "":
                                    var v16 uint16 = 0
                                    _ = v16
                                case "foo bar":
                                    fmt.Println(fi14, p0)
                                }
                            case 4294967295:
                                fmt.Println(fi14)
                            }
                        } else {
                            var v17 int8 = -128
                            _ = v17
                        }
                        p0 = p0
                    }
                    var v18 uint64 = 415
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    _ = fi14
                    _ = v18
                }
                var v19 float64 = (float64(1e10) + 0.1)
                _ = v19
            } else {
                v20 := GenS2{F0: false, F1: 0}
                fmt.Println(p0)
                _ = v20
            }
            _ = v12
        }
        var v21 uint16 = 2
        _ = v21
    } else {
        var v22 float32 = 1e10
        v22 = 1e10
        _ = v22
    }
    var v23 uint64 = 42
    v24 := GenS3{F0: 2, F1: 1, F2: 2}
    _ = p0
    _ = v23
    _ = v24
    return 1
}

func main() {
    var v1 uint16 = 452
    var v2 int8 = -4
    var v3 uint8 = 224
    for fi25 := 0; fi25 < 12; fi25++ {
        v3 = (v3 + (v3 & (v3 >> ((65535) & 63))))
        fmt.Println(v2, fi25, fi25)
        _ = fi25
    }
    var v26 int8 = v2
    if (1 != 0) {
        if (1 < 1) {
            var v27 float32 = float32(3.14159)
            fmt.Println(v2, v3, v3)
            _ = v27
        }
    }
    var v28 uint64 = 1
    switch 32767 {
    case 415:
        var v29 int32 = 0
        _ = v29
    case 14:
        var v30 uint8 = (v3 % ((v3) | 1))
        _ = v30
    case 769:
        if (false || true) {
            v26 -= v2
        }
    }
    if true {
        fmt.Println(v1, v3)
        fmt.Println(v26, v28)
    }
    v3 = 255
    for fi31 := 0; fi31 < 4; fi31++ {
        v2 = (-(v26))
        var v32 int8 = (v26 % (((v2 & (v2 * 52))) | 1))
        _ = fi31
        _ = v32
    }
    var v33 int32 = 0
    var v34 int8 = (v2 >> ((65535) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v33
    _ = v34
}
