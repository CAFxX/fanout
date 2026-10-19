package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 string; F2 int16 }
type GenS3 struct { F0 int; F1 int64; F2 int8 }
type GenS4 struct { F0 uint64; F1 uint8; F2 string }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 int32; F1 int32; F2 int8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8; F1 uint8; F2 int64 }
func genFn0(p0 uint64, p1 int) int32 {
    fmt.Println(p0, p1)
    var v6 int64 = 42
    var v7 uint16 = 782
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -2147483648
}
func genFn1(p0 uint16, p1 uint, p2 uint) uint32 {
    p2 += (p2 << ((p2) & 63))
    if false {
        if (!(("foo bar" == ("foo bar" + "foo bar")) || ((genFn0(1, 555) != -425) && false))) {
            switch "hello" {
            case "foo bar":
                v8 := GenS1{F0: "hello"}
                _ = v8
            case "a":
                fmt.Println(p0, p1)
            }
            for fi9 := 0; fi9 < 12; fi9++ {
                fmt.Println(fi9, p2)
                _ = fi9
            }
            if (!(!true)) {
                var v10 int8 = 0
                var v11 float32 = 0.0
                if (false || (p2 >= p1)) {
                    fmt.Println(p1)
                } else {
                    v10 += (v10 / ((-128) | 1))
                }
                _ = v10
                _ = v11
            } else {
                v12 := GenS2{F0: 719, F1: "", F2: -32768}
                var v13 int = 2147483647
                _ = v12
                _ = v13
            }
        } else {
            switch int(p0) {
            case 845:
                var v14 uint32 = 65535
                _ = v14
            case 500:
                v15 := []int{1, -888, -939}
                _ = v15
            case -318:
                v16 := GenS1{F0: "hello"}
                _ = v16
            }
            v17 := []int{}
            _ = v17
        }
    } else {
        p0 -= p0
        if true {
            var v18 int8 = -81
            var v19 int8 = v18
            var v20 int64 = -79
            _ = v18
            _ = v19
            _ = v20
        } else {
            fmt.Println(p2)
        }
    }
    p0 -= 65535
    v21 := []int{-1, 98, 32767, 556}
    var v22 string = ("\n\t" + "foo bar")
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    return 2147483648
}

