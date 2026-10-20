package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 string }
type GenS2 struct { F0 uint; F1 uint16; F2 int8 }
func genFn0(p0 uint, p1 int32, p2 uint8) string {
    switch 9223372036854775807 {
    case 4294967295:
        var v7 uint16 = 42
        _ = v7
    default:
        for fi8 := 0; fi8 < 13; fi8++ {
            if (42 > (p0 << (((p0 * 4294967295)) & 63))) {
                switch (fi8 + (-(fi8))) {
                case -287:
                    fmt.Println(p1, p0, p2)
                case -9223372036854775808:
                    var v9 int8 = -2
                    _ = v9
                case 32768:
                    p0 = (p0 >> ((42) & 63))
                }
                v10 := []int{}
                _ = v10
            } else {
                v11 := []int{}
                _ = v11
            }
            v12 := []int{}
            _ = fi8
            _ = v12
        }
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    return ("" + ("a" + (("foo bar" + "\n\t") + "a")))
}

func main() {
    var v1 uint32 = 2147483648
    var v2 int16 = -32768
    var v3 float64 = 1e10
    v4 := make(map[string]bool)
    v5 := []uint32{v1, (v1 & v1), (^(v1))}
    var v6 bool = true
    v13 := GenS1{F0: v2, F1: (0 <= 787), F2: "\n\t"}
    fmt.Println(v1)
    var v14 uint = 255
    v15 := GenS2{F0: v14, F1: 53, F2: 1}
    _ = len(v5)
    var v16 float64 = float64(3.14159)
    v17 := GenS2{F0: (^(v14)), F1: uint16(v1), F2: -128}
    v6 = v6
    v2 = v2
    var v18 int32 = 42
    var v19 string = ("" + "\n\t")
    var v20 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
