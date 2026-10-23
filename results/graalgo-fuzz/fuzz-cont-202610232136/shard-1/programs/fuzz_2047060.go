package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int32 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS4 struct { F0 int64; F1 uint16; F2 int16 }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64; F1 int32; F2 int16 }
type GenS8 struct { F0 int16; F1 int32; F2 int }
func genFn0(p0 uint, p1 int, p2 int16) int32 {
    var v4 int32 = 2147483647
    var v5 float32 = 0.0
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            fmt.Println(fi7, v5, p0)
            p1 += -1
            p1 += fi6
            _ = fi7
        }
        v8 := []int{}
        switch 0 {
        case 4294967295:
            if len(v8) > 0 {
                v8[0] = fi6
            }
        case 9223372036854775807:
            var v9 float64 = 0.0
            _ = v9
        }
        _ = fi6
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return v4
}
func genFn1(p0 int32, p1 int, p2 int32) string {
    p2 = p2
    for fi10 := 0; fi10 < 10; fi10++ {
        switch "\n\t" {
        case "a":
            var v11 string = "\n\t"
            _ = v11
        default:
            switch "" {
            case "\n\t":
                fmt.Println(p1)
            default:
                p1 = p1
            }
        }
        var v12 int64 = -1
        v13 := GenS1{F0: 0, F1: 65535, F2: p2}
        _ = fi10
        _ = v12
        _ = v13
    }
    fmt.Println(p0, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    return ("" + ("foo bar" + "hello"))
}

func main() {
    var v1 uint16 = 665
    var v2 uint16 = (v1 << ((584) & 63))
    v3 := make(map[int]string)
    v2 = v2
    var v14 int16 = 719
    fmt.Println(v2, v1, v14)
    if true {
        if true {
            if false {
                if ((v14 - v14) != 0) {
                    v14 -= 1000
                    v14 = (v14 - -32768)
                } else {
                    if (255 != 2) {
                        switch 1 {
                        case 42:
                            v15 := v3[-1]
                            _ = v15
                        }
                        v16 := GenS2{F0: 663, F1: v14}
                        var v17 float32 = 3.14159
                        _ = v16
                        _ = v17
                    }
                    switch 128 {
                    case -795:
                        v14 = (v14 ^ v14)
                    case -373:
                        for fi18 := 0; fi18 < 4; fi18++ {
                            var v19 uint64 = 0
                            if (true && ((-(v14)) == v14)) {
                                var v20 int64 = int64(v19)
                                if (254 > 54) {
                                    fmt.Println(v19, v19)
                                    var v21 int8 = 94
                                    v22 := GenS3{F0: 65535, F1: 4294967295, F2: 692}
                                    _ = v21
                                    _ = v22
                                } else {
                                    var v23 bool = true
                                    _ = v23
                                }
                                _ = v20
                            }
                            _ = fi18
                            _ = v19
                        }
                    case 127:
                        var v24 uint32 = 2147483648
                        _ = v24
                    default:
                        v2 = 65535
                    }
                }
                if true {
                    if ((!(630 > v14)) || false) {
                        v14 = 32767
                    }
                } else {
                    v2 = 1
                    fmt.Println(v14)
                    fmt.Println(v2, v14, v1)
                }
            }
        }
        v25 := []int{128, 127, -838}
        _ = v25
    } else {
        var v26 uint = 255
        fmt.Println(v1)
        if false {
            switch "foo bar" {
            case "hello":
                switch "" {
                case "foo bar":
                    var v27 int32 = -797
                    _ = v27
                case "":
                    v28 := GenS4{F0: -466, F1: v2, F2: 0}
                    _ = v28
                }
            case "":
                v29 := GenS4{F0: int64(v14), F1: v2, F2: v14}
                _ = v29
            case "foo bar":
                v30 := []int{32768, 2147483647, -714, 32768}
                _ = v30
            default:
                if false {
                    v26 = (v26 >> (((v26 >> ((v26) & 63))) & 63))
                } else {
                    fmt.Println(v1, v1, v1)
                }
            }
        } else {
            var v31 string = "foo bar"
            for fi32 := 0; fi32 < 4; fi32++ {
                var v33 int16 = v14
                v14 += v14
                _ = fi32
                _ = v33
            }
            switch 2147483647 {
            case -198:
                switch 408 {
                case 627:
                    switch v31 {
                    case "\n\t":
                        var v34 string = "a"
                        _ = v34
                    default:
                        var v35 int64 = 1
                        _ = v35
                    }
                case 2147483647:
                    if (42 <= 1) {
                        if (!(199 == -32768)) {
                            for fi36 := 0; fi36 < 10; fi36++ {
                                v14 -= (v14 >> (((v26 & 255)) & 63))
                                var v37 int16 = -32768
                                _ = fi36
                                _ = v37
                            }
                        }
                        var v38 float32 = -2.25
                        _ = v38
                    }
                case 32767:
                    var v39 uint64 = 466
                    _ = v39
                }
            case -1:
                if true {
                    v26 = 1
                    var v40 int16 = (^(v14))
                    _ = v40
                } else {
                    fmt.Println(v1, v26, v1)
                    v14 += v14
                    v41 := []int{9223372036854775807, 127, -676}
                    _ = v41
                }
            }
            _ = v31
        }
        _ = v26
    }
    v42 := v3[9223372036854775807]
    for fi43 := 0; fi43 < 5; fi43++ {
        if (((3.14159 + (1.5 + 0.0)) / 1.5) == 0.0) {
            v14 = (v14 - (^(v14)))
            v42 = "a"
            v44 := GenS2{F0: (fi43 ^ (fi43 | (fi43 ^ (fi43 << ((4294967295) & 63))))), F1: (v14 + v14)}
            _ = v44
        } else {
            var v45 int32 = -1
            v42 = (("hello" + (v42 + v42)) + "a")
            v46 := GenS4{F0: 42, F1: (^(v1)), F2: 236}
            _ = v45
            _ = v46
        }
        _ = fi43
    }
    if (42 == 9223372036854775807) {
        if (798 == 109) {
            fmt.Println(v14, v14, v42)
            fmt.Println(v14, v42, v2)
        } else {
            fmt.Println(v1, v2)
        }
    }
    var v47 bool = true
    var v48 float64 = (1e10 * (((((0.1 - ((0.0 - float64((1.5 + float32(((1.5 + ((((0.1 + 3.14159) - 0.0) / 1.5) + ((1.5 + float64(1e10)) + float64(0.1)))) - (1.5 - ((1.5 + (3.14159 * 0.0)) - ((3.14159 + 1e10) + -2.25)))))))) / 3.14159)) / -2.25) - (1e10 / 1.0)) * 1.5) + 3.14159))
    if (!v47) {
        v14 = -1
    }
    switch ("hello" + v42) {
    case "\n\t":
        for fi49 := 0; fi49 < 2; fi49++ {
            for fi50 := 0; fi50 < 15; fi50++ {
                var v51 int = fi50
                if ((-51 == 100) || (!(!(false || ((1e10 == (3.14159 + 3.14159)) || v47))))) {
                    if (true || v47) {
                        v14 = (v14 + v14)
                        switch -555 {
                        case -705:
                            v52 := []int{(fi49 >> ((42) & 63))}
                            _ = v52
                        case 42:
                            v1 = (v1 & (v2 - (v2 << ((255) & 63))))
                        case -1:
                            v53 := GenS4{F0: 259, F1: v1, F2: (-(v14))}
                            _ = v53
                        default:
                            fmt.Println(v47)
                        }
                    } else {
                        v14 -= (v14 + v14)
                        v54 := GenS4{F0: 9223372036854775807, F1: (v2 % (((v2 ^ (v1 * (v2 / ((929) | 1))))) | 1)), F2: (-(v14))}
                        _ = v54
                    }
                    var v55 bool = (255 >= 241)
                    _ = len(v3)
                    v3[fi50] = v42
                    v56 := GenS2{F0: 128, F1: v14}
                    _ = v55
                    _ = v56
                }
                _ = fi50
                _ = v51
            }
            var v57 uint32 = 79
            v58 := GenS5{F0: (v42 + "foo bar"), F1: (v14 << ((1) & 63))}
            _ = fi49
            _ = v57
            _ = v58
        }
    case "hello":
        switch 4294967295 {
        case -2147483648:
            v59 := GenS6{F0: -1}
            _ = v59
        case 2147483647:
            if true {
                v47 = true
                if (!(0 != genFn0(800, 4294967295, -836))) {
                    var v60 int = 0
                    if (((v47 && (1 == 2147483647)) || (65535 < 18446744073709551615)) || v47) {
                        v14 -= v14
                        var v61 float32 = -2.25
                        _ = v61
                    }
                    for fi62 := 0; fi62 < 19; fi62++ {
                        var v63 int8 = 0
                        v64 := GenS7{F0: -9223372036854775808, F1: 2147483647, F2: v14}
                        switch v60 {
                        case -1:
                            var v65 uint = 65535
                            _ = v65
                        case 9223372036854775807:
                            var v66 uint8 = 128
                            _ = v66
                        case 32768:
                            for fi67 := 0; fi67 < 2; fi67++ {
                                v68 := GenS6{F0: 1}
                                _ = fi67
                                _ = v68
                            }
                        }
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                    _ = v60
                }
                var v69 uint = 42
                _ = v69
            }
        default:
            delete(v3, 127)
        }
    }
    var v70 int16 = v14
    v71 := GenS8{F0: (v14 + (v14 - v70)), F1: 0, F2: -2147483648}
    var v72 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v70:", v70)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v42
    _ = v47
    _ = v48
    _ = v70
    _ = v71
    _ = v72
}
