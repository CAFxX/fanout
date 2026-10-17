package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 bool }
type GenS2 struct { F0 int8; F1 uint }
func genFn0(p0 int8, p1 int64, p2 int32) uint64 {
    var v7 float64 = -2.25
    p0 = int8(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 int64, p1 uint) int {
    if false {
        v8 := []int{}
        if (!((2147483647 >= -802) && (!(!true)))) {
            for fi9 := 0; fi9 < 10; fi9++ {
                v10 := GenS1{F0: 1, F1: 1000, F2: (-1 == 943)}
                var v11 int8 = -128
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        _ = v8
    } else {
        var v12 uint32 = 15
        v13 := []int{-9223372036854775808}
        _ = v12
        _ = v13
    }
    p1 = (^(p1))
    v14 := []int{}
    _ = p0
    _ = p1
    _ = v14
    return -2147483648
}

func main() {
    var v1 int8 = -1
    v2 := make(map[string]uint32)
    var v3 uint8 = 0
    var v4 bool = (42 == 42)
    var v5 float64 = 0.0
    v6 := []int{}
    switch "" {
    case "hello":
        switch 9223372036854775807 {
        case 2147483647:
            v15 := GenS1{F0: 42, F1: 1, F2: v4}
            _ = v15
        default:
            if false {
                if v4 {
                    var v16 uint16 = 825
                    _ = v16
                } else {
                    fmt.Println(v4, v3)
                    switch "a" {
                    case "":
                        var v17 uint8 = v3
                        _ = v17
                    }
                    v18 := GenS1{F0: 1, F1: 32767, F2: false}
                    _ = v18
                }
                switch "a" {
                case "a":
                    for fi19 := 0; fi19 < 17; fi19++ {
                        v20 := GenS1{F0: -477, F1: 1000, F2: v4}
                        v4 = (true || (-998 == -882))
                        _ = fi19
                        _ = v20
                    }
                default:
                    v21 := GenS1{F0: -2147483648, F1: -32768, F2: v4}
                    _ = v21
                }
                var v22 int32 = -471
                _ = v22
            } else {
                v6 = append(v6, 0)
            }
        }
    case "a":
        v1 -= v1
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        v24 := v2["hello"]
        v3 = 0
        _ = fi23
        _ = v24
    }
    switch 42 {
    case 32768:
        switch ("foo bar" + (("" + "\n\t") + "\n\t")) {
        case "a":
            v3 = (v3 & v3)
        }
    default:
        if v4 {
            v2["a"] = 2147483648
            for fi25 := 0; fi25 < 5; fi25++ {
                v3 = (^(v3))
                v6 = append(v6, (fi25 >> ((42) & 63)))
                _ = fi25
            }
            v1 += v1
        }
    }
    v26 := GenS2{F0: (^(v1)), F1: 18446744073709551615}
    fmt.Println(v1, v3)
    if len(v6) > 0 {
        v6[0] = genFn1(254, 65535)
    }
    var v27 int32 = -228
    for fi28 := 0; fi28 < 10; fi28++ {
        v4 = (false && ((v3 % ((128) | 1)) < 190))
        _ = fi28
    }
    if (v4 || (255 == 255)) {
        var v29 int8 = (v1 ^ (v1 ^ v1))
        var v30 int16 = -281
        _ = v29
        _ = v30
    } else {
        var v31 int = -274
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
