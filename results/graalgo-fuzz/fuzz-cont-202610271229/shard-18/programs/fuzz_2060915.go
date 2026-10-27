package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 string; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 bool; F2 int32 }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 bool; F1 int; F2 int32 }

func main() {
    var v1 float32 = (3.14159 * float32(float64(1.5)))
    var v2 float32 = 1.5
    var v3 uint = 0
    v4 := make(map[int]bool)
    v5 := []uint32{2147483647, 1, 358, 2147483648}
    v6 := GenS1{F0: 1, F1: 0}
    var v7 string = "hello"
    v1 = v2
    fmt.Println(v1)
    var v8 int = 4294967295
    var v9 string = v7
    for fi10 := 0; fi10 < 5; fi10++ {
        v2 -= (float32(3.14159) / 1.0)
        var v11 uint64 = 18446744073709551615
        _ = len(v4)
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        if (1 != -128) {
            v1 = float32(3.14159)
            v4[v8] = false
        }
        v1 = 1.5
        _ = fi12
    }
    v13 := GenS2{F0: "foo bar", F1: 42, F2: 4294967295}
    var v14 bool = false
    if ((3.14159 > (float64((v1 - (v1 / v2))) * 1e10)) && v14) {
        var v15 float64 = float64(v1)
        switch ((v7 + v9) + ("\n\t" + "")) {
        case "a":
            if v14 {
                switch (-(v8)) {
                case -809:
                    for fi16 := 0; fi16 < 10; fi16++ {
                        var v17 int64 = 655
                        _ = fi16
                        _ = v17
                    }
                case 128:
                    v18 := GenS2{F0: "", F1: 1, F2: 65535}
                    _ = v18
                }
            } else {
                var v19 float32 = v2
                for fi20 := 0; fi20 < 14; fi20++ {
                    if len(v5) > 0 {
                        v5[0] = 523
                    }
                    v21 := GenS3{F0: 9223372036854775807}
                    _ = fi20
                    _ = v21
                }
                _ = v19
            }
        case "hello":
            if ((v15 == v15) && (v14 && (-828 > (-(v8))))) {
                switch "\n\t" {
                case "foo bar":
                    if (0 != 107) {
                        for fi22 := 0; fi22 < 13; fi22++ {
                            var v23 float32 = v2
                            switch v9 {
                            case "a":
                                v24 := GenS3{F0: 0}
                                _ = v24
                            default:
                                v15 = (v15 - ((1.5 - v15) - 0.0))
                            }
                            _ = fi22
                            _ = v23
                        }
                        _ = len(v5)
                        for fi25 := 0; fi25 < 11; fi25++ {
                            if false {
                                var v26 int8 = -1
                                var v27 int = (^(fi25))
                                v27 += v8
                                _ = v26
                                _ = v27
                            } else {
                                fmt.Println(v3, fi25)
                                switch fi25 {
                                case 9223372036854775807:
                                    if v14 {
                                        var v28 uint = (^(v3))
                                        v8 = fi25
                                        var v29 uint16 = 40
                                        _ = v28
                                        _ = v29
                                    } else {
                                        for fi30 := 0; fi30 < 9; fi30++ {
                                            if true {
                                                v31 := GenS2{F0: (v7 + v7), F1: 255, F2: 0}
                                                _ = v31
                                            }
                                            _ = fi30
                                        }
                                        v32 := GenS4{F0: -513}
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = v32
                                    }
                                case 2147483647:
                                    v15 += v15
                                case -9223372036854775808:
                                    v4[fi25] = (-1 == -128)
                                default:
                                    v14 = true
                                }
                            }
                            if ((3 > 32767) || ((!v14) || v14)) {
                                v8 += v8
                            } else {
                                if (-835 != 0) {
                                    switch (v8 - 4294967295) {
                                    case 4294967295:
                                        var v33 int32 = 2147483647
                                        _ = v33
                                    case -9223372036854775808:
                                        if true {
                                            fmt.Println(v9, v15)
                                            var v34 uint16 = 0
                                            var v35 uint64 = 4294967295
                                            _ = v34
                                            _ = v35
                                        } else {
                                            var v36 string = ((v9 + "a") + ("foo bar" + "hello"))
                                            switch (v36 + (v36 + "\n\t")) {
                                            case "a":
                                                if true {
                                                    var v37 int = v8
                                                    if (true || v14) {
                                                        if true {
                                                            switch (v36 + ("\n\t" + "")) {
                                                            case "\n\t":
                                                                var v38 uint32 = 4294967295
                                                                _ = v38
                                                            case "a":
                                                                if v14 {
                                                                    for fi39 := 0; fi39 < 19; fi39++ {
                                                                        var v40 float64 = (float64(v1) - (v15 + (v15 * ((v15 - 1e10) * (((3.14159 / (-2.25 * v15)) + v15) * ((v15 * 1.5) + (1e10 * v15)))))))
                                                                        v41 := GenS4{F0: -9223372036854775808}
                                                                        v7 = v36
                                                                        _ = fi39
                                                                        _ = v40
                                                                        _ = v41
                                                                    }
                                                                    v7 = v9
                                                                }
                                                            default:
                                                                switch v36 {
                                                                case "a":
                                                                    switch (v8 - (fi25 + (-(v8)))) {
                                                                    case 128:
                                                                        var v42 int = 692
                                                                        _ = v42
                                                                    case 127:
                                                                        var v43 string = v7
                                                                        _ = v43
                                                                    }
                                                                case "foo bar":
                                                                    v44 := GenS5{F0: -32768}
                                                                    _ = v44
                                                                }
                                                            }
                                                            if v14 {
                                                                v37 += (^(fi25))
                                                            }
                                                        } else {
                                                            v8 = (v8 - v37)
                                                            fmt.Println(fi25)
                                                            fmt.Println(v1, v37)
                                                        }
                                                        var v45 int16 = -91
                                                        fmt.Println(v15, v3)
                                                        _ = v45
                                                    } else {
                                                        var v46 string = v7
                                                        _ = v46
                                                    }
                                                    _ = v37
                                                }
                                            case "\n\t":
                                                if (("\n\t" + v9) != v7) {
                                                    v47 := GenS1{F0: 1, F1: (fi25 - (^(v8)))}
                                                    v48 := GenS1{F0: -1, F1: (v8 - -1)}
                                                    _ = v47
                                                    _ = v48
                                                }
                                            }
                                            v49 := GenS6{F0: 65535, F1: true, F2: -1}
                                            _ = v36
                                            _ = v49
                                        }
                                    case 127:
                                        var v50 float32 = v1
                                        _ = v50
                                    }
                                    for fi51 := 0; fi51 < 11; fi51++ {
                                        v52 := GenS7{F0: 2, F1: v9}
                                        if (0 != 1) {
                                            v53 := GenS7{F0: 18446744073709551615, F1: "foo bar"}
                                            v7 = (v7 + (v7 + ""))
                                            _ = v53
                                        } else {
                                            fmt.Println(fi25, fi25)
                                            if (80 != 2) {
                                                var v54 float64 = ((v15 * v15) + v15)
                                                v55 := GenS3{F0: -624}
                                                if (1 == -1) {
                                                    v15 = ((0.1 + float64(v2)) - (v54 * 3.14159))
                                                    v15 = v54
                                                    v1 -= v1
                                                }
                                                _ = v54
                                                _ = v55
                                            }
                                            if (((v2 * 0.0) + ((float32((0.1 * v15)) * -2.25) + v2)) > float32(-2.25)) {
                                                v7 = v9
                                                v56 := GenS8{F0: (39 <= 1), F1: (v8 % (((fi51 >> ((65535) & 63))) | 1)), F2: -714}
                                                if true {
                                                    if true {
                                                        if v14 {
                                                            _ = len(v5)
                                                            v3 += v3
                                                        }
                                                        fmt.Println(fi25, v14, fi25)
                                                        var v57 float64 = float64(1.5)
                                                        _ = v57
                                                    }
                                                } else {
                                                    var v58 float64 = (((((v15 + float64((v1 - ((0.1 / v2) + (1e10 * v2))))) * v15) * float64(float32(v15))) - ((v15 + ((v15 + ((1.5 - v15) / v15)) / 0.1)) + 1e10)) * float64(v2))
                                                    v7 = "foo bar"
                                                    _ = v58
                                                }
                                                _ = v56
                                            }
                                        }
                                        _ = fi51
                                        _ = v52
                                    }
                                    fmt.Println(v9)
                                } else {
                                    if v14 {
                                        v59 := GenS8{F0: v14, F1: (-(v8)), F2: 2147483647}
                                        _ = v59
                                    } else {
                                        v15 += v15
                                    }
                                    for fi60 := 0; fi60 < 7; fi60++ {
                                        if (124 <= -1) {
                                            var v61 uint = v3
                                            _ = v61
                                        } else {
                                            var v62 float64 = float64(1e10)
                                            _ = v62
                                        }
                                        var v63 int32 = 1
                                        if (!(false && ((v3 * 18446744073709551615) > 4294967295))) {
                                            fmt.Println(v7, v3)
                                            fmt.Println(v1, v7)
                                        }
                                        v9 = v9
                                        _ = fi60
                                        _ = v63
                                    }
                                    v64 := GenS1{F0: -24, F1: -9223372036854775808}
                                    _ = v64
                                }
                            }
                            _ = fi25
                        }
                    } else {
                        var v65 uint8 = 128
                        _ = v65
                    }
                default:
                    switch v9 {
                    case "":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                }
            }
        case "foo bar":
            v15 -= v15
        }
        _ = v15
    }
    v3 += 241
    if (v14 || v14) {
        v3 = 1
        fmt.Println(v3, v3, v3)
    }
    v2 = v1
    v2 += v2
    var v66 uint = v3
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v66:", v66)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v7
    _ = v8
    _ = v9
}
