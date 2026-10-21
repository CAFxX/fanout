package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int8; F2 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 int8; F1 uint32; F2 int8 }
type GenS7 struct { F0 uint16; F1 uint8 }
type GenS8 struct { F0 uint16; F1 int16 }
func genFn0(p0 bool) bool {
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: 1, F1: "", F2: 87}
        v9 := GenS1{F0: 2147483647, F1: ("a" + "foo bar"), F2: 42}
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 = false
    if (p0 || p0) {
        v10 := []int{0, 42}
        _ = v10
    }
    _ = p0
    return p0
}

func main() {
    var v1 uint = 65535
    var v2 uint32 = 193
    var v3 bool = (0 <= -483)
    var v4 int32 = -2147483648
    var v5 uint = 4294967295
    var v6 float32 = ((0.0 / -2.25) * (0.1 * 0.0))
    fmt.Println(v6, v1)
    v4 -= (v4 / (((v4 >> (((v5 >> (((v1 - 65535)) & 63))) & 63))) | 1))
    v11 := []int{}
    v5 += 255
    if v3 {
        if (!(v3 && (!v3))) {
            v6 = float32(0.0)
        }
        var v12 uint = (v5 * v1)
        _ = v12
    }
    if true {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        v13 := GenS2{F0: 127}
        var v14 int8 = 127
        _ = v13
        _ = v14
    }
    v5 = (^(v5))
    fmt.Println(v2)
    var v15 uint = uint(v2)
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 uint64 = 4294967295
        if false {
            var v18 uint8 = 42
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = v18
        } else {
            fmt.Println(fi16, v4)
            if ((fi16 - fi16) > fi16) {
                if v3 {
                    _ = len(v11)
                }
                if v3 {
                    _ = len(v11)
                } else {
                    var v19 float32 = -2.25
                    _ = v19
                }
                if (v3 || v3) {
                    v3 = true
                    if v3 {
                        if v3 {
                            var v20 float64 = 1.5
                            for fi21 := 0; fi21 < 10; fi21++ {
                                var v22 float64 = v20
                                _ = fi21
                                _ = v22
                            }
                            _ = v20
                        } else {
                            v23 := GenS3{F0: 100, F1: 42, F2: v15}
                            _ = v23
                        }
                    }
                }
            } else {
                if v3 {
                    v15 += v15
                } else {
                    var v24 uint16 = 2
                    _ = v24
                }
                if v3 {
                    v25 := GenS4{F0: 1000}
                    for fi26 := 0; fi26 < 12; fi26++ {
                        var v27 bool = ("a" < "\n\t")
                        _ = fi26
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 16; fi28++ {
                        var v29 uint32 = 0
                        _ = fi28
                        _ = v29
                    }
                    _ = v25
                }
                for fi30 := 0; fi30 < 4; fi30++ {
                    v5 = v1
                    switch "" {
                    case "foo bar":
                        var v31 uint64 = v17
                        if len(v11) > 0 {
                            v11[0] = fi16
                        }
                        v5 -= v15
                        _ = v31
                    case "hello":
                        fmt.Println(v15)
                    case "\n\t":
                        if false {
                            v1 = v1
                        } else {
                            var v32 bool = (!v3)
                            var v33 uint8 = 1
                            v17 -= v17
                            _ = v32
                            _ = v33
                        }
                    }
                    var v34 uint8 = 127
                    _ = fi30
                    _ = v34
                }
            }
            var v35 uint64 = 244
            _ = v35
        }
        switch "foo bar" {
        case "a":
            for fi36 := 0; fi36 < 5; fi36++ {
                switch "" {
                case "":
                    fmt.Println(v6)
                case "\n\t":
                    if len(v11) > 0 {
                        v11[0] = 0
                    }
                case "hello":
                    if v3 {
                        var v37 bool = ((32767 > 787) || v3)
                        _ = v37
                    }
                    v5 = (^(v1))
                    v38 := GenS5{F0: 127, F1: "\n\t"}
                    v39 := GenS3{F0: 100, F1: 1, F2: v1}
                    if v3 {
                        fmt.Println(v17, v2)
                        fmt.Println(v17)
                        if v3 {
                            v1 -= 4294967295
                        } else {
                            switch "\n\t" {
                            case "hello":
                                switch "\n\t" {
                                case "hello":
                                    var v40 int32 = (v4 % ((v4) | 1))
                                    _ = v40
                                case "a":
                                    if false {
                                        var v41 uint64 = 18446744073709551615
                                        _ = v41
                                    } else {
                                        if false {
                                            if (18446744073709551615 >= v17) {
                                                if v3 {
                                                    if v3 {
                                                        if (919 < fi36) {
                                                            v3 = (0.0 < 1.5)
                                                        }
                                                    }
                                                    if len(v11) >= 2 {
                                                        v11 = v11[1:len(v11)]
                                                    }
                                                    switch ("\n\t" + "hello") {
                                                    case "":
                                                        var v42 string = "hello"
                                                        _ = v42
                                                    case "\n\t":
                                                        v43 := GenS6{F0: -76, F1: (^(v2)), F2: 1}
                                                        _ = v43
                                                    default:
                                                        fmt.Println(v15, v3, v15)
                                                    }
                                                } else {
                                                    fmt.Println(v17)
                                                    v3 = (32767 == 362)
                                                    v15 -= (v15 | v1)
                                                }
                                                if (!(550 == 65535)) {
                                                    v17 = v17
                                                    if v3 {
                                                        switch 9223372036854775807 {
                                                        case -2147483648:
                                                            switch "foo bar" {
                                                            case "":
                                                                v5 -= (v5 / ((18446744073709551615) | 1))
                                                            default:
                                                                switch "\n\t" {
                                                                case "":
                                                                    v1 += (^(v15))
                                                                case "\n\t":
                                                                    var v44 uint = v1
                                                                    v45 := GenS6{F0: 1, F1: 65535, F2: -1}
                                                                    if len(v11) >= 2 {
                                                                        v11 = v11[1:len(v11)]
                                                                    }
                                                                    _ = v44
                                                                    _ = v45
                                                                case "a":
                                                                    var v46 string = "\n\t"
                                                                    _ = v46
                                                                }
                                                            }
                                                        case -587:
                                                            _ = len(v11)
                                                        }
                                                        switch "" {
                                                        case "":
                                                            if v3 {
                                                                var v47 uint8 = 128
                                                                _ = v47
                                                            }
                                                        case "\n\t":
                                                            if false {
                                                                var v48 uint = v15
                                                                _ = v48
                                                            }
                                                        case "a":
                                                            var v49 int = -2147483648
                                                            _ = v49
                                                        }
                                                    } else {
                                                        v3 = v3
                                                        var v50 int64 = 9223372036854775807
                                                        _ = v50
                                                    }
                                                } else {
                                                    var v51 float64 = (0.0 * ((float64((v6 / v6)) / float64(v6)) - (((1.5 - 3.14159) + 0.1) - ((1.5 + (-2.25 * 1.5)) * ((1.5 * ((1.5 * 3.14159) - 1e10)) / float64(1.5))))))
                                                    if len(v11) > 0 {
                                                        v11[0] = -1
                                                    }
                                                    _ = v51
                                                }
                                                v1 += (v5 / ((uint(v2)) | 1))
                                            } else {
                                                v3 = v3
                                                fmt.Println(v15, v1)
                                            }
                                        }
                                    }
                                default:
                                    var v52 uint32 = v2
                                    _ = v52
                                }
                            }
                            if (714 == 772) {
                                var v53 int8 = 127
                                if (v3 || v3) {
                                    if ((v15 | 4294967295) == 180) {
                                        switch ("" + "hello") {
                                        case "":
                                            if v3 {
                                                var v54 int32 = (v4 & v4)
                                                v55 := GenS7{F0: 1, F1: uint8(v15)}
                                                if false {
                                                    v1 = 957
                                                    var v56 float32 = -2.25
                                                    _ = v56
                                                }
                                                _ = v54
                                                _ = v55
                                            } else {
                                                var v57 uint16 = 0
                                                _ = v57
                                            }
                                        case "\n\t":
                                            v58 := GenS2{F0: fi36}
                                            _ = v58
                                        }
                                        v15 -= (v15 ^ 255)
                                    } else {
                                        if v3 {
                                            v17 += v17
                                        }
                                        fmt.Println(v5, v53)
                                    }
                                    fmt.Println(fi36)
                                }
                                var v59 uint8 = 128
                                _ = v53
                                _ = v59
                            } else {
                                fmt.Println(v3)
                                v60 := GenS4{F0: 682}
                                v5 = (^(v15))
                                _ = v60
                            }
                        }
                    }
                    var v61 uint16 = 42
                    v4 = v4
                    _ = v38
                    _ = v39
                    _ = v61
                default:
                    v15 = (^(v5))
                    var v62 float32 = float32(-2.25)
                    v11 = append(v11, fi16)
                    _ = v62
                }
                _ = fi36
            }
        }
        _ = fi16
        _ = v17
    }
    v15 = (v5 | 4294967295)
    v63 := GenS8{F0: 2, F1: 1}
    fmt.Println(v1, v3)
    v15 -= v15
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
}
