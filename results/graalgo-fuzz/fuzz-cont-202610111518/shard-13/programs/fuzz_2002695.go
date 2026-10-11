package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int16; F2 int16 }
type GenS3 struct { F0 string; F1 int; F2 uint32 }
type GenS4 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS5 struct { F0 int16; F1 uint }

func main() {
    var v1 uint8 = 36
    var v2 uint = 1
    var v3 uint = 2
    var v4 bool = (("" + ("foo bar" + "foo bar")) == (("foo bar" + "foo bar") + "\n\t"))
    v5 := make(map[string]bool)
    var v6 uint8 = v1
    v7 := GenS1{F0: uint(v6)}
    for fi8 := 0; fi8 < 2; fi8++ {
        if v4 {
            fmt.Println(v6, fi8)
            if (!v4) {
                for fi9 := 0; fi9 < 4; fi9++ {
                    if false {
                        v10 := v5[("foo bar" + (("hello" + ("\n\t" + ((("" + ("" + ("hello" + ("a" + "a")))) + (("\n\t" + ("" + ((("hello" + "\n\t") + (((("hello" + "") + "a") + ("\n\t" + "\n\t")) + "foo bar")) + "a"))) + "hello")) + "foo bar"))) + "a"))]
                        _ = v10
                    } else {
                        delete(v5, "")
                        var v11 float64 = (((-2.25 + 0.0) + (0.1 * 1e10)) * float64((1.5 + -2.25)))
                        if (411 >= 645) {
                            var v12 bool = (v4 || (v6 >= v6))
                            fmt.Println(fi8, fi9, v6)
                            _ = v12
                        }
                        _ = v11
                    }
                    _ = fi9
                }
                v2 += (v3 >> ((v3) & 63))
            }
            v13 := GenS1{F0: (v3 | 42)}
            _ = v13
        } else {
            v14 := []int{60, 1}
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 int32 = 42
                v17 := GenS1{F0: (^(v2))}
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v14
        }
        v18 := GenS2{F0: v4, F1: 1000, F2: 1}
        var v19 int16 = -1
        _ = fi8
        _ = v18
        _ = v19
    }
    var v20 string = ("\n\t" + "")
    if true {
        v21 := GenS1{F0: v2}
        v22 := GenS3{F0: "\n\t", F1: 42, F2: 65535}
        for fi23 := 0; fi23 < 7; fi23++ {
            var v24 bool = (!v4)
            v25 := []int{4294967295, 128, fi23}
            if (!v4) {
                switch (fi23 - -107) {
                case 32767:
                    v6 += uint8(v3)
                case 0:
                    if (!(255 >= 1)) {
                        fmt.Println(fi23, v3, v1)
                        v26 := GenS3{F0: (v20 + v20), F1: (fi23 / ((fi23) | 1)), F2: 2147483647}
                        for fi27 := 0; fi27 < 10; fi27++ {
                            fmt.Println(v1, fi23)
                            fmt.Println(v6, v4)
                            v28 := GenS3{F0: ("hello" + ""), F1: fi27, F2: 2147483648}
                            v25 = append(v25, -839)
                            _ = fi27
                            _ = v28
                        }
                        _ = v26
                    } else {
                        v1 = (v6 | v6)
                        var v29 int16 = -1
                        _ = v29
                    }
                }
                var v30 int = fi23
                _ = v30
            } else {
                v1 += v6
                for fi31 := 0; fi31 < 15; fi31++ {
                    var v32 int8 = 73
                    var v33 float32 = 1.5
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                fmt.Println(fi23, v2)
            }
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = v21
        _ = v22
    }
    v6 -= (v6 | v6)
    var v34 bool = v4
    v34 = (v4 || (!v4))
    v35 := GenS4{F0: 2, F1: (^(v1)), F2: -128}
    v36 := GenS2{F0: v4, F1: 1000, F2: -927}
    v34 = v4
    if ((("hello" + (((v20 + "\n\t") + (("hello" + ("" + (v20 + v20))) + v20)) + "foo bar")) + (v20 + ((v20 + ("\n\t" + (((v20 + (v20 + v20)) + v20) + v20))) + v20))) == "foo bar") {
        v37 := []int{832, 468}
        for fi38 := 0; fi38 < 13; fi38++ {
            fmt.Println(v2, v4)
            if (!v4) {
                v39 := GenS2{F0: (-128 <= 63), F1: 0, F2: -446}
                var v40 uint64 = 512
                v41 := GenS1{F0: v2}
                _ = v39
                _ = v40
                _ = v41
            }
            _ = fi38
        }
        v1 = 74
        _ = v37
    } else {
        v42 := []int{9223372036854775807, 668, 1}
        _ = v42
    }
    v43 := GenS2{F0: v34, F1: -418, F2: -1}
    for fi44 := 0; fi44 < 19; fi44++ {
        v45 := []int{}
        switch "a" {
        case "hello":
            var v46 int32 = -2147483648
            _ = v46
        case "\n\t":
            fmt.Println(v34, v34)
        case "":
            for fi47 := 0; fi47 < 19; fi47++ {
                var v48 string = ("" + ("a" + "hello"))
                _ = fi47
                _ = v48
            }
        default:
            v49 := GenS5{F0: -89, F1: (v3 * (v2 >> (((v2 % ((2) | 1))) & 63)))}
            _ = v49
        }
        _ = fi44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
