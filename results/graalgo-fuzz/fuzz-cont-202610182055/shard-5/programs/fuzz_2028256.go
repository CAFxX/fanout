package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int; F2 int8 }
type GenS4 struct { F0 int; F1 int32; F2 int16 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 float32 = (1e10 * 1.5)
    var v3 uint = 380
    var v4 int64 = v1
    var v5 float64 = 0.1
    switch (((("hello" + ("" + "hello")) + ("\n\t" + (("a" + ((("foo bar" + ("a" + "hello")) + "\n\t") + "")) + "\n\t"))) + "\n\t") + ("foo bar" + "\n\t")) {
    case "hello":
        if false {
            fmt.Println(v2, v1)
            v6 := []int{1, int(v4), 42}
            switch ("\n\t" + ("" + "hello")) {
            case "":
                switch "hello" {
                case "":
                    if false {
                        v7 := GenS1{F0: 32768}
                        _ = v7
                    } else {
                        var v8 int = 2147483647
                        _ = len(v6)
                        _ = v8
                    }
                case "\n\t":
                    if ((128 == 127) || true) {
                        if true {
                            for fi9 := 0; fi9 < 4; fi9++ {
                                var v10 float64 = float64((0.1 / 1.5))
                                if true {
                                    v11 := GenS1{F0: int(v3)}
                                    v6 = append(v6, (fi9 ^ (fi9 / ((fi9) | 1))))
                                    for fi12 := 0; fi12 < 14; fi12++ {
                                        if len(v6) > 0 {
                                            v6[0] = 0
                                        }
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = fi12
                                    }
                                    _ = v11
                                }
                                if len(v6) > 0 {
                                    v6[0] = (^(fi9))
                                }
                                _ = fi9
                                _ = v10
                            }
                        }
                    }
                case "foo bar":
                    var v13 float64 = v5
                    _ = v13
                }
            default:
                if len(v6) > 0 {
                    v6[0] = -998
                }
            }
            _ = v6
        }
    case "":
        for fi14 := 0; fi14 < 20; fi14++ {
            fmt.Println(v3, v3, fi14)
            v2 += (0.0 + 3.14159)
            v15 := []int{fi14}
            _ = fi14
            _ = v15
        }
    case "foo bar":
        switch "hello" {
        case "\n\t":
            v1 -= v4
        default:
            v16 := GenS2{F0: v3}
            _ = v16
        }
    }
    v17 := GenS2{F0: v3}
    var v18 uint32 = 2
    var v19 uint64 = 18446744073709551615
    var v20 float64 = float64(v2)
    if (!(!true)) {
        switch "foo bar" {
        case "hello":
            v1 = (v1 / (((v1 % ((v4) | 1))) | 1))
        case "":
            var v21 bool = ("foo bar" < ("\n\t" + "\n\t"))
            _ = v21
        }
        fmt.Println(v3, v5)
        var v22 float64 = (v20 / v5)
        _ = v22
    }
    var v23 int16 = 0
    v23 = v23
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 uint = uint(v18)
        _ = fi24
        _ = v25
    }
    if (!(true && false)) {
        v18 = v18
        v20 = v5
    } else {
        var v26 uint64 = uint64(v3)
        var v27 uint32 = (v18 | v18)
        _ = v26
        _ = v27
    }
    v18 = 666
    v28 := []int{42, 834}
    var v29 int8 = 87
    if (float32(v20) <= v2) {
        var v30 int16 = (v23 * (v23 >> ((18446744073709551615) & 63)))
        _ = v30
    } else {
        var v31 uint16 = 255
        if true {
            v32 := GenS1{F0: 127}
            v33 := GenS1{F0: 1}
            switch int(v29) {
            case 4294967295:
                v3 = (v3 / ((v3) | 1))
            case 2147483647:
                for fi34 := 0; fi34 < 10; fi34++ {
                    v29 = -128
                    v35 := GenS3{F0: uint8(v19), F1: fi34, F2: v29}
                    _ = fi34
                    _ = v35
                }
            default:
                for fi36 := 0; fi36 < 10; fi36++ {
                    if true {
                        _ = len(v28)
                        if true {
                            v5 += float64(v2)
                            fmt.Println(fi36)
                        }
                        switch (fi36 | fi36) {
                        case 937:
                            var v37 int32 = -1
                            _ = v37
                        case -1:
                            v19 += 42
                        case 2147483647:
                            if false {
                                if true {
                                    var v38 int8 = v29
                                    var v39 bool = (false && true)
                                    var v40 int16 = (v23 | 0)
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                } else {
                                    v4 = v4
                                    v1 -= (v4 % (((v1 / (((v1 >> (((v3 ^ (v3 / ((v3) | 1)))) & 63))) | 1))) | 1))
                                }
                            } else {
                                var v41 string = "a"
                                if len(v28) >= 2 {
                                    v28 = v28[1:len(v28)]
                                }
                                fmt.Println(fi36, v3, v20)
                                _ = v41
                            }
                        }
                    } else {
                        var v42 uint8 = 104
                        _ = v42
                    }
                    v43 := GenS4{F0: int(v23), F1: 0, F2: int16(fi36)}
                    var v44 int16 = (v23 % ((v23) | 1))
                    _ = fi36
                    _ = v43
                    _ = v44
                }
            }
            _ = v32
            _ = v33
        } else {
            var v45 uint8 = 128
            var v46 string = ("foo bar" + "hello")
            _ = v45
            _ = v46
        }
        if ((42 != -168) || true) {
            v47 := GenS4{F0: 205, F1: 0, F2: (-(v23))}
            v18 = (v18 % ((uint32(v3)) | 1))
            var v48 int32 = -1
            _ = v47
            _ = v48
        } else {
            if (!(v2 > 1.5)) {
                var v49 uint64 = v19
                v29 += v29
                _ = v49
            }
            var v50 uint32 = (v18 << (((^(v3))) & 63))
            for fi51 := 0; fi51 < 1; fi51++ {
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
                if (!false) {
                    var v52 uint = (v3 - (v3 + 255))
                    switch "hello" {
                    case "":
                        fmt.Println(fi51)
                    default:
                        for fi53 := 0; fi53 < 4; fi53++ {
                            v23 -= 377
                            _ = fi53
                        }
                    }
                    var v54 int32 = 939
                    _ = v52
                    _ = v54
                } else {
                    var v55 uint16 = uint16(v3)
                    var v56 float32 = v2
                    var v57 int16 = v23
                    _ = v55
                    _ = v56
                    _ = v57
                }
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
                _ = fi51
            }
            _ = v50
        }
        _ = v31
    }
    if (128 > -9223372036854775808) {
        v29 = v29
    } else {
        switch 0 {
        case 699:
            v58 := GenS2{F0: (v3 + (v3 + (v3 | v3)))}
            _ = v58
        default:
            v20 = v5
        }
    }
    v18 = v18
    v18 = v18
    for fi59 := 0; fi59 < 19; fi59++ {
        if true {
            fmt.Println(v5, v23)
        } else {
            v23 -= (^(v23))
            var v60 int16 = -1
            if (463 < v18) {
                for fi61 := 0; fi61 < 4; fi61++ {
                    var v62 int = 904
                    v18 = 2147483648
                    _ = fi61
                    _ = v62
                }
            }
            _ = v60
        }
        _ = fi59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
