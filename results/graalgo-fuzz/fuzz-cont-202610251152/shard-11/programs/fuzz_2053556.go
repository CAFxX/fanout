package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS5 struct { F0 int8; F1 string; F2 string }
func genFn0(p0 string, p1 int8) uint8 {
    v4 := GenS1{F0: -1, F1: 18446744073709551615}
    var v5 uint8 = 1
    var v6 uint8 = (^(v5))
    p1 -= p1
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(v6, v5)
        fmt.Println(p0, p0)
        var v8 int8 = p1
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return (v5 - (v6 - (v6 | v6)))
}
func genFn1() uint64 {
    switch -9223372036854775808 {
    case -1:
        var v9 float64 = (1e10 + 1e10)
        _ = v9
    case 32767:
        var v10 uint = 1
        _ = v10
    }
    v11 := GenS2{F0: 42, F1: 730, F2: 18446744073709551615}
    var v12 string = "\n\t"
    _ = v11
    _ = v12
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 string = "hello"
    v3 := []string{(v2 + "a"), "hello"}
    if false {
        v3 = append(v3, v2)
        var v13 uint16 = 0
        _ = v13
    } else {
        v2 = "\n\t"
        switch -868 {
        case -461:
            var v14 float64 = (0.0 / 1.0)
            _ = v14
        case 32768:
            fmt.Println(v2)
        case 9223372036854775807:
            var v15 uint32 = 438
            _ = v15
        }
        fmt.Println(v2)
    }
    v3 = append(v3, "\n\t")
    v2 = v2
    fmt.Println(v2)
    v2 = v2
    if ("" > "foo bar") {
        for fi16 := 0; fi16 < 13; fi16++ {
            if (-75 > 127) {
                _ = len(v3)
            } else {
                fmt.Println(fi16, v2)
                v2 = (v2 + v2)
                v1[(-(fi16))] = (v2 + v2)
            }
            _ = fi16
        }
    } else {
        v17 := GenS3{F0: 1, F1: 156}
        for fi18 := 0; fi18 < 11; fi18++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch v2 {
            case "a":
                fmt.Println(v2, v2)
            case "\n\t":
                switch (v2 + (v2 + (v2 + "foo bar"))) {
                case "foo bar":
                    var v19 int64 = 570
                    _ = v19
                }
            default:
                v2 = "\n\t"
            }
            _ = fi18
        }
        _ = v17
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        v2 = (v2 + "foo bar")
        v2 = "\n\t"
        _ = fi20
    }
    v21 := GenS4{F0: 0, F1: -1, F2: 4294967295}
    if false {
        fmt.Println(v2)
        if (0 < 1) {
            switch v2 {
            case "a":
                fmt.Println(v2)
            }
            var v22 int64 = 42
            _ = v22
        }
        fmt.Println(v2)
    } else {
        v23 := GenS5{F0: 127, F1: ("foo bar" + v2), F2: v2}
        _ = v23
    }
    switch (v2 + v2) {
    case "\n\t":
        var v24 uint64 = 422
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
}
