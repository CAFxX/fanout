package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 uint; F2 bool }
type GenS6 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS7 struct { F0 int8; F1 uint16; F2 int16 }
type GenS8 struct { F0 int32 }

func main() {
    var v1 string = "a"
    v2 := []string{v1, (v1 + v1), "a", ("\n\t" + v1)}
    var v3 int8 = 37
    var v4 int64 = 1
    var v5 string = ("a" + ("" + ((v1 + v1) + "foo bar")))
    v6 := make(map[int]uint32)
    v7 := v6[-9223372036854775808]
    fmt.Println(v7, v3)
    v8 := GenS1{F0: 0, F1: v7}
    v7 = v7
    var v9 bool = (((v3 >> ((65535) & 63)) == v3) && false)
    if v9 {
        switch 236 {
        case 2147483647:
            for fi10 := 0; fi10 < 1; fi10++ {
                if v9 {
                    v11 := v6[(fi10 | (fi10 >> ((1) & 63)))]
                    if (0 >= 2) {
                        var v12 uint16 = 74
                        for fi13 := 0; fi13 < 8; fi13++ {
                            fmt.Println(fi13)
                            _ = fi13
                        }
                        var v14 float32 = (1.5 * ((float32(0.1) + (1.5 - (1e10 / -2.25))) * 0.0))
                        _ = v12
                        _ = v14
                    } else {
                        switch fi10 {
                        case -9223372036854775808:
                            for fi15 := 0; fi15 < 2; fi15++ {
                                v16 := GenS2{F0: 277, F1: true, F2: 143}
                                v5 = v5
                                v17 := GenS3{F0: 0}
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                        case 914:
                            v18 := GenS1{F0: uint(v11), F1: (v7 << ((383) & 63))}
                            _ = v18
                        case -1:
                            for fi19 := 0; fi19 < 13; fi19++ {
                                fmt.Println(v5, fi10, v4)
                                if (((!v9) || (1.5 >= float64(float32((1.5 / (((1e10 - (3.14159 + 1e10)) - 0.0) - (float64((1.5 + ((1.5 + -2.25) / 1.0))) * 0.1))))))) || false) {
                                    var v20 int64 = 492
                                    switch (fi19 + 2147483647) {
                                    case 286:
                                        if true {
                                            v5 = "a"
                                            v1 = "\n\t"
                                            var v21 uint32 = v11
                                            switch 805 {
                                            case -157:
                                                var v22 bool = ((false || (v9 || false)) || ((1e10 == 0.0) || (v9 && false)))
                                                _ = v22
                                            case 128:
                                                if v9 {
                                                    v1 = "a"
                                                    v5 = "hello"
                                                } else {
                                                    v3 = 100
                                                }
                                            case 32768:
                                                if len(v2) > 0 {
                                                    v2[0] = v1
                                                }
                                            default:
                                                v1 = (("" + "") + ((v5 + "\n\t") + v1))
                                            }
                                            v23 := GenS2{F0: 1, F1: (!(3.14159 >= 0.1)), F2: 128}
                                            _ = v21
                                            _ = v23
                                        } else {
                                            var v24 int64 = (v20 >> ((317) & 63))
                                            _ = len(v2)
                                            _ = v24
                                        }
                                    case 1:
                                        var v25 string = (v1 + "foo bar")
                                        if v9 {
                                            v6[32767] = v7
                                            if v9 {
                                                var v26 float32 = 0.0
                                                v20 += (^(v20))
                                                v27 := GenS2{F0: 18446744073709551615, F1: v9, F2: 0}
                                                _ = v26
                                                _ = v27
                                            } else {
                                                fmt.Println(v1, v7)
                                                var v28 int64 = -857
                                                if (v1 > v5) {
                                                    var v29 int64 = (v28 >> ((210) & 63))
                                                    _ = v29
                                                }
                                                _ = v28
                                            }
                                            var v30 uint64 = 42
                                            _ = v30
                                        } else {
                                            var v31 int16 = 1000
                                            _ = v31
                                        }
                                        if v9 {
                                            v5 = v25
                                            v3 = v3
                                            var v32 uint64 = 18446744073709551615
                                            _ = v32
                                        }
                                        _ = v25
                                    case 32767:
                                        var v33 uint64 = 4294967295
                                        _ = v33
                                    default:
                                        var v34 uint16 = 1
                                        _ = v34
                                    }
                                    if (-1 != -1) {
                                        fmt.Println(v20)
                                        if true {
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v35 bool = true
                                                _ = v35
                                            case "":
                                                var v36 int64 = (v20 * (v20 - v4))
                                                _ = v36
                                            }
                                        } else {
                                            _ = len(v2)
                                        }
                                        if (v3 < 1) {
                                            var v37 uint16 = 1
                                            _ = v37
                                        }
                                        if v9 {
                                            fmt.Println(v20)
                                            if (false || (true && (!v9))) {
                                                v38 := GenS4{F0: 9223372036854775807}
                                                _ = v38
                                            } else {
                                                v9 = (!(v9 || ((((v1 + v5) + v1) < "\n\t") && (v9 && v9))))
                                                var v39 uint16 = 13
                                                v40 := GenS4{F0: -9223372036854775808}
                                                _ = v39
                                                _ = v40
                                            }
                                            var v41 uint32 = v11
                                            _ = v41
                                        }
                                    } else {
                                        var v42 uint32 = v7
                                        _ = v42
                                    }
                                    v43 := GenS4{F0: (v4 % (((v20 | -1)) | 1))}
                                    fmt.Println(v4, v1, v3)
                                    fmt.Println(v1, v4)
                                    if ((v11 & (v7 / ((v7) | 1))) < v11) {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        fmt.Println(v20, v11)
                                    }
                                    v6[fi19] = v7
                                    _ = v20
                                    _ = v43
                                }
                                _ = fi19
                            }
                        default:
                            fmt.Println(v1, v3, v5)
                        }
                        v44 := GenS4{F0: (v4 & 1)}
                        v3 = (v3 - (^(v3)))
                        _ = v44
                    }
                    _ = v11
                } else {
                    v45 := GenS5{F0: (v3 / ((v3) | 1)), F1: 0, F2: true}
                    _ = len(v6)
                    _ = v45
                }
                if v9 {
                    var v46 uint8 = 0
                    v9 = true
                    var v47 string = ("\n\t" + ((v1 + "a") + ""))
                    _ = v46
                    _ = v47
                } else {
                    v9 = (!v9)
                    switch 0 {
                    case -902:
                        if v9 {
                            v6[0] = (v7 ^ v7)
                            switch 0 {
                            case -1:
                                for fi48 := 0; fi48 < 2; fi48++ {
                                    switch (v1 + v5) {
                                    case "foo bar":
                                        v5 = v1
                                    case "a":
                                        v7 += (^(v7))
                                    }
                                    _ = fi48
                                }
                            case 4294967295:
                                v4 += (v4 % ((v4) | 1))
                            case 32768:
                                v49 := GenS6{F0: 127, F1: 4294967295, F2: -1}
                                _ = v49
                            }
                        }
                    case 717:
                        v50 := GenS6{F0: 42, F1: v7, F2: 1}
                        _ = v50
                    case 9223372036854775807:
                        v3 += int8(fi10)
                    default:
                        var v51 uint64 = 64
                        _ = v51
                    }
                    var v52 int = fi10
                    _ = v52
                }
                switch ("foo bar" + (v1 + v1)) {
                case "a":
                    v5 = ("foo bar" + v5)
                case "\n\t":
                    var v53 int32 = -2147483648
                    _ = v53
                }
                _ = fi10
            }
        case 32767:
            if (true || v9) {
                fmt.Println(v7, v9, v3)
                switch v1 {
                case "hello":
                    v2 = append(v2, (v1 + v5))
                default:
                    if v9 {
                        if true {
                            var v54 uint = 253
                            for fi55 := 0; fi55 < 15; fi55++ {
                                var v56 float64 = float64(0.1)
                                if v9 {
                                    fmt.Println(v54)
                                    var v57 uint8 = 42
                                    _ = v57
                                }
                                if v9 {
                                    for fi58 := 0; fi58 < 1; fi58++ {
                                        v1 = "a"
                                        if v9 {
                                            v9 = ((v9 || false) && v9)
                                        }
                                        _ = fi58
                                    }
                                    for fi59 := 0; fi59 < 7; fi59++ {
                                        var v60 uint32 = (v7 - (^(v7)))
                                        _ = fi59
                                        _ = v60
                                    }
                                    delete(v6, fi55)
                                }
                                _ = fi55
                                _ = v56
                            }
                            fmt.Println(v5, v54)
                            _ = v54
                        } else {
                            var v61 int32 = -790
                            var v62 uint16 = 1
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v61
                            _ = v62
                        }
                        v1 = v5
                    }
                }
            }
        case 1:
            if (1e10 >= (1.5 + 1.5)) {
                var v63 int64 = 366
                _ = v63
            }
        default:
            if false {
                v7 -= v7
            }
        }
        v64 := GenS1{F0: 876, F1: (v7 >> ((2) & 63))}
        var v65 int16 = 32767
        _ = v64
        _ = v65
    } else {
        if (1 == 4294967295) {
            v4 += v4
            var v66 bool = (9223372036854775807 <= -786)
            _ = v66
        }
        switch v5 {
        case "a":
            switch v5 {
            case "a":
                var v67 float64 = (1e10 - 0.0)
                _ = v67
            case "hello":
                for fi68 := 0; fi68 < 10; fi68++ {
                    var v69 int16 = 673
                    switch v1 {
                    case "a":
                        var v70 uint = 4294967295
                        _ = v70
                    case "foo bar":
                        var v71 bool = (false && (!v9))
                        _ = v71
                    case "hello":
                        var v72 uint16 = 188
                        _ = v72
                    }
                    _ = fi68
                    _ = v69
                }
            }
        case "hello":
            var v73 int8 = v3
            _ = v73
        }
        v74 := GenS3{F0: 255}
        _ = v74
    }
    fmt.Println(v3)
    v4 += (v4 % ((v4) | 1))
    v75 := GenS7{F0: v3, F1: 32767, F2: 42}
    fmt.Println(v1)
    v76 := GenS8{F0: -987}
    var v77 string = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v77:", v77)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v75
    _ = v76
    _ = v77
    _ = v8
    _ = v9
}
