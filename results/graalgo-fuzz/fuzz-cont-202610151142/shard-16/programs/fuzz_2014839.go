package main

import (
    "fmt"
)


func main() {
    var v1 int32 = -833
    var v2 bool = false
    var v3 int = 32767
    var v4 int16 = 42
    v5 := []int{}
    v2 = (128 <= 58)
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int = 2147483647
        v3 = (^(v7))
        var v8 uint8 = 127
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(v3, v2)
    var v9 float64 = ((-2.25 / 1.0) * (-2.25 / float64(0.1)))
    var v10 uint16 = 660
    if (!(526 <= 703)) {
        var v11 int64 = 42
        _ = v11
    }
    v5 = append(v5, (^(v3)))
    fmt.Println(v4)
    if v2 {
        if len(v5) > 0 {
            v5[0] = v3
        }
        v2 = (!v2)
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        switch "hello" {
        case "hello":
            _ = len(v5)
        case "a":
            v2 = false
        }
        _ = fi12
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 string = "\n\t"
        v2 = v2
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi13
        _ = v14
    }
    v4 -= (v4 % ((v4) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
