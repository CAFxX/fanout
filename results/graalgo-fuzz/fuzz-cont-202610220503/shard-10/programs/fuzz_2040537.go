package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint, p1 uint16) int8 {
    if true {
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 int = -2147483648
            _ = fi6
            _ = v7
        }
        fmt.Println(p1, p1)
    }
    p1 = (p1 + 42)
    if false {
        v8 := GenS1{F0: 818, F1: 133, F2: 53}
        if false {
            switch 128 {
            case -2147483648:
                p0 = 996
            case 188:
                var v9 string = ("\n\t" + "hello")
                _ = v9
            case 4294967295:
                p1 += (^(p1))
            }
        } else {
            switch "" {
            case "foo bar":
                if true {
                    p1 = (^(p1))
                    var v10 int16 = 42
                    fmt.Println(v10)
                    _ = v10
                }
            case "a":
                fmt.Println(p0)
            case "":
                var v11 uint64 = 42
                _ = v11
            default:
                fmt.Println(p1, p1)
            }
            for fi12 := 0; fi12 < 11; fi12++ {
                p0 += (p0 & (^(p0)))
                _ = fi12
            }
        }
        v13 := []int{32768, -662, 1}
        _ = v13
        _ = v8
    }
    v14 := []int{}
    _ = p0
    _ = p1
    _ = v14
    return -1
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]bool)
    var v3 uint32 = 4294967295
    var v4 uint8 = 38
    var v5 uint64 = 4294967295
    var v15 uint64 = 0
    for fi16 := 0; fi16 < 6; fi16++ {
        v3 = v3
        _ = fi16
    }
    var v17 float32 = v1
    if true {
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 int8 = 1
            _ = fi18
            _ = v19
        }
    } else {
        for fi20 := 0; fi20 < 3; fi20++ {
            switch "foo bar" {
            case "":
                v2[(fi20 / ((fi20) | 1))] = false
            default:
                var v21 bool = (true && true)
                _ = v21
            }
            v17 = v1
            _ = fi20
        }
        var v22 uint32 = v3
        for fi23 := 0; fi23 < 17; fi23++ {
            v17 -= v1
            fmt.Println(v4)
            v5 += 268
            _ = fi23
        }
        _ = v22
    }
    v24 := GenS1{F0: 42, F1: (^(v4)), F2: -32768}
    v1 = v17
    v25 := GenS1{F0: 831, F1: (v4 >> ((65535) & 63)), F2: 0}
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 uint64 = (^(v15))
        v28 := GenS1{F0: -1, F1: v4, F2: -709}
        var v29 int64 = -9223372036854775808
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    if ((("" + "foo bar") + ("a" + "\n\t")) != ("" + ("\n\t" + "foo bar"))) {
        var v30 float32 = 0.1
        _ = v30
    }
    v31 := GenS2{F0: v4}
    fmt.Println(v4, v5, v5)
    var v32 uint = 778
    v33 := []int{127}
    fmt.Println(v5, v3)
    var v34 uint32 = uint32(v15)
    v1 = float32(0.1)
    fmt.Println(v15)
    v17 = float32(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
