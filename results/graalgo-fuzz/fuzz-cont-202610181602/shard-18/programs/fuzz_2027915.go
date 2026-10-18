package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint32 }

func main() {
    var v1 uint8 = 241
    var v2 uint32 = 2
    var v3 bool = ((false && true) && true)
    switch 1 {
    case 4294967295:
        fmt.Println(v1, v3, v3)
    }
    for fi4 := 0; fi4 < 14; fi4++ {
        v2 += (v2 >> ((740) & 63))
        _ = fi4
    }
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 uint8 = v1
        v6 = v6
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: "a", F1: "\n\t", F2: 0}
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            v2 = 568
            v10 := []int{fi8, 539}
            _ = fi9
            _ = v10
        }
        v2 = 4294967295
        v2 += v2
        _ = fi8
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 uint16 = 96
        _ = fi11
        _ = v12
    }
    v13 := []int{127, 4294967295}
    if (v3 || ((!(0 == 18446744073709551615)) || (1.5 >= ((1.5 - 1.5) - 0.1)))) {
        var v14 int = 865
        _ = v14
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v7
}
