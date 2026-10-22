package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint; F1 int32; F2 uint }
type GenS4 struct { F0 int32; F1 bool; F2 uint64 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 uint64, p1 uint64, p2 int32) string {
    fmt.Println(p2)
    switch "a" {
    case "foo bar":
        if (((p0 * (p0 + (p1 >> ((255) & 63)))) >= p0) || true) {
            for fi7 := 0; fi7 < 19; fi7++ {
                v8 := GenS1{F0: "\n\t"}
                v9 := GenS2{F0: 2, F1: ((p1 << ((4294967295) & 63)) >= p1), F2: 1}
                _ = fi7
                _ = v8
                _ = v9
            }
            fmt.Println(p1)
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 int64 = -1
                switch fi10 {
                case 4294967295:
                    if false {
                        switch "hello" {
                        case "hello":
                            var v12 bool = ((4294967295 < 0) || (true || ((p1 ^ (p1 + 4294967295)) != (p1 << ((255) & 63)))))
                            _ = v12
                        }
                    } else {
                        v13 := []int{fi10}
                        if true {
                            var v14 string = ("" + "foo bar")
                            _ = v14
                        }
                        _ = v13
                    }
                default:
                    fmt.Println(v11, v11)
                }
                _ = fi10
                _ = v11
            }
        }
    case "hello":
        var v15 float64 = (1e10 + (((((-2.25 * 0.0) - (1.5 - 0.1)) - 0.0) / 0.1) / 1.5))
        _ = v15
    case "":
        v16 := GenS2{F0: 152, F1: ((p1 << ((18446744073709551615) & 63)) == p0), F2: 1}
        _ = v16
    default:
        if true {
            v17 := []int{2147483647, 32768}
            _ = v17
        } else {
            for fi18 := 0; fi18 < 12; fi18++ {
                if ((p0 | (^(p1))) == p0) {
                    v19 := GenS3{F0: 255, F1: p2, F2: 0}
                    switch "" {
                    case "\n\t":
                        p0 = p0
                    case "hello":
                        var v20 uint16 = 799
                        _ = v20
                    default:
                        v21 := []int{}
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 7; fi22++ {
                        if true {
                            var v23 int16 = 32767
                            fmt.Println(p0, p1)
                            if (false || (0 != v23)) {
                                fmt.Println(v23, p1)
                                switch "hello" {
                                case "foo bar":
                                    v24 := []int{698, fi18, 4294967295, (fi22 << ((4294967295) & 63))}
                                    fmt.Println(v23, fi18)
                                    if ((!true) && ((!false) && true)) {
                                        v25 := GenS4{F0: p2, F1: (-9223372036854775808 >= 9223372036854775807), F2: (p1 ^ p1)}
                                        _ = v25
                                    }
                                    _ = v24
                                default:
                                    p0 += (p0 / (((p1 * 1)) | 1))
                                }
                            } else {
                                var v26 int64 = -544
                                v26 = (v26 & 1)
                                switch "" {
                                case "foo bar":
                                    var v27 float32 = -2.25
                                    _ = v27
                                case "\n\t":
                                    v26 = v26
                                case "hello":
                                    v23 += v23
                                default:
                                    if (1 == 822) {
                                        if true {
                                            var v28 int8 = -1
                                            _ = v28
                                        } else {
                                            p2 = (p2 | (p2 + p2))
                                        }
                                    } else {
                                        v29 := []int{-725, 259, (fi18 >> ((239) & 63)), (fi18 & (fi18 | (-(fi18))))}
                                        _ = v29
                                    }
                                }
                                _ = v26
                            }
                            _ = v23
                        }
                        _ = fi22
                    }
                    _ = v19
                }
                _ = fi18
            }
            for fi30 := 0; fi30 < 8; fi30++ {
                var v31 float32 = 0.1
                v32 := []int{fi30, fi30}
                _ = fi30
                _ = v31
                _ = v32
            }
        }
    }
    if (!false) {
        v33 := GenS3{F0: 891, F1: p2, F2: 65535}
        switch int(p2) {
        case 2147483647:
            for fi34 := 0; fi34 < 20; fi34++ {
                p2 = (p2 >> ((539) & 63))
                var v35 int16 = 0
                _ = fi34
                _ = v35
            }
        default:
            switch (("a" + (("hello" + (("hello" + "\n\t") + "foo bar")) + "hello")) + "") {
            case "\n\t":
                v36 := []int{-254, 42}
                _ = v36
            case "foo bar":
                v37 := []int{}
                _ = v37
            default:
                v38 := GenS5{F0: (p0 + p1), F1: 792}
                _ = v38
            }
        }
        _ = v33
    }
    p0 -= p1
    var v39 int16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v39
    return ("foo bar" + ("a" + "foo bar"))
}
func genFn1(p0 uint, p1 int8) int64 {
    for fi40 := 0; fi40 < 16; fi40++ {
        p1 = p1
        _ = fi40
    }
    var v41 uint8 = 1
    p1 -= (p1 | p1)
    var v42 float32 = 0.0
    _ = p0
    _ = p1
    _ = v41
    _ = v42
    return 1
}

