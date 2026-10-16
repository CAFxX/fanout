package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 0
    var v3 uint32 = 670
    v4 := make(map[int]string)
    var v5 float64 = -2.25
    switch 32768 {
    case -9223372036854775808:
        v6 := GenS1{F0: 1, F1: v3, F2: 2}
        _ = v6
    case 289:
        fmt.Println(v3, v3)
    case 0:
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 uint16 = 157
            _ = fi7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        switch 321 {
        case 127:
            for fi10 := 0; fi10 < 1; fi10++ {
                switch (("a" + "foo bar") + "hello") {
                case "hello":
                    var v11 float64 = v5
                    _ = v11
                case "\n\t":
                    if (0 > 0) {
                        v12 := v1["foo bar"]
                        var v13 float64 = ((-2.25 / ((v5 * ((v5 / v5) + float64((0.1 / 1e10)))) + v5)) / -2.25)
                        _ = v12
                        _ = v13
                    } else {
                        if (!true) {
                            var v14 int8 = 42
                            if (1 <= 65535) {
                                v15 := GenS1{F0: 9223372036854775807, F1: (v3 & (v3 % (((v3 & (v3 & (v3 ^ v3)))) | 1))), F2: 2}
                                fmt.Println(v14, fi10, v14)
                                _ = v15
                            }
                            _ = v14
                        } else {
                            switch "\n\t" {
                            case "hello":
                                var v16 bool = true
                                _ = v16
                            }
                            fmt.Println(fi9)
                        }
                    }
                default:
                    var v17 float64 = v5
                    _ = v17
                }
                _ = fi10
            }
        case -2147483648:
            v18 := GenS1{F0: -9223372036854775808, F1: v3, F2: (v2 * v2)}
            _ = v18
        case -32:
            for fi19 := 0; fi19 < 6; fi19++ {
                v20 := GenS1{F0: 42, F1: (v3 * 65535), F2: v2}
                _ = fi19
                _ = v20
            }
        default:
            v3 = (v3 | v3)
        }
        var v21 uint16 = 255
        var v22 int16 = 1
        _ = fi9
        _ = v21
        _ = v22
    }
    v5 += 3.14159
    v5 = 0.0
    var v23 bool = true
    if (!(42 < 100)) {
        switch "a" {
        case "foo bar":
            for fi24 := 0; fi24 < 14; fi24++ {
                v25 := GenS1{F0: 42, F1: 860, F2: (v2 & v2)}
                v26 := GenS1{F0: 1, F1: 0, F2: 2}
                var v27 int = 32767
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
        }
        v28 := v4[-1]
        _ = v28
    }
    var v29 int16 = 42
    _ = len(v4)
    for fi30 := 0; fi30 < 5; fi30++ {
        switch fi30 {
        case -9223372036854775808:
            v31 := []int{(-(fi30))}
            _ = v31
        case -622:
            v32 := []int{(-(fi30)), (fi30 % ((fi30) | 1)), fi30}
            _ = v32
        case 32768:
            v5 += (-2.25 + (v5 + v5))
        default:
            fmt.Println(fi30, fi30, fi30)
        }
        _ = fi30
    }
    v33 := GenS1{F0: -877, F1: (v3 * (v3 & (v3 >> ((65535) & 63)))), F2: (v2 * 1)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
