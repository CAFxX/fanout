package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 uint32 }
func genFn0(p0 int) uint8 {
    if (true && false) {
        var v4 float64 = (0.0 * ((1.5 - (((-2.25 + 3.14159) + ((1.5 + float64(0.1)) + 3.14159)) + float64(3.14159))) / 0.1))
        _ = v4
    } else {
        v5 := GenS1{F0: 73, F1: -9223372036854775808}
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 uint = 18446744073709551615
            v8 := GenS2{F0: -1, F1: 202}
            v9 := []int{}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        v10 := GenS2{F0: 2147483647, F1: 246}
        _ = v10
        _ = v5
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 int32 = -2147483648
        switch 42 {
        case 1:
            var v13 uint16 = 65535
            _ = v13
        case 482:
            var v14 int8 = 100
            _ = v14
        case -2147483648:
            v12 = (-(v12))
        default:
            for fi15 := 0; fi15 < 16; fi15++ {
                if (!(!(1000 < 316))) {
                    v12 = v12
                    var v16 int16 = -32768
                    switch p0 {
                    case 32767:
                        var v17 int32 = v12
                        if (6 >= 58) {
                            v16 -= v16
                        }
                        _ = v17
                    case -9223372036854775808:
                        if ((fi11 | (^(fi11))) < fi15) {
                            v12 -= 0
                            var v18 float32 = -2.25
                            var v19 int16 = 952
                            p0 = fi15
                            _ = v18
                            _ = v19
                        }
                    default:
                        v20 := GenS3{F0: 0}
                        _ = v20
                    }
                    _ = v16
                }
                fmt.Println(fi15, fi11, fi11)
                _ = fi15
            }
        }
        var v21 int32 = (v12 & 0)
        _ = fi11
        _ = v12
        _ = v21
    }
    p0 = (p0 ^ (p0 * (p0 << ((65535) & 63))))
    for fi22 := 0; fi22 < 1; fi22++ {
        for fi23 := 0; fi23 < 17; fi23++ {
            var v24 uint8 = 52
            p0 = 32767
            switch ("hello" + ("a" + "")) {
            case "hello":
                if false {
                    v24 -= (v24 / (((v24 % ((v24) | 1))) | 1))
                    v25 := []int{42, fi22, (fi22 / ((fi23) | 1))}
                    _ = v25
                } else {
                    v26 := []int{32767, -692}
                    p0 -= (p0 & fi22)
                    _ = len(v26)
                    _ = v26
                }
            default:
                v27 := []int{(-(fi22))}
                _ = v27
            }
            _ = fi23
            _ = v24
        }
        var v28 int = (fi22 | 0)
        var v29 int32 = -1
        _ = fi22
        _ = v28
        _ = v29
    }
    var v30 int = (^(p0))
    _ = p0
    _ = v30
    return 9
}

func main() {
    v1 := []int{}
    var v2 string = "hello"
    var v3 uint16 = 0
    if false {
        v3 = (v3 * 65535)
    }
    if false {
        if len(v1) > 0 {
            v1[0] = -930
        }
        v31 := GenS2{F0: 0, F1: genFn0(-199)}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v31
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi32 := 0; fi32 < 7; fi32++ {
            if len(v1) > 0 {
                v1[0] = (fi32 - (fi32 >> ((0) & 63)))
            }
            _ = fi32
        }
        var v33 string = "a"
        _ = v33
    }
    var v34 uint64 = 21
    for fi35 := 0; fi35 < 4; fi35++ {
        for fi36 := 0; fi36 < 5; fi36++ {
            v37 := GenS1{F0: -2147483648, F1: 287}
            var v38 bool = (uint32(v34) > 0)
            var v39 uint8 = 170
            var v40 int16 = -32768
            v1 = append(v1, (fi35 ^ 4294967295))
            if (v38 && false) {
                switch (fi36 >> ((0) & 63)) {
                case 614:
                    fmt.Println(v39)
                case -366:
                    var v41 uint16 = 0
                    _ = v41
                case 127:
                    fmt.Println(v2, v2, v38)
                }
                if (1 < 0) {
                    fmt.Println(fi35)
                }
            }
            var v42 string = ("a" + v2)
            _ = fi36
            _ = v37
            _ = v38
            _ = v39
            _ = v40
            _ = v42
        }
        var v43 bool = false
        if true {
            fmt.Println(v2, v43)
            _ = len(v1)
            for fi44 := 0; fi44 < 6; fi44++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi44
            }
        } else {
            for fi45 := 0; fi45 < 17; fi45++ {
                switch ("hello" + "foo bar") {
                case "a":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "\n\t":
                    fmt.Println(v43, v34, v3)
                case "hello":
                    _ = len(v1)
                }
                _ = fi45
            }
            v3 = v3
            v3 += (v3 + (v3 << ((0) & 63)))
        }
        _ = fi35
        _ = v43
    }
    _ = len(v1)
    var v46 int16 = 1000
    v46 = 662
    v47 := GenS4{F0: 0, F1: 52}
    v34 = (^(v34))
    v48 := GenS1{F0: 0, F1: 9223372036854775807}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v46
    _ = v47
    _ = v48
}
