package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint16; F2 int64 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 bool; F1 uint64; F2 uint }
type GenS5 struct { F0 int; F1 bool; F2 uint32 }
type GenS6 struct { F0 uint; F1 uint32; F2 int16 }
type GenS7 struct { F0 int8; F1 uint32; F2 bool }

func main() {
    var v1 uint32 = 356
    v2 := []int{}
    var v3 int64 = 9223372036854775807
    var v4 int = -9223372036854775808
    var v5 int64 = (v3 / (((v3 * v3)) | 1))
    v6 := make(map[int]string)
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 string = ""
        v9 := GenS1{F0: (v1 % ((0) | 1))}
        v1 = 42
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: (v1 * (v1 ^ (v1 >> ((2) & 63))))}
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS1{F0: (v1 >> ((4294967295) & 63))}
        _ = len(v2)
        for fi13 := 0; fi13 < 7; fi13++ {
            v3 += (v3 * v3)
            switch "hello" {
            case "":
                v6[(^(fi11))] = "\n\t"
            case "foo bar":
                v14 := GenS2{F0: (false || (940 < 384)), F1: 568, F2: v3}
                _ = v14
            case "a":
                v15 := GenS2{F0: ((!(true && false)) || false), F1: 255, F2: v3}
                _ = v15
            default:
                if (false || ((true || false) && ((2 != 128) || true))) {
                    var v16 int8 = -55
                    var v17 int16 = 0
                    _ = v16
                    _ = v17
                } else {
                    v3 += (v5 % ((-1) | 1))
                    v18 := GenS1{F0: 692}
                    _ = v18
                }
            }
            if ((v1 + (^(v1))) > (v1 * (v1 + 975))) {
                v1 += v1
                fmt.Println(v1, v3)
                var v19 string = (("\n\t" + "hello") + ("foo bar" + "hello"))
                _ = v19
            }
            switch -690 {
            case -9223372036854775808:
                v5 = v5
                v2 = append(v2, 352)
            }
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    if len(v2) > 0 {
        v2[0] = v4
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 += 378
    var v20 bool = (42 <= 32767)
    v21 := GenS2{F0: (float32(float64(-2.25)) != 1e10), F1: 2, F2: (v5 ^ v5)}
    fmt.Println(v1, v5)
    var v22 int = (v4 - (v4 - v4))
    v3 += (v3 | 199)
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 int64 = (^(v5))
        for fi25 := 0; fi25 < 10; fi25++ {
            fmt.Println(v4, v4, fi23)
            var v26 uint = 322
            v24 = v3
            var v27 float32 = (1.5 + float32(1.5))
            _ = fi25
            _ = v26
            _ = v27
        }
        if v20 {
            v28 := GenS2{F0: v20, F1: 32767, F2: -1}
            _ = v28
        } else {
            for fi29 := 0; fi29 < 11; fi29++ {
                v30 := GenS3{F0: ((("" + "a") + ((("hello" + "") + (("\n\t" + "hello") + ("foo bar" + ""))) + "\n\t")) + "foo bar"), F1: -1}
                _ = fi29
                _ = v30
            }
            for fi31 := 0; fi31 < 14; fi31++ {
                v32 := GenS4{F0: ("hello" <= (("a" + "") + "foo bar")), F1: 177, F2: 1}
                var v33 uint64 = 79
                if true {
                    switch v4 {
                    case -852:
                        var v34 uint32 = v1
                        if (true || (true && (-32768 < 32767))) {
                            _ = len(v2)
                        }
                        _ = v34
                    }
                    if v20 {
                        fmt.Println(fi31, v22)
                        if (!(1 <= 127)) {
                            var v35 uint32 = (v1 | (^(v1)))
                            _ = v35
                        }
                        v36 := GenS2{F0: v20, F1: 2, F2: 42}
                        var v37 uint16 = 526
                        _ = v36
                        _ = v37
                    } else {
                        var v38 uint64 = (^(v33))
                        v6[(fi31 * v4)] = "a"
                        var v39 int32 = 0
                        _ = v38
                        _ = v39
                    }
                    v4 += 801
                }
                _ = fi31
                _ = v32
                _ = v33
            }
        }
        _ = fi23
        _ = v24
    }
    for fi40 := 0; fi40 < 7; fi40++ {
        v41 := GenS4{F0: false, F1: 0, F2: 65535}
        var v42 float64 = 3.14159
        v22 = v22
        _ = fi40
        _ = v41
        _ = v42
    }
    switch "a" {
    case "a":
        if len(v2) > 0 {
            v2[0] = v22
        }
    case "\n\t":
        v2 = append(v2, (v4 * (^(v22))))
    case "hello":
        v43 := GenS5{F0: (v4 & 127), F1: v20, F2: 4294967295}
        _ = v43
    }
    for fi44 := 0; fi44 < 3; fi44++ {
        if false {
            var v45 int64 = v3
            if ((!v20) || (((("\n\t" + "a") + ("foo bar" + "a")) + "foo bar") <= ((("foo bar" + "hello") + "") + "a"))) {
                var v46 float32 = 3.14159
                _ = v46
            } else {
                for fi47 := 0; fi47 < 13; fi47++ {
                    if (23 > 0) {
                        var v48 uint32 = 2147483648
                        var v49 uint16 = 32767
                        v50 := GenS3{F0: "", F1: -1}
                        var v51 int32 = -263
                        if (false && true) {
                            var v52 bool = v20
                            _ = v52
                        }
                        _ = v48
                        _ = v49
                        _ = v50
                        _ = v51
                    }
                    _ = fi47
                }
                v53 := GenS6{F0: 255, F1: v1, F2: -909}
                v54 := GenS7{F0: -128, F1: v1, F2: (!(v1 != v1))}
                _ = v53
                _ = v54
            }
            _ = v45
        }
        v55 := GenS3{F0: "hello", F1: 0}
        delete(v6, (v22 | 32767))
        _ = fi44
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
