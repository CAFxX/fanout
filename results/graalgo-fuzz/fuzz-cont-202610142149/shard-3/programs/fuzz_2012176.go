package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int16; F1 uint32; F2 int64 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 uint32; F1 uint64; F2 string }
type GenS7 struct { F0 uint64; F1 uint; F2 bool }
type GenS8 struct { F0 int8; F1 string; F2 int64 }

func main() {
    var v1 uint16 = 42
    var v2 uint64 = 1
    var v3 int32 = -185
    switch 9223372036854775807 {
    case 128:
        if (-2147483648 <= v3) {
            v4 := []int{}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v5 := GenS1{F0: v1, F1: 128}
            _ = v4
            _ = v5
        }
    case -672:
        v1 -= v1
    case -2147483648:
        var v6 int8 = -128
        _ = v6
    default:
        var v7 uint32 = 4294967295
        _ = v7
    }
    v8 := []int{496, 42, -9223372036854775808}
    v2 -= 21
    v9 := GenS1{F0: v1, F1: 39}
    if true {
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := GenS1{F0: (v1 ^ v1), F1: fi10}
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            v12 := GenS2{F0: 2, F1: uint64(v1), F2: (v1 >> ((255) & 63))}
            _ = fi10
            _ = v11
            _ = v12
        }
    } else {
        _ = len(v8)
        v13 := GenS3{F0: 570, F1: 65535, F2: 1}
        var v14 uint8 = 128
        _ = v13
        _ = v14
    }
    _ = len(v8)
    switch "foo bar" {
    case "hello":
        var v15 int8 = 42
        _ = v15
    case "":
        if true {
            v1 -= v1
            v1 = v1
        } else {
            if true {
                var v16 int32 = (^(v3))
                _ = v16
            }
            switch 2147483647 {
            case -152:
                v8 = append(v8, -2147483648)
            case 42:
                for fi17 := 0; fi17 < 12; fi17++ {
                    var v18 uint16 = v1
                    _ = fi17
                    _ = v18
                }
            case 128:
                switch "a" {
                case "":
                    v19 := GenS2{F0: 274, F1: (v2 << ((18446744073709551615) & 63)), F2: 0}
                    _ = v19
                case "\n\t":
                    v20 := GenS4{F0: int(v2), F1: -603}
                    _ = v20
                }
            default:
                v8 = append(v8, -9223372036854775808)
            }
            for fi21 := 0; fi21 < 14; fi21++ {
                var v22 string = "foo bar"
                _ = fi21
                _ = v22
            }
        }
    default:
        for fi23 := 0; fi23 < 3; fi23++ {
            if len(v8) > 0 {
                v8[0] = fi23
            }
            switch "a" {
            case "hello":
                if (true && (v1 >= 1)) {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                } else {
                    var v24 int16 = 1
                    var v25 uint = 18446744073709551615
                    fmt.Println(v2)
                    _ = v24
                    _ = v25
                }
            case "foo bar":
                switch "" {
                case "a":
                    var v26 int = (fi23 % (((fi23 | 2147483647)) | 1))
                    _ = v26
                default:
                    var v27 int8 = 91
                    _ = v27
                }
            case "a":
                if (true || false) {
                    fmt.Println(v1)
                }
            }
            var v28 int = fi23
            _ = fi23
            _ = v28
        }
    }
    v3 = (-(v3))
    if (false || ((v1 | (^(v1))) == (^(v1)))) {
        v3 += (v3 ^ v3)
        _ = len(v8)
    }
    v1 -= v1
    v8 = append(v8, 273)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v29 int = 2147483647
    for fi30 := 0; fi30 < 14; fi30++ {
        if true {
            v31 := GenS1{F0: v1, F1: v29}
            v3 = (v3 * -2147483648)
            _ = v31
        } else {
            if (!false) {
                switch "foo bar" {
                case "foo bar":
                    if (true && (0.0 > (1e10 / 0.1))) {
                        var v32 bool = (("\n\t" + ("" + "foo bar")) <= ("foo bar" + ((("foo bar" + "a") + "foo bar") + "hello")))
                        _ = v32
                    }
                case "":
                    fmt.Println(v2, v29)
                case "hello":
                    if len(v8) > 0 {
                        v8[0] = fi30
                    }
                default:
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                }
                var v33 bool = false
                _ = len(v8)
                _ = v33
            } else {
                if (-128 != 13) {
                    switch v29 {
                    case -530:
                        v3 = v3
                    case 443:
                        fmt.Println(v29, v2, fi30)
                    default:
                        _ = len(v8)
                    }
                    v34 := GenS3{F0: -300, F1: 2147483647, F2: -114}
                    for fi35 := 0; fi35 < 16; fi35++ {
                        var v36 uint16 = v1
                        v3 -= v3
                        fmt.Println(fi30, v1)
                        _ = fi35
                        _ = v36
                    }
                    _ = v34
                }
                switch -771 {
                case 75:
                    switch (fi30 >> ((18446744073709551615) & 63)) {
                    case 127:
                        v3 -= v3
                    case 23:
                        v37 := GenS5{F0: "hello", F1: true}
                        _ = v37
                    default:
                        switch int(v1) {
                        case 4294967295:
                            for fi38 := 0; fi38 < 1; fi38++ {
                                var v39 int16 = -581
                                var v40 int16 = v39
                                _ = fi38
                                _ = v39
                                _ = v40
                            }
                        case -585:
                            if true {
                                v41 := GenS6{F0: 811, F1: v2, F2: "a"}
                                v42 := GenS2{F0: 2, F1: v2, F2: v1}
                                var v43 int64 = -1
                                _ = v41
                                _ = v42
                                _ = v43
                            } else {
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                            }
                        case 331:
                            var v44 uint = 460
                            _ = v44
                        }
                    }
                default:
                    if true {
                        var v45 string = "a"
                        var v46 uint = 755
                        var v47 string = "foo bar"
                        _ = v45
                        _ = v46
                        _ = v47
                    } else {
                        switch -9223372036854775808 {
                        case 199:
                            for fi48 := 0; fi48 < 18; fi48++ {
                                v29 -= 4294967295
                                _ = fi48
                            }
                        case 32768:
                            fmt.Println(v2, v2, v29)
                        }
                        if true {
                            var v49 string = "\n\t"
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            v8 = append(v8, v29)
                            _ = v49
                        } else {
                            v50 := GenS5{F0: ("a" + ("" + (("hello" + "a") + ("a" + ("a" + ("\n\t" + ("a" + ("\n\t" + "foo bar")))))))), F1: (v3 != v3)}
                            _ = v50
                        }
                    }
                }
            }
            v51 := GenS7{F0: v2, F1: 255, F2: false}
            v3 = 42
            _ = v51
        }
        if false {
            fmt.Println(v3, v1)
        } else {
            v52 := GenS8{F0: 1, F1: (((((((("" + "") + ("a" + "foo bar")) + ("hello" + (("hello" + "a") + "a"))) + (("" + "\n\t") + "")) + "a") + "") + "\n\t") + "foo bar"), F2: 1}
            v29 += (-(fi30))
            _ = v52
        }
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v8
    _ = v9
}
