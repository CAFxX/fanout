package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint64, p1 int16) int64 {
    var v7 int16 = (p1 + p1)
    switch "foo bar" {
    case "hello":
        v7 += -936
    case "\n\t":
        v8 := GenS1{F0: 149}
        _ = v8
    default:
        v9 := []int{128, 128}
        _ = v9
    }
    var v10 uint = 1
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint32 = 778
    v3 := []uint32{}
    var v4 uint = 255
    v5 := []int{}
    var v6 int64 = 42
    var v11 string = "hello"
    if len(v3) > 0 {
        v3[0] = (v2 | (v2 >> ((18446744073709551615) & 63)))
    }
    _ = len(v5)
    v11 = ("foo bar" + (v11 + (v11 + ((v11 + "\n\t") + v11))))
    v11 = ""
    var v12 int8 = 42
    var v13 string = v11
    switch v11 {
    case "\n\t":
        v6 -= v6
    default:
        var v14 int32 = int32(v1)
        _ = v14
    }
    if ((-2.25 + float64((0.1 + (float32(1.5) / ((float32(float64(0.1)) - -2.25) / (1e10 - 3.14159)))))) > (float64(0.0) - float64((0.0 / 3.14159)))) {
        v2 = v2
        fmt.Println(v2)
    } else {
        var v15 int64 = -1
        var v16 uint16 = 1
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 11}
    v4 = v4
    for fi18 := 0; fi18 < 18; fi18++ {
        if false {
            var v19 float64 = 1.5
            _ = v19
        }
        fmt.Println(v11)
        if true {
            v6 = 0
        } else {
            v2 += (v2 & (v2 & (^(v2))))
            v3 = append(v3, uint32(v4))
            if (!true) {
                var v20 uint64 = (v1 | 0)
                _ = v20
            }
        }
        _ = fi18
    }
    v12 -= (-(v12))
    v2 += (v2 << (((v4 << (((^(v4))) & 63))) & 63))
    if (v11 > (("a" + "\n\t") + v13)) {
        var v21 uint32 = (v2 & 2147483647)
        _ = v21
    } else {
        fmt.Println(v4, v13)
    }
    var v22 uint32 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
