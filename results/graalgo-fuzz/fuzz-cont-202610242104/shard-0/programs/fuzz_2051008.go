package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int16; F2 string }
type GenS3 struct { F0 uint }
func genFn0(p0 bool, p1 bool) int64 {
    var v7 int = 856
    v8 := []int{}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = -2.25
    var v3 uint = 18446744073709551615
    var v4 float32 = float32(v2)
    var v5 int16 = 0
    v6 := make(map[string]int)
    if (true && true) {
        v9 := GenS1{F0: 150}
        var v10 int = 127
        fmt.Println(v10, v2, v5)
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 int16 = 1
            v4 = (((((v4 + v4) / ((v4 + (1e10 * v4)) * (v4 * float32((1.5 - v2))))) - (v4 / v4)) + float32((v2 / float64((((0.1 - v4) + v4) / float32(((0.1 / ((v2 - 0.0) * v2)) - v2))))))) + float32(0.0))
            switch fi11 {
            case -1:
                var v13 int = (fi11 / (((fi11 ^ -813)) | 1))
                _ = v13
            default:
                v14 := []int{}
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        var v15 float64 = ((1.5 + v2) - float64(v4))
        _ = v15
    }
    v3 -= v3
    for fi16 := 0; fi16 < 20; fi16++ {
        if true {
            var v17 uint64 = 2
            _ = v17
        }
        var v18 float32 = v4
        _ = fi16
        _ = v18
    }
    switch ("a" + ("" + ("hello" + ""))) {
    case "\n\t":
        for fi19 := 0; fi19 < 10; fi19++ {
            fmt.Println(fi19, v3)
            _ = fi19
        }
    case "":
        v5 += v5
    case "foo bar":
        var v20 uint = 110
        _ = v20
    }
    v5 = (v5 | -794)
    v21 := v6[("hello" + "\n\t")]
    v21 = v21
    v5 = (v5 % ((416) | 1))
    for fi22 := 0; fi22 < 12; fi22++ {
        if (127 <= 100) {
            var v23 int64 = 9223372036854775807
            _ = v23
        } else {
            var v24 bool = (42 > v5)
            var v25 int = (-(v21))
            switch ("hello" + "\n\t") {
            case "foo bar":
                var v26 uint8 = 77
                _ = v26
            case "\n\t":
                var v27 uint16 = 1
                _ = v27
            case "hello":
                fmt.Println(v3, fi22)
            }
            _ = v24
            _ = v25
        }
        v28 := GenS2{F0: fi22, F1: (v5 ^ 42), F2: "foo bar"}
        _ = fi22
        _ = v28
    }
    var v29 int = 0
    v4 -= float32(((float64(-2.25) + v2) + ((((1e10 - 3.14159) * -2.25) + v2) / (1e10 / 1e10))))
    v1[65535] = true
    for fi30 := 0; fi30 < 10; fi30++ {
        var v31 float32 = v4
        switch "\n\t" {
        case "":
            switch "foo bar" {
            case "\n\t":
                switch fi30 {
                case 1:
                    var v32 int16 = v5
                    _ = v32
                case -9223372036854775808:
                    v2 = (float64(v4) * (float64(v4) + v2))
                case 849:
                    v3 += v3
                }
            case "":
                v29 = v21
            }
        case "a":
            v33 := GenS3{F0: v3}
            _ = v33
        }
        _ = fi30
        _ = v31
    }
    v34 := []int{0, v21}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
