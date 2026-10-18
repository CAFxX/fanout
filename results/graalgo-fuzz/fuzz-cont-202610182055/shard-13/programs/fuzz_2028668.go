package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 uint; F1 bool; F2 string }
type GenS9 struct { F0 int16; F1 uint64; F2 uint64 }
func genFn0(p0 uint, p1 string) uint {
    p0 = (p0 | (p0 >> (((p0 & 65535)) & 63)))
    if true {
        p1 = "hello"
    } else {
        var v4 int8 = 100
        v5 := []int{286}
        _ = v4
        _ = v5
    }
    var v6 uint16 = 693
    var v7 uint8 = 205
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint, p1 uint32) uint {
    p0 += (^(p0))
    v8 := []int{-9223372036854775808, -595, 2147483647}
    _ = p0
    _ = p1
    _ = v8
    return genFn0(p0, "\n\t")
}

func main() {
    var v1 uint64 = 0
    v2 := []string{(("a" + "a") + "foo bar"), "hello"}
    v3 := make(map[string]string)
    v1 -= (v1 % ((v1) | 1))
    v9 := GenS1{F0: 4294967295, F1: -1, F2: 0}
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 bool = (true && (!((float32(0.0) + 1.5) < 1.5)))
        fmt.Println(fi10, fi10)
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 float64 = ((float64(float32(1.5)) * float64(1.5)) + (-2.25 + ((1.5 * 1e10) / float64(1.5))))
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    _ = len(v3)
    if true {
        v14 := GenS1{F0: v1, F1: 477, F2: 128}
        _ = v14
    }
    switch "a" {
    case "foo bar":
        v2 = append(v2, "")
    default:
        switch 4294967295 {
        case -9223372036854775808:
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 uint64 = v1
                switch (-(fi15)) {
                case 389:
                    if false {
                        v17 := GenS2{F0: "", F1: (false || false)}
                        v18 := GenS2{F0: ("" + "foo bar"), F1: (42 != 159)}
                        _ = v17
                        _ = v18
                    } else {
                        for fi19 := 0; fi19 < 6; fi19++ {
                            switch (("hello" + "") + "hello") {
                            case "foo bar":
                                v20 := GenS3{F0: 102, F1: ""}
                                _ = v20
                            default:
                                fmt.Println(v16)
                            }
                            var v21 string = "a"
                            v22 := GenS4{F0: 42}
                            _ = fi19
                            _ = v21
                            _ = v22
                        }
                        v16 += (v16 / (((v16 | 18446744073709551615)) | 1))
                    }
                case 127:
                    v16 += (v1 | v1)
                case 2147483647:
                    switch "foo bar" {
                    case "hello":
                        var v23 uint32 = 2147483647
                        _ = v23
                    case "":
                        v24 := GenS3{F0: 100, F1: "\n\t"}
                        _ = v24
                    }
                }
                for fi25 := 0; fi25 < 11; fi25++ {
                    v16 -= v16
                    if false {
                        v16 += (v16 % ((v16) | 1))
                        v26 := GenS5{F0: 65535, F1: 128}
                        _ = v26
                    }
                    fmt.Println(fi25, fi15)
                    if (-898 == -1) {
                        if false {
                            fmt.Println(fi15, v16, fi25)
                            if ((^(v1)) != v1) {
                                var v27 uint = 42
                                switch "foo bar" {
                                case "":
                                    v28 := GenS1{F0: v1, F1: -523, F2: 128}
                                    _ = v28
                                default:
                                    var v29 float32 = -2.25
                                    _ = v29
                                }
                                switch 128 {
                                case 4294967295:
                                    var v30 bool = (2 < (v27 & (v27 ^ v27)))
                                    _ = v30
                                case 32768:
                                    v1 -= (v16 >> (((^(v27))) & 63))
                                    v1 += 4294967295
                                case 128:
                                    v16 -= 42
                                default:
                                    v31 := GenS4{F0: 32767}
                                    _ = v31
                                }
                                v16 += v1
                                switch -985 {
                                case 0:
                                    if (42 <= 2147483647) {
                                        if true {
                                            fmt.Println(fi15, v1)
                                            var v32 bool = false
                                            v16 = v1
                                            var v33 uint64 = (v16 * v16)
                                            v34 := GenS6{F0: 1, F1: 65535, F2: 675}
                                            _ = v32
                                            _ = v33
                                            _ = v34
                                        }
                                        v1 -= v1
                                    } else {
                                        var v35 int8 = 127
                                        v36 := GenS5{F0: 0, F1: 128}
                                        _ = v35
                                        _ = v36
                                    }
                                case 820:
                                    switch (("\n\t" + "hello") + "\n\t") {
                                    case "\n\t":
                                        fmt.Println(fi15)
                                    default:
                                        var v37 uint16 = 943
                                        _ = v37
                                    }
                                case -824:
                                    v27 += v27
                                default:
                                    _ = len(v3)
                                }
                                _ = v27
                            }
                            v1 -= (v1 >> ((1) & 63))
                            if false {
                                fmt.Println(v16, v1, v1)
                                v38 := GenS1{F0: v1, F1: 9223372036854775807, F2: 182}
                                v16 += 4294967295
                                if (0 >= 4294967295) {
                                    var v39 int16 = -32768
                                    var v40 int8 = 16
                                    if (42 > 649) {
                                        v3[""] = "foo bar"
                                        v41 := GenS7{F0: (v39 / (((v39 >> ((297) & 63))) | 1)), F1: 128}
                                        _ = v41
                                    }
                                    _ = v39
                                    _ = v40
                                } else {
                                    var v42 uint16 = 1
                                    if true {
                                        v43 := GenS8{F0: 1, F1: false, F2: "hello"}
                                        if (1 == -36) {
                                            var v44 float32 = (0.0 / 1e10)
                                            var v45 uint16 = (v42 * (v42 + v42))
                                            var v46 uint = 795
                                            var v47 uint = genFn1(v46, 526)
                                            v48 := GenS3{F0: 42, F1: (("foo bar" + "foo bar") + ("hello" + "hello"))}
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                            _ = v47
                                            _ = v48
                                        } else {
                                            v49 := GenS4{F0: -1}
                                            if true {
                                                var v50 string = ""
                                                v16 = (v1 & v16)
                                                v51 := GenS1{F0: (v1 - v1), F1: 1, F2: 127}
                                                _ = v50
                                                _ = v51
                                            }
                                            var v52 int64 = 931
                                            _ = v49
                                            _ = v52
                                        }
                                        _ = v43
                                    }
                                    _ = v42
                                }
                                fmt.Println(fi25)
                                _ = v38
                            }
                        }
                    }
                    _ = fi25
                }
                _ = fi15
                _ = v16
            }
        case 127:
            var v53 float64 = (-2.25 + ((1e10 * 0.0) + ((1e10 - 0.1) + 1.5)))
            _ = v53
        case 42:
            _ = len(v2)
        }
    }
    v54 := GenS8{F0: 112, F1: false, F2: "hello"}
    _ = len(v3)
    var v55 float32 = 1e10
    if true {
        delete(v3, "\n\t")
    } else {
        v56 := GenS3{F0: 105, F1: "hello"}
        v3[(("foo bar" + ("foo bar" + "hello")) + "\n\t")] = "hello"
        if (42 != (v1 - v1)) {
            for fi57 := 0; fi57 < 13; fi57++ {
                v1 = v1
                v1 = 992
                _ = fi57
            }
        } else {
            v58 := GenS9{F0: 32767, F1: v1, F2: v1}
            _ = v58
        }
        _ = v56
    }
    fmt.Println(v1, v55)
    fmt.Println(v1)
    fmt.Println(v1, v55)
    v59 := GenS8{F0: 323, F1: false, F2: "foo bar"}
    switch "\n\t" {
    case "hello":
        delete(v3, "")
    }
    var v60 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v55:", v55)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v54
    _ = v55
    _ = v59
    _ = v60
    _ = v9
}
