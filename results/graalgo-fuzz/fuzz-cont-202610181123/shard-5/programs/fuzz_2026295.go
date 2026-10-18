package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int64; F1 int64; F2 string }
type GenS3 struct { F0 int32; F1 int32; F2 int16 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 uint8 }
func genFn0(p0 int16, p1 uint8, p2 string) int {
    v4 := GenS1{F0: 643, F1: 255, F2: 2147483647}
    v5 := []int{0, 127, 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 128
}

func main() {
    var v1 uint8 = 42
    v2 := []string{"", "foo bar", ("hello" + "foo bar"), ("\n\t" + "a")}
    v3 := make(map[string]bool)
    if false {
        var v6 int16 = 32767
        v3["foo bar"] = true
        _ = v6
    }
    v3["\n\t"] = (v1 >= (^(v1)))
    switch (("foo bar" + ("hello" + "hello")) + "a") {
    case "\n\t":
        for fi7 := 0; fi7 < 3; fi7++ {
            fmt.Println(fi7, fi7)
            v3["a"] = true
            _ = len(v2)
            _ = fi7
        }
    case "foo bar":
        fmt.Println(v1)
    default:
        var v8 int8 = 70
        _ = v8
    }
    _ = len(v2)
    fmt.Println(v1)
    var v9 int64 = -1
    v10 := GenS2{F0: 9223372036854775807, F1: 42, F2: "hello"}
    if ((true && (-128 > 127)) || false) {
        v11 := GenS2{F0: (v9 * v9), F1: (v9 >> ((42) & 63)), F2: ((("" + "foo bar") + "hello") + "a")}
        v12 := GenS2{F0: (-(v9)), F1: (v9 | v9), F2: ((("\n\t" + "a") + ("hello" + "\n\t")) + "foo bar")}
        switch 9223372036854775807 {
        case 4294967295:
            for fi13 := 0; fi13 < 15; fi13++ {
                _ = len(v3)
                var v14 float64 = ((float64((-2.25 - -2.25)) * float64(3.14159)) + -2.25)
                _ = fi13
                _ = v14
            }
        case -832:
            _ = len(v3)
        case 213:
            fmt.Println(v1, v9)
        default:
            var v15 uint32 = 4294967295
            _ = v15
        }
        _ = v11
        _ = v12
    }
    var v16 uint16 = 884
    v9 = (v9 & v9)
    fmt.Println(v1, v16)
    var v17 float64 = -2.25
    var v18 uint8 = 0
    v19 := GenS3{F0: 42, F1: 117, F2: 1000}
    var v20 int16 = 261
    if (!false) {
        v20 = int16(v16)
        var v21 uint16 = v16
        if true {
            v18 += (v1 % ((v1) | 1))
            v21 = (v21 << ((18446744073709551615) & 63))
            if true {
                v22 := GenS4{F0: (false || true), F1: (v9 % (((v9 % ((-591) | 1))) | 1))}
                _ = v22
            }
        } else {
            for fi23 := 0; fi23 < 16; fi23++ {
                if (0 != 42) {
                    var v24 uint32 = 1
                    if (-123 <= -122) {
                        var v25 uint16 = (^(v21))
                        for fi26 := 0; fi26 < 18; fi26++ {
                            switch 42 {
                            case -2147483648:
                                fmt.Println(v21)
                            case -606:
                                switch int(v20) {
                                case 237:
                                    if (false && ((v20 << ((0) & 63)) == -583)) {
                                        v9 += 42
                                        var v27 uint16 = (v25 % ((42) | 1))
                                        fmt.Println(v20, v20, v18)
                                        _ = v27
                                    }
                                case -2147483648:
                                    v17 -= v17
                                }
                            case 4294967295:
                                var v28 int32 = 154
                                _ = v28
                            default:
                                v29 := GenS1{F0: (v24 << ((18446744073709551615) & 63)), F1: (v25 % ((v16) | 1)), F2: (v24 ^ v24)}
                                _ = v29
                            }
                            _ = fi26
                        }
                        fmt.Println(fi23)
                        _ = v25
                    }
                    switch "foo bar" {
                    case "\n\t":
                        v16 -= (v16 - (v16 / (((v21 & (v16 / ((65535) | 1)))) | 1)))
                    case "foo bar":
                        var v30 float64 = 0.0
                        _ = v30
                    case "hello":
                        switch fi23 {
                        case 32768:
                            fmt.Println(fi23, v20, v1)
                        case 42:
                            var v31 int8 = 100
                            _ = v31
                        default:
                            if true {
                                switch (fi23 + 42) {
                                case 2147483647:
                                    v24 = v24
                                case 440:
                                    var v32 string = "hello"
                                    _ = v32
                                case 32767:
                                    if (-505 <= (-(v9))) {
                                        if ((genFn0(v20, (v1 >> ((255) & 63)), "hello") != 24) && ((v21 != (^(v16))) && false)) {
                                            for fi33 := 0; fi33 < 13; fi33++ {
                                                switch "" {
                                                case "a":
                                                    switch ((((("\n\t" + ("" + "foo bar")) + "a") + "\n\t") + "\n\t") + "\n\t") {
                                                    case "a":
                                                        switch ((("a" + (("" + ("\n\t" + (("foo bar" + "a") + ("a" + "\n\t")))) + "a")) + "a") + "foo bar") {
                                                        case "":
                                                            if len(v2) > 0 {
                                                                v2[0] = "foo bar"
                                                            }
                                                        case "\n\t":
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        default:
                                                            if (2147483647 != 1) {
                                                                fmt.Println(v1, v1)
                                                            } else {
                                                                var v34 uint = 42
                                                                if false {
                                                                    switch fi33 {
                                                                    case 658:
                                                                        var v35 bool = false
                                                                        _ = v35
                                                                    }
                                                                    v2 = append(v2, "")
                                                                }
                                                                var v36 int64 = 9
                                                                v9 -= (v9 * (v9 >> (((v34 - v34)) & 63)))
                                                                _ = v34
                                                                _ = v36
                                                            }
                                                        }
                                                    }
                                                }
                                                v37 := GenS2{F0: (v9 ^ v9), F1: int64(v1), F2: "\n\t"}
                                                v38 := GenS1{F0: 0, F1: 255, F2: (^(v24))}
                                                switch "hello" {
                                                case "hello":
                                                    var v39 float64 = (float64(-2.25) * (0.1 - v17))
                                                    var v40 int64 = (v9 | (v9 + v9))
                                                    v24 -= (v24 ^ (v24 | (v24 << ((18446744073709551615) & 63))))
                                                    _ = v39
                                                    _ = v40
                                                case "\n\t":
                                                    v18 = v1
                                                }
                                                _ = fi33
                                                _ = v37
                                                _ = v38
                                            }
                                        } else {
                                            switch (^(fi23)) {
                                            case 4294967295:
                                                v1 = v1
                                            case -1:
                                                var v41 int64 = -9223372036854775808
                                                _ = v41
                                            default:
                                                v42 := GenS5{F0: fi23}
                                                _ = v42
                                            }
                                            var v43 int16 = v20
                                            if len(v2) > 0 {
                                                v2[0] = "\n\t"
                                            }
                                            _ = v43
                                        }
                                        var v44 float32 = float32(v17)
                                        _ = v44
                                    }
                                default:
                                    for fi45 := 0; fi45 < 19; fi45++ {
                                        v17 += 0.0
                                        v2 = append(v2, "\n\t")
                                        var v46 int64 = (v9 * (v9 | (^(v9))))
                                        _ = fi45
                                        _ = v46
                                    }
                                }
                            } else {
                                v47 := GenS5{F0: genFn0(v20, v18, "")}
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v48 := GenS6{F0: true, F1: false}
                                _ = v47
                                _ = v48
                            }
                        }
                    }
                    _ = v24
                } else {
                    if false {
                        var v49 uint32 = 711
                        v50 := GenS6{F0: false, F1: false}
                        _ = v49
                        _ = v50
                    }
                }
                _ = fi23
            }
        }
        _ = v21
    }
    v51 := GenS5{F0: 4294967295}
    switch 2147483647 {
    case -653:
        v52 := GenS7{F0: v1}
        _ = v52
    case 9223372036854775807:
        var v53 uint16 = v16
        _ = v53
    case 293:
        v9 = (-(v9))
    default:
        for fi54 := 0; fi54 < 20; fi54++ {
            v9 += v9
            _ = fi54
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v51
    _ = v9
}
