package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 int; F2 int16 }
func genFn0(p0 int64, p1 int16) int64 {
    v5 := GenS1{F0: true, F1: 21}
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            switch fi6 {
            case 0:
                var v8 float64 = 3.14159
                var v9 int64 = (p0 & p0)
                _ = v8
                _ = v9
            case 1:
                v10 := []int{-9223372036854775808, fi6, -603}
                _ = v10
            }
            _ = fi7
        }
        p1 -= 1
        _ = fi6
    }
    switch ((("hello" + ("" + "hello")) + ("a" + "foo bar")) + ((("foo bar" + "foo bar") + (("a" + "a") + ("a" + "foo bar"))) + "")) {
    case "\n\t":
        p0 = -10
    case "a":
        if (!(127 >= 0)) {
            fmt.Println(p0, p1)
            var v11 float64 = 3.14159
            _ = v11
        } else {
            if false {
                fmt.Println(p1, p1)
                var v12 uint32 = 2147483648
                var v13 int = -1
                _ = v12
                _ = v13
            } else {
                v14 := GenS1{F0: true, F1: 1}
                _ = v14
            }
        }
    case "hello":
        if (((204 > 2147483648) || true) && false) {
            v15 := []int{699, 997, int(p0)}
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return (p0 * (p0 & p0))
}
func genFn1(p0 int8) int {
    p0 = (p0 + (p0 % (((p0 ^ p0)) | 1)))
    v16 := []int{-2147483648}
    var v17 uint16 = 2
    _ = p0
    _ = v16
    _ = v17
    return 127
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint64 = 4294967295
    v3 := make(map[int]int)
    var v4 uint64 = (v2 << ((572) & 63))
    var v18 int8 = 1
    var v19 int64 = v1
    delete(v3, genFn1(v18))
    switch "\n\t" {
    case "foo bar":
        for fi20 := 0; fi20 < 19; fi20++ {
            v3[(fi20 << ((255) & 63))] = -1
            v21 := GenS1{F0: true, F1: 0}
            v22 := []int{}
            _ = fi20
            _ = v21
            _ = v22
        }
    case "hello":
        if (false || true) {
            switch "a" {
            case "hello":
                var v23 int64 = v19
                _ = v23
            case "\n\t":
                switch "\n\t" {
                case "\n\t":
                    var v24 string = "hello"
                    _ = v24
                case "a":
                    v25 := []int{-1, 1}
                    _ = v25
                case "foo bar":
                    for fi26 := 0; fi26 < 8; fi26++ {
                        switch "a" {
                        case "":
                            v2 = (v4 & (^(v4)))
                        case "hello":
                            v27 := v3[(^(fi26))]
                            _ = v27
                        }
                        v19 -= 1
                        v28 := GenS1{F0: false, F1: v18}
                        _ = fi26
                        _ = v28
                    }
                default:
                    v29 := GenS2{F0: 42, F1: (v1 & 0)}
                    _ = v29
                }
            case "":
                var v30 int32 = 2147483647
                _ = v30
            }
            switch 127 {
            case 32768:
                v2 = 948
            default:
                for fi31 := 0; fi31 < 16; fi31++ {
                    var v32 int16 = 1000
                    var v33 int64 = (v1 - -1)
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            }
        } else {
            v34 := []int{-9223372036854775808, 9223372036854775807, -2147483648, 4294967295}
            _ = v34
        }
    }
    v35 := GenS3{F0: (!true)}
    var v36 float64 = (((3.14159 - (1.5 / 0.1)) - 1e10) + 0.0)
    var v37 float32 = 1e10
    v18 = 89
    if true {
        for fi38 := 0; fi38 < 20; fi38++ {
            for fi39 := 0; fi39 < 6; fi39++ {
                var v40 uint = uint(v4)
                _ = len(v3)
                switch fi39 {
                case 32768:
                    v1 = (v19 ^ (-(v19)))
                case 2147483647:
                    _ = len(v3)
                }
                switch ("a" + (("hello" + "") + ("foo bar" + ""))) {
                case "":
                    switch fi39 {
                    case 47:
                        v40 = v40
                    default:
                        var v41 int64 = (v1 + (v1 + (v19 * v19)))
                        _ = v41
                    }
                case "hello":
                    v42 := []int{408, fi39, 2147483647}
                    _ = v42
                default:
                    v19 -= (^(v1))
                }
                var v43 float32 = 3.14159
                switch genFn1(-128) {
                case -1:
                    if ((v4 + v4) <= 2) {
                        v36 = ((v36 / 1.0) - v36)
                        fmt.Println(v37, v18, v2)
                        fmt.Println(v1)
                    }
                }
                _ = fi39
                _ = v40
                _ = v43
            }
            for fi44 := 0; fi44 < 3; fi44++ {
                if true {
                    var v45 uint32 = 1
                    var v46 string = "foo bar"
                    _ = v45
                    _ = v46
                }
                var v47 uint64 = v2
                _ = fi44
                _ = v47
            }
            _ = fi38
        }
        v48 := GenS2{F0: 65535, F1: genFn0((v1 / ((v19) | 1)), 469)}
        v49 := GenS4{F0: true, F1: 9223372036854775807, F2: -32768}
        _ = v48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
