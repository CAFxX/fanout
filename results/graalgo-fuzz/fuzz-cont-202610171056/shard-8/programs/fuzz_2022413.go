package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int8; F1 uint32 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 255
    v3 := []uint32{964}
    v4 := []string{""}
    v2 = v2
    if ("a" < ("a" + (("foo bar" + "\n\t") + ""))) {
        v2 += 128
    }
    var v5 uint32 = 2
    if false {
        v4 = append(v4, "")
        if (false && false) {
            v2 = (v2 * v2)
        } else {
            fmt.Println(v2)
            switch "" {
            case "\n\t":
                for fi6 := 0; fi6 < 8; fi6++ {
                    v7 := GenS1{F0: ((("hello" + "a") + "") + ("\n\t" + "foo bar")), F1: v2}
                    for fi8 := 0; fi8 < 10; fi8++ {
                        fmt.Println(fi8, fi6, fi6)
                        _ = fi8
                    }
                    _ = fi6
                    _ = v7
                }
            case "":
                for fi9 := 0; fi9 < 12; fi9++ {
                    var v10 uint8 = v2
                    var v11 float64 = 1e10
                    switch "foo bar" {
                    case "a":
                        var v12 int8 = 0
                        _ = v12
                    case "foo bar":
                        if len(v3) > 0 {
                            v3[0] = 4294967295
                        }
                    case "":
                        var v13 uint = 2
                        _ = v13
                    }
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            default:
                fmt.Println(v5, v2)
            }
            v2 -= 0
        }
    } else {
        var v14 float32 = 1.5
        v5 = 65535
        _ = v14
    }
    v2 += 2
    fmt.Println(v2)
    for fi15 := 0; fi15 < 4; fi15++ {
        switch 42 {
        case 1:
            v2 += v2
        case 42:
            if (!false) {
                if (!(float64(0.0) <= 0.0)) {
                    fmt.Println(v2, v5)
                    switch "a" {
                    case "hello":
                        if (2 != 399) {
                            for fi16 := 0; fi16 < 17; fi16++ {
                                var v17 bool = true
                                _ = len(v3)
                                if true {
                                    fmt.Println(fi15, fi15, fi15)
                                    v5 += (v5 >> ((1) & 63))
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v18 int16 = -32768
                                        _ = v18
                                    case "foo bar":
                                        v19 := GenS1{F0: ("\n\t" + ("foo bar" + "a")), F1: (v2 | v2)}
                                        _ = v19
                                    default:
                                        v2 = v2
                                    }
                                }
                                _ = fi16
                                _ = v17
                            }
                        } else {
                            v20 := GenS1{F0: "\n\t", F1: v2}
                            _ = v20
                        }
                    case "a":
                        v2 = (v2 / ((v2) | 1))
                    default:
                        var v21 uint16 = 42
                        _ = v21
                    }
                    switch "\n\t" {
                    case "hello":
                        v1[32767] = (42 >= 0)
                    default:
                        var v22 float64 = float64(1e10)
                        _ = v22
                    }
                } else {
                    if true {
                        v23 := GenS2{F0: 949, F1: 100}
                        var v24 bool = true
                        _ = v23
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 12; fi25++ {
                            if false {
                                fmt.Println(v5)
                                fmt.Println(fi15)
                                if ((("" + "foo bar") + "hello") >= "hello") {
                                    v26 := GenS3{F0: 1, F1: uint32(v2)}
                                    v27 := GenS3{F0: -1, F1: v5}
                                    _ = v26
                                    _ = v27
                                } else {
                                    if ("foo bar" >= "") {
                                        var v28 int8 = 108
                                        v2 -= (v2 >> ((18446744073709551615) & 63))
                                        switch (^(fi15)) {
                                        case 0:
                                            _ = len(v1)
                                        case 348:
                                            delete(v1, (fi25 / (((fi25 / ((fi15) | 1))) | 1)))
                                        case 127:
                                            fmt.Println(v2, fi25, v2)
                                        default:
                                            v2 += 1
                                        }
                                        _ = v28
                                    }
                                }
                            } else {
                                var v29 bool = true
                                fmt.Println(fi25, v29, fi25)
                                _ = v29
                            }
                            v5 += (v5 | (v5 >> ((65535) & 63)))
                            switch fi25 {
                            case -2147483648:
                                fmt.Println(v2, v5)
                            case 42:
                                fmt.Println(fi15, fi15)
                            default:
                                if false {
                                    var v30 uint64 = 69
                                    _ = v30
                                } else {
                                    var v31 int32 = -572
                                    v31 = (v31 - -2147483648)
                                    fmt.Println(v31)
                                    _ = v31
                                }
                            }
                            _ = fi25
                        }
                    }
                    v32 := GenS2{F0: 1, F1: 101}
                    for fi33 := 0; fi33 < 19; fi33++ {
                        v34 := GenS3{F0: -1, F1: (v5 / (((v5 * v5)) | 1))}
                        v2 = (v2 * v2)
                        _ = fi33
                        _ = v34
                    }
                    _ = v32
                }
                v35 := GenS3{F0: 127, F1: (^(v5))}
                _ = v35
            }
        case -732:
            var v36 int8 = -1
            _ = v36
        }
        v2 = (v2 / (((v2 ^ (^(v2)))) | 1))
        _ = fi15
    }
    v2 = uint8(v5)
    var v37 uint = 42
    fmt.Println("maplen v1:", len(v1))
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
