package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS4 struct { F0 uint16; F1 int8; F2 string }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 float32 = 0.1
    v2 := []int{}
    var v3 float32 = v1
    var v4 int32 = 42
    v5 := make(map[int]bool)
    var v6 string = ("" + "\n\t")
    var v7 uint32 = 2
    v8 := GenS1{F0: uint8(v7), F1: 86, F2: 255}
    var v9 uint = 4294967295
    v9 += (v9 - 42)
    v9 = (^(v9))
    v1 = v3
    fmt.Println(v1)
    var v10 bool = true
    for fi11 := 0; fi11 < 20; fi11++ {
        if (249 >= 42) {
            if v10 {
                fmt.Println(v9, v3, v1)
            }
            v1 = (v3 + v1)
            var v12 uint8 = 0
            _ = v12
        } else {
            if (((false || v10) || v10) || false) {
                var v13 uint8 = 2
                _ = v13
            }
            for fi14 := 0; fi14 < 2; fi14++ {
                fmt.Println(fi14, v6)
                _ = fi14
            }
            fmt.Println(v9, v10, v3)
        }
        v15 := v5[fi11]
        _ = fi11
        _ = v15
    }
    v6 = v6
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v4, v4)
    _ = len(v5)
    switch -9223372036854775808 {
    case 32768:
        switch 524 {
        case 1:
            v5[935] = (!v10)
        case 127:
            switch -9223372036854775808 {
            case -1:
                var v16 int64 = 942
                _ = v16
            case 32768:
                switch 314 {
                case 4294967295:
                    v4 -= v4
                }
            }
        case 389:
            switch "hello" {
            case "foo bar":
                var v17 int = -242
                _ = v17
            case "hello":
                fmt.Println(v1, v6)
            case "":
                if ((float64((v3 - float32((1e10 / 1.5)))) + 1.5) != (-2.25 + -2.25)) {
                    var v18 uint8 = 255
                    v9 += (v9 - (^(v9)))
                    _ = v18
                }
            default:
                switch (v6 + v6) {
                case "":
                    v19 := GenS1{F0: 42, F1: -115, F2: 32767}
                    _ = v19
                }
            }
        default:
            v6 = ((((v6 + v6) + v6) + ("hello" + "hello")) + ("foo bar" + (v6 + (v6 + v6))))
        }
    case -1:
        var v20 uint32 = (v7 >> (((v9 | 618)) & 63))
        _ = v20
    }
    v6 = ""
    if (49 <= 2) {
        switch 193 {
        case 128:
            for fi21 := 0; fi21 < 12; fi21++ {
                switch v6 {
                case "a":
                    switch (fi21 | 127) {
                    case 127:
                        v2 = append(v2, (fi21 | fi21))
                    case -366:
                        var v22 bool = v10
                        _ = v22
                    default:
                        v3 -= (((3.14159 + v1) - (v3 - ((3.14159 - float32(1.5)) + ((1.5 + v3) + (v1 + v1))))) + v1)
                    }
                case "\n\t":
                    if (fi21 <= fi21) {
                        v23 := v5[1]
                        v23 = false
                        var v24 bool = ((^(v9)) == v9)
                        _ = v23
                        _ = v24
                    } else {
                        var v25 float32 = 0.0
                        v26 := GenS1{F0: 255, F1: 0, F2: 0}
                        _ = v25
                        _ = v26
                    }
                default:
                    v27 := GenS2{F0: ((v6 + v6) + v6), F1: 0}
                    _ = v27
                }
                v28 := v5[(fi21 | 42)]
                if v28 {
                    switch fi21 {
                    case 126:
                        if (!(!v10)) {
                            for fi29 := 0; fi29 < 6; fi29++ {
                                if true {
                                    v3 += float32(1e10)
                                } else {
                                    fmt.Println(fi29, v6, fi21)
                                    var v30 uint64 = 42
                                    if v10 {
                                        v28 = ((v1 - float32(1.5)) < 3.14159)
                                        v31 := GenS3{F0: 1, F1: 255, F2: 826}
                                        if ((-(v4)) <= (v4 % ((1) | 1))) {
                                            v3 += 1.5
                                            v32 := GenS4{F0: 255, F1: -49, F2: v6}
                                            v33 := GenS3{F0: 122, F1: 0, F2: uint16(v7)}
                                            if (!(663 != 363)) {
                                                switch ("a" + v6) {
                                                case "\n\t":
                                                    _ = len(v2)
                                                case "":
                                                    fmt.Println(fi29, v4)
                                                default:
                                                    fmt.Println(v6)
                                                }
                                            } else {
                                                if v28 {
                                                    var v34 uint32 = 2147483648
                                                    var v35 int16 = 32767
                                                    v36 := GenS3{F0: 116, F1: 742, F2: 42}
                                                    _ = v34
                                                    _ = v35
                                                    _ = v36
                                                } else {
                                                    if (!(true || (!(305 >= 0)))) {
                                                        _ = len(v5)
                                                        delete(v5, (^(fi29)))
                                                    } else {
                                                        if v28 {
                                                            _ = len(v2)
                                                        } else {
                                                            v37 := GenS5{F0: (^(v7))}
                                                            switch -9223372036854775808 {
                                                            case 759:
                                                                fmt.Println(v6, fi29)
                                                            default:
                                                                fmt.Println(v3)
                                                            }
                                                            switch 32768 {
                                                            case -9223372036854775808:
                                                                v38 := GenS2{F0: v6, F1: 0}
                                                                _ = v38
                                                            case -789:
                                                                fmt.Println(v7, v1, v4)
                                                            }
                                                            _ = v37
                                                        }
                                                        var v39 uint64 = 440
                                                        if (v9 >= (v9 % ((uint(v7)) | 1))) {
                                                            if true {
                                                                if v28 {
                                                                    v40 := GenS2{F0: v6, F1: 42}
                                                                    v7 -= (v7 | (v7 / ((uint32(v9)) | 1)))
                                                                    fmt.Println(fi29, v7)
                                                                    _ = v40
                                                                }
                                                            }
                                                            switch -674 {
                                                            case 32767:
                                                                var v41 bool = false
                                                                _ = v41
                                                            }
                                                        } else {
                                                            v3 = v3
                                                            v42 := GenS5{F0: v7}
                                                            _ = v42
                                                        }
                                                        v43 := v5[(fi21 >> ((v9) & 63))]
                                                        _ = v39
                                                        _ = v43
                                                    }
                                                    fmt.Println(fi21)
                                                }
                                                var v44 uint32 = v7
                                                v30 -= 2
                                                _ = v44
                                            }
                                            var v45 int = (fi29 + fi29)
                                            _ = v32
                                            _ = v33
                                            _ = v45
                                        }
                                        var v46 float32 = (v1 - (0.1 + (1e10 * v3)))
                                        _ = v31
                                        _ = v46
                                    } else {
                                        v2 = append(v2, (fi29 ^ (^(fi21))))
                                        v9 = (v9 + 392)
                                        fmt.Println(v28)
                                        if v10 {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v47 := GenS4{F0: 65535, F1: -1, F2: ("" + v6)}
                                            fmt.Println(fi21)
                                            _ = v47
                                        } else {
                                            v4 += v4
                                        }
                                    }
                                    _ = v30
                                }
                                _ = fi29
                            }
                            v3 = 1.5
                        }
                    case -9223372036854775808:
                        var v48 int8 = -44
                        _ = v48
                    }
                    for fi49 := 0; fi49 < 4; fi49++ {
                        var v50 float64 = ((3.14159 + 1.5) * 1e10)
                        var v51 string = (((v6 + v6) + "hello") + "")
                        v6 = v51
                        v10 = (!(!(false && v28)))
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                    v28 = ((!(161 < 1)) && true)
                }
                _ = fi21
                _ = v28
            }
        default:
            if (255 == 65535) {
                var v52 bool = (-548 == 0)
                var v53 uint16 = 255
                _ = v52
                _ = v53
            }
        }
        v6 = ("foo bar" + v6)
    } else {
        v7 += 2147483648
        v54 := GenS4{F0: 42, F1: -58, F2: (v6 + "hello")}
        _ = v54
    }
    var v55 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
