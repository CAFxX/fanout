package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint32; F2 int }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 bool; F1 int32; F2 int }
type GenS6 struct { F0 int; F1 int32 }

func main() {
    v1 := []string{("a" + ((("a" + "hello") + ("foo bar" + "hello")) + "\n\t"))}
    var v2 float32 = 1.5
    v3 := []string{"foo bar", ("a" + ""), "hello"}
    if false {
        fmt.Println(v2)
        v2 += 3.14159
    } else {
        v4 := GenS1{F0: -1}
        for fi5 := 0; fi5 < 12; fi5++ {
            for fi6 := 0; fi6 < 17; fi6++ {
                var v7 uint = 573
                var v8 int = (fi6 + fi5)
                if false {
                    if (0 > 2) {
                        v2 = v2
                        v9 := GenS2{F0: 128, F1: 42}
                        var v10 float32 = -2.25
                        if false {
                            var v11 int = 4294967295
                            v12 := GenS2{F0: 128, F1: 73}
                            _ = v11
                            _ = v12
                        }
                        switch fi6 {
                        case 1:
                            v13 := GenS3{F0: 1, F1: 603, F2: fi5}
                            _ = v13
                        case -1:
                            var v14 uint16 = 32767
                            _ = v14
                        default:
                            var v15 uint16 = 42
                            _ = v15
                        }
                        if (fi5 >= (fi6 / ((4294967295) | 1))) {
                            v2 = (v2 + (v10 - float32(0.0)))
                            if ("" > "foo bar") {
                                var v16 uint64 = 4294967295
                                _ = v16
                            } else {
                                var v17 uint8 = 255
                                v18 := GenS3{F0: -2147483648, F1: 42, F2: 127}
                                var v19 int8 = int8(v17)
                                if ((v17 >> ((v7) & 63)) < v17) {
                                    var v20 float32 = (float32(1.5) - (v10 - v10))
                                    var v21 int8 = (^(v19))
                                    _ = v20
                                    _ = v21
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = "foo bar"
                                    }
                                    if false {
                                        if (!(!true)) {
                                            v22 := GenS3{F0: -1, F1: 4294967295, F2: v8}
                                            _ = v22
                                        }
                                        var v23 float64 = -2.25
                                        v23 = (float64(((v10 - float32(v23)) * 1e10)) / 1e10)
                                        _ = v23
                                    } else {
                                        if (65535 < 2) {
                                            v10 = v2
                                            v24 := GenS2{F0: (v17 & v17), F1: (^(v17))}
                                            if ((!false) && false) {
                                                var v25 uint64 = 0
                                                v26 := GenS4{F0: ("a" + ("a" + "\n\t")), F1: -9223372036854775808}
                                                _ = v25
                                                _ = v26
                                            }
                                            _ = v24
                                        } else {
                                            var v27 float32 = (1e10 - float32(float64((1e10 + v10))))
                                            fmt.Println(v7)
                                            v28 := GenS2{F0: (v17 & v17), F1: 0}
                                            fmt.Println(v10, fi5, v19)
                                            var v29 int32 = -1
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        }
                                        fmt.Println(v17)
                                        if false {
                                            switch 378 {
                                            case 646:
                                                v30 := GenS5{F0: false, F1: 637, F2: v8}
                                                var v31 uint8 = 101
                                                var v32 uint32 = 591
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                            default:
                                                v19 = v19
                                            }
                                            _ = len(v3)
                                        }
                                    }
                                    var v33 uint8 = (v17 % (((v17 ^ v17)) | 1))
                                    _ = v33
                                }
                                v17 += (v17 % (((v17 - (v17 + (^(v17))))) | 1))
                                _ = v17
                                _ = v18
                                _ = v19
                            }
                        }
                        _ = v10
                        _ = v9
                    } else {
                        var v34 float64 = ((float64(0.1) / (-2.25 / 1.0)) + float64((3.14159 - (v2 + ((0.0 * v2) * v2)))))
                        _ = v34
                    }
                    switch fi6 {
                    case 9223372036854775807:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case 4294967295:
                        switch ((("hello" + "hello") + ("" + "foo bar")) + "") {
                        case "a":
                            var v35 int = (fi5 ^ (fi5 / (((fi5 * (^(fi5)))) | 1)))
                            _ = v35
                        case "foo bar":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "hello":
                            v2 += float32((3.14159 * 3.14159))
                        }
                    default:
                        var v36 string = ("\n\t" + "a")
                        _ = v36
                    }
                }
                _ = fi6
                _ = v7
                _ = v8
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi5
        }
        var v37 bool = (42 >= 0)
        _ = v37
        _ = v4
    }
    switch "foo bar" {
    case "foo bar":
        v38 := GenS4{F0: ((("" + "hello") + "\n\t") + "\n\t"), F1: 9223372036854775807}
        _ = v38
    }
    switch "a" {
    case "foo bar":
        v2 = (v2 / v2)
    case "a":
        v39 := GenS6{F0: 4294967295, F1: -2147483648}
        _ = v39
    case "hello":
        fmt.Println(v2)
    }
    for fi40 := 0; fi40 < 12; fi40++ {
        var v41 int16 = 0
        fmt.Println(v2)
        _ = len(v3)
        _ = fi40
        _ = v41
    }
    if true {
        var v42 uint32 = 760
        _ = v42
    }
    var v43 uint16 = 2
    v2 = v2
    for fi44 := 0; fi44 < 9; fi44++ {
        v2 = (v2 - float32(float64(v2)))
        v3 = append(v3, "a")
        switch (("foo bar" + "\n\t") + "\n\t") {
        case "hello":
            var v45 int8 = 19
            _ = v45
        }
        _ = fi44
    }
    var v46 uint8 = 76
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v46
}
