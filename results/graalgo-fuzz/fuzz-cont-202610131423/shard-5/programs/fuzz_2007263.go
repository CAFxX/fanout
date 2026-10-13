package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 uint16; F2 uint32 }
type GenS6 struct { F0 bool; F1 uint64; F2 int64 }
type GenS7 struct { F0 int32; F1 int; F2 uint16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 bool }
func genFn0(p0 uint8) uint64 {
    fmt.Println(p0)
    var v4 int = 127
    _ = p0
    _ = v4
    return 189
}
func genFn1(p0 uint16, p1 int64) bool {
    if true {
        var v5 string = "a"
        v5 = (v5 + "foo bar")
        if true {
            for fi6 := 0; fi6 < 1; fi6++ {
                p0 -= p0
                var v7 int16 = 0
                _ = fi6
                _ = v7
            }
        } else {
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 int16 = 1
                if false {
                    fmt.Println(v9, v5, p0)
                    for fi10 := 0; fi10 < 8; fi10++ {
                        v9 = (v9 & v9)
                        v11 := []int{(fi10 ^ int(p0)), 4294967295}
                        _ = fi10
                        _ = v11
                    }
                } else {
                    fmt.Println(p0)
                    v9 = (v9 * (v9 | (^(v9))))
                    v9 -= 32767
                }
                _ = fi8
                _ = v9
            }
            for fi12 := 0; fi12 < 8; fi12++ {
                v13 := GenS1{F0: 2}
                _ = fi12
                _ = v13
            }
        }
        _ = v5
    }
    p0 = (^(p0))
    if (!((p0 >> ((0) & 63)) < 0)) {
        p0 = p0
    }
    var v14 uint = 166
    _ = p0
    _ = p1
    _ = v14
    return true
}

