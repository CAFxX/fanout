package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int; F1 uint32; F2 bool }
type GenS3 struct { F0 int8; F1 uint64; F2 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 int16; F2 string }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint; F1 uint32; F2 int64 }
type GenS9 struct { F0 uint32; F1 uint64; F2 string }
func genFn0(p0 int8) int16 {
    for fi5 := 0; fi5 < 2; fi5++ {
        fmt.Println(p0, fi5)
        _ = fi5
    }
    v6 := []int{2147483647, 125}
    p0 += (^(p0))
    v7 := GenS1{F0: -748, F1: 42, F2: 42}
    _ = p0
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint8, p1 uint8, p2 uint64) int64 {
    fmt.Println(p0, p2)
    v8 := GenS1{F0: -9223372036854775808, F1: 443, F2: 293}
    var v9 int16 = 1000
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 9223372036854775807
}

func main() {
    v1 := []string{"hello", "hello"}
    v2 := make(map[uint32]int)
    var v3 uint = 255
    var v4 bool = (2 < v3)
    var v10 int64 = 115
    fmt.Println(v4)
    v4 = true
    switch (("a" + "a") + ("foo bar" + "\n\t")) {
    case "":
        if false {
            for fi11 := 0; fi11 < 7; fi11++ {
                v3 = 0
                var v12 int32 = 0
                _ = fi11
                _ = v12
            }
        }
    case "a":
        var v13 uint = 42
        _ = v13
    case "hello":
        if (-67 == -1) {
            if (32767 <= 108) {
                v1 = append(v1, ("hello" + "foo bar"))
                if v4 {
                    if v4 {
                        var v14 uint8 = 220
                        _ = v14
                    } else {
                        v1 = append(v1, "a")
                        v3 -= 949
                        var v15 bool = v4
                        _ = v15
                    }
                    var v16 float32 = (-2.25 - ((3.14159 * float32(0.0)) * ((((1.5 - 1.5) + (3.14159 / 3.14159)) + float32((1.5 / float64(3.14159)))) / 3.14159)))
                    if v4 {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        v10 = genFn1(146, 195, 42)
                    }
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 18; fi17++ {
                        v18 := GenS2{F0: (^(fi17)), F1: 65535, F2: (0.0 > (1.5 * 3.14159))}
                        for fi19 := 0; fi19 < 16; fi19++ {
                            var v20 bool = ((!v4) || ((4294967295 <= 35) && (710 != (fi19 | int(v3)))))
                            if (2147483647 < 2147483648) {
                                switch "hello" {
                                case "\n\t":
                                    switch (-(fi17)) {
                                    case 4294967295:
                                        if v4 {
                                            fmt.Println(fi19)
                                            var v21 uint64 = 42
                                            _ = v21
                                        } else {
                                            var v22 int64 = (v10 ^ v10)
                                            v2[4294967295] = fi17
                                            _ = v22
                                        }
                                    case -223:
                                        var v23 float64 = 1e10
                                        if (((("hello" + "\n\t") + "foo bar") + "a") != "\n\t") {
                                            v24 := GenS3{F0: 1, F1: 18446744073709551615, F2: 114}
                                            v10 = (v10 | v10)
                                            _ = v24
                                        }
                                        if ((fi17 << ((593) & 63)) >= fi19) {
                                            var v25 int32 = 42
                                            fmt.Println(v20)
                                            _ = v25
                                        } else {
                                            var v26 uint = (v3 * (^(v3)))
                                            _ = v26
                                        }
                                        delete(v2, 0)
                                        _ = v23
                                    default:
                                        v10 = (v10 / (((v10 - genFn1(uint8(v3), 42, 42))) | 1))
                                    }
                                case "a":
                                    fmt.Println(v20, v10)
                                case "hello":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                default:
                                    v4 = (!v4)
                                }
                            } else {
                                v27 := GenS4{F0: (v10 >> (((v3 % ((v3) | 1))) & 63))}
                                _ = v27
                            }
                            v1 = append(v1, "\n\t")
                            v20 = (42 == (v3 % ((748) | 1)))
                            if ((!(!(int16(fi19) <= 0))) || v4) {
                                var v28 uint32 = 65535
                                _ = v28
                            } else {
                                var v29 int16 = 42
                                fmt.Println(v29, v3, v29)
                                _ = v29
                            }
                            _ = fi19
                            _ = v20
                        }
                        for fi30 := 0; fi30 < 11; fi30++ {
                            if (140 > 255) {
                                if (-2147483648 >= 829) {
                                    v10 = 761
                                    var v31 float32 = float32(((0.0 * 0.1) / 1.0))
                                    switch (fi17 - fi17) {
                                    case 2147483647:
                                        if (v3 != v3) {
                                            v3 += v3
                                            fmt.Println(fi17, v10)
                                            if v4 {
                                                var v32 uint32 = 74
                                                fmt.Println(fi17, fi17, fi30)
                                                _ = v32
                                            }
                                        }
                                    case 42:
                                        v31 = float32((3.14159 + ((((3.14159 * -2.25) + float64(0.1)) - 0.0) - -2.25)))
                                    }
                                    _ = v31
                                }
                                var v33 uint = v3
                                _ = v33
                            } else {
                                v34 := GenS5{F0: 127, F1: 32767, F2: "hello"}
                                v4 = v4
                                v3 += 0
                                _ = v34
                            }
                            _ = fi30
                        }
                        _ = fi17
                        _ = v18
                    }
                    var v35 int32 = -2147483648
                    _ = v35
                }
            } else {
                v4 = v4
            }
            switch -2147483648 {
            case -111:
                if ("\n\t" != "hello") {
                    switch -675 {
                    case -9223372036854775808:
                        for fi36 := 0; fi36 < 15; fi36++ {
                            var v37 int = fi36
                            fmt.Println(v4, v37, fi36)
                            if v4 {
                                if (true || v4) {
                                    v38 := GenS5{F0: 86, F1: 42, F2: ("\n\t" + (("foo bar" + "") + ("hello" + (("foo bar" + "\n\t") + "a"))))}
                                    _ = v38
                                } else {
                                    v10 = v10
                                    v3 = (v3 & (v3 / (((v3 % ((v3) | 1))) | 1)))
                                }
                                if (-76 < 125) {
                                    var v39 int32 = 830
                                    switch "" {
                                    case "hello":
                                        var v40 int8 = 1
                                        _ = v40
                                    case "\n\t":
                                        _ = len(v1)
                                    default:
                                        v4 = true
                                    }
                                    if (!v4) {
                                        fmt.Println(v37, v10, v10)
                                        var v41 int64 = v10
                                        var v42 uint16 = 255
                                        _ = v41
                                        _ = v42
                                    } else {
                                        v43 := GenS2{F0: 32768, F1: 0, F2: v4}
                                        var v44 float32 = 1.5
                                        if v4 {
                                            if (v3 >= v3) {
                                                switch v37 {
                                                case 1:
                                                    if (9223372036854775807 != fi36) {
                                                        var v45 int16 = 32767
                                                        for fi46 := 0; fi46 < 20; fi46++ {
                                                            if (v4 && v4) {
                                                                fmt.Println(v44, v39)
                                                            }
                                                            _ = fi46
                                                        }
                                                        _ = v45
                                                    } else {
                                                        if (!v4) {
                                                            fmt.Println(v4)
                                                            v47 := GenS6{F0: v3, F1: 278}
                                                            _ = v47
                                                        }
                                                    }
                                                case 127:
                                                    fmt.Println(v39)
                                                case -610:
                                                    v48 := GenS7{F0: 4294967295}
                                                    _ = v48
                                                default:
                                                    if false {
                                                        if (!(v4 && v4)) {
                                                            if v4 {
                                                                v44 = v44
                                                            }
                                                        } else {
                                                            v44 = v44
                                                            v49 := GenS3{F0: 1, F1: 2, F2: 127}
                                                            switch "" {
                                                            case "":
                                                                v50 := GenS5{F0: -98, F1: genFn0(127), F2: ("hello" + (("a" + "foo bar") + "hello"))}
                                                                _ = v50
                                                            case "a":
                                                                fmt.Println(v37)
                                                            case "hello":
                                                                fmt.Println(v3, v4, v39)
                                                            }
                                                            _ = v49
                                                        }
                                                    }
                                                }
                                            } else {
                                                switch (("" + ("" + "")) + "a") {
                                                case "a":
                                                    v37 += (fi36 & (v37 & fi36))
                                                }
                                                for fi51 := 0; fi51 < 19; fi51++ {
                                                    v1 = append(v1, ("hello" + ("\n\t" + ("foo bar" + "hello"))))
                                                    _ = fi51
                                                }
                                                v39 += (v39 << ((18446744073709551615) & 63))
                                            }
                                        }
                                        _ = v43
                                        _ = v44
                                    }
                                    _ = v39
                                }
                            }
                            _ = fi36
                            _ = v37
                        }
                    default:
                        v10 = -1
                    }
                }
            default:
                switch "foo bar" {
                case "a":
                    var v52 uint64 = 2
                    _ = v52
                case "":
                    _ = len(v2)
                default:
                    var v53 uint = v3
                    _ = v53
                }
            }
            var v54 int = 1
            _ = v54
        } else {
            var v55 string = (("foo bar" + "a") + "\n\t")
            _ = v55
        }
    }
    switch -2147483648 {
    case 32768:
        v56 := v2[1]
        _ = v56
    case -9223372036854775808:
        var v57 uint16 = 957
        _ = v57
    case 843:
        var v58 string = "hello"
        _ = v58
    }
    if (("" <= "") && (!v4)) {
        _ = len(v1)
    } else {
        fmt.Println(v3, v4)
        for fi59 := 0; fi59 < 9; fi59++ {
            switch 32768 {
            case -784:
                v60 := GenS8{F0: (v3 * 43), F1: 2147483647, F2: v10}
                _ = v60
            case -687:
                for fi61 := 0; fi61 < 18; fi61++ {
                    var v62 int16 = -1
                    _ = fi61
                    _ = v62
                }
            }
            _ = fi59
        }
    }
    var v63 float32 = 3.14159
    v4 = v4
    _ = len(v1)
    fmt.Println(v63, v63)
    v1 = append(v1, "\n\t")
    if (182 != 0) {
        v64 := GenS8{F0: (v3 - 4294967295), F1: 2147483647, F2: v10}
        v65 := GenS9{F0: 0, F1: 18446744073709551615, F2: "hello"}
        _ = v64
        _ = v65
    }
    if v4 {
        var v66 int = -1
        _ = v66
    }
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v63
}
