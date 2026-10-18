package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint32 }
func genFn0() int32 {
    if false {
        var v5 int8 = 100
        v5 -= v5
        _ = v5
    }
    v6 := GenS1{F0: true, F1: 0, F2: 2147483648}
    _ = v6
    return 42
}
func genFn1(p0 uint, p1 int64) int64 {
    v7 := GenS1{F0: (-128 < 19), F1: 1, F2: 4294967295}
    switch 0 {
    case 32767:
        p1 = 0
    case 343:
        var v8 int = 580
        _ = v8
    case -9223372036854775808:
        for fi9 := 0; fi9 < 5; fi9++ {
            for fi10 := 0; fi10 < 2; fi10++ {
                v11 := []int{}
                v11 = append(v11, fi9)
                var v12 int = (fi10 + (fi9 * (fi9 / ((-899) | 1))))
                _ = fi10
                _ = v11
                _ = v12
            }
            switch "a" {
            case "foo bar":
                var v13 float64 = float64(-2.25)
                _ = v13
            case "hello":
                switch (fi9 - fi9) {
                case -1:
                    p0 += 42
                case -58:
                    switch "\n\t" {
                    case "\n\t":
                        var v14 int64 = (p1 << ((p0) & 63))
                        _ = v14
                    case "hello":
                        var v15 bool = (!true)
                        _ = v15
                    default:
                        var v16 uint16 = 32767
                        _ = v16
                    }
                case 9223372036854775807:
                    v17 := []int{fi9, 42, fi9}
                    _ = v17
                default:
                    if (!(!(42 != 127))) {
                        var v18 uint = (p0 ^ (^(p0)))
                        _ = v18
                    } else {
                        p1 += (-(p1))
                        var v19 uint8 = 128
                        _ = v19
                    }
                }
            default:
                switch -825 {
                case 32767:
                    v20 := []int{(-(fi9))}
                    _ = v20
                case 32768:
                    p1 = (p1 - (p1 ^ (-(p1))))
                }
            }
            for fi21 := 0; fi21 < 13; fi21++ {
                p0 = p0
                if false {
                    p0 -= (p0 % (((p0 - (p0 & p0))) | 1))
                    var v22 int64 = (p1 - p1)
                    p0 = 18446744073709551615
                    _ = v22
                } else {
                    fmt.Println(p1, fi21)
                    var v23 string = ("a" + "\n\t")
                    _ = v23
                }
                v24 := GenS1{F0: false, F1: 0, F2: 2147483648}
                v25 := []int{(fi9 ^ (fi21 | fi9))}
                _ = fi21
                _ = v24
                _ = v25
            }
            _ = fi9
        }
    default:
        p1 += (^(p1))
    }
    v26 := GenS1{F0: false, F1: -103, F2: 100}
    _ = p0
    _ = p1
    _ = v26
    _ = v7
    return (p1 | p1)
}

func main() {
    var v1 int8 = 127
    var v2 float64 = 0.0
    var v3 uint = 1
    var v4 uint = (v3 << ((v3) & 63))
    var v27 int8 = (v1 / ((v1) | 1))
    v28 := GenS1{F0: false, F1: -1, F2: 42}
    for fi29 := 0; fi29 < 12; fi29++ {
        fmt.Println(v2, v3, fi29)
        v30 := GenS1{F0: true, F1: 96, F2: 829}
        _ = fi29
        _ = v30
    }
    var v31 float64 = 1e10
    var v32 uint8 = 1
    v33 := GenS1{F0: (uint32(v3) <= 839), F1: -32768, F2: 4294967295}
    if true {
        v1 = (v1 & 100)
        v34 := []int{-820, 334}
        var v35 int64 = -157
        _ = v34
        _ = v35
    } else {
        v36 := []int{}
        var v37 int = -1
        if len(v36) >= 2 {
            v36 = v36[1:len(v36)]
        }
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        for fi39 := 0; fi39 < 4; fi39++ {
            var v40 float32 = 3.14159
            v41 := []int{int(v3), (-(fi39)), (fi38 / (((fi39 << (((v3 % (((v4 % ((v4) | 1))) | 1))) & 63))) | 1)), 4294967295}
            switch "\n\t" {
            case "":
                if ((("" != "foo bar") || ((v40 + v40) == 1.5)) && true) {
                    v42 := GenS1{F0: (65535 <= 583), F1: 1, F2: 403}
                    _ = v42
                } else {
                    v3 += 364
                    v43 := GenS1{F0: false, F1: 0, F2: uint32(v3)}
                    v2 -= v2
                    _ = v43
                }
            case "foo bar":
                v3 = (^(v3))
                var v44 int64 = 42
                _ = v44
            default:
                fmt.Println(v2)
            }
            _ = fi39
            _ = v40
            _ = v41
        }
        var v45 string = "a"
        v46 := []int{(^(fi38)), (fi38 ^ 2147483647), fi38}
        _ = fi38
        _ = v45
        _ = v46
    }
    for fi47 := 0; fi47 < 9; fi47++ {
        var v48 uint16 = 481
        _ = fi47
        _ = v48
    }
    var v49 int32 = 42
    switch 128 {
    case -422:
        var v50 int = -577
        _ = v50
    case -1:
        var v51 uint = 73
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v49
}
