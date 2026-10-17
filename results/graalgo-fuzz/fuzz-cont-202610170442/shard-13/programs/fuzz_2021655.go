package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int8 }
type GenS2 struct { F0 int; F1 int; F2 uint16 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS4 struct { F0 bool; F1 uint32; F2 string }
type GenS5 struct { F0 string; F1 bool; F2 int }
func genFn0() int16 {
    var v5 uint = 943
    v6 := GenS1{F0: 1, F1: 14, F2: 42}
    switch ("a" + "hello") {
    case "foo bar":
        if (681 >= 2) {
            v5 -= 2
            v7 := []int{-677}
            v7 = append(v7, 128)
            _ = v7
        } else {
            for fi8 := 0; fi8 < 1; fi8++ {
                v5 = 42
                _ = fi8
            }
        }
    }
    _ = v5
    _ = v6
    return 1000
}
func genFn1(p0 string, p1 string, p2 string) int64 {
    p2 = (p2 + (p0 + p2))
    p2 = p0
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := []string{(((("foo bar" + "\n\t") + "foo bar") + "\n\t") + "foo bar"), ("foo bar" + "a"), "\n\t"}
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]bool)
    switch 1 {
    case 2147483647:
        var v9 uint16 = 42
        _ = v9
    }
    var v10 int32 = -2147483648
    fmt.Println(v10)
    v10 = -635
    fmt.Println(v10)
    for fi11 := 0; fi11 < 12; fi11++ {
        v10 -= (v10 + v10)
        if false {
            v10 += v10
        } else {
            var v12 uint16 = 892
            _ = len(v1)
            _ = v12
        }
        if (((0.0 / 3.14159) + 0.1) > 1e10) {
            var v13 int16 = genFn0()
            if (true || (863 <= v13)) {
                for fi14 := 0; fi14 < 11; fi14++ {
                    v13 = 0
                    _ = fi14
                }
                if false {
                    var v15 int64 = 0
                    if false {
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                        fmt.Println(v15, v10)
                        if false {
                            v13 = int16(v15)
                            for fi16 := 0; fi16 < 14; fi16++ {
                                v13 = (v13 >> ((271) & 63))
                                _ = fi16
                            }
                        } else {
                            fmt.Println(v13, fi11, v15)
                            if true {
                                switch 9223372036854775807 {
                                case 9223372036854775807:
                                    v13 -= (v13 & (-(v13)))
                                case 4294967295:
                                    var v17 int = fi11
                                    _ = v17
                                }
                                v13 = v13
                            } else {
                                v15 = v15
                            }
                            v18 := GenS1{F0: 42, F1: 1, F2: 127}
                            _ = v18
                        }
                    } else {
                        for fi19 := 0; fi19 < 11; fi19++ {
                            var v20 int8 = 0
                            v20 = (v20 & (^(v20)))
                            _ = fi19
                            _ = v20
                        }
                        v21 := GenS1{F0: 1, F1: -1, F2: int8(v15)}
                        v3[2147483648] = "hello"
                        _ = v21
                    }
                    _ = v15
                }
                var v22 int16 = (^(v13))
                _ = v22
            } else {
                v10 = v10
                if true {
                    switch "foo bar" {
                    case "hello":
                        for fi23 := 0; fi23 < 7; fi23++ {
                            switch "" {
                            case "":
                                switch fi11 {
                                case -442:
                                    v24 := GenS2{F0: (fi23 | -9223372036854775808), F1: fi23, F2: 42}
                                    _ = v24
                                default:
                                    v25 := GenS2{F0: fi23, F1: fi23, F2: 42}
                                    _ = v25
                                }
                                if false {
                                    switch ("hello" + ("foo bar" + "a")) {
                                    case "a":
                                        v26 := GenS1{F0: 578, F1: 42, F2: 127}
                                        _ = v26
                                    default:
                                        delete(v3, 874)
                                    }
                                    v27 := GenS3{F0: 2, F1: 65535, F2: 42}
                                    _ = v27
                                }
                                var v28 int16 = -32768
                                _ = v28
                            case "hello":
                                var v29 int8 = 1
                                _ = v29
                            default:
                                delete(v2, "hello")
                            }
                            _ = fi23
                        }
                    case "\n\t":
                        for fi30 := 0; fi30 < 8; fi30++ {
                            switch fi30 {
                            case 342:
                                if false {
                                    v10 += v10
                                } else {
                                    switch (("a" + (("a" + "") + "foo bar")) + "hello") {
                                    case "":
                                        if true {
                                            var v31 uint32 = 2
                                            switch 217 {
                                            case 128:
                                                var v32 int64 = 0
                                                _ = v32
                                            case 127:
                                                var v33 uint64 = 1
                                                _ = v33
                                            case 0:
                                                v10 = v10
                                            }
                                            _ = v31
                                        }
                                    default:
                                        switch ((((("a" + (("foo bar" + "hello") + "a")) + "a") + "") + "") + "foo bar") {
                                        case "a":
                                            switch fi30 {
                                            case -9223372036854775808:
                                                v2["hello"] = fi11
                                            case 128:
                                                fmt.Println(fi30)
                                            }
                                            var v34 uint64 = 2
                                            _ = v34
                                        default:
                                            var v35 float64 = float64(3.14159)
                                            _ = v35
                                        }
                                    }
                                }
                            case 1:
                                v10 = 2147483647
                                v10 = v10
                                v10 = v10
                            case 128:
                                v10 = -130
                            }
                            _ = fi30
                        }
                    default:
                        v36 := GenS4{F0: (false && true), F1: 65535, F2: ("hello" + "a")}
                        _ = v36
                    }
                } else {
                    v10 += v10
                    v37 := GenS5{F0: "", F1: false, F2: (-(fi11))}
                    switch fi11 {
                    case 32767:
                        for fi38 := 0; fi38 < 15; fi38++ {
                            v13 = 1000
                            _ = fi38
                        }
                    case 9223372036854775807:
                        switch fi11 {
                        case 128:
                            var v39 int64 = 0
                            _ = v39
                        }
                    }
                    _ = v37
                }
            }
            _ = v13
        }
        _ = fi11
    }
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    v10 = (v10 ^ v10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
