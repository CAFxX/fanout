package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 string; F1 int }
func genFn0(p0 int64, p1 uint32) int8 {
    var v6 int = 32768
    p1 += p1
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1() uint {
    var v7 uint8 = 0
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: (v7 << ((0) & 63)), F1: 941}
        v7 -= 42
        _ = fi8
        _ = v9
    }
    if (255 > 32767) {
        if (v7 < 36) {
            switch ("\n\t" + (((("a" + "a") + "a") + "\n\t") + "a")) {
            case "":
                if false {
                    switch 0 {
                    case 32767:
                        var v10 int16 = 0
                        _ = v10
                    case 42:
                        v7 -= 255
                    default:
                        if (2147483647 > 4294967295) {
                            var v11 int16 = -32768
                            v12 := []int{792}
                            var v13 float32 = -2.25
                            _ = v11
                            _ = v12
                            _ = v13
                        } else {
                            for fi14 := 0; fi14 < 15; fi14++ {
                                v7 -= (v7 << ((524) & 63))
                                _ = fi14
                            }
                        }
                    }
                } else {
                    v15 := []int{4294967295, 42}
                    v7 = 255
                    _ = v15
                }
            default:
                var v16 uint16 = 0
                _ = v16
            }
        } else {
            if ("\n\t" > "") {
                for fi17 := 0; fi17 < 5; fi17++ {
                    v7 = v7
                    _ = fi17
                }
                var v18 uint = 1
                _ = v18
            }
            v7 = 128
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            v7 += (v7 | v7)
            _ = fi19
        }
        v20 := []int{9223372036854775807, -2147483648, 127, 42}
        _ = v20
    } else {
        switch "\n\t" {
        case "":
            if false {
                var v21 uint8 = 83
                v22 := GenS2{F0: 911}
                v21 = 42
                _ = v21
                _ = v22
            }
        case "a":
            var v23 float64 = (0.0 + 0.0)
            _ = v23
        case "\n\t":
            if true {
                v7 = v7
                v24 := GenS3{F0: 42, F1: 42, F2: -32768}
                _ = v24
            } else {
                var v25 int = -82
                v7 = v7
                v26 := GenS1{F0: v7, F1: 0}
                _ = v25
                _ = v26
            }
        }
        v27 := []int{32768, -2147483648, 0}
        for fi28 := 0; fi28 < 5; fi28++ {
            if len(v27) > 0 {
                v27[0] = fi28
            }
            _ = fi28
        }
        _ = v27
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        v30 := GenS3{F0: 2147483647, F1: uint64(v7), F2: -92}
        v31 := []int{fi29, fi29, (fi29 % ((fi29) | 1))}
        fmt.Println(v7)
        _ = fi29
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 11; fi32++ {
        v7 -= v7
        if false {
            for fi33 := 0; fi33 < 18; fi33++ {
                v7 = (v7 + (v7 / ((v7) | 1)))
                _ = fi33
            }
            for fi34 := 0; fi34 < 19; fi34++ {
                if (!(9223372036854775807 >= 1)) {
                    v7 -= 38
                    var v35 uint8 = v7
                    v36 := []int{fi34, (fi34 | (fi32 | 128))}
                    _ = v35
                    _ = v36
                }
                v7 += (v7 - (v7 & v7))
                v7 += (v7 / ((v7) | 1))
                _ = fi34
            }
            for fi37 := 0; fi37 < 13; fi37++ {
                fmt.Println(v7)
                v7 = (v7 - 128)
                _ = fi37
            }
        } else {
            fmt.Println(v7, v7)
            switch "foo bar" {
            case "foo bar":
                v38 := GenS2{F0: (fi32 % ((32767) | 1))}
                _ = v38
            case "hello":
                v39 := GenS4{F0: 0, F1: 1}
                _ = v39
            default:
                switch "hello" {
                case "foo bar":
                    switch (fi32 & 0) {
                    case 0:
                        v40 := []int{int(v7), 572}
                        _ = v40
                    case 32768:
                        var v41 uint = 572
                        _ = v41
                    default:
                        var v42 int32 = 1
                        _ = v42
                    }
                case "\n\t":
                    for fi43 := 0; fi43 < 2; fi43++ {
                        if (42 < 380) {
                            var v44 string = ("hello" + "a")
                            v45 := GenS4{F0: 594, F1: 0}
                            switch "a" {
                            case "foo bar":
                                v44 = "foo bar"
                            }
                            _ = v44
                            _ = v45
                        } else {
                            fmt.Println(v7, fi43, v7)
                            v7 = v7
                        }
                        var v46 uint16 = 42
                        _ = fi43
                        _ = v46
                    }
                case "":
                    if (true || (!(!(1.5 != float32(0.0))))) {
                        switch "foo bar" {
                        case "a":
                            var v47 int = fi32
                            _ = v47
                        case "":
                            fmt.Println(v7, fi32)
                        case "hello":
                            fmt.Println(v7, fi32)
                        }
                        v48 := []int{fi32, (-(fi32)), -221}
                        switch "foo bar" {
                        case "foo bar":
                            v49 := GenS5{F0: ((true || (!true)) && false), F1: 65535}
                            _ = v49
                        default:
                            for fi50 := 0; fi50 < 1; fi50++ {
                                v7 = (v7 / ((v7) | 1))
                                _ = len(v48)
                                v48 = append(v48, fi50)
                                var v51 int16 = 263
                                if (42 == 0) {
                                    fmt.Println(v51, v7, fi50)
                                }
                                v51 -= (v51 * v51)
                                _ = fi50
                                _ = v51
                            }
                        }
                        _ = v48
                    }
                }
            }
        }
        v7 = v7
        _ = fi32
    }
    _ = v7
    return 0
}

func main() {
    v1 := []string{"", (((((("\n\t" + "\n\t") + ("" + "hello")) + ("a" + "hello")) + ("" + "")) + (("hello" + "\n\t") + "a")) + ""), "a", "foo bar"}
    v2 := []string{("foo bar" + ""), ((("foo bar" + "\n\t") + (("\n\t" + "") + "")) + ("foo bar" + "foo bar"))}
    var v3 float64 = 1e10
    var v4 int16 = 42
    var v5 int64 = 0
    _ = len(v2)
    _ = len(v2)
    for fi52 := 0; fi52 < 1; fi52++ {
        for fi53 := 0; fi53 < 1; fi53++ {
            _ = len(v2)
            fmt.Println(v3)
            if (!(0 >= 0)) {
                var v54 uint = 42
                v3 = 3.14159
                if false {
                    fmt.Println(fi53)
                    var v55 int64 = 9223372036854775807
                    _ = v55
                } else {
                    if false {
                        v54 = 65535
                    } else {
                        fmt.Println(v3, v4, v54)
                        v5 -= (-(v5))
                        if false {
                            v54 = v54
                        } else {
                            v5 = v5
                            var v56 uint8 = 219
                            _ = v56
                        }
                    }
                }
                _ = v54
            } else {
                v57 := GenS6{F0: "\n\t", F1: (fi52 << ((65535) & 63))}
                _ = v57
            }
            _ = fi53
        }
        _ = fi52
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v58 string = "hello"
    var v59 float64 = v3
    v59 = v3
    v60 := GenS5{F0: false, F1: 975}
    switch 42 {
    case 32767:
        v3 = v59
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v60
}
