package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint16; F1 int16; F2 int8 }
type GenS5 struct { F0 int8; F1 uint16 }
func genFn0(p0 int8, p1 uint16) uint64 {
    if true {
        fmt.Println(p1, p0)
        p1 = (^(p1))
        v6 := []int{2147483647}
        _ = v6
    } else {
        v7 := []int{2147483647}
        switch "" {
        case "hello":
            switch 1 {
            case 257:
                switch "" {
                case "":
                    for fi8 := 0; fi8 < 11; fi8++ {
                        for fi9 := 0; fi9 < 6; fi9++ {
                            switch 1 {
                            case 2:
                                if len(v7) > 0 {
                                    v7[0] = (-(fi9))
                                }
                            case -565:
                                switch ("foo bar" + ("a" + (("" + ("a" + ("\n\t" + "a"))) + ""))) {
                                case "a":
                                    switch 221 {
                                    case 324:
                                        var v10 uint32 = 0
                                        _ = v10
                                    case -266:
                                        fmt.Println(fi9, p0)
                                    case 173:
                                        switch "hello" {
                                        case "":
                                            fmt.Println(p1)
                                            var v11 uint32 = uint32(p1)
                                            _ = v11
                                        default:
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                        }
                                    }
                                    var v12 float32 = 0.0
                                    _ = v12
                                case "":
                                    switch (fi9 / (((fi8 + 212)) | 1)) {
                                    case 1:
                                        p0 = (^(p0))
                                    case 42:
                                        var v13 int8 = (p0 >> ((649) & 63))
                                        v7 = append(v7, (fi8 & fi8))
                                        _ = v13
                                    case -597:
                                        var v14 int32 = 42
                                        _ = v14
                                    default:
                                        var v15 uint = 4294967295
                                        _ = v15
                                    }
                                }
                            case 32767:
                                var v16 int32 = -399
                                v7 = append(v7, (fi9 << ((480) & 63)))
                                _ = v16
                            }
                            _ = fi9
                        }
                        _ = fi8
                    }
                }
            default:
                p1 = 930
            }
        case "foo bar":
            v17 := GenS1{F0: (^(p1))}
            _ = v17
        default:
            v7 = append(v7, -769)
        }
        _ = v7
    }
    var v18 float32 = (1e10 + 1e10)
    switch "" {
    case "foo bar":
        var v19 int = 1
        _ = v19
    case "":
        if ("" <= "foo bar") {
            var v20 int32 = 0
            _ = v20
        }
    default:
        for fi21 := 0; fi21 < 15; fi21++ {
            switch (fi21 | fi21) {
            case 4294967295:
                v22 := GenS2{F0: ("hello" + "hello"), F1: (fi21 >> ((65535) & 63)), F2: 4294967295}
                _ = v22
            case -678:
                p1 += (p1 >> ((2) & 63))
            case -882:
                p1 = (p1 & (p1 - (p1 / (((p1 << ((65535) & 63))) | 1))))
            default:
                for fi23 := 0; fi23 < 1; fi23++ {
                    v24 := []int{(fi23 / ((2147483647) | 1))}
                    if len(v24) > 0 {
                        v24[0] = (fi23 & 2147483647)
                    }
                    _ = fi23
                    _ = v24
                }
            }
            fmt.Println(fi21, fi21)
            v25 := GenS2{F0: "foo bar", F1: fi21, F2: 42}
            _ = fi21
            _ = v25
        }
    }
    p1 += (p1 % (((^(p1))) | 1))
    var v26 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = v18
    _ = v26
    return 42
}
func genFn1(p0 int64) string {
    v27 := GenS3{F0: 4294967295, F1: p0}
    for fi28 := 0; fi28 < 17; fi28++ {
        fmt.Println(fi28)
        var v29 uint = 2
        _ = fi28
        _ = v29
    }
    _ = p0
    _ = v27
    return "foo bar"
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{}
    var v3 int64 = -1
    v4 := make(map[string]bool)
    v5 := make(map[uint32]int)
    v3 += 42
    if (false && false) {
        v30 := GenS3{F0: 18446744073709551615, F1: (v3 - 0)}
        fmt.Println(v3)
        _ = v30
    }
    var v31 uint = 2
    v32 := v4["hello"]
    var v33 int = 42
    var v34 bool = (((v32 && (2147483648 > 73)) && v32) && (1 <= 42))
    v35 := GenS4{F0: 42, F1: -1, F2: -128}
    switch "\n\t" {
    case "hello":
        var v36 uint8 = 2
        _ = v36
    default:
        v34 = (true || (v34 || v34))
    }
    fmt.Println(v32, v33)
    if len(v2) > 0 {
        v2[0] = v33
    }
    var v37 bool = (32767 > -32768)
    v3 = v3
    fmt.Println(v31)
    var v38 int64 = v3
    v37 = v32
    if (65535 < 4294967295) {
        var v39 int = v33
        v40 := GenS1{F0: 173}
        v41 := GenS5{F0: -69, F1: 1}
        _ = v39
        _ = v40
        _ = v41
    } else {
        var v42 uint64 = 265
        _ = v42
    }
    var v43 uint8 = 105
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v5
}
