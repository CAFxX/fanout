package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint8, p1 int16) uint32 {
    var v7 uint8 = 229
    v8 := []int{42, -688, -880}
    var v9 int16 = (p1 % ((32767) | 1))
    fmt.Println(p1, p0)
    if (2147483647 == -1) {
        fmt.Println(v9, p1, v9)
        var v10 int32 = 555
        v9 = (v9 - p1)
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 907
}

func main() {
    var v1 int = -1
    var v2 bool = (!false)
    v3 := []string{}
    var v4 int16 = 42
    var v5 uint16 = 298
    v6 := make(map[string]int)
    v11 := GenS1{F0: 2147483647}
    var v12 string = ("hello" + "")
    v13 := v6[v12]
    switch v1 {
    case 503:
        for fi14 := 0; fi14 < 18; fi14++ {
            _ = len(v6)
            for fi15 := 0; fi15 < 4; fi15++ {
                v16 := v6["a"]
                fmt.Println(v16)
                v12 = v12
                _ = fi15
                _ = v16
            }
            fmt.Println(v4)
            _ = fi14
        }
    case 519:
        var v17 uint64 = 42
        _ = v17
    default:
        if ((float32(-2.25) < 1.5) && (false && true)) {
            var v18 int32 = -686
            for fi19 := 0; fi19 < 11; fi19++ {
                v12 = (("\n\t" + v12) + ((("hello" + v12) + (v12 + v12)) + (("a" + v12) + "hello")))
                _ = fi19
            }
            _ = v18
        }
    }
    v20 := GenS2{F0: 0, F1: (^(v5))}
    fmt.Println(v12, v4, v2)
    v5 -= (v5 << ((0) & 63))
    var v21 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
