package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int8; F1 bool; F2 int32 }
type GenS5 struct { F0 int16 }
func genFn0(p0 int64) uint {
    var v7 bool = true
    v8 := []int{128, 368}
    _ = p0
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 int64) int32 {
    if (1e10 > -2.25) {
        var v9 uint = 911
        _ = v9
    } else {
        p0 = (p0 | -9223372036854775808)
        p0 = p0
    }
    v10 := []int{-591, -1, -2147483648, 635}
    for fi11 := 0; fi11 < 20; fi11++ {
        if ("\n\t" >= "") {
            p0 = (p0 - p0)
        }
        _ = fi11
    }
    switch "" {
    case "a":
        fmt.Println(p0)
    case "hello":
        p0 += (p0 | p0)
    }
    if false {
        var v12 uint = 481
        _ = v12
    }
    _ = p0
    _ = v10
    return -2147483648
}

func main() {
    var v1 bool = false
    var v2 int16 = -1
    var v3 uint8 = 1
    var v4 int8 = -128
    v5 := make(map[int]string)
    var v6 int64 = 42
    for fi13 := 0; fi13 < 18; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            v15 := v5[fi13]
            var v16 int8 = v4
            _ = fi14
            _ = v15
            _ = v16
        }
        v17 := GenS1{F0: 488}
        _ = fi13
        _ = v17
    }
    if (v6 == v6) {
        v6 = -1
        var v18 uint32 = 859
        _ = v18
    } else {
        v6 += (v6 / (((-(v6))) | 1))
        if true {
            var v19 int8 = 127
            _ = v19
        } else {
            for fi20 := 0; fi20 < 9; fi20++ {
                for fi21 := 0; fi21 < 13; fi21++ {
                    switch "foo bar" {
                    case "":
                        if (v6 >= v6) {
                            switch int(v2) {
                            case 601:
                                v3 = v3
                                v22 := []int{127, fi21, (fi21 / (((fi20 + (fi21 - (fi21 | (fi20 - 32768))))) | 1)), (fi21 & (-(fi21)))}
                                v1 = false
                                _ = v22
                            default:
                                v23 := GenS2{F0: 18446744073709551615}
                                _ = v23
                            }
                            v1 = (2 < 1)
                            var v24 int = (fi20 + int(v3))
                            _ = v24
                        } else {
                            var v25 float64 = -2.25
                            var v26 string = "\n\t"
                            _ = v25
                            _ = v26
                        }
                    default:
                        fmt.Println(v4, fi20)
                    }
                    v3 = 128
                    v27 := GenS3{F0: 37, F1: genFn1(v6)}
                    _ = fi21
                    _ = v27
                }
                var v28 int8 = (v4 << ((genFn0((v6 & v6))) & 63))
                _ = fi20
                _ = v28
            }
            if (v3 == 127) {
                v29 := []int{}
                _ = v29
            }
            v30 := []int{295, 128, 1}
            _ = v30
        }
        fmt.Println(v3)
    }
    fmt.Println(v6, v3)
    v1 = (42 == 42)
    v31 := []int{342, 32767, -9223372036854775808}
    switch 4294967295 {
    case -960:
        v32 := GenS2{F0: 649}
        _ = v32
    case -1:
        v33 := GenS3{F0: 1, F1: genFn1(v6)}
        _ = v33
    default:
        v3 = (v3 * v3)
    }
    if (v1 || (((!true) && (3.14159 <= float64((1.5 / 3.14159)))) || (1.5 <= 3.14159))) {
        if v1 {
            var v34 int8 = -128
            v1 = (233 != 947)
            _ = v34
        }
    } else {
        v35 := GenS4{F0: (v4 + (^(v4))), F1: ((^(v3)) < (v3 - (^(v3)))), F2: genFn1((-(v6)))}
        v5[698] = ("" + ("hello" + ""))
        _ = v35
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        for fi37 := 0; fi37 < 14; fi37++ {
            if (!((float64(1e10) / -2.25) != 1e10)) {
                v38 := GenS5{F0: (v2 ^ (v2 / (((v2 ^ (v2 / ((-777) | 1)))) | 1)))}
                _ = v38
            } else {
                fmt.Println(v1, v2, v2)
            }
            _ = fi37
        }
        var v39 bool = (819 == 523)
        fmt.Println(v3)
        _ = fi36
        _ = v39
    }
    var v40 float32 = (1.5 * float32(0.0))
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        v2 -= v2
    case "":
        v41 := v5[671]
        _ = v41
    case "hello":
        if len(v31) > 0 {
            v31[0] = -1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
