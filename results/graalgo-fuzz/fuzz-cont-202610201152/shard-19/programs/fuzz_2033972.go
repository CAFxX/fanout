package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint8; F2 int16 }
type GenS3 struct { F0 int8; F1 string }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    var v4 float32 = (1.5 - float32((v1 / v1)))
    v5 := make(map[int]int)
    v6 := make(map[uint32]int)
    if ((2 == 255) || true) {
        var v7 uint8 = 255
        for fi8 := 0; fi8 < 17; fi8++ {
            for fi9 := 0; fi9 < 8; fi9++ {
                v7 = (v7 % (((v7 | 239)) | 1))
                _ = fi9
            }
            _ = fi8
        }
        _ = v7
    } else {
        v10 := v6[0]
        _ = v10
    }
    var v11 int8 = 1
    v12 := []int{32768}
    for fi13 := 0; fi13 < 20; fi13++ {
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 int8 = v11
            if (1 != 142) {
                var v16 uint = 42
                _ = v16
            } else {
                var v17 int8 = (v15 + v15)
                _ = v17
            }
            var v18 uint = 1
            _ = fi14
            _ = v15
            _ = v18
        }
        v19 := GenS1{F0: 2, F1: 128}
        switch "foo bar" {
        case "foo bar":
            var v20 int = fi13
            _ = v20
        case "a":
            if (1 <= 255) {
                v4 = v4
                switch "foo bar" {
                case "\n\t":
                    v21 := GenS2{F0: 2147483647, F1: 42, F2: 42}
                    _ = v21
                default:
                    var v22 float32 = 1e10
                    _ = v22
                }
                v11 -= 100
            } else {
                for fi23 := 0; fi23 < 4; fi23++ {
                    switch fi23 {
                    case 42:
                        v24 := GenS2{F0: 350, F1: 255, F2: int16(v11)}
                        _ = v24
                    default:
                        v25 := GenS1{F0: 4294967295, F1: 42}
                        _ = v25
                    }
                    v4 += v4
                    _ = fi23
                }
                var v26 string = ("" + ("" + ("a" + "foo bar")))
                v27 := GenS2{F0: -1, F1: 197, F2: 32767}
                _ = v26
                _ = v27
            }
        case "\n\t":
            v4 += float32((v1 / 1e10))
        }
        _ = fi13
        _ = v19
    }
    if (!(true || (false && true))) {
        var v28 float32 = v4
        var v29 uint16 = 42
        _ = v28
        _ = v29
    }
    switch "foo bar" {
    case "a":
        v1 += (-2.25 * v1)
    case "foo bar":
        if (!(((65535 >= 2147483648) && (112 <= 2)) || false)) {
            v30 := v2[128]
            if false {
                fmt.Println(v30)
                for fi31 := 0; fi31 < 11; fi31++ {
                    v30 = 103
                    v32 := GenS2{F0: 1, F1: 42, F2: -32768}
                    v1 = ((v1 + 1.5) + float64((v4 - (3.14159 - v4))))
                    _ = fi31
                    _ = v32
                }
            }
            if (-1 != v11) {
                if (1 <= 42) {
                    v11 = (v11 >> ((65535) & 63))
                    switch "hello" {
                    case "hello":
                        if false {
                            if len(v12) > 0 {
                                v12[0] = v30
                            }
                            v12 = append(v12, 32768)
                            var v33 uint64 = 0
                            _ = v33
                        } else {
                            for fi34 := 0; fi34 < 19; fi34++ {
                                delete(v5, (v30 / (((^(v30))) | 1)))
                                v11 -= v11
                                for fi35 := 0; fi35 < 13; fi35++ {
                                    var v36 string = "hello"
                                    if false {
                                        switch (fi35 >> ((255) & 63)) {
                                        case 0:
                                            fmt.Println(fi34, v30, fi35)
                                        case 9223372036854775807:
                                            fmt.Println(v11)
                                        }
                                        v37 := GenS1{F0: 2, F1: 2}
                                        _ = v37
                                    }
                                    _ = fi35
                                    _ = v36
                                }
                                _ = fi34
                            }
                            var v38 uint = 308
                            _ = v38
                        }
                    default:
                        fmt.Println(v1, v4, v4)
                    }
                } else {
                    if len(v12) > 0 {
                        v12[0] = (v30 - v30)
                    }
                }
                v1 += 0.1
                var v39 uint32 = 65535
                _ = v39
            } else {
                if true {
                    if false {
                        fmt.Println(v30)
                    }
                } else {
                    v40 := v2[127]
                    _ = v40
                }
                fmt.Println(v11)
            }
            _ = v30
        } else {
            var v41 float64 = (v1 * (v1 / ((v1 / 1e10) - v1)))
            for fi42 := 0; fi42 < 3; fi42++ {
                if (!(fi42 < fi42)) {
                    var v43 int8 = v11
                    _ = v43
                }
                if (702 <= 781) {
                    v1 = (v1 / v1)
                    switch (fi42 >> ((4294967295) & 63)) {
                    case -288:
                        fmt.Println(v11, v11)
                    }
                }
                v44 := GenS1{F0: 18446744073709551615, F1: 172}
                _ = fi42
                _ = v44
            }
            _ = v41
        }
    case "":
        v11 = (v11 << ((0) & 63))
    }
    v11 = v11
    switch 128 {
    case 128:
        fmt.Println(v4)
    }
    v11 = 0
    fmt.Println(v11, v1)
    v11 = (^(v11))
    v45 := GenS2{F0: 1, F1: 127, F2: 525}
    var v46 float64 = 0.0
    v46 = 0.0
    switch "\n\t" {
    case "a":
        fmt.Println(v11, v11)
    case "\n\t":
        if false {
            v47 := GenS3{F0: (v11 | (v11 * v11)), F1: "foo bar"}
            v11 = 0
            var v48 uint = 604
            _ = v47
            _ = v48
        }
    default:
        fmt.Println(v11, v1, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
