package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int16; F2 int64 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint64; F1 int; F2 uint64 }

func main() {
    v1 := []int{86, 127}
    v2 := make(map[int]uint32)
    var v3 uint8 = 42
    v3 = v3
    v3 += v3
    for fi4 := 0; fi4 < 19; fi4++ {
        v5 := GenS1{F0: 1}
        var v6 uint = 509
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 0}
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint8 = (^(v3))
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v11 float32 = float32((float64(float32(1e10)) / 1e10))
    v12 := GenS2{F0: -1, F1: 427, F2: -1}
    var v13 int16 = -783
    if true {
        var v14 uint32 = 2147483647
        for fi15 := 0; fi15 < 15; fi15++ {
            v16 := GenS3{F0: (false && false), F1: (v14 & 32), F2: 444}
            v17 := GenS4{F0: 0, F1: (fi15 & fi15), F2: 18446744073709551615}
            _ = fi15
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 18; fi18++ {
            if (65535 > 544) {
                switch "" {
                case "foo bar":
                    var v19 float64 = (0.0 + float64(v11))
                    _ = v19
                case "\n\t":
                    if (fi18 == (fi18 * 127)) {
                        v20 := v2[fi18]
                        _ = v20
                    }
                case "hello":
                    if (42 != uint64(v14)) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v21 int16 = -251
                        v22 := GenS1{F0: (v13 - (v21 | v21))}
                        _ = v21
                        _ = v22
                    } else {
                        fmt.Println(fi18, v14, fi18)
                        fmt.Println(v11)
                        var v23 float64 = 1.5
                        _ = v23
                    }
                }
                if (0 <= -2147483648) {
                    v24 := GenS3{F0: false, F1: (v14 - (^(v14))), F2: 787}
                    _ = v24
                } else {
                    _ = len(v2)
                    v3 += v3
                }
            }
            for fi25 := 0; fi25 < 16; fi25++ {
                var v26 float64 = (0.0 / 0.1)
                var v27 uint8 = v3
                var v28 bool = false
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
            _ = fi18
        }
        _ = v14
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        fmt.Println(v11, v3)
        var v30 float32 = (v11 * v11)
        _ = fi29
        _ = v30
    }
    var v31 uint8 = (^(v3))
    var v32 uint64 = 62
    for fi33 := 0; fi33 < 11; fi33++ {
        v34 := v2[int(v32)]
        _ = fi33
        _ = v34
    }
    v11 = (v11 / 1.5)
    fmt.Println(v13, v32, v13)
    switch 724 {
    case 148:
        v35 := GenS4{F0: (v32 >> ((191) & 63)), F1: 706, F2: v32}
        _ = v35
    case 127:
        v32 -= v32
    case 270:
        var v36 string = ""
        _ = v36
    default:
        v13 -= -32768
    }
    var v37 uint8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v37
    _ = v7
}
