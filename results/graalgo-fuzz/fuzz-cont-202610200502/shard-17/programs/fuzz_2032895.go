package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint16; F1 int8; F2 int8 }
type GenS5 struct { F0 int8; F1 int64 }
type GenS6 struct { F0 int; F1 uint32; F2 string }
type GenS7 struct { F0 uint8; F1 int64; F2 int16 }
type GenS8 struct { F0 int32; F1 string }
func genFn0(p0 uint, p1 int8) int {
    v4 := GenS1{F0: 191}
    v5 := GenS2{F0: 127, F1: (p0 + p0), F2: 4294967295}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 777
}
func genFn1(p0 string, p1 uint64) uint32 {
    v6 := GenS3{F0: (p1 / (((p1 % (((p1 << ((386) & 63))) | 1))) | 1)), F1: 2, F2: -1}
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            p1 = p1
            p0 = (p0 + "a")
            _ = fi8
        }
        v9 := []int{fi7, 9223372036854775807}
        _ = fi7
        _ = v9
    }
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = v6
    return 266
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint8 = 37
    v3 := []uint32{2, 2147483647, 4294967295, 2147483648}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2, v2)
    switch ("hello" + "") {
    case "":
        v1 -= 42
    case "hello":
        v10 := GenS4{F0: 65535, F1: 42, F2: -104}
        _ = v10
    case "a":
        v2 = v2
    default:
        for fi11 := 0; fi11 < 7; fi11++ {
            fmt.Println(v2)
            _ = fi11
        }
    }
    v2 -= v2
    v2 += v2
    fmt.Println(v1, v1)
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 float64 = 0.1
            fmt.Println(fi13, v1)
            v1 = (v1 % (((-(v1))) | 1))
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    var v15 float32 = -2.25
    for fi16 := 0; fi16 < 7; fi16++ {
        v17 := GenS2{F0: 2147483647, F1: 18446744073709551615, F2: 1}
        fmt.Println(v1)
        var v18 uint32 = genFn1(("" + "a"), 4294967295)
        _ = fi16
        _ = v17
        _ = v18
    }
    v15 += v15
    if len(v3) > 0 {
        v3[0] = 1
    }
    if (!(false || false)) {
        var v19 float32 = v15
        var v20 uint32 = 0
        switch ("" + (("\n\t" + "foo bar") + "a")) {
        case "\n\t":
            _ = len(v3)
        case "a":
            var v21 bool = (false && ((0 != 625) && true))
            _ = v21
        }
        _ = v19
        _ = v20
    } else {
        fmt.Println(v2, v2)
        fmt.Println(v15, v15, v2)
    }
    if (false || false) {
        var v22 bool = true
        v22 = (!(true && (2 != 132)))
        var v23 uint16 = 32767
        _ = v22
        _ = v23
    } else {
        var v24 uint8 = (^(v2))
        switch "hello" {
        case "":
            switch ("a" + "\n\t") {
            case "hello":
                if true {
                    for fi25 := 0; fi25 < 8; fi25++ {
                        for fi26 := 0; fi26 < 8; fi26++ {
                            switch (-(fi25)) {
                            case 713:
                                v2 = (^(v2))
                                if (-177 > (fi25 + (fi25 / ((fi25) | 1)))) {
                                    var v27 int = fi26
                                    switch ("foo bar" + "a") {
                                    case "foo bar":
                                        v28 := GenS5{F0: 127, F1: v1}
                                        _ = v28
                                    case "":
                                        var v29 float32 = float32(((-2.25 - float64(v15)) / -2.25))
                                        _ = v29
                                    case "\n\t":
                                        v30 := GenS6{F0: (fi25 & v27), F1: genFn1(("" + "\n\t"), 42), F2: "hello"}
                                        _ = v30
                                    }
                                    _ = v27
                                } else {
                                    switch (-(fi25)) {
                                    case 128:
                                        fmt.Println(v1, fi25, v1)
                                    case -157:
                                        var v31 uint64 = 0
                                        _ = v31
                                    case -1:
                                        var v32 float64 = float64((v15 - ((v15 + (v15 - v15)) * v15)))
                                        _ = v32
                                    default:
                                        var v33 int32 = -2147483648
                                        _ = v33
                                    }
                                    if ((-124 != -458) && false) {
                                        fmt.Println(v24)
                                        var v34 int32 = 42
                                        _ = v34
                                    } else {
                                        var v35 string = "a"
                                        switch v35 {
                                        case "":
                                            v15 = 0.1
                                        }
                                        var v36 uint8 = 175
                                        _ = v35
                                        _ = v36
                                    }
                                }
                            case 0:
                                v1 -= v1
                            case 127:
                                v2 = (v2 ^ 1)
                                fmt.Println(fi26, fi26, v15)
                            }
                            if (-1 == -766) {
                                if false {
                                    v24 = v2
                                    if ("\n\t" != "") {
                                        fmt.Println(v15)
                                    }
                                } else {
                                    switch (-(fi25)) {
                                    case 9223372036854775807:
                                        fmt.Println(v24)
                                        if (!((0.0 < 0.0) || (false && false))) {
                                            var v37 float64 = 1.5
                                            v15 = 1.5
                                            _ = v37
                                        } else {
                                            _ = len(v3)
                                            _ = len(v3)
                                            var v38 int64 = -1
                                            _ = v38
                                        }
                                        fmt.Println(v24, v15, v1)
                                    }
                                    var v39 int64 = v1
                                    var v40 int16 = 32767
                                    _ = v39
                                    _ = v40
                                }
                            } else {
                                var v41 float64 = (1.5 + 1e10)
                                switch "" {
                                case "a":
                                    if len(v3) > 0 {
                                        v3[0] = genFn1("hello", 1)
                                    }
                                case "foo bar":
                                    v24 = v24
                                case "\n\t":
                                    var v42 float32 = 0.0
                                    _ = v42
                                default:
                                    var v43 int64 = v1
                                    v2 = v2
                                    var v44 uint8 = (v24 ^ (v24 | 255))
                                    _ = v43
                                    _ = v44
                                }
                                v45 := GenS6{F0: (fi25 % (((fi26 >> ((2) & 63))) | 1)), F1: 1, F2: ""}
                                _ = v41
                                _ = v45
                            }
                            _ = fi26
                        }
                        _ = fi25
                    }
                    v1 += (v1 - (v1 | -1))
                    for fi46 := 0; fi46 < 20; fi46++ {
                        v1 += (v1 + (v1 + v1))
                        _ = fi46
                    }
                }
            default:
                v1 -= (v1 & (v1 * (v1 ^ (v1 << ((4294967295) & 63)))))
            }
        case "hello":
            for fi47 := 0; fi47 < 5; fi47++ {
                if len(v3) > 0 {
                    v3[0] = 803
                }
                if true {
                    v1 += 42
                    v2 = v24
                    v48 := GenS7{F0: v24, F1: (v1 * (v1 >> ((18446744073709551615) & 63))), F2: -434}
                    _ = v48
                } else {
                    v24 -= v24
                    if (42 != 0) {
                        if false {
                            switch (^(fi47)) {
                            case 4294967295:
                                var v49 int32 = 1
                                _ = v49
                            case -2147483648:
                                if false {
                                    if false {
                                        var v50 uint16 = 2
                                        if (2 >= 2147483648) {
                                            switch fi47 {
                                            case 127:
                                                v51 := GenS8{F0: 59, F1: "a"}
                                                _ = v51
                                            case 2147483647:
                                                if true {
                                                    for fi52 := 0; fi52 < 4; fi52++ {
                                                        v50 = 255
                                                        _ = fi52
                                                    }
                                                    v53 := GenS8{F0: -1, F1: ""}
                                                    _ = v53
                                                }
                                            case -9223372036854775808:
                                                fmt.Println(v15)
                                            default:
                                                fmt.Println(v15)
                                            }
                                            fmt.Println(v50)
                                        } else {
                                            v24 -= v2
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v54 := GenS6{F0: (^(fi47)), F1: 0, F2: ("hello" + "hello")}
                                            _ = v54
                                        }
                                        for fi55 := 0; fi55 < 3; fi55++ {
                                            var v56 uint32 = 2
                                            fmt.Println(v24, v24)
                                            var v57 uint16 = (v50 ^ (^(v50)))
                                            v56 = v56
                                            _ = fi55
                                            _ = v56
                                            _ = v57
                                        }
                                        _ = v50
                                    }
                                    v24 = v2
                                    v15 += (v15 + (v15 / float32(3.14159)))
                                } else {
                                    v15 = (3.14159 * (v15 + (-2.25 * v15)))
                                }
                            case -885:
                                v1 = (^(v1))
                            default:
                                v24 = v24
                            }
                        }
                    } else {
                        v1 = (v1 - v1)
                        var v58 uint16 = 88
                        _ = v58
                    }
                    if (true && (-2147483648 <= 375)) {
                        for fi59 := 0; fi59 < 19; fi59++ {
                            v1 = 1
                            var v60 string = ((("hello" + "foo bar") + ("a" + ("hello" + "foo bar"))) + ("" + "hello"))
                            _ = fi59
                            _ = v60
                        }
                        v24 = v24
                        var v61 int = (fi47 + genFn0(1, 36))
                        _ = v61
                    } else {
                        fmt.Println(v24, v24)
                        var v62 uint16 = 812
                        _ = v62
                    }
                }
                switch ("\n\t" + ("foo bar" + "")) {
                case "foo bar":
                    for fi63 := 0; fi63 < 6; fi63++ {
                        _ = len(v3)
                        v64 := GenS7{F0: (v2 + (v2 >> ((267) & 63))), F1: v1, F2: 32767}
                        _ = fi63
                        _ = v64
                    }
                }
                _ = fi47
            }
        }
        _ = v24
    }
    _ = len(v3)
    fmt.Println(v2, v15, v2)
    var v65 uint32 = 32
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v65
}
