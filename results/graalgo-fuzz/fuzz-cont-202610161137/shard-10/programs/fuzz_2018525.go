package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 bool; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint16; F1 string; F2 uint8 }
type GenS4 struct { F0 uint; F1 bool; F2 int16 }
func genFn0(p0 int8, p1 uint8) int64 {
    var v6 float64 = 1.5
    fmt.Println(v6, p1, p0)
    _ = p0
    _ = p1
    _ = v6
    return 9223372036854775807
}

func main() {
    v1 := []uint32{}
    var v2 string = "foo bar"
    var v3 float64 = 0.1
    var v4 uint = 31
    v5 := make(map[string]uint32)
    switch "foo bar" {
    case "foo bar":
        var v7 int = -86
        _ = v7
    default:
        fmt.Println(v4)
    }
    v8 := GenS1{F0: -91, F1: 974}
    for fi9 := 0; fi9 < 5; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := GenS1{F0: 42, F1: -541}
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    var v12 uint64 = 2
    if true {
        var v13 uint = v4
        _ = v13
    }
    v5[(v2 + v2)] = 65535
    var v14 uint32 = 1
    v12 = (v12 - (v12 + (^(v12))))
    v15 := GenS1{F0: 10, F1: -1}
    v3 = float64((float32(v3) + -2.25))
    v2 = ("hello" + v2)
    var v16 uint16 = 339
    v17 := GenS2{F0: (v2 == ("a" + "foo bar")), F1: false, F2: 4294967295}
    var v18 bool = (genFn0(0, 2) == -465)
    var v19 float64 = v3
    v20 := GenS3{F0: (v16 & v16), F1: (((v2 + v2) + v2) + v2), F2: 255}
    switch 127 {
    case 2147483647:
        fmt.Println(v19)
    case 1:
        v21 := GenS4{F0: v4, F1: v18, F2: -32768}
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
