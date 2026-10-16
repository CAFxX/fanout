package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint; F1 uint; F2 uint64 }

func main() {
    var v1 uint32 = 4294967295
    var v2 bool = ((!(42 >= -1)) && (65535 <= 1))
    v3 := []uint32{v1, v1, v1, 1}
    var v4 float32 = 0.0
    v5 := make(map[string]uint32)
    var v6 uint32 = v1
    if ((v1 * 42) < (v1 % ((v1) | 1))) {
        fmt.Println(v6, v2)
        v7 := GenS1{F0: v2}
        fmt.Println(v6, v4, v1)
        _ = v7
    }
    var v8 bool = true
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 -= v1
    v9 := v5[("hello" + "hello")]
    if ((v6 * v1) < 1) {
        for fi10 := 0; fi10 < 16; fi10++ {
            if (!v8) {
                v11 := v5[((("hello" + "\n\t") + "\n\t") + ("\n\t" + ("a" + "\n\t")))]
                for fi12 := 0; fi12 < 6; fi12++ {
                    v1 = (v1 | v1)
                    if false {
                        switch 32768 {
                        case 4294967295:
                            var v13 uint16 = 32767
                            _ = v13
                        case -2147483648:
                            fmt.Println(v6, v1)
                        case 127:
                            switch -999 {
                            case -2147483648:
                                var v14 uint = 4294967295
                                _ = v14
                            case -369:
                                v6 = (v6 - 1)
                            }
                        default:
                            if (42 >= 128) {
                                v3 = append(v3, (v1 * (v1 / (((v1 * 65535)) | 1))))
                                var v15 int16 = -1
                                _ = v15
                            }
                            switch -2147483648 {
                            case 1:
                                var v16 string = "a"
                                _ = v16
                            case -968:
                                v11 = (v11 - (^(v11)))
                            default:
                                var v17 int16 = 0
                                _ = v17
                            }
                            switch "\n\t" {
                            case "\n\t":
                                var v18 uint32 = v1
                                _ = v18
                            case "":
                                var v19 float32 = v4
                                _ = v19
                            case "foo bar":
                                _ = len(v5)
                            default:
                                v8 = false
                            }
                        }
                        if (!(-543 >= -113)) {
                            v2 = (!(fi10 >= 0))
                            v11 -= v6
                            var v20 int64 = 782
                            _ = v20
                        }
                        v1 = 0
                    } else {
                        fmt.Println(fi12)
                        v6 += v6
                        var v21 int16 = -946
                        _ = v21
                    }
                    v8 = (v2 && v2)
                    var v22 int64 = 1
                    _ = fi12
                    _ = v22
                }
                _ = len(v5)
                _ = v11
            }
            v23 := GenS2{F0: 255, F1: 792}
            v24 := GenS2{F0: 255, F1: 364}
            _ = fi10
            _ = v23
            _ = v24
        }
        var v25 float64 = 1.5
        _ = v25
    } else {
        if len(v3) > 0 {
            v3[0] = (^(v6))
        }
        v9 = (v6 - (v1 % (((v1 & (^(v6)))) | 1)))
        v26 := GenS1{F0: (1 <= 2147483647)}
        _ = v26
    }
    var v27 uint16 = 255
    v6 = v1
    var v28 int = 127
    for fi29 := 0; fi29 < 14; fi29++ {
        switch ((("" + ("a" + "")) + "\n\t") + (("\n\t" + "foo bar") + "")) {
        case "\n\t":
            if v8 {
                var v30 uint16 = v27
                if false {
                    fmt.Println(v1, v1, fi29)
                    v31 := GenS2{F0: 0, F1: v30}
                    _ = v31
                } else {
                    var v32 int16 = 42
                    v33 := GenS3{F0: 42, F1: 1, F2: 805}
                    _ = v32
                    _ = v33
                }
                for fi34 := 0; fi34 < 10; fi34++ {
                    if v8 {
                        v35 := GenS1{F0: (v27 <= uint16(v1))}
                        v36 := GenS4{F0: 626, F1: 4294967295, F2: 0}
                        v2 = ((v4 * v4) > float32((0.0 * 0.1)))
                        v9 += 2147483647
                        if true {
                            var v37 int64 = 946
                            _ = v37
                        } else {
                            _ = len(v3)
                            if true {
                                var v38 float32 = float32(1e10)
                                _ = v38
                            }
                            var v39 int32 = 355
                            _ = v39
                        }
                        _ = v35
                        _ = v36
                    }
                    fmt.Println(v30)
                    _ = len(v3)
                    v6 = (v1 >> ((684) & 63))
                    v27 -= v27
                    _ = fi34
                }
                _ = v30
            }
        case "foo bar":
            v4 = 1e10
        }
        for fi40 := 0; fi40 < 7; fi40++ {
            if ((^(v9)) < 2147483648) {
                switch "\n\t" {
                case "hello":
                    fmt.Println(v9, fi40, v27)
                case "foo bar":
                    if v8 {
                        if v8 {
                            v27 -= (^(v27))
                            var v41 uint64 = uint64(v27)
                            _ = v41
                        }
                        var v42 uint = 0
                        var v43 uint32 = (v6 * 266)
                        _ = v42
                        _ = v43
                    } else {
                        if true {
                            v44 := GenS2{F0: 128, F1: 42}
                            if (int32(v9) >= 0) {
                                switch (^(fi40)) {
                                case -1:
                                    delete(v5, ("foo bar" + "a"))
                                default:
                                    if (42 >= -1) {
                                        switch "hello" {
                                        case "":
                                            v45 := GenS3{F0: 65535, F1: 1, F2: v9}
                                            _ = v45
                                        default:
                                            if ((((2 <= 18446744073709551615) || ((v8 && (!v2)) && ((v8 || (v8 && ((1000 > -1) || false))) || (!(!(!false)))))) || ((v4 + float32((3.14159 * (0.1 / float64(v4))))) > 1.5)) && (664 > 42)) {
                                                v2 = (!true)
                                                if false {
                                                    v6 -= v9
                                                    _ = len(v5)
                                                }
                                            } else {
                                                v27 -= (v27 << ((9) & 63))
                                                if v2 {
                                                    switch 1 {
                                                    case 9223372036854775807:
                                                        v9 += 621
                                                    default:
                                                        switch ("a" + "foo bar") {
                                                        case "":
                                                            delete(v5, ("\n\t" + (((("" + "") + "a") + "") + ("a" + "hello"))))
                                                        case "foo bar":
                                                            switch (("foo bar" + "foo bar") + ("foo bar" + (("hello" + "\n\t") + ("foo bar" + "a")))) {
                                                            case "hello":
                                                                var v46 int = v28
                                                                _ = v46
                                                            case "":
                                                                v47 := GenS3{F0: v9, F1: 840, F2: v9}
                                                                _ = v47
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    var v48 uint = 2
                                                    _ = v48
                                                }
                                            }
                                        }
                                        switch "hello" {
                                        case "a":
                                            var v49 int64 = -302
                                            _ = v49
                                        case "\n\t":
                                            var v50 uint8 = 141
                                            _ = v50
                                        default:
                                            v8 = (v2 || v2)
                                        }
                                    }
                                }
                                var v51 int = -9223372036854775808
                                if true {
                                    var v52 uint8 = 0
                                    v27 = uint16(v9)
                                    _ = v52
                                } else {
                                    var v53 int32 = 2147483647
                                    var v54 uint32 = v1
                                    _ = v53
                                    _ = v54
                                }
                                fmt.Println(v9, v8)
                                _ = v51
                            } else {
                                fmt.Println(fi29, fi29, v6)
                            }
                            var v55 uint64 = 4294967295
                            _ = v44
                            _ = v55
                        }
                        v56 := GenS3{F0: 517, F1: 214, F2: v1}
                        _ = v56
                    }
                default:
                    var v57 int32 = -2147483648
                    _ = v57
                }
                if (!(v4 > 0.1)) {
                    var v58 int64 = 9223372036854775807
                    _ = v58
                }
                var v59 uint = 0
                _ = v59
            } else {
                var v60 uint8 = 52
                if v2 {
                    if ((!(0 == 0)) && (!(true && true))) {
                        var v61 uint = 0
                        _ = len(v3)
                        _ = v61
                    } else {
                        var v62 int = v28
                        _ = v62
                    }
                    if (-2.25 != (3.14159 + 1e10)) {
                        fmt.Println(v27, v8)
                        v6 += v9
                        if true {
                            var v63 int16 = 0
                            if len(v3) > 0 {
                                v3[0] = v1
                            }
                            var v64 int16 = (-(v63))
                            _ = v63
                            _ = v64
                        } else {
                            var v65 uint8 = uint8(v6)
                            var v66 int32 = -2147483648
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v60 -= (^(v60))
                            _ = v65
                            _ = v66
                        }
                        switch "a" {
                        case "":
                            var v67 string = "a"
                            _ = v67
                        case "\n\t":
                            fmt.Println(v9, fi40, fi29)
                        }
                    }
                    var v68 int16 = 1
                    _ = v68
                }
                var v69 uint32 = (v9 >> ((6) & 63))
                v2 = ((1.5 - 0.0) == 3.14159)
                var v70 uint64 = 27
                if (!v2) {
                    delete(v5, "\n\t")
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v2 = (v8 || ((1 < 255) && (443 <= (^(v70)))))
                    v71 := GenS2{F0: 2, F1: v27}
                    _ = v71
                }
                var v72 int = (fi29 >> ((588) & 63))
                _ = v60
                _ = v69
                _ = v70
                _ = v72
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi40
        }
        _ = fi29
    }
    var v73 uint = 18446744073709551615
    var v74 int16 = -19
    fmt.Println(v28, v8)
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v74
    _ = v8
    _ = v9
}
