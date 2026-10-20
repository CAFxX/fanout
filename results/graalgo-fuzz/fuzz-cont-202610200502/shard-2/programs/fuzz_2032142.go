package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    var v3 uint = 25
    v4 := []uint32{}
    v5 := make(map[int]string)
    var v6 uint = (^(v3))
    for fi7 := 0; fi7 < 5; fi7++ {
        v2[4294967295] = (("a" + "foo bar") + ("\n\t" + ""))
        for fi8 := 0; fi8 < 18; fi8++ {
            if len(v4) > 0 {
                v4[0] = 1
            }
            v9 := GenS1{F0: 42, F1: 891, F2: fi7}
            v6 -= 18446744073709551615
            var v10 uint16 = 14
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v11 uint8 = 255
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 int64 = 1
        _ = fi12
        _ = v13
    }
    var v14 uint16 = 523
    var v15 int = 42
    var v16 int64 = 42
    v16 = v16
    var v17 uint32 = 2147483647
    v6 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
