package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
func genFn0(p0 int16) uint {
    var v5 bool = false
    v6 := GenS1{F0: false}
    _ = p0
    _ = v5
    _ = v6
    return 255
}
func genFn1() bool {
    v7 := GenS1{F0: true}
    var v8 uint32 = 284
    v9 := []int{-1}
    _ = v7
    _ = v8
    _ = v9
    return false
}

func main() {
    v1 := []int{4294967295, 454, -163, 0}
    v2 := []int{32767, 42, -2147483648, 42}
    var v3 int32 = -231
    v4 := []int{401}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v10 float64 = 1.5
    if true {
        var v11 int16 = 1
        _ = v11
    }
    v3 -= (v3 + -120)
    for fi12 := 0; fi12 < 10; fi12++ {
        fmt.Println(v10, v3, v10)
        var v13 int16 = 0
        for fi14 := 0; fi14 < 10; fi14++ {
            v1 = append(v1, 32768)
            v13 = -32768
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v15 float64 = (1.5 - float64(1.5))
    var v16 int8 = -20
    v17 := GenS2{F0: 127}
    for fi18 := 0; fi18 < 4; fi18++ {
        v15 += 1.5
        _ = len(v4)
        var v19 int32 = 42
        _ = fi18
        _ = v19
    }
    switch "foo bar" {
    case "hello":
        if true {
            fmt.Println(v10, v15)
        }
    case "foo bar":
        var v20 string = "\n\t"
        _ = v20
    case "\n\t":
        v21 := GenS2{F0: (v16 - v16)}
        _ = v21
    default:
        v2 = append(v2, -260)
    }
    v10 = ((3.14159 - v15) / 0.1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
