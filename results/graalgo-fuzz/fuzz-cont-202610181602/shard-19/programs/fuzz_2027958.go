package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 bool) uint64 {
    v6 := []int{-2147483648, -543, 9223372036854775807, 128}
    var v7 int16 = 0
    v8 := GenS1{F0: 18446744073709551615}
    for fi9 := 0; fi9 < 18; fi9++ {
        if (0 >= 0) {
            switch (fi9 * -959) {
            case -9223372036854775808:
                var v10 uint16 = 666
                _ = v10
            case 910:
                var v11 uint = 1
                _ = v11
            case -2147483648:
                v12 := GenS1{F0: 1}
                _ = v12
            default:
                fmt.Println(v7, v7)
            }
            fmt.Println(fi9, p0)
        }
        p0 = (p0 || (843 != 379))
        _ = fi9
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 42
    var v3 int64 = (^(v2))
    v4 := []uint32{670, 1}
    v5 := make(map[uint32]string)
    v3 -= (v2 * v3)
    fmt.Println(v2, v2)
    var v13 int16 = -154
    for fi14 := 0; fi14 < 3; fi14++ {
        v2 -= v2
        var v15 uint = 1
        if false {
            v3 += 9223372036854775807
            _ = len(v4)
        } else {
            v15 -= v15
            fmt.Println(v2, fi14, v2)
            var v16 int32 = 1
            _ = v16
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v13, v2, v13)
    var v17 int8 = 127
    var v18 bool = (false && true)
    v2 += v3
    var v19 int32 = -2147483648
    var v20 uint64 = 1
    var v21 int = 32768
    switch -895 {
    case -9223372036854775808:
        for fi22 := 0; fi22 < 2; fi22++ {
            switch "" {
            case "":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "hello":
                var v23 uint = 326
                _ = v23
            default:
                var v24 int64 = 375
                _ = v24
            }
            _ = fi22
        }
    case -1:
        var v25 int = (v21 | v21)
        _ = v25
    case 32768:
        var v26 uint32 = 42
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
