package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint32 }
type GenS2 struct { F0 string; F1 uint16 }
func genFn0(p0 int) int {
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{(fi7 * p0), (fi7 << ((728) & 63)), -9223372036854775808, 869}
        p0 = fi7
        _ = fi7
        _ = v8
    }
    p0 += (p0 & p0)
    p0 += (p0 | (p0 / (((-(p0))) | 1)))
    var v9 uint16 = 0
    if (false && false) {
        var v10 int = 32768
        var v11 int16 = 1
        var v12 int = (v10 & -956)
        _ = v10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v9
    return (p0 * -1)
}
func genFn1(p0 uint16, p1 int32) int8 {
    v13 := GenS1{F0: true, F1: 0, F2: 1}
    p0 = 2
    p0 = p0
    _ = p0
    _ = p1
    _ = v13
    return 77
}

func main() {
    var v1 int16 = -32768
    var v2 uint16 = 264
    var v3 uint8 = 127
    v4 := []uint32{447, 4294967295}
    v5 := make(map[int]uint32)
    var v6 uint = 456
    for fi14 := 0; fi14 < 10; fi14++ {
        v1 += v1
        v2 = (v2 * 65535)
        _ = fi14
    }
    var v15 uint64 = 4294967295
    fmt.Println(v2)
    if len(v4) > 0 {
        v4[0] = 920
    }
    var v16 float32 = (1.5 + -2.25)
    var v17 uint8 = v3
    v6 = 1
    v6 = (v6 * v6)
    if true {
        v18 := GenS1{F0: (v16 > v16), F1: v1, F2: 4294967295}
        _ = v18
    } else {
        var v19 float32 = 1.5
        _ = v19
    }
    switch -237 {
    case 877:
        if ((v2 & 32767) < (v2 << ((797) & 63))) {
            switch "\n\t" {
            case "foo bar":
                v3 += v17
            case "a":
                switch ("a" + (("hello" + "") + ((("a" + ("\n\t" + "foo bar")) + "\n\t") + ""))) {
                case "hello":
                    var v20 uint32 = uint32(v15)
                    _ = v20
                case "":
                    for fi21 := 0; fi21 < 4; fi21++ {
                        v16 = float32(0.1)
                        _ = fi21
                    }
                }
            case "\n\t":
                for fi22 := 0; fi22 < 4; fi22++ {
                    v23 := GenS1{F0: (42 > -766), F1: v1, F2: 65535}
                    _ = fi22
                    _ = v23
                }
            default:
                if ((v6 * v6) >= (v6 % ((uint(v15)) | 1))) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            }
        } else {
            v24 := GenS1{F0: true, F1: -1, F2: 2}
            fmt.Println(v1, v16, v16)
            _ = v24
        }
    }
    var v25 int32 = -2147483648
    if ((v25 * v25) > int32(v2)) {
        fmt.Println(v2, v15, v3)
    } else {
        v26 := GenS1{F0: false, F1: 42, F2: 969}
        var v27 uint64 = v15
        _ = v26
        _ = v27
    }
    if ((v16 / (v16 - v16)) > (1.5 / (v16 + (1e10 - (((1e10 * -2.25) + float32(0.1)) + v16))))) {
        fmt.Println(v1, v6)
    } else {
        v4 = append(v4, 575)
        var v28 float64 = -2.25
        _ = v28
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 bool = false
        if v30 {
            if false {
                var v31 int16 = (v1 >> (((v6 ^ 255)) & 63))
                _ = v31
            }
        }
        _ = fi29
        _ = v30
    }
    fmt.Println(v16, v16)
    v25 += v25
    if (v16 != v16) {
        for fi32 := 0; fi32 < 12; fi32++ {
            if (false || ("foo bar" != ("\n\t" + ((("foo bar" + ("\n\t" + ("a" + "foo bar"))) + "\n\t") + "a")))) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                v3 -= (v3 ^ 19)
                var v33 uint32 = 2
                v17 += (^(v3))
                _ = v33
            }
            var v34 uint32 = 85
            if (v6 != 575) {
                v35 := GenS1{F0: false, F1: v1, F2: (v34 + uint32(v15))}
                switch "\n\t" {
                case "hello":
                    fmt.Println(v34)
                case "":
                    v16 = v16
                case "foo bar":
                    v2 = (^(v2))
                }
                _ = v35
            } else {
                if (v1 > 603) {
                    var v36 float64 = ((1.5 * float64((v16 * float32(1.5)))) * (((1.5 + 0.0) * 1e10) / 3.14159))
                    var v37 string = "\n\t"
                    _ = v36
                    _ = v37
                }
            }
            _ = fi32
            _ = v34
        }
        for fi38 := 0; fi38 < 7; fi38++ {
            v2 += v2
            v39 := GenS2{F0: "foo bar", F1: (v2 % ((v2) | 1))}
            _ = fi38
            _ = v39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
