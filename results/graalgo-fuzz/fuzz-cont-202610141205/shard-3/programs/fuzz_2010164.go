package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint32; F1 int64 }
func genFn0(p0 int16) int32 {
    var v7 int32 = 42
    var v8 int = 1
    v9 := []int{}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (v7 - v7)
}

func main() {
    v1 := []string{"\n\t", "foo bar", ((("foo bar" + ((("\n\t" + "a") + "foo bar") + ("foo bar" + "hello"))) + "foo bar") + "\n\t")}
    var v2 int16 = -32768
    v3 := []uint32{4294967295, 588, 2147483647}
    var v4 float32 = (-2.25 + ((1.5 - float32(1.5)) / 1.5))
    var v5 uint8 = 127
    var v6 int8 = -111
    var v10 float64 = 1.5
    fmt.Println(v10)
    fmt.Println(v2, v10)
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    switch "a" {
    case "foo bar":
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 int = fi11
            _ = len(v3)
            _ = fi11
            _ = v12
        }
    default:
        var v13 uint32 = 42
        _ = v13
    }
    switch "\n\t" {
    case "hello":
        var v14 bool = false
        _ = v14
    case "\n\t":
        var v15 uint8 = 42
        _ = v15
    }
    var v16 uint32 = 65535
    for fi17 := 0; fi17 < 3; fi17++ {
        v4 += 0.1
        fmt.Println(fi17)
        _ = fi17
    }
    fmt.Println(v10, v6)
    fmt.Println(v16, v16)
    if false {
        v3 = append(v3, (v16 - (v16 | 1)))
    } else {
        v18 := GenS1{F0: (v16 & v16), F1: 2147483648, F2: v6}
        _ = v18
    }
    if ((!(!((^(v5)) == 2))) && true) {
        var v19 int32 = genFn0((v2 << ((42) & 63)))
        if (v2 >= v2) {
            v4 += v4
        }
        if (32767 < 958) {
            _ = len(v1)
            var v20 float64 = ((v10 + (-2.25 + (v10 * 0.1))) * ((((-2.25 - v10) * (v10 / v10)) - (3.14159 - 1e10)) / ((v10 - (v10 - (((v10 / (v10 - float64(v4))) * v10) + (((-2.25 - v10) / 1.0) + v10)))) * 1e10)))
            _ = v20
        } else {
            var v21 uint8 = 42
            v4 = v4
            _ = len(v3)
            _ = v21
        }
        _ = v19
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        v10 -= ((v10 - 1.5) + v10)
        v23 := GenS2{F0: v16, F1: -1}
        _ = fi22
        _ = v23
    }
    if (!false) {
        if false {
            v5 += (v5 * (^(v5)))
        }
        v2 -= (v2 - (v2 | 0))
    } else {
        v24 := GenS2{F0: v16, F1: 9223372036854775807}
        var v25 uint64 = 4294967295
        _ = v24
        _ = v25
    }
    v4 -= v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
