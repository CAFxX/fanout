package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS3 struct { F0 string; F1 uint8; F2 uint }

func main() {
    var v1 int64 = -970
    var v2 int = 9223372036854775807
    v3 := make(map[int]uint32)
    fmt.Println(v2, v2)
    v2 = 4294967295
    var v4 int64 = 0
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 bool = (42 > (v2 & fi5))
        v7 := []int{(-(fi5)), (fi5 * (-(v2))), (fi5 << ((406) & 63)), (-(fi5))}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: true, F1: 127}
    switch "hello" {
    case "foo bar":
        v4 += -9223372036854775808
    case "hello":
        v2 = (v2 % (((v2 ^ v2)) | 1))
    case "\n\t":
        v4 -= v1
    }
    v9 := GenS2{F0: 255, F1: 1, F2: 2}
    v10 := []int{(v2 << ((255) & 63)), (v2 << ((975) & 63)), 0, v2}
    v2 = v2
    switch ("foo bar" + "") {
    case "":
        v4 += v1
    default:
        v11 := GenS3{F0: "\n\t", F1: 2, F2: 862}
        _ = v11
    }
    v1 = (v4 + v4)
    for fi12 := 0; fi12 < 11; fi12++ {
        if true {
            v13 := GenS3{F0: ("" + "\n\t"), F1: 0, F2: 18446744073709551615}
            v14 := GenS3{F0: "foo bar", F1: 255, F2: 18446744073709551615}
            _ = v13
            _ = v14
        }
        v2 += v2
        var v15 float64 = (1e10 + 3.14159)
        _ = fi12
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        v17 := GenS2{F0: 42, F1: 42, F2: 821}
        _ = fi16
        _ = v17
    }
    v4 = v4
    v1 = (-(v4))
    var v18 bool = false
    for fi19 := 0; fi19 < 1; fi19++ {
        v18 = v18
        v4 += 1
        _ = fi19
    }
    switch "hello" {
    case "hello":
        v20 := GenS2{F0: 127, F1: 1, F2: 312}
        _ = v20
    case "\n\t":
        var v21 uint32 = 687
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
