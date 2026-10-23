package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 bool }
type GenS2 struct { F0 uint64; F1 int; F2 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 int32; F2 int16 }
type GenS5 struct { F0 int16; F1 int16; F2 uint8 }
type GenS6 struct { F0 int16; F1 int16 }
type GenS7 struct { F0 uint32; F1 int16 }

func main() {
    var v1 int64 = 42
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    var v4 int64 = v1
    v5 := GenS1{F0: 79, F1: 2, F2: false}
    switch 9223372036854775807 {
    case 127:
        switch ("foo bar" + (("" + ("foo bar" + "a")) + "")) {
        case "":
            for fi6 := 0; fi6 < 17; fi6++ {
                switch "hello" {
                case "a":
                    v7 := []int{}
                    _ = v7
                case "":
                    v8 := v2[fi6]
                    _ = v8
                default:
                    if ((1e10 > (0.1 * ((0.1 - 0.0) / 1.5))) || false) {
                        for fi9 := 0; fi9 < 18; fi9++ {
                            if (!(1 < int16(v1))) {
                                v10 := GenS2{F0: 710, F1: (fi9 << ((2) & 63)), F2: true}
                                var v11 int = (fi9 - (fi9 + fi9))
                                fmt.Println(v4)
                                _ = v10
                                _ = v11
                            }
                            _ = fi9
                        }
                        v12 := []int{(fi6 << ((18446744073709551615) & 63)), fi6, fi6}
                        _ = v12
                    } else {
                        for fi13 := 0; fi13 < 6; fi13++ {
                            if (635 >= 890) {
                                var v14 int32 = 1
                                switch "foo bar" {
                                case "\n\t":
                                    v15 := GenS2{F0: 4294967295, F1: (fi13 / ((1) | 1)), F2: true}
                                    _ = v15
                                case "a":
                                    v16 := v2[-9223372036854775808]
                                    _ = v16
                                }
                                switch ("foo bar" + "") {
                                case "":
                                    var v17 float64 = 0.1
                                    _ = v17
                                case "foo bar":
                                    if (466 != 612) {
                                        v4 += -1
                                    }
                                default:
                                    if true {
                                        var v18 bool = true
                                        v19 := GenS3{F0: 154}
                                        _ = v18
                                        _ = v19
                                    }
                                }
                                _ = v14
                            } else {
                                delete(v2, (fi13 / ((1) | 1)))
                                switch (fi13 % ((fi13) | 1)) {
                                case 2147483647:
                                    if (392 > 42) {
                                        fmt.Println(fi13, v4)
                                        if ((fi13 % ((fi13) | 1)) < fi6) {
                                            v1 -= (v4 + (v4 >> ((758) & 63)))
                                            if (!false) {
                                                v20 := GenS4{F0: 0, F1: int32(fi13), F2: 1}
                                                var v21 uint64 = 1
                                                v3[("" + ("a" + ""))] = fi13
                                                var v22 int16 = 0
                                                switch (fi6 << ((20) & 63)) {
                                                case 42:
                                                    if ("" == "hello") {
                                                        var v23 uint8 = 128
                                                        var v24 float64 = 3.14159
                                                        _ = v23
                                                        _ = v24
                                                    } else {
                                                        v2[-2147483648] = "\n\t"
                                                        var v25 float32 = 3.14159
                                                        v4 = (-(v4))
                                                        _ = v25
                                                    }
                                                case 1:
                                                    var v26 int64 = v1
                                                    _ = v26
                                                }
                                                _ = v20
                                                _ = v21
                                                _ = v22
                                            } else {
                                                if false {
                                                    v2[(-(fi6))] = ("hello" + "hello")
                                                } else {
                                                    var v27 uint = 2
                                                    var v28 int8 = -6
                                                    _ = v27
                                                    _ = v28
                                                }
                                                v1 = v1
                                            }
                                        } else {
                                            switch "\n\t" {
                                            case "a":
                                                v1 += v1
                                            }
                                        }
                                        v29 := GenS4{F0: 2147483647, F1: 0, F2: 621}
                                        _ = v29
                                    }
                                case 1:
                                    var v30 float64 = (3.14159 * 0.0)
                                    _ = v30
                                }
                                v2[(fi13 << ((65535) & 63))] = ""
                                v4 = v1
                                switch ((("" + "foo bar") + (("" + "a") + ("a" + "foo bar"))) + ("\n\t" + "")) {
                                case "":
                                    v31 := GenS5{F0: 1000, F1: 998, F2: 161}
                                    _ = v31
                                }
                            }
                            fmt.Println(fi6, fi6)
                            v32 := []int{(fi13 ^ (-(fi6))), fi13, fi13, fi6}
                            _ = fi13
                            _ = v32
                        }
                    }
                }
                v4 = v1
                _ = fi6
            }
        case "hello":
            var v33 string = "foo bar"
            _ = v33
        }
    case 32768:
        if false {
            var v34 uint64 = 4294967295
            if false {
                var v35 bool = false
                var v36 uint16 = 255
                v37 := []int{9223372036854775807, 442}
                _ = v35
                _ = v36
                _ = v37
            }
            _ = v34
        }
    case -559:
        var v38 uint64 = 1
        _ = v38
    default:
        v39 := GenS6{F0: 42, F1: -32768}
        _ = v39
    }
    if true {
        v1 = v4
        v4 = 0
    } else {
        if false {
            var v40 int32 = 109
            fmt.Println(v4)
            v1 = (v4 ^ (^(v1)))
            _ = v40
        } else {
            v41 := []int{612, -307, 42}
            v1 = v1
            if (-114 >= 1) {
                switch "foo bar" {
                case "\n\t":
                    for fi42 := 0; fi42 < 11; fi42++ {
                        v1 = v1
                        v43 := GenS6{F0: 32767, F1: -32768}
                        _ = fi42
                        _ = v43
                    }
                }
            } else {
                var v44 float64 = (3.14159 + 1.5)
                v45 := GenS4{F0: -791, F1: -1, F2: -32768}
                _ = v44
                _ = v45
            }
            _ = v41
        }
        var v46 float32 = 1e10
        for fi47 := 0; fi47 < 9; fi47++ {
            var v48 int32 = -1
            _ = fi47
            _ = v48
        }
        _ = v46
    }
    var v49 int8 = 1
    v50 := []int{0}
    var v51 bool = false
    v52 := GenS7{F0: 1, F1: 954}
    var v53 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
}
