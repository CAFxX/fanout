package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() uint {
    var v5 bool = true
    var v6 uint16 = 0
    v7 := []int{42, -1, 127}
    var v8 uint64 = 2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    v4 := make(map[uint32]string)
    var v9 int8 = 127
    if false {
        for fi10 := 0; fi10 < 9; fi10++ {
            v9 -= v9
            _ = fi10
        }
        v1 = v1
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 4; fi12++ {
            v13 := []int{32767, 670, -686, fi12}
            v4[2147483648] = "a"
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    if false {
        var v14 uint8 = 143
        switch "" {
        case "\n\t":
            for fi15 := 0; fi15 < 3; fi15++ {
                fmt.Println(v1, v14)
                v9 = -1
                _ = fi15
            }
        case "foo bar":
            v1 += 42
        }
        _ = v14
    } else {
        v16 := GenS1{F0: 0}
        _ = v16
    }
    var v17 uint8 = 89
    for fi18 := 0; fi18 < 8; fi18++ {
        v1 -= 0
        var v19 uint32 = 2
        var v20 uint = 223
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := []int{-511, 867, 32767, -490}
    if len(v21) > 0 {
        v21[0] = 127
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
