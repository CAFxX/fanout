package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 string; F1 uint16; F2 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 100
    var v3 int16 = 32767
    var v4 string = ("a" + "hello")
    if true {
        switch -1 {
        case -9223372036854775808:
            switch "\n\t" {
            case "":
                var v5 int8 = -1
                _ = v5
            case "\n\t":
                v2 = v2
            case "hello":
                v6 := GenS1{F0: "", F1: 127}
                _ = v6
            }
        case 127:
            v7 := []int{4294967295, 9223372036854775807}
            _ = v7
        case 32767:
            fmt.Println(v4)
        default:
            v3 -= (v3 - v3)
        }
        var v8 int32 = -374
        _ = v8
    }
    if false {
        var v9 uint32 = 4294967295
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(v4, v2)
        var v11 int = (fi10 + (fi10 | (fi10 - fi10)))
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        _ = len(v1)
        v2 = -128
        for fi13 := 0; fi13 < 5; fi13++ {
            v14 := GenS2{F0: v4, F1: 0, F2: 2147483647}
            var v15 int32 = -621
            fmt.Println(v15, fi12)
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    var v16 int = 127
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 uint16 = 0
        var v19 int8 = (v2 & (v2 / ((-128) | 1)))
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println(v2)
    fmt.Println(v2, v16, v16)
    var v20 uint = 255
    if (((0.1 * 1e10) / ((-2.25 + 0.1) - float64(-2.25))) >= 1e10) {
        var v21 float32 = 1e10
        v22 := v1[v4]
        _ = v21
        _ = v22
    }
    fmt.Println(v2)
    v23 := GenS2{F0: v4, F1: 83, F2: 377}
    var v24 uint8 = 2
    v25 := v1[("\n\t" + v4)]
    var v26 uint = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
