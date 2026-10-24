package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint64 }

func main() {
    var v1 int8 = 1
    var v2 uint64 = 2
    var v3 float64 = 1.5
    v4 := make(map[int]string)
    var v5 float32 = 1.5
    v4[32768] = ((("a" + "a") + "a") + "")
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int32 = 1
        fmt.Println(v1, fi6)
        _ = fi6
        _ = v7
    }
    v3 += ((v3 / v3) - 1.5)
    v5 = float32(v3)
    v2 += (v2 ^ 42)
    v1 = v1
    fmt.Println(v3, v1)
    var v8 int8 = v1
    if ((1.5 / (v5 * v5)) < v5) {
        v9 := GenS1{F0: 65535, F1: (-(v8)), F2: (v2 % (((v2 + (^(v2)))) | 1))}
        v10 := GenS1{F0: 575, F1: v8, F2: (v2 + 825)}
        for fi11 := 0; fi11 < 7; fi11++ {
            v12 := GenS1{F0: 255, F1: (v8 & (v1 << ((255) & 63))), F2: (v2 >> ((229) & 63))}
            if (431 >= 2) {
                v1 += v8
                var v13 string = ""
                _ = v13
            }
            var v14 uint64 = (v2 ^ 407)
            _ = fi11
            _ = v12
            _ = v14
        }
        _ = v10
        _ = v9
    } else {
        var v15 int16 = 0
        var v16 uint64 = 2
        v1 = (v8 >> ((901) & 63))
        _ = v15
        _ = v16
    }
    delete(v4, 2147483647)
    v8 = -85
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
