package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
func genFn0(p0 uint8) uint8 {
    for fi6 := 0; fi6 < 16; fi6++ {
        if false {
            var v7 uint64 = 18446744073709551615
            _ = v7
        }
        _ = fi6
    }
    fmt.Println(p0)
    p0 += (^(p0))
    _ = p0
    return p0
}
func genFn1(p0 string, p1 int8) string {
    switch -9223372036854775808 {
    case 2147483647:
        if (false || (0 != 1)) {
            var v8 uint32 = 0
            _ = v8
        }
    default:
        var v9 bool = (!true)
        _ = v9
    }
    p1 += 0
    v10 := []int{938, 976, 2147483647, 32767}
    for fi11 := 0; fi11 < 13; fi11++ {
        if true {
            if false {
                var v12 uint8 = genFn0(10)
                var v13 uint16 = 65535
                _ = v12
                _ = v13
            }
        } else {
            if len(v10) > 0 {
                v10[0] = fi11
            }
        }
        p1 -= (^(p1))
        var v14 int8 = p1
        _ = fi11
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v10
    return (p0 + (p0 + p0))
}

func main() {
    var v1 uint = 2
    var v2 uint = v1
    var v3 int64 = 815
    var v4 string = "a"
    var v5 uint32 = 454
    var v15 int16 = -976
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 uint32 = (v5 << (((v1 ^ 779)) & 63))
        _ = fi16
        _ = v17
    }
    switch v4 {
    case "":
        v3 = v3
    default:
        v3 = -1
    }
    if true {
        fmt.Println(v1)
        v4 = v4
    }
    var v18 uint64 = 42
    switch v4 {
    case "hello":
        v15 = v15
    default:
        for fi19 := 0; fi19 < 6; fi19++ {
            v1 = (v2 >> (((^(v2))) & 63))
            v1 -= v2
            v18 += 993
            _ = fi19
        }
    }
    var v20 uint8 = 2
    v21 := GenS1{F0: 1, F1: v4}
    var v22 float64 = 1.5
    v23 := []int{2147483647, 606}
    v2 = (^(v1))
    v15 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
