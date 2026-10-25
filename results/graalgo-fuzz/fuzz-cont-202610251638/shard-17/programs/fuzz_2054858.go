package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint }
type GenS3 struct { F0 uint; F1 uint8; F2 uint }
type GenS4 struct { F0 uint8; F1 int16; F2 int8 }
func genFn0(p0 int16, p1 int, p2 int32) int16 {
    p0 -= p0
    if true {
        v5 := GenS1{F0: (p1 - (p1 + p1)), F1: 127, F2: (p2 ^ (p2 % (((p2 >> ((4294967295) & 63))) | 1)))}
        var v6 int8 = 0
        _ = v5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return -32768
}
func genFn1() int16 {
    if ((((0.1 + float32((0.1 * ((-2.25 * 1e10) + (0.0 * (0.1 * -2.25)))))) / 1.0) / 1.0) < 1.5) {
        var v7 int = -244
        fmt.Println(v7)
        _ = v7
    } else {
        var v8 uint64 = 669
        fmt.Println(v8)
        _ = v8
    }
    var v9 int16 = 374
    v10 := []int{}
    _ = v10
    _ = v9
    return (v9 ^ (v9 | (v9 & v9)))
}

func main() {
    v1 := []uint32{42}
    var v2 uint64 = 1
    v3 := []string{"foo bar", ("a" + "a"), ("foo bar" + "a"), "\n\t"}
    v4 := []uint32{65535, 549, 1}
    var v11 int16 = 32767
    var v12 uint = 18446744073709551615
    fmt.Println(v2, v12, v2)
    _ = len(v3)
    var v13 float32 = (float32(3.14159) + 0.1)
    v14 := GenS1{F0: 0, F1: 1, F2: 781}
    if ((226 <= uint8(v12)) || true) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        v12 = (v12 % (((v12 >> ((v12) & 63))) | 1))
    }
    if (1 > 107) {
        if true {
            v15 := GenS2{F0: 1, F1: 2, F2: 1}
            var v16 int16 = genFn1()
            _ = v15
            _ = v16
        }
    } else {
        switch "\n\t" {
        case "a":
            var v17 int8 = 42
            _ = v17
        case "foo bar":
            v11 += v11
        }
        var v18 int64 = -1
        _ = v18
    }
    v19 := GenS3{F0: (v12 / ((v12) | 1)), F1: 0, F2: (v12 % (((v12 - (v12 >> ((v12) & 63)))) | 1))}
    if (9223372036854775807 != 42) {
        switch ("hello" + ("" + "hello")) {
        case "hello":
            v20 := GenS4{F0: 255, F1: 1000, F2: 42}
            _ = v20
        default:
            var v21 int = -2147483648
            _ = v21
        }
        fmt.Println(v11)
        v3 = append(v3, ((("a" + "hello") + ("a" + "a")) + "\n\t"))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
