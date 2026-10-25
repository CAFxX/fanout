package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint8; F1 int32; F2 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint; F1 bool; F2 uint64 }
type GenS6 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int; F1 uint; F2 uint16 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    v3 := []uint32{2}
    var v4 int64 = -1
    var v5 uint64 = 2
    var v6 float64 = 0.1
    fmt.Println(v4)
    switch "\n\t" {
    case "foo bar":
        v4 = v4
    case "\n\t":
        fmt.Println(v5, v4)
    case "a":
        fmt.Println(v6, v4)
    default:
        if ((false && true) && false) {
            switch "hello" {
            case "a":
                var v7 int64 = -9223372036854775808
                _ = v7
            case "foo bar":
                fmt.Println(v4, v4, v4)
            case "\n\t":
                v8 := GenS1{F0: 707, F1: 725, F2: 2}
                _ = v8
            default:
                delete(v2, "foo bar")
            }
            v9 := v1[""]
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 string = ""
                v3 = append(v3, 65535)
                _ = fi10
                _ = v11
            }
            _ = v9
        } else {
            v12 := GenS2{F0: 0, F1: 0}
            _ = v12
        }
    }
    _ = len(v1)
    if false {
        if false {
            var v13 bool = false
            if false {
                v14 := GenS2{F0: -16, F1: 1}
                switch 0 {
                case 2147483647:
                    v4 -= -9223372036854775808
                case 32768:
                    for fi15 := 0; fi15 < 9; fi15++ {
                        v16 := GenS3{F0: 128, F1: 0, F2: ""}
                        v1[("" + ("hello" + "hello"))] = (-2147483648 <= 2147483647)
                        _ = fi15
                        _ = v16
                    }
                default:
                    v17 := GenS4{F0: 42}
                    _ = v17
                }
                for fi18 := 0; fi18 < 17; fi18++ {
                    v19 := GenS5{F0: 990, F1: (v13 && (1 != 2)), F2: (v5 % (((v5 - v5)) | 1))}
                    _ = fi18
                    _ = v19
                }
                _ = v14
            }
            _ = v13
        }
        v5 = v5
    }
    v20 := GenS6{F0: v5, F1: true, F2: 255}
    v21 := GenS2{F0: 1, F1: 0}
    v22 := GenS2{F0: 1, F1: -228}
    for fi23 := 0; fi23 < 12; fi23++ {
        if (-1 < 100) {
            switch fi23 {
            case 1:
                v6 = 1e10
            case -2147483648:
                if true {
                    v3 = append(v3, 536)
                    switch "" {
                    case "\n\t":
                        v24 := GenS2{F0: -590, F1: 42}
                        _ = v24
                    default:
                        _ = len(v2)
                    }
                } else {
                    if (-30 == 42) {
                        var v25 uint = 65535
                        _ = v25
                    }
                    if false {
                        _ = len(v3)
                        v26 := GenS3{F0: 42, F1: 14, F2: ""}
                        _ = v26
                    } else {
                        if false {
                            v4 = (v4 | (v4 - (v4 + -9223372036854775808)))
                            var v27 float64 = v6
                            if true {
                                var v28 uint16 = 255
                                if (("hello" + ("" + "a")) < "a") {
                                    for fi29 := 0; fi29 < 3; fi29++ {
                                        v30 := GenS7{F0: 12}
                                        var v31 float64 = 0.0
                                        _ = fi29
                                        _ = v30
                                        _ = v31
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                } else {
                                    var v32 float32 = float32(v6)
                                    for fi33 := 0; fi33 < 12; fi33++ {
                                        var v34 uint32 = 65535
                                        v35 := GenS2{F0: 0, F1: -676}
                                        _ = fi33
                                        _ = v34
                                        _ = v35
                                    }
                                    var v36 string = ("" + "a")
                                    _ = v32
                                    _ = v36
                                }
                                _ = v28
                            }
                            _ = v27
                        } else {
                            fmt.Println(v4, v4, fi23)
                        }
                        var v37 uint16 = 636
                        var v38 int32 = 0
                        _ = v37
                        _ = v38
                    }
                    if false {
                        v4 = v4
                        if true {
                            var v39 int8 = 1
                            _ = v39
                        } else {
                            switch "" {
                            case "\n\t":
                                if false {
                                    v40 := GenS2{F0: -32768, F1: 1}
                                    for fi41 := 0; fi41 < 19; fi41++ {
                                        switch "" {
                                        case "":
                                            if true {
                                                var v42 float64 = float64(((1.5 * -2.25) / 1.0))
                                                _ = v42
                                            } else {
                                                if (!false) {
                                                    v3 = append(v3, 1)
                                                    fmt.Println(v6, v6, v5)
                                                    _ = len(v3)
                                                } else {
                                                    var v43 uint32 = 2
                                                    fmt.Println(v4, fi41)
                                                    if false {
                                                        var v44 int32 = 1
                                                        var v45 float64 = (float64(0.1) - 1e10)
                                                        _ = v44
                                                        _ = v45
                                                    } else {
                                                        fmt.Println(v43, v5, v5)
                                                        switch (("foo bar" + "hello") + ((("hello" + "hello") + "") + "a")) {
                                                        case "hello":
                                                            var v46 int16 = 91
                                                            _ = v46
                                                        case "a":
                                                            var v47 uint8 = 34
                                                            _ = v47
                                                        }
                                                    }
                                                    _ = v43
                                                }
                                                v48 := GenS7{F0: fi23}
                                                _ = v48
                                            }
                                        case "a":
                                            var v49 int32 = 354
                                            _ = v49
                                        case "\n\t":
                                            v6 += 0.0
                                            switch "hello" {
                                            case "foo bar":
                                                v5 += (v5 / ((4294967295) | 1))
                                            case "hello":
                                                if false {
                                                    var v50 uint64 = v5
                                                    var v51 uint64 = 2
                                                    _ = v50
                                                    _ = v51
                                                }
                                            case "\n\t":
                                                v3 = append(v3, 4294967295)
                                            default:
                                                var v52 int16 = 42
                                                _ = v52
                                            }
                                            delete(v1, "")
                                        }
                                        v53 := GenS8{F0: 127}
                                        _ = fi41
                                        _ = v53
                                    }
                                    _ = v40
                                }
                            case "a":
                                switch (fi23 ^ fi23) {
                                case 0:
                                    v6 -= ((v6 * (3.14159 + v6)) * v6)
                                case 127:
                                    v5 -= 4294967295
                                case -9223372036854775808:
                                    delete(v2, "foo bar")
                                default:
                                    v5 += (v5 % ((v5) | 1))
                                }
                            default:
                                v5 += 18446744073709551615
                            }
                            if (false && ((fi23 >> ((27) & 63)) <= (fi23 & -9223372036854775808))) {
                                var v54 uint8 = 16
                                var v55 int32 = 0
                                _ = v54
                                _ = v55
                            }
                        }
                        switch ("" + "") {
                        case "hello":
                            v56 := GenS6{F0: v5, F1: false, F2: 42}
                            _ = v56
                        case "":
                            var v57 uint64 = (v5 % (((v5 << ((0) & 63))) | 1))
                            _ = v57
                        case "a":
                            if (true || (0 >= 0)) {
                                var v58 string = "foo bar"
                                v6 -= (v6 + -2.25)
                                if true {
                                    v59 := GenS9{F0: fi23, F1: 525, F2: 255}
                                    v4 = v4
                                    _ = v59
                                }
                                _ = v58
                            }
                        default:
                            v5 = v5
                        }
                    } else {
                        var v60 uint = 113
                        v3 = append(v3, 2147483648)
                        _ = v60
                    }
                }
            case 458:
                fmt.Println(v5)
            default:
                v6 += -2.25
            }
        } else {
            v4 = (v4 ^ (^(v4)))
        }
        var v61 uint32 = uint32(v5)
        _ = fi23
        _ = v61
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
