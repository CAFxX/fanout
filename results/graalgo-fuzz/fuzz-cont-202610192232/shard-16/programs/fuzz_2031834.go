package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 uint16; F2 int }
func genFn0(p0 int16) int8 {
    fmt.Println(p0)
    p0 -= p0
    _ = p0
    return 17
}
func genFn1(p0 uint32, p1 uint32, p2 int) int16 {
    var v5 int32 = 556
    switch (p2 ^ (-(p2))) {
    case 2:
        if ((true && ((p1 / (((p0 - (p0 >> ((255) & 63)))) | 1)) == (p1 / ((1) | 1)))) || false) {
            var v6 float32 = ((((3.14159 * 1.5) + (0.1 + 0.1)) / 1.0) * 0.0)
            _ = v6
        } else {
            if true {
                var v7 int8 = 1
                v8 := GenS1{F0: false}
                var v9 int32 = 42
                _ = v7
                _ = v8
                _ = v9
            }
            v10 := []int{}
            _ = len(v10)
            _ = v10
        }
    }
    v11 := []int{p2, p2, p2, (p2 % ((p2) | 1))}
    var v12 bool = (128 < 246)
    p1 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v5
    return 710
}

func main() {
    v1 := make(map[string]int)
    var v2 int = -1
    v3 := make(map[uint32]bool)
    v4 := []string{}
    if (1 >= 2) {
        if (!((1.5 - float64(3.14159)) <= -2.25)) {
            var v13 float32 = (-2.25 * (1e10 / float32(1.5)))
            if true {
                var v14 bool = false
                switch "hello" {
                case "":
                    var v15 int = 128
                    _ = v15
                default:
                    if (-97 != 100) {
                        if false {
                            delete(v1, "\n\t")
                            v14 = (2 != 357)
                        } else {
                            v2 = (v2 >> ((786) & 63))
                            if ((1000 == genFn1(65535, 2147483648, (v2 % (((v2 - 329)) | 1)))) && (-599 == (v2 ^ (v2 - (v2 + v2))))) {
                                v16 := GenS2{F0: 18446744073709551615}
                                fmt.Println(v13, v13)
                                var v17 int32 = 42
                                _ = v16
                                _ = v17
                            }
                            for fi18 := 0; fi18 < 15; fi18++ {
                                fmt.Println(v14, v14)
                                v14 = (1 <= 168)
                                _ = fi18
                            }
                        }
                    } else {
                        v19 := GenS1{F0: (1e10 > 3.14159)}
                        v14 = v14
                        _ = v19
                    }
                }
                if true {
                    v2 -= 1
                } else {
                    if v14 {
                        v14 = false
                        var v20 uint64 = 112
                        _ = v20
                    }
                    if len(v4) > 0 {
                        v4[0] = "foo bar"
                    }
                    for fi21 := 0; fi21 < 15; fi21++ {
                        for fi22 := 0; fi22 < 1; fi22++ {
                            var v23 int = 0
                            if (-32768 > -393) {
                                switch "hello" {
                                case "a":
                                    var v24 int = 874
                                    _ = v24
                                case "hello":
                                    v25 := GenS1{F0: v14}
                                    _ = v25
                                }
                                v2 += v2
                            }
                            var v26 float64 = 3.14159
                            _ = fi22
                            _ = v23
                            _ = v26
                        }
                        delete(v3, 65535)
                        _ = fi21
                    }
                }
                _ = v14
            } else {
                var v27 uint8 = 1
                _ = v27
            }
            _ = v13
        } else {
            v2 = (^(v2))
            var v28 int8 = 0
            if len(v4) > 0 {
                v4[0] = "a"
            }
            _ = v28
        }
        switch ("" + "foo bar") {
        case "hello":
            fmt.Println(v2)
        case "a":
            v29 := GenS1{F0: (42 == 127)}
            _ = v29
        }
        delete(v3, 2147483648)
    }
    fmt.Println(v2)
    v2 += 32767
    fmt.Println(v2)
    switch (v2 ^ (v2 ^ (v2 >> ((65535) & 63)))) {
    case 4294967295:
        switch "foo bar" {
        case "hello":
            var v30 uint = 398
            _ = v30
        case "\n\t":
            var v31 uint8 = 128
            _ = v31
        case "":
            var v32 string = "foo bar"
            _ = v32
        }
    case -9223372036854775808:
        v2 = (v2 * v2)
    default:
        v2 = v2
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        if (0 <= 0) {
            if (255 < 0) {
                var v34 int = fi33
                v35 := GenS3{F0: 121, F1: 1, F2: (v34 << ((65535) & 63))}
                switch ("a" + "") {
                case "\n\t":
                    var v36 int32 = 884
                    _ = v36
                case "":
                    v34 += (v2 + (v2 + 32768))
                case "foo bar":
                    v37 := GenS3{F0: genFn0(-32768), F1: 463, F2: (fi33 << ((317) & 63))}
                    _ = v37
                }
                _ = v34
                _ = v35
            } else {
                var v38 float32 = 1.5
                _ = v38
            }
        }
        _ = fi33
    }
    v2 = (^(v2))
    for fi39 := 0; fi39 < 14; fi39++ {
        v2 -= v2
        for fi40 := 0; fi40 < 7; fi40++ {
            if (42 < 2147483648) {
                v2 += v2
                if false {
                    switch ("" + "hello") {
                    case "hello":
                        v41 := GenS1{F0: (!true)}
                        _ = v41
                    case "\n\t":
                        var v42 float64 = -2.25
                        _ = v42
                    }
                }
            }
            v43 := v3[42]
            _ = fi40
            _ = v43
        }
        _ = fi39
    }
    if (561 != -9223372036854775808) {
        var v44 string = "\n\t"
        _ = v44
    }
    fmt.Println(v2)
    for fi45 := 0; fi45 < 14; fi45++ {
        fmt.Println(fi45, v2)
        _ = fi45
    }
    v46 := GenS3{F0: -1, F1: 255, F2: 127}
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
}
