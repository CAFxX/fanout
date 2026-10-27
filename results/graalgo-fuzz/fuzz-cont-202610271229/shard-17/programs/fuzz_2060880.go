package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint }
type GenS2 struct { F0 int64; F1 uint; F2 bool }
func genFn0(p0 uint, p1 uint32) int32 {
    v4 := GenS1{F0: 18446744073709551615, F1: 42, F2: (p0 >> ((p0) & 63))}
    switch "a" {
    case "hello":
        p0 = p0
    }
    v5 := []int{-2147483648, 2147483647}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}

func main() {
    v1 := []int{}
    var v2 uint32 = 1
    var v3 uint8 = 255
    if false {
        if (!true) {
            if (!(!((true && ((-226 <= 0) && false)) && ((42 == -32768) || false)))) {
                fmt.Println(v2, v3)
                v3 = 2
                for fi6 := 0; fi6 < 16; fi6++ {
                    if false {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        for fi7 := 0; fi7 < 13; fi7++ {
                            var v8 uint = 0
                            v2 = v2
                            _ = fi7
                            _ = v8
                        }
                    }
                    v1 = append(v1, (fi6 | 1))
                    v2 -= (v2 ^ v2)
                    _ = fi6
                }
            } else {
                if ((v3 | 42) < 42) {
                    if true {
                        if (!(907 >= 65535)) {
                            v2 -= (v2 + (v2 - 0))
                        } else {
                            switch 4294967295 {
                            case 453:
                                v2 -= (v2 << ((65535) & 63))
                            case 105:
                                v2 += (v2 * v2)
                            case 65:
                                v3 += (v3 >> ((961) & 63))
                            }
                            v9 := GenS1{F0: 42, F1: -1, F2: 255}
                            fmt.Println(v2)
                            _ = v9
                        }
                        v10 := GenS1{F0: 779, F1: 930, F2: 1}
                        if false {
                            for fi11 := 0; fi11 < 10; fi11++ {
                                v3 -= (v3 % (((v3 % ((v3) | 1))) | 1))
                                var v12 uint32 = 2
                                _ = len(v1)
                                _ = fi11
                                _ = v12
                            }
                            var v13 uint16 = 1
                            _ = v13
                        }
                        _ = v10
                    }
                    var v14 int8 = -128
                    var v15 float32 = float32((1.5 * 3.14159))
                    _ = v14
                    _ = v15
                } else {
                    v16 := GenS1{F0: 335, F1: -947, F2: 4294967295}
                    fmt.Println(v2)
                    _ = v16
                }
                if len(v1) > 0 {
                    v1[0] = 577
                }
                v1 = append(v1, -290)
            }
        }
        if true {
            var v17 uint = 2
            var v18 uint8 = 1
            _ = v17
            _ = v18
        }
        var v19 uint8 = uint8(v2)
        _ = v19
    }
    if (84 != 1) {
        v2 += v2
        var v20 uint32 = (v2 & (v2 * 997))
        _ = v20
    } else {
        var v21 uint64 = 42
        v1 = append(v1, -9223372036854775808)
        _ = v21
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v22 uint64 = 4294967295
    switch "foo bar" {
    case "foo bar":
        v23 := GenS2{F0: 42, F1: 18, F2: false}
        _ = v23
    case "hello":
        for fi24 := 0; fi24 < 1; fi24++ {
            if true {
                var v25 uint32 = v2
                v1 = append(v1, fi24)
                _ = v25
            } else {
                var v26 int8 = int8(v3)
                _ = v26
            }
            for fi27 := 0; fi27 < 9; fi27++ {
                v1 = append(v1, fi24)
                if len(v1) > 0 {
                    v1[0] = fi27
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi27
            }
            _ = fi24
        }
    default:
        fmt.Println(v22, v22)
    }
    var v28 string = ("hello" + "")
    var v29 float64 = 0.1
    v2 -= v2
    if (!false) {
        fmt.Println(v2)
    }
    v30 := GenS2{F0: 298, F1: 144, F2: ((^(v22)) <= 2)}
    for fi31 := 0; fi31 < 4; fi31++ {
        v3 -= v3
        for fi32 := 0; fi32 < 17; fi32++ {
            fmt.Println(v29, v2)
            _ = fi32
        }
        _ = fi31
    }
    v33 := GenS1{F0: v22, F1: genFn0(313, 4294967295), F2: 2}
    for fi34 := 0; fi34 < 4; fi34++ {
        var v35 float32 = ((0.1 + 0.1) + (1e10 * -2.25))
        for fi36 := 0; fi36 < 20; fi36++ {
            var v37 int16 = -562
            var v38 float64 = (v29 * 1.5)
            var v39 int32 = genFn0(268, (v2 << ((65535) & 63)))
            _ = fi36
            _ = v37
            _ = v38
            _ = v39
        }
        _ = fi34
        _ = v35
    }
    var v40 int32 = -614
    var v41 uint64 = 733
    var v42 uint16 = 255
    var v43 string = ""
    fmt.Println(v29, v41, v43)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
