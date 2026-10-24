package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int; F2 int64 }
type GenS3 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS4 struct { F0 bool; F1 int16; F2 uint8 }
type GenS5 struct { F0 string; F1 string; F2 string }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 uint16; F1 uint; F2 uint }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint64; F1 uint; F2 int32 }
func genFn0(p0 uint8) uint8 {
    v4 := GenS1{F0: -886}
    v5 := GenS2{F0: 604, F1: 128, F2: 481}
    _ = p0
    _ = v4
    _ = v5
    return (p0 / (((p0 & (p0 - (p0 | 0)))) | 1))
}

func main() {
    var v1 int16 = 42
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v6 float64 = 0.1
    if (0 <= -9223372036854775808) {
        switch -263 {
        case -917:
            v7 := GenS3{F0: v1, F1: 263, F2: 65535}
            _ = v7
        case 363:
            if ((true && (2 < 0)) && true) {
                v8 := []int{-9223372036854775808}
                switch -2147483648 {
                case 128:
                    for fi9 := 0; fi9 < 4; fi9++ {
                        fmt.Println(v1, fi9)
                        v10 := GenS2{F0: 1, F1: fi9, F2: -9223372036854775808}
                        _ = fi9
                        _ = v10
                    }
                case -2147483648:
                    if (100 < -84) {
                        v6 -= v6
                        if true {
                            fmt.Println(v1, v1)
                            for fi11 := 0; fi11 < 8; fi11++ {
                                for fi12 := 0; fi12 < 9; fi12++ {
                                    v6 += (v6 * v6)
                                    _ = fi12
                                }
                                _ = fi11
                            }
                        }
                    } else {
                        var v13 float32 = 0.1
                        v14 := GenS4{F0: false, F1: 0, F2: 180}
                        if false {
                            var v15 string = ""
                            _ = len(v8)
                            _ = v15
                        } else {
                            v13 = float32(v6)
                        }
                        _ = v13
                        _ = v14
                    }
                case 171:
                    var v16 int16 = (v1 | v1)
                    _ = v16
                default:
                    switch "a" {
                    case "hello":
                        for fi17 := 0; fi17 < 8; fi17++ {
                            v18 := GenS5{F0: "foo bar", F1: "a", F2: ""}
                            v1 += (^(v1))
                            _ = len(v2)
                            _ = fi17
                            _ = v18
                        }
                    case "":
                        for fi19 := 0; fi19 < 2; fi19++ {
                            fmt.Println(fi19, v6)
                            var v20 string = ("hello" + (("" + "hello") + "foo bar"))
                            fmt.Println(fi19, fi19, fi19)
                            _ = fi19
                            _ = v20
                        }
                    case "\n\t":
                        _ = len(v8)
                    }
                }
                _ = v8
            }
        }
    } else {
        var v21 int8 = -1
        var v22 int32 = 2147483647
        _ = v21
        _ = v22
    }
    v6 = 3.14159
    var v23 int32 = 2147483647
    if ((v6 > v6) || false) {
        if true {
            var v24 bool = false
            v1 -= v1
            v23 = v23
            _ = v24
        } else {
            v6 += v6
            v25 := []int{0}
            v3[2147483647] = "a"
            _ = v25
        }
        for fi26 := 0; fi26 < 15; fi26++ {
            for fi27 := 0; fi27 < 1; fi27++ {
                var v28 float64 = (0.1 - 3.14159)
                _ = fi27
                _ = v28
            }
            v29 := GenS1{F0: -9223372036854775808}
            switch (^(fi26)) {
            case 32767:
                fmt.Println(fi26, v23)
            case 1:
                v23 += 1
            default:
                var v30 string = "foo bar"
                _ = v30
            }
            _ = fi26
            _ = v29
        }
    } else {
        if ((v6 / -2.25) == v6) {
            for fi31 := 0; fi31 < 9; fi31++ {
                fmt.Println(v6, v6, v1)
                v32 := GenS1{F0: 1}
                var v33 float64 = (3.14159 + (0.1 * (3.14159 * (((v6 / v6) / v6) + v6))))
                _ = fi31
                _ = v32
                _ = v33
            }
            var v34 float64 = v6
            v1 = v1
            _ = v34
        } else {
            v6 = ((float64((0.0 + (3.14159 * 3.14159))) + (v6 - v6)) * v6)
            if true {
                fmt.Println(v6, v23, v23)
            }
        }
    }
    fmt.Println(v23, v23, v1)
    var v35 uint16 = 52
    v23 = v23
    for fi36 := 0; fi36 < 4; fi36++ {
        switch "foo bar" {
        case "foo bar":
            switch fi36 {
            case 2147483647:
                v37 := GenS4{F0: (!(v1 < (v1 - (v1 / ((-1) | 1))))), F1: int16(fi36), F2: 128}
                _ = v37
            case 32767:
                if true {
                    v1 = -726
                    for fi38 := 0; fi38 < 15; fi38++ {
                        switch ("\n\t" + "a") {
                        case "foo bar":
                            v39 := []int{-936, (fi36 / ((fi38) | 1))}
                            _ = v39
                        case "\n\t":
                            v40 := GenS6{F0: 1000, F1: "a"}
                            _ = v40
                        case "a":
                            v23 -= v23
                        }
                        _ = fi38
                    }
                    for fi41 := 0; fi41 < 19; fi41++ {
                        v35 -= (^(v35))
                        var v42 uint8 = 1
                        fmt.Println(v42, v6)
                        _ = fi41
                        _ = v42
                    }
                }
            case -859:
                for fi43 := 0; fi43 < 2; fi43++ {
                    fmt.Println(v35, v35, v6)
                    v44 := GenS7{F0: v35, F1: 991, F2: 624}
                    var v45 int16 = (v1 + (v1 >> ((1) & 63)))
                    v46 := GenS6{F0: (v1 / ((v1) | 1)), F1: (("a" + ("" + "a")) + "a")}
                    fmt.Println(v35, fi36)
                    if (((306 < 0) || false) && ("\n\t" <= "\n\t")) {
                        if false {
                            v47 := []int{(-(fi36)), 1, 783}
                            switch (fi36 / (((fi43 % (((fi43 & fi43)) | 1))) | 1)) {
                            case 127:
                                switch fi36 {
                                case 9223372036854775807:
                                    var v48 bool = true
                                    _ = v48
                                default:
                                    v3[780] = "hello"
                                }
                                fmt.Println(fi43, fi36)
                                v49 := GenS5{F0: ((("\n\t" + (("" + "a") + "foo bar")) + "a") + "a"), F1: "foo bar", F2: "hello"}
                                _ = v49
                            }
                            switch -249 {
                            case 32767:
                                v45 = -32768
                            case -413:
                                v6 = (3.14159 * float64(3.14159))
                                if len(v47) > 0 {
                                    v47[0] = fi43
                                }
                            default:
                                switch (^(fi36)) {
                                case 4294967295:
                                    v6 = 1e10
                                }
                            }
                            _ = v47
                        } else {
                            if true {
                                v50 := GenS8{F0: 2147483647}
                                _ = v50
                            }
                            v51 := GenS9{F0: uint64(v35), F1: 42, F2: v23}
                            _ = v51
                        }
                    }
                    fmt.Println(v23, v45, v35)
                    _ = fi43
                    _ = v44
                    _ = v45
                    _ = v46
                }
            default:
                fmt.Println(v6)
            }
        default:
            if (1e10 > -2.25) {
                v52 := GenS9{F0: 67, F1: uint(v35), F2: (v23 * (v23 & v23))}
                var v53 int64 = 42
                _ = v52
                _ = v53
            }
        }
        var v54 uint32 = 2147483647
        _ = fi36
        _ = v54
    }
    if true {
        switch "\n\t" {
        case "hello":
            for fi55 := 0; fi55 < 9; fi55++ {
                var v56 uint64 = 18446744073709551615
                var v57 uint = 0
                v23 += (v23 | 1)
                _ = fi55
                _ = v56
                _ = v57
            }
        }
    } else {
        var v58 bool = false
        var v59 uint = 2
        v6 = (v6 * (v6 / 1.5))
        _ = v58
        _ = v59
    }
    var v60 int32 = v23
    v61 := GenS6{F0: v1, F1: "foo bar"}
    v1 -= v1
    if (!true) {
        v6 += 0.0
    }
    if true {
        var v62 uint64 = 1
        _ = v62
    }
    for fi63 := 0; fi63 < 8; fi63++ {
        v35 += (v35 ^ 42)
        for fi64 := 0; fi64 < 2; fi64++ {
            var v65 uint8 = 252
            switch ("foo bar" + "") {
            case "hello":
                var v66 uint64 = 632
                _ = v66
            }
            fmt.Println(v6, v35, v6)
            _ = fi64
            _ = v65
        }
        _ = fi63
    }
    fmt.Println(v23, v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v35
    _ = v6
    _ = v60
    _ = v61
}
