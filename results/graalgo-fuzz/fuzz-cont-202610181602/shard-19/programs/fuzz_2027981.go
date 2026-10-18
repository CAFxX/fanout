package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int }
func genFn0() uint8 {
    var v6 uint8 = 2
    v6 += v6
    switch -970 {
    case 356:
        fmt.Println(v6)
    case 9223372036854775807:
        var v7 uint8 = 206
        _ = v7
    }
    if true {
        switch ("foo bar" + "hello") {
        case "foo bar":
            if false {
                v6 = v6
                var v8 int = 9223372036854775807
                fmt.Println(v8)
                _ = v8
            } else {
                var v9 bool = (!(-2.25 < 3.14159))
                _ = v9
            }
        case "hello":
            var v10 int8 = 127
            _ = v10
        case "a":
            v6 += (v6 << ((460) & 63))
        default:
            v6 = v6
        }
    } else {
        if true {
            var v11 int16 = 442
            var v12 int32 = -2147483648
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 1; fi13++ {
                v6 = (v6 >> ((65535) & 63))
                for fi14 := 0; fi14 < 13; fi14++ {
                    v6 -= v6
                    _ = fi14
                }
                _ = fi13
            }
        }
        v6 = 0
        for fi15 := 0; fi15 < 17; fi15++ {
            switch ("foo bar" + (("a" + "hello") + "foo bar")) {
            case "a":
                if (!(119 < 100)) {
                    v16 := []int{1, fi15, (fi15 >> ((uint(v6)) & 63)), (fi15 << ((0) & 63))}
                    _ = v16
                } else {
                    v6 = (^(v6))
                    for fi17 := 0; fi17 < 16; fi17++ {
                        v18 := GenS1{F0: (-(fi15)), F1: -2147483648, F2: 32767}
                        fmt.Println(v6, fi15, fi17)
                        var v19 uint16 = 0
                        if false {
                            fmt.Println(fi15)
                        } else {
                            if (true && (true || false)) {
                                v20 := []int{}
                                if ((v6 | (v6 % (((v6 >> ((18446744073709551615) & 63))) | 1))) <= 0) {
                                    v6 += v6
                                } else {
                                    fmt.Println(v19)
                                    v20 = append(v20, fi17)
                                    v6 += v6
                                    fmt.Println(v19, fi17, fi15)
                                    v21 := GenS1{F0: (fi15 << ((2) & 63)), F1: -1, F2: -1}
                                    v19 = 878
                                    _ = v21
                                }
                                var v22 bool = (-638 <= fi15)
                                _ = v20
                                _ = v22
                            }
                            fmt.Println(fi17, fi15, v6)
                        }
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
            case "hello":
                if ((192 <= 18446744073709551615) || ((fi15 % (((fi15 * 32768)) | 1)) < -601)) {
                    v6 -= v6
                    v6 += (v6 << ((681) & 63))
                }
            default:
                for fi23 := 0; fi23 < 20; fi23++ {
                    var v24 int8 = 8
                    var v25 int16 = -625
                    var v26 int8 = v24
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
            }
            _ = fi15
        }
    }
    _ = v6
    return v6
}
func genFn1(p0 string, p1 uint32, p2 int8) uint32 {
    fmt.Println(p1, p2)
    v27 := GenS1{F0: 621, F1: 2147483647, F2: 32767}
    p1 -= 2
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return (p1 / (((p1 << ((4294967295) & 63))) | 1))
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"hello"}
    var v3 int8 = 42
    var v4 float32 = 3.14159
    var v5 float64 = (-2.25 - 0.1)
    v4 += (float32(1e10) / (0.0 - ((v4 + 3.14159) + 1.5)))
    v5 -= v5
    var v28 int16 = -32768
    v3 = v3
    var v29 int16 = v28
    if (("" + ("hello" + "foo bar")) != (("\n\t" + "foo bar") + ("hello" + "hello"))) {
        v30 := v1[2147483648]
        var v31 int8 = v3
        _ = v30
        _ = v31
    } else {
        for fi32 := 0; fi32 < 11; fi32++ {
            var v33 uint32 = 507
            if (false && (false && false)) {
                if false {
                    fmt.Println(v29, v3)
                    for fi34 := 0; fi34 < 20; fi34++ {
                        fmt.Println(v5, v33, v29)
                        var v35 float32 = (v4 + ((float32(v5) / v4) + 1.5))
                        switch fi34 {
                        case 2147483647:
                            if false {
                                v33 = 0
                            } else {
                                var v36 uint = 255
                                var v37 uint8 = 2
                                v36 -= v36
                                fmt.Println(v35, v28, v28)
                                _ = v36
                                _ = v37
                            }
                        case -326:
                            switch int(v28) {
                            case 127:
                                var v38 bool = (v33 >= v33)
                                _ = v38
                            case 128:
                                switch "foo bar" {
                                case "a":
                                    v39 := GenS1{F0: (fi34 - fi34), F1: 2147483647, F2: v29}
                                    _ = v39
                                case "\n\t":
                                    var v40 float64 = (v5 - v5)
                                    _ = v40
                                }
                            case 919:
                                switch 32767 {
                                case -110:
                                    if (!false) {
                                        if false {
                                            var v41 uint64 = 18446744073709551615
                                            _ = v41
                                        } else {
                                            if ((v29 | 282) != v29) {
                                                var v42 float32 = ((-2.25 - (v4 + float32((-2.25 + (((1e10 + ((-2.25 - ((v5 * 1.5) - ((0.0 / (float64(v35) - (v5 * -2.25))) * (((v5 * (float64(1.5) * v5)) - ((v5 + 3.14159) * 1e10)) / (((3.14159 + (v5 / (v5 - (v5 - (v5 * -2.25))))) / v5) - (v5 * (float64(float32(v5)) * 0.0))))))) * ((v5 - v5) * v5))) - v5) * (0.0 * float64(3.14159))))))) + v4)
                                                _ = v42
                                            } else {
                                                switch (fi34 >> ((255) & 63)) {
                                                case -1:
                                                    var v43 int32 = -2147483648
                                                    _ = v43
                                                default:
                                                    v44 := GenS1{F0: fi34, F1: -29, F2: (v28 % (((v29 - -32768)) | 1))}
                                                    _ = v44
                                                }
                                                v45 := GenS1{F0: (fi32 | (fi34 | (fi32 / (((fi34 % ((fi32) | 1))) | 1)))), F1: 42, F2: 1000}
                                                _ = v45
                                            }
                                            if ((^(v33)) != (^(v33))) {
                                                v3 += (^(v3))
                                            }
                                        }
                                    } else {
                                        var v46 float32 = 0.1
                                        _ = v46
                                    }
                                case 32768:
                                    v5 -= float64((float32(v5) / float32(((v5 - (v5 / (v5 + (float64(1.5) + (0.1 - (v5 - (0.0 * ((v5 / -2.25) - float64((0.0 + v4)))))))))) / (v5 - float64((3.14159 - v4)))))))
                                case 4294967295:
                                    v35 += v4
                                    if (false && false) {
                                        v4 += v35
                                        fmt.Println(v29, v35, fi32)
                                        fmt.Println(v33, fi32, v4)
                                    } else {
                                        v35 = v35
                                        fmt.Println(v28)
                                        var v47 int = (fi32 - (fi32 & fi32))
                                        _ = v47
                                    }
                                    var v48 int64 = 678
                                    _ = v48
                                }
                            default:
                                var v49 bool = (v3 == v3)
                                _ = v49
                            }
                        case 32768:
                            fmt.Println(fi32)
                        }
                        if (true || (uint8(v33) <= 0)) {
                            v33 -= (v33 >> ((0) & 63))
                            v1[318] = ""
                        } else {
                            fmt.Println(fi32, v5, fi32)
                            switch "a" {
                            case "":
                                switch "" {
                                case "hello":
                                    var v50 int16 = (v29 + 32767)
                                    _ = v50
                                case "a":
                                    if (-2.25 == float64(0.1)) {
                                        v29 = 1
                                    } else {
                                        if (!(false || (v5 < (3.14159 - float64(float32(1e10)))))) {
                                            v51 := GenS2{F0: 65535, F1: 301, F2: genFn1("\n\t", (v33 - genFn1("", (^(v33)), (v3 - 127))), v3)}
                                            fmt.Println(v4, fi34, v5)
                                            if (!(!(v5 <= 0.1))) {
                                                v3 = (v3 * v3)
                                                var v52 uint8 = 135
                                                v5 = float64(v35)
                                                if true {
                                                    var v53 uint8 = (v52 + v52)
                                                    fmt.Println(fi32)
                                                    _ = v53
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v33 += 2
                                                }
                                                _ = v52
                                            }
                                            _ = v51
                                        }
                                    }
                                default:
                                    var v54 int32 = -410
                                    _ = v54
                                }
                            case "\n\t":
                                if true {
                                    fmt.Println(v4, v4)
                                } else {
                                    if true {
                                        var v55 uint16 = 32767
                                        _ = v55
                                    }
                                    v35 += v4
                                }
                            default:
                                switch (fi34 * (fi34 % ((-30) | 1))) {
                                case 4294967295:
                                    v33 -= genFn1("", 1, 100)
                                }
                            }
                        }
                        _ = fi34
                        _ = v35
                    }
                    var v56 uint64 = 18446744073709551615
                    _ = v56
                } else {
                    for fi57 := 0; fi57 < 7; fi57++ {
                        v58 := GenS3{F0: int64(v33)}
                        v4 = 1.5
                        _ = fi57
                        _ = v58
                    }
                    switch fi32 {
                    case 127:
                        v59 := GenS4{F0: (v3 >> ((4294967295) & 63)), F1: (fi32 - fi32)}
                        _ = v59
                    }
                    fmt.Println(fi32)
                }
                var v60 int32 = 2147483647
                var v61 string = "foo bar"
                _ = v60
                _ = v61
            }
            switch "hello" {
            case "hello":
                v4 = (v4 - float32(v5))
            case "\n\t":
                for fi62 := 0; fi62 < 9; fi62++ {
                    v63 := GenS1{F0: 4294967295, F1: 515, F2: -32768}
                    var v64 uint8 = genFn0()
                    _ = fi62
                    _ = v63
                    _ = v64
                }
            default:
                for fi65 := 0; fi65 < 11; fi65++ {
                    switch (fi32 & fi65) {
                    case -376:
                        var v66 uint = 65535
                        _ = v66
                    }
                    _ = fi65
                }
            }
            _ = fi32
            _ = v33
        }
        for fi67 := 0; fi67 < 9; fi67++ {
            var v68 uint16 = 899
            if true {
                var v69 int16 = -7
                _ = v69
            } else {
                v4 = -2.25
            }
            fmt.Println(v28, v28)
            _ = fi67
            _ = v68
        }
    }
    if ((65535 <= 279) || (true && (2147483647 >= 42))) {
        var v70 int32 = 377
        v3 = (v3 - 100)
        var v71 uint64 = 4294967295
        _ = v70
        _ = v71
    }
    var v72 string = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v72
}
