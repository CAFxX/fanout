package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint16; F1 int8; F2 string }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint) int {
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 int = fi6
        var v8 float64 = 1.5
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 float32 = (0.1 / 1.5)
    p0 = (p0 / ((p0) | 1))
    for fi10 := 0; fi10 < 13; fi10++ {
        switch "foo bar" {
        case "\n\t":
            for fi11 := 0; fi11 < 20; fi11++ {
                var v12 int64 = 42
                p0 -= (p0 >> ((p0) & 63))
                if false {
                    p0 -= 350
                    v12 -= v12
                } else {
                    v9 = v9
                }
                _ = fi11
                _ = v12
            }
        case "":
            if false {
                v13 := []int{fi10, (fi10 / ((4294967295) | 1)), fi10, fi10}
                var v14 float64 = 1.5
                var v15 uint = (p0 / ((p0) | 1))
                _ = v13
                _ = v14
                _ = v15
            }
        }
        p0 = (p0 % ((2) | 1))
        _ = fi10
    }
    _ = p0
    _ = v9
    return -1
}
func genFn1(p0 int, p1 int64, p2 bool) string {
    for fi16 := 0; fi16 < 12; fi16++ {
        switch ("a" + "\n\t") {
        case "":
            switch "foo bar" {
            case "\n\t":
                var v17 string = "foo bar"
                _ = v17
            case "a":
                p2 = p2
            default:
                if true {
                    p1 -= p1
                    var v18 int = p0
                    _ = v18
                }
            }
        case "foo bar":
            p1 = p1
        default:
            switch fi16 {
            case -553:
                var v19 int32 = 1
                _ = v19
            default:
                if false {
                    if (true || true) {
                        p0 += p0
                        p1 = p1
                        switch ("a" + "\n\t") {
                        case "hello":
                            for fi20 := 0; fi20 < 16; fi20++ {
                                var v21 uint16 = 730
                                var v22 int8 = 16
                                switch ("foo bar" + "a") {
                                case "a":
                                    p1 -= (p1 + p1)
                                default:
                                    switch (fi16 + (-(fi20))) {
                                    case -144:
                                        v21 = v21
                                    case 515:
                                        var v23 int16 = 846
                                        _ = v23
                                    }
                                }
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        case "a":
                            switch "\n\t" {
                            case "a":
                                if p2 {
                                    p0 = p0
                                } else {
                                    v24 := []int{p0, fi16, genFn0(941), fi16}
                                    switch (^(fi16)) {
                                    case 127:
                                        var v25 uint = 255
                                        _ = v25
                                    default:
                                        switch p0 {
                                        case 128:
                                            fmt.Println(p1, p0, p2)
                                        }
                                    }
                                    _ = v24
                                }
                            case "foo bar":
                                p1 = p1
                            case "":
                                v26 := GenS1{F0: 42, F1: 1, F2: p1}
                                _ = v26
                            default:
                                var v27 uint64 = 0
                                _ = v27
                            }
                        }
                    } else {
                        var v28 float64 = 3.14159
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 2; fi29++ {
                        if (3.14159 > float64(1.5)) {
                            var v30 int = (fi16 | (p0 << ((65535) & 63)))
                            switch "" {
                            case "":
                                if (65535 != 2) {
                                    v31 := GenS1{F0: 255, F1: 444, F2: (p1 ^ (p1 - 9223372036854775807))}
                                    switch fi16 {
                                    case 0:
                                        v32 := GenS2{F0: -1}
                                        _ = v32
                                    case 4294967295:
                                        var v33 uint32 = 2147483647
                                        _ = v33
                                    }
                                    p1 = p1
                                    _ = v31
                                } else {
                                    if p2 {
                                        switch "a" {
                                        case "foo bar":
                                            var v34 uint32 = 2147483647
                                            _ = v34
                                        case "\n\t":
                                            v35 := []int{-123}
                                            _ = v35
                                        default:
                                            p0 = v30
                                        }
                                        fmt.Println(v30, fi29)
                                        if p2 {
                                            var v36 uint32 = 998
                                            _ = v36
                                        } else {
                                            v30 += (fi16 ^ (fi29 | genFn0(255)))
                                            var v37 int8 = -128
                                            _ = v37
                                        }
                                    } else {
                                        var v38 int32 = 1
                                        v39 := GenS2{F0: -124}
                                        var v40 float32 = 0.1
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    }
                                    v41 := GenS3{F0: 0, F1: 9223372036854775807}
                                    _ = v41
                                }
                            default:
                                fmt.Println(fi16, p2, v30)
                            }
                            fmt.Println(v30, p1, fi16)
                            var v42 int8 = 1
                            _ = v30
                            _ = v42
                        } else {
                            var v43 int64 = (p1 | (p1 ^ 9223372036854775807))
                            var v44 uint32 = 2147483648
                            v45 := GenS1{F0: 0, F1: 65535, F2: p1}
                            if ((^(fi29)) == genFn0(uint(v44))) {
                                v43 = (p1 | v43)
                            }
                            v44 = (^(v44))
                            var v46 uint64 = 0
                            _ = v43
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                        var v47 uint32 = 543
                        _ = fi29
                        _ = v47
                    }
                }
            }
        }
        var v48 int = (fi16 / (((fi16 % (((fi16 * 32768)) | 1))) | 1))
        _ = fi16
        _ = v48
    }
    p1 += p1
    if p2 {
        p1 = p1
        var v49 uint64 = 4294967295
        fmt.Println(p2)
        _ = v49
    } else {
        var v50 uint64 = 460
        var v51 uint8 = 181
        _ = v50
        _ = v51
    }
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}

func main() {
    var v1 bool = (1 == 861)
    v2 := make(map[int]int)
    var v3 float32 = 0.1
    var v4 bool = (32767 > -986)
    v5 := []uint32{2147483647, 2147483647, 1}
    v1 = true
    delete(v2, -366)
    v5 = append(v5, 42)
    _ = len(v5)
    v52 := GenS1{F0: 1, F1: 255, F2: -1}
    v1 = v1
    v1 = v4
    if v4 {
        var v53 string = ""
        v54 := GenS3{F0: 289, F1: 0}
        _ = v53
        _ = v54
    }
    v3 += (1.5 + (v3 - (v3 / v3)))
    v4 = v4
    if true {
        if true {
            v1 = (!(42 <= 42))
            v55 := GenS4{F0: 255, F1: 122, F2: "\n\t"}
            v4 = true
            _ = v55
        } else {
            for fi56 := 0; fi56 < 5; fi56++ {
                v1 = v1
                _ = fi56
            }
        }
        delete(v2, 127)
    }
    var v57 uint64 = 338
    switch ("a" + "\n\t") {
    case "hello":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        v57 = 0
    }
    if false {
        if ((1 == 843) || false) {
            v58 := GenS5{F0: 65535}
            v3 -= 1e10
            v59 := GenS3{F0: 2147483648, F1: int64(v57)}
            _ = v58
            _ = v59
        } else {
            var v60 uint = 163
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v60
        }
        v61 := GenS3{F0: 42, F1: -735}
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v57
}
