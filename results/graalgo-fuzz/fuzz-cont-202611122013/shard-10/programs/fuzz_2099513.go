package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int16; F1 int32; F2 string }
type GenS3 struct { F0 int64; F1 uint32; F2 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 uint; F1 uint32 }
func genFn0(p0 uint8, p1 int, p2 string) uint8 {
    v4 := GenS1{F0: 29, F1: (false || (false || true))}
    p2 = ("\n\t" + "hello")
    v5 := GenS1{F0: 4294967295, F1: (((p1 | p1) < (p1 % ((p1) | 1))) || false)}
    switch "foo bar" {
    case "":
        switch "a" {
        case "foo bar":
            var v6 uint64 = 18446744073709551615
            _ = v6
        case "hello":
            p1 = p1
        case "a":
            v7 := GenS1{F0: 65535, F1: false}
            _ = v7
        }
    }
    var v8 uint8 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v8
    return 255
}

func main() {
    var v1 int16 = 1
    var v2 int16 = (v1 & 42)
    var v3 uint8 = 127
    for fi9 := 0; fi9 < 3; fi9++ {
        switch fi9 {
        case -2147483648:
            for fi10 := 0; fi10 < 11; fi10++ {
                switch (("hello" + "a") + "\n\t") {
                case "foo bar":
                    var v11 float32 = 1e10
                    if true {
                        switch (fi10 >> ((65535) & 63)) {
                        case -1:
                            var v12 int32 = -1
                            _ = v12
                        case 2147483647:
                            v13 := GenS1{F0: 0, F1: (v1 >= -32768)}
                            _ = v13
                        case 720:
                            switch "foo bar" {
                            case "hello":
                                switch fi9 {
                                case -478:
                                    var v14 uint = 484
                                    _ = v14
                                case 209:
                                    v15 := GenS1{F0: 18446744073709551615, F1: (!((42 == 42) && ((genFn0((v3 << ((1) & 63)), (fi10 * fi9), "a") > v3) || false)))}
                                    _ = v15
                                default:
                                    var v16 int16 = (v2 & -847)
                                    _ = v16
                                }
                            case "":
                                v17 := []int{}
                                fmt.Println(v11, v2, fi9)
                                _ = v17
                            case "a":
                                var v18 int = -3
                                _ = v18
                            default:
                                v19 := GenS2{F0: -597, F1: 1, F2: "\n\t"}
                                if false {
                                    switch "hello" {
                                    case "hello":
                                        v1 += (v1 % ((42) | 1))
                                    }
                                    v2 = v1
                                    fmt.Println(v2)
                                } else {
                                    var v20 float32 = ((((v11 * (v11 + v11)) * (v11 + v11)) + v11) * v11)
                                    v21 := GenS2{F0: v2, F1: -831, F2: "a"}
                                    switch "" {
                                    case "\n\t":
                                        var v22 int64 = -9223372036854775808
                                        _ = v22
                                    case "foo bar":
                                        var v23 float32 = ((v11 * (float32(1.5) * v11)) + (((v20 + (-2.25 - (v20 / v20))) / v20) - -2.25))
                                        fmt.Println(fi10)
                                        _ = v23
                                    default:
                                        var v24 float32 = v11
                                        _ = v24
                                    }
                                    _ = v20
                                    _ = v21
                                }
                                _ = v19
                            }
                        }
                    }
                    if (-1 > 20) {
                        fmt.Println(fi9, fi10, v3)
                        v1 += v1
                        var v25 int16 = (v1 % (((v2 + v1)) | 1))
                        _ = v25
                    }
                    _ = v11
                default:
                    v26 := GenS3{F0: 454, F1: 42, F2: (false || ((1e10 - (0.1 - 0.1)) > (3.14159 + float32(3.14159))))}
                    _ = v26
                }
                var v27 int8 = 127
                switch fi9 {
                case 32768:
                    switch "hello" {
                    case "hello":
                        var v28 uint64 = 18446744073709551615
                        _ = v28
                    case "foo bar":
                        var v29 int8 = 118
                        _ = v29
                    case "\n\t":
                        fmt.Println(fi9, v27, fi10)
                        v30 := GenS1{F0: 885, F1: ((1.5 - ((1e10 * 0.1) + 0.1)) > (float64(1e10) / 0.1))}
                        _ = v30
                    default:
                        fmt.Println(fi9, fi10)
                        if (!true) {
                            v31 := []int{140}
                            var v32 bool = (((v27 + (-(v27))) < (v27 >> ((363) & 63))) && (!(42 >= 42)))
                            _ = v31
                            _ = v32
                        }
                        v33 := []int{}
                        _ = v33
                    }
                case 245:
                    fmt.Println(v1)
                case -1:
                    v27 = v27
                default:
                    v34 := []int{fi10, (fi9 >> ((1) & 63)), fi10, (fi9 & int(v1))}
                    _ = v34
                }
                _ = fi10
                _ = v27
            }
        default:
            v35 := GenS1{F0: 583, F1: (!true)}
            _ = v35
        }
        var v36 string = (("foo bar" + ("hello" + "hello")) + "foo bar")
        v37 := GenS3{F0: 0, F1: 83, F2: ("a" < ((v36 + (v36 + (v36 + v36))) + v36))}
        _ = fi9
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 12; fi38++ {
        v3 = v3
        _ = fi38
    }
    if ((-(v1)) <= -1) {
        v3 += (^(v3))
    } else {
        fmt.Println(v2, v3)
        switch -1 {
        case 9223372036854775807:
            v39 := []int{}
            _ = v39
        case 361:
            v40 := GenS4{F0: 255}
            _ = v40
        case -577:
            fmt.Println(v3, v3, v3)
        default:
            v2 -= (v1 / ((-871) | 1))
        }
    }
    if false {
        var v41 int64 = -54
        _ = v41
    }
    switch -9223372036854775808 {
    case 4294967295:
        for fi42 := 0; fi42 < 16; fi42++ {
            v1 = (v1 ^ v1)
            v2 = (v1 + (v1 / ((int16(v3)) | 1)))
            var v43 int = 32767
            _ = fi42
            _ = v43
        }
    case 427:
        v2 = v1
    case 32767:
        v44 := GenS5{F0: "hello"}
        _ = v44
    default:
        v45 := []int{9223372036854775807, 0, 32768}
        _ = v45
    }
    switch "" {
    case "":
        for fi46 := 0; fi46 < 2; fi46++ {
            if false {
                fmt.Println(fi46, v2, v1)
                fmt.Println(v2, v1, v3)
            }
            var v47 int64 = 9223372036854775807
            _ = fi46
            _ = v47
        }
    case "\n\t":
        v3 = genFn0(v3, 4294967295, ("foo bar" + "foo bar"))
    default:
        v48 := GenS2{F0: v1, F1: -463, F2: (("\n\t" + "\n\t") + "")}
        _ = v48
    }
    if false {
        v49 := GenS6{F0: 0, F1: 1}
        _ = v49
    }
    v3 += (v3 & 2)
    v50 := []int{-86, 1, 553, 1}
    var v51 float32 = -2.25
    fmt.Println(v51)
    if ((-35 >= (v2 & v2)) && (true && (!true))) {
        _ = len(v50)
        v1 = v2
        var v52 int16 = (v2 << ((42) & 63))
        _ = v52
    }
    if false {
        v53 := GenS2{F0: v1, F1: 591, F2: "foo bar"}
        _ = v53
    } else {
        v2 += (v1 % (((v2 * -1)) | 1))
        var v54 int32 = -2147483648
        if true {
            switch 32767 {
            case 157:
                if false {
                    var v55 int32 = 797
                    _ = v55
                }
            case -9223372036854775808:
                if (962 >= 0) {
                    var v56 uint32 = 2147483647
                    _ = v56
                } else {
                    v57 := GenS1{F0: 18446744073709551615, F1: false}
                    _ = len(v50)
                    _ = v57
                }
            }
            v58 := GenS4{F0: 2}
            _ = v58
        } else {
            fmt.Println(v1, v2, v3)
            v59 := GenS6{F0: v54, F1: 169}
            _ = v59
        }
        _ = v54
    }
    for fi60 := 0; fi60 < 10; fi60++ {
        for fi61 := 0; fi61 < 12; fi61++ {
            if len(v50) >= 2 {
                v50 = v50[1:len(v50)]
            }
            _ = fi61
        }
        v62 := GenS7{F0: 0, F1: 355}
        _ = fi60
        _ = v62
    }
    if len(v50) >= 2 {
        v50 = v50[1:len(v50)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v50
    _ = v51
}
