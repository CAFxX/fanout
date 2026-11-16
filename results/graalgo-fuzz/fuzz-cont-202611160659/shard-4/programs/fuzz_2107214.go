package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint }
func genFn0() string {
    if true {
        v4 := GenS1{F0: 0, F1: true}
        var v5 float64 = 3.14159
        switch "" {
        case "foo bar":
            v6 := []int{}
            _ = v6
        case "":
            switch "foo bar" {
            case "":
                v7 := GenS1{F0: 808, F1: false}
                _ = v7
            case "a":
                v5 = v5
            case "\n\t":
                for fi8 := 0; fi8 < 1; fi8++ {
                    fmt.Println(fi8, v5)
                    var v9 bool = ((1.5 * float64((1e10 - (0.1 * 1e10)))) >= v5)
                    _ = fi8
                    _ = v9
                }
            }
        case "a":
            v5 = (v5 * 3.14159)
        }
        _ = v4
        _ = v5
    } else {
        v10 := []int{-2147483648, 2147483647}
        var v11 bool = true
        _ = v10
        _ = v11
    }
    var v12 uint = 255
    var v13 int16 = -32768
    _ = v12
    _ = v13
    return "hello"
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint8 = 1
    var v3 int8 = -128
    v14 := GenS1{F0: 0, F1: false}
    var v15 int16 = 0
    v1 = ((v1 - v1) - ((v1 * (v1 / v1)) + (v1 + v1)))
    var v16 bool = false
    v17 := GenS1{F0: 1, F1: true}
    for fi18 := 0; fi18 < 14; fi18++ {
        switch fi18 {
        case 1:
            fmt.Println(v16, v1)
        }
        if v16 {
            switch "hello" {
            case "hello":
                var v19 uint16 = 32767
                _ = v19
            default:
                switch "\n\t" {
                case "":
                    if (v16 && true) {
                        v16 = (!(0.0 < v1))
                        switch "" {
                        case "foo bar":
                            if v16 {
                                switch "\n\t" {
                                case "":
                                    var v20 uint32 = 2147483647
                                    _ = v20
                                case "a":
                                    var v21 uint = uint(v2)
                                    _ = v21
                                default:
                                    for fi22 := 0; fi22 < 19; fi22++ {
                                        v23 := GenS1{F0: 564, F1: (v3 > (v3 | -128))}
                                        v24 := GenS1{F0: 315, F1: v16}
                                        switch "foo bar" {
                                        case "\n\t":
                                            switch fi22 {
                                            case 2147483647:
                                                var v25 uint16 = 207
                                                _ = v25
                                            case -540:
                                                if (0 < 745) {
                                                    v26 := GenS1{F0: 954, F1: v16}
                                                    var v27 int64 = -9223372036854775808
                                                    _ = v26
                                                    _ = v27
                                                }
                                            case 674:
                                                var v28 float32 = -2.25
                                                _ = v28
                                            default:
                                                v29 := []int{-2147483648}
                                                _ = v29
                                            }
                                        case "hello":
                                            var v30 uint32 = 1
                                            _ = v30
                                        case "":
                                            switch (((((("\n\t" + ("hello" + ("foo bar" + "\n\t"))) + "foo bar") + ("foo bar" + (("hello" + ("foo bar" + "\n\t")) + ("a" + "")))) + ((("\n\t" + ("foo bar" + "")) + (("hello" + "\n\t") + "foo bar")) + "\n\t")) + "\n\t") + "foo bar") {
                                            case "hello":
                                                if ("\n\t" <= "") {
                                                    v31 := []int{}
                                                    _ = v31
                                                } else {
                                                    v2 += 157
                                                }
                                            default:
                                                v32 := []int{(fi18 + fi18)}
                                                _ = v32
                                            }
                                        }
                                        _ = fi22
                                        _ = v23
                                        _ = v24
                                    }
                                }
                                var v33 int64 = 42
                                _ = v33
                            }
                        case "hello":
                            switch "a" {
                            case "foo bar":
                                for fi34 := 0; fi34 < 3; fi34++ {
                                    switch fi34 {
                                    case -926:
                                        v3 = v3
                                    case 0:
                                        v3 += (v3 * (v3 * (^(v3))))
                                    case -9223372036854775808:
                                        var v35 int64 = -1
                                        _ = v35
                                    default:
                                        v15 = (v15 | 42)
                                    }
                                    fmt.Println(v2, v16, v2)
                                    fmt.Println(v1, v2)
                                    _ = fi34
                                }
                            case "":
                                v36 := GenS2{F0: (^(v2)), F1: -2147483648}
                                _ = v36
                            case "hello":
                                v3 = 73
                            }
                        case "":
                            v37 := GenS3{F0: 42}
                            _ = v37
                        }
                        for fi38 := 0; fi38 < 14; fi38++ {
                            fmt.Println(v15)
                            var v39 int = (fi18 + fi18)
                            _ = fi38
                            _ = v39
                        }
                    }
                case "foo bar":
                    v40 := GenS2{F0: v2, F1: (fi18 ^ 4294967295)}
                    _ = v40
                default:
                    v41 := GenS1{F0: 930, F1: v16}
                    _ = v41
                }
            }
        }
        fmt.Println(v1, fi18, v1)
        _ = fi18
    }
    for fi42 := 0; fi42 < 4; fi42++ {
        v43 := GenS1{F0: 42, F1: v16}
        v16 = true
        _ = fi42
        _ = v43
    }
    var v44 int = 32768
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v44
}
