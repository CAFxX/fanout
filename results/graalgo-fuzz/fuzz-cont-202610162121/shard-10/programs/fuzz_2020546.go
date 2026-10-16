package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint64 }
type GenS2 struct { F0 string }
func genFn0(p0 int16, p1 string) int8 {
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int = 2147483647
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 4294967295
    var v10 uint8 = 128
    var v11 int32 = -295
    if (true && true) {
        v12 := []int{int(p0), 183, -1}
        _ = v12
    } else {
        v13 := GenS1{F0: 1, F1: -1, F2: 4294967295}
        var v14 uint32 = 240
        v9 -= (v9 & (v14 | 72))
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 1
}
func genFn1(p0 uint64, p1 int32) uint {
    v15 := GenS2{F0: "hello"}
    if (18446744073709551615 < 18446744073709551615) {
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = v15
    return 968
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    var v3 float32 = 0.0
    v4 := []string{"\n\t", "foo bar", "a"}
    var v5 uint = 65535
    var v6 int32 = 0
    var v16 int8 = -92
    if len(v4) > 0 {
        v4[0] = (("" + "") + "a")
    }
    v1[2147483648] = int(v5)
    v17 := GenS1{F0: (-(v16)), F1: 115, F2: 18446744073709551615}
    _ = len(v2)
    v6 = (v6 * int32(v16))
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 int64 = 468
        _ = fi18
        _ = v19
    }
    v16 = (v16 * v16)
    var v20 uint64 = 2
    switch 128 {
    case 154:
        if true {
            if (!true) {
                var v21 int = 420
                v16 += v16
                fmt.Println(v5)
                _ = v21
            }
        }
    case -9223372036854775808:
        var v22 uint32 = 93
        _ = v22
    case 32768:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v23 := GenS2{F0: "\n\t"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
