package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string }
func genFn0() bool {
    var v4 int64 = 0
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: 0, F1: 42, F2: (v4 + 9223372036854775807)}
        _ = fi5
        _ = v6
    }
    var v7 float32 = (0.1 - ((0.1 / 1.5) - (1.5 / 0.1)))
    _ = v4
    _ = v7
    return ((false || true) && ((v4 + 0) != (-(v4))))
}
func genFn1(p0 int8) uint8 {
    v8 := []int{4294967295, 2147483647, 4294967295, -482}
    var v9 uint16 = 255
    var v10 uint8 = 0
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return (^(v10))
}

func main() {
    var v1 uint64 = 0
    var v2 uint32 = 966
    var v3 float64 = (1.5 - float64(-2.25))
    switch 128 {
    case -682:
        v11 := GenS1{F0: (v1 / ((144) | 1)), F1: 65535, F2: 9223372036854775807}
        _ = v11
    case 757:
        v2 = (v2 % ((995) | 1))
    }
    switch "\n\t" {
    case "":
        v12 := []int{-156, 0, 4294967295, -752}
        _ = v12
    default:
        var v13 uint16 = 533
        _ = v13
    }
    var v14 bool = (-274 == 0)
    for fi15 := 0; fi15 < 12; fi15++ {
        fmt.Println(fi15)
        var v16 uint64 = v1
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: (v1 | v1), F1: 1, F2: -9223372036854775808}
    if (!(0 < -1)) {
        if (42 != 103) {
            v18 := []int{}
            _ = v18
        }
        v19 := GenS1{F0: 781, F1: 65535, F2: -678}
        _ = v19
    } else {
        v3 = float64((3.14159 / 1.0))
        v20 := GenS1{F0: (v1 << ((18446744073709551615) & 63)), F1: 255, F2: 9223372036854775807}
        _ = v20
    }
    if ((!v14) || (-1 > 1)) {
        if true {
            v21 := []int{-2147483648, int(v2), 128}
            _ = v21
        }
        if (true || (!(0.1 <= 0.1))) {
            v14 = v14
        }
        if (v14 && false) {
            var v22 int16 = -1
            v23 := GenS2{F0: 42}
            _ = v22
            _ = v23
        } else {
            var v24 string = ("" + (("hello" + "") + "foo bar"))
            _ = v24
        }
    } else {
        if (!(42 > -1)) {
            for fi25 := 0; fi25 < 4; fi25++ {
                v26 := []int{}
                v3 = v3
                v27 := GenS2{F0: 42}
                _ = fi25
                _ = v26
                _ = v27
            }
            if (-1 >= 1) {
                var v28 int16 = 0
                v2 -= (v2 / ((35) | 1))
                _ = v28
            }
            v3 += v3
        } else {
            if (!v14) {
                for fi29 := 0; fi29 < 20; fi29++ {
                    fmt.Println(v1, v2)
                    var v30 int16 = 0
                    for fi31 := 0; fi31 < 4; fi31++ {
                        var v32 int = (fi31 >> ((652) & 63))
                        var v33 bool = ((32767 == uint16(v1)) || (!true))
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    _ = fi29
                    _ = v30
                }
            }
            v34 := GenS2{F0: 1}
            switch "hello" {
            case "":
                v3 -= 1e10
            case "hello":
                for fi35 := 0; fi35 < 2; fi35++ {
                    v2 = v2
                    switch "a" {
                    case "\n\t":
                        v36 := []int{(fi35 * -181)}
                        _ = v36
                    case "":
                        v1 = 333
                    case "a":
                        var v37 uint8 = 42
                        _ = v37
                    default:
                        for fi38 := 0; fi38 < 19; fi38++ {
                            switch ("a" + (("hello" + "hello") + ("a" + "a"))) {
                            case "hello":
                                var v39 uint8 = 5
                                v40 := []int{-749, fi35, (fi35 >> ((2) & 63)), fi38}
                                _ = v39
                                _ = v40
                            case "foo bar":
                                if (-494 > -9223372036854775808) {
                                    fmt.Println(v3)
                                    v41 := GenS3{F0: 0}
                                    _ = v41
                                } else {
                                    var v42 uint = 601
                                    v43 := GenS1{F0: v1, F1: 255, F2: -224}
                                    _ = v42
                                    _ = v43
                                }
                            default:
                                v1 -= 42
                            }
                            v44 := []int{134, 1}
                            var v45 float32 = 1e10
                            _ = len(v44)
                            v45 -= (v45 + ((v45 + 1.5) + v45))
                            if (v14 || v14) {
                                v46 := GenS4{F0: 42}
                                if len(v44) > 0 {
                                    v44[0] = (fi38 >> ((211) & 63))
                                }
                                _ = v46
                            } else {
                                fmt.Println(v2, v1)
                            }
                            v47 := GenS3{F0: 149}
                            var v48 int64 = 780
                            var v49 uint16 = 0
                            _ = fi38
                            _ = v44
                            _ = v45
                            _ = v47
                            _ = v48
                            _ = v49
                        }
                    }
                    var v50 float64 = (v3 * v3)
                    _ = fi35
                    _ = v50
                }
            }
            _ = v34
        }
        var v51 uint32 = (v2 & (v2 + v2))
        _ = v51
    }
    v1 = 0
    switch 128 {
    case 530:
        v3 = 1.5
    }
    v52 := []int{-1, 32768}
    for fi53 := 0; fi53 < 5; fi53++ {
        var v54 int64 = 9223372036854775807
        for fi55 := 0; fi55 < 12; fi55++ {
            var v56 string = ""
            v54 -= (^(v54))
            _ = fi55
            _ = v56
        }
        for fi57 := 0; fi57 < 20; fi57++ {
            v52 = append(v52, 596)
            if v14 {
                v58 := GenS5{F0: 0, F1: 583}
                var v59 int8 = 127
                var v60 uint = 352
                _ = v58
                _ = v59
                _ = v60
            }
            if v14 {
                var v61 int16 = 1000
                _ = v61
            } else {
                var v62 string = (("hello" + "\n\t") + ("hello" + "hello"))
                v63 := GenS6{F0: genFn1(127)}
                var v64 bool = (2 <= 42)
                _ = v62
                _ = v63
                _ = v64
            }
            switch ("" + "") {
            case "":
                if (!((v14 || (463 >= 0)) && v14)) {
                    v65 := GenS4{F0: v54}
                    v52 = append(v52, 2147483647)
                    v66 := GenS7{F0: ("foo bar" + ((("" + ("a" + "\n\t")) + "\n\t") + ("a" + "a")))}
                    _ = v65
                    _ = v66
                } else {
                    var v67 float32 = float32(v3)
                    _ = v67
                }
            default:
                var v68 float32 = (0.0 + 3.14159)
                _ = v68
            }
            _ = fi57
        }
        _ = fi53
        _ = v54
    }
    v2 -= (v2 * 589)
    v1 -= 4294967295
    v69 := GenS4{F0: 781}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v52:", len(v52), v52)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v52
    _ = v69
}
