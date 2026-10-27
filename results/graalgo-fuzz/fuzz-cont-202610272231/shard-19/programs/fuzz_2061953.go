package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int }
type GenS2 struct { F0 int; F1 uint8; F2 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 uint8 }
func genFn0(p0 int64) int64 {
    switch ("\n\t" + "") {
    case "a":
        var v5 bool = (false || false)
        _ = v5
    case "foo bar":
        v6 := GenS1{F0: 42, F1: 0, F2: -496}
        _ = v6
    case "\n\t":
        var v7 uint = 18446744073709551615
        _ = v7
    default:
        v8 := GenS2{F0: 32768, F1: 42, F2: 1000}
        _ = v8
    }
    v9 := GenS2{F0: 128, F1: 255, F2: -1}
    if ((!true) && (-59 < 42)) {
        v10 := GenS2{F0: 549, F1: 71, F2: 1000}
        v11 := GenS2{F0: -40, F1: 0, F2: 0}
        _ = v10
        _ = v11
    } else {
        p0 = (p0 - (p0 << ((248) & 63)))
        v12 := []int{938, -9223372036854775808, 32767}
        switch 73 {
        case 0:
            for fi13 := 0; fi13 < 12; fi13++ {
                for fi14 := 0; fi14 < 3; fi14++ {
                    if len(v12) > 0 {
                        v12[0] = (fi13 * (fi14 + fi13))
                    }
                    _ = fi14
                }
                for fi15 := 0; fi15 < 8; fi15++ {
                    v16 := GenS2{F0: -403, F1: 1, F2: -1}
                    fmt.Println(p0, fi13)
                    _ = fi15
                    _ = v16
                }
                _ = fi13
            }
        }
        _ = v12
    }
    var v17 float64 = float64(3.14159)
    _ = p0
    _ = v17
    _ = v9
    return p0
}

