package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint32, p1 bool, p2 uint8) uint64 {
    v7 := GenS1{F0: (p2 / ((0) | 1)), F1: "", F2: -71}
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    p2 = (p2 - 203)
    v10 := GenS1{F0: (p2 & 2), F1: "foo bar", F2: 7}
    switch 1 {
    case 32768:
        v11 := GenS1{F0: p2, F1: (("" + ("" + "")) + ""), F2: 100}
        _ = v11
    default:
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 float32 = ((-2.25 - 1e10) - float32(0.0))
            v13 -= v13
            _ = fi12
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 276
}
func genFn1(p0 uint64, p1 uint16, p2 string) int {
    p2 = "foo bar"
    v14 := []int{-646, -68, -473, 293}
    var v15 int = 1
    switch (("\n\t" + (("\n\t" + (("hello" + p2) + (((("foo bar" + p2) + (("a" + (p2 + p2)) + (p2 + ""))) + ("" + p2)) + "\n\t"))) + p2)) + (p2 + p2)) {
    case "foo bar":
        var v16 uint8 = 255
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return (v15 - (v15 * -9223372036854775808))
}

func main() {
    v1 := make(map[int]int)
    var v2 float32 = (1e10 / 0.1)
    v3 := make(map[uint32]int)
    var v4 int64 = 0
    v5 := []int{0, 0, 128}
    var v6 int16 = 0
    switch 32767 {
    case 0:
        v17 := GenS2{F0: 2147483647}
        _ = v17
    }
    var v18 int = 337
    fmt.Println(v6, v2, v6)
    var v19 bool = (1 == 32767)
    for fi20 := 0; fi20 < 19; fi20++ {
        v19 = true
        for fi21 := 0; fi21 < 9; fi21++ {
            if (v19 && false) {
                v22 := GenS3{F0: 42}
                _ = v22
            } else {
                v2 -= -2.25
            }
            switch 32768 {
            case -606:
                if v19 {
                    var v23 int = 128
                    if v19 {
                        v5 = append(v5, int(v6))
                        var v24 uint64 = 42
                        _ = v24
                    }
                    var v25 int = genFn1(449, 608, "a")
                    var v26 string = "a"
                    var v27 uint64 = 2
                    v6 += v6
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v23
                    _ = v25
                    _ = v26
                    _ = v27
                } else {
                    v28 := GenS3{F0: 705}
                    v2 -= v2
                    _ = v28
                }
            }
            v6 -= v6
            _ = fi21
        }
        v19 = ("" <= ("\n\t" + "a"))
        _ = fi20
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        v18 = genFn1(1, 271, "foo bar")
        _ = fi29
    }
    v2 += v2
    if (false || v19) {
        v18 = v18
    }
    v30 := GenS3{F0: 0}
    var v31 string = ("a" + "foo bar")
    v5 = append(v5, (^(v18)))
    if ((2 < 593) || (v19 || v19)) {
        v6 += v6
        if (((false || (0 <= 1)) || v19) && (!(42 <= 118))) {
            var v32 uint32 = 890
            var v33 uint32 = v32
            var v34 int = (v18 * 10)
            _ = v32
            _ = v33
            _ = v34
        } else {
            if ((v6 & v6) > (v6 - (v6 / ((980) | 1)))) {
                switch v18 {
                case 32768:
                    for fi35 := 0; fi35 < 3; fi35++ {
                        var v36 float64 = 0.1
                        for fi37 := 0; fi37 < 2; fi37++ {
                            var v38 int64 = v4
                            fmt.Println(fi35)
                            _ = fi37
                            _ = v38
                        }
                        _ = fi35
                        _ = v36
                    }
                default:
                    var v39 int32 = 0
                    _ = v39
                }
                var v40 uint8 = 194
                _ = v40
            }
        }
        var v41 int64 = (v4 % ((0) | 1))
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
