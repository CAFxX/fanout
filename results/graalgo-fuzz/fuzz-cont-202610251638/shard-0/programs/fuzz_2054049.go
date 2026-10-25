package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint16; F2 int16 }
type GenS3 struct { F0 int8; F1 int32; F2 uint8 }
type GenS4 struct { F0 int64; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool; F1 uint16 }
func genFn0() int8 {
    var v5 float32 = 0.0
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 int8 = 1
        _ = fi6
        _ = v7
    }
    if (42 == 719) {
        var v8 uint32 = 2
        _ = v8
    }
    _ = v5
    return 100
}

func main() {
    var v1 uint8 = 114
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]bool)
    v4 := []string{("hello" + (((("a" + ("a" + "a")) + "\n\t") + "hello") + "hello")), "foo bar", "foo bar"}
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 float32 = 0.1
        _ = fi9
        _ = v10
    }
    var v11 uint = 1
    var v12 int32 = -1
    if (true || false) {
        for fi13 := 0; fi13 < 8; fi13++ {
            v11 += 255
            var v14 int = 32767
            _ = fi13
            _ = v14
        }
        fmt.Println(v12, v1)
        if false {
            v11 -= 18446744073709551615
        } else {
            var v15 uint16 = 1
            var v16 float32 = (0.1 * ((0.0 / ((3.14159 / 3.14159) / 1.0)) - float32(3.14159)))
            v4 = append(v4, "")
            _ = v15
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 12; fi17++ {
            if (!false) {
                fmt.Println(v11)
            }
            if (440 == 372) {
                switch (fi17 & fi17) {
                case 32768:
                    v12 -= v12
                case 127:
                    v1 = (v1 >> ((v11) & 63))
                default:
                    v18 := GenS1{F0: (^(v1))}
                    _ = v18
                }
                if true {
                    v1 = (^(v1))
                } else {
                    switch (fi17 - fi17) {
                    case -902:
                        fmt.Println(fi17)
                    case 9223372036854775807:
                        for fi19 := 0; fi19 < 2; fi19++ {
                            switch "a" {
                            case "":
                                if false {
                                    delete(v3, 65535)
                                    if (-2.25 != (3.14159 / 1.0)) {
                                        v20 := GenS2{F0: false, F1: 326, F2: -32768}
                                        _ = v20
                                    }
                                    switch (^(fi17)) {
                                    case -9223372036854775808:
                                        fmt.Println(fi17)
                                    case 32767:
                                        if true {
                                            v21 := v2[2]
                                            fmt.Println(v21)
                                            if true {
                                                var v22 uint64 = 292
                                                var v23 string = (v21 + (((v21 + (("" + v21) + "foo bar")) + ("foo bar" + v21)) + ((v21 + v21) + v21)))
                                                v24 := GenS3{F0: -98, F1: -1, F2: 2}
                                                _ = v22
                                                _ = v23
                                                _ = v24
                                            }
                                            fmt.Println(fi17, v21, v11)
                                            _ = v21
                                        } else {
                                            if (42 >= 1) {
                                                v11 -= (v11 + 393)
                                                v25 := GenS4{F0: -9223372036854775808, F1: (^(v1)), F2: 800}
                                                v12 = (v12 + (v12 - v12))
                                                _ = v25
                                            } else {
                                                if (!(9223372036854775807 == 66)) {
                                                    var v26 float32 = float32(3.14159)
                                                    if ((-32768 > 1) || true) {
                                                        v4 = append(v4, "")
                                                    } else {
                                                        v12 = v12
                                                        var v27 uint32 = 2147483647
                                                        _ = v27
                                                    }
                                                    fmt.Println(v12)
                                                    _ = v26
                                                }
                                                v12 += (v12 >> ((v11) & 63))
                                                if ((!false) || false) {
                                                    switch fi17 {
                                                    case 9223372036854775807:
                                                        fmt.Println(v11, fi17)
                                                    case -83:
                                                        v12 = (-(v12))
                                                    default:
                                                        if true {
                                                            switch fi19 {
                                                            case 1:
                                                                var v28 bool = ("" != "a")
                                                                _ = v28
                                                            default:
                                                                v11 = v11
                                                            }
                                                            v29 := GenS5{F0: 42}
                                                            var v30 uint64 = 0
                                                            v30 = uint64(v1)
                                                            _ = v29
                                                            _ = v30
                                                        }
                                                    }
                                                    v31 := GenS2{F0: false, F1: 65535, F2: -1}
                                                    _ = v31
                                                }
                                            }
                                        }
                                    case -335:
                                        fmt.Println(fi19)
                                    default:
                                        v1 += v1
                                    }
                                } else {
                                    v32 := GenS6{F0: 880}
                                    if len(v4) > 0 {
                                        v4[0] = (("" + "\n\t") + "\n\t")
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = "foo bar"
                                    }
                                    _ = v32
                                }
                            }
                            switch (("" + "a") + "\n\t") {
                            case "":
                                if len(v4) > 0 {
                                    v4[0] = (("a" + "hello") + ("a" + ("hello" + ("foo bar" + "hello"))))
                                }
                            case "\n\t":
                                var v33 int = fi19
                                _ = v33
                            default:
                                _ = len(v2)
                            }
                            _ = fi19
                        }
                    }
                }
                v1 += v1
            } else {
                if true {
                    switch "\n\t" {
                    case "foo bar":
                        for fi34 := 0; fi34 < 4; fi34++ {
                            switch fi34 {
                            case -530:
                                switch ("" + "foo bar") {
                                case "hello":
                                    v1 += (v1 | 1)
                                case "foo bar":
                                    var v35 bool = false
                                    _ = v35
                                default:
                                    v1 += 77
                                }
                            case 4294967295:
                                delete(v3, 2)
                            default:
                                delete(v3, 665)
                            }
                            fmt.Println(v1, v12, fi17)
                            _ = fi34
                        }
                    default:
                        switch (^(fi17)) {
                        case 4294967295:
                            var v36 uint = 65535
                            _ = v36
                        }
                    }
                }
                if false {
                    var v37 bool = (!(!((v12 + (-(v12))) != (v12 & (v12 % (((v12 | (v12 / ((2147483647) | 1)))) | 1))))))
                    for fi38 := 0; fi38 < 18; fi38++ {
                        var v39 uint64 = 42
                        switch fi17 {
                        case 128:
                            if v37 {
                                if (int16(v11) < 0) {
                                    if false {
                                        v11 += uint(v1)
                                        var v40 uint8 = 177
                                        _ = v40
                                    } else {
                                        if len(v4) > 0 {
                                            v4[0] = ("" + "hello")
                                        }
                                        var v41 uint = v11
                                        v12 -= (-(v12))
                                        _ = v41
                                    }
                                    var v42 uint16 = 2
                                    _ = v42
                                }
                                var v43 uint = (^(v11))
                                if true {
                                    var v44 int16 = 1000
                                    _ = v44
                                } else {
                                    _ = len(v4)
                                    delete(v3, 602)
                                    if len(v4) > 0 {
                                        v4[0] = "hello"
                                    }
                                }
                                _ = v43
                            } else {
                                fmt.Println(v1, v12, v1)
                                v45 := GenS3{F0: genFn0(), F1: v12, F2: v1}
                                _ = v45
                            }
                        default:
                            var v46 int8 = -1
                            _ = v46
                        }
                        _ = fi38
                        _ = v39
                    }
                    _ = v37
                }
            }
            v12 = 42
            _ = fi17
        }
    }
    var v47 uint32 = 2147483647
    if false {
        if true {
            if true {
                fmt.Println(v12)
                _ = len(v3)
            } else {
                v48 := v2[v47]
                v49 := GenS6{F0: v47}
                _ = v48
                _ = v49
            }
            var v50 uint16 = 65535
            v47 += v47
            _ = v50
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        delete(v3, v47)
        v4 = append(v4, "a")
    }
    for fi51 := 0; fi51 < 1; fi51++ {
        fmt.Println(v1)
        if true {
            var v52 bool = true
            for fi53 := 0; fi53 < 16; fi53++ {
                var v54 uint16 = 255
                _ = fi53
                _ = v54
            }
            _ = v52
        }
        _ = fi51
    }
    v1 += v1
    v55 := GenS7{F0: false, F1: 2}
    var v56 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v55
    _ = v56
}
