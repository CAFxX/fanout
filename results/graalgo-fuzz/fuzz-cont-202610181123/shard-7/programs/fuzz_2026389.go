package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint16; F1 string; F2 uint32 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 int32; F1 int; F2 int8 }
type GenS5 struct { F0 uint; F1 uint32; F2 int16 }

func main() {
    var v1 int8 = -128
    v2 := []int{-9223372036854775808, 429}
    v3 := make(map[uint32]string)
    if true {
        if true {
            switch "foo bar" {
            case "":
                v4 := GenS1{F0: 1, F1: 15}
                _ = v4
            default:
                v5 := GenS1{F0: -972, F1: 18446744073709551615}
                _ = v5
            }
            if false {
                v6 := GenS1{F0: 0, F1: 494}
                switch "hello" {
                case "hello":
                    fmt.Println(v1)
                case "a":
                    if false {
                        for fi7 := 0; fi7 < 4; fi7++ {
                            v1 -= (v1 >> ((18446744073709551615) & 63))
                            var v8 uint = 1
                            _ = fi7
                            _ = v8
                        }
                        var v9 float64 = 0.0
                        _ = v9
                    }
                }
                _ = v6
            }
        } else {
            if ((0.0 >= 1.5) || false) {
                var v10 bool = true
                switch ("\n\t" + "a") {
                case "a":
                    v11 := GenS1{F0: 42, F1: 0}
                    _ = v11
                case "hello":
                    var v12 int8 = v1
                    _ = v12
                }
                v1 = 42
                _ = v10
            }
        }
    } else {
        v13 := GenS2{F0: 2, F1: ("" + ("a" + "")), F2: 0}
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        v2 = append(v2, fi14)
        _ = fi14
    }
    v15 := GenS1{F0: -1, F1: 897}
    for fi16 := 0; fi16 < 11; fi16++ {
        v2 = append(v2, fi16)
        _ = len(v3)
        _ = fi16
    }
    v1 = (v1 % ((v1) | 1))
    if false {
        for fi17 := 0; fi17 < 4; fi17++ {
            if false {
                v3[1] = "\n\t"
                v1 = v1
                v18 := GenS1{F0: -1, F1: 2}
                _ = v18
            }
            v19 := GenS3{F0: 100, F1: -1}
            _ = fi17
            _ = v19
        }
        if true {
            var v20 uint32 = 65535
            _ = v20
        } else {
            for fi21 := 0; fi21 < 15; fi21++ {
                v22 := GenS4{F0: -2147483648, F1: 234, F2: v1}
                v23 := GenS1{F0: -726, F1: 2}
                v24 := GenS2{F0: 42, F1: "hello", F2: 661}
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            fmt.Println(v1)
        }
    } else {
        v1 += (v1 / ((101) | 1))
        v1 -= 42
    }
    if (-745 >= -32768) {
        for fi25 := 0; fi25 < 12; fi25++ {
            v1 -= (v1 - (v1 ^ 52))
            v1 += (v1 & 1)
            _ = fi25
        }
        if false {
            switch ("\n\t" + "\n\t") {
            case "":
                for fi26 := 0; fi26 < 2; fi26++ {
                    v1 = (v1 + (v1 ^ (v1 & (v1 - v1))))
                    fmt.Println(v1)
                    for fi27 := 0; fi27 < 15; fi27++ {
                        _ = len(v3)
                        v1 += (v1 >> ((65535) & 63))
                        v28 := GenS2{F0: 32767, F1: "\n\t", F2: 42}
                        _ = fi27
                        _ = v28
                    }
                    _ = fi26
                }
            case "hello":
                switch -397 {
                case 1:
                    fmt.Println(v1)
                case 4294967295:
                    if false {
                        for fi29 := 0; fi29 < 2; fi29++ {
                            var v30 int8 = (-(v1))
                            var v31 float32 = float32(((0.0 / 1e10) - 1e10))
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                    } else {
                        var v32 uint64 = 1
                        _ = v32
                    }
                default:
                    for fi33 := 0; fi33 < 3; fi33++ {
                        v34 := GenS3{F0: 42, F1: -1}
                        _ = fi33
                        _ = v34
                    }
                }
            case "a":
                var v35 int64 = 1
                _ = v35
            default:
                fmt.Println(v1)
            }
            for fi36 := 0; fi36 < 12; fi36++ {
                v37 := GenS4{F0: -119, F1: fi36, F2: 127}
                switch fi36 {
                case 209:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case -9223372036854775808:
                    v1 += (-(v1))
                case 2147483647:
                    var v38 int = (fi36 % ((fi36) | 1))
                    _ = v38
                }
                _ = fi36
                _ = v37
            }
            if true {
                var v39 string = "\n\t"
                _ = v39
            } else {
                v1 = -1
                var v40 int32 = -382
                _ = v40
            }
        }
        var v41 uint64 = 2
        _ = v41
    }
    v1 = v1
    switch (("foo bar" + "hello") + "") {
    case "":
        if true {
            v1 -= (-(v1))
            _ = len(v2)
        } else {
            v42 := v3[698]
            for fi43 := 0; fi43 < 14; fi43++ {
                if true {
                    if (false && (-1 <= 293)) {
                        switch v42 {
                        case "a":
                            v44 := GenS5{F0: 1, F1: 1, F2: 925}
                            _ = v44
                        case "hello":
                            for fi45 := 0; fi45 < 14; fi45++ {
                                v46 := GenS3{F0: (v1 % (((v1 << ((2) & 63))) | 1)), F1: -9223372036854775808}
                                v42 = (((v42 + (v42 + "hello")) + "foo bar") + "a")
                                fmt.Println(fi43)
                                _ = len(v3)
                                _ = fi45
                                _ = v46
                            }
                        default:
                            switch (v42 + (v42 + v42)) {
                            case "\n\t":
                                var v47 float32 = 1e10
                                _ = v47
                            case "foo bar":
                                v42 = (((v42 + v42) + v42) + ((v42 + ("\n\t" + v42)) + (v42 + v42)))
                            }
                        }
                        if ((v42 > (("hello" + (v42 + "")) + (v42 + v42))) && false) {
                            if false {
                                var v48 uint = 14
                                var v49 int64 = -516
                                if ((-2.25 * (3.14159 - 0.1)) != float64(1e10)) {
                                    v1 -= 127
                                    v1 -= v1
                                    v48 += v48
                                } else {
                                    var v50 uint64 = 18446744073709551615
                                    _ = v50
                                }
                                _ = v48
                                _ = v49
                            } else {
                                v1 += v1
                            }
                        }
                        fmt.Println(v1)
                    } else {
                        fmt.Println(v1, v42, v1)
                        if false {
                            v51 := v3[2147483648]
                            var v52 int32 = 0
                            _ = v51
                            _ = v52
                        } else {
                            switch (fi43 - fi43) {
                            case 32767:
                                var v53 string = ((("foo bar" + ((v42 + "\n\t") + v42)) + v42) + (v42 + (v42 + "")))
                                _ = v53
                            case -768:
                                var v54 int8 = (^(v1))
                                _ = v54
                            case -9223372036854775808:
                                v1 = (v1 + 42)
                            }
                            v55 := GenS4{F0: -1, F1: (-(fi43)), F2: v1}
                            for fi56 := 0; fi56 < 17; fi56++ {
                                fmt.Println(v42, fi56, v42)
                                _ = fi56
                            }
                            _ = v55
                        }
                        if (false && (true && false)) {
                            var v57 bool = ((v1 % (((v1 ^ -1)) | 1)) > v1)
                            v1 -= v1
                            var v58 int = fi43
                            _ = v57
                            _ = v58
                        } else {
                            fmt.Println(v1, v42, fi43)
                        }
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                v42 = (("\n\t" + v42) + (v42 + (((v42 + "hello") + (v42 + "a")) + "foo bar")))
                _ = fi43
            }
            v42 = (v42 + (v42 + v42))
            _ = v42
        }
    }
    var v59 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v59:", v59)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v59
}
