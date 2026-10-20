package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint64; F1 bool; F2 int8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint, p1 uint16) uint {
    p1 = p1
    v7 := []int{-1}
    for fi8 := 0; fi8 < 4; fi8++ {
        if true {
            if len(v7) > 0 {
                v7[0] = fi8
            }
            v9 := GenS1{F0: (fi8 + (fi8 >> ((2) & 63))), F1: 980, F2: 40}
            _ = v9
        } else {
            fmt.Println(fi8, fi8, p1)
            fmt.Println(fi8)
        }
        _ = fi8
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = v7
    return (p0 | (p0 - 255))
}
func genFn1(p0 uint, p1 uint32) uint {
    if false {
        p1 = p1
        v10 := []int{-2147483648, 200, 2147483647}
        fmt.Println(p0)
        _ = v10
    }
    v11 := GenS2{F0: 2, F1: true, F2: 42}
    var v12 bool = (!false)
    var v13 int64 = 1
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return genFn0(42, 0)
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[string]uint32)
    v3 := []string{}
    v4 := make(map[uint32]string)
    v5 := make(map[string]string)
    var v6 int32 = -2147483648
    v3 = append(v3, ("foo bar" + ("" + "\n\t")))
    var v14 uint64 = 42
    v1 -= v1
    v3 = append(v3, ((("" + "") + "\n\t") + "a"))
    var v15 uint32 = 2
    fmt.Println(v15, v6, v1)
    if ((int16(v1) != 1000) && false) {
        switch 0 {
        case 42:
            v1 = 1
        }
        switch ("" + "foo bar") {
        case "foo bar":
            v5["foo bar"] = "foo bar"
        case "hello":
            if false {
                v16 := GenS3{F0: 489}
                _ = v16
            } else {
                var v17 int8 = -99
                var v18 int8 = -128
                _ = v17
                _ = v18
            }
        case "a":
            if (genFn1(uint(v15), (v15 >> ((178) & 63))) <= 255) {
                for fi19 := 0; fi19 < 16; fi19++ {
                    switch "hello" {
                    case "\n\t":
                        if false {
                            var v20 float32 = 0.0
                            var v21 uint64 = v14
                            var v22 uint16 = v1
                            _ = v20
                            _ = v21
                            _ = v22
                        } else {
                            for fi23 := 0; fi23 < 16; fi23++ {
                                v24 := GenS3{F0: -936}
                                _ = fi23
                                _ = v24
                            }
                            switch fi19 {
                            case -1:
                                fmt.Println(v1, v15, v6)
                            case 32767:
                                v6 = v6
                            case -2147483648:
                                v25 := GenS3{F0: -32768}
                                _ = v25
                            default:
                                switch "\n\t" {
                                case "":
                                    for fi26 := 0; fi26 < 6; fi26++ {
                                        v1 += v1
                                        v27 := GenS1{F0: fi19, F1: v15, F2: (v1 / (((v1 % (((v1 - v1)) | 1))) | 1))}
                                        v14 -= v14
                                        _ = fi26
                                        _ = v27
                                    }
                                case "hello":
                                    v1 = 255
                                case "\n\t":
                                    for fi28 := 0; fi28 < 11; fi28++ {
                                        if len(v3) > 0 {
                                            v3[0] = ("hello" + ("foo bar" + (("a" + "\n\t") + "hello")))
                                        }
                                        v29 := GenS2{F0: (^(v14)), F1: (v6 > v6), F2: -1}
                                        v30 := GenS1{F0: 32767, F1: (^(v15)), F2: 255}
                                        _ = fi28
                                        _ = v29
                                        _ = v30
                                    }
                                }
                            }
                        }
                    case "":
                        var v31 int32 = 865
                        _ = v31
                    case "hello":
                        switch (("\n\t" + "hello") + ("" + "hello")) {
                        case "foo bar":
                            fmt.Println(v14, fi19)
                        case "a":
                            fmt.Println(v6, v6)
                        }
                    }
                    v32 := GenS1{F0: fi19, F1: (v15 / (((v15 % ((v15) | 1))) | 1)), F2: v1}
                    _ = fi19
                    _ = v32
                }
                for fi33 := 0; fi33 < 20; fi33++ {
                    switch fi33 {
                    case 42:
                        var v34 bool = true
                        _ = v34
                    case 0:
                        v14 += (^(v14))
                    }
                    var v35 int32 = v6
                    _ = fi33
                    _ = v35
                }
            }
        default:
            v36 := v5[(("foo bar" + "") + ("\n\t" + "\n\t"))]
            _ = v36
        }
    }
    v4[v15] = "foo bar"
    var v37 uint32 = 0
    v15 += (v37 << ((0) & 63))
    _ = len(v2)
    if ((0 >= 255) && (-128 < 0)) {
        var v38 int = 4294967295
        switch (("" + ("" + "a")) + "a") {
        case "\n\t":
            v39 := GenS3{F0: 1000}
            _ = v39
        case "":
            v37 += (v15 >> ((255) & 63))
        default:
            var v40 int8 = 127
            _ = v40
        }
        for fi41 := 0; fi41 < 18; fi41++ {
            v42 := GenS2{F0: 18446744073709551615, F1: (!((^(v1)) != 42)), F2: 118}
            fmt.Println(v38)
            var v43 float64 = 0.1
            _ = fi41
            _ = v42
            _ = v43
        }
        _ = v38
    }
    v5["hello"] = (("" + "") + ("\n\t" + ("\n\t" + "\n\t")))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
