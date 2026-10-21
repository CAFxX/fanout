package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint64; F1 int8 }
func genFn0() int64 {
    v6 := []int{}
    var v7 float64 = 1.5
    _ = v6
    _ = v7
    return -456
}
func genFn1(p0 int64, p1 int, p2 int64) string {
    if false {
        p2 = genFn0()
    }
    if false {
        p2 += 426
    }
    var v8 bool = false
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := []int{}
        if (1 > 0) {
            if true {
                switch (p1 & p1) {
                case 42:
                    if true {
                        if (v8 || (p1 != fi9)) {
                            for fi11 := 0; fi11 < 9; fi11++ {
                                v8 = true
                                var v12 int = (fi9 << ((712) & 63))
                                _ = fi11
                                _ = v12
                            }
                            for fi13 := 0; fi13 < 5; fi13++ {
                                if (1000 > -202) {
                                    var v14 uint64 = 456
                                    v15 := GenS1{F0: 915}
                                    _ = v14
                                    _ = v15
                                }
                                v8 = false
                                fmt.Println(p0)
                                _ = fi13
                            }
                            var v16 uint = 42
                            _ = v16
                        } else {
                            switch "\n\t" {
                            case "":
                                var v17 float32 = 1.5
                                _ = v17
                            }
                            for fi18 := 0; fi18 < 14; fi18++ {
                                v19 := GenS1{F0: -2147483648}
                                v8 = v8
                                fmt.Println(p2, v8, fi9)
                                _ = fi18
                                _ = v19
                            }
                            fmt.Println(p2, p1, v8)
                        }
                    }
                case -322:
                    for fi20 := 0; fi20 < 9; fi20++ {
                        p2 = p0
                        var v21 uint64 = 1
                        v22 := GenS2{F0: v8, F1: (((("foo bar" + ("a" + ((("foo bar" + ("a" + (("\n\t" + "a") + "\n\t"))) + "hello") + "foo bar"))) + "") + ("" + "a")) + ("" + "hello"))}
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                default:
                    var v23 int8 = 100
                    _ = v23
                }
                var v24 int64 = (p2 / ((p2) | 1))
                _ = v24
            } else {
                v25 := GenS1{F0: 0}
                v26 := GenS2{F0: v8, F1: "a"}
                var v27 int8 = 100
                _ = v25
                _ = v26
                _ = v27
            }
        } else {
            for fi28 := 0; fi28 < 13; fi28++ {
                if len(v10) > 0 {
                    v10[0] = (^(fi9))
                }
                var v29 uint64 = 0
                v29 -= v29
                var v30 uint8 = 2
                v31 := GenS3{F0: uint8(v29), F1: uint64(v30)}
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
            var v32 string = "a"
            for fi33 := 0; fi33 < 5; fi33++ {
                p1 -= fi33
                var v34 int8 = 100
                _ = len(v10)
                _ = fi33
                _ = v34
            }
            _ = v32
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return ("\n\t" + ("hello" + "foo bar"))
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 961
    v3 := []string{"\n\t", "\n\t"}
    v4 := make(map[string]bool)
    v5 := make(map[uint32]bool)
    v2 -= 42
    v35 := GenS4{F0: 2, F1: -128}
    for fi36 := 0; fi36 < 5; fi36++ {
        v2 -= v2
        _ = fi36
    }
    v37 := GenS3{F0: 1, F1: 0}
    v2 = (v2 / ((0) | 1))
    fmt.Println(v2)
    switch ("hello" + ("" + (("" + "") + ("foo bar" + "hello")))) {
    case "a":
        v2 -= (v2 - 42)
    case "\n\t":
        v2 -= (v2 % (((-(v2))) | 1))
    case "hello":
        for fi38 := 0; fi38 < 3; fi38++ {
            v39 := v1[2147483648]
            fmt.Println(fi38, v2, v2)
            switch 9223372036854775807 {
            case 525:
                for fi40 := 0; fi40 < 10; fi40++ {
                    v41 := GenS4{F0: 722, F1: 100}
                    switch fi38 {
                    case 924:
                        if (!(976 > 264)) {
                            v3 = append(v3, "\n\t")
                        } else {
                            switch "a" {
                            case "a":
                                v39 = (1 > -107)
                            case "":
                                var v42 int16 = 0
                                _ = v42
                            }
                        }
                    case -831:
                        fmt.Println(v2, fi40)
                    case 1:
                        var v43 uint64 = 2
                        _ = v43
                    default:
                        var v44 int16 = (v2 - v2)
                        _ = v44
                    }
                    v45 := GenS1{F0: -2147483648}
                    _ = fi40
                    _ = v41
                    _ = v45
                }
            case -2147483648:
                fmt.Println(v39, v39)
            default:
                v39 = (32767 > 484)
            }
            _ = fi38
            _ = v39
        }
    default:
        var v46 int8 = -128
        _ = v46
    }
    switch "a" {
    case "foo bar":
        v2 += v2
    case "a":
        v2 -= (v2 % ((v2) | 1))
    case "hello":
        v2 = (v2 + v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v4
    _ = v5
}
