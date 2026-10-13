package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int16 }

func main() {
    var v1 float64 = float64(3.14159)
    v2 := []string{"", ("\n\t" + "a"), ""}
    v3 := make(map[string]bool)
    v4 := []int{}
    var v5 int64 = -1
    var v6 float64 = v1
    var v7 uint32 = 2147483647
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 string = ("a" + "hello")
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            v6 = (0.1 / 3.14159)
            v5 -= (^(v5))
            var v12 uint16 = 548
            var v13 uint8 = 127
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    v7 -= 0
    for fi14 := 0; fi14 < 20; fi14++ {
        v5 += 506
        v15 := GenS1{F0: -1, F1: 42, F2: 42}
        _ = fi14
        _ = v15
    }
    var v16 uint16 = 255
    var v17 string = ("hello" + "a")
    if true {
        var v18 bool = false
        _ = v18
    } else {
        var v19 int16 = 42
        if len(v2) > 0 {
            v2[0] = ((v17 + "a") + v17)
        }
        for fi20 := 0; fi20 < 19; fi20++ {
            v5 -= (^(v5))
            v21 := GenS1{F0: 0, F1: 42, F2: (v19 ^ v19)}
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    fmt.Println(v16, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