func main() {
    v1 := []string{}
    var v2 uint = 204
    v3 := make(map[string]int)
    fmt.Println(v2)
    switch "foo bar" {
    case "a":
        for fi15 := 0; fi15 < 20; fi15++ {
            if false {
                fmt.Println(v2, fi15)
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 float32 = ((0.1 - -2.25) * (0.0 * 0.0))
                    _ = fi16
                    _ = v17
                }
                fmt.Println(v2)
            } else {
                v2 -= 42
            }
            _ = len(v1)
            _ = fi15
        }
    case "\n\t":
        var v18 int64 = 9223372036854775807
        _ = v18
    case "hello":
        for fi19 := 0; fi19 < 6; fi19++ {
            v2 = (^(v2))
            if true {
                v20 := GenS1{F0: 531}
                for fi21 := 0; fi21 < 2; fi21++ {
                    fmt.Println(fi19, fi21, fi21)
                    fmt.Println(fi21)
                    _ = fi21
                }
                var v22 uint = (v2 - v2)
                _ = v20
                _ = v22
            } else {
                fmt.Println(v2)
            }
            switch "foo bar" {
            case "a":
                for fi23 := 0; fi23 < 13; fi23++ {
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                    var v24 uint64 = 1
                    _ = fi23
                    _ = v24
                }
            case "foo bar":
                v2 = v2
            case "hello":
                v2 = (v2 >> ((115) & 63))
            }
            _ = fi19
        }
    }
    v2 += (v2 % (((v2 >> ((42) & 63))) | 1))
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 19; fi26++ {
            v2 -= (v2 / (((v2 >> ((v2) & 63))) | 1))
            v2 = (^(v2))
            _ = fi26
        }
        var v27 int64 = 664
        fmt.Println(fi25, fi25, fi25)
        _ = fi25
        _ = v27
    }
    fmt.Println(v2)
    v28 := GenS2{F0: 0}
    switch 42 {
    case -753:
        var v29 int32 = -1
        _ = v29
    case -1:
        if true {
            for fi30 := 0; fi30 < 11; fi30++ {
                switch (fi30 ^ fi30) {
                case -2147483648:
                    for fi31 := 0; fi31 < 3; fi31++ {
                        switch (-(fi30)) {
                        case 85:
                            if true {
                                var v32 uint64 = genFn0(42)
                                v33 := GenS3{F0: (("a" + ("hello" + (((("" + "") + "\n\t") + "\n\t") + "a"))) + ("" + "\n\t"))}
                                v34 := GenS4{F0: v32}
                                fmt.Println(v32)
                                var v35 int = (fi30 & int(v2))
                                v35 += (fi30 - (fi31 * (v35 / ((fi31) | 1))))
                                var v36 uint64 = v32
                                _ = v32
                                _ = v33
                                _ = v34
                                _ = v35
                                _ = v36
                            } else {
                                if (!false) {
                                    switch "foo bar" {
                                    case "hello":
                                        switch ("hello" + "foo bar") {
                                        case "foo bar":
                                            var v37 uint64 = genFn0(133)
                                            _ = v37
                                        case "a":
                                            v2 += (v2 % ((v2) | 1))
                                        case "hello":
                                            v3[("hello" + "")] = 0
                                            v2 -= 65535
                                            if false {
                                                switch (-(fi30)) {
                                                case 367:
                                                    var v38 uint16 = 32767
                                                    _ = v38
                                                default:
                                                    v39 := GenS4{F0: 366}
                                                    _ = v39
                                                }
                                                v2 = (v2 - v2)
                                            }
                                        }
                                    case "":
                                        v2 = (^(v2))
                                    }
                                } else {
                                    switch "foo bar" {
                                    case "a":
                                        if (161 != 0) {
                                            v40 := v3["foo bar"]
                                            if ((2147483647 > 2147483647) || true) {
                                                v40 += v40
                                            } else {
                                                var v41 int = (^(v40))
                                                switch "hello" {
                                                case "a":
                                                    var v42 float32 = float32((float64(-2.25) * 0.1))
                                                    _ = v42
                                                }
                                                _ = v41
                                            }
                                            v40 -= 32768
                                            v43 := GenS5{F0: "hello", F1: 1, F2: 4294967295}
                                            _ = len(v3)
                                            _ = v40
                                            _ = v43
                                        } else {
                                            fmt.Println(fi30)
                                            var v44 int16 = -565
                                            var v45 string = "\n\t"
                                            _ = v44
                                            _ = v45
                                        }
                                    case "\n\t":
                                        switch "\n\t" {
                                        case "a":
                                            v2 += v2
                                            switch ("a" + "foo bar") {
                                            case "\n\t":
                                                v2 -= (v2 + v2)
                                            case "hello":
                                                var v46 uint8 = 0
                                                _ = v46
                                            case "":
                                                v47 := GenS6{F0: false, F1: 1, F2: -9223372036854775808}
                                                _ = v47
                                            default:
                                                if false {
                                                    v48 := GenS5{F0: "", F1: 42, F2: 2147483647}
                                                    v2 = (v2 / ((18446744073709551615) | 1))
                                                    _ = v48
                                                } else {
                                                    if len(v1) > 0 {
                                                        v1[0] = ""
                                                    }
                                                    v49 := v3["a"]
                                                    switch v49 {
                                                    case 32768:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                    _ = v49
                                                }
                                            }
                                        case "hello":
                                            v2 += v2
                                        default:
                                            v2 -= 255
                                        }
                                    default:
                                        fmt.Println(fi30)
                                    }
                                }
                                delete(v3, "")
                            }
                        case 2147483647:
                            fmt.Println(v2, fi31)
                        }
                        if (699 >= v2) {
                            v2 += v2
                            v50 := GenS1{F0: 4294967295}
                            v51 := GenS4{F0: 1}
                            _ = v50
                            _ = v51
                        }
                        v52 := GenS5{F0: "hello", F1: 65535, F2: 68}
                        _ = fi31
                        _ = v52
                    }
                case 282:
                    fmt.Println(fi30)
                case -436:
                    v2 = 878
                }
                _ = fi30
            }
            v53 := GenS7{F0: 459, F1: 1, F2: 2}
            v54 := GenS8{F0: 127}
            _ = v53
            _ = v54
        } else {
            var v55 float64 = 0.1
            fmt.Println(v55, v2)
            _ = v55
        }
    case 572:
        _ = len(v3)
    }
    var v56 int8 = 42
    switch ("" + "hello") {
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "":
        fmt.Println(v56)
    case "foo bar":
        v1 = append(v1, "hello")
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v57 := GenS9{F0: true}
    v56 = v56
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v56
    _ = v57
}
