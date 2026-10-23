package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 4294967295
    v3 := make(map[int]bool)
    v1[-9223372036854775808] = ("\n\t" + "")
    var v4 int32 = 602
    v4 -= v4
    var v5 int64 = 0
    fmt.Println(v2, v5, v5)
    if (1 != 2) {
        for fi6 := 0; fi6 < 17; fi6++ {
            if (-9223372036854775808 >= fi6) {
                v7 := GenS1{F0: 65535, F1: ((!true) || true), F2: v5}
                _ = v7
            } else {
                v8 := GenS1{F0: 0, F1: (true && true), F2: 0}
                var v9 uint = 65535
                _ = v8
                _ = v9
            }
            v10 := []int{fi6, -1, fi6}
            _ = fi6
            _ = v10
        }
        v11 := GenS1{F0: 0, F1: false, F2: (v5 % (((v5 * (-(v5)))) | 1))}
        fmt.Println(v4, v4)
        _ = v11
    }
    switch 0 {
    case -2147483648:
        fmt.Println(v2)
    case 9223372036854775807:
        v12 := []int{4294967295, 9223372036854775807}
        _ = v12
    default:
        var v13 uint64 = 898
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        fmt.Println(v2)
        switch "" {
        case "\n\t":
            v15 := GenS1{F0: 42, F1: ((65535 != 0) && true), F2: v5}
            _ = v15
        case "a":
            var v16 uint = 65535
            _ = v16
        default:
            var v17 int = fi14
            _ = v17
        }
        _ = fi14
    }
    delete(v1, int(v4))
    if false {
        v5 = (v5 >> ((2) & 63))
        v18 := []int{}
        _ = v18
    }
    v19 := GenS1{F0: 383, F1: true, F2: v5}
    v20 := []int{-887, 32767, 32768, -1}
    _ = len(v20)
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
