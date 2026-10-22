package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int64; F1 int16; F2 int16 }
type GenS4 struct { F0 bool; F1 uint16; F2 uint16 }
func genFn0(p0 int, p1 uint16) int8 {
    var v7 uint32 = 0
    for fi8 := 0; fi8 < 13; fi8++ {
        if (683 < 4294967295) {
            v9 := GenS1{F0: "foo bar", F1: "foo bar"}
            if ((fi8 << ((22) & 63)) < p0) {
                fmt.Println(v7, p0)
                fmt.Println(p0)
            }
            _ = v9
        } else {
            for fi10 := 0; fi10 < 19; fi10++ {
                var v11 int64 = 9223372036854775807
                _ = fi10
                _ = v11
            }
        }
        v12 := []int{42, p0, (p0 ^ (p0 * 759))}
        _ = fi8
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    return 100
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 42
    var v3 bool = false
    var v4 uint16 = 849
    var v5 uint8 = 87
    v6 := make(map[string]bool)
    v5 += (^(v5))
    fmt.Println(v2, v4)
    v13 := GenS1{F0: "", F1: "hello"}
    v14 := GenS1{F0: "", F1: ("\n\t" + (("a" + ("hello" + ("\n\t" + "hello"))) + "foo bar"))}
    v3 = (!v3)
    v2 += 32767
    v4 -= 935
    if (907 == v4) {
        v4 += v4
        if (v5 > 1) {
            for fi15 := 0; fi15 < 19; fi15++ {
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 float32 = 1e10
                    switch (v2 / ((int(v4)) | 1)) {
                    case -718:
                        v2 -= v2
                    case -218:
                        fmt.Println(v5, v5)
                    case -1:
                        if (v3 || (376 >= 51)) {
                            v18 := []int{(fi16 | v2)}
                            _ = v18
                        } else {
                            v17 += float32((1e10 / ((float64((1.5 - 0.0)) / -2.25) / (float64((-2.25 + v17)) - -2.25))))
                        }
                    default:
                        v19 := GenS1{F0: "hello", F1: ("foo bar" + "")}
                        _ = v19
                    }
                    _ = fi16
                    _ = v17
                }
                if (v3 && false) {
                    v4 = 32767
                }
                _ = fi15
            }
            v20 := GenS1{F0: "\n\t", F1: "\n\t"}
            fmt.Println(v3, v4, v3)
            _ = v20
        } else {
            if v3 {
                v3 = (!v3)
                fmt.Println(v4, v2, v5)
                v2 += v2
            } else {
                for fi21 := 0; fi21 < 3; fi21++ {
                    v22 := GenS1{F0: "a", F1: (("hello" + (("\n\t" + "") + "foo bar")) + "hello")}
                    _ = fi21
                    _ = v22
                }
                for fi23 := 0; fi23 < 20; fi23++ {
                    switch "hello" {
                    case "":
                        for fi24 := 0; fi24 < 10; fi24++ {
                            v25 := []int{128, (fi24 << ((65535) & 63))}
                            if v3 {
                                if v3 {
                                    if v3 {
                                        if (!(v3 || ((3.14159 / (1.5 - (float32((0.0 + float64((1.5 / 3.14159)))) / 1e10))) < float32((0.1 + ((1e10 / 1e10) * 1.5)))))) {
                                            fmt.Println(v2, v4)
                                            var v26 bool = false
                                            v2 = (fi24 + 127)
                                            if len(v25) >= 2 {
                                                v25 = v25[1:len(v25)]
                                            }
                                            v27 := GenS1{F0: "foo bar", F1: "hello"}
                                            _ = v26
                                            _ = v27
                                        }
                                        v2 -= (fi24 % (((-(fi23))) | 1))
                                        if ((((("a" + ("\n\t" + "foo bar")) + "a") + (((("a" + "hello") + (("\n\t" + "a") + (("" + "") + "hello"))) + (("a" + "foo bar") + "")) + "a")) + "\n\t") == "a") {
                                            fmt.Println(fi23)
                                            v1[(((("" + "") + "\n\t") + "\n\t") + "")] = 960
                                        } else {
                                            v25 = append(v25, v2)
                                            var v28 uint32 = 65535
                                            _ = v28
                                        }
                                    }
                                    v29 := v1[("\n\t" + "\n\t")]
                                    v30 := GenS2{F0: v5, F1: 1000}
                                    _ = v29
                                    _ = v30
                                }
                            }
                            v6["a"] = v3
                            _ = fi24
                            _ = v25
                        }
                    default:
                        v2 = fi23
                    }
                    fmt.Println(v5, v3)
                    var v31 float64 = ((((0.0 + float64((float32(0.0) / 1.0))) / float64((1.5 - (((-2.25 - 0.1) * 1e10) / (1.5 + 0.0))))) - 1e10) + 3.14159)
                    _ = fi23
                    _ = v31
                }
                switch -404 {
                case 782:
                    v32 := GenS3{F0: 0, F1: 1000, F2: -32768}
                    _ = v32
                default:
                    v1["a"] = 310
                }
            }
            if v3 {
                v33 := []int{-566, v2}
                if (42 < -2147483648) {
                    for fi34 := 0; fi34 < 3; fi34++ {
                        for fi35 := 0; fi35 < 9; fi35++ {
                            if len(v33) >= 2 {
                                v33 = v33[1:len(v33)]
                            }
                            _ = fi35
                        }
                        for fi36 := 0; fi36 < 12; fi36++ {
                            if ((fi34 ^ (fi36 % (((fi34 << ((1) & 63))) | 1))) != v2) {
                                v6[("a" + "")] = (0 == 2147483647)
                                var v37 bool = (v3 || (-9223372036854775808 <= 1))
                                v5 += v5
                                v5 += 112
                                _ = v37
                            }
                            if (42 >= 0) {
                                if v3 {
                                    v4 += 937
                                    var v38 uint64 = 2
                                    _ = v38
                                }
                                _ = len(v33)
                            }
                            _ = fi36
                        }
                        _ = fi34
                    }
                    var v39 uint = 42
                    switch "\n\t" {
                    case "a":
                        var v40 bool = false
                        _ = v40
                    case "":
                        v41 := GenS4{F0: true, F1: (^(v4)), F2: 816}
                        _ = v41
                    case "\n\t":
                        _ = len(v33)
                    }
                    _ = v39
                } else {
                    if (0.0 <= -2.25) {
                        var v42 uint64 = 4294967295
                        var v43 bool = v3
                        _ = v42
                        _ = v43
                    }
                }
                _ = v33
            } else {
                var v44 uint32 = 2147483648
                _ = v44
            }
            for fi45 := 0; fi45 < 20; fi45++ {
                if v3 {
                    v6["\n\t"] = (!(!false))
                    if v3 {
                        fmt.Println(v5)
                        for fi46 := 0; fi46 < 6; fi46++ {
                            var v47 uint = 0
                            v48 := GenS3{F0: -9223372036854775808, F1: int16(v47), F2: 0}
                            _ = fi46
                            _ = v47
                            _ = v48
                        }
                    }
                    v49 := []int{42, (v2 >> ((95) & 63))}
                    _ = v49
                }
                var v50 int32 = -909
                var v51 uint8 = v5
                _ = fi45
                _ = v50
                _ = v51
            }
        }
    } else {
        v4 += 0
        v4 = (v4 / ((255) | 1))
    }
    var v52 uint8 = 128
    for fi53 := 0; fi53 < 5; fi53++ {
        if v3 {
            v52 += (v52 * 0)
        } else {
            v52 += (v5 % ((128) | 1))
            switch -934 {
            case -2147483648:
                delete(v6, "\n\t")
            case -9223372036854775808:
                var v54 uint16 = v4
                _ = v54
            }
            fmt.Println(v52, v3, v52)
        }
        var v55 int8 = 1
        _ = fi53
        _ = v55
    }
    v5 = (^(v5))
    for fi56 := 0; fi56 < 10; fi56++ {
        v57 := GenS2{F0: (^(v52)), F1: -1}
        v4 += 619
        if (!(true && v3)) {
            var v58 int = v2
            _ = v58
        } else {
            var v59 float32 = 0.0
            _ = v59
        }
        _ = fi56
        _ = v57
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
