package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int32) int16 {
    fmt.Println(p0)
    var v5 int16 = 42
    for fi6 := 0; fi6 < 11; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := []int{9223372036854775807}
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi7
            _ = v8
        }
        var v9 int64 = -1
        p0 += p0
        _ = fi6
        _ = v9
    }
    _ = p0
    _ = v5
    return v5
}
func genFn1(p0 uint16) bool {
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 uint32 = 404
        _ = fi10
        _ = v11
    }
    fmt.Println(p0)
    for fi12 := 0; fi12 < 2; fi12++ {
        fmt.Println(p0, fi12)
        var v13 uint8 = 128
        _ = fi12
        _ = v13
    }
    p0 += 255
    _ = p0
    return (((!(!false)) && (82 >= 0)) && true)
}

func main() {
    var v1 uint8 = 1
    var v2 int8 = 101
    var v3 int64 = -804
    var v4 uint16 = 255
    switch 1 {
    case -1:
        var v14 float32 = ((0.0 - (((3.14159 * 0.1) * 0.1) / 1.0)) / 0.1)
        _ = v14
    default:
        switch "a" {
        case "hello":
            v15 := GenS1{F0: 962}
            _ = v15
        }
    }
    v16 := GenS2{F0: 255, F1: (!false)}
    var v17 int = 127
    v3 += 82
    var v18 float32 = 1.5
    var v19 bool = false
    var v20 uint8 = 42
    switch v17 {
    case 32768:
        if true {
            fmt.Println(v4, v19)
        }
    }
    v3 += (v3 & -1)
    var v21 int32 = 42
    var v22 float64 = float64((v18 * v18))
    for fi23 := 0; fi23 < 6; fi23++ {
        for fi24 := 0; fi24 < 17; fi24++ {
            v25 := GenS1{F0: v17}
            var v26 int = fi23
            _ = fi24
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 10; fi27++ {
            v28 := GenS3{F0: v21, F1: 0}
            var v29 int64 = int64(v4)
            _ = fi27
            _ = v28
            _ = v29
        }
        for fi30 := 0; fi30 < 14; fi30++ {
            var v31 uint8 = 2
            _ = fi30
            _ = v31
        }
        _ = fi23
    }
    fmt.Println(v1)
    v32 := GenS4{F0: 240}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v4
}
