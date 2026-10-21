package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 string }
type GenS2 struct { F0 int8; F1 uint64 }

func main() {
    var v1 int32 = 0
    var v2 int32 = (-(v1))
    var v3 uint16 = 982
    var v4 uint32 = 953
    v5 := GenS1{F0: true, F1: 0, F2: "\n\t"}
    var v6 uint = 1
    v2 = (v1 % ((261) | 1))
    fmt.Println(v6, v4, v1)
    v7 := GenS2{F0: 127, F1: 18446744073709551615}
    v8 := []int{}
    v8 = append(v8, -9223372036854775808)
    v1 += (v2 & 753)
    v8 = append(v8, -851)
    switch 32768 {
    case 32768:
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 uint = 0
            _ = fi9
            _ = v10
        }
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    switch ("foo bar" + "\n\t") {
    case "a":
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 uint32 = v4
            _ = fi11
            _ = v12
        }
    case "hello":
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 uint64 = 0
            _ = fi13
            _ = v14
        }
    }
    var v15 uint8 = 255
    if (v1 > 2147483647) {
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 int = 9223372036854775807
            if (float64(0.1) >= ((0.0 - 1.5) / (3.14159 / -2.25))) {
                var v18 int64 = 0
                var v19 uint = v6
                fmt.Println(v1, v2, v18)
                _ = v18
                _ = v19
            }
            var v20 float32 = ((0.0 + -2.25) + 3.14159)
            _ = fi16
            _ = v17
            _ = v20
        }
        var v21 int32 = v2
        _ = v21
    } else {
        v22 := GenS2{F0: 100, F1: 2}
        for fi23 := 0; fi23 < 13; fi23++ {
            var v24 int = fi23
            v4 += (v4 * v4)
            _ = fi23
            _ = v24
        }
        var v25 bool = false
        _ = v22
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
