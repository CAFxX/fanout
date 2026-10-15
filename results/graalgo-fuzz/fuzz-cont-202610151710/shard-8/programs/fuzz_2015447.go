package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int; F2 string }
type GenS5 struct { F0 int; F1 int }
func genFn0(p0 int16) int {
    var v7 int16 = (p0 % (((p0 + (p0 << ((4294967295) & 63)))) | 1))
    fmt.Println(v7, v7)
    for fi8 := 0; fi8 < 3; fi8++ {
        v7 -= p0
        var v9 int8 = 100
        for fi10 := 0; fi10 < 9; fi10++ {
            switch "a" {
            case "a":
                v11 := []int{(fi10 & fi10), -591}
                _ = v11
            case "foo bar":
                var v12 int64 = -1
                _ = v12
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v7
    return 32768
}
func genFn1() int8 {
    switch 32768 {
    case -582:
        var v13 int8 = -14
        _ = v13
    }
    v14 := []int{genFn0(622), 575, 2147483647}
    var v15 string = (("hello" + "a") + (("hello" + "hello") + ("a" + "foo bar")))
    var v16 uint = 1
    v15 = "\n\t"
    _ = v14
    _ = v15
    _ = v16
    return 127
}

func main() {
    var v1 int8 = 33
    var v2 int64 = 0
    v3 := make(map[uint32]bool)
    var v4 int8 = 111
    v5 := make(map[uint32]bool)
    v6 := []int{}
    v6 = append(v6, -9223372036854775808)
    var v17 bool = (42 <= 44)
    switch ("\n\t" + ("" + "\n\t")) {
    case "foo bar":
        v3[42] = ((!(!v17)) || (false || v17))
    case "a":
        v18 := GenS1{F0: 255}
        _ = v18
    case "hello":
        v2 = v2
    }
    v19 := GenS1{F0: 1}
    for fi20 := 0; fi20 < 19; fi20++ {
        switch ("foo bar" + ("hello" + "\n\t")) {
        case "":
            var v21 float32 = (float32(1e10) + -2.25)
            _ = v21
        case "foo bar":
            for fi22 := 0; fi22 < 16; fi22++ {
                if (("a" + "a") > ("hello" + "a")) {
                    v23 := GenS2{F0: 0, F1: -822}
                    _ = v23
                }
                if ((65535 >= 203) || ((float32(-2.25) + 3.14159) <= 1.5)) {
                    var v24 float32 = (float32(0.0) - 0.1)
                    var v25 uint8 = 127
                    v4 += (v4 ^ v1)
                    _ = v24
                    _ = v25
                }
                if false {
                    v2 = (v2 / ((v2) | 1))
                }
                _ = fi22
            }
        case "a":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        switch (fi20 >> ((2) & 63)) {
        case -109:
            var v26 float32 = -2.25
            _ = v26
        case 32767:
            if v17 {
                if true {
                    for fi27 := 0; fi27 < 20; fi27++ {
                        v28 := GenS3{F0: 42}
                        v4 -= (v1 & 122)
                        _ = fi27
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 9; fi29++ {
                        if v17 {
                            v2 -= 984
                        }
                        _ = fi29
                    }
                } else {
                    v2 = v2
                    switch genFn0(-32768) {
                    case -800:
                        v1 -= (v1 + 42)
                    case -1:
                        switch (fi20 * fi20) {
                        case 42:
                            v4 += (v4 - (v1 % (((v1 - (v1 & genFn1()))) | 1)))
                        case -9223372036854775808:
                            v2 = -1
                        default:
                            for fi30 := 0; fi30 < 9; fi30++ {
                                v31 := GenS1{F0: 42}
                                v32 := GenS2{F0: 1, F1: -1}
                                switch (fi30 % ((127) | 1)) {
                                case 1:
                                    v33 := GenS4{F0: v17, F1: fi20, F2: "\n\t"}
                                    _ = v33
                                case -865:
                                    switch fi20 {
                                    case 32767:
                                        v1 = (v1 % (((v1 + genFn1())) | 1))
                                    case 639:
                                        v6 = append(v6, fi20)
                                    case 1:
                                        var v34 uint64 = 2
                                        _ = v34
                                    default:
                                        fmt.Println(v1, fi20)
                                    }
                                case 0:
                                    fmt.Println(v17, v1, v2)
                                    if (100 != (v1 * (v1 | (v1 * v1)))) {
                                        fmt.Println(fi30, v17, v2)
                                    }
                                    fmt.Println(v1)
                                }
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        }
                    default:
                        v5[2147483647] = ("hello" != "hello")
                    }
                }
                for fi35 := 0; fi35 < 2; fi35++ {
                    v36 := GenS1{F0: 110}
                    fmt.Println(fi35, fi20, fi35)
                    _ = fi35
                    _ = v36
                }
            } else {
                v37 := GenS1{F0: 42}
                _ = v37
            }
        default:
            v1 = v1
        }
        _ = fi20
    }
    v2 = 1
    if (4294967295 >= 255) {
        for fi38 := 0; fi38 < 13; fi38++ {
            v2 += v2
            _ = fi38
        }
        var v39 int32 = -1
        _ = v39
    } else {
        for fi40 := 0; fi40 < 18; fi40++ {
            var v41 float64 = (float64((1.5 + 1e10)) / (float64(0.0) + float64(3.14159)))
            _ = fi40
            _ = v41
        }
        v42 := GenS3{F0: 255}
        if true {
            v43 := GenS5{F0: 4294967295, F1: -9223372036854775808}
            _ = v43
        }
        _ = v42
    }
    v17 = v17
    for fi44 := 0; fi44 < 8; fi44++ {
        v45 := GenS5{F0: genFn0(42), F1: 1}
        _ = fi44
        _ = v45
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v17 = ("" != "foo bar")
    fmt.Println(v1, v17, v17)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
