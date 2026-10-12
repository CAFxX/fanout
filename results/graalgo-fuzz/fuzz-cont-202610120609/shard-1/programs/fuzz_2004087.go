package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 42
    var v3 int16 = -32768
    v4 := []int{32767}
    if ((!(-532 > 0)) || true) {
        delete(v1, 261)
    }
    v5 := GenS1{F0: 255, F1: 51, F2: 132}
    v6 := GenS1{F0: 42, F1: 1, F2: 127}
    for fi7 := 0; fi7 < 20; fi7++ {
        switch 9223372036854775807 {
        case 2147483647:
            delete(v1, 2147483647)
        case -2147483648:
            var v8 int8 = 63
            _ = v8
        default:
            if len(v4) > 0 {
                v4[0] = (v2 & (-(fi7)))
            }
        }
        for fi9 := 0; fi9 < 1; fi9++ {
            if false {
                if len(v4) > 0 {
                    v4[0] = fi9
                }
                v4 = append(v4, fi9)
                delete(v1, 1)
            }
            _ = fi9
        }
        _ = fi7
    }
    v2 += (v2 << ((652) & 63))
    v2 = (v2 % (((v2 | 32768)) | 1))
    v10 := v1[1]
    if false {
        var v11 uint16 = 1
        _ = v11
    }
    v2 += v2
    fmt.Println(v10, v10, v3)
    if (42 == 334) {
        var v12 uint8 = 42
        _ = v12
    } else {
        var v13 string = "hello"
        _ = v13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
