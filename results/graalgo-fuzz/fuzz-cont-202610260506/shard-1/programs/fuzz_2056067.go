package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint, p1 uint64) uint32 {
    v4 := []int{128}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v4
    return uint32(p0)
}
func genFn1(p0 int8) uint8 {
    var v5 uint32 = 698
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 uint = 18446744073709551615
        var v8 int16 = 32767
        var v9 int64 = -9223372036854775808
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    fmt.Println(p0, v5)
    var v10 string = ("\n\t" + "a")
    _ = p0
    _ = v10
    _ = v5
    return 128
}

func main() {
    var v1 float64 = 0.0
    var v2 int64 = 1
    v3 := make(map[string]uint32)
    v1 += float64((1e10 * 1.5))
    var v11 float64 = ((-2.25 * v1) * v1)
    v1 -= (v11 - 0.1)
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := GenS1{F0: 433}
        switch fi12 {
        case 32768:
            var v14 uint = 18446744073709551615
            _ = v14
        case 9223372036854775807:
            v11 -= v1
        case 1:
            var v15 uint8 = 42
            _ = v15
        default:
            for fi16 := 0; fi16 < 18; fi16++ {
                v11 = 1e10
                _ = fi16
            }
        }
        v2 = 1
        _ = fi12
        _ = v13
    }
    fmt.Println(v1, v1)
    v11 += 0.1
    var v17 bool = true
    v2 += (v2 * v2)
    v11 = ((v11 / ((v1 - 3.14159) * (float64(0.0) - 1e10))) * float64((-2.25 - (0.1 - (-2.25 - 0.1)))))
    v11 += (v1 / float64(3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
}
