package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int8) int64 {
    switch "foo bar" {
    case "":
        var v5 bool = false
        _ = v5
    case "hello":
        var v6 uint64 = 1
        _ = v6
    case "foo bar":
        p0 = (^(p0))
    default:
        p0 += (p0 & p0)
    }
    p0 = (p0 % ((p0) | 1))
    switch (("\n\t" + "a") + "foo bar") {
    case "\n\t":
        p0 = (p0 | 0)
    case "a":
        v7 := GenS1{F0: true, F1: 18446744073709551615}
        _ = v7
    default:
        var v8 int32 = 908
        _ = v8
    }
    if false {
        if true {
            var v9 string = "hello"
            _ = v9
        }
    } else {
        if (-48 < 0) {
            if (-9223372036854775808 == -331) {
                for fi10 := 0; fi10 < 5; fi10++ {
                    var v11 int16 = 42
                    _ = fi10
                    _ = v11
                }
                v12 := []int{}
                if false {
                    p0 = -1
                    p0 -= p0
                } else {
                    p0 -= -128
                }
                _ = v12
            }
        }
        var v13 string = ""
        _ = v13
    }
    if (18446744073709551615 >= 4294967295) {
        p0 -= 42
    } else {
        var v14 uint32 = 65535
        _ = v14
    }
    _ = p0
    return 42
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    var v4 int = -9223372036854775808
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS2{F0: 2, F1: (941 >= 255), F2: 0}
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 int = 128
            v19 := GenS2{F0: 128, F1: v1, F2: 65535}
            var v20 int8 = 127
            v1 = true
            var v21 int64 = -9223372036854775808
            v4 = v18
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
            _ = v21
        }
        fmt.Println(v1)
        _ = fi15
        _ = v16
    }
    v4 -= v4
    v22 := GenS3{F0: 2147483648}
    fmt.Println(v1, v4)
    fmt.Println(v1, v1)
    v1 = (v4 >= 4294967295)
    v23 := GenS1{F0: ((2147483648 >= 2147483647) && (-1 != 42)), F1: 0}
    for fi24 := 0; fi24 < 12; fi24++ {
        for fi25 := 0; fi25 < 12; fi25++ {
            v26 := GenS4{F0: -629, F1: v4}
            v4 = (v4 & fi24)
            v3["hello"] = (("\n\t" + ("" + (("\n\t" + ("a" + "\n\t")) + "a"))) + "hello")
            _ = fi25
            _ = v26
        }
        v27 := GenS4{F0: -975, F1: (v4 / ((fi24) | 1))}
        if false {
            delete(v2, 2147483647)
            var v28 uint16 = 0
            _ = v28
        }
        _ = fi24
        _ = v27
    }
    if true {
        for fi29 := 0; fi29 < 18; fi29++ {
            var v30 int16 = 1
            _ = fi29
            _ = v30
        }
        for fi31 := 0; fi31 < 19; fi31++ {
            v32 := []int{}
            if ((-2.25 + (3.14159 - 0.0)) > -2.25) {
                var v33 bool = true
                _ = v33
            } else {
                if (!true) {
                    delete(v2, 65535)
                }
                for fi34 := 0; fi34 < 18; fi34++ {
                    v4 = 127
                    if true {
                        v1 = (("\n\t" <= "a") || v1)
                        v1 = v1
                        var v35 bool = v1
                        if (2 != 2147483648) {
                            v1 = (((42 < -1) || (0 >= 0)) || (63 == -128))
                            v1 = v35
                            switch -1 {
                            case -712:
                                var v36 uint32 = 2
                                _ = v36
                            case -1:
                                var v37 int8 = 1
                                _ = v37
                            }
                        } else {
                            if v1 {
                                v35 = v1
                                v35 = (!v35)
                                if true {
                                    var v38 int = -225
                                    _ = v38
                                } else {
                                    v35 = ("hello" > ("" + ("a" + ("foo bar" + "hello"))))
                                }
                            } else {
                                v4 = (v4 | (fi34 * fi34))
                                v4 = (v4 * fi34)
                            }
                        }
                        _ = v35
                    } else {
                        fmt.Println(fi34, fi34)
                        v4 = (fi31 | (^(fi34)))
                        fmt.Println(v1, fi34)
                        v39 := GenS4{F0: 42, F1: v4}
                        _ = v39
                    }
                    _ = fi34
                }
                v3[(("" + "") + "\n\t")] = "\n\t"
            }
            _ = fi31
            _ = v32
        }
        var v40 float32 = (float32(3.14159) * 0.1)
        _ = v40
    } else {
        v41 := GenS3{F0: 1}
        if (true && v1) {
            for fi42 := 0; fi42 < 14; fi42++ {
                var v43 int = fi42
                _ = fi42
                _ = v43
            }
            var v44 float32 = -2.25
            _ = v44
        } else {
            var v45 float64 = (0.1 / (3.14159 - -2.25))
            for fi46 := 0; fi46 < 1; fi46++ {
                _ = len(v2)
                switch (v4 | v4) {
                case -680:
                    var v47 uint16 = 483
                    _ = v47
                case 127:
                    v48 := GenS3{F0: 65535}
                    _ = v48
                case 0:
                    fmt.Println(v1, v4, v1)
                default:
                    fmt.Println(v4, v4, fi46)
                }
                _ = fi46
            }
            _ = v45
        }
        v49 := GenS4{F0: -270, F1: v4}
        _ = v41
        _ = v49
    }
    if v1 {
        v50 := []int{}
        if false {
            v51 := v3["hello"]
            v50 = append(v50, (v4 | (v4 | -851)))
            _ = len(v50)
            _ = v51
        } else {
            var v52 int64 = 0
            if v1 {
                var v53 int64 = v52
                v54 := GenS1{F0: v1, F1: 422}
                v1 = ((-2.25 <= 3.14159) && v1)
                _ = v53
                _ = v54
            }
            _ = v52
        }
        _ = v50
    } else {
        switch "\n\t" {
        case "\n\t":
            var v55 bool = v1
            _ = v55
        case "":
            if v1 {
                for fi56 := 0; fi56 < 11; fi56++ {
                    switch "\n\t" {
                    case "a":
                        v2[995] = (("a" + ("foo bar" + ("hello" + "a"))) + ("\n\t" + "hello"))
                    case "":
                        v57 := GenS5{F0: 253}
                        _ = v57
                    }
                    _ = fi56
                }
            }
        case "hello":
            var v58 float32 = float32(1.5)
            _ = v58
        default:
            v59 := GenS1{F0: (v1 && v1), F1: 18446744073709551615}
            _ = v59
        }
        v60 := []int{0, v4}
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
