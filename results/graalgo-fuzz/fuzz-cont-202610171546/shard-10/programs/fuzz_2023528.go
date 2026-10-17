package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint }
type GenS3 struct { F0 uint16; F1 uint32; F2 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint32; F1 string }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int32 }
func genFn0() int8 {
    switch ("foo bar" + (("foo bar" + "a") + (((((("foo bar" + ("hello" + "hello")) + ("a" + "")) + ("hello" + "hello")) + (("a" + ("a" + "a")) + "\n\t")) + "hello") + ("hello" + ("foo bar" + ("\n\t" + (("hello" + "") + "a"))))))) {
    case "foo bar":
        for fi6 := 0; fi6 < 10; fi6++ {
            switch "foo bar" {
            case "hello":
                v7 := GenS1{F0: 4294967295, F1: 196, F2: 54}
                _ = v7
            case "\n\t":
                v8 := GenS1{F0: 4294967295, F1: 65535, F2: 1}
                _ = v8
            default:
                var v9 uint = 42
                _ = v9
            }
            v10 := GenS1{F0: 311, F1: 2, F2: 0}
            v11 := []int{(fi6 ^ (fi6 ^ (fi6 - fi6)))}
            _ = fi6
            _ = v10
            _ = v11
        }
    case "a":
        var v12 bool = (!true)
        _ = v12
    case "hello":
        var v13 uint = 781
        _ = v13
    default:
        var v14 uint8 = 127
        _ = v14
    }
    v15 := []int{-620}
    if (false || true) {
        var v16 int16 = 42
        v16 += v16
        var v17 int = 9223372036854775807
        _ = v16
        _ = v17
    } else {
        var v18 bool = true
        v18 = v18
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 int8 = 0
            v21 := GenS2{F0: 65, F1: 0, F2: 42}
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = v18
    }
    _ = v15
    return 0
}
func genFn1() uint16 {
    var v22 uint16 = 865
    var v23 uint = 413
    _ = v22
    _ = v23
    return 2
}

func main() {
    var v1 int64 = -267
    v2 := []string{"hello", ("" + "hello"), "a", "hello"}
    v3 := []int{0, -2147483648}
    var v4 uint = 18446744073709551615
    v5 := []uint32{2147483648, 862}
    fmt.Println(v4, v1)
    fmt.Println(v1, v4)
    if false {
        v4 = v4
        v4 = 184
    }
    var v24 uint16 = 65535
    if false {
        _ = len(v5)
    }
    var v25 uint8 = 253
    fmt.Println(v24, v1)
    v24 = 2
    if len(v5) > 0 {
        v5[0] = 2147483648
    }
    var v26 uint8 = (^(v25))
    v4 = (v4 << (((^(v4))) & 63))
    var v27 int64 = (v1 | v1)
    v5 = append(v5, 2147483647)
    if false {
        switch 32768 {
        case -9223372036854775808:
            for fi28 := 0; fi28 < 10; fi28++ {
                v24 = (^(v24))
                _ = len(v2)
                v1 -= 42
                _ = fi28
            }
        case -2147483648:
            _ = len(v3)
        default:
            v1 += v1
        }
        if len(v5) > 0 {
            v5[0] = 215
        }
        for fi29 := 0; fi29 < 4; fi29++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v30 := GenS3{F0: v24, F1: 2147483647, F2: "\n\t"}
            var v31 float64 = ((0.0 + (((0.1 + 1e10) - ((float64(3.14159) / -2.25) / 0.1)) / float64(1.5))) * -2.25)
            _ = fi29
            _ = v30
            _ = v31
        }
    } else {
        v32 := GenS4{F0: true}
        if true {
            for fi33 := 0; fi33 < 13; fi33++ {
                fmt.Println(v27)
                v34 := GenS2{F0: (v25 - v26), F1: 4294967295, F2: 527}
                _ = fi33
                _ = v34
            }
            if false {
                if false {
                    fmt.Println(v24)
                    switch 127 {
                    case 618:
                        for fi35 := 0; fi35 < 10; fi35++ {
                            for fi36 := 0; fi36 < 17; fi36++ {
                                var v37 int8 = 118
                                if (!true) {
                                    v37 -= (^(v37))
                                } else {
                                    fmt.Println(v24)
                                    v38 := GenS5{F0: v1, F1: 1}
                                    var v39 uint = 908
                                    _ = v38
                                    _ = v39
                                }
                                var v40 string = "\n\t"
                                _ = fi36
                                _ = v37
                                _ = v40
                            }
                            v41 := GenS6{F0: uint16(v4)}
                            v42 := GenS7{F0: (^(v4))}
                            _ = fi35
                            _ = v41
                            _ = v42
                        }
                    case -9223372036854775808:
                        if len(v3) > 0 {
                            v3[0] = -2147483648
                        }
                    default:
                        for fi43 := 0; fi43 < 4; fi43++ {
                            v44 := GenS6{F0: v24}
                            var v45 float32 = -2.25
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                    }
                    if false {
                        v27 = (v1 << ((v4) & 63))
                        switch 128 {
                        case 1:
                            var v46 int32 = -1
                            _ = v46
                        case 128:
                            if false {
                                fmt.Println(v25)
                                fmt.Println(v1, v1)
                            }
                        }
                        switch -363 {
                        case 128:
                            var v47 uint16 = (v24 / ((v24) | 1))
                            _ = v47
                        default:
                            fmt.Println(v1)
                        }
                    } else {
                        var v48 int16 = 32767
                        for fi49 := 0; fi49 < 6; fi49++ {
                            v50 := GenS8{F0: 0, F1: ("foo bar" + (("a" + "") + "a"))}
                            for fi51 := 0; fi51 < 2; fi51++ {
                                v52 := GenS8{F0: 0, F1: "hello"}
                                _ = fi51
                                _ = v52
                            }
                            _ = fi49
                            _ = v50
                        }
                        var v53 uint = 697
                        _ = v48
                        _ = v53
                    }
                } else {
                    v1 -= 42
                }
            } else {
                for fi54 := 0; fi54 < 5; fi54++ {
                    v25 += 1
                    var v55 int = (^(fi54))
                    if false {
                        var v56 uint64 = 0
                        fmt.Println(v24, v56, fi54)
                        _ = v56
                    } else {
                        var v57 float32 = 0.0
                        v58 := GenS9{F0: (v25 * v26)}
                        v27 = -9223372036854775808
                        _ = v57
                        _ = v58
                    }
                    _ = fi54
                    _ = v55
                }
                v59 := GenS10{F0: 0}
                _ = v59
            }
        } else {
            v27 = 42
            v24 = v24
            fmt.Println(v25, v27)
        }
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
