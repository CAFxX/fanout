package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint32 }

func main() {
    var v1 float32 = float32(1.5)
    var v2 int8 = -128
    v3 := make(map[int]uint32)
    v4 := []uint32{4294967295, 2147483647, 2147483648}
    v5 := []string{("a" + ""), "a", "hello"}
    var v6 uint32 = 42
    v3[0] = (v6 ^ (v6 | v6))
    if (!true) {
        v4 = append(v4, 2147483647)
    }
    var v7 uint32 = 2147483647
    fmt.Println(v7, v2, v1)
    var v8 uint16 = 601
    fmt.Println(v1, v7)
    v9 := GenS1{F0: (false || (false || (!true))), F1: int32(v8), F2: v6}
    fmt.Println(v6)
    v10 := GenS1{F0: (71 <= 2), F1: 42, F2: 202}
    switch ("hello" + "hello") {
    case "a":
        if (-9223372036854775808 == -9223372036854775808) {
            for fi11 := 0; fi11 < 12; fi11++ {
                fmt.Println(v1)
                _ = fi11
            }
            var v12 int8 = -1
            _ = v12
        }
    case "\n\t":
        var v13 bool = ((640 == 18446744073709551615) || (29 == 255))
        _ = v13
    case "":
        v14 := GenS1{F0: true, F1: 2147483647, F2: uint32(v8)}
        _ = v14
    default:
        var v15 int32 = -2147483648
        _ = v15
    }
    var v16 string = ("foo bar" + (("hello" + "a") + ("foo bar" + (("a" + ("foo bar" + "hello")) + "hello"))))
    if true {
        if ((v8 / ((0) | 1)) < 32767) {
            _ = len(v4)
            fmt.Println(v8)
        }
    }
    v3[32768] = v7
    v8 = 1
    v1 -= v1
    var v17 uint64 = 485
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
