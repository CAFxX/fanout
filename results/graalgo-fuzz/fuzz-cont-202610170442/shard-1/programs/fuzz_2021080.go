package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint32 }
func genFn0(p0 int64, p1 string, p2 uint) uint {
    var v6 uint = p2
    var v7 float64 = 0.1
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := []int{}
            _ = fi9
            _ = v10
        }
        p1 = ("hello" + (("a" + ("hello" + p1)) + "foo bar"))
        var v11 int8 = 127
        _ = fi8
        _ = v11
    }
    var v12 uint8 = 0
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 int32 = -2147483648
        v15 := GenS1{F0: -483, F1: 42, F2: 2}
        v7 += 1.5
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v6
    _ = v7
    return uint(v12)
}
func genFn1(p0 int64, p1 string, p2 uint32) uint16 {
    p1 = ("a" + p1)
    var v16 uint64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 65535
}

func main() {
    var v1 uint8 = 0
    v2 := []int{2147483647, 42}
    v3 := make(map[string]string)
    v4 := []int{-615, 775}
    var v5 uint64 = 957
    v5 = (v5 % ((200) | 1))
    v1 = v1
    switch -9223372036854775808 {
    case -247:
        v17 := GenS1{F0: -850, F1: 2, F2: 2}
        _ = v17
    }
    if (((102 <= v1) && (false && true)) || true) {
        for fi18 := 0; fi18 < 18; fi18++ {
            var v19 uint = 42
            _ = fi18
            _ = v19
        }
        v20 := GenS1{F0: 1000, F1: genFn1(-9223372036854775808, (("" + (((("a" + ("" + ("foo bar" + (("" + (("foo bar" + "") + "\n\t")) + ("foo bar" + "a"))))) + ("\n\t" + "")) + "") + "a")) + "foo bar"), 399), F2: genFn1(-9223372036854775808, "\n\t", 2147483648)}
        _ = v20
    }
    if (356 >= genFn0(42, ("" + ("foo bar" + "foo bar")), 65535)) {
        v21 := GenS2{F0: 1}
        _ = v21
    } else {
        var v22 uint32 = 2147483648
        _ = v22
    }
    v23 := v3["a"]
    for fi24 := 0; fi24 < 6; fi24++ {
        switch v23 {
        case "":
            v4 = append(v4, (-(fi24)))
        case "a":
            var v25 float64 = 3.14159
            _ = v25
        }
        _ = fi24
    }
    var v26 uint8 = v1
    v5 += (v5 + v5)
    if (true && (1000 != -32768)) {
        if (false && true) {
            for fi27 := 0; fi27 < 12; fi27++ {
                switch v23 {
                case "hello":
                    var v28 string = v23
                    _ = v28
                case "a":
                    v5 -= v5
                case "foo bar":
                    var v29 int8 = 0
                    _ = v29
                default:
                    for fi30 := 0; fi30 < 4; fi30++ {
                        var v31 uint32 = 2147483647
                        v1 -= 0
                        _ = fi30
                        _ = v31
                    }
                }
                _ = fi27
            }
        }
        _ = len(v2)
    }
    if ((-2.25 * (0.1 * 0.1)) != 0.0) {
        for fi32 := 0; fi32 < 11; fi32++ {
            for fi33 := 0; fi33 < 16; fi33++ {
                v26 -= 42
                if false {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(fi33, v26, fi32)
                } else {
                    var v34 int16 = 32767
                    _ = v34
                }
                switch 32768 {
                case -726:
                    v35 := GenS2{F0: 255}
                    _ = v35
                default:
                    v36 := GenS3{F0: ((1.5 / 3.14159) >= (-2.25 / 3.14159))}
                    _ = v36
                }
                _ = fi33
            }
            var v37 int = fi32
            v26 = v1
            _ = fi32
            _ = v37
        }
        v1 += uint8(v5)
        delete(v3, v23)
    }
    fmt.Println(v1, v1)
    if true {
        var v38 float64 = -2.25
        if true {
            var v39 int16 = -1
            _ = v39
        } else {
            v40 := v3[v23]
            var v41 uint16 = 0
            _ = v40
            _ = v41
        }
        _ = v38
    }
    var v42 uint8 = (v26 ^ v1)
    switch -9223372036854775808 {
    case 42:
        var v43 bool = false
        _ = v43
    case 127:
        var v44 int = -705
        _ = v44
    }
    v45 := GenS4{F0: 955, F1: 364, F2: 1}
    var v46 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v42
    _ = v45
    _ = v46
    _ = v5
}
