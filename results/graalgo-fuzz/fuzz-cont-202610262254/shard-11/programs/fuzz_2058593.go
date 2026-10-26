package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int16 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 int16; F2 uint8 }
type GenS5 struct { F0 string; F1 int }

func main() {
    var v1 int32 = 512
    v2 := []string{}
    v3 := make(map[string]string)
    v4 := make(map[uint32]bool)
    v5 := []int{560, -434, -545}
    v1 = (v1 | v1)
    var v6 int16 = 32767
    fmt.Println(v1, v1)
    v1 = 1
    v6 = 42
    for fi7 := 0; fi7 < 8; fi7++ {
        if false {
            v6 = (v6 & (v6 - v6))
        } else {
            switch ("a" + "foo bar") {
            case "foo bar":
                v6 += (v6 + (v6 * -1))
            default:
                fmt.Println(v1, v1, fi7)
            }
            var v8 int8 = 100
            if (!false) {
                _ = len(v3)
                if (4294967295 == 491) {
                    for fi9 := 0; fi9 < 2; fi9++ {
                        var v10 uint32 = 1
                        fmt.Println(v8, v10)
                        if true {
                            var v11 string = ("\n\t" + "foo bar")
                            fmt.Println(fi7, v1, v8)
                            _ = v11
                        } else {
                            v12 := GenS1{F0: 799, F1: 0, F2: (v6 / ((1) | 1))}
                            if (18446744073709551615 > 803) {
                                var v13 bool = (!((!(!true)) || false))
                                var v14 string = "foo bar"
                                var v15 int = (fi9 >> ((2) & 63))
                                _ = v13
                                _ = v14
                                _ = v15
                            } else {
                                switch ("" + "\n\t") {
                                case "hello":
                                    var v16 int = (fi9 * (fi7 & (fi9 + (fi7 >> ((255) & 63)))))
                                    _ = v16
                                case "":
                                    fmt.Println(v8)
                                    var v17 int8 = (v8 + (v8 % ((-128) | 1)))
                                    v18 := GenS1{F0: 1, F1: 2147483647, F2: v6}
                                    v6 -= (v6 / (((v6 * -282)) | 1))
                                    _ = v17
                                    _ = v18
                                case "\n\t":
                                    _ = len(v4)
                                default:
                                    switch 32768 {
                                    case 9223372036854775807:
                                        switch fi7 {
                                        case 4294967295:
                                            var v19 float32 = ((3.14159 + float32((-2.25 / 1.5))) / 0.1)
                                            _ = v19
                                        case 9223372036854775807:
                                            var v20 uint16 = 255
                                            _ = v20
                                        default:
                                            if (true || ((v6 | int16(v1)) <= v6)) {
                                                fmt.Println(v8)
                                            }
                                        }
                                        v21 := GenS2{F0: (v1 >> ((1) & 63)), F1: v10}
                                        switch "foo bar" {
                                        case "":
                                            var v22 uint32 = v10
                                            _ = v22
                                        case "hello":
                                            v6 = v6
                                        default:
                                            var v23 int64 = 1
                                            _ = v23
                                        }
                                        _ = v21
                                    default:
                                        delete(v3, "foo bar")
                                    }
                                }
                            }
                            switch fi9 {
                            case -1:
                                v24 := GenS2{F0: 2147483647, F1: 2147483648}
                                _ = v24
                            case 2147483647:
                                switch "foo bar" {
                                case "foo bar":
                                    var v25 uint8 = 42
                                    _ = v25
                                case "\n\t":
                                    _ = len(v5)
                                case "":
                                    _ = len(v4)
                                default:
                                    if ((fi7 >> ((4294967295) & 63)) < 32767) {
                                        switch "" {
                                        case "hello":
                                            v8 += (v8 ^ (-(v8)))
                                            v10 -= (v10 << ((402) & 63))
                                        case "a":
                                            v6 += v6
                                        }
                                        v10 += (v10 + v10)
                                    } else {
                                        v8 -= (-(v8))
                                    }
                                }
                            default:
                                switch "hello" {
                                case "\n\t":
                                    v26 := GenS2{F0: v1, F1: (v10 + (v10 % (((v10 & 63)) | 1)))}
                                    _ = v26
                                case "a":
                                    var v27 uint8 = 87
                                    _ = v27
                                default:
                                    if len(v2) > 0 {
                                        v2[0] = "hello"
                                    }
                                }
                            }
                            var v28 int16 = -73
                            v29 := GenS2{F0: v1, F1: v10}
                            _ = v12
                            _ = v28
                            _ = v29
                        }
                        v30 := GenS3{F0: 143}
                        if false {
                            var v31 int16 = v6
                            v31 -= v31
                            v32 := GenS1{F0: 2, F1: (-(fi7)), F2: 42}
                            _ = v31
                            _ = v32
                        } else {
                            switch (^(fi9)) {
                            case -445:
                                if (!false) {
                                    var v33 int = fi9
                                    v34 := GenS3{F0: 65535}
                                    _ = v33
                                    _ = v34
                                }
                            }
                            switch fi7 {
                            case 9223372036854775807:
                                switch "foo bar" {
                                case "a":
                                    if len(v2) > 0 {
                                        v2[0] = ("foo bar" + "\n\t")
                                    }
                                }
                            case -2147483648:
                                v1 += 0
                            case 4294967295:
                                var v35 uint64 = 42
                                _ = v35
                            default:
                                var v36 uint16 = 42
                                _ = v36
                            }
                        }
                        v37 := GenS4{F0: (v8 >> ((2) & 63)), F1: (v6 + int16(v8)), F2: 128}
                        _ = fi9
                        _ = v10
                        _ = v30
                        _ = v37
                    }
                } else {
                    var v38 bool = (0.1 >= 3.14159)
                    var v39 float32 = 3.14159
                    _ = v38
                    _ = v39
                }
                switch fi7 {
                case 23:
                    v40 := GenS5{F0: ("hello" + "\n\t"), F1: (fi7 + (fi7 / (((fi7 ^ fi7)) | 1)))}
                    _ = v40
                case 880:
                    switch ("a" + "foo bar") {
                    case "\n\t":
                        var v41 float32 = 1e10
                        _ = v41
                    }
                default:
                    if (!false) {
                        var v42 string = "hello"
                        for fi43 := 0; fi43 < 12; fi43++ {
                            switch v42 {
                            case "":
                                var v44 int8 = (-(v8))
                                _ = v44
                            case "\n\t":
                                fmt.Println(v8, v1)
                            case "a":
                                var v45 float32 = 0.0
                                _ = v45
                            default:
                                switch ((v42 + ("a" + v42)) + (v42 + v42)) {
                                case "":
                                    fmt.Println(fi43)
                                case "\n\t":
                                    var v46 float32 = (float32(1e10) - (0.0 - ((3.14159 / 0.1) / float32(1e10))))
                                    _ = v46
                                }
                            }
                            _ = fi43
                        }
                        _ = v42
                    }
                }
            }
            _ = v8
        }
        v47 := GenS5{F0: "foo bar", F1: (-(fi7))}
        v6 -= (v6 * v6)
        _ = fi7
        _ = v47
    }
    var v48 float64 = 3.14159
    v1 -= (v1 % ((1) | 1))
    delete(v3, ("foo bar" + "hello"))
    for fi49 := 0; fi49 < 2; fi49++ {
        v6 += 896
        var v50 uint = 315
        _ = fi49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
