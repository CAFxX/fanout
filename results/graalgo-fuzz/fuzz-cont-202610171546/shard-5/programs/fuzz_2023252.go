package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 uint8; F1 uint16 }
func genFn0() int32 {
    var v7 uint = 226
    if (-125 > 42) {
        if false {
            switch "hello" {
            case "hello":
                v8 := []int{32768, 128}
                _ = v8
            case "a":
                switch "\n\t" {
                case "hello":
                    v9 := []int{0, 32767, 127}
                    _ = v9
                case "foo bar":
                    switch 127 {
                    case 469:
                        v7 += 795
                    }
                case "\n\t":
                    if (128 >= 32767) {
                        var v10 uint32 = 777
                        _ = v10
                    }
                default:
                    v11 := GenS1{F0: -925, F1: 32767, F2: 4294967295}
                    _ = v11
                }
            default:
                var v12 bool = true
                _ = v12
            }
            if false {
                var v13 int16 = -838
                if (v7 > 393) {
                    if true {
                        v7 += v7
                    } else {
                        v14 := GenS2{F0: 70}
                        v15 := GenS3{F0: (("foo bar" + "a") + (("" + "foo bar") + "hello"))}
                        _ = v14
                        _ = v15
                    }
                    if false {
                        var v16 uint = v7
                        _ = v16
                    } else {
                        v17 := []int{-2147483648}
                        v18 := GenS2{F0: 42}
                        _ = len(v17)
                        _ = v17
                        _ = v18
                    }
                    switch ("foo bar" + "a") {
                    case "foo bar":
                        for fi19 := 0; fi19 < 9; fi19++ {
                            for fi20 := 0; fi20 < 2; fi20++ {
                                var v21 uint8 = 128
                                fmt.Println(fi20, v13)
                                v22 := []int{}
                                var v23 uint64 = 817
                                if true {
                                    v24 := GenS4{F0: 255, F1: v21}
                                    _ = v24
                                } else {
                                    var v25 int = int(v13)
                                    if false {
                                        fmt.Println(v13, fi20, v21)
                                    } else {
                                        if len(v22) >= 2 {
                                            v22 = v22[1:len(v22)]
                                        }
                                    }
                                    var v26 int32 = 2147483647
                                    _ = v25
                                    _ = v26
                                }
                                _ = fi20
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            v13 = (v13 | (-(v13)))
                            _ = fi19
                        }
                    default:
                        if false {
                            v27 := []int{-1}
                            v13 += v13
                            _ = len(v27)
                            _ = v27
                        } else {
                            v7 = (v7 | (^(v7)))
                            for fi28 := 0; fi28 < 16; fi28++ {
                                var v29 uint = (v7 << (((v7 / ((v7) | 1))) & 63))
                                _ = fi28
                                _ = v29
                            }
                        }
                    }
                }
                _ = v13
            } else {
                v30 := GenS4{F0: 255, F1: 0}
                _ = v30
            }
        }
    } else {
        v7 -= 691
        v7 += v7
        var v31 float32 = 0.1
        _ = v31
    }
    var v32 bool = (true && false)
    var v33 string = ("hello" + "\n\t")
    _ = v32
    _ = v33
    _ = v7
    return 2147483647
}
func genFn1(p0 int64, p1 uint16) uint8 {
    switch ("hello" + ("\n\t" + "")) {
    case "hello":
        fmt.Println(p0, p0)
    case "":
        if true {
            p0 -= -1
        } else {
            fmt.Println(p1)
            v34 := []int{9223372036854775807, -9223372036854775808, -9223372036854775808, -1}
            if len(v34) >= 2 {
                v34 = v34[1:len(v34)]
            }
            _ = v34
        }
    default:
        var v35 int64 = p0
        _ = v35
    }
    p1 -= (p1 & p1)
    p0 = (p0 - p0)
    _ = p0
    _ = p1
    return 2
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]int)
    var v3 int8 = 1
    var v4 uint16 = 255
    v5 := []int{736, 1, 128}
    var v6 string = ("a" + "foo bar")
    fmt.Println(v3, v4)
    v3 = v3
    v3 = (v3 / (((v3 + v3)) | 1))
    fmt.Println(v4, v6, v4)
    if true {
        v6 = "foo bar"
        v3 += 1
    } else {
        fmt.Println(v6)
    }
    fmt.Println(v6, v6)
    v6 = "foo bar"
    var v36 int32 = 42
    fmt.Println(v36)
    fmt.Println(v4)
    _ = len(v2)
    v6 = v6
    switch (("hello" + "a") + (("hello" + (("a" + "a") + (("\n\t" + v6) + v6))) + v6)) {
    case "\n\t":
        var v37 uint64 = 189
        _ = v37
    case "foo bar":
        var v38 int = 0
        _ = v38
    case "a":
        if false {
            switch "foo bar" {
            case "foo bar":
                var v39 uint64 = 18446744073709551615
                _ = v39
            case "hello":
                v36 -= genFn0()
            case "a":
                for fi40 := 0; fi40 < 16; fi40++ {
                    var v41 uint8 = 128
                    _ = fi40
                    _ = v41
                }
            default:
                if false {
                    v6 = v6
                }
            }
            v42 := GenS5{F0: genFn1(-1, v4), F1: 0}
            _ = v42
        } else {
            v6 = "foo bar"
            if (!false) {
                if (genFn0() <= (v36 * (v36 + (v36 ^ (-(v36)))))) {
                    v43 := GenS3{F0: v6}
                    var v44 int64 = -834
                    for fi45 := 0; fi45 < 13; fi45++ {
                        fmt.Println(v3, v6, v44)
                        v6 = "foo bar"
                        v46 := GenS5{F0: 127, F1: v4}
                        _ = fi45
                        _ = v46
                    }
                    _ = v43
                    _ = v44
                }
                v47 := v2[0]
                _ = v47
            }
            delete(v2, 0)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
