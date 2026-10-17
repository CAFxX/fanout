package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 uint16; F2 string }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint16; F1 int; F2 uint8 }
func genFn0(p0 int16) uint32 {
    var v4 uint16 = 0
    var v5 float32 = 0.0
    v6 := []int{-9223372036854775808}
    if len(v6) > 0 {
        v6[0] = -502
    }
    v7 := GenS1{F0: (!(42 > 42))}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 int32 = 774
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    var v8 bool = true
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 float64 = ((1e10 / 1e10) / 1.0)
        _ = fi9
        _ = v10
    }
    var v11 int32 = (v1 ^ (v1 & (v1 & v1)))
    if (0 <= 2) {
        for fi12 := 0; fi12 < 11; fi12++ {
            for fi13 := 0; fi13 < 19; fi13++ {
                v14 := GenS1{F0: v8}
                v15 := GenS2{F0: (fi12 % ((fi13) | 1)), F1: 65535, F2: ""}
                _ = fi13
                _ = v14
                _ = v15
            }
            v16 := GenS3{F0: 71, F1: (fi12 % (((fi12 & 9223372036854775807)) | 1))}
            _ = fi12
            _ = v16
        }
        switch "\n\t" {
        case "\n\t":
            var v17 uint8 = 196
            _ = v17
        default:
            v8 = true
        }
        fmt.Println(v8)
    } else {
        var v18 bool = v8
        _ = v18
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        v11 = 0
        switch "hello" {
        case "foo bar":
            v20 := []int{}
            _ = v20
        case "hello":
            _ = len(v2)
        case "":
            _ = len(v2)
        default:
            v21 := []int{(fi19 + fi19), fi19, fi19, fi19}
            _ = v21
        }
        v11 = (v1 + (v1 & 0))
        _ = fi19
    }
    _ = len(v2)
    v1 = v11
    var v22 bool = v8
    if true {
        var v23 uint64 = 1
        var v24 uint8 = 255
        _ = v23
        _ = v24
    } else {
        switch ((("foo bar" + "\n\t") + "a") + "") {
        case "":
            fmt.Println(v22, v8)
        }
    }
    var v25 float32 = 0.0
    v26 := GenS2{F0: 32768, F1: 391, F2: "foo bar"}
    v22 = true
    v2[-35] = -9223372036854775808
    var v27 float32 = (((float32(0.0) * (0.0 + float32(((float64(float32((-2.25 / 0.1))) + 1e10) + 1e10)))) * ((float32(3.14159) / (float32(1.5) * (v25 + (1.5 * v25)))) * ((0.1 / v25) - (v25 + v25)))) - ((float32(1.5) * 1.5) * (float32(3.14159) - (((v25 - 1e10) * ((v25 - v25) + (v25 * (3.14159 + float32(1e10))))) / v25))))
    if true {
        fmt.Println(v27, v8, v11)
        fmt.Println(v1)
        v28 := GenS4{F0: 42, F1: 127, F2: 60}
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v8
}
