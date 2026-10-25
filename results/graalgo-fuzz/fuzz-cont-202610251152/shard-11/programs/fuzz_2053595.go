package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int16; F1 string; F2 int16 }

func main() {
    var v1 uint64 = 0
    v2 := []uint32{}
    v3 := make(map[int]uint32)
    var v4 int32 = -878
    var v5 uint32 = 2147483647
    var v6 uint = 63
    var v7 int = 4294967295
    fmt.Println(v6, v4, v6)
    for fi8 := 0; fi8 < 12; fi8++ {
        fmt.Println(v7, v1)
        var v9 int16 = 512
        var v10 uint8 = 42
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: (v1 + uint64(v6)), F1: 437, F2: 42}
    fmt.Println(v5)
    v1 -= v1
    if ((v4 > 42) && (!(549 == (v6 - v6)))) {
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 int = (fi12 ^ v7)
            _ = fi12
            _ = v13
        }
        if (!false) {
            var v14 int64 = -903
            _ = v14
        }
    } else {
        if false {
            v15 := GenS2{F0: v6}
            _ = v15
        } else {
            fmt.Println(v5, v6)
            switch ("foo bar" + "a") {
            case "a":
                v1 = (v1 % (((v1 + (v1 & (^(v1))))) | 1))
            case "foo bar":
                var v16 uint = (v6 * (v6 | v6))
                _ = v16
            default:
                if true {
                    v5 += 2
                    var v17 float32 = (0.1 + 1e10)
                    switch "foo bar" {
                    case "\n\t":
                        var v18 uint64 = 1
                        _ = v18
                    }
                    _ = v17
                }
            }
            v1 = v1
        }
    }
    if true {
        delete(v3, v7)
        var v19 int16 = 643
        v2 = append(v2, (v5 >> ((255) & 63)))
        _ = v19
    } else {
        v7 += (v7 ^ (v7 >> ((v6) & 63)))
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 int8 = 1
        v7 = 42
        v22 := GenS3{F0: v4, F1: 42}
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        v24 := GenS1{F0: v1, F1: 17, F2: 255}
        _ = fi23
        _ = v24
    }
    v25 := GenS4{F0: -32768, F1: "\n\t", F2: -1}
    var v26 uint8 = 127
    if true {
        var v27 int32 = (v4 + v4)
        v27 -= (v27 & (v27 + v27))
        for fi28 := 0; fi28 < 13; fi28++ {
            v6 += v6
            v29 := GenS4{F0: 1, F1: "a", F2: 1}
            for fi30 := 0; fi30 < 4; fi30++ {
                var v31 int64 = 0
                v5 += 2147483648
                var v32 string = ""
                switch (^(fi28)) {
                case -13:
                    var v33 float32 = 3.14159
                    _ = v33
                case -1:
                    if false {
                        var v34 int8 = -27
                        _ = v34
                    } else {
                        if (uint(v26) <= v6) {
                            var v35 int64 = -9223372036854775808
                            var v36 float32 = -2.25
                            _ = v35
                            _ = v36
                        }
                    }
                }
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi28
            _ = v29
        }
        _ = v27
    }
    v5 -= 2147483647
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
