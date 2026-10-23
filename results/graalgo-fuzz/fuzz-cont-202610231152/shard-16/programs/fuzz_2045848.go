package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 uint; F2 int32 }
type GenS5 struct { F0 uint64; F1 uint }
func genFn0() int8 {
    var v6 int8 = -75
    if true {
        v7 := GenS1{F0: 42, F1: 255}
        v6 = (v6 << ((255) & 63))
        _ = v7
    } else {
        v6 += 1
        var v8 float32 = 3.14159
        v6 -= v6
        _ = v8
    }
    var v9 uint8 = 1
    v10 := GenS1{F0: 0, F1: 935}
    if true {
        v6 = (v6 ^ (v6 % (((v6 ^ (v6 * (v6 << ((0) & 63))))) | 1)))
    }
    _ = v10
    _ = v6
    _ = v9
    return 100
}
func genFn1(p0 uint, p1 bool) int16 {
    if (243 < 714) {
        if (1 > 264) {
            switch ("a" + ((("" + "\n\t") + (("" + ((("\n\t" + "hello") + (("a" + "foo bar") + "\n\t")) + "hello")) + "foo bar")) + "")) {
            case "a":
                var v11 string = ("foo bar" + "a")
                _ = v11
            case "foo bar":
                fmt.Println(p0, p1)
            case "hello":
                v12 := []int{-128, 4294967295}
                _ = v12
            }
        }
    } else {
        var v13 uint = p0
        var v14 uint16 = 403
        _ = v13
        _ = v14
    }
    if (858 != 1000) {
        v15 := GenS2{F0: 65535}
        _ = v15
    } else {
        for fi16 := 0; fi16 < 5; fi16++ {
            p1 = (!((898 > 0) && (!p1)))
            var v17 int16 = 1000
            for fi18 := 0; fi18 < 3; fi18++ {
                switch 1 {
                case 1:
                    v19 := GenS3{F0: uint32(p0)}
                    _ = v19
                case 4294967295:
                    v20 := GenS2{F0: 2147483648}
                    _ = v20
                case -2147483648:
                    v21 := GenS2{F0: 277}
                    _ = v21
                default:
                    if p1 {
                        v17 -= v17
                        var v22 uint = (p0 / ((0) | 1))
                        v23 := []int{42}
                        _ = v22
                        _ = v23
                    } else {
                        v17 += v17
                        if ((0.0 / (1e10 + -2.25)) != 0.1) {
                            v24 := GenS1{F0: 0, F1: 0}
                            v25 := GenS4{F0: (^(fi18)), F1: p0, F2: 1}
                            _ = v24
                            _ = v25
                        }
                        p1 = p1
                    }
                }
                switch ("" + "foo bar") {
                case "hello":
                    switch fi16 {
                    case -991:
                        switch -96 {
                        case -694:
                            var v26 string = "\n\t"
                            _ = v26
                        }
                    case 964:
                        p1 = (((-2.25 - 1.5) / 3.14159) > 1e10)
                    }
                default:
                    if p1 {
                        switch (^(fi18)) {
                        case 0:
                            switch 4294967295 {
                            case 1:
                                v17 -= v17
                            case 696:
                                p0 -= 0
                            default:
                                p0 = p0
                            }
                        case 1:
                            var v27 uint16 = 255
                            _ = v27
                        case -997:
                            var v28 bool = (403 == 495)
                            _ = v28
                        }
                        p1 = p1
                        p0 = 2
                        v29 := []int{fi16, -2147483648, fi18}
                        _ = v29
                    } else {
                        if p1 {
                            switch (fi18 + (-(fi18))) {
                            case 32768:
                                v30 := []int{(fi18 >> (((p0 | p0)) & 63)), (-(fi16))}
                                _ = v30
                            case 128:
                                if (255 >= 960) {
                                    v31 := GenS5{F0: 42, F1: p0}
                                    _ = v31
                                } else {
                                    var v32 uint64 = 125
                                    v17 -= (v17 << ((p0) & 63))
                                    var v33 int64 = -946
                                    _ = v32
                                    _ = v33
                                }
                            case -9223372036854775808:
                                if (!(p1 || (!((p0 / (((p0 & p0)) | 1)) != 823)))) {
                                    switch fi16 {
                                    case -2147483648:
                                        if (13 == uint8(p0)) {
                                            p0 = 21
                                        }
                                    case 32767:
                                        fmt.Println(p1)
                                    case 1:
                                        p1 = ((0 > 127) && p1)
                                    }
                                    fmt.Println(fi16, v17)
                                }
                            }
                            if (p1 || (!(18446744073709551615 > 18446744073709551615))) {
                                fmt.Println(fi18)
                            }
                            var v34 int = fi16
                            _ = v34
                        }
                    }
                }
                var v35 uint32 = 708
                p0 += (p0 + p0)
                var v36 int16 = (^(v17))
                _ = fi18
                _ = v35
                _ = v36
            }
            _ = fi16
            _ = v17
        }
        p0 = (p0 & 65535)
        var v37 int64 = 455
        _ = v37
    }
    var v38 bool = true
    _ = p0
    _ = p1
    _ = v38
    return 42
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = -32768
    v3 := []string{"foo bar", ("foo bar" + "a"), "hello", "a"}
    v4 := []string{("" + "hello"), "hello"}
    var v5 float32 = 0.1
    var v39 uint64 = 0
    for fi40 := 0; fi40 < 7; fi40++ {
        fmt.Println(v5)
        for fi41 := 0; fi41 < 5; fi41++ {
            v42 := GenS5{F0: v39, F1: 4294967295}
            if false {
                switch int(v2) {
                case 0:
                    fmt.Println(v2, v5, v5)
                case -324:
                    v5 = (1.5 / 1.0)
                case -104:
                    fmt.Println(v39)
                default:
                    v39 = 17
                }
                v43 := GenS1{F0: 9223372036854775807, F1: 32767}
                _ = v43
            } else {
                v44 := GenS3{F0: 829}
                _ = v44
            }
            if true {
                var v45 uint32 = 2
                _ = v45
            }
            fmt.Println(v5, v5, v5)
            v2 = (v2 & 0)
            _ = fi41
            _ = v42
        }
        _ = fi40
    }
    v46 := GenS3{F0: 4294967295}
    v39 += (v39 >> ((979) & 63))
    fmt.Println(v2, v2)
    switch 2147483647 {
    case -2147483648:
        var v47 uint64 = (v39 << ((0) & 63))
        _ = v47
    case 2147483647:
        var v48 uint = 972
        _ = v48
    }
    fmt.Println(v39)
    var v49 int = 32768
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v50 string = "foo bar"
    v51 := GenS5{F0: (v39 / ((v39) | 1)), F1: 1}
    var v52 int32 = -2147483648
    v4 = append(v4, "foo bar")
    v53 := GenS2{F0: 154}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v46
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
}
