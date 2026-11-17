package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
func genFn0() uint64 {
    v6 := GenS1{F0: 127}
    v7 := []int{9223372036854775807}
    var v8 bool = (true || true)
    if (32768 >= 121) {
        v9 := GenS1{F0: 0}
        switch "a" {
        case "foo bar":
            fmt.Println(v8)
        default:
            v10 := GenS1{F0: -128}
            _ = v10
        }
        _ = v9
    }
    var v11 int8 = 100
    _ = v11
    _ = v6
    _ = v7
    _ = v8
    return 18446744073709551615
}
func genFn1(p0 uint64) uint8 {
    fmt.Println(p0)
    if (true && false) {
        var v12 float32 = ((float32(1.5) / 0.1) + (float32((0.1 + 0.1)) * (0.1 - (float32((1e10 - (1.5 + 0.0))) + (0.0 + (1e10 + 1.5))))))
        v12 += (float32((0.1 / float64(v12))) + 1.5)
        _ = v12
    }
    fmt.Println(p0)
    p0 += (p0 * p0)
    _ = p0
    return 82
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    var v3 float32 = 1.5
    var v4 float32 = v3
    v5 := make(map[int]uint32)
    _ = len(v5)
    if false {
        v4 -= v4
        for fi13 := 0; fi13 < 3; fi13++ {
            fmt.Println(v4)
            v4 = (float32(float64(v3)) / (v4 / 1.0))
            _ = fi13
        }
        var v14 float32 = float32((float64(0.1) - 0.0))
        _ = v14
    } else {
        for fi15 := 0; fi15 < 17; fi15++ {
            v16 := []int{395, -9223372036854775808, 26}
            fmt.Println(v3)
            _ = fi15
            _ = v16
        }
        if false {
            if false {
                fmt.Println(v3, v4)
            } else {
                switch 127 {
                case 4294967295:
                    v4 = float32(0.0)
                }
                v17 := GenS1{F0: 127}
                _ = v17
            }
            var v18 int16 = 1000
            _ = v18
        } else {
            v3 += 1e10
            for fi19 := 0; fi19 < 15; fi19++ {
                var v20 int = (fi19 + 1)
                _ = fi19
                _ = v20
            }
            v21 := GenS2{F0: (0 >= 2147483647)}
            _ = v21
        }
    }
    v22 := GenS2{F0: (1.5 != (1e10 / 1e10))}
    v23 := GenS1{F0: -1}
    v4 = -2.25
    fmt.Println(v4, v4)
    v4 = (v3 - (3.14159 / v3))
    for fi24 := 0; fi24 < 2; fi24++ {
        v4 = (3.14159 - float32(-2.25))
        _ = fi24
    }
    fmt.Println(v3, v3)
    switch ("a" + "a") {
    case "hello":
        v3 += v3
    case "":
        v25 := []int{128, 1, 128}
        _ = v25
    }
    if (!(((true || (((!(0 > -128)) && false) || (32767 < 295))) && false) && (!false))) {
        fmt.Println(v3, v3)
        v4 = 3.14159
    }
    var v26 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
