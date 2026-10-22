package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 uint }

func main() {
    v1 := []uint32{2147483647, 42}
    v2 := make(map[int]int)
    var v3 uint = 255
    var v4 int32 = 0
    var v5 string = "hello"
    v6 := v2[42]
    switch v6 {
    case 32767:
        var v7 int8 = 0
        _ = v7
    default:
        if false {
            switch "\n\t" {
            case "a":
                var v8 int = v6
                _ = v8
            case "":
                delete(v2, -88)
            case "\n\t":
                var v9 float64 = (1.5 + 0.0)
                _ = v9
            default:
                var v10 uint64 = 42
                _ = v10
            }
            switch (v6 << (((v3 / (((v3 - v3)) | 1))) & 63)) {
            case 128:
                v3 = v3
            case 127:
                var v11 uint8 = 67
                _ = v11
            case -319:
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    v12 := GenS1{F0: v4}
                    _ = v12
                }
            default:
                v4 = v4
            }
        } else {
            for fi13 := 0; fi13 < 3; fi13++ {
                v5 = v5
                fmt.Println(fi13, v6)
                v6 += (-(v6))
                _ = fi13
            }
            if false {
                for fi14 := 0; fi14 < 5; fi14++ {
                    if false {
                        var v15 bool = (true && (((0.1 - 3.14159) - float32(0.0)) < 3.14159))
                        for fi16 := 0; fi16 < 3; fi16++ {
                            if v15 {
                                v17 := GenS1{F0: (v4 & (v4 | -110))}
                                _ = v17
                            }
                            var v18 float64 = 0.0
                            if true {
                                var v19 float32 = (((1e10 * -2.25) - 3.14159) * 1.5)
                                v20 := GenS1{F0: v4}
                                _ = len(v2)
                                _ = v19
                                _ = v20
                            }
                            _ = fi16
                            _ = v18
                        }
                        if (-1 < 0) {
                            fmt.Println(fi14)
                            fmt.Println(v3, v4)
                            v21 := GenS2{F0: -38, F1: (^(v6))}
                            _ = v21
                        }
                        _ = v15
                    } else {
                        var v22 int8 = 1
                        for fi23 := 0; fi23 < 5; fi23++ {
                            switch (fi14 ^ (v6 / (((v6 & 178)) | 1))) {
                            case 4294967295:
                                v6 -= fi23
                            default:
                                var v24 uint8 = 33
                                _ = v24
                            }
                            if (!true) {
                                v2[(fi23 << ((v3) & 63))] = fi23
                                if true {
                                    v25 := GenS1{F0: 1}
                                    _ = v25
                                } else {
                                    var v26 string = "a"
                                    _ = v26
                                }
                            }
                            delete(v2, (v6 + (v6 | fi23)))
                            _ = fi23
                        }
                        v27 := v2[(fi14 * 9223372036854775807)]
                        _ = v22
                        _ = v27
                    }
                    _ = len(v2)
                    fmt.Println(v3)
                    _ = fi14
                }
                var v28 float64 = 1.5
                if true {
                    if (193 == 166) {
                        if (834 == -913) {
                            var v29 int8 = -128
                            if false {
                                if (0 != 183) {
                                    var v30 uint32 = 504
                                    var v31 uint32 = (^(v30))
                                    _ = v30
                                    _ = v31
                                } else {
                                    v32 := GenS3{F0: 42, F1: 2}
                                    fmt.Println(v28)
                                    if false {
                                        _ = len(v1)
                                        var v33 int8 = (v29 % (((v29 - v29)) | 1))
                                        _ = v33
                                    } else {
                                        v6 = v6
                                        v3 -= (v3 + (v3 & 414))
                                    }
                                    _ = v32
                                }
                            }
                            _ = v29
                        }
                        v2[2147483647] = 1
                        var v34 int = 32768
                        _ = v34
                    } else {
                        var v35 int64 = 9223372036854775807
                        var v36 float64 = 1e10
                        v4 = (v4 >> ((v3) & 63))
                        _ = v35
                        _ = v36
                    }
                } else {
                    var v37 uint64 = 4294967295
                    if len(v1) > 0 {
                        v1[0] = 515
                    }
                    _ = v37
                }
                _ = v28
            }
            var v38 bool = (42 != 32767)
            _ = v38
        }
    }
    v2[(^(v6))] = v6
    for fi39 := 0; fi39 < 4; fi39++ {
        v3 = v3
        if ((v6 + 42) == 4294967295) {
            for fi40 := 0; fi40 < 1; fi40++ {
                switch v6 {
                case 32767:
                    fmt.Println(v6, v5, v5)
                case 4294967295:
                    var v41 int16 = 0
                    _ = v41
                case 128:
                    v42 := GenS3{F0: 42, F1: 255}
                    _ = v42
                default:
                    v43 := GenS2{F0: 42, F1: (fi40 ^ v6)}
                    _ = v43
                }
                v4 = v4
                switch (v5 + v5) {
                case "":
                    v6 += fi40
                case "foo bar":
                    if (true && ((^(fi39)) < v6)) {
                        var v44 int = fi39
                        var v45 int64 = 647
                        v46 := GenS4{F0: int32(v3), F1: 2147483647}
                        _ = v44
                        _ = v45
                        _ = v46
                    } else {
                        v5 = (v5 + "\n\t")
                    }
                case "hello":
                    v47 := GenS4{F0: v4, F1: 0}
                    _ = v47
                default:
                    fmt.Println(fi40)
                }
                _ = fi40
            }
        } else {
            v2[v6] = (fi39 << (((v3 << (((^(v3))) & 63))) & 63))
            v48 := GenS4{F0: (v4 - (v4 | v4)), F1: v4}
            _ = v48
        }
        _ = fi39
    }
    v49 := GenS3{F0: 2, F1: 144}
    v4 = 2147483647
    if (!(-1 > -1)) {
        if false {
            var v50 int16 = -1
            for fi51 := 0; fi51 < 10; fi51++ {
                for fi52 := 0; fi52 < 2; fi52++ {
                    var v53 float32 = (-2.25 + float32((float64(1.5) - (3.14159 * 1.5))))
                    var v54 uint16 = 373
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi52
                    _ = v53
                    _ = v54
                }
                _ = fi51
            }
            delete(v2, v6)
            _ = v50
        } else {
            v55 := GenS2{F0: 823, F1: v6}
            _ = v55
        }
        for fi56 := 0; fi56 < 15; fi56++ {
            for fi57 := 0; fi57 < 20; fi57++ {
                switch -614 {
                case 473:
                    v4 -= v4
                case 128:
                    var v58 uint32 = uint32(v3)
                    _ = v58
                case 4294967295:
                    v5 = ((v5 + (v5 + "a")) + ("a" + "a"))
                }
                _ = fi57
            }
            _ = fi56
        }
        for fi59 := 0; fi59 < 1; fi59++ {
            switch (("" + "hello") + v5) {
            case "":
                v60 := v2[(-(v6))]
                _ = v60
            case "foo bar":
                fmt.Println(v6, fi59, v6)
            }
            v2[-9223372036854775808] = fi59
            v4 -= (v4 + 42)
            _ = fi59
        }
    } else {
        fmt.Println(v6, v5, v5)
        v5 = "\n\t"
        for fi61 := 0; fi61 < 4; fi61++ {
            v62 := GenS5{F0: v3}
            _ = fi61
            _ = v62
        }
    }
    switch (v5 + (v5 + "\n\t")) {
    case "foo bar":
        var v63 int32 = (v4 * (^(v4)))
        _ = v63
    case "a":
        var v64 uint = 255
        _ = v64
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
