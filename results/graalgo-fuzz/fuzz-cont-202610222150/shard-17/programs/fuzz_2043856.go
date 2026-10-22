package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 string }
func genFn0(p0 int) int32 {
    for fi4 := 0; fi4 < 14; fi4++ {
        v5 := GenS1{F0: 54}
        fmt.Println(p0, p0)
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 int16 = 32767
            switch "\n\t" {
            case "a":
                v7 -= v7
            }
            _ = fi6
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    var v8 int8 = 127
    v9 := []int{p0, (p0 >> ((0) & 63)), (p0 & (p0 % ((p0) | 1))), (p0 - (^(p0)))}
    fmt.Println(v8)
    _ = p0
    _ = v8
    _ = v9
    return -108
}
func genFn1(p0 string) bool {
    fmt.Println(p0)
    switch p0 {
    case "":
        if (-1 <= 4294967295) {
            var v10 int8 = 1
            v11 := []int{-9223372036854775808}
            switch 9223372036854775807 {
            case 127:
                _ = len(v11)
            case 378:
                v12 := GenS1{F0: 573}
                _ = v12
            case -347:
                v13 := GenS2{F0: 606, F1: 128}
                _ = v13
            default:
                p0 = "a"
            }
            _ = v10
            _ = v11
        }
    default:
        p0 = ""
    }
    v14 := GenS1{F0: 920}
    _ = p0
    _ = v14
    return true
}

func main() {
    var v1 int8 = 1
    var v2 uint8 = 128
    var v3 bool = false
    var v15 uint16 = 32767
    if v3 {
        var v16 int16 = 0
        fmt.Println(v16, v16)
        _ = v16
    } else {
        var v17 float64 = 0.0
        var v18 uint16 = v15
        _ = v17
        _ = v18
    }
    if v3 {
        if false {
            for fi19 := 0; fi19 < 3; fi19++ {
                v20 := []int{32767}
                v3 = v3
                _ = fi19
                _ = v20
            }
        } else {
            switch -1 {
            case 591:
                v21 := GenS1{F0: 0}
                _ = v21
            case 2147483647:
                for fi22 := 0; fi22 < 11; fi22++ {
                    var v23 bool = false
                    _ = fi22
                    _ = v23
                }
            case 9223372036854775807:
                switch "foo bar" {
                case "foo bar":
                    v24 := GenS2{F0: 2147483647, F1: -845}
                    _ = v24
                case "\n\t":
                    var v25 int16 = -34
                    _ = v25
                default:
                    v26 := GenS1{F0: 2}
                    _ = v26
                }
            }
            v1 = (v1 / ((-1) | 1))
            var v27 uint8 = (v2 & 1)
            _ = v27
        }
        for fi28 := 0; fi28 < 19; fi28++ {
            v1 += (v1 % ((100) | 1))
            _ = fi28
        }
    } else {
        switch ("" + "") {
        case "a":
            v2 += (v2 - uint8(v15))
        }
    }
    fmt.Println(v1)
    for fi29 := 0; fi29 < 11; fi29++ {
        fmt.Println(v15, v3)
        switch (("" + "hello") + ("hello" + "")) {
        case "hello":
            var v30 uint64 = 1
            _ = v30
        default:
            var v31 float32 = ((0.0 / (0.1 + 1.5)) + 3.14159)
            _ = v31
        }
        for fi32 := 0; fi32 < 18; fi32++ {
            fmt.Println(fi29)
            var v33 uint16 = v15
            v34 := GenS2{F0: 429, F1: -9223372036854775808}
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = fi29
    }
    fmt.Println(v2, v1)
    if v3 {
        v35 := GenS1{F0: uint64(v15)}
        var v36 float64 = 0.1
        _ = v35
        _ = v36
    } else {
        v37 := GenS3{F0: "foo bar"}
        _ = v37
    }
    v1 = (v1 ^ (^(v1)))
    v38 := []int{-313, 1}
    for fi39 := 0; fi39 < 5; fi39++ {
        var v40 int = (fi39 * (fi39 + (fi39 & fi39)))
        _ = fi39
        _ = v40
    }
    v1 = (v1 - (^(v1)))
    if len(v38) > 0 {
        v38[0] = -1
    }
    if len(v38) > 0 {
        v38[0] = 32768
    }
    v1 = (-(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v38
}
