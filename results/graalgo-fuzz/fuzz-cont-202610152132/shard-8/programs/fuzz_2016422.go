package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int64; F1 int64 }
func genFn0(p0 int, p1 int64) int32 {
    v4 := GenS1{F0: 42, F1: 2147483647, F2: -537}
    p0 = 1
    p1 = (p1 * p1)
    _ = p0
    _ = p1
    _ = v4
    return 181
}
func genFn1(p0 uint64, p1 int8, p2 int32) uint64 {
    v5 := []int{-9223372036854775808}
    p2 = (p2 >> ((364) & 63))
    if false {
        for fi6 := 0; fi6 < 7; fi6++ {
            switch (fi6 & (fi6 | fi6)) {
            case 32767:
                var v7 int = -2147483648
                _ = v7
            }
            var v8 int8 = p1
            _ = fi6
            _ = v8
        }
        switch "" {
        case "":
            var v9 int16 = -32768
            _ = v9
        case "foo bar":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case "hello":
            for fi10 := 0; fi10 < 19; fi10++ {
                p0 = p0
                _ = fi10
            }
        default:
            if (1 < 378) {
                for fi11 := 0; fi11 < 11; fi11++ {
                    var v12 float64 = 1e10
                    if (!(255 >= 1)) {
                        v13 := GenS2{F0: -128}
                        v5 = append(v5, 855)
                        p0 = p0
                        _ = v13
                    }
                    if (!((-1 >= -1) && (genFn0(fi11, 256) == genFn0(fi11, -111)))) {
                        p0 = p0
                    }
                    _ = fi11
                    _ = v12
                }
            }
        }
    } else {
        if (462 >= -1) {
            var v14 bool = (65535 != 4294967295)
            _ = v14
        } else {
            if (true && ((58 < 255) || false)) {
                switch ("a" + "foo bar") {
                case "foo bar":
                    var v15 string = "a"
                    _ = v15
                default:
                    if true {
                        p1 -= p1
                    }
                }
                if false {
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v5 = append(v5, 9223372036854775807)
                        _ = fi16
                    }
                    fmt.Println(p2, p1, p2)
                }
                fmt.Println(p0, p0)
            } else {
                p1 -= (p1 / ((1) | 1))
                for fi17 := 0; fi17 < 7; fi17++ {
                    for fi18 := 0; fi18 < 16; fi18++ {
                        p1 += (p1 << ((255) & 63))
                        var v19 uint32 = 1
                        var v20 float32 = ((((float32((1e10 * 1e10)) * ((float32(0.1) / (-2.25 / (-2.25 * 1e10))) / (float32(-2.25) - 0.0))) * 0.0) * (-2.25 - float32((-2.25 - 1.5)))) * (float32((float64(0.1) - 1.5)) - ((1e10 / -2.25) * (1e10 + 1e10))))
                        _ = len(v5)
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    if len(v5) > 0 {
                        v5[0] = (fi17 & (fi17 % ((128) | 1)))
                    }
                    _ = fi17
                }
                if (!((160 >= 0) || false)) {
                    switch 128 {
                    case 128:
                        p1 = (p1 * p1)
                    case 42:
                        if (false || (((3.14159 / float32(0.1)) + float32(3.14159)) != (3.14159 - (0.0 + (1.5 + 0.1))))) {
                            v21 := GenS2{F0: -128}
                            var v22 int16 = -1
                            _ = v21
                            _ = v22
                        }
                    case 631:
                        var v23 uint = 2
                        _ = v23
                    default:
                        if false {
                            var v24 uint32 = 221
                            v25 := GenS1{F0: 42, F1: (p2 - (p2 + p2)), F2: -952}
                            switch "hello" {
                            case "\n\t":
                                p2 = (p2 ^ p2)
                            }
                            _ = v24
                            _ = v25
                        }
                    }
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p0
}

func main() {
    var v1 int = 947
    v2 := make(map[uint32]int)
    var v3 int64 = -9223372036854775808
    for fi26 := 0; fi26 < 2; fi26++ {
        switch "foo bar" {
        case "":
            if false {
                v27 := []int{v1, v1, 127, 32767}
                for fi28 := 0; fi28 < 11; fi28++ {
                    switch 2147483647 {
                    case 32767:
                        v29 := GenS2{F0: -1}
                        _ = v29
                    case 375:
                        fmt.Println(v3)
                    case 0:
                        v1 = (fi26 & fi26)
                    }
                    if false {
                        v30 := GenS1{F0: 486, F1: -2147483648, F2: 1}
                        switch fi28 {
                        case -267:
                            if false {
                                var v31 string = "foo bar"
                                _ = v31
                            } else {
                                fmt.Println(fi26, v1)
                                v3 -= v3
                            }
                        }
                        _ = v30
                    } else {
                        var v32 uint16 = 32767
                        fmt.Println(fi28, v32)
                        _ = v32
                    }
                    v1 = fi26
                    _ = fi28
                }
                delete(v2, 2147483647)
                _ = v27
            }
        }
        v33 := GenS2{F0: -61}
        _ = fi26
        _ = v33
    }
    v3 -= 42
    switch "a" {
    case "":
        v34 := []int{(^(v1)), v1, 0, (v1 % (((v1 ^ 128)) | 1))}
        _ = v34
    }
    v35 := GenS3{F0: (v1 / ((4294967295) | 1)), F1: 42}
    switch "" {
    case "a":
        switch ("hello" + "hello") {
        case "":
            switch (v1 | v1) {
            case -2147483648:
                if (42 == 42) {
                    switch v1 {
                    case 4294967295:
                        v36 := v2[0]
                        _ = v36
                    case 9223372036854775807:
                        switch 485 {
                        case -1:
                            if (genFn0((v1 | v1), -1) <= genFn0((v1 ^ (v1 / ((999) | 1))), v3)) {
                                if true {
                                    switch ("a" + "") {
                                    case "foo bar":
                                        v3 = v3
                                    }
                                } else {
                                    v1 -= v1
                                }
                            }
                        default:
                            v37 := GenS1{F0: 42, F1: -2147483648, F2: 533}
                            _ = v37
                        }
                    case -706:
                        v38 := []int{(v1 - (v1 / ((v1) | 1)))}
                        _ = v38
                    }
                    v39 := []int{v1, (v1 % ((v1) | 1)), v1, v1}
                    _ = v39
                } else {
                    var v40 int = -803
                    _ = v40
                }
            case 0:
                v41 := GenS4{F0: v3, F1: (v3 * v3)}
                _ = v41
            default:
                v42 := []int{(v1 << ((0) & 63)), (v1 / (((v1 << ((0) & 63))) | 1)), v1, (v1 + v1)}
                _ = v42
            }
        case "\n\t":
            var v43 int32 = 2147483647
            _ = v43
        case "a":
            var v44 bool = false
            _ = v44
        }
    case "hello":
        var v45 float32 = ((-2.25 + (((0.1 + ((0.1 * 0.1) + 1e10)) - 1.5) / 0.1)) * (float32(0.1) - -2.25))
        _ = v45
    default:
        v3 -= v3
    }
    if (42 == 151) {
        v3 = (^(v3))
        v3 += (v3 + v3)
    }
    v1 -= (^(v1))
    v1 = v1
    switch (v1 | v1) {
    case -1:
        var v46 float32 = 1e10
        _ = v46
    }
    v47 := GenS4{F0: -1, F1: (-(v3))}
    var v48 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v47
    _ = v48
}
