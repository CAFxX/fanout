package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint }
func genFn0(p0 uint32) uint64 {
    v6 := GenS1{F0: 128}
    v7 := GenS2{F0: 2147483647, F1: false, F2: 18446744073709551615}
    _ = p0
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 uint = 1
    var v2 int8 = 127
    var v3 float64 = ((-2.25 + 0.1) * (0.1 - float64(1.5)))
    var v4 float64 = v3
    v5 := make(map[string]bool)
    if false {
        switch 681 {
        case 32768:
            v8 := GenS3{F0: 0, F1: v1}
            _ = v8
        default:
            if true {
                if (true || (true || (((0 <= 9223372036854775807) || (!true)) || false))) {
                    fmt.Println(v3, v4, v4)
                }
                var v9 int32 = 351
                v3 -= float64(0.1)
                _ = v9
            } else {
                v4 += 3.14159
                v10 := []int{9223372036854775807, 1}
                _ = v10
            }
        }
    } else {
        if false {
            var v11 bool = false
            v11 = (738 < 2147483647)
            var v12 int32 = -276
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 uint8 = 128
            _ = fi13
            _ = v14
        }
        if (42 > -1) {
            v15 := v5["\n\t"]
            var v16 uint8 = 194
            _ = v15
            _ = v16
        } else {
            var v17 uint32 = 862
            var v18 string = ""
            _ = v17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        fmt.Println(v3, v3)
        _ = fi19
    }
    switch ((("\n\t" + "") + "\n\t") + "foo bar") {
    case "hello":
        switch "foo bar" {
        case "\n\t":
            v20 := []int{}
            _ = v20
        }
    }
    v2 = -128
    var v21 uint32 = 2147483648
    v22 := GenS3{F0: 0, F1: 430}
    delete(v5, "\n\t")
    if (((v21 & v21) <= (v21 >> ((504) & 63))) && false) {
        var v23 uint32 = (v21 | (v21 << ((v1) & 63)))
        fmt.Println(v4)
        var v24 uint8 = 2
        _ = v23
        _ = v24
    }
    v25 := GenS2{F0: -1, F1: true, F2: 42}
    v2 += (v2 % ((88) | 1))
    v4 = float64(0.0)
    var v26 int64 = 1
    v27 := []int{-1, 823, -1, -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
