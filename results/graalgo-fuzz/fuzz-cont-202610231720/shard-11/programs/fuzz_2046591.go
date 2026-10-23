package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "a"
    var v3 bool = false
    var v4 int = 2147483647
    if (1000 <= 0) {
        v4 = v4
        v5 := GenS1{F0: 912, F1: 65535}
        _ = v5
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint64 = 18446744073709551615
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v9 := GenS2{F0: -1}
    fmt.Println(v3, v2)
    var v10 bool = v3
    fmt.Println(v2, v2)
    v11 := GenS1{F0: 0, F1: 908}
    switch "hello" {
    case "hello":
        var v12 uint16 = 42
        _ = v12
    case "\n\t":
        if (-883 < -2147483648) {
            v13 := []int{}
            var v14 uint8 = 1
            v15 := GenS1{F0: 2, F1: 248}
            _ = v13
            _ = v14
            _ = v15
        } else {
            v16 := GenS3{F0: 32767, F1: 255}
            v10 = true
            v2 = v2
            _ = v16
        }
    default:
        var v17 float32 = (1.5 - 3.14159)
        _ = v17
    }
    v18 := []int{32767, (-(v4)), 128}
    switch "hello" {
    case "hello":
        v2 = ((v2 + ((v2 + (v2 + ("foo bar" + ("a" + v2)))) + (v2 + (v2 + "\n\t")))) + ((((v2 + (v2 + v2)) + v2) + v2) + ""))
    case "":
        v19 := GenS4{F0: -2147483648}
        _ = v19
    default:
        var v20 int = (-(v4))
        _ = v20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        if v3 {
            for fi22 := 0; fi22 < 2; fi22++ {
                _ = len(v18)
                var v23 int16 = 460
                _ = fi22
                _ = v23
            }
            var v24 int32 = -1
            var v25 int8 = 0
            _ = v24
            _ = v25
        }
        _ = fi21
    }
    if (100 == -35) {
        switch v2 {
        case "":
            fmt.Println(v2, v4)
        case "a":
            fmt.Println(v10, v10)
        }
        v26 := GenS2{F0: 42}
        v27 := GenS3{F0: 42, F1: 127}
        _ = v26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
