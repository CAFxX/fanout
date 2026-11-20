package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 int32; F2 uint16 }
func genFn0() int16 {
    switch 688 {
    case 2147483647:
        var v6 uint8 = 127
        _ = v6
    default:
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int32 = -887
            v8 += v8
            _ = fi7
            _ = v8
        }
    }
    var v9 uint16 = 65535
    _ = v9
    return 42
}

func main() {
    var v1 string = "hello"
    var v2 float64 = 1.5
    var v3 uint8 = 0
    var v4 float64 = ((((v2 + ((3.14159 / 1.0) + ((v2 + (v2 - v2)) - ((float64(0.1) / (0.1 - 1e10)) * v2)))) * v2) + (1e10 - (float64(0.1) + 1.5))) - ((v2 * v2) / ((v2 - v2) / float64(3.14159))))
    var v5 uint = 351
    v10 := []int{127}
    if ("" == (("\n\t" + v1) + (((v1 + (v1 + "a")) + v1) + ((v1 + ("a" + v1)) + v1)))) {
        v11 := GenS1{F0: -115}
        v12 := GenS2{F0: 704, F1: 809, F2: 255}
        _ = v11
        _ = v12
    } else {
        if len(v10) > 0 {
            v10[0] = 703
        }
        if true {
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 float32 = -2.25
                for fi15 := 0; fi15 < 5; fi15++ {
                    switch (fi13 + (fi13 - -392)) {
                    case 32767:
                        var v16 string = v1
                        _ = v16
                    case -9223372036854775808:
                        var v17 string = v1
                        _ = v17
                    case 362:
                        var v18 bool = (((genFn0() == 32767) || true) || true)
                        _ = v18
                    default:
                        var v19 uint8 = (v3 * (^(v3)))
                        _ = v19
                    }
                    _ = fi15
                }
                _ = fi13
                _ = v14
            }
        } else {
            fmt.Println(v4)
            v1 = ""
            for fi20 := 0; fi20 < 10; fi20++ {
                v4 -= (v2 + v4)
                fmt.Println(v5, fi20)
                _ = fi20
            }
        }
    }
    var v21 uint16 = 65535
    var v22 bool = true
    switch (("foo bar" + (v1 + v1)) + v1) {
    case "hello":
        switch -2147483648 {
        case -178:
            if ((0.1 / 1.5) <= 0.1) {
                v21 = v21
                var v23 uint8 = (v3 * 127)
                _ = v23
            }
        default:
            if false {
                var v24 float64 = 1e10
                var v25 bool = false
                _ = v24
                _ = v25
            } else {
                var v26 uint64 = 2
                _ = v26
            }
        }
    case "foo bar":
        v4 -= ((v2 / (((float64(float32(v2)) - 3.14159) * v4) + 3.14159)) - 1e10)
    case "":
        var v27 bool = v22
        _ = v27
    default:
        fmt.Println(v1)
    }
    if v22 {
        v28 := GenS2{F0: 18446744073709551615, F1: -569, F2: v21}
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        if (!(4294967295 > 18446744073709551615)) {
            for fi29 := 0; fi29 < 6; fi29++ {
                switch (fi29 % ((127) | 1)) {
                case 1:
                    switch v1 {
                    case "foo bar":
                        var v30 uint32 = 2
                        _ = v30
                    }
                }
                v31 := GenS3{F0: (fi29 >> ((v5) & 63))}
                _ = fi29
                _ = v31
            }
            if v22 {
                var v32 uint8 = (^(v3))
                v2 += (((v2 - (0.1 / ((3.14159 / v4) * v2))) + (v4 + ((0.0 + 1e10) - 0.0))) + v4)
                v21 = v21
                _ = v32
            }
            fmt.Println(v4, v4, v1)
        } else {
            v4 = 0.1
        }
        _ = v28
    } else {
        fmt.Println(v5)
        if len(v10) > 0 {
            v10[0] = 9223372036854775807
        }
        v33 := GenS1{F0: -105}
        _ = v33
    }
    switch v1 {
    case "foo bar":
        if true {
            if (v22 && v22) {
                fmt.Println(v3, v1)
            } else {
                var v34 uint8 = v3
                switch -514 {
                case 814:
                    var v35 bool = (!((v4 / float64((0.1 / (-2.25 * 3.14159)))) > v2))
                    _ = v35
                case 0:
                    v3 -= (v3 * v3)
                }
                var v36 string = v1
                _ = v34
                _ = v36
            }
        } else {
            for fi37 := 0; fi37 < 3; fi37++ {
                fmt.Println(v5, v3)
                _ = fi37
            }
            fmt.Println(v1, v3, v2)
        }
    case "a":
        fmt.Println(v5, v1, v4)
    }
    switch (v1 + "hello") {
    case "a":
        var v38 int = 128
        _ = v38
    case "hello":
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    case "":
        fmt.Println(v4, v4)
    default:
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    }
    v39 := GenS2{F0: 404, F1: 487, F2: (v21 >> (((v5 - v5)) & 63))}
    switch ("a" + ("foo bar" + (v1 + ((v1 + v1) + v1)))) {
    case "a":
        v10 = append(v10, 2147483647)
    case "hello":
        v22 = ((v5 > 2) || v22)
    case "foo bar":
        if v22 {
            for fi40 := 0; fi40 < 1; fi40++ {
                v41 := GenS1{F0: -128}
                _ = fi40
                _ = v41
            }
            _ = len(v10)
            v4 = v4
        } else {
            for fi42 := 0; fi42 < 10; fi42++ {
                for fi43 := 0; fi43 < 13; fi43++ {
                    fmt.Println(v1, v2)
                    fmt.Println(fi43, v2, v21)
                    _ = fi43
                }
                for fi44 := 0; fi44 < 4; fi44++ {
                    var v45 float64 = (1e10 - float64(1e10))
                    _ = fi44
                    _ = v45
                }
                v46 := GenS4{F0: -23, F1: 95, F2: (v21 << ((v5) & 63))}
                _ = fi42
                _ = v46
            }
            var v47 uint = v5
            var v48 string = v1
            _ = v47
            _ = v48
        }
    default:
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    }
    for fi49 := 0; fi49 < 5; fi49++ {
        _ = len(v10)
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        for fi50 := 0; fi50 < 10; fi50++ {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            v51 := GenS4{F0: 0, F1: 642, F2: uint16(v3)}
            _ = fi50
            _ = v51
        }
        _ = fi49
    }
    var v52 float64 = ((float64(0.0) / 0.1) * 0.1)
    for fi53 := 0; fi53 < 10; fi53++ {
        for fi54 := 0; fi54 < 5; fi54++ {
            var v55 uint32 = 366
            _ = fi54
            _ = v55
        }
        _ = fi53
    }
    _ = len(v10)
    for fi56 := 0; fi56 < 19; fi56++ {
        switch (fi56 | fi56) {
        case 2147483647:
            fmt.Println(v2, v1, v21)
        case 42:
            fmt.Println(fi56)
        case -1:
            if true {
                if (!v22) {
                    fmt.Println(fi56, v52)
                    v1 = v1
                    v57 := GenS2{F0: 1, F1: 848, F2: uint16(v5)}
                    _ = v57
                } else {
                    var v58 float32 = 3.14159
                    _ = v58
                }
                var v59 int64 = int64(v3)
                _ = v59
            } else {
                if false {
                    fmt.Println(v21)
                } else {
                    for fi60 := 0; fi60 < 7; fi60++ {
                        if len(v10) > 0 {
                            v10[0] = fi56
                        }
                        var v61 float64 = (v52 + ((v52 * -2.25) * -2.25))
                        _ = fi60
                        _ = v61
                    }
                }
            }
        default:
            switch fi56 {
            case 32768:
                v4 += ((3.14159 / (v52 - v2)) + v4)
            default:
                v2 = (1.5 + v52)
            }
        }
        _ = fi56
    }
    for fi62 := 0; fi62 < 13; fi62++ {
        for fi63 := 0; fi63 < 9; fi63++ {
            var v64 float64 = (v4 / (((v2 * v52) / v4) + v2))
            v2 = 1e10
            v65 := GenS1{F0: 0}
            switch 128 {
            case 1:
                v5 -= v5
            case -599:
                if len(v10) > 0 {
                    v10[0] = fi62
                }
            }
            var v66 bool = ((v1 + v1) >= (v1 + (v1 + v1)))
            _ = fi63
            _ = v64
            _ = v65
            _ = v66
        }
        for fi67 := 0; fi67 < 10; fi67++ {
            v52 += float64(0.1)
            v68 := GenS1{F0: 1}
            v52 = 1.5
            _ = fi67
            _ = v68
        }
        var v69 bool = (!(0 == -128))
        _ = fi62
        _ = v69
    }
    var v70 uint32 = 587
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v52
    _ = v70
}
