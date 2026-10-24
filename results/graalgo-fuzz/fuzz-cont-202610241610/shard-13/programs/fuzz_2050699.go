package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int }

func main() {
    var v1 float32 = 3.14159
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]string)
    v4 := []string{"hello", ""}
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]string)
    var v7 int = -2147483648
    v8 := GenS1{F0: 0, F1: 42, F2: 127}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v1, v2)
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int8 = 100
        switch 32767 {
        case -93:
            fmt.Println(v2, v1)
        case 42:
            var v11 uint16 = 65535
            _ = v11
        case 32768:
            var v12 uint16 = 42
            _ = v12
        default:
            for fi13 := 0; fi13 < 10; fi13++ {
                var v14 string = (("" + ("hello" + ("a" + "\n\t"))) + ("\n\t" + "foo bar"))
                _ = fi13
                _ = v14
            }
        }
        _ = fi9
        _ = v10
    }
    var v15 int8 = 0
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 bool = (true && (2147483647 < 2147483647))
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    var v19 int = v7
    delete(v6, (v7 + 256))
    for fi20 := 0; fi20 < 6; fi20++ {
        if true {
            if (42 > 255) {
                if (!((3.14159 < 3.14159) || true)) {
                    if ("\n\t" != "\n\t") {
                        if len(v4) > 0 {
                            v4[0] = ("a" + "")
                        }
                    }
                    if (!(25 == 0)) {
                        switch "\n\t" {
                        case "a":
                            v2 = v2
                        case "":
                            if (84 == 0) {
                                fmt.Println(v7, fi20)
                                if (37 != 176) {
                                    switch "hello" {
                                    case "foo bar":
                                        var v21 bool = false
                                        _ = v21
                                    case "a":
                                        switch "foo bar" {
                                        case "\n\t":
                                            v2 = 18446744073709551615
                                        }
                                    case "hello":
                                        if true {
                                            var v22 int = fi20
                                            fmt.Println(v22)
                                            switch v19 {
                                            case 127:
                                                v15 += (v15 << ((18446744073709551615) & 63))
                                            case 9223372036854775807:
                                                _ = len(v3)
                                            case -545:
                                                switch 128 {
                                                case 9223372036854775807:
                                                    var v23 bool = false
                                                    _ = v23
                                                default:
                                                    for fi24 := 0; fi24 < 12; fi24++ {
                                                        switch (v7 % ((32767) | 1)) {
                                                        case 42:
                                                            v15 = int8(v2)
                                                        case -9223372036854775808:
                                                            v22 = (fi24 % ((v22) | 1))
                                                        }
                                                        _ = fi24
                                                    }
                                                }
                                            }
                                            _ = v22
                                        } else {
                                            if false {
                                                v25 := GenS1{F0: (v15 - 1), F1: 4294967295, F2: 2}
                                                v26 := v6[v19]
                                                _ = v25
                                                _ = v26
                                            }
                                            var v27 string = ("hello" + ("\n\t" + "\n\t"))
                                            v28 := GenS1{F0: v15, F1: 18446744073709551615, F2: 114}
                                            _ = v27
                                            _ = v28
                                        }
                                    default:
                                        switch ("a" + "a") {
                                        case "foo bar":
                                            v4 = append(v4, ("\n\t" + ("\n\t" + ("" + "hello"))))
                                        case "a":
                                            v7 = int(v2)
                                        case "":
                                            var v29 uint16 = 1
                                            _ = v29
                                        }
                                    }
                                    if (false && true) {
                                        var v30 float64 = ((0.1 - -2.25) * 3.14159)
                                        v31 := v6[v19]
                                        _ = v30
                                        _ = v31
                                    }
                                } else {
                                    v1 = v1
                                    var v32 int16 = -32768
                                    fmt.Println(v2, v19)
                                    _ = v32
                                }
                                if false {
                                    v33 := v5[972]
                                    var v34 uint8 = 219
                                    _ = v33
                                    _ = v34
                                }
                            } else {
                                var v35 uint8 = 154
                                _ = v35
                            }
                        }
                        var v36 int = (^(v7))
                        _ = v36
                    }
                    v19 -= 9223372036854775807
                } else {
                    if true {
                        var v37 float64 = 1.5
                        switch (-(v7)) {
                        case -9223372036854775808:
                            for fi38 := 0; fi38 < 1; fi38++ {
                                if true {
                                    if false {
                                        switch ("hello" + ("" + "hello")) {
                                        case "a":
                                            var v39 string = "a"
                                            _ = v39
                                        case "":
                                            var v40 bool = (4294967295 >= 2147483648)
                                            _ = v40
                                        }
                                        if false {
                                            v37 += v37
                                            var v41 float64 = 0.0
                                            v19 += (v19 | 42)
                                            _ = v41
                                        } else {
                                            _ = len(v4)
                                        }
                                        if (1 != 1) {
                                            v4 = append(v4, "a")
                                            v2 += (^(v2))
                                            var v42 int32 = 0
                                            _ = v42
                                        }
                                    }
                                } else {
                                    switch ("a" + ((("\n\t" + "a") + "a") + "")) {
                                    case "hello":
                                        var v43 uint = 929
                                        fmt.Println(v37, v2)
                                        var v44 int64 = -9223372036854775808
                                        _ = v43
                                        _ = v44
                                    }
                                    v45 := GenS1{F0: -128, F1: (v2 & (v2 | v2)), F2: 1}
                                    var v46 uint = 255
                                    _ = v45
                                    _ = v46
                                }
                                v4 = append(v4, (("hello" + "") + "a"))
                                v47 := GenS1{F0: (v15 | v15), F1: 18446744073709551615, F2: 146}
                                _ = fi38
                                _ = v47
                            }
                        default:
                            fmt.Println(v15, fi20)
                        }
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                        _ = v37
                    }
                    if false {
                        var v48 uint8 = 234
                        var v49 uint = 577
                        fmt.Println(v15, v1, v1)
                        _ = v48
                        _ = v49
                    } else {
                        v19 += v7
                    }
                }
            }
            for fi50 := 0; fi50 < 12; fi50++ {
                switch (("a" + "") + ("foo bar" + ("" + "\n\t"))) {
                case "a":
                    switch (v7 * fi50) {
                    case 127:
                        v1 = float32(1.5)
                    case -9223372036854775808:
                        delete(v6, -583)
                    case -812:
                        switch v7 {
                        case -636:
                            switch v19 {
                            case 0:
                                var v51 float64 = ((3.14159 + 3.14159) + 3.14159)
                                _ = v51
                            default:
                                v52 := GenS1{F0: v15, F1: (v2 | (^(v2))), F2: 241}
                                _ = v52
                            }
                        case 558:
                            v15 -= 42
                        }
                    default:
                        fmt.Println(v2, fi50, fi50)
                    }
                case "foo bar":
                    var v53 int16 = 1
                    _ = v53
                case "":
                    if (882 > 766) {
                        var v54 int32 = 42
                        v55 := GenS1{F0: 37, F1: (v2 / (((v2 - 484)) | 1)), F2: uint8(v2)}
                        var v56 string = (((((("foo bar" + "foo bar") + ("hello" + ("" + "a"))) + "foo bar") + "foo bar") + "a") + "\n\t")
                        _ = v54
                        _ = v55
                        _ = v56
                    } else {
                        v57 := GenS2{F0: (v7 >> ((18446744073709551615) & 63))}
                        v4 = append(v4, "")
                        v58 := GenS1{F0: (v15 + v15), F1: 4294967295, F2: 255}
                        _ = v57
                        _ = v58
                    }
                default:
                    var v59 uint = 42
                    _ = v59
                }
                v4 = append(v4, (("" + "foo bar") + (("a" + "") + "a")))
                _ = fi50
            }
        } else {
            v60 := v3[2]
            _ = v60
        }
        v7 = fi20
        var v61 uint64 = (^(v2))
        _ = fi20
        _ = v61
    }
    var v62 uint16 = 0
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v62:", v62)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v7
    _ = v8
}
