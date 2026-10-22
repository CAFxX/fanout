package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 uint16 }
func genFn0() bool {
    if (!false) {
        v5 := GenS1{F0: 42}
        v6 := GenS2{F0: 1000}
        if (1 > 1) {
            v7 := []int{32768, 127, -367}
            var v8 int16 = 412
            if len(v7) > 0 {
                v7[0] = 2147483647
            }
            _ = v7
            _ = v8
        }
        _ = v5
        _ = v6
    }
    var v9 float32 = float32(-2.25)
    v10 := []int{-915, 32768, 2147483647}
    var v11 uint64 = 956
    v9 -= v9
    _ = v10
    _ = v11
    _ = v9
    return true
}
func genFn1(p0 string, p1 int8) uint64 {
    if false {
        v12 := []int{32767}
        fmt.Println(p1)
        switch p0 {
        case "a":
            p1 += 42
        default:
            if (true && false) {
                if ((p1 ^ (^(p1))) >= (p1 % ((0) | 1))) {
                    for fi13 := 0; fi13 < 10; fi13++ {
                        fmt.Println(p1, p1, fi13)
                        _ = fi13
                    }
                    v12 = append(v12, 0)
                } else {
                    v14 := GenS2{F0: -127}
                    var v15 uint32 = 65535
                    _ = v14
                    _ = v15
                }
            } else {
                if (true && true) {
                    for fi16 := 0; fi16 < 7; fi16++ {
                        for fi17 := 0; fi17 < 19; fi17++ {
                            if (1e10 == ((3.14159 + ((-2.25 + 1e10) * 0.1)) + (0.0 + 1.5))) {
                                v12 = append(v12, fi17)
                                fmt.Println(p0, p1, fi17)
                            } else {
                                v12 = append(v12, 0)
                                fmt.Println(fi16, fi17)
                                switch p0 {
                                case "a":
                                    p0 = "hello"
                                case "hello":
                                    v18 := GenS1{F0: 4294967295}
                                    _ = v18
                                case "":
                                    _ = len(v12)
                                }
                            }
                            switch ("hello" + (("foo bar" + p0) + "")) {
                            case "":
                                var v19 int64 = -832
                                _ = v19
                            }
                            _ = fi17
                        }
                        for fi20 := 0; fi20 < 7; fi20++ {
                            if ((!false) || true) {
                                p0 = p0
                                if (0.1 == 0.1) {
                                    var v21 int16 = 1000
                                    _ = v21
                                } else {
                                    var v22 uint16 = 65535
                                    var v23 uint32 = 94
                                    _ = v22
                                    _ = v23
                                }
                            }
                            var v24 uint64 = 910
                            if (!false) {
                                v12 = append(v12, 127)
                                if len(v12) > 0 {
                                    v12[0] = fi20
                                }
                            } else {
                                var v25 int64 = 387
                                _ = v25
                            }
                            _ = fi20
                            _ = v24
                        }
                        _ = fi16
                    }
                    var v26 int64 = 1
                    _ = v26
                } else {
                    if false {
                        fmt.Println(p1)
                        var v27 int8 = (p1 * 90)
                        _ = v27
                    } else {
                        var v28 float64 = 3.14159
                        _ = v28
                    }
                    switch ("foo bar" + "foo bar") {
                    case "":
                        p1 -= p1
                    default:
                        switch -9223372036854775808 {
                        case 32767:
                            var v29 float64 = float64(1.5)
                            _ = v29
                        default:
                            switch 1 {
                            case -949:
                                _ = len(v12)
                            }
                        }
                    }
                    var v30 int = 0
                    _ = v30
                }
                var v31 int32 = 857
                p0 = p0
                _ = v31
            }
        }
        _ = v12
    } else {
        for fi32 := 0; fi32 < 17; fi32++ {
            v33 := []int{fi32, fi32, fi32, 128}
            fmt.Println(p1, p1)
            switch ((p0 + p0) + p0) {
            case "a":
                if false {
                    for fi34 := 0; fi34 < 8; fi34++ {
                        if (-2147483648 > 42) {
                            var v35 string = p0
                            v36 := GenS1{F0: 1}
                            p0 = "a"
                            var v37 string = p0
                            _ = v35
                            _ = v36
                            _ = v37
                        } else {
                            p0 = "a"
                        }
                        switch (^(fi34)) {
                        case 342:
                            _ = len(v33)
                        }
                        _ = fi34
                    }
                    for fi38 := 0; fi38 < 7; fi38++ {
                        _ = len(v33)
                        _ = fi38
                    }
                    fmt.Println(p0, p0, fi32)
                }
            case "hello":
                p1 -= (^(p1))
            case "":
                p1 = 127
            }
            _ = fi32
            _ = v33
        }
        p0 = p0
    }
    p0 = (((p0 + p0) + (p0 + "\n\t")) + p0)
    var v39 uint = 1
    fmt.Println(v39, v39, p1)
    switch p0 {
    case "\n\t":
        fmt.Println(v39)
    case "":
        p1 = (p1 ^ p1)
    }
    _ = p0
    _ = p1
    _ = v39
    return 243
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 421
    var v3 float32 = 1.5
    var v4 int64 = -507
    for fi40 := 0; fi40 < 3; fi40++ {
        v2 -= (v2 % (((v2 * v2)) | 1))
        _ = fi40
    }
    for fi41 := 0; fi41 < 2; fi41++ {
        v42 := GenS1{F0: genFn1((("foo bar" + "hello") + "\n\t"), 1)}
        switch 0 {
        case 0:
            if (!(v3 >= (v3 + 0.0))) {
                v2 = v2
                if (!false) {
                    v3 = v3
                }
                if true {
                    v2 = (v2 - (v2 / (((^(v2))) | 1)))
                }
            }
        case 2147483647:
            v1[fi41] = (fi41 % ((fi41) | 1))
        }
        v2 = 1
        _ = fi41
        _ = v42
    }
    var v43 float64 = 3.14159
    v44 := GenS2{F0: -558}
    if ((!(v3 > v3)) || false) {
        v45 := GenS3{F0: 67, F1: 2147483648}
        _ = v45
    }
    for fi46 := 0; fi46 < 2; fi46++ {
        v47 := []int{(fi46 | (fi46 << ((0) & 63))), 9223372036854775807, (fi46 << ((358) & 63)), 284}
        v47 = append(v47, fi46)
        _ = fi46
        _ = v47
    }
    switch ((("\n\t" + "") + "") + "\n\t") {
    case "hello":
        for fi48 := 0; fi48 < 6; fi48++ {
            delete(v1, fi48)
            switch (fi48 - 32767) {
            case 32767:
                for fi49 := 0; fi49 < 19; fi49++ {
                    v50 := v1[fi49]
                    _ = fi49
                    _ = v50
                }
            case 42:
                delete(v1, int(v4))
            }
            fmt.Println(v4, v4, v4)
            _ = fi48
        }
    case "\n\t":
        v4 = (v4 + -1)
    case "a":
        var v51 uint64 = v2
        _ = v51
    default:
        if false {
            if (!true) {
                for fi52 := 0; fi52 < 10; fi52++ {
                    for fi53 := 0; fi53 < 16; fi53++ {
                        fmt.Println(v4, v2)
                        v54 := GenS4{F0: 42}
                        _ = fi53
                        _ = v54
                    }
                    fmt.Println(v43)
                    _ = fi52
                }
            }
        } else {
            v55 := v1[0]
            if true {
                v56 := GenS5{F0: (v2 + (v2 << ((18446744073709551615) & 63))), F1: 1}
                var v57 string = "\n\t"
                _ = v56
                _ = v57
            }
            v58 := []int{4294967295, 127}
            _ = v55
            _ = v58
        }
    }
    v43 += 0.1
    for fi59 := 0; fi59 < 2; fi59++ {
        var v60 int64 = -1
        fmt.Println(v2)
        _ = fi59
        _ = v60
    }
    var v61 int64 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v61
}
