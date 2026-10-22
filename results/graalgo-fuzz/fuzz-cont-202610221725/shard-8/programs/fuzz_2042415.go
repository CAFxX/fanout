package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int32 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int32; F1 uint8; F2 int }
func genFn0(p0 int16, p1 int64, p2 uint32) string {
    v6 := []int{}
    v7 := GenS1{F0: 3, F1: 740, F2: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "hello"
}
func genFn1() uint32 {
    var v8 uint = 18446744073709551615
    var v9 int64 = -9223372036854775808
    v8 = v8
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 int = -1
    var v2 int = v1
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]bool)
    v5 := []string{"", "\n\t", ("\n\t" + ("\n\t" + "hello"))}
    for fi10 := 0; fi10 < 14; fi10++ {
        v1 = (v1 & (v2 >> ((4294967295) & 63)))
        fmt.Println(v2)
        _ = fi10
    }
    if (0.0 > (float32(float64(0.0)) / 1e10)) {
        var v11 uint16 = 626
        if (true || true) {
            var v12 int16 = 54
            v13 := GenS1{F0: int8(v12), F1: 4294967295, F2: -2147483648}
            for fi14 := 0; fi14 < 6; fi14++ {
                v11 -= 640
                _ = fi14
            }
            _ = v12
            _ = v13
        }
        _ = len(v4)
        _ = v11
    }
    var v15 int64 = -9223372036854775808
    var v16 uint = 0
    _ = len(v5)
    v1 = v2
    var v17 uint32 = 517
    var v18 uint64 = 1
    var v19 int64 = 42
    if true {
        if true {
            fmt.Println(v19)
        }
        v20 := GenS2{F0: ((("\n\t" + ("hello" + "\n\t")) + "foo bar") + ("hello" + ("a" + ("" + "hello")))), F1: 42}
        v21 := GenS3{F0: -831, F1: 30, F2: (v2 | -1)}
        _ = v20
        _ = v21
    } else {
        var v22 uint64 = v18
        if ((-(v15)) <= v15) {
            var v23 int32 = 782
            _ = v23
        }
        _ = v22
    }
    var v24 uint64 = (v18 + (v18 >> ((2) & 63)))
    for fi25 := 0; fi25 < 10; fi25++ {
        v19 = 264
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
