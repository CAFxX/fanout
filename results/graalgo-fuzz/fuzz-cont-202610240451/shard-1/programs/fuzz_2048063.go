package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint16; F1 bool; F2 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 int32 = -649
    var v2 int32 = v1
    v3 := make(map[string]string)
    v4 := make(map[uint32]string)
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int8 = 127
        v7 := GenS1{F0: ("foo bar" + ("a" + "a"))}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{127}
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            v2 = v2
            if true {
                v11 := GenS1{F0: "foo bar"}
                _ = v11
            }
            var v12 float64 = -2.25
            _ = fi10
            _ = v12
        }
        if false {
            var v13 float32 = ((float32(3.14159) + float32(1.5)) * 3.14159)
            v13 -= float32((((1.5 - 1.5) / (1e10 - 1.5)) + 0.0))
            _ = v13
        } else {
            var v14 float64 = -2.25
            v1 = (v2 * (v2 >> ((18446744073709551615) & 63)))
            v15 := GenS1{F0: ("foo bar" + "\n\t")}
            _ = v14
            _ = v15
        }
        switch "foo bar" {
        case "hello":
            var v16 float64 = 1e10
            _ = v16
        case "\n\t":
            if true {
                v4[2147483648] = "foo bar"
                _ = len(v3)
            }
        default:
            v17 := GenS2{F0: 42, F1: (false && (1.5 >= (-2.25 + float32((1.5 - float64(1e10))))))}
            _ = v17
        }
        _ = fi9
    }
    fmt.Println(v1, v2)
    v8 = append(v8, -9223372036854775808)
    v18 := GenS1{F0: "\n\t"}
    fmt.Println(v2)
    for fi19 := 0; fi19 < 6; fi19++ {
        v20 := GenS3{F0: 0, F1: true, F2: (fi19 ^ fi19)}
        v21 := GenS4{F0: 1}
        if true {
            for fi22 := 0; fi22 < 17; fi22++ {
                v23 := GenS5{F0: 1}
                _ = fi22
                _ = v23
            }
            fmt.Println(v1, fi19)
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    var v24 uint = 2
    if true {
        var v25 uint8 = 2
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v8
}
