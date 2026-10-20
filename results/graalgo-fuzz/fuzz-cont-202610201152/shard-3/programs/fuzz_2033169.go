package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS5 struct { F0 uint64; F1 int32 }
func genFn0(p0 uint8) uint {
    v4 := GenS1{F0: 2147483647, F1: 592, F2: 42}
    fmt.Println(p0)
    for fi5 := 0; fi5 < 17; fi5++ {
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 uint32 = 2147483647
            v7 -= v7
            var v8 string = "a"
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    _ = p0
    _ = v4
    return 18446744073709551615
}
func genFn1(p0 uint32, p1 int, p2 int) uint8 {
    if (-128 > 1) {
        v9 := GenS1{F0: (p0 / (((p0 << ((1) & 63))) | 1)), F1: 42, F2: 18446744073709551615}
        var v10 uint16 = 0
        _ = v10
        _ = v9
    } else {
        v11 := []int{(p1 + 42), -35}
        _ = v11
    }
    var v12 int64 = 1
    v13 := []int{p2}
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 float32 = 1e10
        switch ("" + "foo bar") {
        case "":
            if (true && true) {
                fmt.Println(p1, p2, p1)
            }
        case "a":
            switch ("foo bar" + ("a" + "foo bar")) {
            case "\n\t":
                p2 += (-(fi14))
            case "":
                v16 := GenS1{F0: 0, F1: (p0 - (p0 | p0)), F2: 65535}
                _ = v16
            case "hello":
                var v17 int32 = 989
                _ = v17
            default:
                for fi18 := 0; fi18 < 5; fi18++ {
                    v19 := GenS1{F0: (^(p0)), F1: 262, F2: 255}
                    v12 = v12
                    _ = fi18
                    _ = v19
                }
            }
        case "foo bar":
            var v20 int64 = int64(p0)
            _ = v20
        default:
            var v21 float32 = v15
            _ = v21
        }
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 234
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 0.1
    var v3 uint16 = 255
    var v22 int16 = -1
    var v23 uint8 = 226
    for fi24 := 0; fi24 < 11; fi24++ {
        v25 := v1[2]
        if (((v3 / (((v3 >> ((2) & 63))) | 1)) > 65535) || v25) {
            v26 := GenS2{F0: (v3 & (v3 << ((1) & 63)))}
            v23 -= (v23 / ((genFn1(1, fi24, 32768)) | 1))
            _ = v26
        } else {
            v23 = v23
        }
        v23 -= (^(v23))
        _ = fi24
        _ = v25
    }
    if false {
        v27 := []int{}
        _ = v27
    }
    fmt.Println(v23, v3, v2)
    v28 := GenS3{F0: v22, F1: -32768}
    if ((v3 & (v3 - (v3 / ((255) | 1)))) == (^(v3))) {
        v29 := GenS3{F0: v22, F1: (v22 & 1)}
        _ = v29
    } else {
        var v30 uint16 = v3
        if true {
            var v31 uint8 = v23
            v32 := GenS4{F0: (v3 & v3), F1: false, F2: 65535}
            for fi33 := 0; fi33 < 16; fi33++ {
                v34 := GenS2{F0: uint16(v31)}
                _ = fi33
                _ = v34
            }
            _ = v31
            _ = v32
        } else {
            v35 := v1[2]
            v36 := GenS5{F0: uint64(v3), F1: 1}
            var v37 int16 = v22
            _ = v35
            _ = v36
            _ = v37
        }
        _ = v30
    }
    var v38 uint8 = (^(v23))
    _ = len(v1)
    if true {
        v22 = v22
        var v39 string = "hello"
        var v40 int8 = -74
        _ = v39
        _ = v40
    }
    v3 = (v3 + 255)
    fmt.Println(v23)
    v41 := []int{-29}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v38
    _ = v41
}
