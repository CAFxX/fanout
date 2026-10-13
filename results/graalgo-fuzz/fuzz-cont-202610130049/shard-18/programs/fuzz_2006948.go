package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int64 }
func genFn0(p0 int, p1 uint64, p2 uint) int8 {
    if true {
        if (p1 <= p1) {
            for fi5 := 0; fi5 < 9; fi5++ {
                v6 := []int{}
                p2 -= 18446744073709551615
                var v7 int32 = 42
                _ = fi5
                _ = v6
                _ = v7
            }
        }
    }
    p2 = p2
    p0 += (p0 / (((p0 * 127)) | 1))
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := GenS1{F0: 192, F1: 1000, F2: 9223372036854775807}
        v10 := GenS1{F0: -9223372036854775808, F1: 42, F2: -1}
        fmt.Println(p1)
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return -128
}

func main() {
    var v1 string = ((("hello" + "foo bar") + ("\n\t" + (("foo bar" + ("" + "hello")) + "hello"))) + "")
    v2 := make(map[int]string)
    var v3 uint16 = 1
    var v4 int16 = -1
    if (1 > 4294967295) {
        fmt.Println(v3, v3, v4)
    }
    v4 = v4
    var v11 uint64 = 0
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 uint8 = 244
            v15 := GenS1{F0: -1, F1: (v4 + 0), F2: -669}
            _ = len(v2)
            _ = fi13
            _ = v14
            _ = v15
        }
        fmt.Println(v4)
        _ = fi12
    }
    v16 := []int{-9223372036854775808, -1}
    if true {
        switch (v1 + v1) {
        case "":
            if (false && (true || false)) {
                v4 = v4
                v1 = v1
                switch (("\n\t" + ((("\n\t" + "\n\t") + "hello") + v1)) + (v1 + "a")) {
                case "a":
                    _ = len(v2)
                default:
                    var v17 uint32 = uint32(v3)
                    _ = v17
                }
            } else {
                _ = len(v16)
                v11 = 4294967295
            }
        case "\n\t":
            fmt.Println(v1, v11)
        case "hello":
            var v18 uint64 = (v11 << ((453) & 63))
            _ = v18
        }
        var v19 uint64 = (v11 & uint64(v3))
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 int = (fi20 * (fi20 * fi20))
        var v22 uint = 103
        v1 = "a"
        _ = fi20
        _ = v21
        _ = v22
    }
    v3 = v3
    v4 = v4
    for fi23 := 0; fi23 < 5; fi23++ {
        var v24 uint32 = 342
        v2[(fi23 + 32768)] = "a"
        switch (v1 + (v1 + "\n\t")) {
        case "foo bar":
            fmt.Println(fi23, v24, v1)
        case "hello":
            fmt.Println(v1, fi23)
        default:
            for fi25 := 0; fi25 < 6; fi25++ {
                var v26 uint16 = v3
                v11 += (v11 >> ((4294967295) & 63))
                var v27 bool = true
                var v28 int64 = -890
                if len(v16) > 0 {
                    v16[0] = (fi25 & (fi23 % ((int(v26)) | 1)))
                }
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
        }
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
