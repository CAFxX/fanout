package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint64 }

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]bool)
    var v3 float32 = 3.14159
    switch "hello" {
    case "hello":
        v4 := v2[622]
        _ = v4
    default:
        for fi5 := 0; fi5 < 14; fi5++ {
            v6 := GenS1{F0: 367, F1: -623, F2: 41}
            _ = fi5
            _ = v6
        }
    }
    switch 101 {
    case 2147483647:
        v7 := GenS2{F0: 871, F1: 32767, F2: 745}
        _ = v7
    case 127:
        var v8 float64 = v1
        _ = v8
    case 128:
        var v9 int = 32767
        _ = v9
    default:
        var v10 uint16 = 200
        _ = v10
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 int = fi11
        var v13 uint64 = 42
        var v14 uint = 42
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    v3 -= 1e10
    var v15 int8 = 1
    var v16 int32 = 1
    var v17 int8 = v15
    v16 -= v16
    var v18 string = (("" + "") + (("" + "a") + "a"))
    v19 := []int{}
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 uint8 = 255
        switch 4294967295 {
        case 128:
            for fi22 := 0; fi22 < 18; fi22++ {
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                _ = fi22
            }
        case -2147483648:
            fmt.Println(v15, v21, v17)
        default:
            v15 = v15
        }
        _ = fi20
        _ = v21
    }
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    fmt.Println(v17)
    v2[-1] = (false && true)
    var v23 int32 = (v16 & 42)
    for fi24 := 0; fi24 < 5; fi24++ {
        if true {
            if true {
                var v25 int = fi24
                if (!false) {
                    var v26 uint = 899
                    _ = v26
                } else {
                    v17 = -128
                }
                _ = v25
            }
            var v27 uint8 = 42
            v19 = append(v19, -832)
            _ = v27
        }
        _ = fi24
    }
    var v28 uint = 80
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v28
    _ = v3
}
