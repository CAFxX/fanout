package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int }

func main() {
    var v1 uint16 = 813
    var v2 uint16 = (v1 << ((4294967295) & 63))
    var v3 int64 = 668
    var v4 uint32 = 42
    fmt.Println(v4, v4)
    v5 := GenS1{F0: ""}
    fmt.Println(v2, v4)
    for fi6 := 0; fi6 < 18; fi6++ {
        v1 = v1
        _ = fi6
    }
    if (2147483648 >= v4) {
        v2 -= 65535
        switch 32767 {
        case 42:
            v4 = 42
        }
    }
    fmt.Println(v1, v2)
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(fi7)
        if (1e10 >= 0.1) {
            for fi8 := 0; fi8 < 4; fi8++ {
                fmt.Println(v1)
                v2 -= v2
                var v9 uint32 = (^(v4))
                _ = fi8
                _ = v9
            }
            v10 := []int{}
            v2 = (^(v1))
            _ = v10
        }
        _ = fi7
    }
    var v11 string = "foo bar"
    v3 += (v3 * (v3 + v3))
    v11 = "foo bar"
    v12 := GenS2{F0: int16(v3), F1: 32768}
    v13 := []int{700, -1, 128, 128}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
