package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 uint16; F2 uint8 }
type GenS7 struct { F0 int8; F1 int16; F2 int8 }
type GenS8 struct { F0 bool; F1 uint8; F2 uint16 }
func genFn0(p0 bool) uint32 {
    var v6 uint = 18446744073709551615
    v6 += 42
    for fi7 := 0; fi7 < 16; fi7++ {
        p0 = (((("foo bar" + "hello") + "foo bar") + "a") != ((("foo bar" + "") + (("\n\t" + "foo bar") + "hello")) + "a"))
        _ = fi7
    }
    _ = p0
    _ = v6
    return 42
}

func main() {
    v1 := []string{(("\n\t" + "foo bar") + "foo bar"), "foo bar"}
    var v2 uint8 = 127
    var v3 uint16 = 65535
    v4 := make(map[string]int)
    var v5 uint32 = 2
    for fi8 := 0; fi8 < 15; fi8++ {
        if (!false) {
            for fi9 := 0; fi9 < 1; fi9++ {
                v10 := GenS1{F0: -2147483648, F1: true, F2: ("hello" + ("a" + "a"))}
                switch ("\n\t" + "foo bar") {
                case "\n\t":
                    fmt.Println(fi9, v5)
                case "foo bar":
                    fmt.Println(v3, fi8)
                }
                if false {
                    var v11 int64 = -1
                    _ = v11
                }
                switch "" {
                case "\n\t":
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                }
                if true {
                    v1 = append(v1, "a")
                    v5 = 581
                } else {
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                    if (1000 != 0) {
                        v12 := GenS2{F0: (fi9 >> ((42) & 63))}
                        _ = len(v1)
                        _ = v12
                    }
                    v13 := v4[""]
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v14 := GenS2{F0: (v13 << ((uint(v3)) & 63))}
                    _ = v13
                    _ = v14
                }
                v2 += (v2 >> ((42) & 63))
                _ = fi9
                _ = v10
            }
            if ((-2.25 + -2.25) <= 0.1) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
        fmt.Println(v2, v5, fi8)
        _ = fi8
    }
    v15 := GenS3{F0: 0, F1: 18446744073709551615}
    switch 32768 {
    case 2147483647:
        if true {
            _ = len(v4)
        } else {
            v16 := GenS4{F0: 32767, F1: 546}
            _ = len(v4)
            var v17 int8 = 42
            _ = v16
            _ = v17
        }
    case 128:
        v18 := GenS5{F0: 32767}
        _ = v18
    case -9223372036854775808:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        v19 := GenS6{F0: 0, F1: v3, F2: (v2 | (v2 ^ 128))}
        _ = v19
    }
    v5 = (v5 / ((genFn0(((v5 >> ((1) & 63)) > uint32(v2)))) | 1))
    v2 -= (v2 + 255)
    fmt.Println(v3)
    for fi20 := 0; fi20 < 14; fi20++ {
        var v21 uint8 = (v2 / (((^(v2))) | 1))
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 int = (fi20 << ((42) & 63))
            var v24 int16 = -295
            _ = fi22
            _ = v23
            _ = v24
        }
        switch "foo bar" {
        case "hello":
            switch fi20 {
            case 32768:
                delete(v4, "hello")
            case -975:
                var v25 float64 = 0.1
                _ = v25
            case 1:
                if false {
                    if false {
                        switch -9223372036854775808 {
                        case -852:
                            var v26 int = (fi20 << ((956) & 63))
                            _ = v26
                        case -499:
                            fmt.Println(v3)
                        case 128:
                            for fi27 := 0; fi27 < 8; fi27++ {
                                var v28 float64 = ((float64(0.0) * 1e10) + ((-2.25 - ((1.5 / -2.25) + 0.0)) - 1.5))
                                switch (fi27 - fi20) {
                                case -2147483648:
                                    _ = len(v4)
                                case -49:
                                    v28 += v28
                                case -9223372036854775808:
                                    _ = len(v1)
                                }
                                v21 = (v2 | (v2 | (v2 ^ 255)))
                                _ = fi27
                                _ = v28
                            }
                        }
                        fmt.Println(v3, v21)
                        for fi29 := 0; fi29 < 20; fi29++ {
                            var v30 int8 = 127
                            v1 = append(v1, "foo bar")
                            v31 := GenS7{F0: (v30 - (v30 ^ (-(v30)))), F1: -895, F2: v30}
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                    } else {
                        switch (("foo bar" + "") + ("foo bar" + "foo bar")) {
                        case "":
                            var v32 float32 = (-2.25 + 1.5)
                            _ = v32
                        default:
                            var v33 int = (fi20 - (-(fi20)))
                            _ = v33
                        }
                        fmt.Println(v3, v21, v21)
                    }
                    var v34 int64 = int64(v3)
                    _ = v34
                } else {
                    v2 += v2
                }
            }
        case "":
            _ = len(v4)
        default:
            var v35 uint64 = 2
            _ = v35
        }
        _ = fi20
        _ = v21
    }
    for fi36 := 0; fi36 < 6; fi36++ {
        var v37 int = fi36
        for fi38 := 0; fi38 < 9; fi38++ {
            v39 := GenS8{F0: true, F1: (^(v2)), F2: (^(v3))}
            v37 = int(v2)
            if false {
                if (true && (0 > 0)) {
                    if false {
                        v3 -= 65535
                    }
                    v3 = 2
                    switch v37 {
                    case -246:
                        var v40 int8 = -1
                        v40 -= 100
                        var v41 float32 = 1e10
                        var v42 int8 = (v40 * (v40 - -1))
                        _ = v40
                        _ = v41
                        _ = v42
                    case -2147483648:
                        fmt.Println(v2, v37)
                    }
                }
            }
            _ = fi38
            _ = v39
        }
        _ = fi36
        _ = v37
    }
    v3 = v3
    v2 += 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
