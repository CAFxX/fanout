package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 uint16; F1 int32; F2 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    v3 := []string{("\n\t" + "\n\t")}
    v4 := make(map[string]string)
    if (((1e10 / 1.0) <= 1e10) || (!true)) {
        var v5 int8 = -1
        _ = v5
    } else {
        var v6 int = -286
        _ = v6
    }
    switch (("a" + "\n\t") + "hello") {
    case "":
        for fi7 := 0; fi7 < 16; fi7++ {
            if ((651 != 65535) && (!true)) {
                fmt.Println(fi7)
                fmt.Println(fi7)
            } else {
                if (42 > 2) {
                    if (2 == 1) {
                        var v8 uint16 = 1
                        v8 -= (v8 | 42)
                        _ = v8
                    } else {
                        var v9 uint8 = 2
                        _ = v9
                    }
                }
            }
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            v10 := GenS1{F0: 1, F1: (fi7 + fi7), F2: fi7}
            _ = fi7
            _ = v10
        }
    case "\n\t":
        for fi11 := 0; fi11 < 6; fi11++ {
            if (!false) {
                if false {
                    var v12 int = (^(fi11))
                    _ = v12
                }
                v13 := GenS2{F0: fi11, F1: fi11}
                for fi14 := 0; fi14 < 7; fi14++ {
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                    var v15 uint32 = 336
                    if (!false) {
                        v16 := v2[v15]
                        v16 = (v15 / ((v16) | 1))
                        if false {
                            var v17 int = -802
                            var v18 int64 = 9223372036854775807
                            v17 += v17
                            _ = v17
                            _ = v18
                        }
                        _ = v16
                    }
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
            fmt.Println(fi11)
            _ = fi11
        }
    case "a":
        var v19 uint64 = 88
        _ = v19
    }
    var v20 int8 = 100
    for fi21 := 0; fi21 < 6; fi21++ {
        for fi22 := 0; fi22 < 7; fi22++ {
            _ = len(v3)
            if (!(0 < 42)) {
                if (!true) {
                    switch (fi21 + (fi22 - 2147483647)) {
                    case 127:
                        v20 += -124
                    }
                    v20 += v20
                }
                v20 = v20
                v20 += 94
                switch "" {
                case "":
                    v20 += (v20 | (v20 >> ((255) & 63)))
                case "a":
                    switch ("foo bar" + "") {
                    case "hello":
                        v20 += 0
                    case "a":
                        if (true || true) {
                            v20 = (-(v20))
                            switch ("" + "a") {
                            case "hello":
                                var v23 uint32 = 2147483647
                                _ = v23
                            case "foo bar":
                                switch ("\n\t" + "hello") {
                                case "":
                                    if len(v3) > 0 {
                                        v3[0] = ("" + "\n\t")
                                    }
                                case "a":
                                    fmt.Println(fi21, v20, fi22)
                                }
                            case "":
                                fmt.Println(fi21, fi21, v20)
                            default:
                                v20 = v20
                            }
                        } else {
                            switch "foo bar" {
                            case "a":
                                var v24 int8 = v20
                                _ = v24
                            case "hello":
                                if true {
                                    var v25 uint = 130
                                    switch "\n\t" {
                                    case "foo bar":
                                        v26 := GenS2{F0: -1, F1: (fi21 + fi21)}
                                        _ = v26
                                    case "\n\t":
                                        var v27 float64 = -2.25
                                        _ = v27
                                    }
                                    if true {
                                        fmt.Println(fi22, fi21)
                                    }
                                    var v28 bool = true
                                    _ = v25
                                    _ = v28
                                }
                            case "\n\t":
                                var v29 uint8 = 42
                                _ = v29
                            default:
                                var v30 float32 = 1.5
                                _ = v30
                            }
                        }
                    case "\n\t":
                        v20 = -113
                    }
                case "foo bar":
                    if len(v3) > 0 {
                        v3[0] = ("hello" + "hello")
                    }
                    switch fi22 {
                    case 4294967295:
                        v20 = v20
                    case 32767:
                        _ = len(v3)
                        var v31 uint = 236
                        v31 = (v31 << ((v31) & 63))
                        _ = v31
                    case 32768:
                        fmt.Println(fi22, fi22)
                    }
                    var v32 uint16 = 425
                    fmt.Println(fi22, v32)
                    delete(v2, 65535)
                    fmt.Println(v20, v20, fi21)
                    _ = v32
                }
                delete(v2, 569)
                v33 := GenS3{F0: -104, F1: 2}
                var v34 float32 = ((-2.25 - ((float32(1.5) * float32(3.14159)) + 0.0)) + float32(float64(1e10)))
                _ = v33
                _ = v34
            } else {
                var v35 int8 = 17
                if true {
                    fmt.Println(v35, fi21, fi22)
                } else {
                    fmt.Println(v35, v35, v35)
                    var v36 int8 = -1
                    _ = len(v3)
                    _ = v36
                }
                _ = len(v3)
                if ((!true) && (-32768 != 42)) {
                    switch fi22 {
                    case -2147483648:
                        _ = len(v2)
                    case 42:
                        if (fi21 == (^(fi21))) {
                            switch (fi22 << ((18446744073709551615) & 63)) {
                            case 2147483647:
                                v37 := GenS3{F0: (-(v35)), F1: 32767}
                                _ = v37
                            case 128:
                                fmt.Println(v20, v20, v20)
                            default:
                                v38 := GenS4{F0: 222, F1: 2147483647, F2: -653}
                                _ = v38
                            }
                            v35 = v35
                        }
                    case 32767:
                        var v39 int32 = 1
                        switch fi21 {
                        case -2147483648:
                            var v40 float32 = 1.5
                            _ = v40
                        default:
                            var v41 float64 = 1e10
                            _ = v41
                        }
                        _ = v39
                    default:
                        var v42 int16 = -1
                        v20 = (^(v20))
                        _ = v42
                    }
                    var v43 uint64 = 42
                    _ = v43
                } else {
                    if false {
                        v44 := GenS5{F0: ("\n\t" + "a")}
                        v35 -= 127
                        _ = v44
                    } else {
                        fmt.Println(fi22)
                    }
                    var v45 int8 = 100
                    _ = v45
                }
                var v46 int8 = 1
                _ = v35
                _ = v46
            }
            fmt.Println(fi22, fi21, fi21)
            _ = fi22
        }
        var v47 float32 = float32(1e10)
        _ = fi21
        _ = v47
    }
    if false {
        v20 -= (^(v20))
    } else {
        var v48 bool = (1 >= 4294967295)
        v49 := v1[128]
        _ = v48
        _ = v49
    }
    if (v20 >= 42) {
        fmt.Println(v20)
    } else {
        v20 += v20
    }
    _ = len(v1)
    v20 += (^(v20))
    v20 = (v20 & v20)
    v50 := GenS4{F0: 32767, F1: 0, F2: 1000}
    fmt.Println(v20)
    v3 = append(v3, "")
    v20 = v20
    var v51 uint64 = 0
    v52 := GenS1{F0: 42, F1: 2147483647, F2: 910}
    v20 += (v20 * v20)
    v53 := GenS6{F0: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v52
    _ = v53
}
