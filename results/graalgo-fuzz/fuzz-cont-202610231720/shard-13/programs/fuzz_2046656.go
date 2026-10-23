package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 bool; F1 string }
type GenS9 struct { F0 int; F1 string }
type GenS10 struct { F0 uint32; F1 uint64; F2 string }
func genFn0() uint32 {
    var v5 uint = 18446744073709551615
    if (true && false) {
        v5 = (v5 & (v5 << ((v5) & 63)))
        v6 := []int{-2147483648, -1}
        _ = v6
    }
    _ = v5
    return 65535
}
func genFn1(p0 uint) int64 {
    p0 += (p0 >> (((p0 - 255)) & 63))
    if (2 <= 0) {
        p0 = (p0 + p0)
        p0 -= p0
    }
    _ = p0
    return -701
}

func main() {
    var v1 uint64 = 83
    var v2 uint8 = 255
    var v3 int16 = 42
    var v4 int64 = 0
    fmt.Println(v1)
    v1 = v1
    v4 += -1000
    v7 := GenS1{F0: int8(v3), F1: -2147483648}
    v8 := GenS1{F0: 127, F1: -2147483648}
    var v9 int32 = 42
    var v10 int = 32767
    var v11 bool = false
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 string = "hello"
        if (v4 < genFn1(39)) {
            fmt.Println(v4, v13, v11)
            if ((128 >= (v2 - v2)) && false) {
                v14 := []int{v10, 127}
                v3 += 32767
                for fi15 := 0; fi15 < 9; fi15++ {
                    v3 -= v3
                    var v16 int = -2147483648
                    fmt.Println(v13, fi15)
                    _ = fi15
                    _ = v16
                }
                _ = v14
            } else {
                switch fi12 {
                case 1:
                    for fi17 := 0; fi17 < 2; fi17++ {
                        switch v13 {
                        case "":
                            fmt.Println(v3, fi12)
                            var v18 int64 = (v4 & (v4 + -1))
                            _ = v18
                        case "\n\t":
                            v4 += v4
                        case "hello":
                            if (v13 >= v13) {
                                v2 += (v2 + 2)
                                if true {
                                    v4 -= (v4 & 9223372036854775807)
                                } else {
                                    v19 := []int{fi12, (v10 + -9223372036854775808), -1}
                                    if v11 {
                                        v20 := GenS2{F0: 798}
                                        v10 = (v10 >> ((255) & 63))
                                        switch v13 {
                                        case "a":
                                            v10 -= -957
                                        case "foo bar":
                                            v13 = (((v13 + "hello") + ((v13 + "a") + "\n\t")) + (v13 + "a"))
                                        case "hello":
                                            v3 = (v3 * (v3 * 42))
                                        default:
                                            v10 = (fi17 << ((18446744073709551615) & 63))
                                        }
                                        _ = v20
                                    }
                                    switch "" {
                                    case "hello":
                                        fmt.Println(v10, v2, fi17)
                                        var v21 int64 = v4
                                        v1 = v1
                                        var v22 int64 = -1
                                        _ = v21
                                        _ = v22
                                    }
                                    fmt.Println(v3, v9, v11)
                                    _ = v19
                                }
                            } else {
                                v23 := GenS3{F0: (^(v10))}
                                _ = v23
                            }
                        }
                        _ = fi17
                    }
                case 32767:
                    var v24 int32 = v9
                    _ = v24
                case 4294967295:
                    v11 = v11
                default:
                    var v25 uint64 = 0
                    _ = v25
                }
                if true {
                    v10 += (v10 / ((fi12) | 1))
                } else {
                    v26 := GenS4{F0: (v9 | -2147483648), F1: v13}
                    _ = v26
                }
            }
        } else {
            v27 := GenS3{F0: (v10 | fi12)}
            switch v13 {
            case "hello":
                v28 := GenS5{F0: 42}
                _ = v28
            case "a":
                v29 := GenS6{F0: 724, F1: -80}
                _ = v29
            default:
                for fi30 := 0; fi30 < 2; fi30++ {
                    v4 = 9223372036854775807
                    _ = fi30
                }
            }
            if v11 {
                fmt.Println(v1)
                v31 := GenS6{F0: (v1 << ((773) & 63)), F1: 100}
                _ = v31
            } else {
                for fi32 := 0; fi32 < 5; fi32++ {
                    v3 = (v3 ^ -196)
                    v33 := GenS1{F0: 100, F1: -1}
                    v34 := GenS7{F0: 780}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                v2 = v2
            }
            _ = v27
        }
        switch "" {
        case "":
            switch v13 {
            case "":
                fmt.Println(v1, v13, v9)
            case "hello":
                for fi35 := 0; fi35 < 14; fi35++ {
                    switch fi12 {
                    case -9223372036854775808:
                        v36 := []int{(-(v10)), (fi35 & (fi12 * (^(fi35)))), v10, v10}
                        _ = v36
                    default:
                        var v37 int32 = v9
                        _ = v37
                    }
                    if (!false) {
                        if ((v9 % ((v9) | 1)) <= int32(fi12)) {
                            switch v13 {
                            case "foo bar":
                                var v38 bool = (!((v10 % (((fi12 % ((39) | 1))) | 1)) != fi35))
                                _ = v38
                            case "\n\t":
                                v3 += v3
                            default:
                                v39 := GenS6{F0: v1, F1: 127}
                                _ = v39
                            }
                            switch ("" + v13) {
                            case "\n\t":
                                var v40 int32 = (v9 / ((-904) | 1))
                                _ = v40
                            case "foo bar":
                                v41 := GenS8{F0: v11, F1: v13}
                                _ = v41
                            case "":
                                if true {
                                    v42 := GenS7{F0: 65535}
                                    _ = v42
                                } else {
                                    v43 := []int{-2147483648, (v10 << ((0) & 63)), -9223372036854775808, (fi12 | 32768)}
                                    var v44 uint64 = v1
                                    var v45 float64 = -2.25
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                }
                            default:
                                v11 = (!(v4 >= 1))
                            }
                            if (1 >= 255) {
                                fmt.Println(fi35, v10, v1)
                                switch v13 {
                                case "foo bar":
                                    fmt.Println(fi35)
                                default:
                                    fmt.Println(v10, v13, v13)
                                }
                                if v11 {
                                    v9 -= v9
                                    v9 = v9
                                    switch "foo bar" {
                                    case "hello":
                                        if v11 {
                                            if (v11 || v11) {
                                                fmt.Println(v1, v9)
                                                v46 := []int{(v10 / ((fi35) | 1))}
                                                _ = v46
                                            } else {
                                                v4 = (v4 << ((42) & 63))
                                            }
                                            v47 := []int{fi12, fi12, (^(fi35)), (fi35 | fi35)}
                                            var v48 bool = v11
                                            v47 = append(v47, (fi35 - fi35))
                                            _ = v47
                                            _ = v48
                                        } else {
                                            v49 := GenS1{F0: 0, F1: (v9 ^ (v9 & (^(v9))))}
                                            fmt.Println(v11, v4)
                                            var v50 uint64 = 42
                                            v51 := GenS6{F0: (v50 >> ((255) & 63)), F1: 100}
                                            _ = v49
                                            _ = v50
                                            _ = v51
                                        }
                                    case "foo bar":
                                        v52 := []int{}
                                        _ = v52
                                    case "a":
                                        v53 := []int{179, -2147483648}
                                        _ = v53
                                    }
                                    v9 = (v9 << ((42) & 63))
                                } else {
                                    v54 := []int{fi35, (^(v10)), fi12, (v10 | fi12)}
                                    _ = v54
                                }
                            }
                        }
                        v55 := GenS9{F0: 128, F1: v13}
                        v11 = (v11 && (!(v11 && (2 >= 1))))
                        _ = v55
                    }
                    v56 := []int{fi12}
                    _ = fi35
                    _ = v56
                }
            case "foo bar":
                v57 := GenS7{F0: 2147483648}
                _ = v57
            default:
                if (76 < 69) {
                    v58 := GenS10{F0: 559, F1: (v1 & uint64(v2)), F2: (v13 + v13)}
                    _ = v58
                }
            }
        case "\n\t":
            v59 := []int{fi12, (fi12 >> ((679) & 63)), (v10 ^ v10), -438}
            _ = v59
        case "foo bar":
            var v60 uint = 860
            _ = v60
        }
        _ = fi12
        _ = v13
    }
    v61 := GenS8{F0: false, F1: ""}
    v62 := GenS9{F0: (-(v10)), F1: "hello"}
    var v63 uint16 = 937
    v9 = v9
    var v64 int64 = v4
    for fi65 := 0; fi65 < 1; fi65++ {
        switch "\n\t" {
        case "a":
            switch 42 {
            case 128:
                v10 -= (v10 >> ((65535) & 63))
            }
        case "":
            fmt.Println(v2, v11, v4)
        case "\n\t":
            fmt.Println(fi65, v1)
        }
        v3 = -32768
        if v11 {
            v11 = v11
            fmt.Println(v9)
        }
        _ = fi65
    }
    fmt.Println(v9, v64, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v7
    _ = v8
    _ = v9
}
