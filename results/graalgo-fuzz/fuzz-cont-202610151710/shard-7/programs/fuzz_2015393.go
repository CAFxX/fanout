package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint64; F1 string; F2 string }
type GenS3 struct { F0 int32; F1 int64 }

func main() {
    var v1 uint32 = 2
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    for fi4 := 0; fi4 < 17; fi4++ {
        for fi5 := 0; fi5 < 2; fi5++ {
            v1 = v1
            v6 := GenS1{F0: 2, F1: 2147483647}
            _ = fi5
            _ = v6
        }
        var v7 float32 = (-2.25 / (-2.25 / 3.14159))
        _ = fi4
        _ = v7
    }
    if ((false && true) || (100 >= 42)) {
        v8 := []int{824, 4294967295, -177}
        var v9 uint32 = v1
        _ = v8
        _ = v9
    }
    delete(v3, -9223372036854775808)
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 6; fi11++ {
            v1 = (v1 - (v1 % (((^(v1))) | 1)))
            switch "" {
            case "hello":
                v1 = v1
                v1 = v1
                var v12 uint64 = 1
                _ = v12
            }
            v1 += v1
            _ = fi11
        }
        _ = fi10
    }
    switch "foo bar" {
    case "foo bar":
        if ((^(v1)) >= (v1 >> ((255) & 63))) {
            var v13 int64 = 0
            for fi14 := 0; fi14 < 14; fi14++ {
                v15 := GenS2{F0: 761, F1: "", F2: (("hello" + "a") + "")}
                _ = fi14
                _ = v15
            }
            v1 = (v1 / ((v1) | 1))
            _ = v13
        } else {
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v1)
            default:
                v1 = v1
            }
            var v16 float32 = 0.0
            if false {
                fmt.Println(v1, v16)
                v16 += float32(((1.5 - 0.0) / float64((3.14159 / v16))))
                for fi17 := 0; fi17 < 8; fi17++ {
                    if (198 != 42) {
                        var v18 uint64 = 4294967295
                        var v19 uint64 = (v18 / ((v18) | 1))
                        v20 := GenS1{F0: 4, F1: 2147483647}
                        _ = v18
                        _ = v19
                        _ = v20
                    } else {
                        var v21 int64 = 0
                        v3[fi17] = fi17
                        if true {
                            if true {
                                if false {
                                    var v22 uint64 = uint64(v1)
                                    v22 = v22
                                    if true {
                                        var v23 int16 = 1
                                        if false {
                                            fmt.Println(v16, v1)
                                        } else {
                                            v24 := v2["foo bar"]
                                            _ = v24
                                        }
                                        var v25 int16 = v23
                                        _ = v23
                                        _ = v25
                                    }
                                    _ = v22
                                } else {
                                    v26 := GenS2{F0: 72, F1: "\n\t", F2: "hello"}
                                    switch fi17 {
                                    case 733:
                                        v27 := GenS1{F0: 2, F1: -223}
                                        _ = v27
                                    }
                                    _ = v26
                                }
                                v21 = 361
                                switch "a" {
                                case "hello":
                                    var v28 int = fi17
                                    _ = v28
                                default:
                                    v1 += (^(v1))
                                }
                            } else {
                                v1 = (v1 + (v1 & (^(v1))))
                            }
                        }
                        _ = v21
                    }
                    _ = fi17
                }
            }
            _ = v16
        }
    case "a":
        v1 += (v1 | v1)
    default:
        switch -14 {
        case 9223372036854775807:
            v1 = v1
        case 0:
            fmt.Println(v1)
        }
    }
    var v29 int = 0
    v3[v29] = v29
    for fi30 := 0; fi30 < 9; fi30++ {
        for fi31 := 0; fi31 < 12; fi31++ {
            fmt.Println(fi31, fi31, fi31)
            v32 := GenS3{F0: -133, F1: 1}
            if (true && false) {
                var v33 uint8 = 35
                v34 := GenS2{F0: 18446744073709551615, F1: "", F2: "a"}
                _ = v33
                _ = v34
            }
            _ = fi31
            _ = v32
        }
        v35 := GenS1{F0: uint8(v1), F1: 870}
        if (!(!false)) {
            v3[9223372036854775807] = v29
        } else {
            if true {
                var v36 uint = 525
                if true {
                    fmt.Println(v1)
                }
                _ = v36
            }
            _ = len(v3)
            v3[-93] = (fi30 * 128)
        }
        _ = fi30
        _ = v35
    }
    var v37 uint = 18446744073709551615
    var v38 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v38
}
