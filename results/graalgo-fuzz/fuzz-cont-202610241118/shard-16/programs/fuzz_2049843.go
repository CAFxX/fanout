package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int32; F1 bool; F2 uint64 }
type GenS5 struct { F0 bool; F1 int8; F2 uint16 }

func main() {
    var v1 uint32 = 42
    var v2 int32 = 42
    var v3 uint64 = 141
    var v4 int32 = v2
    v5 := []string{(("hello" + "foo bar") + "foo bar"), "a", "hello", "hello"}
    var v6 uint64 = 2
    v5 = append(v5, "\n\t")
    v2 -= 42
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int64 = 855
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: -1, F1: 42}
    v10 := GenS1{F0: -108, F1: 409}
    for fi11 := 0; fi11 < 12; fi11++ {
        fmt.Println(v1, v3)
        var v12 float64 = float64((-2.25 + 1.5))
        switch 537 {
        case 1:
            v13 := GenS2{F0: 947}
            _ = v13
        case 128:
            _ = len(v5)
        case 42:
            if ((fi11 & fi11) > fi11) {
                var v14 uint32 = 241
                if ((false && false) || (true && (0 <= 0))) {
                    switch (fi11 << ((584) & 63)) {
                    case -2147483648:
                        fmt.Println(v3, v4)
                    }
                }
                v15 := GenS2{F0: (v6 ^ v3)}
                _ = v14
                _ = v15
            }
        }
        _ = fi11
        _ = v12
    }
    if true {
        if ((v2 + (v4 & v2)) == (v4 >> ((65535) & 63))) {
            if (132 > 42) {
                for fi16 := 0; fi16 < 15; fi16++ {
                    var v17 string = "a"
                    v18 := GenS3{F0: -801, F1: -1}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                switch "a" {
                case "hello":
                    v6 -= v6
                case "a":
                    v1 = (v1 + 576)
                case "foo bar":
                    var v19 int = 32767
                    _ = v19
                default:
                    v5 = append(v5, "foo bar")
                }
            }
            v4 = -1
        }
        v6 -= (v6 & 893)
        var v20 int = 214
        _ = v20
    } else {
        v3 += (^(v6))
        v2 -= 2147483647
        if true {
            var v21 uint32 = (v1 / (((^(v1))) | 1))
            v5 = append(v5, "a")
            for fi22 := 0; fi22 < 6; fi22++ {
                _ = len(v5)
                var v23 int32 = 1
                _ = fi22
                _ = v23
            }
            _ = v21
        } else {
            switch 2147483647 {
            case 4294967295:
                v6 = 0
            case 1:
                if len(v5) > 0 {
                    v5[0] = "hello"
                }
            case 32768:
                var v24 uint32 = v1
                _ = v24
            default:
                var v25 int64 = 42
                _ = v25
            }
            if (42 == 388) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                for fi26 := 0; fi26 < 16; fi26++ {
                    var v27 float64 = -2.25
                    _ = fi26
                    _ = v27
                }
            }
        }
    }
    if ((v4 / ((0) | 1)) > int32(v6)) {
        if false {
            var v28 uint16 = 255
            v5 = append(v5, "")
            _ = v28
        } else {
            v1 = v1
            for fi29 := 0; fi29 < 19; fi29++ {
                for fi30 := 0; fi30 < 6; fi30++ {
                    v31 := GenS4{F0: (v2 - (v4 >> ((118) & 63))), F1: true, F2: uint64(v1)}
                    v3 = (v6 - (^(v3)))
                    _ = fi30
                    _ = v31
                }
                v5 = append(v5, (("foo bar" + "hello") + "a"))
                _ = fi29
            }
        }
    }
    v1 -= (v1 ^ (^(v1)))
    switch "hello" {
    case "\n\t":
        switch "foo bar" {
        case "":
            _ = len(v5)
        case "hello":
            v2 += v2
        case "a":
            if false {
                var v32 int32 = 932
                _ = v32
            }
        default:
            for fi33 := 0; fi33 < 15; fi33++ {
                if false {
                    v34 := GenS4{F0: v2, F1: true, F2: 1}
                    fmt.Println(v1, v2, v6)
                    _ = v34
                }
                _ = fi33
            }
        }
    }
    if false {
        var v35 float64 = (0.1 * (3.14159 / 1e10))
        switch 2147483647 {
        case 4294967295:
            v6 = v3
        case 9223372036854775807:
            for fi36 := 0; fi36 < 16; fi36++ {
                v5 = append(v5, "hello")
                fmt.Println(fi36, v2)
                v37 := GenS5{F0: false, F1: 127, F2: 32767}
                _ = fi36
                _ = v37
            }
        }
        switch "hello" {
        case "hello":
            v5 = append(v5, "\n\t")
        case "foo bar":
            fmt.Println(v6, v3)
        }
        _ = v35
    }
    fmt.Println(v4, v4)
    switch "\n\t" {
    case "foo bar":
        v1 = v1
    case "a":
        var v38 int64 = 1
        _ = v38
    case "\n\t":
        for fi39 := 0; fi39 < 6; fi39++ {
            fmt.Println(v1, v6)
            _ = fi39
        }
    }
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
