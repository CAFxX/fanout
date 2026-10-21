package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint32 }
func genFn0(p0 int32) uint {
    p0 -= (-(p0))
    v7 := GenS1{F0: 2}
    var v8 int = -2147483648
    v9 := GenS2{F0: (414 < 286)}
    v10 := []int{}
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}
func genFn1(p0 uint32) int8 {
    p0 = (p0 | 977)
    p0 = 65535
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := GenS1{F0: 42}
        var v13 uint16 = 0
        v14 := GenS1{F0: 42}
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    _ = p0
    return 0
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{"", ("hello" + "a")}
    var v3 float32 = 0.1
    var v4 uint64 = 2
    var v5 string = ("" + "\n\t")
    var v6 uint16 = 149
    switch (v5 + ((v5 + v5) + v5)) {
    case "":
        v6 += (^(v6))
    default:
        switch v5 {
        case "hello":
            var v15 int32 = -879
            _ = v15
        case "a":
            v4 = (v4 + 1)
        default:
            fmt.Println(v4)
        }
    }
    var v16 string = (v5 + v5)
    v17 := v1[451]
    _ = len(v1)
    v17 = v5
    if true {
        v3 -= (1e10 + v3)
        var v18 string = v5
        var v19 uint32 = 690
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: (v6 % (((^(v6))) | 1))}
    if (0 != -9223372036854775808) {
        if (!false) {
            if len(v2) > 0 {
                v2[0] = v16
            }
            v3 -= v3
            var v21 uint16 = 42
            _ = v21
        } else {
            if (!(!true)) {
                var v22 uint = 574
                v23 := GenS3{F0: (v6 * v6), F1: 2147483648, F2: 42}
                _ = v22
                _ = v23
            }
            switch -1 {
            case 9223372036854775807:
                v17 = v17
            case 127:
                fmt.Println(v16, v4)
            default:
                for fi24 := 0; fi24 < 6; fi24++ {
                    v1[fi24] = "foo bar"
                    v3 = ((3.14159 / v3) * v3)
                    var v25 int64 = 436
                    _ = fi24
                    _ = v25
                }
            }
            v26 := GenS2{F0: (1 >= 65535)}
            _ = v26
        }
        var v27 uint64 = (v4 >> ((uint(v6)) & 63))
        _ = v27
    }
    _ = len(v1)
    _ = len(v1)
    v3 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
