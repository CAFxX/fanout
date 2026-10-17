package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS5 struct { F0 uint16; F1 uint16; F2 int16 }
func genFn0() int64 {
    v7 := []int{32768}
    var v8 bool = (42 < -9223372036854775808)
    var v9 uint8 = 127
    v9 += 255
    for fi10 := 0; fi10 < 4; fi10++ {
        fmt.Println(fi10)
        if (-1 != 2147483647) {
            for fi11 := 0; fi11 < 9; fi11++ {
                if len(v7) > 0 {
                    v7[0] = fi10
                }
                switch "foo bar" {
                case "foo bar":
                    var v12 int64 = -9223372036854775808
                    _ = v12
                case "":
                    fmt.Println(v9, v8)
                default:
                    if true {
                        var v13 uint64 = 0
                        v14 := GenS1{F0: 1, F1: "hello"}
                        _ = v13
                        _ = v14
                    }
                }
                _ = fi11
            }
            switch (fi10 - fi10) {
            case -9223372036854775808:
                v9 = v9
            case 32767:
                switch 128 {
                case -879:
                    if len(v7) > 0 {
                        v7[0] = (fi10 + fi10)
                    }
                case 746:
                    v15 := GenS1{F0: 507, F1: "foo bar"}
                    _ = v15
                }
            default:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
            var v16 uint8 = v9
            _ = v16
        } else {
            for fi17 := 0; fi17 < 9; fi17++ {
                switch (("\n\t" + ("a" + "hello")) + ("a" + "")) {
                case "hello":
                    v8 = false
                case "foo bar":
                    if v8 {
                        if len(v7) > 0 {
                            v7[0] = (fi10 | fi17)
                        }
                        if v8 {
                            var v18 int8 = 127
                            _ = v18
                        } else {
                            fmt.Println(fi17, v8)
                            v7 = append(v7, 128)
                        }
                    } else {
                        var v19 int = fi17
                        v20 := GenS1{F0: 32767, F1: ""}
                        var v21 float64 = (0.0 - 3.14159)
                        if len(v7) > 0 {
                            v7[0] = fi17
                        }
                        v22 := GenS2{F0: 1000, F1: -9223372036854775808}
                        _ = v19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                case "a":
                    var v23 float64 = 0.1
                    _ = v23
                default:
                    fmt.Println(fi17, v9, v9)
                }
                v24 := GenS3{F0: (1 <= 2), F1: "foo bar"}
                switch ("foo bar" + "") {
                case "\n\t":
                    var v25 float32 = (float32(float64(-2.25)) - ((0.0 - (-2.25 * -2.25)) / 1.0))
                    _ = v25
                }
                var v26 float32 = -2.25
                v9 = 1
                if ((true || (-720 > 324)) || v8) {
                    v26 -= -2.25
                } else {
                    v27 := GenS2{F0: 32767, F1: 9223372036854775807}
                    var v28 uint8 = v9
                    _ = v27
                    _ = v28
                }
                _ = fi17
                _ = v24
                _ = v26
            }
        }
        _ = fi10
    }
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int64, p1 int64) uint8 {
    for fi29 := 0; fi29 < 16; fi29++ {
        p0 = 604
        p1 = genFn0()
        if false {
            p0 = -401
        }
        _ = fi29
    }
    v30 := []int{-2147483648, -895}
    for fi31 := 0; fi31 < 1; fi31++ {
        if (!("a" <= "")) {
            var v32 uint8 = 114
            fmt.Println(v32, p1, p1)
            _ = v32
        }
        _ = fi31
    }
    var v33 int16 = 0
    var v34 int = 4294967295
    _ = p0
    _ = p1
    _ = v30
    _ = v33
    _ = v34
    return 127
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = -96
    v3 := []int{}
    var v4 uint = 42
    var v5 int16 = 291
    var v6 float32 = (1.5 / -2.25)
    if false {
        v2 += v2
        fmt.Println(v6, v6)
    }
    v2 += v2
    fmt.Println(v2)
    if true {
        for fi35 := 0; fi35 < 12; fi35++ {
            v36 := GenS3{F0: true, F1: "\n\t"}
            _ = fi35
            _ = v36
        }
        v37 := GenS4{F0: 2147483647, F1: 0, F2: 65535}
        _ = v37
    } else {
        var v38 int64 = 42
        for fi39 := 0; fi39 < 14; fi39++ {
            delete(v1, 42)
            _ = fi39
        }
        _ = v38
    }
    fmt.Println(v4)
    if (4294967295 < 0) {
        v40 := GenS1{F0: (v5 * int16(v2)), F1: ""}
        v41 := GenS2{F0: 42, F1: 726}
        var v42 float64 = 3.14159
        _ = v40
        _ = v41
        _ = v42
    } else {
        v4 = (v4 << ((4294967295) & 63))
    }
    var v43 uint16 = 0
    var v44 uint64 = 247
    v45 := GenS5{F0: (v43 - uint16(v44)), F1: v43, F2: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
