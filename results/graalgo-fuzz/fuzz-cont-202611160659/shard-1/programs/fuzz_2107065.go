package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int16 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 int16; F2 uint8 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 uint16 }
func genFn0() int16 {
    var v5 int8 = -1
    v6 := GenS1{F0: -2147483648, F1: 77, F2: -1}
    for fi7 := 0; fi7 < 14; fi7++ {
        if true {
            var v8 int = 4294967295
            _ = v8
        } else {
            if (!(18446744073709551615 <= 114)) {
                v9 := GenS1{F0: 210, F1: 255, F2: 42}
                _ = v9
            } else {
                v10 := GenS2{F0: -1, F1: 1}
                if false {
                    var v11 uint32 = 2
                    var v12 uint32 = (v11 / ((v11) | 1))
                    _ = v11
                    _ = v12
                } else {
                    var v13 int = fi7
                    _ = v13
                }
                _ = v10
            }
        }
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := []int{fi14}
            _ = fi14
            _ = v15
        }
        _ = fi7
    }
    v16 := []int{32768}
    if true {
        switch "a" {
        case "foo bar":
            if true {
                switch 4294967295 {
                case 127:
                    v5 += v5
                default:
                    v5 -= (v5 | (-(v5)))
                }
                _ = len(v16)
            } else {
                fmt.Println(v5)
                v16 = append(v16, 127)
            }
        case "\n\t":
            v5 = (v5 & 100)
        case "":
            v5 = v5
        default:
            v5 = v5
        }
        var v17 uint8 = 255
        fmt.Println(v5, v17)
        _ = v17
    } else {
        var v18 int32 = 42
        if len(v16) > 0 {
            v16[0] = 1
        }
        for fi19 := 0; fi19 < 14; fi19++ {
            v20 := GenS3{F0: "foo bar"}
            _ = fi19
            _ = v20
        }
        _ = v18
    }
    _ = v16
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint, p1 uint16) int {
    for fi21 := 0; fi21 < 17; fi21++ {
        var v22 int = (fi21 * (fi21 / ((fi21) | 1)))
        fmt.Println(fi21, p1, p0)
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        fmt.Println(p1, fi23)
        var v24 int8 = 42
        fmt.Println(v24)
        _ = fi23
        _ = v24
    }
    p1 = p1
    switch "\n\t" {
    case "\n\t":
        if (((423 >= (^(p0))) && (!false)) && false) {
            p1 = p1
            for fi25 := 0; fi25 < 20; fi25++ {
                p1 = (^(p1))
                for fi26 := 0; fi26 < 6; fi26++ {
                    var v27 uint64 = 2
                    p0 = 0
                    p0 -= p0
                    _ = fi26
                    _ = v27
                }
                _ = fi25
            }
            v28 := GenS2{F0: -500, F1: 1}
            _ = v28
        }
    case "foo bar":
        var v29 int = 9223372036854775807
        _ = v29
    case "hello":
        var v30 uint32 = 1
        _ = v30
    default:
        if (p0 <= (p0 * p0)) {
            v31 := GenS2{F0: 9223372036854775807, F1: uint8(p0)}
            switch 1 {
            case -111:
                var v32 bool = false
                _ = v32
            case 32767:
                v33 := []int{42, 32768, 0}
                _ = v33
            case 0:
                var v34 uint64 = 4294967295
                _ = v34
            }
            var v35 uint16 = (p1 << ((p0) & 63))
            _ = v31
            _ = v35
        }
    }
    _ = p0
    _ = p1
    return 32768
}

func main() {
    var v1 int8 = 100
    var v2 string = ""
    var v3 float32 = 1e10
    v4 := []string{"foo bar", (v2 + v2)}
    for fi36 := 0; fi36 < 5; fi36++ {
        v37 := GenS3{F0: ("\n\t" + v2)}
        _ = fi36
        _ = v37
    }
    if len(v4) > 0 {
        v4[0] = ((v2 + "") + v2)
    }
    switch v2 {
    case "":
        v38 := GenS3{F0: ("a" + "")}
        _ = v38
    case "a":
        var v39 int64 = 756
        _ = v39
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v3, v3, v2)
    v40 := GenS4{F0: v2, F1: 673, F2: 127}
    fmt.Println(v2, v3, v3)
    if false {
        if true {
            var v41 uint = 838
            var v42 uint64 = 4294967295
            v2 = "a"
            _ = v41
            _ = v42
        }
    } else {
        v3 += float32(0.0)
        v3 += (v3 * v3)
        v43 := GenS2{F0: 9223372036854775807, F1: 42}
        _ = v43
    }
    switch -1 {
    case 128:
        for fi44 := 0; fi44 < 18; fi44++ {
            fmt.Println(v3, fi44, v1)
            _ = fi44
        }
    case -1:
        for fi45 := 0; fi45 < 6; fi45++ {
            fmt.Println(v2)
            if true {
                var v46 int8 = 127
                for fi47 := 0; fi47 < 18; fi47++ {
                    v48 := GenS5{F0: 468, F1: -818}
                    _ = fi47
                    _ = v48
                }
                _ = v46
            } else {
                var v49 int = (fi45 / (((fi45 - fi45)) | 1))
                var v50 int = genFn1(65535, 32767)
                fmt.Println(v50)
                _ = v49
                _ = v50
            }
            _ = fi45
        }
    }
    if ((42 > 1) && false) {
        v4 = append(v4, v2)
        v1 -= -128
    } else {
        v3 -= (0.0 + 3.14159)
    }
    if false {
        if (0 < 100) {
            switch 32768 {
            case -773:
                for fi51 := 0; fi51 < 6; fi51++ {
                    switch "hello" {
                    case "hello":
                        fmt.Println(fi51)
                    case "a":
                        v52 := GenS6{F0: 42}
                        _ = v52
                    }
                    var v53 bool = (false && true)
                    _ = fi51
                    _ = v53
                }
            default:
                v1 = (v1 - (v1 - v1))
            }
            _ = len(v4)
            var v54 int = -2147483648
            _ = v54
        }
    }
    switch 4294967295 {
    case 0:
        for fi55 := 0; fi55 < 7; fi55++ {
            var v56 float64 = (-2.25 + (1e10 / (0.1 * 3.14159)))
            v57 := GenS5{F0: 719, F1: 42}
            _ = fi55
            _ = v56
            _ = v57
        }
    }
    switch v2 {
    case "\n\t":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "foo bar":
        v58 := GenS7{F0: 843, F1: 42}
        _ = v58
    case "hello":
        v3 += float32((((3.14159 * 3.14159) + ((float64(3.14159) - 0.1) * 1.5)) / 1.5))
    }
    var v59 int = 829
    switch -487 {
    case 32768:
        var v60 int32 = 878
        _ = v60
    case 164:
        v3 = v3
    default:
        v3 += float32(float64(v3))
    }
    var v61 float32 = float32((-2.25 + -2.25))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v59:", v59)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v59
    _ = v61
}
