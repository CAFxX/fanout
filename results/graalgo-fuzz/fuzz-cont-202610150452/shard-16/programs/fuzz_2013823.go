package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint32 }

func main() {
    var v1 uint64 = 42
    var v2 int64 = 9223372036854775807
    var v3 int = 288
    fmt.Println(v1, v3, v1)
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        for fi4 := 0; fi4 < 2; fi4++ {
            v5 := GenS1{F0: (v2 | v2)}
            for fi6 := 0; fi6 < 7; fi6++ {
                var v7 float32 = (3.14159 - -2.25)
                _ = fi6
                _ = v7
            }
            v8 := []int{42, fi4}
            _ = fi4
            _ = v5
            _ = v8
        }
    case "\n\t":
        fmt.Println(v1, v3)
    case "a":
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 int16 = 32767
            v11 := []int{(v3 * fi9), (v3 / ((v3) | 1)), int(v1)}
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        fmt.Println(v1, v3)
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        v13 := GenS1{F0: v2}
        if true {
            var v14 int8 = 0
            var v15 uint8 = 42
            _ = v14
            _ = v15
        } else {
            switch (-(fi12)) {
            case -2147483648:
                var v16 uint64 = 108
                _ = v16
            case 64:
                v2 = v2
            }
        }
        if ((1 > 0) || ((((0 != (v2 | v2)) && (-1 < -103)) || true) && true)) {
            var v17 float64 = 1e10
            _ = v17
        }
        _ = fi12
        _ = v13
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        if (411 >= 572) {
            switch (fi18 / ((-365) | 1)) {
            case 32767:
                var v19 uint8 = 128
                _ = v19
            case 4294967295:
                v3 -= 42
            case 1:
                v2 = (v2 - (v2 - v2))
            default:
                for fi20 := 0; fi20 < 4; fi20++ {
                    var v21 uint64 = v1
                    v22 := []int{(v3 | (-(v3))), 127, int(v2), -9223372036854775808}
                    var v23 int8 = 42
                    fmt.Println(v21, v21)
                    if (-902 <= -2147483648) {
                        v1 += 42
                    }
                    fmt.Println(v3, v21, v21)
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
            }
        }
        var v24 bool = false
        var v25 int64 = (v2 % ((-165) | 1))
        _ = fi18
        _ = v24
        _ = v25
    }
    v26 := []int{(v3 * (v3 - (v3 >> ((65535) & 63)))), 127, v3, v3}
    var v27 uint = 786
    fmt.Println(v27, v1)
    v28 := GenS1{F0: v2}
    if true {
        v2 += v2
    }
    v29 := GenS2{F0: v2, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
