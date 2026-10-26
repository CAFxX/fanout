package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int64; F2 int8 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 int16 }

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 1
    v3 := []int{}
    var v4 bool = ((44 <= 59) && true)
    if (4294967295 > 888) {
        if true {
            v4 = (1 <= 571)
        } else {
            var v5 float32 = ((0.1 * 0.0) - 3.14159)
            _ = v5
        }
        if (127 < 211) {
            var v6 uint8 = 2
            _ = v6
        }
    } else {
        if v4 {
            switch -746 {
            case -2147483648:
                v3 = append(v3, 42)
            }
            for fi7 := 0; fi7 < 11; fi7++ {
                _ = len(v1)
                var v8 int8 = 100
                _ = fi7
                _ = v8
            }
            if v4 {
                var v9 uint64 = 106
                v9 = 42
                delete(v1, 9223372036854775807)
                _ = v9
            }
        } else {
            v4 = (!(128 == 255))
        }
        var v10 int32 = 2147483647
        _ = v10
    }
    v11 := GenS1{F0: 9223372036854775807}
    switch (("" + "a") + "foo bar") {
    case "\n\t":
        v2 -= 42
    default:
        var v12 bool = v4
        _ = v12
    }
    _ = len(v1)
    v13 := GenS1{F0: v2}
    v1[32767] = (42 < 32768)
    switch -1 {
    case 42:
        switch "a" {
        case "hello":
            switch "" {
            case "\n\t":
                var v14 int32 = -813
                _ = v14
            default:
                if ("\n\t" == "") {
                    if v4 {
                        var v15 uint64 = 827
                        _ = v15
                    }
                    v16 := GenS2{F0: ("foo bar" + (((("\n\t" + "foo bar") + "") + "\n\t") + ("hello" + ("\n\t" + "")))), F1: (v2 / ((v2) | 1)), F2: 42}
                    v17 := v1[829]
                    _ = v16
                    _ = v17
                }
            }
        case "foo bar":
            fmt.Println(v4, v2)
        default:
            switch "a" {
            case "":
                var v18 int16 = -1
                _ = v18
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "foo bar":
                if len(v3) > 0 {
                    v3[0] = 0
                }
            }
        }
    case 587:
        if v4 {
            if (42 <= 1) {
                v19 := GenS3{F0: 127, F1: 42}
                v2 -= (v2 | v2)
                _ = v19
            }
            for fi20 := 0; fi20 < 2; fi20++ {
                if len(v3) > 0 {
                    v3[0] = fi20
                }
                _ = fi20
            }
        } else {
            if len(v3) > 0 {
                v3[0] = 341
            }
            fmt.Println(v2, v4)
        }
    case 195:
        var v21 string = "foo bar"
        _ = v21
    default:
        if v4 {
            v22 := GenS4{F0: -1}
            v2 += v2
            _ = v22
        } else {
            v2 -= v2
        }
    }
    fmt.Println(v4, v2)
    var v23 float64 = 1e10
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 uint = 255
        switch (fi24 << ((803) & 63)) {
        case -1:
            if v4 {
                if (v4 || (false && (("foo bar" <= "") || (1000 == 1)))) {
                    if (v4 || false) {
                        _ = len(v3)
                    }
                    v23 = float64(((((-2.25 + float32(float64((float32(v23) / 1.0)))) / 1.5) + 0.0) * 0.0))
                } else {
                    v4 = true
                }
                for fi26 := 0; fi26 < 20; fi26++ {
                    v4 = v4
                    v27 := GenS5{F0: 100, F1: -1}
                    switch 4294967295 {
                    case 745:
                        var v28 uint64 = 42
                        _ = v28
                    default:
                        if ((v2 <= v2) && (-32768 > -197)) {
                            _ = len(v1)
                            switch "" {
                            case "a":
                                var v29 int8 = 42
                                _ = v29
                            default:
                                fmt.Println(fi26)
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                    _ = fi26
                    _ = v27
                }
                var v30 int = fi24
                _ = v30
            } else {
                var v31 uint8 = 1
                var v32 bool = true
                _ = v31
                _ = v32
            }
        case -9223372036854775808:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            var v33 bool = false
            _ = v33
        }
        _ = fi24
        _ = v25
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        fmt.Println(fi34, v4)
        for fi35 := 0; fi35 < 9; fi35++ {
            var v36 int = 1
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    switch "\n\t" {
    case "hello":
        v2 -= (v2 % (((^(v2))) | 1))
    }
    v1[42] = v4
    v4 = v4
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
