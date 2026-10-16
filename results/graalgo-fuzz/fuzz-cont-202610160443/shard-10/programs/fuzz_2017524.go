package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 int8; F2 int32 }
type GenS5 struct { F0 int8; F1 uint16; F2 string }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 uint8; F1 uint32; F2 uint16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    var v3 int16 = 233
    var v4 float32 = 0.0
    if ((((1 > 1) || (true || (-2147483648 != 0))) || (-1 >= 0)) || true) {
        var v5 float32 = v4
        _ = v5
    }
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int64 = 0
        var v8 uint32 = 640
        var v9 uint = 42
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    fmt.Println(v3, v4)
    var v10 int8 = 1
    var v11 int = -2147483648
    var v12 float32 = 0.0
    switch (v11 % ((v11) | 1)) {
    case 128:
        for fi13 := 0; fi13 < 9; fi13++ {
            delete(v1, 127)
            _ = fi13
        }
    case 32767:
        for fi14 := 0; fi14 < 13; fi14++ {
            fmt.Println(v4, v12, v11)
            v15 := []int{32767, (v11 | (fi14 & (v11 << ((289) & 63)))), fi14, (-(v11))}
            v12 = v12
            _ = fi14
            _ = v15
        }
    }
    fmt.Println(v12)
    if true {
        for fi16 := 0; fi16 < 9; fi16++ {
            for fi17 := 0; fi17 < 2; fi17++ {
                v3 = 0
                if true {
                    v18 := v2[0]
                    switch ("a" + ("" + ("\n\t" + ""))) {
                    case "foo bar":
                        v10 = v10
                    case "\n\t":
                        var v19 float32 = (v12 * (3.14159 - (0.0 / v4)))
                        _ = v19
                    default:
                        v2[2147483647] = (v18 / (((v18 % ((v18) | 1))) | 1))
                    }
                    v20 := GenS1{F0: true, F1: true, F2: v3}
                    _ = v18
                    _ = v20
                } else {
                    v21 := v1[127]
                    v22 := []int{fi17, fi17, fi16, int(v21)}
                    _ = v21
                    _ = v22
                }
                _ = fi17
            }
            v23 := GenS1{F0: true, F1: (!(!true)), F2: (v3 - v3)}
            v24 := []int{fi16}
            _ = fi16
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 6; fi25++ {
            for fi26 := 0; fi26 < 16; fi26++ {
                v10 += (v10 / ((0) | 1))
                var v27 int = fi25
                _ = fi26
                _ = v27
            }
            for fi28 := 0; fi28 < 15; fi28++ {
                v29 := []int{fi28, fi28, (-(fi28))}
                v30 := GenS1{F0: (837 != int32(v10)), F1: (2 == 42), F2: 32767}
                switch (fi25 & -2147483648) {
                case 128:
                    v4 += (1.5 + v12)
                case 1:
                    v31 := GenS1{F0: (-856 == -2147483648), F1: ((0 <= 9223372036854775807) && true), F2: v3}
                    _ = v31
                case 2147483647:
                    switch 32768 {
                    case 2147483647:
                        fmt.Println(v11, v12)
                    case -1:
                        v32 := GenS1{F0: ((1e10 + (0.0 - -2.25)) >= 3.14159), F1: false, F2: 0}
                        _ = v32
                    case 1:
                        v12 = v12
                    }
                default:
                    if false {
                        if true {
                            fmt.Println(fi25)
                            var v33 int = (fi28 + fi28)
                            _ = v33
                        } else {
                            fmt.Println(v4, v12, fi25)
                            if true {
                                var v34 uint32 = 2
                                v4 = ((1e10 - -2.25) / v12)
                                var v35 uint8 = 173
                                if false {
                                    delete(v1, 32767)
                                    v34 = 1
                                } else {
                                    var v36 bool = true
                                    _ = v36
                                }
                                _ = v34
                                _ = v35
                            } else {
                                var v37 uint64 = 2
                                _ = v37
                            }
                            var v38 uint16 = 2
                            switch "a" {
                            case "\n\t":
                                fmt.Println(v11, fi28, v10)
                            case "":
                                v39 := GenS2{F0: (v3 | (v3 & v3))}
                                _ = v39
                            default:
                                _ = len(v1)
                            }
                            _ = v38
                        }
                        if (true || false) {
                            v40 := GenS2{F0: (-(v3))}
                            v41 := GenS1{F0: false, F1: (-663 < 2147483647), F2: v3}
                            _ = v40
                            _ = v41
                        }
                        var v42 bool = false
                        _ = v42
                    } else {
                        if len(v29) > 0 {
                            v29[0] = (fi28 & v11)
                        }
                        var v43 float32 = (v4 * 0.1)
                        _ = v43
                    }
                }
                if (939 >= 65535) {
                    if (1 != 4294967295) {
                        v44 := GenS3{F0: false}
                        var v45 int8 = (v10 - v10)
                        var v46 uint32 = 747
                        _ = v44
                        _ = v45
                        _ = v46
                    }
                    var v47 int8 = v10
                    fmt.Println(v10, v3, v10)
                    _ = v47
                }
                var v48 int = (fi25 % ((fi25) | 1))
                v12 = v4
                _ = fi28
                _ = v29
                _ = v30
                _ = v48
            }
            for fi49 := 0; fi49 < 3; fi49++ {
                fmt.Println(fi25, fi49)
                var v50 uint64 = 530
                v12 += (v4 + (-2.25 / float32(-2.25)))
                v10 = (v10 ^ (v10 >> ((915) & 63)))
                v12 = 3.14159
                fmt.Println(fi25, fi49)
                _ = fi49
                _ = v50
            }
            _ = fi25
        }
    }
    for fi51 := 0; fi51 < 18; fi51++ {
        if true {
            var v52 int8 = v10
            v10 += v52
            var v53 uint32 = 42
            _ = v52
            _ = v53
        } else {
            fmt.Println(v12, v12, v11)
        }
        _ = fi51
    }
    if ((v3 & v3) >= (^(v3))) {
        v54 := GenS4{F0: 65535, F1: (v10 % ((v10) | 1)), F2: -2147483648}
        for fi55 := 0; fi55 < 4; fi55++ {
            if ((65535 < 32767) || (true && (false || (!false)))) {
                var v56 int64 = 351
                fmt.Println(v10, fi55)
                v57 := GenS3{F0: true}
                _ = v56
                _ = v57
            } else {
                var v58 int = (fi55 - 480)
                v59 := GenS5{F0: v10, F1: 2, F2: ("foo bar" + ("" + "\n\t"))}
                switch (^(v58)) {
                case 871:
                    v10 += v10
                case 0:
                    if (-557 >= 42) {
                        fmt.Println(fi55, v3)
                        if false {
                            v2[1] = 735
                            if (((("" + "a") + "\n\t") == "\n\t") && true) {
                                fmt.Println(v58, v4)
                                switch (v58 + (v58 ^ fi55)) {
                                case 32767:
                                    v10 -= (v10 / ((v10) | 1))
                                case 32768:
                                    v12 = v4
                                default:
                                    var v60 string = ""
                                    _ = v60
                                }
                                for fi61 := 0; fi61 < 5; fi61++ {
                                    var v62 int = (^(fi55))
                                    fmt.Println(v58, v62)
                                    _ = fi61
                                    _ = v62
                                }
                            }
                        } else {
                            fmt.Println(v10)
                            v63 := GenS6{F0: 1, F1: v10}
                            if false {
                                var v64 float32 = v12
                                if (0 <= 127) {
                                    for fi65 := 0; fi65 < 3; fi65++ {
                                        switch "hello" {
                                        case "a":
                                            v66 := GenS5{F0: v10, F1: 117, F2: ("a" + (("hello" + ("" + ("\n\t" + "a"))) + ("hello" + "foo bar")))}
                                            _ = v66
                                        case "\n\t":
                                            v67 := []int{}
                                            _ = v67
                                        case "":
                                            v11 = (v11 >> ((838) & 63))
                                        }
                                        v4 += (v64 - v12)
                                        v68 := GenS7{F0: 0, F1: 65535, F2: 255}
                                        var v69 int8 = v10
                                        _ = fi65
                                        _ = v68
                                        _ = v69
                                    }
                                    if (-2147483648 == 42) {
                                        var v70 bool = (18446744073709551615 >= 2)
                                        v71 := []int{}
                                        _ = v70
                                        _ = v71
                                    } else {
                                        if true {
                                            var v72 float64 = 1e10
                                            v72 += v72
                                            v73 := []int{}
                                            _ = v72
                                            _ = v73
                                        }
                                        v74 := GenS3{F0: (!true)}
                                        v12 += v12
                                        _ = v74
                                    }
                                }
                                var v75 float64 = -2.25
                                _ = v64
                                _ = v75
                            }
                            _ = v63
                        }
                        v12 -= 0.1
                    }
                default:
                    switch -9223372036854775808 {
                    case 1:
                        v3 -= (v3 >> ((809) & 63))
                    }
                }
                _ = v58
                _ = v59
            }
            _ = fi55
        }
        _ = v54
    } else {
        if true {
            var v76 uint16 = 2
            v77 := GenS7{F0: 2, F1: 2147483647, F2: 824}
            v78 := GenS4{F0: v76, F1: v10, F2: 1}
            _ = v76
            _ = v77
            _ = v78
        } else {
            v79 := []int{(v11 / ((32768) | 1)), (v11 & 32767), v11}
            _ = v79
        }
        var v80 int16 = 0
        var v81 uint8 = 1
        _ = v80
        _ = v81
    }
    var v82 int8 = v10
    v83 := v1[v11]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v82
    _ = v83
}
