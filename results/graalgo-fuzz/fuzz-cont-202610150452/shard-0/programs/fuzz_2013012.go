package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int8, p1 int32) uint16 {
    p0 = (p0 | 100)
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 int32 = (p1 | p1)
        p1 += (v6 - (v6 * (p1 ^ p1)))
        if (!true) {
            var v7 float32 = ((((0.0 / 1.0) - ((1e10 / -2.25) + float32(3.14159))) / (0.1 + 0.1)) - 3.14159)
            if true {
                switch 9223372036854775807 {
                case 128:
                    fmt.Println(v6, p0)
                case 568:
                    v7 += 0.1
                case 32768:
                    var v8 uint64 = 2
                    _ = v8
                }
                for fi9 := 0; fi9 < 11; fi9++ {
                    fmt.Println(v6, p1, fi9)
                    p1 -= (v6 << ((4294967295) & 63))
                    _ = fi9
                }
            } else {
                v10 := GenS1{F0: -9223372036854775808, F1: 171}
                switch -2147483648 {
                case -40:
                    var v11 float32 = v7
                    _ = v11
                }
                _ = v10
            }
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS1{F0: -1, F1: 255}
        if true {
            v14 := []int{}
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    p0 -= (p0 * (p0 ^ 127))
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 float32 = 0.0
        if false {
            for fi17 := 0; fi17 < 8; fi17++ {
                v16 += float32((((-2.25 - 0.0) + 1e10) * 3.14159))
                if true {
                    v18 := []int{}
                    var v19 uint64 = 18446744073709551615
                    _ = len(v18)
                    _ = v18
                    _ = v19
                }
                if true {
                    v20 := []int{}
                    if false {
                        fmt.Println(v16, fi15)
                    }
                    _ = v20
                }
                _ = fi17
            }
            p1 += (p1 >> ((18446744073709551615) & 63))
            if true {
                p0 = p0
            } else {
                var v21 uint16 = 255
                _ = v21
            }
        }
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    return 17
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{1, 880, 2147483648, 2}
    var v3 uint64 = 451
    var v4 int16 = -418
    v22 := v1["foo bar"]
    var v23 uint8 = 150
    v4 = 32767
    v3 = 31
    if true {
        var v24 uint8 = 127
        var v25 int64 = 1
        _ = v24
        _ = v25
    } else {
        if v22 {
            fmt.Println(v23)
            _ = len(v2)
        }
    }
    v26 := GenS1{F0: -1, F1: 18446744073709551615}
    if ((-1 != 0) || (2147483647 > -982)) {
        if (!(0.1 >= 0.0)) {
            v22 = ((!true) || (((1 >= 4294967295) || (960 >= 0)) || (-552 != 314)))
        } else {
            v22 = true
        }
        delete(v1, "foo bar")
        fmt.Println(v23)
    }
    v22 = (uint(v3) < 789)
    for fi27 := 0; fi27 < 10; fi27++ {
        var v28 uint = 636
        v23 = 0
        _ = fi27
        _ = v28
    }
    if (!v22) {
        v29 := GenS2{F0: 508}
        var v30 int32 = 2147483647
        var v31 int16 = (v4 << ((255) & 63))
        _ = v29
        _ = v30
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
}
