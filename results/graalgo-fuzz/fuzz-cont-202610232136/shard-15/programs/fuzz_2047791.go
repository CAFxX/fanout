package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 bool }
type GenS2 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint32; F1 string; F2 int16 }
type GenS4 struct { F0 bool; F1 uint; F2 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 uint8; F2 int8 }
type GenS7 struct { F0 int64; F1 uint16; F2 string }
type GenS8 struct { F0 uint8; F1 int16; F2 int8 }
type GenS9 struct { F0 int8; F1 bool; F2 bool }
func genFn0(p0 uint) uint64 {
    for fi4 := 0; fi4 < 19; fi4++ {
        p0 = (p0 / ((795) | 1))
        _ = fi4
    }
    var v5 uint8 = 42
    _ = p0
    _ = v5
    return 18446744073709551615
}
func genFn1(p0 uint16) int8 {
    p0 += p0
    p0 += 32767
    switch 32767 {
    case -2147483648:
        switch -2147483648 {
        case -137:
            var v6 int = 32768
            _ = v6
        default:
            var v7 int32 = 2147483647
            _ = v7
        }
    case 128:
        for fi8 := 0; fi8 < 12; fi8++ {
            p0 = 159
            v9 := GenS1{F0: 2147483647, F1: 59, F2: false}
            p0 -= 42
            _ = fi8
            _ = v9
        }
    default:
        switch "a" {
        case "hello":
            switch 2147483647 {
            case 0:
                v10 := []int{2147483647, -9223372036854775808, 32768, 4294967295}
                _ = v10
            default:
                fmt.Println(p0)
            }
        case "\n\t":
            v11 := GenS2{F0: 18446744073709551615, F1: 0, F2: p0}
            _ = v11
        default:
            v12 := GenS2{F0: genFn0(109), F1: -1, F2: (^(p0))}
            _ = v12
        }
    }
    _ = p0
    return 100
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint = 4294967295
    v3 := []int{-184, -707}
    v2 += v2
    if (((("a" + ("foo bar" + "\n\t")) <= ("" + "")) && ("\n\t" < "\n\t")) && true) {
        var v13 uint8 = 255
        _ = len(v3)
        var v14 string = ("\n\t" + ("a" + "\n\t"))
        _ = v13
        _ = v14
    }
    v15 := GenS3{F0: 65535, F1: (("a" + "a") + ""), F2: -32768}
    var v16 float64 = (v1 + 3.14159)
    if false {
        fmt.Println(v1, v2)
        for fi17 := 0; fi17 < 9; fi17++ {
            v18 := GenS2{F0: 216, F1: 42, F2: 1}
            var v19 int8 = 89
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    if (true && false) {
        if true {
            v20 := GenS4{F0: ((2147483648 < 186) || true), F1: v2, F2: 220}
            for fi21 := 0; fi21 < 18; fi21++ {
                if true {
                    for fi22 := 0; fi22 < 17; fi22++ {
                        var v23 float64 = (float64(1e10) - 1.5)
                        v2 += (v2 * (v2 | 0))
                        _ = fi22
                        _ = v23
                    }
                    v24 := GenS5{F0: -1}
                    v16 = -2.25
                    _ = v24
                } else {
                    if false {
                        v3 = append(v3, (fi21 | (fi21 >> (((^(v2))) & 63))))
                    }
                    for fi25 := 0; fi25 < 17; fi25++ {
                        var v26 int8 = 42
                        if true {
                            v2 -= v2
                        }
                        var v27 uint16 = 167
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                }
                _ = fi21
            }
            _ = v20
        } else {
            fmt.Println(v1, v2)
        }
        if ((v2 | (v2 & (v2 << ((v2) & 63)))) < (v2 / (((^(v2))) | 1))) {
            if true {
                var v28 int8 = genFn1(960)
                _ = v28
            }
            if (42 < 2147483647) {
                var v29 int32 = -337
                v30 := GenS6{F0: "", F1: 0, F2: 100}
                _ = v29
                _ = v30
            }
            fmt.Println(v2, v2)
        } else {
            var v31 float32 = (1e10 - (0.1 + 1e10))
            _ = v31
        }
        v16 = (v16 / (v1 / 3.14159))
    } else {
        for fi32 := 0; fi32 < 2; fi32++ {
            var v33 uint = (v2 + (v2 + 415))
            v16 = ((((float64(0.1) - (((float64((1.5 + 1e10)) - (3.14159 / 1.5)) - (float64((float32((v16 + v1)) - 3.14159)) / 1.0)) + (((v16 + 3.14159) - 3.14159) - v16))) / (v1 / (v1 * 0.0))) + v1) + 0.0)
            _ = fi32
            _ = v33
        }
        switch "\n\t" {
        case "foo bar":
            if true {
                v34 := GenS2{F0: 0, F1: -1, F2: 255}
                _ = v34
            }
        default:
            var v35 uint8 = 60
            _ = v35
        }
    }
    v2 = v2
    v36 := GenS7{F0: 0, F1: 1, F2: ((("foo bar" + "foo bar") + "\n\t") + ((("\n\t" + "a") + ("hello" + ("hello" + "a"))) + "\n\t"))}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi37 := 0; fi37 < 14; fi37++ {
        fmt.Println(v2)
        v38 := GenS7{F0: -229, F1: 876, F2: "\n\t"}
        for fi39 := 0; fi39 < 15; fi39++ {
            if true {
                if (false && false) {
                    v1 += float64(float32(float64(1e10)))
                }
            } else {
                v16 -= (-2.25 - 0.1)
            }
            var v40 uint = (v2 | v2)
            _ = fi39
            _ = v40
        }
        _ = fi37
        _ = v38
    }
    for fi41 := 0; fi41 < 18; fi41++ {
        if false {
            v16 += v16
            if ((false || (4294967295 <= 251)) || false) {
                var v42 uint32 = 0
                v43 := GenS8{F0: 128, F1: -890, F2: 42}
                _ = v42
                _ = v43
            }
            v44 := GenS9{F0: 42, F1: true, F2: true}
            _ = v44
        }
        var v45 string = ("a" + "foo bar")
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi41
        _ = v45
    }
    fmt.Println(v1, v2)
    var v46 uint = v2
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v46
}
