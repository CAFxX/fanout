package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int; F1 uint; F2 uint }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint; F1 uint }

func main() {
    v1 := []int{-1}
    var v2 int16 = 1
    var v3 bool = (-9223372036854775808 <= 42)
    v4 := make(map[uint32]bool)
    var v5 bool = v3
    v6 := GenS1{F0: 0, F1: 1}
    var v7 int = 908
    v5 = (-1 >= 1)
    fmt.Println(v3, v5)
    v8 := GenS1{F0: 89, F1: 127}
    if (154 > 39) {
        if v3 {
            var v9 int = v7
            _ = v9
        } else {
            switch "foo bar" {
            case "":
                var v10 uint = 603
                _ = v10
            case "hello":
                for fi11 := 0; fi11 < 20; fi11++ {
                    var v12 int8 = -33
                    v3 = true
                    _ = fi11
                    _ = v12
                }
            case "foo bar":
                switch "hello" {
                case "hello":
                    v13 := GenS1{F0: 945, F1: -1}
                    _ = v13
                case "":
                    v4[1] = ((((v2 >= v2) && true) && v5) && false)
                case "foo bar":
                    if (v3 && false) {
                        var v14 int8 = -1
                        if len(v1) > 0 {
                            v1[0] = -1
                        }
                        _ = v14
                    }
                default:
                    for fi15 := 0; fi15 < 8; fi15++ {
                        if true {
                            for fi16 := 0; fi16 < 15; fi16++ {
                                v7 = 32767
                                var v17 uint8 = 188
                                if (0 > 9223372036854775807) {
                                    v18 := GenS2{F0: v7, F1: 0, F2: 255}
                                    _ = v18
                                } else {
                                    if ("foo bar" <= "") {
                                        fmt.Println(v17, v7)
                                        v19 := GenS3{F0: 777, F1: v2}
                                        _ = v19
                                    } else {
                                        switch ((("a" + ("\n\t" + "")) + "a") + ("foo bar" + "foo bar")) {
                                        case "":
                                            v2 = (-(v2))
                                        case "hello":
                                            var v20 float64 = 0.0
                                            _ = v20
                                        case "foo bar":
                                            if len(v1) > 0 {
                                                v1[0] = -1
                                            }
                                        }
                                        var v21 uint32 = 2147483647
                                        v1 = append(v1, fi16)
                                        var v22 float32 = -2.25
                                        _ = v21
                                        _ = v22
                                    }
                                }
                                _ = fi16
                                _ = v17
                            }
                            var v23 float32 = (1e10 + (-2.25 + (0.1 + 0.1)))
                            _ = v23
                        } else {
                            var v24 int8 = 1
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 11; fi25++ {
                            v3 = (!(!v5))
                            if len(v1) > 0 {
                                v1[0] = 42
                            }
                            switch "foo bar" {
                            case "\n\t":
                                v1 = append(v1, (^(fi15)))
                            case "foo bar":
                                v26 := GenS4{F0: (v2 | (v2 & 281))}
                                _ = v26
                            case "hello":
                                var v27 uint64 = 0
                                _ = v27
                            default:
                                var v28 int16 = 946
                                _ = v28
                            }
                            _ = fi25
                        }
                        fmt.Println(fi15, v5, v2)
                        _ = fi15
                    }
                }
            }
        }
    }
    fmt.Println(v5)
    v4[4294967295] = (!(2 > 202))
    if v5 {
        v1 = append(v1, v7)
        if v5 {
            _ = len(v1)
            fmt.Println(v5, v2)
            v7 += (v7 >> ((1) & 63))
        } else {
            if (1 <= -1) {
                if ((0.0 != float64(-2.25)) || (3.14159 > (float64(0.0) * -2.25))) {
                    var v29 int32 = 1
                    _ = v29
                }
            }
            var v30 int32 = -1
            _ = v30
        }
        if v5 {
            v31 := GenS3{F0: 176, F1: (v2 % (((v2 / ((v2) | 1))) | 1))}
            _ = v31
        }
    }
    v32 := GenS2{F0: (v7 ^ (v7 * (v7 - (v7 % ((42) | 1))))), F1: 1, F2: 255}
    v5 = v3
    var v33 uint = 42
    var v34 int32 = 42
    v7 = (v7 * v7)
    if v5 {
        v2 += v2
        var v35 int8 = 42
        _ = v35
    }
    switch (v7 << (((v33 >> ((v33) & 63))) & 63)) {
    case 32768:
        if (3.14159 < 3.14159) {
            v36 := GenS1{F0: 42, F1: -128}
            _ = v36
        } else {
            v37 := GenS2{F0: 42, F1: (v33 % ((65535) | 1)), F2: (v33 | v33)}
            for fi38 := 0; fi38 < 7; fi38++ {
                var v39 int16 = int16(v33)
                v40 := GenS1{F0: 32767, F1: 100}
                if (true || v3) {
                    v39 = (v39 / (((v39 << ((v33) & 63))) | 1))
                    fmt.Println(fi38, v5, v3)
                    for fi41 := 0; fi41 < 9; fi41++ {
                        var v42 int32 = (v34 % ((int32(fi41)) | 1))
                        _ = fi41
                        _ = v42
                    }
                } else {
                    var v43 int64 = 1
                    for fi44 := 0; fi44 < 11; fi44++ {
                        var v45 uint8 = 0
                        if (!true) {
                            if (!(!v3)) {
                                v1 = append(v1, 1)
                            } else {
                                var v46 string = (("foo bar" + "foo bar") + (("" + ((("\n\t" + "foo bar") + "a") + "foo bar")) + "\n\t"))
                                v47 := GenS5{F0: (v33 % (((v33 / (((^(v33))) | 1))) | 1)), F1: v33}
                                var v48 int32 = (^(v34))
                                _ = v46
                                _ = v47
                                _ = v48
                            }
                            var v49 float32 = (1.5 / 0.1)
                            fmt.Println(v33)
                            _ = v49
                        }
                        var v50 string = ("foo bar" + "hello")
                        v7 -= -510
                        _ = fi44
                        _ = v45
                        _ = v50
                    }
                    _ = v43
                }
                _ = fi38
                _ = v39
                _ = v40
            }
            v51 := GenS1{F0: 2, F1: 42}
            _ = v37
            _ = v51
        }
    case 2147483647:
        if (!(1 >= 968)) {
            var v52 uint32 = 2147483648
            v53 := GenS1{F0: uint16(v33), F1: -128}
            _ = v52
            _ = v53
        } else {
            if v5 {
                var v54 int = 463
                for fi55 := 0; fi55 < 17; fi55++ {
                    v56 := v4[0]
                    _ = fi55
                    _ = v56
                }
                v54 = (v54 >> ((v33) & 63))
                _ = v54
            }
        }
    case 0:
        switch v7 {
        case 9223372036854775807:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 267:
            fmt.Println(v33)
        default:
            v7 = -393
        }
    }
    v57 := GenS5{F0: v33, F1: v33}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v7
    _ = v8
}
