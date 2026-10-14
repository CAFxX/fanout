package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int32 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 string }
func genFn0(p0 int8) string {
    v5 := GenS1{F0: 0, F1: 42, F2: 42}
    if (!(!(!(((949 > 42) || (-2147483648 != -9223372036854775808)) || false)))) {
        switch ("hello" + "") {
        case "a":
            fmt.Println(p0)
        }
        var v6 uint16 = 332
        switch ("\n\t" + "") {
        case "hello":
            var v7 uint = 42
            _ = v7
        case "a":
            v8 := GenS1{F0: 2147483647, F1: -528, F2: 1}
            _ = v8
        default:
            var v9 float64 = ((((-2.25 + -2.25) + 0.1) + -2.25) / 1.0)
            _ = v9
        }
        _ = v6
    } else {
        p0 -= (p0 >> ((0) & 63))
        p0 = (p0 >> ((153) & 63))
        p0 -= p0
    }
    var v10 bool = (!true)
    if ((!v10) || v10) {
        p0 += (^(p0))
        var v11 bool = (("\n\t" + (("hello" + "foo bar") + ("hello" + ("foo bar" + "foo bar")))) <= ("" + ""))
        _ = v11
    }
    v10 = (!v10)
    _ = p0
    _ = v10
    _ = v5
    return ("hello" + "")
}
func genFn1(p0 uint, p1 int32) int64 {
    p1 -= p1
    v12 := GenS2{F0: ("\n\t" + "foo bar"), F1: ((false && false) && (42 != 42))}
    switch -76 {
    case 1:
        p1 = p1
    case 42:
        v13 := []int{9223372036854775807, 128}
        _ = v13
    case 487:
        p1 -= 2147483647
    default:
        v14 := []int{}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v12
    return 42
}

func main() {
    var v1 int8 = 127
    var v2 bool = (42 <= -478)
    var v3 int = -372
    v4 := []string{((("\n\t" + "\n\t") + "hello") + ""), "foo bar"}
    if (false || v2) {
        if true {
            v15 := GenS3{F0: v3, F1: v2}
            v16 := GenS3{F0: 4294967295, F1: true}
            _ = v15
            _ = v16
        }
        v1 -= (-(v1))
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        fmt.Println(v3, v1)
        _ = fi17
    }
    v2 = true
    v2 = ((v2 && ((!(true || v2)) || ((("foo bar" + "hello") + "") < "foo bar"))) && (65535 >= 338))
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 int8 = v1
        v1 += v1
        _ = fi18
        _ = v19
    }
    v1 = v1
    if (100 < v1) {
        if v2 {
            v20 := GenS4{F0: 65535, F1: (65535 >= 662)}
            if false {
                v3 -= (v3 * v3)
            } else {
                if len(v4) > 0 {
                    v4[0] = ("foo bar" + (("hello" + ("foo bar" + "hello")) + ("hello" + "")))
                }
                for fi21 := 0; fi21 < 13; fi21++ {
                    if v2 {
                        var v22 uint16 = 2
                        _ = v22
                    }
                    _ = fi21
                }
                v3 = v3
            }
            var v23 uint = 255
            _ = v20
            _ = v23
        }
    } else {
        v2 = (!false)
    }
    v24 := GenS5{F0: "hello"}
    fmt.Println(v3, v1)
    for fi25 := 0; fi25 < 8; fi25++ {
        switch "foo bar" {
        case "foo bar":
            v2 = ((255 == 255) && false)
        }
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
