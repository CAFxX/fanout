package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int64 }
type GenS2 struct { F0 int32; F1 int16; F2 string }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint8; F1 int64; F2 int16 }
func genFn0(p0 int, p1 int16, p2 uint8) uint8 {
    p1 -= p1
    p1 = (p1 / ((p1) | 1))
    for fi5 := 0; fi5 < 8; fi5++ {
        p0 -= fi5
        _ = fi5
    }
    var v6 uint8 = (p2 - (p2 & (p2 + p2)))
    var v7 float64 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (^(p2))
}
func genFn1(p0 bool, p1 uint32) int32 {
    p0 = p0
    if (!p0) {
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS1{F0: 42, F1: 423, F2: 1}
            var v10 int64 = -1
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 int8 = 54
        _ = v11
    }
    p1 = (p1 % ((p1) | 1))
    if true {
        v12 := GenS2{F0: -2147483648, F1: 615, F2: "\n\t"}
        for fi13 := 0; fi13 < 20; fi13++ {
            p0 = (3.14159 < (0.1 - (1.5 * 1.5)))
            _ = fi13
        }
        for fi14 := 0; fi14 < 4; fi14++ {
            v15 := []int{}
            if p0 {
                switch (fi14 ^ (fi14 & fi14)) {
                case -9223372036854775808:
                    v16 := GenS2{F0: -1, F1: -32768, F2: (("foo bar" + "\n\t") + ("a" + ("hello" + ((("hello" + "hello") + "") + "hello"))))}
                    _ = v16
                case 128:
                    switch "\n\t" {
                    case "a":
                        v17 := GenS2{F0: -300, F1: 32767, F2: "\n\t"}
                        _ = v17
                    case "hello":
                        v18 := GenS2{F0: 0, F1: 1000, F2: ""}
                        _ = v18
                    default:
                        var v19 uint16 = 255
                        _ = v19
                    }
                }
                for fi20 := 0; fi20 < 11; fi20++ {
                    var v21 float32 = -2.25
                    _ = fi20
                    _ = v21
                }
                fmt.Println(fi14)
            } else {
                p1 += (p1 + 668)
                switch "foo bar" {
                case "":
                    var v22 uint64 = 0
                    _ = v22
                case "a":
                    var v23 int8 = -128
                    _ = v23
                }
            }
            if (3.14159 < ((3.14159 - 0.1) - 3.14159)) {
                var v24 int = fi14
                var v25 uint32 = (p1 - (p1 - 1))
                _ = v24
                _ = v25
            } else {
                if (0 != 18446744073709551615) {
                    _ = len(v15)
                    if p0 {
                        _ = len(v15)
                        p1 = p1
                    } else {
                        if (false && false) {
                            var v26 int = fi14
                            v27 := GenS2{F0: 0, F1: 1, F2: "a"}
                            var v28 int64 = -1
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                    }
                } else {
                    fmt.Println(fi14, p0, p0)
                    if (0.1 >= (3.14159 - ((1.5 + 0.0) / 1.5))) {
                        p1 += p1
                    } else {
                        p0 = (!(((false && ((p0 || (p0 && ((100 > -1) || p0))) || false)) && (p0 || p0)) || p0))
                    }
                }
                v29 := GenS2{F0: 42, F1: -230, F2: "a"}
                _ = v29
            }
            _ = fi14
            _ = v15
        }
        _ = v12
    } else {
        v30 := []int{}
        p1 = p1
        _ = v30
    }
    p0 = (3.14159 > 1.5)
    _ = p0
    _ = p1
    return -705
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{"a"}
    var v3 int64 = -1
    v4 := []string{"\n\t", "hello", "", "\n\t"}
    v3 += -9223372036854775808
    var v31 bool = (2 >= 0)
    v31 = false
    if len(v2) > 0 {
        v2[0] = "a"
    }
    if (v31 && v31) {
        var v32 int32 = -1
        var v33 float32 = ((float32((float64(0.1) + float64(-2.25))) * ((float32(((1e10 * (float64((-2.25 * float32((float64(((3.14159 / 3.14159) + 1.5)) - 1e10)))) - 1.5)) - (((-2.25 / 1e10) * 0.1) + 1e10))) - -2.25) - float32(-2.25))) / (((3.14159 / 1.0) - -2.25) / (0.0 + 3.14159)))
        v34 := GenS3{F0: 127, F1: int32(v3)}
        _ = v32
        _ = v33
        _ = v34
    } else {
        v35 := GenS4{F0: genFn0(343, -32768, 83), F1: v3, F2: -32768}
        _ = v35
    }
    if len(v2) > 0 {
        v2[0] = ""
    }
    fmt.Println(v3, v3)
    if true {
        fmt.Println(v3)
        if true {
            var v36 string = "\n\t"
            _ = v36
        } else {
            for fi37 := 0; fi37 < 19; fi37++ {
                var v38 float64 = 0.1
                _ = fi37
                _ = v38
            }
        }
        fmt.Println(v31, v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}
