package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16; F1 uint16; F2 int64 }
type GenS8 struct { F0 int32; F1 uint32; F2 string }
func genFn0(p0 uint16) string {
    p0 = (p0 + (p0 ^ (p0 >> ((2) & 63))))
    var v6 uint = 2
    v7 := GenS1{F0: true, F1: 2, F2: -2147483648}
    _ = p0
    _ = v6
    _ = v7
    return ("foo bar" + "")
}
func genFn1() bool {
    var v8 float64 = 0.1
    v9 := GenS1{F0: (true || false), F1: 65535, F2: 2147483647}
    for fi10 := 0; fi10 < 14; fi10++ {
        v8 -= float64(1e10)
        if false {
            for fi11 := 0; fi11 < 9; fi11++ {
                fmt.Println(fi10)
                switch "\n\t" {
                case "\n\t":
                    var v12 int16 = -1
                    _ = v12
                case "hello":
                    if (18446744073709551615 > 1) {
                        v8 = v8
                        if true {
                            var v13 int32 = -2147483648
                            _ = v13
                        } else {
                            var v14 uint32 = 890
                            _ = v14
                        }
                        if (!true) {
                            var v15 uint64 = 1
                            _ = v15
                        } else {
                            if (!(50 == -28)) {
                                if (128 == 6) {
                                    v16 := []int{1, fi11, (^(fi10))}
                                    v17 := GenS2{F0: 0, F1: (fi10 << ((255) & 63))}
                                    v18 := GenS3{F0: 42, F1: 2, F2: 892}
                                    v16 = append(v16, fi10)
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                }
                                fmt.Println(fi11, fi11, fi11)
                            }
                            v19 := []int{fi11, fi11, (fi10 | (^(fi10)))}
                            _ = v19
                        }
                        var v20 float32 = 0.0
                        var v21 int = (^(fi11))
                        _ = v20
                        _ = v21
                    } else {
                        v8 += 3.14159
                        switch ("foo bar" + "hello") {
                        case "a":
                            fmt.Println(fi10)
                        default:
                            v8 -= 3.14159
                        }
                    }
                case "":
                    if (!true) {
                        v8 -= -2.25
                        v22 := GenS4{F0: 691}
                        _ = v22
                    }
                default:
                    v8 = ((v8 - v8) * (v8 + v8))
                }
                _ = fi11
            }
            var v23 int = fi10
            switch (v23 ^ (fi10 ^ (fi10 >> ((65535) & 63)))) {
            case 1:
                v24 := GenS3{F0: 42, F1: 91, F2: 0}
                _ = v24
            case 42:
                var v25 uint16 = 995
                _ = v25
            case 0:
                for fi26 := 0; fi26 < 4; fi26++ {
                    v27 := []int{}
                    if true {
                        v8 = v8
                        v8 = v8
                    }
                    var v28 float32 = -2.25
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
            _ = v23
        } else {
            var v29 bool = false
            _ = v29
        }
        switch (fi10 | -9223372036854775808) {
        case 4294967295:
            v8 = (v8 - ((v8 / (v8 * (v8 + v8))) - v8))
        case 332:
            v30 := GenS4{F0: 637}
            _ = v30
        }
        _ = fi10
    }
    _ = v8
    _ = v9
    return (!true)
}

func main() {
    var v1 bool = (!true)
    var v2 int16 = -32768
    v3 := make(map[string]bool)
    var v4 int32 = 0
    var v5 uint = 51
    if v1 {
        fmt.Println(v1, v4)
        fmt.Println(v2, v2)
        v5 = v5
    } else {
        switch ("\n\t" + ("foo bar" + "foo bar")) {
        case "foo bar":
            for fi31 := 0; fi31 < 7; fi31++ {
                var v32 uint32 = 1
                _ = fi31
                _ = v32
            }
        case "a":
            v2 += 0
        case "hello":
            v33 := GenS3{F0: 87, F1: 32767, F2: 32767}
            _ = v33
        }
        for fi34 := 0; fi34 < 16; fi34++ {
            fmt.Println(fi34)
            _ = fi34
        }
        if (1 != 0) {
            if (128 >= -1) {
                var v35 int32 = v4
                if v1 {
                    v36 := GenS5{F0: 42}
                    var v37 float32 = (3.14159 / -2.25)
                    fmt.Println(v4)
                    _ = v36
                    _ = v37
                }
                _ = v35
            } else {
                v38 := GenS1{F0: false, F1: 1, F2: (v4 + (v4 << (((^(v5))) & 63)))}
                _ = v38
            }
            if (-2147483648 >= v4) {
                if (v1 || true) {
                    v39 := []int{-9223372036854775808, 32767}
                    _ = v39
                }
            } else {
                _ = len(v3)
                var v40 int32 = (v4 / (((^(v4))) | 1))
                _ = v40
            }
        } else {
            switch -400 {
            case 0:
                var v41 uint = v5
                _ = v41
            case 32767:
                for fi42 := 0; fi42 < 12; fi42++ {
                    if v1 {
                        v43 := GenS6{F0: -9223372036854775808}
                        fmt.Println(v1, v4)
                        var v44 string = ("" + "")
                        _ = v43
                        _ = v44
                    } else {
                        for fi45 := 0; fi45 < 4; fi45++ {
                            var v46 int = fi45
                            v3["\n\t"] = v1
                            if (!(-2.25 > (-2.25 * (1e10 - 0.0)))) {
                                v2 = 201
                            }
                            _ = fi45
                            _ = v46
                        }
                        if (v4 >= v4) {
                            v47 := []int{fi42}
                            _ = v47
                        } else {
                            v48 := GenS7{F0: 42, F1: 0, F2: -1}
                            switch "a" {
                            case "foo bar":
                                v2 = v2
                            case "hello":
                                switch (("a" + ("hello" + ((("foo bar" + "foo bar") + "\n\t") + ("" + "foo bar")))) + ("hello" + "a")) {
                                case "":
                                    v2 = v2
                                case "\n\t":
                                    if (65535 < 2147483647) {
                                        v4 += 2147483647
                                        fmt.Println(fi42, v5, v4)
                                        fmt.Println(v2, v2, v5)
                                    } else {
                                        var v49 int32 = (v4 ^ -1)
                                        v1 = true
                                        v50 := v3[("" + ("\n\t" + ("a" + "\n\t")))]
                                        _ = v49
                                        _ = v50
                                    }
                                }
                            case "":
                                var v51 float64 = 3.14159
                                _ = v51
                            default:
                                fmt.Println(v4, v4, v5)
                            }
                            _ = v48
                        }
                    }
                    _ = fi42
                }
            case -90:
                v52 := []int{2147483647, 2147483647, -106, 32767}
                _ = v52
            }
        }
    }
    for fi53 := 0; fi53 < 2; fi53++ {
        for fi54 := 0; fi54 < 8; fi54++ {
            v55 := GenS8{F0: 42, F1: 4294967295, F2: ""}
            delete(v3, ("hello" + ("foo bar" + "hello")))
            v5 = (v5 / ((65535) | 1))
            _ = fi54
            _ = v55
        }
        if (!v1) {
            fmt.Println(v2, v4, v1)
        } else {
            v1 = true
        }
        _ = fi53
    }
    var v56 float32 = 1e10
    v57 := GenS8{F0: v4, F1: 2147483648, F2: ("hello" + (("foo bar" + "") + "foo bar"))}
    switch 32767 {
    case -544:
        v58 := []int{128, 590}
        _ = v58
    case 109:
        fmt.Println(v2, v4)
    case -1:
        for fi59 := 0; fi59 < 4; fi59++ {
            var v60 int8 = -109
            var v61 uint32 = 259
            _ = fi59
            _ = v60
            _ = v61
        }
    }
    var v62 int32 = -371
    v2 += v2
    var v63 uint64 = 4294967295
    v2 = (v2 | v2)
    var v64 string = ("" + ((("a" + "hello") + "a") + (("a" + ("hello" + "a")) + "hello")))
    fmt.Println(v56, v64)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v62
    _ = v63
    _ = v64
}
