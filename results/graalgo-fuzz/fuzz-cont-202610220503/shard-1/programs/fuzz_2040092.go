package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 string }
type GenS2 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32 }
func genFn0() string {
    var v7 uint = 255
    v8 := []int{-9223372036854775808, 2147483647, -2147483648, -6}
    var v9 float32 = float32((3.14159 / 1e10))
    var v10 bool = true
    if v10 {
        v11 := GenS1{F0: "a", F1: 18446744073709551615, F2: (("hello" + ("" + ("foo bar" + ((("" + "\n\t") + (("\n\t" + "a") + "a")) + "")))) + "a")}
        var v12 int = 828
        _ = v11
        _ = v12
    }
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return ""
}
func genFn1(p0 uint8) uint16 {
    for fi13 := 0; fi13 < 8; fi13++ {
        p0 = 42
        v14 := []int{fi13, (fi13 << ((603) & 63)), (fi13 << ((18446744073709551615) & 63)), 374}
        _ = fi13
        _ = v14
    }
    var v15 float32 = 0.1
    var v16 bool = ("a" >= "")
    var v17 float64 = 0.0
    p0 += (p0 - (^(p0)))
    _ = p0
    _ = v15
    _ = v16
    _ = v17
    return 65535
}

func main() {
    v1 := []string{"a", "\n\t", "a"}
    var v2 uint16 = 255
    var v3 bool = false
    var v4 bool = true
    var v5 float32 = ((0.1 - (3.14159 - (((0.0 + ((1e10 * ((3.14159 / (3.14159 + 1.5)) * 0.1)) + 1.5)) - float32(1e10)) / 0.1))) + ((((-2.25 - -2.25) / 0.1) + (1.5 + 1.5)) / 3.14159))
    var v6 uint32 = 750
    var v18 bool = (0 < 246)
    v6 = 2147483647
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 float64 = -2.25
        if v4 {
            var v21 float32 = v5
            v18 = (998 < -2147483648)
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    var v22 uint8 = 127
    v22 = v22
    var v23 float32 = -2.25
    v5 = 1.5
    var v24 float32 = (v23 - float32(((float64(v5) / (-2.25 - ((float64(0.1) + 0.0) + (1.5 * 0.0)))) * -2.25)))
    v6 = v6
    if (v6 == 26) {
        _ = len(v1)
    } else {
        var v25 float32 = v24
        _ = v25
    }
    var v26 float32 = ((v24 - (float32(0.1) * v23)) + 1e10)
    v27 := GenS2{F0: 2147483648, F1: 0, F2: (^(v2))}
    switch -2147483648 {
    case -241:
        v26 = v5
    case 399:
        if v4 {
            fmt.Println(v24, v3, v4)
        } else {
            var v28 int = 32768
            switch "hello" {
            case "foo bar":
                if v3 {
                    switch "\n\t" {
                    case "":
                        var v29 string = ("foo bar" + "\n\t")
                        _ = v29
                    case "\n\t":
                        v26 = (v23 * 0.1)
                    case "a":
                        for fi30 := 0; fi30 < 7; fi30++ {
                            _ = len(v1)
                            v23 = v24
                            _ = fi30
                        }
                    }
                    if len(v1) > 0 {
                        v1[0] = ("foo bar" + "\n\t")
                    }
                    fmt.Println(v5)
                } else {
                    fmt.Println(v5, v24, v3)
                }
            case "\n\t":
                v31 := GenS3{F0: 0, F1: 4294967295, F2: 32767}
                _ = v31
            case "a":
                v32 := GenS2{F0: 786, F1: 1, F2: (v2 / (((v2 | (^(v2)))) | 1))}
                _ = v32
            default:
                v33 := GenS4{F0: 0}
                _ = v33
            }
            for fi34 := 0; fi34 < 15; fi34++ {
                if (-1 != 2147483647) {
                    v35 := GenS5{F0: 854}
                    if (438 <= 2) {
                        v28 -= (v28 * v28)
                        switch 42 {
                        case 2147483647:
                            _ = len(v1)
                        }
                        v1 = append(v1, "hello")
                    }
                    v6 = 842
                    _ = v35
                }
                v36 := GenS5{F0: (v6 / ((v6) | 1))}
                _ = fi34
                _ = v36
            }
            _ = v28
        }
    case 0:
        if (!false) {
            var v37 bool = (-7 <= -1)
            var v38 uint16 = 255
            _ = v37
            _ = v38
        }
    default:
        for fi39 := 0; fi39 < 4; fi39++ {
            var v40 uint8 = v22
            _ = fi39
            _ = v40
        }
    }
    if v3 {
        _ = len(v1)
    }
    v23 = 0.0
    v41 := GenS5{F0: (^(v6))}
    var v42 uint16 = (^(v2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
