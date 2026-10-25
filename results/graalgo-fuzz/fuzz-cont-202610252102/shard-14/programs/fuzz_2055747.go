package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 int16; F2 uint }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int16; F1 bool }

func main() {
    var v1 uint8 = 127
    v2 := []string{"\n\t"}
    var v3 int8 = 125
    v4 := GenS1{F0: 126}
    var v5 uint = 2
    switch 878 {
    case 2147483647:
        var v6 float32 = 1e10
        _ = v6
    case 42:
        if true {
            v3 = v3
        }
    default:
        for fi7 := 0; fi7 < 3; fi7++ {
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 int16 = -1
                v10 := GenS1{F0: (^(v5))}
                v1 -= (v1 << ((v5) & 63))
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
    }
    v1 -= v1
    var v11 float32 = (-2.25 * (0.0 - (3.14159 + float32(((((0.0 - 0.1) + ((0.0 + (1e10 - float64(0.0))) + 0.1)) * 0.0) / 1.0)))))
    var v12 string = "foo bar"
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 uint64 = 4294967295
        for fi15 := 0; fi15 < 7; fi15++ {
            switch -2147483648 {
            case 4294967295:
                v16 := GenS2{F0: (v3 >> ((v5) & 63)), F1: 32767, F2: 609}
                var v17 float32 = -2.25
                var v18 uint64 = (v14 ^ (v14 & (v14 ^ (^(v14)))))
                _ = v16
                _ = v17
                _ = v18
            case -710:
                v19 := GenS2{F0: v3, F1: 574, F2: 1}
                _ = v19
            case 42:
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v11 -= v11
                }
            default:
                v2 = append(v2, ("foo bar" + (v12 + "foo bar")))
            }
            v1 = v1
            var v20 string = "a"
            switch v20 {
            case "foo bar":
                if ((fi15 / (((fi13 - (fi13 | fi15))) | 1)) == (fi13 / ((fi15) | 1))) {
                    v14 -= v14
                }
                fmt.Println(v14)
                v21 := GenS1{F0: 777}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = len(v2)
                _ = v21
            case "a":
                if true {
                    if (false && (int(v3) == (-(fi15)))) {
                        v20 = (v20 + "")
                        var v22 int8 = v3
                        _ = len(v2)
                        v23 := GenS3{F0: (^(v1)), F1: 2147483647, F2: 0}
                        var v24 int32 = -824
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    if (32767 > (fi13 ^ (fi13 & int(v14)))) {
                        switch (((v20 + v20) + "\n\t") + v12) {
                        case "a":
                            if true {
                                v25 := GenS4{F0: v20, F1: (fi15 - int(v14))}
                                if false {
                                    var v26 int64 = 42
                                    switch fi13 {
                                    case -733:
                                        v11 -= -2.25
                                    case 1:
                                        v20 = ("\n\t" + v12)
                                        var v27 uint8 = (^(v1))
                                        v14 += v14
                                        var v28 int8 = v3
                                        if false {
                                            v2 = append(v2, ((("foo bar" + v20) + ("\n\t" + v12)) + "\n\t"))
                                            var v29 int8 = (v3 - (v28 / ((v3) | 1)))
                                            var v30 uint64 = (^(v14))
                                            _ = v29
                                            _ = v30
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = ("a" + v20)
                                        }
                                        var v31 int64 = (v26 + (v26 + (-(v26))))
                                        _ = v27
                                        _ = v28
                                        _ = v31
                                    default:
                                        v12 = (v20 + v12)
                                    }
                                    _ = v26
                                }
                                switch (fi13 | fi15) {
                                case -12:
                                    var v32 float64 = 1.5
                                    _ = v32
                                default:
                                    var v33 uint32 = 42
                                    if true {
                                        if true {
                                            if true {
                                                v2 = append(v2, (v12 + ((v20 + "hello") + v20)))
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v34 := GenS2{F0: v3, F1: -999, F2: 1}
                                            _ = v34
                                        } else {
                                            v35 := GenS3{F0: (v1 - v1), F1: 42, F2: 65535}
                                            fmt.Println(v12, v3, v11)
                                            _ = v35
                                        }
                                    } else {
                                        v36 := GenS5{F0: 477}
                                        if true {
                                            fmt.Println(fi15, fi15, v20)
                                            v14 = (v14 >> ((v5) & 63))
                                        } else {
                                            var v37 int32 = -2147483648
                                            _ = v37
                                        }
                                        _ = v36
                                    }
                                    v3 -= -69
                                    _ = v33
                                }
                                _ = v25
                            } else {
                                v38 := GenS3{F0: (^(v1)), F1: -1, F2: uint16(v1)}
                                fmt.Println(v5, v3)
                                v1 = (v1 * (v1 & v1))
                                switch fi13 {
                                case -825:
                                    var v39 int64 = 591
                                    _ = v39
                                default:
                                    if (!false) {
                                        v11 += (-2.25 - float32(1e10))
                                    } else {
                                        v2 = append(v2, v20)
                                        if (float64(0.0) == float64(-2.25)) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            if (!(((v11 / (-2.25 - (float32(0.1) / -2.25))) - (-2.25 + 3.14159)) < v11)) {
                                                v5 = (v5 + v5)
                                                var v40 uint8 = (^(v1))
                                                var v41 uint = 65535
                                                _ = v40
                                                _ = v41
                                            }
                                        } else {
                                            switch fi15 {
                                            case 42:
                                                v12 = v20
                                            default:
                                                switch 552 {
                                                case 42:
                                                    var v42 string = v20
                                                    _ = v42
                                                case 522:
                                                    var v43 string = ((v12 + ("\n\t" + "foo bar")) + "a")
                                                    fmt.Println(v3, v12)
                                                    v44 := GenS2{F0: v3, F1: 42, F2: 32767}
                                                    _ = v43
                                                    _ = v44
                                                }
                                            }
                                        }
                                        v20 = (("hello" + v20) + (((("hello" + "a") + "\n\t") + "foo bar") + v20))
                                    }
                                }
                                var v45 int8 = 82
                                _ = v38
                                _ = v45
                            }
                        }
                        switch fi13 {
                        case 127:
                            if (false && (9223372036854775807 >= 9223372036854775807)) {
                                switch 42 {
                                case 686:
                                    v46 := GenS6{F0: true, F1: 42, F2: 475}
                                    var v47 uint8 = 128
                                    _ = v46
                                    _ = v47
                                default:
                                    v2 = append(v2, ((("" + v12) + v12) + v20))
                                }
                            }
                        }
                    } else {
                        _ = len(v2)
                    }
                    if true {
                        if (255 <= 42) {
                            v48 := GenS4{F0: "a", F1: fi13}
                            _ = v48
                        }
                        v49 := GenS7{F0: v20}
                        _ = len(v2)
                        _ = v49
                    }
                } else {
                    v11 = 1e10
                    if false {
                        if (false || (true && false)) {
                            if (42 <= 42) {
                                v11 = ((-2.25 / float32(float64(v11))) + (((v11 - v11) / v11) - v11))
                            }
                            var v50 uint8 = (v1 / ((128) | 1))
                            v50 = v50
                            _ = v50
                        }
                    }
                    var v51 uint32 = 2147483647
                    _ = v51
                }
            case "":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                if true {
                    if true {
                        if false {
                            v11 -= (v11 / v11)
                            v12 = v20
                            v52 := GenS5{F0: 32767}
                            _ = v52
                        } else {
                            v5 += v5
                            if (0.0 <= float64(((((0.1 + -2.25) - 3.14159) * v11) + v11))) {
                                fmt.Println(v14, v12)
                                var v53 uint = uint(v14)
                                var v54 bool = false
                                var v55 int = (fi15 & (fi15 - (fi15 & fi13)))
                                var v56 uint = 2
                                v11 -= 1e10
                                switch (fi13 / (((fi15 + (fi13 & v55))) | 1)) {
                                case 32768:
                                    var v57 int16 = 1000
                                    _ = v57
                                case 0:
                                    v53 = (v53 ^ v56)
                                case 128:
                                    fmt.Println(v12, v53, v5)
                                default:
                                    var v58 uint8 = v1
                                    _ = v58
                                }
                                _ = v53
                                _ = v54
                                _ = v55
                                _ = v56
                            }
                            var v59 int = fi15
                            _ = v59
                        }
                        v1 = v1
                    }
                    v60 := GenS8{F0: int16(fi15), F1: (173 != 65535)}
                    _ = v60
                } else {
                    v20 = (v20 + (("" + ((v20 + v20) + "")) + v20))
                    var v61 bool = (!true)
                    var v62 float32 = (1.5 - v11)
                    _ = v61
                    _ = v62
                }
            }
            _ = fi15
            _ = v20
        }
        v63 := GenS6{F0: (((v5 % (((^(v5))) | 1)) < (v5 * (v5 + (v5 / (((v5 * 1)) | 1))))) || true), F1: 42, F2: 224}
        _ = fi13
        _ = v14
        _ = v63
    }
    for fi64 := 0; fi64 < 12; fi64++ {
        var v65 int = fi64
        _ = fi64
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