func main() {
    var v1 uint = 255
    var v2 uint32 = 2
    var v3 uint = v1
    v4 := make(map[string]bool)
    var v5 float32 = 0.0
    v6 := make(map[int]string)
    v5 += 0.1
    var v43 uint = v1
    v1 += (v1 / ((v1) | 1))
    v43 = v1
    switch "" {
    case "a":
        v44 := []int{-2147483648, 74, 467, -771}
        _ = v44
    case "\n\t":
        fmt.Println(v43, v1)
    case "hello":
        var v45 bool = false
        _ = v45
    default:
        for fi46 := 0; fi46 < 16; fi46++ {
            for fi47 := 0; fi47 < 6; fi47++ {
                var v48 float32 = v5
                v49 := GenS6{F0: 4294967295, F1: 113, F2: 59}
                var v50 uint16 = 42
                _ = len(v4)
                _ = fi47
                _ = v48
                _ = v49
                _ = v50
            }
            for fi51 := 0; fi51 < 12; fi51++ {
                fmt.Println(v43, v43, fi46)
                _ = fi51
            }
            _ = fi46
        }
    }
    switch -1 {
    case 9223372036854775807:
        fmt.Println(v5, v2, v5)
    default:
        var v52 int = 9223372036854775807
        _ = v52
    }
    v5 = 1.5
    var v53 uint8 = uint8(v43)
    v54 := []int{128, 579, 9223372036854775807}
    v2 -= 0
    v53 = 1
    v43 -= 65535
    if (2 > 0) {
        delete(v4, ("hello" + (("hello" + "\n\t") + ("hello" + "\n\t"))))
    } else {
        switch -1 {
        case -9223372036854775808:
            v55 := GenS3{F0: v3, F1: -2147483648, F2: v43}
            _ = v55
        case 4294967295:
            v43 -= 694
        default:
            for fi56 := 0; fi56 < 12; fi56++ {
                var v57 float64 = 0.1
                _ = fi56
                _ = v57
            }
        }
    }
    var v58 uint = 65535
    var v59 int16 = 1
    switch ((("hello" + "hello") + "a") + ((("\n\t" + (("\n\t" + "") + "\n\t")) + "\n\t") + "foo bar")) {
    case "":
        if true {
            for fi60 := 0; fi60 < 20; fi60++ {
                if len(v54) >= 2 {
                    v54 = v54[1:len(v54)]
                }
                v61 := GenS5{F0: 243, F1: 975}
                _ = fi60
                _ = v61
            }
            for fi62 := 0; fi62 < 17; fi62++ {
                if (9223372036854775807 > 307) {
                    switch fi62 {
                    case 2147483647:
                        if (128 >= 127) {
                            var v63 uint16 = 575
                            _ = v63
                        } else {
                            if (42 == -2147483648) {
                                if false {
                                    v58 -= v1
                                } else {
                                    fmt.Println(v59, v2)
                                }
                                var v64 int64 = 142
                                _ = v64
                            }
                            var v65 uint64 = 55
                            v2 = (v2 | 328)
                            _ = v65
                        }
                    case 9223372036854775807:
                        _ = len(v6)
                    case 0:
                        v3 -= v43
                    }
                    var v66 float32 = v5
                    fmt.Println(v43)
                    _ = v66
                } else {
                    v67 := GenS1{F0: "foo bar"}
                    for fi68 := 0; fi68 < 6; fi68++ {
                        v69 := GenS7{F0: 65535}
                        v1 += 1
                        fmt.Println(v2, v43)
                        _ = fi68
                        _ = v69
                    }
                    var v70 uint64 = 2
                    _ = v67
                    _ = v70
                }
                v54 = append(v54, fi62)
                _ = fi62
            }
        }
    }
    var v71 uint64 = 938
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v53
    _ = v54
    _ = v58
    _ = v59
    _ = v6
    _ = v71
}
