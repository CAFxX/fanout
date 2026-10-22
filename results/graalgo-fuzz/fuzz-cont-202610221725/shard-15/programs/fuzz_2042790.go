package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int) uint8 {
    for fi5 := 0; fi5 < 1; fi5++ {
        p0 -= (p0 * 127)
        var v6 uint = 0
        _ = fi5
        _ = v6
    }
    var v7 bool = ((1000 == 1) && true)
    switch p0 {
    case 0:
        p0 -= p0
    }
    _ = p0
    _ = v7
    return 160
}

func main() {
    var v1 int8 = 42
    var v2 string = (("a" + "\n\t") + "a")
    var v3 uint64 = 1
    var v4 float64 = 3.14159
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 uint32 = 713
        v10 := GenS1{F0: 42, F1: -1}
        _ = fi8
        _ = v10
        _ = v9
    }
    v1 += (v1 % (((v1 + (v1 % ((v1) | 1)))) | 1))
    v4 = v4
    var v11 float64 = v4
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 string = (v2 + "a")
        _ = fi12
        _ = v13
    }
    fmt.Println(v2, v11)
    v14 := GenS2{F0: 2}
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
