package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int16 = 32767
    v2 := []uint32{}
    var v3 string = "hello"
    var v4 int = 128
    var v5 uint64 = 1
    var v6 bool = false
    var v7 int64 = 1
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := GenS1{F0: 842}
        v10 := GenS1{F0: (v5 | (v5 - (v5 * v5)))}
        v1 += v1
        _ = fi8
        _ = v10
        _ = v9
    }
    if len(v2) > 0 {
        v2[0] = 28
    }
    v5 += v5
    fmt.Println(v6, v4)
    for fi11 := 0; fi11 < 6; fi11++ {
        fmt.Println(v7)
        if (false && (!v6)) {
            v5 -= v5
        } else {
            for fi12 := 0; fi12 < 6; fi12++ {
                switch "foo bar" {
                case "":
                    var v13 int32 = -2147483648
                    _ = v13
                }
                _ = fi12
            }
        }
        var v14 uint8 = 127
        _ = fi11
        _ = v14
    }
    fmt.Println(v5, v3)
    v5 = (v5 ^ 516)
    fmt.Println(v6, v3)
    v15 := GenS1{F0: (v5 & v5)}
    v1 += -156
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) > 0 {
        v2[0] = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
