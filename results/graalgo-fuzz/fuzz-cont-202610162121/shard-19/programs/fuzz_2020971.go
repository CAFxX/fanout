package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int8; F2 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 int16; F2 int64 }
type GenS5 struct { F0 int8; F1 uint64 }

func main() {
    var v1 uint64 = 352
    v2 := []uint32{176, uint32(v1)}
    v3 := make(map[uint32]int)
    var v4 uint = 18446744073709551615
    switch ("hello" + ("" + ("hello" + (("" + ("" + "a")) + ("a" + "foo bar"))))) {
    case "\n\t":
        v5 := GenS1{F0: 2147483647}
        _ = v5
    case "hello":
        v6 := GenS2{F0: 100, F1: 27, F2: false}
        _ = v6
    }
    if ((1.5 / -2.25) >= 3.14159) {
        if ((!false) && (float32(1.5) <= 3.14159)) {
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v1 = (v1 % ((2) | 1))
                var v7 int8 = 100
                _ = v7
            }
            var v8 float32 = (float32(1.5) * (float32(0.0) - 1e10))
            _ = v8
        } else {
            switch 1 {
            case -2147483648:
                var v9 int8 = 39
                _ = v9
            case 42:
                switch 1 {
                case 2147483647:
                    var v10 uint8 = 0
                    _ = v10
                case -9223372036854775808:
                    v1 -= (v1 >> (((v4 << (((v4 - (v4 / (((v4 * v4)) | 1)))) & 63))) & 63))
                case 1:
                    var v11 int64 = int64(v1)
                    _ = v11
                default:
                    if false {
                        switch 32768 {
                        case -14:
                            switch "" {
                            case "":
                                var v12 int64 = -1
                                _ = v12
                            }
                        case 0:
                            switch 0 {
                            case 32767:
                                v13 := GenS3{F0: ("foo bar" == (("" + "foo bar") + "foo bar"))}
                                _ = v13
                            default:
                                if (42 > 2) {
                                    for fi14 := 0; fi14 < 20; fi14++ {
                                        var v15 int64 = 1
                                        _ = fi14
                                        _ = v15
                                    }
                                    v4 = 4294967295
                                    var v16 float32 = 0.0
                                    _ = v16
                                } else {
                                    v1 = 601
                                }
                            }
                        case 666:
                            v1 = v1
                        }
                    } else {
                        var v17 int64 = -386
                        for fi18 := 0; fi18 < 20; fi18++ {
                            switch (-(fi18)) {
                            case 2147483647:
                                var v19 int64 = v17
                                _ = v19
                            }
                            var v20 uint64 = (^(v1))
                            _ = fi18
                            _ = v20
                        }
                        v21 := GenS4{F0: 52, F1: -954, F2: 42}
                        _ = v17
                        _ = v21
                    }
                }
            default:
                switch "" {
                case "hello":
                    for fi22 := 0; fi22 < 3; fi22++ {
                        if true {
                            v1 += 0
                            for fi23 := 0; fi23 < 17; fi23++ {
                                var v24 uint = v4
                                _ = fi23
                                _ = v24
                            }
                        }
                        var v25 uint64 = 71
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi22
                        _ = v25
                    }
                case "foo bar":
                    fmt.Println(v1)
                case "\n\t":
                    fmt.Println(v1, v1)
                default:
                    v4 -= 4294967295
                }
            }
            v1 -= (v1 << ((v4) & 63))
            var v26 float32 = float32(1e10)
            _ = v26
        }
        var v27 uint = 42
        for fi28 := 0; fi28 < 11; fi28++ {
            for fi29 := 0; fi29 < 6; fi29++ {
                var v30 float32 = (1.5 / float32(1.5))
                v30 -= (v30 + 3.14159)
                var v31 bool = false
                _ = fi29
                _ = v30
                _ = v31
            }
            var v32 uint8 = 1
            _ = fi28
            _ = v32
        }
        _ = v27
    }
    switch "hello" {
    case "hello":
        fmt.Println(v1, v4)
    case "\n\t":
        v4 += 0
    }
    v1 += (v1 * 42)
    for fi33 := 0; fi33 < 7; fi33++ {
        v34 := v3[65535]
        _ = fi33
        _ = v34
    }
    v4 = v4
    v4 = 2
    if false {
        if (1 != 623) {
            v35 := GenS2{F0: -1, F1: 127, F2: true}
            for fi36 := 0; fi36 < 3; fi36++ {
                v37 := GenS5{F0: 100, F1: v1}
                switch (fi36 & 2147483647) {
                case 32767:
                    var v38 uint = v4
                    _ = v38
                case -2147483648:
                    var v39 float64 = (float64(0.0) + -2.25)
                    _ = v39
                }
                _ = len(v2)
                _ = fi36
                _ = v37
            }
            fmt.Println(v1)
            _ = v35
        }
        var v40 string = ("hello" + "")
        _ = len(v2)
        _ = v40
    } else {
        switch ("a" + "foo bar") {
        case "\n\t":
            var v41 int64 = 422
            _ = v41
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v4, v1)
    }
    for fi42 := 0; fi42 < 8; fi42++ {
        var v43 int32 = 2147483647
        var v44 int64 = -980
        _ = fi42
        _ = v43
        _ = v44
    }
    switch ("a" + "hello") {
    case "hello":
        _ = len(v2)
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
