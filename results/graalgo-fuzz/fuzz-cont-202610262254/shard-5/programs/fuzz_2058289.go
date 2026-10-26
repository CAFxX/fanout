package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 int; F2 uint16 }
type GenS5 struct { F0 int32 }

func main() {
    var v1 string = ("hello" + ("a" + ("" + (("hello" + ((("hello" + "foo bar") + ("hello" + ("a" + "foo bar"))) + "hello")) + "foo bar"))))
    v2 := make(map[string]string)
    v3 := make(map[string]uint32)
    var v4 uint32 = 65535
    v5 := make(map[int]string)
    if true {
        fmt.Println(v4)
    }
    var v6 uint32 = (v4 % ((v4) | 1))
    var v7 int64 = -603
    for fi8 := 0; fi8 < 12; fi8++ {
        if ((((v1 + v1) + (v1 + (v1 + ((("foo bar" + (v1 + (("" + "hello") + v1))) + "foo bar") + v1)))) <= (v1 + v1)) || false) {
            var v9 int64 = v7
            _ = v9
        } else {
            v1 = v1
            v10 := GenS1{F0: v1}
            _ = v10
        }
        var v11 int8 = 6
        var v12 int64 = v7
        _ = fi8
        _ = v11
        _ = v12
    }
    v7 = (^(v7))
    v6 = 4294967295
    v13 := []int{127, -240, 9223372036854775807, 32768}
    for fi14 := 0; fi14 < 12; fi14++ {
        v7 -= (v7 * v7)
        var v15 int = 211
        _ = fi14
        _ = v15
    }
    v6 = 2
    if false {
        var v16 int8 = 127
        _ = v16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 int64 = (^(v7))
        fmt.Println(v1, fi17)
        _ = fi17
        _ = v18
    }
    var v19 int16 = 259
    v20 := GenS2{F0: 42}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v21 := GenS3{F0: 32767}
    switch "" {
    case "":
        for fi22 := 0; fi22 < 16; fi22++ {
            fmt.Println(v19, fi22)
            if (((!(v19 == -1)) && (1 > 2)) && (42 <= 139)) {
                var v23 uint16 = 65535
                if (((1.5 - 1.5) <= 0.1) || false) {
                    v24 := GenS3{F0: (^(fi22))}
                    v2["foo bar"] = ""
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 11; fi25++ {
                        var v26 uint = 4294967295
                        if (false || ("a" != "")) {
                            var v27 float64 = (3.14159 * ((((((((3.14159 * 0.0) * 3.14159) * float64(0.0)) / 1e10) - 1e10) / float64((((0.1 * 1.5) - float32((-2.25 / (1.5 - (1.5 * 0.0))))) / 1e10))) / 3.14159) - 1.5))
                            fmt.Println(v1)
                            fmt.Println(fi25, v26)
                            switch 4294967295 {
                            case 4294967295:
                                switch ("" + "a") {
                                case "hello":
                                    if true {
                                        v13 = append(v13, fi25)
                                    } else {
                                        v7 += 482
                                        if len(v13) > 0 {
                                            v13[0] = 4294967295
                                        }
                                    }
                                default:
                                    v26 += v26
                                }
                                v28 := GenS1{F0: v1}
                                v29 := GenS1{F0: v1}
                                switch fi22 {
                                case 0:
                                    fmt.Println(v6)
                                case 32767:
                                    var v30 int = 127
                                    _ = v30
                                default:
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(v26)
                                    case "hello":
                                        if (!(!(2 < v4))) {
                                            var v31 float32 = ((float32(v27) + 1.5) + float32(1.5))
                                            if false {
                                                fmt.Println(v27, v23, v31)
                                                fmt.Println(v31, v6)
                                                v32 := GenS3{F0: fi25}
                                                _ = v32
                                            } else {
                                                var v33 int = (^(fi25))
                                                _ = len(v2)
                                                var v34 int64 = v7
                                                _ = v33
                                                _ = v34
                                            }
                                            _ = v31
                                        }
                                        var v35 string = "hello"
                                        _ = v35
                                    case "a":
                                        switch v1 {
                                        case "a":
                                            if true {
                                                if true {
                                                    switch ((v1 + v1) + (v1 + (("" + "a") + v1))) {
                                                    case "\n\t":
                                                        var v36 float32 = (0.0 + float32(float64(3.14159)))
                                                        _ = v36
                                                    case "hello":
                                                        _ = len(v13)
                                                    case "foo bar":
                                                        var v37 int = fi22
                                                        _ = v37
                                                    default:
                                                        if (2147483647 < 1) {
                                                            v26 = 1
                                                            var v38 int16 = 952
                                                            _ = v38
                                                        } else {
                                                            v39 := GenS2{F0: 4294967295}
                                                            v6 = v6
                                                            _ = v39
                                                        }
                                                    }
                                                    var v40 int16 = (v19 % ((v19) | 1))
                                                    _ = v40
                                                }
                                                v41 := GenS3{F0: fi25}
                                                v27 -= v27
                                                _ = v41
                                            } else {
                                                v42 := GenS2{F0: v4}
                                                fmt.Println(v23, v26, v1)
                                                v43 := v5[(fi25 - fi22)]
                                                _ = v42
                                                _ = v43
                                            }
                                        case "":
                                            if false {
                                                if ("a" < "hello") {
                                                    if (v19 <= (-(v19))) {
                                                        var v44 string = v1
                                                        v45 := GenS4{F0: 65535, F1: (fi25 | fi25), F2: (^(v23))}
                                                        _ = v44
                                                        _ = v45
                                                    } else {
                                                        fmt.Println(v19, fi22, v19)
                                                    }
                                                    delete(v3, v1)
                                                }
                                            }
                                        }
                                    }
                                }
                                _ = v28
                                _ = v29
                            case 2147483647:
                                v46 := GenS5{F0: 0}
                                _ = v46
                            case 128:
                                v27 = ((v27 + v27) - (-2.25 * v27))
                            default:
                                if false {
                                    v47 := GenS4{F0: (v4 / ((uint32(v26)) | 1)), F1: (fi25 * (fi22 % ((int(v23)) | 1))), F2: (v23 << (((^(v26))) & 63))}
                                    v1 = v1
                                    v7 -= v7
                                    if ((true || true) && true) {
                                        _ = len(v3)
                                        _ = len(v2)
                                    }
                                    _ = v47
                                } else {
                                    fmt.Println(v7, fi22)
                                    fmt.Println(v19, fi22)
                                    v26 = 18446744073709551615
                                }
                            }
                            fmt.Println(v7)
                            _ = v27
                        } else {
                            fmt.Println(v26)
                            v3["foo bar"] = (^(v4))
                            v19 = -119
                        }
                        _ = fi25
                        _ = v26
                    }
                    if len(v13) > 0 {
                        v13[0] = 765
                    }
                }
                fmt.Println(v23, v7, v6)
                _ = v23
            } else {
                var v48 float64 = (float64(0.1) * (3.14159 - (0.0 * (3.14159 - float64((1e10 + (1e10 / 1.5)))))))
                if true {
                    for fi49 := 0; fi49 < 13; fi49++ {
                        switch v1 {
                        case "hello":
                            v50 := v2[v1]
                            _ = v50
                        default:
                            v51 := v2["foo bar"]
                            _ = v51
                        }
                        _ = fi49
                    }
                    v4 = (v6 << ((65535) & 63))
                    switch -131 {
                    case 32767:
                        v52 := GenS1{F0: (v1 + ("hello" + "a"))}
                        _ = v52
                    }
                } else {
                    fmt.Println(v4)
                    if (65535 != 65535) {
                        switch "hello" {
                        case "":
                            v53 := v2[(v1 + v1)]
                            _ = v53
                        case "a":
                            if false {
                                fmt.Println(v4)
                            }
                        case "foo bar":
                            var v54 string = (v1 + "a")
                            _ = v54
                        default:
                            _ = len(v13)
                        }
                    } else {
                        v1 = (v1 + "hello")
                        for fi55 := 0; fi55 < 9; fi55++ {
                            v4 -= 1
                            switch -658 {
                            case -2147483648:
                                _ = len(v13)
                            }
                            _ = fi55
                        }
                    }
                }
                _ = v48
            }
            _ = fi22
        }
    default:
        v13 = append(v13, -2147483648)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
