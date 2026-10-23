package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []uint32{}
    v2 := make(map[int]int)
    v3 := []int{4294967295, -77}
    var v4 uint32 = 0
    var v5 int16 = -32768
    v5 = 1000
    var v6 uint8 = 255
    switch 127 {
    case 333:
        switch 4294967295 {
        case 32767:
            if false {
                var v7 uint8 = 0
                _ = v7
            } else {
                v4 += (v4 % ((65535) | 1))
                v4 = (v4 % (((v4 * (^(v4)))) | 1))
            }
        case 496:
            v4 = v4
        case 434:
            fmt.Println(v5)
        }
    }
    v5 = v5
    v3 = append(v3, -2147483648)
    v4 = (v4 + (v4 * (v4 << ((121) & 63))))
    v8 := GenS1{F0: true}
    var v9 uint = 255
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 float64 = 3.14159
        var v12 uint = v9
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint8 = (^(v6))
    v14 := GenS1{F0: false}
    var v15 int32 = 291
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
