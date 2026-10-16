package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int64; F2 string }

func main() {
    v1 := make(map[string]int)
    v2 := []string{"\n\t", "\n\t", "foo bar"}
    v3 := make(map[uint32]bool)
    v4 := make(map[string]uint32)
    var v5 uint64 = 18446744073709551615
    var v6 int64 = 172
    if ((v5 | v5) <= v5) {
        switch -804 {
        case -2147483648:
            if (false || (!true)) {
                var v7 int = 32768
                var v8 float32 = -2.25
                _ = v7
                _ = v8
            } else {
                for fi9 := 0; fi9 < 10; fi9++ {
                    var v10 string = "foo bar"
                    var v11 int64 = (v6 ^ (v6 | v6))
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
        case -272:
            v3[329] = (-1 >= -651)
        }
    } else {
        for fi12 := 0; fi12 < 7; fi12++ {
            if false {
                if false {
                    for fi13 := 0; fi13 < 9; fi13++ {
                        var v14 string = ("a" + "a")
                        _ = fi13
                        _ = v14
                    }
                    if false {
                        var v15 int = 398
                        var v16 bool = (11 < 2)
                        _ = v15
                        _ = v16
                    }
                    if (fi12 > fi12) {
                        v6 = 191
                        fmt.Println(v6)
                    } else {
                        var v17 float32 = (1.5 - 0.0)
                        v6 = v6
                        v6 += v6
                        _ = v17
                    }
                } else {
                    for fi18 := 0; fi18 < 14; fi18++ {
                        var v19 int8 = 1
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v6 += v6
                        _ = fi18
                        _ = v19
                    }
                    v20 := GenS1{F0: v6}
                    _ = v20
                }
            }
            var v21 int32 = 344
            _ = fi12
            _ = v21
        }
        delete(v3, 2147483647)
        if ((((127 > 1) || true) && true) && true) {
            for fi22 := 0; fi22 < 9; fi22++ {
                v23 := GenS1{F0: (v6 / ((0) | 1))}
                if ((1 > 326) || (2 <= (v5 >> ((558) & 63)))) {
                    var v24 int8 = int8(fi22)
                    fmt.Println(v24, v5, v24)
                    var v25 int16 = 1000
                    _ = v24
                    _ = v25
                }
                _ = len(v3)
                _ = fi22
                _ = v23
            }
            if false {
                v5 -= (^(v5))
                v26 := GenS1{F0: (v6 * (^(v6)))}
                var v27 int16 = 32767
                _ = v26
                _ = v27
            } else {
                v5 += (v5 | (v5 | v5))
                fmt.Println(v6, v5)
                var v28 float32 = 0.1
                _ = v28
            }
        }
    }
    switch -9223372036854775808 {
    case 127:
        var v29 uint = 65535
        _ = v29
    }
    v5 += v5
    if false {
        v30 := GenS2{F0: 1000}
        var v31 uint = 784
        _ = v30
        _ = v31
    }
    v6 -= 246
    switch ("hello" + "foo bar") {
    case "\n\t":
        if (((-2.25 / 1e10) / ((-2.25 + ((1e10 * ((1.5 / ((0.1 + 1e10) / 1.5)) / (-2.25 * 3.14159))) * (float32(float64(-2.25)) - (float32(1e10) / 1e10)))) - (1e10 / (0.0 - 3.14159)))) <= 3.14159) {
            fmt.Println(v6, v6)
        } else {
            for fi32 := 0; fi32 < 3; fi32++ {
                v6 = v6
                var v33 string = ("foo bar" + "\n\t")
                fmt.Println(v6, v33, v5)
                _ = fi32
                _ = v33
            }
        }
    case "foo bar":
        if true {
            v34 := GenS3{F0: (^(v6)), F1: 42, F2: "foo bar"}
            _ = v34
        }
    default:
        v5 += v5
    }
    var v35 float64 = (float64(0.1) - 0.0)
    v3[2147483647] = (42 <= int32(v5))
    var v36 uint8 = 160
    _ = len(v2)
    v35 -= float64(0.1)
    var v37 uint8 = (v36 & v36)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
