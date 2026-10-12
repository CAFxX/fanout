package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int32; F1 int32; F2 uint64 }

func main() {
    var v1 bool = (!false)
    var v2 uint16 = 113
    v3 := []int{42, 42, -1, 128}
    v4 := make(map[int]bool)
    var v5 uint16 = v2
    var v6 uint8 = 182
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint8 = (v6 << ((65535) & 63))
        var v9 int64 = 1
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 4294967295, F1: 18446744073709551615, F2: "hello"}
    var v11 uint32 = 2147483647
    v12 := GenS1{F0: (v11 << ((790) & 63)), F1: 42, F2: ""}
    fmt.Println(v2, v11, v2)
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := v4[(fi13 + (fi13 + (fi13 - fi13)))]
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        v6 -= v6
        v16 := GenS1{F0: uint32(v5), F1: 255, F2: ""}
        fmt.Println(v6)
        _ = fi15
        _ = v16
    }
    if true {
        if v1 {
            switch (("a" + "a") + ("" + "a")) {
            case "hello":
                var v17 uint16 = v2
                _ = v17
            }
        }
    }
    v18 := GenS2{F0: v2}
    v4[-9223372036854775808] = v1
    var v19 int8 = -56
    fmt.Println(v5)
    v11 -= v11
    switch 0 {
    case 32768:
        var v20 int64 = -715
        _ = v20
    case -1:
        fmt.Println(v11, v6, v6)
    case 128:
        var v21 uint8 = v6
        _ = v21
    default:
        if v1 {
            for fi22 := 0; fi22 < 2; fi22++ {
                for fi23 := 0; fi23 < 3; fi23++ {
                    switch (("a" + "foo bar") + "a") {
                    case "foo bar":
                        var v24 float32 = (3.14159 * 0.0)
                        var v25 int8 = (-(v19))
                        var v26 float64 = 3.14159
                        _ = v24
                        _ = v25
                        _ = v26
                    case "":
                        switch (fi23 & (-(fi22))) {
                        case 436:
                            if (true && v1) {
                                var v27 uint32 = (v11 + 2147483648)
                                if (v5 == (v2 - v2)) {
                                    v28 := GenS3{F0: 0, F1: 410}
                                    _ = v28
                                }
                                var v29 float64 = -2.25
                                _ = v27
                                _ = v29
                            } else {
                                if (0 >= 0) {
                                    if (false && v1) {
                                        v5 = v5
                                    } else {
                                        v1 = (907 > v11)
                                        delete(v4, -487)
                                    }
                                    v30 := GenS2{F0: v2}
                                    fmt.Println(fi22, v2, v19)
                                    _ = v30
                                } else {
                                    switch "foo bar" {
                                    case "\n\t":
                                        _ = len(v4)
                                    case "foo bar":
                                        v31 := GenS4{F0: fi22}
                                        _ = v31
                                    }
                                    v2 -= (v2 % (((v2 * 65535)) | 1))
                                    if v1 {
                                        var v32 uint64 = 245
                                        switch (((((("foo bar" + ("\n\t" + "a")) + "foo bar") + "foo bar") + "a") + "") + ("a" + (("\n\t" + ("a" + "hello")) + (((("\n\t" + "\n\t") + "\n\t") + "hello") + "\n\t")))) {
                                        case "":
                                            v6 = v6
                                        case "a":
                                            if v1 {
                                                if true {
                                                    var v33 int16 = -32768
                                                    _ = v33
                                                } else {
                                                    v3 = append(v3, 718)
                                                }
                                                switch 423 {
                                                case -1000:
                                                    v5 -= 65535
                                                case -1:
                                                    v5 += (v5 & uint16(v11))
                                                default:
                                                    v34 := GenS4{F0: fi23}
                                                    _ = v34
                                                }
                                                v35 := GenS5{F0: 1, F1: (v19 >> ((0) & 63)), F2: (v32 - v32)}
                                                _ = v35
                                            } else {
                                                if v1 {
                                                    switch (^(fi23)) {
                                                    case -9223372036854775808:
                                                        switch "\n\t" {
                                                        case "hello":
                                                            fmt.Println(v19, v2)
                                                        case "foo bar":
                                                            fmt.Println(v2)
                                                        case "\n\t":
                                                            fmt.Println(v19, v5, v19)
                                                        }
                                                    case 1:
                                                        switch int(v32) {
                                                        case 1:
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                v11 = uint32(v6)
                                                            default:
                                                                fmt.Println(v2, v6, v5)
                                                                if ((^(v11)) > (v11 - (v11 >> ((255) & 63)))) {
                                                                    v36 := v4[fi23]
                                                                    if (-526 != (fi23 + fi22)) {
                                                                        fmt.Println(v11, v1)
                                                                        v2 = 1
                                                                    } else {
                                                                        v37 := GenS6{F0: 0}
                                                                        _ = v37
                                                                    }
                                                                    fmt.Println(v2, v5)
                                                                    _ = v36
                                                                }
                                                                var v38 string = "a"
                                                                _ = v38
                                                            }
                                                        case 9223372036854775807:
                                                            v39 := GenS7{F0: -79, F1: -1, F2: v32}
                                                            _ = v39
                                                        default:
                                                            var v40 float64 = (float64(0.0) - float64((0.0 - ((float32((1.5 - ((-2.25 - 0.1) * float64((-2.25 + ((0.0 * 3.14159) / 1.0)))))) * (0.0 * 0.0)) - 0.1))))
                                                            _ = v40
                                                        }
                                                    case 0:
                                                        delete(v4, (fi23 & -279))
                                                    }
                                                    switch fi23 {
                                                    case 128:
                                                        v6 += 2
                                                    default:
                                                        switch 4294967295 {
                                                        case -672:
                                                            v4[fi22] = (!("" > ("foo bar" + ("\n\t" + ""))))
                                                        default:
                                                            var v41 float64 = 0.1
                                                            _ = v41
                                                        }
                                                    }
                                                }
                                                var v42 int8 = v19
                                                if v1 {
                                                    v43 := GenS3{F0: v32, F1: v32}
                                                    _ = v43
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = (fi22 & (-(fi23)))
                                                    }
                                                    v44 := GenS6{F0: 2147483647}
                                                    var v45 uint64 = 18446744073709551615
                                                    _ = v44
                                                    _ = v45
                                                }
                                                _ = v42
                                            }
                                        }
                                        _ = v32
                                    }
                                }
                                _ = len(v3)
                                fmt.Println(fi23)
                            }
                        case 231:
                            v6 -= v6
                        case 128:
                            v5 = 65535
                        }
                    }
                    if v1 {
                        if v1 {
                            _ = len(v3)
                        }
                    }
                    v46 := GenS1{F0: (v11 >> ((65535) & 63)), F1: 4294967295, F2: ("foo bar" + ("hello" + ("a" + "foo bar")))}
                    _ = fi23
                    _ = v46
                }
                delete(v4, fi22)
                _ = fi22
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