func main() {
    var v1 uint64 = 2
    var v2 float64 = (float64(((1e10 + 3.14159) + 0.1)) - float64(((((-2.25 - 3.14159) / 1.0) * (-2.25 - 0.0)) / float32(1.5))))
    var v3 uint = 65535
    v4 := []string{("foo bar" + "hello"), "foo bar", (("" + "\n\t") + ("foo bar" + ("" + ("hello" + "a")))), ("\n\t" + (("a" + "") + ("\n\t" + "\n\t")))}
    fmt.Println(v2)
    if (v2 >= v2) {
        switch ("" + "foo bar") {
        case "":
            v1 -= 18446744073709551615
        case "foo bar":
            switch "hello" {
            case "foo bar":
                switch "hello" {
                case "a":
                    fmt.Println(v2, v2)
                }
            case "a":
                v2 = (0.1 * v2)
            case "\n\t":
                for fi18 := 0; fi18 < 1; fi18++ {
                    v3 -= 65535
                    v3 += 4294967295
                    fmt.Println(fi18, fi18)
                    _ = fi18
                }
            default:
                if (!true) {
                    var v19 int16 = 1000
                    fmt.Println(v19, v2)
                    _ = v19
                }
            }
        default:
            var v20 int32 = -468
            _ = v20
        }
        v2 = 1e10
    }
    if true {
        if ((!(-601 < 42)) && (-108 < 0)) {
            var v21 int64 = 275
            var v22 int8 = 13
            _ = v21
            _ = v22
        }
        var v23 int64 = genFn0(1)
        for fi24 := 0; fi24 < 3; fi24++ {
            var v25 string = "hello"
            v26 := GenS3{F0: (v23 << (((v3 / ((425) | 1))) & 63))}
            switch ((v25 + v25) + v25) {
            case "":
                v27 := GenS2{F0: fi24, F1: 255, F2: -1}
                _ = v27
            case "hello":
                for fi28 := 0; fi28 < 13; fi28++ {
                    v4 = append(v4, "")
                    _ = fi28
                }
            case "a":
                fmt.Println(v3)
            default:
                _ = len(v4)
            }
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = v23
    } else {
        v2 += ((-2.25 - v2) / float64(0.1))
        v3 += v3
        var v29 bool = false
        _ = v29
    }
    v3 = 33
    switch ("\n\t" + ((("foo bar" + (((("hello" + "a") + "a") + "") + "\n\t")) + ("a" + "foo bar")) + "hello")) {
    case "":
        var v30 int16 = 379
        _ = v30
    case "\n\t":
        for fi31 := 0; fi31 < 4; fi31++ {
            switch fi31 {
            case -870:
                fmt.Println(v1)
            default:
                fmt.Println(fi31, v2)
            }
            for fi32 := 0; fi32 < 13; fi32++ {
                _ = len(v4)
                var v33 int8 = 127
                _ = fi32
                _ = v33
            }
            var v34 int = int(v3)
            _ = fi31
            _ = v34
        }
    case "hello":
        for fi35 := 0; fi35 < 11; fi35++ {
            for fi36 := 0; fi36 < 1; fi36++ {
                var v37 uint = 18446744073709551615
                _ = fi36
                _ = v37
            }
            _ = fi35
        }
    default:
        switch "foo bar" {
        case "hello":
            if true {
                _ = len(v4)
            }
        case "a":
            for fi38 := 0; fi38 < 13; fi38++ {
                fmt.Println(v2)
                var v39 int16 = -32768
                var v40 uint64 = (v1 | (v1 & 4294967295))
                _ = fi38
                _ = v39
                _ = v40
            }
        case "foo bar":
            _ = len(v4)
        }
    }
    var v41 uint16 = 129
    for fi42 := 0; fi42 < 10; fi42++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = len(v4)
        fmt.Println(fi42, fi42)
        _ = fi42
    }
    if true {
        var v43 uint8 = 42
        for fi44 := 0; fi44 < 17; fi44++ {
            fmt.Println(v41)
            switch -9223372036854775808 {
            case -2147483648:
                if false {
                    for fi45 := 0; fi45 < 19; fi45++ {
                        fmt.Println(v41)
                        _ = fi45
                    }
                    v46 := GenS4{F0: true, F1: (v43 >> ((v3) & 63))}
                    _ = v46
                } else {
                    for fi47 := 0; fi47 < 11; fi47++ {
                        switch fi47 {
                        case -1:
                            var v48 uint16 = v41
                            _ = v48
                        }
                        _ = len(v4)
                        _ = fi47
                    }
                    if (-9223372036854775808 <= 9223372036854775807) {
                        v4 = append(v4, "\n\t")
                        var v49 uint64 = 0
                        _ = v49
                    } else {
                        v4 = append(v4, ("" + "\n\t"))
                    }
                    v50 := GenS1{F0: v1, F1: fi44, F2: 505}
                    _ = v50
                }
            default:
                switch ("foo bar" + "") {
                case "a":
                    if false {
                        var v51 int16 = 492
                        var v52 uint64 = 2
                        fmt.Println(v43)
                        _ = v51
                        _ = v52
                    } else {
                        v43 = v43
                    }
                default:
                    v43 = v43
                }
            }
            _ = fi44
        }
        fmt.Println(v41, v3)
        _ = v43
    } else {
        v1 = v1
    }
    v53 := GenS1{F0: (v1 << (((v3 / (((v3 | (v3 | 759))) | 1))) & 63)), F1: 128, F2: 42}
    fmt.Println(v3, v3, v41)
    if (32767 < -441) {
        v41 += 755
        fmt.Println(v2, v2)
        for fi54 := 0; fi54 < 17; fi54++ {
            var v55 uint16 = (v41 & (v41 % ((2) | 1)))
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v56 uint32 = 2147483647
            _ = fi54
            _ = v55
            _ = v56
        }
    }
    v3 -= 65535
    v3 += (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v53
}
