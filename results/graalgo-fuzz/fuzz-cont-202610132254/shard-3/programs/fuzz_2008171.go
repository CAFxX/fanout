package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int64; F2 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 uint32; F2 int }
type GenS5 struct { F0 uint8; F1 int64; F2 bool }
type GenS6 struct { F0 string; F1 int; F2 int32 }
type GenS7 struct { F0 int8; F1 int; F2 uint }
func genFn0() bool {
    v4 := GenS1{F0: 128}
    switch -9223372036854775808 {
    case 4294967295:
        v5 := GenS2{F0: 58, F1: -1, F2: -1}
        _ = v5
    default:
        var v6 int = 19
        _ = v6
    }
    v7 := GenS2{F0: 42, F1: 866, F2: -9223372036854775808}
    _ = v4
    _ = v7
    return (0 == -886)
}
func genFn1() uint64 {
    switch ("hello" + "") {
    case "hello":
        if (18446744073709551615 != 781) {
            if (true && true) {
                switch ("hello" + (("a" + "hello") + ("foo bar" + ""))) {
                case "a":
                    switch 128 {
                    case 948:
                        v8 := GenS3{F0: 0}
                        _ = v8
                    default:
                        v9 := []int{676}
                        _ = v9
                    }
                case "":
                    for fi10 := 0; fi10 < 19; fi10++ {
                        var v11 uint32 = 909
                        v12 := []int{fi10}
                        fmt.Println(v11, v11)
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                default:
                    var v13 int32 = 0
                    _ = v13
                }
            }
            switch 127 {
            case 127:
                for fi14 := 0; fi14 < 18; fi14++ {
                    var v15 bool = false
                    v15 = (("a" > "hello") && v15)
                    _ = fi14
                    _ = v15
                }
            case 2147483647:
                var v16 int = -2147483648
                _ = v16
            default:
                var v17 string = ("\n\t" + "hello")
                _ = v17
            }
        }
    }
    v18 := GenS4{F0: 4294967295, F1: 2147483648, F2: 9223372036854775807}
    var v19 int64 = -857
    var v20 int = -1
    _ = v18
    _ = v19
    _ = v20
    return 802
}

func main() {
    var v1 int = -437
    var v2 float64 = (0.1 + ((1e10 / 1.0) - (1.5 - 0.1)))
    v3 := make(map[string]bool)
    for fi21 := 0; fi21 < 11; fi21++ {
        v22 := GenS4{F0: 2147483648, F1: 833, F2: (fi21 * (-(fi21)))}
        if false {
            var v23 string = "hello"
            _ = v23
        }
        for fi24 := 0; fi24 < 14; fi24++ {
            v1 = (fi24 + fi21)
            _ = fi24
        }
        _ = fi21
        _ = v22
    }
    fmt.Println(v1, v1)
    for fi25 := 0; fi25 < 9; fi25++ {
        switch (-(fi25)) {
        case 128:
            for fi26 := 0; fi26 < 4; fi26++ {
                v1 += 9223372036854775807
                var v27 int32 = -1
                v28 := GenS2{F0: 32767, F1: -9223372036854775808, F2: 1}
                _ = fi26
                _ = v27
                _ = v28
            }
        default:
            for fi29 := 0; fi29 < 6; fi29++ {
                var v30 uint16 = 873
                var v31 uint = 42
                var v32 int8 = -128
                v33 := GenS5{F0: 181, F1: 1, F2: false}
                v31 += (^(v31))
                switch "\n\t" {
                case "":
                    v1 += (v1 / ((v1) | 1))
                case "foo bar":
                    v1 = fi25
                }
                var v34 float32 = float32(float64(((1e10 / 1e10) + ((float32(1e10) / float32(v2)) / 3.14159))))
                var v35 float32 = (v34 / ((v34 / v34) / v34))
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
                _ = v33
                _ = v34
                _ = v35
            }
        }
        _ = fi25
    }
    for fi36 := 0; fi36 < 13; fi36++ {
        var v37 float64 = v2
        for fi38 := 0; fi38 < 3; fi38++ {
            switch "foo bar" {
            case "foo bar":
                v1 = (fi38 >> ((204) & 63))
            case "":
                v2 += v37
            case "\n\t":
                var v39 int64 = 9223372036854775807
                _ = v39
            }
            var v40 int32 = 1
            switch fi38 {
            case 42:
                v3[("a" + "foo bar")] = (!(false || false))
            default:
                var v41 float64 = (((float64(-2.25) / (v2 * float64(float32(0.0)))) - ((v37 * (float64((-2.25 + -2.25)) / float64((((0.1 / float32(v2)) * 3.14159) / (3.14159 / (1e10 + (0.1 - (0.0 / (3.14159 - 1e10))))))))) * 1.5)) / (float64(float32(v2)) / v37))
                _ = v41
            }
            switch "foo bar" {
            case "":
                switch -118 {
                case 127:
                    var v42 int32 = v40
                    v42 = 2147483647
                    _ = v42
                }
            case "a":
                switch (fi36 & (-(fi36))) {
                case 2147483647:
                    v43 := GenS6{F0: "\n\t", F1: 4294967295, F2: (v40 * (v40 & v40))}
                    _ = v43
                case 32767:
                    if false {
                        if (v40 > -1) {
                            var v44 float32 = -2.25
                            var v45 string = "foo bar"
                            _ = v44
                            _ = v45
                        } else {
                            var v46 string = "\n\t"
                            if false {
                                var v47 int = (fi36 >> ((1) & 63))
                                var v48 string = v46
                                _ = v47
                                _ = v48
                            } else {
                                v49 := v3[v46]
                                _ = v49
                            }
                            _ = v46
                        }
                        fmt.Println(v37)
                    }
                default:
                    var v50 string = "\n\t"
                    _ = v50
                }
                fmt.Println(fi38)
            case "hello":
                var v51 int8 = 42
                _ = v51
            default:
                if true {
                    v40 = v40
                }
            }
            var v52 uint64 = 2
            if false {
                v40 += 1
                switch "\n\t" {
                case "a":
                    var v53 bool = true
                    _ = v53
                case "\n\t":
                    v54 := GenS2{F0: 255, F1: 121, F2: -845}
                    _ = v54
                }
                var v55 uint = 65535
                v56 := GenS7{F0: 100, F1: (fi36 - (fi38 & 32768)), F2: v55}
                if false {
                    var v57 bool = false
                    _ = v57
                } else {
                    fmt.Println(v52, fi38)
                    fmt.Println(v40, v52)
                    if false {
                        var v58 uint8 = 127
                        var v59 int16 = 1000
                        var v60 float64 = v37
                        switch ("a" + "") {
                        case "foo bar":
                            var v61 float64 = -2.25
                            _ = v61
                        default:
                            var v62 uint16 = 0
                            _ = v62
                        }
                        v58 = v58
                        v63 := []int{(fi36 % (((fi38 ^ fi36)) | 1)), (fi38 & 9223372036854775807), (v1 * (v1 - (v1 | 2147483647)))}
                        _ = v58
                        _ = v59
                        _ = v60
                        _ = v63
                    }
                }
                _ = v55
                _ = v56
            } else {
                var v64 uint64 = 436
                _ = v64
            }
            _ = fi38
            _ = v40
            _ = v52
        }
        v1 -= 42
        _ = fi36
        _ = v37
    }
    var v65 string = ("foo bar" + ("foo bar" + ("a" + "foo bar")))
    var v66 bool = false
    var v67 uint64 = 2
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v65
    _ = v66
    _ = v67
}
