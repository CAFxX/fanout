package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 bool; F2 int8 }
func genFn0() uint64 {
    var v5 int8 = 1
    v6 := []int{-2147483648, -1, 127}
    _ = len(v6)
    v5 = (v5 & (v5 / ((v5) | 1)))
    v7 := GenS1{F0: 0}
    _ = v5
    _ = v6
    _ = v7
    return 971
}
func genFn1(p0 int, p1 uint64, p2 uint8) uint16 {
    switch (p0 - p0) {
    case -849:
        if (-43 > 127) {
            for fi8 := 0; fi8 < 9; fi8++ {
                for fi9 := 0; fi9 < 10; fi9++ {
                    switch ("\n\t" + ("a" + "foo bar")) {
                    case "\n\t":
                        p1 -= (^(p1))
                    default:
                        if false {
                            var v10 uint8 = p2
                            _ = v10
                        }
                    }
                    p0 += 32768
                    if false {
                        if false {
                            switch fi9 {
                            case 1:
                                var v11 uint16 = 32767
                                _ = v11
                            default:
                                v12 := []int{(fi9 * fi9), 0}
                                p1 = genFn0()
                                _ = v12
                            }
                        }
                    }
                    v13 := GenS2{F0: ("" + "a"), F1: ("" == "\n\t"), F2: 100}
                    _ = fi9
                    _ = v13
                }
                fmt.Println(p0, fi8, fi8)
                _ = fi8
            }
            for fi14 := 0; fi14 < 10; fi14++ {
                v15 := []int{4294967295, 89}
                switch 816 {
                case 0:
                    p1 += p1
                case 32768:
                    fmt.Println(p1, p2)
                }
                _ = fi14
                _ = v15
            }
        } else {
            p0 = p0
            v16 := []int{-1, (p0 ^ p0)}
            _ = v16
        }
    default:
        v17 := []int{p0}
        _ = v17
    }
    var v18 bool = ((((p0 | p0) > (p0 - (p0 - p0))) && (true || (true || true))) || false)
    var v19 int32 = 155
    v20 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    _ = v20
    return 255
}

func main() {
    v1 := []int{128, 433, -2147483648}
    var v2 int16 = 42
    v3 := make(map[int]uint32)
    var v4 string = ("foo bar" + "hello")
    for fi21 := 0; fi21 < 1; fi21++ {
        v22 := GenS2{F0: v4, F1: ((!false) || (!(true || (0 <= 42)))), F2: -6}
        var v23 float64 = (((1.5 + 1.5) / 1.0) + 1e10)
        if len(v1) > 0 {
            v1[0] = (fi21 - fi21)
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    v2 -= (v2 - v2)
    var v24 uint8 = 2
    v4 = ""
    var v25 float64 = -2.25
    var v26 string = "foo bar"
    v27 := GenS1{F0: 42}
    v25 = 0.1
    switch 32767 {
    case -2147483648:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        var v28 uint64 = 4294967295
        _ = v28
    }
    v29 := GenS2{F0: "a", F1: ((^(v24)) > 1), F2: 127}
    v30 := GenS1{F0: -794}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
