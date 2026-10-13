package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 bool }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint; F1 int8; F2 int64 }
type GenS5 struct { F0 int64; F1 uint32; F2 int16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int; F1 int32 }

func main() {
    var v1 int = 1
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    var v4 float32 = 0.0
    v5 := GenS1{F0: 903, F1: 42, F2: true}
    v6 := []int{(v1 * 0), -479, v1}
    var v7 int64 = -1
    var v8 int64 = v7
    var v9 float64 = -2.25
    switch v1 {
    case 9223372036854775807:
        fmt.Println(v9, v1, v7)
    case 283:
        for fi10 := 0; fi10 < 15; fi10++ {
            if (128 == 2) {
                for fi11 := 0; fi11 < 17; fi11++ {
                    v7 = v7
                    _ = fi11
                }
                var v12 int32 = -2147483648
                fmt.Println(v7)
                _ = v12
            } else {
                var v13 float32 = (((0.1 - (v4 + float32(v9))) - (((0.1 / 1.0) / v4) + (float32(v9) / 1.5))) - v4)
                _ = v13
            }
            v14 := GenS2{F0: (^(v1)), F1: 224}
            _ = fi10
            _ = v14
        }
    case 4294967295:
        fmt.Println(v7)
    default:
        if ((v7 != int64(v1)) || false) {
            if false {
                var v15 uint16 = 0
                if true {
                    v7 -= (v8 & 0)
                    fmt.Println(v7)
                    var v16 bool = (42 < -32768)
                    _ = v16
                }
                _ = v15
            } else {
                var v17 int8 = -128
                v1 -= v1
                fmt.Println(v17, v4)
                _ = v17
            }
        } else {
            var v18 int64 = v8
            for fi19 := 0; fi19 < 4; fi19++ {
                for fi20 := 0; fi20 < 20; fi20++ {
                    var v21 uint32 = 2
                    var v22 int = -2147483648
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                for fi23 := 0; fi23 < 20; fi23++ {
                    switch "foo bar" {
                    case "\n\t":
                        var v24 float64 = v9
                        _ = v24
                    }
                    switch 697 {
                    case 0:
                        if false {
                            v7 = v7
                            var v25 int = v1
                            if (4294967295 >= 679) {
                                fmt.Println(v4)
                                var v26 uint8 = 2
                                if false {
                                    v4 += 0.0
                                    if (!true) {
                                        var v27 uint = 2
                                        _ = v27
                                    }
                                    v28 := GenS3{F0: "", F1: 18446744073709551615}
                                    _ = v28
                                } else {
                                    if (true || false) {
                                        if false {
                                            v29 := GenS4{F0: 12, F1: 41, F2: v8}
                                            v30 := GenS4{F0: 0, F1: -9, F2: 679}
                                            switch (fi19 << ((2) & 63)) {
                                            case -1:
                                                v31 := GenS4{F0: 709, F1: 100, F2: (v18 << ((65535) & 63))}
                                                _ = v31
                                            }
                                            _ = v29
                                            _ = v30
                                        } else {
                                            switch fi19 {
                                            case -429:
                                                var v32 int64 = -9223372036854775808
                                                var v33 int16 = -32768
                                                v7 += -9223372036854775808
                                                var v34 int8 = -25
                                                v35 := GenS3{F0: ("" + ""), F1: 2}
                                                _ = v32
                                                _ = v33
                                                _ = v34
                                                _ = v35
                                            }
                                            var v36 string = (("\n\t" + (("a" + "hello") + "a")) + "foo bar")
                                            switch v36 {
                                            case "foo bar":
                                                var v37 uint64 = 0
                                                _ = v37
                                            }
                                            _ = v36
                                        }
                                        v38 := GenS5{F0: (v7 - -667), F1: 4294967295, F2: -1}
                                        fmt.Println(v25, fi19)
                                        _ = v38
                                    }
                                    var v39 bool = (("foo bar" + (("a" + "hello") + ((("\n\t" + ("\n\t" + "hello")) + "\n\t") + "hello"))) == ("hello" + "a"))
                                    _ = v39
                                }
                                v25 = 128
                                var v40 bool = (!(v8 <= (-(v8))))
                                v26 = (v26 - (v26 - (v26 / ((1) | 1))))
                                v41 := GenS4{F0: 1, F1: -32, F2: (v18 / ((v18) | 1))}
                                _ = v26
                                _ = v40
                                _ = v41
                            } else {
                                if ((!(1 == 2)) && true) {
                                    var v42 uint16 = 42
                                    var v43 uint8 = 0
                                    _ = len(v3)
                                    v7 = v18
                                    var v44 uint16 = (v42 ^ v42)
                                    var v45 uint8 = 42
                                    if len(v6) > 0 {
                                        v6[0] = 128
                                    }
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                }
                            }
                            _ = v25
                        } else {
                            switch fi23 {
                            case 42:
                                v18 = v7
                            case 1:
                                v46 := GenS3{F0: "", F1: 42}
                                _ = v46
                            case 128:
                                var v47 uint32 = 4294967295
                                _ = v47
                            }
                        }
                    case 127:
                        var v48 int32 = 2147483647
                        if ((v48 ^ (v48 + v48)) != (v48 & 42)) {
                            v49 := GenS6{F0: 1}
                            if ((v9 - v9) <= float64((v4 * 1e10))) {
                                var v50 int = fi23
                                if ((!false) && false) {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                } else {
                                    switch (v1 + -704) {
                                    case 4294967295:
                                        if false {
                                            var v51 uint8 = 1
                                            if true {
                                                var v52 float64 = 0.1
                                                fmt.Println(v51, v9, v4)
                                                if true {
                                                    var v53 uint64 = 808
                                                    var v54 float64 = v9
                                                    _ = v53
                                                    _ = v54
                                                }
                                                _ = v52
                                            }
                                            _ = v51
                                        }
                                    default:
                                        var v55 int = (fi19 - (-(fi19)))
                                        var v56 uint = 65535
                                        _ = v55
                                        _ = v56
                                    }
                                }
                                var v57 int8 = 117
                                _ = v50
                                _ = v57
                            }
                            _ = v49
                        } else {
                            var v58 int8 = 0
                            if (false || true) {
                                v58 -= int8(v7)
                                v8 += v18
                                v59 := GenS7{F0: (fi23 | 32768), F1: v48}
                                _ = v59
                            }
                            _ = v58
                        }
                        _ = v48
                    case -114:
                        var v60 uint64 = 0
                        _ = v60
                    }
                    var v61 uint16 = 27
                    _ = fi23
                    _ = v61
                }
                switch "foo bar" {
                case "foo bar":
                    var v62 int32 = 878
                    _ = v62
                case "\n\t":
                    var v63 int8 = 20
                    _ = v63
                case "":
                    if ((-1 < 19) && (!(710 == 4294967295))) {
                        if ((v18 % (((v18 - v8)) | 1)) <= 1) {
                            v64 := v2[765]
                            switch "hello" {
                            case "":
                                if ((float32(v9) - float32(3.14159)) == (v4 - v4)) {
                                    var v65 int32 = 42
                                    v66 := GenS2{F0: 42, F1: 255}
                                    _ = v65
                                    _ = v66
                                } else {
                                    switch v1 {
                                    case 9223372036854775807:
                                        v67 := GenS7{F0: v1, F1: 0}
                                        _ = v67
                                    default:
                                        _ = len(v6)
                                    }
                                    var v68 bool = true
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = v68
                                }
                            }
                            var v69 int8 = 1
                            _ = v64
                            _ = v69
                        }
                    } else {
                        var v70 int64 = (v7 | 870)
                        _ = v70
                    }
                }
                _ = fi19
            }
            _ = v18
        }
    }
    if true {
        v7 = v7
    }
    var v71 uint16 = 65535
    if len(v6) > 0 {
        v6[0] = (v1 << ((0) & 63))
    }
    v8 += 693
    _ = len(v6)
    v6 = append(v6, v1)
    switch (v1 << ((305) & 63)) {
    case -649:
        for fi72 := 0; fi72 < 17; fi72++ {
            v9 += 0.1
            fmt.Println(v1, v9, fi72)
            _ = fi72
        }
    case 42:
        switch "" {
        case "\n\t":
            v73 := GenS6{F0: 880}
            _ = v73
        }
    case 32767:
        for fi74 := 0; fi74 < 11; fi74++ {
            var v75 int = 1
            var v76 float64 = v9
            _ = fi74
            _ = v75
            _ = v76
        }
    default:
        var v77 string = "foo bar"
        _ = v77
    }
    switch ((("foo bar" + "a") + "") + "foo bar") {
    case "\n\t":
        switch v1 {
        case 9223372036854775807:
            v9 = v9
        }
    default:
        v9 -= (0.1 * (v9 / (v9 + v9)))
    }
    v1 -= (v1 << ((18446744073709551615) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v71:", v71)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v71
    _ = v8
    _ = v9
}
