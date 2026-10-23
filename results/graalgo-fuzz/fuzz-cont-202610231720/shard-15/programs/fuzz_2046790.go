package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 uint }
type GenS7 struct { F0 uint32 }
func genFn0(p0 uint64, p1 uint16) string {
    v4 := []int{-9223372036854775808, 9223372036854775807, 128}
    v5 := GenS1{F0: 4294967295, F1: 0, F2: 195}
    if len(v4) > 0 {
        v4[0] = 933
    }
    var v6 int32 = -249
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return "\n\t"
}
func genFn1() string {
    switch "" {
    case "hello":
        var v7 string = "hello"
        _ = v7
    case "a":
        var v8 uint8 = 34
        _ = v8
    case "\n\t":
        switch "" {
        case "a":
            var v9 uint = 2
            _ = v9
        default:
            v10 := []int{42, 9223372036854775807}
            _ = v10
        }
    }
    v11 := []int{32767, 32767}
    _ = v11
    return ""
}

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]int)
    var v3 uint32 = 500
    var v12 uint32 = 902
    v3 += v3
    v13 := GenS2{F0: 1, F1: false}
    fmt.Println(v1, v12, v12)
    var v14 int8 = 1
    v15 := []int{-1}
    if true {
        for fi16 := 0; fi16 < 4; fi16++ {
            _ = len(v2)
            var v17 int = (-(fi16))
            if (false && (1 != 255)) {
                switch (((("a" + "\n\t") + "hello") + (("\n\t" + (("a" + "a") + ("foo bar" + "\n\t"))) + ((("hello" + "a") + "") + "\n\t"))) + "foo bar") {
                case "\n\t":
                    if false {
                        v1 = v1
                        if ((!(!false)) && (255 != 2)) {
                            if (true && (!(true && true))) {
                                var v18 string = "foo bar"
                                for fi19 := 0; fi19 < 2; fi19++ {
                                    var v20 string = "a"
                                    _ = fi19
                                    _ = v20
                                }
                                if false {
                                    for fi21 := 0; fi21 < 2; fi21++ {
                                        var v22 uint = 366
                                        v15 = append(v15, 127)
                                        _ = fi21
                                        _ = v22
                                    }
                                    v14 += v14
                                    v18 = v18
                                }
                                _ = v18
                            }
                            switch "\n\t" {
                            case "hello":
                                var v23 uint32 = (v3 ^ (v3 ^ v12))
                                _ = v23
                            case "":
                                if true {
                                    v14 += v14
                                    var v24 int = v17
                                    if false {
                                        v12 = v3
                                        v15 = append(v15, (v17 % ((int(v1)) | 1)))
                                    } else {
                                        var v25 int16 = v1
                                        switch ("foo bar" + ("" + ("\n\t" + ""))) {
                                        case "":
                                            var v26 int64 = 49
                                            _ = v26
                                        case "\n\t":
                                            v27 := GenS3{F0: -1}
                                            _ = v27
                                        case "a":
                                            if (42 != 2) {
                                                v14 = (-(v14))
                                                _ = len(v2)
                                            } else {
                                                v28 := GenS3{F0: 0}
                                                switch ((("foo bar" + "") + "a") + "hello") {
                                                case "foo bar":
                                                    _ = len(v15)
                                                case "":
                                                    var v29 bool = true
                                                    _ = v29
                                                default:
                                                    v3 += 999
                                                }
                                                if true {
                                                    v30 := GenS3{F0: -3}
                                                    _ = v30
                                                } else {
                                                    var v31 int = 32767
                                                    v32 := GenS4{F0: "foo bar", F1: v3}
                                                    _ = v31
                                                    _ = v32
                                                }
                                                _ = v28
                                            }
                                        default:
                                            var v33 string = "hello"
                                            _ = v33
                                        }
                                        _ = v25
                                    }
                                    _ = v24
                                }
                            case "\n\t":
                                _ = len(v15)
                            }
                        }
                        if false {
                            v17 -= (v17 / ((v17) | 1))
                            for fi34 := 0; fi34 < 15; fi34++ {
                                v15 = append(v15, (fi34 << ((714) & 63)))
                                if true {
                                    if true {
                                        var v35 int16 = (v1 - v1)
                                        fmt.Println(v12, fi34)
                                        var v36 int = (^(v17))
                                        _ = v35
                                        _ = v36
                                    }
                                    v17 -= -682
                                }
                                var v37 uint16 = 0
                                v17 += int(v37)
                                v38 := v2[(v3 / ((v3) | 1))]
                                _ = fi34
                                _ = v37
                                _ = v38
                            }
                        }
                    } else {
                        switch "hello" {
                        case "hello":
                            if (602 == 2) {
                                fmt.Println(fi16, v14)
                                _ = len(v15)
                                v17 = (fi16 & (^(v17)))
                            }
                        case "":
                            fmt.Println(fi16, v3)
                        case "\n\t":
                            v39 := GenS5{F0: -603}
                            _ = v39
                        }
                    }
                case "":
                    v14 = v14
                default:
                    var v40 int = (-(v17))
                    _ = v40
                }
                switch "a" {
                case "foo bar":
                    var v41 uint8 = 0
                    _ = v41
                case "hello":
                    v42 := GenS5{F0: 32768}
                    _ = v42
                }
                if (2 > 274) {
                    v1 = (v1 * v1)
                    for fi43 := 0; fi43 < 12; fi43++ {
                        if (false || false) {
                            v44 := v2[(v3 % ((1) | 1))]
                            var v45 uint16 = 0
                            v46 := GenS6{F0: (v45 << ((255) & 63)), F1: 18446744073709551615}
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                        _ = fi43
                    }
                }
            }
            _ = fi16
            _ = v17
        }
        if (0 != 812) {
            v3 += (v3 & v3)
        }
    }
    if (65535 != 65535) {
        v14 = v14
        v47 := GenS7{F0: v3}
        var v48 int16 = (v1 ^ (^(v1)))
        _ = v47
        _ = v48
    } else {
        for fi49 := 0; fi49 < 18; fi49++ {
            var v50 uint = 0
            _ = fi49
            _ = v50
        }
    }
    var v51 bool = ((!(!((v1 ^ v1) <= v1))) || false)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v51
}
