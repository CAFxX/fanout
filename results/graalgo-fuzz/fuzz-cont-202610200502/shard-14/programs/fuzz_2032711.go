package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() int16 {
    var v5 string = "\n\t"
    if false {
        var v6 int64 = -1
        _ = v6
    }
    var v7 uint32 = 0
    if false {
        fmt.Println(v5, v5)
    } else {
        for fi8 := 0; fi8 < 6; fi8++ {
            if false {
                var v9 uint32 = (v7 % ((860) | 1))
                fmt.Println(v5)
                _ = v9
            }
            _ = fi8
        }
    }
    var v10 uint = 13
    _ = v10
    _ = v5
    _ = v7
    return 0
}
func genFn1(p0 int32) int32 {
    v11 := GenS1{F0: -1}
    p0 += (p0 * (p0 << ((255) & 63)))
    _ = p0
    _ = v11
    return (p0 >> ((255) & 63))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 9223372036854775807
    var v3 uint8 = 138
    v4 := make(map[string]int)
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 uint32 = 47
        _ = fi12
        _ = v13
    }
    v14 := []int{v2, v2}
    var v15 uint64 = 2
    delete(v1, 837)
    var v16 int = (v2 | 42)
    if ((v16 + v2) < 2147483647) {
        v16 -= (v16 + -2147483648)
        switch 32767 {
        case 1:
            if (uint16(v3) < 42) {
                var v17 float64 = (1.5 + (3.14159 - (0.0 - 0.0)))
                var v18 uint8 = (v3 + v3)
                var v19 float32 = 1.5
                _ = v17
                _ = v18
                _ = v19
            }
        case 127:
            fmt.Println(v3, v16, v15)
        default:
            v20 := GenS1{F0: 32767}
            _ = v20
        }
        v3 += v3
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        v1[101] = "foo bar"
        _ = fi21
    }
    v2 = (^(v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
