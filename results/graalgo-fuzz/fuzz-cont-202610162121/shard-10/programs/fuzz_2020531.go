package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 bool; F2 uint8 }
type GenS5 struct { F0 int32; F1 int8; F2 string }

func main() {
    var v1 bool = (true && false)
    var v2 int = 127
    v3 := []string{"\n\t", "foo bar", "\n\t", "a"}
    switch v2 {
    case -9223372036854775808:
        var v4 int32 = -2147483648
        _ = v4
    case 260:
        if v1 {
            fmt.Println(v1)
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v5 := GenS1{F0: "a"}
            _ = v5
        }
    case 4294967295:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        _ = len(v3)
        _ = fi6
    }
    switch (("a" + ("a" + "hello")) + "") {
    case "a":
        if (1 <= -672) {
            for fi7 := 0; fi7 < 8; fi7++ {
                v8 := GenS1{F0: "a"}
                for fi9 := 0; fi9 < 1; fi9++ {
                    v2 += v2
                    _ = fi9
                }
                var v10 uint8 = 0
                _ = fi7
                _ = v10
                _ = v8
            }
            v11 := GenS2{F0: 22, F1: -154}
            var v12 int32 = 2147483647
            _ = v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := GenS3{F0: 380}
        _ = fi13
        _ = v14
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v15 int8 = -1
    fmt.Println(v1)
    var v16 string = "\n\t"
    switch ("\n\t" + v16) {
    case "\n\t":
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 float32 = 3.14159
            if false {
                var v19 int32 = 42
                v15 = v15
                _ = v19
            }
            _ = fi17
            _ = v18
        }
    case "hello":
        var v20 float64 = -2.25
        _ = v20
    case "a":
        switch v2 {
        case 42:
            v21 := GenS3{F0: 9223372036854775807}
            _ = v21
        case 127:
            v1 = v1
        case 128:
            if (float32(1e10) >= (-2.25 / -2.25)) {
                v22 := GenS3{F0: 897}
                for fi23 := 0; fi23 < 12; fi23++ {
                    _ = len(v3)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if (!((v15 == (v15 % (((v15 - v15)) | 1))) && (0.1 > 0.1))) {
                        if ((1e10 - ((1e10 / 0.1) / 1.0)) >= 1e10) {
                            if v1 {
                                var v24 uint = 256
                                _ = v24
                            } else {
                                var v25 bool = v1
                                if v1 {
                                    fmt.Println(fi23, fi23)
                                } else {
                                    for fi26 := 0; fi26 < 15; fi26++ {
                                        v25 = (0 >= 1)
                                        var v27 uint = 4294967295
                                        _ = fi26
                                        _ = v27
                                    }
                                    if (564 != 18446744073709551615) {
                                        if ((127 > 185) && true) {
                                            v16 = v16
                                            var v28 float64 = 3.14159
                                            _ = v28
                                        } else {
                                            fmt.Println(v2, v2, fi23)
                                            _ = len(v3)
                                            fmt.Println(v16, v1)
                                        }
                                        var v29 uint64 = 99
                                        for fi30 := 0; fi30 < 15; fi30++ {
                                            var v31 int32 = -707
                                            _ = fi30
                                            _ = v31
                                        }
                                        _ = v29
                                    } else {
                                        for fi32 := 0; fi32 < 19; fi32++ {
                                            var v33 int8 = (v15 + v15)
                                            if v25 {
                                                _ = len(v3)
                                                if ((^(fi23)) != 583) {
                                                    var v34 bool = ((!(1 <= 128)) && (980 == 2))
                                                    v15 += v15
                                                    _ = v34
                                                }
                                            }
                                            fmt.Println(fi23, v1, fi32)
                                            _ = fi32
                                            _ = v33
                                        }
                                        if (v1 || v25) {
                                            for fi35 := 0; fi35 < 9; fi35++ {
                                                if v25 {
                                                    _ = len(v3)
                                                    v3 = append(v3, v16)
                                                } else {
                                                    var v36 float64 = (1.5 / ((3.14159 - ((-2.25 - 0.0) + 1.5)) * 0.1))
                                                    v36 = (v36 * 1.5)
                                                    var v37 uint16 = 873
                                                    switch v16 {
                                                    case "":
                                                        var v38 bool = v1
                                                        _ = v38
                                                    }
                                                    _ = v36
                                                    _ = v37
                                                }
                                                fmt.Println(v2, v25, fi23)
                                                v16 = "\n\t"
                                                _ = fi35
                                            }
                                            v39 := GenS2{F0: -225, F1: 2147483647}
                                            _ = v39
                                        }
                                    }
                                }
                                _ = v25
                            }
                            _ = len(v3)
                        } else {
                            if (!false) {
                                var v40 bool = false
                                fmt.Println(fi23)
                                fmt.Println(v15, v15, v2)
                                _ = v40
                            }
                            v2 -= -9223372036854775808
                        }
                    }
                    _ = fi23
                }
                _ = v22
            } else {
                switch ("" + "foo bar") {
                case "hello":
                    for fi41 := 0; fi41 < 6; fi41++ {
                        var v42 uint32 = 1
                        switch (v16 + "hello") {
                        case "\n\t":
                            _ = len(v3)
                        case "hello":
                            if len(v3) > 0 {
                                v3[0] = v16
                            }
                        }
                        if ((239 >= 2) && v1) {
                            var v43 int64 = 0
                            for fi44 := 0; fi44 < 20; fi44++ {
                                if v1 {
                                    v3 = append(v3, v16)
                                } else {
                                    if true {
                                        if v1 {
                                            v15 = v15
                                        } else {
                                            v45 := GenS1{F0: (v16 + v16)}
                                            _ = v45
                                        }
                                        var v46 uint8 = 109
                                        _ = v46
                                    }
                                    v47 := GenS4{F0: 32767, F1: (fi44 < 32767), F2: 42}
                                    _ = len(v3)
                                    _ = v47
                                }
                                fmt.Println(v1)
                                _ = fi44
                            }
                            _ = v43
                        }
                        _ = fi41
                        _ = v42
                    }
                case "a":
                    var v48 string = ("hello" + "foo bar")
                    _ = v48
                case "":
                    v49 := GenS5{F0: -1, F1: (v15 % (((v15 / (((v15 >> ((1) & 63))) | 1))) | 1)), F2: v16}
                    _ = v49
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                v50 := GenS5{F0: 1, F1: (v15 - v15), F2: "a"}
                if v1 {
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                }
                _ = v50
            }
        default:
            var v51 string = ("" + v16)
            _ = v51
        }
    default:
        v3 = append(v3, v16)
    }
    v15 = (v15 % ((v15) | 1))
    if len(v3) > 0 {
        v3[0] = (((v16 + v16) + v16) + "\n\t")
    }
    v1 = (1 < 42)
    for fi52 := 0; fi52 < 20; fi52++ {
        if v1 {
            fmt.Println(fi52, v15, v2)
            var v53 float32 = 3.14159
            _ = v53
        } else {
            v54 := GenS2{F0: -893, F1: 1}
            var v55 int64 = 9223372036854775807
            _ = v54
            _ = v55
        }
        fmt.Println(fi52, v1, v2)
        var v56 uint8 = 129
        _ = fi52
        _ = v56
    }
    if v1 {
        if (-9223372036854775808 <= 917) {
            for fi57 := 0; fi57 < 4; fi57++ {
                v15 = (v15 - (v15 % (((v15 | -1)) | 1)))
                v16 = ("foo bar" + v16)
                _ = fi57
            }
            fmt.Println(v2)
        }
        var v58 bool = (0 >= 4294967295)
        var v59 uint16 = 1
        _ = v58
        _ = v59
    }
    for fi60 := 0; fi60 < 20; fi60++ {
        v61 := GenS5{F0: -2147483648, F1: v15, F2: v16}
        var v62 uint = 18446744073709551615
        _ = fi60
        _ = v61
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
