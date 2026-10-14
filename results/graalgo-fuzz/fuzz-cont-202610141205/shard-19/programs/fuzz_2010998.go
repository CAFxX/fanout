package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint32; F1 int8 }

func main() {
    var v1 int = -851
    v2 := []string{}
    var v3 int8 = 127
    v4 := make(map[uint32]int)
    for fi5 := 0; fi5 < 17; fi5++ {
        v4[418] = -9223372036854775808
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 uint8 = 1
            v7 += v7
            var v8 int32 = 42
            switch fi5 {
            case 9223372036854775807:
                var v9 uint = 65535
                _ = v9
            case 127:
                fmt.Println(v7, v3, fi5)
            case 42:
                var v10 int8 = v3
                _ = v10
            default:
                v11 := GenS1{F0: 4294967295, F1: 42}
                _ = v11
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        v1 = 170
        _ = fi5
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 float64 = float64(0.0)
        var v14 int = fi12
        if (0 > 2) {
            v15 := GenS1{F0: 2147483648, F1: 2147483647}
            _ = v15
        } else {
            var v16 float32 = (3.14159 - 1e10)
            var v17 uint = 61
            switch "a" {
            case "hello":
                v18 := GenS2{F0: 1, F1: 89}
                _ = v18
            case "":
                var v19 float64 = -2.25
                _ = v19
            }
            _ = v16
            _ = v17
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = len(v4)
    switch ("hello" + ("foo bar" + (("foo bar" + "hello") + ("a" + "a")))) {
    case "foo bar":
        fmt.Println(v1, v1)
    case "":
        var v20 uint8 = 1
        _ = v20
    default:
        v3 += 42
    }
    var v21 int16 = 1000
    v21 = (v21 % ((v21) | 1))
    if true {
        for fi22 := 0; fi22 < 5; fi22++ {
            v2 = append(v2, ("foo bar" + "hello"))
            v23 := GenS1{F0: 4294967295, F1: -2147483648}
            var v24 bool = true
            _ = fi22
            _ = v23
            _ = v24
        }
        v3 += (^(v3))
    } else {
        v2 = append(v2, ("" + "foo bar"))
        v25 := GenS1{F0: 65535, F1: -1}
        var v26 int8 = v3
        _ = v25
        _ = v26
    }
    if (!(false && false)) {
        if false {
            v27 := GenS3{F0: 2, F1: -128}
            var v28 uint = 329
            _ = v27
            _ = v28
        } else {
            v29 := GenS3{F0: 4294967295, F1: (v3 - (v3 / ((-128) | 1)))}
            v3 = (v3 ^ (v3 >> ((18446744073709551615) & 63)))
            v21 += v21
            _ = v29
        }
        for fi30 := 0; fi30 < 20; fi30++ {
            v31 := GenS2{F0: (v3 + v3), F1: 1}
            var v32 int8 = v3
            v1 = (v1 << ((1) & 63))
            _ = fi30
            _ = v31
            _ = v32
        }
    }
    var v33 int16 = (v21 & (-(v21)))
    switch "" {
    case "hello":
        if ((1 >= (v3 % ((100) | 1))) || false) {
            v21 += 1000
        } else {
            v21 = v33
            v21 -= v33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v33
    _ = v4
}
