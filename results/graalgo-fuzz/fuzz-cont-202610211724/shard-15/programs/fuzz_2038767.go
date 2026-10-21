package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 string; F1 uint8 }
func genFn0() int16 {
    if true {
        var v4 int32 = 42
        _ = v4
    } else {
        var v5 int8 = 42
        v5 += v5
        var v6 int = -188
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 18446744073709551615, F1: ("a" + "foo bar")}
    v8 := []int{}
    v9 := GenS1{F0: 242, F1: "\n\t"}
    var v10 int64 = 1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 359
}
func genFn1(p0 string, p1 uint8, p2 string) int32 {
    for fi11 := 0; fi11 < 1; fi11++ {
        p1 = 246
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 float32 = (0.1 / 3.14159)
            v14 := GenS1{F0: 0, F1: p0}
            v15 := GenS1{F0: 0, F1: p0}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi11
    }
    var v16 int64 = -255
    v17 := []int{}
    var v18 int32 = -1
    switch 866 {
    case 0:
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
    case 4294967295:
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
    case 9223372036854775807:
        v19 := GenS1{F0: 4294967295, F1: "\n\t"}
        _ = v19
    default:
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v18
    return v18
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 42
    v3 := make(map[int]bool)
    switch 128 {
    case 683:
        v2 -= v2
    case 460:
        if true {
            v2 = 0
            var v20 uint = 42
            v21 := GenS2{F0: 9223372036854775807, F1: 42}
            _ = v20
            _ = v21
        } else {
            var v22 uint32 = 65535
            _ = v22
        }
    default:
        for fi23 := 0; fi23 < 11; fi23++ {
            if true {
                var v24 int = fi23
                _ = v24
            }
            _ = fi23
        }
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        v2 += (v2 + v2)
        _ = fi25
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        if (true || (!true)) {
            v27 := GenS3{F0: 414}
            _ = v27
        }
        _ = fi26
    }
    var v28 string = ""
    switch -9223372036854775808 {
    case 128:
        v28 = ("a" + (v28 + "a"))
    case 2147483647:
        for fi29 := 0; fi29 < 16; fi29++ {
            v30 := []int{(fi29 << ((255) & 63)), (fi29 / ((fi29) | 1))}
            v28 = (v28 + "")
            if false {
                for fi31 := 0; fi31 < 20; fi31++ {
                    var v32 string = (("foo bar" + "\n\t") + ("foo bar" + ("\n\t" + (("hello" + v28) + v28))))
                    var v33 uint16 = 65535
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                for fi34 := 0; fi34 < 8; fi34++ {
                    _ = len(v30)
                    v28 = v28
                    _ = fi34
                }
            }
            _ = fi29
            _ = v30
        }
    case 127:
        var v35 uint32 = 1
        _ = v35
    }
    if ((74 > v2) && true) {
        var v36 uint = 65535
        _ = v36
    } else {
        fmt.Println(v28)
    }
    var v37 uint8 = (v2 - 255)
    if (2147483648 <= 286) {
        if (true && false) {
            if true {
                v2 += v2
                switch (v28 + v28) {
                case "hello":
                    v28 = v28
                }
                var v38 float32 = -2.25
                _ = v38
            } else {
                v28 = ((v28 + (v28 + v28)) + v28)
                var v39 uint16 = 0
                _ = v39
            }
            v40 := []int{42}
            v41 := GenS4{F0: ((-128 != 1) && true)}
            _ = v40
            _ = v41
        } else {
            switch 1 {
            case 9223372036854775807:
                var v42 uint16 = 546
                _ = v42
            }
            var v43 int16 = 0
            _ = v43
        }
    }
    var v44 uint8 = (v2 * v2)
    if (42 <= 1) {
        if true {
            if (!((92 < 87) && true)) {
                if false {
                    if (float64(0.0) <= 1.5) {
                        for fi45 := 0; fi45 < 18; fi45++ {
                            var v46 float32 = 1.5
                            if (122 != 100) {
                                v44 = (v37 ^ (v2 << ((886) & 63)))
                                fmt.Println(v37)
                                v46 = 0.1
                            }
                            var v47 bool = true
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                        var v48 uint32 = 0
                        _ = v48
                    } else {
                        var v49 int = -796
                        _ = v49
                    }
                    fmt.Println(v37, v44, v28)
                    v50 := v3[-803]
                    _ = v50
                } else {
                    v44 = (v2 + 204)
                    if true {
                        v51 := v1[277]
                        if true {
                            v28 = v51
                            if false {
                                v52 := []int{42, 2147483647, 0}
                                _ = v52
                            } else {
                                if (42 < 164) {
                                    if true {
                                        for fi53 := 0; fi53 < 10; fi53++ {
                                            v37 -= v2
                                            _ = fi53
                                        }
                                        var v54 uint = uint(v44)
                                        var v55 uint32 = 0
                                        _ = v54
                                        _ = v55
                                    }
                                    var v56 string = v51
                                    _ = v56
                                }
                                var v57 bool = false
                                var v58 float64 = 3.14159
                                _ = v57
                                _ = v58
                            }
                            v59 := v1[2147483648]
                            _ = v59
                        }
                        for fi60 := 0; fi60 < 3; fi60++ {
                            switch (fi60 >> ((255) & 63)) {
                            case -770:
                                v51 = v28
                            default:
                                if false {
                                    v44 = 64
                                    if (-47 > 75) {
                                        v61 := GenS5{F0: 455, F1: v28}
                                        _ = v61
                                    }
                                    var v62 bool = false
                                    _ = v62
                                } else {
                                    var v63 float64 = 1e10
                                    _ = v63
                                }
                            }
                            _ = fi60
                        }
                        _ = v51
                    }
                    var v64 int64 = -1
                    _ = v64
                }
            }
            v28 = (("\n\t" + (v28 + v28)) + "\n\t")
        }
    } else {
        fmt.Println(v2, v44)
        switch v28 {
        case "hello":
            fmt.Println(v2, v2, v44)
        case "\n\t":
            for fi65 := 0; fi65 < 4; fi65++ {
                v66 := GenS6{F0: v28, F1: (v37 >> ((614) & 63))}
                for fi67 := 0; fi67 < 8; fi67++ {
                    var v68 uint8 = (v2 / ((v44) | 1))
                    var v69 int = (fi65 - fi67)
                    delete(v3, (v69 + (fi67 % (((fi65 % (((fi65 % (((fi65 & 127)) | 1))) | 1))) | 1))))
                    _ = fi67
                    _ = v68
                    _ = v69
                }
                _ = fi65
                _ = v66
            }
        case "":
            v2 = (^(v37))
        default:
            switch 2147483647 {
            case 2147483647:
                for fi70 := 0; fi70 < 16; fi70++ {
                    v71 := []int{(fi70 * fi70), 1, (fi70 % ((128) | 1))}
                    _ = fi70
                    _ = v71
                }
            case 32768:
                _ = len(v3)
            case 872:
                fmt.Println(v28)
            }
        }
    }
    v1[0] = ((((v28 + "foo bar") + v28) + v28) + v28)
    v37 = v2
    var v72 float32 = 1e10
    if (42 <= 760) {
        fmt.Println(v2)
        if (-2147483648 != 376) {
            v37 = v37
            for fi73 := 0; fi73 < 19; fi73++ {
                var v74 uint32 = 182
                for fi75 := 0; fi75 < 6; fi75++ {
                    if false {
                        v76 := v1[v74]
                        _ = v76
                    } else {
                        var v77 uint32 = 61
                        v37 = v44
                        var v78 float32 = (float32((3.14159 - ((-2.25 + 3.14159) + 1.5))) / float32(-2.25))
                        _ = v77
                        _ = v78
                    }
                    _ = fi75
                }
                _ = fi73
                _ = v74
            }
            var v79 int16 = 0
            _ = v79
        } else {
            _ = len(v3)
        }
        var v80 bool = true
        _ = v80
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v44:", v44)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v37
    _ = v44
    _ = v72
}
