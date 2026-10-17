package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 bool; F2 int8 }
func genFn0(p0 int, p1 uint32, p2 bool) uint16 {
    v6 := []int{32768}
    if true {
        var v7 uint = 1
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        fmt.Println(p2, p1, p1)
        _ = fi8
    }
    var v9 int8 = -1
    var v10 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v9
    return 42
}
func genFn1(p0 uint8, p1 uint32, p2 int64) int16 {
    fmt.Println(p2, p2)
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 uint8 = 128
    var v2 int = 594
    var v3 float64 = 1.5
    v4 := []int{v2}
    var v5 uint16 = 382
    v1 = (^(v1))
    v2 -= v2
    v11 := GenS1{F0: v5}
    v3 -= v3
    v2 = v2
    switch 9223372036854775807 {
    case 9223372036854775807:
        v12 := GenS1{F0: genFn0(v2, 4294967295, false)}
        _ = v12
    case 128:
        switch (v2 >> ((1) & 63)) {
        case 9223372036854775807:
            if true {
                switch 0 {
                case 4294967295:
                    switch (("" + "hello") + ("\n\t" + "foo bar")) {
                    case "":
                        v1 += 1
                    case "foo bar":
                        var v13 float32 = 0.0
                        _ = v13
                    default:
                        var v14 bool = ((v2 & 179) < v2)
                        _ = v14
                    }
                default:
                    _ = len(v4)
                }
                fmt.Println(v5)
            }
        case 564:
            v15 := GenS2{F0: v5, F1: true, F2: 42}
            _ = v15
        }
    case 225:
        if len(v4) > 0 {
            v4[0] = v2
        }
    default:
        fmt.Println(v5)
    }
    if (v3 > (0.0 - (float64((3.14159 + (0.0 - 3.14159))) + (1.5 / v3)))) {
        switch ("hello" + ("foo bar" + "hello")) {
        case "foo bar":
            v2 -= v2
        case "\n\t":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            var v16 string = ""
            _ = v16
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            v1 += v1
            var v18 float64 = (v3 * (v3 - 3.14159))
            _ = fi17
            _ = v18
        }
        var v19 bool = (288 < 743)
        _ = v19
    }
    var v20 bool = true
    var v21 float32 = (0.1 / 1.0)
    v3 = 3.14159
    var v22 int = (v2 / ((v2) | 1))
    for fi23 := 0; fi23 < 15; fi23++ {
        for fi24 := 0; fi24 < 18; fi24++ {
            v25 := GenS2{F0: v5, F1: (42 != 1), F2: -128}
            v26 := GenS1{F0: (v5 * (v5 ^ (v5 ^ 987)))}
            switch "\n\t" {
            case "a":
                v1 += 127
            case "\n\t":
                var v27 uint64 = 42
                _ = v27
            default:
                fmt.Println(v22)
            }
            _ = fi24
            _ = v25
            _ = v26
        }
        for fi28 := 0; fi28 < 15; fi28++ {
            v22 = fi23
            _ = fi28
        }
        _ = fi23
    }
    switch (v22 * (-(v22))) {
    case 32768:
        var v29 uint8 = v1
        _ = v29
    case 1:
        var v30 string = "foo bar"
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
