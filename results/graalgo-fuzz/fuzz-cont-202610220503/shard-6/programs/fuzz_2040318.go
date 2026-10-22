package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 uint64 }
func genFn0() int {
    switch 948 {
    case 128:
        for fi7 := 0; fi7 < 2; fi7++ {
            if false {
                var v8 int64 = 42
                _ = v8
            }
            var v9 int32 = 2147483647
            _ = fi7
            _ = v9
        }
    case -2147483648:
        var v10 bool = true
        _ = v10
    case 127:
        var v11 bool = (true || (true && false))
        _ = v11
    default:
        var v12 uint16 = 1
        _ = v12
    }
    if false {
        v13 := []int{32767, 732, 4294967295, 42}
        var v14 uint = 351
        _ = v13
        _ = v14
    } else {
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 uint = 756
            _ = fi15
            _ = v16
        }
        if false {
            var v17 int16 = 1
            var v18 uint = 4294967295
            for fi19 := 0; fi19 < 9; fi19++ {
                fmt.Println(v18, fi19, v17)
                _ = fi19
            }
            _ = v17
            _ = v18
        } else {
            for fi20 := 0; fi20 < 12; fi20++ {
                if true {
                    var v21 uint32 = 42
                    _ = v21
                } else {
                    var v22 int = (fi20 - (fi20 / ((fi20) | 1)))
                    var v23 float64 = 0.1
                    if true {
                        if true {
                            v22 -= fi20
                            v24 := GenS1{F0: 875, F1: 4294967295}
                            _ = v24
                        }
                        var v25 bool = false
                        v23 += 3.14159
                        _ = v25
                    }
                    _ = v22
                    _ = v23
                }
                for fi26 := 0; fi26 < 15; fi26++ {
                    var v27 uint32 = 4294967295
                    _ = fi26
                    _ = v27
                }
                switch fi20 {
                case 1:
                    var v28 uint8 = 255
                    _ = v28
                case 2147483647:
                    v29 := []int{fi20, (fi20 / ((fi20) | 1))}
                    _ = v29
                case 9223372036854775807:
                    if true {
                        if (42 <= -1) {
                            if true {
                                var v30 int = 1
                                var v31 bool = (false && true)
                                _ = v30
                                _ = v31
                            }
                        }
                        fmt.Println(fi20)
                        v32 := []int{670, fi20, (fi20 << ((51) & 63)), (fi20 | fi20)}
                        _ = v32
                    }
                }
                _ = fi20
            }
            for fi33 := 0; fi33 < 3; fi33++ {
                switch fi33 {
                case 1:
                    v34 := []int{}
                    _ = v34
                case -9223372036854775808:
                    var v35 float32 = -2.25
                    _ = v35
                default:
                    var v36 int16 = -675
                    _ = v36
                }
                _ = fi33
            }
        }
    }
    return 2147483647
}
func genFn1(p0 int32, p1 int32) int8 {
    for fi37 := 0; fi37 < 19; fi37++ {
        v38 := []int{genFn0()}
        switch fi37 {
        case -2147483648:
            fmt.Println(p1, p0, fi37)
        case -1:
            v38 = append(v38, fi37)
        default:
            fmt.Println(p0, p1, fi37)
        }
        var v39 float32 = 3.14159
        _ = fi37
        _ = v38
        _ = v39
    }
    if false {
        for fi40 := 0; fi40 < 8; fi40++ {
            p1 -= p0
            var v41 float64 = (1e10 + float64((float32((1e10 * 0.1)) / (((3.14159 * -2.25) / 1.0) + 0.0))))
            if false {
                v42 := GenS2{F0: 625, F1: 739, F2: 1}
                _ = v42
            } else {
                p1 = p0
                v43 := GenS3{F0: "", F1: (fi40 ^ fi40)}
                var v44 bool = true
                _ = v43
                _ = v44
            }
            _ = fi40
            _ = v41
        }
    } else {
        var v45 int32 = -2147483648
        _ = v45
    }
    _ = p0
    _ = p1
    return 127
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{223, 28, 777}
    var v3 float64 = (0.1 / (((1e10 + -2.25) + 1e10) / float64(((0.1 - float32(((0.1 / 3.14159) + 1.5))) + (1e10 / float32(-2.25))))))
    var v4 int32 = 0
    var v5 uint64 = 822
    var v6 float64 = v3
    v46 := GenS4{F0: 65535, F1: "a"}
    if true {
        v47 := GenS5{F0: 65535, F1: true}
        var v48 bool = false
        _ = v47
        _ = v48
    }
    var v49 bool = false
    v6 += (v6 * (0.0 * -2.25))
    var v50 uint64 = (v5 + (v5 ^ v5))
    v50 = 18446744073709551615
    if v49 {
        v1[127] = 32767
        if (!false) {
            var v51 int = -9223372036854775808
            if (false && v49) {
                fmt.Println(v50)
                for fi52 := 0; fi52 < 14; fi52++ {
                    var v53 string = ("a" + ("\n\t" + ("foo bar" + ("" + ""))))
                    _ = fi52
                    _ = v53
                }
                v54 := v1[v51]
                _ = v54
            } else {
                switch ("a" + (("" + "a") + "")) {
                case "\n\t":
                    v55 := GenS1{F0: 651, F1: 18446744073709551615}
                    _ = v55
                case "a":
                    var v56 uint64 = v5
                    _ = v56
                case "foo bar":
                    v6 = v6
                }
                v57 := GenS3{F0: (("hello" + "hello") + ("" + ("foo bar" + "hello"))), F1: int(v4)}
                var v58 int = (^(v51))
                _ = v57
                _ = v58
            }
            _ = v51
        } else {
            switch "\n\t" {
            case "":
                fmt.Println(v49, v5)
            default:
                if false {
                    if (32767 <= 23) {
                        v59 := GenS2{F0: v5, F1: (v50 - (v50 ^ v50)), F2: 0}
                        switch 349 {
                        case -611:
                            for fi60 := 0; fi60 < 8; fi60++ {
                                v2 = append(v2, 2147483648)
                                var v61 int16 = 1000
                                for fi62 := 0; fi62 < 8; fi62++ {
                                    v6 -= v3
                                    if len(v2) > 0 {
                                        v2[0] = 122
                                    }
                                    var v63 int32 = 0
                                    _ = fi62
                                    _ = v63
                                }
                                _ = fi60
                                _ = v61
                            }
                        case -811:
                            v64 := GenS1{F0: 42, F1: v50}
                            _ = v64
                        }
                        for fi65 := 0; fi65 < 17; fi65++ {
                            var v66 float32 = 0.1
                            _ = fi65
                            _ = v66
                        }
                        _ = v59
                    }
                } else {
                    fmt.Println(v49, v50, v4)
                    v67 := GenS4{F0: 657, F1: "a"}
                    v68 := GenS6{F0: (v5 * (^(v5)))}
                    _ = v67
                    _ = v68
                }
            }
            _ = len(v1)
        }
        if ((1e10 == (-2.25 / (v6 / 1.0))) || false) {
            for fi69 := 0; fi69 < 14; fi69++ {
                switch "" {
                case "hello":
                    var v70 string = ("foo bar" + ("foo bar" + ("hello" + (("\n\t" + "foo bar") + "a"))))
                    _ = v70
                }
                _ = fi69
            }
        } else {
            for fi71 := 0; fi71 < 13; fi71++ {
                v50 = v5
                var v72 float64 = 1e10
                _ = fi71
                _ = v72
            }
            var v73 int64 = -1
            _ = len(v1)
            _ = v73
        }
    } else {
        var v74 string = ("a" + ((("" + "\n\t") + ("a" + "hello")) + ("a" + ("hello" + ""))))
        _ = v74
    }
    fmt.Println(v3, v49, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
