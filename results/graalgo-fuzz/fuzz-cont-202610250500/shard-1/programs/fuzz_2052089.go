package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int32; F2 int32 }
type GenS4 struct { F0 string; F1 int64; F2 uint32 }
type GenS5 struct { F0 int32; F1 uint32; F2 uint32 }

func main() {
    var v1 uint64 = 2
    var v2 bool = (42 != -1)
    var v3 uint64 = (v1 | v1)
    v4 := []int{9223372036854775807, 32768, 9223372036854775807}
    var v5 int32 = 0
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi6 := 0; fi6 < 18; fi6++ {
        v5 += v5
        v5 += 1
        for fi7 := 0; fi7 < 4; fi7++ {
            v1 += (v1 % ((42) | 1))
            _ = fi7
        }
        _ = fi6
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v5 += (v5 + 2147483647)
    if (!true) {
        fmt.Println(v5)
    }
    v5 -= (v5 - v5)
    v8 := GenS1{F0: -1, F1: 18446744073709551615}
    var v9 string = ("" + ((("foo bar" + "\n\t") + "hello") + "a"))
    v3 = (v3 ^ v1)
    for fi10 := 0; fi10 < 20; fi10++ {
        if len(v4) > 0 {
            v4[0] = (fi10 >> ((4294967295) & 63))
        }
        if true {
            if (v2 && (420 == 18446744073709551615)) {
                v11 := GenS1{F0: fi10, F1: 266}
                var v12 string = v9
                _ = v11
                _ = v12
            } else {
                if v2 {
                    if v2 {
                        var v13 uint8 = 36
                        v14 := GenS2{F0: 42}
                        _ = v13
                        _ = v14
                    } else {
                        var v15 uint = 0
                        v15 = v15
                        _ = v15
                    }
                    fmt.Println(v5, v1, v1)
                }
                var v16 bool = v2
                if v16 {
                    if false {
                        if (42 != 2) {
                            v17 := GenS3{F0: uint32(v1), F1: v5, F2: (v5 / (((v5 ^ v5)) | 1))}
                            if v2 {
                                var v18 uint16 = 1
                                fmt.Println(v18)
                                _ = v18
                            } else {
                                var v19 uint32 = 68
                                var v20 uint32 = 0
                                _ = v19
                                _ = v20
                            }
                            v21 := GenS4{F0: v9, F1: -1, F2: 690}
                            _ = v17
                            _ = v21
                        }
                        var v22 float32 = 3.14159
                        _ = v22
                    } else {
                        v2 = (0 >= 128)
                        _ = len(v4)
                    }
                    v1 += 2
                    v16 = (v16 && true)
                } else {
                    v1 = 993
                    if v16 {
                        fmt.Println(v3)
                        var v23 string = "hello"
                        _ = v23
                    } else {
                        v2 = true
                        for fi24 := 0; fi24 < 6; fi24++ {
                            v1 += v3
                            v25 := GenS1{F0: fi24, F1: 4294967295}
                            fmt.Println(v2, fi24)
                            if false {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if (30 >= -1) {
                                    v26 := GenS3{F0: 4294967295, F1: -406, F2: 2147483647}
                                    v4 = append(v4, (fi24 | (fi24 ^ fi10)))
                                    v4 = append(v4, (fi24 * (^(fi10))))
                                    v27 := GenS1{F0: -1, F1: 4294967295}
                                    _ = v26
                                    _ = v27
                                } else {
                                    fmt.Println(fi24, v9)
                                    v4 = append(v4, 9223372036854775807)
                                    fmt.Println(fi10, fi24)
                                }
                                fmt.Println(v3)
                            }
                            _ = fi24
                            _ = v25
                        }
                        if (false && (42 == 0)) {
                            var v28 uint64 = 4294967295
                            _ = v28
                        }
                    }
                    _ = len(v4)
                }
                _ = v16
            }
            for fi29 := 0; fi29 < 12; fi29++ {
                v30 := GenS4{F0: v9, F1: -580, F2: 0}
                v2 = true
                v9 = ((((("" + "foo bar") + v9) + v9) + v9) + v9)
                _ = fi29
                _ = v30
            }
        }
        var v31 uint = 967
        _ = fi10
        _ = v31
    }
    if (v2 && (false && (180 != -1))) {
        v5 = (v5 & v5)
        if v2 {
            var v32 float64 = 0.0
            v2 = v2
            switch "a" {
            case "\n\t":
                var v33 uint8 = 0
                _ = v33
            case "a":
                v4 = append(v4, 32767)
            case "foo bar":
                switch ((v9 + ("a" + "a")) + v9) {
                case "hello":
                    var v34 int32 = -2147483648
                    _ = v34
                case "foo bar":
                    var v35 uint = 4294967295
                    _ = v35
                case "a":
                    fmt.Println(v2)
                }
            default:
                v9 = v9
            }
            _ = v32
        } else {
            v4 = append(v4, 0)
        }
        _ = len(v4)
    } else {
        v2 = v2
        v36 := GenS5{F0: 42, F1: 564, F2: 0}
        _ = v36
    }
    var v37 bool = v2
    for fi38 := 0; fi38 < 8; fi38++ {
        var v39 uint16 = 65535
        v2 = true
        for fi40 := 0; fi40 < 19; fi40++ {
            switch v9 {
            case "hello":
                var v41 bool = v2
                v37 = (v39 != (v39 << ((4294967295) & 63)))
                _ = v41
            case "a":
                if len(v4) > 0 {
                    v4[0] = 32767
                }
            default:
                fmt.Println(v39, v1)
            }
            var v42 uint = 274
            switch v9 {
            case "hello":
                v37 = v37
            case "":
                v43 := GenS5{F0: v5, F1: 2147483648, F2: 890}
                _ = v43
            case "a":
                if v37 {
                    var v44 uint16 = v39
                    v42 = (v42 / ((v42) | 1))
                    _ = v44
                }
            }
            _ = fi40
            _ = v42
        }
        _ = fi38
        _ = v39
    }
    switch -113 {
    case -199:
        if v2 {
            fmt.Println(v2)
            v4 = append(v4, -510)
        }
    case 1:
        for fi45 := 0; fi45 < 6; fi45++ {
            var v46 string = "a"
            fmt.Println(v3)
            for fi47 := 0; fi47 < 15; fi47++ {
                switch fi45 {
                case -2147483648:
                    v48 := GenS3{F0: 2, F1: v5, F2: v5}
                    _ = v48
                case 128:
                    v46 = v9
                default:
                    v37 = v2
                }
                fmt.Println(v37)
                _ = len(v4)
                _ = fi47
            }
            _ = fi45
            _ = v46
        }
    default:
        if len(v4) > 0 {
            v4[0] = -9223372036854775808
        }
    }
    var v49 int64 = 9223372036854775807
    for fi50 := 0; fi50 < 8; fi50++ {
        var v51 uint32 = 0
        var v52 string = v9
        _ = fi50
        _ = v51
        _ = v52
    }
    var v53 bool = (v49 < (v49 - (v49 * (v49 * (v49 ^ (^(v49)))))))
    v5 -= 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v49
    _ = v5
    _ = v53
    _ = v8
    _ = v9
}
