package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int64, p1 string) int64 {
    switch -1 {
    case -2147483648:
        v7 := GenS1{F0: false, F1: p1, F2: 42}
        _ = v7
    case -9223372036854775808:
        var v8 int16 = 1
        _ = v8
    }
    var v9 uint8 = 2
    if ((p0 & p0) < 0) {
        var v10 uint8 = (v9 / (((v9 ^ v9)) | 1))
        p1 = p1
        _ = v10
    } else {
        p1 = (("foo bar" + ("hello" + p1)) + ("foo bar" + "foo bar"))
        v11 := GenS1{F0: true, F1: ((p1 + p1) + (("foo bar" + "hello") + p1)), F2: (v9 ^ (v9 % (((v9 << ((18446744073709551615) & 63))) | 1)))}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v9
    return 72
}

func main() {
    var v1 int64 = 0
    var v2 int16 = -1
    v3 := make(map[uint32]bool)
    v4 := make(map[int]int)
    var v5 uint = 255
    var v6 int16 = v2
    switch 32768 {
    case 0:
        v12 := GenS1{F0: false, F1: ("a" + "a"), F2: 38}
        _ = v12
    case 1:
        fmt.Println(v2, v1, v1)
    case 151:
        v6 = (-(v6))
    }
    v5 = (v5 + (^(v5)))
    var v13 bool = true
    var v14 uint64 = 2
    if (255 == 1) {
        var v15 int16 = (v2 - (v6 << ((18446744073709551615) & 63)))
        if v13 {
            v16 := GenS2{F0: (v14 - v14), F1: 2}
            var v17 int32 = 42
            _ = v16
            _ = v17
        }
        _ = v15
    }
    v14 = v14
    switch "" {
    case "":
        if v13 {
            v5 = (v5 - (v5 / (((v5 * v5)) | 1)))
        } else {
            v6 -= 31
        }
    }
    v18 := []int{-9223372036854775808, 2147483647, -376, 1}
    var v19 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
