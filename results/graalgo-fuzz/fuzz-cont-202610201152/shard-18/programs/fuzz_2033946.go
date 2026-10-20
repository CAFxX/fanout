package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint }
func genFn0(p0 uint8) int64 {
    var v5 uint8 = 185
    var v6 int16 = -32768
    v6 = (v6 - (v6 % ((42) | 1)))
    v7 := GenS1{F0: -19, F1: 255, F2: 470}
    var v8 int8 = 33
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 947
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 4294967295
    v3 := []uint32{}
    var v4 uint = 18446744073709551615
    v9 := GenS1{F0: -113, F1: 127, F2: 0}
    for fi10 := 0; fi10 < 18; fi10++ {
        switch -1 {
        case -933:
            var v11 int16 = 1
            _ = v11
        case 128:
            var v12 uint8 = 42
            _ = v12
        default:
            fmt.Println(fi10)
        }
        var v13 uint = 255
        _ = fi10
        _ = v13
    }
    var v14 uint64 = 42
    var v15 int16 = 966
    if true {
        switch 32768 {
        case 4294967295:
            var v16 uint64 = (^(v14))
            _ = v16
        case 2147483647:
            var v17 int64 = -9223372036854775808
            _ = v17
        }
        if false {
            v18 := GenS1{F0: 127, F1: 42, F2: 1}
            v19 := GenS2{F0: 2, F1: (^(v4))}
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 8; fi20++ {
            v15 += v15
            for fi21 := 0; fi21 < 4; fi21++ {
                v22 := v1[(fi21 * (fi20 + fi20))]
                _ = fi21
                _ = v22
            }
            var v23 string = (("foo bar" + "hello") + "")
            _ = fi20
            _ = v23
        }
    }
    var v24 string = "hello"
    _ = len(v3)
    v25 := v1[42]
    var v26 int32 = -2147483648
    for fi27 := 0; fi27 < 19; fi27++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi27
    }
    v28 := GenS2{F0: 127, F1: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v9
}
