package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint16 }
type GenS2 struct { F0 int64 }
func genFn0() uint64 {
    var v5 uint = 540
    if (false && (0 > 199)) {
        if (1e10 <= 0.0) {
            switch (("hello" + ("\n\t" + "foo bar")) + "hello") {
            case "hello":
                var v6 int64 = -9223372036854775808
                _ = v6
            case "\n\t":
                fmt.Println(v5)
            }
            v5 = v5
        } else {
            fmt.Println(v5)
            var v7 uint8 = 170
            v8 := GenS1{F0: (false || (false || true)), F1: -32768, F2: 690}
            _ = v7
            _ = v8
        }
    } else {
        v9 := GenS1{F0: (uint16(v5) <= 42), F1: 32767, F2: 65535}
        _ = v9
    }
    v10 := []int{}
    _ = v10
    _ = v5
    return 2
}
func genFn1(p0 uint8) uint8 {
    v11 := GenS2{F0: 42}
    fmt.Println(p0)
    if (-1 != 42) {
        p0 = p0
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := []int{9223372036854775807}
            _ = fi12
            _ = v13
        }
        var v14 uint = 698
        _ = v14
    }
    _ = p0
    _ = v11
    return (p0 / (((p0 | 0)) | 1))
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    var v4 uint16 = 2
    if true {
        v15 := []int{-1}
        var v16 int8 = -101
        _ = v15
        _ = v16
    } else {
        v4 += v4
        v4 = v4
        var v17 int64 = 0
        _ = v17
    }
    v1 = v1
    if true {
        var v18 bool = ((32767 < v4) && (v1 >= (v1 - (-2.25 + v1))))
        switch 1 {
        case 42:
            v19 := GenS2{F0: 718}
            _ = v19
        case -529:
            fmt.Println(v1, v18)
        case 32768:
            delete(v3, "a")
        }
        var v20 uint16 = v4
        _ = v18
        _ = v20
    } else {
        switch -324 {
        case -455:
            v1 -= (v1 + (v1 + v1))
        case 128:
            var v21 uint8 = 68
            _ = v21
        case 32767:
            v4 -= (v4 * v4)
        default:
            var v22 string = "\n\t"
            _ = v22
        }
        fmt.Println(v4, v4)
    }
    v23 := GenS1{F0: true, F1: 42, F2: v4}
    v1 -= float32(0.0)
    v4 = (v4 << ((255) & 63))
    v1 = -2.25
    fmt.Println(v1)
    v4 -= v4
    var v24 uint32 = 2147483647
    if false {
        var v25 uint32 = v24
        v24 -= (v24 >> ((42) & 63))
        _ = v25
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        v27 := GenS1{F0: true, F1: -1, F2: (v4 & v4)}
        var v28 int32 = 1
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
