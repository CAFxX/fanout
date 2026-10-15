package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 string }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int64; F1 int16 }
func genFn0(p0 uint32) string {
    v4 := GenS1{F0: false, F1: 32767, F2: "hello"}
    v5 := []int{42, 42, -42, -1}
    _ = p0
    _ = v4
    _ = v5
    return ""
}
func genFn1() int32 {
    v6 := []int{127}
    if false {
        if false {
            var v7 bool = true
            _ = v7
        }
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 int32 = 0
            v9 = v9
            _ = fi8
            _ = v9
        }
    } else {
        var v10 int32 = -967
        fmt.Println(v10)
        v11 := GenS1{F0: true, F1: -32768, F2: (("\n\t" + "foo bar") + ("hello" + ("foo bar" + ("a" + "foo bar"))))}
        _ = v10
        _ = v11
    }
    if false {
        var v12 uint64 = 34
        v12 = (^(v12))
        v12 = v12
        _ = v12
    } else {
        var v13 int64 = 0
        switch 128 {
        case -860:
            var v14 int16 = -32768
            _ = v14
        case 127:
            for fi15 := 0; fi15 < 14; fi15++ {
                fmt.Println(v13)
                v13 -= (-(v13))
                _ = fi15
            }
        default:
            var v16 string = "\n\t"
            _ = v16
        }
        _ = v13
    }
    var v17 int = -53
    _ = v17
    _ = v6
    return -530
}

func main() {
    v1 := []int{32768, 907, 1, 2147483647}
    var v2 int = 2147483647
    var v3 uint = 2
    fmt.Println(v2, v2)
    v3 = 634
    fmt.Println(v3, v2)
    v18 := GenS1{F0: false, F1: 1, F2: "\n\t"}
    if true {
        var v19 string = ("foo bar" + "a")
        if (1 <= 384) {
            v20 := GenS1{F0: true, F1: 42, F2: v19}
            _ = v20
        } else {
            if len(v1) > 0 {
                v1[0] = v2
            }
            _ = len(v1)
            v3 -= 113
        }
        _ = v19
    }
    switch "foo bar" {
    case "\n\t":
        var v21 int32 = 2147483647
        _ = v21
    }
    if (-698 >= -32768) {
        var v22 uint32 = 575
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := GenS1{F0: false, F1: -78, F2: ((("hello" + "\n\t") + ("hello" + "")) + (("" + (("a" + "") + "foo bar")) + "foo bar"))}
            _ = fi23
            _ = v24
        }
        _ = v22
    } else {
        v2 = v2
    }
    var v25 uint32 = 1
    var v26 int8 = 127
    switch (("\n\t" + "") + ("hello" + "")) {
    case "a":
        if false {
            var v27 int32 = 1
            v28 := GenS1{F0: (v25 >= uint32(v3)), F1: 42, F2: ""}
            var v29 int8 = v26
            _ = v27
            _ = v28
            _ = v29
        } else {
            v3 = (^(v3))
        }
    case "\n\t":
        v26 = -90
    case "hello":
        _ = len(v1)
    default:
        v26 = v26
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        v3 = v3
        for fi31 := 0; fi31 < 19; fi31++ {
            v26 = 0
            _ = fi31
        }
        if ((1.5 >= -2.25) || false) {
            v32 := GenS1{F0: (1 >= 2), F1: -318, F2: "a"}
            fmt.Println(v25, v3, fi30)
            _ = v32
        } else {
            v33 := GenS1{F0: false, F1: 1000, F2: "foo bar"}
            _ = v33
        }
        _ = fi30
    }
    if len(v1) > 0 {
        v1[0] = v2
    }
    var v34 int = v2
    if (!((v26 & (v26 & v26)) > 0)) {
        if ((true || (true && false)) || false) {
            v35 := GenS1{F0: true, F1: -32768, F2: "a"}
            _ = v35
        }
    } else {
        if true {
            if false {
                for fi36 := 0; fi36 < 6; fi36++ {
                    for fi37 := 0; fi37 < 5; fi37++ {
                        var v38 uint32 = 65535
                        switch "\n\t" {
                        case "\n\t":
                            if (!(567 != 0)) {
                                fmt.Println(v2)
                                var v39 string = "hello"
                                v26 -= v26
                                _ = v39
                            } else {
                                v26 += (v26 - v26)
                            }
                        case "":
                            if len(v1) > 0 {
                                v1[0] = v34
                            }
                        default:
                            v40 := GenS2{F0: 42, F1: 750, F2: 2}
                            _ = v40
                        }
                        v25 += v38
                        _ = fi37
                        _ = v38
                    }
                    var v41 float32 = 0.1
                    for fi42 := 0; fi42 < 15; fi42++ {
                        _ = len(v1)
                        v41 += v41
                        v43 := GenS1{F0: (fi42 >= fi42), F1: -1, F2: ""}
                        _ = fi42
                        _ = v43
                    }
                    _ = fi36
                    _ = v41
                }
            } else {
                var v44 int64 = -1
                _ = v44
            }
            for fi45 := 0; fi45 < 9; fi45++ {
                v34 = (fi45 / (((v34 >> ((v3) & 63))) | 1))
                if (-813 > 9223372036854775807) {
                    v46 := GenS3{F0: 42, F1: (!("foo bar" != "hello"))}
                    _ = v46
                } else {
                    var v47 int8 = int8(v3)
                    switch "foo bar" {
                    case "hello":
                        var v48 int8 = (v47 ^ v47)
                        _ = v48
                    case "foo bar":
                        if false {
                            _ = len(v1)
                            if (81 <= 42) {
                                fmt.Println(v26, fi45, v47)
                                for fi49 := 0; fi49 < 10; fi49++ {
                                    v34 = int(v25)
                                    _ = fi49
                                }
                                v1 = append(v1, v2)
                            }
                            var v50 uint = v3
                            _ = v50
                        } else {
                            v34 = v34
                            v51 := GenS4{F0: -160, F1: 1}
                            for fi52 := 0; fi52 < 1; fi52++ {
                                var v53 float64 = ((-2.25 - float64((1e10 * -2.25))) + ((3.14159 + 0.0) * 3.14159))
                                _ = fi52
                                _ = v53
                            }
                            _ = v51
                        }
                    }
                    switch v2 {
                    case 4294967295:
                        v47 = -69
                    default:
                        if false {
                            var v54 uint = v3
                            fmt.Println(v34, v2)
                            fmt.Println(v3)
                            _ = v54
                        }
                    }
                    _ = v47
                }
                _ = fi45
            }
            var v55 string = "hello"
            _ = v55
        }
        fmt.Println(v34, v2, v34)
    }
    v1 = append(v1, 128)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v34
}
