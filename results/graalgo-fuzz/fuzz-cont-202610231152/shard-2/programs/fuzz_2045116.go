package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 int32; F1 uint64; F2 uint32 }
func genFn0(p0 string, p1 string) uint64 {
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 1; fi6++ {
            switch fi6 {
            case -9223372036854775808:
                v7 := GenS1{F0: -128}
                _ = v7
            case 32768:
                p0 = "a"
            case 9223372036854775807:
                if ("foo bar" == "") {
                    p0 = (p0 + p1)
                    var v8 uint64 = 42
                    if true {
                        v9 := []int{127, fi5}
                        _ = v9
                    }
                    _ = v8
                } else {
                    v10 := []int{-9223372036854775808, 228, fi6, 2147483647}
                    if false {
                        var v11 string = p1
                        var v12 int64 = -9223372036854775808
                        switch fi5 {
                        case -9223372036854775808:
                            fmt.Println(v11, fi5, p1)
                        case 32768:
                            var v13 int16 = 886
                            _ = v13
                        case 251:
                            v14 := GenS1{F0: 127}
                            _ = v14
                        default:
                            var v15 uint16 = 32767
                            _ = v15
                        }
                        _ = v11
                        _ = v12
                    } else {
                        var v16 float64 = ((-2.25 / 1e10) - (1e10 + 1e10))
                        _ = v16
                    }
                    _ = v10
                }
                if (!(float64((float32(3.14159) + (3.14159 - (float32(0.1) - 0.0)))) > 0.0)) {
                    p1 = (("\n\t" + "a") + (p1 + "foo bar"))
                    v17 := []int{fi6}
                    var v18 bool = (1 <= -9223372036854775808)
                    _ = v17
                    _ = v18
                }
                v19 := GenS1{F0: 1}
                _ = v19
            default:
                var v20 int = fi5
                _ = v20
            }
            var v21 int16 = 1
            var v22 int8 = -1
            _ = fi6
            _ = v21
            _ = v22
        }
        if true {
            if (128 > 127) {
                var v23 bool = (4294967295 != 839)
                _ = v23
            }
            for fi24 := 0; fi24 < 4; fi24++ {
                var v25 uint64 = 4294967295
                v26 := GenS1{F0: 0}
                if true {
                    v25 = v25
                    v27 := GenS2{F0: ((p1 + ("" + p1)) + p0)}
                    var v28 string = ""
                    var v29 uint8 = 141
                    var v30 uint32 = 42
                    var v31 uint32 = 892
                    switch p0 {
                    case "hello":
                        v30 += (^(v30))
                    default:
                        if false {
                            var v32 uint16 = 0
                            v33 := GenS1{F0: -1}
                            _ = v32
                            _ = v33
                        }
                    }
                    _ = v27
                    _ = v28
                    _ = v29
                    _ = v30
                    _ = v31
                } else {
                    var v34 float64 = 3.14159
                    _ = v34
                }
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        _ = fi5
    }
    var v35 uint8 = 28
    p1 = p0
    if (v35 != 1) {
        if false {
            v35 -= v35
        }
    }
    _ = p0
    _ = p1
    _ = v35
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 156
    var v3 int = 32767
    var v4 int32 = 1
    v2 += 2
    v2 = v2
    _ = len(v1)
    v1[(v3 << ((255) & 63))] = (v3 | v3)
    if ((2 < 224) && (1.5 >= float32(0.0))) {
        if false {
            v3 = v3
            v36 := GenS2{F0: ("\n\t" + "")}
            _ = v36
        }
        if (((0.0 - ((1e10 + 0.1) * 0.0)) / (0.1 / ((3.14159 - 1e10) * 1e10))) >= -2.25) {
            v4 = (v4 - v4)
            v37 := []int{v3, -296, (v3 >> ((535) & 63))}
            if false {
                v38 := v1[(v3 * v3)]
                v2 = v2
                for fi39 := 0; fi39 < 5; fi39++ {
                    if true {
                        for fi40 := 0; fi40 < 14; fi40++ {
                            if (false || true) {
                                v2 += v2
                            }
                            _ = fi40
                        }
                        var v41 float64 = ((1e10 * 0.1) + 1.5)
                        var v42 float64 = v41
                        _ = v41
                        _ = v42
                    } else {
                        if len(v37) > 0 {
                            v37[0] = v38
                        }
                        for fi43 := 0; fi43 < 12; fi43++ {
                            v44 := GenS2{F0: "\n\t"}
                            var v45 bool = true
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                        v38 -= v38
                    }
                    if (!(true && (!true))) {
                        fmt.Println(v38, v3, v3)
                        fmt.Println(fi39, v2)
                    } else {
                        if false {
                            v3 = v3
                        } else {
                            _ = len(v1)
                            v3 += (v38 ^ (v38 << ((177) & 63)))
                        }
                    }
                    _ = fi39
                }
                _ = v38
            }
            _ = v37
        }
    } else {
        var v46 int = (v3 << ((87) & 63))
        fmt.Println(v46)
        _ = v46
    }
    v3 -= (v3 | (v3 << ((439) & 63)))
    v3 -= (v3 % ((384) | 1))
    v3 -= v3
    v47 := GenS3{F0: (2 >= 1)}
    fmt.Println(v4)
    if true {
        var v48 int = (v3 * (v3 ^ v3))
        v49 := []int{-380, (v3 << ((4294967295) & 63))}
        _ = v48
        _ = v49
    } else {
        v50 := GenS4{F0: 572, F1: false}
        for fi51 := 0; fi51 < 16; fi51++ {
            var v52 uint32 = 65535
            _ = fi51
            _ = v52
        }
        _ = v50
    }
    if ((!false) || false) {
        if false {
            v53 := GenS3{F0: ((1 < 389) || true)}
            if ((1e10 - -2.25) < 1.5) {
                switch 899 {
                case 0:
                    v4 -= (v4 >> ((57) & 63))
                case 4294967295:
                    var v54 uint = 65535
                    _ = v54
                }
                v2 = 72
            } else {
                var v55 string = "hello"
                v4 += v4
                fmt.Println(v55, v3, v3)
                _ = v55
            }
            if false {
                v56 := GenS5{F0: (-(v4)), F1: 293, F2: 2147483647}
                for fi57 := 0; fi57 < 5; fi57++ {
                    for fi58 := 0; fi58 < 13; fi58++ {
                        v59 := GenS1{F0: -1}
                        _ = fi58
                        _ = v59
                    }
                    _ = fi57
                }
                _ = v56
            }
            _ = v53
        } else {
            switch "\n\t" {
            case "":
                for fi60 := 0; fi60 < 16; fi60++ {
                    for fi61 := 0; fi61 < 4; fi61++ {
                        var v62 uint = 323
                        v3 = (fi61 >> (((^(v62))) & 63))
                        _ = fi61
                        _ = v62
                    }
                    v2 += v2
                    _ = fi60
                }
            }
            v63 := GenS3{F0: false}
            var v64 float32 = (0.1 * 1e10)
            _ = v63
            _ = v64
        }
        _ = len(v1)
        switch (v3 % (((v3 % (((v3 % ((v3) | 1))) | 1))) | 1)) {
        case 32767:
            for fi65 := 0; fi65 < 19; fi65++ {
                v66 := []int{}
                var v67 uint8 = v2
                _ = fi65
                _ = v66
                _ = v67
            }
        case 0:
            v2 -= v2
        }
    }
    for fi68 := 0; fi68 < 3; fi68++ {
        var v69 int8 = -73
        var v70 uint16 = 0
        _ = fi68
        _ = v69
        _ = v70
    }
    v71 := GenS4{F0: 42, F1: (34 != 0)}
    for fi72 := 0; fi72 < 10; fi72++ {
        for fi73 := 0; fi73 < 12; fi73++ {
            var v74 string = "a"
            _ = fi73
            _ = v74
        }
        if (("hello" + "foo bar") > "a") {
            for fi75 := 0; fi75 < 11; fi75++ {
                var v76 bool = (!(65535 >= 2))
                var v77 uint8 = (v2 / ((127) | 1))
                _ = fi75
                _ = v76
                _ = v77
            }
        }
        _ = fi72
    }
    v78 := GenS3{F0: false}
    v3 += int(v2)
    v79 := []int{(v3 % (((v3 & (^(v3)))) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v79:", len(v79), v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v71
    _ = v78
    _ = v79
}
