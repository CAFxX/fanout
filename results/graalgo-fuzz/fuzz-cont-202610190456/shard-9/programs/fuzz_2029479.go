package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
func genFn0() bool {
    var v6 uint = 18446744073709551615
    v6 += v6
    v7 := GenS1{F0: 0}
    _ = v6
    _ = v7
    return true
}

func main() {
    v1 := []uint32{348, 2147483647}
    var v2 float64 = 1.5
    var v3 bool = (!(-9223372036854775808 < -2147483648))
    v4 := []uint32{65535, 65535, 4294967295}
    var v5 int64 = 1
    v2 += float64(-2.25)
    v8 := GenS2{F0: 482}
    v3 = true
    if true {
        var v9 float64 = (1e10 - float64(0.1))
        if len(v1) > 0 {
            v1[0] = 7
        }
        var v10 int8 = 1
        _ = v10
        _ = v9
    } else {
        if len(v4) > 0 {
            v4[0] = 2147483648
        }
        var v11 int = 127
        if (42 != 0) {
            v5 = (v5 ^ (-(v5)))
            var v12 uint = 537
            for fi13 := 0; fi13 < 8; fi13++ {
                fmt.Println(v2)
                _ = fi13
            }
            _ = v12
        }
        _ = v11
    }
    var v14 uint8 = 1
    if (v14 > (^(v14))) {
        switch "hello" {
        case "foo bar":
            for fi15 := 0; fi15 < 20; fi15++ {
                _ = len(v1)
                _ = fi15
            }
        case "\n\t":
            var v16 int8 = -128
            _ = v16
        case "a":
            v3 = ((v2 - 1e10) > 1e10)
        default:
            fmt.Println(v3)
        }
    } else {
        _ = len(v1)
    }
    if len(v4) > 0 {
        v4[0] = 978
    }
    v2 += ((1e10 - (0.0 - (0.0 + (v2 * (1.5 * v2))))) + v2)
    _ = len(v4)
    var v17 int16 = -1
    var v18 float32 = (((0.0 * -2.25) * 1.5) - (float32(1.5) / 0.1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
