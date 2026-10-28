package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint; F1 uint64; F2 int64 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 uint64; F1 uint64; F2 bool }
type GenS7 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS8 struct { F0 uint32; F1 int8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int }

func main() {
    var v1 string = ""
    var v2 string = "foo bar"
    v3 := make(map[string]bool)
    var v4 float64 = 3.14159
    var v5 int32 = 0
    var v6 bool = ((float32(v4) > 1e10) && (354 >= 313))
    switch 32767 {
    case 127:
        var v7 uint16 = 1
        _ = v7
    case 1:
        var v8 float32 = -2.25
        _ = v8
    case -1:
        var v9 uint16 = 42
        _ = v9
    default:
        v10 := GenS1{F0: -128}
        _ = v10
    }
    var v11 float64 = v4
    for fi12 := 0; fi12 < 1; fi12++ {
        v1 = ((v1 + ("" + v1)) + ((v2 + v1) + "\n\t"))
        v13 := GenS1{F0: -1}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        if v6 {
            v1 = (v2 + v2)
            v15 := []int{fi14, fi14}
            _ = v15
        } else {
            if (v6 || (4294967295 == 18446744073709551615)) {
                var v16 string = (v2 + "\n\t")
                var v17 int64 = 9223372036854775807
                _ = v16
                _ = v17
            } else {
                v18 := GenS2{F0: 2, F1: 0}
                _ = v18
            }
            v19 := []int{(fi14 + (fi14 >> ((347) & 63))), 32768, (fi14 % ((42) | 1))}
            if false {
                var v20 int64 = 745
                for fi21 := 0; fi21 < 10; fi21++ {
                    if v6 {
                        if false {
                            fmt.Println(v6, v4, v20)
                            if v6 {
                                switch 127 {
                                case -793:
                                    v22 := GenS3{F0: 42, F1: 42, F2: 794}
                                    _ = v22
                                }
                                var v23 uint16 = 1
                                fmt.Println(v23)
                                _ = v23
                            }
                        }
                        v20 -= v20
                        var v24 uint32 = 2147483648
                        _ = v24
                    } else {
                        switch 32768 {
                        case 127:
                            v20 -= v20
                            delete(v3, "a")
                            v25 := GenS4{F0: 255, F1: 400, F2: -297}
                            _ = v25
                        case 216:
                            var v26 uint = 2
                            _ = v26
                        case -2147483648:
                            if (!v6) {
                                v3[v2] = v6
                                v1 = ((v2 + (v1 + v1)) + (((v2 + v1) + (("foo bar" + "hello") + "\n\t")) + v1))
                            } else {
                                switch (v1 + v1) {
                                case "hello":
                                    fmt.Println(v6)
                                case "foo bar":
                                    switch fi14 {
                                    case 128:
                                        v27 := GenS5{F0: v5, F1: v5}
                                        _ = v27
                                    default:
                                        v28 := GenS5{F0: (v5 / ((v5) | 1)), F1: v5}
                                        _ = v28
                                    }
                                }
                                var v29 uint8 = 255
                                v19 = append(v19, -9223372036854775808)
                                _ = v29
                            }
                        default:
                            switch fi14 {
                            case 1:
                                if len(v19) > 0 {
                                    v19[0] = fi21
                                }
                            case 4294967295:
                                var v30 uint = 255
                                _ = v30
                            case 0:
                                var v31 float64 = v11
                                _ = v31
                            default:
                                var v32 bool = ((255 < 65535) && ((42 < -128) && true))
                                _ = v32
                            }
                        }
                        delete(v3, v1)
                        if v6 {
                            v33 := GenS4{F0: 255, F1: 756, F2: (v20 & 0)}
                            var v34 uint32 = 2147483648
                            if v6 {
                                v35 := GenS3{F0: 361, F1: 65535, F2: 18446744073709551615}
                                switch (fi14 % ((fi14) | 1)) {
                                case -2147483648:
                                    var v36 uint64 = 1
                                    v37 := GenS2{F0: 4294967295, F1: 92}
                                    _ = v36
                                    _ = v37
                                }
                                _ = v35
                            } else {
                                fmt.Println(v5, v2)
                                v11 += float64((float32(-2.25) + 1.5))
                                var v38 int8 = -1
                                _ = v38
                            }
                            delete(v3, v2)
                            _ = v33
                            _ = v34
                        } else {
                            var v39 string = "\n\t"
                            fmt.Println(v5)
                            fmt.Println(fi14, v39)
                            _ = v39
                        }
                    }
                    v40 := GenS3{F0: -444, F1: 18446744073709551615, F2: 948}
                    delete(v3, ("\n\t" + v2))
                    v4 -= v4
                    v41 := GenS3{F0: 32767, F1: 18446744073709551615, F2: 2}
                    _ = fi21
                    _ = v40
                    _ = v41
                }
                var v42 int32 = (v5 << ((255) & 63))
                _ = v20
                _ = v42
            } else {
                if (v6 || (v6 && true)) {
                    if (4294967295 < 1) {
                        v43 := GenS6{F0: 1, F1: 2, F2: (!(("foo bar" + v1) == "foo bar"))}
                        for fi44 := 0; fi44 < 9; fi44++ {
                            v45 := GenS1{F0: -1}
                            fmt.Println(fi44, v2, v4)
                            _ = fi44
                            _ = v45
                        }
                        var v46 uint8 = 128
                        _ = v43
                        _ = v46
                    }
                }
                for fi47 := 0; fi47 < 15; fi47++ {
                    var v48 string = v2
                    v6 = false
                    _ = fi47
                    _ = v48
                }
                v49 := GenS7{F0: 112, F1: 2147483647, F2: 2}
                _ = v49
            }
            _ = v19
        }
        for fi50 := 0; fi50 < 5; fi50++ {
            v51 := GenS2{F0: 0, F1: 1}
            v52 := GenS5{F0: 110, F1: (^(v5))}
            fmt.Println(v5)
            _ = fi50
            _ = v51
            _ = v52
        }
        _ = fi14
    }
    var v53 uint16 = 1
    v54 := []int{-543, -2147483648}
    if len(v54) > 0 {
        v54[0] = 128
    }
    if len(v54) >= 2 {
        v54 = v54[1:len(v54)]
    }
    delete(v3, (v2 + (v2 + v1)))
    for fi55 := 0; fi55 < 19; fi55++ {
        v56 := GenS8{F0: 2, F1: 127}
        v57 := GenS8{F0: 51, F1: 0}
        v58 := GenS9{F0: (v6 && (!v6))}
        _ = fi55
        _ = v56
        _ = v57
        _ = v58
    }
    v59 := GenS10{F0: ((((v11 + (v4 + float64(float32(v11)))) + v11) + (v4 + float64((((1.5 + -2.25) / 1.5) + 0.0)))) < float64(0.1))}
    fmt.Println(v53)
    v1 = v1
    delete(v3, "")
    v60 := v3[v2]
    if (v60 && v6) {
        if false {
            if false {
                if (!true) {
                    if (v11 <= v11) {
                        for fi61 := 0; fi61 < 5; fi61++ {
                            v53 -= (^(v53))
                            _ = len(v54)
                            _ = fi61
                        }
                    } else {
                        v4 -= 1.5
                        switch v1 {
                        case "":
                            for fi62 := 0; fi62 < 15; fi62++ {
                                fmt.Println(v6, v6, v4)
                                var v63 bool = (("hello" + v1) < "a")
                                _ = fi62
                                _ = v63
                            }
                        case "\n\t":
                            if v6 {
                                v53 = 0
                                v53 += v53
                            }
                        }
                    }
                    if v60 {
                        v64 := GenS3{F0: 0, F1: uint(v53), F2: 716}
                        if v60 {
                            v65 := GenS11{F0: -2147483648}
                            _ = v65
                        }
                        _ = v64
                    }
                }
                fmt.Println(v6)
                v54 = append(v54, -2147483648)
            }
        }
        for fi66 := 0; fi66 < 13; fi66++ {
            var v67 float32 = -2.25
            v68 := GenS2{F0: 18446744073709551615, F1: 93}
            var v69 int8 = -1
            _ = fi66
            _ = v67
            _ = v68
            _ = v69
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v59
    _ = v6
    _ = v60
}
