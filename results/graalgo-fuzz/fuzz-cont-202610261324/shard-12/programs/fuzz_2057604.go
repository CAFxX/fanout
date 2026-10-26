package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 uint16; F1 int64; F2 string }
func genFn0(p0 bool, p1 int64) uint32 {
    p1 += p1
    fmt.Println(p0, p0)
    p0 = (42 >= -80)
    if (!((!(1 < 255)) || (false && (p0 && (p0 || (9223372036854775807 <= -9223372036854775808)))))) {
        v4 := []int{158, 42, 0}
        _ = v4
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 int32, p1 uint64, p2 int32) int64 {
    if (!false) {
        var v5 int32 = p0
        var v6 uint64 = (^(p1))
        switch 32768 {
        case 32767:
            var v7 uint = 65535
            _ = v7
        case 1:
            fmt.Println(p1, p1, v6)
        case 32768:
            if (!false) {
                var v8 int8 = -75
                _ = v8
            }
        }
        _ = v5
        _ = v6
    }
    p0 = 0
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := []int{fi9, (fi9 >> ((42) & 63)), fi9}
        var v11 int8 = -128
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 798
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -1
    v3 := make(map[string]int)
    if true {
        v2 = v2
    }
    v2 += 135
    for fi12 := 0; fi12 < 1; fi12++ {
        switch -262 {
        case 127:
            v13 := GenS1{F0: true, F1: 2}
            _ = v13
        case -2147483648:
            v14 := v3["a"]
            _ = v14
        case 2147483647:
            v15 := v1[fi12]
            _ = v15
        default:
            if (((0.1 / 3.14159) / 1.0) > (((1.5 * ((((float64(float32(0.1)) + ((3.14159 + (3.14159 * float64(1e10))) / -2.25)) - (1.5 * 1e10)) / (float64(0.1) / 1e10)) * 0.1)) - 0.1) * 1e10)) {
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 string = "\n\t"
                    _ = fi16
                    _ = v17
                }
            } else {
                v2 = (-(v2))
                var v18 bool = true
                switch fi12 {
                case -2147483648:
                    var v19 int32 = 103
                    _ = v19
                case -1:
                    var v20 float64 = ((3.14159 * 1e10) - ((0.1 - 1e10) * 1e10))
                    _ = v20
                }
                _ = v18
            }
        }
        _ = fi12
    }
    v2 += v2
    var v21 int8 = 101
    switch -1 {
    case -2147483648:
        var v22 int = -2147483648
        _ = v22
    }
    fmt.Println(v21, v2)
    v23 := []int{-2147483648}
    if true {
        switch 128 {
        case 2147483647:
            var v24 int64 = -1
            _ = v24
        case 32768:
            switch "a" {
            case "hello":
                switch (("\n\t" + ("\n\t" + ("foo bar" + "foo bar"))) + "\n\t") {
                case "hello":
                    if false {
                        v25 := GenS1{F0: (920 != 1), F1: 42}
                        _ = len(v3)
                        _ = v25
                    }
                }
            case "a":
                if false {
                    v2 = 32767
                }
            }
        }
    } else {
        for fi26 := 0; fi26 < 16; fi26++ {
            v23 = append(v23, fi26)
            var v27 bool = true
            var v28 uint16 = 0
            _ = fi26
            _ = v27
            _ = v28
        }
        v29 := GenS1{F0: true, F1: 1}
        _ = v29
    }
    var v30 uint64 = 18446744073709551615
    for fi31 := 0; fi31 < 1; fi31++ {
        var v32 bool = (0.1 == float32(1.5))
        _ = fi31
        _ = v32
    }
    if (255 >= 65535) {
        v33 := GenS2{F0: true, F1: false}
        _ = v33
    } else {
        var v34 string = "\n\t"
        v34 = (v34 + "a")
        var v35 int64 = 9223372036854775807
        _ = v34
        _ = v35
    }
    var v36 uint8 = 255
    for fi37 := 0; fi37 < 17; fi37++ {
        if true {
            switch "\n\t" {
            case "a":
                v30 = 109
            case "":
                var v38 int64 = genFn1(0, (v30 & v30), int32(v21))
                _ = v38
            default:
                var v39 uint32 = 2147483648
                _ = v39
            }
            switch "a" {
            case "foo bar":
                if (fi37 <= int(v2)) {
                    switch "" {
                    case "hello":
                        var v40 int32 = 42
                        _ = v40
                    case "\n\t":
                        switch (("hello" + "hello") + (("a" + ("foo bar" + (("foo bar" + "hello") + "\n\t"))) + "foo bar")) {
                        case "a":
                            v36 += (^(v36))
                        case "":
                            switch "a" {
                            case "a":
                                switch fi37 {
                                case 127:
                                    var v41 uint16 = 808
                                    _ = v41
                                default:
                                    fmt.Println(v2, v36)
                                }
                            }
                        }
                    default:
                        for fi42 := 0; fi42 < 16; fi42++ {
                            if true {
                                v43 := GenS2{F0: true, F1: true}
                                if false {
                                    fmt.Println(v21)
                                } else {
                                    v44 := GenS3{F0: 226, F1: genFn1(2147483647, v30, -2147483648), F2: "foo bar"}
                                    v23 = append(v23, fi37)
                                    var v45 string = ("" + ("" + "foo bar"))
                                    _ = v44
                                    _ = v45
                                }
                                _ = v43
                            } else {
                                var v46 uint8 = uint8(v30)
                                v47 := GenS1{F0: false, F1: 916}
                                var v48 uint16 = 830
                                _ = v46
                                _ = v47
                                _ = v48
                            }
                            if false {
                                var v49 string = "\n\t"
                                fmt.Println(fi42, v21)
                                _ = v49
                            } else {
                                v50 := GenS1{F0: true, F1: 65535}
                                var v51 int8 = (v21 << ((18446744073709551615) & 63))
                                _ = v50
                                _ = v51
                            }
                            _ = fi42
                        }
                    }
                    var v52 int64 = 0
                    for fi53 := 0; fi53 < 2; fi53++ {
                        switch fi37 {
                        case 315:
                            fmt.Println(v30, v36)
                        default:
                            var v54 uint64 = v30
                            _ = v54
                        }
                        _ = fi53
                    }
                    _ = v52
                } else {
                    v21 -= int8(v30)
                }
            case "\n\t":
                var v55 uint32 = 495
                _ = v55
            default:
                var v56 uint64 = (v30 - v30)
                _ = v56
            }
            if ((v21 < (v21 - v21)) || true) {
                for fi57 := 0; fi57 < 18; fi57++ {
                    var v58 float32 = 3.14159
                    _ = fi57
                    _ = v58
                }
            } else {
                for fi59 := 0; fi59 < 4; fi59++ {
                    v2 -= (v2 >> ((997) & 63))
                    _ = fi59
                }
                fmt.Println(fi37)
            }
        } else {
            v60 := v3["a"]
            v36 = 255
            _ = v60
        }
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v30
    _ = v36
}
