package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 uint64; F1 int64; F2 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint32 }

func main() {
    var v1 int32 = 99
    v2 := make(map[uint32]string)
    var v3 int64 = -1
    var v4 int = 0
    var v5 uint32 = 2147483648
    v6 := make(map[string]bool)
    v7 := []int{127}
    v5 += 2
    if false {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        var v8 int = (v4 % (((-(v4))) | 1))
        _ = v8
    }
    if (("\n\t" + (("hello" + ("a" + ("hello" + "a"))) + "a")) < (("hello" + "\n\t") + "a")) {
        v9 := GenS1{F0: (!false), F1: 0}
        v10 := GenS2{F0: 2}
        for fi11 := 0; fi11 < 2; fi11++ {
            v5 = v5
            fmt.Println(v5, fi11, v1)
            _ = fi11
        }
        _ = v10
        _ = v9
    } else {
        if false {
            v12 := GenS3{F0: 4294967295, F1: v1, F2: 127}
            v4 = v4
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 uint16 = 2
                _ = fi13
                _ = v14
            }
            _ = v12
        }
        v5 -= 0
    }
    var v15 uint = 459
    switch ("foo bar" + "foo bar") {
    case "a":
        if (false && (!(false || true))) {
            var v16 uint64 = 18446744073709551615
            fmt.Println(v5)
            _ = v16
        } else {
            v7 = append(v7, v4)
            v3 -= v3
            for fi17 := 0; fi17 < 15; fi17++ {
                _ = len(v6)
                _ = fi17
            }
        }
    case "foo bar":
        if true {
            fmt.Println(v5, v3)
            v18 := GenS2{F0: 887}
            var v19 uint8 = 151
            _ = v18
            _ = v19
        } else {
            fmt.Println(v3, v3)
            if false {
                var v20 int16 = 35
                switch "foo bar" {
                case "":
                    var v21 int64 = (v3 ^ v3)
                    _ = v21
                case "a":
                    if false {
                        v22 := GenS2{F0: 18446744073709551615}
                        v4 += 127
                        _ = v22
                    }
                case "foo bar":
                    if false {
                        var v23 uint32 = (v5 % ((675) | 1))
                        if (("\n\t" + ("" + ("a" + "hello"))) < ((("hello" + "\n\t") + "\n\t") + "hello")) {
                            v24 := GenS1{F0: (!(false || true)), F1: 4294967295}
                            v23 = 994
                            _ = v24
                        } else {
                            var v25 int8 = 42
                            _ = v25
                        }
                        _ = v23
                    }
                }
                _ = v20
            } else {
                if (uint(v5) < v15) {
                    if (!false) {
                        if true {
                            delete(v6, "")
                        } else {
                            if false {
                                fmt.Println(v5, v5, v4)
                                v2[65535] = ("foo bar" + "hello")
                                v3 = (-(v3))
                            } else {
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(v5)
                                case "hello":
                                    v26 := GenS3{F0: 1, F1: (-(v1)), F2: 2}
                                    _ = v26
                                default:
                                    v3 = -9223372036854775808
                                }
                                fmt.Println(v5)
                                v3 += 534
                            }
                            var v27 string = ""
                            _ = v27
                        }
                        var v28 int32 = 1
                        for fi29 := 0; fi29 < 3; fi29++ {
                            fmt.Println(v28, v5)
                            if (!false) {
                                if (true || true) {
                                    v7 = append(v7, -199)
                                    v30 := GenS4{F0: 0, F1: v15}
                                    _ = v30
                                } else {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                }
                            } else {
                                v31 := GenS5{F0: 0, F1: true}
                                var v32 uint16 = 751
                                _ = v31
                                _ = v32
                            }
                            if true {
                                v33 := GenS6{F0: 0, F1: 1, F2: "a"}
                                _ = v33
                            }
                            _ = fi29
                        }
                        _ = v28
                    }
                } else {
                    var v34 int64 = (^(v3))
                    if true {
                        var v35 float64 = 1.5
                        _ = v35
                    } else {
                        v36 := v6[("" + "a")]
                        _ = v36
                    }
                    _ = v34
                }
            }
            for fi37 := 0; fi37 < 16; fi37++ {
                v15 -= 1
                v15 += v15
                var v38 int32 = v1
                _ = fi37
                _ = v38
            }
        }
    case "":
        v5 += (v5 | (v5 - v5))
    default:
        v2[2147483647] = ""
    }
    v15 += v15
    v39 := GenS7{F0: 1000}
    v5 -= v5
    var v40 int = (^(v4))
    if (!true) {
        v7 = append(v7, -1)
        v15 = v15
        v41 := GenS8{F0: 1}
        _ = v41
    } else {
        fmt.Println(v5, v40)
        v1 += v1
    }
    var v42 uint16 = 2
    if false {
        for fi43 := 0; fi43 < 9; fi43++ {
            _ = len(v2)
            v42 += uint16(v5)
            _ = fi43
        }
    }
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
