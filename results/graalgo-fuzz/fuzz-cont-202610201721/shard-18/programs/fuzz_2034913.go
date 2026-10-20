package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int; F2 int16 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16; F1 uint32; F2 uint16 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[uint32]uint32)
    var v3 int8 = -56
    v3 -= (v3 | v3)
    v3 += v3
    delete(v2, 2147483648)
    v3 += (-(v3))
    v4 := GenS1{F0: 947}
    if true {
        var v5 int16 = 1
        _ = v5
    } else {
        if ((("\n\t" + "foo bar") + ("" + "foo bar")) > ("a" + ("hello" + "\n\t"))) {
            v3 = (v3 * v3)
            var v6 bool = true
            _ = v6
        }
    }
    if true {
        for fi7 := 0; fi7 < 2; fi7++ {
            delete(v2, 1)
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 uint8 = 218
                switch fi7 {
                case -59:
                    var v10 int32 = -601
                    _ = v10
                case 258:
                    v3 = (-(v3))
                default:
                    v11 := []int{fi8, (fi8 / ((fi7) | 1))}
                    _ = v11
                }
                v9 = 140
                fmt.Println(fi8)
                _ = fi8
                _ = v9
            }
            var v12 float32 = -2.25
            _ = fi7
            _ = v12
        }
        v13 := GenS1{F0: 861}
        _ = v13
    }
    if false {
        var v14 float64 = 3.14159
        v1 += (v1 % (((v1 / ((v1) | 1))) | 1))
        if true {
            v1 -= (v1 | v1)
            _ = len(v2)
        }
        _ = v14
    } else {
        var v15 int32 = 1
        var v16 string = "foo bar"
        var v17 uint64 = 70
        _ = v15
        _ = v16
        _ = v17
    }
    if false {
        v18 := GenS1{F0: 675}
        if (!(-1 < -407)) {
            var v19 int64 = v1
            v2[2] = 2
            _ = v19
        } else {
            v20 := []int{127}
            if false {
                _ = len(v20)
                if true {
                    for fi21 := 0; fi21 < 19; fi21++ {
                        for fi22 := 0; fi22 < 7; fi22++ {
                            var v23 uint8 = 255
                            _ = fi22
                            _ = v23
                        }
                        v3 -= v3
                        _ = fi21
                    }
                    var v24 float64 = (-2.25 * 1.5)
                    _ = v24
                }
            } else {
                v3 += 127
            }
            v3 = (v3 & v3)
            _ = v20
        }
        _ = v18
    } else {
        if false {
            switch ("" + ("hello" + "\n\t")) {
            case "a":
                switch 128 {
                case 0:
                    for fi25 := 0; fi25 < 17; fi25++ {
                        v1 -= (v1 >> ((42) & 63))
                        _ = fi25
                    }
                case 690:
                    for fi26 := 0; fi26 < 6; fi26++ {
                        var v27 int16 = -1
                        if (true && (!true)) {
                            v28 := GenS2{F0: 4294967295, F1: fi26, F2: v27}
                            _ = v28
                        }
                        _ = fi26
                        _ = v27
                    }
                }
            case "\n\t":
                if false {
                    var v29 uint8 = 2
                    _ = v29
                }
            case "hello":
                var v30 uint16 = 255
                _ = v30
            }
            switch 2147483647 {
            case 32767:
                fmt.Println(v1, v1)
            case -366:
                for fi31 := 0; fi31 < 3; fi31++ {
                    _ = len(v2)
                    var v32 uint16 = 32767
                    switch "" {
                    case "":
                        if false {
                            switch (fi31 << ((65535) & 63)) {
                            case 128:
                                v33 := GenS1{F0: 32767}
                                _ = v33
                            case 583:
                                fmt.Println(v3, v1)
                            default:
                                fmt.Println(v3, fi31, fi31)
                            }
                            fmt.Println(v3, v1)
                            if (true || false) {
                                if false {
                                    switch (fi31 << ((4294967295) & 63)) {
                                    case 32767:
                                        var v34 string = (("hello" + ("a" + ("foo bar" + "hello"))) + "")
                                        _ = v34
                                    case -2147483648:
                                        switch fi31 {
                                        case -1:
                                            if (v32 >= (v32 + (v32 ^ v32))) {
                                                var v35 int8 = int8(v1)
                                                _ = v35
                                            }
                                        }
                                    }
                                }
                                if (2 <= 1) {
                                    var v36 bool = false
                                    _ = v36
                                }
                                var v37 int = fi31
                                _ = v37
                            }
                        } else {
                            v1 = v1
                        }
                    case "a":
                        var v38 uint = 42
                        _ = v38
                    }
                    _ = fi31
                    _ = v32
                }
            case 2147483647:
                v1 += v1
            }
            var v39 uint64 = 574
            _ = v39
        }
        for fi40 := 0; fi40 < 17; fi40++ {
            var v41 uint64 = 1
            if true {
                if (true || (true && (2 > 1))) {
                    fmt.Println(v3)
                    v3 -= int8(v1)
                    for fi42 := 0; fi42 < 9; fi42++ {
                        var v43 int64 = v1
                        v3 = 30
                        if (42 >= 255) {
                            var v44 uint64 = (v41 >> ((4294967295) & 63))
                            _ = v44
                        } else {
                            v43 -= (v43 | -9223372036854775808)
                            switch fi40 {
                            case 32767:
                                v2[65535] = 65535
                            case 127:
                                v45 := GenS3{F0: (^(v1)), F1: v41}
                                _ = v45
                            case 1:
                                fmt.Println(v3)
                            default:
                                if false {
                                    var v46 float64 = -2.25
                                    _ = v46
                                } else {
                                    v2[199] = 65535
                                }
                            }
                            v3 += 94
                            if ((("" <= ("foo bar" + ("" + "foo bar"))) && false) || false) {
                                switch fi40 {
                                case 4294967295:
                                    var v47 uint8 = 255
                                    _ = v47
                                case -9223372036854775808:
                                    v2[0] = 248
                                default:
                                    if (-32768 <= 0) {
                                        if (fi40 == (fi40 >> ((4294967295) & 63))) {
                                            v48 := []int{(fi42 / (((fi42 / ((int(v43)) | 1))) | 1)), (^(fi42))}
                                            var v49 uint64 = v41
                                            _ = v48
                                            _ = v49
                                        } else {
                                            v3 = -55
                                            switch ("\n\t" + (("a" + "\n\t") + "hello")) {
                                            case "":
                                                v50 := GenS1{F0: 255}
                                                _ = v50
                                            default:
                                                v43 += v43
                                            }
                                        }
                                    } else {
                                        v2[65535] = 1
                                        if (1 >= 900) {
                                            v2[65535] = 470
                                            v51 := GenS1{F0: 32767}
                                            v41 = v41
                                            _ = v51
                                        }
                                        fmt.Println(v3, fi42)
                                    }
                                }
                                v43 -= v1
                            } else {
                                v52 := GenS1{F0: 2}
                                v1 = 897
                                v2[2147483647] = 2147483648
                                _ = v52
                            }
                            v3 = -128
                            fmt.Println(v43, v41, v43)
                        }
                        _ = fi42
                        _ = v43
                    }
                }
            }
            for fi53 := 0; fi53 < 7; fi53++ {
                switch "" {
                case "\n\t":
                    if true {
                        v54 := []int{}
                        _ = v54
                    }
                case "a":
                    var v55 uint8 = 159
                    _ = v55
                case "foo bar":
                    if false {
                        v3 = v3
                        var v56 float64 = float64((3.14159 + -2.25))
                        switch "a" {
                        case "foo bar":
                            var v57 bool = (false && true)
                            _ = v57
                        case "hello":
                            v58 := GenS4{F0: (!true)}
                            _ = v58
                        default:
                            v41 += (v41 & (^(v41)))
                        }
                        _ = v56
                    } else {
                        v59 := GenS5{F0: 32767, F1: 0, F2: 747}
                        v41 -= (v41 & (v41 % ((v41) | 1)))
                        var v60 uint16 = 42
                        _ = v59
                        _ = v60
                    }
                }
                if (0 == 255) {
                    if (false && ((("\n\t" + ("\n\t" + "")) < "foo bar") && true)) {
                        var v61 int32 = 42
                        _ = v61
                    } else {
                        var v62 int = 32767
                        v1 = (v1 << ((42) & 63))
                        delete(v2, 1)
                        _ = v62
                    }
                    var v63 float64 = (0.1 - ((1e10 * 0.1) - (float64(1.5) / 0.1)))
                    _ = v63
                }
                _ = fi53
            }
            _ = fi40
            _ = v41
        }
    }
    if ("a" != "foo bar") {
        for fi64 := 0; fi64 < 16; fi64++ {
            var v65 int64 = v1
            var v66 int32 = 42
            var v67 float32 = float32((((-2.25 / (float64((float32(0.0) - (-2.25 * 0.1))) - 0.1)) + ((float64((((((1e10 - float32(1.5)) / 1.0) / -2.25) * (float32(3.14159) / float32(float64((1e10 - 0.0))))) / (-2.25 / -2.25))) * -2.25) * 3.14159)) / float64(1e10)))
            _ = fi64
            _ = v65
            _ = v66
            _ = v67
        }
        v3 = (v3 + -1)
        var v68 int = 0
        _ = v68
    }
    fmt.Println(v3)
    v69 := []int{}
    var v70 int8 = v3
    if (1000 >= 0) {
        var v71 float32 = 0.1
        if (true && false) {
            v70 = (v3 * (v70 >> ((965) & 63)))
            var v72 int32 = 42
            fmt.Println(v72)
            _ = v72
        }
        _ = v71
    } else {
        v70 = v3
    }
    v70 = (v3 >> ((296) & 63))
    fmt.Println(v70, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v69:", len(v69), v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v69
    _ = v70
}