func main() {
    var v1 int16 = -1
    v2 := []string{("" + ((("hello" + ("hello" + "hello")) + ("hello" + "a")) + ("" + ""))), ("a" + ("a" + ("\n\t" + "\n\t"))), "hello"}
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]int)
    v5 := make(map[int]uint32)
    for fi23 := 0; fi23 < 12; fi23++ {
        v1 -= (^(v1))
        _ = fi23
    }
    fmt.Println(v1)
    v24 := GenS3{F0: 32767, F1: 0, F2: -38}
    v25 := GenS4{F0: 476, F1: 127, F2: ((("\n\t" + "\n\t") + "hello") + "hello")}
    if false {
        fmt.Println(v1)
    } else {
        v26 := GenS1{F0: "a"}
        if (2147483647 < 2147483648) {
            switch "" {
            case "hello":
                var v27 int16 = (v1 - (v1 + (v1 * (v1 - v1))))
                _ = v27
            case "":
                if false {
                    var v28 int32 = -2147483648
                    v29 := GenS5{F0: 65535, F1: 42}
                    fmt.Println(v1, v28)
                    _ = v28
                    _ = v29
                } else {
                    v30 := GenS4{F0: 41, F1: 32, F2: ("" + "a")}
                    var v31 uint16 = 507
                    _ = v30
                    _ = v31
                }
            case "\n\t":
                var v32 bool = false
                _ = v32
            default:
                v1 = (v1 * 1)
            }
        } else {
            for fi33 := 0; fi33 < 9; fi33++ {
                switch "hello" {
                case "a":
                    v34 := GenS2{F0: genFn1(0, 835, 2), F1: "foo bar", F2: 32767}
                    _ = v34
                }
                v2 = append(v2, ("" + ""))
                _ = fi33
            }
            var v35 uint32 = 2147483647
            fmt.Println(v35, v1)
            _ = v35
        }
        _ = v26
    }
    v1 = (-(v1))
    v36 := GenS5{F0: 65535, F1: -36}
    for fi37 := 0; fi37 < 8; fi37++ {
        var v38 int = (fi37 >> ((42) & 63))
        _ = fi37
        _ = v38
    }
    var v39 uint32 = 1
    if (!(!(3.14159 != ((-2.25 * 1.5) * 0.1)))) {
        var v40 int8 = 42
        if false {
            if false {
                if false {
                    v41 := GenS6{F0: 1, F1: 1, F2: (v40 * v40)}
                    var v42 uint = 42
                    for fi43 := 0; fi43 < 15; fi43++ {
                        var v44 int = int(v39)
                        switch ("a" + (("foo bar" + "") + "hello")) {
                        case "foo bar":
                            v39 = 594
                        case "":
                            if (!(0.0 > 1.5)) {
                                fmt.Println(v40, v44, v42)
                            } else {
                                v44 = fi43
                            }
                        case "a":
                            _ = len(v2)
                        }
                        v44 -= 51
                        _ = fi43
                        _ = v44
                    }
                    _ = v41
                    _ = v42
                }
                if true {
                    for fi45 := 0; fi45 < 7; fi45++ {
                        var v46 float32 = (1e10 - (1.5 * ((3.14159 / 1.0) - 0.0)))
                        v47 := GenS1{F0: "hello"}
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                    for fi48 := 0; fi48 < 6; fi48++ {
                        switch (("" + "foo bar") + "hello") {
                        case "":
                            for fi49 := 0; fi49 < 13; fi49++ {
                                var v50 int = fi48
                                v2 = append(v2, "foo bar")
                                _ = fi49
                                _ = v50
                            }
                        case "\n\t":
                            for fi51 := 0; fi51 < 20; fi51++ {
                                if (1.5 < (1e10 / 1.5)) {
                                    v52 := v3[(v39 ^ v39)]
                                    var v53 uint = 435
                                    _ = v52
                                    _ = v53
                                }
                                v1 = (v1 + -1)
                                fmt.Println(fi48, v1)
                                _ = fi51
                            }
                        case "hello":
                            if (float64(-2.25) != 1e10) {
                                if false {
                                    var v54 string = ""
                                    _ = v54
                                }
                                v55 := GenS7{F0: 1}
                                _ = v55
                            }
                        }
                        v2 = append(v2, ((("" + ("foo bar" + "")) + "a") + ""))
                        _ = fi48
                    }
                    var v56 float64 = float64((((1.5 + 1e10) * float32(((0.1 * 3.14159) * 0.0))) + ((0.0 + ((1.5 + 1e10) * ((0.1 / (1e10 / float32((1.5 / 1e10)))) / float32(1e10)))) + (-2.25 * -2.25))))
                    _ = v56
                }
            }
            fmt.Println(v39, v1)
        } else {
            v4[(v39 ^ 42)] = -9223372036854775808
        }
        switch "hello" {
        case "\n\t":
            v57 := GenS4{F0: 4294967295, F1: 127, F2: ("a" + "")}
            _ = v57
        case "foo bar":
            switch -297 {
            case 127:
                var v58 int = 4294967295
                _ = v58
            case -2147483648:
                v39 += 4294967295
            }
        }
        _ = v40
    }
    for fi59 := 0; fi59 < 17; fi59++ {
        for fi60 := 0; fi60 < 13; fi60++ {
            v1 = v1
            if true {
                v61 := GenS8{F0: -64, F1: 127, F2: -120}
                if (!false) {
                    v62 := GenS5{F0: 544, F1: 48}
                    v63 := GenS1{F0: ("\n\t" + ("\n\t" + ""))}
                    _ = v62
                    _ = v63
                }
                var v64 int = (^(fi60))
                v65 := GenS8{F0: int8(fi60), F1: 0, F2: -9223372036854775808}
                fmt.Println(v1, v39, v64)
                v1 -= (-(v1))
                _ = v61
                _ = v64
                _ = v65
            }
            _ = fi60
        }
        _ = fi59
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v5
}
