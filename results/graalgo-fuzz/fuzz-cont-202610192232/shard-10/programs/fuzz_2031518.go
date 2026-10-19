package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int32; F1 int; F2 int16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint }

func main() {
    var v1 int16 = -994
    var v2 string = ("" + "\n\t")
    var v3 bool = true
    var v4 int8 = 0
    var v5 float32 = 0.1
    if v3 {
        var v6 uint32 = 42
        _ = v6
    }
    if v3 {
        v7 := []int{-2147483648, -506}
        v8 := GenS1{F0: ("\n\t" + v2), F1: -369, F2: 904}
        var v9 uint64 = 18446744073709551615
        _ = v7
        _ = v8
        _ = v9
    } else {
        if (0 == 42) {
            var v10 uint8 = 128
            v1 = (v1 & v1)
            v11 := []int{32768, 42, -212}
            _ = v10
            _ = v11
        } else {
            if (((4294967295 == 255) && v3) || (255 < 255)) {
                v1 = (v1 >> ((255) & 63))
            }
        }
        v1 = v1
        var v12 string = v2
        _ = v12
    }
    v13 := GenS2{F0: 0}
    if (!(742 >= 18446744073709551615)) {
        switch "a" {
        case "\n\t":
            if (false && (v4 != 127)) {
                if v3 {
                    var v14 int32 = -82
                    v15 := []int{1, 128, int(v14), -255}
                    fmt.Println(v2, v5)
                    _ = v14
                    _ = v15
                }
                for fi16 := 0; fi16 < 18; fi16++ {
                    if (fi16 != 1) {
                        switch v2 {
                        case "":
                            for fi17 := 0; fi17 < 4; fi17++ {
                                v3 = v3
                                fmt.Println(v4, v1, v2)
                                var v18 uint32 = 2147483648
                                _ = fi17
                                _ = v18
                            }
                        }
                        for fi19 := 0; fi19 < 1; fi19++ {
                            var v20 uint32 = 42
                            _ = fi19
                            _ = v20
                        }
                        fmt.Println(v1, v5, v5)
                    } else {
                        v21 := GenS2{F0: 0}
                        v22 := []int{(fi16 << ((193) & 63))}
                        fmt.Println(v3, fi16)
                        _ = v21
                        _ = v22
                    }
                    v2 = v2
                    _ = fi16
                }
                fmt.Println(v5, v2)
            }
        case "":
            for fi23 := 0; fi23 < 7; fi23++ {
                v24 := []int{fi23}
                _ = fi23
                _ = v24
            }
        case "foo bar":
            var v25 int64 = 1
            _ = v25
        }
    }
    fmt.Println(v5, v4)
    v2 = ("foo bar" + (v2 + (v2 + v2)))
    if v3 {
        v26 := GenS1{F0: "", F1: 0, F2: -1}
        _ = v26
    }
    if (!true) {
        v27 := GenS1{F0: v2, F1: 9223372036854775807, F2: 2147483647}
        fmt.Println(v5)
        v3 = ((!v3) || v3)
        _ = v27
    } else {
        v28 := []int{}
        fmt.Println(v1, v2)
        v28 = append(v28, 32767)
        _ = v28
    }
    v29 := []int{42, 128, 128}
    if (-2147483648 > -1) {
        var v30 uint = 1
        switch 128 {
        case 128:
            for fi31 := 0; fi31 < 3; fi31++ {
                v30 += (v30 % ((v30) | 1))
                _ = fi31
            }
        }
        switch v2 {
        case "foo bar":
            v2 = "a"
        case "a":
            v1 += (v1 - -1)
        default:
            for fi32 := 0; fi32 < 10; fi32++ {
                var v33 int64 = 1
                v29 = append(v29, 42)
                _ = fi32
                _ = v33
            }
        }
        _ = v30
    } else {
        if true {
            v5 = (v5 / ((float32((1e10 / float64(((v5 / v5) / 1.0)))) - float32(3.14159)) - float32((0.0 * 1e10))))
            v29 = append(v29, 32767)
            var v34 float64 = 0.0
            _ = v34
        } else {
            v35 := GenS3{F0: 32768, F1: 42}
            v36 := GenS4{F0: -2147483648, F1: 793, F2: v1}
            v37 := GenS1{F0: (v2 + (v2 + ("foo bar" + (((v2 + v2) + ((v2 + v2) + ("a" + v2))) + ("\n\t" + v2))))), F1: 0, F2: 1}
            _ = v35
            _ = v36
            _ = v37
        }
        for fi38 := 0; fi38 < 20; fi38++ {
            if (!v3) {
                switch v2 {
                case "foo bar":
                    if (!((false && ((!v3) || (42 >= -9223372036854775808))) || (18446744073709551615 != 1))) {
                        if ((v1 ^ v1) >= v1) {
                            if false {
                                switch (fi38 + fi38) {
                                case 128:
                                    v1 -= (v1 % ((v1) | 1))
                                case 845:
                                    v39 := GenS3{F0: (^(fi38)), F1: -9223372036854775808}
                                    _ = v39
                                case 1:
                                    var v40 uint8 = 255
                                    _ = v40
                                }
                            }
                            for fi41 := 0; fi41 < 10; fi41++ {
                                if (66 != 42) {
                                    fmt.Println(v2, fi41, v3)
                                    v29 = append(v29, (fi38 << ((84) & 63)))
                                } else {
                                    var v42 string = v2
                                    v29 = append(v29, (fi41 % ((fi41) | 1)))
                                    v5 = (1.5 - (v5 + float32(float64(((1.5 / ((v5 + v5) / (v5 - 0.0))) + ((float32(1.5) * v5) + float32((float64(v5) + 1e10))))))))
                                    v5 -= (0.1 + -2.25)
                                    if (v3 || v3) {
                                        v5 = (((v5 - (v5 + 0.1)) / float32(-2.25)) - v5)
                                        v43 := GenS5{F0: 399}
                                        _ = v43
                                    }
                                    _ = v42
                                }
                                v44 := GenS3{F0: (^(fi38)), F1: 644}
                                _ = fi41
                                _ = v44
                            }
                        } else {
                            var v45 uint16 = 10
                            _ = v45
                        }
                    } else {
                        for fi46 := 0; fi46 < 11; fi46++ {
                            if len(v29) >= 2 {
                                v29 = v29[1:len(v29)]
                            }
                            v4 -= 0
                            fmt.Println(v1, v3)
                            v47 := GenS3{F0: 127, F1: -9223372036854775808}
                            _ = fi46
                            _ = v47
                        }
                        var v48 float64 = 1e10
                        _ = v48
                    }
                case "":
                    v4 += v4
                }
                for fi49 := 0; fi49 < 14; fi49++ {
                    v5 -= v5
                    _ = len(v29)
                    v50 := GenS6{F0: 4294967295}
                    _ = fi49
                    _ = v50
                }
                switch 386 {
                case -1:
                    if len(v29) >= 2 {
                        v29 = v29[1:len(v29)]
                    }
                }
            }
            v4 = 48
            v3 = true
            _ = fi38
        }
    }
    for fi51 := 0; fi51 < 17; fi51++ {
        for fi52 := 0; fi52 < 9; fi52++ {
            var v53 float32 = (0.1 + v5)
            _ = fi52
            _ = v53
        }
        v54 := GenS4{F0: 42, F1: fi51, F2: v1}
        _ = fi51
        _ = v54
    }
    v55 := GenS6{F0: 274}
    var v56 bool = (0.0 < -2.25)
    v57 := GenS3{F0: 0, F1: 9223372036854775807}
    for fi58 := 0; fi58 < 20; fi58++ {
        switch (fi58 | (fi58 << ((274) & 63))) {
        case 42:
            v59 := GenS6{F0: 2}
            _ = v59
        case 32767:
            for fi60 := 0; fi60 < 15; fi60++ {
                v4 -= v4
                _ = fi60
            }
        default:
            if len(v29) > 0 {
                v29[0] = (fi58 / ((-1) | 1))
            }
        }
        v56 = false
        for fi61 := 0; fi61 < 13; fi61++ {
            v2 = "foo bar"
            _ = fi61
        }
        _ = fi58
    }
    v62 := GenS3{F0: 150, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v62
}
