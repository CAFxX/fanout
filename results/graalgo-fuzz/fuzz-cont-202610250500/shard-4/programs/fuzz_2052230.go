package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 uint8; F1 uint16 }
func genFn0(p0 int64, p1 uint16, p2 int64) int64 {
    p2 = 42
    var v7 bool = true
    var v8 int8 = 1
    switch "hello" {
    case "foo bar":
        v9 := []int{-608}
        _ = v9
    }
    fmt.Println(p2, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 int16, p1 uint16, p2 int64) uint64 {
    fmt.Println(p1, p0, p2)
    var v10 uint = 203
    v11 := []int{42, 416, 127, 0}
    var v12 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{"hello", (("a" + "foo bar") + ("" + "foo bar")), "\n\t", "a"}
    var v3 int32 = -1
    v4 := make(map[string]string)
    v5 := []string{}
    var v6 bool = (false && (1 == 307))
    for fi13 := 0; fi13 < 20; fi13++ {
        if false {
            var v14 uint = 4294967295
            for fi15 := 0; fi15 < 18; fi15++ {
                if true {
                    var v16 int16 = -32768
                    _ = v16
                } else {
                    var v17 int16 = 1000
                    _ = v17
                }
                var v18 float32 = (-2.25 + float32(float64(0.1)))
                v19 := GenS1{F0: 628, F1: fi15}
                _ = fi15
                _ = v18
                _ = v19
            }
            _ = v14
        }
        v6 = (-742 > 153)
        _ = fi13
    }
    var v20 int = 2147483647
    for fi21 := 0; fi21 < 9; fi21++ {
        for fi22 := 0; fi22 < 5; fi22++ {
            if v6 {
                v23 := GenS2{F0: (("foo bar" + "a") + ("" + "hello"))}
                _ = v23
            }
            _ = fi22
        }
        if (924 >= 622) {
            fmt.Println(fi21, v20, v6)
            v3 += (v3 | (v3 & (v3 ^ 1)))
        }
        var v24 uint = 653
        _ = fi21
        _ = v24
    }
    switch 42 {
    case 42:
        v20 = v20
    }
    fmt.Println(v3, v3, v20)
    v6 = false
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 uint16 = 1
        var v27 bool = false
        var v28 uint16 = v26
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    delete(v4, ("foo bar" + ""))
    for fi29 := 0; fi29 < 10; fi29++ {
        var v30 int32 = v3
        switch fi29 {
        case 1:
            v31 := GenS1{F0: 65535, F1: (v20 / ((32767) | 1))}
            _ = v31
        default:
            switch ("hello" + (("\n\t" + "\n\t") + "")) {
            case "\n\t":
                v32 := GenS2{F0: "foo bar"}
                _ = v32
            default:
                _ = len(v5)
            }
        }
        if false {
            _ = len(v2)
        } else {
            if (false && (31 <= 1)) {
                v30 = v3
            } else {
                v3 += (v3 << ((1) & 63))
                _ = len(v5)
            }
            switch "a" {
            case "":
                var v33 uint8 = 2
                _ = v33
            default:
                var v34 bool = v6
                _ = v34
            }
            if (!(((420 != 1) || (-32768 <= 797)) || (v20 != fi29))) {
                fmt.Println(v20)
                switch ("" + ("a" + "foo bar")) {
                case "a":
                    v35 := GenS3{F0: 0, F1: genFn0(0, 42, 1)}
                    _ = v35
                case "hello":
                    switch "\n\t" {
                    case "a":
                        _ = len(v4)
                    default:
                        v3 += v3
                    }
                }
                if (!v6) {
                    v20 = (-(v20))
                } else {
                    fmt.Println(v20, fi29)
                    for fi36 := 0; fi36 < 15; fi36++ {
                        if (!(!(51 != 255))) {
                            var v37 uint = 18446744073709551615
                            var v38 int16 = 42
                            _ = v37
                            _ = v38
                        } else {
                            v39 := GenS4{F0: 94}
                            v40 := GenS5{F0: 0, F1: 1}
                            v20 += v20
                            var v41 uint32 = 1
                            v3 += (v3 + -2147483648)
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                        v42 := GenS4{F0: 70}
                        var v43 uint32 = 65535
                        v44 := GenS6{F0: 0, F1: 32767}
                        fmt.Println(fi29)
                        v45 := GenS4{F0: 2}
                        fmt.Println(v20)
                        _ = fi36
                        _ = v42
                        _ = v43
                        _ = v44
                        _ = v45
                    }
                    v46 := v4["hello"]
                    _ = v46
                }
            } else {
                v4[""] = ""
            }
        }
        _ = fi29
        _ = v30
    }
    switch "foo bar" {
    case "hello":
        if v6 {
            for fi47 := 0; fi47 < 19; fi47++ {
                switch "" {
                case "\n\t":
                    v6 = true
                case "":
                    for fi48 := 0; fi48 < 17; fi48++ {
                        v20 = (v20 << ((0) & 63))
                        _ = fi48
                    }
                case "foo bar":
                    for fi49 := 0; fi49 < 3; fi49++ {
                        v50 := GenS2{F0: "foo bar"}
                        _ = fi49
                        _ = v50
                    }
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if v6 {
                    for fi51 := 0; fi51 < 3; fi51++ {
                        fmt.Println(fi47, fi47, v3)
                        if (456 <= 2) {
                            delete(v4, (((("a" + "a") + "hello") + "foo bar") + "hello"))
                            switch ("a" + ("a" + "\n\t")) {
                            case "foo bar":
                                var v52 int64 = -1
                                _ = v52
                            default:
                                var v53 float64 = 3.14159
                                _ = v53
                            }
                            v6 = ((v3 % (((^(v3))) | 1)) == -78)
                            var v54 string = "a"
                            _ = v54
                        } else {
                            v55 := GenS5{F0: 0, F1: 42}
                            if (2 == 143) {
                                var v56 uint8 = 130
                                v20 -= fi51
                                v57 := GenS5{F0: 4294967295, F1: 0}
                                v6 = (v56 > (v56 | (v56 / ((1) | 1))))
                                v6 = v6
                                v2 = append(v2, "")
                                _ = v56
                                _ = v57
                            }
                            v4[("a" + "hello")] = "hello"
                            _ = v55
                        }
                        v3 = 0
                        _ = fi51
                    }
                    fmt.Println(v6, v3, fi47)
                }
                _ = fi47
            }
            v6 = v6
            var v58 int32 = (v3 / (((v3 | int32(v20))) | 1))
            _ = v58
        } else {
            var v59 int16 = 32767
            v20 = 1
            if ((-(v3)) == 42) {
                _ = len(v1)
                v2 = append(v2, ("hello" + (("a" + "") + (("\n\t" + "") + ("a" + "")))))
            } else {
                switch (v20 << ((4294967295) & 63)) {
                case -9223372036854775808:
                    var v60 uint32 = 2147483648
                    _ = v60
                case 32768:
                    v61 := GenS4{F0: 0}
                    _ = v61
                }
                var v62 float64 = ((1.5 / float64(1e10)) / float64(3.14159))
                v63 := GenS3{F0: 9223372036854775807, F1: genFn0(-320, 146, -1)}
                _ = v62
                _ = v63
            }
            _ = v59
        }
    default:
        v64 := GenS1{F0: 88, F1: 42}
        _ = v64
    }
    if (v20 != v20) {
        if v6 {
            v3 += v3
            fmt.Println(v3, v20, v3)
        } else {
            v65 := GenS1{F0: 65535, F1: (v20 / ((v20) | 1))}
            var v66 uint32 = 1
            _ = v65
            _ = v66
        }
        for fi67 := 0; fi67 < 17; fi67++ {
            for fi68 := 0; fi68 < 10; fi68++ {
                var v69 string = ""
                fmt.Println(fi67, fi68, v69)
                switch (^(fi67)) {
                case 777:
                    v3 -= v3
                default:
                    var v70 bool = (0 < 711)
                    _ = v70
                }
                _ = fi68
                _ = v69
            }
            _ = fi67
        }
    }
    switch "\n\t" {
    case "":
        for fi71 := 0; fi71 < 18; fi71++ {
            fmt.Println(v3, v20, v20)
            delete(v4, "a")
            _ = fi71
        }
    default:
        v72 := v1[(v20 >> ((42) & 63))]
        _ = v72
    }
    var v73 int16 = -32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
}
