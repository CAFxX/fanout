package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() bool {
    var v6 uint16 = 796
    fmt.Println(v6)
    v7 := []int{}
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 uint32 = 1
        _ = fi8
        _ = v9
    }
    _ = v6
    _ = v7
    return false
}

func main() {
    var v1 int = 32767
    var v2 int16 = -32768
    var v3 float64 = (1e10 / 1.0)
    var v4 float32 = -2.25
    v5 := make(map[int]int)
    var v10 uint64 = 524
    v2 = v2
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 7; fi12++ {
            v3 = ((0.1 / (float64(0.0) * float64((v4 - v4)))) + (-2.25 - float64(1.5)))
            _ = fi12
        }
        _ = fi11
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        if true {
            fmt.Println(v10)
            v4 = float32(float64(((v4 + (v4 - v4)) - ((v4 - v4) * (v4 + (v4 - (v4 * 1e10)))))))
            var v14 string = "foo bar"
            _ = v14
        }
        v10 += 2
        _ = fi13
    }
    var v15 uint8 = 128
    v2 += v2
    v16 := GenS1{F0: 42}
    switch "foo bar" {
    case "a":
        var v17 uint16 = 255
        _ = v17
    case "\n\t":
        v4 += float32((3.14159 + 3.14159))
    case "hello":
        v3 += 1e10
    }
    v18 := GenS1{F0: 1}
    var v19 uint16 = 447
    switch (v1 ^ (-(v1))) {
    case 2147483647:
        var v20 uint = 42
        _ = v20
    default:
        fmt.Println(v1, v3, v19)
    }
    v21 := GenS1{F0: 905}
    switch "foo bar" {
    case "a":
        var v22 bool = ((^(v1)) != (v1 & v1))
        _ = v22
    case "\n\t":
        for fi23 := 0; fi23 < 2; fi23++ {
            var v24 int8 = -78
            var v25 int32 = 631
            v19 = v19
            _ = fi23
            _ = v24
            _ = v25
        }
    case "foo bar":
        v26 := GenS1{F0: -1}
        _ = v26
    }
    v5[1] = v1
    switch "\n\t" {
    case "":
        v19 += uint16(v15)
    case "a":
        v5[32768] = v1
    case "hello":
        for fi27 := 0; fi27 < 12; fi27++ {
            v15 = v15
            _ = fi27
        }
    }
    v10 -= v10
    v3 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
