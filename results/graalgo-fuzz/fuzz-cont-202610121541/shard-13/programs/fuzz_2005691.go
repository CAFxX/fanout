package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 uint; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 string; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int, p1 uint) int16 {
    v4 := []int{p0, p0, (p0 % ((p0) | 1)), p0}
    switch p0 {
    case 1:
        p0 = (^(p0))
    case 2147483647:
        p0 += (p0 % ((4294967295) | 1))
    }
    _ = p0
    _ = p1
    _ = v4
    return 32767
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = 3.14159
    v3 := make(map[uint32]uint32)
    v2 = float64(0.1)
    switch "a" {
    case "\n\t":
        v2 = 1.5
    case "a":
        v5 := GenS1{F0: (127 < 128)}
        _ = v5
    }
    v2 += v2
    v6 := GenS2{F0: 127, F1: 946, F2: 792}
    if false {
        if false {
            v2 = v2
            v7 := GenS1{F0: (false || (true || false))}
            switch 127 {
            case 854:
                v8 := GenS1{F0: false}
                _ = v8
            case 4294967295:
                if true {
                    v9 := GenS2{F0: -396, F1: 65535, F2: 65535}
                    v2 += (v2 - (1.5 * (v2 + ((v2 * (((1e10 / (1.5 / v2)) * v2) / v2)) - ((float64(-2.25) - 1e10) + (3.14159 + 1.5))))))
                    _ = v9
                }
            default:
                switch "hello" {
                case "hello":
                    for fi10 := 0; fi10 < 15; fi10++ {
                        for fi11 := 0; fi11 < 9; fi11++ {
                            v2 -= (0.0 + v2)
                            v12 := []int{-907, (fi10 * fi10), fi11}
                            _ = len(v12)
                            _ = fi11
                            _ = v12
                        }
                        switch fi10 {
                        case 4294967295:
                            switch "foo bar" {
                            case "hello":
                                v2 = float64((1e10 - (((0.0 / 1.0) + 0.1) - ((-2.25 + 0.1) + 0.1))))
                            case "a":
                                switch ("\n\t" + "hello") {
                                case "":
                                    var v13 float32 = -2.25
                                    _ = v13
                                case "foo bar":
                                    var v14 int = (fi10 >> ((216) & 63))
                                    _ = v14
                                case "\n\t":
                                    var v15 int64 = -862
                                    _ = v15
                                }
                            case "foo bar":
                                for fi16 := 0; fi16 < 18; fi16++ {
                                    var v17 float64 = v2
                                    _ = fi16
                                    _ = v17
                                }
                            default:
                                v18 := []int{(fi10 << ((255) & 63))}
                                _ = v18
                            }
                        case 42:
                            for fi19 := 0; fi19 < 13; fi19++ {
                                switch (-(fi10)) {
                                case -1:
                                    fmt.Println(fi10, v2, fi19)
                                case 217:
                                    var v20 uint64 = 751
                                    _ = v20
                                }
                                switch "foo bar" {
                                case "a":
                                    fmt.Println(v2, fi19, fi10)
                                case "\n\t":
                                    switch 128 {
                                    case 127:
                                        if (!(!false)) {
                                            fmt.Println(fi10, fi19)
                                            v2 -= v2
                                        } else {
                                            var v21 float64 = (v2 / float64((0.1 + -2.25)))
                                            _ = v21
                                        }
                                    case 647:
                                        if false {
                                            fmt.Println(fi19)
                                            v22 := GenS1{F0: (genFn0(fi10, 407) >= 0)}
                                            if true {
                                                v2 += (3.14159 * 3.14159)
                                                v2 += 0.1
                                            } else {
                                                switch 517 {
                                                case -1:
                                                    v2 = v2
                                                case -865:
                                                    fmt.Println(fi10)
                                                case -854:
                                                    v23 := GenS1{F0: false}
                                                    _ = v23
                                                }
                                                v24 := GenS3{F0: -51, F1: 481}
                                                _ = v24
                                            }
                                            switch (fi10 | fi10) {
                                            case 32768:
                                                v25 := []int{fi10}
                                                _ = v25
                                            case -9223372036854775808:
                                                var v26 int64 = 42
                                                _ = v26
                                            case 4294967295:
                                                v27 := v3[42]
                                                _ = v27
                                            }
                                            _ = v22
                                        }
                                    default:
                                        v2 -= (-2.25 - v2)
                                    }
                                    v2 = float64(-2.25)
                                case "foo bar":
                                    if (0 == 1) {
                                        fmt.Println(fi10)
                                        delete(v3, 42)
                                        fmt.Println(fi19)
                                    } else {
                                        switch (fi19 % ((fi10) | 1)) {
                                        case 2147483647:
                                            delete(v3, 4294967295)
                                        case 688:
                                            switch -2147483648 {
                                            case -475:
                                                var v28 int8 = -128
                                                _ = v28
                                            case 881:
                                                var v29 bool = true
                                                _ = v29
                                            default:
                                                var v30 int8 = 100
                                                _ = v30
                                            }
                                        case -1:
                                            v31 := []int{(fi19 ^ -217), 786, fi10, fi10}
                                            _ = v31
                                        }
                                        v2 += (((v2 / v2) - ((float64((3.14159 + 3.14159)) + v2) + -2.25)) + 0.0)
                                        v2 = v2
                                    }
                                }
                                if ((fi10 << ((18446744073709551615) & 63)) == (fi19 - (fi10 >> ((255) & 63)))) {
                                    v32 := []int{-986, fi10, (fi10 - 607)}
                                    var v33 int16 = 32767
                                    _ = v32
                                    _ = v33
                                }
                                _ = fi19
                            }
                        }
                        _ = fi10
                    }
                default:
                    for fi34 := 0; fi34 < 17; fi34++ {
                        v2 -= 0.0
                        var v35 uint8 = 255
                        v36 := []int{}
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                }
            }
            _ = v7
        } else {
            if (false || true) {
                v2 -= -2.25
                v37 := GenS4{F0: ("\n\t" + ("hello" + "")), F1: 1, F2: 42}
                _ = v37
            } else {
                switch "foo bar" {
                case "a":
                    v1[42] = 51
                case "hello":
                    var v38 uint = 18446744073709551615
                    _ = v38
                case "\n\t":
                    var v39 int = 32768
                    _ = v39
                }
            }
            if true {
                v40 := GenS4{F0: "", F1: 53, F2: 422}
                var v41 float64 = ((v2 - v2) / v2)
                _ = v40
                _ = v41
            } else {
                if false {
                    var v42 string = "foo bar"
                    _ = v42
                }
                v2 = ((v2 - (3.14159 + 0.1)) + float64(((((0.1 - 0.0) / -2.25) / (0.1 + (3.14159 - 0.0))) + 3.14159)))
                switch "hello" {
                case "foo bar":
                    v43 := GenS5{F0: 0, F1: 657}
                    _ = v43
                default:
                    v44 := []int{}
                    _ = v44
                }
            }
        }
        if (false || true) {
            for fi45 := 0; fi45 < 20; fi45++ {
                for fi46 := 0; fi46 < 4; fi46++ {
                    var v47 uint64 = 42
                    if false {
                        v48 := GenS5{F0: 1, F1: 42}
                        if false {
                            v49 := GenS1{F0: true}
                            if false {
                                fmt.Println(v47)
                            } else {
                                v47 = v47
                                var v50 float32 = 0.1
                                var v51 uint16 = 855
                                _ = v50
                                _ = v51
                            }
                            _ = v49
                        } else {
                            if true {
                                v52 := []int{fi46, 32767, 9223372036854775807}
                                var v53 uint16 = 656
                                fmt.Println(v53, fi45, v53)
                                _ = v52
                                _ = v53
                            } else {
                                v47 = (v47 >> ((42) & 63))
                                var v54 string = "foo bar"
                                _ = v54
                            }
                        }
                        _ = v48
                    } else {
                        v55 := []int{}
                        v47 += (^(v47))
                        _ = v55
                    }
                    var v56 uint = 255
                    _ = fi46
                    _ = v47
                    _ = v56
                }
                _ = fi45
            }
            if (255 < 127) {
                var v57 bool = (100 == -38)
                fmt.Println(v57)
                for fi58 := 0; fi58 < 3; fi58++ {
                    v59 := []int{831}
                    fmt.Println(v57)
                    _ = fi58
                    _ = v59
                }
                _ = v57
            }
            var v60 string = ("a" + ("hello" + (("a" + "a") + (("foo bar" + "foo bar") + "hello"))))
            _ = v60
        } else {
            if (true && true) {
                v61 := []int{42}
                _ = v61
            } else {
                var v62 int32 = -2147483648
                var v63 uint32 = 42
                if false {
                    switch -154 {
                    case 0:
                        v62 = -2147483648
                    default:
                        var v64 uint32 = 2
                        _ = v64
                    }
                    var v65 uint32 = 1
                    _ = v65
                }
                _ = v62
                _ = v63
            }
        }
    }
    for fi66 := 0; fi66 < 2; fi66++ {
        _ = len(v1)
        if (188 < 1) {
            var v67 int64 = -1
            var v68 uint8 = 236
            _ = v67
            _ = v68
        } else {
            v69 := GenS3{F0: -87, F1: 2}
            _ = v69
        }
        _ = fi66
    }
    for fi70 := 0; fi70 < 6; fi70++ {
        v2 = 3.14159
        v2 = v2
        var v71 int32 = -245
        _ = fi70
        _ = v71
    }
    v72 := []int{}
    var v73 int = -2147483648
    switch (v73 * (v73 / ((492) | 1))) {
    case -1:
        if len(v72) > 0 {
            v72[0] = (v73 | v73)
        }
    default:
        var v74 int64 = 1
        _ = v74
    }
    switch -2147483648 {
    case 33:
        for fi75 := 0; fi75 < 2; fi75++ {
            v2 = ((v2 + v2) - (v2 / 0.1))
            _ = fi75
        }
    }
    var v76 uint16 = 724
    fmt.Println(v2, v76, v73)
    var v77 uint16 = v76
    for fi78 := 0; fi78 < 3; fi78++ {
        fmt.Println(v73)
        if false {
            switch (v73 & (^(v73))) {
            case 1:
                v73 -= v73
            default:
                fmt.Println(v73, v76, fi78)
            }
            v76 = v77
        }
        _ = fi78
    }
    for fi79 := 0; fi79 < 12; fi79++ {
        if true {
            for fi80 := 0; fi80 < 11; fi80++ {
                switch fi79 {
                case 251:
                    v73 = (v73 / ((fi79) | 1))
                    v72 = append(v72, fi79)
                default:
                    v73 = 9223372036854775807
                }
                _ = fi80
            }
            var v81 uint = 4294967295
            _ = v81
        } else {
            for fi82 := 0; fi82 < 8; fi82++ {
                v76 = (^(v77))
                _ = len(v1)
                var v83 int32 = -2147483648
                _ = fi82
                _ = v83
            }
        }
        if len(v72) >= 2 {
            v72 = v72[1:len(v72)]
        }
        _ = fi79
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v72:", len(v72), v72)
    fmt.Println("v73:", v73)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v72
    _ = v73
    _ = v76
    _ = v77
}
