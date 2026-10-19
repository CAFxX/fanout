package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS5 struct { F0 uint16; F1 string }

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{}
    var v3 uint64 = 398
    var v4 int16 = -32768
    v5 := []string{"foo bar", ("foo bar" + (("foo bar" + "hello") + "a"))}
    v6 := GenS1{F0: 2, F1: -268}
    v7 := v1[2]
    var v8 int = 127
    v4 += 42
    v8 -= v8
    if true {
        for fi9 := 0; fi9 < 6; fi9++ {
            switch "" {
            case "":
                _ = len(v5)
            case "\n\t":
                var v10 int8 = -1
                _ = v10
            case "hello":
                var v11 uint8 = 128
                _ = v11
            }
            if (97 <= 100) {
                switch (-(v8)) {
                case 9223372036854775807:
                    v12 := GenS2{F0: 127}
                    _ = v12
                case 0:
                    switch "hello" {
                    case "foo bar":
                        for fi13 := 0; fi13 < 6; fi13++ {
                            v14 := GenS2{F0: 164}
                            var v15 uint32 = 2
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    default:
                        v3 = 42
                    }
                case -1:
                    for fi16 := 0; fi16 < 3; fi16++ {
                        switch fi9 {
                        case -2147483648:
                            v17 := v1[uint32(v3)]
                            _ = v17
                        case 0:
                            v18 := GenS1{F0: 1, F1: -2147483648}
                            _ = v18
                        case 557:
                            var v19 int64 = 0
                            _ = v19
                        default:
                            var v20 int16 = v4
                            _ = v20
                        }
                        v8 += (v8 | v8)
                        var v21 float64 = 0.1
                        _ = len(v1)
                        if false {
                            v3 += v3
                            v3 = v3
                            if true {
                                v8 -= 32767
                                v4 += v4
                                v3 = 1
                            }
                        }
                        _ = fi16
                        _ = v21
                    }
                }
                for fi22 := 0; fi22 < 1; fi22++ {
                    if false {
                        var v23 int64 = 9223372036854775807
                        var v24 int32 = 42
                        _ = v23
                        _ = v24
                    }
                    _ = fi22
                }
                if len(v5) > 0 {
                    v5[0] = "hello"
                }
            }
            _ = fi9
        }
        var v25 uint16 = 122
        if true {
            fmt.Println(v3, v7)
            v4 = -836
            if false {
                v7 = v7
                var v26 uint8 = 0
                if len(v2) > 0 {
                    v2[0] = 2
                }
                _ = v26
            } else {
                for fi27 := 0; fi27 < 15; fi27++ {
                    v3 += (^(v3))
                    var v28 uint = 805
                    v8 = fi27
                    _ = fi27
                    _ = v28
                }
                var v29 int32 = 834
                _ = v29
            }
        } else {
            var v30 uint32 = 572
            _ = v30
        }
        _ = v25
    } else {
        for fi31 := 0; fi31 < 19; fi31++ {
            if false {
                var v32 int16 = (v4 | v4)
                fmt.Println(v32)
                delete(v1, 42)
                _ = v32
            } else {
                var v33 string = v7
                v34 := GenS2{F0: 1}
                for fi35 := 0; fi35 < 9; fi35++ {
                    delete(v1, 440)
                    switch "a" {
                    case "":
                        if ((0.0 / (3.14159 - (float64((3.14159 * 3.14159)) - 1e10))) > (1e10 * (0.1 - ((-2.25 - 0.1) * (1.5 / (3.14159 + 0.1)))))) {
                            var v36 uint64 = v3
                            v37 := GenS3{F0: 9223372036854775807}
                            var v38 int16 = (v4 * (v4 >> ((65535) & 63)))
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                    case "foo bar":
                        v33 = v33
                    case "hello":
                        var v39 int32 = -225
                        _ = v39
                    default:
                        if true {
                            var v40 uint8 = 42
                            v40 += v40
                            if len(v2) > 0 {
                                v2[0] = 2
                            }
                            _ = v40
                        } else {
                            var v41 bool = false
                            _ = v41
                        }
                    }
                    fmt.Println(fi31, v4)
                    _ = fi35
                }
                _ = v33
                _ = v34
            }
            _ = fi31
        }
    }
    _ = len(v5)
    var v42 int32 = 2147483647
    for fi43 := 0; fi43 < 8; fi43++ {
        v44 := GenS1{F0: 4294967295, F1: (^(v42))}
        v2 = append(v2, 0)
        _ = fi43
        _ = v44
    }
    switch v8 {
    case -2147483648:
        v45 := GenS4{F0: 18446744073709551615, F1: -580, F2: 2}
        _ = v45
    default:
        if true {
            switch int(v42) {
            case 32767:
                if false {
                    v7 = v7
                    if true {
                        v7 = (v7 + "hello")
                        _ = len(v1)
                        var v46 int = v8
                        _ = v46
                    } else {
                        for fi47 := 0; fi47 < 12; fi47++ {
                            if ((v8 % ((v8) | 1)) > (fi47 ^ (fi47 * 9223372036854775807))) {
                                var v48 int16 = (v4 % ((0) | 1))
                                _ = len(v1)
                                _ = v48
                            }
                            var v49 int = 1
                            _ = fi47
                            _ = v49
                        }
                        v42 -= -1
                        v2 = append(v2, 42)
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    for fi50 := 0; fi50 < 2; fi50++ {
                        if true {
                            var v51 uint32 = 0
                            var v52 int = fi50
                            fmt.Println(v3, v7, v51)
                            _ = v51
                            _ = v52
                        } else {
                            if (false && false) {
                                fmt.Println(v3, v8, v42)
                            } else {
                                v2 = append(v2, 244)
                                v53 := GenS3{F0: fi50}
                                _ = v53
                            }
                            switch "" {
                            case "\n\t":
                                for fi54 := 0; fi54 < 8; fi54++ {
                                    v1[1] = v7
                                    _ = fi54
                                }
                            case "":
                                switch v7 {
                                case "hello":
                                    switch -702 {
                                    case 127:
                                        var v55 int32 = -2147483648
                                        _ = v55
                                    case 42:
                                        v1[42] = (v7 + "a")
                                    default:
                                        if false {
                                            v7 = ""
                                            v8 = (v8 << ((255) & 63))
                                        } else {
                                            for fi56 := 0; fi56 < 18; fi56++ {
                                                var v57 int8 = -1
                                                _ = fi56
                                                _ = v57
                                            }
                                            if true {
                                                _ = len(v1)
                                            }
                                        }
                                    }
                                case "a":
                                    switch (v7 + v7) {
                                    case "foo bar":
                                        switch (((v7 + (((v7 + v7) + ("hello" + (v7 + v7))) + "\n\t")) + v7) + "foo bar") {
                                        case "":
                                            var v58 int8 = -1
                                            _ = v58
                                        case "a":
                                            if (true || true) {
                                                var v59 int64 = -1
                                                if true {
                                                    if false {
                                                        var v60 uint = 99
                                                        _ = v60
                                                    }
                                                    v61 := GenS5{F0: 590, F1: "a"}
                                                    _ = v61
                                                }
                                                var v62 int = (-(v8))
                                                _ = v59
                                                _ = v62
                                            }
                                        default:
                                            if ((-1 > 9223372036854775807) || false) {
                                                var v63 int64 = -9223372036854775808
                                                v63 = v63
                                                _ = v63
                                            }
                                        }
                                    case "hello":
                                        v8 = v8
                                    }
                                default:
                                    if false {
                                        v4 = 0
                                    }
                                }
                            case "a":
                                for fi64 := 0; fi64 < 10; fi64++ {
                                    _ = len(v1)
                                    v3 = v3
                                    _ = fi64
                                }
                            default:
                                var v65 float32 = ((-2.25 * 0.0) / float32(1e10))
                                _ = v65
                            }
                        }
                        if (42 == 255) {
                            v42 -= v42
                        } else {
                            for fi66 := 0; fi66 < 18; fi66++ {
                                var v67 int = -388
                                _ = fi66
                                _ = v67
                            }
                            var v68 bool = (-1 >= int8(v4))
                            _ = v68
                        }
                        var v69 uint16 = 32767
                        _ = fi50
                        _ = v69
                    }
                    var v70 uint = 174
                    for fi71 := 0; fi71 < 6; fi71++ {
                        _ = len(v1)
                        v42 += (-(v42))
                        for fi72 := 0; fi72 < 15; fi72++ {
                            v42 = (v42 * v42)
                            if (float64((3.14159 - 0.1)) != (1e10 + 0.0)) {
                                v3 = (^(v3))
                            }
                            if (v3 >= v3) {
                                v2 = append(v2, 741)
                                fmt.Println(v8, v7, v7)
                            } else {
                                var v73 string = "foo bar"
                                if (0.0 >= -2.25) {
                                    v74 := GenS1{F0: 1, F1: 42}
                                    _ = v74
                                } else {
                                    v7 = ((("\n\t" + "hello") + "foo bar") + v73)
                                    v42 = (v42 | v42)
                                    fmt.Println(v4)
                                }
                                _ = v73
                            }
                            _ = fi72
                        }
                        _ = fi71
                    }
                    _ = v70
                }
            default:
                v7 = v7
            }
            v75 := GenS5{F0: 2, F1: ("hello" + v7)}
            _ = v75
        } else {
            v3 += v3
            if false {
                var v76 int8 = -89
                _ = v76
            }
            if false {
                v42 += (v42 >> ((0) & 63))
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
