package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint8; F1 string; F2 int8 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS4 struct { F0 bool; F1 uint64; F2 uint16 }

func main() {
    v1 := []uint32{4294967295, 960}
    v2 := make(map[uint32]int)
    var v3 uint16 = 32767
    var v4 uint64 = 284
    var v5 bool = ((("" >= "a") && (true || false)) || false)
    v3 = v3
    v3 += v3
    if len(v1) > 0 {
        v1[0] = 2
    }
    delete(v2, 2)
    v5 = false
    if (!(18446744073709551615 >= (^(v4)))) {
        v4 -= 18446744073709551615
    }
    var v6 float64 = float64(((((((0.0 - float32(3.14159)) * (0.1 + 3.14159)) - 1e10) / 1.0) * (0.0 + 1.5)) + -2.25))
    if (v5 || true) {
        var v7 uint16 = 42
        switch "foo bar" {
        case "foo bar":
            var v8 uint8 = 2
            _ = v8
        case "\n\t":
            for fi9 := 0; fi9 < 19; fi9++ {
                for fi10 := 0; fi10 < 18; fi10++ {
                    if (-646 != 42) {
                        v11 := GenS1{F0: 0, F1: uint16(v4), F2: v4}
                        v12 := v2[2]
                        var v13 uint8 = 136
                        if v5 {
                            var v14 uint8 = uint8(v4)
                            var v15 int8 = 0
                            if len(v1) > 0 {
                                v1[0] = 393
                            }
                            var v16 int = fi10
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                        _ = v11
                        _ = v12
                        _ = v13
                    } else {
                        v3 = v3
                        v4 = (v4 / ((v4) | 1))
                    }
                    var v17 int16 = -32768
                    v5 = v5
                    _ = fi10
                    _ = v17
                }
                _ = fi9
            }
        case "hello":
            switch "" {
            case "foo bar":
                v7 -= v3
            case "hello":
                var v18 int64 = 737
                _ = v18
            case "a":
                v2[65535] = 1
            default:
                if (1.5 > (3.14159 + v6)) {
                    if v5 {
                        delete(v2, 2147483648)
                        v19 := GenS1{F0: 42, F1: v7, F2: (v4 | v4)}
                        v20 := GenS2{F0: 0, F1: "\n\t", F2: 100}
                        _ = v19
                        _ = v20
                    } else {
                        switch -1 {
                        case 0:
                            v6 = (0.0 * v6)
                        }
                    }
                    fmt.Println(v3, v6, v3)
                }
            }
        default:
            var v21 int64 = 42
            _ = v21
        }
        _ = v7
    }
    var v22 float64 = (v6 / float64(3.14159))
    v2[2147483647] = 42
    var v23 int8 = 1
    _ = len(v1)
    for fi24 := 0; fi24 < 4; fi24++ {
        if (2147483647 > -710) {
            fmt.Println(fi24, v5)
        }
        if v5 {
            if false {
                var v25 int16 = 0
                _ = v25
            } else {
                fmt.Println(v5, v4)
                fmt.Println(v6, v23, v23)
            }
            var v26 uint32 = 42
            var v27 uint32 = v26
            _ = v26
            _ = v27
        } else {
            switch 2147483647 {
            case -9223372036854775808:
                if len(v1) > 0 {
                    v1[0] = 293
                }
            case 128:
                v28 := GenS3{F0: 255, F1: (^(v23)), F2: 4294967295}
                _ = v28
            case 127:
                for fi29 := 0; fi29 < 20; fi29++ {
                    fmt.Println(v5, v23)
                    if (v4 < v4) {
                        if (0.0 >= 1.5) {
                            var v30 uint = 691
                            if false {
                                if (v5 && (!(v4 == (v4 / ((954) | 1))))) {
                                    if (709 >= 2147483647) {
                                        var v31 int32 = 2147483647
                                        if (-9223372036854775808 <= fi24) {
                                            v31 = v31
                                        } else {
                                            if (1000 <= 259) {
                                                var v32 string = ("\n\t" + ("" + "a"))
                                                _ = v32
                                            } else {
                                                if false {
                                                    v33 := GenS4{F0: v5, F1: (v4 >> ((v30) & 63)), F2: (v3 - v3)}
                                                    if v5 {
                                                        var v34 int64 = 0
                                                        v23 -= 127
                                                        _ = v34
                                                    } else {
                                                        var v35 uint = 842
                                                        switch "hello" {
                                                        case "hello":
                                                            switch 9223372036854775807 {
                                                            case 2147483647:
                                                                delete(v2, 0)
                                                            case 447:
                                                                v2[33] = 961
                                                            }
                                                        case "":
                                                            var v36 int32 = -2147483648
                                                            _ = v36
                                                        case "a":
                                                            v1 = append(v1, 42)
                                                        default:
                                                            if true {
                                                                var v37 float32 = 0.0
                                                                v37 = (v37 + float32(v6))
                                                                switch (fi24 % (((-(fi29))) | 1)) {
                                                                case 9223372036854775807:
                                                                    var v38 uint32 = 65535
                                                                    _ = v38
                                                                case 127:
                                                                    _ = len(v2)
                                                                }
                                                                _ = v37
                                                            }
                                                        }
                                                        _ = v35
                                                    }
                                                    var v39 uint = v30
                                                    _ = v33
                                                    _ = v39
                                                }
                                            }
                                            v1 = append(v1, 0)
                                            v5 = (731 >= -32768)
                                            if ((v6 / (v22 * float64(0.1))) >= ((v6 - (-2.25 * v6)) + (v22 - -2.25))) {
                                                if v5 {
                                                    var v40 string = ("" + "")
                                                    v22 = (v6 + (1e10 * 1.5))
                                                    _ = v40
                                                }
                                            }
                                            var v41 int32 = (v31 / (((v31 ^ v31)) | 1))
                                            _ = v41
                                        }
                                        delete(v2, 65535)
                                        _ = v31
                                    }
                                }
                            }
                            if v5 {
                                v42 := GenS3{F0: 255, F1: (v23 + v23), F2: v4}
                                if false {
                                    var v43 uint64 = uint64(v30)
                                    _ = v43
                                } else {
                                    var v44 uint16 = 93
                                    v45 := GenS1{F0: 93, F1: v44, F2: (^(v4))}
                                    fmt.Println(v4, v44, v6)
                                    _ = v44
                                    _ = v45
                                }
                                _ = v42
                            } else {
                                v5 = (!(false || v5))
                                fmt.Println(v6, v23)
                                v5 = true
                            }
                            _ = v30
                        } else {
                            fmt.Println(v4)
                            if v5 {
                                var v46 uint32 = 65535
                                _ = v46
                            }
                        }
                    }
                    v47 := GenS4{F0: true, F1: v4, F2: 32767}
                    _ = fi29
                    _ = v47
                }
            }
            var v48 uint = 763
            _ = v48
        }
        _ = fi24
    }
    v23 += 100
    var v49 string = ""
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
