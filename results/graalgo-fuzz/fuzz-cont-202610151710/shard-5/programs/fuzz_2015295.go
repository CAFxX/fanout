package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 2147483647
    var v3 uint64 = 744
    v4 := make(map[int]bool)
    v5 := make(map[uint32]int)
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 int = fi6
        _ = fi6
        _ = v7
    }
    v8 := []int{1, 4294967295, -9223372036854775808, 32767}
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 int32 = v2
        var v11 uint16 = 1
        var v12 uint8 = 156
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 bool = true
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 int8 = -1
        if v13 {
            v8 = append(v8, fi14)
        } else {
            var v16 int32 = (v2 << ((955) & 63))
            _ = v16
        }
        v3 = 107
        _ = fi14
        _ = v15
    }
    switch ("hello" + ("foo bar" + "a")) {
    case "":
        if v13 {
            fmt.Println(v3, v3)
            v17 := GenS1{F0: 42, F1: v3}
            _ = v17
        }
    case "foo bar":
        var v18 int = 9223372036854775807
        _ = v18
    case "hello":
        v2 += v2
    default:
        var v19 int16 = -32768
        _ = v19
    }
    v3 = (v3 << ((2) & 63))
    _ = len(v4)
    for fi20 := 0; fi20 < 10; fi20++ {
        if (4294967295 <= 232) {
            fmt.Println(v13, v2)
            var v21 float32 = 1e10
            v22 := GenS1{F0: 1, F1: (v3 & (v3 >> ((308) & 63)))}
            _ = v21
            _ = v22
        }
        _ = fi20
    }
    switch "\n\t" {
    case "hello":
        var v23 uint32 = 838
        _ = v23
    case "":
        fmt.Println(v3, v2)
    }
    var v24 int8 = -47
    v25 := GenS2{F0: -185}
    var v26 int16 = 0
    v2 = (v2 - 0)
    for fi27 := 0; fi27 < 15; fi27++ {
        v13 = (4294967295 == 65535)
        switch (fi27 % ((fi27) | 1)) {
        case 361:
            v24 = (v24 % (((^(v24))) | 1))
        case -2147483648:
            if (((127 > 255) && false) || ((uint(v3) >= 65535) && v13)) {
                var v28 int16 = (v26 << ((0) & 63))
                var v29 int64 = 203
                _ = v28
                _ = v29
            } else {
                for fi30 := 0; fi30 < 16; fi30++ {
                    v5[4294967295] = 1
                    _ = fi30
                }
                v31 := GenS1{F0: 161, F1: v3}
                _ = v31
            }
        default:
            for fi32 := 0; fi32 < 13; fi32++ {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                v3 -= 18446744073709551615
                _ = fi32
            }
        }
        switch fi27 {
        case 32768:
            if v13 {
                for fi33 := 0; fi33 < 19; fi33++ {
                    v3 = (^(v3))
                    v1[2147483648] = (false || true)
                    _ = fi33
                }
            }
        default:
            var v34 int32 = -2147483648
            _ = v34
        }
        _ = fi27
    }
    v35 := GenS2{F0: -487}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v8
}
