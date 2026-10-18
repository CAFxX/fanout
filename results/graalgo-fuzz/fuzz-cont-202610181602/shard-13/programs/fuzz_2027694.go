package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 float64 = float64(1e10)
    var v2 bool = ((-9223372036854775808 < 0) && false)
    var v3 float64 = v1
    v4 := make(map[uint32]int)
    var v5 int8 = 1
    fmt.Println(v5, v1)
    if v2 {
        v6 := v4[278]
        v7 := GenS1{F0: 402, F1: 4294967295}
        _ = v6
        _ = v7
    } else {
        switch "" {
        case "\n\t":
            if v2 {
                v2 = v2
                fmt.Println(v1, v3)
            } else {
                var v8 string = ("" + "a")
                v1 += 0.1
                _ = v8
            }
        case "hello":
            var v9 uint64 = 18446744073709551615
            _ = v9
        case "a":
            for fi10 := 0; fi10 < 8; fi10++ {
                v3 += (((float64(((-2.25 / float32((float64(-2.25) - (v3 / (v3 + v3))))) * float32(v3))) / (0.0 + 1.5)) - (1e10 - v1)) - ((v3 * ((3.14159 - v1) * v3)) * v3))
                if (275 >= 2) {
                    v11 := GenS1{F0: 0, F1: 1}
                    var v12 int64 = 0
                    _ = v11
                    _ = v12
                }
                if (!(!(!false))) {
                    v3 = v3
                    fmt.Println(fi10, v3, fi10)
                    var v13 uint16 = 255
                    _ = v13
                }
                _ = fi10
            }
        }
    }
    v5 += v5
    fmt.Println(v5, v2, v2)
    var v14 float32 = float32(v1)
    fmt.Println(v3)
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 uint16 = 65535
        _ = fi15
        _ = v16
    }
    var v17 uint16 = 956
    for fi18 := 0; fi18 < 10; fi18++ {
        if v2 {
            v4[2147483647] = (fi18 >> ((2) & 63))
            if v2 {
                v19 := GenS1{F0: v17, F1: 4294967295}
                _ = v19
            }
            fmt.Println(v5, v5, v2)
        } else {
            var v20 int16 = 1000
            fmt.Println(v1, v1, v5)
            fmt.Println(v20, v1)
            _ = v20
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
