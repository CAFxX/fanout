package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 int64; F2 uint }
func genFn0(p0 int64, p1 uint32, p2 uint16) uint8 {
    if true {
        v5 := GenS1{F0: "foo bar", F1: p2}
        _ = v5
    }
    p2 = 0
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 uint32 = 637
    var v2 uint = 18446744073709551615
    v3 := []int{127, -985}
    var v4 int8 = -48
    v4 = (v4 * (-(v4)))
    v1 = v1
    v2 = v2
    for fi6 := 0; fi6 < 15; fi6++ {
        if (true || false) {
            if (1e10 < 3.14159) {
                v7 := GenS2{F0: 65535, F1: genFn0(567, v1, 42)}
                _ = v7
            } else {
                if false {
                    var v8 int64 = 42
                    fmt.Println(v4, fi6)
                    _ = v8
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        var v9 float32 = 1.5
                        _ = v9
                    case "a":
                        var v10 float64 = (1.5 + (-2.25 - -2.25))
                        _ = v10
                    default:
                        if false {
                            switch (^(fi6)) {
                            case 434:
                                _ = len(v3)
                            default:
                                for fi11 := 0; fi11 < 10; fi11++ {
                                    fmt.Println(v4, v4, fi6)
                                    if true {
                                        var v12 int8 = v4
                                        if true {
                                            v13 := GenS2{F0: v2, F1: 128}
                                            fmt.Println(v4, fi11, v2)
                                            _ = v13
                                        } else {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v14 := GenS2{F0: v2, F1: genFn0(865, v1, 255)}
                                            v15 := GenS1{F0: ("hello" + ""), F1: 65535}
                                            _ = v14
                                            _ = v15
                                        }
                                        _ = v12
                                    }
                                    _ = fi11
                                }
                            }
                            var v16 int64 = -9223372036854775808
                            if (false || true) {
                                var v17 uint64 = 2
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v17
                            }
                            _ = v16
                        }
                    }
                }
                var v18 float32 = (((1e10 + (-2.25 / ((float32(((1.5 / 3.14159) * 1e10)) * ((-2.25 + -2.25) * float32((1e10 / 0.1)))) / 1.0))) + -2.25) / 0.1)
                _ = v18
            }
        }
        v19 := GenS3{F0: fi6}
        _ = fi6
        _ = v19
    }
    if ((255 != 0) || (3.14159 <= 3.14159)) {
        var v20 float32 = 0.0
        for fi21 := 0; fi21 < 13; fi21++ {
            var v22 int64 = 732
            _ = fi21
            _ = v22
        }
        if (false || (!true)) {
            v23 := GenS3{F0: 536}
            v24 := GenS1{F0: "", F1: 1}
            _ = v23
            _ = v24
        } else {
            switch 714 {
            case 982:
                v20 -= 1.5
            default:
                fmt.Println(v1, v4)
            }
            fmt.Println(v20, v4)
            v25 := GenS4{F0: "", F1: 51}
            _ = v25
        }
        _ = v20
    }
    var v26 uint16 = 255
    fmt.Println(v4)
    var v27 uint = 865
    v28 := GenS5{F0: 1}
    v29 := GenS6{F0: (v4 / (((-(v4))) | 1)), F1: 42, F2: v27}
    if len(v3) > 0 {
        v3[0] = -227
    }
    var v30 uint16 = (^(v26))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
