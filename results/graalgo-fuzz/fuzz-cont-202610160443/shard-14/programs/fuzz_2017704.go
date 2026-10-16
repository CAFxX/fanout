package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int8 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int32; F1 int8 }

func main() {
    v1 := []int{0, 127}
    var v2 uint32 = 4294967295
    var v3 bool = (!(1 == -789))
    var v4 int32 = 2147483647
    var v5 uint64 = 416
    v6 := []int{}
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint = 570
        _ = fi7
        _ = v8
    }
    var v9 uint = 99
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            v2 = 544
            var v12 int = fi11
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    if (1e10 != float64(0.1)) {
        if ((!v3) || (!v3)) {
            _ = len(v1)
            if v3 {
                var v13 int = 42
                fmt.Println(v3, v2, v13)
                var v14 uint8 = 1
                _ = v13
                _ = v14
            } else {
                switch "" {
                case "":
                    for fi15 := 0; fi15 < 14; fi15++ {
                        if (-205 > 32767) {
                            if (!true) {
                                v5 += v5
                                var v16 uint64 = 4294967295
                                v17 := GenS1{F0: 299, F1: 42, F2: int8(fi15)}
                                _ = v16
                                _ = v17
                            } else {
                                switch ("\n\t" + "\n\t") {
                                case "\n\t":
                                    if (2 != 774) {
                                        v2 += 65535
                                    }
                                case "":
                                    if (-77 <= -15) {
                                        v2 -= v2
                                    } else {
                                        var v18 uint8 = 164
                                        _ = v18
                                    }
                                case "foo bar":
                                    if true {
                                        var v19 uint64 = (v5 + v5)
                                        _ = v19
                                    }
                                default:
                                    v3 = false
                                }
                                switch "" {
                                case "hello":
                                    v2 = 0
                                case "a":
                                    var v20 uint = (v9 << ((v9) & 63))
                                    _ = v20
                                case "":
                                    v21 := GenS1{F0: 707, F1: 32767, F2: 127}
                                    _ = v21
                                default:
                                    var v22 uint16 = 676
                                    _ = v22
                                }
                            }
                            var v23 int8 = -1
                            _ = v23
                        } else {
                            v5 = 18446744073709551615
                            switch 42 {
                            case -1:
                                var v24 uint16 = 924
                                _ = v24
                            case 32768:
                                if v3 {
                                    var v25 uint8 = 1
                                    for fi26 := 0; fi26 < 15; fi26++ {
                                        fmt.Println(v9, v5)
                                        _ = fi26
                                    }
                                    v27 := GenS1{F0: uint16(v2), F1: -32768, F2: -128}
                                    _ = v25
                                    _ = v27
                                } else {
                                    fmt.Println(v3, v5)
                                }
                            default:
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                            var v28 int = (fi15 - fi15)
                            _ = v28
                        }
                        if v3 {
                            for fi29 := 0; fi29 < 12; fi29++ {
                                v9 -= v9
                                var v30 int8 = 1
                                fmt.Println(fi15, v4)
                                _ = fi29
                                _ = v30
                            }
                        } else {
                            var v31 uint64 = (^(v5))
                            v32 := GenS1{F0: 255, F1: 1000, F2: -128}
                            _ = v31
                            _ = v32
                        }
                        _ = fi15
                    }
                case "a":
                    var v33 uint = 18446744073709551615
                    _ = v33
                case "foo bar":
                    var v34 int16 = 93
                    _ = v34
                default:
                    var v35 int = 822
                    _ = v35
                }
            }
        } else {
            var v36 bool = ((float64(1.5) / float64((1.5 + -2.25))) > 3.14159)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 13; fi37++ {
        var v38 int16 = -32768
        for fi39 := 0; fi39 < 18; fi39++ {
            v4 += (v4 + (v4 - v4))
            if v3 {
                v1 = append(v1, (fi37 ^ (^(fi37))))
                fmt.Println(fi39, v5, v5)
                if (("" >= ("a" + ("a" + ("hello" + "foo bar")))) || v3) {
                    v9 -= v9
                    var v40 int64 = -1
                    _ = v40
                } else {
                    v5 += 42
                    v4 = (v4 + (v4 | (^(v4))))
                    if (!(v3 || (false && v3))) {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v41 int = fi39
                        var v42 float32 = (0.0 - 3.14159)
                        _ = v41
                        _ = v42
                    }
                }
            }
            var v43 string = ""
            _ = fi39
            _ = v43
        }
        var v44 int16 = v38
        _ = fi37
        _ = v38
        _ = v44
    }
    var v45 int32 = (v4 << (((v9 / ((v9) | 1))) & 63))
    v46 := GenS2{F0: 1, F1: v4}
    v4 += 42
    if true {
        var v47 int8 = -37
        if v3 {
            var v48 uint8 = 255
            if true {
                if len(v6) > 0 {
                    v6[0] = 9223372036854775807
                }
                var v49 float64 = ((1e10 + 3.14159) - 0.1)
                var v50 string = ("foo bar" + "")
                _ = v49
                _ = v50
            }
            switch "foo bar" {
            case "":
                var v51 string = "hello"
                _ = v51
            case "\n\t":
                switch 128 {
                case 4294967295:
                    v47 += 100
                case 32767:
                    fmt.Println(v47)
                }
            default:
                for fi52 := 0; fi52 < 6; fi52++ {
                    var v53 uint16 = 255
                    _ = fi52
                    _ = v53
                }
            }
            _ = v48
        } else {
            for fi54 := 0; fi54 < 14; fi54++ {
                var v55 int8 = (v47 - v47)
                if false {
                    if v3 {
                        if ((v3 && v3) || true) {
                            v55 += (-(v47))
                        }
                        var v56 uint64 = 1
                        var v57 uint8 = 80
                        _ = v56
                        _ = v57
                    }
                    for fi58 := 0; fi58 < 11; fi58++ {
                        fmt.Println(v47, v47)
                        v4 -= (-(v4))
                        v45 = -1
                        switch (fi58 - (fi58 | (fi54 * fi54))) {
                        case 2147483647:
                            var v59 int = (fi58 ^ 32768)
                            if true {
                                v5 -= v5
                            } else {
                                var v60 float32 = -2.25
                                _ = v60
                            }
                            var v61 int = (v59 / (((^(fi54))) | 1))
                            _ = v59
                            _ = v61
                        case -853:
                            v62 := GenS3{F0: 0, F1: v47}
                            _ = v62
                        case 319:
                            var v63 uint64 = 299
                            _ = v63
                        }
                        fmt.Println(v2, v9, v47)
                        _ = fi58
                    }
                    v64 := GenS2{F0: 927, F1: 782}
                    _ = v64
                } else {
                    var v65 uint = 42
                    var v66 int16 = 1
                    fmt.Println(v2, v45, v66)
                    _ = v65
                    _ = v66
                }
                fmt.Println(v4)
                _ = fi54
                _ = v55
            }
            v5 -= v5
        }
        _ = v47
    } else {
        var v67 uint = 18446744073709551615
        for fi68 := 0; fi68 < 1; fi68++ {
            if v3 {
                for fi69 := 0; fi69 < 4; fi69++ {
                    v67 -= v9
                    _ = fi69
                }
            }
            v6 = append(v6, fi68)
            _ = fi68
        }
        _ = v67
    }
    fmt.Println(v2, v9)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v9
}
