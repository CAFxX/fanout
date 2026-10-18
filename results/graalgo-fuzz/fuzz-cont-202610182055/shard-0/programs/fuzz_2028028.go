package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int16 }
func genFn0(p0 string) uint8 {
    if (!(774 < 52)) {
        var v5 int64 = -9223372036854775808
        fmt.Println(p0, v5)
        _ = v5
    }
    if false {
        p0 = p0
    }
    if true {
        switch 283 {
        case -259:
            var v6 uint8 = 128
            _ = v6
        default:
            fmt.Println(p0)
        }
        switch 9223372036854775807 {
        case 0:
            v7 := GenS1{F0: -1, F1: 1}
            _ = v7
        case 473:
            p0 = (p0 + p0)
        default:
            v8 := GenS1{F0: -864, F1: 518}
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 string = p0
        _ = fi9
        _ = v10
    }
    p0 = (p0 + "")
    _ = p0
    return 104
}
func genFn1(p0 int64, p1 string, p2 bool) uint8 {
    fmt.Println(p1)
    var v11 bool = false
    if (0 != 65535) {
        p2 = false
    }
    v12 := GenS2{F0: 2147483647, F1: 2}
    for fi13 := 0; fi13 < 16; fi13++ {
        v14 := []int{(fi13 | (fi13 + 9223372036854775807)), (fi13 - fi13), (fi13 >> ((18446744073709551615) & 63))}
        v15 := GenS3{F0: 916}
        var v16 float32 = float32((0.1 + float64(0.1)))
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 127
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 42
    var v3 int64 = 0
    var v4 uint64 = 4294967295
    fmt.Println(v4, v3, v3)
    v1[4294967295] = (232 <= genFn1(v3, ((("hello" + "hello") + ((("a" + "\n\t") + ("foo bar" + "\n\t")) + ("foo bar" + "foo bar"))) + "hello"), false))
    var v17 int = -471
    v3 = -926
    v4 = v4
    v2 = v2
    var v18 bool = ((0.0 + 0.0) < 3.14159)
    fmt.Println(v2)
    v19 := []int{}
    var v20 float64 = (0.1 + (((0.0 + 1.5) * (0.0 * (0.0 + (((1.5 + 3.14159) + 3.14159) + (0.1 - 0.1))))) / 1e10))
    var v21 uint16 = 231
    fmt.Println(v17, v20, v2)
    if len(v19) > 0 {
        v19[0] = (v17 | (^(v17)))
    }
    var v22 uint = uint(v21)
    var v23 uint32 = 2147483648
    v18 = false
    var v24 float64 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
