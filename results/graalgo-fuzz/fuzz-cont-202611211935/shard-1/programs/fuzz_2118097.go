package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
func genFn0(p0 uint32) uint8 {
    fmt.Println(p0)
    var v4 int32 = 42
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 uint64 = 0
        fmt.Println(v4, v6, v6)
        _ = fi5
        _ = v6
    }
    if ((1.5 + 1.5) >= ((-2.25 / 0.1) * 1.5)) {
        v7 := []int{-907, 748, 2147483647, 0}
        p0 -= (p0 >> ((1) & 63))
        v8 := GenS1{F0: 793, F1: 42}
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 981, F1: -9223372036854775808}
    _ = p0
    _ = v4
    _ = v9
    return 58
}

func main() {
    var v1 uint32 = 2
    var v2 int32 = 42
    var v3 uint64 = 1
    v10 := []int{670}
    switch 1 {
    case 32768:
        if len(v10) > 0 {
            v10[0] = -1
        }
    default:
        if true {
            v3 += 18446744073709551615
        } else {
            if false {
                if true {
                    v11 := GenS1{F0: (v2 * v2), F1: 793}
                    _ = v11
                } else {
                    v10 = append(v10, -433)
                }
            }
        }
    }
    _ = len(v10)
    var v12 uint8 = 128
    v13 := GenS1{F0: (v2 + int32(v3)), F1: -9223372036854775808}
    var v14 int64 = -9223372036854775808
    v3 = 2
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 int = (fi15 / ((fi15) | 1))
        _ = fi15
        _ = v16
    }
    if (0 < v1) {
        fmt.Println(v14, v2)
    } else {
        if (!true) {
            fmt.Println(v14, v2)
            fmt.Println(v12)
        } else {
            switch 9223372036854775807 {
            case 747:
                var v17 int64 = 0
                _ = v17
            default:
                _ = len(v10)
            }
            v1 -= (v1 - 760)
            v10 = append(v10, -1)
        }
        v18 := GenS1{F0: -2147483648, F1: (-(v14))}
        var v19 uint64 = uint64(v12)
        _ = v18
        _ = v19
    }
    if (!(false || false)) {
        if len(v10) > 0 {
            v10[0] = 128
        }
        _ = len(v10)
    }
    var v20 uint16 = uint16(v3)
    var v21 uint32 = 711
    v2 -= (v2 & int32(v14))
    var v22 string = ("\n\t" + "hello")
    var v23 uint32 = v1
    v22 = "a"
    v24 := GenS1{F0: v2, F1: v14}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
