package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int32; F1 string; F2 int }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint32, p1 string, p2 uint8) int16 {
    v4 := []int{128, 66}
    var v5 int16 = -1
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(p2)
        if false {
            v7 := GenS1{F0: (fi6 / (((fi6 % ((fi6) | 1))) | 1)), F1: 2}
            v8 := GenS2{F0: 604, F1: p1, F2: fi6}
            _ = v7
            _ = v8
        }
        _ = fi6
    }
    if (18446744073709551615 > 42) {
        p1 = ("foo bar" + p1)
        var v9 uint = uint(p0)
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (v5 - (^(v5)))
}

func main() {
    var v1 bool = false
    v2 := []string{"foo bar", "hello"}
    var v3 uint8 = 255
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 uint64 = 808
        switch "hello" {
        case "a":
            var v12 int16 = -32768
            _ = v12
        }
        switch (("a" + "foo bar") + "a") {
        case "":
            v11 -= (v11 - v11)
        case "a":
            var v13 float64 = 0.0
            _ = v13
        default:
            fmt.Println(v1, v1, v11)
        }
        _ = fi10
        _ = v11
    }
    if (9223372036854775807 >= -307) {
        if (-1 == 907) {
            if (!v1) {
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                fmt.Println(v1, v3)
                for fi14 := 0; fi14 < 5; fi14++ {
                    switch (((((("" + "\n\t") + "\n\t") + ("foo bar" + ("a" + "\n\t"))) + "foo bar") + "") + (("a" + ("" + ("hello" + "\n\t"))) + ("a" + ("\n\t" + "hello")))) {
                    case "":
                        if true {
                            for fi15 := 0; fi15 < 1; fi15++ {
                                v1 = v1
                                _ = fi15
                            }
                            v1 = (773 < 65535)
                        }
                    default:
                        fmt.Println(v1, v1, v1)
                    }
                    var v16 int8 = 100
                    _ = fi14
                    _ = v16
                }
            }
            v3 += v3
            var v17 float32 = 1e10
            _ = v17
        } else {
            v3 += v3
            _ = len(v2)
        }
        var v18 float32 = 0.0
        var v19 int64 = -784
        _ = v18
        _ = v19
    }
    if (42 >= 0) {
        var v20 int8 = -1
        _ = v20
    } else {
        var v21 uint16 = 1
        v1 = v1
        _ = v21
    }
    v2 = append(v2, "hello")
    if v1 {
        if v1 {
            switch 0 {
            case -2147483648:
                v2 = append(v2, ("foo bar" + "\n\t"))
            case 405:
                v1 = (v1 || false)
            case 795:
                fmt.Println(v1)
            }
            v3 -= (^(v3))
            v1 = (!v1)
        }
        v22 := GenS2{F0: -1, F1: "\n\t", F2: 9223372036854775807}
        _ = v22
    } else {
        for fi23 := 0; fi23 < 18; fi23++ {
            fmt.Println(v1)
            v3 = 1
            _ = fi23
        }
    }
    switch "\n\t" {
    case "":
        var v24 int64 = -9223372036854775808
        _ = v24
    case "a":
        if ("a" != "foo bar") {
            if len(v2) > 0 {
                v2[0] = (("hello" + "foo bar") + "")
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v25 := GenS3{F0: 824, F1: 1}
            _ = v25
        } else {
            if (!v1) {
                v1 = (v1 && v1)
                _ = len(v2)
            } else {
                v1 = v1
                var v26 int8 = 0
                var v27 bool = v1
                _ = v26
                _ = v27
            }
        }
    }
    v28 := GenS4{F0: 481}
    if v1 {
        v3 += 255
    }
    var v29 float64 = (0.0 + 1.5)
    var v30 int64 = 1
    var v31 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
