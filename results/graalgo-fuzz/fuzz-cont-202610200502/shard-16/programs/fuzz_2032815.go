package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int32 }
func genFn0(p0 int8, p1 int8) int64 {
    var v7 uint16 = 1
    v8 := []int{}
    var v9 uint16 = (v7 + (v7 % ((v7) | 1)))
    v10 := GenS1{F0: "", F1: 775}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -254
}
func genFn1() uint64 {
    var v11 int32 = -2147483648
    fmt.Println(v11)
    var v12 int64 = 42
    var v13 uint64 = 4294967295
    _ = v11
    _ = v12
    _ = v13
    return v13
}

func main() {
    var v1 float32 = 0.1
    var v2 float64 = 3.14159
    v3 := []int{}
    var v4 bool = (true || ((-2.25 >= 0.0) || false))
    var v5 uint32 = 0
    v6 := make(map[string]int)
    v5 += (v5 ^ v5)
    fmt.Println(v5, v2, v5)
    fmt.Println(v4, v1, v1)
    var v14 int64 = 9223372036854775807
    if len(v3) > 0 {
        v3[0] = 851
    }
    fmt.Println(v2, v2)
    fmt.Println(v5, v1)
    fmt.Println(v1)
    if v4 {
        v14 += (v14 << ((42) & 63))
    }
    switch -778 {
    case 0:
        var v15 int16 = -1
        _ = v15
    case 9223372036854775807:
        fmt.Println(v4)
    case 42:
        fmt.Println(v5)
    }
    var v16 int32 = 1
    for fi17 := 0; fi17 < 1; fi17++ {
        v18 := GenS1{F0: "foo bar", F1: 65535}
        for fi19 := 0; fi19 < 15; fi19++ {
            v2 = (v2 - v2)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v16 = (v16 | (-(v16)))
    delete(v6, (("hello" + ("hello" + "a")) + "\n\t"))
    v20 := GenS2{F0: v16}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
