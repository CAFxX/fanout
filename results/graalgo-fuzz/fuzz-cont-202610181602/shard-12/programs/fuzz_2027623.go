package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8; F1 bool; F2 bool }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int16; F1 int8 }
func genFn0(p0 uint32, p1 uint16) int64 {
    for fi4 := 0; fi4 < 4; fi4++ {
        p1 += (^(p1))
        for fi5 := 0; fi5 < 12; fi5++ {
            fmt.Println(fi5, fi5)
            if (!(1e10 < ((0.0 + (1.5 * 0.1)) * 3.14159))) {
                fmt.Println(p1)
                v6 := GenS1{F0: "a", F1: 553}
                _ = v6
            }
            if false {
                if (true || (!false)) {
                    p0 = p0
                    p0 -= 513
                    p0 = 42
                }
                switch "a" {
                case "hello":
                    v7 := GenS2{F0: 100, F1: (p1 >> ((65535) & 63))}
                    _ = v7
                case "foo bar":
                    var v8 uint8 = 1
                    _ = v8
                default:
                    if (!true) {
                        v9 := GenS2{F0: 100, F1: p1}
                        var v10 bool = (9223372036854775807 > 42)
                        _ = v10
                        _ = v9
                    }
                }
                fmt.Println(p0)
            }
            _ = fi5
        }
        _ = fi4
    }
    var v11 uint64 = 815
    _ = p0
    _ = p1
    _ = v11
    return -151
}
func genFn1(p0 int, p1 int16, p2 int32) bool {
    if (!(true && (!(!(127 > 1))))) {
        switch 9223372036854775807 {
        case 42:
            if (!true) {
                switch (("hello" + "") + ("" + ("\n\t" + (("a" + "hello") + "foo bar")))) {
                case "\n\t":
                    var v12 float64 = 3.14159
                    _ = v12
                case "hello":
                    if false {
                        var v13 int16 = (p1 << ((4294967295) & 63))
                        v14 := []int{p0, 32768}
                        var v15 int16 = (p1 / ((v13) | 1))
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                default:
                    for fi16 := 0; fi16 < 1; fi16++ {
                        for fi17 := 0; fi17 < 9; fi17++ {
                            if (65535 != 1) {
                                switch p0 {
                                case 2147483647:
                                    var v18 bool = false
                                    _ = v18
                                default:
                                    v19 := GenS2{F0: -61, F1: 2}
                                    _ = v19
                                }
                                fmt.Println(fi16)
                            }
                            _ = fi17
                        }
                        switch ("a" + (("\n\t" + "foo bar") + "\n\t")) {
                        case "hello":
                            p2 = (p2 * 2147483647)
                        default:
                            v20 := GenS3{F0: true}
                            _ = v20
                        }
                        switch "foo bar" {
                        case "hello":
                            if false {
                                p1 -= p1
                            } else {
                                v21 := GenS4{F0: 19}
                                _ = v21
                            }
                        default:
                            p2 = p2
                        }
                        _ = fi16
                    }
                }
                var v22 int64 = -711
                if (int64(p0) != (v22 % (((v22 / (((v22 ^ genFn0(1, 507))) | 1))) | 1))) {
                    fmt.Println(p2, p1)
                    v23 := []int{}
                    var v24 uint8 = 186
                    _ = v23
                    _ = v24
                } else {
                    p2 += 42
                }
                _ = v22
            }
        case -1:
            fmt.Println(p0, p0)
        case 9223372036854775807:
            v25 := GenS5{F0: 0, F1: (817 == 0), F2: true}
            _ = v25
        }
        v26 := GenS6{F0: p2}
        _ = v26
    } else {
        for fi27 := 0; fi27 < 17; fi27++ {
            v28 := GenS7{F0: (p0 << ((4294967295) & 63))}
            v29 := []int{(p0 >> ((4294967295) & 63))}
            if false {
                v30 := GenS3{F0: false}
                var v31 int32 = p2
                _ = v30
                _ = v31
            } else {
                for fi32 := 0; fi32 < 9; fi32++ {
                    var v33 int = (^(fi27))
                    p1 += p1
                    if len(v29) >= 2 {
                        v29 = v29[1:len(v29)]
                    }
                    _ = fi32
                    _ = v33
                }
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        v34 := []int{}
        _ = v34
    }
    var v35 uint8 = 127
    v36 := []int{709, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v35
    _ = v36
    return false
}

func main() {
    var v1 uint = 115
    v2 := []string{"", ("hello" + ("hello" + ""))}
    v3 := []int{-51, 32768}
    var v37 int16 = 0
    for fi38 := 0; fi38 < 4; fi38++ {
        var v39 int8 = 0
        v40 := GenS7{F0: (fi38 % ((32768) | 1))}
        _ = fi38
        _ = v39
        _ = v40
    }
    var v41 float64 = -2.25
    if len(v2) > 0 {
        v2[0] = ("a" + "foo bar")
    }
    for fi42 := 0; fi42 < 10; fi42++ {
        var v43 uint32 = 42
        if (false && (255 > 127)) {
            v37 += v37
            v1 -= (v1 << (((v1 % ((1) | 1))) & 63))
        }
        _ = fi42
        _ = v43
    }
    var v44 uint8 = 2
    for fi45 := 0; fi45 < 12; fi45++ {
        var v46 int16 = (v37 | (v37 << ((v1) & 63)))
        v47 := GenS8{F0: (v46 >> ((561) & 63)), F1: 100}
        var v48 int8 = 127
        _ = fi45
        _ = v46
        _ = v47
        _ = v48
    }
    var v49 float32 = 1e10
    v44 -= v44
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v41
    _ = v44
    _ = v49
}
