package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint64 }
type GenS2 struct { F0 uint; F1 int16; F2 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 bool; F2 uint8 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int64; F1 int16; F2 int16 }
func genFn0(p0 int16) bool {
    var v6 bool = false
    v6 = v6
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(v6, fi7, p0)
        var v8 int16 = 1000
        v9 := []int{-341}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v6
    return v6
}

func main() {
    var v1 bool = false
    var v2 int16 = 1000
    v3 := []string{"\n\t", ("a" + "a"), "", ("a" + ("" + ""))}
    var v4 int32 = -1
    var v5 int32 = 1
    v10 := GenS1{F0: -9223372036854775808, F1: "", F2: 1}
    var v11 int64 = 38
    if v1 {
        for fi12 := 0; fi12 < 20; fi12++ {
            v13 := GenS1{F0: 0, F1: "\n\t", F2: 2}
            _ = fi12
            _ = v13
        }
        fmt.Println(v2, v2, v1)
        v2 -= (v2 | v2)
    } else {
        v3 = append(v3, "foo bar")
        for fi14 := 0; fi14 < 16; fi14++ {
            v15 := GenS1{F0: (v11 + v11), F1: "a", F2: 1}
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 10; fi16++ {
            if (225 < 2) {
                switch "foo bar" {
                case "a":
                    v2 = (v2 % ((v2) | 1))
                case "":
                    fmt.Println(fi16)
                default:
                    if v1 {
                        v5 = (v5 | v4)
                        if (!(1.5 != 1e10)) {
                            switch "" {
                            case "a":
                                v1 = v1
                            case "":
                                for fi17 := 0; fi17 < 7; fi17++ {
                                    if false {
                                        var v18 int64 = v11
                                        fmt.Println(fi17, v1, v5)
                                        fmt.Println(v18, fi17)
                                        _ = v18
                                    } else {
                                        v19 := GenS1{F0: -733, F1: "foo bar", F2: 966}
                                        var v20 bool = v1
                                        _ = v19
                                        _ = v20
                                    }
                                    _ = fi17
                                }
                            }
                            fmt.Println(fi16, v2)
                            if v1 {
                                fmt.Println(v11, fi16)
                                var v21 uint16 = 255
                                _ = v21
                            }
                        }
                    }
                }
            } else {
                switch "a" {
                case "a":
                    for fi22 := 0; fi22 < 10; fi22++ {
                        fmt.Println(v4, v1)
                        if (true || false) {
                            if (0 < 65535) {
                                var v23 bool = true
                                v24 := GenS1{F0: (v11 & v11), F1: ("" + "\n\t"), F2: 0}
                                var v25 int16 = v2
                                _ = v23
                                _ = v24
                                _ = v25
                            } else {
                                if (128 != 255) {
                                    var v26 int8 = -128
                                    if v1 {
                                        var v27 uint32 = 753
                                        _ = v27
                                    } else {
                                        v1 = (false || v1)
                                        v1 = (true || v1)
                                        v4 -= (v4 & (-(v5)))
                                    }
                                    var v28 uint8 = 30
                                    fmt.Println(v5, v28, v26)
                                    var v29 uint16 = 0
                                    _ = v26
                                    _ = v28
                                    _ = v29
                                } else {
                                    var v30 uint16 = 42
                                    _ = v30
                                }
                                v4 -= (v4 / ((v4) | 1))
                                switch fi16 {
                                case 0:
                                    v2 = 1
                                case 2147483647:
                                    v2 = (v2 % (((v2 << ((255) & 63))) | 1))
                                }
                                v11 -= v11
                                var v31 uint64 = 470
                                var v32 int64 = v11
                                _ = v31
                                _ = v32
                            }
                            if (v1 || true) {
                                var v33 string = "foo bar"
                                v33 = "hello"
                                v11 -= v11
                                _ = v33
                            }
                        } else {
                            v34 := GenS2{F0: 481, F1: -840, F2: v1}
                            var v35 int = (fi16 << ((18446744073709551615) & 63))
                            var v36 bool = v1
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                        var v37 int64 = (-(v11))
                        _ = fi22
                        _ = v37
                    }
                case "\n\t":
                    switch 128 {
                    case 32767:
                        var v38 uint8 = 255
                        _ = v38
                    case 0:
                        switch (fi16 | (fi16 | 506)) {
                        case 2147483647:
                            fmt.Println(v11, v1, v2)
                        case 174:
                            fmt.Println(v1)
                        case 4294967295:
                            v39 := GenS3{F0: (v2 + v2)}
                            _ = v39
                        }
                    default:
                        for fi40 := 0; fi40 < 1; fi40++ {
                            v2 = 0
                            switch "foo bar" {
                            case "\n\t":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case "foo bar":
                                if len(v3) > 0 {
                                    v3[0] = ("hello" + ("\n\t" + ("hello" + (("a" + "a") + ""))))
                                }
                                v41 := GenS2{F0: 949, F1: (v2 | (v2 >> ((65535) & 63))), F2: (2 == 1)}
                                switch (("foo bar" + ("" + "hello")) + ("a" + "a")) {
                                case "foo bar":
                                    v2 += 42
                                case "\n\t":
                                    var v42 uint8 = 2
                                    _ = v42
                                case "a":
                                    if v1 {
                                        fmt.Println(v4)
                                        switch (("a" + "hello") + "") {
                                        case "foo bar":
                                            v4 = (-(v4))
                                        }
                                        var v43 float32 = 3.14159
                                        _ = v43
                                    } else {
                                        var v44 int16 = (^(v2))
                                        _ = v44
                                    }
                                default:
                                    var v45 uint64 = 4294967295
                                    _ = v45
                                }
                                _ = v41
                            case "hello":
                                v4 -= 42
                            }
                            switch fi40 {
                            case 1:
                                v4 = 312
                            }
                            var v46 int32 = (v5 * v4)
                            _ = fi40
                            _ = v46
                        }
                    }
                case "":
                    for fi47 := 0; fi47 < 20; fi47++ {
                        v11 = (v11 - (^(v11)))
                        var v48 string = (("\n\t" + "\n\t") + ((("\n\t" + ("" + "a")) + "") + "foo bar"))
                        if (0 > 372) {
                            if (false && true) {
                                var v49 int16 = (v2 | v2)
                                _ = v49
                            } else {
                                v4 -= 42
                                var v50 string = ("" + v48)
                                var v51 uint16 = 412
                                v50 = "a"
                                var v52 int = (fi47 % ((fi16) | 1))
                                _ = v50
                                _ = v51
                                _ = v52
                            }
                        } else {
                            if true {
                                fmt.Println(v2, fi16, fi16)
                                switch ("\n\t" + v48) {
                                case "a":
                                    var v53 uint16 = 65535
                                    _ = v53
                                case "":
                                    fmt.Println(v4)
                                    v2 += 32767
                                    var v54 uint64 = 1
                                    _ = v54
                                case "\n\t":
                                    if true {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                default:
                                    if (false || true) {
                                        v3 = append(v3, (v48 + v48))
                                        v55 := GenS1{F0: 935, F1: (((v48 + ((v48 + "") + "a")) + "foo bar") + v48), F2: 1}
                                        switch v48 {
                                        case "\n\t":
                                            var v56 uint16 = 0
                                            _ = v56
                                        case "hello":
                                            if len(v3) > 0 {
                                                v3[0] = v48
                                            }
                                        case "a":
                                            v3 = append(v3, v48)
                                        default:
                                            if (539 > 1) {
                                                var v57 string = ((v48 + v48) + v48)
                                                var v58 int16 = int16(fi16)
                                                v3 = append(v3, v48)
                                                _ = v57
                                                _ = v58
                                            } else {
                                                var v59 uint64 = 640
                                                var v60 uint16 = 65535
                                                switch ("hello" + v48) {
                                                case "hello":
                                                    v61 := GenS3{F0: (v2 >> ((255) & 63))}
                                                    _ = v61
                                                }
                                                _ = v59
                                                _ = v60
                                            }
                                        }
                                        _ = v55
                                    }
                                }
                            } else {
                                if false {
                                    if (!(v2 != v2)) {
                                        _ = len(v3)
                                        fmt.Println(fi16, v11, fi16)
                                        switch (v48 + v48) {
                                        case "foo bar":
                                            var v62 uint32 = 80
                                            _ = v62
                                        case "":
                                            var v63 uint = 4294967295
                                            _ = v63
                                        default:
                                            if v1 {
                                                v11 += v11
                                            }
                                        }
                                    } else {
                                        v2 += (v2 >> ((2) & 63))
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                v64 := GenS3{F0: v2}
                                v2 -= v2
                                _ = v64
                            }
                            fmt.Println(v5)
                        }
                        v2 -= v2
                        _ = fi47
                        _ = v48
                    }
                default:
                    if (((fi16 <= fi16) && (65 > 1)) && v1) {
                        for fi65 := 0; fi65 < 2; fi65++ {
                            v5 += (v5 - (v5 * (-(v5))))
                            _ = fi65
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v4 = -196
                    }
                }
                switch "\n\t" {
                case "":
                    var v66 float32 = float32(-2.25)
                    _ = v66
                case "hello":
                    var v67 uint16 = 556
                    _ = v67
                default:
                    for fi68 := 0; fi68 < 16; fi68++ {
                        v5 -= v4
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if len(v3) > 0 {
                            v3[0] = ("" + "a")
                        }
                        _ = fi68
                    }
                }
            }
            _ = fi16
        }
    }
    v69 := GenS4{F0: v4, F1: (true && (v4 != (-(v5)))), F2: 255}
    v1 = v1
    for fi70 := 0; fi70 < 15; fi70++ {
        v2 -= (v2 % (((v2 & v2)) | 1))
        var v71 float32 = (((3.14159 + 0.1) - -2.25) / 1e10)
        _ = fi70
        _ = v71
    }
    for fi72 := 0; fi72 < 2; fi72++ {
        var v73 float32 = 0.1
        for fi74 := 0; fi74 < 5; fi74++ {
            v2 = v2
            v4 += v5
            if v1 {
                v5 -= v5
                var v75 uint8 = 135
                _ = v75
            }
            var v76 float64 = -2.25
            _ = fi74
            _ = v76
        }
        _ = fi72
        _ = v73
    }
    v3 = append(v3, "foo bar")
    var v77 int64 = v11
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    var v78 int16 = (v2 * 1000)
    if false {
        var v79 float32 = 3.14159
        var v80 uint16 = 42
        v2 = (v78 + (v78 << ((42) & 63)))
        _ = v79
        _ = v80
    } else {
        var v81 int16 = -947
        _ = v81
    }
    var v82 float32 = (float32((float64((3.14159 + -2.25)) + (float64(1e10) * ((0.0 * (1.5 * (1.5 / (3.14159 / float64(0.1))))) * 0.1)))) * 1e10)
    _ = len(v3)
    if v1 {
        for fi83 := 0; fi83 < 2; fi83++ {
            v84 := GenS5{F0: 1, F1: 1}
            v85 := GenS5{F0: (fi83 + (fi83 * int(v2))), F1: 42}
            v3 = append(v3, "a")
            _ = fi83
            _ = v84
            _ = v85
        }
        switch 32768 {
        case 530:
            switch "foo bar" {
            case "":
                for fi86 := 0; fi86 < 1; fi86++ {
                    v77 = v11
                    fmt.Println(v5, v2, v4)
                    _ = fi86
                }
            }
        case -2147483648:
            switch -1 {
            case -2147483648:
                var v87 uint32 = 2147483648
                _ = v87
            case 714:
                v88 := GenS6{F0: (v1 && ((-(v77)) >= v77))}
                _ = v88
            case -1:
                switch "a" {
                case "hello":
                    if v1 {
                        v5 = v5
                        var v89 uint64 = 0
                        _ = v89
                    } else {
                        if v1 {
                            var v90 uint32 = 2
                            _ = v90
                        }
                        var v91 float64 = ((1.5 - ((-2.25 - 1.5) * ((3.14159 * 1.5) - 1.5))) + 0.0)
                        _ = v91
                    }
                case "a":
                    v92 := GenS7{F0: v11, F1: 393, F2: 42}
                    _ = v92
                }
            }
        default:
            switch "\n\t" {
            case "":
                v4 += v4
            case "hello":
                for fi93 := 0; fi93 < 19; fi93++ {
                    v1 = v1
                    v78 = (v78 - int16(v77))
                    _ = fi93
                }
            }
        }
        v1 = (299 <= (v2 >> ((1) & 63)))
    } else {
        var v94 float32 = v82
        for fi95 := 0; fi95 < 20; fi95++ {
            var v96 int64 = v77
            var v97 string = "foo bar"
            _ = fi95
            _ = v96
            _ = v97
        }
        _ = v94
    }
    if len(v3) > 0 {
        v3[0] = ""
    }
    fmt.Println(v82)
    for fi98 := 0; fi98 < 17; fi98++ {
        if true {
            for fi99 := 0; fi99 < 10; fi99++ {
                switch (fi98 % ((fi99) | 1)) {
                case 292:
                    var v100 int = int(v5)
                    _ = v100
                case 50:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case 0:
                    v11 = (v77 / (((v77 + v11)) | 1))
                }
                _ = fi99
            }
            var v101 int = fi98
            _ = len(v3)
            _ = v101
        }
        var v102 uint8 = 1
        var v103 uint16 = 1
        _ = fi98
        _ = v102
        _ = v103
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v69
    _ = v77
    _ = v78
    _ = v82
}
