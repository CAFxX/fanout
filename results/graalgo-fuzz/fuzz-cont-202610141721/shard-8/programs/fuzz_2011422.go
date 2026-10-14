package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 string }

func main() {
    var v1 uint16 = 32767
    var v2 int64 = -107
    v3 := make(map[uint32]string)
    var v4 int64 = v2
    v5 := []uint32{2, 367, 4294967295}
    var v6 int64 = (v2 >> ((0) & 63))
    var v7 int8 = 42
    for fi8 := 0; fi8 < 7; fi8++ {
        fmt.Println(v1, v2)
        _ = fi8
    }
    v1 = (v1 ^ 255)
    fmt.Println(v7, v7, v7)
    if (false || true) {
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 uint32 = 4294967295
            _ = fi9
            _ = v10
        }
        if false {
            fmt.Println(v7)
        } else {
            var v11 string = ""
            v6 -= (^(v4))
            _ = len(v3)
            _ = v11
        }
        var v12 int32 = -1
        _ = v12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := GenS1{F0: 219, F1: 42}
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: 255, F1: 1}
    v16 := GenS1{F0: 451, F1: 715}
    v7 = v7
    if false {
        v3[2147483647] = "\n\t"
        for fi17 := 0; fi17 < 17; fi17++ {
            fmt.Println(v1)
            switch "hello" {
            case "":
                var v18 uint32 = 2147483647
                _ = v18
            case "\n\t":
                if (-962 < int64(v7)) {
                    if false {
                        if len(v5) > 0 {
                            v5[0] = 1
                        }
                    }
                }
            default:
                if (true || false) {
                    v4 += (^(v6))
                }
            }
            switch fi17 {
            case 42:
                v19 := GenS2{F0: 42, F1: 255}
                _ = v19
            case -9223372036854775808:
                v1 -= (v1 * (^(v1)))
            case 127:
                var v20 uint16 = (v1 << ((585) & 63))
                _ = v20
            default:
                for fi21 := 0; fi21 < 5; fi21++ {
                    v22 := GenS3{F0: int8(fi17), F1: 2147483647}
                    fmt.Println(fi17, v7)
                    v2 = (v4 ^ v2)
                    _ = fi21
                    _ = v22
                }
            }
            _ = fi17
        }
        var v23 int16 = 1000
        _ = v23
    }
    fmt.Println(v4, v4)
    _ = len(v5)
    v5 = append(v5, 4294967295)
    fmt.Println(v4, v6)
    var v24 string = "foo bar"
    v25 := GenS4{F0: v24}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
