package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 bool }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 bool; F1 int; F2 uint32 }
type GenS6 struct { F0 int32; F1 int16; F2 int16 }
func genFn0(p0 uint8, p1 int16) bool {
    if (797 <= -2147483648) {
        for fi5 := 0; fi5 < 8; fi5++ {
            p1 = (p1 - 1)
            var v6 int64 = 42
            _ = fi5
            _ = v6
        }
        if (874 == -150) {
            v7 := GenS1{F0: -2147483648, F1: 42, F2: true}
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 int16 = (p1 * p1)
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 6; fi10++ {
                p0 += (p0 * (p0 ^ (p0 & p0)))
                fmt.Println(p0, p1)
                p0 -= p0
                _ = fi10
            }
            _ = v7
        }
        for fi11 := 0; fi11 < 1; fi11++ {
            switch "\n\t" {
            case "\n\t":
                for fi12 := 0; fi12 < 17; fi12++ {
                    if true {
                        var v13 bool = true
                        _ = v13
                    }
                    if true {
                        p1 -= p1
                        var v14 uint32 = 1
                        _ = v14
                    }
                    if ((!false) && false) {
                        var v15 bool = false
                        v15 = (!v15)
                        _ = v15
                    }
                    _ = fi12
                }
            case "hello":
                p1 = p1
            case "a":
                p1 += p1
            }
            p0 -= p0
            _ = fi11
        }
    } else {
        var v16 uint32 = 4294967295
        _ = v16
    }
    if (255 > 42) {
        fmt.Println(p1)
        fmt.Println(p0, p1)
        v17 := []int{-169, 0}
        _ = v17
    }
    if true {
        v18 := GenS2{F0: ("a" + (("" + "hello") + (("a" + "a") + "a"))), F1: 0}
        if (!((42 >= 32767) && true)) {
            var v19 string = ("" + "foo bar")
            var v20 string = (v19 + "foo bar")
            p1 -= p1
            _ = v19
            _ = v20
        }
        _ = v18
    } else {
        for fi21 := 0; fi21 < 20; fi21++ {
            v22 := GenS3{F0: true}
            _ = fi21
            _ = v22
        }
    }
    _ = p0
    _ = p1
    return true
}
func genFn1(p0 bool, p1 uint) int16 {
    v23 := GenS4{F0: "", F1: -32768}
    var v24 int8 = -21
    if (65535 <= 725) {
        if (2 >= 2) {
            p1 -= (p1 & (p1 % ((18446744073709551615) | 1)))
            if p0 {
                p0 = (590 < int32(v24))
            }
        } else {
            fmt.Println(p0, p1)
            var v25 float32 = 0.1
            switch 455 {
            case 477:
                v26 := []int{1, 1}
                _ = v26
            case 0:
                v27 := []int{4294967295, 4294967295, 2147483647}
                _ = v27
            case 1:
                v28 := []int{}
                _ = v28
            default:
                v29 := GenS1{F0: 4294967295, F1: 1, F2: ("foo bar" < "foo bar")}
                _ = v29
            }
            _ = v25
        }
        v24 -= v24
    }
    v30 := GenS4{F0: (("a" + "a") + ("\n\t" + "\n\t")), F1: 499}
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    _ = v30
    return 42
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int16 = 1000
    var v3 float64 = (-2.25 + 1e10)
    v4 := []string{"hello", "hello", (("a" + ("\n\t" + ("a" + ("\n\t" + "a")))) + "\n\t"), "foo bar"}
    for fi31 := 0; fi31 < 12; fi31++ {
        _ = len(v4)
        v3 -= (3.14159 + (v3 - (v3 - v3)))
        _ = fi31
    }
    v32 := GenS5{F0: true, F1: v1, F2: 99}
    switch "\n\t" {
    case "\n\t":
        var v33 uint8 = 0
        _ = v33
    case "":
        for fi34 := 0; fi34 < 4; fi34++ {
            v3 = ((0.0 - v3) - v3)
            switch "a" {
            case "foo bar":
                var v35 int = 4294967295
                _ = v35
            default:
                var v36 uint64 = 4294967295
                _ = v36
            }
            _ = fi34
        }
    case "a":
        switch (v1 - 32767) {
        case 947:
            var v37 uint64 = 1
            _ = v37
        default:
            v1 = (v1 >> ((18446744073709551615) & 63))
        }
    }
    fmt.Println(v1)
    for fi38 := 0; fi38 < 17; fi38++ {
        if (!true) {
            fmt.Println(v1)
            v39 := GenS4{F0: (("foo bar" + ("foo bar" + "")) + "\n\t"), F1: (-(v2))}
            for fi40 := 0; fi40 < 17; fi40++ {
                v4 = append(v4, "\n\t")
                fmt.Println(fi38)
                _ = len(v4)
                if (false || (fi38 >= (-(fi40)))) {
                    var v41 uint8 = 1
                    _ = v41
                } else {
                    var v42 int32 = 42
                    _ = v42
                }
                _ = fi40
            }
            _ = v39
        } else {
            v43 := GenS6{F0: int32(v1), F1: genFn1(((true && (!true)) && true), 1), F2: 117}
            for fi44 := 0; fi44 < 1; fi44++ {
                v1 -= 32768
                _ = fi44
            }
            _ = v43
        }
        _ = fi38
    }
    for fi45 := 0; fi45 < 20; fi45++ {
        var v46 float32 = -2.25
        _ = len(v4)
        _ = fi45
        _ = v46
    }
    v47 := GenS6{F0: 0, F1: (-(v2)), F2: (v2 % (((v2 / ((v2) | 1))) | 1))}
    var v48 int64 = -825
    var v49 bool = (false && false)
    fmt.Println(v48)
    v49 = v49
    for fi50 := 0; fi50 < 15; fi50++ {
        var v51 uint = 255
        for fi52 := 0; fi52 < 20; fi52++ {
            var v53 uint64 = 4294967295
            var v54 uint64 = v53
            if (!((2147483648 != 2) && true)) {
                var v55 uint32 = 838
                var v56 string = "a"
                if len(v4) > 0 {
                    v4[0] = ""
                }
                _ = v55
                _ = v56
            }
            _ = fi52
            _ = v53
            _ = v54
        }
        v57 := GenS6{F0: 0, F1: (v2 & (v2 << ((1) & 63))), F2: v2}
        _ = fi50
        _ = v51
        _ = v57
    }
    fmt.Println(v49, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v47
    _ = v48
    _ = v49
}
