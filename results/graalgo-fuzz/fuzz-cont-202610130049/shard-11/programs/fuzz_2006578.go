package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int32; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint }

func main() {
    var v1 int32 = 208
    v2 := []uint32{687}
    v3 := []string{"foo bar", "hello", "", ("" + (("a" + "hello") + "hello"))}
    var v4 int8 = 98
    v5 := make(map[int]string)
    v1 += 2147483647
    var v6 uint8 = 0
    v7 := GenS1{F0: 142, F1: 42}
    if (2 <= 4294967295) {
        v5[-166] = ("foo bar" + "foo bar")
    } else {
        v8 := GenS1{F0: 403, F1: 1}
        v5[1] = "\n\t"
        var v9 uint16 = 42
        _ = v8
        _ = v9
    }
    v4 += v4
    v4 = (v4 | 0)
    if (64 < 1) {
        v10 := GenS1{F0: -32768, F1: 0}
        for fi11 := 0; fi11 < 4; fi11++ {
            fmt.Println(v1)
            var v12 uint8 = (v6 << ((4294967295) & 63))
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        for fi13 := 0; fi13 < 3; fi13++ {
            if (!true) {
                for fi14 := 0; fi14 < 14; fi14++ {
                    delete(v5, (fi14 / ((fi14) | 1)))
                    v15 := GenS2{F0: (v1 >> ((18446744073709551615) & 63)), F1: (v1 * 2147483647), F2: 881}
                    _ = fi14
                    _ = v15
                }
                v5[(fi13 % (((fi13 ^ (^(fi13)))) | 1))] = "foo bar"
                var v16 int8 = v4
                _ = v16
            } else {
                for fi17 := 0; fi17 < 12; fi17++ {
                    v1 -= 42
                    v5[(fi13 & 32768)] = (("a" + "hello") + "foo bar")
                    _ = fi17
                }
            }
            _ = fi13
        }
        if true {
            v1 -= 0
            v18 := GenS2{F0: (^(v1)), F1: v1, F2: 117}
            _ = v18
        }
    }
    switch 9223372036854775807 {
    case 364:
        var v19 int8 = (v4 % ((v4) | 1))
        _ = v19
    }
    var v20 float64 = (1e10 / -2.25)
    v21 := GenS3{F0: 1}
    v5[9223372036854775807] = "a"
    var v22 float32 = 3.14159
    var v23 uint8 = (v6 / ((2) | 1))
    for fi24 := 0; fi24 < 2; fi24++ {
        var v25 uint16 = 0
        _ = fi24
        _ = v25
    }
    if (0 > 851) {
        var v26 string = "a"
        var v27 int64 = 42
        _ = v26
        _ = v27
    }
    v5[439] = ("" + "a")
    v28 := GenS2{F0: v1, F1: (v1 & -879), F2: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
