package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int }
func genFn0(p0 uint64, p1 int) uint {
    p1 += 128
    v5 := GenS1{F0: -9223372036854775808, F1: 44, F2: p1}
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1(p0 uint64) int16 {
    v6 := []int{}
    p0 = (p0 & p0)
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 uint64 = (^(p0))
        _ = fi7
        _ = v8
    }
    v6 = append(v6, -337)
    _ = p0
    _ = v6
    return 32767
}

func main() {
    var v1 int8 = 127
    var v2 uint16 = 42
    var v3 uint16 = 2
    var v4 int64 = -609
    var v9 uint16 = 65535
    var v10 int32 = -87
    if false {
        v2 -= 2
    }
    fmt.Println(v10, v9, v9)
    if (!false) {
        for fi11 := 0; fi11 < 16; fi11++ {
            if (!false) {
                switch fi11 {
                case -862:
                    fmt.Println(v10, v4, v2)
                }
            } else {
                v12 := []int{(fi11 - fi11), int(v4), fi11}
                for fi13 := 0; fi13 < 16; fi13++ {
                    v12 = append(v12, fi13)
                    _ = fi13
                }
                v14 := GenS1{F0: (-(v4)), F1: 65535, F2: -1}
                _ = v12
                _ = v14
            }
            _ = fi11
        }
        var v15 float32 = float32((3.14159 - 1e10))
        switch "" {
        case "hello":
            if false {
                var v16 int64 = 9223372036854775807
                v15 -= ((v15 + (((1.5 * v15) / v15) + v15)) / float32(3.14159))
                var v17 int8 = int8(v4)
                _ = v16
                _ = v17
            }
        case "foo bar":
            v1 += v1
        case "":
            v18 := []int{}
            _ = v18
        default:
            switch 0 {
            case 1:
                v19 := []int{9223372036854775807, 572, -559, 0}
                _ = v19
            }
        }
        _ = v15
    }
    v10 -= 2147483647
    var v20 bool = (!false)
    v4 = -1
    v3 -= 2
    v9 -= 65535
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
