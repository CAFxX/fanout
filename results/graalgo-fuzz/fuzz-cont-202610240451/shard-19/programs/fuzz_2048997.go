package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint; F2 uint64 }
type GenS3 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 int32; F1 string }

func main() {
    var v1 int16 = -644
    var v2 uint16 = 255
    v3 := make(map[int]bool)
    fmt.Println(v2)
    v4 := GenS1{F0: (v1 << ((uint(v2)) & 63)), F1: 2}
    v5 := []int{-1, 842}
    v2 = v2
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 string = "a"
        var v8 int32 = 0
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 float64 = 0.0
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := GenS1{F0: v1, F1: 650}
            var v13 float64 = -2.25
            _ = fi11
            _ = v12
            _ = v13
        }
        v9 -= (v9 - 0.0)
        _ = fi10
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 int = fi14
        switch v15 {
        case 9223372036854775807:
            _ = len(v3)
        case 499:
            v2 = 0
        case 127:
            var v16 uint8 = 255
            _ = v16
        default:
            if (false || (!(!true))) {
                var v17 uint16 = 255
                _ = v17
            } else {
                v18 := GenS1{F0: v1, F1: 147}
                for fi19 := 0; fi19 < 19; fi19++ {
                    var v20 float64 = float64(1e10)
                    _ = fi19
                    _ = v20
                }
                var v21 float32 = 1e10
                _ = v18
                _ = v21
            }
        }
        if false {
            if (!(65535 == 355)) {
                v15 -= (-(v15))
                var v22 int8 = 100
                v23 := GenS2{F0: (false || true), F1: 0, F2: 148}
                _ = v22
                _ = v23
            }
            v24 := GenS3{F0: 0, F1: v2, F2: 0}
            if true {
                for fi25 := 0; fi25 < 12; fi25++ {
                    if (-2147483648 < 1) {
                        v2 -= v2
                        v15 = (^(v15))
                        if len(v5) > 0 {
                            v5[0] = fi25
                        }
                        if (1 <= 353) {
                            v9 = v9
                            var v26 float64 = v9
                            fmt.Println(fi14, v1, v26)
                            _ = v26
                        } else {
                            var v27 int8 = 0
                            v27 = v27
                            v28 := GenS2{F0: (-222 >= int32(v1)), F1: 4294967295, F2: 344}
                            _ = v27
                            _ = v28
                        }
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    switch "" {
                    case "foo bar":
                        var v29 int16 = (-(v1))
                        _ = v29
                    default:
                        var v30 int16 = int16(v2)
                        _ = v30
                    }
                    _ = fi25
                }
                var v31 bool = (1 < 61)
                _ = v31
            }
            _ = v24
        }
        _ = fi14
        _ = v15
    }
    v2 += v2
    v32 := GenS4{F0: -34, F1: (1 <= -840)}
    v33 := GenS5{F0: 1, F1: (("foo bar" + (((("hello" + "") + "foo bar") + "") + "\n\t")) + "foo bar")}
    var v34 string = ("" + (("" + ("" + "foo bar")) + "hello"))
    var v35 int = int(v1)
    if false {
        if (true && (4294967295 <= 292)) {
            var v36 int = (v35 + v35)
            _ = v36
        } else {
            var v37 bool = true
            _ = v37
        }
    }
    switch -454 {
    case 9223372036854775807:
        var v38 uint32 = 2147483647
        _ = v38
    case 4294967295:
        v3[(v35 | 9223372036854775807)] = (!false)
    default:
        v5 = append(v5, -1)
    }
    for fi39 := 0; fi39 < 11; fi39++ {
        v40 := v3[fi39]
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v41 uint16 = v2
        _ = fi39
        _ = v40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v9
}
