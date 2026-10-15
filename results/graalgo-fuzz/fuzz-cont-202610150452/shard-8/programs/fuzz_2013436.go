package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint64; F2 uint }
type GenS3 struct { F0 int; F1 int64; F2 uint64 }
func genFn0(p0 int32, p1 int64) int8 {
    p0 += p0
    p1 -= p1
    _ = p0
    _ = p1
    return 100
}
func genFn1() uint64 {
    v7 := GenS1{F0: "\n\t"}
    var v8 int16 = 0
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint64 = 387
    var v2 uint = 255
    var v3 int16 = 472
    var v4 int8 = 0
    var v5 uint16 = 32767
    var v6 uint8 = 127
    fmt.Println(v4, v4, v5)
    if false {
        var v9 uint16 = v5
        var v10 int16 = v3
        _ = v10
        _ = v9
    } else {
        v1 += v1
        fmt.Println(v6, v6)
        switch "hello" {
        case "a":
            v11 := []int{262, 9223372036854775807, int(v4), -113}
            _ = v11
        case "\n\t":
            switch 4294967295 {
            case 973:
                v12 := []int{-813}
                _ = v12
            case -9223372036854775808:
                v13 := GenS1{F0: ("" + ("\n\t" + "hello"))}
                _ = v13
            }
        case "foo bar":
            v14 := GenS1{F0: "hello"}
            _ = v14
        default:
            if true {
                if false {
                    if (297 != -110) {
                        v5 -= 42
                    }
                } else {
                    fmt.Println(v4, v1, v4)
                    v2 = 255
                    var v15 int64 = -9223372036854775808
                    _ = v15
                }
            }
        }
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        v2 -= 221
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 string = ("hello" + "\n\t")
            var v19 uint = (^(v2))
            var v20 uint16 = v5
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        v21 := []int{127, (fi16 - fi16), (fi16 + (fi16 | (fi16 ^ (fi16 - fi16))))}
        _ = fi16
        _ = v21
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        v4 = (v4 & v4)
        if (v1 >= 609) {
            var v23 int16 = (v3 % ((v3) | 1))
            var v24 int8 = v4
            v25 := GenS2{F0: 0, F1: 0, F2: (v2 << ((uint(v5)) & 63))}
            _ = v23
            _ = v24
            _ = v25
        } else {
            v26 := GenS1{F0: ""}
            if true {
                var v27 float32 = ((1.5 + 1.5) / 1.0)
                v28 := []int{}
                switch "\n\t" {
                case "":
                    var v29 uint32 = 1
                    _ = v29
                case "a":
                    v3 += (v3 >> ((18446744073709551615) & 63))
                }
                _ = v27
                _ = v28
            } else {
                v30 := GenS2{F0: 1, F1: v1, F2: v2}
                _ = v30
            }
            _ = v26
        }
        v6 += (v6 ^ 2)
        _ = fi22
    }
    if false {
        v3 -= (v3 << ((uint(v1)) & 63))
        for fi31 := 0; fi31 < 11; fi31++ {
            switch (fi31 | 9223372036854775807) {
            case -9223372036854775808:
                v32 := GenS3{F0: (fi31 - int(v2)), F1: 675, F2: (v1 / (((v1 << (((v2 | (v2 | v2))) & 63))) | 1))}
                _ = v32
            default:
                fmt.Println(v1, v2, v4)
            }
            fmt.Println(v2, v2)
            _ = fi31
        }
        if ((0.1 / float64(1.5)) > (0.0 * (1.5 + (0.0 - 1e10)))) {
            for fi33 := 0; fi33 < 2; fi33++ {
                for fi34 := 0; fi34 < 2; fi34++ {
                    switch fi34 {
                    case 32767:
                        var v35 int64 = -212
                        _ = v35
                    }
                    v36 := []int{(fi33 + 0), fi34, 42, (fi33 | (fi33 >> (((^(v2))) & 63)))}
                    var v37 int16 = (v3 + -32768)
                    _ = fi34
                    _ = v36
                    _ = v37
                }
                _ = fi33
            }
            for fi38 := 0; fi38 < 8; fi38++ {
                for fi39 := 0; fi39 < 13; fi39++ {
                    v40 := []int{}
                    v4 = v4
                    _ = fi39
                    _ = v40
                }
                v1 = (v1 * (^(v1)))
                _ = fi38
            }
            v3 = (v3 * v3)
        }
    }
    v6 += (v6 << ((1) & 63))
    v41 := []int{32768, 150}
    switch (("hello" + "foo bar") + ((("" + "foo bar") + (("" + "hello") + (("" + "") + "hello"))) + ("a" + "foo bar"))) {
    case "foo bar":
        for fi42 := 0; fi42 < 12; fi42++ {
            v6 = (^(v6))
            _ = fi42
        }
    case "\n\t":
        if (2147483647 >= uint32(v6)) {
            fmt.Println(v6, v2)
            var v43 string = "\n\t"
            _ = v43
        }
    default:
        v4 = v4
    }
    if (!true) {
        switch "" {
        case "\n\t":
            if len(v41) >= 2 {
                v41 = v41[1:len(v41)]
            }
        }
        fmt.Println(v6, v1)
        switch -790 {
        case 32768:
            _ = len(v41)
        default:
            var v44 uint = 654
            _ = v44
        }
    } else {
        fmt.Println(v3)
        if (v2 > 18446744073709551615) {
            for fi45 := 0; fi45 < 11; fi45++ {
                var v46 int64 = 0
                _ = fi45
                _ = v46
            }
            var v47 int64 = 9223372036854775807
            _ = v47
        }
        for fi48 := 0; fi48 < 3; fi48++ {
            fmt.Println(v4)
            var v49 uint8 = 32
            v50 := GenS2{F0: 1, F1: v1, F2: 65535}
            _ = fi48
            _ = v49
            _ = v50
        }
    }
    switch -327 {
    case 968:
        v6 = v6
    default:
        switch 127 {
        case -403:
            var v51 bool = (v1 < v1)
            _ = v51
        }
    }
    switch (("" + "hello") + "hello") {
    case "foo bar":
        if len(v41) >= 2 {
            v41 = v41[1:len(v41)]
        }
    }
    fmt.Println(v1)
    v3 = v3
    switch "hello" {
    case "a":
        if false {
            for fi52 := 0; fi52 < 15; fi52++ {
                switch ("\n\t" + ("\n\t" + "\n\t")) {
                case "hello":
                    v53 := GenS1{F0: "hello"}
                    _ = v53
                case "\n\t":
                    var v54 bool = (65535 > 191)
                    _ = v54
                case "":
                    v6 = v6
                default:
                    var v55 bool = (true || true)
                    _ = v55
                }
                _ = len(v41)
                _ = fi52
            }
        } else {
            v6 -= v6
        }
    case "hello":
        var v56 uint = v2
        _ = v56
    case "\n\t":
        var v57 uint16 = v5
        _ = v57
    }
    v41 = append(v41, 566)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
