package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8; F1 string }
type GenS7 struct { F0 int16 }
func genFn0(p0 uint16, p1 int, p2 uint32) bool {
    p2 = (p2 % (((p2 >> ((290) & 63))) | 1))
    var v6 string = "a"
    if (42 != 18446744073709551615) {
        v7 := GenS1{F0: -1, F1: -1}
        _ = v7
    } else {
        v6 = v6
        var v8 uint32 = p2
        _ = v8
    }
    switch v6 {
    case "hello":
        v9 := GenS1{F0: 76, F1: 28}
        _ = v9
    default:
        if (682 >= 2) {
            if (4294967295 <= 65535) {
                if (!true) {
                    v6 = ((v6 + (((v6 + v6) + v6) + (v6 + v6))) + ((v6 + v6) + "foo bar"))
                    v10 := []int{p1, (p1 % ((147) | 1)), 128}
                    _ = v10
                }
                var v11 uint32 = p2
                _ = v11
            } else {
                v12 := []int{p1, p1, (p1 / (((p1 % (((p1 ^ (p1 & -354))) | 1))) | 1))}
                _ = v12
            }
        } else {
            p2 += p2
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return true
}
func genFn1() int64 {
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int = fi13
        v15 := []int{-839, 42}
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint32 = 340
    switch -592 {
    case -2147483648:
        v17 := GenS2{F0: 178}
        _ = v17
    case -9223372036854775808:
        v16 += 42
    }
    _ = v16
    return 1
}

func main() {
    var v1 uint8 = 1
    v2 := []string{}
    v3 := make(map[string]bool)
    var v4 uint64 = 42
    var v5 float64 = 1.5
    v18 := GenS1{F0: 100, F1: -72}
    if true {
        fmt.Println(v4, v5)
        v4 -= 2
    } else {
        v19 := GenS3{F0: 127, F1: 255}
        _ = v19
    }
    fmt.Println(v5, v1)
    switch "\n\t" {
    case "foo bar":
        switch int(v4) {
        case -842:
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 string = "hello"
                _ = fi20
                _ = v21
            }
        }
    case "":
        _ = len(v3)
    case "\n\t":
        if (!(42 != 2)) {
            var v22 string = ("foo bar" + ((("a" + "foo bar") + "hello") + "foo bar"))
            if (!true) {
                _ = len(v2)
            }
            for fi23 := 0; fi23 < 3; fi23++ {
                v24 := GenS4{F0: -128}
                var v25 int16 = 0
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = v22
        }
    }
    switch ("a" + "foo bar") {
    case "hello":
        switch "hello" {
        case "foo bar":
            v5 -= (v5 + -2.25)
        case "a":
            switch 32767 {
            case 128:
                var v26 uint32 = 0
                _ = v26
            default:
                for fi27 := 0; fi27 < 16; fi27++ {
                    v3["hello"] = (((v1 < (v1 << ((4294967295) & 63))) && true) || (9223372036854775807 >= -780))
                    switch "" {
                    case "\n\t":
                        for fi28 := 0; fi28 < 17; fi28++ {
                            v29 := GenS5{F0: 659}
                            _ = fi28
                            _ = v29
                        }
                    }
                    var v30 float64 = 0.1
                    _ = fi27
                    _ = v30
                }
            }
        default:
            fmt.Println(v1, v4, v4)
        }
    case "\n\t":
        v5 += v5
    case "":
        v31 := GenS6{F0: -1, F1: ""}
        _ = v31
    }
    var v32 int8 = -128
    v4 = (^(v4))
    v5 = v5
    if ((v4 % ((v4) | 1)) > v4) {
        delete(v3, "hello")
        var v33 int = 9223372036854775807
        _ = v33
    } else {
        _ = len(v2)
        if len(v2) > 0 {
            v2[0] = ("hello" + "a")
        }
        var v34 int64 = -9223372036854775808
        _ = v34
    }
    if (float64(1e10) < float64(3.14159)) {
        for fi35 := 0; fi35 < 16; fi35++ {
            v1 -= 42
            v36 := GenS7{F0: 42}
            if (-1 == -32768) {
                fmt.Println(v5)
                fmt.Println(v1)
                v4 += 4294967295
            }
            _ = fi35
            _ = v36
        }
        var v37 int32 = 42
        v5 -= (v5 / 1e10)
        _ = v37
    } else {
        for fi38 := 0; fi38 < 2; fi38++ {
            fmt.Println(fi38, v1)
            v5 -= v5
            switch "\n\t" {
            case "a":
                switch "" {
                case "a":
                    switch ("\n\t" + "hello") {
                    case "a":
                        var v39 bool = false
                        _ = v39
                    case "hello":
                        for fi40 := 0; fi40 < 20; fi40++ {
                            fmt.Println(v1, v32)
                            var v41 uint8 = v1
                            fmt.Println(v5, fi40, v4)
                            _ = fi40
                            _ = v41
                        }
                    }
                case "hello":
                    if true {
                        v42 := GenS1{F0: (v32 - (v32 | (v32 << ((616) & 63)))), F1: 100}
                        _ = v42
                    }
                case "":
                    if (((((v32 >> ((237) & 63)) > (^(v32))) && false) || (879 > 1)) && true) {
                        fmt.Println(v5)
                        switch "a" {
                        case "foo bar":
                            fmt.Println(v5)
                        case "a":
                            v1 = 42
                        case "hello":
                            if false {
                                fmt.Println(v1)
                                if len(v2) > 0 {
                                    v2[0] = (("" + "hello") + "\n\t")
                                }
                                v1 -= (v1 ^ 1)
                            }
                        }
                    } else {
                        if false {
                            var v43 int8 = (v32 << ((4294967295) & 63))
                            v43 += (v43 % ((v32) | 1))
                            var v44 float64 = 1e10
                            _ = v43
                            _ = v44
                        } else {
                            v4 = v4
                        }
                        v45 := v3["foo bar"]
                        v46 := GenS1{F0: 0, F1: v32}
                        _ = v45
                        _ = v46
                    }
                default:
                    var v47 bool = (!(1 <= 336))
                    _ = v47
                }
            case "foo bar":
                delete(v3, (("a" + (("\n\t" + "foo bar") + "")) + "\n\t"))
            default:
                if ("" >= ("hello" + "hello")) {
                    v48 := GenS6{F0: 84, F1: ((("hello" + "") + "") + ("\n\t" + ("" + (("" + ("\n\t" + "a")) + ""))))}
                    _ = v48
                } else {
                    v5 += v5
                    var v49 uint64 = v4
                    var v50 int32 = 1
                    _ = v49
                    _ = v50
                }
            }
            _ = fi38
        }
        v3["foo bar"] = true
        v51 := GenS3{F0: (^(v32)), F1: (v1 + v1)}
        _ = v51
    }
    v52 := GenS1{F0: v32, F1: v32}
    v53 := GenS3{F0: (v32 * (v32 - v32)), F1: 127}
    v54 := GenS3{F0: (v32 >> ((2) & 63)), F1: 127}
    var v55 string = "a"
    var v56 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v56
}
