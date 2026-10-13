package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 int16 }
func genFn0() int16 {
    v5 := []int{}
    var v6 int8 = -128
    var v7 uint8 = 1
    fmt.Println(v7, v7)
    var v8 int64 = -9223372036854775808
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 32767
}
func genFn1() uint {
    if (false && (65535 == 2)) {
        var v9 uint = 720
        if (-842 != -1) {
            fmt.Println(v9)
        }
        v10 := GenS1{F0: 127}
        _ = v10
        _ = v9
    }
    var v11 uint = 0
    _ = v11
    return (v11 ^ (v11 | v11))
}

func main() {
    var v1 bool = (float32(1e10) != 1e10)
    var v2 uint64 = 653
    var v3 uint8 = 42
    var v4 uint = 483
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 int = fi12
        switch ("hello" + "hello") {
        case "\n\t":
            switch -2147483648 {
            case 83:
                v14 := []int{}
                _ = v14
            case 9223372036854775807:
                fmt.Println(v2)
            }
        case "a":
            var v15 bool = (v2 >= v2)
            _ = v15
        case "foo bar":
            v16 := []int{v13, v13, fi12}
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    v17 := []int{-172, 1, 850}
    v4 = 4294967295
    var v18 int8 = -54
    for fi19 := 0; fi19 < 1; fi19++ {
        if v1 {
            if v1 {
                switch fi19 {
                case 127:
                    if len(v17) > 0 {
                        v17[0] = fi19
                    }
                }
                var v20 int64 = 1
                _ = v20
            } else {
                _ = len(v17)
                var v21 uint64 = 4294967295
                _ = v21
            }
            v22 := GenS2{F0: ("foo bar" + "\n\t"), F1: -859}
            switch (fi19 / ((int(v4)) | 1)) {
            case 128:
                v23 := GenS2{F0: ("a" + ""), F1: int32(v18)}
                _ = v23
            case 32768:
                _ = len(v17)
            default:
                v24 := GenS1{F0: 0}
                _ = v24
            }
            _ = v22
        }
        var v25 int32 = 1
        if true {
            if ((!(1 < 1)) && ("hello" <= ("\n\t" + "\n\t"))) {
                v3 = v3
                _ = len(v17)
                var v26 int16 = 42
                _ = v26
            }
            var v27 uint = 0
            fmt.Println(v18, v2)
            _ = v27
        } else {
            if len(v17) > 0 {
                v17[0] = (^(fi19))
            }
            if (2 <= (^(v4))) {
                var v28 string = "\n\t"
                _ = v28
            }
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
        }
        _ = fi19
        _ = v25
    }
    if ((127 >= -1) || true) {
        var v29 int32 = 42
        _ = v29
    } else {
        switch (("\n\t" + (("hello" + ("\n\t" + (("hello" + ("foo bar" + "hello")) + "\n\t"))) + "hello")) + "\n\t") {
        case "foo bar":
            if v1 {
                fmt.Println(v3, v3, v18)
                v30 := GenS1{F0: (v3 ^ (v3 | 198))}
                for fi31 := 0; fi31 < 4; fi31++ {
                    v32 := GenS1{F0: (v3 % (((v3 << ((18446744073709551615) & 63))) | 1))}
                    if false {
                        for fi33 := 0; fi33 < 20; fi33++ {
                            v1 = (!(!((v1 && (float32(1.5) <= 0.0)) || (!(!false)))))
                            fmt.Println(v3, fi33, v4)
                            _ = fi33
                        }
                        v18 = v18
                        if (2147483648 < 2) {
                            v34 := GenS1{F0: (v3 >> (((v4 & v4)) & 63))}
                            fmt.Println(fi31)
                            switch (fi31 & (fi31 - (fi31 / (((fi31 << (((v4 / (((v4 << ((genFn1()) & 63))) | 1))) & 63))) | 1)))) {
                            case 0:
                                switch "\n\t" {
                                case "foo bar":
                                    v35 := GenS1{F0: (v3 >> ((uint(v3)) & 63))}
                                    _ = v35
                                case "a":
                                    v36 := GenS2{F0: ("foo bar" + "\n\t"), F1: -223}
                                    _ = v36
                                case "\n\t":
                                    v18 = -128
                                }
                            case 127:
                                var v37 uint64 = v2
                                _ = v37
                            }
                            _ = v34
                        } else {
                            switch ("a" + "foo bar") {
                            case "a":
                                var v38 int16 = 42
                                _ = v38
                            case "":
                                var v39 int64 = -828
                                _ = v39
                            }
                            v18 = -128
                            v40 := GenS1{F0: v3}
                            _ = v40
                        }
                    }
                    _ = fi31
                    _ = v32
                }
                _ = v30
            } else {
                var v41 int = -2147483648
                v17 = append(v17, v41)
                _ = v41
            }
        case "a":
            v3 = (v3 / ((v3) | 1))
        }
        for fi42 := 0; fi42 < 20; fi42++ {
            v43 := GenS3{F0: v2, F1: 65535}
            _ = fi42
            _ = v43
        }
        for fi44 := 0; fi44 < 12; fi44++ {
            v1 = v1
            _ = fi44
        }
    }
    if (genFn1() >= v4) {
        var v45 float64 = -2.25
        v45 = (v45 - 3.14159)
        _ = v45
    } else {
        v3 = (^(v3))
        v1 = v1
        v17 = append(v17, 4294967295)
    }
    var v46 uint8 = 42
    fmt.Println(v1, v4, v18)
    for fi47 := 0; fi47 < 18; fi47++ {
        v48 := GenS2{F0: "hello", F1: 2147483647}
        v49 := GenS4{F0: 1000}
        for fi50 := 0; fi50 < 10; fi50++ {
            v4 -= (^(v4))
            switch -1 {
            case 32767:
                if ((0 < -700) && true) {
                    var v51 int32 = 1
                    v51 -= v51
                    v3 = v46
                    var v52 float32 = (1e10 * 0.1)
                    _ = v51
                    _ = v52
                }
            case -9:
                v18 = -1
            case -9223372036854775808:
                if v1 {
                    var v53 bool = (v18 < v18)
                    fmt.Println(v18, v46, v4)
                    _ = v53
                } else {
                    switch (("\n\t" + "hello") + ("hello" + "foo bar")) {
                    case "a":
                        var v54 float32 = ((1.5 - float32(0.0)) - (0.1 * -2.25))
                        v55 := GenS2{F0: "\n\t", F1: -677}
                        _ = v54
                        _ = v55
                    case "\n\t":
                        v1 = v1
                    }
                }
            default:
                switch ("" + ("hello" + "\n\t")) {
                case "\n\t":
                    v3 += (^(v46))
                case "foo bar":
                    var v56 int32 = 42
                    _ = v56
                case "":
                    var v57 int32 = -2147483648
                    _ = v57
                }
            }
            v1 = false
            var v58 int16 = 42
            _ = len(v17)
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            v59 := GenS4{F0: (v58 ^ 124)}
            var v60 string = ("hello" + "hello")
            _ = fi50
            _ = v58
            _ = v59
            _ = v60
        }
        _ = fi47
        _ = v48
        _ = v49
    }
    var v61 float64 = ((3.14159 / 3.14159) / -2.25)
    v17 = append(v17, 9223372036854775807)
    fmt.Println(v3)
    var v62 float64 = ((v61 + float64(1.5)) + v61)
    var v63 int64 = 966
    if len(v17) > 0 {
        v17[0] = 513
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v61
    _ = v62
    _ = v63
}
