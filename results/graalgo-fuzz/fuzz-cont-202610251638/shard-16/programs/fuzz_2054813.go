package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int32; F1 uint; F2 uint8 }
func genFn0() int32 {
    v6 := GenS1{F0: 4294967295, F1: -405}
    v7 := GenS1{F0: 2147483647, F1: 0}
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 uint64 = 762
    var v2 uint16 = 956
    v3 := []uint32{65535}
    v4 := make(map[string]int)
    var v5 bool = (false || (!true))
    v2 = v2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "a" {
    case "":
        v1 = 2
    }
    if false {
        for fi8 := 0; fi8 < 9; fi8++ {
            v5 = ((!v5) && (42 <= -1))
            _ = fi8
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if v5 {
        var v9 uint16 = v2
        _ = v9
    } else {
        fmt.Println(v1, v1)
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        v2 = v2
        _ = fi10
    }
    var v11 int16 = 1000
    fmt.Println(v2, v5)
    v1 += 1
    v12 := GenS2{F0: 305, F1: 857, F2: 255}
    var v13 uint64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
