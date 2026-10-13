package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int64; F1 string; F2 uint8 }

func main() {
    var v1 uint64 = 664
    var v2 uint32 = 682
    v3 := []uint32{v2, 405, v2}
    var v4 float32 = 1e10
    var v5 uint = 18446744073709551615
    for fi6 := 0; fi6 < 13; fi6++ {
        if len(v3) > 0 {
            v3[0] = (v2 << ((255) & 63))
        }
        fmt.Println(v2)
        _ = fi6
    }
    v7 := GenS1{F0: 42, F1: 42}
    v2 += v2
    v4 = 1.5
    if len(v3) > 0 {
        v3[0] = (v2 ^ v2)
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS2{F0: 42, F1: "a", F2: 42}
        _ = fi8
        _ = v9
    }
    v2 = 729
    v1 -= 731
    v5 -= 4294967295
    var v10 int8 = -76
    if len(v3) > 0 {
        v3[0] = 0
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := GenS2{F0: 96, F1: "hello", F2: 42}
        v2 -= v2
        if true {
            v5 -= v5
            for fi13 := 0; fi13 < 6; fi13++ {
                if ((v10 ^ -128) < v10) {
                    if true {
                        v1 = 0
                        v3 = append(v3, v2)
                        v14 := GenS2{F0: 9223372036854775807, F1: "hello", F2: 42}
                        _ = v14
                    } else {
                        fmt.Println(v1, v4)
                    }
                }
                fmt.Println(v10)
                switch ("a" + "a") {
                case "foo bar":
                    v1 = v1
                }
                _ = fi13
            }
        } else {
            v4 -= v4
        }
        _ = fi11
        _ = v12
    }
    if (true && true) {
        var v15 int16 = 1
        _ = v15
    }
    var v16 float64 = (float64((v4 - float32(float64(((v4 * float32(0.1)) + 1.5))))) - 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
