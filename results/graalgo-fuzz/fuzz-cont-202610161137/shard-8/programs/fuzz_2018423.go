package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
func genFn0(p0 uint, p1 string) int16 {
    if (42 >= 33) {
        if (182 > 42) {
            for fi7 := 0; fi7 < 10; fi7++ {
                p1 = p1
                var v8 uint = p0
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 12; fi9++ {
                var v10 int64 = 956
                var v11 int8 = -72
                _ = fi9
                _ = v10
                _ = v11
            }
        } else {
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 int16 = -446
                _ = fi12
                _ = v13
            }
            var v14 int8 = 42
            fmt.Println(p1, v14)
            _ = v14
        }
        for fi15 := 0; fi15 < 6; fi15++ {
            p1 = p1
            v16 := []int{}
            if true {
                var v17 uint = 18446744073709551615
                _ = v17
            } else {
                fmt.Println(fi15, p0, p1)
            }
            _ = fi15
            _ = v16
        }
    } else {
        var v18 string = p1
        _ = v18
    }
    p1 = p1
    var v19 float32 = (0.1 * 0.0)
    if true {
        v20 := GenS1{F0: (p0 / (((p0 | (p0 - 18446744073709551615))) | 1)), F1: 2147483647}
        if (18446744073709551615 >= 4294967295) {
            v21 := GenS1{F0: p0, F1: 2147483648}
            _ = v21
        }
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v19
    return 1000
}
func genFn1(p0 int8, p1 string) uint32 {
    var v22 float64 = (0.0 + float64(3.14159))
    var v23 uint16 = 668
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return 155
}

func main() {
    v1 := []uint32{2}
    var v2 uint32 = 1
    v3 := make(map[int]string)
    var v4 uint8 = 15
    v5 := []string{"\n\t", "\n\t"}
    v6 := []int{}
    var v24 int32 = 69
    var v25 float64 = (0.1 * float64(float32(0.1)))
    v26 := GenS1{F0: 255, F1: v2}
    for fi27 := 0; fi27 < 20; fi27++ {
        for fi28 := 0; fi28 < 15; fi28++ {
            if false {
                var v29 string = "hello"
                v29 = "a"
                _ = v29
            }
            _ = fi28
        }
        var v30 float64 = v25
        v31 := v3[127]
        _ = fi27
        _ = v30
        _ = v31
    }
    var v32 int = 324
    var v33 int8 = 1
    v34 := GenS1{F0: 4294967295, F1: v2}
    fmt.Println(v32)
    fmt.Println(v33, v4, v4)
    var v35 uint = 18446744073709551615
    var v36 int64 = -123
    if ("foo bar" < "foo bar") {
        _ = len(v3)
        v33 = v33
        _ = len(v3)
    } else {
        var v37 int8 = -128
        _ = v37
    }
    switch v32 {
    case -9223372036854775808:
        switch ("a" + ("foo bar" + "a")) {
        case "hello":
            var v38 uint32 = v2
            _ = v38
        case "foo bar":
            switch ("foo bar" + ("\n\t" + "hello")) {
            case "a":
                v6 = append(v6, v32)
            case "\n\t":
                var v39 uint16 = 94
                _ = v39
            default:
                for fi40 := 0; fi40 < 11; fi40++ {
                    v41 := GenS1{F0: (v35 - v35), F1: uint32(v35)}
                    _ = fi40
                    _ = v41
                }
            }
        default:
            var v42 bool = (v33 <= (v33 % (((v33 + -1)) | 1)))
            _ = v42
        }
    default:
        v24 += (-(v24))
    }
    switch "\n\t" {
    case "foo bar":
        v4 = (v4 ^ (v4 >> ((v35) & 63)))
    }
    var v43 float64 = v25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
