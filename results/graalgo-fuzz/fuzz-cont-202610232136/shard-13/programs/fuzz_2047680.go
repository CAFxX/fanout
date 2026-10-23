package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint }
type GenS2 struct { F0 uint8; F1 uint8; F2 bool }
func genFn0(p0 int32, p1 string, p2 int16) bool {
    switch 4294967295 {
    case -56:
        if (2147483647 < p0) {
            v5 := GenS1{F0: 0, F1: p1, F2: 255}
            _ = v5
        } else {
            switch 1 {
            case 2147483647:
                v6 := []int{-1, -2147483648, 128}
                _ = v6
            }
            for fi7 := 0; fi7 < 3; fi7++ {
                p0 -= 42
                v8 := GenS2{F0: 51, F1: 0, F2: true}
                var v9 uint64 = 0
                _ = fi7
                _ = v8
                _ = v9
            }
            for fi10 := 0; fi10 < 6; fi10++ {
                p1 = p1
                _ = fi10
            }
        }
    case 42:
        var v11 uint = 1
        _ = v11
    }
    var v12 uint16 = 32767
    p0 = (p0 | (p0 << ((42) & 63)))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return (false || true)
}

func main() {
    v1 := []int{127}
    var v2 uint = 18446744073709551615
    var v3 int64 = -1
    v4 := []uint32{2147483647, 528, 4294967295}
    if true {
        if false {
            if len(v1) > 0 {
                v1[0] = 127
            }
            var v13 uint64 = 4294967295
            _ = v13
        } else {
            for fi14 := 0; fi14 < 9; fi14++ {
                switch "foo bar" {
                case "\n\t":
                    v3 -= (v3 << ((18446744073709551615) & 63))
                default:
                    fmt.Println(fi14, v3, v3)
                }
                _ = fi14
            }
            fmt.Println(v3, v3)
            var v15 int64 = v3
            _ = v15
        }
        var v16 float64 = (1.5 + 3.14159)
        v2 += 42
        _ = v16
    }
    v1 = append(v1, 4294967295)
    var v17 int64 = (v3 | v3)
    if (true && (255 == 121)) {
        _ = len(v1)
        var v18 int64 = (v3 & v17)
        _ = v18
    }
    var v19 uint64 = 18446744073709551615
    var v20 int64 = (v3 | 0)
    var v21 int32 = -2147483648
    v21 = (-(v21))
    v3 += 102
    switch "a" {
    case "":
        if (42 > 1) {
            v22 := GenS2{F0: 128, F1: 100, F2: (3.14159 > 1e10)}
            _ = v22
        } else {
            var v23 uint64 = 42
            _ = v23
        }
    }
    v2 += (v2 - v2)
    var v24 int8 = 127
    v25 := GenS1{F0: 65535, F1: "a", F2: (^(v2))}
    var v26 int = 128
    fmt.Println(v2, v26, v24)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
