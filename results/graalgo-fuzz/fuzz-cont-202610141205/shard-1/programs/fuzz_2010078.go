package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint16 }
func genFn0(p0 string, p1 int) uint8 {
    fmt.Println(p1, p0)
    var v5 int8 = 0
    var v6 int16 = 1
    switch "hello" {
    case "foo bar":
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := GenS1{F0: 65535, F1: (v5 % ((1) | 1)), F2: -128}
            var v9 int32 = int32(fi7)
            for fi10 := 0; fi10 < 18; fi10++ {
                v11 := []int{(^(fi10))}
                if (0 != 2) {
                    v9 = (v9 ^ (^(v9)))
                } else {
                    var v12 bool = (v5 != (v5 ^ v5))
                    v9 += v9
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    case "hello":
        v13 := []int{(p1 | p1), (p1 * (p1 - p1)), (p1 - p1), (p1 / (((p1 % ((p1) | 1))) | 1))}
        _ = v13
    case "a":
        switch "hello" {
        case "hello":
            for fi14 := 0; fi14 < 14; fi14++ {
                p1 -= int(v6)
                v15 := GenS2{F0: -355, F1: (-(v5))}
                if (!true) {
                    for fi16 := 0; fi16 < 3; fi16++ {
                        fmt.Println(p0, v5)
                        var v17 int = (p1 & -1)
                        _ = fi16
                        _ = v17
                    }
                }
                _ = fi14
                _ = v15
            }
        case "foo bar":
            fmt.Println(v6)
        case "":
            var v18 float32 = (1e10 + (1e10 * ((0.1 / (((0.1 - float32(0.1)) + 0.1) + (0.0 / -2.25))) * 3.14159)))
            _ = v18
        }
    default:
        var v19 int64 = 0
        _ = v19
    }
    var v20 int64 = -1
    _ = p0
    _ = p1
    _ = v20
    _ = v5
    _ = v6
    return 128
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{}
    v3 := make(map[uint32]string)
    var v4 bool = true
    v3[1] = "foo bar"
    fmt.Println(v4)
    for fi21 := 0; fi21 < 12; fi21++ {
        if (1 != 0) {
            fmt.Println(v4, fi21)
            v4 = v4
        } else {
            var v22 uint64 = 42
            _ = v22
        }
        v4 = v4
        var v23 float64 = (0.1 / (-2.25 - 1.5))
        _ = fi21
        _ = v23
    }
    delete(v3, 659)
    var v24 int8 = -1
    _ = len(v3)
    for fi25 := 0; fi25 < 15; fi25++ {
        v3[477] = ""
        _ = fi25
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        v24 = (v24 & v24)
        _ = fi26
    }
    if (9223372036854775807 <= -599) {
        if (1 <= 133) {
            for fi27 := 0; fi27 < 10; fi27++ {
                for fi28 := 0; fi28 < 3; fi28++ {
                    v29 := v3[250]
                    _ = fi28
                    _ = v29
                }
                _ = fi27
            }
            for fi30 := 0; fi30 < 11; fi30++ {
                fmt.Println(fi30, v4)
                var v31 int8 = v24
                _ = fi30
                _ = v31
            }
            v4 = (v4 && (!v4))
        } else {
            fmt.Println(v24, v24)
            delete(v1, 87)
        }
        for fi32 := 0; fi32 < 12; fi32++ {
            v33 := GenS1{F0: 42, F1: v24, F2: v24}
            _ = fi32
            _ = v33
        }
    }
    var v34 uint64 = 0
    var v35 int64 = 740
    for fi36 := 0; fi36 < 16; fi36++ {
        var v37 uint8 = 214
        v3[242] = ("hello" + "a")
        _ = fi36
        _ = v37
    }
    fmt.Println(v24)
    v38 := GenS3{F0: "\n\t"}
    v35 -= -9223372036854775808
    fmt.Println(v4)
    if len(v2) > 0 {
        v2[0] = 1
    }
    for fi39 := 0; fi39 < 15; fi39++ {
        if true {
            var v40 uint32 = 2147483648
            _ = len(v2)
            if v4 {
                v2 = append(v2, fi39)
                v41 := GenS3{F0: ("foo bar" + "\n\t")}
                _ = v41
            } else {
                for fi42 := 0; fi42 < 2; fi42++ {
                    if false {
                        switch (fi42 * 127) {
                        case 9223372036854775807:
                            v43 := GenS3{F0: "\n\t"}
                            _ = v43
                        case 842:
                            v2 = append(v2, -9223372036854775808)
                        }
                        switch "" {
                        case "":
                            v40 += v40
                        case "a":
                            v34 -= 42
                        default:
                            fmt.Println(fi39)
                        }
                    } else {
                        v44 := GenS4{F0: 255, F1: genFn0("a", fi39), F2: 2}
                        v35 = v35
                        if len(v2) > 0 {
                            v2[0] = fi39
                        }
                        _ = v44
                    }
                    v45 := GenS2{F0: 32767, F1: v24}
                    _ = fi42
                    _ = v45
                }
            }
            _ = v40
        } else {
            var v46 uint32 = 2147483648
            _ = v46
        }
        v47 := v1[1]
        _ = fi39
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v4
}
