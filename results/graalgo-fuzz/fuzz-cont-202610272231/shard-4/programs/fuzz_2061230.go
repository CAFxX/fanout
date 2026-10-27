package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint; F1 string }

func main() {
    var v1 string = "foo bar"
    var v2 bool = (0 == 42)
    var v3 int8 = -128
    v4 := make(map[int]int)
    var v5 uint = 1
    var v6 uint32 = 4294967295
    for fi7 := 0; fi7 < 18; fi7++ {
        v3 = (v3 / ((v3) | 1))
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 127, F1: 127}
    if false {
        v5 = (^(v5))
    }
    v10 := []int{-392}
    for fi11 := 0; fi11 < 10; fi11++ {
        v2 = true
        var v12 uint16 = 255
        _ = fi11
        _ = v12
    }
    v5 -= v5
    fmt.Println(v1)
    v1 = v1
    fmt.Println(v1)
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            v2 = false
            v15 := v4[128]
            v16 := v4[fi13]
            _ = fi14
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 15; fi17++ {
            v3 -= 42
            _ = fi17
        }
        _ = fi13
    }
    v2 = (((v1 + ("" + "\n\t")) + v1) == v1)
    v18 := GenS2{F0: (v5 & (^(v5))), F1: "hello"}
    var v19 uint32 = v6
    fmt.Println(v2, v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
