package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint8; F1 int8; F2 uint16 }
func genFn0() uint16 {
    switch (("hello" + "a") + (("foo bar" + ("foo bar" + ("" + "\n\t"))) + ("hello" + ("" + "hello")))) {
    case "a":
        var v5 int32 = -370
        _ = v5
    case "":
        if false {
            var v6 uint = 2
            v7 := GenS1{F0: "a", F1: 42, F2: 380}
            _ = v6
            _ = v7
        } else {
            var v8 string = ""
            fmt.Println(v8)
            switch 1 {
            case 128:
                var v9 int = -9223372036854775808
                _ = v9
            case 42:
                fmt.Println(v8)
            case 9223372036854775807:
                if (4294967295 != 9223372036854775807) {
                    v8 = v8
                    var v10 int8 = 100
                    v8 = v8
                    _ = v10
                } else {
                    switch -319 {
                    case 42:
                        fmt.Println(v8)
                    case 32767:
                        v8 = (v8 + v8)
                    }
                    switch -991 {
                    case -1:
                        for fi11 := 0; fi11 < 17; fi11++ {
                            var v12 uint = 120
                            fmt.Println(fi11, v12)
                            switch -261 {
                            case 32768:
                                fmt.Println(v8, v12, v12)
                            case -768:
                                fmt.Println(fi11)
                            }
                            _ = fi11
                            _ = v12
                        }
                    default:
                        switch "hello" {
                        case "\n\t":
                            var v13 int16 = 0
                            _ = v13
                        default:
                            var v14 uint32 = 6
                            _ = v14
                        }
                    }
                }
            default:
                v15 := []int{}
                _ = v15
            }
            _ = v8
        }
    case "foo bar":
        v16 := GenS2{F0: 111, F1: 4294967295, F2: -9223372036854775808}
        _ = v16
    default:
        var v17 int8 = -128
        _ = v17
    }
    var v18 uint32 = 2
    v18 += (v18 % ((v18) | 1))
    _ = v18
    return 42
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]uint32)
    var v4 uint = 255
    var v19 uint = v4
    for fi20 := 0; fi20 < 5; fi20++ {
        for fi21 := 0; fi21 < 12; fi21++ {
            if false {
                if true {
                    var v22 uint64 = 42
                    v23 := GenS2{F0: 139, F1: 515, F2: 9223372036854775807}
                    delete(v3, uint32(v22))
                    _ = v22
                    _ = v23
                } else {
                    switch (fi21 % ((fi21) | 1)) {
                    case 4294967295:
                        var v24 uint64 = 642
                        _ = v24
                    case -904:
                        v19 = (v4 - v19)
                    case 127:
                        v4 += (^(v4))
                        var v25 uint16 = (^(v1))
                        _ = v25
                    }
                    v4 = v19
                }
            } else {
                switch 32767 {
                case -486:
                    if false {
                        switch fi21 {
                        case -9223372036854775808:
                            fmt.Println(v19, v4, fi21)
                        default:
                            if false {
                                fmt.Println(v19, fi21)
                                v4 -= (v19 / ((255) | 1))
                                var v26 int32 = -2147483648
                                _ = v26
                            } else {
                                v27 := GenS2{F0: 90, F1: 0, F2: 42}
                                v1 = uint16(v4)
                                _ = v27
                            }
                        }
                        var v28 int16 = -32768
                        v29 := v3[4294967295]
                        _ = v28
                        _ = v29
                    }
                default:
                    var v30 int = -1
                    _ = v30
                }
            }
            _ = fi21
        }
        for fi31 := 0; fi31 < 8; fi31++ {
            v32 := []int{(fi20 * 9223372036854775807), fi31}
            _ = fi31
            _ = v32
        }
        v33 := GenS2{F0: 1, F1: 2147483647, F2: 1}
        _ = fi20
        _ = v33
    }
    switch "foo bar" {
    case "hello":
        v19 += v4
    case "\n\t":
        switch -271 {
        case 1:
            fmt.Println(v1, v1)
        case -2147483648:
            fmt.Println(v1)
        default:
            var v34 float32 = (((0.0 - (1.5 * 1e10)) * (0.0 - (((-2.25 * (1.5 - 1e10)) / 1.0) * 1.5))) - 3.14159)
            _ = v34
        }
    case "foo bar":
        switch ("" + "a") {
        case "hello":
            fmt.Println(v1)
        }
    }
    v35 := GenS3{F0: 1, F1: 42}
    switch "hello" {
    case "\n\t":
        for fi36 := 0; fi36 < 10; fi36++ {
            var v37 bool = (false && ((fi36 ^ fi36) < (fi36 / ((fi36) | 1))))
            _ = fi36
            _ = v37
        }
    default:
        if true {
            fmt.Println(v4)
        }
    }
    var v38 int = 430
    v39 := []int{}
    if len(v39) >= 2 {
        v39 = v39[1:len(v39)]
    }
    for fi40 := 0; fi40 < 1; fi40++ {
        if len(v39) >= 2 {
            v39 = v39[1:len(v39)]
        }
        var v41 bool = false
        _ = fi40
        _ = v41
    }
    for fi42 := 0; fi42 < 15; fi42++ {
        fmt.Println(v38)
        _ = fi42
    }
    if (v38 != v38) {
        if len(v39) >= 2 {
            v39 = v39[1:len(v39)]
        }
        v43 := GenS3{F0: int64(v4), F1: 328}
        for fi44 := 0; fi44 < 2; fi44++ {
            v45 := GenS4{F0: 0, F1: 0, F2: genFn0()}
            v46 := v3[0]
            v47 := GenS4{F0: 0, F1: -128, F2: v1}
            _ = fi44
            _ = v45
            _ = v46
            _ = v47
        }
        _ = v43
    } else {
        v39 = append(v39, 765)
        v38 += (v38 ^ 32767)
    }
    var v48 uint64 = 273
    v4 -= (v19 ^ v19)
    var v49 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v48
    _ = v49
}
