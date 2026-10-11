package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int16; F1 int64; F2 int16 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 string; F1 bool; F2 uint32 }

func main() {
    var v1 uint64 = 363
    var v2 float32 = (3.14159 * ((1.5 - 1.5) / 1.0))
    v3 := make(map[int]bool)
    var v4 uint = 18446744073709551615
    var v5 uint16 = 1
    var v6 int16 = 916
    v7 := GenS1{F0: ("hello" + "hello"), F1: 2147483647}
    v8 := []int{}
    v2 = (v2 * (((float32(((1e10 / (-2.25 * 0.1)) - 0.0)) / v2) + (3.14159 / 0.1)) * v2))
    var v9 float32 = v2
    v2 += 1.5
    v9 += 1e10
    var v10 int16 = v6
    v11 := v3[-121]
    for fi12 := 0; fi12 < 14; fi12++ {
        if v11 {
            if len(v8) > 0 {
                v8[0] = -278
            }
            switch ("\n\t" + ("\n\t" + "a")) {
            case "hello":
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            case "\n\t":
                for fi13 := 0; fi13 < 4; fi13++ {
                    v4 -= (^(v4))
                    fmt.Println(v5)
                    _ = fi13
                }
            default:
                v10 = v6
            }
            v14 := GenS1{F0: "foo bar", F1: (fi12 * -743)}
            _ = v14
        } else {
            var v15 uint64 = (v1 | 118)
            _ = v15
        }
        var v16 int8 = 100
        _ = fi12
        _ = v16
    }
    v9 = 0.0
    var v17 uint8 = 127
    for fi18 := 0; fi18 < 4; fi18++ {
        v19 := GenS1{F0: "a", F1: fi18}
        _ = fi18
        _ = v19
    }
    var v20 float32 = (v9 + v2)
    var v21 uint = (v4 - (v4 - 255))
    if v11 {
        var v22 int32 = 2147483647
        for fi23 := 0; fi23 < 6; fi23++ {
            v2 += (0.1 * 0.0)
            if (((v5 - v5) <= v5) && false) {
                if v11 {
                    v24 := GenS2{F0: -867, F1: -9223372036854775808, F2: v6}
                    var v25 int16 = (v10 ^ v10)
                    v26 := v3[(fi23 % ((fi23) | 1))]
                    _ = v24
                    _ = v25
                    _ = v26
                }
                if len(v8) > 0 {
                    v8[0] = (fi23 % (((-(fi23))) | 1))
                }
                switch ("hello" + ("" + "hello")) {
                case "hello":
                    var v27 bool = (!(!((!(((-2.25 - float64(v2)) * float64(v9)) != (((1.5 - 0.0) - 0.1) * 1e10))) || ((v11 && (!(true && (v11 && ((v5 | (v5 + v5)) == v5))))) || (float64(v9) == 1.5)))))
                    _ = v27
                case "foo bar":
                    _ = len(v8)
                default:
                    if (v11 || ((988 > v5) || v11)) {
                        switch (("foo bar" + "") + "") {
                        case "foo bar":
                            v28 := GenS3{F0: 352, F1: v4}
                            _ = v28
                        default:
                            if (uint8(v1) < (^(v17))) {
                                fmt.Println(v22, v20, v2)
                                v8 = append(v8, fi23)
                            }
                        }
                        _ = len(v8)
                    }
                }
            }
            if v11 {
                if (0.1 <= 1.5) {
                    v21 = (^(v4))
                } else {
                    for fi29 := 0; fi29 < 11; fi29++ {
                        _ = len(v3)
                        _ = fi29
                    }
                }
                var v30 int32 = (v22 + (v22 - v22))
                v31 := GenS4{F0: "foo bar", F1: v11, F2: 4294967295}
                _ = v30
                _ = v31
            }
            _ = fi23
        }
        _ = v22
    } else {
        fmt.Println(v17, v4, v9)
        if false {
            for fi32 := 0; fi32 < 6; fi32++ {
                v33 := v3[fi32]
                if v11 {
                    _ = len(v3)
                    _ = len(v3)
                    switch (fi32 | fi32) {
                    case 128:
                        var v34 int = (-(fi32))
                        _ = v34
                    case -9223372036854775808:
                        _ = len(v8)
                    case -57:
                        if true {
                            var v35 uint16 = v5
                            var v36 int64 = -207
                            _ = v35
                            _ = v36
                        }
                    }
                } else {
                    v17 -= 0
                    _ = len(v8)
                }
                _ = fi32
                _ = v33
            }
            var v37 float32 = -2.25
            var v38 bool = true
            _ = v37
            _ = v38
        } else {
            switch 9223372036854775807 {
            case 128:
                v39 := GenS4{F0: "a", F1: v11, F2: 1}
                _ = v39
            case 0:
                var v40 uint = (v4 | (v21 / (((v21 ^ 43)) | 1)))
                _ = v40
            }
            v5 += v5
        }
    }
    var v41 uint8 = (v17 % (((v17 + (v17 + (v17 / ((255) | 1))))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
