package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    var v3 int8 = 100
    var v4 int16 = 0
    v5 := []int{}
    _ = len(v5)
    var v6 uint32 = 116
    v7 := GenS1{F0: 4294967295}
    v8 := GenS1{F0: 260}
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 uint16 = 42
        _ = fi9
        _ = v10
    }
    v5 = append(v5, -694)
    v3 -= (-(v3))
    if true {
        v3 -= -1
        var v11 uint8 = 42
        if false {
            for fi12 := 0; fi12 < 6; fi12++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi12
            }
        } else {
            var v13 uint = 255
            _ = v13
        }
        _ = v11
    } else {
        var v14 int16 = 1000
        v3 -= -55
        for fi15 := 0; fi15 < 18; fi15++ {
            v4 -= 106
            var v16 uint = 255
            fmt.Println(v14, v16, v1)
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        if v1 {
            var v18 string = "a"
            _ = v18
        } else {
            for fi19 := 0; fi19 < 10; fi19++ {
                v20 := GenS1{F0: 4294967295}
                v21 := GenS2{F0: 42}
                switch "a" {
                case "a":
                    v22 := GenS2{F0: 221}
                    _ = v22
                case "\n\t":
                    v5 = append(v5, fi17)
                case "hello":
                    v23 := GenS1{F0: 998}
                    _ = v23
                }
                v1 = ((!(0 <= 65535)) || v1)
                _ = fi19
                _ = v20
                _ = v21
            }
            v4 += int16(fi17)
        }
        var v24 int16 = (-(v4))
        _ = fi17
        _ = v24
    }
    if len(v5) > 0 {
        v5[0] = 32767
    }
    fmt.Println(v4)
    v3 -= v3
    v1 = true
    v6 = v6
    switch "" {
    case "\n\t":
        v3 += (v3 * (v3 >> ((2) & 63)))
    case "a":
        var v25 int8 = (v3 << ((0) & 63))
        _ = v25
    default:
        for fi26 := 0; fi26 < 18; fi26++ {
            var v27 int32 = 0
            _ = fi26
            _ = v27
        }
    }
    v6 = 2147483647
    if true {
        v28 := GenS2{F0: 128}
        v4 -= (^(v4))
        _ = v28
    } else {
        if (!(v1 && false)) {
            if false {
                for fi29 := 0; fi29 < 3; fi29++ {
                    v30 := GenS3{F0: 38, F1: 615}
                    _ = fi29
                    _ = v30
                }
                switch "" {
                case "":
                    v31 := GenS4{F0: v4}
                    _ = v31
                case "foo bar":
                    var v32 int64 = -401
                    _ = v32
                case "a":
                    fmt.Println(v1, v3, v6)
                }
                v6 = (v6 % (((v6 >> ((1) & 63))) | 1))
            }
            var v33 int8 = v3
            var v34 float32 = (((float32(1e10) - 0.1) - (-2.25 - 0.0)) - (1e10 * (float32(1e10) - 1.5)))
            _ = v33
            _ = v34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
