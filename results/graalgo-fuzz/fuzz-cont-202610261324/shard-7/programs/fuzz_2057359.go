package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 string; F2 uint8 }

func main() {
    var v1 uint = 0
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    v4 := []string{"foo bar", "a"}
    var v5 float32 = -2.25
    var v6 float64 = 3.14159
    if (false && (!("\n\t" <= ("hello" + (("\n\t" + "") + ("" + ((("foo bar" + "a") + "\n\t") + (("foo bar" + "hello") + "")))))))) {
        v6 += v6
        fmt.Println(v1, v6, v1)
    } else {
        for fi7 := 0; fi7 < 11; fi7++ {
            if false {
                v8 := GenS1{F0: 4294967295, F1: 42}
                switch fi7 {
                case -1:
                    fmt.Println(v1)
                default:
                    var v9 uint64 = 1
                    _ = v9
                }
                v5 = ((((v5 + 1.5) / (((-2.25 * 3.14159) + v5) / 1.5)) * 0.0) * (0.1 + (((v5 + (0.1 - (1.5 / (3.14159 - v5)))) / v5) + (float32((3.14159 * v6)) / -2.25))))
                _ = v8
            } else {
                v6 += 0.0
                v5 += (0.1 - float32((1e10 * (v6 + 3.14159))))
                fmt.Println(v5, v5)
            }
            fmt.Println(v6, fi7, v6)
            _ = fi7
        }
    }
    var v10 uint32 = 2147483647
    v11 := GenS1{F0: 2, F1: 2}
    fmt.Println(v1, v5, v10)
    if (1 <= -1) {
        v4 = append(v4, ("foo bar" + "a"))
    }
    v12 := GenS2{F0: v1}
    v13 := GenS3{F0: (v10 / ((v10) | 1))}
    switch (("a" + "\n\t") + ("\n\t" + "\n\t")) {
    case "foo bar":
        var v14 int32 = -1
        _ = v14
    case "hello":
        switch 9223372036854775807 {
        case 42:
            for fi15 := 0; fi15 < 1; fi15++ {
                if (!((v10 + (v10 | (v10 | 777))) == (v10 + v10))) {
                    var v16 int16 = 326
                    v17 := GenS4{F0: 42}
                    switch fi15 {
                    case 4294967295:
                        if false {
                            var v18 uint64 = 18446744073709551615
                            v1 = v1
                            _ = v18
                        }
                    case 0:
                        var v19 int64 = 0
                        _ = v19
                    }
                    _ = v16
                    _ = v17
                }
                switch (fi15 | fi15) {
                case 355:
                    var v20 string = "foo bar"
                    _ = v20
                case 2147483647:
                    if ((((v6 + (v6 * v6)) * v6) * v6) == (v6 * v6)) {
                        var v21 uint8 = 42
                        var v22 uint16 = 0
                        _ = v21
                        _ = v22
                    } else {
                        var v23 uint8 = 255
                        for fi24 := 0; fi24 < 15; fi24++ {
                            v25 := GenS4{F0: 847}
                            var v26 int16 = 42
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        _ = v23
                    }
                }
                var v27 uint8 = 1
                _ = fi15
                _ = v27
            }
        default:
            fmt.Println(v1)
        }
    case "\n\t":
        fmt.Println(v5, v5)
    default:
        v28 := GenS5{F0: 42, F1: ("hello" + "foo bar"), F2: 147}
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
