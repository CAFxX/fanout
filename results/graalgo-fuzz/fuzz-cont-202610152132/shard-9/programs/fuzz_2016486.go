package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint; F2 int32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int8, p1 bool) int8 {
    fmt.Println(p1, p1)
    v7 := []int{1, 181, -2147483648, 9223372036854775807}
    var v8 int32 = 42
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p0
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = 0
    var v3 int32 = 387
    v4 := make(map[string]string)
    v5 := make(map[int]int)
    var v6 uint = 65535
    v2 -= v2
    if (1 <= 289) {
        v3 += 1
        switch (v2 * v2) {
        case -9223372036854775808:
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 int32 = v3
                v11 := GenS1{F0: -128, F1: "foo bar", F2: 2}
                _ = fi9
                _ = v10
                _ = v11
            }
        default:
            var v12 bool = true
            _ = v12
        }
    }
    var v13 string = ("a" + "hello")
    var v14 uint = v6
    switch -1 {
    case 1:
        v15 := v5[2147483647]
        _ = v15
    default:
        if (0.1 != float32(-2.25)) {
            var v16 int64 = 42
            _ = v16
        } else {
            var v17 float64 = 1.5
            v2 += (v2 >> ((1) & 63))
            v18 := GenS2{F0: 221, F1: (v14 % ((v6) | 1)), F2: (v3 >> (((^(v6))) & 63))}
            _ = v17
            _ = v18
        }
    }
    v19 := GenS3{F0: 1}
    for fi20 := 0; fi20 < 19; fi20++ {
        var v21 bool = false
        if ((2147483648 < 42) || v21) {
            var v22 uint = (v6 | v6)
            var v23 uint32 = 2147483648
            _ = v22
            _ = v23
        }
        switch "\n\t" {
        case "a":
            fmt.Println(v3, v13)
        case "":
            var v24 int32 = 0
            _ = v24
        }
        _ = fi20
        _ = v21
    }
    var v25 int32 = (v3 | (v3 % ((v3) | 1)))
    var v26 uint32 = 0
    fmt.Println(v3)
    v25 += (v25 >> ((42) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
