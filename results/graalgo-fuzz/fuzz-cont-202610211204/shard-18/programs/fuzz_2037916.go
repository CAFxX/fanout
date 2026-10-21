package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 uint16; F1 uint16; F2 int8 }
func genFn0(p0 int8) int64 {
    v5 := GenS1{F0: 635}
    var v6 float32 = (0.0 / float32(0.1))
    _ = p0
    _ = v5
    _ = v6
    return -590
}

func main() {
    var v1 bool = false
    var v2 int8 = 42
    v3 := []string{"hello", "foo bar", ((("hello" + "hello") + ("" + ("\n\t" + "foo bar"))) + ("hello" + ""))}
    v4 := make(map[uint32]uint32)
    v2 = v2
    if (0 <= 128) {
        var v7 bool = (float64((-2.25 + (-2.25 / 0.1))) == 1.5)
        _ = v7
    }
    var v8 int8 = (v2 >> ((4294967295) & 63))
    v1 = v1
    if v1 {
        if false {
            var v9 int8 = 100
            _ = v9
        }
        v10 := GenS2{F0: ("a" + "")}
        fmt.Println(v1, v2)
        _ = v10
    } else {
        switch "hello" {
        case "hello":
            var v11 uint64 = 1
            _ = v11
        case "\n\t":
            for fi12 := 0; fi12 < 14; fi12++ {
                v8 -= (-(v2))
                _ = fi12
            }
        }
        v8 += (-(v2))
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 uint32 = 2
            switch "hello" {
            case "":
                if false {
                    if ((v1 || v1) && v1) {
                        v14 += v14
                        fmt.Println(v8, v8)
                        switch (fi13 + (fi13 / (((fi13 & fi13)) | 1))) {
                        case 4294967295:
                            v15 := GenS2{F0: "a"}
                            _ = v15
                        case 42:
                            fmt.Println(v8, v8, v8)
                        }
                    }
                    v4[2] = v14
                } else {
                    v16 := GenS2{F0: ("a" + ("hello" + "\n\t"))}
                    v17 := GenS2{F0: "\n\t"}
                    _ = v16
                    _ = v17
                }
            case "hello":
                if true {
                    fmt.Println(v2, fi13)
                    v2 = (v2 / ((-128) | 1))
                    v14 += 930
                } else {
                    fmt.Println(v8, v14, v2)
                    v18 := GenS3{F0: -1}
                    if true {
                        switch 127 {
                        case 1:
                            v2 -= -128
                        case -2147483648:
                            switch "\n\t" {
                            case "foo bar":
                                v8 += (v8 % ((0) | 1))
                            case "a":
                                v2 = -128
                            default:
                                v19 := GenS3{F0: -2147483648}
                                _ = v19
                            }
                        case 683:
                            v20 := GenS2{F0: ""}
                            _ = v20
                        }
                    } else {
                        if true {
                            var v21 int8 = v8
                            var v22 uint8 = 0
                            v14 = 820
                            _ = v21
                            _ = v22
                        } else {
                            v14 -= 163
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(v14, v8, fi13)
                        }
                    }
                    _ = v18
                }
            case "a":
                var v23 int16 = -153
                _ = v23
            default:
                v8 = v8
            }
            v24 := GenS4{F0: 32767, F1: "a"}
            _ = fi13
            _ = v14
            _ = v24
        }
    }
    v8 -= v2
    for fi25 := 0; fi25 < 7; fi25++ {
        switch 32767 {
        case 32768:
            switch "" {
            case "foo bar":
                v2 = v2
            default:
                v26 := GenS1{F0: 873}
                _ = v26
            }
        case 1:
            v8 += (^(v2))
        }
        _ = len(v4)
        _ = fi25
    }
    v27 := GenS1{F0: 2}
    switch 1 {
    case -9223372036854775808:
        if (true || ((0.1 < 0.0) || (v2 == (v2 + -128)))) {
            for fi28 := 0; fi28 < 5; fi28++ {
                if (-1 <= 0) {
                    switch "" {
                    case "\n\t":
                        v8 = v2
                    case "a":
                        if (v1 && v1) {
                            for fi29 := 0; fi29 < 9; fi29++ {
                                var v30 uint16 = 1
                                var v31 int = fi28
                                switch fi29 {
                                case 9223372036854775807:
                                    var v32 bool = false
                                    _ = v32
                                case -1:
                                    switch "foo bar" {
                                    case "\n\t":
                                        v31 += fi28
                                    }
                                }
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                            if len(v3) > 0 {
                                v3[0] = ((("foo bar" + (("foo bar" + "foo bar") + "")) + "") + "a")
                            }
                            v8 -= 100
                        }
                    case "":
                        fmt.Println(v8, v1, fi28)
                    }
                }
                switch "" {
                case "hello":
                    v8 += (v2 >> ((18446744073709551615) & 63))
                }
                _ = fi28
            }
        } else {
            for fi33 := 0; fi33 < 8; fi33++ {
                var v34 string = "a"
                if len(v3) > 0 {
                    v3[0] = ("a" + (v34 + (v34 + ("foo bar" + "\n\t"))))
                }
                v35 := GenS5{F0: v34, F1: 42}
                _ = fi33
                _ = v34
                _ = v35
            }
        }
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v2, v2, v8)
    case "\n\t":
        v36 := GenS5{F0: "\n\t", F1: 42}
        _ = v36
    case "hello":
        for fi37 := 0; fi37 < 6; fi37++ {
            if (("foo bar" < "hello") || (!v1)) {
                fmt.Println(v8)
            }
            var v38 int16 = -32768
            if (1 <= 174) {
                if (1e10 == ((0.0 * float32(-2.25)) * 0.0)) {
                    switch "" {
                    case "hello":
                        v1 = true
                    case "foo bar":
                        v1 = ((!true) && (2 < 18446744073709551615))
                    case "a":
                        var v39 int = fi37
                        _ = v39
                    }
                } else {
                    var v40 int64 = 927
                    _ = v40
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = (("\n\t" + "") + "")
                }
                _ = len(v3)
            }
            _ = fi37
            _ = v38
        }
    default:
        v3 = append(v3, "\n\t")
    }
    switch "" {
    case "\n\t":
        if (0 > 18446744073709551615) {
            var v41 uint64 = 2
            for fi42 := 0; fi42 < 12; fi42++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v43 := GenS2{F0: ""}
                _ = fi42
                _ = v43
            }
            _ = v41
        } else {
            delete(v4, 0)
            v44 := GenS2{F0: "hello"}
            _ = v44
        }
    case "hello":
        v45 := GenS3{F0: -1}
        _ = v45
    default:
        var v46 int16 = -346
        _ = v46
    }
    _ = len(v4)
    v47 := GenS6{F0: 42, F1: 42, F2: v2}
    var v48 uint8 = 128
    if (!false) {
        v8 = -128
    } else {
        var v49 uint = 196
        v50 := GenS3{F0: 2147483647}
        fmt.Println(v2, v1)
        _ = v49
        _ = v50
    }
    fmt.Println(v2)
    switch "hello" {
    case "a":
        v1 = false
    default:
        for fi51 := 0; fi51 < 17; fi51++ {
            v1 = v1
            _ = fi51
        }
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v8
}
