package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int32; F1 string; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint32) int32 {
    if true {
        v4 := GenS1{F0: 146, F1: 291}
        fmt.Println(p0)
        _ = v4
    }
    p0 = 345
    if ("hello" > "") {
        switch 32768 {
        case 32767:
            v5 := GenS1{F0: 128, F1: 4294967295}
            _ = v5
        case -9223372036854775808:
            var v6 float32 = 1.5
            _ = v6
        }
    } else {
        switch -9223372036854775808 {
        case 172:
            switch ("foo bar" + "foo bar") {
            case "":
                if true {
                    var v7 uint32 = 65535
                    fmt.Println(p0)
                    fmt.Println(p0, p0)
                    _ = v7
                }
            case "\n\t":
                var v8 uint16 = 42
                _ = v8
            }
        case 0:
            for fi9 := 0; fi9 < 3; fi9++ {
                for fi10 := 0; fi10 < 18; fi10++ {
                    switch fi9 {
                    case 0:
                        switch "hello" {
                        case "":
                            fmt.Println(fi9)
                        case "\n\t":
                            v11 := GenS1{F0: 255, F1: 2}
                            _ = v11
                        case "hello":
                            p0 = p0
                        }
                    }
                    if false {
                        if false {
                            p0 = 661
                            var v12 int16 = 42
                            fmt.Println(fi10)
                            v13 := []int{(fi10 + (fi10 | -2147483648))}
                            _ = v12
                            _ = v13
                        }
                        fmt.Println(p0)
                        fmt.Println(p0, p0)
                        v14 := GenS2{F0: -936, F1: "\n\t", F2: 18446744073709551615}
                        v15 := GenS1{F0: 150, F1: 584}
                        _ = v14
                        _ = v15
                    } else {
                        if (((4294967295 > 967) || true) && true) {
                            var v16 uint32 = (p0 / ((0) | 1))
                            v17 := GenS1{F0: 0, F1: 1}
                            p0 = v16
                            _ = v16
                            _ = v17
                        } else {
                            p0 = 0
                            var v18 int = fi9
                            switch 174 {
                            case 127:
                                if ((0.0 - 1e10) <= (0.1 * 1e10)) {
                                    var v19 int64 = 42
                                    fmt.Println(p0, v18)
                                    _ = v19
                                } else {
                                    if false {
                                        v20 := GenS3{F0: 4294967295, F1: 32767}
                                        switch (^(v18)) {
                                        case 32768:
                                            v21 := []int{fi10, -513}
                                            _ = v21
                                        case 0:
                                            var v22 int16 = 1000
                                            _ = v22
                                        case -912:
                                            v23 := GenS1{F0: 128, F1: 255}
                                            _ = v23
                                        default:
                                            v24 := []int{(fi9 & fi10), (fi10 - fi10), (^(fi9))}
                                            _ = v24
                                        }
                                        var v25 uint64 = 109
                                        _ = v20
                                        _ = v25
                                    }
                                    switch (fi9 - fi9) {
                                    case 128:
                                        fmt.Println(v18, p0)
                                    case 936:
                                        v26 := []int{42}
                                        _ = v26
                                    case 1:
                                        p0 += p0
                                    }
                                }
                            default:
                                p0 += p0
                            }
                            v18 += (fi10 << ((2) & 63))
                            if (false && true) {
                                var v27 uint32 = (p0 >> ((1) & 63))
                                _ = v27
                            }
                            _ = v18
                        }
                    }
                    _ = fi10
                }
                for fi28 := 0; fi28 < 6; fi28++ {
                    var v29 uint32 = p0
                    v29 += (^(v29))
                    var v30 float64 = (0.1 - ((float64(1.5) + -2.25) + (1e10 + float64(0.0))))
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = fi9
            }
        }
    }
    var v31 uint64 = 147
    var v32 float32 = -2.25
    _ = p0
    _ = v31
    _ = v32
    return -2147483648
}
func genFn1(p0 bool, p1 uint64) int16 {
    if p0 {
        fmt.Println(p1, p1)
        switch "a" {
        case "":
            switch ("hello" + ("" + ((("" + "hello") + "") + "foo bar"))) {
            case "\n\t":
                var v33 float64 = 1e10
                _ = v33
            case "hello":
                v34 := []int{128, 2147483647, 2147483647}
                _ = v34
            default:
                var v35 uint64 = 42
                _ = v35
            }
        case "hello":
            for fi36 := 0; fi36 < 6; fi36++ {
                for fi37 := 0; fi37 < 16; fi37++ {
                    var v38 int32 = genFn0(2147483648)
                    var v39 float64 = 1e10
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                _ = fi36
            }
        }
    } else {
        p1 -= (p1 << ((2) & 63))
        var v40 uint8 = 255
        _ = v40
    }
    var v41 int32 = 2147483647
    var v42 uint64 = (p1 | (p1 ^ 194))
    p0 = false
    _ = p0
    _ = p1
    _ = v41
    _ = v42
    return -1
}

func main() {
    var v1 uint = 530
    v2 := make(map[uint32]int)
    v3 := []string{"hello", ("foo bar" + ("foo bar" + ("" + "hello")))}
    if true {
        var v43 string = (("foo bar" + "\n\t") + "")
        switch 876 {
        case 32768:
            v1 = 0
        case -1:
            var v44 uint8 = 92
            _ = v44
        }
        _ = v43
    } else {
        var v45 uint8 = 55
        var v46 int8 = 127
        _ = v45
        _ = v46
    }
    var v47 int16 = 1
    switch "hello" {
    case "":
        fmt.Println(v1, v1)
    case "foo bar":
        if true {
            _ = len(v2)
            if (true && true) {
                fmt.Println(v47, v1)
                v47 = -820
                v1 = v1
            } else {
                v47 = -32768
                var v48 bool = true
                switch 2147483647 {
                case -1:
                    v2[42] = -2147483648
                case 42:
                    var v49 uint = (v1 & v1)
                    _ = v49
                case 32767:
                    _ = len(v3)
                }
                _ = v48
            }
        } else {
            var v50 string = ("\n\t" + "a")
            for fi51 := 0; fi51 < 8; fi51++ {
                var v52 uint16 = 375
                fmt.Println(fi51, v1)
                for fi53 := 0; fi53 < 12; fi53++ {
                    if false {
                        var v54 int8 = -1
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v55 uint8 = 75
                        fmt.Println(fi51, v55)
                        _ = v54
                        _ = v55
                    } else {
                        v47 += (v47 * genFn1((-9223372036854775808 > 9223372036854775807), 149))
                        if (!false) {
                            if false {
                                delete(v2, 373)
                                if ((1e10 + 1e10) != -2.25) {
                                    var v56 uint = uint(v52)
                                    _ = v56
                                }
                            }
                            v1 = 18446744073709551615
                            v52 += v52
                        }
                    }
                    _ = fi53
                }
                _ = fi51
                _ = v52
            }
            var v57 int32 = genFn0(2)
            _ = v50
            _ = v57
        }
    case "a":
        var v58 uint32 = 42
        _ = v58
    default:
        v59 := GenS4{F0: false, F1: 232}
        _ = v59
    }
    v47 -= 585
    for fi60 := 0; fi60 < 4; fi60++ {
        v61 := GenS4{F0: true, F1: 127}
        delete(v2, 2147483648)
        v1 = v1
        _ = fi60
        _ = v61
    }
    var v62 string = "foo bar"
    switch "foo bar" {
    case "a":
        delete(v2, 1)
    }
    for fi63 := 0; fi63 < 1; fi63++ {
        var v64 float32 = -2.25
        _ = fi63
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v47:", v47)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
    _ = v62
}
