package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint16 }
type GenS2 struct { F0 int32; F1 string; F2 int8 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 int16; F1 uint16; F2 string }
func genFn0(p0 uint64) uint32 {
    p0 -= 4294967295
    fmt.Println(p0)
    var v7 uint64 = p0
    _ = p0
    _ = v7
    return 65535
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{4294967295, 470, 2147483647}
    var v3 int32 = -819
    v4 := make(map[int]bool)
    var v5 uint8 = 92
    v6 := make(map[string]int)
    if (true || false) {
        for fi8 := 0; fi8 < 17; fi8++ {
            if (!(127 < 0)) {
                if false {
                    for fi9 := 0; fi9 < 6; fi9++ {
                        fmt.Println(fi9, fi9, v5)
                        v5 -= v5
                        _ = fi9
                    }
                    var v10 float64 = 1e10
                    var v11 uint8 = v5
                    _ = v10
                    _ = v11
                }
                fmt.Println(v5, v5)
                if false {
                    v3 += v3
                    v5 = (v5 + 127)
                    var v12 bool = ((float64(-2.25) != 3.14159) || (((!(32767 >= -48)) && false) || false))
                    _ = v12
                }
            } else {
                fmt.Println(v5, v5, v3)
                v13 := GenS1{F0: "a", F1: fi8, F2: 65535}
                _ = v13
            }
            _ = fi8
        }
    }
    var v14 int32 = (v3 << ((4294967295) & 63))
    if false {
        if true {
            switch 42 {
            case 1:
                delete(v1, int(v3))
            case 128:
                var v15 uint16 = 42
                _ = v15
            case -1:
                v16 := GenS2{F0: v14, F1: "\n\t", F2: -21}
                _ = v16
            }
            v14 = v14
            v5 = (^(v5))
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v17 := GenS3{F0: 410, F1: false}
            _ = v17
        }
        v5 = v5
    } else {
        var v18 string = "\n\t"
        _ = v18
    }
    if true {
        if false {
            v4[-2147483648] = (-914 >= 42)
            _ = len(v1)
        } else {
            v19 := GenS4{F0: (491 <= -573), F1: 1}
            var v20 string = ""
            for fi21 := 0; fi21 < 15; fi21++ {
                v20 = v20
                var v22 bool = (genFn0(18446744073709551615) < 2147483648)
                v23 := GenS1{F0: "a", F1: fi21, F2: 32767}
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = v19
            _ = v20
        }
        fmt.Println(v3, v3, v14)
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        if true {
            v25 := GenS3{F0: 32767, F1: (uint64(v5) == 0)}
            if (1 >= 32767) {
                v3 = (-(v3))
            } else {
                v26 := GenS5{F0: 32767, F1: 1, F2: ""}
                v5 = (v5 & 211)
                var v27 uint = 4294967295
                _ = v26
                _ = v27
            }
            _ = v25
        } else {
            v14 = int32(v5)
            var v28 int16 = -1
            _ = v28
        }
        var v29 string = ("\n\t" + "hello")
        _ = fi24
        _ = v29
    }
    var v30 uint64 = 42
    for fi31 := 0; fi31 < 5; fi31++ {
        if (68 != 118) {
            fmt.Println(v3, fi31, v14)
            v3 -= 674
        }
        _ = fi31
    }
    if true {
        for fi32 := 0; fi32 < 1; fi32++ {
            v33 := GenS1{F0: "foo bar", F1: (fi32 & fi32), F2: 725}
            var v34 uint64 = (v30 ^ 0)
            var v35 uint16 = 65535
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
    }
    var v36 bool = (("hello" + "hello") < "foo bar")
    v36 = (!true)
    fmt.Println(v30, v14)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "foo bar" {
    case "\n\t":
        v36 = ((false && (v36 || ((((-2.25 / -2.25) / -2.25) + 1.5) > 3.14159))) || ((v30 / ((v30) | 1)) < v30))
    case "foo bar":
        if true {
            for fi37 := 0; fi37 < 12; fi37++ {
                if ((v5 / (((^(v5))) | 1)) <= 2) {
                    var v38 float64 = float64(1e10)
                    _ = len(v4)
                    fmt.Println(v38, v38, v5)
                    _ = v38
                }
                switch (fi37 / ((fi37) | 1)) {
                case 128:
                    for fi39 := 0; fi39 < 16; fi39++ {
                        v40 := GenS2{F0: 146, F1: "hello", F2: 23}
                        v4[-738] = true
                        _ = fi39
                        _ = v40
                    }
                default:
                    v30 = v30
                }
                switch "\n\t" {
                case "\n\t":
                    v3 = (v3 & 0)
                case "":
                    v36 = ((!(false && false)) || v36)
                }
                _ = fi37
            }
            v6[("a" + ("foo bar" + ("" + ("a" + ""))))] = 9223372036854775807
        }
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        for fi41 := 0; fi41 < 6; fi41++ {
            for fi42 := 0; fi42 < 4; fi42++ {
                switch "a" {
                case "":
                    switch (("hello" + (("hello" + "a") + ("" + ("foo bar" + ("" + "\n\t"))))) + (((("a" + ("foo bar" + "foo bar")) + (("\n\t" + ((("hello" + ("hello" + "\n\t")) + "hello") + "")) + "a")) + "foo bar") + ("a" + "hello"))) {
                    case "\n\t":
                        var v43 int64 = 0
                        _ = v43
                    default:
                        switch ("\n\t" + (("hello" + "hello") + "a")) {
                        case "a":
                            switch 128 {
                            case 128:
                                if true {
                                    fmt.Println(v36, v14, fi42)
                                    if (!(v36 && true)) {
                                        if v36 {
                                            fmt.Println(v30)
                                            var v44 int8 = 124
                                            _ = v44
                                        }
                                        v45 := GenS4{F0: v36, F1: 411}
                                        var v46 int = 346
                                        _ = v45
                                        _ = v46
                                    }
                                }
                            case -2147483648:
                                v30 = (^(v30))
                            }
                        case "foo bar":
                            switch "a" {
                            case "":
                                if v36 {
                                    var v47 uint32 = 42
                                    _ = v47
                                }
                            }
                        case "":
                            v30 = v30
                        default:
                            switch "\n\t" {
                            case "\n\t":
                                switch ("foo bar" + "foo bar") {
                                case "a":
                                    v36 = v36
                                case "foo bar":
                                    v30 = (^(v30))
                                default:
                                    var v48 string = "hello"
                                    _ = v48
                                }
                            default:
                                var v49 int8 = -10
                                _ = v49
                            }
                        }
                    }
                default:
                    v50 := v6["a"]
                    _ = v50
                }
                _ = fi42
            }
            v14 -= (^(v14))
            _ = fi41
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
