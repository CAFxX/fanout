package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int64; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 int; F1 bool; F2 int16 }
type GenS6 struct { F0 int; F1 int16 }
type GenS7 struct { F0 int64; F1 uint64 }
type GenS8 struct { F0 uint; F1 bool }
type GenS9 struct { F0 uint; F1 int8; F2 int64 }
func genFn0() int {
    var v4 float32 = -2.25
    var v5 bool = ((628 <= 42) || (!false))
    v6 := []int{9223372036854775807}
    v7 := GenS1{F0: "hello"}
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return -2147483648
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint = 162
    var v3 string = "foo bar"
    v8 := GenS2{F0: 85, F1: -9223372036854775808, F2: 0}
    var v9 uint64 = (^(v1))
    v10 := GenS1{F0: "hello"}
    switch "foo bar" {
    case "a":
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 uint = v2
            v13 := GenS2{F0: genFn0(), F1: 42, F2: -1}
            v2 += v2
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    if (2 <= 1) {
        var v14 uint16 = 65535
        _ = v14
    }
    if (255 < 106) {
        fmt.Println(v2, v1, v2)
        if (v9 <= uint64(v2)) {
            for fi15 := 0; fi15 < 8; fi15++ {
                for fi16 := 0; fi16 < 13; fi16++ {
                    var v17 int8 = 100
                    v18 := []int{586}
                    if (!false) {
                        v1 += (v9 >> ((uint(v1)) & 63))
                    } else {
                        if ((255 != 90) || true) {
                            if len(v18) > 0 {
                                v18[0] = (fi15 << ((v2) & 63))
                            }
                            v19 := GenS3{F0: 0}
                            switch fi16 {
                            case 485:
                                var v20 uint = (v2 ^ v2)
                                _ = v20
                            case 913:
                                _ = len(v18)
                            default:
                                switch (fi16 << ((v2) & 63)) {
                                case 0:
                                    var v21 float64 = ((1.5 / 1.0) * 3.14159)
                                    _ = v21
                                case -921:
                                    v22 := GenS2{F0: 32767, F1: -722, F2: 32767}
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                    _ = v22
                                case -1:
                                    switch fi16 {
                                    case 0:
                                        var v23 string = ("foo bar" + (v3 + v3))
                                        _ = v23
                                    case -41:
                                        var v24 uint64 = v1
                                        _ = v24
                                    case 1:
                                        v17 -= (v17 ^ (v17 / ((v17) | 1)))
                                    default:
                                        v1 += v1
                                    }
                                default:
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                    v25 := GenS4{F0: "foo bar", F1: (v17 - v17)}
                                    var v26 float32 = 0.1
                                    _ = v25
                                    _ = v26
                                }
                            }
                            if true {
                                switch 2147483647 {
                                case 127:
                                    if false {
                                        v27 := GenS2{F0: fi16, F1: 9223372036854775807, F2: 778}
                                        _ = v27
                                    } else {
                                        v28 := GenS4{F0: v3, F1: (v17 / (((^(v17))) | 1))}
                                        _ = v28
                                    }
                                case 2147483647:
                                    v3 = v3
                                case 236:
                                    fmt.Println(v9, v9, fi16)
                                    v17 -= 127
                                    v3 = ""
                                default:
                                    switch v3 {
                                    case "a":
                                        if false {
                                            v1 += v9
                                        } else {
                                            fmt.Println(fi15, fi16, fi16)
                                            var v29 bool = (!false)
                                            v9 = 2
                                            _ = v29
                                        }
                                    default:
                                        _ = len(v18)
                                    }
                                }
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                                var v30 int32 = 1
                                _ = v30
                            } else {
                                if false {
                                    v9 = (v9 & v9)
                                    if false {
                                        switch v3 {
                                        case "\n\t":
                                            switch "" {
                                            case "a":
                                                fmt.Println(fi15, v9)
                                            case "\n\t":
                                                v31 := GenS5{F0: (fi16 / (((fi16 << ((v2) & 63))) | 1)), F1: (2147483647 != 2), F2: -787}
                                                _ = v31
                                            default:
                                                fmt.Println(v17)
                                            }
                                        case "foo bar":
                                            var v32 float32 = 1.5
                                            _ = v32
                                        }
                                        if true {
                                            if true {
                                                if false {
                                                    v33 := GenS1{F0: v3}
                                                    _ = v33
                                                } else {
                                                    v34 := GenS2{F0: fi16, F1: -1, F2: -1}
                                                    _ = v34
                                                }
                                            } else {
                                                v1 = 0
                                            }
                                            var v35 uint32 = 2
                                            var v36 int32 = 2147483647
                                            _ = v35
                                            _ = v36
                                        } else {
                                            v3 = (v3 + ((((("hello" + "a") + v3) + "foo bar") + ("foo bar" + ("a" + (v3 + ("a" + "hello"))))) + ""))
                                            var v37 uint32 = 65535
                                            if len(v18) >= 2 {
                                                v18 = v18[1:len(v18)]
                                            }
                                            _ = v37
                                        }
                                        v38 := GenS6{F0: genFn0(), F1: -498}
                                        _ = v38
                                    }
                                    v18 = append(v18, fi15)
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                }
                                var v39 bool = false
                                v1 = v1
                                _ = v39
                            }
                            v1 += (v1 * v1)
                            v18 = append(v18, -9223372036854775808)
                            _ = v19
                        } else {
                            v40 := GenS7{F0: -921, F1: 42}
                            if (!(!(-2147483648 >= (fi15 % ((fi15) | 1))))) {
                                var v41 string = v3
                                if true {
                                    var v42 uint = (v2 - v2)
                                    var v43 float64 = 1.5
                                    _ = v42
                                    _ = v43
                                }
                                v17 = -1
                                _ = v41
                            } else {
                                switch (v3 + v3) {
                                case "":
                                    var v44 uint64 = (v9 / ((42) | 1))
                                    _ = v44
                                case "hello":
                                    v45 := GenS8{F0: (v2 + 1), F1: (0 == (v17 * (-(v17))))}
                                    _ = v45
                                case "foo bar":
                                    switch "\n\t" {
                                    case "foo bar":
                                        if len(v18) > 0 {
                                            v18[0] = (fi16 >> ((v2) & 63))
                                        }
                                    default:
                                        var v46 int8 = (-(v17))
                                        _ = v46
                                    }
                                }
                                v1 = (v1 ^ (v9 ^ v1))
                            }
                            _ = len(v18)
                            _ = v40
                        }
                    }
                    v2 -= 42
                    if (!true) {
                        switch (fi15 % ((fi16) | 1)) {
                        case 608:
                            switch 4294967295 {
                            case -618:
                                fmt.Println(v9, v17, v1)
                            }
                        }
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                v47 := GenS9{F0: (v2 * 255), F1: 100, F2: 0}
                _ = fi15
                _ = v47
            }
            v48 := GenS6{F0: 103, F1: -701}
            _ = v48
        } else {
            if (v3 != "\n\t") {
                if false {
                    v3 = (v3 + "")
                    var v49 string = v3
                    v50 := []int{int(v2)}
                    _ = v49
                    _ = v50
                }
            } else {
                fmt.Println(v9, v3)
                switch v3 {
                case "\n\t":
                    for fi51 := 0; fi51 < 2; fi51++ {
                        var v52 uint64 = (v1 | 4294967295)
                        _ = fi51
                        _ = v52
                    }
                }
            }
            switch 128 {
            case 821:
                var v53 uint = 42
                _ = v53
            case 32768:
                v2 -= (v2 << (((v2 / ((580) | 1))) & 63))
            default:
                v54 := GenS8{F0: (^(v2)), F1: ((4294967295 > 32768) || ((1.5 / 0.1) == (((1.5 / (3.14159 - (-2.25 + 1.5))) / 1.5) * 0.0)))}
                _ = v54
            }
        }
    }
    for fi55 := 0; fi55 < 20; fi55++ {
        var v56 int8 = 127
        if (3.14159 == 0.1) {
            v57 := []int{fi55, (fi55 | fi55), 127}
            v3 = "\n\t"
            _ = v57
        } else {
            var v58 uint16 = 373
            _ = v58
        }
        v59 := []int{}
        _ = fi55
        _ = v56
        _ = v59
    }
    v1 = v1
    switch v3 {
    case "":
        var v60 float32 = float32(1.5)
        _ = v60
    case "a":
        for fi61 := 0; fi61 < 5; fi61++ {
            v62 := []int{(fi61 << ((v2) & 63))}
            for fi63 := 0; fi63 < 8; fi63++ {
                v1 = 1
                var v64 int16 = -920
                _ = fi63
                _ = v64
            }
            switch (fi61 | fi61) {
            case -2147483648:
                var v65 int32 = 42
                _ = v65
            case 32768:
                var v66 int8 = -1
                _ = v66
            case -9223372036854775808:
                v62 = append(v62, (fi61 | fi61))
            }
            _ = fi61
            _ = v62
        }
    case "foo bar":
        switch (((("hello" + v3) + "") + v3) + v3) {
        case "a":
            switch -2147483648 {
            case 1:
                v67 := []int{}
                _ = v67
            case 0:
                v68 := GenS9{F0: (v2 % (((^(v2))) | 1)), F1: -1, F2: 6}
                _ = v68
            }
        default:
            fmt.Println(v9)
        }
    default:
        var v69 float64 = -2.25
        _ = v69
    }
    v9 += (v9 / ((v1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
