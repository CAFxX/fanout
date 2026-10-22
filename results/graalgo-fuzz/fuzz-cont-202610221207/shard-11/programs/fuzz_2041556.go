package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint64; F2 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 string; F2 uint }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 string; F1 uint8 }
type GenS10 struct { F0 int64; F1 string }
type GenS11 struct { F0 bool }
func genFn0(p0 int16, p1 uint64) string {
    var v5 int8 = -54
    p1 = p1
    _ = p0
    _ = p1
    _ = v5
    return "hello"
}
func genFn1(p0 bool) int64 {
    var v6 int16 = -32768
    v7 := GenS1{F0: 0, F1: -972, F2: 42}
    v8 := []int{32768}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int64 = 1
    var v2 uint = 2
    v3 := []int{}
    var v4 uint16 = 32767
    for fi9 := 0; fi9 < 19; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 bool = (true || false)
            _ = fi10
            _ = v11
        }
        var v12 int64 = 0
        var v13 uint64 = 4294967295
        _ = fi9
        _ = v12
        _ = v13
    }
    switch 1 {
    case 536:
        v3 = append(v3, 127)
    case 32768:
        var v14 bool = true
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 bool = true
        _ = len(v3)
        if (1 < 42) {
            _ = len(v3)
            var v17 uint = (v2 - 255)
            _ = v17
        } else {
            switch (fi15 & fi15) {
            case -151:
                v18 := GenS1{F0: 8, F1: -9223372036854775808, F2: 128}
                _ = v18
            case 128:
                switch "" {
                case "\n\t":
                    v3 = append(v3, 9223372036854775807)
                default:
                    switch fi15 {
                    case 563:
                        if false {
                            v19 := GenS2{F0: "\n\t"}
                            if (!v16) {
                                var v20 int = fi15
                                for fi21 := 0; fi21 < 4; fi21++ {
                                    _ = len(v3)
                                    if ((v16 || (v2 < 255)) && false) {
                                        if ((v1 * v1) >= v1) {
                                            v22 := GenS3{F0: 1000, F1: 42, F2: genFn1(false)}
                                            v23 := GenS2{F0: ("hello" + "a")}
                                            switch v20 {
                                            case 313:
                                                v16 = false
                                            default:
                                                fmt.Println(v2, v4)
                                            }
                                            v24 := GenS4{F0: 1}
                                            fmt.Println(v4)
                                            v4 += v4
                                            _ = v22
                                            _ = v23
                                            _ = v24
                                        }
                                    } else {
                                        v2 = (^(v2))
                                        v25 := GenS2{F0: "hello"}
                                        _ = v25
                                    }
                                    _ = fi21
                                }
                                v26 := GenS2{F0: ("hello" + (("" + "a") + ("foo bar" + "hello")))}
                                _ = v20
                                _ = v26
                            } else {
                                var v27 bool = v16
                                if v27 {
                                    fmt.Println(v16)
                                    for fi28 := 0; fi28 < 16; fi28++ {
                                        v4 = (v4 | 934)
                                        v16 = v27
                                        v29 := GenS2{F0: ("foo bar" + "\n\t")}
                                        _ = fi28
                                        _ = v29
                                    }
                                } else {
                                    var v30 int64 = 1
                                    _ = v30
                                }
                                _ = v27
                            }
                            _ = v19
                        } else {
                            if v16 {
                                v2 -= (v2 & v2)
                            } else {
                                var v31 uint16 = v4
                                _ = v31
                            }
                            v3 = append(v3, (fi15 << ((v2) & 63)))
                            fmt.Println(v2, v16)
                        }
                    case 1:
                        fmt.Println(v1, v4, v4)
                    default:
                        v3 = append(v3, fi15)
                    }
                }
            case 424:
                switch (("" + "") + "hello") {
                case "foo bar":
                    if true {
                        fmt.Println(v1, v4)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v1 -= (v1 | v1)
                    } else {
                        switch fi15 {
                        case 2147483647:
                            v3 = append(v3, fi15)
                        case -2147483648:
                            fmt.Println(v4)
                        }
                    }
                case "\n\t":
                    for fi32 := 0; fi32 < 2; fi32++ {
                        if v16 {
                            var v33 uint8 = 27
                            fmt.Println(fi15, v2)
                            _ = v33
                        } else {
                            v1 = -749
                        }
                        switch 42 {
                        case 9223372036854775807:
                            v4 = 255
                        case 2147483647:
                            v2 = 18446744073709551615
                            v16 = v16
                        case 128:
                            v34 := GenS4{F0: -2147483648}
                            _ = v34
                        }
                        _ = fi32
                    }
                }
            }
            v3 = append(v3, fi15)
        }
        _ = fi15
        _ = v16
    }
    switch (((((("" + "") + ("" + (("\n\t" + "a") + "\n\t"))) + "hello") + "\n\t") + "") + ("foo bar" + ("hello" + ("" + "a")))) {
    case "\n\t":
        fmt.Println(v2, v2, v2)
    case "":
        var v35 bool = (true && false)
        _ = v35
    case "hello":
        fmt.Println(v4)
    default:
        if true {
            var v36 float64 = (1e10 * -2.25)
            _ = len(v3)
            _ = v36
        }
    }
    v37 := GenS5{F0: 1000}
    if (false && (-698 != -2147483648)) {
        if (584 == -823) {
            var v38 uint16 = v4
            fmt.Println(v4)
            _ = v38
        }
        var v39 float32 = 0.0
        var v40 int32 = 778
        _ = v39
        _ = v40
    }
    fmt.Println(v1, v2, v1)
    var v41 uint32 = 0
    if len(v3) > 0 {
        v3[0] = -1
    }
    v2 = 1
    for fi42 := 0; fi42 < 4; fi42++ {
        for fi43 := 0; fi43 < 11; fi43++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi43
        }
        _ = fi42
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v41)
    v2 += v2
    if true {
        fmt.Println(v41, v4, v2)
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            switch "\n\t" {
            case "":
                _ = len(v3)
            case "a":
                v44 := GenS6{F0: (v41 | (v41 >> ((1) & 63))), F1: "foo bar", F2: v2}
                _ = v44
            }
        case "\n\t":
            var v45 uint8 = 30
            _ = v45
        }
        for fi46 := 0; fi46 < 16; fi46++ {
            v47 := GenS3{F0: -123, F1: uint64(v4), F2: (v1 >> ((1) & 63))}
            v41 += v41
            if false {
                v48 := GenS7{F0: 65535}
                v49 := GenS8{F0: -128}
                _ = v48
                _ = v49
            }
            _ = fi46
            _ = v47
        }
    }
    switch 32767 {
    case 127:
        switch "hello" {
        case "foo bar":
            v2 -= (v2 ^ v2)
        case "\n\t":
            v4 = uint16(v41)
        default:
            var v50 uint8 = 128
            _ = v50
        }
    case -2147483648:
        if (127 >= 1) {
            v51 := GenS1{F0: 1, F1: genFn1((-2.25 >= (0.0 / -2.25))), F2: uint8(v2)}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v1 -= -9223372036854775808
            _ = v51
        } else {
            var v52 int64 = v1
            if false {
                switch -223 {
                case 646:
                    v1 = -1
                case 128:
                    if false {
                        switch 127 {
                        case -2147483648:
                            v41 -= (v41 >> ((v2) & 63))
                        case 42:
                            if (false || true) {
                                if (true && (!true)) {
                                    v53 := GenS9{F0: "", F1: uint8(v2)}
                                    _ = v53
                                } else {
                                    var v54 uint8 = 42
                                    _ = v54
                                }
                            }
                        }
                        var v55 int = -851
                        v3 = append(v3, v55)
                        _ = v55
                    }
                }
                var v56 int8 = 1
                v57 := GenS10{F0: 42, F1: "a"}
                _ = v56
                _ = v57
            }
            _ = v52
        }
    case 113:
        _ = len(v3)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch ("foo bar" + "\n\t") {
    case "a":
        v58 := GenS11{F0: ((v1 & (v1 - genFn1((!true)))) <= -657)}
        _ = v58
    case "hello":
        for fi59 := 0; fi59 < 5; fi59++ {
            fmt.Println(v1, v4)
            if false {
                var v60 int32 = -1
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v61 float64 = 3.14159
                _ = v60
                _ = v61
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            v62 := GenS10{F0: v1, F1: "hello"}
            _ = fi59
            _ = v62
        }
    case "\n\t":
        for fi63 := 0; fi63 < 15; fi63++ {
            v64 := GenS6{F0: v41, F1: "a", F2: v2}
            _ = fi63
            _ = v64
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
}
