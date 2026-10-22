package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint }
type GenS2 struct { F0 uint; F1 uint64; F2 bool }
type GenS3 struct { F0 int16; F1 int; F2 string }
func genFn0(p0 int) int32 {
    p0 -= p0
    var v6 uint16 = 0
    _ = p0
    _ = v6
    return 792
}
func genFn1(p0 int16) string {
    var v7 int32 = 42
    p0 += -417
    _ = p0
    _ = v7
    return "\n\t"
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 663
    v3 := make(map[int]int)
    var v4 int64 = 0
    v5 := make(map[int]uint32)
    _ = len(v3)
    switch 128 {
    case -2147483648:
        v4 = v4
    case 2147483647:
        v4 -= 1
    default:
        var v8 int = 32768
        _ = v8
    }
    v9 := GenS1{F0: -513, F1: -859, F2: 196}
    var v10 bool = false
    for fi11 := 0; fi11 < 14; fi11++ {
        v2 += 4294967295
        for fi12 := 0; fi12 < 15; fi12++ {
            v4 = -1
            var v13 string = "a"
            var v14 float64 = 3.14159
            _ = fi12
            _ = v13
            _ = v14
        }
        v15 := GenS2{F0: 593, F1: v2, F2: (!v10)}
        _ = fi11
        _ = v15
    }
    v16 := GenS3{F0: 1, F1: 822, F2: ("" + "a")}
    var v17 int = 9223372036854775807
    if (127 > (v17 & v17)) {
        v18 := GenS2{F0: 65535, F1: v2, F2: (156 > 0)}
        v19 := []int{(^(v17)), 732}
        _ = v18
        _ = v19
    } else {
        v4 -= (-(v4))
        v20 := GenS3{F0: -32768, F1: v17, F2: ("foo bar" + ("\n\t" + "a"))}
        var v21 string = "foo bar"
        _ = v20
        _ = v21
    }
    v10 = ("a" >= "")
    v22 := []int{v17, (v17 - int(v4)), v17, v17}
    var v23 uint8 = 42
    if (866 == 65535) {
        fmt.Println(v10)
        if v10 {
            fmt.Println(v2, v2)
            for fi24 := 0; fi24 < 20; fi24++ {
                var v25 float64 = 0.0
                _ = fi24
                _ = v25
            }
        }
        v17 -= 2147483647
    }
    v26 := v3[(v17 + -444)]
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
