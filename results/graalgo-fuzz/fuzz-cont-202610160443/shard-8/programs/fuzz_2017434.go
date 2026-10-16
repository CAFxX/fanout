package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 bool; F2 int16 }
type GenS5 struct { F0 uint; F1 uint32; F2 int16 }
type GenS6 struct { F0 int16 }

func main() {
    var v1 string = ("a" + "foo bar")
    var v2 uint16 = 255
    v3 := []uint32{uint32(v2), 812, 65535, 2147483648}
    v2 = v2
    if true {
        v2 = 586
    }
    if (989 < 996) {
        var v4 int = 2147483647
        _ = v4
    }
    switch -389 {
    case -9223372036854775808:
        switch 852 {
        case 127:
            var v5 bool = false
            _ = v5
        case 32768:
            var v6 uint16 = 605
            _ = v6
        case 128:
            if ("\n\t" == v1) {
                v3 = append(v3, 42)
                for fi7 := 0; fi7 < 20; fi7++ {
                    _ = len(v3)
                    v1 = (v1 + v1)
                    var v8 int = fi7
                    _ = fi7
                    _ = v8
                }
                if false {
                    v3 = append(v3, 2147483647)
                    var v9 int = 32767
                    v2 = 1
                    _ = v9
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
        default:
            v10 := GenS1{F0: 2}
            _ = v10
        }
    case 4294967295:
        if true {
            v1 = v1
            if true {
                switch 127 {
                case 42:
                    v11 := GenS2{F0: -32, F1: 1, F2: 0}
                    _ = v11
                default:
                    if (110 != 176) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        if (9 == 1) {
                            v12 := GenS3{F0: 952}
                            for fi13 := 0; fi13 < 10; fi13++ {
                                var v14 string = "foo bar"
                                for fi15 := 0; fi15 < 20; fi15++ {
                                    var v16 float64 = ((float64(0.0) - 0.0) + -2.25)
                                    switch "hello" {
                                    case "foo bar":
                                        var v17 string = ("\n\t" + v14)
                                        _ = v17
                                    case "":
                                        var v18 string = (v1 + v1)
                                        _ = v18
                                    }
                                    _ = fi15
                                    _ = v16
                                }
                                var v19 int8 = -1
                                _ = fi13
                                _ = v14
                                _ = v19
                            }
                            switch 1 {
                            case -930:
                                for fi20 := 0; fi20 < 11; fi20++ {
                                    for fi21 := 0; fi21 < 10; fi21++ {
                                        if len(v3) > 0 {
                                            v3[0] = uint32(v2)
                                        }
                                        var v22 int32 = 2
                                        var v23 uint16 = v2
                                        if true {
                                            var v24 float64 = (-2.25 * 0.0)
                                            _ = v24
                                        } else {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            if true {
                                                if false {
                                                    v25 := GenS2{F0: -1, F1: 42, F2: 42}
                                                    _ = v25
                                                } else {
                                                    if true {
                                                        if (!true) {
                                                            if true {
                                                                if true {
                                                                    switch (v1 + v1) {
                                                                    case "hello":
                                                                        v26 := GenS3{F0: 0}
                                                                        _ = v26
                                                                    case "\n\t":
                                                                        _ = len(v3)
                                                                    case "a":
                                                                        var v27 string = v1
                                                                        _ = v27
                                                                    default:
                                                                        if true {
                                                                            _ = len(v3)
                                                                        } else {
                                                                            switch "\n\t" {
                                                                            case "a":
                                                                                var v28 bool = (v23 >= 2)
                                                                                _ = v28
                                                                            case "foo bar":
                                                                                v29 := GenS2{F0: 127, F1: v22, F2: 42}
                                                                                _ = v29
                                                                            default:
                                                                                var v30 uint64 = 42
                                                                                _ = v30
                                                                            }
                                                                            fmt.Println(v2, fi21)
                                                                            var v31 uint = 71
                                                                            _ = v31
                                                                        }
                                                                    }
                                                                    if false {
                                                                        v32 := GenS4{F0: v1, F1: true, F2: 42}
                                                                        var v33 uint8 = 1
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        v33 += (v33 >> ((255) & 63))
                                                                        switch v1 {
                                                                        case "":
                                                                            var v34 int32 = (v22 >> ((4294967295) & 63))
                                                                            _ = v34
                                                                        case "hello":
                                                                            switch "\n\t" {
                                                                            case "foo bar":
                                                                                var v35 float32 = -2.25
                                                                                _ = v35
                                                                            case "a":
                                                                                v36 := GenS2{F0: -128, F1: (v22 / ((42) | 1)), F2: v33}
                                                                                _ = v36
                                                                            case "\n\t":
                                                                                if (-2147483648 >= v22) {
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                    v33 += 255
                                                                                }
                                                                            }
                                                                        default:
                                                                            if (v1 > (v1 + (v1 + v1))) {
                                                                                v37 := GenS1{F0: 2}
                                                                                var v38 int8 = -44
                                                                                _ = v37
                                                                                _ = v38
                                                                            } else {
                                                                                if true {
                                                                                    v39 := GenS1{F0: 1}
                                                                                    _ = v39
                                                                                } else {
                                                                                    var v40 uint = 0
                                                                                    var v41 int = (fi21 | 0)
                                                                                    _ = v40
                                                                                    _ = v41
                                                                                }
                                                                                v33 += v33
                                                                                v23 += (v2 << ((3) & 63))
                                                                            }
                                                                        }
                                                                        _ = v32
                                                                        _ = v33
                                                                    }
                                                                } else {
                                                                    var v42 uint8 = 1
                                                                    _ = v42
                                                                }
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            } else {
                                                                if (true && false) {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 2147483647
                                                                    }
                                                                }
                                                                var v43 int64 = 1
                                                                var v44 float64 = 1.5
                                                                _ = v43
                                                                _ = v44
                                                            }
                                                            var v45 uint32 = 701
                                                            _ = v45
                                                        } else {
                                                            if true {
                                                                v1 = v1
                                                                switch "a" {
                                                                case "a":
                                                                    v46 := GenS2{F0: -128, F1: v22, F2: 255}
                                                                    _ = v46
                                                                }
                                                                v3 = append(v3, 145)
                                                                if ((!((false || false) || false)) && false) {
                                                                    v3 = append(v3, 65535)
                                                                    v47 := GenS5{F0: 468, F1: 0, F2: 1000}
                                                                    fmt.Println(v2, fi21)
                                                                    _ = v47
                                                                } else {
                                                                    var v48 uint16 = (^(v23))
                                                                    v1 = v1
                                                                    v48 = (v23 >> ((539) & 63))
                                                                    _ = v48
                                                                }
                                                            }
                                                            if len(v3) > 0 {
                                                                v3[0] = 0
                                                            }
                                                        }
                                                        var v49 uint8 = uint8(v2)
                                                        _ = v49
                                                    }
                                                    var v50 int16 = 32767
                                                    _ = v50
                                                }
                                                v51 := GenS4{F0: v1, F1: true, F2: 320}
                                                var v52 int32 = v22
                                                _ = v51
                                                _ = v52
                                            }
                                            _ = len(v3)
                                        }
                                        _ = fi21
                                        _ = v22
                                        _ = v23
                                    }
                                    v53 := GenS6{F0: 1000}
                                    _ = fi20
                                    _ = v53
                                }
                            case 42:
                                if len(v3) > 0 {
                                    v3[0] = 2147483648
                                }
                            case -195:
                                var v54 uint64 = 42
                                _ = v54
                            }
                            _ = v12
                        }
                        var v55 int = -9223372036854775808
                        v55 -= (^(v55))
                        _ = v55
                    }
                }
                switch ((v1 + v1) + "\n\t") {
                case "":
                    for fi56 := 0; fi56 < 20; fi56++ {
                        if (2147483647 == 775) {
                            v57 := GenS3{F0: 76}
                            _ = v57
                        } else {
                            switch (fi56 ^ (fi56 << ((42) & 63))) {
                            case 4294967295:
                                v3 = append(v3, 772)
                            default:
                                fmt.Println(fi56, fi56)
                            }
                            for fi58 := 0; fi58 < 15; fi58++ {
                                var v59 bool = false
                                if v59 {
                                    var v60 int = 42
                                    _ = len(v3)
                                    v1 = "hello"
                                    _ = v60
                                }
                                _ = fi58
                                _ = v59
                            }
                        }
                        v1 = ("foo bar" + (v1 + ("" + v1)))
                        _ = fi56
                    }
                }
                var v61 int = -666
                _ = v61
            }
        }
    default:
        var v62 uint = 255
        _ = v62
    }
    v63 := GenS4{F0: v1, F1: ((v1 + "") == v1), F2: 1000}
    v1 = (v1 + v1)
    fmt.Println(v2)
    if len(v3) > 0 {
        v3[0] = 2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v63
}
