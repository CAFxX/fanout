package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 22
    var v3 float64 = -2.25
    var v4 int = 4294967295
    switch 32768 {
    case 4294967295:
        v5 := []int{}
        _ = v5
    default:
        v3 = 1e10
    }
    var v6 float32 = 1e10
    v7 := GenS1{F0: (v2 | v2)}
    var v8 int32 = 1
    v6 += 0.0
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 bool = (597 > 1)
        switch v1 {
        case "foo bar":
            v11 := []int{v4, (v4 & 127), (v4 * 32767)}
            _ = v11
        case "":
            v12 := GenS1{F0: (^(v2))}
            _ = v12
        case "a":
            v4 = 42
        default:
            v4 += (-(v4))
        }
        var v13 uint8 = 2
        _ = fi9
        _ = v10
        _ = v13
    }
    var v14 uint64 = 4294967295
    var v15 bool = false
    if (!v15) {
        v16 := GenS2{F0: 0}
        _ = v16
    } else {
        var v17 uint8 = v2
        var v18 int64 = 1
        _ = v17
        _ = v18
    }
    var v19 uint16 = 255
    var v20 int16 = 42
    for fi21 := 0; fi21 < 4; fi21++ {
        for fi22 := 0; fi22 < 18; fi22++ {
            v23 := GenS2{F0: v14}
            _ = fi22
            _ = v23
        }
        var v24 int8 = -1
        _ = fi21
        _ = v24
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        var v26 int8 = -49
        _ = fi25
        _ = v26
    }
    v8 = v8
    for fi27 := 0; fi27 < 17; fi27++ {
        v2 += v2
        v6 = (v6 - (v6 / v6))
        v28 := GenS2{F0: uint64(v2)}
        _ = fi27
        _ = v28
    }
    var v29 uint64 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
