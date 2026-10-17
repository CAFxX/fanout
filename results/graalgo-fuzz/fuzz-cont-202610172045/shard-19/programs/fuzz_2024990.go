package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 uint8 }
func genFn0(p0 bool, p1 int8, p2 uint64) uint32 {
    for fi4 := 0; fi4 < 14; fi4++ {
        var v5 uint = 525
        p2 -= p2
        switch "hello" {
        case "\n\t":
            switch -259 {
            case 9223372036854775807:
                if p0 {
                    v6 := []int{-1, 32767, -861, 9223372036854775807}
                    if (!(1e10 == -2.25)) {
                        if (!false) {
                            if (!(p0 && true)) {
                                p1 = -1
                                if len(v6) > 0 {
                                    v6[0] = fi4
                                }
                                v5 = (^(v5))
                            }
                            p2 = p2
                        } else {
                            v7 := GenS1{F0: 2, F1: -32768}
                            v8 := GenS1{F0: 2, F1: -535}
                            _ = v7
                            _ = v8
                        }
                        v9 := GenS2{F0: 32767, F1: (^(p2))}
                        var v10 int = fi4
                        _ = v10
                        _ = v9
                    } else {
                        v11 := GenS3{F0: (v5 % (((v5 << ((65535) & 63))) | 1))}
                        fmt.Println(v5, p0, p1)
                        _ = v11
                    }
                    _ = v6
                }
            case 128:
                switch "\n\t" {
                case "a":
                    v5 -= (v5 | v5)
                }
            default:
                if p0 {
                    v12 := GenS4{F0: int32(p1)}
                    fmt.Println(v5)
                    v5 += 42
                    _ = v12
                }
            }
        case "":
            fmt.Println(p0)
        default:
            var v13 int32 = 0
            _ = v13
        }
        _ = fi4
        _ = v5
    }
    v14 := GenS1{F0: 4294967295, F1: 42}
    v15 := []int{}
    for fi16 := 0; fi16 < 19; fi16++ {
        if false {
            for fi17 := 0; fi17 < 7; fi17++ {
                v18 := GenS2{F0: int16(p1), F1: 42}
                _ = fi17
                _ = v18
            }
        }
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 2
}
func genFn1() int16 {
    for fi19 := 0; fi19 < 5; fi19++ {
        v20 := GenS4{F0: -834}
        _ = fi19
        _ = v20
    }
    v21 := []int{}
    v22 := GenS3{F0: 255}
    _ = v21
    _ = v22
    return 968
}

func main() {
    var v1 bool = true
    var v2 string = "foo bar"
    var v3 uint64 = 320
    var v23 int = 32768
    var v24 uint32 = 42
    for fi25 := 0; fi25 < 11; fi25++ {
        for fi26 := 0; fi26 < 5; fi26++ {
            var v27 uint16 = 65535
            if false {
                v28 := GenS3{F0: 18446744073709551615}
                v29 := GenS5{F0: 14, F1: true}
                var v30 uint = 65535
                if true {
                    v31 := []int{v23, v23, fi25, (fi26 * (-(fi25)))}
                    var v32 int8 = -81
                    _ = v31
                    _ = v32
                } else {
                    var v33 float32 = 3.14159
                    if v1 {
                        fmt.Println(v27, v33)
                    }
                    _ = v33
                }
                _ = v28
                _ = v29
                _ = v30
            }
            _ = fi26
            _ = v27
        }
        if false {
            v34 := []int{}
            _ = v34
        }
        _ = fi25
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        switch (v2 + (v2 + v2)) {
        case "\n\t":
            v1 = false
        case "a":
            fmt.Println(v3, v2, v23)
        case "foo bar":
            var v36 int = (v23 - 1)
            _ = v36
        default:
            v24 += uint32(v3)
        }
        _ = fi35
    }
    fmt.Println(v1, v2, v3)
    var v37 int16 = 0
    for fi38 := 0; fi38 < 7; fi38++ {
        var v39 uint8 = 2
        _ = fi38
        _ = v39
    }
    v40 := []int{519, (v23 ^ 9223372036854775807)}
    for fi41 := 0; fi41 < 13; fi41++ {
        for fi42 := 0; fi42 < 1; fi42++ {
            v2 = (("foo bar" + "") + v2)
            var v43 float64 = 0.0
            _ = fi42
            _ = v43
        }
        _ = fi41
    }
    v3 += v3
    v44 := GenS6{F0: 84}
    v45 := GenS5{F0: 50, F1: true}
    fmt.Println(v24, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v37
    _ = v40
    _ = v44
    _ = v45
}
