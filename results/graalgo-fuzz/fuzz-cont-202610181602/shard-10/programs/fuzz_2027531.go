package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS7 struct { F0 int16; F1 string }
type GenS8 struct { F0 string; F1 int32; F2 bool }

func main() {
    v1 := []int{}
    var v2 int = 0
    var v3 uint64 = 816
    if (!(false && true)) {
        fmt.Println(v3)
        var v4 int8 = -42
        v5 := GenS1{F0: -32768}
        _ = v4
        _ = v5
    } else {
        v6 := GenS2{F0: v3, F1: 959}
        switch (-(v2)) {
        case 872:
            fmt.Println(v3, v3)
        case 128:
            fmt.Println(v2)
        case 1:
            switch ("a" + ("\n\t" + "hello")) {
            case "hello":
                switch v2 {
                case 1:
                    fmt.Println(v3)
                case -1:
                    if true {
                        if (!(false && (1.5 != 1e10))) {
                            if len(v1) > 0 {
                                v1[0] = v2
                            }
                        } else {
                            _ = len(v1)
                            for fi7 := 0; fi7 < 17; fi7++ {
                                if false {
                                    if ((42 != 85) || true) {
                                        fmt.Println(v2, v2, v2)
                                    } else {
                                        fmt.Println(v3, v3)
                                    }
                                }
                                v8 := GenS3{F0: 0, F1: "a"}
                                _ = fi7
                                _ = v8
                            }
                        }
                        if true {
                            for fi9 := 0; fi9 < 18; fi9++ {
                                for fi10 := 0; fi10 < 4; fi10++ {
                                    v3 -= v3
                                    v2 = (fi10 + 32767)
                                    _ = fi10
                                }
                                v3 = (v3 ^ (v3 * (v3 & v3)))
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = fi9
                            }
                            var v11 int32 = -2147483648
                            _ = v11
                        } else {
                            var v12 int8 = 0
                            if (-422 >= 9223372036854775807) {
                                if (true || (2147483647 <= 2147483647)) {
                                    v13 := GenS4{F0: 200, F1: 11, F2: 900}
                                    _ = v13
                                }
                                v12 -= v12
                            }
                            _ = v12
                        }
                    } else {
                        v2 = -45
                    }
                default:
                    v2 = v2
                }
            case "\n\t":
                fmt.Println(v3, v3)
            }
        }
        _ = v6
    }
    switch "hello" {
    case "a":
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 int64 = -580
            for fi16 := 0; fi16 < 10; fi16++ {
                fmt.Println(fi14)
                var v17 int32 = 42
                v1 = append(v1, 127)
                switch (-(fi16)) {
                case -709:
                    var v18 float64 = (1e10 * 3.14159)
                    _ = v18
                case -2147483648:
                    if (true || (v3 > (^(v3)))) {
                        var v19 int16 = 32767
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if false {
                                v2 += -2147483648
                            }
                        }
                        _ = v19
                    }
                }
                v17 = v17
                _ = fi16
                _ = v17
            }
            _ = fi14
            _ = v15
        }
    case "hello":
        var v20 int8 = -9
        _ = v20
    case "\n\t":
        for fi21 := 0; fi21 < 18; fi21++ {
            fmt.Println(v3, v3)
            v3 = 608
            _ = fi21
        }
    default:
        v22 := GenS4{F0: 151, F1: 42, F2: 2}
        _ = v22
    }
    v23 := GenS5{F0: 2, F1: 722}
    switch v2 {
    case -9223372036854775808:
        var v24 uint = 18446744073709551615
        _ = v24
    case 1:
        v25 := GenS3{F0: 100, F1: "foo bar"}
        _ = v25
    case -930:
        switch (v2 | v2) {
        case 32768:
            v26 := GenS1{F0: 42}
            _ = v26
        case -1:
            v27 := GenS3{F0: 127, F1: "a"}
            _ = v27
        case -2147483648:
            var v28 int64 = -726
            _ = v28
        default:
            v29 := GenS6{F0: 4294967295, F1: -1, F2: 255}
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 9; fi30++ {
        v3 -= 1
        var v31 uint16 = 2
        _ = fi30
        _ = v31
    }
    var v32 uint16 = 255
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = (v2 % (((v2 ^ v2)) | 1))
    }
    _ = len(v1)
    var v33 int = -9223372036854775808
    v32 = (^(v32))
    if (0.1 >= (3.14159 / 1.0)) {
        if true {
            v32 = v32
            v3 = 18446744073709551615
        }
        for fi34 := 0; fi34 < 12; fi34++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi34
        }
    } else {
        for fi35 := 0; fi35 < 10; fi35++ {
            fmt.Println(v33, v33)
            var v36 string = ""
            _ = fi35
            _ = v36
        }
        v32 = 0
    }
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        var v37 bool = (9223372036854775807 >= -1)
        _ = v37
    case "hello":
        for fi38 := 0; fi38 < 9; fi38++ {
            fmt.Println(v3, v2, fi38)
            var v39 uint32 = 4294967295
            _ = fi38
            _ = v39
        }
    case "":
        for fi40 := 0; fi40 < 8; fi40++ {
            if (((("" + ("hello" + ("hello" + ("foo bar" + "foo bar")))) + "a") + "") != "a") {
                fmt.Println(v33, fi40)
                for fi41 := 0; fi41 < 17; fi41++ {
                    switch (("hello" + "hello") + "\n\t") {
                    case "hello":
                        var v42 uint16 = v32
                        _ = v42
                    case "foo bar":
                        v43 := GenS7{F0: int16(v3), F1: "hello"}
                        _ = v43
                    case "\n\t":
                        var v44 uint64 = v3
                        _ = v44
                    }
                    var v45 uint16 = (v32 * (^(v32)))
                    v46 := GenS4{F0: 128, F1: 65535, F2: 42}
                    fmt.Println(fi40, v32)
                    var v47 string = ("\n\t" + "foo bar")
                    _ = fi41
                    _ = v45
                    _ = v46
                    _ = v47
                }
                v3 -= v3
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            if false {
                v48 := GenS1{F0: 0}
                _ = v48
            } else {
                v3 = (v3 << ((573) & 63))
            }
            _ = fi40
        }
    }
    v2 += v2
    if false {
        if true {
            for fi49 := 0; fi49 < 9; fi49++ {
                var v50 bool = (42 <= uint32(v32))
                _ = fi49
                _ = v50
            }
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
        var v51 int64 = 0
        var v52 int8 = -128
        _ = v51
        _ = v52
    } else {
        switch (v2 * 0) {
        case -9223372036854775808:
            if false {
                var v53 uint64 = v3
                _ = v53
            }
        case 0:
            v33 = int(v32)
        case 4294967295:
            var v54 uint8 = 128
            _ = v54
        default:
            v55 := GenS8{F0: "foo bar", F1: -1, F2: ((v32 | (v32 - 2)) >= (v32 << ((170) & 63)))}
            _ = v55
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v33
}
