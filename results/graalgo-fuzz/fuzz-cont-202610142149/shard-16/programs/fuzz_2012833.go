package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int16 }
type GenS2 struct { F0 int32; F1 string; F2 uint }
type GenS3 struct { F0 string; F1 uint; F2 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 int32 }
func genFn0() int8 {
    v5 := []int{-943, 9223372036854775807}
    switch "a" {
    case "foo bar":
        var v6 int32 = -474
        _ = v6
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    _ = v5
    return -3
}
func genFn1() int16 {
    v7 := []int{4294967295, 4294967295}
    if (!(float64(0.1) <= ((1e10 + 1e10) * 0.1))) {
        var v8 int = -152
        _ = v8
    } else {
        var v9 uint8 = 255
        if len(v7) > 0 {
            v7[0] = -9223372036854775808
        }
        v10 := GenS1{F0: 65535, F1: -1, F2: -32768}
        _ = v10
        _ = v9
    }
    var v11 int64 = 42
    v11 -= v11
    if false {
        v12 := GenS1{F0: 193, F1: -2147483648, F2: 1}
        v13 := GenS1{F0: 4294967295, F1: 961, F2: 1000}
        switch -468 {
        case -32:
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 int = (fi14 << ((65535) & 63))
                for fi16 := 0; fi16 < 2; fi16++ {
                    var v17 string = "a"
                    var v18 bool = (1 != -1)
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = fi14
                _ = v15
            }
        case -240:
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
        _ = v12
        _ = v13
    } else {
        v11 = 9223372036854775807
        fmt.Println(v11)
        var v19 int8 = 100
        _ = v19
    }
    _ = v11
    _ = v7
    return 0
}

func main() {
    v1 := []int{522}
    v2 := []string{}
    v3 := []string{"foo bar", "hello"}
    var v4 int64 = 42
    if (707 <= 0) {
        if false {
            if true {
                v20 := GenS2{F0: 0, F1: "a", F2: 255}
                _ = v20
            } else {
                switch "\n\t" {
                case "\n\t":
                    _ = len(v2)
                case "a":
                    var v21 int16 = genFn1()
                    _ = v21
                default:
                    _ = len(v3)
                }
            }
            _ = len(v3)
            for fi22 := 0; fi22 < 20; fi22++ {
                switch (fi22 / ((fi22) | 1)) {
                case 26:
                    var v23 uint64 = 990
                    _ = v23
                default:
                    v3 = append(v3, (("a" + "hello") + "\n\t"))
                }
                for fi24 := 0; fi24 < 18; fi24++ {
                    if (986 >= 42) {
                        var v25 int32 = -736
                        switch "\n\t" {
                        case "":
                            if false {
                                _ = len(v1)
                                if false {
                                    if false {
                                        var v26 int32 = v25
                                        var v27 uint = 959
                                        _ = v26
                                        _ = v27
                                    }
                                    v28 := GenS3{F0: (("hello" + "a") + ("\n\t" + ("a" + "foo bar"))), F1: 360, F2: "foo bar"}
                                    _ = v28
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v4 = -9223372036854775808
                                    switch "a" {
                                    case "a":
                                        fmt.Println(v4, fi22)
                                    case "hello":
                                        fmt.Println(v4, fi22)
                                    case "\n\t":
                                        fmt.Println(fi22)
                                    default:
                                        fmt.Println(fi22, fi22)
                                    }
                                    var v29 uint64 = 2
                                    _ = v29
                                }
                                v25 = (v25 ^ 0)
                                v25 = v25
                            } else {
                                var v30 float32 = 1.5
                                _ = v30
                            }
                        }
                        _ = v25
                    }
                    v31 := GenS2{F0: -380, F1: "\n\t", F2: 255}
                    v4 += (v4 & (v4 | v4))
                    _ = fi24
                    _ = v31
                }
                _ = fi22
            }
        }
        v2 = append(v2, "a")
        switch "foo bar" {
        case "\n\t":
            _ = len(v2)
        case "a":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "":
            v32 := GenS4{F0: 294}
            _ = v32
        }
    } else {
        v33 := GenS1{F0: 1, F1: -929, F2: -505}
        _ = v33
    }
    var v34 string = "foo bar"
    if false {
        if (!(2 > 291)) {
            var v35 int8 = 100
            var v36 uint32 = 881
            _ = v35
            _ = v36
        }
        switch v34 {
        case "foo bar":
            switch v34 {
            case "":
                var v37 int16 = 1000
                _ = v37
            }
        case "\n\t":
            var v38 float64 = 0.1
            _ = v38
        case "hello":
            v4 = v4
        default:
            v39 := GenS5{F0: 68, F1: 629}
            _ = v39
        }
    }
    var v40 int = int(v4)
    v40 = 128
    v34 = ("hello" + "hello")
    switch v40 {
    case -463:
        v2 = append(v2, v34)
    case 127:
        if (716 == 4294967295) {
            v40 = (v40 % (((v40 << ((673) & 63))) | 1))
        } else {
            var v41 string = "a"
            fmt.Println(v40, v4)
            if (1 == -2147483648) {
                if len(v1) > 0 {
                    v1[0] = 42
                }
                for fi42 := 0; fi42 < 1; fi42++ {
                    if (!(42 <= 67)) {
                        fmt.Println(v4, v40)
                        for fi43 := 0; fi43 < 3; fi43++ {
                            var v44 uint = 255
                            _ = fi43
                            _ = v44
                        }
                        _ = len(v2)
                    } else {
                        v40 = fi42
                    }
                    _ = fi42
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                var v45 int = (-(v40))
                v46 := GenS5{F0: (v40 + (v40 ^ v40)), F1: 113}
                _ = v45
                _ = v46
            }
            _ = v41
        }
    case 886:
        v47 := GenS4{F0: 1}
        _ = v47
    }
    switch ("a" + "a") {
    case "\n\t":
        v4 = v4
    case "hello":
        var v48 uint16 = 30
        _ = v48
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v40
}
