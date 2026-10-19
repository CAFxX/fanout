package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint }
type GenS2 struct { F0 string; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 int; F1 int8; F2 string }
type GenS5 struct { F0 string; F1 int32; F2 int8 }
type GenS6 struct { F0 int64; F1 uint32; F2 uint }
type GenS7 struct { F0 uint64; F1 int32; F2 bool }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint16; F1 int; F2 int8 }
type GenS10 struct { F0 string }
func genFn0(p0 uint16, p1 int8) int64 {
    p0 = 785
    v6 := []int{9223372036854775807, 371, -2147483648}
    var v7 float64 = ((((0.1 * 1e10) + (0.1 * (1e10 / -2.25))) - (-2.25 - -2.25)) * 1.5)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = ("\n\t" + "\n\t")
    var v3 uint8 = 127
    v4 := make(map[string]uint32)
    var v5 int32 = 2147483647
    v2 = "foo bar"
    v2 = ((("hello" + "hello") + ("hello" + v2)) + (v2 + ("foo bar" + "hello")))
    v5 = -265
    v8 := GenS1{F0: 776, F1: "hello", F2: 18446744073709551615}
    var v9 uint8 = 42
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 uint8 = 77
            v9 -= 1
            var v13 int64 = 0
            _ = fi11
            _ = v12
            _ = v13
        }
        v14 := v1[65535]
        for fi15 := 0; fi15 < 5; fi15++ {
            v3 = v9
            _ = fi15
        }
        _ = fi10
        _ = v14
    }
    if (v9 < (v9 | (v9 & 32))) {
        v16 := GenS1{F0: -252, F1: ("foo bar" + "\n\t"), F2: 837}
        switch 1 {
        case 32768:
            for fi17 := 0; fi17 < 3; fi17++ {
                v18 := v1[974]
                v5 -= (v5 + v5)
                _ = fi17
                _ = v18
            }
        case -2147483648:
            v19 := GenS1{F0: -903, F1: (("\n\t" + v2) + "hello"), F2: 122}
            _ = v19
        case -373:
            for fi20 := 0; fi20 < 1; fi20++ {
                v2 = ("a" + v2)
                _ = fi20
            }
        default:
            if (!false) {
                delete(v1, 1)
                var v21 uint32 = 74
                v21 = (v21 - v21)
                _ = v21
            } else {
                var v22 uint8 = 28
                var v23 int8 = -1
                v24 := GenS1{F0: int16(v3), F1: (v2 + v2), F2: 65535}
                _ = v22
                _ = v23
                _ = v24
            }
        }
        if true {
            v25 := []int{2147483647}
            var v26 string = "a"
            var v27 int32 = v5
            _ = v25
            _ = v26
            _ = v27
        }
        _ = v16
    }
    switch v2 {
    case "":
        var v28 float32 = 3.14159
        _ = v28
    case "hello":
        fmt.Println(v3)
    default:
        v29 := []int{32767, 127}
        _ = v29
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        v5 -= v5
        var v31 int = fi30
        _ = fi30
        _ = v31
    }
    if false {
        var v32 uint8 = (v3 >> ((989) & 63))
        v33 := GenS1{F0: 42, F1: v2, F2: 749}
        _ = v32
        _ = v33
    } else {
        v34 := v1[285]
        v5 = v5
        for fi35 := 0; fi35 < 2; fi35++ {
            switch fi35 {
            case 128:
                v36 := GenS2{F0: ("a" + (v2 + "")), F1: 808, F2: 659}
                _ = v36
            case 32767:
                if (int8(v5) <= 0) {
                    fmt.Println(fi35, v2, v3)
                    v2 = ("hello" + v2)
                } else {
                    fmt.Println(v3)
                }
            default:
                fmt.Println(v2, v9, fi35)
            }
            for fi37 := 0; fi37 < 2; fi37++ {
                if v34 {
                    if (!(!((!((v34 && v34) || (v34 && (1e10 == 3.14159)))) && (384 >= 967)))) {
                        v9 += v3
                        switch 127 {
                        case 32768:
                            switch ((("foo bar" + (v2 + v2)) + v2) + "a") {
                            case "":
                                v38 := []int{}
                                _ = v38
                            case "hello":
                                switch (fi35 >> ((2) & 63)) {
                                case 42:
                                    fmt.Println(v9)
                                case -1:
                                    var v39 int32 = (-(v5))
                                    _ = v39
                                }
                            default:
                                if true {
                                    v9 = v3
                                    v40 := GenS3{F0: 32767, F1: (!(v34 && (-867 == 42)))}
                                    var v41 uint = 65535
                                    var v42 string = v2
                                    var v43 uint64 = 18446744073709551615
                                    v9 = v9
                                    v44 := v1[1]
                                    v45 := []int{}
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                } else {
                                    switch (fi35 >> ((65535) & 63)) {
                                    case 1:
                                        switch v2 {
                                        case "\n\t":
                                            v46 := GenS4{F0: 127, F1: -1, F2: ""}
                                            _ = v46
                                        case "a":
                                            v5 += v5
                                        case "hello":
                                            var v47 int32 = (v5 % ((v5) | 1))
                                            _ = v47
                                        }
                                    case 127:
                                        if v34 {
                                            v48 := GenS5{F0: (v2 + v2), F1: 1, F2: 42}
                                            var v49 uint16 = 42
                                            _ = v48
                                            _ = v49
                                        } else {
                                            var v50 uint8 = (v9 - (v9 >> ((42) & 63)))
                                            _ = v50
                                        }
                                        var v51 int = fi37
                                        _ = v51
                                    case 128:
                                        v3 = (v9 & v9)
                                    }
                                }
                            }
                        case 709:
                            if v34 {
                                if v34 {
                                    var v52 bool = v34
                                    _ = v52
                                }
                                v5 = v5
                                v53 := GenS6{F0: genFn0(32767, 100), F1: 734, F2: 0}
                                _ = v53
                            }
                            var v54 int = 42
                            _ = v54
                        case 32767:
                            if (!(!((((v2 + v2) + (v2 + v2)) < v2) && v34))) {
                                if v34 {
                                    switch (fi35 << ((162) & 63)) {
                                    case 32768:
                                        fmt.Println(v3, v5)
                                    default:
                                        switch "hello" {
                                        case "foo bar":
                                            fmt.Println(v34, v3)
                                        case "\n\t":
                                            v55 := []int{32767}
                                            v34 = v34
                                            if len(v55) >= 2 {
                                                v55 = v55[1:len(v55)]
                                            }
                                            _ = v55
                                        default:
                                            v56 := GenS7{F0: 39, F1: (^(v5)), F2: (v5 >= (v5 << ((891) & 63)))}
                                            _ = v56
                                        }
                                    }
                                    delete(v4, (v2 + "\n\t"))
                                    fmt.Println(fi35, fi37)
                                    fmt.Println(v5)
                                    if (!(2147483647 >= 146)) {
                                        v57 := GenS8{F0: v34}
                                        var v58 uint16 = 907
                                        _ = v57
                                        _ = v58
                                    } else {
                                        var v59 string = (("" + "a") + "")
                                        var v60 uint32 = 113
                                        _ = v59
                                        _ = v60
                                    }
                                } else {
                                    v2 = "hello"
                                    var v61 int = (fi37 + (fi35 | -876))
                                    _ = v61
                                }
                            }
                        }
                        v62 := []int{(fi37 | fi35), (fi35 * 836), fi35, fi37}
                        _ = v62
                    }
                } else {
                    v63 := []int{-175}
                    _ = v63
                }
                _ = fi37
            }
            _ = fi35
        }
        _ = v34
    }
    for fi64 := 0; fi64 < 5; fi64++ {
        fmt.Println(v2)
        _ = fi64
    }
    var v65 int8 = -128
    v66 := []int{127, -1, 769, 32768}
    if false {
        v67 := GenS5{F0: (v2 + "\n\t"), F1: (v5 >> ((255) & 63)), F2: v65}
        for fi68 := 0; fi68 < 18; fi68++ {
            if (false && ((v65 + (v65 % (((v65 | v65)) | 1))) != v65)) {
                v4[(v2 + v2)] = 65535
                for fi69 := 0; fi69 < 14; fi69++ {
                    v5 = 42
                    _ = fi69
                }
            } else {
                v70 := GenS9{F0: 255, F1: (fi68 & (^(fi68))), F2: (v65 << ((42) & 63))}
                var v71 uint32 = 195
                for fi72 := 0; fi72 < 12; fi72++ {
                    var v73 float32 = float32(0.1)
                    switch (fi68 % (((fi72 / ((fi72) | 1))) | 1)) {
                    case 9223372036854775807:
                        if false {
                            var v74 uint32 = v71
                            _ = v74
                        }
                    case -809:
                        if (fi72 == fi68) {
                            if true {
                                switch (fi68 >> ((859) & 63)) {
                                case 2147483647:
                                    v71 = v71
                                case 32767:
                                    v75 := GenS10{F0: v2}
                                    _ = v75
                                case 42:
                                    v9 = (v9 >> ((255) & 63))
                                }
                            }
                            if false {
                                var v76 uint = 395
                                switch v2 {
                                case "a":
                                    v2 = "\n\t"
                                }
                                _ = v76
                            }
                        }
                    case 32767:
                        v3 = v3
                    default:
                        if len(v66) >= 2 {
                            v66 = v66[1:len(v66)]
                        }
                    }
                    fmt.Println(v73, v73, v71)
                    _ = fi72
                    _ = v73
                }
                _ = v70
                _ = v71
            }
            if len(v66) >= 2 {
                v66 = v66[1:len(v66)]
            }
            v3 -= 1
            _ = fi68
        }
        _ = v67
    } else {
        var v77 int64 = int64(v3)
        _ = v77
    }
    if true {
        switch 196 {
        case 128:
            switch 558 {
            case -2147483648:
                v5 = v5
            default:
                fmt.Println(v5)
            }
        case 32768:
            var v78 string = (v2 + ("foo bar" + "hello"))
            _ = v78
        }
        if false {
            fmt.Println(v5, v3, v5)
            switch (v2 + "hello") {
            case "hello":
                switch 1 {
                case 127:
                    fmt.Println(v9, v2, v3)
                case 32767:
                    fmt.Println(v3, v5, v65)
                }
            default:
                var v79 uint16 = 663
                _ = v79
            }
            for fi80 := 0; fi80 < 2; fi80++ {
                fmt.Println(fi80, v2, v5)
                _ = fi80
            }
        } else {
            v9 = v3
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    fmt.Println("slicelen v66:", len(v66), v66)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
    _ = v8
    _ = v9
}
