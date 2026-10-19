package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint, p1 int) uint16 {
    fmt.Println(p1, p1)
    var v7 int8 = 100
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 uint = 0
    var v2 uint8 = 148
    v3 := []uint32{209, 2147483647}
    v4 := make(map[string]bool)
    var v5 uint8 = v2
    var v6 float64 = 1.5
    v8 := GenS1{F0: (v1 * v1)}
    var v9 float32 = (0.0 - 0.0)
    v9 += ((v9 / 1.0) * float32(v6))
    switch (("\n\t" + (("\n\t" + ("hello" + "")) + "foo bar")) + "a") {
    case "foo bar":
        _ = len(v4)
    case "\n\t":
        if false {
            v10 := GenS1{F0: v1}
            _ = v10
        }
    default:
        _ = len(v3)
    }
    switch -2147483648 {
    case -982:
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 float32 = 0.0
            delete(v4, "\n\t")
            var v13 int8 = 0
            _ = fi11
            _ = v12
            _ = v13
        }
    case -9223372036854775808:
        if false {
            v9 = v9
            fmt.Println(v2)
        }
    }
    if true {
        v14 := GenS2{F0: v1, F1: 0}
        v2 += 22
        _ = v14
    } else {
        var v15 int64 = -9223372036854775808
        v5 = v5
        var v16 int8 = 127
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v5 = (v5 >> (((^(v1))) & 63))
        _ = fi17
    }
    var v18 int32 = 1
    fmt.Println(v6, v9)
    v2 += (v5 % ((v5) | 1))
    switch ("" + "") {
    case "hello":
        v1 = 988
    case "":
        for fi19 := 0; fi19 < 17; fi19++ {
            fmt.Println(fi19)
            _ = fi19
        }
    }
    v6 -= 0.1
    var v20 bool = true
    if v20 {
        switch -1 {
        case 142:
            v20 = v20
        case 121:
            v21 := GenS3{F0: 399}
            _ = v21
        default:
            v22 := GenS2{F0: (v1 - v1), F1: 9223372036854775807}
            _ = v22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
