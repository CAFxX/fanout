package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int16 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 string; F1 int16; F2 uint16 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS5 struct { F0 int8; F1 uint8; F2 int64 }
func genFn0() uint16 {
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint = 58
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 int8 = -128
            fmt.Println(fi9)
            fmt.Println(fi9, v11)
            _ = fi10
            _ = v11
        }
        var v12 int32 = -1
        _ = fi9
        _ = v12
    }
    return 674
}
func genFn1(p0 bool, p1 int64) uint32 {
    p0 = true
    v13 := []int{-2147483648, -685, 4294967295, 32767}
    var v14 uint16 = 32767
    var v15 int8 = 42
    switch -272 {
    case 2147483647:
        if p0 {
            p0 = (-2.25 != 1.5)
            v15 = -1
            for fi16 := 0; fi16 < 17; fi16++ {
                v14 -= 42
                _ = fi16
            }
        }
    case 32768:
        v17 := GenS1{F0: v14, F1: p1, F2: 38}
        _ = v17
    default:
        fmt.Println(v15)
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 174
    var v3 float32 = 3.14159
    var v4 int16 = 1
    var v5 int32 = 2147483647
    v6 := []string{"foo bar"}
    if true {
        v18 := GenS2{F0: false, F1: "a"}
        v19 := GenS1{F0: 255, F1: 1, F2: v4}
        v3 = float32(float64(-2.25))
        _ = v18
        _ = v19
    } else {
        for fi20 := 0; fi20 < 8; fi20++ {
            for fi21 := 0; fi21 < 8; fi21++ {
                v2 -= (v2 << ((18446744073709551615) & 63))
                var v22 uint32 = 2147483647
                _ = fi21
                _ = v22
            }
            v23 := v1[fi20]
            _ = fi20
            _ = v23
        }
        var v24 uint64 = 42
        v4 = v4
        _ = v24
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        var v26 bool = true
        _ = len(v6)
        fmt.Println(v3, v2)
        _ = fi25
        _ = v26
    }
    if false {
        var v27 int32 = (-(v5))
        _ = v27
    }
    if (1 < 42) {
        for fi28 := 0; fi28 < 4; fi28++ {
            switch "\n\t" {
            case "a":
                var v29 uint16 = 2
                _ = v29
            case "\n\t":
                v30 := GenS3{F0: "foo bar", F1: (v4 + (-(v4))), F2: genFn0()}
                _ = v30
            }
            _ = fi28
        }
        fmt.Println(v3, v5, v5)
        v5 = 2147483647
    }
    switch "foo bar" {
    case "":
        var v31 uint = 2
        _ = v31
    case "\n\t":
        _ = len(v1)
    case "hello":
        var v32 bool = false
        _ = v32
    default:
        v2 = v2
    }
    var v33 string = "\n\t"
    var v34 float32 = v3
    if (true && false) {
        if (!false) {
            v35 := GenS4{F0: 108, F1: genFn0(), F2: 0}
            v5 = (v5 - (v5 % ((v5) | 1)))
            _ = len(v1)
            _ = v35
        } else {
            if true {
                v3 += 0.1
                if len(v6) > 0 {
                    v6[0] = v33
                }
            } else {
                v36 := GenS2{F0: (1 >= 42), F1: "a"}
                _ = v36
            }
            fmt.Println(v3, v4)
            if true {
                if false {
                    var v37 uint8 = v2
                    fmt.Println(v37, v3, v2)
                    _ = v37
                }
                v2 -= (v2 << ((0) & 63))
            }
        }
        v38 := GenS1{F0: genFn0(), F1: -9223372036854775808, F2: (v4 / (((v4 >> ((11) & 63))) | 1))}
        _ = v38
    } else {
        var v39 uint32 = 2
        _ = v39
    }
    if ((-996 > 32767) && false) {
        for fi40 := 0; fi40 < 20; fi40++ {
            v41 := GenS3{F0: "\n\t", F1: v4, F2: 255}
            for fi42 := 0; fi42 < 2; fi42++ {
                v34 = v34
                var v43 uint32 = 604
                if (v43 == genFn1(false, 367)) {
                    v44 := GenS5{F0: 42, F1: (v2 | 127), F2: 0}
                    var v45 int8 = 1
                    _ = v44
                    _ = v45
                } else {
                    var v46 int8 = 19
                    _ = len(v1)
                    var v47 string = ((((v33 + (v33 + v33)) + "") + "") + "")
                    v5 -= (v5 ^ v5)
                    _ = v46
                    _ = v47
                }
                _ = fi42
                _ = v43
            }
            _ = fi40
            _ = v41
        }
        switch ("a" + v33) {
        case "\n\t":
            var v48 string = "foo bar"
            _ = v48
        case "a":
            v49 := GenS1{F0: 32767, F1: -9223372036854775808, F2: int16(v5)}
            _ = v49
        }
    } else {
        var v50 float32 = v34
        var v51 float32 = v50
        _ = v50
        _ = v51
    }
    v4 = (v4 << ((603) & 63))
    v52 := v1[4294967295]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
