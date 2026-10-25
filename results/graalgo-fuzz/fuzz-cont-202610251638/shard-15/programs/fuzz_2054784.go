package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
func genFn0(p0 uint64, p1 int32) int {
    var v6 uint8 = 177
    v6 = v6
    v7 := GenS1{F0: 1, F1: "a"}
    fmt.Println(p0, p1, v6)
    for fi8 := 0; fi8 < 7; fi8++ {
        v6 -= (v6 << ((9) & 63))
        var v9 uint32 = 65535
        var v10 int32 = (p1 * -213)
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 128
}
func genFn1(p0 bool, p1 string) string {
    var v11 int8 = 127
    var v12 int = 0
    for fi13 := 0; fi13 < 2; fi13++ {
        fmt.Println(v12, v11, v12)
        var v14 int8 = v11
        if (27 >= 0) {
            if ((fi13 & fi13) < fi13) {
                v15 := GenS1{F0: v14, F1: p1}
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 uint16 = 32767
                    v18 := []int{(^(v12)), (v12 * (fi13 * (fi13 - v12))), -1}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = v15
            }
        } else {
            var v19 float32 = 0.1
            v14 = v14
            _ = v19
        }
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return ("" + p1)
}

func main() {
    var v1 int64 = 257
    var v2 uint64 = 899
    v3 := make(map[int]uint32)
    var v4 uint64 = 18446744073709551615
    var v5 uint16 = 1
    v5 += v5
    var v20 uint64 = v4
    v21 := GenS1{F0: -128, F1: "hello"}
    v20 = v2
    _ = len(v3)
    switch "hello" {
    case "hello":
        for fi22 := 0; fi22 < 13; fi22++ {
            var v23 uint64 = 42
            _ = fi22
            _ = v23
        }
    case "":
        v5 = 786
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        if false {
            v2 = v2
            v25 := GenS1{F0: 42, F1: ""}
            _ = v25
        }
        switch (fi24 << ((255) & 63)) {
        case -2147483648:
            var v26 uint8 = 42
            _ = v26
        case 127:
            v3[fi24] = 65535
        case 0:
            for fi27 := 0; fi27 < 13; fi27++ {
                var v28 uint16 = (v5 ^ (v5 >> ((1) & 63)))
                if ((v4 >> ((uint(v28)) & 63)) != 0) {
                    v29 := GenS1{F0: 40, F1: "hello"}
                    var v30 bool = (1e10 < 1.5)
                    _ = v29
                    _ = v30
                }
                _ = fi27
                _ = v28
            }
        default:
            var v31 int16 = 32767
            _ = v31
        }
        _ = fi24
    }
    fmt.Println(v1)
    v20 = v20
    fmt.Println(v2, v2, v20)
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
