package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 string }
func genFn0(p0 int) int8 {
    if false {
        var v4 int32 = -1
        v5 := []int{32767}
        p0 = -1
        _ = v4
        _ = v5
    } else {
        v6 := []int{(p0 ^ (p0 << ((0) & 63)))}
        v7 := GenS1{F0: 2, F1: -1}
        _ = v6
        _ = v7
    }
    p0 = (p0 - p0)
    if true {
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 float64 = 3.14159
            switch "\n\t" {
            case "":
                v10 := GenS1{F0: 1, F1: 2147483647}
                _ = v10
            default:
                v11 := []int{(fi8 ^ (fi8 / (((-(p0))) | 1))), p0, fi8}
                _ = v11
            }
            fmt.Println(v9, fi8)
            _ = fi8
            _ = v9
        }
    } else {
        p0 = -1
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        v13 := GenS1{F0: 800, F1: 449}
        _ = fi12
        _ = v13
    }
    v14 := []int{(-(p0)), (^(p0)), (p0 ^ (p0 % ((p0) | 1)))}
    _ = p0
    _ = v14
    return 100
}
func genFn1(p0 int8) bool {
    for fi15 := 0; fi15 < 17; fi15++ {
        v16 := GenS1{F0: 4294967295, F1: -1}
        _ = fi15
        _ = v16
    }
    v17 := []int{789, 128}
    _ = p0
    _ = v17
    return ((true || (!true)) && false)
}

func main() {
    var v1 int = -2147483648
    var v2 bool = false
    v3 := make(map[string]string)
    v1 = (^(v1))
    v1 -= (v1 ^ v1)
    v2 = true
    v2 = (1.5 < float32((3.14159 + ((((0.1 * 1.5) / 3.14159) + (((0.0 - 0.0) * (0.0 + 3.14159)) / 1.0)) + 3.14159))))
    v2 = false
    v18 := []int{-1, v1, v1, 32768}
    switch (("\n\t" + "\n\t") + "a") {
    case "foo bar":
        if v2 {
            var v19 int16 = 0
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 uint16 = 255
                switch (v1 << ((0) & 63)) {
                case 32768:
                    v22 := v3[(("\n\t" + "hello") + "\n\t")]
                    _ = v22
                case -1:
                    var v23 int8 = -1
                    _ = v23
                case 843:
                    for fi24 := 0; fi24 < 2; fi24++ {
                        var v25 uint16 = v21
                        _ = fi24
                        _ = v25
                    }
                }
                _ = fi20
                _ = v21
            }
            _ = v19
        }
    case "":
        _ = len(v18)
    }
    v2 = (v2 || (2 != 63))
    v26 := GenS1{F0: 42, F1: -1}
    v1 += (v1 >> ((342) & 63))
    v1 = (^(v1))
    for fi27 := 0; fi27 < 12; fi27++ {
        var v28 int16 = 1
        if v2 {
            var v29 float32 = 0.1
            if ((0.1 >= 1.5) || true) {
                v1 -= 353
            }
            v30 := GenS1{F0: 2147483647, F1: 42}
            _ = v29
            _ = v30
        }
        _ = fi27
        _ = v28
    }
    if v2 {
        if len(v18) > 0 {
            v18[0] = (v1 | (-(v1)))
        }
        for fi31 := 0; fi31 < 6; fi31++ {
            v32 := GenS2{F0: ("a" + ("a" + ""))}
            var v33 int64 = 0
            _ = fi31
            _ = v32
            _ = v33
        }
    } else {
        fmt.Println(v2, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
}
