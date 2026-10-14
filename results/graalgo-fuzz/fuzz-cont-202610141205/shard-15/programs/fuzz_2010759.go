package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int32 }
func genFn0() uint16 {
    var v4 uint16 = 1
    if (false || ((true || (1 > 18446744073709551615)) && true)) {
        v5 := GenS1{F0: -108, F1: "a", F2: 828}
        fmt.Println(v4)
        _ = v5
    }
    _ = v4
    return 701
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[int]int)
    var v3 int64 = 102
    for fi6 := 0; fi6 < 5; fi6++ {
        v1 = (v1 ^ 2)
        v1 += v1
        if (42 != -1) {
            v2[fi6] = (-(fi6))
        }
        _ = fi6
    }
    var v7 uint = 132
    v8 := []int{42, 1, 2147483647, -1}
    delete(v2, 9223372036854775807)
    if true {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint8 = uint8(v1)
            switch fi9 {
            case -2147483648:
                var v11 uint8 = v10
                _ = v11
            case 0:
                for fi12 := 0; fi12 < 19; fi12++ {
                    var v13 int32 = -1
                    var v14 uint32 = 0
                    var v15 uint = v7
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
            default:
                var v16 uint8 = (^(v10))
                _ = v16
            }
            _ = fi9
            _ = v10
        }
        var v17 uint64 = 2
        var v18 int = 951
        _ = v17
        _ = v18
    } else {
        switch -849 {
        case 9223372036854775807:
            var v19 uint64 = 2
            _ = v19
        case -2147483648:
            v1 += v1
        default:
            var v20 int64 = (v3 * 42)
            _ = v20
        }
        var v21 int64 = 9223372036854775807
        _ = v21
    }
    var v22 string = ""
    v22 = "hello"
    var v23 uint64 = 4294967295
    v7 += (v7 << ((uint(v1)) & 63))
    fmt.Println(v7, v7)
    var v24 uint16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v7
    _ = v8
}
