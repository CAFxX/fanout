package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint16 = 2
    var v2 int64 = 0
    var v3 float64 = 3.14159
    v4 := []int{465, int(v2), 730}
    v5 := []int{310, 571}
    var v6 int32 = 0
    var v7 string = "a"
    for fi8 := 0; fi8 < 11; fi8++ {
        fmt.Println(v6)
        if len(v4) > 0 {
            v4[0] = (fi8 - 2147483647)
        }
        var v9 string = "a"
        _ = fi8
        _ = v9
    }
    if true {
        v4 = append(v4, 32768)
        v1 -= v1
        v10 := GenS1{F0: ("hello" + (((v7 + v7) + v7) + (v7 + v7)))}
        _ = v10
    }
    var v11 int8 = 7
    var v12 uint16 = v1
    v13 := GenS1{F0: (v7 + v7)}
    for fi14 := 0; fi14 < 5; fi14++ {
        v3 -= (3.14159 + (1.5 - (0.0 * v3)))
        for fi15 := 0; fi15 < 11; fi15++ {
            v2 -= 326
            fmt.Println(v3, fi14, v12)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi15
        }
        _ = fi14
    }
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
