package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 string }
func genFn0(p0 int16) uint32 {
    var v7 int64 = 324
    v8 := []int{}
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint32 = 177
    var v2 int32 = 2147483647
    var v3 float32 = 0.1
    v4 := []int{-2147483648}
    var v5 float32 = (v3 / (v3 / 1.0))
    var v6 int16 = -1
    for fi9 := 0; fi9 < 1; fi9++ {
        if true {
            for fi10 := 0; fi10 < 16; fi10++ {
                v5 -= v3
                v5 -= v5
                _ = fi10
            }
        }
        _ = fi9
    }
    if ((v5 + (float32(0.0) - float32(1.5))) != v3) {
        v5 = v5
        switch "foo bar" {
        case "\n\t":
            for fi11 := 0; fi11 < 20; fi11++ {
                var v12 int64 = 1
                _ = fi11
                _ = v12
            }
        case "":
            v3 -= float32(3.14159)
        case "foo bar":
            v6 -= -15
        }
        v2 = 2147483647
    }
    v6 = 32767
    v2 = v2
    if (0 != 0) {
        var v13 uint64 = 187
        _ = v13
    }
    var v14 float64 = (-2.25 + 0.1)
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 float64 = (-2.25 - v14)
        _ = fi15
        _ = v16
    }
    var v17 int16 = v6
    var v18 uint64 = 2
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 int64 = -536
        var v21 int64 = (v20 + (-(v20)))
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 int16 = (v6 - v6)
    v23 := GenS1{F0: v1, F1: 32767, F2: ((("a" + "foo bar") + "hello") + "\n\t")}
    v3 = float32(v14)
    v24 := GenS1{F0: (^(v1)), F1: 42, F2: "a"}
    var v25 string = (("a" + "\n\t") + "\n\t")
    v1 = (v1 / (((v1 - (v1 & 2147483647))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
