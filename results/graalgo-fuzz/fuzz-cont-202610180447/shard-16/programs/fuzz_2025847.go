package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int32; F1 int; F2 uint32 }
type GenS3 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint32, p1 string, p2 int64) uint16 {
    var v6 int8 = 1
    if (false || (p2 == (p2 ^ p2))) {
        v7 := []int{}
        _ = len(v7)
        if true {
            if (42 < 41) {
                v8 := GenS1{F0: true, F1: (v6 % ((v6) | 1))}
                var v9 bool = (!(-70 == -315))
                _ = v8
                _ = v9
            } else {
                switch 243 {
                case 128:
                    if true {
                        p2 -= p2
                    } else {
                        fmt.Println(v6, v6, v6)
                        fmt.Println(p1, v6)
                    }
                default:
                    var v10 int8 = v6
                    _ = v10
                }
            }
            if len(v7) > 0 {
                v7[0] = 2147483647
            }
        }
        _ = v7
    } else {
        if false {
            var v11 int8 = v6
            fmt.Println(v6, p1, p0)
            p1 = p1
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}
func genFn1(p0 uint64) uint {
    v12 := []int{127, -1, 42, 0}
    var v13 uint32 = 721
    _ = len(v12)
    fmt.Println(p0)
    _ = p0
    _ = v12
    _ = v13
    return 18446744073709551615
}

func main() {
    v1 := []uint32{419}
    var v2 string = "hello"
    var v3 int = 9223372036854775807
    v4 := []string{v2, (v2 + ""), "\n\t", v2}
    v5 := make(map[uint32]int)
    v14 := GenS2{F0: 1, F1: v3, F2: 49}
    switch v2 {
    case "a":
        fmt.Println(v2, v2)
    case "":
        var v15 int = (-(v3))
        _ = v15
    case "\n\t":
        v16 := GenS3{F0: 759, F1: 890}
        _ = v16
    }
    if len(v4) > 0 {
        v4[0] = v2
    }
    v2 = v2
    v2 = "foo bar"
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 uint32 = 2147483647
        v19 := GenS2{F0: 42, F1: (fi17 | fi17), F2: v18}
        v3 -= 400
        _ = fi17
        _ = v18
        _ = v19
    }
    switch v3 {
    case -9223372036854775808:
        v2 = v2
    case 974:
        v20 := v5[1]
        _ = v20
    }
    fmt.Println(v2)
    fmt.Println(v2)
    if false {
        var v21 uint8 = 255
        for fi22 := 0; fi22 < 10; fi22++ {
            v2 = "a"
            _ = fi22
        }
        _ = v21
    }
    var v23 float32 = (0.0 + 1.5)
    v3 = (-(v3))
    v3 -= (v3 ^ (v3 & (v3 & (-(v3)))))
    var v24 uint8 = 114
    v23 = (float32(1e10) * v23)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
