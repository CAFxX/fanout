package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int64 }
func genFn0(p0 int32, p1 int16, p2 int64) uint16 {
    for fi7 := 0; fi7 < 18; fi7++ {
        if true {
            v8 := GenS1{F0: (p2 ^ (-(p2)))}
            var v9 int8 = 69
            _ = v8
            _ = v9
        } else {
            switch (fi7 + fi7) {
            case -473:
                v10 := GenS2{F0: 42, F1: p2}
                _ = v10
            case 128:
                switch (fi7 % ((fi7) | 1)) {
                case -1:
                    p1 = 32767
                default:
                    var v11 int64 = p2
                    _ = v11
                }
            }
        }
        _ = fi7
    }
    fmt.Println(p0, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    return 87
}
func genFn1(p0 uint32) uint {
    for fi12 := 0; fi12 < 11; fi12++ {
        if true {
            var v13 float64 = 0.0
            if false {
                var v14 float32 = ((float32(3.14159) * float32(((v13 / v13) - v13))) * (-2.25 * ((1e10 - float32(1.5)) * -2.25)))
                v14 += 1.5
                switch 2147483647 {
                case -888:
                    v15 := GenS1{F0: 1}
                    _ = v15
                case 42:
                    v16 := []int{(^(fi12)), 42}
                    _ = v16
                default:
                    v17 := GenS2{F0: 127, F1: -1}
                    _ = v17
                }
                _ = v14
            }
            _ = v13
        } else {
            if false {
                p0 += 2147483647
            }
            fmt.Println(p0, p0)
        }
        _ = fi12
    }
    var v18 bool = (false && (((9223372036854775807 != 42) || (42 < -571)) && true))
    v19 := []int{-9223372036854775808}
    var v20 string = "\n\t"
    _ = p0
    _ = v18
    _ = v19
    _ = v20
    return 0
}

func main() {
    var v1 int64 = -202
    var v2 int8 = 11
    v3 := []string{"foo bar", "a"}
    var v4 uint8 = 130
    var v5 string = (("\n\t" + "hello") + "hello")
    var v6 int8 = -1
    var v21 int8 = v6
    v4 = v4
    switch 0 {
    case -2147483648:
        v22 := GenS1{F0: v1}
        _ = v22
    default:
        var v23 int16 = 0
        _ = v23
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        if true {
            for fi25 := 0; fi25 < 16; fi25++ {
                fmt.Println(v6, fi24)
                _ = fi25
            }
        }
        v6 = 1
        var v26 uint = 255
        _ = fi24
        _ = v26
    }
    fmt.Println(v1)
    if true {
        v27 := GenS2{F0: v21, F1: v1}
        v5 = ""
        _ = v27
    }
    v28 := GenS2{F0: (-(v21)), F1: 419}
    v21 = -40
    v29 := GenS2{F0: v6, F1: (v1 | -9223372036854775808)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
