package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }
func genFn0() string {
    var v5 int8 = 57
    v6 := []int{128, 130}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = v5
    _ = v6
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = -1
    v3 := []int{9223372036854775807, 4294967295}
    var v4 bool = true
    switch ("a" + "foo bar") {
    case "a":
        var v7 int = 146
        _ = v7
    }
    _ = len(v3)
    var v8 int = 32768
    var v9 float32 = 1e10
    for fi10 := 0; fi10 < 3; fi10++ {
        fmt.Println(fi10, v2)
        v11 := GenS1{F0: 435, F1: -1, F2: (!false)}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 int32 = (v2 % (((v2 * v2)) | 1))
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    var v15 uint = 4294967295
    fmt.Println(v2)
    var v16 uint = v15
    v4 = (!v4)
    switch v8 {
    case -9223372036854775808:
        _ = len(v3)
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if ((!(!(2 <= uint32(v16)))) || (((v2 | v2) == (v2 - (^(v2)))) || v4)) {
        for fi17 := 0; fi17 < 7; fi17++ {
            for fi18 := 0; fi18 < 9; fi18++ {
                fmt.Println(v2, v4)
                v19 := GenS1{F0: 0, F1: int16(fi17), F2: v4}
                _ = fi18
                _ = v19
            }
            switch "foo bar" {
            case "a":
                v1[2147483648] = (-(v8))
            case "foo bar":
                var v20 string = "\n\t"
                _ = v20
            case "\n\t":
                if v4 {
                    v21 := GenS2{F0: 204}
                    if (false && v4) {
                        fmt.Println(v4, v16)
                        v22 := GenS3{F0: (v2 / (((v2 / ((617) | 1))) | 1))}
                        _ = v22
                    } else {
                        fmt.Println(v15, v8)
                        for fi23 := 0; fi23 < 10; fi23++ {
                            var v24 uint64 = uint64(v15)
                            _ = fi23
                            _ = v24
                        }
                        v25 := GenS2{F0: 2}
                        _ = v25
                    }
                    fmt.Println(v8, v4)
                    _ = v21
                } else {
                    var v26 int = (v8 ^ v8)
                    _ = v26
                }
            default:
                var v27 uint32 = 354
                _ = v27
            }
            fmt.Println(v16, fi17)
            _ = fi17
        }
    } else {
        var v28 string = "foo bar"
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
