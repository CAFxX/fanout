package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 float32 = (1.5 * (1.5 - 0.0))
    var v2 uint16 = 2
    v3 := []uint32{2147483648, 2147483647}
    var v4 float64 = float64(3.14159)
    v5 := make(map[uint32]bool)
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            v1 -= 1e10
            _ = fi7
        }
        switch ((("foo bar" + "a") + "\n\t") + "\n\t") {
        case "":
            _ = len(v5)
        case "\n\t":
            var v8 uint32 = 2147483648
            _ = v8
        case "a":
            var v9 uint64 = 4294967295
            _ = v9
        }
        _ = fi6
    }
    if (-9223372036854775808 != 0) {
        v10 := GenS1{F0: 58, F1: -1}
        _ = v10
    }
    v1 += v1
    switch "hello" {
    case "foo bar":
        v11 := GenS1{F0: 127, F1: 0}
        _ = v11
    case "hello":
        for fi12 := 0; fi12 < 18; fi12++ {
            for fi13 := 0; fi13 < 1; fi13++ {
                v2 = 0
                _ = fi13
            }
            switch "foo bar" {
            case "a":
                v2 -= (v2 - (v2 * v2))
            case "":
                v4 += (0.1 + -2.25)
            case "hello":
                var v14 uint32 = 2147483648
                _ = v14
            }
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 float32 = -2.25
                v1 -= v16
                v16 = v16
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v17 := GenS1{F0: -16, F1: 0}
                fmt.Println(v16, v16)
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi12
        }
    case "a":
        fmt.Println(v4)
    default:
        v18 := GenS2{F0: "a"}
        _ = v18
    }
    var v19 int = 1
    fmt.Println(v19)
    fmt.Println(v4, v1, v4)
    var v20 uint64 = 4294967295
    v4 -= (1e10 * 0.0)
    v20 -= 0
    fmt.Println(v2)
    var v21 uint = 2
    if false {
        var v22 string = "\n\t"
        if (float32(v4) >= 1e10) {
            switch v19 {
            case 32767:
                v23 := GenS2{F0: "hello"}
                _ = v23
            case -2147483648:
                var v24 uint32 = 0
                _ = v24
            }
            v22 = (v22 + (v22 + "foo bar"))
        } else {
            var v25 uint64 = 0
            switch v22 {
            case "hello":
                v26 := GenS3{F0: (v25 ^ v25)}
                _ = v26
            case "":
                var v27 int32 = 939
                _ = v27
            default:
                var v28 float32 = float32(1e10)
                _ = v28
            }
            _ = v25
        }
        for fi29 := 0; fi29 < 20; fi29++ {
            v4 = v4
            _ = fi29
        }
        _ = v22
    }
    fmt.Println(v20)
    var v30 int16 = -246
    var v31 uint32 = 2147483647
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
