package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 int32; F2 bool }

func main() {
    var v1 int32 = -452
    var v2 float32 = float32(float64(0.0))
    v3 := []uint32{2147483647, 2147483648, 65535, 415}
    var v4 int16 = -879
    switch 393 {
    case 128:
        _ = len(v3)
    case 2147483647:
        for fi5 := 0; fi5 < 16; fi5++ {
            for fi6 := 0; fi6 < 8; fi6++ {
                switch fi5 {
                case 127:
                    fmt.Println(v4)
                case 2147483647:
                    if true {
                        var v7 uint = 2
                        var v8 uint64 = 0
                        _ = v7
                        _ = v8
                    } else {
                        if true {
                            switch "\n\t" {
                            case "hello":
                                if (false || false) {
                                    switch "hello" {
                                    case "foo bar":
                                        var v9 int8 = 1
                                        _ = v9
                                    default:
                                        fmt.Println(fi5, v2, v4)
                                    }
                                    fmt.Println(v2, v2)
                                    if (true || false) {
                                        if (!(((9223372036854775807 < 42) || false) || true)) {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v1 = v1
                                        } else {
                                            v1 = 1
                                            switch "" {
                                            case "a":
                                                fmt.Println(fi5, fi5, v4)
                                            }
                                        }
                                    } else {
                                        if true {
                                            v10 := GenS1{F0: 1}
                                            _ = v10
                                        }
                                        fmt.Println(v2, fi6, v4)
                                        if false {
                                            v1 = (v1 | v1)
                                        } else {
                                            v2 = v2
                                            if (-2.25 >= -2.25) {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v11 string = "foo bar"
                                                _ = v11
                                            }
                                            v12 := GenS1{F0: 42}
                                            _ = v12
                                        }
                                    }
                                }
                            case "foo bar":
                                v13 := GenS1{F0: 0}
                                _ = v13
                            default:
                                var v14 uint = 1
                                _ = v14
                            }
                            if false {
                                var v15 uint8 = 0
                                v16 := GenS2{F0: 127, F1: (v4 - (^(v4)))}
                                var v17 uint = 627
                                var v18 int16 = v4
                                _ = v15
                                _ = v16
                                _ = v17
                                _ = v18
                            } else {
                                if ((42 > 42) && (2 > 280)) {
                                    v19 := GenS2{F0: int8(v4), F1: v4}
                                    if (1 < 1) {
                                        v2 -= (0.1 / -2.25)
                                        if true {
                                            switch 604 {
                                            case -2147483648:
                                                if true {
                                                    v2 -= float32(0.1)
                                                    v20 := GenS1{F0: 59}
                                                    _ = v20
                                                } else {
                                                    v4 += 649
                                                }
                                            case 143:
                                                switch 127 {
                                                case 9223372036854775807:
                                                    v21 := GenS3{F0: 0, F1: "foo bar"}
                                                    _ = v21
                                                default:
                                                    v1 += (v1 & (v1 + (v1 - (v1 / (((v1 % (((v1 ^ -548)) | 1))) | 1)))))
                                                }
                                            }
                                            var v22 int32 = (v1 + 2147483647)
                                            var v23 int = fi6
                                            v24 := GenS3{F0: 0, F1: (("hello" + (("\n\t" + ("foo bar" + ("foo bar" + "foo bar"))) + "")) + ("a" + ((("\n\t" + ("hello" + ("" + "\n\t"))) + ("\n\t" + "hello")) + "\n\t")))}
                                            _ = len(v3)
                                            _ = v22
                                            _ = v23
                                            _ = v24
                                        }
                                    } else {
                                        v4 += v4
                                        v25 := GenS1{F0: 42}
                                        var v26 float32 = ((((1e10 - v2) * 1.5) + float32((0.1 / 1.5))) - v2)
                                        _ = v25
                                        _ = v26
                                    }
                                    var v27 uint64 = 2
                                    if (!((!false) || true)) {
                                        var v28 uint = 4294967295
                                        _ = len(v3)
                                        _ = v28
                                    }
                                    var v29 uint = 1
                                    v30 := GenS4{F0: 65535}
                                    _ = v19
                                    _ = v27
                                    _ = v29
                                    _ = v30
                                } else {
                                    var v31 int32 = (v1 << ((18446744073709551615) & 63))
                                    switch fi5 {
                                    case 466:
                                        if (978 == 0) {
                                            switch "foo bar" {
                                            case "hello":
                                                v32 := GenS4{F0: 2}
                                                _ = v32
                                            default:
                                                v2 += ((v2 - v2) * (v2 + (v2 - v2)))
                                            }
                                            var v33 uint8 = 1
                                            v31 = v1
                                            _ = v33
                                        }
                                    case -9223372036854775808:
                                        var v34 uint16 = 32767
                                        _ = v34
                                    case 9223372036854775807:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    _ = len(v3)
                                    _ = v31
                                }
                                v35 := GenS5{F0: 2, F1: -1, F2: true}
                                _ = v35
                            }
                            v1 = (v1 - v1)
                        } else {
                            var v36 uint16 = 65535
                            switch 1 {
                            case -9223372036854775808:
                                fmt.Println(fi6, fi6, v36)
                            case 128:
                                if (!true) {
                                    if false {
                                        v37 := GenS3{F0: 659, F1: "foo bar"}
                                        _ = v37
                                    } else {
                                        var v38 string = (("" + "hello") + "foo bar")
                                        v38 = v38
                                        var v39 uint8 = 42
                                        _ = v38
                                        _ = v39
                                    }
                                    if ((^(v36)) >= 1) {
                                        switch -1 {
                                        case -1:
                                            if ("\n\t" != ("foo bar" + "foo bar")) {
                                                var v40 uint32 = 708
                                                _ = v40
                                            } else {
                                                switch fi5 {
                                                case 1:
                                                    if (953 == 1) {
                                                        if ((v4 >> ((4294967295) & 63)) > v4) {
                                                            v4 = -167
                                                        }
                                                        if len(v3) > 0 {
                                                            v3[0] = 42
                                                        }
                                                        var v41 uint = 18446744073709551615
                                                        _ = v41
                                                    }
                                                default:
                                                    var v42 int16 = (v4 << ((622) & 63))
                                                    _ = v42
                                                }
                                            }
                                        default:
                                            if ((42 > 2) && true) {
                                                fmt.Println(v4, v4, v2)
                                            }
                                        }
                                        v1 = (v1 & (v1 & (v1 << ((4294967295) & 63))))
                                    } else {
                                        fmt.Println(v2)
                                        var v43 int16 = int16(v1)
                                        _ = v43
                                    }
                                    var v44 bool = (127 > 110)
                                    _ = v44
                                } else {
                                    fmt.Println(fi6, v1, v4)
                                }
                            default:
                                v4 = v4
                            }
                            _ = v36
                        }
                        var v45 uint = 0
                        switch "foo bar" {
                        case "\n\t":
                            v2 = (v2 - v2)
                        case "hello":
                            if (1.5 != (3.14159 - (0.0 / 3.14159))) {
                                fmt.Println(fi6, v45, fi5)
                            } else {
                                _ = len(v3)
                                var v46 uint8 = 42
                                _ = v46
                            }
                        }
                        v3 = append(v3, 117)
                        _ = v45
                    }
                case 32767:
                    v1 = (v1 % ((v1) | 1))
                }
                _ = fi6
            }
            _ = fi5
        }
    case 0:
        v4 -= -1
    }
    v4 = v4
    var v47 int32 = -2147483648
    v48 := GenS4{F0: 4294967295}
    fmt.Println(v1, v1)
    var v49 uint = 255
    var v50 int16 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
