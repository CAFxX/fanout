package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int32; F2 int }
type GenS4 struct { F0 uint64; F1 uint32; F2 string }
type GenS5 struct { F0 int }
func genFn0(p0 uint) int8 {
    var v7 int64 = 0
    v8 := GenS1{F0: p0, F1: -1}
    switch 128 {
    case -9223372036854775808:
        fmt.Println(v7, p0)
    }
    var v9 int8 = 104
    if (0 <= 18) {
        if true {
            if (-1 <= 716) {
                v10 := []int{128, 1, -1}
                var v11 string = ""
                if len(v10) > 0 {
                    v10[0] = -9223372036854775808
                }
                _ = v10
                _ = v11
            }
            var v12 uint = (p0 / (((p0 & (p0 / (((^(p0))) | 1)))) | 1))
            _ = v12
        } else {
            var v13 int8 = -1
            switch "a" {
            case "hello":
                var v14 uint32 = 0
                _ = v14
            case "foo bar":
                var v15 int16 = -32768
                _ = v15
            default:
                p0 = p0
            }
            for fi16 := 0; fi16 < 15; fi16++ {
                v7 -= v7
                _ = fi16
            }
            _ = v13
        }
    } else {
        var v17 bool = (float32(3.14159) == 1.5)
        v18 := GenS1{F0: (p0 % ((785) | 1)), F1: -716}
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 uint8 = 2
            fmt.Println(v17, v20, p0)
            _ = fi19
            _ = v20
        }
        _ = v17
        _ = v18
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (v9 | (-(v9)))
}
func genFn1(p0 int) int32 {
    p0 -= 0
    p0 += p0
    _ = p0
    return -2147483648
}

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = float64(0.0)
    v3 := make(map[int]string)
    var v4 uint16 = 42
    v5 := []uint32{120, 0, 909, 42}
    var v6 uint = 255
    v21 := GenS1{F0: v6, F1: 32768}
    for fi22 := 0; fi22 < 15; fi22++ {
        if ((!true) && false) {
            v2 -= ((float64(1.5) + v2) + 0.0)
        }
        fmt.Println(v6, v6)
        _ = fi22
    }
    v23 := GenS2{F0: true}
    if true {
        for fi24 := 0; fi24 < 14; fi24++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v25 := GenS3{F0: fi24, F1: -2147483648, F2: fi24}
            _ = fi24
            _ = v25
        }
        _ = len(v3)
    }
    _ = len(v3)
    v5 = append(v5, 42)
    switch (("\n\t" + "\n\t") + "") {
    case "":
        v26 := GenS3{F0: 32768, F1: 168, F2: 128}
        _ = v26
    case "a":
        v27 := GenS4{F0: 1, F1: 2, F2: "foo bar"}
        _ = v27
    }
    var v28 uint = 51
    var v29 int = 32768
    for fi30 := 0; fi30 < 17; fi30++ {
        var v31 float64 = ((v2 - (((v2 * float64((float32(1e10) - ((1.5 + -2.25) / ((0.0 - (float32(float64(0.1)) * 1.5)) * 1e10))))) / ((((1.5 / 1.0) + v2) + v2) / 1.5)) / (v2 / float64((((1.5 - (1.5 + 3.14159)) + 1e10) - 0.1))))) * v2)
        if true {
            v31 += v2
            var v32 uint8 = 128
            v33 := GenS5{F0: (fi30 - (-(v29)))}
            _ = v32
            _ = v33
        }
        _ = fi30
        _ = v31
    }
    var v34 uint = v28
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
