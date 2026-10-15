package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{1}
    var v3 int16 = 0
    var v4 uint8 = 1
    var v5 int16 = v3
    v4 = 128
    v3 = (v5 & 32767)
    if (false || ((42 < -936) || true)) {
        v4 = (v4 * v4)
        fmt.Println(v5, v4, v5)
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := GenS1{F0: 900}
        var v8 uint8 = 1
        for fi9 := 0; fi9 < 4; fi9++ {
            fmt.Println(fi6)
            v4 += v8
            _ = fi9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v10 := GenS2{F0: -9223372036854775808}
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    var v11 int64 = 1
    for fi12 := 0; fi12 < 20; fi12++ {
        if (!true) {
            v4 = (^(v4))
            if ((true || false) || (0 != 698)) {
                switch fi12 {
                case 9223372036854775807:
                    if (1 < (v11 * (v11 / ((1) | 1)))) {
                        var v13 uint16 = 2
                        v1[("hello" + ("\n\t" + "hello"))] = (-805 >= fi12)
                        _ = v13
                    }
                }
                for fi14 := 0; fi14 < 15; fi14++ {
                    v15 := GenS2{F0: (v11 % ((866) | 1))}
                    var v16 bool = true
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            v1[""] = true
        }
        v4 = v4
        _ = fi12
    }
    var v17 uint16 = 32767
    var v18 uint8 = 128
    for fi19 := 0; fi19 < 3; fi19++ {
        fmt.Println(v5)
        v11 = v11
        v5 = 32767
        _ = fi19
    }
    _ = len(v1)
    fmt.Println(v11)
    if len(v2) > 0 {
        v2[0] = 909
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
