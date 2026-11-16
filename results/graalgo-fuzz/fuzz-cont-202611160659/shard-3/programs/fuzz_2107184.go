package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 uint; F1 int64 }

func main() {
    var v1 int8 = 92
    var v2 int64 = 9223372036854775807
    v3 := []uint32{}
    v4 := []int{}
    var v5 uint32 = 2
    if false {
        v6 := GenS1{F0: v1}
        v7 := GenS2{F0: ("foo bar" + "\n\t")}
        _ = v6
        _ = v7
    } else {
        if (5 != 0) {
            v8 := GenS3{F0: 42, F1: 2}
            switch 2147483647 {
            case -9223372036854775808:
                for fi9 := 0; fi9 < 9; fi9++ {
                    v1 = (v1 % ((-41) | 1))
                    switch "" {
                    case "\n\t":
                        v1 = v1
                    }
                    _ = fi9
                }
            case -384:
                v10 := GenS3{F0: 42, F1: 991}
                _ = v10
            default:
                v11 := GenS4{F0: ("" + "foo bar"), F1: false}
                _ = v11
            }
            _ = v8
        }
    }
    switch -2147483648 {
    case 4294967295:
        if ((false || ((20 > 32767) && (2 < 0))) && true) {
            var v12 uint16 = 32767
            switch ("\n\t" + ("hello" + (("a" + "a") + ""))) {
            case "hello":
                for fi13 := 0; fi13 < 13; fi13++ {
                    if (0 < 42) {
                        switch (^(fi13)) {
                        case -882:
                            if false {
                                if false {
                                    switch "" {
                                    case "\n\t":
                                        v5 = 2147483648
                                    case "":
                                        switch "" {
                                        case "":
                                            v3 = append(v3, v5)
                                        case "hello":
                                            v1 -= (v1 & 42)
                                        case "foo bar":
                                            if (((v12 | (v12 / ((255) | 1))) != v12) && false) {
                                                var v14 uint = 4294967295
                                                var v15 int64 = v2
                                                var v16 uint = (v14 & 573)
                                                _ = v14
                                                _ = v15
                                                _ = v16
                                            }
                                        }
                                    case "foo bar":
                                        v5 = (v5 - (v5 % ((v5) | 1)))
                                    }
                                }
                                for fi17 := 0; fi17 < 10; fi17++ {
                                    var v18 uint32 = (v5 / (((v5 + 2147483648)) | 1))
                                    var v19 int64 = (v2 & -564)
                                    switch fi17 {
                                    case 2147483647:
                                        var v20 uint16 = 726
                                        _ = v20
                                    case 42:
                                        fmt.Println(v1)
                                    case 4294967295:
                                        var v21 uint32 = v5
                                        _ = v21
                                    }
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                }
                            } else {
                                for fi22 := 0; fi22 < 19; fi22++ {
                                    var v23 string = "a"
                                    v5 += (v5 ^ (v5 & (v5 ^ (^(v5)))))
                                    _ = fi22
                                    _ = v23
                                }
                                _ = len(v3)
                            }
                        case 1:
                            switch ("" + "foo bar") {
                            case "\n\t":
                                var v24 float32 = 0.0
                                _ = v24
                            case "":
                                if (!true) {
                                    fmt.Println(v5)
                                    var v25 uint = 559
                                    var v26 int8 = v1
                                    _ = v25
                                    _ = v26
                                } else {
                                    if (v5 >= v5) {
                                        v1 = -128
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = 317
                                        }
                                        v27 := GenS5{F0: 2, F1: v2}
                                        if false {
                                            if false {
                                                v5 = v5
                                                fmt.Println(v5)
                                            }
                                        }
                                        _ = v27
                                    }
                                }
                            case "a":
                                v28 := GenS2{F0: ("foo bar" + "a")}
                                _ = v28
                            }
                        }
                    }
                    var v29 float64 = -2.25
                    _ = fi13
                    _ = v29
                }
            case "\n\t":
                var v30 float64 = ((1e10 * 1.5) - 1e10)
                _ = v30
            default:
                var v31 float32 = 0.1
                _ = v31
            }
            fmt.Println(v5, v1, v2)
            _ = v12
        } else {
            var v32 float32 = ((3.14159 - 0.1) + 3.14159)
            var v33 float64 = float64(3.14159)
            _ = v32
            _ = v33
        }
    case 359:
        fmt.Println(v1, v2, v2)
    case 654:
        var v34 uint64 = 2
        _ = v34
    default:
        var v35 string = "foo bar"
        _ = v35
    }
    v1 -= -128
    v1 += (-(v1))
    switch 32768 {
    case 90:
        var v36 bool = false
        _ = v36
    case -1:
        v1 -= -128
    }
    v4 = append(v4, 32768)
    var v37 int8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
