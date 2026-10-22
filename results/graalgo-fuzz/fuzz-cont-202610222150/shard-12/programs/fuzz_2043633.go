package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
func genFn0(p0 int16, p1 uint8) int16 {
    var v4 int32 = 1
    var v5 string = (("\n\t" + ("\n\t" + ("foo bar" + ""))) + "a")
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (p0 | (p0 >> ((18446744073709551615) & 63)))
}
func genFn1(p0 int64, p1 string, p2 int16) uint64 {
    for fi6 := 0; fi6 < 13; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    p0 = 1
    var v7 uint8 = 1
    fmt.Println(p1, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}

func main() {
    var v1 string = ""
    v2 := []string{"a", v1, "", v1}
    v3 := make(map[int]int)
    var v8 int = 127
    if (genFn0(-1, 127) < 1000) {
        delete(v3, -2147483648)
        var v9 uint32 = 4294967295
        _ = v9
    } else {
        if true {
            v8 += (-(v8))
        } else {
            var v10 int16 = 748
            _ = v10
        }
    }
    switch v8 {
    case 32767:
        switch v1 {
        case "hello":
            fmt.Println(v8, v1)
        default:
            v2 = append(v2, (("foo bar" + v1) + v1))
        }
    case 128:
        fmt.Println(v8)
    }
    if false {
        var v11 float32 = 0.0
        var v12 uint64 = 1
        fmt.Println(v8, v1, v1)
        _ = v11
        _ = v12
    }
    var v13 uint32 = 0
    var v14 uint64 = genFn1(621, v1, -1)
    v15 := GenS1{F0: genFn0(-32768, 0), F1: int8(v8)}
    v16 := GenS1{F0: -32768, F1: -1}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
}
