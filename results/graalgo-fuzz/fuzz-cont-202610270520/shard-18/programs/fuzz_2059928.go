package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint }
func genFn0(p0 uint8, p1 int64) uint32 {
    switch "foo bar" {
    case "a":
        v6 := []int{1, -432, 9223372036854775807, -659}
        _ = v6
    case "hello":
        for fi7 := 0; fi7 < 2; fi7++ {
            fmt.Println(p1, fi7, fi7)
            _ = fi7
        }
    case "":
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := GenS1{F0: 2147483648}
            switch "a" {
            case "foo bar":
                switch "" {
                case "a":
                    var v10 uint32 = 65535
                    _ = v10
                }
            case "\n\t":
                fmt.Println(p1, p0, p0)
            default:
                v11 := GenS2{F0: 95, F1: 65535}
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    default:
        p0 = p0
    }
    if (!true) {
        p0 = (p0 % ((1) | 1))
    } else {
        for fi12 := 0; fi12 < 3; fi12++ {
            switch 128 {
            case 211:
                var v13 float32 = 0.0
                _ = v13
            }
            _ = fi12
        }
        var v14 string = "a"
        _ = v14
    }
    var v15 string = "\n\t"
    v16 := []int{1, 0}
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return 1
}
func genFn1(p0 bool) int32 {
    fmt.Println(p0)
    p0 = true
    var v17 float64 = -2.25
    v18 := GenS1{F0: 65535}
    _ = p0
    _ = v17
    _ = v18
    return 2147483647
}

func main() {
    var v1 uint32 = 1
    var v2 int8 = 108
    v3 := []string{}
    v4 := make(map[int]int)
    v5 := []int{9223372036854775807, 42, 945}
    var v19 int8 = (v2 + (-(v2)))
    if false {
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 float64 = 0.1
            for fi22 := 0; fi22 < 18; fi22++ {
                switch ("foo bar" + ("" + ("foo bar" + ""))) {
                case "":
                    switch "\n\t" {
                    case "foo bar":
                        if false {
                            var v23 uint16 = 524
                            v2 += v2
                            fmt.Println(v23, v21)
                            _ = v23
                        } else {
                            var v24 uint16 = 65535
                            _ = v24
                        }
                    }
                default:
                    _ = len(v4)
                    var v25 uint = 0
                    switch ("\n\t" + "a") {
                    case "hello":
                        if false {
                            delete(v4, (fi22 - (-(fi22))))
                        }
                    case "":
                        var v26 uint = 1
                        _ = v26
                    case "a":
                        _ = len(v4)
                    }
                    _ = v25
                }
                _ = fi22
            }
            v1 = (^(v1))
            _ = fi20
            _ = v21
        }
    }
    if true {
        v19 -= 97
        switch "a" {
        case "foo bar":
            v2 = v19
        }
    } else {
        v19 = 34
        v19 -= v2
        var v27 uint8 = 0
        _ = v27
    }
    fmt.Println(v19, v19, v1)
    _ = len(v5)
    v28 := v4[853]
    v2 = v2
    v2 -= 127
    if (-2.25 < 1e10) {
        switch ((("" + "\n\t") + "hello") + "foo bar") {
        case "foo bar":
            v2 -= 84
        case "\n\t":
            _ = len(v4)
        case "hello":
            var v29 int = -669
            _ = v29
        default:
            v28 += v28
        }
        if false {
            var v30 float32 = 3.14159
            fmt.Println(v1, v30, v28)
            var v31 bool = false
            _ = v30
            _ = v31
        }
        v1 -= v1
    } else {
        var v32 uint8 = 0
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
