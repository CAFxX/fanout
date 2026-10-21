package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint; F1 uint16; F2 bool }
type GenS4 struct { F0 uint; F1 int32; F2 int64 }
func genFn0(p0 int8, p1 uint) uint64 {
    v4 := []int{-473}
    var v5 uint16 = 2
    for fi6 := 0; fi6 < 9; fi6++ {
        p1 = p1
        var v7 float64 = 1.5
        if true {
            switch 389 {
            case 0:
                var v8 uint = 4294967295
                _ = v8
            case 4294967295:
                fmt.Println(p0, v5, v5)
            case -2147483648:
                _ = len(v4)
            default:
                v7 = v7
            }
            if (!(!true)) {
                v7 = float64(1.5)
                if (470 != 2) {
                    for fi9 := 0; fi9 < 7; fi9++ {
                        var v10 uint = (^(p1))
                        var v11 uint = (p1 | 42)
                        p1 = v11
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    v4 = append(v4, (fi6 - 32767))
                } else {
                    p0 -= 100
                    fmt.Println(p0, fi6)
                }
            }
            _ = len(v4)
        }
        _ = fi6
        _ = v7
    }
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 928
}
func genFn1() bool {
    var v12 int8 = -128
    v13 := GenS1{F0: 42, F1: -32768}
    for fi14 := 0; fi14 < 17; fi14++ {
        v12 = v12
        var v15 float32 = 0.1
        for fi16 := 0; fi16 < 8; fi16++ {
            switch (("hello" + "") + "\n\t") {
            case "":
                v12 = 100
            case "\n\t":
                v17 := GenS1{F0: 120, F1: -1}
                _ = v17
            }
            fmt.Println(v15, v15, fi14)
            var v18 uint64 = 455
            _ = fi16
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    _ = v12
    _ = v13
    return (777 >= -1)
}

func main() {
    v1 := []string{("a" + ""), "a", "foo bar"}
    v2 := make(map[uint32]int)
    var v3 uint = 86
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if len(v1) > 0 {
            v1[0] = "hello"
        }
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        v3 += v3
        var v20 int = (-(fi19))
        fmt.Println(v20)
        _ = fi19
        _ = v20
    }
    var v21 bool = (!false)
    if v21 {
        if v21 {
            switch "\n\t" {
            case "a":
                var v22 int16 = -829
                _ = v22
            }
            v21 = (-482 >= 2147483647)
        } else {
            for fi23 := 0; fi23 < 11; fi23++ {
                for fi24 := 0; fi24 < 6; fi24++ {
                    if (!(0.0 != 0.0)) {
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(v21, v3)
                        case "hello":
                            v25 := GenS2{F0: -97, F1: 1}
                            _ = v25
                        default:
                            var v26 uint = v3
                            _ = v26
                        }
                        var v27 int16 = 1
                        var v28 uint16 = 2
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                        _ = v27
                        _ = v28
                    } else {
                        switch "hello" {
                        case "hello":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case "\n\t":
                            fmt.Println(fi23, fi24)
                        case "foo bar":
                            v29 := GenS1{F0: 42, F1: -1}
                            _ = v29
                        default:
                            switch (-(fi23)) {
                            case -1:
                                fmt.Println(v21, v21)
                            case 128:
                                v21 = v21
                            case 2147483647:
                                _ = len(v1)
                            default:
                                var v30 float32 = 0.1
                                _ = v30
                            }
                        }
                        v31 := GenS3{F0: (v3 & v3), F1: 14, F2: (1 == -16)}
                        delete(v2, 2147483647)
                        _ = v31
                    }
                    if v21 {
                        var v32 uint32 = 2147483648
                        switch "" {
                        case "":
                            switch "\n\t" {
                            case "\n\t":
                                switch "hello" {
                                case "":
                                    var v33 int8 = 1
                                    _ = v33
                                case "a":
                                    v3 += (v3 << ((v3) & 63))
                                case "hello":
                                    if ((true && v21) && false) {
                                        switch fi24 {
                                        case 1:
                                            var v34 uint16 = 42
                                            _ = v34
                                        default:
                                            var v35 int64 = 0
                                            _ = v35
                                        }
                                    }
                                default:
                                    var v36 string = "hello"
                                    _ = v36
                                }
                            case "foo bar":
                                v37 := GenS2{F0: 9223372036854775807, F1: -810}
                                _ = v37
                            case "a":
                                fmt.Println(v3, fi23)
                                fmt.Println(fi23, fi24, fi23)
                            }
                        default:
                            var v38 int16 = 1000
                            _ = v38
                        }
                        var v39 int64 = 0
                        fmt.Println(fi23)
                        if ((-79 < -1) || (fi23 < (fi24 + fi24))) {
                            v2[v32] = (-(fi24))
                            v40 := v2[(v32 % (((^(v32))) | 1))]
                            _ = v40
                        }
                        v3 = (^(v3))
                        _ = v32
                        _ = v39
                    } else {
                        fmt.Println(fi24)
                        var v41 int16 = 32767
                        _ = v41
                    }
                    v42 := GenS1{F0: 310, F1: -130}
                    _ = fi24
                    _ = v42
                }
                _ = fi23
            }
            if v21 {
                if v21 {
                    v3 = 255
                    _ = len(v2)
                } else {
                    v21 = v21
                }
                v1 = append(v1, "a")
                var v43 uint16 = 255
                _ = v43
            } else {
                v21 = v21
            }
        }
        switch 614 {
        case 1:
            var v44 uint = (^(v3))
            _ = v44
        case 42:
            v21 = (42 < (v3 / ((v3) | 1)))
        case -1:
            var v45 int8 = 1
            _ = v45
        }
    }
    v46 := GenS4{F0: 18446744073709551615, F1: 42, F2: -9223372036854775808}
    _ = len(v2)
    v21 = (((1e10 > 1e10) && false) || ((!(v21 || (((3.14159 + ((1e10 - 0.0) * float64((1.5 * (0.1 - (3.14159 / 0.1)))))) - 0.0) != (-2.25 + 0.1)))) && v21))
    for fi47 := 0; fi47 < 17; fi47++ {
        v2[958] = 725
        _ = fi47
    }
    var v48 float64 = (1.5 / 0.1)
    switch -976 {
    case 127:
        v2[733] = -435
    case 9223372036854775807:
        v48 -= 3.14159
    case 674:
        v21 = v21
    default:
        var v49 int16 = 1
        _ = v49
    }
    v2[965] = 4294967295
    delete(v2, 4294967295)
    switch "" {
    case "\n\t":
        fmt.Println(v3)
    case "":
        v50 := GenS2{F0: 0, F1: -1}
        _ = v50
    }
    v48 -= v48
    fmt.Println(v3, v48, v3)
    v21 = ((9223372036854775807 != -9223372036854775808) && (!false))
    if false {
        if v21 {
            var v51 int16 = 1
            v52 := v2[42]
            if (v21 && v21) {
                var v53 string = "foo bar"
                _ = v53
            } else {
                v51 -= (v51 + v51)
            }
            _ = v51
            _ = v52
        } else {
            var v54 int16 = 1
            var v55 uint16 = 264
            _ = v54
            _ = v55
        }
        fmt.Println(v48, v21)
        var v56 uint8 = 127
        _ = v56
    } else {
        var v57 string = "a"
        _ = v57
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v46
    _ = v48
}
