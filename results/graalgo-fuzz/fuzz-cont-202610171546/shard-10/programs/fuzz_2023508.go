package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 uint64; F2 string }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int8; F1 uint16; F2 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    v3 := []string{"", ("\n\t" + "\n\t"), "hello"}
    v4 := []uint32{}
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 int32 = -1
        _ = fi5
        _ = v6
    }
    var v7 float32 = 0.0
    if true {
        v7 = v7
        if (42 <= 1) {
            fmt.Println(v7)
            var v8 uint = 18446744073709551615
            _ = v8
        }
    } else {
        v7 += (v7 - v7)
        if true {
            switch -9223372036854775808 {
            case 752:
                for fi9 := 0; fi9 < 7; fi9++ {
                    var v10 uint16 = 65535
                    fmt.Println(fi9, fi9)
                    _ = fi9
                    _ = v10
                }
            case 942:
                v11 := GenS1{F0: 9223372036854775807}
                _ = v11
            }
            v7 += (v7 - 1e10)
            v7 += ((v7 - v7) / 3.14159)
        }
        var v12 uint16 = 32767
        _ = v12
    }
    var v13 int16 = -41
    switch -934 {
    case 128:
        if true {
            fmt.Println(v13, v7)
            var v14 int32 = -798
            if (false && false) {
                for fi15 := 0; fi15 < 2; fi15++ {
                    fmt.Println(v14)
                    v13 -= v13
                    _ = fi15
                }
            }
            _ = v14
        }
    case -913:
        v16 := GenS2{F0: v13, F1: false}
        _ = v16
    }
    var v17 int32 = 0
    var v18 int16 = (v13 + v13)
    for fi19 := 0; fi19 < 20; fi19++ {
        if true {
            v20 := GenS3{F0: 0}
            v17 = v17
            v21 := GenS2{F0: v13, F1: true}
            _ = v20
            _ = v21
        } else {
            var v22 uint16 = 2
            _ = v22
        }
        if (42 <= 2) {
            v23 := GenS4{F0: (-(fi19))}
            _ = v23
        }
        if (2 >= 2) {
            var v24 uint64 = 42
            switch ("" + ((("\n\t" + "hello") + ("foo bar" + ((("foo bar" + ("foo bar" + "foo bar")) + "foo bar") + "foo bar"))) + ((("" + "\n\t") + ("" + "a")) + "\n\t"))) {
            case "foo bar":
                var v25 int = fi19
                _ = v25
            case "a":
                fmt.Println(v18, v18)
            case "\n\t":
                for fi26 := 0; fi26 < 20; fi26++ {
                    v18 -= 42
                    _ = fi26
                }
            }
            v27 := GenS5{F0: true}
            _ = v24
            _ = v27
        }
        _ = fi19
    }
    fmt.Println(v13)
    var v28 uint8 = 128
    if (9223372036854775807 >= -9223372036854775808) {
        var v29 uint = 686
        if (2 > 490) {
            if (true || true) {
                v17 += (v17 / (((v17 - v17)) | 1))
                switch "foo bar" {
                case "":
                    v30 := v1[39]
                    _ = v30
                default:
                    if (!(-1 > 127)) {
                        var v31 uint64 = 0
                        _ = v31
                    } else {
                        if (0 >= v17) {
                            var v32 int16 = v13
                            switch (("a" + ("a" + "")) + ("" + "foo bar")) {
                            case "a":
                                fmt.Println(v28)
                            case "\n\t":
                                switch -550 {
                                case -57:
                                    switch 32768 {
                                    case 127:
                                        v33 := v1[0]
                                        _ = v33
                                    }
                                }
                            default:
                                v34 := GenS6{F0: ("a" + "\n\t"), F1: 18446744073709551615, F2: "a"}
                                _ = v34
                            }
                            _ = v32
                        } else {
                            _ = len(v3)
                        }
                        switch (("a" + ("hello" + (("\n\t" + ("hello" + ("hello" + ("hello" + "")))) + "foo bar"))) + "\n\t") {
                        case "\n\t":
                            v29 = v29
                        default:
                            for fi35 := 0; fi35 < 19; fi35++ {
                                v36 := GenS2{F0: (v18 - v18), F1: (true && true)}
                                var v37 float32 = (v7 + float32(float64(v7)))
                                for fi38 := 0; fi38 < 19; fi38++ {
                                    var v39 uint64 = 1
                                    v28 = (v28 % ((v28) | 1))
                                    var v40 uint32 = 65535
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    delete(v1, 0)
                                    _ = fi38
                                    _ = v39
                                    _ = v40
                                }
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                        }
                    }
                }
            } else {
                var v41 string = "\n\t"
                switch ("" + v41) {
                case "":
                    switch 32767 {
                    case -2147483648:
                        v42 := GenS7{F0: -492}
                        _ = v42
                    case 1:
                        switch ("" + v41) {
                        case "\n\t":
                            for fi43 := 0; fi43 < 7; fi43++ {
                                _ = len(v4)
                                _ = fi43
                            }
                        case "hello":
                            v41 = (v41 + (v41 + v41))
                        }
                    case -876:
                        v29 = v29
                    default:
                        var v44 float64 = ((1e10 * (((0.1 / ((((0.1 - 0.0) - (float64(((v7 * v7) * v7)) + 1.5)) - -2.25) + ((3.14159 * 1e10) - 3.14159))) * ((0.1 + -2.25) / 1.0)) - -2.25)) * 1e10)
                        _ = v44
                    }
                case "a":
                    var v45 float32 = float32(3.14159)
                    _ = v45
                default:
                    v17 = v17
                }
                switch -296 {
                case 512:
                    fmt.Println(v28)
                case 42:
                    if len(v3) > 0 {
                        v3[0] = ("hello" + "foo bar")
                    }
                default:
                    v7 -= (((v7 - v7) * -2.25) + (0.0 + (float32(0.1) / ((1.5 - 1e10) * 3.14159))))
                }
                _ = v41
            }
            v46 := GenS1{F0: 0}
            _ = v46
        }
        _ = v29
    } else {
        var v47 int8 = 42
        _ = v47
    }
    _ = len(v3)
    switch "\n\t" {
    case "":
        v48 := GenS6{F0: "a", F1: 2, F2: ""}
        _ = v48
    }
    switch 0 {
    case 168:
        _ = len(v2)
    }
    fmt.Println(v18, v13, v28)
    switch -409 {
    case 32767:
        var v49 float64 = 1e10
        _ = v49
    case -631:
        delete(v1, 1)
    default:
        v50 := GenS8{F0: 100, F1: 65535, F2: 4294967295}
        _ = v50
    }
    v7 = 3.14159
    var v51 float32 = v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v51:", v51)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v51
    _ = v7
}
