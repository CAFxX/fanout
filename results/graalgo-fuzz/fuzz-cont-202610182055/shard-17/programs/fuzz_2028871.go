package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int8 }
func genFn0(p0 string, p1 uint, p2 int8) bool {
    v4 := []int{-2147483648, 127, -1}
    v5 := GenS1{F0: 1}
    fmt.Println(p0, p2, p2)
    switch ("a" + p0) {
    case "\n\t":
        var v6 float64 = -2.25
        _ = v6
    case "foo bar":
        if false {
            switch (p0 + ("" + "hello")) {
            case "foo bar":
                switch 4294967295 {
                case -1:
                    for fi7 := 0; fi7 < 6; fi7++ {
                        fmt.Println(fi7, p0)
                        switch fi7 {
                        case -9223372036854775808:
                            var v8 uint8 = 42
                            _ = v8
                        case 128:
                            switch p0 {
                            case "\n\t":
                                v9 := GenS1{F0: 42}
                                _ = v9
                            case "hello":
                                if (!(false || false)) {
                                    var v10 bool = ((false || (!false)) && false)
                                    _ = v10
                                }
                            }
                        case -2147483648:
                            p0 = ("\n\t" + p0)
                        }
                        _ = fi7
                    }
                case 614:
                    p0 = "a"
                case -9223372036854775808:
                    v11 := GenS2{F0: 18446744073709551615, F1: p2}
                    _ = v11
                }
            default:
                if true {
                    if true {
                        var v12 uint8 = 0
                        _ = v12
                    }
                    if ((!false) && true) {
                        p1 -= p1
                        var v13 uint32 = 1
                        _ = v13
                    }
                } else {
                    var v14 string = p0
                    var v15 uint32 = 1
                    _ = v14
                    _ = v15
                }
            }
            p2 -= (p2 * -11)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return true
}
func genFn1(p0 int, p1 bool, p2 int32) int8 {
    var v16 float64 = 1.5
    v16 -= 0.1
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return -1
}

func main() {
    v1 := []int{4294967295}
    var v2 uint = 2
    var v3 float64 = 0.0
    v3 = v3
    var v17 uint32 = 65535
    v3 = v3
    var v18 int8 = 0
    var v19 float64 = (float64(float32((v3 + (-2.25 / (0.0 + v3))))) - ((v3 + float64(3.14159)) - float64(0.1)))
    fmt.Println(v3)
    fmt.Println(v19)
    if (!true) {
        for fi20 := 0; fi20 < 13; fi20++ {
            for fi21 := 0; fi21 < 17; fi21++ {
                _ = len(v1)
                _ = fi21
            }
            _ = fi20
        }
    }
    switch (((("hello" + "") + "hello") + (("foo bar" + "foo bar") + "")) + "hello") {
    case "foo bar":
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 bool = (42 > 982)
            var v24 int8 = (-(v18))
            var v25 int32 = -1
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
    }
    if (uint64(v17) == 4294967295) {
        if true {
            if (1000 >= -295) {
                v2 = (^(v2))
                v3 += 0.0
                switch "a" {
                case "foo bar":
                    if true {
                        v19 = (-2.25 * 1e10)
                        for fi26 := 0; fi26 < 8; fi26++ {
                            if len(v1) > 0 {
                                v1[0] = -1
                            }
                            _ = fi26
                        }
                        var v27 uint32 = 1
                        _ = v27
                    }
                case "hello":
                    v18 += v18
                }
            }
            fmt.Println(v3)
            if (!(v17 >= v17)) {
                var v28 int16 = 42
                for fi29 := 0; fi29 < 9; fi29++ {
                    if len(v1) > 0 {
                        v1[0] = (fi29 + (fi29 / ((fi29) | 1)))
                    }
                    fmt.Println(v28)
                    var v30 float64 = 0.1
                    _ = fi29
                    _ = v30
                }
                var v31 uint8 = 0
                _ = v28
                _ = v31
            }
        }
        if true {
            var v32 int64 = 42
            _ = v32
        }
    } else {
        fmt.Println(v19)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
