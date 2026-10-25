package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int; F2 uint }
type GenS3 struct { F0 int32; F1 uint16; F2 int }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 int16 }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{745, 42, 1, 884}
    var v3 float32 = 1e10
    var v4 int64 = 9223372036854775807
    v5 := make(map[string]int)
    var v6 float32 = 1e10
    for fi7 := 0; fi7 < 16; fi7++ {
        v6 = v6
        for fi8 := 0; fi8 < 18; fi8++ {
            switch (fi8 << ((18446744073709551615) & 63)) {
            case -1:
                if (((-2.25 - 1.5) < (((1e10 * 0.1) / 0.1) / (0.0 - 1.5))) && (true || (42 == 0))) {
                    fmt.Println(v4, fi8, fi7)
                    switch (("\n\t" + ("foo bar" + "")) + "a") {
                    case "hello":
                        switch (fi8 & fi7) {
                        case 127:
                            v4 += v4
                        case 475:
                            var v9 int8 = 93
                            _ = v9
                        }
                    case "":
                        fmt.Println(fi8)
                    case "\n\t":
                        var v10 string = "a"
                        _ = v10
                    default:
                        fmt.Println(v3, v6)
                        switch ("\n\t" + ("a" + "a")) {
                        case "foo bar":
                            v4 = 42
                        }
                        switch "hello" {
                        case "":
                            v11 := GenS1{F0: 755}
                            _ = v11
                        case "\n\t":
                            if true {
                                var v12 int64 = v4
                                _ = v12
                            } else {
                                var v13 uint16 = 42
                                fmt.Println(fi7, v13)
                                if ((v13 / (((v13 ^ (^(v13)))) | 1)) < v13) {
                                    if true {
                                        v2 = append(v2, 673)
                                        var v14 uint8 = 229
                                        if false {
                                            switch (-(fi8)) {
                                            case 0:
                                                v3 = 0.1
                                            }
                                        }
                                        _ = v14
                                    }
                                }
                                _ = v13
                            }
                        }
                    }
                } else {
                    if (true && ("\n\t" == "foo bar")) {
                        switch (("\n\t" + "a") + ("\n\t" + ("\n\t" + "a"))) {
                        case "a":
                            v6 += ((v6 - (v6 * 3.14159)) * (1.5 * 3.14159))
                        case "foo bar":
                            if len(v2) > 0 {
                                v2[0] = 4294967295
                            }
                        case "hello":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v6 = 1.5
                            v4 = 244
                        }
                        v6 += v3
                        fmt.Println(v4, fi7)
                        switch ("hello" + ("" + "\n\t")) {
                        case "":
                            var v15 uint64 = 1
                            _ = v15
                        case "hello":
                            fmt.Println(fi7, fi8)
                            if len(v2) > 0 {
                                v2[0] = 4294967295
                            }
                        case "\n\t":
                            if true {
                                switch "\n\t" {
                                case "a":
                                    v4 += v4
                                default:
                                    var v16 int = (fi8 >> ((506) & 63))
                                    _ = v16
                                }
                                v3 = v6
                                v17 := GenS2{F0: 127, F1: (-(fi7)), F2: 42}
                                v3 = 0.1
                                _ = v17
                            }
                            var v18 uint8 = 127
                            if false {
                                v6 = v3
                                if (true || ((float64(v3) * ((1e10 / 1e10) - (0.0 / 0.1))) > -2.25)) {
                                    var v19 int = fi8
                                    _ = v19
                                } else {
                                    v18 -= 255
                                    v6 = v3
                                }
                            }
                            _ = v18
                        default:
                            var v20 string = ("" + "")
                            _ = v20
                        }
                        v21 := v5[((("\n\t" + "\n\t") + "hello") + ((("\n\t" + ((("\n\t" + "\n\t") + "hello") + "hello")) + (((("\n\t" + ("hello" + (("\n\t" + "") + ""))) + "\n\t") + "a") + "\n\t")) + "a"))]
                        _ = v21
                    }
                    v22 := GenS3{F0: 2147483647, F1: 499, F2: (-(fi8))}
                    if false {
                        var v23 int8 = 42
                        _ = v23
                    } else {
                        var v24 uint8 = 1
                        fmt.Println(v4, v3)
                        _ = v24
                    }
                    _ = v22
                }
            case 32768:
                v3 -= 1.5
            case 0:
                var v25 uint32 = 2
                _ = v25
            default:
                v26 := GenS4{F0: 1, F1: 0}
                _ = v26
            }
            if (2 < 65535) {
                var v27 int64 = (v4 ^ (v4 / ((137) | 1)))
                v6 = (float32(1.5) + ((v6 - (v6 / (1e10 - v3))) / ((1e10 * 0.0) * (v6 + v6))))
                if len(v2) > 0 {
                    v2[0] = 2147483648
                }
                _ = v27
            }
            v3 = 1.5
            _ = fi8
        }
        _ = fi7
    }
    switch -1 {
    case 4294967295:
        v28 := GenS5{F0: 42}
        _ = v28
    default:
        var v29 int32 = -1
        _ = v29
    }
    var v30 float64 = 3.14159
    v6 -= (v3 / 3.14159)
    var v31 string = "\n\t"
    v32 := GenS3{F0: 42, F1: 65535, F2: 1}
    fmt.Println(v3, v6, v6)
    for fi33 := 0; fi33 < 9; fi33++ {
        switch (fi33 | 128) {
        case 920:
            switch (fi33 | fi33) {
            case -140:
                var v34 float64 = ((((v30 + ((3.14159 / v30) - ((((-2.25 / (v30 / (0.1 * 1e10))) + 3.14159) + (float64((v6 + v3)) - 3.14159)) + v30))) / (((1.5 * (float64(float32(v30)) / ((v30 - 3.14159) + v30))) + 1e10) / 1.0)) + (v30 + v30)) + v30)
                _ = v34
            default:
                var v35 float32 = (float32((float64((v6 + (((((-2.25 - float32((v30 / (v30 / float64((v6 + v6)))))) * v6) * (v3 * 1.5)) * v6) + (v3 * 0.1)))) * (v30 * v30))) / v6)
                _ = v35
            }
        }
        v36 := GenS4{F0: -341, F1: -1}
        _ = fi33
        _ = v36
    }
    var v37 int = -9223372036854775808
    v6 = (((((v3 + float32(v30)) + v3) / v6) - float32(float64(v6))) + v6)
    if false {
        var v38 int32 = -1
        switch v37 {
        case -1:
            v39 := GenS2{F0: 42, F1: v37, F2: 1}
            _ = v39
        default:
            var v40 int32 = 0
            _ = v40
        }
        _ = v38
    }
    switch "\n\t" {
    case "":
        v41 := GenS4{F0: 42, F1: 0}
        _ = v41
    case "a":
        for fi42 := 0; fi42 < 3; fi42++ {
            var v43 int16 = 833
            v3 -= v6
            var v44 uint16 = 2
            _ = fi42
            _ = v43
            _ = v44
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
