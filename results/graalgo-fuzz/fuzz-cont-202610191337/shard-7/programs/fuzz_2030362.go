package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 int16 = 32767
    var v2 string = ("foo bar" + (("a" + "") + ("" + "a")))
    var v3 uint8 = 127
    v4 := []uint32{509, 65535}
    v5 := make(map[uint32]uint32)
    var v6 int32 = -248
    var v7 int8 = -128
    fmt.Println(v6, v1, v1)
    var v8 int16 = (v1 / ((v1) | 1))
    v7 = (v7 % ((0) | 1))
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 uint64 = 105
            v7 = (v7 + (v7 % ((v7) | 1)))
            if ((v2 + "\n\t") >= "foo bar") {
                v6 += (^(v6))
            }
            v7 -= 42
            fmt.Println(fi10)
            var v12 uint16 = 2
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    var v13 bool = true
    var v14 float64 = 0.0
    _ = len(v4)
    var v15 float32 = -2.25
    for fi16 := 0; fi16 < 13; fi16++ {
        if len(v4) > 0 {
            v4[0] = 2
        }
        _ = fi16
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        switch (fi17 << ((uint(v3)) & 63)) {
        case -710:
            v6 -= int32(fi17)
        case 127:
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 uint16 = 873
                v20 := GenS1{F0: "foo bar"}
                v6 += (v6 - v6)
                _ = fi18
                _ = v19
                _ = v20
            }
        case -2147483648:
            if v13 {
                v1 = -934
                if v13 {
                    fmt.Println(v8, v2)
                    for fi21 := 0; fi21 < 9; fi21++ {
                        v5[998] = 4294967295
                        _ = fi21
                    }
                } else {
                    var v22 uint32 = 989
                    _ = v22
                }
            } else {
                fmt.Println(v3)
                v14 = v14
                if (2147483647 != 0) {
                    v23 := GenS2{F0: 42, F1: 1}
                    v24 := GenS3{F0: 1}
                    _ = v23
                    _ = v24
                }
            }
        default:
            v25 := GenS2{F0: 2147483648, F1: 42}
            _ = v25
        }
        var v26 bool = ((fi17 * int(v1)) <= (fi17 * 2147483647))
        if v13 {
            var v27 int = fi17
            _ = v27
        }
        _ = fi17
        _ = v26
    }
    if (1 < -1) {
        var v28 int = -144
        v2 = "a"
        _ = v28
    } else {
        var v29 float64 = v14
        var v30 int8 = (^(v7))
        var v31 float32 = 1e10
        _ = v29
        _ = v30
        _ = v31
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
