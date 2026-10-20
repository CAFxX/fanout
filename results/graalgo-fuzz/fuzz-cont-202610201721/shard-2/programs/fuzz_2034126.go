package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 int16; F2 bool }
type GenS4 struct { F0 uint }

func main() {
    var v1 bool = false
    var v2 float64 = (-2.25 - 0.0)
    var v3 uint64 = 18446744073709551615
    var v4 int32 = 292
    v5 := []int{}
    var v6 float32 = 1e10
    v2 += (0.1 * v2)
    if (97 <= -9223372036854775808) {
        if v1 {
            if v1 {
                var v7 float32 = v6
                v8 := GenS1{F0: "", F1: "", F2: 18446744073709551615}
                _ = v7
                _ = v8
            }
            switch 2147483647 {
            case 32768:
                v6 = v6
            case -2147483648:
                var v9 uint = 255
                _ = v9
            case 353:
                v4 += v4
            }
            if false {
                var v10 uint8 = 255
                _ = v10
            }
        } else {
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 bool = v1
                if len(v5) > 0 {
                    v5[0] = 32768
                }
                _ = fi11
                _ = v12
            }
        }
        var v13 uint8 = 11
        _ = v13
    } else {
        switch "foo bar" {
        case "hello":
            for fi14 := 0; fi14 < 19; fi14++ {
                v6 += (v6 + v6)
                v15 := GenS2{F0: -1}
                _ = fi14
                _ = v15
            }
        case "a":
            var v16 uint8 = 42
            _ = v16
        case "\n\t":
            v17 := GenS3{F0: 255, F1: 750, F2: false}
            _ = v17
        default:
            v18 := GenS3{F0: 255, F1: -976, F2: v1}
            _ = v18
        }
        v4 -= v4
    }
    var v19 int = 268
    v3 -= (^(v3))
    var v20 uint32 = 2147483647
    if ((v4 + 226) != v4) {
        for fi21 := 0; fi21 < 7; fi21++ {
            v3 = v3
            _ = fi21
        }
        if ("foo bar" == ("hello" + "")) {
            v6 += float32(0.1)
            v2 += (v2 * (v2 + (float64((v6 + (-2.25 / ((float32(v2) / (float32(v2) * (float32(0.0) - ((v6 - 0.1) - v6)))) - v6)))) / v2)))
            var v22 string = ("a" + "foo bar")
            _ = v22
        } else {
            _ = len(v5)
        }
        v23 := GenS4{F0: uint(v3)}
        _ = v23
    }
    var v24 string = "\n\t"
    switch "hello" {
    case "foo bar":
        var v25 int32 = 2147483647
        _ = v25
    case "\n\t":
        v26 := GenS2{F0: -697}
        _ = v26
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
