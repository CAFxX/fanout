package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint; F1 int32; F2 int8 }
type GenS3 struct { F0 int64; F1 string; F2 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 uint; F1 uint8 }
type GenS7 struct { F0 int16 }

func main() {
    var v1 uint16 = 252
    var v2 bool = (0.1 == -2.25)
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    v5 := make(map[string]int)
    var v6 uint = 255
    for fi7 := 0; fi7 < 17; fi7++ {
        switch "a" {
        case "a":
            v6 = v6
        }
        _ = fi7
    }
    switch 940 {
    case 128:
        var v8 uint = 42
        _ = v8
    case 0:
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := []int{}
            if len(v10) > 0 {
                v10[0] = 1
            }
            if true {
                for fi11 := 0; fi11 < 8; fi11++ {
                    switch ("" + "foo bar") {
                    case "hello":
                        v12 := GenS1{F0: 127, F1: v6}
                        _ = v12
                    case "a":
                        var v13 int16 = 1
                        _ = v13
                    default:
                        delete(v4, "\n\t")
                    }
                    switch "\n\t" {
                    case "hello":
                        v1 += (v1 | v1)
                    case "foo bar":
                        v14 := GenS1{F0: 255, F1: (v6 | (v6 + v6))}
                        var v15 uint16 = v1
                        _ = v14
                        _ = v15
                    }
                    if v2 {
                        delete(v5, (("hello" + "a") + "foo bar"))
                        v16 := v5[(("\n\t" + "hello") + ("" + "foo bar"))]
                        _ = v16
                    } else {
                        if (42 > 0) {
                            v17 := GenS1{F0: 2, F1: (v6 << ((18446744073709551615) & 63))}
                            if true {
                                v2 = v2
                                switch (^(fi9)) {
                                case 856:
                                    v1 += (^(v1))
                                default:
                                    v1 = 1
                                }
                                v1 = v1
                                v18 := GenS2{F0: v6, F1: 522, F2: 42}
                                _ = v18
                            } else {
                                var v19 int = (fi9 % ((0) | 1))
                                if (!v2) {
                                    v2 = (!v2)
                                } else {
                                    v2 = v2
                                    var v20 string = "hello"
                                    v21 := GenS3{F0: 1, F1: v20, F2: -776}
                                    _ = len(v10)
                                    v6 = 795
                                    _ = v20
                                    _ = v21
                                }
                                v1 += (^(v1))
                                _ = v19
                            }
                            if ((v1 >> (((v6 | 2)) & 63)) >= (v1 + 32767)) {
                                var v22 int32 = 2147483647
                                if (uint64(v1) >= 2) {
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                    v23 := GenS4{F0: 127}
                                    v2 = true
                                    _ = v23
                                }
                                var v24 uint32 = 2147483648
                                _ = v22
                                _ = v24
                            } else {
                                v2 = v2
                                v6 -= 2
                                var v25 int = 9223372036854775807
                                _ = v25
                            }
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            var v26 string = "a"
                            _ = v17
                            _ = v26
                        }
                    }
                    _ = fi11
                }
            }
            _ = fi9
            _ = v10
        }
    case 4294967295:
        var v27 int = 1
        _ = v27
    }
    switch "a" {
    case "foo bar":
        switch -22 {
        case 42:
            for fi28 := 0; fi28 < 17; fi28++ {
                for fi29 := 0; fi29 < 12; fi29++ {
                    v2 = v2
                    switch "hello" {
                    case "foo bar":
                        var v30 int8 = 42
                        _ = v30
                    case "":
                        v2 = (0.0 < float64(0.1))
                    default:
                        var v31 uint64 = 4294967295
                        _ = v31
                    }
                    var v32 float64 = 1e10
                    _ = fi29
                    _ = v32
                }
                for fi33 := 0; fi33 < 5; fi33++ {
                    v34 := GenS1{F0: 127, F1: v6}
                    _ = fi33
                    _ = v34
                }
                var v35 float32 = ((float32((float64(1.5) - float64(0.1))) - (-2.25 - ((3.14159 / 1.5) / float32(-2.25)))) * 3.14159)
                _ = fi28
                _ = v35
            }
        }
    }
    v1 -= (v1 & (v1 >> ((108) & 63)))
    var v36 uint8 = 38
    v2 = true
    v37 := []int{}
    for fi38 := 0; fi38 < 7; fi38++ {
        var v39 uint16 = (v1 + (v1 << ((18446744073709551615) & 63)))
        if ((0 < 65535) || (v2 && (!(v36 <= uint8(v6))))) {
            if (fi38 != (fi38 << (((v6 * (v6 % ((598) | 1)))) & 63))) {
                for fi40 := 0; fi40 < 6; fi40++ {
                    v37 = append(v37, 32767)
                    v37 = append(v37, fi38)
                    v36 -= v36
                    delete(v5, (((("\n\t" + "") + "hello") + "a") + "foo bar"))
                    var v41 uint = (^(v6))
                    if (1 > (v6 >> ((780) & 63))) {
                        v41 += v41
                        v42 := GenS5{F0: (v41 << (((^(v41))) & 63)), F1: (fi38 >> ((v6) & 63))}
                        var v43 uint64 = 470
                        _ = v42
                        _ = v43
                    } else {
                        v36 = (v36 | (v36 ^ (v36 / (((v36 - v36)) | 1))))
                        var v44 uint64 = 0
                        v44 = v44
                        _ = v44
                    }
                    var v45 string = "\n\t"
                    _ = fi40
                    _ = v41
                    _ = v45
                }
                if ("a" != "a") {
                    if v2 {
                        var v46 uint8 = 42
                        _ = v46
                    }
                    if (int8(v6) >= 100) {
                        for fi47 := 0; fi47 < 3; fi47++ {
                            if ((!v2) && (142 <= (v36 + (^(v36))))) {
                                v36 += 0
                                v36 -= v36
                                fmt.Println(v36, fi38, v6)
                            } else {
                                var v48 bool = (225 <= 2147483648)
                                var v49 int8 = 100
                                switch fi38 {
                                case 0:
                                    v50 := GenS6{F0: (^(v6)), F1: 128}
                                    _ = v50
                                case 128:
                                    v49 = -128
                                    v49 += 100
                                }
                                _ = v48
                                _ = v49
                            }
                            var v51 uint8 = (^(v36))
                            _ = fi47
                            _ = v51
                        }
                        for fi52 := 0; fi52 < 10; fi52++ {
                            fmt.Println(v6, v2)
                            _ = fi52
                        }
                        switch ("\n\t" + "") {
                        case "foo bar":
                            fmt.Println(v6, v39, v6)
                        case "hello":
                            var v53 bool = true
                            _ = v53
                        case "\n\t":
                            v54 := GenS7{F0: 1}
                            _ = v54
                        }
                    }
                    v2 = (v2 || (!((v1 << (((v6 >> ((255) & 63))) & 63)) >= v39)))
                } else {
                    v3[uint32(v36)] = 65535
                }
            } else {
                var v55 uint8 = (v36 + 1)
                var v56 bool = (1 > 0)
                _ = v55
                _ = v56
            }
            fmt.Println(v1, v6, v39)
            if (0 != 824) {
                var v57 uint64 = 42
                _ = v57
            } else {
                v58 := GenS5{F0: v6, F1: -2147483648}
                _ = v58
            }
        }
        _ = fi38
        _ = v39
    }
    v37 = append(v37, -695)
    var v59 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v59
    _ = v6
}
