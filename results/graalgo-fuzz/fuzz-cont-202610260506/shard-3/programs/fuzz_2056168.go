package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint16 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint16 = 324
    v3 := make(map[int]bool)
    var v4 float32 = 0.0
    var v5 int64 = -1
    v4 += (float32(1.5) - v4)
    var v6 int32 = -1
    switch 9223372036854775807 {
    case 86:
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 int = (fi7 ^ fi7)
            var v9 string = "\n\t"
            var v10 uint = v1
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
    default:
        fmt.Println(v1, v2, v5)
    }
    v1 = v1
    v11 := []int{32767}
    var v12 int32 = (v6 * (v6 & 1))
    if true {
        for fi13 := 0; fi13 < 2; fi13++ {
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 int16 = 0
                _ = fi14
                _ = v15
            }
            var v16 int64 = -1
            _ = fi13
            _ = v16
        }
        var v17 int8 = 127
        _ = v17
    } else {
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 float32 = v4
            _ = fi18
            _ = v19
        }
        v20 := GenS1{F0: 220, F1: 0, F2: 971}
        if ((((("hello" + ("" + "\n\t")) + "hello") + "\n\t") + ("\n\t" + "a")) >= "foo bar") {
            var v21 int64 = v5
            var v22 uint32 = 1
            v11 = append(v11, 9223372036854775807)
            _ = v21
            _ = v22
        }
        _ = v20
    }
    var v23 int8 = 1
    v23 = 127
    fmt.Println(v5, v1)
    v3[-2147483648] = (!false)
    v1 += (v1 + v1)
    v1 = (v1 >> ((v1) & 63))
    v24 := GenS2{F0: 882, F1: (v2 ^ 0)}
    var v25 uint32 = 65535
    fmt.Println(v2, v4, v23)
    fmt.Println(v5, v6, v4)
    if (v1 >= v1) {
        if (0 >= -56) {
            _ = len(v3)
            if ((!false) && true) {
                fmt.Println(v4, v23, v23)
            } else {
                var v26 float64 = (((((1e10 + 3.14159) + ((0.0 * 0.1) - (1.5 + 1.5))) * 0.0) / 1e10) - (0.1 - 1e10))
                if false {
                    var v27 uint64 = 4294967295
                    _ = v27
                }
                _ = v26
            }
            v4 = v4
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
