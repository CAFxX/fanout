package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
func genFn0(p0 string, p1 int16, p2 int) uint8 {
    fmt.Println(p2, p2, p2)
    fmt.Println(p0, p0)
    var v6 int8 = 100
    var v7 string = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 127
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 65535
    v3 := []int{}
    var v4 string = ("a" + "a")
    var v5 uint64 = 42
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint8 = 42
        _ = fi8
        _ = v9
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        delete(v1, (^(fi10)))
        _ = fi10
    }
    var v11 int32 = -1
    fmt.Println(v2)
    if len(v3) > 0 {
        v3[0] = 819
    }
    v12 := GenS1{F0: 0, F1: 42}
    var v13 uint64 = (^(v5))
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 11; fi15++ {
            v16 := GenS1{F0: 0, F1: 9223372036854775807}
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    var v17 uint = 255
    v13 = v5
    v11 -= v11
    v17 -= v17
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
