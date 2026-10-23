package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint8; F1 int }

func main() {
    v1 := []string{"a", (("a" + "hello") + ("" + "hello"))}
    var v2 float32 = float32((float64(-2.25) / float64((float32((-2.25 + (0.1 - (float64((0.0 - 0.1)) + (1e10 * 1.5))))) * -2.25))))
    var v3 int32 = 2147483647
    var v4 int = 9223372036854775807
    v5 := make(map[string]int)
    var v6 uint8 = 179
    var v7 uint64 = 18446744073709551615
    var v8 bool = (0 == 1)
    if (v4 != 32767) {
        if ((v2 / (float32(((1e10 / 1.0) + 1.5)) / (1.5 * v2))) == (((((3.14159 / v2) / (v2 / (v2 / v2))) + 1e10) + (v2 - (v2 + 1e10))) / 1e10)) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            v8 = (42 > 32767)
            switch "a" {
            case "foo bar":
                var v9 uint32 = 2147483647
                _ = v9
            }
        }
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        v2 += ((float32(1e10) - v2) - (v2 - 3.14159))
        var v11 int32 = (v3 & v3)
        _ = fi10
        _ = v11
    }
    var v12 int64 = -9223372036854775808
    v13 := GenS1{F0: 42, F1: 255, F2: -1}
    v14 := GenS1{F0: 255, F1: v6, F2: 42}
    if (2 > 4294967295) {
        v15 := GenS1{F0: 969, F1: (v6 - (v6 + v6)), F2: 61}
        fmt.Println(v4, v2, v4)
        _ = v15
    }
    var v16 int = v4
    if v8 {
        v16 -= 32768
        if len(v1) > 0 {
            v1[0] = "a"
        }
    }
    var v17 uint16 = 37
    v18 := v5[""]
    v16 = (-(v16))
    v19 := GenS2{F0: (v6 | (v6 ^ v6)), F1: (-(v4))}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
