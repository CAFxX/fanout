package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint64; F1 int16 }
func genFn0(p0 int16, p1 string) string {
    var v6 int16 = p0
    var v7 int32 = 1
    if ((!(true && false)) && true) {
        for fi8 := 0; fi8 < 2; fi8++ {
            switch fi8 {
            case 42:
                switch 1 {
                case -9223372036854775808:
                    for fi9 := 0; fi9 < 10; fi9++ {
                        v10 := []int{(fi9 % ((2147483647) | 1)), 1, fi8}
                        v11 := GenS1{F0: 42}
                        v12 := GenS2{F0: 0}
                        v13 := GenS3{F0: 65535, F1: 390, F2: 4294967295}
                        v6 = (^(p0))
                        _ = fi9
                        _ = v10
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                default:
                    if (true && true) {
                        var v14 uint8 = 0
                        _ = v14
                    } else {
                        for fi15 := 0; fi15 < 10; fi15++ {
                            var v16 float32 = 1.5
                            if false {
                                v17 := []int{0, 68, (fi8 >> ((0) & 63))}
                                _ = v17
                            }
                            _ = fi15
                            _ = v16
                        }
                        var v18 string = "foo bar"
                        switch 127 {
                        case 783:
                            v19 := GenS2{F0: -1}
                            _ = v19
                        case 32767:
                            fmt.Println(v6)
                        case 0:
                            fmt.Println(p0, p0)
                        }
                        _ = v18
                    }
                }
            case -9223372036854775808:
                var v20 bool = (0 != 215)
                _ = v20
            }
            _ = fi8
        }
        var v21 uint64 = 0
        _ = v21
    }
    if true {
        v22 := GenS4{F0: p0}
        var v23 int64 = 42
        v24 := []int{2147483647}
        _ = v22
        _ = v23
        _ = v24
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p1 + (((((("hello" + p1) + p1) + "hello") + p1) + ("foo bar" + "hello")) + p1))
}
func genFn1(p0 int8, p1 uint8, p2 int64) string {
    v25 := GenS3{F0: uint32(p1), F1: 2, F2: 1}
    p1 = p1
    for fi26 := 0; fi26 < 4; fi26++ {
        if false {
            v27 := []int{}
            var v28 int64 = -922
            _ = v27
            _ = v28
        }
        var v29 int8 = (p0 << ((453) & 63))
        _ = fi26
        _ = v29
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return "a"
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{(("a" + "") + ((("foo bar" + (((("" + "foo bar") + "foo bar") + "a") + ("foo bar" + "a"))) + ("a" + "hello")) + "foo bar")), ("hello" + "foo bar"), "foo bar"}
    var v3 uint8 = 0
    var v4 int8 = 0
    var v5 uint32 = 42
    v4 += (v4 & v4)
    if len(v2) > 0 {
        v2[0] = ("foo bar" + "hello")
    }
    if (false && false) {
        v30 := GenS1{F0: 694}
        _ = v30
    } else {
        _ = len(v2)
        switch "hello" {
        case "a":
            if true {
                v4 += (^(v4))
            }
        case "\n\t":
            v4 += 1
        }
    }
    v4 = v4
    var v31 uint = 18446744073709551615
    switch -9223372036854775808 {
    case 32768:
        switch "foo bar" {
        case "\n\t":
            for fi32 := 0; fi32 < 17; fi32++ {
                var v33 int32 = -1
                v34 := GenS5{F0: fi32, F1: fi32}
                _ = fi32
                _ = v33
                _ = v34
            }
        case "":
            v35 := GenS6{F0: -997}
            _ = v35
        }
    case -812:
        switch "\n\t" {
        case "":
            fmt.Println(v31, v31, v4)
        case "\n\t":
            if (2 == 32767) {
                var v36 string = ""
                _ = v36
            } else {
                var v37 int8 = v4
                _ = v37
            }
        case "hello":
            var v38 int64 = 0
            _ = v38
        }
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        for fi40 := 0; fi40 < 2; fi40++ {
            v41 := GenS7{F0: ("\n\t" + "\n\t")}
            switch fi40 {
            case 128:
                if (!true) {
                    v3 += 2
                    var v42 uint = 208
                    if ((uint16(v3) != 32767) || true) {
                        v43 := GenS4{F0: int16(v5)}
                        if (true || false) {
                            fmt.Println(v5, v5, v31)
                            v31 = v42
                            if false {
                                delete(v1, fi40)
                            } else {
                                if (32767 <= 54) {
                                    fmt.Println(fi40)
                                    v44 := GenS4{F0: int16(v5)}
                                    _ = len(v1)
                                    fmt.Println(v31, v42)
                                    v45 := v1[fi40]
                                    fmt.Println(v5)
                                    _ = v44
                                    _ = v45
                                }
                            }
                        } else {
                            var v46 int32 = 1
                            _ = v46
                        }
                        _ = v43
                    } else {
                        var v47 int64 = 651
                        delete(v1, 0)
                        _ = v47
                    }
                    _ = v42
                }
                var v48 uint16 = 423
                if (-128 < (^(v4))) {
                    switch "" {
                    case "hello":
                        v4 = v4
                    case "foo bar":
                        var v49 uint = (^(v31))
                        _ = v49
                    default:
                        var v50 float32 = ((3.14159 - -2.25) * 1.5)
                        _ = v50
                    }
                } else {
                    if (((v5 >> ((772) & 63)) < (v5 << (((v31 / (((^(v31))) | 1))) & 63))) || true) {
                        var v51 int64 = 333
                        var v52 uint16 = 65535
                        v53 := GenS8{F0: 1, F1: 1000}
                        _ = v51
                        _ = v52
                        _ = v53
                    }
                    v5 = v5
                }
                _ = v48
            }
            _ = fi40
            _ = v41
        }
        v3 -= v3
        _ = fi39
    }
    fmt.Println(v31)
    fmt.Println(v4, v31)
    fmt.Println(v5, v5, v4)
    for fi54 := 0; fi54 < 4; fi54++ {
        if false {
            v31 = (v31 & 253)
        } else {
            fmt.Println(v31)
        }
        var v55 int = (fi54 & (fi54 ^ 32767))
        _ = fi54
        _ = v55
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
