package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint64; F1 int8; F2 uint64 }
func genFn0(p0 uint64) bool {
    p0 = p0
    if false {
        if false {
            v7 := []int{32768}
            var v8 uint32 = 2
            _ = v7
            _ = v8
        } else {
            fmt.Println(p0)
            p0 += p0
            if true {
                var v9 string = "foo bar"
                _ = v9
            } else {
                v10 := GenS1{F0: 1, F1: 1, F2: 1}
                _ = v10
            }
        }
        for fi11 := 0; fi11 < 6; fi11++ {
            fmt.Println(p0, p0)
            for fi12 := 0; fi12 < 18; fi12++ {
                v13 := GenS1{F0: 2147483647, F1: -109, F2: 18446744073709551615}
                v14 := GenS2{F0: 2147483648}
                _ = fi12
                _ = v13
                _ = v14
            }
            switch (("a" + ("a" + "hello")) + ((("foo bar" + "a") + "foo bar") + ("hello" + "foo bar"))) {
            case "foo bar":
                v15 := GenS2{F0: 219}
                _ = v15
            case "a":
                for fi16 := 0; fi16 < 8; fi16++ {
                    if true {
                        fmt.Println(fi16, p0, fi16)
                        var v17 int8 = 0
                        fmt.Println(fi16, fi16)
                        _ = v17
                    } else {
                        v18 := GenS3{F0: -128, F1: (true || ("a" >= "a"))}
                        fmt.Println(fi11, fi11)
                        _ = v18
                    }
                    var v19 uint8 = 42
                    _ = fi16
                    _ = v19
                }
            default:
                var v20 int32 = 2147483647
                _ = v20
            }
            _ = fi11
        }
    } else {
        for fi21 := 0; fi21 < 6; fi21++ {
            if (-9223372036854775808 >= 0) {
                v22 := []int{995, 32767, (fi21 * fi21), fi21}
                for fi23 := 0; fi23 < 17; fi23++ {
                    var v24 uint32 = 42
                    if (true || (42 != 8)) {
                        var v25 uint32 = v24
                        v26 := GenS4{F0: -9223372036854775808, F1: (fi21 / ((fi23) | 1))}
                        var v27 string = "foo bar"
                        var v28 bool = true
                        _ = v25
                        _ = v26
                        _ = v27
                        _ = v28
                    } else {
                        v24 += 2147483648
                        var v29 int8 = 79
                        var v30 uint = 18446744073709551615
                        _ = v29
                        _ = v30
                    }
                    var v31 uint = 1
                    _ = fi23
                    _ = v24
                    _ = v31
                }
                var v32 uint8 = 238
                _ = v22
                _ = v32
            } else {
                var v33 uint16 = 0
                v34 := []int{}
                _ = v33
                _ = v34
            }
            _ = fi21
        }
    }
    _ = p0
    return true
}

