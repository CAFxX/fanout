package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint8; F2 uint }
type GenS4 struct { F0 int16; F1 int32; F2 uint16 }
type GenS5 struct { F0 int32; F1 int16; F2 uint }

func main() {
    var v1 uint32 = 254
    var v2 uint = 2
    v3 := []string{"", "a", ("a" + ("foo bar" + "foo bar")), "foo bar"}
    var v4 float64 = 0.0
    v5 := make(map[uint32]uint32)
    var v6 int32 = 937
    for fi7 := 0; fi7 < 5; fi7++ {
        if false {
            v3 = append(v3, ("foo bar" + (("hello" + "a") + ("foo bar" + "a"))))
            var v8 int8 = int8(v6)
            _ = v8
        }
        v6 = (v6 << (((v2 * v2)) & 63))
        _ = fi7
    }
    if true {
        fmt.Println(v1, v4)
        switch "hello" {
        case "foo bar":
            switch (("a" + ("hello" + "hello")) + "\n\t") {
            case "a":
                for fi9 := 0; fi9 < 3; fi9++ {
                    fmt.Println(v2, v4, fi9)
                    _ = fi9
                }
            case "":
                var v10 uint64 = 1
                _ = v10
            default:
                if ((v1 | v1) == (v1 % (((v1 ^ v1)) | 1))) {
                    if ((true && (561 > 9223372036854775807)) && false) {
                        var v11 int8 = -1
                        v12 := GenS1{F0: (v2 & v2), F1: 1}
                        var v13 bool = false
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    switch 32767 {
                    case 4294967295:
                        if false {
                            var v14 uint = (v2 ^ v2)
                            _ = v14
                        } else {
                            for fi15 := 0; fi15 < 20; fi15++ {
                                v2 = v2
                                _ = fi15
                            }
                        }
                    }
                } else {
                    if (false || true) {
                        v6 -= -2147483648
                        var v16 int8 = -128
                        _ = len(v5)
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 11; fi17++ {
                            v6 = (v6 * -2147483648)
                            if false {
                                var v18 int16 = 688
                                v6 = v6
                                _ = v18
                            } else {
                                v19 := GenS1{F0: v2, F1: 413}
                                var v20 bool = (!(v1 < v1))
                                var v21 int32 = v6
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            for fi22 := 0; fi22 < 11; fi22++ {
                                v23 := GenS2{F0: (-32768 >= 1000)}
                                fmt.Println(v4, v1)
                                v3 = append(v3, "\n\t")
                                _ = fi22
                                _ = v23
                            }
                            _ = fi17
                        }
                        v24 := GenS1{F0: (v2 ^ (v2 + v2)), F1: 0}
                        v25 := GenS3{F0: 2147483647, F1: 83, F2: v2}
                        _ = v24
                        _ = v25
                    }
                    v26 := GenS2{F0: ("\n\t" < (("a" + "\n\t") + ("a" + ("\n\t" + ""))))}
                    v1 = v1
                    _ = v26
                }
            }
        case "hello":
            for fi27 := 0; fi27 < 10; fi27++ {
                if ((v6 == v6) && true) {
                    var v28 uint32 = 42
                    _ = v28
                }
                v2 -= v2
                _ = fi27
            }
        case "a":
            if len(v3) > 0 {
                v3[0] = "a"
            }
        default:
            var v29 uint64 = 4294967295
            _ = v29
        }
        for fi30 := 0; fi30 < 12; fi30++ {
            if false {
                v4 += 0.0
            } else {
                fmt.Println(v4, v2)
                fmt.Println(v4, v1)
            }
            if (((v6 << (((v2 / ((uint(v1)) | 1))) & 63)) != (v6 + 42)) || (-9223372036854775808 <= 43)) {
                var v31 uint = (v2 >> ((v2) & 63))
                _ = v31
            } else {
                fmt.Println(fi30, fi30)
                v4 -= float64((-2.25 - (0.0 / 1e10)))
                for fi32 := 0; fi32 < 8; fi32++ {
                    fmt.Println(v4, v2, fi32)
                    v33 := GenS4{F0: -1, F1: (v6 << ((v2) & 63)), F2: 2}
                    var v34 int = (fi30 ^ (fi30 & fi30))
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
            _ = fi30
        }
    } else {
        if false {
            var v35 uint8 = 2
            var v36 float64 = (1e10 / v4)
            for fi37 := 0; fi37 < 6; fi37++ {
                fmt.Println(v6, v4, fi37)
                for fi38 := 0; fi38 < 7; fi38++ {
                    var v39 int = 0
                    switch "foo bar" {
                    case "foo bar":
                        var v40 bool = false
                        _ = len(v3)
                        if (2 > 127) {
                            var v41 int = 0
                            var v42 float64 = ((0.1 / (((float64(float32(float64(3.14159))) - v4) * (v4 / float64(0.1))) - ((v4 + v4) * v36))) * (float64((0.1 / 1e10)) / v36))
                            _ = v41
                            _ = v42
                        } else {
                            switch (^(fi38)) {
                            case -513:
                                switch (-(v39)) {
                                case 127:
                                    v43 := GenS5{F0: 0, F1: 32767, F2: v2}
                                    _ = v43
                                }
                                switch (v39 * (^(fi37))) {
                                case 32768:
                                    switch "a" {
                                    case "":
                                        v5[161] = v1
                                    case "\n\t":
                                        var v44 float32 = float32(float64(float32(((v36 / v4) * float64((-2.25 + 0.0))))))
                                        _ = v44
                                    case "a":
                                        var v45 int = v39
                                        _ = v45
                                    }
                                case 247:
                                    var v46 float64 = float64((-2.25 * 1e10))
                                    _ = v46
                                }
                                v47 := GenS3{F0: v6, F1: v35, F2: v2}
                                var v48 uint8 = v35
                                fmt.Println(fi38, v4, v1)
                                _ = v47
                                _ = v48
                            default:
                                if ((!v40) && (true && v40)) {
                                    v3 = append(v3, "hello")
                                } else {
                                    var v49 uint32 = 638
                                    _ = v49
                                }
                            }
                        }
                        _ = v40
                    default:
                        v2 -= v2
                    }
                    var v50 int = (v39 ^ v39)
                    _ = fi38
                    _ = v39
                    _ = v50
                }
                fmt.Println(v1)
                _ = fi37
            }
            _ = v35
            _ = v36
        }
        v51 := GenS3{F0: 891, F1: uint8(v1), F2: 18446744073709551615}
        if false {
            for fi52 := 0; fi52 < 16; fi52++ {
                var v53 int64 = 9223372036854775807
                _ = fi52
                _ = v53
            }
        } else {
            v2 -= v2
            var v54 uint32 = (v1 - (v1 / (((v1 / ((v1) | 1))) | 1)))
            _ = v54
        }
        _ = v51
    }
    var v55 uint64 = 0
    v1 += v1
    v2 = v2
    switch 32767 {
    case -2147483648:
        fmt.Println(v6, v4, v55)
    }
    if false {
        fmt.Println(v4, v2, v1)
    }
    fmt.Println(v2, v2, v55)
    v3 = append(v3, "foo bar")
    fmt.Println(v55, v2)
    fmt.Println(v6, v4, v2)
    v6 += (v6 / ((1) | 1))
    if (true && true) {
        var v56 int64 = 42
        v55 = v55
        _ = v56
    } else {
        v1 += 65535
        var v57 int64 = 951
        for fi58 := 0; fi58 < 7; fi58++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if ((3.14159 * 0.0) <= -2.25) {
                if ((v57 >> ((v2) & 63)) >= v57) {
                    delete(v5, 2147483647)
                    var v59 bool = (false || (false && false))
                    v60 := GenS5{F0: v6, F1: -1, F2: 255}
                    _ = v59
                    _ = v60
                }
            } else {
                delete(v5, (^(v1)))
                v61 := GenS2{F0: (228 <= 42)}
                _ = v61
            }
            _ = fi58
        }
        _ = v57
    }
    v55 -= v55
    var v62 int8 = int8(v6)
    fmt.Println(v4, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
    _ = v62
}
