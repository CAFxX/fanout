package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 bool; F1 int8; F2 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32 }

func main() {
    var v1 uint32 = 1
    var v2 int8 = -40
    var v3 int8 = v2
    v4 := make(map[int]uint32)
    v5 := []string{"\n\t"}
    var v6 uint32 = 925
    v7 := GenS1{F0: 2147483647, F1: 32767}
    switch 0 {
    case -1:
        v8 := GenS1{F0: (v1 ^ v6), F1: 111}
        _ = v8
    case 128:
        if (!true) {
            v9 := GenS2{F0: -1}
            if (193 <= 2) {
                for fi10 := 0; fi10 < 9; fi10++ {
                    _ = len(v5)
                    _ = fi10
                }
                switch -9223372036854775808 {
                case 601:
                    var v11 string = "a"
                    _ = v11
                case 128:
                    v12 := GenS1{F0: 950, F1: 574}
                    _ = v12
                case 127:
                    var v13 int8 = v2
                    _ = v13
                default:
                    var v14 uint16 = uint16(v6)
                    _ = v14
                }
                v6 = v6
            } else {
                var v15 float32 = 0.1
                var v16 string = ""
                fmt.Println(v15, v6)
                _ = v15
                _ = v16
            }
            _ = v9
        } else {
            fmt.Println(v6, v6)
            v17 := v4[0]
            _ = v17
        }
    default:
        for fi18 := 0; fi18 < 11; fi18++ {
            if (fi18 >= (fi18 >> ((65535) & 63))) {
                fmt.Println(v1, fi18, v6)
            } else {
                v3 -= (v3 + v2)
                var v19 uint64 = 42
                switch (("" + "") + ((("foo bar" + "foo bar") + "hello") + "a")) {
                case "foo bar":
                    if true {
                        v19 = v19
                        if true {
                            for fi20 := 0; fi20 < 3; fi20++ {
                                var v21 float32 = (1.5 * -2.25)
                                _ = fi20
                                _ = v21
                            }
                            for fi22 := 0; fi22 < 19; fi22++ {
                                var v23 uint16 = 435
                                var v24 uint = 255
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        } else {
                            fmt.Println(v6, v3)
                            v19 -= v19
                            switch "foo bar" {
                            case "":
                                var v25 int16 = 0
                                _ = v25
                            default:
                                v26 := GenS2{F0: 8}
                                _ = v26
                            }
                        }
                    }
                default:
                    v3 += (v3 / ((0) | 1))
                }
                _ = v19
            }
            _ = fi18
        }
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        v28 := GenS2{F0: 0}
        if false {
            var v29 int8 = v3
            for fi30 := 0; fi30 < 13; fi30++ {
                v5 = append(v5, "")
                var v31 string = "a"
                var v32 int64 = 9223372036854775807
                _ = fi30
                _ = v31
                _ = v32
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v29
        } else {
            v3 -= -82
        }
        _ = fi27
        _ = v28
    }
    v33 := GenS3{F0: 779, F1: v3}
    for fi34 := 0; fi34 < 17; fi34++ {
        switch "\n\t" {
        case "a":
            v35 := GenS4{F0: (333 <= 652), F1: (-(v2)), F2: 705}
            _ = v35
        default:
            v36 := GenS5{F0: true}
            _ = v36
        }
        v3 = -111
        var v37 uint = 0
        _ = fi34
        _ = v37
    }
    var v38 int = 721
    if (9223372036854775807 <= 9223372036854775807) {
        fmt.Println(v1)
    }
    var v39 int64 = 1
    for fi40 := 0; fi40 < 15; fi40++ {
        switch v38 {
        case 32767:
            var v41 int64 = v39
            _ = v41
        case 1:
            var v42 uint = 174
            _ = v42
        default:
            for fi43 := 0; fi43 < 12; fi43++ {
                var v44 uint32 = v1
                v39 = 687
                switch "" {
                case "a":
                    if (845 <= 18446744073709551615) {
                        var v45 int32 = 1
                        v44 -= (^(v6))
                        switch "hello" {
                        case "":
                            if (!(516 == 1)) {
                                v39 += -9223372036854775808
                                v46 := GenS6{F0: v44}
                                _ = v46
                            } else {
                                v45 -= v45
                                var v47 uint = 461
                                _ = v47
                            }
                        case "\n\t":
                            v38 -= (fi43 % ((int(v45)) | 1))
                        }
                        _ = v45
                    } else {
                        if (2147483647 < -2147483648) {
                            var v48 uint = 18446744073709551615
                            v39 -= v39
                            _ = v48
                        }
                    }
                default:
                    var v49 int32 = int32(v44)
                    _ = v49
                }
                _ = fi43
                _ = v44
            }
        }
        _ = fi40
    }
    var v50 int32 = 1
    var v51 float32 = float32((0.0 * 1e10))
    v50 = 2147483647
    v38 = (v38 - (v38 | 4294967295))
    v39 -= v39
    var v52 string = "hello"
    v52 = v52
    for fi53 := 0; fi53 < 17; fi53++ {
        var v54 string = v52
        _ = fi53
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
    _ = v7
}
