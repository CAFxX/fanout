package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint16; F1 int; F2 int16 }
type GenS3 struct { F0 int8; F1 uint32 }
func genFn0(p0 uint32) uint32 {
    var v5 int64 = -756
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: 42, F1: 1}
        var v8 string = ("a" + ("foo bar" + "\n\t"))
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v5
    return (p0 * (p0 & 70))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = 0.0
    var v3 int64 = 43
    v4 := []uint32{362, 404, 2}
    var v9 int32 = 1
    switch int(v9) {
    case 9223372036854775807:
        var v10 string = ("hello" + "a")
        _ = v10
    default:
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := GenS1{F0: 4294967295, F1: -1}
            v3 += v3
            _ = fi11
            _ = v12
        }
    }
    v9 += v9
    v13 := GenS2{F0: 0, F1: -320, F2: 1}
    if (!true) {
        _ = len(v1)
        switch 4294967295 {
        case 32767:
            var v14 int64 = (v3 | v3)
            _ = v14
        case 168:
            var v15 int64 = (v3 * (v3 & (v3 * v3)))
            _ = v15
        case 42:
            for fi16 := 0; fi16 < 5; fi16++ {
                for fi17 := 0; fi17 < 15; fi17++ {
                    fmt.Println(v9, v2, fi16)
                    v18 := GenS1{F0: 869, F1: -2147483648}
                    if (42 != 128) {
                        if (4294967295 <= 42) {
                            fmt.Println(v2)
                            v9 += v9
                        }
                        v19 := GenS1{F0: 137, F1: -2147483648}
                        v3 = (v3 | v3)
                        _ = v19
                    } else {
                        if (((("hello" + ("hello" + "hello")) + "\n\t") + (("foo bar" + "foo bar") + ("" + ((("a" + "") + "hello") + "")))) == ("\n\t" + "")) {
                            var v20 uint = 152
                            _ = v20
                        } else {
                            _ = len(v1)
                        }
                        switch -1 {
                        case 1:
                            var v21 float32 = v2
                            _ = v21
                        case 127:
                            var v22 int16 = 42
                            _ = v22
                        default:
                            v23 := GenS3{F0: 100, F1: 0}
                            _ = v23
                        }
                    }
                    _ = fi17
                    _ = v18
                }
                v3 += (v3 / (((v3 | (v3 % ((0) | 1)))) | 1))
                _ = fi16
            }
        }
        v24 := v1[127]
        _ = v24
    }
    v25 := GenS1{F0: 4294967295, F1: (^(v9))}
    for fi26 := 0; fi26 < 8; fi26++ {
        fmt.Println(v2, v2)
        var v27 uint32 = 65535
        _ = fi26
        _ = v27
    }
    var v28 float32 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v9
}
