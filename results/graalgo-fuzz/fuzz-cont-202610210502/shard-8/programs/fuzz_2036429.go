package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 string; F1 int64; F2 int8 }
type GenS4 struct { F0 int16 }
func genFn0() uint64 {
    if true {
        var v6 uint32 = 1
        var v7 uint16 = 2
        _ = v6
        _ = v7
    }
    switch "foo bar" {
    case "hello":
        var v8 string = ("foo bar" + "\n\t")
        _ = v8
    case "\n\t":
        var v9 int8 = 1
        _ = v9
    }
    var v10 float64 = (0.0 + 0.0)
    for fi11 := 0; fi11 < 2; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            switch (("a" + ("a" + "\n\t")) + (("a" + ("\n\t" + "")) + "a")) {
            case "hello":
                var v13 int32 = 2147483647
                fmt.Println(v10, fi12)
                if true {
                    v14 := GenS1{F0: -128, F1: v13, F2: 1}
                    _ = v14
                } else {
                    if (2 < 42) {
                        v10 = (((float64(-2.25) - v10) * ((v10 - (v10 + 1e10)) + v10)) / (v10 * (((float64((((3.14159 / (float32(0.0) + float32(3.14159))) * 1e10) - ((0.0 - 0.1) - (0.0 + (1e10 - (-2.25 / 1e10)))))) + float64(3.14159)) * 1e10) / v10)))
                        if (1000 > 42) {
                            var v15 uint64 = 86
                            var v16 uint16 = 42
                            _ = v15
                            _ = v16
                        }
                        var v17 uint8 = 0
                        _ = v17
                    } else {
                        var v18 int32 = 1
                        if (1.5 >= -2.25) {
                            switch ("" + "") {
                            case "hello":
                                var v19 string = "\n\t"
                                _ = v19
                            }
                            v20 := []int{2147483647}
                            _ = v20
                        }
                        v13 += 2147483647
                        _ = v18
                    }
                }
                var v21 bool = true
                v22 := GenS2{F0: 1, F1: -1}
                _ = v13
                _ = v21
                _ = v22
            case "a":
                var v23 float32 = (1.5 + (1e10 - (1e10 * 1.5)))
                _ = v23
            default:
                if (3.14159 == ((float64(3.14159) + (float64(-2.25) + v10)) - v10)) {
                    v24 := GenS2{F0: 33, F1: 9223372036854775807}
                    _ = v24
                }
            }
            _ = fi12
        }
        _ = fi11
    }
    _ = v10
    return 0
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]bool)
    var v3 bool = true
    var v4 uint32 = 2147483648
    v5 := make(map[uint32]int)
    v4 = 4294967295
    for fi25 := 0; fi25 < 2; fi25++ {
        v4 = 2
        _ = fi25
    }
    var v26 float64 = ((0.0 / 1e10) - (0.0 * (0.0 / ((0.1 - 3.14159) * 1.5))))
    for fi27 := 0; fi27 < 10; fi27++ {
        for fi28 := 0; fi28 < 20; fi28++ {
            fmt.Println(v4)
            _ = fi28
        }
        _ = fi27
    }
    v29 := []int{-2147483648, 727}
    var v30 float32 = ((((3.14159 / (float32(float64(-2.25)) * -2.25)) / 1.5) * float32(0.1)) - 0.0)
    v26 -= 3.14159
    _ = len(v5)
    v31 := GenS3{F0: v1, F1: 0, F2: 127}
    v30 += 0.0
    for fi32 := 0; fi32 < 17; fi32++ {
        v33 := GenS4{F0: 0}
        v3 = v3
        _ = fi32
        _ = v33
    }
    v29 = append(v29, 32768)
    v34 := GenS2{F0: 42, F1: -1}
    if (!(788 != 4294967295)) {
        v35 := GenS4{F0: 32767}
        _ = v35
    } else {
        fmt.Println(v26)
    }
    v26 = (v26 - v26)
    switch -9223372036854775808 {
    case 220:
        fmt.Println(v26, v3, v26)
    case -86:
        v30 = 3.14159
    }
    var v36 int8 = 42
    if (((!v3) && true) || (!((v36 % ((v36) | 1)) < v36))) {
        v26 -= (v26 * ((((v26 / v26) * (v26 * v26)) * v26) / 1.0))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v36
    _ = v4
    _ = v5
}
