package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint64, p1 int64, p2 uint) uint32 {
    v6 := []int{221}
    if (p2 >= 18446744073709551615) {
        v7 := GenS1{F0: 32767, F1: -231}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 15; fi8++ {
            for fi9 := 0; fi9 < 9; fi9++ {
                v6 = append(v6, (fi9 * fi8))
                if (0.0 == 3.14159) {
                    v10 := GenS1{F0: -32768, F1: -1}
                    var v11 string = "\n\t"
                    _ = v10
                    _ = v11
                }
                v12 := GenS1{F0: -638, F1: -1}
                _ = fi9
                _ = v12
            }
            if len(v6) > 0 {
                v6[0] = 127
            }
            _ = fi8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2147483647
}
func genFn1(p0 uint16, p1 string, p2 int64) uint {
    for fi13 := 0; fi13 < 12; fi13++ {
        if false {
            v14 := []int{(fi13 % ((fi13) | 1)), 32768, int(p2)}
            fmt.Println(p1, p2)
            _ = v14
        } else {
            p2 = p2
        }
        _ = fi13
    }
    if false {
        v15 := []int{}
        if (4294967295 <= 731) {
            if (1 > 4294967295) {
                if (4294967295 < 4294967295) {
                    for fi16 := 0; fi16 < 15; fi16++ {
                        if (false && false) {
                            p2 += (p2 | (p2 / ((p2) | 1)))
                            var v17 bool = false
                            _ = v17
                        } else {
                            p1 = (p1 + p1)
                        }
                        v18 := GenS2{F0: (fi16 + fi16), F1: (p1 + ("" + p1))}
                        _ = fi16
                        _ = v18
                    }
                    p2 = p2
                }
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
            }
            p2 = (^(p2))
            v19 := GenS3{F0: 121, F1: p0}
            _ = v19
        } else {
            if len(v15) > 0 {
                v15[0] = 2147483647
            }
            fmt.Println(p2)
        }
        _ = v15
    }
    if true {
        switch p1 {
        case "a":
            var v20 float64 = 0.1
            _ = v20
        case "foo bar":
            if (-2.25 != 1e10) {
                switch -2147483648 {
                case 9223372036854775807:
                    var v21 int8 = 100
                    _ = v21
                case 128:
                    fmt.Println(p1)
                }
            }
        default:
            if true {
                p0 += 42
                p0 = (p0 + (p0 >> ((0) & 63)))
                v22 := []int{-2147483648, 62, 0, 32767}
                _ = v22
            } else {
                for fi23 := 0; fi23 < 8; fi23++ {
                    fmt.Println(p1, p2)
                    if true {
                        p1 = p1
                        p1 = "a"
                        switch (^(fi23)) {
                        case -9223372036854775808:
                            fmt.Println(p0, p2, p0)
                        default:
                            var v24 int = (fi23 % ((32767) | 1))
                            _ = v24
                        }
                    }
                    _ = fi23
                }
            }
        }
        if true {
            p0 = p0
            var v25 int16 = 42
            _ = v25
        } else {
            if (42 != 0) {
                p0 -= p0
                p0 += p0
            }
            v26 := GenS3{F0: 100, F1: 65535}
            _ = v26
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]uint32)
    var v3 uint64 = 4294967295
    v4 := make(map[int]string)
    var v5 int8 = 127
    var v27 int32 = (v1 ^ 0)
    v27 = v1
    v28 := GenS3{F0: (v5 / (((v5 | v5)) | 1)), F1: 32767}
    if true {
        var v29 bool = (65535 < 137)
        var v30 uint8 = 2
        _ = v29
        _ = v30
    }
    v5 = 127
    v1 = (v27 ^ 504)
    v27 = 530
    delete(v2, 237)
    for fi31 := 0; fi31 < 9; fi31++ {
        var v32 int16 = 42
        for fi33 := 0; fi33 < 10; fi33++ {
            switch "hello" {
            case "hello":
                v3 = v3
            default:
                v2[538] = genFn0(v3, -269, 1)
                if (((v5 >= v5) && true) || (float64((1e10 + -2.25)) >= 0.0)) {
                    var v34 int16 = v32
                    var v35 int32 = (v1 % (((v27 << ((288) & 63))) | 1))
                    _ = v34
                    _ = v35
                }
            }
            switch "foo bar" {
            case "a":
                v36 := []int{(fi31 - int(v3)), -9223372036854775808}
                _ = v36
            }
            fmt.Println(v27, fi31)
            switch 284 {
            case 127:
                if true {
                    v4[(fi31 & fi31)] = ("hello" + ("\n\t" + "a"))
                    var v37 int = fi33
                    v1 = (v27 ^ (v1 / ((v27) | 1)))
                    _ = v37
                }
            default:
                fmt.Println(v3)
            }
            _ = fi33
        }
        switch (-(fi31)) {
        case 0:
            v27 -= 0
        }
        _ = fi31
        _ = v32
    }
    switch ("hello" + "\n\t") {
    case "a":
        for fi38 := 0; fi38 < 13; fi38++ {
            var v39 string = "a"
            v39 = "\n\t"
            _ = fi38
            _ = v39
        }
    case "foo bar":
        fmt.Println(v5)
    case "hello":
        v1 += (-(v27))
    }
    switch "\n\t" {
    case "a":
        var v40 int = -1
        _ = v40
    }
    switch ("" + ("\n\t" + "\n\t")) {
    case "\n\t":
        v5 = -1
    default:
        if true {
            var v41 uint64 = (^(v3))
            if (!true) {
                fmt.Println(v1, v27)
                v41 -= (v41 % ((v3) | 1))
                _ = len(v4)
            }
            v42 := GenS2{F0: 2147483647, F1: "foo bar"}
            _ = v41
            _ = v42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
