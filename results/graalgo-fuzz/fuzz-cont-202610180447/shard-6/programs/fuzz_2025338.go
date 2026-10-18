package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int, p1 int32) uint {
    var v7 uint = 956
    fmt.Println(v7, v7, v7)
    p0 = (p0 / ((p0) | 1))
    for fi8 := 0; fi8 < 14; fi8++ {
        if ((0.0 * -2.25) > 0.1) {
            switch fi8 {
            case 42:
                fmt.Println(v7, p1, p0)
            case -382:
                if true {
                    v7 -= 42
                    p0 = p0
                }
            case 866:
                var v9 float32 = 0.1
                _ = v9
            default:
                p1 = 42
            }
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := GenS1{F0: 42, F1: 113, F2: -128}
                if false {
                    v12 := GenS2{F0: int64(v7), F1: 0}
                    _ = v12
                } else {
                    var v13 string = ""
                    var v14 uint32 = 2147483647
                    _ = v13
                    _ = v14
                }
                _ = fi10
                _ = v11
            }
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v7
    return (v7 >> ((v7) & 63))
}

func main() {
    var v1 int8 = 100
    v2 := []uint32{4294967295, 519}
    var v3 uint8 = 128
    var v4 int16 = 335
    v5 := make(map[uint32]bool)
    var v6 uint = 0
    if true {
        fmt.Println(v6, v3, v6)
    } else {
        var v15 int64 = 1
        var v16 uint = v6
        _ = v15
        _ = v16
    }
    var v17 uint16 = 32767
    switch 9223372036854775807 {
    case 908:
        v6 -= genFn0(643, -2147483648)
    default:
        switch 1 {
        case 265:
            _ = len(v5)
        }
    }
    fmt.Println(v6, v3)
    switch "" {
    case "foo bar":
        switch "" {
        case "":
            var v18 float32 = 0.1
            _ = v18
        default:
            v6 = genFn0(2147483647, 994)
        }
    case "a":
        if ((v1 ^ -128) != (v1 * (v1 & v1))) {
            fmt.Println(v6, v1, v1)
            _ = len(v2)
            var v19 float32 = (((0.1 / float32((1e10 / float64((1e10 * (-2.25 + ((-2.25 / 1.0) / 1e10))))))) - 0.1) - (((1.5 / 0.1) + ((1.5 + 3.14159) / 1.0)) - 1.5))
            _ = v19
        }
    }
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    v3 -= (v3 % ((128) | 1))
    var v20 uint = (v6 * (v6 - (v6 % (((^(v6))) | 1))))
    var v21 int64 = 42
    v22 := GenS3{F0: v1}
    switch -543 {
    case -2147483648:
        v3 = v3
    default:
        if true {
            if len(v2) > 0 {
                v2[0] = 294
            }
            v4 += v4
        }
    }
    if (42 > 2147483647) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v23 := GenS2{F0: v21, F1: 128}
        v1 = 127
        _ = v23
    } else {
        v4 -= (v4 / (((v4 - v4)) | 1))
        for fi24 := 0; fi24 < 2; fi24++ {
            v25 := GenS2{F0: v21, F1: (^(v3))}
            delete(v5, 2147483647)
            _ = fi24
            _ = v25
        }
    }
    if false {
        switch "foo bar" {
        case "":
            _ = len(v2)
        case "\n\t":
            v20 -= 0
        case "hello":
            v1 = (v1 | (v1 + 0))
        default:
            fmt.Println(v17)
        }
        if false {
            if (((0 >= 18446744073709551615) && ((^(v17)) < uint16(v20))) && false) {
                var v26 uint16 = (^(v17))
                for fi27 := 0; fi27 < 9; fi27++ {
                    v17 = (v26 - v17)
                    fmt.Println(v21)
                    var v28 int32 = int32(v3)
                    _ = fi27
                    _ = v28
                }
                _ = v26
            }
        }
        v6 -= (v6 + v20)
    } else {
        var v29 uint32 = 320
        _ = v29
    }
    if (true && (!(106 >= 34))) {
        var v30 uint64 = uint64(v17)
        if true {
            var v31 string = ("hello" + "\n\t")
            v31 = (v31 + ("hello" + v31))
            var v32 float32 = (0.1 / (-2.25 * 1e10))
            _ = v31
            _ = v32
        } else {
            var v33 int16 = 32767
            v5[581] = false
            fmt.Println(v1)
            _ = v33
        }
        _ = v30
    } else {
        v4 = v4
    }
    fmt.Println(v17, v1, v21)
    var v34 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
