package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int32; F1 bool; F2 int16 }
func genFn0(p0 uint8, p1 bool, p2 string) uint64 {
    fmt.Println(p2, p0, p1)
    switch "\n\t" {
    case "":
        for fi4 := 0; fi4 < 15; fi4++ {
            if false {
                p1 = (!false)
            } else {
                p0 = 1
                v5 := []int{2147483647}
                if p1 {
                    p2 = p2
                } else {
                    var v6 uint16 = 255
                    _ = v6
                }
                _ = v5
            }
            _ = fi4
        }
    }
    var v7 string = "\n\t"
    fmt.Println(v7)
    p1 = (42 == 0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 uint16) uint32 {
    v8 := []int{499, -2147483648, -278}
    switch 2147483647 {
    case -9223372036854775808:
        var v9 int = -206
        _ = v9
    }
    var v10 uint16 = (p0 * p0)
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 uint16 = (^(v10))
        v10 = 32767
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 string = "foo bar"
        v14 = (v14 + v14)
        if (false && (true && true)) {
            v15 := GenS1{F0: -1, F1: (-(fi13))}
            if (0.0 >= -2.25) {
                fmt.Println(v14, fi13)
            }
            fmt.Println(v14, fi13, v10)
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = v10
    _ = v8
    return 1
}

func main() {
    v1 := []int{1, 32767, 2147483647, 0}
    var v2 int32 = 42
    var v3 string = "foo bar"
    fmt.Println(v3, v2)
    v2 = 1
    var v16 uint16 = 2
    switch 127 {
    case 128:
        fmt.Println(v16, v2, v2)
    }
    var v17 bool = (true || true)
    v18 := GenS2{F0: int32(v16), F1: v17, F2: 1000}
    if v17 {
        var v19 int32 = -2147483648
        v1 = append(v1, 1)
        _ = v19
    } else {
        var v20 float32 = (3.14159 - -2.25)
        _ = v20
    }
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = 0
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
