package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint }
type GenS2 struct { F0 uint64; F1 int32; F2 uint }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 uint; F2 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string; F1 uint; F2 int }
func genFn0(p0 uint, p1 bool) uint8 {
    for fi5 := 0; fi5 < 19; fi5++ {
        switch "\n\t" {
        case "a":
            p1 = ((^(p0)) <= 637)
        case "\n\t":
            for fi6 := 0; fi6 < 17; fi6++ {
                p0 -= p0
                v7 := []int{-9223372036854775808, 626, fi5}
                if p1 {
                    fmt.Println(fi6)
                }
                switch "a" {
                case "\n\t":
                    p0 -= p0
                case "hello":
                    p1 = true
                case "foo bar":
                    var v8 bool = false
                    _ = v8
                default:
                    p0 = 18446744073709551615
                }
                _ = fi6
                _ = v7
            }
        case "hello":
            p0 = 392
        }
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 string = "\n\t"
            v11 := GenS1{F0: 0, F1: fi9, F2: p0}
            if p1 {
                if (65535 < 4294967295) {
                    var v12 int64 = 1
                    if p1 {
                        var v13 int = 2147483647
                        v14 := GenS1{F0: 42, F1: 2147483647, F2: p0}
                        _ = v13
                        _ = v14
                    } else {
                        v15 := []int{}
                        _ = v15
                    }
                    fmt.Println(p0, p1, p1)
                    v16 := GenS2{F0: 752, F1: -574, F2: (p0 ^ p0)}
                    _ = v12
                    _ = v16
                }
                var v17 uint8 = 224
                if (((!p1) && (!(0.0 >= (1e10 * (float64(float32(0.0)) * 3.14159))))) && p1) {
                    p1 = (1 > 4294967295)
                    v18 := GenS2{F0: 0, F1: 0, F2: 643}
                    v19 := GenS3{F0: v10, F1: 42}
                    _ = v18
                    _ = v19
                }
                _ = v17
            } else {
                v20 := GenS4{F0: (fi9 * (fi9 % (((-(fi5))) | 1)))}
                switch fi5 {
                case 146:
                    var v21 float32 = 3.14159
                    v22 := GenS5{F0: 65535, F1: 65535, F2: (!p1)}
                    p0 -= 18446744073709551615
                    _ = v21
                    _ = v22
                case 0:
                    p1 = p1
                }
                var v23 int8 = 1
                var v24 uint16 = 65535
                var v25 int = -9223372036854775808
                _ = v20
                _ = v23
                _ = v24
                _ = v25
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi5
    }
    fmt.Println(p0, p1)
    var v26 uint64 = 18446744073709551615
    v27 := []int{}
    switch "foo bar" {
    case "hello":
        var v28 uint8 = 0
        _ = v28
    case "foo bar":
        p1 = true
    default:
        _ = len(v27)
    }
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    return 255
}
func genFn1(p0 int16, p1 int32, p2 int32) int {
    var v29 uint32 = 189
    var v30 float64 = (0.0 * -2.25)
    if false {
        v31 := GenS6{F0: p0}
        p2 -= (^(p2))
        for fi32 := 0; fi32 < 7; fi32++ {
            switch "a" {
            case "a":
                p0 = p0
            case "\n\t":
                v33 := []int{}
                _ = v33
            case "hello":
                p2 -= (^(p2))
            default:
                if false {
                    p1 += p1
                    var v34 uint64 = 799
                    v30 = v30
                    _ = v34
                }
            }
            _ = fi32
        }
        _ = v31
    } else {
        p0 -= p0
        v30 = v30
    }
    fmt.Println(p1)
    var v35 uint32 = 866
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    _ = v30
    _ = v35
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 0.0
    var v3 uint16 = 42
    var v4 int64 = -244
    if false {
        delete(v1, 0)
        v36 := []int{}
        _ = len(v1)
        _ = v36
    } else {
        var v37 int = 32768
        _ = v37
    }
    v38 := GenS1{F0: 127, F1: 9223372036854775807, F2: 4294967295}
    fmt.Println(v3, v2)
    for fi39 := 0; fi39 < 12; fi39++ {
        v40 := GenS7{F0: "foo bar", F1: 255, F2: (fi39 % ((fi39) | 1))}
        v41 := GenS6{F0: 204}
        for fi42 := 0; fi42 < 15; fi42++ {
            v43 := GenS4{F0: fi39}
            v44 := []int{(-(fi42)), genFn1(1, 0, -2147483648), (fi39 ^ (-(fi39)))}
            v45 := GenS7{F0: (("\n\t" + "foo bar") + "a"), F1: 0, F2: (fi42 ^ 32767)}
            _ = fi42
            _ = v43
            _ = v44
            _ = v45
        }
        _ = fi39
        _ = v40
        _ = v41
    }
    for fi46 := 0; fi46 < 14; fi46++ {
        v3 = 1
        _ = fi46
    }
    switch genFn1(-814, 800, 0) {
    case -420:
        switch "a" {
        case "a":
            delete(v1, 2147483647)
        case "\n\t":
            v2 -= 3.14159
        case "":
            switch 9223372036854775807 {
            case 1:
                var v47 uint16 = v3
                _ = v47
            }
        default:
            v48 := v1[375]
            _ = v48
        }
    }
    if (!(!(634 <= -567))) {
        var v49 bool = true
        _ = v49
    }
    var v50 int32 = -1
    if true {
        fmt.Println(v4, v3, v2)
    } else {
        var v51 float64 = v2
        _ = v51
    }
    v1[65535] = ("\n\t" + "")
    var v52 float32 = float32(v2)
    _ = len(v1)
    v53 := GenS3{F0: "", F1: (v50 / (((v50 << ((880) & 63))) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v50
    _ = v52
    _ = v53
}
