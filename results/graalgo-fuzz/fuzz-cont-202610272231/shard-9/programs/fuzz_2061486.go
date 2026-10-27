package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint; F2 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint16; F1 bool; F2 bool }
func genFn0() int32 {
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint = 65535
        for fi9 := 0; fi9 < 16; fi9++ {
            v8 += 65535
            var v10 float64 = 1.5
            var v11 int8 = 42
            v11 -= 1
            fmt.Println(fi9, v10)
            v11 += (v11 ^ v11)
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 string = ("\n\t" + ("foo bar" + ""))
        _ = fi7
        _ = v12
        _ = v8
    }
    var v13 string = ""
    _ = v13
    return 2147483647
}
func genFn1(p0 uint8, p1 int) int64 {
    for fi14 := 0; fi14 < 20; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 int16 = 0
            _ = fi15
            _ = v16
        }
        p1 = fi14
        _ = fi14
    }
    p1 = (p1 / ((42) | 1))
    p0 = (p0 / (((^(p0))) | 1))
    v17 := GenS1{F0: p1}
    _ = p0
    _ = p1
    _ = v17
    return -575
}

func main() {
    var v1 int = 128
    v2 := []uint32{}
    var v3 uint64 = 18446744073709551615
    v4 := []string{"foo bar", "a"}
    v5 := []uint32{42, 2, 65535}
    var v6 uint32 = 2147483648
    if true {
        if false {
            v18 := GenS1{F0: v1}
            for fi19 := 0; fi19 < 3; fi19++ {
                fmt.Println(v6, fi19)
                v6 += v6
                _ = fi19
            }
            switch v1 {
            case 422:
                var v20 uint32 = (v6 >> ((uint(v3)) & 63))
                _ = v20
            case 128:
                if false {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v4 = append(v4, "a")
                } else {
                    v21 := GenS2{F0: 13}
                    v1 += 469
                    _ = v21
                }
            case 0:
                var v22 int = (v1 | (v1 & v1))
                _ = v22
            default:
                var v23 float32 = (0.1 - (((-2.25 + float32(-2.25)) * (1.5 / ((0.1 / 1e10) / -2.25))) + 3.14159))
                _ = v23
            }
            _ = v18
        }
    }
    if true {
        switch ("" + "") {
        case "hello":
            v24 := GenS3{F0: 0, F1: 255, F2: true}
            _ = v24
        case "a":
            var v25 int = 42
            _ = v25
        case "foo bar":
            v1 = (^(v1))
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    if true {
        v3 = v3
        var v26 int64 = 0
        _ = v26
    } else {
        for fi27 := 0; fi27 < 7; fi27++ {
            var v28 int = v1
            fmt.Println(v6, v6, v3)
            _ = fi27
            _ = v28
        }
    }
    if true {
        var v29 uint = 599
        _ = v29
    } else {
        fmt.Println(v1)
        for fi30 := 0; fi30 < 8; fi30++ {
            var v31 bool = true
            fmt.Println(v3)
            _ = fi30
            _ = v31
        }
    }
    switch "" {
    case "\n\t":
        if true {
            if true {
                v3 = 481
                var v32 int8 = 95
                _ = v32
            }
        } else {
            var v33 string = ("a" + "foo bar")
            _ = v33
        }
    }
    v3 += 4294967295
    _ = len(v4)
    for fi34 := 0; fi34 < 2; fi34++ {
        v3 += v3
        _ = fi34
    }
    var v35 uint8 = uint8(v3)
    var v36 uint8 = v35
    v3 += v3
    fmt.Println(v3, v3, v36)
    for fi37 := 0; fi37 < 5; fi37++ {
        for fi38 := 0; fi38 < 7; fi38++ {
            var v39 uint8 = 165
            if false {
                if len(v4) > 0 {
                    v4[0] = ""
                }
                if (-584 > 1) {
                    v2 = append(v2, (v6 & 0))
                    fmt.Println(v39, fi37)
                    if (1 > 23) {
                        if (v6 <= v6) {
                            switch ("a" + (("foo bar" + "hello") + ("a" + "\n\t"))) {
                            case "hello":
                                if ((2 < 92) || (42 <= -9223372036854775808)) {
                                    if true {
                                        v40 := GenS1{F0: 127}
                                        _ = v40
                                    } else {
                                        v41 := GenS4{F0: (-158 == 42)}
                                        if len(v5) > 0 {
                                            v5[0] = (v6 / (((v6 & 0)) | 1))
                                        }
                                        switch (fi38 ^ int(v3)) {
                                        case 0:
                                            var v42 uint64 = (^(v3))
                                            _ = v42
                                        case 128:
                                            var v43 int = fi38
                                            _ = v43
                                        case 492:
                                            switch "foo bar" {
                                            case "hello":
                                                var v44 string = ("" + ("" + "a"))
                                                _ = v44
                                            }
                                        default:
                                            if (-1 > 42) {
                                                v1 += 42
                                                if true {
                                                    v3 -= 4294967295
                                                    var v45 int64 = 9223372036854775807
                                                    _ = v45
                                                } else {
                                                    v36 += (v39 % (((v35 | (^(v39)))) | 1))
                                                    v46 := GenS5{F0: ("" + ("foo bar" + ("\n\t" + "foo bar")))}
                                                    _ = v46
                                                }
                                            }
                                        }
                                        switch "hello" {
                                        case "\n\t":
                                            var v47 int32 = 0
                                            _ = v47
                                        default:
                                            var v48 bool = (255 == 783)
                                            _ = v48
                                        }
                                        var v49 int = v1
                                        var v50 float64 = 1.5
                                        if false {
                                            v1 = 9223372036854775807
                                        } else {
                                            v51 := GenS6{F0: (v3 % (((^(v3))) | 1)), F1: 127}
                                            _ = v51
                                        }
                                        v52 := GenS7{F0: 42}
                                        v6 = (v6 >> ((790) & 63))
                                        _ = v41
                                        _ = v49
                                        _ = v50
                                        _ = v52
                                    }
                                    var v53 uint16 = 2
                                    v53 = 327
                                    _ = v53
                                } else {
                                    var v54 int16 = 449
                                    if ((v3 >= v3) && (0 == -1)) {
                                        fmt.Println(v35, v54)
                                        v36 = 255
                                        if false {
                                            v6 += 2147483647
                                            var v55 uint16 = 777
                                            var v56 int32 = -2147483648
                                            v3 += uint64(v35)
                                            var v57 uint8 = (v35 / ((2) | 1))
                                            _ = v55
                                            _ = v56
                                            _ = v57
                                        }
                                        var v58 bool = (!false)
                                        v5 = append(v5, (^(v6)))
                                        _ = v58
                                    }
                                    v5 = append(v5, v6)
                                    _ = v54
                                }
                            case "foo bar":
                                if len(v2) > 0 {
                                    v2[0] = uint32(v3)
                                }
                            default:
                                v59 := GenS1{F0: (^(fi38))}
                                _ = v59
                            }
                            fmt.Println(fi37, v3)
                        }
                        v6 += v6
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        v60 := GenS5{F0: ("a" + "\n\t")}
                        _ = v60
                    }
                    var v61 float64 = 1e10
                    if len(v4) > 0 {
                        v4[0] = ("\n\t" + ("hello" + ""))
                    }
                    v36 += 46
                    if (true && true) {
                        switch 244 {
                        case 0:
                            var v62 int8 = 38
                            _ = v62
                        case 897:
                            v61 = v61
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        if (1 != 1) {
                            fmt.Println(fi37, v61, v3)
                            switch (-(fi38)) {
                            case 2147483647:
                                if len(v2) > 0 {
                                    v2[0] = (v6 ^ (v6 % (((v6 - 4294967295)) | 1)))
                                }
                            }
                            v63 := GenS7{F0: 42}
                            var v64 uint32 = 110
                            _ = v63
                            _ = v64
                        } else {
                            fmt.Println(v3)
                            switch (fi38 << ((2) & 63)) {
                            case 128:
                                v1 = (-(fi38))
                                v35 = v36
                            case 42:
                                switch ("hello" + ("" + (("foo bar" + "") + ("hello" + "hello")))) {
                                case "\n\t":
                                    if true {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v3 = (^(v3))
                                    }
                                }
                            }
                        }
                    }
                    _ = v61
                }
            } else {
                v5 = append(v5, (v6 & 65535))
                v3 = (v3 >> ((18446744073709551615) & 63))
                v65 := GenS8{F0: 234, F1: ((^(v36)) < v35), F2: (42 != 0)}
                _ = v65
            }
            _ = fi38
            _ = v39
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
