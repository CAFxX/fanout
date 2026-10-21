package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := []uint32{}
    v2 := []int{-991, 32768, -2147483648}
    var v3 float64 = 1e10
    v4 := make(map[string]string)
    var v5 int16 = -170
    switch 1 {
    case -1:
        var v6 bool = (!false)
        _ = v6
    case 0:
        var v7 uint = 141
        _ = v7
    case -636:
        switch 32767 {
        case -435:
            switch 4294967295 {
            case -9223372036854775808:
                var v8 int64 = -1
                _ = v8
            case 1:
                var v9 uint8 = 2
                _ = v9
            }
        }
    default:
        v5 = v5
    }
    v5 = v5
    switch ((((("\n\t" + "") + "") + ("\n\t" + "a")) + "a") + ("a" + ("a" + ""))) {
    case "foo bar":
        fmt.Println(v3, v3)
    default:
        _ = len(v2)
    }
    v5 = v5
    v5 -= 63
    if true {
        for fi10 := 0; fi10 < 14; fi10++ {
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 int32 = 662
                _ = fi11
                _ = v12
            }
            switch fi10 {
            case 236:
                var v13 int8 = -2
                _ = v13
            case 0:
                v14 := GenS1{F0: (fi10 * fi10)}
                _ = v14
            case -9223372036854775808:
                var v15 int = (fi10 | fi10)
                _ = v15
            default:
                var v16 int64 = -511
                _ = v16
            }
            _ = fi10
        }
    } else {
        var v17 uint = 4294967295
        _ = v17
    }
    _ = len(v1)
    fmt.Println(v5, v5)
    _ = len(v4)
    v5 -= v5
    v3 = (float64(0.1) - v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
