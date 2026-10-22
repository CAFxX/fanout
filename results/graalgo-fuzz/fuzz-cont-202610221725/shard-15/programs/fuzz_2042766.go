package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 uint32; F1 int64; F2 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 string; F1 uint8 }
func genFn0(p0 bool, p1 bool) bool {
    fmt.Println(p1, p0)
    if (519 > -32768) {
        p1 = true
        v6 := []int{}
        _ = v6
    } else {
        v7 := []int{-1, -9223372036854775808, 0}
        var v8 uint = 1
        v8 = 42
        _ = v7
        _ = v8
    }
    switch "" {
    case "\n\t":
        var v9 uint8 = 0
        _ = v9
    case "hello":
        v10 := []int{127, 1}
        _ = v10
    case "a":
        fmt.Println(p1, p0)
    default:
        p0 = (768 < 18446744073709551615)
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 uint8 = 114
        v13 := GenS1{F0: 4294967295, F1: (-277 == -32768), F2: 541}
        p1 = p0
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int64 = 42
    _ = p0
    _ = p1
    _ = v14
    return false
}
func genFn1(p0 int32, p1 uint32) int32 {
    if false {
        p0 += (p0 | (p0 >> ((4294967295) & 63)))
    } else {
        fmt.Println(p0, p1)
        p1 += p1
    }
    if true {
        switch -1 {
        case 128:
            for fi15 := 0; fi15 < 10; fi15++ {
                switch ("hello" + "") {
                case "\n\t":
                    fmt.Println(fi15, fi15)
                }
                _ = fi15
            }
        default:
            var v16 uint64 = 392
            _ = v16
        }
        p0 = 1
    }
    var v17 bool = false
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 int16 = 601
        fmt.Println(v17)
        p1 = (p1 + p1)
        _ = fi18
        _ = v19
    }
    if (true || true) {
        if (18446744073709551615 == 4294967295) {
            switch ((("foo bar" + "foo bar") + "hello") + "") {
            case "":
                switch -9223372036854775808 {
                case -612:
                    v20 := []int{32767, 32768}
                    _ = v20
                default:
                    v21 := GenS1{F0: 0, F1: v17, F2: 931}
                    _ = v21
                }
            case "a":
                v22 := GenS2{F0: uint8(p1)}
                _ = v22
            case "foo bar":
                var v23 int64 = -9223372036854775808
                _ = v23
            }
            var v24 float64 = float64(((0.1 + 0.0) * 3.14159))
            _ = v24
        }
        fmt.Println(p0)
        if true {
            if (-38 != 6) {
                fmt.Println(p0, p0, p1)
                if true {
                    for fi25 := 0; fi25 < 14; fi25++ {
                        var v26 int16 = 32767
                        fmt.Println(p0, fi25)
                        _ = fi25
                        _ = v26
                    }
                    switch ("hello" + "\n\t") {
                    case "\n\t":
                        v27 := []int{-2147483648, -636, -1}
                        _ = v27
                    case "":
                        p1 = (p1 | (p1 * 2))
                    default:
                        for fi28 := 0; fi28 < 9; fi28++ {
                            var v29 uint64 = 42
                            _ = fi28
                            _ = v29
                        }
                    }
                    v17 = ((0.0 + (0.0 - (0.0 - -2.25))) == ((float32((1e10 + (1e10 / 1.5))) - 0.0) * 1e10))
                } else {
                    for fi30 := 0; fi30 < 12; fi30++ {
                        v17 = v17
                        v31 := []int{fi30, 9223372036854775807, (fi30 ^ (^(fi30))), (fi30 << ((18446744073709551615) & 63))}
                        _ = fi30
                        _ = v31
                    }
                    v32 := GenS3{F0: 9223372036854775807, F1: 0}
                    _ = v32
                }
            }
            switch "" {
            case "":
                v33 := []int{-438, 4294967295, 42}
                _ = v33
            case "hello":
                p1 -= (p1 ^ (^(p1)))
            }
        } else {
            switch "\n\t" {
            case "foo bar":
                for fi34 := 0; fi34 < 19; fi34++ {
                    v35 := GenS4{F0: 867, F1: p1}
                    v17 = ((!(!(v17 || v17))) && (127 > 28))
                    var v36 bool = (255 < 0)
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            case "hello":
                v17 = false
            default:
                for fi37 := 0; fi37 < 19; fi37++ {
                    v38 := []int{(^(fi37))}
                    var v39 int8 = 0
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            }
            fmt.Println(p0, p0)
        }
    } else {
        fmt.Println(p0, p0, v17)
    }
    _ = p0
    _ = p1
    _ = v17
    return -1
}

func main() {
    v1 := []string{((("foo bar" + ("" + ("hello" + ("" + "a")))) + "\n\t") + ("hello" + "\n\t"))}
    var v2 uint16 = 979
    var v3 int16 = 42
    var v4 int64 = 1
    var v5 int32 = -1
    v40 := GenS5{F0: 959, F1: v4, F2: int8(v3)}
    v41 := GenS6{F0: "a"}
    v42 := GenS3{F0: (-(v4)), F1: (v3 ^ v3)}
    v43 := GenS7{F0: -2147483648}
    for fi44 := 0; fi44 < 7; fi44++ {
        v45 := GenS6{F0: ("a" + "foo bar")}
        v46 := GenS5{F0: 65535, F1: v4, F2: -106}
        _ = fi44
        _ = v45
        _ = v46
    }
    if ("hello" == "") {
        v1 = append(v1, ("foo bar" + "hello"))
    } else {
        v3 = v3
        for fi47 := 0; fi47 < 15; fi47++ {
            switch "hello" {
            case "hello":
                var v48 string = "hello"
                _ = v48
            }
            v49 := GenS1{F0: 0, F1: false, F2: 97}
            switch "a" {
            case "a":
                fmt.Println(fi47)
            case "hello":
                v50 := GenS3{F0: -58, F1: v3}
                _ = v50
            case "":
                if false {
                    v5 += (v5 & -3)
                    var v51 int = (fi47 + (-(fi47)))
                    _ = v51
                } else {
                    for fi52 := 0; fi52 < 4; fi52++ {
                        if false {
                            v53 := GenS3{F0: v4, F1: (-(v3))}
                            _ = v53
                        } else {
                            if true {
                                var v54 uint = 583
                                v2 = v2
                                fmt.Println(fi52, v54, fi47)
                                _ = v54
                            }
                        }
                        v2 = v2
                        _ = len(v1)
                        _ = fi52
                    }
                }
            default:
                _ = len(v1)
            }
            _ = fi47
            _ = v49
        }
        v4 += (^(v4))
    }
    var v55 int16 = v3
    if false {
        v56 := GenS5{F0: 65535, F1: (v4 >> ((18) & 63)), F2: 100}
        var v57 bool = false
        v58 := GenS4{F0: v3, F1: 42}
        _ = v56
        _ = v57
        _ = v58
    }
    for fi59 := 0; fi59 < 18; fi59++ {
        fmt.Println(v4, v3, v4)
        v1 = append(v1, (((("foo bar" + "\n\t") + "a") + ((("a" + "") + (("" + "a") + "hello")) + (("hello" + (("\n\t" + "a") + ((("a" + "foo bar") + ("\n\t" + ("foo bar" + "foo bar"))) + "hello"))) + "hello"))) + ("" + "\n\t")))
        var v60 bool = (false && true)
        _ = fi59
        _ = v60
    }
    fmt.Println(v5, v5)
    var v61 uint = 42
    switch 197 {
    case 1:
        if false {
            fmt.Println(v55, v2, v61)
        }
    default:
        v62 := GenS8{F0: "a", F1: 128}
        _ = v62
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v55
    _ = v61
}
