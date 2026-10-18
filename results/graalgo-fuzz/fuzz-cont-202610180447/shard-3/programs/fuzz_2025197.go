package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 string; F1 uint64; F2 string }
func genFn0() uint {
    var v4 uint16 = 2
    switch ("a" + ("\n\t" + "foo bar")) {
    case "foo bar":
        for fi5 := 0; fi5 < 3; fi5++ {
            switch 32767 {
            case 2147483647:
                if (!(3.14159 >= 0.0)) {
                    for fi6 := 0; fi6 < 17; fi6++ {
                        v7 := []int{(fi5 % (((fi6 >> ((65535) & 63))) | 1)), fi5, fi5}
                        _ = fi6
                        _ = v7
                    }
                    v8 := []int{(fi5 + 42), (-(fi5)), -2147483648}
                    var v9 int = fi5
                    _ = v8
                    _ = v9
                } else {
                    var v10 float64 = 3.14159
                    var v11 uint8 = 0
                    _ = v10
                    _ = v11
                }
            case -2147483648:
                fmt.Println(v4)
            default:
                if true {
                    switch "\n\t" {
                    case "a":
                        v12 := GenS1{F0: (^(fi5)), F1: 392}
                        _ = v12
                    case "\n\t":
                        v4 = v4
                    case "foo bar":
                        v13 := []int{9223372036854775807, 127, (fi5 - (-(fi5))), 163}
                        _ = v13
                    default:
                        for fi14 := 0; fi14 < 19; fi14++ {
                            v15 := []int{-2147483648, fi5}
                            _ = fi14
                            _ = v15
                        }
                    }
                    v16 := []int{128, 42, -2147483648}
                    _ = v16
                } else {
                    v17 := []int{fi5, fi5}
                    v18 := GenS1{F0: 4294967295, F1: 731}
                    _ = v17
                    _ = v18
                }
            }
            var v19 uint64 = uint64(v4)
            if false {
                v20 := []int{(^(fi5))}
                v19 = 597
                for fi21 := 0; fi21 < 3; fi21++ {
                    var v22 float32 = 3.14159
                    _ = fi21
                    _ = v22
                }
                _ = v20
            } else {
                switch 1 {
                case -9223372036854775808:
                    if false {
                        switch ("hello" + "") {
                        case "":
                            v23 := GenS1{F0: (fi5 ^ fi5), F1: 0}
                            _ = v23
                        }
                    }
                case 156:
                    var v24 int64 = -1
                    _ = v24
                }
            }
            _ = fi5
            _ = v19
        }
    }
    v4 += v4
    v4 -= v4
    v4 = (v4 >> ((614) & 63))
    _ = v4
    return 731
}
func genFn1(p0 int64, p1 uint8) string {
    p1 -= (^(p1))
    var v25 uint8 = 42
    p1 += v25
    _ = p0
    _ = p1
    _ = v25
    return "foo bar"
}

func main() {
    var v1 int8 = 0
    var v2 float64 = 0.1
    var v3 bool = ("a" != "hello")
    v26 := []int{}
    var v27 float32 = (1e10 / -2.25)
    v26 = append(v26, 171)
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    if ((52 < 395) || v3) {
        v28 := GenS2{F0: "foo bar", F1: 2, F2: "foo bar"}
        v29 := GenS2{F0: ("a" + ("foo bar" + "a")), F1: 18446744073709551615, F2: ((("" + "a") + "a") + ((("foo bar" + "hello") + ("foo bar" + "hello")) + ("" + "hello")))}
        _ = v28
        _ = v29
    }
    v1 = (v1 * v1)
    var v30 int16 = 32767
    v1 -= v1
    for fi31 := 0; fi31 < 12; fi31++ {
        switch fi31 {
        case 9223372036854775807:
            var v32 string = "a"
            _ = v32
        default:
            var v33 uint32 = 42
            _ = v33
        }
        var v34 uint64 = 735
        v1 += (v1 % (((-(v1))) | 1))
        _ = fi31
        _ = v34
    }
    fmt.Println(v30, v3)
    for fi35 := 0; fi35 < 10; fi35++ {
        var v36 string = (("foo bar" + "a") + "a")
        v27 -= -2.25
        fmt.Println(fi35, v2)
        _ = fi35
        _ = v36
    }
    var v37 int16 = (v30 - v30)
    var v38 int64 = -430
    fmt.Println(v2, v37)
    for fi39 := 0; fi39 < 9; fi39++ {
        var v40 int32 = 42
        var v41 float32 = v27
        _ = fi39
        _ = v40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v37
    _ = v38
}
