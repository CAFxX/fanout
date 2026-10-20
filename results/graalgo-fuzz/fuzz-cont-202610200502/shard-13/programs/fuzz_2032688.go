package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int; F1 int16; F2 int64 }
type GenS3 struct { F0 int; F1 uint; F2 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 uint; F2 int16 }
func genFn0() uint64 {
    var v7 uint64 = 18446744073709551615
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint = 42
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        v7 += (^(v7))
        _ = fi10
    }
    var v11 uint64 = v7
    _ = v11
    _ = v7
    return v7
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{}
    var v3 uint32 = 600
    var v4 int8 = 42
    var v5 float32 = 0.0
    var v6 uint = uint(v3)
    if (true && (0.0 < (((0.0 + 1.5) - float64(v5)) * -2.25))) {
        v12 := GenS1{F0: v4, F1: 234}
        switch "a" {
        case "\n\t":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        fmt.Println(v4, v3)
        _ = v12
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v13 uint16 = 65535
    if false {
        var v14 uint = 18446744073709551615
        fmt.Println(v14, v5, v5)
        if false {
            v5 = v5
            var v15 float64 = 1.5
            _ = v15
        } else {
            if true {
                for fi16 := 0; fi16 < 18; fi16++ {
                    v6 -= v6
                    _ = fi16
                }
            }
        }
        _ = v14
    } else {
        fmt.Println(v4)
    }
    fmt.Println(v6, v13, v5)
    v4 = v4
    if len(v2) > 0 {
        v2[0] = -947
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        switch "" {
        case "hello":
            var v18 int16 = 1
            _ = v18
        default:
            for fi19 := 0; fi19 < 7; fi19++ {
                fmt.Println(v5, fi17)
                if (true || true) {
                    var v20 bool = true
                    fmt.Println(v5, v3, fi17)
                    _ = v20
                }
                switch fi17 {
                case -9223372036854775808:
                    var v21 int8 = (v4 * v4)
                    _ = v21
                }
                _ = fi19
            }
        }
        fmt.Println(v5, fi17, v6)
        _ = fi17
    }
    v22 := v1[32767]
    v22 = (1 >= 2147483647)
    for fi23 := 0; fi23 < 11; fi23++ {
        _ = len(v2)
        if (255 > 255) {
            fmt.Println(fi23, v6, v22)
            var v24 float32 = v5
            _ = v24
        } else {
            v25 := GenS2{F0: (-(fi23)), F1: 42, F2: -1}
            _ = v25
        }
        _ = fi23
    }
    var v26 uint32 = v3
    v27 := GenS1{F0: v4, F1: (v13 + v13)}
    switch -2147483648 {
    case 32768:
        if ((0.1 * v5) > 3.14159) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = len(v1)
            v6 -= v6
        }
    }
    v6 += v6
    v5 += 0.1
    if ("hello" <= ("\n\t" + "hello")) {
        if v22 {
            v28 := v1[int(v3)]
            var v29 uint = v6
            _ = v28
            _ = v29
        } else {
            delete(v1, 2147483647)
            var v30 uint64 = 18446744073709551615
            v6 -= (v6 + 0)
            _ = v30
        }
        switch 0 {
        case -615:
            if ((v22 && (!v22)) || false) {
                switch "foo bar" {
                case "foo bar":
                    var v31 int8 = v4
                    _ = v31
                case "hello":
                    v22 = (!false)
                }
            } else {
                for fi32 := 0; fi32 < 19; fi32++ {
                    v33 := GenS3{F0: (fi32 / (((fi32 * -906)) | 1)), F1: v6, F2: 42}
                    if true {
                        fmt.Println(v13, v22)
                    } else {
                        delete(v1, int(v13))
                        var v34 string = "a"
                        v6 = (^(v6))
                        _ = v34
                    }
                    fmt.Println(v13, v4, v4)
                    _ = fi32
                    _ = v33
                }
            }
        case 799:
            fmt.Println(v26, v13)
        case 32768:
            fmt.Println(v5)
        }
        v6 = (v6 + (v6 >> (((v6 - (v6 * 4294967295))) & 63)))
    }
    if false {
        v26 = 288
        if v22 {
            if false {
                if v22 {
                    fmt.Println(v5)
                    var v35 uint64 = 2
                    for fi36 := 0; fi36 < 20; fi36++ {
                        if true {
                            var v37 uint64 = 1
                            for fi38 := 0; fi38 < 12; fi38++ {
                                fmt.Println(fi36)
                                if (222 != 42) {
                                    _ = len(v2)
                                    fmt.Println(v3, v3)
                                    var v39 uint16 = (v13 - (v13 | uint16(v6)))
                                    _ = v39
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    switch "a" {
                                    case "a":
                                        v40 := GenS1{F0: (^(v4)), F1: (v13 ^ v13)}
                                        _ = v40
                                    default:
                                        v13 += (^(v13))
                                    }
                                    fmt.Println(v6, fi38)
                                }
                                _ = fi38
                            }
                            switch (fi36 >> ((uint(v13)) & 63)) {
                            case -554:
                                switch ("\n\t" + "foo bar") {
                                case "foo bar":
                                    var v41 uint16 = v13
                                    _ = v41
                                default:
                                    if (42 <= -1) {
                                        var v42 int16 = int16(v26)
                                        _ = v42
                                    } else {
                                        var v43 uint8 = 170
                                        var v44 float64 = 0.0
                                        _ = v43
                                        _ = v44
                                    }
                                }
                            case -463:
                                for fi45 := 0; fi45 < 17; fi45++ {
                                    var v46 uint8 = 21
                                    _ = fi45
                                    _ = v46
                                }
                            case -2147483648:
                                var v47 int64 = 0
                                _ = v47
                            }
                            _ = v37
                        } else {
                            var v48 int32 = -188
                            switch ("" + "a") {
                            case "hello":
                                v2 = append(v2, (fi36 >> ((v6) & 63)))
                            case "\n\t":
                                v1[98] = v22
                            case "foo bar":
                                v49 := GenS4{F0: "a"}
                                _ = v49
                            }
                            v35 = (v35 & v35)
                            _ = v48
                        }
                        if ((true && v22) || true) {
                            v26 = v3
                        }
                        v6 = 4294967295
                        _ = fi36
                    }
                    _ = v35
                }
                var v50 int32 = 2147483647
                switch ("foo bar" + ("hello" + (("" + ((("" + (("\n\t" + "foo bar") + "hello")) + "hello") + (("foo bar" + "foo bar") + ""))) + ("a" + "\n\t")))) {
                case "\n\t":
                    fmt.Println(v26, v6)
                case "":
                    fmt.Println(v26)
                case "hello":
                    v51 := v1[-421]
                    _ = v51
                default:
                    v4 -= (v4 & 88)
                }
                _ = v50
            } else {
                for fi52 := 0; fi52 < 4; fi52++ {
                    v53 := GenS2{F0: -141, F1: 0, F2: -1}
                    _ = fi52
                    _ = v53
                }
            }
        }
    } else {
        for fi54 := 0; fi54 < 16; fi54++ {
            for fi55 := 0; fi55 < 13; fi55++ {
                fmt.Println(v5, v22, v6)
                v56 := GenS5{F0: v22, F1: (v6 | v6), F2: 42}
                _ = fi55
                _ = v56
            }
            _ = fi54
        }
        for fi57 := 0; fi57 < 3; fi57++ {
            v2 = append(v2, fi57)
            switch "a" {
            case "":
                v58 := GenS2{F0: (fi57 / (((fi57 | (fi57 >> ((65535) & 63)))) | 1)), F1: 512, F2: 0}
                _ = v58
            case "foo bar":
                var v59 int8 = (v4 ^ v4)
                _ = v59
            }
            _ = fi57
        }
        switch 9223372036854775807 {
        case 1:
            v60 := GenS2{F0: -844, F1: 363, F2: -1}
            _ = v60
        default:
            for fi61 := 0; fi61 < 6; fi61++ {
                v62 := GenS1{F0: 127, F1: (v13 & v13)}
                _ = fi61
                _ = v62
            }
        }
    }
    if len(v2) > 0 {
        v2[0] = -513
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
