package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS5 struct { F0 uint }

func main() {
    var v1 float64 = -2.25
    var v2 uint16 = 290
    var v3 float64 = (1e10 - v1)
    var v4 uint32 = 2
    var v5 bool = true
    v1 = (v3 / 3.14159)
    v5 = (!v5)
    if (v5 && false) {
        for fi6 := 0; fi6 < 10; fi6++ {
            fmt.Println(v3)
            v2 = v2
            _ = fi6
        }
        var v7 string = (("\n\t" + ("foo bar" + ("hello" + "hello"))) + "")
        _ = v7
    }
    var v8 uint32 = v4
    fmt.Println(v8, v1)
    if false {
        v9 := GenS1{F0: 0}
        var v10 int64 = 0
        var v11 uint8 = 128
        _ = v10
        _ = v11
        _ = v9
    } else {
        switch ("\n\t" + ("" + "foo bar")) {
        case "foo bar":
            switch 32768 {
            case -953:
                for fi12 := 0; fi12 < 4; fi12++ {
                    v2 = v2
                    _ = fi12
                }
            case 319:
                v13 := []int{-705, 112, 32767, 9223372036854775807}
                _ = v13
            case 2147483647:
                if (670 == 699) {
                    for fi14 := 0; fi14 < 9; fi14++ {
                        var v15 int32 = -548
                        var v16 int8 = -22
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    v2 = v2
                    switch 128 {
                    case -1:
                        var v17 float32 = 3.14159
                        _ = v17
                    case 1:
                        v18 := GenS1{F0: 2}
                        _ = v18
                    case 32767:
                        switch "\n\t" {
                        case "foo bar":
                            switch (("a" + ("" + "a")) + ("foo bar" + "foo bar")) {
                            case "hello":
                                switch "hello" {
                                case "\n\t":
                                    if true {
                                        if ((1 >= 0) && v5) {
                                            switch ((("foo bar" + "") + (("a" + (("\n\t" + "") + ("" + ""))) + (("a" + ("hello" + "hello")) + ((("foo bar" + "hello") + ("foo bar" + ((("hello" + "a") + ("\n\t" + "a")) + "foo bar"))) + "hello")))) + "") {
                                            case "hello":
                                                if true {
                                                    if true {
                                                        var v19 uint8 = 1
                                                        _ = v19
                                                    }
                                                } else {
                                                    switch "a" {
                                                    case "a":
                                                        fmt.Println(v5, v3, v3)
                                                    case "foo bar":
                                                        var v20 int16 = -1
                                                        _ = v20
                                                    case "\n\t":
                                                        if v5 {
                                                            fmt.Println(v5)
                                                        } else {
                                                            fmt.Println(v2)
                                                        }
                                                    }
                                                }
                                            case "\n\t":
                                                var v21 uint = 1
                                                _ = v21
                                            case "":
                                                switch "\n\t" {
                                                case "a":
                                                    switch -1 {
                                                    case 127:
                                                        var v22 int64 = 1
                                                        _ = v22
                                                    default:
                                                        v8 += v4
                                                    }
                                                default:
                                                    var v23 uint = 42
                                                    _ = v23
                                                }
                                            default:
                                                if v5 {
                                                    v24 := GenS1{F0: 2}
                                                    _ = v24
                                                } else {
                                                    var v25 int16 = 442
                                                    var v26 bool = false
                                                    v27 := GenS1{F0: 255}
                                                    _ = v25
                                                    _ = v26
                                                    _ = v27
                                                }
                                            }
                                            fmt.Println(v1)
                                            v8 = v4
                                        }
                                        switch ("a" + (("hello" + "") + ("hello" + "a"))) {
                                        case "":
                                            var v28 int64 = 42
                                            _ = v28
                                        default:
                                            for fi29 := 0; fi29 < 3; fi29++ {
                                                v1 = (1.5 * (v3 + float64(-2.25)))
                                                v30 := []int{fi29, fi29, -409}
                                                _ = fi29
                                                _ = v30
                                            }
                                        }
                                        for fi31 := 0; fi31 < 4; fi31++ {
                                            var v32 float64 = (v1 / v1)
                                            if v5 {
                                                var v33 uint64 = 980
                                                v34 := GenS2{F0: (v33 - 4294967295)}
                                                _ = v33
                                                _ = v34
                                            }
                                            _ = fi31
                                            _ = v32
                                        }
                                    }
                                case "":
                                    var v35 int16 = -436
                                    _ = v35
                                default:
                                    for fi36 := 0; fi36 < 17; fi36++ {
                                        var v37 float32 = 0.1
                                        if v5 {
                                            switch "\n\t" {
                                            case "foo bar":
                                                v38 := []int{(fi36 >> ((469) & 63)), 0}
                                                _ = v38
                                            case "\n\t":
                                                v39 := []int{(fi36 + (fi36 >> ((18446744073709551615) & 63))), -1}
                                                _ = v39
                                            case "a":
                                                v3 -= v3
                                            }
                                        }
                                        _ = fi36
                                        _ = v37
                                    }
                                }
                            case "":
                                if ((36 == 42) || v5) {
                                    var v40 uint = 566
                                    v1 = v1
                                    _ = v40
                                }
                            case "a":
                                v4 += 987
                            default:
                                v41 := []int{418, int(v4), 54}
                                _ = v41
                            }
                        default:
                            v2 = (v2 ^ (v2 >> ((42) & 63)))
                        }
                    }
                }
            }
        case "a":
            v42 := []int{-164, 0}
            _ = v42
        case "":
            for fi43 := 0; fi43 < 14; fi43++ {
                if (((true && true) || v5) && (v2 == (v2 ^ (v2 / ((32767) | 1))))) {
                    for fi44 := 0; fi44 < 3; fi44++ {
                        if false {
                            if ((-2.25 * (3.14159 / 1.0)) > -2.25) {
                                var v45 int32 = -1
                                _ = v45
                            } else {
                                var v46 uint16 = (^(v2))
                                var v47 int8 = 127
                                var v48 uint8 = uint8(v4)
                                _ = v46
                                _ = v47
                                _ = v48
                            }
                            if false {
                                v5 = true
                                v4 = v4
                                v49 := GenS3{F0: 42, F1: "hello"}
                                var v50 float64 = v3
                                _ = v49
                                _ = v50
                            } else {
                                if v5 {
                                    v51 := GenS4{F0: 769, F1: 2, F2: (^(v2))}
                                    v52 := []int{59, (fi43 - 1), fi44, fi43}
                                    v2 = (v2 << ((395) & 63))
                                    v52 = append(v52, (fi43 - fi43))
                                    v2 = v2
                                    v1 = 0.0
                                    _ = v51
                                    _ = v52
                                }
                            }
                        } else {
                            var v53 uint16 = 65535
                            var v54 uint32 = 65535
                            var v55 string = "a"
                            _ = v53
                            _ = v54
                            _ = v55
                        }
                        var v56 uint8 = 127
                        switch "foo bar" {
                        case "":
                            v57 := GenS5{F0: 2}
                            _ = v57
                        case "a":
                            var v58 int16 = 954
                            _ = v58
                        }
                        _ = fi44
                        _ = v56
                    }
                } else {
                    switch "a" {
                    case "\n\t":
                        for fi59 := 0; fi59 < 5; fi59++ {
                            v60 := GenS5{F0: 65535}
                            _ = fi59
                            _ = v60
                        }
                    case "a":
                        var v61 uint = 394
                        _ = v61
                    case "":
                        v8 = (v8 % ((v4) | 1))
                    }
                    var v62 uint16 = 42
                    _ = v62
                }
                v63 := []int{fi43, (fi43 / (((-(fi43))) | 1))}
                _ = fi43
                _ = v63
            }
        }
        switch "" {
        case "\n\t":
            switch -9223372036854775808 {
            case 824:
                v2 = (v2 % ((v2) | 1))
            default:
                if (1 != 219) {
                    v5 = true
                } else {
                    v64 := GenS5{F0: 18446744073709551615}
                    _ = v64
                }
            }
        }
        v3 = (((v3 + 1e10) - 0.1) * 1e10)
    }
    var v65 uint8 = 255
    v66 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    fmt.Println("slicelen v66:", len(v66), v66)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
    _ = v8
}
