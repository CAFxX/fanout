package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 bool; F1 bool }
func genFn0(p0 string, p1 uint) uint {
    var v5 uint32 = 65535
    switch 128 {
    case -2147483648:
        p1 -= 446
    }
    _ = p0
    _ = p1
    _ = v5
    return p1
}
func genFn1(p0 int32) uint8 {
    v6 := []int{285}
    p0 = (p0 + (^(p0)))
    switch ("foo bar" + ("hello" + "foo bar")) {
    case "\n\t":
        var v7 int = -1
        _ = v7
    case "":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case "a":
        p0 -= p0
    default:
        fmt.Println(p0)
    }
    switch -9223372036854775808 {
    case 32767:
        v8 := GenS1{F0: "hello"}
        _ = v8
    case 951:
        var v9 float32 = 0.0
        _ = v9
    }
    _ = p0
    _ = v6
    return 69
}

func main() {
    v1 := []string{"\n\t", "foo bar", "\n\t", ("a" + (("a" + "foo bar") + "foo bar"))}
    var v2 int64 = 42
    var v3 int16 = 0
    v4 := make(map[uint32]bool)
    v2 += (v2 << ((4294967295) & 63))
    fmt.Println(v3)
    switch -1 {
    case -807:
        switch -112 {
        case -2147483648:
            switch "a" {
            case "":
                v4[2] = false
            case "a":
                v10 := GenS2{F0: (((0.1 - (0.0 + -2.25)) < (1e10 / (3.14159 + (-2.25 - 1.5)))) || false), F1: ("a" > "foo bar")}
                _ = v10
            case "foo bar":
                if false {
                    v11 := GenS3{F0: 1000, F1: v2}
                    v3 -= 0
                    _ = v11
                }
            }
        case -1:
            if false {
                switch 32768 {
                case 284:
                    var v12 uint32 = 138
                    _ = v12
                case 32767:
                    v2 += (^(v2))
                case 4294967295:
                    for fi13 := 0; fi13 < 5; fi13++ {
                        v3 -= v3
                        _ = fi13
                    }
                }
                if false {
                    var v14 int32 = -159
                    _ = v14
                } else {
                    var v15 uint16 = 65535
                    fmt.Println(v2, v3, v2)
                    for fi16 := 0; fi16 < 7; fi16++ {
                        fmt.Println(v2, fi16)
                        fmt.Println(v2)
                        v2 += v2
                        _ = fi16
                    }
                    _ = v15
                }
                fmt.Println(v2, v3)
            } else {
                v17 := GenS4{F0: 1}
                v18 := GenS5{F0: genFn1(0)}
                v19 := GenS6{F0: false, F1: 87}
                _ = v17
                _ = v18
                _ = v19
            }
        case 32768:
            for fi20 := 0; fi20 < 18; fi20++ {
                for fi21 := 0; fi21 < 19; fi21++ {
                    var v22 int64 = -9223372036854775808
                    _ = fi21
                    _ = v22
                }
                _ = fi20
            }
        default:
            for fi23 := 0; fi23 < 11; fi23++ {
                if true {
                    var v24 int8 = 116
                    switch fi23 {
                    case 46:
                        if len(v1) > 0 {
                            v1[0] = ("foo bar" + "hello")
                        }
                    case 444:
                        v25 := GenS6{F0: false, F1: 496}
                        _ = v25
                    }
                    _ = v24
                } else {
                    v26 := GenS4{F0: -274}
                    v2 = v2
                    v3 = 1000
                    _ = v26
                }
                _ = len(v4)
                _ = fi23
            }
        }
    case 1:
        switch -323 {
        case 550:
            fmt.Println(v3)
        case -555:
            v3 = (v3 ^ (v3 - (v3 * v3)))
        default:
            v2 = (v2 + (v2 << ((0) & 63)))
        }
    default:
        if false {
            _ = len(v4)
        }
    }
    if (v2 < v2) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v2 -= v2
        v3 += -32768
    }
    v27 := GenS7{F0: (!true), F1: true}
    v1 = append(v1, "\n\t")
    var v28 int16 = (v3 * v3)
    var v29 float64 = 3.14159
    var v30 uint8 = 128
    if false {
        v31 := v4[556]
        v32 := GenS3{F0: v3, F1: -9223372036854775808}
        v28 += v3
        _ = v31
        _ = v32
    } else {
        var v33 uint32 = 2
        v34 := v4[(^(v33))]
        v2 = 0
        _ = v33
        _ = v34
    }
    fmt.Println(v29, v28, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
