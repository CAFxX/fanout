package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int32; F1 bool }
func genFn0(p0 string, p1 uint32, p2 string) int {
    p1 = (p1 & p1)
    for fi7 := 0; fi7 < 10; fi7++ {
        p1 = p1
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}
func genFn1(p0 int16, p1 string, p2 uint64) int16 {
    switch 946 {
    case 32767:
        var v8 bool = false
        _ = v8
    case 4294967295:
        p1 = p1
    case 1:
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 int8 = 42
            if true {
                fmt.Println(p2)
            }
            _ = fi9
            _ = v10
        }
    default:
        var v11 float64 = 0.1
        _ = v11
    }
    v12 := []int{1}
    if true {
        var v13 uint32 = 846
        p1 = p1
        if (!true) {
            p1 = p1
            if ("hello" >= "foo bar") {
                for fi14 := 0; fi14 < 16; fi14++ {
                    for fi15 := 0; fi15 < 6; fi15++ {
                        v16 := GenS1{F0: 186, F1: true, F2: 42}
                        _ = fi15
                        _ = v16
                    }
                    var v17 int8 = 127
                    _ = fi14
                    _ = v17
                }
                for fi18 := 0; fi18 < 9; fi18++ {
                    var v19 int64 = -9223372036854775808
                    _ = fi18
                    _ = v19
                }
                v20 := GenS1{F0: 42, F1: ((!(!(42 >= 1))) && (p1 > p1)), F2: 0}
                _ = v20
            }
            for fi21 := 0; fi21 < 1; fi21++ {
                p0 -= (p0 % (((p0 * p0)) | 1))
                _ = fi21
            }
        }
        _ = v13
    }
    v22 := GenS1{F0: uint8(p2), F1: false, F2: 42}
    if false {
        if (-535 == 9223372036854775807) {
            var v23 uint16 = 516
            if ((v23 % (((v23 - (v23 % (((v23 & v23)) | 1)))) | 1)) != v23) {
                var v24 int16 = p0
                switch p1 {
                case "\n\t":
                    p0 -= v24
                default:
                    p1 = "hello"
                }
                _ = v24
            } else {
                v23 -= (v23 & (v23 | (v23 + 32767)))
                fmt.Println(p0)
            }
            if ((false || (false || true)) && (!(true && true))) {
                v23 = (v23 >> ((138) & 63))
                fmt.Println(p0, p2, p0)
            }
            _ = v23
        } else {
            fmt.Println(p2, p1)
        }
        for fi25 := 0; fi25 < 13; fi25++ {
            fmt.Println(fi25)
            var v26 int64 = -562
            _ = fi25
            _ = v26
        }
    } else {
        for fi27 := 0; fi27 < 2; fi27++ {
            v28 := GenS1{F0: 128, F1: true, F2: 0}
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            if ((1.5 - -2.25) <= 3.14159) {
                var v29 float32 = 1e10
                _ = v29
            } else {
                if true {
                    var v30 float32 = 0.1
                    _ = v30
                } else {
                    fmt.Println(p1, fi27, p1)
                }
                var v31 int8 = 1
                _ = v31
            }
            _ = fi27
            _ = v28
        }
        for fi32 := 0; fi32 < 14; fi32++ {
            for fi33 := 0; fi33 < 10; fi33++ {
                var v34 uint32 = 359
                fmt.Println(p0, p0, p1)
                fmt.Println(v34, v34)
                var v35 int8 = -15
                _ = fi33
                _ = v34
                _ = v35
            }
            v36 := GenS2{F0: (p0 / ((p0) | 1)), F1: -128}
            v37 := GenS3{F0: 0, F1: (!false)}
            _ = fi32
            _ = v36
            _ = v37
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v22
    return p0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    var v4 int16 = -1
    var v5 uint16 = 42
    var v6 uint16 = (^(v5))
    var v38 float64 = 0.1
    v6 = (v5 * 21)
    if true {
        v4 += int16(v6)
        v3[2147483647] = -2147483648
        var v39 int8 = 100
        _ = v39
    }
    if (-1 >= 328) {
        fmt.Println(v38)
    }
    v38 += float64((((-2.25 / 3.14159) * 0.0) * 1.5))
    fmt.Println(v4, v4, v38)
    fmt.Println(v38, v5)
    v4 += v4
    v4 += v4
    var v40 int32 = -70
    var v41 string = ("hello" + ("foo bar" + ("a" + "foo bar")))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
