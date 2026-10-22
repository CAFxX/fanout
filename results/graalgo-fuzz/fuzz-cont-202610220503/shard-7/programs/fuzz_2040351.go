package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint8; F1 bool; F2 int8 }
type GenS4 struct { F0 uint64; F1 uint; F2 uint32 }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 int8; F1 uint16 }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 uint64; F1 bool }
type GenS9 struct { F0 bool; F1 uint32 }

func main() {
    var v1 bool = false
    v2 := make(map[int]bool)
    var v3 uint = 495
    var v4 int16 = -582
    var v5 bool = ((10 == 128) || (v1 && v1))
    var v6 int8 = 1
    v7 := v2[0]
    if v7 {
        v1 = v1
    } else {
        v8 := v2[267]
        _ = v8
    }
    var v9 uint32 = 2147483648
    if true {
        var v10 int16 = (v4 / ((401) | 1))
        _ = v10
    } else {
        var v11 int = 42
        fmt.Println(v5, v9)
        _ = v11
    }
    var v12 uint8 = 1
    if true {
        fmt.Println(v9, v6, v1)
        v9 = v9
        var v13 int32 = 0
        _ = v13
    } else {
        if (!("\n\t" == ("foo bar" + "hello"))) {
            v14 := []int{42}
            _ = v14
        }
        fmt.Println(v5, v9, v12)
        v9 -= (v9 - 2)
    }
    switch 127 {
    case 636:
        for fi15 := 0; fi15 < 1; fi15++ {
            for fi16 := 0; fi16 < 5; fi16++ {
                if v5 {
                    var v17 int32 = -981
                    var v18 uint32 = v9
                    _ = v17
                    _ = v18
                }
                var v19 uint64 = uint64(v12)
                _ = fi16
                _ = v19
            }
            v12 -= (^(v12))
            _ = fi15
        }
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        if (v3 == (v3 ^ v3)) {
            var v21 uint32 = v9
            fmt.Println(fi20, v4, v1)
            _ = v21
        }
        for fi22 := 0; fi22 < 18; fi22++ {
            if false {
                if ((v4 << (((v3 << ((42) & 63))) & 63)) == (v4 - (v4 / ((32767) | 1)))) {
                    if v5 {
                        var v23 uint32 = 2
                        _ = v23
                    }
                    if v5 {
                        if v5 {
                            var v24 int = fi20
                            if v5 {
                                var v25 string = "foo bar"
                                v26 := GenS1{F0: 0}
                                v27 := GenS2{F0: 2, F1: fi20}
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                            _ = v24
                        } else {
                            switch fi22 {
                            case 127:
                                if true {
                                    v28 := GenS1{F0: 378}
                                    v12 += 128
                                    delete(v2, 526)
                                    _ = v28
                                }
                            case 9223372036854775807:
                                v29 := []int{fi20, 128, (fi20 | -899), (fi22 % ((fi22) | 1))}
                                _ = v29
                            case -1:
                                v30 := GenS3{F0: (v12 >> (((^(v3))) & 63)), F1: true, F2: int8(v3)}
                                _ = v30
                            default:
                                v9 += v9
                            }
                            var v31 int32 = -867
                            _ = v31
                        }
                        var v32 float32 = (0.1 / 1.0)
                        switch fi20 {
                        case 4294967295:
                            var v33 int = (fi20 >> ((v3) & 63))
                            _ = v33
                        case -522:
                            var v34 float32 = 3.14159
                            _ = v34
                        }
                        delete(v2, fi20)
                        if (!v5) {
                            v35 := []int{(-(fi22))}
                            v3 = v3
                            var v36 bool = ((((v32 - ((((((0.1 / v32) - (v32 - -2.25)) / v32) + 1.5) * 3.14159) / 0.1)) + v32) / 1.5) >= v32)
                            _ = v35
                            _ = v36
                        }
                        _ = v32
                    }
                } else {
                    var v37 int64 = 599
                    var v38 int16 = (v4 % (((v4 - (v4 << ((v3) & 63)))) | 1))
                    _ = v37
                    _ = v38
                }
                if v1 {
                    v12 -= 16
                    _ = len(v2)
                }
                v9 = v9
                var v39 int64 = -431
                fmt.Println(fi22, v5)
                v40 := []int{fi22}
                _ = v39
                _ = v40
            }
            v7 = true
            switch (^(fi22)) {
            case 867:
                v2[fi20] = v1
            case -452:
                v41 := GenS4{F0: 18446744073709551615, F1: (v3 + (v3 - (v3 & 808))), F2: 42}
                _ = v41
            }
            v3 = (v3 ^ (v3 | 103))
            var v42 uint = (v3 % (((v3 >> ((v3) & 63))) | 1))
            _ = fi22
            _ = v42
        }
        _ = fi20
    }
    for fi43 := 0; fi43 < 2; fi43++ {
        v1 = true
        _ = fi43
    }
    v44 := []int{}
    v4 += (v4 << ((302) & 63))
    if len(v44) > 0 {
        v44[0] = 9223372036854775807
    }
    fmt.Println(v7)
    v3 -= v3
    v1 = (835 <= 4294967295)
    var v45 uint8 = 1
    if (-9223372036854775808 < 42) {
        v46 := GenS3{F0: v45, F1: (!((!true) || false)), F2: (v6 / ((-58) | 1))}
        v4 = 42
        for fi47 := 0; fi47 < 19; fi47++ {
            switch "" {
            case "hello":
                var v48 string = ("" + "\n\t")
                _ = v48
            }
            fmt.Println(v4)
            _ = fi47
        }
        _ = v46
    } else {
        if v1 {
            v12 -= v12
            for fi49 := 0; fi49 < 2; fi49++ {
                v50 := GenS5{F0: (v12 * v12), F1: "\n\t"}
                _ = fi49
                _ = v50
            }
            for fi51 := 0; fi51 < 7; fi51++ {
                v44 = append(v44, int(v3))
                switch (fi51 % ((0) | 1)) {
                case -374:
                    var v52 float32 = 3.14159
                    _ = v52
                case 0:
                    v53 := v2[fi51]
                    _ = v53
                default:
                    fmt.Println(v45)
                }
                for fi54 := 0; fi54 < 9; fi54++ {
                    v55 := GenS5{F0: v12, F1: "a"}
                    if (-9223372036854775808 > 601) {
                        v4 = (v4 << ((v3) & 63))
                        if (976 <= 2) {
                            if v7 {
                                _ = len(v44)
                            } else {
                                switch fi54 {
                                case -472:
                                    v6 = int8(v3)
                                default:
                                    v56 := GenS5{F0: (^(v45)), F1: "hello"}
                                    _ = v56
                                }
                                v57 := GenS6{F0: 1, F1: 499}
                                v44 = append(v44, fi51)
                                _ = v57
                            }
                        }
                        fmt.Println(v9)
                        v58 := GenS7{F0: v9, F1: ((1.5 - 0.0) != 1e10)}
                        if false {
                            if (!false) {
                                if true {
                                    v59 := GenS6{F0: -128, F1: 65535}
                                    _ = v59
                                }
                                v1 = v7
                            }
                        }
                        v6 -= v6
                        var v60 float32 = (0.1 - 1.5)
                        _ = v58
                        _ = v60
                    }
                    _ = fi54
                    _ = v55
                }
                _ = fi51
            }
        } else {
            for fi61 := 0; fi61 < 6; fi61++ {
                var v62 float64 = 0.0
                if len(v44) >= 2 {
                    v44 = v44[1:len(v44)]
                }
                v63 := GenS1{F0: 9223372036854775807}
                _ = fi61
                _ = v62
                _ = v63
            }
            var v64 float64 = float64(1.5)
            _ = v64
        }
        if (true || v5) {
            var v65 float64 = 1.5
            _ = v65
        }
        if (807 >= 65535) {
            var v66 uint8 = (v45 | v45)
            if (1.5 == 0.1) {
                v45 = v12
                v7 = ((v3 < 1) || v5)
                v4 = -32768
            }
            for fi67 := 0; fi67 < 9; fi67++ {
                for fi68 := 0; fi68 < 20; fi68++ {
                    if (1 == 32767) {
                        var v69 float64 = float64(0.1)
                        v9 -= v9
                        _ = v69
                    } else {
                        var v70 float64 = 0.0
                        if (-56 <= v6) {
                            fmt.Println(v12)
                            switch fi67 {
                            case 2147483647:
                                switch ("\n\t" + ("" + "foo bar")) {
                                case "\n\t":
                                    v71 := GenS8{F0: 4294967295, F1: (v6 > 100)}
                                    _ = v71
                                case "a":
                                    if (0 == int32(v6)) {
                                        var v72 string = "\n\t"
                                        v73 := GenS9{F0: ((v72 + (("a" + "hello") + (("" + v72) + "hello"))) == v72), F1: 65535}
                                        v70 += (3.14159 / (v70 * v70))
                                        fmt.Println(fi67)
                                        var v74 float32 = ((3.14159 - (((-2.25 + 1.5) - -2.25) - ((1.5 + (-2.25 + ((0.1 / float32(((float64(-2.25) + v70) / 1.0))) + float32(0.0)))) - 1.5))) * 0.1)
                                        _ = v72
                                        _ = v73
                                        _ = v74
                                    }
                                }
                            case 0:
                                v9 = v9
                            }
                        }
                        _ = v70
                    }
                    if v1 {
                        var v75 int32 = 0
                        v76 := v2[fi68]
                        delete(v2, fi68)
                        _ = len(v2)
                        _ = v75
                        _ = v76
                    }
                    _ = fi68
                }
                _ = fi67
            }
            _ = v66
        }
    }
    fmt.Println(v45, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
