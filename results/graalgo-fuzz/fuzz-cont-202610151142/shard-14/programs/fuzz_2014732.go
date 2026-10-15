package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint, p1 int16) int64 {
    v6 := []int{}
    switch -980 {
    case -2147483648:
        for fi7 := 0; fi7 < 7; fi7++ {
            if (((float32(0.0) - 1.5) - 3.14159) > 3.14159) {
                v6 = append(v6, (fi7 ^ (-(fi7))))
                var v8 string = (("foo bar" + "hello") + (("\n\t" + "\n\t") + "hello"))
                _ = v8
            }
            _ = fi7
        }
    default:
        for fi9 := 0; fi9 < 1; fi9++ {
            for fi10 := 0; fi10 < 8; fi10++ {
                var v11 uint = 194
                v12 := GenS1{F0: 0, F1: (p1 * p1)}
                var v13 string = "foo bar"
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            var v14 int64 = -9223372036854775808
            if len(v6) > 0 {
                v6[0] = 127
            }
            _ = fi9
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return -710
}
func genFn1(p0 int32, p1 int32) uint64 {
    var v15 int8 = 42
    var v16 float32 = -2.25
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return 1
}

func main() {
    var v1 uint16 = 1
    v2 := []uint32{498, 65535}
    var v3 int16 = -32768
    var v4 int64 = 1
    var v5 uint = 18446744073709551615
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v17 int16 = v3
    var v18 uint8 = 2
    switch "hello" {
    case "hello":
        switch -2147483648 {
        case 9223372036854775807:
            v17 = v17
        default:
            var v19 float64 = (0.1 + (3.14159 / 1e10))
            _ = v19
        }
    case "":
        switch 32767 {
        case -270:
            v1 = (^(v1))
        case 32767:
            var v20 int8 = -66
            _ = v20
        case 2147483647:
            var v21 int16 = (v3 << (((v5 + v5)) & 63))
            _ = v21
        }
    }
    fmt.Println(v17, v3, v5)
    v22 := GenS1{F0: 922, F1: int16(v5)}
    v3 += -32768
    var v23 bool = (uint32(v1) >= uint32(v1))
    var v24 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
