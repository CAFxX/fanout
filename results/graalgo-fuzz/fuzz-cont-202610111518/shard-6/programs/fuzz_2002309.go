package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 bool }

func main() {
    var v1 float64 = ((3.14159 * float64(0.1)) + 3.14159)
    v2 := make(map[uint32]bool)
    var v3 int32 = 1
    fmt.Println(v1, v1)
    if (1 == -32768) {
        var v4 float32 = float32(0.1)
        switch 127 {
        case 4294967295:
            fmt.Println(v3)
        }
        switch ("\n\t" + "hello") {
        case "a":
            v5 := GenS1{F0: 4294967295, F1: 0, F2: (v3 == (-(v3)))}
            _ = v5
        case "\n\t":
            v6 := []int{}
            _ = v6
        }
        _ = v4
    }
    if true {
        var v7 int64 = -458
        _ = v7
    } else {
        switch 0 {
        case 32767:
            var v8 float32 = 1e10
            _ = v8
        case 0:
            for fi9 := 0; fi9 < 19; fi9++ {
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 bool = (v1 < v1)
                    fmt.Println(fi10, v1, fi9)
                    v3 -= 884
                    v11 = ((v3 << ((4294967295) & 63)) <= v3)
                    _ = fi10
                    _ = v11
                }
                fmt.Println(v3, v1, fi9)
                _ = len(v2)
                _ = fi9
            }
        default:
            var v12 int64 = 682
            _ = v12
        }
        var v13 uint64 = 4294967295
        v3 -= int32(v13)
        _ = v13
    }
    var v14 int32 = (v3 ^ (v3 * (^(v3))))
    v15 := GenS1{F0: -1, F1: 1, F2: true}
    for fi16 := 0; fi16 < 16; fi16++ {
        v14 -= (v14 & 0)
        v17 := GenS1{F0: 383, F1: -1, F2: false}
        _ = fi16
        _ = v17
    }
    var v18 string = ("foo bar" + "hello")
    if false {
        if (952 == -95) {
            v2[0] = false
            v2[120] = false
            var v19 int32 = (^(v3))
            _ = v19
        }
        v1 += (v1 / v1)
        var v20 float32 = -2.25
        _ = v20
    }
    v3 -= -608
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
}
