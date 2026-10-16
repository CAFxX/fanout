package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int16 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint }
type GenS3 struct { F0 int32; F1 int16; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 0
    var v3 uint32 = 510
    var v4 float64 = 1e10
    var v5 uint64 = 614
    var v6 int32 = 42
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint16 = 65535
        v9 := []int{}
        if (false && true) {
            var v10 int16 = 62
            switch (fi7 << ((0) & 63)) {
            case 0:
                v4 -= float64(1e10)
            case -20:
                fmt.Println(v4)
            default:
                var v11 int64 = v2
                _ = v11
            }
            _ = v10
        } else {
            for fi12 := 0; fi12 < 20; fi12++ {
                v5 = (v5 << ((42) & 63))
                var v13 int64 = 649
                _ = fi12
                _ = v13
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    switch -455 {
    case 128:
        v14 := GenS1{F0: 255, F1: 240, F2: -452}
        _ = v14
    case 2147483647:
        switch "\n\t" {
        case "foo bar":
            var v15 uint8 = 0
            _ = v15
        default:
            var v16 float32 = 0.1
            _ = v16
        }
    case -9223372036854775808:
        v2 = (v2 % ((v2) | 1))
    }
    if (false || (!((true && ((true || (true && (v4 > (0.1 * 1.5)))) || (4294967295 >= int(v5)))) && (0 > -9223372036854775808)))) {
        switch -814 {
        case 2147483647:
            v1[-1] = ""
        case -2147483648:
            v17 := []int{32767, -326, -229}
            _ = v17
        case 42:
            v18 := []int{-211}
            _ = v18
        }
        for fi19 := 0; fi19 < 2; fi19++ {
            v20 := []int{(fi19 / (((fi19 | (fi19 << ((161) & 63)))) | 1))}
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 13; fi21++ {
            v22 := []int{fi21, -2147483648, fi21, fi21}
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 uint32 = v3
                v4 += v4
                _ = fi23
                _ = v24
            }
            _ = fi21
            _ = v22
        }
    } else {
        fmt.Println(v4, v4)
        switch 860 {
        case 4294967295:
            var v25 string = "hello"
            _ = v25
        }
    }
    v3 -= (v3 + (v3 * 843))
    if false {
        v26 := GenS1{F0: 4294967295, F1: 367, F2: 1}
        v27 := GenS2{F0: (v5 | (v5 - v5)), F1: v6, F2: 4294967295}
        if true {
            if (1.5 <= 0.1) {
                var v28 int16 = 958
                var v29 string = "foo bar"
                _ = v28
                _ = v29
            }
            var v30 string = ("a" + ((("" + "\n\t") + ("hello" + "hello")) + "\n\t"))
            _ = v30
        } else {
            v31 := GenS1{F0: 986, F1: 255, F2: -707}
            _ = len(v1)
            if false {
                var v32 float32 = ((-2.25 * float32(1e10)) * 3.14159)
                v33 := []int{127, 46}
                _ = v32
                _ = v33
            } else {
                for fi34 := 0; fi34 < 7; fi34++ {
                    if (!true) {
                        if true {
                            v35 := GenS3{F0: (-(v6)), F1: 1000, F2: 65535}
                            if true {
                                if (100 > 59) {
                                    switch (fi34 * 0) {
                                    case -728:
                                        var v36 int8 = -128
                                        _ = v36
                                    default:
                                        var v37 bool = false
                                        _ = v37
                                    }
                                    v6 = (v6 << ((65535) & 63))
                                    var v38 uint64 = (^(v5))
                                    _ = v38
                                } else {
                                    delete(v1, fi34)
                                    var v39 int32 = (v6 ^ -652)
                                    _ = v39
                                }
                                v4 = -2.25
                            }
                            _ = v35
                        } else {
                            var v40 int16 = -650
                            v41 := GenS4{F0: true, F1: v3}
                            v2 += (^(v2))
                            _ = v40
                            _ = v41
                        }
                    }
                    if (v4 < 0.0) {
                        var v42 int = (fi34 - fi34)
                        _ = v42
                    }
                    v3 -= v3
                    _ = fi34
                }
                v43 := GenS4{F0: false, F1: (^(v3))}
                var v44 float32 = -2.25
                _ = v43
                _ = v44
            }
            _ = v31
        }
        _ = v26
        _ = v27
    }
    for fi45 := 0; fi45 < 4; fi45++ {
        var v46 string = "a"
        var v47 int64 = (v2 + -1)
        var v48 float64 = (((v4 / float64((3.14159 / 1.5))) / 0.1) - ((v4 - (v4 / v4)) / ((float64(0.1) + v4) * 3.14159)))
        _ = fi45
        _ = v46
        _ = v47
        _ = v48
    }
    v49 := []int{173}
    _ = len(v49)
    delete(v1, -34)
    fmt.Println(v3, v4, v2)
    var v50 uint = 65535
    if true {
        var v51 int32 = v6
        var v52 uint64 = (v5 ^ uint64(v3))
        _ = v51
        _ = v52
    } else {
        for fi53 := 0; fi53 < 6; fi53++ {
            var v54 float32 = 1.5
            fmt.Println(v4, v6)
            v4 = v4
            _ = fi53
            _ = v54
        }
        v2 -= (^(v2))
    }
    if ((v4 * v4) > v4) {
        if true {
            if false {
                for fi55 := 0; fi55 < 4; fi55++ {
                    var v56 bool = (!true)
                    var v57 uint64 = v5
                    _ = fi55
                    _ = v56
                    _ = v57
                }
                if ((^(v3)) > uint32(v50)) {
                    v1[2147483647] = ("" + "hello")
                    fmt.Println(v2, v4)
                } else {
                    v58 := GenS1{F0: (^(v50)), F1: (v50 + v50), F2: 42}
                    _ = v58
                }
            }
        }
        if ((1 > 128) && false) {
            for fi59 := 0; fi59 < 11; fi59++ {
                v5 += (v5 & (^(v5)))
                fmt.Println(v50)
                var v60 uint16 = 66
                _ = fi59
                _ = v60
            }
            v3 = 111
            switch int(v2) {
            case 4294967295:
                var v61 int32 = (v6 % ((int32(v50)) | 1))
                _ = v61
            case 9223372036854775807:
                if len(v49) >= 2 {
                    v49 = v49[1:len(v49)]
                }
            case 128:
                for fi62 := 0; fi62 < 11; fi62++ {
                    v63 := GenS4{F0: false, F1: (v3 - v3)}
                    _ = fi62
                    _ = v63
                }
            default:
                if len(v49) >= 2 {
                    v49 = v49[1:len(v49)]
                }
            }
        } else {
            if (((0.0 < 0.0) && (!(2 == 255))) || (120 != 127)) {
                v50 -= v50
                if (0 <= 42) {
                    var v64 int16 = 320
                    fmt.Println(v6, v3)
                    switch "\n\t" {
                    case "\n\t":
                        var v65 int = 835
                        _ = v65
                    case "":
                        _ = len(v1)
                    default:
                        v3 = (v3 / ((2) | 1))
                    }
                    _ = v64
                }
            }
            v5 = v5
            v4 = -2.25
        }
        fmt.Println(v3, v50, v4)
    }
    v1[-1] = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
