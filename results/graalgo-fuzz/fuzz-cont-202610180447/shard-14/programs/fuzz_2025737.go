package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int16, p1 int, p2 int8) int64 {
    v5 := GenS1{F0: 1, F1: 65535, F2: 375}
    fmt.Println(p0, p0)
    switch "\n\t" {
    case "a":
        var v6 int32 = 42
        _ = v6
    case "hello":
        p0 -= (p0 % (((p0 | p0)) | 1))
    default:
        var v7 int32 = 0
        _ = v7
    }
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}
func genFn1(p0 int32, p1 string) int16 {
    var v8 uint64 = 42
    v8 -= v8
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 float32 = (0.0 - 3.14159)
        v11 := GenS1{F0: 0, F1: 2, F2: 386}
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v8
    return -32768
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = -696
    var v3 uint = 1
    v4 := make(map[int]string)
    var v12 float32 = 0.0
    for fi13 := 0; fi13 < 12; fi13++ {
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 uint64 = 837
            if (false || (((102 > 0) && true) && false)) {
                if false {
                    v3 -= (^(v3))
                }
                var v16 float32 = (float32(1e10) * float32(0.0))
                _ = v16
            }
            if false {
                var v17 int64 = -1
                v3 = (v3 ^ v3)
                v18 := GenS1{F0: 755, F1: (^(v3)), F2: 111}
                _ = v17
                _ = v18
            }
            v1[(("a" + "hello") + ("a" + (("hello" + "hello") + "a")))] = 4294967295
            v19 := GenS2{F0: 462}
            _ = fi14
            _ = v15
            _ = v19
        }
        _ = fi13
    }
    var v20 float64 = (float64(v12) - (-2.25 * (1e10 * 0.1)))
    var v21 int = v2
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 uint8 = 1
        var v24 int8 = -128
        v3 -= (v3 ^ v3)
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        var v26 uint64 = 713
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        var v28 uint64 = 1
        switch ("a" + "\n\t") {
        case "\n\t":
            var v29 uint = v3
            _ = v29
        }
        _ = fi27
        _ = v28
    }
    if (32767 == -32768) {
        var v30 uint32 = 42
        fmt.Println(v2, v21, v2)
        _ = v30
    }
    v4[-9223372036854775808] = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
