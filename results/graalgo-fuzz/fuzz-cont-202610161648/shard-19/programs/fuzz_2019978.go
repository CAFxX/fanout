package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint8; F2 int8 }

func main() {
    var v1 int64 = 1
    var v2 int64 = v1
    var v3 uint = 65535
    v4 := []string{}
    _ = len(v4)
    fmt.Println(v1)
    fmt.Println(v3)
    for fi5 := 0; fi5 < 18; fi5++ {
        if (!false) {
            var v6 float64 = (0.0 * float64(-2.25))
            for fi7 := 0; fi7 < 9; fi7++ {
                if len(v4) > 0 {
                    v4[0] = ""
                }
                v3 = (v3 - 65535)
                v8 := GenS1{F0: 0, F1: 39}
                v9 := GenS2{F0: 4294967295, F1: 231, F2: 100}
                v6 += (-2.25 * v6)
                _ = fi7
                _ = v8
                _ = v9
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v6
        }
        v3 = v3
        v4 = append(v4, "hello")
        _ = fi5
    }
    fmt.Println(v3, v3)
    switch 4294967295 {
    case 2147483647:
        fmt.Println(v2)
    case 0:
        var v10 uint32 = uint32(v3)
        _ = v10
    case 285:
        fmt.Println(v1, v3, v3)
    }
    v2 = -208
    v1 -= v2
    _ = len(v4)
    v3 += 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
