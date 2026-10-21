package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int = 4294967295
    v2 := []uint32{}
    v3 := []int{v1, (v1 ^ (v1 | (v1 - (v1 % ((128) | 1))))), v1, v1}
    var v4 int16 = 42
    v1 = (v1 << ((18446744073709551615) & 63))
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := GenS1{F0: 2}
        var v7 uint8 = 2
        _ = fi5
        _ = v6
        _ = v7
    }
    if (((0 <= 755) || false) || false) {
        if (!true) {
            var v8 uint = 18446744073709551615
            v9 := GenS1{F0: 0}
            _ = v8
            _ = v9
        }
        if len(v2) > 0 {
            v2[0] = 373
        }
    }
    fmt.Println(v4, v1)
    _ = len(v3)
    fmt.Println(v1, v1)
    v10 := GenS1{F0: 1}
    v1 += 2147483647
    if false {
        v11 := GenS1{F0: 2}
        fmt.Println(v4)
        v3 = append(v3, -917)
        _ = v11
    } else {
        var v12 int16 = 32767
        v2 = append(v2, 384)
        switch ("hello" + (("a" + ("" + "a")) + "a")) {
        case "":
            fmt.Println(v4, v12, v12)
        default:
            for fi13 := 0; fi13 < 18; fi13++ {
                if len(v2) > 0 {
                    v2[0] = 536
                }
                _ = fi13
            }
        }
        _ = v12
    }
    v1 += 128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
