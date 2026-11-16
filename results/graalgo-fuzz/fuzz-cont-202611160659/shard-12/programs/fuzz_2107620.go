package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int32; F1 int; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool; F1 int64 }
type GenS9 struct { F0 uint8 }
func genFn0() uint16 {
    if true {
        var v5 uint32 = 0
        _ = v5
    } else {
        var v6 float64 = 3.14159
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int8 = 127
        v9 := []int{fi7, fi7}
        if true {
            for fi10 := 0; fi10 < 10; fi10++ {
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                if true {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    v8 -= (v8 % (((v8 & (-(v8)))) | 1))
                    v8 += v8
                    switch "\n\t" {
                    case "":
                        v9 = append(v9, fi10)
                    case "hello":
                        switch 32767 {
                        case -517:
                            if false {
                                v9 = append(v9, fi7)
                                if (((fi10 / ((29) | 1)) < (fi7 << ((1) & 63))) || (false || (false && false))) {
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    fmt.Println(fi10)
                                } else {
                                    fmt.Println(fi10, fi7, fi7)
                                    fmt.Println(fi7, fi7)
                                    v11 := GenS1{F0: 32767, F1: v8}
                                    _ = v11
                                }
                            }
                        case 0:
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                        case 4294967295:
                            v12 := GenS1{F0: -32768, F1: (v8 + (v8 % ((127) | 1)))}
                            _ = v12
                        }
                    case "foo bar":
                        var v13 string = "hello"
                        _ = v13
                    }
                    v8 = v8
                }
                v8 -= (v8 - v8)
                _ = fi10
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v14 := []int{-2147483648, 9223372036854775807}
    _ = v14
    return 2
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 1
    var v3 uint = 988
    var v4 uint16 = 32767
    fmt.Println(v2, v2)
    var v15 uint = 2
    v16 := []int{}
    var v17 int64 = 852
    if (genFn0() >= v4) {
        v18 := GenS2{F0: 399, F1: -2147483648, F2: -2147483648}
        if (-128 != 127) {
            switch -9223372036854775808 {
            case -42:
                if (255 >= v15) {
                    v17 += -182
                    v17 = v17
                } else {
                    v2 -= v2
                    _ = len(v16)
                    v16 = append(v16, int(v2))
                }
            case 1:
                if true {
                    switch int(v17) {
                    case -2147483648:
                        v2 -= (v2 / (((v2 ^ (v2 / ((v2) | 1)))) | 1))
                    case -615:
                        var v19 uint16 = uint16(v3)
                        _ = v19
                    case -9223372036854775808:
                        var v20 int32 = int32(v15)
                        _ = v20
                    default:
                        fmt.Println(v2)
                    }
                    var v21 int8 = int8(v4)
                    var v22 float64 = float64(float32(3.14159))
                    _ = v21
                    _ = v22
                }
            case 863:
                v3 += 65535
            default:
                v4 = uint16(v3)
            }
        }
        var v23 float64 = ((3.14159 * 0.0) * 0.0)
        _ = v18
        _ = v23
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        v2 = v2
        for fi25 := 0; fi25 < 8; fi25++ {
            v15 = (v3 >> ((848) & 63))
            if len(v16) > 0 {
                v16[0] = (fi25 * (fi25 % (((fi25 / (((-(fi25))) | 1))) | 1)))
            }
            if true {
                var v26 float64 = -2.25
                switch (fi24 / (((fi25 / ((fi25) | 1))) | 1)) {
                case -690:
                    switch fi24 {
                    case 9223372036854775807:
                        if false {
                            v4 = 0
                        } else {
                            var v27 uint64 = 650
                            if false {
                                var v28 int16 = -1
                                if false {
                                    v26 = v26
                                } else {
                                    v15 += v15
                                }
                                if false {
                                    fmt.Println(v2)
                                    var v29 string = "a"
                                    _ = v29
                                } else {
                                    if (v27 != 18446744073709551615) {
                                        _ = len(v16)
                                        var v30 float64 = v26
                                        var v31 int16 = v28
                                        _ = v30
                                        _ = v31
                                    }
                                    v32 := GenS1{F0: (v28 - v28), F1: int8(v15)}
                                    v17 += v17
                                    var v33 uint = (^(v15))
                                    if (!(v27 > (v27 * (v27 / (((v27 + (v27 - v27))) | 1))))) {
                                        if (false || (!(true || false))) {
                                            if len(v16) > 0 {
                                                v16[0] = fi25
                                            }
                                        }
                                    }
                                    _ = v32
                                    _ = v33
                                }
                                _ = v28
                            } else {
                                switch "hello" {
                                case "hello":
                                    v27 -= v27
                                case "a":
                                    switch ("" + "a") {
                                    case "foo bar":
                                        fmt.Println(v17, v2, fi25)
                                    case "":
                                        v15 = v15
                                    case "\n\t":
                                        var v34 uint8 = 186
                                        _ = v34
                                    }
                                }
                                var v35 int16 = 0
                                _ = v35
                            }
                            var v36 int32 = -901
                            _ = v27
                            _ = v36
                        }
                    default:
                        v37 := GenS3{F0: (v2 ^ v2)}
                        _ = v37
                    }
                case 42:
                    var v38 int64 = v17
                    _ = v38
                default:
                    v15 += 255
                }
                _ = v26
            } else {
                _ = len(v1)
                if (true || (int8(v4) != 100)) {
                    v39 := GenS3{F0: v2}
                    _ = v39
                } else {
                    if ((v4 | 65535) != (v4 | genFn0())) {
                        v40 := v1["hello"]
                        _ = v40
                    }
                    var v41 float32 = -2.25
                    _ = v41
                }
            }
            _ = fi25
        }
        v42 := GenS2{F0: -2147483648, F1: fi24, F2: fi24}
        _ = fi24
        _ = v42
    }
    if ((^(v3)) >= v15) {
        if (v17 < (v17 - (v17 % ((v17) | 1)))) {
            v2 = (^(v2))
            var v43 uint32 = (^(v2))
            _ = v43
        }
        if true {
            var v44 int = -606
            var v45 string = "\n\t"
            _ = v44
            _ = v45
        } else {
            fmt.Println(v15, v17)
            if (!true) {
                var v46 int = 458
                switch "hello" {
                case "\n\t":
                    if true {
                        v47 := GenS1{F0: 1, F1: 50}
                        var v48 int16 = 1000
                        delete(v1, "hello")
                        _ = v47
                        _ = v48
                    }
                case "hello":
                    switch 128 {
                    case 127:
                        for fi49 := 0; fi49 < 9; fi49++ {
                            switch fi49 {
                            case -1:
                                v46 = fi49
                            case 625:
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                            case 32768:
                                for fi50 := 0; fi50 < 3; fi50++ {
                                    var v51 uint64 = 362
                                    _ = fi50
                                    _ = v51
                                }
                            default:
                                delete(v1, ("foo bar" + "a"))
                            }
                            _ = fi49
                        }
                    case 1:
                        v2 = (v2 % ((v2) | 1))
                    }
                }
                if false {
                    var v52 string = "foo bar"
                    var v53 uint8 = 2
                    _ = v52
                    _ = v53
                } else {
                    var v54 uint32 = v2
                    for fi55 := 0; fi55 < 2; fi55++ {
                        if true {
                            if (968 > v3) {
                                if len(v16) > 0 {
                                    v16[0] = (fi55 << ((42) & 63))
                                }
                            } else {
                                switch ("\n\t" + "a") {
                                case "\n\t":
                                    for fi56 := 0; fi56 < 5; fi56++ {
                                        var v57 uint = v15
                                        var v58 uint32 = (v54 - v2)
                                        if (127 <= -107) {
                                            v15 += v3
                                            if false {
                                                _ = len(v16)
                                                switch -346 {
                                                case 115:
                                                    fmt.Println(fi55)
                                                }
                                            }
                                            var v59 uint = v3
                                            _ = v59
                                        }
                                        _ = fi56
                                        _ = v57
                                        _ = v58
                                    }
                                default:
                                    for fi60 := 0; fi60 < 3; fi60++ {
                                        v61 := GenS4{F0: 1, F1: false}
                                        v62 := GenS5{F0: 596, F1: (v4 << ((255) & 63))}
                                        _ = fi60
                                        _ = v61
                                        _ = v62
                                    }
                                }
                                for fi63 := 0; fi63 < 6; fi63++ {
                                    _ = len(v1)
                                    var v64 uint16 = (^(v4))
                                    var v65 uint = (v15 / (((v3 * v15)) | 1))
                                    var v66 int32 = 0
                                    _ = fi63
                                    _ = v64
                                    _ = v65
                                    _ = v66
                                }
                                fmt.Println(v3, v15, v3)
                            }
                        }
                        _ = fi55
                    }
                    _ = v54
                }
                _ = v46
            } else {
                var v67 float64 = 1.5
                v68 := GenS6{F0: 1, F1: v2, F2: v2}
                _ = v67
                _ = v68
            }
            _ = len(v1)
        }
        _ = len(v16)
    }
    if true {
        v3 += 4294967295
        delete(v1, "foo bar")
        switch (("foo bar" + "") + "a") {
        case "hello":
            for fi69 := 0; fi69 < 2; fi69++ {
                var v70 uint = 415
                _ = fi69
                _ = v70
            }
        case "foo bar":
            var v71 int = -9223372036854775808
            _ = v71
        default:
            fmt.Println(v15)
        }
    }
    v2 = v2
    if true {
        v3 -= (v15 - v15)
        delete(v1, ("\n\t" + "a"))
        v16 = append(v16, 32767)
    } else {
        v17 += -1
    }
    for fi72 := 0; fi72 < 11; fi72++ {
        for fi73 := 0; fi73 < 3; fi73++ {
            if false {
                _ = len(v1)
                if true {
                    if (65535 >= v15) {
                        _ = len(v1)
                        if (1 > 42) {
                            var v74 float64 = 0.0
                            v75 := GenS7{F0: "\n\t"}
                            v3 = v15
                            delete(v1, "\n\t")
                            _ = v74
                            _ = v75
                        }
                    } else {
                        _ = len(v16)
                        _ = len(v16)
                    }
                    v15 = v3
                }
                v76 := GenS3{F0: v2}
                _ = v76
            }
            if false {
                v77 := GenS8{F0: (17 > 1), F1: v17}
                _ = v77
            } else {
                fmt.Println(v2, v3, v3)
                v78 := GenS9{F0: 255}
                _ = v78
            }
            var v79 bool = false
            _ = fi73
            _ = v79
        }
        _ = fi72
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
