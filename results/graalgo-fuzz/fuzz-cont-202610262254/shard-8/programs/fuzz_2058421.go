package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint16; F1 uint8; F2 string }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 string; F1 int; F2 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int; F1 uint64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int8; F1 uint32 }
func genFn0(p0 int64, p1 int8, p2 uint) bool {
    if false {
        v5 := GenS1{F0: (912 >= 0)}
        _ = v5
    }
    if true {
        if ("" <= "foo bar") {
            if (((!(p0 <= p0)) || false) && ((p1 & p1) != p1)) {
                fmt.Println(p2, p0)
                v6 := []int{}
                _ = v6
            }
        }
        fmt.Println(p2, p2, p0)
        var v7 int64 = -9223372036854775808
        _ = v7
    }
    fmt.Println(p1, p2, p0)
    v8 := GenS1{F0: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return true
}
func genFn1(p0 int32, p1 uint) uint8 {
    fmt.Println(p1)
    if (((1 > 367) || false) || false) {
        for fi9 := 0; fi9 < 7; fi9++ {
            for fi10 := 0; fi10 < 4; fi10++ {
                p1 -= (p1 | p1)
                _ = fi10
            }
            _ = fi9
        }
        p1 = p1
        if false {
            v11 := GenS2{F0: -57, F1: false}
            _ = v11
        }
    } else {
        if false {
            p0 = (p0 >> ((p1) & 63))
            p0 += 85
            v12 := GenS2{F0: 100, F1: (50 < 430)}
            _ = v12
        } else {
            p0 = 0
        }
        var v13 uint8 = 0
        var v14 float64 = (3.14159 * 0.1)
        _ = v13
        _ = v14
    }
    switch "\n\t" {
    case "foo bar":
        switch "a" {
        case "hello":
            var v15 bool = false
            _ = v15
        default:
            var v16 uint32 = 1
            _ = v16
        }
    case "\n\t":
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 bool = (!false)
            switch fi17 {
            case -433:
                switch fi17 {
                case -939:
                    v19 := []int{(fi17 ^ (fi17 << (((p1 - p1)) & 63))), 778, (fi17 & fi17)}
                    _ = v19
                case 9223372036854775807:
                    for fi20 := 0; fi20 < 6; fi20++ {
                        fmt.Println(v18)
                        v18 = (!(!v18))
                        if (!true) {
                            var v21 uint = p1
                            _ = v21
                        } else {
                            p0 = (p0 & p0)
                            if v18 {
                                v18 = (1 < 9223372036854775807)
                            }
                        }
                        _ = fi20
                    }
                case -870:
                    if false {
                        v22 := GenS1{F0: v18}
                        if v18 {
                            v23 := []int{}
                            fmt.Println(v18, p0)
                            p0 -= p0
                            _ = v23
                        }
                        _ = v22
                    }
                }
            case 32768:
                switch 2147483647 {
                case 128:
                    if v18 {
                        v24 := GenS1{F0: (((0.0 - -2.25) >= 3.14159) && true)}
                        _ = v24
                    } else {
                        fmt.Println(p1, v18, p1)
                        v25 := GenS3{F0: 65535, F1: 42, F2: "foo bar"}
                        if (-77 >= 1) {
                            v26 := []int{fi17, fi17}
                            v27 := GenS4{F0: 127, F1: 130}
                            var v28 int = fi17
                            _ = v26
                            _ = v27
                            _ = v28
                        } else {
                            v29 := []int{fi17, (^(fi17))}
                            _ = v29
                        }
                        _ = v25
                    }
                case -808:
                    var v30 uint32 = 328
                    _ = v30
                case -405:
                    v31 := []int{(fi17 / ((fi17) | 1))}
                    _ = v31
                }
            default:
                for fi32 := 0; fi32 < 10; fi32++ {
                    var v33 uint16 = 2
                    v34 := GenS4{F0: 42, F1: 128}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
            _ = fi17
            _ = v18
        }
    }
    switch 127 {
    case 32767:
        p0 = p0
    }
    v35 := GenS5{F0: "\n\t", F1: 32767, F2: p0}
    _ = p0
    _ = p1
    _ = v35
    return 127
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    var v3 float64 = 0.1
    var v4 uint32 = 2
    v36 := GenS6{F0: -9223372036854775808}
    if true {
        v37 := GenS7{F0: 4294967295, F1: 69}
        var v38 int16 = 0
        _ = v37
        _ = v38
    }
    v3 -= (v3 - 0.1)
    var v39 uint32 = (v4 | v4)
    v39 -= (v39 - (v4 / ((v4) | 1)))
    var v40 uint8 = 42
    switch -978 {
    case 1:
        v41 := GenS6{F0: -733}
        _ = v41
    }
    v42 := GenS8{F0: 1}
    v43 := GenS9{F0: 98, F1: v39}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v43
}
