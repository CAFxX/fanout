package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint64 }

func main() {
    var v1 uint8 = 127
    var v2 int32 = -2147483648
    var v3 string = "hello"
    var v4 uint32 = 983
    v5 := []int{32768}
    var v6 float32 = 1e10
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: 18446744073709551615, F1: (-1 > 14)}
        fmt.Println(v3, fi7)
        v2 -= 0
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 18446744073709551615
    v1 += uint8(v4)
    if false {
        fmt.Println(v3, v2, v4)
    }
    var v10 uint64 = (v9 / ((v9) | 1))
    var v11 int32 = v2
    var v12 uint16 = 282
    v12 += (v12 >> ((uint(v12)) & 63))
    v13 := GenS2{F0: uint32(v12), F1: 2, F2: uint64(v12)}
    v2 += 1
    for fi14 := 0; fi14 < 10; fi14++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if (v12 > v12) {
            var v15 int16 = 42
            _ = v15
        } else {
            v11 = (v11 ^ 42)
            switch "foo bar" {
            case "foo bar":
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 int16 = 0
                    v17 = 32767
                    v5 = append(v5, fi14)
                    _ = fi16
                    _ = v17
                }
            default:
                for fi18 := 0; fi18 < 5; fi18++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi18
                }
            }
        }
        _ = fi14
    }
    var v19 uint64 = v9
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
