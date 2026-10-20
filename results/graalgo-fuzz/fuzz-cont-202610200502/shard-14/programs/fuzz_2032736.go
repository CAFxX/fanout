package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int32 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 uint8; F2 int }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 int32 = 2147483647
    var v4 int8 = 100
    var v5 float32 = ((((1.5 + (1e10 + 0.0)) / 1.0) - ((((0.0 * ((1e10 * 1e10) + 1e10)) / 1.0) + (((1e10 + 0.1) / (1e10 / 0.1)) + 1e10)) * 1.5)) / (1.5 / ((3.14159 - 1e10) + 0.0)))
    v6 := []int{-571, 732, 9223372036854775807, int(v4)}
    v4 = (v4 / ((v4) | 1))
    var v7 float64 = 1e10
    v3 = 2147483647
    for fi8 := 0; fi8 < 19; fi8++ {
        switch fi8 {
        case -2147483648:
            v9 := GenS1{F0: fi8, F1: -859, F2: v3}
            _ = v9
        case 128:
            v3 = -2147483648
        case 0:
            _ = len(v6)
        }
        var v10 int8 = v4
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 int = (-(fi8))
            var v13 uint8 = 128
            var v14 string = "hello"
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi8
        _ = v10
    }
    var v15 uint = 4294967295
    var v16 bool = ("hello" < ("foo bar" + "foo bar"))
    v5 = (v5 * (float32((0.1 - v7)) - (0.1 - v5)))
    if (!v16) {
        delete(v1, (((("a" + ("a" + (("a" + "\n\t") + "foo bar"))) + "foo bar") + "foo bar") + "hello"))
        delete(v1, ("hello" + "foo bar"))
    } else {
        v17 := GenS2{F0: v4, F1: 0}
        var v18 uint = 65535
        _ = v17
        _ = v18
    }
    v19 := GenS3{F0: 0}
    v16 = (342 <= -2147483648)
    v15 = v15
    v4 = 1
    for fi20 := 0; fi20 < 13; fi20++ {
        if (!true) {
            var v21 int16 = 0
            var v22 uint = (v15 << (((^(v15))) & 63))
            switch fi20 {
            case 128:
                var v23 uint16 = 0
                _ = v23
            }
            _ = v21
            _ = v22
        } else {
            if (v16 || (!(!true))) {
                fmt.Println(v4)
                var v24 int16 = 1
                _ = v24
            }
            for fi25 := 0; fi25 < 14; fi25++ {
                switch fi25 {
                case -9223372036854775808:
                    fmt.Println(v4, v5)
                default:
                    var v26 float32 = (v5 / ((((1.5 / 1.0) * (float32(v7) - v5)) + 1.5) / (float32(v7) + v5)))
                    _ = v26
                }
                _ = fi25
            }
            if ((v5 / (float32(v7) / (((v5 - 1e10) * 1e10) - ((3.14159 - ((v5 * v5) * 1e10)) / (v5 * 0.0))))) < (v5 + (float32(3.14159) - (3.14159 - float32((v7 * v7)))))) {
                var v27 int = fi20
                if true {
                    var v28 uint16 = 32767
                    var v29 uint64 = 42
                    var v30 int8 = v4
                    _ = v28
                    _ = v29
                    _ = v30
                }
                _ = v27
            }
        }
        if v16 {
            v31 := GenS2{F0: (v4 - (v4 & (v4 & v4))), F1: -32768}
            _ = len(v2)
            v15 -= (v15 << ((v15) & 63))
            _ = v31
        }
        _ = fi20
    }
    v32 := GenS4{F0: (v4 * v4), F1: 132, F2: -1}
    v16 = (!v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
