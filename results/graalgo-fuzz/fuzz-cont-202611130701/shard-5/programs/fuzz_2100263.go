package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint64 }
type GenS2 struct { F0 string; F1 int; F2 int32 }
type GenS3 struct { F0 uint64; F1 bool; F2 string }
type GenS4 struct { F0 uint8; F1 string }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{0}
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    v5 := []uint32{2147483647, 65535, 797, 305}
    var v6 uint32 = 2147483647
    v6 += 1
    fmt.Println(v6)
    v7 := GenS1{F0: "", F1: 0, F2: 18446744073709551615}
    v6 += 4294967295
    var v8 uint = 18446744073709551615
    if len(v2) > 0 {
        v2[0] = (v6 | (v6 + (v6 + v6)))
    }
    switch "hello" {
    case "\n\t":
        switch ("a" + (("foo bar" + ("hello" + "")) + ("a" + (("" + "\n\t") + "\n\t")))) {
        case "a":
            var v9 uint = (v8 + 491)
            _ = v9
        case "\n\t":
            var v10 int16 = -32768
            _ = v10
        default:
            var v11 float32 = 0.0
            _ = v11
        }
    case "hello":
        if len(v5) > 0 {
            v5[0] = v6
        }
    case "foo bar":
        v6 = 2147483647
    }
    switch "foo bar" {
    case "":
        for fi12 := 0; fi12 < 16; fi12++ {
            for fi13 := 0; fi13 < 4; fi13++ {
                switch "hello" {
                case "\n\t":
                    v8 = v8
                case "hello":
                    if false {
                        if (true && true) {
                            var v14 uint8 = 128
                            if (85 != 957) {
                                v6 = v6
                                switch fi13 {
                                case 4294967295:
                                    _ = len(v4)
                                default:
                                    v14 -= 1
                                }
                            } else {
                                switch (fi12 % (((-(fi12))) | 1)) {
                                case 32768:
                                    var v15 uint16 = 417
                                    _ = v15
                                default:
                                    switch "hello" {
                                    case "a":
                                        switch "foo bar" {
                                        case "foo bar":
                                            v14 = (v14 + v14)
                                        }
                                    case "":
                                        v16 := GenS1{F0: "\n\t", F1: 1, F2: 42}
                                        _ = v16
                                    case "foo bar":
                                        var v17 uint = v8
                                        _ = v17
                                    }
                                }
                                var v18 float32 = float32(1e10)
                                if ((float32(1e10) < (3.14159 * v18)) || true) {
                                    var v19 string = ""
                                    _ = len(v4)
                                    _ = v19
                                } else {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                v6 -= v6
                                var v20 bool = false
                                _ = v18
                                _ = v20
                            }
                            v14 += v14
                            v2 = append(v2, (v6 & (v6 * v6)))
                            v8 = uint(v6)
                            v3[(("a" + "\n\t") + "")] = v6
                            _ = v14
                        }
                    }
                case "":
                    v8 = (v8 << (((v8 / (((^(v8))) | 1))) & 63))
                }
                var v21 uint32 = v6
                switch fi12 {
                case 128:
                    var v22 uint8 = 127
                    _ = v22
                case 2147483647:
                    switch ("hello" + "\n\t") {
                    case "a":
                        var v23 bool = false
                        _ = v23
                    }
                }
                _ = fi13
                _ = v21
            }
            v2 = append(v2, (v6 << ((0) & 63)))
            _ = fi12
        }
    case "foo bar":
        switch "foo bar" {
        case "\n\t":
            if true {
                v8 = v8
            } else {
                if (true && (-1 > -937)) {
                    var v24 int64 = 1
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    switch 0 {
                    case 4294967295:
                        v25 := GenS2{F0: "", F1: 1, F2: 42}
                        _ = v25
                    case 9223372036854775807:
                        v24 = (v24 ^ (v24 - v24))
                    }
                    _ = v24
                }
                v8 = (v8 - (v8 + (v8 + (v8 + v8))))
                v26 := GenS3{F0: 4294967295, F1: false, F2: (((("" + "\n\t") + ("foo bar" + (("a" + ("" + "")) + ("hello" + "foo bar")))) + "a") + (("foo bar" + (("foo bar" + "hello") + "")) + (((("hello" + "foo bar") + ("\n\t" + ("a" + (("a" + "\n\t") + "a")))) + (("foo bar" + ("a" + "")) + "a")) + "foo bar")))}
                _ = v26
            }
        case "":
            var v27 uint16 = 443
            _ = v27
        case "foo bar":
            if (65535 <= 0) {
                v28 := GenS1{F0: "a", F1: 42, F2: 1}
                _ = v28
            } else {
                v8 += (v8 & (v8 - v8))
                for fi29 := 0; fi29 < 6; fi29++ {
                    if false {
                        for fi30 := 0; fi30 < 1; fi30++ {
                            var v31 int8 = -18
                            var v32 int8 = -1
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        fmt.Println(fi29)
                    }
                    if true {
                        v8 = 42
                        v8 += v8
                        var v33 uint16 = 255
                        _ = v33
                    }
                    _ = fi29
                }
                if false {
                    v34 := GenS3{F0: 791, F1: (false || true), F2: (("" + "hello") + ("a" + "hello"))}
                    _ = v34
                }
            }
        default:
            delete(v3, "hello")
        }
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        switch ((("hello" + "hello") + "foo bar") + "foo bar") {
        case "":
            v8 -= 0
        case "a":
            var v36 int32 = 263
            _ = v36
        case "\n\t":
            v37 := v1[fi35]
            _ = v37
        default:
            v6 = (v6 % (((v6 % ((v6) | 1))) | 1))
        }
        v8 = (^(v8))
        switch "\n\t" {
        case "":
            v8 -= (v8 / ((v8) | 1))
        case "hello":
            var v38 uint32 = v6
            _ = v38
        default:
            for fi39 := 0; fi39 < 7; fi39++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v8 = 0
                _ = fi39
            }
        }
        _ = fi35
    }
    var v40 int8 = -1
    if len(v5) > 0 {
        v5[0] = 0
    }
    var v41 bool = false
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi42 := 0; fi42 < 14; fi42++ {
        for fi43 := 0; fi43 < 5; fi43++ {
            v40 = v40
            var v44 int = fi43
            var v45 uint64 = 18446744073709551615
            switch fi43 {
            case 164:
                if ((182 == 38) || ((-2.25 + 0.1) >= 3.14159)) {
                    var v46 uint64 = v45
                    if (9223372036854775807 != -815) {
                        switch "a" {
                        case "foo bar":
                            if false {
                                fmt.Println(v6)
                                var v47 string = ""
                                _ = v47
                            }
                        case "":
                            v1[(v44 % ((1) | 1))] = 9223372036854775807
                        case "\n\t":
                            v6 = v6
                        }
                        var v48 int32 = 928
                        switch 42 {
                        case 1:
                            var v49 uint = v8
                            v41 = v41
                            _ = len(v5)
                            _ = v49
                        default:
                            _ = len(v5)
                        }
                        _ = v48
                    }
                    switch fi43 {
                    case 0:
                        if (((v44 << ((v8) & 63)) < (fi42 - (v44 >> (((v8 / ((18446744073709551615) | 1))) & 63)))) && (0 != 1)) {
                            var v50 int8 = v40
                            v51 := v3[""]
                            _ = v50
                            _ = v51
                        }
                    case 9223372036854775807:
                        var v52 uint8 = 42
                        _ = v52
                    }
                    v41 = (!(!v41))
                    _ = v46
                }
            case -998:
                fmt.Println(v45, fi42)
            default:
                v53 := GenS4{F0: 2, F1: "\n\t"}
                _ = v53
            }
            switch (fi42 | fi42) {
            case 32768:
                v54 := GenS4{F0: 29, F1: "\n\t"}
                _ = v54
            default:
                fmt.Println(v6, v40, v44)
            }
            _ = fi43
            _ = v44
            _ = v45
        }
        _ = fi42
    }
    var v55 int64 = -726
    if len(v2) > 0 {
        v2[0] = (v6 / (((v6 / (((v6 - 299)) | 1))) | 1))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v55
    _ = v6
    _ = v7
    _ = v8
}
