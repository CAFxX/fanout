package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 bool; F2 string }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int; F1 int64; F2 int32 }
type GenS5 struct { F0 int32; F1 uint8; F2 uint8 }
func genFn0() uint {
    var v5 int8 = -56
    v6 := GenS1{F0: 0}
    if (255 != 223) {
        var v7 int32 = -1
        for fi8 := 0; fi8 < 2; fi8++ {
            fmt.Println(v7)
            for fi9 := 0; fi9 < 20; fi9++ {
                switch "hello" {
                case "\n\t":
                    v10 := []int{fi9, (fi8 | fi9), (fi9 ^ -9223372036854775808)}
                    _ = v10
                }
                if (fi9 <= fi8) {
                    v7 += (v7 % (((v7 + 2147483647)) | 1))
                }
                switch "" {
                case "foo bar":
                    fmt.Println(fi8, v7)
                case "\n\t":
                    if ((0.1 + 1.5) != float32(1.5)) {
                        var v11 int16 = -32768
                        if true {
                            v12 := []int{9223372036854775807, (fi9 / ((2147483647) | 1))}
                            if len(v12) > 0 {
                                v12[0] = fi9
                            }
                            _ = v12
                        }
                        _ = v11
                    }
                default:
                    if (3.14159 != 0.1) {
                        var v13 int64 = -9223372036854775808
                        _ = v13
                    } else {
                        v5 = v5
                        var v14 float32 = ((3.14159 - 3.14159) - ((3.14159 * -2.25) * 0.0))
                        v7 = v7
                        _ = v14
                    }
                }
                _ = fi9
            }
            v7 = 0
            _ = fi8
        }
        _ = v7
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        if true {
            v16 := GenS2{F0: 2147483647, F1: true, F2: "\n\t"}
            var v17 int64 = 0
            _ = v16
            _ = v17
        }
        v18 := GenS3{F0: (v5 % (((v5 >> ((671) & 63))) | 1)), F1: "foo bar"}
        _ = fi15
        _ = v18
    }
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 uint, p1 string) uint64 {
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 int64 = -47
        p1 = p1
        var v21 bool = (-892 == 42)
        _ = fi19
        _ = v20
        _ = v21
    }
    if false {
        switch -742 {
        case 2147483647:
            var v22 uint8 = 217
            _ = v22
        case 438:
            v23 := []int{-689, 42, 42}
            _ = v23
        default:
            for fi24 := 0; fi24 < 10; fi24++ {
                if (true && (42 < 42)) {
                    v25 := []int{fi24, (-(fi24)), (fi24 / (((fi24 ^ (fi24 - 4294967295))) | 1)), fi24}
                    _ = v25
                }
                if (-938 >= -9223372036854775808) {
                    v26 := GenS2{F0: -2147483648, F1: (!(!false)), F2: "a"}
                    var v27 int16 = -1
                    _ = v26
                    _ = v27
                } else {
                    var v28 uint32 = 2147483647
                    if (!true) {
                        v29 := GenS3{F0: 79, F1: ""}
                        v28 = (v28 * (v28 / ((v28) | 1)))
                        fmt.Println(p1, p1)
                        _ = v29
                    }
                    v28 += v28
                    _ = v28
                }
                _ = fi24
            }
        }
        p1 = (p1 + (p1 + (p1 + "a")))
    }
    switch p1 {
    case "a":
        var v30 uint64 = 1
        _ = v30
    }
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    v1 := []uint32{42, 2147483647, 65535}
    var v2 float32 = float32(1.5)
    var v3 uint = 18446744073709551615
    var v4 int16 = 531
    v4 += (v4 + v4)
    if true {
        switch 2147483647 {
        case 0:
            fmt.Println(v4, v2)
        case -579:
            v31 := GenS4{F0: 127, F1: 549, F2: 912}
            _ = v31
        case 86:
            var v32 uint8 = 0
            _ = v32
        }
    } else {
        switch "\n\t" {
        case "hello":
            var v33 int64 = 1
            _ = v33
        case "":
            for fi34 := 0; fi34 < 7; fi34++ {
                v3 -= v3
                var v35 int = (fi34 - fi34)
                _ = fi34
                _ = v35
            }
        }
        if true {
            v3 = 515
        } else {
            switch -1 {
            case 102:
                switch 42 {
                case 4294967295:
                    fmt.Println(v2, v4, v2)
                }
            case 2147483647:
                v2 += (0.0 - v2)
            }
            fmt.Println(v2, v3)
            v36 := GenS5{F0: 1, F1: 128, F2: 70}
            _ = v36
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v37 float32 = float32(1.5)
    fmt.Println(v3, v2, v37)
    v38 := GenS3{F0: -45, F1: "hello"}
    v2 += v2
    v2 -= v37
    for fi39 := 0; fi39 < 4; fi39++ {
        v1 = append(v1, 2)
        v3 = genFn0()
        var v40 int32 = 0
        _ = fi39
        _ = v40
    }
    if (255 > 32767) {
        _ = len(v1)
        v37 -= (v37 / v37)
    }
    v37 = 0.0
    v1 = append(v1, 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
