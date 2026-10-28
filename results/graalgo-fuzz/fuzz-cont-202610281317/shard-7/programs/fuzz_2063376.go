package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 bool }
func genFn0(p0 uint32) uint {
    var v4 int8 = -1
    v5 := GenS1{F0: -2147483648, F1: 18446744073709551615}
    if false {
        fmt.Println(v4)
    }
    p0 = p0
    _ = p0
    _ = v4
    _ = v5
    return 1
}
func genFn1() uint {
    var v6 uint32 = 0
    v7 := GenS2{F0: 4294967295, F1: 1000, F2: 1}
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 uint64 = 715
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    var v8 string = ("hello" + "")
    if ((435 == v1) && true) {
        if (false || true) {
            v9 := GenS2{F0: v1, F1: 1, F2: 786}
            _ = v9
        } else {
            v8 = v8
            for fi10 := 0; fi10 < 14; fi10++ {
                if true {
                    v11 := GenS1{F0: int32(v1), F1: v1}
                    v12 := GenS2{F0: (^(v1)), F1: 0, F2: 448}
                    _ = v11
                    _ = v12
                }
                _ = fi10
            }
            if (809 < 42) {
                var v13 uint16 = 499
                v14 := []int{4294967295, 345}
                if len(v14) > 0 {
                    v14[0] = 42
                }
                _ = v13
                _ = v14
            } else {
                for fi15 := 0; fi15 < 12; fi15++ {
                    if true {
                        v16 := GenS3{F0: 226, F1: genFn0(65535), F2: (^(v1))}
                        v1 -= (^(v1))
                        _ = v16
                    } else {
                        v17 := GenS3{F0: 42, F1: uint(v1), F2: (v1 >> ((genFn0(0)) & 63))}
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 17; fi18++ {
                        fmt.Println(v1)
                        var v19 int8 = 127
                        _ = fi18
                        _ = v19
                    }
                    _ = fi15
                }
                if (-104 == -128) {
                    fmt.Println(v8, v8)
                    var v20 uint8 = 255
                    var v21 int16 = -770
                    _ = v20
                    _ = v21
                }
            }
        }
    } else {
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 uint32 = 1
            v24 := GenS2{F0: 0, F1: 903, F2: 2}
            v8 = v8
            _ = fi22
            _ = v23
            _ = v24
        }
        var v25 float64 = 1.5
        for fi26 := 0; fi26 < 1; fi26++ {
            for fi27 := 0; fi27 < 2; fi27++ {
                v28 := v2[v8]
                _ = fi27
                _ = v28
            }
            if (32767 != int16(fi26)) {
                var v29 bool = (0 <= 1)
                _ = v29
            } else {
                if (!(fi26 != (fi26 >> ((354) & 63)))) {
                    var v30 int64 = 1
                    switch 32767 {
                    case 0:
                        for fi31 := 0; fi31 < 2; fi31++ {
                            v32 := []int{9223372036854775807, (fi31 << ((4294967295) & 63)), int(v1)}
                            if true {
                                switch (-(fi31)) {
                                case 42:
                                    var v33 float64 = (v25 * v25)
                                    _ = v33
                                case -1:
                                    switch "a" {
                                    case "hello":
                                        v1 += (v1 * (v1 << ((255) & 63)))
                                    case "":
                                        var v34 int8 = 42
                                        _ = v34
                                    }
                                    fmt.Println(v1, v8, v25)
                                    _ = len(v32)
                                case -2147483648:
                                    if len(v32) >= 2 {
                                        v32 = v32[1:len(v32)]
                                    }
                                }
                                fmt.Println(v30)
                            }
                            switch (v8 + v8) {
                            case "a":
                                var v35 uint64 = 1
                                _ = v35
                            }
                            _ = fi31
                            _ = v32
                        }
                    case 32768:
                        var v36 uint64 = (v1 << ((18446744073709551615) & 63))
                        _ = v36
                    case -9223372036854775808:
                        v37 := []int{fi26, 4294967295, (fi26 + (fi26 / ((-1) | 1)))}
                        _ = v37
                    }
                    _ = v30
                }
                switch (("" + v8) + (v8 + v8)) {
                case "\n\t":
                    for fi38 := 0; fi38 < 8; fi38++ {
                        switch (fi38 & (-(fi38))) {
                        case 32767:
                            v1 += (^(v1))
                        case 1:
                            switch "hello" {
                            case "\n\t":
                                v39 := GenS4{F0: (^(fi26)), F1: 65535}
                                _ = v39
                            case "foo bar":
                                v40 := GenS2{F0: v1, F1: 1000, F2: 491}
                                _ = v40
                            default:
                                v41 := GenS4{F0: fi38, F1: 32767}
                                _ = v41
                            }
                        case 734:
                            switch v8 {
                            case "hello":
                                v42 := GenS4{F0: fi26, F1: 32767}
                                _ = v42
                            default:
                                fmt.Println(v8, fi38, v1)
                                var v43 float64 = (0.0 - (v25 + 0.0))
                                _ = v43
                            }
                        default:
                            switch (((v8 + v8) + (v8 + "a")) + ("" + "a")) {
                            case "foo bar":
                                v44 := v3[(^(fi38))]
                                _ = v44
                            default:
                                v45 := GenS3{F0: 237, F1: 65535, F2: v1}
                                _ = v45
                            }
                        }
                        _ = fi38
                    }
                case "":
                    var v46 int64 = 9223372036854775807
                    _ = v46
                }
            }
            if true {
                _ = len(v3)
                switch "\n\t" {
                case "":
                    var v47 int64 = 42
                    _ = v47
                }
            }
            _ = fi26
        }
        _ = v25
    }
    v48 := []int{2147483647, 0}
    v1 += v1
    v1 = v1
    switch (v8 + v8) {
    case "a":
        var v49 uint64 = v1
        _ = v49
    case "hello":
        v50 := GenS5{F0: 18, F1: -32768}
        _ = v50
    default:
        switch ("a" + ((v8 + "a") + v8)) {
        case "hello":
            v51 := GenS6{F0: v1, F1: 127}
            _ = v51
        default:
            for fi52 := 0; fi52 < 2; fi52++ {
                _ = len(v48)
                _ = fi52
            }
        }
    }
    if false {
        _ = len(v48)
        fmt.Println(v1)
        for fi53 := 0; fi53 < 4; fi53++ {
            switch ("a" + "") {
            case "a":
                var v54 int16 = -905
                _ = v54
            }
            var v55 string = v8
            if (-670 <= fi53) {
                var v56 bool = false
                for fi57 := 0; fi57 < 12; fi57++ {
                    delete(v2, v55)
                    switch fi57 {
                    case 42:
                        switch (v8 + "hello") {
                        case "a":
                            fmt.Println(v8, fi53, fi53)
                        case "":
                            v55 = ((v8 + v55) + "a")
                        }
                    default:
                        v1 += v1
                    }
                    _ = len(v2)
                    fmt.Println(v1)
                    _ = fi57
                }
                v1 += v1
                _ = v56
            } else {
                if false {
                    v55 = ("a" + (v8 + (v55 + v55)))
                }
                fmt.Println(v8, fi53, v8)
            }
            _ = fi53
            _ = v55
        }
    } else {
        if false {
            var v58 int8 = 110
            _ = v58
        } else {
            if (517 == 406) {
                v8 = (v8 + ((("a" + "hello") + v8) + v8))
            }
        }
        switch ("foo bar" + v8) {
        case "\n\t":
            v59 := GenS7{F0: false}
            _ = v59
        case "a":
            switch 9223372036854775807 {
            case 588:
                for fi60 := 0; fi60 < 10; fi60++ {
                    v8 = ("hello" + v8)
                    fmt.Println(fi60)
                    _ = fi60
                }
            case 0:
                v61 := GenS3{F0: 167, F1: 2, F2: 400}
                _ = v61
            case -2147483648:
                v62 := v3[-116]
                _ = v62
            }
        case "":
            fmt.Println(v8)
        default:
            v8 = v8
        }
    }
    _ = len(v3)
    fmt.Println(v1)
    var v63 bool = true
    v63 = (18446744073709551615 < 65535)
    switch (((v8 + v8) + ("\n\t" + v8)) + ((v8 + "foo bar") + "\n\t")) {
    case "a":
        fmt.Println(v1, v8, v1)
    case "\n\t":
        v64 := GenS1{F0: -2147483648, F1: v1}
        _ = v64
    case "foo bar":
        var v65 uint = 65535
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v63:", v63)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v63
    _ = v8
}
