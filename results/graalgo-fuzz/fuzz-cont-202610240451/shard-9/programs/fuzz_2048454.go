package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint8; F2 int16 }
type GenS4 struct { F0 int8; F1 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int = -1
    var v3 string = "a"
    v4 := []int{(v2 | (-(v2))), (v2 / (((v2 << ((18446744073709551615) & 63))) | 1)), 32767, v2}
    var v5 int8 = 100
    v6 := []uint32{2147483647, 0}
    v1[("hello" + (("\n\t" + (((v3 + (v3 + "hello")) + "foo bar") + "a")) + v3))] = 2
    switch v3 {
    case "":
        fmt.Println(v2)
    }
    if (!false) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if (18446744073709551615 != 4294967295) {
        var v7 int8 = (v5 << ((554) & 63))
        switch -9223372036854775808 {
        case 32768:
            v8 := GenS1{F0: 0, F1: 1}
            _ = v8
        case 2147483647:
            v3 = v3
        default:
            var v9 uint8 = 255
            _ = v9
        }
        var v10 float64 = -2.25
        _ = v10
        _ = v7
    } else {
        for fi11 := 0; fi11 < 15; fi11++ {
            for fi12 := 0; fi12 < 6; fi12++ {
                v13 := GenS1{F0: 2, F1: -9223372036854775808}
                _ = fi12
                _ = v13
            }
            fmt.Println(v5, v5, v5)
            _ = fi11
        }
        var v14 uint64 = 0
        v6 = append(v6, 42)
        _ = v14
    }
    if true {
        var v15 string = v3
        fmt.Println(v3, v3)
        _ = v15
    } else {
        var v16 int32 = -2147483648
        _ = v16
    }
    v17 := GenS2{F0: -929}
    var v18 int64 = -985
    if ((0.1 * 3.14159) >= 1e10) {
        v19 := GenS2{F0: 1}
        var v20 string = v3
        v21 := GenS3{F0: (v2 & (v2 + v2)), F1: 128, F2: -32768}
        _ = v19
        _ = v20
        _ = v21
    } else {
        _ = len(v1)
        var v22 uint16 = 0
        var v23 float32 = (float32(1e10) / float32(3.14159))
        _ = v22
        _ = v23
    }
    fmt.Println(v5, v2)
    var v24 uint32 = 4294967295
    if false {
        v25 := GenS4{F0: v5, F1: 42}
        _ = v25
    } else {
        if false {
            var v26 uint = 65535
            _ = v26
        }
        for fi27 := 0; fi27 < 10; fi27++ {
            v18 -= 0
            _ = fi27
        }
        var v28 uint32 = (v24 | (v24 & 4294967295))
        _ = v28
    }
    v29 := GenS2{F0: 0}
    fmt.Println(v5, v3, v3)
    switch 876 {
    case 0:
        switch (v2 << ((1) & 63)) {
        case 2147483647:
            fmt.Println(v5, v3, v2)
        case 729:
            fmt.Println(v24, v3)
        default:
            v2 = -874
        }
    case -9223372036854775808:
        fmt.Println(v24, v2, v2)
    default:
        if false {
            v1[v3] = v24
        } else {
            _ = len(v1)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
