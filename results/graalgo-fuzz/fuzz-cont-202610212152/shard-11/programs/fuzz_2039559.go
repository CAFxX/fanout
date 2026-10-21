package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint; F1 int }
func genFn0() uint16 {
    if true {
        var v5 bool = false
        fmt.Println(v5)
        switch 0 {
        case 42:
            v6 := GenS1{F0: 1, F1: 2}
            _ = v6
        case -671:
            var v7 string = "foo bar"
            _ = v7
        default:
            switch ("" + "a") {
            case "hello":
                if (-453 <= -367) {
                    switch ("" + "foo bar") {
                    case "\n\t":
                        if ((-128 >= 127) || v5) {
                            v5 = v5
                            var v8 bool = v5
                            _ = v8
                        }
                    case "hello":
                        var v9 int64 = 745
                        _ = v9
                    }
                    var v10 string = "hello"
                    _ = v10
                } else {
                    var v11 string = "a"
                    _ = v11
                }
            case "foo bar":
                v5 = ((0 < 1) || (82 != 1))
            default:
                v12 := GenS2{F0: ((("" + "a") + "foo bar") + ""), F1: 128}
                _ = v12
            }
        }
        _ = v5
    } else {
        v13 := []int{-9223372036854775808, 42, -2147483648, -9223372036854775808}
        if len(v13) > 0 {
            v13[0] = 128
        }
        _ = v13
    }
    var v14 float64 = 0.0
    switch 9223372036854775807 {
    case 32767:
        v14 = 1e10
    }
    var v15 int16 = 32767
    v16 := []int{0, 42}
    _ = v14
    _ = v15
    _ = v16
    return 374
}
func genFn1(p0 bool, p1 uint, p2 uint32) string {
    var v17 bool = p0
    v18 := []int{}
    fmt.Println(p2, p1, p0)
    switch ("a" + "") {
    case "foo bar":
        switch "foo bar" {
        case "a":
            p0 = true
        case "":
            var v19 string = ("" + ((((("" + "foo bar") + "\n\t") + ("a" + "")) + "hello") + "foo bar"))
            _ = v19
        default:
            var v20 int = 32768
            _ = v20
        }
    default:
        for fi21 := 0; fi21 < 4; fi21++ {
            v22 := GenS2{F0: "foo bar", F1: 255}
            p0 = v17
            for fi23 := 0; fi23 < 6; fi23++ {
                v18 = append(v18, (^(fi21)))
                v17 = v17
                v18 = append(v18, (fi21 ^ 32767))
                _ = fi23
            }
            _ = fi21
            _ = v22
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return "hello"
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = (("hello" + ("foo bar" + "foo bar")) + "\n\t")
    var v3 uint8 = 1
    v4 := []int{42, 835, 127}
    fmt.Println(v3, v3)
    var v24 float32 = 1e10
    for fi25 := 0; fi25 < 19; fi25++ {
        v26 := GenS2{F0: (v2 + ("hello" + "a")), F1: 127}
        var v27 int8 = 1
        switch ("\n\t" + v2) {
        case "foo bar":
            for fi28 := 0; fi28 < 19; fi28++ {
                switch (-(fi25)) {
                case -153:
                    var v29 bool = (0 != 2)
                    _ = v29
                case -2147483648:
                    v30 := GenS1{F0: 1, F1: 65535}
                    _ = v30
                case -842:
                    fmt.Println(v3, v3, v24)
                }
                _ = fi28
            }
        default:
            var v31 bool = ((false && false) && true)
            _ = v31
        }
        _ = fi25
        _ = v26
        _ = v27
    }
    v32 := GenS2{F0: (v2 + v2), F1: v3}
    v2 = v2
    v24 -= v24
    v33 := GenS2{F0: v2, F1: v3}
    for fi34 := 0; fi34 < 7; fi34++ {
        if ((!true) && false) {
            _ = len(v4)
            switch 436 {
            case -1:
                v35 := GenS2{F0: "\n\t", F1: (v3 ^ 1)}
                _ = v35
            default:
                v2 = v2
            }
        } else {
            for fi36 := 0; fi36 < 13; fi36++ {
                v37 := GenS3{F0: 42, F1: 127}
                _ = fi36
                _ = v37
            }
        }
        _ = fi34
    }
    v2 = "foo bar"
    var v38 float64 = float64(v24)
    for fi39 := 0; fi39 < 7; fi39++ {
        var v40 uint32 = 42
        _ = fi39
        _ = v40
    }
    switch -119 {
    case -900:
        if true {
            for fi41 := 0; fi41 < 7; fi41++ {
                _ = len(v1)
                _ = fi41
            }
            if (2 <= 32767) {
                v24 = v24
                delete(v1, 0)
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        }
    case 4294967295:
        for fi42 := 0; fi42 < 12; fi42++ {
            v4 = append(v4, (fi42 | (fi42 % (((fi42 ^ fi42)) | 1))))
            switch v2 {
            case "a":
                if (true || (((752 == 2) && (!true)) || (!false))) {
                    for fi43 := 0; fi43 < 15; fi43++ {
                        if (0 == 892) {
                            v3 += v3
                            fmt.Println(v2)
                            _ = len(v1)
                        } else {
                            var v44 float64 = v38
                            var v45 bool = (false || (v3 != (v3 % ((v3) | 1))))
                            _ = v44
                            _ = v45
                        }
                        var v46 uint8 = (v3 ^ v3)
                        _ = fi43
                        _ = v46
                    }
                    delete(v1, 2)
                } else {
                    v47 := GenS3{F0: 18446744073709551615, F1: 127}
                    var v48 uint8 = (v3 | (^(v3)))
                    _ = v47
                    _ = v48
                }
            case "foo bar":
                fmt.Println(v38, fi42, v2)
            case "":
                switch v2 {
                case "":
                    if true {
                        var v49 int = (fi42 - 128)
                        var v50 uint8 = 1
                        _ = v49
                        _ = v50
                    }
                case "hello":
                    fmt.Println(v2)
                default:
                    delete(v1, 560)
                }
            default:
                var v51 uint64 = 751
                _ = v51
            }
            var v52 int64 = -9223372036854775808
            _ = fi42
            _ = v52
        }
    case 9223372036854775807:
        var v53 int16 = 1
        _ = v53
    }
    v54 := GenS2{F0: v2, F1: (^(v3))}
    if false {
        var v55 uint16 = genFn0()
        var v56 uint32 = 477
        _ = v55
        _ = v56
    } else {
        if len(v4) > 0 {
            v4[0] = -578
        }
    }
    v57 := GenS3{F0: 255, F1: 9223372036854775807}
    switch 32768 {
    case -997:
        var v58 float64 = v38
        _ = v58
    case 4294967295:
        switch -1 {
        case 4294967295:
            if true {
                for fi59 := 0; fi59 < 7; fi59++ {
                    var v60 uint = 807
                    _ = fi59
                    _ = v60
                }
                v38 += v38
            }
        case -1:
            fmt.Println(v38)
        default:
            fmt.Println(v2, v24)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v54
    _ = v57
}
