package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint }
func genFn0(p0 int32, p1 uint, p2 bool) uint8 {
    if ((42 == -9223372036854775808) && false) {
        for fi6 := 0; fi6 < 1; fi6++ {
            fmt.Println(fi6, p0)
            _ = fi6
        }
        p2 = p2
    }
    switch "\n\t" {
    case "a":
        if p2 {
            v7 := []int{int(p1), 1, 32768, 42}
            _ = v7
        }
    default:
        if (p2 && (127 <= 6)) {
            switch "foo bar" {
            case "":
                var v8 float64 = 0.0
                _ = v8
            }
            p0 -= (p0 + (-(p0)))
        } else {
            for fi9 := 0; fi9 < 12; fi9++ {
                var v10 uint64 = 2
                switch "" {
                case "foo bar":
                    if (65535 <= 615) {
                        fmt.Println(p1)
                        var v11 uint8 = 42
                        _ = v11
                    } else {
                        p2 = (9223372036854775807 != 724)
                        p2 = p2
                    }
                }
                _ = fi9
                _ = v10
            }
            for fi12 := 0; fi12 < 20; fi12++ {
                p0 = 2147483647
                _ = fi12
            }
            var v13 string = ""
            _ = v13
        }
    }
    p0 = (p0 / ((p0) | 1))
    v14 := []int{0, -1}
    var v15 bool = (1 == 42)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 128
}

func main() {
    var v1 uint32 = 248
    var v2 int8 = 127
    v3 := make(map[int]int)
    v4 := []string{}
    v5 := []uint32{(^(v1)), (v1 - (v1 - (v1 | (v1 >> ((65535) & 63)))))}
    var v16 int16 = -284
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := v3[fi17]
        _ = fi17
        _ = v18
    }
    if false {
        for fi19 := 0; fi19 < 13; fi19++ {
            v20 := GenS1{F0: 255}
            v1 = v1
            _ = fi19
            _ = v20
        }
        fmt.Println(v2)
        if true {
            for fi21 := 0; fi21 < 6; fi21++ {
                v22 := GenS1{F0: 2}
                _ = fi21
                _ = v22
            }
        } else {
            fmt.Println(v2, v2, v16)
        }
    } else {
        if true {
            var v23 uint = 2
            _ = v23
        } else {
            var v24 uint16 = 42
            v25 := GenS2{F0: (!(65535 == 859)), F1: 106}
            _ = v24
            _ = v25
        }
        fmt.Println(v1, v16, v2)
    }
    v5 = append(v5, v1)
    v3[42] = 9223372036854775807
    switch "\n\t" {
    case "a":
        var v26 int8 = v2
        _ = v26
    }
    fmt.Println(v2)
    v27 := GenS2{F0: (28 >= 2147483647), F1: 65535}
    v2 = v2
    v16 = (v16 + v16)
    var v28 int32 = -20
    if (0.1 != -2.25) {
        switch ("hello" + "\n\t") {
        case "":
            if len(v5) > 0 {
                v5[0] = (v1 & 65535)
            }
        case "foo bar":
            v1 -= 623
        }
    } else {
        v29 := GenS2{F0: (952 < 0), F1: 255}
        _ = v29
    }
    v16 = v16
    v28 = 654
    v2 = (v2 ^ -85)
    v2 -= v2
    if false {
        var v30 uint8 = 128
        var v31 int8 = (v2 ^ v2)
        _ = v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
