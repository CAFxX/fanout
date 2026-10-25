package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 uint8; F2 int32 }
func genFn0(p0 bool, p1 uint, p2 int64) uint64 {
    fmt.Println(p2)
    if p0 {
        for fi5 := 0; fi5 < 1; fi5++ {
            p1 = p1
            _ = fi5
        }
    } else {
        p2 = p2
    }
    p2 += p2
    p0 = ((0.1 / 1.0) >= (float32(1.5) * ((((1e10 * (((((float32(float64(1.5)) + (-2.25 / 1e10)) * 1.5) + 0.0) / ((((-2.25 - 0.1) / ((1.5 * 0.1) + ((float32(((0.1 * float64(0.0)) * (-2.25 + 1e10))) - 1e10) / 0.1))) - ((((0.0 - (3.14159 * ((((float32(0.0) * 1e10) + 0.1) * ((-2.25 - ((((1.5 / 1.0) * 0.1) * 0.0) * -2.25)) - float32(3.14159))) + 1e10))) - 3.14159) / 0.1) * 0.0)) - (1e10 + (1.5 - 0.1)))) / 0.1)) + (((3.14159 - (3.14159 / 1.0)) * 0.1) - (0.1 - 1.5))) / 1.5) - 1.5)))
    _ = p0
    _ = p1
    _ = p2
    return 379
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 1e10
    v3 := make(map[string]uint32)
    var v4 uint16 = 65535
    v4 = v4
    switch (((("" + ("" + "a")) + "foo bar") + "\n\t") + "") {
    case "hello":
        var v6 uint64 = 888
        _ = v6
    case "":
        var v7 uint16 = (v4 * (v4 ^ v4))
        _ = v7
    default:
        var v8 int8 = 1
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := []int{fi9, -805, fi9}
        for fi11 := 0; fi11 < 7; fi11++ {
            v12 := GenS1{F0: "foo bar", F1: -32768, F2: 2147483647}
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: "\n\t", F1: 42, F2: -2147483648}
        _ = fi9
        _ = v10
        _ = v13
    }
    if (42 >= 705) {
        v14 := []int{-1}
        fmt.Println(v4)
        for fi15 := 0; fi15 < 13; fi15++ {
            fmt.Println(v4, fi15)
            v14 = append(v14, 0)
            _ = fi15
        }
        _ = v14
    }
    v16 := GenS1{F0: "hello", F1: 487, F2: 2147483647}
    var v17 int32 = -650
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 float32 = 1e10
        _ = fi18
        _ = v19
    }
    v4 = v4
    v20 := GenS2{F0: 195}
    switch ("foo bar" + "a") {
    case "hello":
        v21 := GenS3{F0: 299, F1: 2, F2: -2147483648}
        _ = v21
    case "a":
        v3["foo bar"] = 1
    case "":
        v22 := GenS2{F0: 0}
        _ = v22
    }
    if false {
        for fi23 := 0; fi23 < 4; fi23++ {
            var v24 int64 = 42
            v24 -= (v24 - 9223372036854775807)
            _ = fi23
            _ = v24
        }
    } else {
        if false {
            v3[("a" + "foo bar")] = 0
        }
        var v25 int64 = 0
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
