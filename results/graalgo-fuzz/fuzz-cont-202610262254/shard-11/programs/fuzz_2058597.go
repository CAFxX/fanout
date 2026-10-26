package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint16 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64 }
func genFn0(p0 int16, p1 int, p2 int8) int64 {
    v5 := GenS1{F0: -395, F1: p1, F2: 2}
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 string = ""
        p1 = fi6
        _ = fi6
        _ = v7
    }
    var v8 int64 = -9223372036854775808
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 int32 = -2147483648
        v11 := GenS1{F0: (v10 + (v10 + v10)), F1: (fi9 ^ p1), F2: 255}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int64 = (v8 & 0)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v5
    _ = v8
    return v12
}

func main() {
    var v1 float64 = float64(0.0)
    var v2 bool = true
    v3 := []uint32{1, 2147483648}
    v4 := []uint32{237, 1, 0}
    switch 1 {
    case 1:
        for fi13 := 0; fi13 < 2; fi13++ {
            v2 = v2
            switch -9223372036854775808 {
            case -430:
                var v14 float64 = v1
                _ = v14
            case -9223372036854775808:
                v1 = (v1 + (v1 / v1))
            default:
                _ = len(v4)
            }
            _ = fi13
        }
    case 42:
        switch 4294967295 {
        case 4294967295:
            var v15 int8 = 115
            _ = v15
        }
    case -888:
        v1 = float64(1e10)
    default:
        if (-903 != 9223372036854775807) {
            if len(v4) > 0 {
                v4[0] = 2147483647
            }
            for fi16 := 0; fi16 < 8; fi16++ {
                fmt.Println(v2, v2)
                _ = fi16
            }
            var v17 uint8 = 100
            _ = v17
        } else {
            _ = len(v4)
        }
    }
    v1 += 1.5
    v1 += float64((float32((v1 / float64(1.5))) - 0.1))
    v2 = (-1 == 1)
    var v18 int8 = 54
    for fi19 := 0; fi19 < 11; fi19++ {
        for fi20 := 0; fi20 < 2; fi20++ {
            fmt.Println(v18, v2)
            if len(v4) > 0 {
                v4[0] = 2147483647
            }
            v2 = v2
            _ = fi20
        }
        switch "foo bar" {
        case "":
            fmt.Println(v18)
        default:
            if (v18 == 12) {
                if (3.14159 >= (-2.25 - float32(v1))) {
                    if true {
                        switch "hello" {
                        case "a":
                            if (1 == 127) {
                                v2 = ((^(v18)) <= v18)
                                if true {
                                    var v21 uint16 = 391
                                    for fi22 := 0; fi22 < 13; fi22++ {
                                        if len(v4) > 0 {
                                            v4[0] = 2
                                        }
                                        _ = fi22
                                    }
                                    v1 = v1
                                    _ = v21
                                }
                                for fi23 := 0; fi23 < 13; fi23++ {
                                    v1 += v1
                                    var v24 int = -547
                                    fmt.Println(fi19, fi23, v1)
                                    _ = fi23
                                    _ = v24
                                }
                            }
                        default:
                            switch (("a" + (("" + ((("a" + ("a" + "foo bar")) + "\n\t") + ("" + "foo bar"))) + "hello")) + "a") {
                            case "\n\t":
                                for fi25 := 0; fi25 < 9; fi25++ {
                                    v18 += (-(v18))
                                    if true {
                                        var v26 int64 = 42
                                        fmt.Println(v1)
                                        _ = v26
                                    } else {
                                        var v27 float32 = 0.0
                                        _ = v27
                                    }
                                    var v28 uint16 = 1
                                    _ = fi25
                                    _ = v28
                                }
                            case "hello":
                                fmt.Println(v2, fi19, v18)
                            case "a":
                                switch (^(fi19)) {
                                case 32768:
                                    var v29 string = ("hello" + "hello")
                                    _ = v29
                                case -468:
                                    v30 := GenS1{F0: 1, F1: fi19, F2: 584}
                                    _ = v30
                                case 9223372036854775807:
                                    v2 = v2
                                default:
                                    if (!(false || ((255 > 65535) && true))) {
                                        v31 := GenS1{F0: -2147483648, F1: 2147483647, F2: 223}
                                        fmt.Println(v18)
                                        if (v2 && true) {
                                            switch (^(fi19)) {
                                            case 1:
                                                v32 := GenS2{F0: 127, F1: 2147483648}
                                                _ = v32
                                            case -53:
                                                var v33 uint32 = 0
                                                _ = v33
                                            case 32768:
                                                if len(v4) > 0 {
                                                    v4[0] = 2147483648
                                                }
                                            }
                                        } else {
                                            var v34 int8 = 25
                                            _ = v34
                                        }
                                        _ = v31
                                    }
                                }
                            default:
                                if ((65535 >= 65535) || (42 > 18446744073709551615)) {
                                    var v35 uint32 = 65535
                                    _ = v35
                                }
                            }
                        }
                        v1 = (v1 / v1)
                        if (!(9223372036854775807 == genFn0(1000, fi19, (v18 << ((0) & 63))))) {
                            if len(v4) > 0 {
                                v4[0] = 685
                            }
                            v36 := GenS1{F0: -437, F1: (fi19 % (((fi19 - fi19)) | 1)), F2: 0}
                            _ = len(v4)
                            _ = v36
                        }
                    } else {
                        var v37 float32 = float32(v1)
                        _ = v37
                    }
                    v38 := GenS3{F0: 72}
                    _ = v38
                } else {
                    var v39 float64 = (v1 + float64(1.5))
                    _ = v39
                }
                var v40 uint8 = 196
                _ = v40
            } else {
                var v41 uint32 = 4294967295
                _ = v41
            }
        }
        _ = fi19
    }
    if ((((3.14159 - 0.1) * 1.5) * 1.5) > 3.14159) {
        _ = len(v4)
    } else {
        v42 := GenS4{F0: genFn0(42, -1, v18)}
        _ = v42
    }
    fmt.Println(v1)
    var v43 uint16 = 2
    for fi44 := 0; fi44 < 20; fi44++ {
        var v45 uint32 = 191
        v2 = false
        for fi46 := 0; fi46 < 11; fi46++ {
            if len(v4) > 0 {
                v4[0] = 733
            }
            var v47 int16 = 270
            _ = fi46
            _ = v47
        }
        _ = fi44
        _ = v45
    }
    v43 -= v43
    var v48 float64 = (v1 * v1)
    var v49 int = 128
    v48 = v1
    if len(v4) > 0 {
        v4[0] = 1
    }
    var v50 string = ("hello" + "a")
    v48 = 1e10
    for fi51 := 0; fi51 < 2; fi51++ {
        if v2 {
            for fi52 := 0; fi52 < 4; fi52++ {
                if v2 {
                    fmt.Println(fi51)
                }
                if len(v3) > 0 {
                    v3[0] = 1
                }
                if len(v3) > 0 {
                    v3[0] = 868
                }
                v49 -= 4294967295
                _ = fi52
            }
            v53 := GenS5{F0: -1}
            _ = v53
        }
        for fi54 := 0; fi54 < 4; fi54++ {
            switch "" {
            case "foo bar":
                fmt.Println(v49, fi51)
            case "\n\t":
                fmt.Println(v18, v48)
            default:
                var v55 uint32 = 0
                fmt.Println(v55, v1)
                _ = v55
            }
            _ = fi54
        }
        switch (v49 / (((v49 * fi51)) | 1)) {
        case 127:
            if v2 {
                fmt.Println(v18, v43, v1)
            }
        case 9223372036854775807:
            for fi56 := 0; fi56 < 18; fi56++ {
                var v57 uint16 = (^(v43))
                var v58 float32 = 1e10
                _ = fi56
                _ = v57
                _ = v58
            }
        case 941:
            v59 := GenS6{F0: 9223372036854775807}
            _ = v59
        default:
            fmt.Println(v2)
        }
        _ = fi51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v48
    _ = v49
    _ = v50
}
