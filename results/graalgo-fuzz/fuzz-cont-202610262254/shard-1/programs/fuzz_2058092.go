package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }

func main() {
    var v1 uint16 = 574
    var v2 int16 = -32768
    v3 := []int{}
    var v4 string = "foo bar"
    var v5 uint32 = 65535
    if true {
        fmt.Println(v1)
    }
    v4 = v4
    v5 -= (v5 / ((2147483648) | 1))
    switch 4294967295 {
    case -269:
        v6 := GenS1{F0: (^(v1)), F1: -1}
        _ = v6
    case 2147483647:
        if (!(42 <= (v5 ^ uint32(v1)))) {
            v1 += v1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            fmt.Println(v2, v1)
        }
    case 168:
        fmt.Println(v4, v4, v2)
    }
    fmt.Println(v2, v4)
    if len(v3) > 0 {
        v3[0] = 877
    }
    var v7 int16 = v2
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 bool = (!(int16(fi8) == 32767))
        _ = fi8
        _ = v9
    }
    var v10 string = (v4 + (v4 + v4))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
