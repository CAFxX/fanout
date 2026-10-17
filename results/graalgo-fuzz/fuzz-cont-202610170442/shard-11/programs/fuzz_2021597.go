package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint64, p1 int8) int64 {
    p0 = (p0 / (((p0 - p0)) | 1))
    v7 := []int{-2147483648, 2147483647}
    p0 = (p0 + p0)
    switch "\n\t" {
    case "foo bar":
        p1 -= 100
    default:
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 196
}

func main() {
    var v1 uint16 = 0
    var v2 int32 = 2147483647
    var v3 int16 = 1
    var v4 int64 = 9223372036854775807
    v5 := []string{}
    v6 := make(map[uint32]uint32)
    var v8 int64 = v4
    var v9 float32 = -2.25
    v1 += 255
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 uint64 = 932
        _ = len(v5)
        for fi12 := 0; fi12 < 10; fi12++ {
            v1 += 2
            var v13 uint8 = 127
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    v14 := v6[42]
    v6[uint32(v1)] = (v14 | 1)
    var v15 int64 = (v8 << ((18446744073709551615) & 63))
    switch -9223372036854775808 {
    case -599:
        v4 -= 42
    case 4294967295:
        delete(v6, (^(v14)))
    case -356:
        if (127 > 100) {
            v16 := GenS1{F0: (v2 + (^(v2)))}
            _ = v16
        }
    default:
        if len(v5) > 0 {
            v5[0] = ("\n\t" + "\n\t")
        }
    }
    var v17 float64 = ((0.0 + -2.25) + float64(v9))
    v17 = (3.14159 - v17)
    v9 = 0.0
    v9 += 1.5
    v3 = int16(v2)
    var v18 uint64 = 0
    delete(v6, (v14 + v14))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
