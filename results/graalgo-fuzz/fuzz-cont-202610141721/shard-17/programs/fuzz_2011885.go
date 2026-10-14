package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 string }
type GenS2 struct { F0 int16; F1 int; F2 uint }
func genFn0(p0 bool, p1 uint16) string {
    v7 := GenS1{F0: -1, F1: 0, F2: "\n\t"}
    switch -2147483648 {
    case 886:
        var v8 float32 = 3.14159
        _ = v8
    case 128:
        var v9 int64 = -9223372036854775808
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return "hello"
}

func main() {
    var v1 string = "a"
    v2 := []string{v1}
    v3 := []int{4294967295, 42, 9223372036854775807}
    var v4 int = 2147483647
    v5 := []uint32{1, 2147483648, 2}
    v6 := make(map[uint32]string)
    var v10 int8 = 9
    delete(v6, 142)
    var v11 uint8 = 2
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 int64 = 1
        switch -95 {
        case -193:
            for fi14 := 0; fi14 < 12; fi14++ {
                v4 += v4
                fmt.Println(fi14, fi12)
                _ = fi14
            }
        case 42:
            var v15 uint16 = 65535
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    var v16 uint = 0
    switch v4 {
    case -1:
        if true {
            fmt.Println(v16, v4, v4)
            var v17 float64 = ((float64((0.0 - -2.25)) / -2.25) / ((0.0 + 0.1) * -2.25))
            switch (v4 / ((v4) | 1)) {
            case 4294967295:
                if true {
                    var v18 uint32 = 989
                    _ = v18
                }
            }
            _ = v17
        }
    default:
        v6[65535] = "a"
    }
    v19 := GenS2{F0: -93, F1: (v4 << (((v16 - v16)) & 63)), F2: 877}
    v16 += v16
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
