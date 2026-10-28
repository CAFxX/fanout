package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint32; F2 bool }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 bool }
func genFn0(p0 uint16, p1 uint64, p2 int8) int64 {
    switch -828 {
    case 431:
        p1 = (p1 & 42)
    }
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 uint16 = p0
        p0 = v5
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 42}
    v7 := GenS1{F0: -1}
    var v8 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 uint16, p1 uint32) int64 {
    p1 = p1
    switch -631 {
    case 289:
        p0 += p0
    case 705:
        v9 := GenS1{F0: 42}
        _ = v9
    case 128:
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := []int{(-(fi10)), fi10, fi10}
            _ = fi10
            _ = v11
        }
    }
    v12 := GenS1{F0: -2147483648}
    v13 := []int{4294967295, 65, 42, 9223372036854775807}
    v13 = append(v13, 707)
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return -1
}

func main() {
    var v1 uint16 = 234
    var v2 int16 = 42
    v3 := []int{-1, 0, 128}
    var v14 float32 = float32(((0.1 + (((0.0 - ((1e10 * -2.25) * (0.0 / 1.5))) * ((1e10 * 1.5) / (float64(1.5) + (1.5 / 0.1)))) / 1e10)) * (float64(1.5) / ((3.14159 * 1e10) / 1.0))))
    v3 = append(v3, 32767)
    var v15 int8 = -8
    fmt.Println(v15)
    var v16 string = "hello"
    if false {
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 uint8 = 205
            _ = fi17
            _ = v18
        }
        v3 = append(v3, 42)
        v16 = v16
    } else {
        var v19 uint8 = 90
        _ = v19
    }
    if (((0 <= 2) || ((v16 + "hello") <= "hello")) && ((v14 + ((-2.25 + v14) / 0.1)) == float32((((0.1 + -2.25) * -2.25) * (((float64((0.1 - (v14 + (v14 + v14)))) + float64((v14 - v14))) * (-2.25 / 1.5)) * 0.0))))) {
        for fi20 := 0; fi20 < 17; fi20++ {
            v15 = (-(v15))
            v15 += 1
            switch -9223372036854775808 {
            case -2147483648:
                fmt.Println(v1, fi20)
            case 704:
                var v21 bool = false
                _ = v21
            case 32768:
                _ = len(v3)
            default:
                v22 := GenS2{F0: (fi20 | (fi20 / ((128) | 1))), F1: 42, F2: v1}
                _ = v22
            }
            _ = fi20
        }
    } else {
        var v23 string = ""
        if false {
            v24 := GenS3{F0: v1, F1: uint32(v1), F2: (true && false)}
            v16 = v16
            _ = v24
        } else {
            if true {
                var v25 uint64 = 4294967295
                if len(v3) > 0 {
                    v3[0] = 992
                }
                _ = v25
            }
        }
        _ = v23
    }
    if (847 == 789) {
        var v26 uint8 = 1
        v27 := GenS1{F0: 2147483647}
        fmt.Println(v15)
        _ = v26
        _ = v27
    }
    var v28 string = v16
    var v29 int8 = int8(v2)
    if (18446744073709551615 != 1) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi30 := 0; fi30 < 19; fi30++ {
            v31 := GenS2{F0: (fi30 / (((^(fi30))) | 1)), F1: 42, F2: (v1 % ((0) | 1))}
            if true {
                v15 = (v15 << ((4294967295) & 63))
                v32 := GenS1{F0: 2147483647}
                var v33 string = "hello"
                _ = v32
                _ = v33
            } else {
                _ = len(v3)
            }
            _ = fi30
            _ = v31
        }
        switch v16 {
        case "foo bar":
            _ = len(v3)
        case "a":
            v2 -= v2
        default:
            for fi34 := 0; fi34 < 7; fi34++ {
                fmt.Println(v14, v2, v28)
                v14 = v14
                if true {
                    for fi35 := 0; fi35 < 14; fi35++ {
                        v1 = v1
                        if len(v3) > 0 {
                            v3[0] = (fi35 / ((32768) | 1))
                        }
                        _ = fi35
                    }
                } else {
                    v36 := GenS4{F0: v1, F1: true}
                    switch ("hello" + v28) {
                    case "":
                        switch -9223372036854775808 {
                        case -2147483648:
                            _ = len(v3)
                        }
                    }
                    _ = v36
                }
                _ = fi34
            }
        }
    }
    if (!false) {
        v37 := GenS5{F0: 1, F1: 65535}
        if (true || true) {
            fmt.Println(v16, v1, v16)
            if false {
                fmt.Println(v1, v1)
            }
        } else {
            var v38 int8 = 1
            _ = v38
        }
        v3 = append(v3, int(v29))
        _ = v37
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi39 := 0; fi39 < 7; fi39++ {
            v40 := GenS1{F0: -2147483648}
            if true {
                v16 = v28
                if len(v3) > 0 {
                    v3[0] = 32768
                }
            }
            v41 := GenS6{F0: ((!(0 < 32767)) && true)}
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    fmt.Println(v1, v29, v28)
    for fi42 := 0; fi42 < 8; fi42++ {
        var v43 int16 = v2
        _ = fi42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
}
