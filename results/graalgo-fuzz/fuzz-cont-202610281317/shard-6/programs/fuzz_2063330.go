package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int32; F2 string }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint8; F1 uint32; F2 bool }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 int16; F1 uint32 }
func genFn0(p0 uint32) int {
    p0 = (p0 + p0)
    v6 := []int{-2147483648, -8, 374, 2147483647}
    _ = p0
    _ = v6
    return 128
}
func genFn1(p0 uint8) string {
    fmt.Println(p0)
    if (true && false) {
        v7 := GenS1{F0: 9223372036854775807, F1: 42}
        _ = v7
    }
    _ = p0
    return "a"
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{}
    var v3 int64 = 0
    var v4 int16 = 32767
    var v5 int64 = v3
    var v8 float64 = ((3.14159 / -2.25) + 0.0)
    for fi9 := 0; fi9 < 2; fi9++ {
        switch "foo bar" {
        case "a":
            switch ("\n\t" + "foo bar") {
            case "a":
                for fi10 := 0; fi10 < 18; fi10++ {
                    if true {
                        v4 += -32768
                        v2 = append(v2, 880)
                    } else {
                        var v11 bool = true
                        v12 := GenS2{F0: 65535, F1: 2147483647, F2: "a"}
                        fmt.Println(fi9, v3)
                        _ = v11
                        _ = v12
                    }
                    fmt.Println(fi10)
                    switch "\n\t" {
                    case "foo bar":
                        if true {
                            v3 = v3
                        } else {
                            fmt.Println(fi9, fi10)
                            _ = len(v2)
                            switch (fi9 % (((fi10 / ((fi10) | 1))) | 1)) {
                            case 0:
                                if true {
                                    var v13 bool = false
                                    var v14 uint16 = 255
                                    _ = v13
                                    _ = v14
                                } else {
                                    switch fi9 {
                                    case -586:
                                        switch "foo bar" {
                                        case "a":
                                            if (!(4294967295 != 2)) {
                                                var v15 int = (fi9 + fi10)
                                                if true {
                                                    if (11 <= -100) {
                                                        if (100 > 0) {
                                                            if (true && (1.5 <= 0.1)) {
                                                                var v16 uint64 = 1
                                                                v17 := GenS1{F0: genFn0(2147483648), F1: 971}
                                                                _ = v16
                                                                _ = v17
                                                            }
                                                            v15 -= (-(fi10))
                                                            v18 := v1[4294967295]
                                                            _ = v18
                                                        }
                                                        var v19 float64 = 1e10
                                                        _ = v19
                                                    }
                                                }
                                                _ = v15
                                            }
                                        case "\n\t":
                                            var v20 string = ("foo bar" + "a")
                                            _ = v20
                                        case "":
                                            v4 = v4
                                        default:
                                            if false {
                                                v3 = v3
                                            }
                                        }
                                    }
                                    v4 = (v4 % ((32767) | 1))
                                    fmt.Println(v3, fi10, v5)
                                }
                            }
                        }
                    case "":
                        v4 = v4
                    default:
                        if (42 != 127) {
                            var v21 uint16 = 2
                            _ = v21
                        }
                    }
                    _ = fi10
                }
            case "\n\t":
                v8 += v8
            case "":
                v22 := GenS2{F0: 899, F1: -1, F2: "foo bar"}
                _ = v22
            }
        case "foo bar":
            for fi23 := 0; fi23 < 5; fi23++ {
                v24 := GenS1{F0: -141, F1: 204}
                var v25 uint8 = 128
                _ = fi23
                _ = v24
                _ = v25
            }
        case "\n\t":
            for fi26 := 0; fi26 < 13; fi26++ {
                fmt.Println(v5, v4)
                v27 := GenS2{F0: 457, F1: 42, F2: ("foo bar" + "foo bar")}
                _ = fi26
                _ = v27
            }
        default:
            v28 := GenS3{F0: 56, F1: v4}
            _ = v28
        }
        _ = fi9
    }
    _ = len(v1)
    fmt.Println(v4, v8)
    v2 = append(v2, 74)
    v3 += (v3 / ((v5) | 1))
    v4 = v4
    for fi29 := 0; fi29 < 11; fi29++ {
        v30 := GenS4{F0: 128, F1: 0, F2: false}
        v3 = (v5 ^ 42)
        var v31 int = (fi29 / ((887) | 1))
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 bool = ("a" >= "")
    if true {
        for fi33 := 0; fi33 < 9; fi33++ {
            v34 := GenS2{F0: 1, F1: 2147483647, F2: "hello"}
            v35 := GenS5{F0: 32767, F1: 255}
            _ = fi33
            _ = v34
            _ = v35
        }
        switch ("a" + ("hello" + "foo bar")) {
        case "":
            v36 := GenS3{F0: 127, F1: 32767}
            _ = v36
        default:
            v37 := GenS3{F0: 434, F1: (v4 & v4)}
            _ = v37
        }
        _ = len(v2)
    }
    for fi38 := 0; fi38 < 1; fi38++ {
        v32 = false
        for fi39 := 0; fi39 < 2; fi39++ {
            if (0.1 > 0.1) {
                v40 := GenS4{F0: 127, F1: 0, F2: (!(255 >= 220))}
                v5 += v3
                _ = v40
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi39
        }
        _ = fi38
    }
    v41 := GenS5{F0: 163, F1: 194}
    for fi42 := 0; fi42 < 12; fi42++ {
        var v43 bool = (v8 < v8)
        for fi44 := 0; fi44 < 17; fi44++ {
            if false {
                v43 = v32
                switch fi44 {
                case 1:
                    v45 := GenS6{F0: v4, F1: 4294967295}
                    _ = v45
                case -1:
                    delete(v1, 69)
                case -2147483648:
                    v4 += 1
                }
            }
            v43 = (1 > 1000)
            _ = fi44
        }
        _ = fi42
        _ = v43
    }
    v32 = (!v32)
    switch genFn0(1) {
    case 2147483647:
        v3 = (-(v3))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v41
    _ = v5
    _ = v8
}
