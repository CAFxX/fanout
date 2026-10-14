package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int8; F1 int; F2 int32 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 string; F1 int32; F2 int64 }

func main() {
    v1 := []string{}
    var v2 uint16 = 1
    v3 := []int{127, -440, 4294967295}
    var v4 float64 = (0.1 - -2.25)
    var v5 float32 = float32(v4)
    v6 := GenS1{F0: 42, F1: 2}
    var v7 int8 = -118
    fmt.Println(v7)
    v2 = 2
    v4 += float64(v5)
    switch (("" + ("hello" + "")) + "hello") {
    case "foo bar":
        var v8 int8 = 1
        _ = v8
    }
    var v9 uint32 = 4294967295
    _ = len(v3)
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(v4)
        var v11 float32 = 0.0
        v5 -= v5
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: (v7 >> ((65535) & 63)), F1: 520, F2: 1}
    if true {
        var v13 uint = 255
        _ = v13
    } else {
        fmt.Println(v4, v2, v7)
    }
    if true {
        fmt.Println(v5)
        v9 -= 4294967295
    }
    if ((v9 << ((4294967295) & 63)) < 2147483647) {
        if (v2 <= v2) {
            if (false && false) {
                if false {
                    var v14 int16 = -1
                    _ = v14
                }
            }
            var v15 string = "foo bar"
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 int16 = 1
                switch (((("a" + (v15 + "")) + "hello") + "") + v15) {
                case "foo bar":
                    v4 += ((0.1 - v4) * 0.0)
                }
                _ = fi16
                _ = v17
            }
            _ = v15
        } else {
            v18 := GenS3{F0: 4294967295, F1: v2}
            switch 127 {
            case 42:
                if (((!((v5 - 1e10) < float32(v4))) && (127 > 1)) && (((((((1e10 - -2.25) / v4) + v4) * v4) / 1.0) <= v4) && false)) {
                    var v19 int64 = 42
                    switch ("a" + "\n\t") {
                    case "a":
                        for fi20 := 0; fi20 < 5; fi20++ {
                            if (((false && false) && false) || true) {
                                _ = len(v1)
                                v19 = (v19 | -718)
                            }
                            _ = fi20
                        }
                    case "\n\t":
                        _ = len(v1)
                    default:
                        var v21 uint64 = 657
                        _ = v21
                    }
                    _ = v19
                }
            case -704:
                for fi22 := 0; fi22 < 1; fi22++ {
                    for fi23 := 0; fi23 < 9; fi23++ {
                        v9 -= 0
                        switch fi23 {
                        case -195:
                            v24 := GenS3{F0: 1, F1: 2}
                            _ = v24
                        case -284:
                            var v25 int = fi23
                            _ = v25
                        case -1:
                            v26 := GenS3{F0: 18446744073709551615, F1: 788}
                            _ = v26
                        default:
                            if (false || false) {
                                v3 = append(v3, fi23)
                                switch (fi23 * (^(fi23))) {
                                case 0:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case 128:
                                    if len(v3) > 0 {
                                        v3[0] = fi23
                                    }
                                    var v27 uint64 = uint64(v2)
                                    switch "foo bar" {
                                    case "a":
                                        if false {
                                            if (((!false) || true) && false) {
                                                v27 = v27
                                                v28 := GenS2{F0: v7, F1: fi22, F2: int32(v9)}
                                                _ = v28
                                            }
                                            if true {
                                                v27 += 42
                                                v2 -= (v2 - (^(v2)))
                                            } else {
                                                switch "\n\t" {
                                                case "a":
                                                    if len(v1) > 0 {
                                                        v1[0] = (("foo bar" + "foo bar") + "a")
                                                    }
                                                case "hello":
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                case "\n\t":
                                                    if (!((false || false) || (-60 != 59))) {
                                                        fmt.Println(v4)
                                                        v5 = v5
                                                        v29 := GenS2{F0: v7, F1: (^(fi23)), F2: -2147483648}
                                                        _ = v29
                                                    }
                                                    var v30 int16 = 1000
                                                    if len(v1) > 0 {
                                                        v1[0] = ("\n\t" + ("a" + "foo bar"))
                                                    }
                                                    _ = v30
                                                default:
                                                    v5 += v5
                                                }
                                            }
                                        }
                                    default:
                                        v7 = (v7 / (((v7 - (v7 & (v7 ^ int8(v27))))) | 1))
                                    }
                                    _ = v27
                                case 42:
                                    fmt.Println(v4)
                                }
                            }
                        }
                        _ = fi23
                    }
                    _ = fi22
                }
            }
            v31 := GenS3{F0: 553, F1: v2}
            _ = v18
            _ = v31
        }
        v5 = float32(v4)
    } else {
        if len(v3) > 0 {
            v3[0] = 128
        }
        switch -2147483648 {
        case 2147483647:
            v32 := GenS4{F0: "\n\t", F1: int32(v2), F2: 9223372036854775807}
            _ = v32
        case 32768:
            v7 -= v7
        }
        for fi33 := 0; fi33 < 19; fi33++ {
            v34 := GenS1{F0: 4294967295, F1: (^(v2))}
            fmt.Println(v7, v2, v5)
            var v35 string = "a"
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        v4 = 0.0
        _ = fi36
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
