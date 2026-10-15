package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint8 }
func genFn0(p0 uint64, p1 uint8) uint64 {
    p0 = 4294967295
    p1 = p1
    v4 := GenS1{F0: -1, F1: 667, F2: (^(p1))}
    _ = p0
    _ = p1
    _ = v4
    return (p0 >> ((62) & 63))
}
func genFn1(p0 int16, p1 int64) int64 {
    var v5 uint64 = genFn0(457, 1)
    v6 := []int{}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (^(p1))
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]int)
    var v3 float64 = 3.14159
    if true {
        v7 := GenS1{F0: 119, F1: -2147483648, F2: 0}
        if (1 != -1) {
            fmt.Println(v3, v1)
        } else {
            v8 := []int{-292, 32768}
            v1 = (v1 / ((v1) | 1))
            _ = v8
        }
        if true {
            switch -1 {
            case -1:
                for fi9 := 0; fi9 < 19; fi9++ {
                    for fi10 := 0; fi10 < 2; fi10++ {
                        var v11 float64 = 1e10
                        _ = fi10
                        _ = v11
                    }
                    v1 -= (v1 / ((v1) | 1))
                    _ = fi9
                }
            }
        }
        _ = v7
    }
    fmt.Println(v3, v3)
    v3 = v3
    for fi12 := 0; fi12 < 20; fi12++ {
        switch (("" + "\n\t") + "foo bar") {
        case "foo bar":
            v3 += v3
        }
        var v13 float64 = float64(1.5)
        _ = fi12
        _ = v13
    }
    var v14 float64 = (0.0 * float64((3.14159 / ((0.1 / 1.0) + 1e10))))
    fmt.Println(v3, v14)
    v14 -= v14
    v3 += (v14 * ((0.0 * -2.25) + 0.0))
    if (282 < 65535) {
        if (!(-655 != -835)) {
            var v15 uint8 = 218
            if (true || false) {
                v16 := []int{127, -480, -548, 128}
                _ = v16
            } else {
                var v17 uint32 = 829
                fmt.Println(v17, v17, v1)
                v17 = v1
                _ = v17
            }
            v14 -= float64(1e10)
            _ = v15
        }
        var v18 int64 = -580
        v1 -= 4294967295
        _ = v18
    } else {
        v14 -= v3
        switch "a" {
        case "":
            v19 := GenS1{F0: 127, F1: -237, F2: 42}
            _ = v19
        case "\n\t":
            var v20 uint16 = 699
            _ = v20
        default:
            v21 := GenS1{F0: -22, F1: 42, F2: 0}
            _ = v21
        }
    }
    var v22 int = -2147483648
    var v23 int64 = 1
    v24 := GenS1{F0: 42, F1: -889, F2: 15}
    v25 := []int{9223372036854775807, -9223372036854775808, v22, int(v1)}
    v1 -= (v1 / ((0) | 1))
    var v26 uint8 = 41
    v27 := GenS1{F0: 6, F1: 1, F2: v26}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
