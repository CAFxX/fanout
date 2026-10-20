package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int64; F2 int32 }
func genFn0(p0 uint, p1 int32, p2 int32) uint16 {
    var v7 int16 = -922
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 uint16 = 2
        fmt.Println(v7, v7)
        _ = fi8
        _ = v9
    }
    var v10 uint32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 0
}
func genFn1(p0 uint8) int64 {
    p0 += p0
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := []int{(fi11 + fi11), (fi11 << ((42) & 63)), fi11}
        for fi13 := 0; fi13 < 7; fi13++ {
            if len(v12) > 0 {
                v12[0] = 42
            }
            fmt.Println(fi11)
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    return 638
}

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]uint32)
    v3 := []string{"foo bar"}
    var v4 int = 42
    v5 := []string{}
    var v6 int64 = -9223372036854775808
    if false {
        v4 -= 42
    }
    var v14 int32 = 2147483647
    fmt.Println(v14, v4, v4)
    v4 += 0
    var v15 uint16 = 65535
    fmt.Println(v14, v1, v1)
    if false {
        v16 := GenS1{F0: v1}
        v6 = -9223372036854775808
        v14 = (v14 & v14)
        _ = v16
    } else {
        if ((v6 / (((v6 / ((v6) | 1))) | 1)) != -177) {
            v15 = genFn0(81, -2147483648, (v14 ^ v14))
            var v17 uint = 65535
            _ = v17
        }
        if false {
            v3 = append(v3, "foo bar")
            v18 := GenS2{F0: (v4 & v4)}
            _ = v18
        }
        v4 = v4
    }
    v19 := GenS3{F0: 128, F1: v6, F2: -1}
    var v20 string = "\n\t"
    var v21 int64 = v6
    var v22 uint32 = 784
    switch -9223372036854775808 {
    case 4294967295:
        _ = len(v2)
    case 718:
        v23 := v2[(v22 - 2)]
        _ = v23
    case 2147483647:
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 uint64 = 42
            _ = fi24
            _ = v25
        }
    }
    var v26 int64 = (v6 | (v21 - int64(v15)))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
