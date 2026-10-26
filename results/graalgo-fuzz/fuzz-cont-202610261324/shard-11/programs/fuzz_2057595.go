package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int8; F1 uint8; F2 int64 }

func main() {
    var v1 string = "\n\t"
    var v2 string = (v1 + v1)
    var v3 uint16 = 65535
    v4 := []string{v1}
    var v5 float32 = (((((((1e10 * 1.5) * -2.25) + 0.0) * float32(float64(float32(0.1)))) * 0.0) - 1e10) + float32(-2.25))
    v6 := make(map[int]int)
    fmt.Println(v3, v5, v1)
    for fi7 := 0; fi7 < 8; fi7++ {
        v3 = v3
        _ = fi7
    }
    fmt.Println(v3, v3)
    v8 := GenS1{F0: false, F1: 677}
    for fi9 := 0; fi9 < 18; fi9++ {
        if false {
            for fi10 := 0; fi10 < 19; fi10++ {
                _ = len(v4)
                var v11 string = (v1 + "a")
                _ = fi10
                _ = v11
            }
            delete(v6, fi9)
            v2 = v2
        }
        if ((251 > 386) || (!true)) {
            if len(v4) > 0 {
                v4[0] = ""
            }
            v12 := GenS2{F0: 1, F1: 2, F2: -322}
            _ = v12
        } else {
            var v13 uint = 2
            var v14 int16 = -790
            _ = v13
            _ = v14
        }
        _ = fi9
    }
    v15 := GenS1{F0: false, F1: -9223372036854775808}
    fmt.Println(v5, v5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v16 uint16 = (v3 << ((470) & 63))
    if ((true && (!((-2147483648 <= -2147483648) && (-2147483648 < 42)))) || false) {
        fmt.Println(v16)
        fmt.Println(v16)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
