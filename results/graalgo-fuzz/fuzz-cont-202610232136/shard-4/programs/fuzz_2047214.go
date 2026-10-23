package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS2 struct { F0 bool; F1 int8; F2 int16 }
type GenS3 struct { F0 int64; F1 uint32 }
func genFn0(p0 uint16) uint {
    v7 := GenS1{F0: (^(p0)), F1: 42, F2: 1}
    if false {
        v8 := []int{416, 128, 1, 42}
        p0 += p0
        _ = v8
    } else {
        if true {
            v9 := []int{231, 4294967295}
            var v10 uint = 4294967295
            _ = v10
            _ = v9
        }
    }
    _ = p0
    _ = v7
    return 0
}
func genFn1(p0 uint16, p1 int16, p2 string) int64 {
    v11 := []int{32768, 438, -1}
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            v14 := GenS1{F0: (p0 & (p0 + p0)), F1: int32(p1), F2: 128}
            p2 = "foo bar"
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 13; fi15++ {
            if len(v11) > 0 {
                v11[0] = (fi15 << ((0) & 63))
            }
            _ = fi15
        }
        _ = fi12
    }
    var v16 int16 = (p1 >> ((278) & 63))
    var v17 int = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v16
    _ = v17
    return -1
}

func main() {
    var v1 uint16 = 255
    v2 := []uint32{}
    var v3 uint = 42
    var v4 uint64 = 2
    v5 := make(map[uint32]string)
    var v6 uint32 = 42
    switch 0 {
    case 32767:
        for fi18 := 0; fi18 < 7; fi18++ {
            v1 += (^(v1))
            _ = fi18
        }
    case -260:
        if false {
            if true {
                v2 = append(v2, v6)
                var v19 bool = false
                _ = v19
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v20 uint16 = v1
            _ = v20
        }
    case 128:
        fmt.Println(v6, v4)
    }
    v6 -= (v6 % ((0) | 1))
    v21 := v5[(v6 * (^(v6)))]
    v6 = v6
    var v22 uint64 = 644
    v1 = 2
    var v23 string = v21
    switch -566 {
    case -1:
        var v24 uint = v3
        _ = v24
    case 129:
        var v25 uint = v3
        _ = v25
    }
    v26 := GenS1{F0: 80, F1: 1, F2: 208}
    if false {
        fmt.Println(v3, v21, v3)
        switch "\n\t" {
        case "hello":
            fmt.Println(v1, v21)
        case "foo bar":
            fmt.Println(v3, v21)
        case "\n\t":
            var v27 uint64 = (v4 << ((genFn0(v1)) & 63))
            _ = v27
        }
        var v28 int = -2147483648
        _ = v28
    } else {
        switch ("a" + (v21 + ("foo bar" + v21))) {
        case "hello":
            var v29 int64 = -9223372036854775808
            _ = v29
        case "a":
            for fi30 := 0; fi30 < 10; fi30++ {
                v6 = 475
                v2 = append(v2, v6)
                _ = fi30
            }
        default:
            v31 := GenS1{F0: (v1 * (v1 % ((1) | 1))), F1: 1, F2: 128}
            _ = v31
        }
    }
    _ = len(v5)
    v22 -= (v4 ^ (v4 - (v22 + (v4 % ((uint64(v3)) | 1)))))
    var v32 int64 = -259
    for fi33 := 0; fi33 < 4; fi33++ {
        v22 += v22
        v34 := GenS2{F0: false, F1: 13, F2: -32768}
        var v35 uint16 = (^(v1))
        _ = fi33
        _ = v34
        _ = v35
    }
    switch 1 {
    case -470:
        switch ("a" + ("a" + v23)) {
        case "hello":
            _ = len(v5)
        case "\n\t":
            v36 := GenS3{F0: v32, F1: (v6 + v6)}
            _ = v36
        default:
            if len(v2) > 0 {
                v2[0] = v6
            }
        }
    default:
        for fi37 := 0; fi37 < 11; fi37++ {
            v3 = (v3 << (((v3 + 42)) & 63))
            _ = fi37
        }
    }
    var v38 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v32
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
