package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint32; F1 int32; F2 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 int64; F1 uint32; F2 int8 }
func genFn0(p0 bool, p1 int16) int64 {
    v6 := []int{-71, 161, 128, -2147483648}
    var v7 uint64 = 42
    p1 -= 1000
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS1{F0: v7, F1: ("a" + "\n\t")}
        v6 = append(v6, (fi8 - 128))
        p0 = (p0 || (0 < 0))
        _ = fi8
        _ = v9
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = false
    v3 := make(map[int]int)
    v4 := []uint32{141, 520, 0}
    v5 := make(map[uint32]uint32)
    v2 = (v2 && true)
    v10 := GenS2{F0: ("a" + "a"), F1: 81}
    fmt.Println(v2)
    var v11 uint16 = 65535
    v2 = v2
    v11 -= (^(v11))
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 uint16 = (v11 & v11)
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v11, v11)
    if (false && v2) {
        var v15 int16 = 474
        _ = v15
    }
    switch (("" + "a") + "foo bar") {
    case "hello":
        fmt.Println(v2, v11)
    case "\n\t":
        for fi16 := 0; fi16 < 17; fi16++ {
            v17 := GenS3{F0: 2147483648, F1: 42, F2: 274}
            v18 := GenS3{F0: 42, F1: 1, F2: 1}
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    var v19 bool = v2
    switch 32768 {
    case -95:
        if v19 {
            v11 = v11
            var v20 bool = v19
            _ = v20
        } else {
            var v21 int8 = 42
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 uint = 4294967295
                switch -833 {
                case 32768:
                    v2 = v19
                case 283:
                    _ = len(v4)
                }
                _ = fi22
                _ = v23
            }
            _ = v21
        }
    case 32767:
        switch 42 {
        case 2147483647:
            for fi24 := 0; fi24 < 14; fi24++ {
                if v2 {
                    if v2 {
                        for fi25 := 0; fi25 < 17; fi25++ {
                            v4 = append(v4, 4294967295)
                            _ = fi25
                        }
                        if ((true && false) || ((0 <= -532) && (186 == 2))) {
                            for fi26 := 0; fi26 < 17; fi26++ {
                                v27 := GenS2{F0: ("foo bar" + "\n\t"), F1: 619}
                                v19 = (2147483647 <= 0)
                                _ = fi26
                                _ = v27
                            }
                            v1[(("hello" + "") + "")] = 2147483648
                            switch fi24 {
                            case 127:
                                var v28 float32 = (1.5 * 3.14159)
                                _ = v28
                            case 0:
                                if (0 == 4294967295) {
                                    for fi29 := 0; fi29 < 15; fi29++ {
                                        fmt.Println(fi29)
                                        _ = fi29
                                    }
                                    v2 = v2
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = 65535
                                    }
                                }
                            case 32767:
                                v19 = ("hello" < "\n\t")
                            default:
                                var v30 string = ((("" + ("" + "a")) + "\n\t") + ("hello" + ((("foo bar" + "a") + "hello") + (("" + "\n\t") + ((("" + ("\n\t" + "hello")) + "a") + "")))))
                                _ = v30
                            }
                        } else {
                            var v31 int16 = 42
                            for fi32 := 0; fi32 < 1; fi32++ {
                                var v33 uint = 65535
                                var v34 float64 = (0.1 * 0.1)
                                if v2 {
                                    fmt.Println(v2)
                                    var v35 uint8 = 154
                                    fmt.Println(v33)
                                    v4 = append(v4, 42)
                                    var v36 uint32 = 2
                                    _ = v35
                                    _ = v36
                                } else {
                                    v4 = append(v4, 65535)
                                    var v37 int8 = -1
                                    _ = v37
                                }
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                            _ = v31
                        }
                    }
                }
                switch ("" + "foo bar") {
                case "\n\t":
                    fmt.Println(fi24, v19, v19)
                case "foo bar":
                    var v38 uint32 = 65535
                    _ = v38
                case "a":
                    switch (fi24 - fi24) {
                    case 127:
                        for fi39 := 0; fi39 < 3; fi39++ {
                            switch ("foo bar" + ("" + "hello")) {
                            case "a":
                                var v40 string = ("\n\t" + "\n\t")
                                _ = v40
                            }
                            if (-2147483648 <= 2147483647) {
                                switch ("foo bar" + "\n\t") {
                                case "\n\t":
                                    var v41 int = (fi24 / (((fi39 * (fi24 << ((2) & 63)))) | 1))
                                    _ = v41
                                case "foo bar":
                                    v19 = (-1 == -2147483648)
                                case "a":
                                    if (true || (!v19)) {
                                        if v19 {
                                            var v42 uint64 = 908
                                            _ = v42
                                        } else {
                                            if (33 == 255) {
                                                var v43 int64 = 0
                                                var v44 uint64 = 18446744073709551615
                                                fmt.Println(v11)
                                                _ = v43
                                                _ = v44
                                            }
                                        }
                                        v45 := GenS1{F0: uint64(v11), F1: "\n\t"}
                                        v2 = (!(0 < -32768))
                                        fmt.Println(v2, v19, v2)
                                        v46 := GenS4{F0: (v11 >> ((18446744073709551615) & 63))}
                                        _ = v45
                                        _ = v46
                                    } else {
                                        v47 := GenS5{F0: 42, F1: 0}
                                        switch fi24 {
                                        case 0:
                                            var v48 int = 127
                                            _ = v48
                                        case -987:
                                            v49 := GenS6{F0: -1, F1: 584, F2: 100}
                                            _ = v49
                                        case 127:
                                            fmt.Println(v2)
                                        default:
                                            var v50 uint64 = 42
                                            _ = v50
                                        }
                                        _ = len(v4)
                                        _ = v47
                                    }
                                }
                                fmt.Println(fi39, v11, fi39)
                                if false {
                                    switch fi39 {
                                    case 32768:
                                        delete(v3, -9223372036854775808)
                                    default:
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v51 int16 = 621
                                            _ = v51
                                        }
                                    }
                                    var v52 int32 = -188
                                    _ = v52
                                } else {
                                    var v53 string = ""
                                    _ = v53
                                }
                            } else {
                                v11 += (v11 * 9)
                                var v54 uint64 = 194
                                _ = v54
                            }
                            _ = fi39
                        }
                    case 0:
                        fmt.Println(v11, v19, v11)
                    }
                }
                v19 = (!false)
                _ = fi24
            }
        }
    default:
        if v2 {
            var v55 int32 = 2147483647
            _ = v55
        } else {
            for fi56 := 0; fi56 < 7; fi56++ {
                switch fi56 {
                case -9223372036854775808:
                    v11 = 121
                case 128:
                    var v57 uint32 = 550
                    _ = v57
                }
                fmt.Println(v11, fi56)
                var v58 string = "\n\t"
                _ = fi56
                _ = v58
            }
            v2 = ((v11 * (v11 - 0)) >= 65535)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    v19 = (!((42 >= 975) || (1 > 0)))
    v11 = 796
    v2 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
