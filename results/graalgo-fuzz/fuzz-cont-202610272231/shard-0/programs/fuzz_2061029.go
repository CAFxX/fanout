package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int8; F1 string; F2 int32 }
type GenS5 struct { F0 uint32; F1 bool; F2 string }
type GenS6 struct { F0 int64; F1 uint8 }

func main() {
    v1 := []string{"", "foo bar", "\n\t", ("\n\t" + "hello")}
    var v2 int16 = 1
    v3 := make(map[int]uint32)
    switch "\n\t" {
    case "a":
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v2)
        case "":
            fmt.Println(v2)
        default:
            var v4 uint16 = 2
            _ = v4
        }
    }
    fmt.Println(v2)
    v2 += 1
    v2 -= 32767
    if (!(0 > 1)) {
        v5 := GenS1{F0: 128, F1: 32767}
        v2 = v2
        var v6 uint = 65535
        _ = v5
        _ = v6
    } else {
        fmt.Println(v2)
    }
    switch -9223372036854775808 {
    case -2147483648:
        fmt.Println(v2)
    case 2147483647:
        v7 := v3[9223372036854775807]
        _ = v7
    default:
        fmt.Println(v2)
    }
    v8 := v3[32767]
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := GenS1{F0: 1, F1: 32768}
        var v11 int32 = -870
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        switch "hello" {
        case "a":
            v13 := GenS2{F0: -1, F1: 255}
            _ = v13
        case "hello":
            if (4294967295 < (^(v8))) {
                _ = len(v1)
            } else {
                var v14 uint64 = 264
                if (107 <= 127) {
                    var v15 uint16 = 65535
                    switch "foo bar" {
                    case "":
                        v1 = append(v1, "hello")
                    case "a":
                        for fi16 := 0; fi16 < 18; fi16++ {
                            if true {
                                switch fi12 {
                                case -2147483648:
                                    if true {
                                        if false {
                                            switch "hello" {
                                            case "a":
                                                switch -2147483648 {
                                                case -1:
                                                    if len(v1) > 0 {
                                                        v1[0] = ""
                                                    }
                                                default:
                                                    switch ("\n\t" + "\n\t") {
                                                    case "hello":
                                                        v8 -= (v8 * 4294967295)
                                                    case "":
                                                        var v17 uint = 2
                                                        _ = v17
                                                    default:
                                                        _ = len(v3)
                                                    }
                                                }
                                            case "":
                                                var v18 uint = 731
                                                v14 = uint64(v8)
                                                _ = v18
                                            }
                                            v19 := GenS3{F0: true, F1: 4294967295}
                                            _ = v19
                                        }
                                    } else {
                                        fmt.Println(v14, v2)
                                    }
                                case -328:
                                    v20 := GenS1{F0: 42, F1: (fi12 * -9223372036854775808)}
                                    var v21 string = ("foo bar" + "foo bar")
                                    _ = v20
                                    _ = v21
                                default:
                                    var v22 int = fi16
                                    _ = v22
                                }
                                v23 := GenS4{F0: 1, F1: "", F2: 2147483647}
                                _ = v23
                            }
                            if len(v1) > 0 {
                                v1[0] = ""
                            }
                            _ = fi16
                        }
                    default:
                        v14 = 18446744073709551615
                    }
                    v15 = 273
                    _ = v15
                }
                _ = v14
            }
        case "":
            switch fi12 {
            case 1:
                switch (^(fi12)) {
                case 21:
                    v24 := GenS3{F0: true, F1: 1}
                    _ = v24
                case 0:
                    for fi25 := 0; fi25 < 20; fi25++ {
                        v26 := GenS1{F0: 1, F1: (fi12 + (-(fi12)))}
                        if false {
                            if (!(fi12 < fi25)) {
                                v8 = 668
                                v2 += (v2 % ((-261) | 1))
                            }
                        }
                        _ = fi25
                        _ = v26
                    }
                case -301:
                    v8 = v8
                }
            default:
                var v27 uint32 = (v8 * (v8 + (v8 * (v8 + v8))))
                _ = v27
            }
        default:
            switch ("hello" + "foo bar") {
            case "\n\t":
                if len(v1) > 0 {
                    v1[0] = "a"
                }
            case "hello":
                fmt.Println(v2)
            case "":
                if true {
                    v28 := GenS3{F0: true, F1: (v8 - (v8 + v8))}
                    _ = v28
                }
            }
        }
        _ = fi12
    }
    v2 = (^(v2))
    if (false || (42 == 1)) {
        for fi29 := 0; fi29 < 5; fi29++ {
            v2 += v2
            v2 -= (v2 % ((-32768) | 1))
            _ = fi29
        }
        if (283 != (^(v8))) {
            var v30 uint32 = (v8 & (^(v8)))
            v30 -= (v8 - (v30 & v8))
            _ = v30
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        v2 = v2
        var v31 int = 32767
        _ = v31
    }
    v1 = append(v1, ("foo bar" + "foo bar"))
    switch 42 {
    case -46:
        if false {
            for fi32 := 0; fi32 < 15; fi32++ {
                if true {
                    fmt.Println(v2, v2)
                } else {
                    v33 := GenS5{F0: v8, F1: ((42 < 735) || (3.14159 < 0.0)), F2: ""}
                    v34 := v3[fi32]
                    delete(v3, -831)
                    _ = v33
                    _ = v34
                }
                v8 -= (v8 & v8)
                _ = fi32
            }
            var v35 int = 32768
            var v36 string = ("hello" + "foo bar")
            _ = v35
            _ = v36
        }
    case 1:
        v2 = -32768
    default:
        v37 := GenS6{F0: -929, F1: 127}
        _ = v37
    }
    var v38 bool = (242 > 161)
    if v38 {
        var v39 int64 = -1
        _ = v39
    } else {
        switch "" {
        case "foo bar":
            fmt.Println(v2)
        case "hello":
            var v40 uint32 = 234
            _ = v40
        case "\n\t":
            for fi41 := 0; fi41 < 9; fi41++ {
                if (!v38) {
                    var v42 uint64 = 4294967295
                    v2 -= (-(v2))
                    _ = v42
                }
                fmt.Println(fi41, v2)
                switch "a" {
                case "\n\t":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    if (255 > 42) {
                        var v43 float32 = 1e10
                        _ = v43
                    }
                }
                _ = fi41
            }
        }
    }
    v38 = (!false)
    delete(v3, 0)
    v38 = (42 >= 0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v8
}
