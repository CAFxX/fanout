package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 string }

func main() {
    var v1 int8 = 127
    var v2 int8 = v1
    var v3 uint = 18446744073709551615
    var v4 float64 = float64(1e10)
    if (((!false) && true) || ((-728 < 1000) || false)) {
        v5 := GenS1{F0: v2}
        v6 := GenS1{F0: -128}
        _ = v5
        _ = v6
    } else {
        v7 := GenS2{F0: 32767, F1: ("hello" + ("" + "\n\t"))}
        var v8 int32 = -1
        v9 := []int{}
        _ = v7
        _ = v8
        _ = v9
    }
    v1 += v2
    v10 := []int{}
    var v11 float32 = 0.1
    var v12 string = ("" + "")
    fmt.Println(v12, v1, v4)
    var v13 uint16 = 32767
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 int16 = int16(v2)
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: (v1 & v2)}
    var v17 float64 = (0.1 - -2.25)
    v18 := GenS2{F0: 106, F1: ("\n\t" + v12)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
