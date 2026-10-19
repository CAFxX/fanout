package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 uint8; F1 uint }
func genFn0() int16 {
    var v7 int32 = -2147483648
    if (!((false && ((166 <= 1) || (false || (-128 <= -1)))) && true)) {
        v8 := []int{}
        _ = v8
    } else {
        var v9 int16 = -356
        v10 := []int{32768, -1}
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 uint = 255
            var v13 int16 = (-(v9))
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v10
        _ = v9
    }
    var v14 int32 = (v7 | v7)
    var v15 uint16 = 0
    _ = v14
    _ = v15
    _ = v7
    return -32768
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    var v3 string = "foo bar"
    var v4 float64 = 1.5
    var v5 uint = 728
    var v6 float32 = 3.14159
    if (genFn0() == 42) {
        v4 += (v4 * 1.5)
        var v16 uint8 = 1
        v17 := []int{9223372036854775807, 32768, -2147483648, -34}
        _ = v16
        _ = v17
    } else {
        v18 := GenS1{F0: (v3 + v3)}
        _ = v18
    }
    var v19 uint = v5
    fmt.Println(v5, v19)
    var v20 bool = false
    var v21 int32 = -2147483648
    v20 = (4294967295 != 18446744073709551615)
    var v22 bool = v20
    var v23 bool = v22
    if v22 {
        if v20 {
            v23 = ((543 <= 42) || ((v5 % ((42) | 1)) > 65535))
        }
    } else {
        for fi24 := 0; fi24 < 1; fi24++ {
            if v20 {
                var v25 int16 = 1000
                _ = v25
            } else {
                var v26 uint = v19
                fmt.Println(fi24, v22)
                _ = v26
            }
            v27 := GenS2{F0: v3, F1: -9223372036854775808, F2: 830}
            _ = fi24
            _ = v27
        }
    }
    if (v22 && true) {
        switch "foo bar" {
        case "a":
            for fi28 := 0; fi28 < 7; fi28++ {
                var v29 int16 = -1
                _ = fi28
                _ = v29
            }
        case "hello":
            v21 += (-(v21))
        default:
            if v23 {
                v1[9223372036854775807] = "foo bar"
                v20 = (0 > 255)
                v23 = (false && v22)
            } else {
                v6 = float32((float64(v6) / 3.14159))
            }
        }
        v20 = v23
        var v30 string = v3
        _ = v30
    }
    v31 := GenS3{F0: 1, F1: v19}
    v4 -= 1e10
    v6 = v6
    if true {
        for fi32 := 0; fi32 < 10; fi32++ {
            if true {
                for fi33 := 0; fi33 < 19; fi33++ {
                    if (fi33 == (fi33 | (-(fi32)))) {
                        v34 := GenS2{F0: v3, F1: 501, F2: 470}
                        v35 := GenS2{F0: (v3 + v3), F1: -297, F2: 90}
                        _ = v34
                        _ = v35
                    }
                    v19 = (v5 >> (((v5 - 18446744073709551615)) & 63))
                    fmt.Println(v4, v3)
                    v36 := GenS2{F0: "\n\t", F1: 191, F2: 1}
                    var v37 int = fi33
                    _ = fi33
                    _ = v36
                    _ = v37
                }
                fmt.Println(v21, v21)
                v22 = true
            } else {
                for fi38 := 0; fi38 < 6; fi38++ {
                    v20 = true
                    _ = fi38
                }
            }
            v39 := v2[fi32]
            _ = fi32
            _ = v39
        }
        for fi40 := 0; fi40 < 1; fi40++ {
            if true {
                v19 += v19
                var v41 uint = (v19 + 696)
                _ = v41
            }
            v42 := GenS3{F0: 0, F1: v19}
            _ = fi40
            _ = v42
        }
        if (!(1e10 > (1.5 + ((v4 * (1e10 - (v4 - (v4 + v4)))) + ((v4 * 0.0) - float64(v6)))))) {
            if true {
                if (v20 || (v3 != v3)) {
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(v23, v5)
                    case "":
                        if true {
                            v43 := []int{2147483647, 1, 42, -2147483648}
                            v44 := GenS2{F0: v3, F1: 42, F2: 2147483647}
                            if v20 {
                                v43 = append(v43, 2147483647)
                                if (!(-128 < 127)) {
                                    v45 := GenS4{F0: 56, F1: 42}
                                    if v23 {
                                        v22 = (!v22)
                                        v21 = v21
                                    } else {
                                        for fi46 := 0; fi46 < 6; fi46++ {
                                            v47 := GenS5{F0: 1, F1: v19}
                                            switch v3 {
                                            case "foo bar":
                                                var v48 int16 = 42
                                                _ = v48
                                            default:
                                                switch "" {
                                                case "foo bar":
                                                    var v49 int16 = int16(v5)
                                                    _ = v49
                                                }
                                            }
                                            _ = fi46
                                            _ = v47
                                        }
                                        v22 = (v6 < (float32((float64(v6) * float64(3.14159))) - -2.25))
                                    }
                                    v50 := v2[860]
                                    _ = v45
                                    _ = v50
                                } else {
                                    var v51 uint16 = 452
                                    for fi52 := 0; fi52 < 19; fi52++ {
                                        v19 = 2
                                        var v53 int = (fi52 >> ((890) & 63))
                                        v54 := GenS2{F0: "foo bar", F1: 42, F2: 42}
                                        _ = fi52
                                        _ = v53
                                        _ = v54
                                    }
                                    for fi55 := 0; fi55 < 10; fi55++ {
                                        fmt.Println(v20, fi55, v19)
                                        _ = fi55
                                    }
                                    _ = v51
                                }
                            } else {
                                if false {
                                    fmt.Println(v19, v5, v4)
                                    var v56 float64 = 0.0
                                    _ = v56
                                }
                            }
                            _ = v43
                            _ = v44
                        }
                    }
                } else {
                    v3 = ("a" + "")
                }
                var v57 uint32 = 0
                v2[0] = -828
                _ = v57
            }
            v23 = (42 < 2)
            _ = len(v1)
        }
    }
    for fi58 := 0; fi58 < 5; fi58++ {
        v21 = (v21 ^ 0)
        var v59 int32 = v21
        v3 = v3
        _ = fi58
        _ = v59
    }
    if v20 {
        v21 = 42
    } else {
        v19 = v5
    }
    for fi60 := 0; fi60 < 2; fi60++ {
        if v23 {
            for fi61 := 0; fi61 < 5; fi61++ {
                v2[fi61] = fi60
                if (2 < 1) {
                    v19 += (v19 << ((v19) & 63))
                    fmt.Println(v20, v3, fi60)
                }
                if (2147483647 < 489) {
                    v22 = v20
                    v62 := []int{(fi61 & fi61), fi60}
                    fmt.Println(v5)
                    _ = v62
                }
                _ = fi61
            }
        }
        for fi63 := 0; fi63 < 8; fi63++ {
            v20 = v22
            if v22 {
                v20 = ((((^(v21)) <= int32(v5)) || true) || (0 <= genFn0()))
                var v64 int64 = 0
                v65 := []int{(fi63 * (fi63 + fi60)), (fi60 % ((fi60) | 1))}
                _ = v64
                _ = v65
            }
            _ = fi63
        }
        v23 = true
        _ = fi60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
