package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint64; F1 int16; F2 int8 }
type GenS3 struct { F0 uint; F1 int; F2 uint8 }
func genFn0(p0 int, p1 int, p2 uint64) uint32 {
    v4 := []int{}
    v5 := GenS1{F0: 1, F1: 313}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := []string{}
    v2 := []string{("foo bar" + "a"), "hello"}
    v3 := make(map[string]uint32)
    var v6 int8 = 100
    var v7 uint = 0
    v7 += 42
    v8 := GenS1{F0: 27, F1: 32767}
    if false {
        fmt.Println(v6)
    }
    v6 = v6
    v9 := GenS1{F0: 2, F1: 42}
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            if ("hello" != ("hello" + "\n\t")) {
                v6 += (v6 | (v6 + (v6 >> ((v7) & 63))))
                v7 = v7
                switch (fi10 % ((fi10) | 1)) {
                case 127:
                    v7 = 2
                case 9223372036854775807:
                    fmt.Println(fi10, fi11)
                }
                switch "\n\t" {
                case "hello":
                    v7 = (v7 * 296)
                case "\n\t":
                    var v12 bool = true
                    v6 -= v6
                    _ = v12
                case "":
                    v7 = v7
                    v13 := GenS1{F0: 88, F1: 0}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v13
                }
            } else {
                v14 := GenS1{F0: 1, F1: 604}
                _ = v14
            }
            v7 += v7
            _ = fi11
        }
        fmt.Println(fi10)
        switch "a" {
        case "hello":
            if false {
                v15 := GenS1{F0: 246, F1: -32768}
                var v16 int16 = -399
                _ = len(v3)
                _ = v15
                _ = v16
            } else {
                _ = len(v3)
                delete(v3, ("foo bar" + ("hello" + "a")))
            }
        case "\n\t":
            var v17 int64 = 1
            _ = v17
        case "":
            if (fi10 > fi10) {
                var v18 uint64 = 1
                _ = v18
            } else {
                var v19 int64 = 1
                v7 = (^(v7))
                fmt.Println(v6)
                _ = v19
            }
        default:
            _ = len(v2)
        }
        _ = fi10
    }
    var v20 int64 = -290
    if (1e10 < 3.14159) {
        fmt.Println(v7)
    } else {
        if false {
            var v21 uint = (v7 ^ 65535)
            for fi22 := 0; fi22 < 3; fi22++ {
                v21 += (v7 & 195)
                _ = fi22
            }
            _ = v21
        } else {
            switch ("hello" + "") {
            case "":
                v20 -= v20
            case "foo bar":
                switch 1 {
                case -776:
                    switch ("\n\t" + "hello") {
                    case "foo bar":
                        if false {
                            v20 -= (v20 << ((2) & 63))
                        } else {
                            v3[(("" + ("" + "a")) + "\n\t")] = 2147483647
                            var v23 uint32 = 2147483647
                            _ = v23
                        }
                    case "hello":
                        for fi24 := 0; fi24 < 14; fi24++ {
                            v6 -= v6
                            switch ("" + "") {
                            case "":
                                v25 := GenS1{F0: 1, F1: -1}
                                _ = v25
                            case "\n\t":
                                if false {
                                    if len(v1) > 0 {
                                        v1[0] = ""
                                    }
                                    fmt.Println(v6)
                                    switch "" {
                                    case "\n\t":
                                        v1 = append(v1, ("a" + "\n\t"))
                                    }
                                }
                            case "hello":
                                delete(v3, "foo bar")
                            }
                            _ = fi24
                        }
                    }
                case -9223372036854775808:
                    for fi26 := 0; fi26 < 20; fi26++ {
                        if len(v2) > 0 {
                            v2[0] = "\n\t"
                        }
                        _ = fi26
                    }
                case -781:
                    var v27 uint16 = 0
                    _ = v27
                }
            case "a":
                if (391 > 4294967295) {
                    for fi28 := 0; fi28 < 2; fi28++ {
                        if false {
                            if (false && ((v7 * v7) != 0)) {
                                v29 := GenS2{F0: 259, F1: 150, F2: v6}
                                _ = v29
                            } else {
                                var v30 uint8 = 127
                                var v31 uint16 = 255
                                for fi32 := 0; fi32 < 18; fi32++ {
                                    if false {
                                        var v33 float64 = -2.25
                                        _ = v33
                                    }
                                    fmt.Println(v6)
                                    v34 := v3["\n\t"]
                                    _ = fi32
                                    _ = v34
                                }
                                _ = v30
                                _ = v31
                            }
                            for fi35 := 0; fi35 < 16; fi35++ {
                                v6 = v6
                                v7 = (^(v7))
                                _ = fi35
                            }
                            for fi36 := 0; fi36 < 3; fi36++ {
                                if (v6 > (v6 * (v6 - v6))) {
                                    switch int(v7) {
                                    case 71:
                                        fmt.Println(v20, v6)
                                    }
                                    var v37 int64 = v20
                                    var v38 uint32 = genFn0((fi36 / ((4294967295) | 1)), (fi36 >> (((v7 - (v7 << ((v7) & 63)))) & 63)), 1)
                                    _ = v37
                                    _ = v38
                                } else {
                                    v6 = (v6 ^ v6)
                                    v39 := GenS3{F0: (v7 - 42), F1: (^(fi28)), F2: 0}
                                    _ = v39
                                }
                                _ = fi36
                            }
                        } else {
                            var v40 float64 = 0.1
                            _ = v40
                        }
                        _ = fi28
                    }
                } else {
                    v3["hello"] = 185
                    for fi41 := 0; fi41 < 9; fi41++ {
                        v6 = (v6 | (v6 / ((v6) | 1)))
                        var v42 uint16 = 42
                        _ = fi41
                        _ = v42
                    }
                }
            }
            delete(v3, "hello")
        }
        v7 = v7
    }
    v7 -= (v7 | v7)
    v43 := GenS1{F0: 42, F1: 567}
    v7 -= v7
    switch 512 {
    case 9223372036854775807:
        v20 = (v20 / ((-9223372036854775808) | 1))
    case 4294967295:
        if (v7 < (v7 & v7)) {
            switch 4294967295 {
            case 4294967295:
                var v44 int32 = 0
                _ = v44
            case 42:
                var v45 uint8 = 2
                _ = v45
            default:
                for fi46 := 0; fi46 < 6; fi46++ {
                    for fi47 := 0; fi47 < 3; fi47++ {
                        v48 := GenS1{F0: 0, F1: 1000}
                        if false {
                            var v49 int = -553
                            v50 := GenS1{F0: 255, F1: -32768}
                            _ = v49
                            _ = v50
                        }
                        v51 := GenS2{F0: 0, F1: -659, F2: v6}
                        _ = fi47
                        _ = v48
                        _ = v51
                    }
                    switch fi46 {
                    case 374:
                        if false {
                            switch ("\n\t" + "a") {
                            case "hello":
                                var v52 int64 = (v20 ^ -323)
                                _ = v52
                            case "":
                                fmt.Println(fi46, v20)
                            default:
                                for fi53 := 0; fi53 < 9; fi53++ {
                                    if true {
                                        v54 := GenS3{F0: 0, F1: 4294967295, F2: 255}
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v55 bool = false
                                        _ = v54
                                        _ = v55
                                    }
                                    v56 := GenS1{F0: 255, F1: -32768}
                                    var v57 bool = false
                                    fmt.Println(fi53)
                                    _ = fi53
                                    _ = v56
                                    _ = v57
                                }
                            }
                            if true {
                                switch "hello" {
                                case "":
                                    if true {
                                        fmt.Println(v20)
                                        fmt.Println(v6, fi46, fi46)
                                    }
                                case "hello":
                                    fmt.Println(fi46, fi46)
                                case "foo bar":
                                    fmt.Println(v20, v20, v6)
                                default:
                                    v3[(("foo bar" + ("a" + "hello")) + "hello")] = 0
                                }
                                var v58 bool = true
                                if true {
                                    var v59 bool = v58
                                    var v60 int16 = 577
                                    _ = v59
                                    _ = v60
                                }
                                _ = v58
                            } else {
                                fmt.Println(v7)
                                var v61 float32 = 0.0
                                var v62 float32 = (3.14159 / v61)
                                _ = v61
                                _ = v62
                            }
                            v20 = -9223372036854775808
                        }
                    case -2147483648:
                        v2 = append(v2, ("" + ""))
                    case 32768:
                        if false {
                            v63 := GenS2{F0: 427, F1: 0, F2: (v6 << (((v7 * 131)) & 63))}
                            v64 := GenS2{F0: 684, F1: -1, F2: int8(v20)}
                            _ = v63
                            _ = v64
                        } else {
                            v6 = -80
                        }
                    default:
                        switch 2147483647 {
                        case 32768:
                            v7 = (v7 - 255)
                        case 953:
                            if ((("\n\t" + "\n\t") + "\n\t") >= "foo bar") {
                                fmt.Println(v6, v20, v20)
                            } else {
                                fmt.Println(fi46, v6)
                                var v65 int32 = int32(v7)
                                _ = v65
                            }
                        default:
                            v3[""] = 2
                        }
                    }
                    _ = fi46
                }
            }
            v20 += v20
        } else {
            fmt.Println(v20)
            fmt.Println(v20, v7, v7)
            _ = len(v1)
        }
    default:
        for fi66 := 0; fi66 < 9; fi66++ {
            fmt.Println(v20, v20)
            v7 -= (^(v7))
            for fi67 := 0; fi67 < 9; fi67++ {
                v20 -= v20
                _ = fi67
            }
            _ = fi66
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v43
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
