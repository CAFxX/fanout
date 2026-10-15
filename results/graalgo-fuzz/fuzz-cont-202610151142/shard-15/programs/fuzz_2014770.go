package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
func genFn0() int8 {
    v4 := []int{2147483647, 2147483647, -796}
    var v5 int8 = 42
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 string = ("hello" + "foo bar")
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 string = "\n\t"
        if len(v4) > 0 {
            v4[0] = fi9
        }
        _ = fi9
        _ = v10
    }
    _ = v4
    _ = v5
    return 91
}
func genFn1(p0 int8) int64 {
    switch "" {
    case "\n\t":
        var v11 float64 = (((3.14159 / 3.14159) - -2.25) / 1.0)
        _ = v11
    case "a":
        p0 += (-(p0))
    case "foo bar":
        if true {
            var v12 uint8 = 2
            _ = v12
        } else {
            p0 -= 1
            v13 := GenS1{F0: 87, F1: 2}
            p0 = (p0 / ((p0) | 1))
            _ = v13
        }
    }
    p0 -= 100
    var v14 float32 = -2.25
    v14 -= 0.1
    _ = p0
    _ = v14
    return 9223372036854775807
}

func main() {
    var v1 uint64 = 222
    var v2 float64 = (0.1 / 0.1)
    var v3 int = 127
    if true {
        var v15 float64 = v2
        var v16 string = "foo bar"
        _ = v15
        _ = v16
    } else {
        if true {
            var v17 float64 = (-2.25 - v2)
            _ = v17
        }
    }
    v18 := GenS1{F0: 552, F1: 0}
    fmt.Println(v3, v3, v1)
    if ((255 != 65535) && true) {
        v19 := []int{v3, 92, (v3 / ((v3) | 1))}
        var v20 uint64 = 1
        _ = v19
        _ = v20
    } else {
        v21 := GenS1{F0: 2, F1: v1}
        fmt.Println(v1, v3, v3)
        _ = v21
    }
    v1 += 0
    v1 += (v1 << ((65535) & 63))
    var v22 int8 = -128
    var v23 int16 = 32767
    for fi24 := 0; fi24 < 3; fi24++ {
        v25 := GenS1{F0: 0, F1: v1}
        v2 += float64(1e10)
        for fi26 := 0; fi26 < 2; fi26++ {
            v23 -= -32768
            _ = fi26
        }
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
