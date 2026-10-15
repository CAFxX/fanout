package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int8 = -1
    var v2 uint = 4294967295
    var v3 float64 = 0.1
    v2 += v2
    v4 := []int{19, -620, 4294967295}
    v4 = append(v4, 1)
    if false {
        v3 = 0.0
        if ((4294967295 != 1) || (!(42 <= 0))) {
            if len(v4) > 0 {
                v4[0] = -9223372036854775808
            }
            var v5 float64 = float64((3.14159 / ((-2.25 / 3.14159) * 1.5)))
            v2 -= v2
            _ = v5
        } else {
            fmt.Println(v1)
            var v6 uint64 = 18446744073709551615
            _ = v6
        }
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 uint16 = 32767
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(v2, v2, v1)
            _ = fi7
            _ = v8
        }
    }
    v1 = 65
    var v9 uint64 = 583
    fmt.Println(v9, v3)
    switch "hello" {
    case "a":
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 uint64 = (v9 << (((v2 + v2)) & 63))
            v12 := GenS1{F0: v11}
            _ = fi10
            _ = v11
            _ = v12
        }
    default:
        var v13 int16 = 1
        _ = v13
    }
    v4 = append(v4, 4294967295)
    var v14 uint8 = 1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
