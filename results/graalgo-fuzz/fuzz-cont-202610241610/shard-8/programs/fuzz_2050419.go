package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
func genFn0(p0 int64, p1 uint8) uint {
    v7 := GenS1{F0: "a", F1: 831}
    switch 1 {
    case 409:
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint16 = 210
            _ = fi8
            _ = v9
        }
    case -33:
        p1 = 127
    case -9223372036854775808:
        if true {
            if ((255 != (p1 >> ((4294967295) & 63))) && (false && (!false))) {
                switch -1 {
                case 0:
                    v10 := GenS1{F0: "", F1: -830}
                    _ = v10
                case -651:
                    switch ("a" + "") {
                    case "\n\t":
                        v11 := []int{}
                        _ = v11
                    }
                }
            }
        }
    default:
        p0 -= p0
    }
    if true {
        p1 = 2
        if (!(1.5 >= 0.0)) {
            v12 := []int{9223372036854775807}
            v13 := GenS1{F0: "a", F1: -146}
            v14 := GenS1{F0: (("a" + ((("\n\t" + "\n\t") + ("foo bar" + ("" + ("" + "foo bar")))) + "hello")) + ((("hello" + "foo bar") + "") + "hello")), F1: 128}
            _ = v12
            _ = v13
            _ = v14
        } else {
            for fi15 := 0; fi15 < 3; fi15++ {
                p0 -= p0
                var v16 uint8 = p1
                fmt.Println(p1)
                _ = fi15
                _ = v16
            }
            v17 := GenS1{F0: "foo bar", F1: 128}
            p0 -= -9223372036854775808
            _ = v17
        }
        fmt.Println(p0, p0)
    }
    var v18 int16 = 1
    _ = p0
    _ = p1
    _ = v18
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = ((3.14159 * 1.5) / -2.25)
    var v3 int = 128
    v4 := make(map[string]int)
    v5 := make(map[int]int)
    var v6 float32 = 0.0
    switch (v3 / ((v3) | 1)) {
    case 189:
        v1["foo bar"] = "foo bar"
    case -9223372036854775808:
        v19 := GenS1{F0: "a", F1: v3}
        _ = v19
    default:
        for fi20 := 0; fi20 < 5; fi20++ {
            for fi21 := 0; fi21 < 5; fi21++ {
                fmt.Println(v2, v2)
                var v22 bool = false
                _ = fi21
                _ = v22
            }
            var v23 bool = (9223372036854775807 <= 1)
            var v24 bool = ((-1 > 785) && (!(true && false)))
            _ = fi20
            _ = v23
            _ = v24
        }
    }
    v25 := GenS1{F0: (("" + ("\n\t" + "foo bar")) + ("foo bar" + "hello")), F1: 1}
    v3 = (v3 | 2147483647)
    var v26 uint64 = 0
    v26 = 42
    v27 := []int{v3, (v3 << ((1) & 63))}
    switch ("a" + "a") {
    case "\n\t":
        v27 = append(v27, v3)
    case "hello":
        fmt.Println(v2, v26, v3)
    case "a":
        v4[("hello" + ("foo bar" + (("\n\t" + "foo bar") + ("a" + ""))))] = v3
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        var v29 string = "foo bar"
        delete(v4, (v29 + (((v29 + ((("foo bar" + "\n\t") + ((v29 + ("hello" + ("" + (v29 + v29)))) + ((("hello" + v29) + v29) + v29))) + "hello")) + (v29 + (((v29 + ("foo bar" + (v29 + ((v29 + v29) + (("foo bar" + "foo bar") + "foo bar"))))) + "hello") + ((v29 + (v29 + v29)) + v29)))) + (v29 + ""))))
        _ = len(v5)
        _ = fi28
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
