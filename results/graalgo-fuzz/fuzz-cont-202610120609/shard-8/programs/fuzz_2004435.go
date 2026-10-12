package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 bool; F1 bool; F2 int32 }
func genFn0() uint16 {
    var v6 uint16 = 42
    v6 = v6
    _ = v6
    return v6
}
func genFn1(p0 int, p1 string) int {
    switch 128 {
    case -2147483648:
        p1 = "a"
    case -1:
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 int8 = 0
            _ = fi7
            _ = v8
        }
    }
    v9 := []int{(^(p0)), (p0 - (p0 >> ((0) & 63))), (p0 & p0)}
    var v10 uint8 = 127
    var v11 float32 = 0.0
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return p0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]bool)
    var v3 int = 32768
    var v4 uint64 = 311
    v5 := make(map[int]uint32)
    switch "" {
    case "foo bar":
        var v12 uint32 = 2147483648
        _ = v12
    case "a":
        for fi13 := 0; fi13 < 2; fi13++ {
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 uint = 0
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
    case "hello":
        v3 = (v3 & v3)
    default:
        var v16 uint64 = v4
        _ = v16
    }
    v17 := GenS1{F0: 4294967295, F1: 255}
    var v18 bool = true
    v4 = (v4 - v4)
    var v19 bool = (v18 || v18)
    if v18 {
        v20 := GenS2{F0: 2147483647, F1: (-32768 != -616)}
        v21 := GenS3{F0: v3, F1: 1}
        switch 2147483647 {
        case 532:
            v19 = v18
        default:
            v22 := []int{-1, 32767, 127, v3}
            _ = v22
        }
        _ = v20
        _ = v21
    }
    if (v18 && (-128 < -42)) {
        switch v3 {
        case -1:
            for fi23 := 0; fi23 < 18; fi23++ {
                for fi24 := 0; fi24 < 12; fi24++ {
                    var v25 float32 = (((((3.14159 / (3.14159 / float32(1.5))) * 3.14159) / (3.14159 - (0.1 + 1.5))) + 0.0) - (1e10 - 0.1))
                    v26 := GenS3{F0: fi23, F1: 9223372036854775807}
                    var v27 float32 = (v25 / ((3.14159 - v25) - 1e10))
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                _ = fi23
            }
        case 501:
            v2[v3] = true
        }
        _ = len(v1)
    } else {
        v19 = (v3 != (v3 % ((genFn1(v3, "a")) | 1)))
    }
    var v28 int64 = 42
    v29 := GenS4{F0: (-2147483648 < (v3 / ((2147483647) | 1))), F1: (!(619 > 610)), F2: -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
