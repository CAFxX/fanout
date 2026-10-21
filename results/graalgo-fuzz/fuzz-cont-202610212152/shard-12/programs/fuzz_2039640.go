package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 uint64; F2 int }
func genFn0(p0 uint16, p1 int8) int8 {
    var v7 uint = 2
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            switch fi8 {
            case 4294967295:
                p1 += (p1 & 42)
            case -1:
                fmt.Println(fi9, p0, p1)
            case -749:
                var v10 uint32 = 2147483647
                _ = v10
            }
            fmt.Println(v7, v7, fi8)
            _ = fi9
        }
        p1 -= p1
        _ = fi8
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        if true {
            v12 := []int{fi11, -1}
            _ = v12
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v7
    return p1
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[string]uint32)
    var v3 uint32 = 990
    var v4 float64 = 3.14159
    v5 := []uint32{v3, 1, 2147483648}
    var v6 uint32 = v3
    v13 := GenS1{F0: -360, F1: 910, F2: 114}
    v1 = 0
    v4 += v4
    v14 := GenS1{F0: -173, F1: uint(v1), F2: -32768}
    var v15 uint16 = 42
    v16 := GenS1{F0: 42, F1: 2, F2: -1}
    v15 += (^(v15))
    v17 := GenS1{F0: -1, F1: 901, F2: -452}
    switch -502 {
    case 0:
        for fi18 := 0; fi18 < 8; fi18++ {
            if (false || true) {
                delete(v2, "\n\t")
            }
            var v19 float64 = v4
            var v20 uint8 = 27
            _ = fi18
            _ = v19
            _ = v20
        }
    case -2147483648:
        v4 -= (v4 + v4)
    default:
        var v21 int = 127
        _ = v21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        v23 := GenS2{F0: 42}
        for fi24 := 0; fi24 < 9; fi24++ {
            v25 := GenS2{F0: (fi24 + (fi24 / (((^(fi24))) | 1)))}
            var v26 int = fi24
            switch ("foo bar" + "") {
            case "a":
                delete(v2, "a")
            case "\n\t":
                if ((v4 - float64((-2.25 * 0.0))) <= (1.5 * float64(-2.25))) {
                    v3 += 1
                } else {
                    switch fi24 {
                    case 32768:
                        v26 = v26
                    default:
                        v27 := v2[""]
                        _ = v27
                    }
                    var v28 bool = (!(65535 <= 255))
                    _ = v28
                }
            case "foo bar":
                if (26 > 37) {
                    if false {
                        fmt.Println(v1)
                    }
                    _ = len(v2)
                    var v29 int = (v26 & fi24)
                    _ = v29
                } else {
                    v6 -= v6
                    v30 := GenS2{F0: fi22}
                    var v31 int16 = 42
                    var v32 float32 = -2.25
                    fmt.Println(v26)
                    if true {
                        if false {
                            if false {
                                v33 := GenS1{F0: -798, F1: 4294967295, F2: 32767}
                                _ = v33
                            } else {
                                if ((v1 * 18446744073709551615) == v1) {
                                    if false {
                                        v3 = 0
                                    }
                                } else {
                                    switch 878 {
                                    case -1:
                                        v34 := GenS3{F0: (("foo bar" + "foo bar") + "foo bar")}
                                        _ = v34
                                    case -91:
                                        if true {
                                            v35 := GenS1{F0: 1, F1: 589, F2: v31}
                                            fmt.Println(v1, v31)
                                            v6 = (v6 % (((v6 + 0)) | 1))
                                            _ = v35
                                        } else {
                                            fmt.Println(v31, v3)
                                            if false {
                                                _ = len(v2)
                                                if true {
                                                    v36 := v2[("\n\t" + (("" + "\n\t") + ("a" + "foo bar")))]
                                                    fmt.Println(v15, v32, v1)
                                                    v15 = (^(v15))
                                                    _ = v36
                                                }
                                                v15 = (v15 >> ((42) & 63))
                                            } else {
                                                var v37 uint16 = v15
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                _ = v37
                                            }
                                        }
                                    case -9223372036854775808:
                                        v3 -= (v3 ^ v3)
                                    default:
                                        fmt.Println(v1, v32, v31)
                                        if false {
                                            if len(v5) > 0 {
                                                v5[0] = (v6 % (((v3 / ((v3) | 1))) | 1))
                                            }
                                        } else {
                                            v26 = fi22
                                            if true {
                                                v26 -= fi24
                                                fmt.Println(v6, v31)
                                            } else {
                                                switch 712 {
                                                case 128:
                                                    v38 := GenS3{F0: "\n\t"}
                                                    _ = v38
                                                case 0:
                                                    v3 = v6
                                                }
                                                var v39 float64 = -2.25
                                                _ = v39
                                            }
                                            v2["foo bar"] = v3
                                            v15 += (v15 / (((v15 + v15)) | 1))
                                            _ = len(v2)
                                        }
                                    }
                                    v15 -= 112
                                }
                                v40 := GenS1{F0: 0, F1: 2, F2: 688}
                                _ = v40
                            }
                        }
                        fmt.Println(v1, fi22)
                        fmt.Println(v3)
                    } else {
                        var v41 uint = 255
                        _ = v41
                    }
                    fmt.Println(v6)
                    _ = v30
                    _ = v31
                    _ = v32
                }
            }
            _ = fi24
            _ = v25
            _ = v26
        }
        for fi42 := 0; fi42 < 4; fi42++ {
            fmt.Println(v4, fi42)
            v43 := GenS3{F0: ""}
            _ = fi42
            _ = v43
        }
        _ = fi22
        _ = v23
    }
    v15 = v15
    switch -518 {
    case 203:
        var v44 int16 = 32767
        _ = v44
    case 128:
        var v45 uint = 684
        _ = v45
    default:
        v46 := GenS2{F0: 36}
        _ = v46
    }
    for fi47 := 0; fi47 < 7; fi47++ {
        var v48 uint32 = 500
        v49 := GenS3{F0: ((("foo bar" + "hello") + "") + "foo bar")}
        if (42 == 1) {
            for fi50 := 0; fi50 < 5; fi50++ {
                fmt.Println(v6, v15)
                if ((v48 << ((393) & 63)) == (^(v48))) {
                    fmt.Println(v48)
                    var v51 uint32 = (v48 << ((18446744073709551615) & 63))
                    if false {
                        var v52 int8 = 43
                        _ = v52
                    }
                    _ = v51
                }
                fmt.Println(fi50, fi47, v4)
                v53 := GenS1{F0: -1, F1: 255, F2: -32768}
                _ = fi50
                _ = v53
            }
            v6 = v6
            v54 := GenS1{F0: -2147483648, F1: 465, F2: 379}
            _ = v54
        } else {
            v55 := GenS2{F0: fi47}
            v48 += (v3 % (((^(v3))) | 1))
            v1 += v1
            _ = v55
        }
        _ = fi47
        _ = v48
        _ = v49
    }
    v56 := v2["hello"]
    for fi57 := 0; fi57 < 11; fi57++ {
        var v58 int32 = 42
        var v59 uint64 = v1
        if ((0.1 + 1e10) < -2.25) {
            v2["a"] = 65535
            switch (fi57 + (-(fi57))) {
            case -804:
                fmt.Println(v59, v4)
            case 2147483647:
                v60 := GenS1{F0: v58, F1: 276, F2: 42}
                _ = v60
            case 242:
                if (!true) {
                    var v61 int16 = 1000
                    _ = v61
                } else {
                    if (955 >= 42) {
                        v62 := GenS4{F0: -128, F1: v59, F2: -88}
                        _ = len(v2)
                        _ = v62
                    } else {
                        fmt.Println(v59, v6)
                        v59 += uint64(v15)
                        var v63 int = fi57
                        _ = v63
                    }
                    v64 := GenS1{F0: v58, F1: 864, F2: int16(v15)}
                    _ = len(v2)
                    _ = v64
                }
            default:
                var v65 uint16 = (v15 + v15)
                _ = v65
            }
        } else {
            fmt.Println(v3)
            if (65535 >= (v6 ^ (^(v6)))) {
                switch (fi57 % ((128) | 1)) {
                case -2147483648:
                    v66 := GenS1{F0: (-(v58)), F1: 255, F2: -290}
                    _ = v66
                }
                fmt.Println(v1, v4)
                v59 += 353
            }
            v67 := GenS1{F0: v58, F1: 612, F2: 1000}
            _ = v67
        }
        _ = fi57
        _ = v58
        _ = v59
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v68 int64 = 0
    var v69 uint16 = (v15 / (((^(v15))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
    _ = v68
    _ = v69
}
