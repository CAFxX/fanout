package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int32, p1 int8, p2 string) int32 {
    switch p2 {
    case "":
        p2 = "\n\t"
    case "foo bar":
        fmt.Println(p1, p2)
    default:
        if false {
            var v6 float64 = 1e10
            if (1 >= 42) {
                p0 = (p0 + -413)
            } else {
                p0 -= -95
                fmt.Println(v6, p1)
            }
            switch 2147483647 {
            case 2147483647:
                var v7 string = (p2 + p2)
                _ = v7
            case 239:
                p1 += p1
            }
            _ = v6
        }
    }
    v8 := []int{128, -9223372036854775808, int(p0), 32768}
    var v9 uint16 = 172
    v9 += v9
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return (p0 | (p0 << ((282) & 63)))
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    v3 := []uint32{42}
    var v4 bool = true
    var v5 int = 929
    v10 := GenS1{F0: 983}
    if v4 {
        var v11 uint64 = 4294967295
        _ = len(v1)
        var v12 uint32 = 42
        _ = v11
        _ = v12
    }
    v4 = ((v4 || v4) && true)
    v1[65535] = 1
    var v13 float32 = 1e10
    v14 := GenS1{F0: -2147483648}
    var v15 int16 = 1
    switch ("" + "a") {
    case "a":
        v16 := v2[65535]
        _ = v16
    case "\n\t":
        v17 := GenS1{F0: 0}
        _ = v17
    }
    var v18 string = "a"
    v4 = (2 >= 4294967295)
    v1[42] = 2
    v15 += 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
