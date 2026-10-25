package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint32 }
func genFn0(p0 int8, p1 uint16) int64 {
    fmt.Println(p0, p1)
    fmt.Println(p1, p1)
    v5 := []int{}
    p0 -= p0
    switch "" {
    case "foo bar":
        if true {
            var v6 bool = true
            v6 = v6
            _ = v6
        }
    default:
        var v7 uint8 = 127
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    var v1 uint32 = 4294967295
    v2 := []string{"foo bar", "a"}
    var v3 uint16 = 65535
    v4 := make(map[string]int)
    v3 += v3
    fmt.Println(v1)
    v1 = (v1 << ((255) & 63))
    if ((((3.14159 - float64(float32(1.5))) - (1e10 + 1e10)) * -2.25) < 0.0) {
        var v8 uint16 = (v3 | (v3 ^ v3))
        if (1000 > 384) {
            switch 299 {
            case 0:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case -598:
                v9 := GenS1{F0: 1, F1: 2147483647, F2: v1}
                _ = v9
            }
        }
        v3 -= 704
        _ = v8
    }
    var v10 uint32 = v1
    if false {
        if (-1 != -219) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v2 = append(v2, "a")
        } else {
            v3 += uint16(v10)
        }
        fmt.Println(v10)
        var v11 int16 = 32767
        _ = v11
    }
    v4["hello"] = 203
    var v12 float64 = 0.1
    fmt.Println(v1, v12)
    v12 = float64(3.14159)
    v13 := GenS1{F0: 127, F1: 42, F2: 293}
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
