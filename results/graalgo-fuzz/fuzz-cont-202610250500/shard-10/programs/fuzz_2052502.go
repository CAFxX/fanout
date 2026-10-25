package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint = 441
    v2 := make(map[int]string)
    v3 := make(map[uint32]int)
    var v4 bool = (2 <= 105)
    v5 := make(map[uint32]bool)
    var v6 int64 = 9223372036854775807
    v7 := GenS1{F0: -128}
    switch 0 {
    case -2147483648:
        var v8 float32 = 1e10
        _ = v8
    case -807:
        fmt.Println(v1, v1)
    case 1:
        for fi9 := 0; fi9 < 12; fi9++ {
            fmt.Println(v4, v6, v6)
            if (1e10 < (-2.25 + 0.1)) {
                for fi10 := 0; fi10 < 14; fi10++ {
                    fmt.Println(v1, fi10)
                    _ = fi10
                }
                var v11 bool = true
                v6 = (v6 >> ((v1) & 63))
                _ = v11
            }
            _ = fi9
        }
    default:
        var v12 uint64 = 90
        _ = v12
    }
    if v4 {
        var v13 string = "hello"
        _ = v13
    }
    fmt.Println(v6, v6, v1)
    for fi14 := 0; fi14 < 3; fi14++ {
        v4 = ((1.5 * ((1e10 * -2.25) - ((0.1 * ((float32(((((1.5 - float64(0.0)) - 1.5) / 1.5) - 1.5)) + 1e10) * 1e10)) - 0.1))) != 1.5)
        _ = len(v5)
        _ = fi14
    }
    _ = len(v2)
    if (32767 > -957) {
        var v15 uint16 = 359
        v4 = v4
        v16 := []int{876}
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 4; fi17++ {
            var v18 uint8 = 42
            var v19 int8 = 0
            _ = fi17
            _ = v18
            _ = v19
        }
        v1 = v1
    }
    v4 = v4
    fmt.Println(v6, v4)
    v20 := []int{}
    switch -179 {
    case 603:
        fmt.Println(v1, v4)
    case -438:
        var v21 uint16 = 65535
        _ = v21
    }
    var v22 bool = v4
    v23 := GenS1{F0: 0}
    for fi24 := 0; fi24 < 16; fi24++ {
        fmt.Println(v4, v1, v1)
        _ = fi24
    }
    var v25 uint32 = 906
    if len(v20) > 0 {
        v20[0] = 32768
    }
    if ((true || (!true)) && (v22 && ((!(v4 && v22)) && (!(v25 == v25))))) {
        var v26 uint16 = 32767
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
