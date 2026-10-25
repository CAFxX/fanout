package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int; F1 uint32 }

func main() {
    var v1 uint16 = 65535
    v2 := []uint32{}
    var v3 int16 = -987
    v4 := make(map[uint32]bool)
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 uint64 = 18446744073709551615
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS1{F0: ""}
            if (false || true) {
                switch ("a" + "foo bar") {
                case "foo bar":
                    v9 := GenS2{F0: 1000, F1: -2147483648}
                    _ = v9
                case "\n\t":
                    v6 -= 0
                default:
                    switch "foo bar" {
                    case "foo bar":
                        _ = len(v2)
                    case "hello":
                        switch "foo bar" {
                        case "hello":
                            v3 -= (v3 << ((890) & 63))
                        default:
                            switch (-(fi5)) {
                            case 256:
                                delete(v4, 2147483647)
                            case 2147483647:
                                v3 = -1
                                v10 := GenS1{F0: ""}
                                if true {
                                    v11 := GenS2{F0: -32768, F1: 109}
                                    fmt.Println(fi5, fi5)
                                    v12 := v4[42]
                                    fmt.Println(fi7)
                                    _ = v11
                                    _ = v12
                                } else {
                                    v6 += v6
                                    var v13 float32 = 1e10
                                    _ = v13
                                }
                                _ = v10
                            case 32768:
                                var v14 int32 = 1
                                _ = v14
                            default:
                                if true {
                                    v15 := GenS3{F0: (fi5 >> ((2) & 63)), F1: 242}
                                    v6 -= v6
                                    v3 += (-(v3))
                                    _ = v15
                                }
                            }
                        }
                    case "\n\t":
                        var v16 int = fi5
                        _ = v16
                    }
                }
                v3 = 1
            }
            var v17 int = fi7
            v18 := GenS2{F0: v3, F1: 2147483647}
            _ = fi7
            _ = v17
            _ = v18
            _ = v8
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi5
        _ = v6
    }
    v19 := GenS1{F0: "hello"}
    v20 := GenS1{F0: ""}
    var v21 float64 = (1.5 - -2.25)
    v22 := v4[2147483647]
    if false {
        var v23 int16 = (v3 ^ v3)
        v3 = 42
        if v22 {
            v21 += (((0.1 / ((v21 + (-2.25 / 1.5)) + (v21 + v21))) / (v21 - v21)) - float64(0.0))
            v3 -= v23
        } else {
            var v24 uint64 = 1
            if len(v2) > 0 {
                v2[0] = 42
            }
            _ = v24
        }
        _ = v23
    } else {
        var v25 int16 = (v3 * (v3 ^ v3))
        _ = v25
    }
    v21 = 1.5
    var v26 uint8 = 71
    if v22 {
        _ = len(v2)
        v27 := GenS2{F0: v3, F1: -2147483648}
        _ = v27
    } else {
        var v28 float64 = v21
        _ = v28
    }
    fmt.Println(v3, v26, v3)
    v29 := GenS3{F0: 1, F1: 204}
    var v30 int32 = 436
    v3 = (v3 % ((v3) | 1))
    v21 += (1e10 * (((float64(-2.25) - (3.14159 / v21)) * float64(0.0)) - 0.0))
    var v31 int = 9223372036854775807
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v32 := GenS3{F0: 8, F1: 2}
    for fi33 := 0; fi33 < 3; fi33++ {
        v1 = (v1 * (v1 & 464))
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
