package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint32; F1 uint; F2 int }
type GenS4 struct { F0 bool; F1 string; F2 int16 }
type GenS5 struct { F0 uint32 }
func genFn0() int32 {
    var v7 uint64 = 4294967295
    v8 := GenS1{F0: -1, F1: -49}
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = -2147483648
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    var v5 uint64 = 42
    var v6 int32 = v2
    v6 = (v2 & v2)
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 string = ""
        var v11 int8 = -128
        _ = fi9
        _ = v10
        _ = v11
    }
    if true {
        v12 := []int{0, 9223372036854775807, 0, 128}
        switch -1 {
        case 1:
            v6 += v6
        case -9223372036854775808:
            _ = len(v4)
        case 271:
            switch "\n\t" {
            case "hello":
                for fi13 := 0; fi13 < 9; fi13++ {
                    var v14 float32 = 0.0
                    _ = fi13
                    _ = v14
                }
            case "foo bar":
                v5 += 42
            }
        default:
            fmt.Println(v6, v5, v5)
        }
        v2 += (v6 % (((v6 - (v6 / (((-(v2))) | 1)))) | 1))
        _ = v12
    }
    v15 := GenS2{F0: 255, F1: 0}
    for fi16 := 0; fi16 < 7; fi16++ {
        for fi17 := 0; fi17 < 3; fi17++ {
            if ((v5 % ((v5) | 1)) != v5) {
                switch fi16 {
                case -2147483648:
                    v6 = v6
                default:
                    v18 := GenS3{F0: 4294967295, F1: 65535, F2: fi17}
                    _ = v18
                }
                v19 := GenS2{F0: 0, F1: 65535}
                _ = v19
            } else {
                var v20 uint32 = 4294967295
                var v21 uint32 = 295
                _ = v20
                _ = v21
            }
            v2 = (v2 * (v6 + v6))
            fmt.Println(fi17, v6, v5)
            _ = fi17
        }
        var v22 int8 = 100
        if (fi16 == fi16) {
            v22 -= 66
            for fi23 := 0; fi23 < 16; fi23++ {
                if (!(!(1000 <= 42))) {
                    var v24 int = fi23
                    _ = v24
                } else {
                    v25 := GenS3{F0: 0, F1: 42, F2: -482}
                    if (0 <= 118) {
                        v2 -= -600
                        var v26 int32 = v2
                        v27 := GenS1{F0: v6, F1: (v22 | v22)}
                        _ = v26
                        _ = v27
                    } else {
                        switch "a" {
                        case "\n\t":
                            if false {
                                var v28 int = fi23
                                _ = v28
                            } else {
                                fmt.Println(v2, fi23, v5)
                            }
                        }
                        if true {
                            var v29 uint8 = 2
                            v2 = (v2 - (^(v2)))
                            switch ("a" + ("" + "\n\t")) {
                            case "foo bar":
                                v22 -= 91
                            case "\n\t":
                                var v30 int8 = -105
                                _ = v30
                            default:
                                v31 := GenS4{F0: false, F1: "\n\t", F2: 1}
                                _ = v31
                            }
                            _ = v29
                        }
                        if (fi23 >= fi16) {
                            var v32 uint16 = 255
                            v6 = v6
                            _ = v32
                        }
                        v33 := GenS3{F0: 42, F1: 369, F2: 9223372036854775807}
                        if true {
                            var v34 int16 = -1
                            _ = v34
                        } else {
                            v22 += v22
                            if (true || true) {
                                v35 := []int{2147483647, (fi16 >> ((2) & 63)), (fi16 - (fi16 + fi23)), (^(fi23))}
                                _ = v35
                            } else {
                                switch ("" + ("\n\t" + "foo bar")) {
                                case "hello":
                                    delete(v1, 32767)
                                }
                            }
                            fmt.Println(v2)
                        }
                        _ = v33
                    }
                    v2 = (^(v6))
                    var v36 uint16 = 1
                    _ = v25
                    _ = v36
                }
                v2 = v6
                _ = fi23
            }
        } else {
            var v37 float32 = (float32((3.14159 - 0.1)) * 1e10)
            fmt.Println(v5, v37, v5)
            _ = v37
        }
        _ = fi16
        _ = v22
    }
    var v38 float64 = (3.14159 - (0.0 / (((0.1 * -2.25) - (0.1 + (float64((0.0 + 3.14159)) - 0.0))) - (1e10 + -2.25))))
    v2 = (v6 >> ((4294967295) & 63))
    if true {
        v39 := GenS5{F0: 798}
        v6 = genFn0()
        var v40 int32 = 430
        _ = v39
        _ = v40
    } else {
        var v41 uint8 = uint8(v5)
        _ = v41
    }
    var v42 float64 = v38
    for fi43 := 0; fi43 < 12; fi43++ {
        if (1000 >= 336) {
            fmt.Println(fi43)
        } else {
            var v44 float64 = v42
            _ = v44
        }
        for fi45 := 0; fi45 < 11; fi45++ {
            var v46 int32 = (v2 % ((-1) | 1))
            fmt.Println(fi45, fi45, v46)
            v5 -= (^(v5))
            _ = fi45
            _ = v46
        }
        _ = fi43
    }
    v5 -= (v5 * (v5 ^ (v5 - 835)))
    if ((9223372036854775807 != 42) && (((-2.25 * ((float32((v42 + 0.1)) - float32(((((v42 / v42) / (((v42 + v42) + v38) * 1e10)) / (v38 / (0.1 / v38))) + 0.0))) - (((3.14159 + 1.5) * float32(v42)) * (3.14159 / ((1.5 - (1e10 + (0.1 - 3.14159))) * (3.14159 * (1.5 / 1e10))))))) / (1e10 + 3.14159)) == (1e10 - 1e10))) {
        var v47 uint = 163
        switch "" {
        case "a":
            v38 += 1e10
        default:
            var v48 uint32 = 38
            _ = v48
        }
        _ = v47
    }
    for fi49 := 0; fi49 < 18; fi49++ {
        var v50 uint8 = 80
        v4[("foo bar" + "")] = 65535
        _ = fi49
        _ = v50
    }
    for fi51 := 0; fi51 < 16; fi51++ {
        v52 := GenS5{F0: 4294967295}
        v53 := []int{fi51, (fi51 | int(v5)), (fi51 & (fi51 >> ((18446744073709551615) & 63)))}
        _ = fi51
        _ = v52
        _ = v53
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