func main() {
    var v1 uint64 = 2
    var v2 float64 = (-2.25 + 3.14159)
    v3 := []uint32{65535, 2}
    var v4 uint32 = 2147483648
    v5 := make(map[uint32]string)
    var v6 string = "a"
    _ = len(v3)
    v1 -= v1
    var v35 uint32 = 13
    if false {
        if (false && ((v6 + ("" + "foo bar")) == ("" + ""))) {
            for fi36 := 0; fi36 < 18; fi36++ {
                _ = len(v3)
                _ = fi36
            }
            if (false && true) {
                var v37 float32 = ((0.0 / 1.5) * ((1e10 / (0.0 + (0.1 - ((1e10 * (1e10 + 0.0)) * (1.5 + 1e10))))) / (3.14159 - (0.1 + 1e10))))
                var v38 float64 = 1.5
                var v39 float32 = v37
                _ = v37
                _ = v38
                _ = v39
            }
        }
        v3 = append(v3, 0)
        fmt.Println(v35, v35)
    }
    var v40 uint32 = (^(v35))
    var v41 int8 = 28
    v42 := GenS1{F0: v40, F1: -490, F2: 0}
    for fi43 := 0; fi43 < 5; fi43++ {
        if true {
            v44 := GenS5{F0: 255}
            v4 = v4
            var v45 string = (((v6 + (v6 + "foo bar")) + "a") + "")
            _ = v44
            _ = v45
        } else {
            var v46 uint64 = v1
            for fi47 := 0; fi47 < 6; fi47++ {
                var v48 uint = 1
                v49 := GenS3{F0: 0, F1: true}
                switch fi43 {
                case 127:
                    if false {
                        v2 = 0.1
                    }
                case 277:
                    switch ("\n\t" + v6) {
                    case "hello":
                        v4 = v4
                    default:
                        var v50 int64 = 172
                        _ = v50
                    }
                }
                _ = fi47
                _ = v48
                _ = v49
            }
            for fi51 := 0; fi51 < 1; fi51++ {
                var v52 float64 = 3.14159
                _ = fi51
                _ = v52
            }
            _ = v46
        }
        if false {
            switch (-(fi43)) {
            case 1:
                v4 = (v40 * (^(v4)))
            case 9223372036854775807:
                fmt.Println(v41, v6)
            case 127:
                var v53 bool = (v35 > 2)
                _ = v53
            default:
                v54 := v5[(v40 % (((v35 | (v40 / (((v40 ^ (v4 % ((2) | 1)))) | 1)))) | 1))]
                _ = v54
            }
        } else {
            var v55 float64 = v2
            v40 = 128
            for fi56 := 0; fi56 < 7; fi56++ {
                if (2147483647 > 765) {
                    v40 += v4
                    if (v4 <= (v4 << ((42) & 63))) {
                        v57 := GenS3{F0: (v41 >> ((uint(v1)) & 63)), F1: true}
                        _ = v57
                    } else {
                        v3 = append(v3, (^(v4)))
                    }
                }
                _ = fi56
            }
            _ = v55
        }
        _ = fi43
    }
    for fi58 := 0; fi58 < 15; fi58++ {
        for fi59 := 0; fi59 < 17; fi59++ {
            v1 += v1
            _ = fi59
        }
        if true {
            var v60 string = v6
            _ = v60
        } else {
            v2 += v2
        }
        _ = fi58
    }
    if false {
        if ((-1 <= 32767) || false) {
            v61 := GenS2{F0: v35}
            v62 := GenS6{F0: 42}
            _ = v61
            _ = v62
        } else {
            _ = len(v5)
            v35 += v35
        }
    }
    v35 = 131
    switch -9223372036854775808 {
    case -754:
        if (!(false || true)) {
            var v63 uint32 = 679
            v64 := v5[v63]
            var v65 bool = (0.0 != 3.14159)
            _ = v63
            _ = v64
            _ = v65
        } else {
            switch -2147483648 {
            case 914:
                fmt.Println(v40, v40, v4)
            }
            if (false || (0 >= 255)) {
                for fi66 := 0; fi66 < 5; fi66++ {
                    if true {
                        var v67 int8 = 1
                        fmt.Println(v35)
                        var v68 int16 = -170
                        _ = v67
                        _ = v68
                    } else {
                        var v69 uint = 255
                        v35 = (v4 | 2147483647)
                        _ = v69
                    }
                    _ = len(v3)
                    for fi70 := 0; fi70 < 8; fi70++ {
                        fmt.Println(v2)
                        if ((((42 < 1) || true) || (!(42 != 181))) || true) {
                            switch (fi70 / (((^(fi66))) | 1)) {
                            case -9223372036854775808:
                                fmt.Println(v41, v2, fi70)
                            default:
                                _ = len(v3)
                            }
                            v71 := GenS6{F0: 539}
                            _ = v71
                        } else {
                            v40 = 753
                        }
                        _ = fi70
                    }
                    _ = fi66
                }
                for fi72 := 0; fi72 < 10; fi72++ {
                    v35 = v4
                    switch "foo bar" {
                    case "hello":
                        v5[(^(v40))] = "foo bar"
                    case "a":
                        var v73 float32 = ((3.14159 + 1e10) / float32(float64(float32(1.5))))
                        _ = v73
                    default:
                        v74 := GenS3{F0: v41, F1: false}
                        _ = v74
                    }
                    var v75 float32 = ((0.1 + 0.1) / 3.14159)
                    _ = fi72
                    _ = v75
                }
            } else {
                var v76 uint = 18446744073709551615
                v77 := GenS5{F0: 788}
                delete(v5, 0)
                _ = v76
                _ = v77
            }
            for fi78 := 0; fi78 < 12; fi78++ {
                v3 = append(v3, v40)
                _ = fi78
            }
        }
    }
    switch v6 {
    case "a":
        switch -2147483648 {
        case 4294967295:
            var v79 int8 = (v41 % (((v41 % (((v41 >> ((1) & 63))) | 1))) | 1))
            _ = v79
        }
    case "hello":
        if (!((float64(1e10) - v2) == 1.5)) {
            if (true || false) {
                v1 = uint64(v4)
                switch -1 {
                case -1:
                    if (!false) {
                        v80 := GenS7{F0: (^(v1)), F1: (v41 >> ((61) & 63)), F2: v1}
                        v4 = v4
                        _ = v80
                    }
                default:
                    if ((((3.14159 * -2.25) + 3.14159) + 1.5) <= 1.5) {
                        _ = len(v3)
                    }
                }
            } else {
                var v81 bool = ((false && (-32768 < 174)) && false)
                _ = v81
            }
            delete(v5, v40)
        }
    case "foo bar":
        v1 -= v1
    }
    var v82 float32 = 1.5
    var v83 uint8 = 42
    if (3.14159 <= (v82 / v82)) {
        switch "a" {
        case "hello":
            v83 += (v83 / ((v83) | 1))
        case "foo bar":
            v83 -= (v83 + 10)
        }
    }
    var v84 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v82
    _ = v83
    _ = v84
}
