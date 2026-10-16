package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 string; F2 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 uint32; F2 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int; F1 uint }
type GenS8 struct { F0 int32; F1 uint; F2 int16 }

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]string)
    v3 := make(map[string]uint32)
    v4 := []uint32{}
    v5 := make(map[string]bool)
    switch 42 {
    case 32768:
        v6 := GenS1{F0: ((("hello" + "foo bar") + "") + (("hello" + (("" + "hello") + "")) + ((("\n\t" + ("a" + "\n\t")) + (("hello" + "") + "foo bar")) + ("" + "\n\t"))))}
        _ = v6
    case 42:
        var v7 float32 = 1.5
        _ = v7
    case 32767:
        if false {
            for fi8 := 0; fi8 < 8; fi8++ {
                fmt.Println(v1, fi8)
                _ = fi8
            }
            var v9 int8 = 0
            v1 -= 544
            _ = v9
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    var v10 uint = 0
    for fi11 := 0; fi11 < 18; fi11++ {
        fmt.Println(v10, fi11)
        if false {
            for fi12 := 0; fi12 < 16; fi12++ {
                fmt.Println(v10, v10, v1)
                _ = fi12
            }
        }
        for fi13 := 0; fi13 < 16; fi13++ {
            if (213 < 63) {
                v1 = v1
                if false {
                    v2[-1] = ("foo bar" + "\n\t")
                    var v14 uint8 = 255
                    _ = v14
                } else {
                    v10 += (v10 | 42)
                    var v15 int16 = (v1 >> (((v10 | (v10 % (((^(v10))) | 1)))) & 63))
                    _ = v15
                }
                switch (((("" + "") + "") + "hello") + "a") {
                case "a":
                    v1 = 0
                case "":
                    var v16 float64 = (1e10 + (3.14159 / 1.0))
                    _ = v16
                case "\n\t":
                    _ = len(v4)
                }
            }
            var v17 uint16 = 739
            _ = fi13
            _ = v17
        }
        _ = fi11
    }
    var v18 string = ""
    if true {
        if len(v4) > 0 {
            v4[0] = 320
        }
        var v19 string = "hello"
        _ = v19
    }
    _ = len(v3)
    var v20 uint16 = 32767
    if true {
        if false {
            v1 -= -160
            v18 = v18
        }
    } else {
        switch 128 {
        case 1:
            switch 32768 {
            case 127:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        case 32768:
            v21 := GenS2{F0: v1}
            _ = v21
        case 42:
            v20 += (v20 & v20)
        }
        for fi22 := 0; fi22 < 11; fi22++ {
            var v23 int = 478
            v20 += (v20 % ((v20) | 1))
            v1 += v1
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 15; fi24++ {
            for fi25 := 0; fi25 < 17; fi25++ {
                v18 = v18
                v26 := GenS2{F0: 1000}
                var v27 float32 = (0.1 - 0.1)
                _ = fi25
                _ = v26
                _ = v27
            }
            switch ("" + ("" + (("hello" + "") + "hello"))) {
            case "hello":
                v1 += 1000
            case "\n\t":
                v1 = (v1 | v1)
            case "foo bar":
                v18 = ""
            }
            switch 128 {
            case 128:
                switch (fi24 + int(v1)) {
                case 32768:
                    if (2 != 877) {
                        v28 := GenS2{F0: v1}
                        for fi29 := 0; fi29 < 2; fi29++ {
                            v1 -= v1
                            if true {
                                v10 = 2
                                switch fi29 {
                                case 128:
                                    v18 = "hello"
                                case 4294967295:
                                    fmt.Println(v1, fi24, v1)
                                case 32768:
                                    v30 := GenS3{F0: 65535, F1: (v18 + (v18 + (("hello" + "hello") + (v18 + "")))), F2: -128}
                                    _ = v30
                                default:
                                    v4 = append(v4, 665)
                                }
                                var v31 uint64 = 141
                                _ = v31
                            }
                            v32 := GenS1{F0: v18}
                            _ = fi29
                            _ = v32
                        }
                        _ = v28
                    } else {
                        v33 := GenS4{F0: 32767}
                        _ = v33
                    }
                case -2147483648:
                    if (((3.14159 * 1.5) + (3.14159 - 3.14159)) <= -2.25) {
                        fmt.Println(v10)
                        if false {
                            var v34 float32 = 0.0
                            if false {
                                v18 = v18
                            }
                            v10 += v10
                            _ = v34
                        }
                    }
                }
            case 9223372036854775807:
                v10 = v10
            case -1:
                v10 = (v10 >> ((2) & 63))
            default:
                for fi35 := 0; fi35 < 7; fi35++ {
                    if (((2147483647 >= 0) && false) || false) {
                        if false {
                            switch (v18 + "\n\t") {
                            case "\n\t":
                                v1 = (v1 - 446)
                            }
                            var v36 int = 32767
                            v18 = "hello"
                            var v37 uint8 = 2
                            if false {
                                v38 := GenS5{F0: (v1 << ((v10) & 63)), F1: 885, F2: uint(v20)}
                                switch (v18 + (("foo bar" + "a") + "")) {
                                case "foo bar":
                                    v39 := GenS6{F0: 199}
                                    _ = v39
                                case "hello":
                                    v40 := GenS5{F0: 0, F1: 65535, F2: (^(v10))}
                                    _ = v40
                                case "a":
                                    v20 -= (v20 % ((v20) | 1))
                                    var v41 uint8 = (v37 ^ (v37 << ((v10) & 63)))
                                    switch (v36 << (((v10 >> (((v10 * (v10 / (((v10 - (v10 + v10))) | 1)))) & 63))) & 63)) {
                                    case 127:
                                        if len(v4) > 0 {
                                            v4[0] = 2
                                        }
                                    }
                                    switch v18 {
                                    case "hello":
                                        v42 := GenS7{F0: (fi35 + (fi35 | (^(fi35)))), F1: v10}
                                        _ = v42
                                    case "":
                                        var v43 int = -424
                                        _ = v43
                                    case "foo bar":
                                        if true {
                                            var v44 uint32 = 0
                                            if false {
                                                fmt.Println(v37, v20)
                                            } else {
                                                var v45 int8 = int8(v1)
                                                fmt.Println(fi35, v1, v45)
                                                var v46 uint64 = 18446744073709551615
                                                v45 = (-(v45))
                                                if (1 != (v45 / ((v45) | 1))) {
                                                    switch "\n\t" {
                                                    case "":
                                                        v1 -= (-(v1))
                                                    }
                                                    var v47 int32 = 497
                                                    fmt.Println(v41, v45)
                                                    _ = v47
                                                }
                                                _ = v45
                                                _ = v46
                                            }
                                            _ = v44
                                        } else {
                                            v37 -= v41
                                            if true {
                                                var v48 uint8 = 2
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = v48
                                            }
                                            if (false && false) {
                                                var v49 float64 = (1e10 / float64(3.14159))
                                                _ = v49
                                            } else {
                                                var v50 int32 = 42
                                                v18 = v18
                                                if true {
                                                    v10 += v10
                                                }
                                                _ = v50
                                            }
                                            delete(v2, v36)
                                        }
                                    }
                                    fmt.Println(fi24, v20, v10)
                                    _ = v41
                                default:
                                    switch (v18 + "hello") {
                                    case "foo bar":
                                        var v51 float64 = -2.25
                                        _ = v51
                                    case "hello":
                                        _ = len(v5)
                                    case "":
                                        if (281 <= v10) {
                                            var v52 uint64 = 4294967295
                                            if len(v4) > 0 {
                                                v4[0] = 65535
                                            }
                                            if true {
                                                v10 += 379
                                                v37 = 42
                                            }
                                            _ = v52
                                        }
                                    default:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                }
                                _ = v38
                            } else {
                                v3[""] = 18
                                v36 = 9223372036854775807
                            }
                            _ = v36
                            _ = v37
                        } else {
                            v53 := GenS3{F0: 137, F1: ("a" + (v18 + v18)), F2: 127}
                            _ = v53
                        }
                        var v54 int16 = 42
                        _ = v54
                    } else {
                        var v55 uint = (v10 % ((v10) | 1))
                        v20 += v20
                        switch (((v18 + "hello") + (v18 + "\n\t")) + "foo bar") {
                        case "a":
                            delete(v5, "")
                            var v56 uint64 = 42
                            var v57 int8 = -81
                            _ = v56
                            _ = v57
                        case "hello":
                            if (0.1 <= 3.14159) {
                                var v58 float32 = 1e10
                                v18 = "a"
                                _ = v58
                            }
                        case "foo bar":
                            v18 = ("hello" + "\n\t")
                        }
                        _ = v55
                    }
                    var v59 int16 = v1
                    _ = len(v2)
                    _ = fi35
                    _ = v59
                }
            }
            _ = fi24
        }
    }
    for fi60 := 0; fi60 < 18; fi60++ {
        v4 = append(v4, 42)
        v5["a"] = (0 < 9223372036854775807)
        _ = fi60
    }
    var v61 int = 127
    fmt.Println(v10, v1)
    v62 := v3[v18]
    v10 += 2
    v63 := GenS8{F0: 152, F1: v10, F2: v1}
    v18 = (v18 + v18)
    v10 = (v10 * (^(v10)))
    var v64 int64 = 9223372036854775807
    var v65 int32 = -390
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
