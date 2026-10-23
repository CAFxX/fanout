package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
func genFn0() bool {
    var v7 int8 = 120
    fmt.Println(v7)
    _ = v7
    return (!(0 == -9223372036854775808))
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 0.0
    v3 := make(map[int]int)
    v4 := make(map[string]string)
    v5 := make(map[int]string)
    var v6 string = "a"
    v6 = v6
    switch 127 {
    case 2147483647:
        v2 += 1.5
    }
    if (v6 < ("hello" + v6)) {
        for fi8 := 0; fi8 < 1; fi8++ {
            fmt.Println(fi8, v2, v6)
            _ = fi8
        }
        v9 := []int{127, 32767, -9223372036854775808}
        _ = v9
    }
    delete(v1, 2)
    fmt.Println(v6, v2)
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 float64 = (float64(v2) + 0.0)
        var v12 uint16 = 2
        if false {
            v13 := GenS1{F0: 42, F1: 0}
            var v14 float32 = v2
            _ = v13
            _ = v14
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    v6 = ("foo bar" + v6)
    var v15 float64 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
