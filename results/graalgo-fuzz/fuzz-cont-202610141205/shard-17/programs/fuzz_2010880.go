package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint16, p1 int8) int64 {
    if true {
        for fi6 := 0; fi6 < 6; fi6++ {
            switch "foo bar" {
            case "\n\t":
                var v7 bool = true
                _ = v7
            case "a":
                for fi8 := 0; fi8 < 17; fi8++ {
                    var v9 string = (("foo bar" + "foo bar") + ("\n\t" + "hello"))
                    _ = fi8
                    _ = v9
                }
            }
            _ = fi6
        }
    } else {
        var v10 float64 = ((((((1e10 * 0.0) + 1.5) / 1.5) - -2.25) - 1e10) + 3.14159)
        _ = v10
    }
    if true {
        p1 = 0
        if true {
            fmt.Println(p0, p0)
        }
        if true {
            var v11 int = 32768
            _ = v11
        }
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 bool, p1 int64) uint16 {
    var v12 string = ((("" + ("\n\t" + "\n\t")) + ("" + "foo bar")) + "\n\t")
    if true {
        v13 := GenS1{F0: 255}
        fmt.Println(p0, p1, v12)
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v12
    return 42
}

func main() {
    var v1 int = 32768
    var v2 string = "\n\t"
    v3 := make(map[int]bool)
    v4 := make(map[int]string)
    var v5 uint = 255
    for fi14 := 0; fi14 < 7; fi14++ {
        v2 = v2
        if (!(v1 < 0)) {
            v1 -= (v1 / (((fi14 % ((0) | 1))) | 1))
            if false {
                var v15 float32 = 0.0
                for fi16 := 0; fi16 < 7; fi16++ {
                    var v17 uint16 = 0
                    var v18 bool = ((-2.25 * v15) <= 0.1)
                    var v19 int = (fi14 & v1)
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
                if false {
                    v2 = ((v2 + v2) + (v2 + v2))
                } else {
                    for fi20 := 0; fi20 < 2; fi20++ {
                        v21 := GenS1{F0: 176}
                        if (!false) {
                            var v22 float32 = float32(1e10)
                            var v23 int32 = 0
                            var v24 uint8 = 2
                            _ = v22
                            _ = v23
                            _ = v24
                        } else {
                            fmt.Println(fi20, v1)
                            v5 += 4294967295
                            v25 := []int{(fi20 << (((v5 << ((v5) & 63))) & 63)), fi14, fi14}
                            if true {
                                switch v1 {
                                case 42:
                                    if true {
                                        fmt.Println(v15, v15)
                                        switch "foo bar" {
                                        case "foo bar":
                                            v25 = append(v25, 0)
                                        case "":
                                            var v26 float32 = -2.25
                                            _ = v26
                                        case "hello":
                                            fmt.Println(v1, v2)
                                        }
                                        switch v2 {
                                        case "a":
                                            var v27 float32 = v15
                                            _ = v27
                                        case "hello":
                                            var v28 uint32 = 636
                                            _ = v28
                                        case "foo bar":
                                            var v29 bool = ((false && true) || (625 < v5))
                                            _ = v29
                                        default:
                                            v30 := GenS2{F0: 100, F1: 18446744073709551615, F2: 255}
                                            _ = v30
                                        }
                                    }
                                default:
                                    var v31 float32 = float32(1e10)
                                    _ = v31
                                }
                                fmt.Println(fi20, fi14)
                            }
                            _ = v25
                        }
                        _ = fi20
                        _ = v21
                    }
                    var v32 uint8 = 127
                    var v33 uint8 = v32
                    _ = v32
                    _ = v33
                }
                _ = v15
            }
        }
        for fi34 := 0; fi34 < 6; fi34++ {
            v35 := GenS3{F0: genFn1((2 != 719), genFn0(genFn1(false, 429), -128))}
            v36 := GenS1{F0: 2}
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi14
    }
    switch "hello" {
    case "":
        v2 = "hello"
    case "\n\t":
        var v37 string = v2
        _ = v37
    case "foo bar":
        var v38 uint8 = 0
        _ = v38
    default:
        v1 = (^(v1))
    }
    var v39 string = ("hello" + v2)
    if (255 > 255) {
        v40 := GenS1{F0: 42}
        if true {
            var v41 float64 = (float64(3.14159) + 0.1)
            _ = v41
        } else {
            v42 := GenS2{F0: 0, F1: 350, F2: 0}
            _ = v42
        }
        _ = v40
    }
    v5 += 0
    var v43 bool = (true || (true || (!true)))
    var v44 uint16 = 419
    for fi45 := 0; fi45 < 5; fi45++ {
        v46 := []int{(v1 + v1)}
        if v43 {
            if false {
                _ = len(v46)
                v1 -= (v1 + fi45)
                if len(v46) > 0 {
                    v46[0] = v1
                }
            }
            v47 := GenS4{F0: true, F1: 2}
            v39 = ""
            _ = v47
        } else {
            var v48 bool = v43
            fmt.Println(v2)
            _ = v48
        }
        var v49 int8 = -128
        _ = fi45
        _ = v46
        _ = v49
    }
    v1 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
