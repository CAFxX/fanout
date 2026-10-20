package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int64; F1 int16; F2 int }

func main() {
    var v1 int64 = -168
    var v2 uint16 = 0
    var v3 bool = true
    var v4 uint32 = 2147483648
    var v5 int32 = 292
    var v6 bool = (v3 || v3)
    var v7 float32 = 1e10
    v8 := []int{-2147483648}
    fmt.Println(v2, v5, v4)
    v4 -= v4
    var v9 uint64 = 1
    switch 4294967295 {
    case -155:
        v10 := GenS1{F0: 0, F1: -2147483648}
        _ = v10
    case -501:
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 bool = v3
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            v13 := GenS1{F0: 127, F1: (-(fi11))}
            _ = fi11
            _ = v12
            _ = v13
        }
    case 245:
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 int = fi14
            switch v15 {
            case 32768:
                v7 -= v7
            case 877:
                for fi16 := 0; fi16 < 12; fi16++ {
                    v2 = v2
                    switch (v15 & (fi16 & (v15 >> ((1) & 63)))) {
                    case -271:
                        switch "" {
                        case "":
                            v17 := GenS1{F0: 128, F1: (^(fi14))}
                            _ = v17
                        case "hello":
                            v18 := GenS1{F0: 255, F1: -480}
                            _ = v18
                        case "a":
                            v9 += v9
                        }
                    }
                    _ = fi16
                }
            case -762:
                v19 := GenS1{F0: 154, F1: v15}
                _ = v19
            default:
                if v6 {
                    if v3 {
                        if ((v3 && v6) && ("hello" == ("foo bar" + ""))) {
                            v3 = (v7 != (v7 * float32(0.0)))
                            fmt.Println(v7)
                        }
                        var v20 float32 = v7
                        _ = v20
                    }
                    v4 = (v4 ^ 578)
                }
            }
            var v21 int32 = v5
            _ = fi14
            _ = v15
            _ = v21
        }
    default:
        v22 := GenS1{F0: 61, F1: 975}
        _ = v22
    }
    switch ("" + "foo bar") {
    case "hello":
        var v23 uint32 = 42
        _ = v23
    case "foo bar":
        v4 = 100
    default:
        switch -895 {
        case -9223372036854775808:
            for fi24 := 0; fi24 < 7; fi24++ {
                v25 := GenS2{F0: (v1 - 15), F1: 891, F2: (fi24 ^ 9223372036854775807)}
                _ = fi24
                _ = v25
            }
        case 9223372036854775807:
            if len(v8) > 0 {
                v8[0] = 313
            }
        default:
            fmt.Println(v1)
        }
    }
    v26 := GenS2{F0: (^(v1)), F1: 0, F2: -1}
    switch "foo bar" {
    case "a":
        if v6 {
            var v27 bool = (!v3)
            _ = v27
        } else {
            var v28 int8 = -1
            v5 -= -1
            fmt.Println(v1)
            _ = v28
        }
    case "\n\t":
        for fi29 := 0; fi29 < 7; fi29++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            switch "foo bar" {
            case "foo bar":
                var v30 uint64 = (v9 - v9)
                _ = v30
            }
            _ = fi29
        }
    case "hello":
        v31 := GenS2{F0: (v1 ^ (v1 * 9223372036854775807)), F1: -464, F2: -333}
        _ = v31
    }
    var v32 int32 = -70
    v9 += v9
    if false {
        switch "hello" {
        case "a":
            var v33 float64 = 0.0
            _ = v33
        case "hello":
            var v34 uint32 = 2
            _ = v34
        case "foo bar":
            v35 := GenS1{F0: 118, F1: 32768}
            _ = v35
        }
    }
    if (!(255 > 128)) {
        var v36 float64 = -2.25
        _ = v36
    } else {
        if (!(42 < 1000)) {
            v3 = ("hello" == "hello")
            if len(v8) > 0 {
                v8[0] = -121
            }
        } else {
            if (!(302 < 961)) {
                var v37 int16 = 1
                _ = v37
            } else {
                v8 = append(v8, -2147483648)
            }
            if v3 {
                var v38 uint64 = 327
                var v39 int = -549
                v4 -= (v4 % (((v4 ^ 728)) | 1))
                _ = v38
                _ = v39
            }
        }
    }
    _ = len(v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
