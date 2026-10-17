package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 bool }
func genFn0(p0 string, p1 bool) uint64 {
    p1 = p1
    v4 := GenS1{F0: 42, F1: 18446744073709551615}
    if true {
        var v5 float32 = 1.5
        _ = v5
    }
    _ = p0
    _ = p1
    _ = v4
    return 0
}

func main() {
    var v1 uint = 2
    var v2 uint16 = 65535
    var v3 float64 = -2.25
    var v6 int32 = -414
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint8 = 159
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := []int{}
        _ = fi9
        _ = v10
    }
    var v11 int16 = -1
    switch -595 {
    case 2147483647:
        v2 += (^(v2))
    default:
        switch 4294967295 {
        case -848:
            for fi12 := 0; fi12 < 7; fi12++ {
                var v13 uint = (v1 % (((v1 * v1)) | 1))
                _ = fi12
                _ = v13
            }
        case 441:
            var v14 float32 = 1e10
            _ = v14
        default:
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 int64 = -1
                var v17 float32 = ((((1.5 + float32(float64(0.0))) / float32((0.1 * 0.1))) * ((0.0 / (1e10 + 1.5)) - 1.5)) - ((float32(1e10) * ((0.1 * (0.0 / 1.0)) / 1.0)) * float32(3.14159)))
                var v18 float64 = v3
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        }
    }
    if true {
        var v19 string = ("a" + "foo bar")
        v20 := GenS2{F0: (v1 % (((v1 % ((v1) | 1))) | 1)), F1: -84}
        switch 42 {
        case -952:
            fmt.Println(v19)
        case 2147483647:
            v21 := GenS1{F0: 13, F1: 2}
            _ = v21
        case -898:
            v2 = v2
        }
        _ = v19
        _ = v20
    }
    switch (("a" + "hello") + "a") {
    case "\n\t":
        var v22 int64 = 781
        _ = v22
    case "":
        v23 := []int{517}
        _ = v23
    default:
        if true {
            for fi24 := 0; fi24 < 3; fi24++ {
                v2 = (v2 | v2)
                _ = fi24
            }
            var v25 bool = true
            _ = v25
        }
    }
    v26 := []int{-533}
    v27 := GenS3{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v6
}
