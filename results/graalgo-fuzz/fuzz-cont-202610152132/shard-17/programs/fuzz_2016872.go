package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int64; F2 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 int; F1 int; F2 int64 }
func genFn0() uint32 {
    var v6 uint = 538
    v6 = (v6 / ((0) | 1))
    _ = v6
    return 42
}
func genFn1(p0 uint, p1 int) bool {
    switch "foo bar" {
    case "":
        if true {
            v7 := GenS1{F0: p0, F1: ""}
            _ = v7
        } else {
            switch 9223372036854775807 {
            case -140:
                p0 = 791
            case 76:
                p0 = 255
            default:
                for fi8 := 0; fi8 < 3; fi8++ {
                    v9 := GenS1{F0: (p0 >> (((p0 >> (((^(p0))) & 63))) & 63)), F1: "hello"}
                    _ = fi8
                    _ = v9
                }
            }
        }
    default:
        switch "foo bar" {
        case "a":
            p1 -= 441
        case "foo bar":
            var v10 int8 = 100
            _ = v10
        case "hello":
            p0 = p0
        default:
            v11 := []int{}
            _ = v11
        }
    }
    var v12 string = "a"
    for fi13 := 0; fi13 < 18; fi13++ {
        p0 = (^(p0))
        var v14 uint = p0
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 uint8 = 128
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    if ((0.0 == 0.0) && false) {
        if true {
            if true {
                v17 := GenS1{F0: (p0 << (((p0 & 576)) & 63)), F1: "\n\t"}
                p1 = p1
                if false {
                    var v18 float32 = 0.1
                    if true {
                        fmt.Println(v18, v12, p1)
                    } else {
                        if (1 <= 42) {
                            switch v12 {
                            case "":
                                v19 := GenS1{F0: (p0 - 1), F1: v12}
                                _ = v19
                            case "a":
                                var v20 int8 = 100
                                _ = v20
                            }
                            for fi21 := 0; fi21 < 1; fi21++ {
                                v22 := GenS2{F0: -1}
                                v23 := []int{fi21, fi21}
                                var v24 string = (("foo bar" + (v12 + ("foo bar" + v12))) + v12)
                                _ = fi21
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                            v25 := GenS3{F0: v12, F1: 9223372036854775807, F2: -939}
                            _ = v25
                        }
                        switch v12 {
                        case "":
                            if true {
                                var v26 uint32 = 954
                                for fi27 := 0; fi27 < 16; fi27++ {
                                    v28 := GenS4{F0: 42}
                                    if false {
                                        fmt.Println(p1)
                                        switch (p1 * (p1 << (((p0 * (p0 / (((p0 / ((4294967295) | 1))) | 1)))) & 63))) {
                                        case 0:
                                            v29 := []int{-1}
                                            _ = v29
                                        case 2147483647:
                                            v30 := []int{}
                                            _ = v30
                                        }
                                        v12 = (("" + v12) + ("a" + ("a" + v12)))
                                    } else {
                                        var v31 int = (p1 * 4294967295)
                                        var v32 float32 = (float32(3.14159) / (0.1 - (v18 / v18)))
                                        _ = v31
                                        _ = v32
                                    }
                                    for fi33 := 0; fi33 < 8; fi33++ {
                                        var v34 uint64 = 1
                                        var v35 float64 = (float64(v18) * 1e10)
                                        _ = fi33
                                        _ = v34
                                        _ = v35
                                    }
                                    _ = fi27
                                    _ = v28
                                }
                                var v36 int = -879
                                _ = v26
                                _ = v36
                            }
                        case "foo bar":
                            v37 := []int{}
                            _ = v37
                        case "hello":
                            v38 := GenS3{F0: "\n\t", F1: -1, F2: -2147483648}
                            _ = v38
                        }
                    }
                    _ = v18
                } else {
                    var v39 int8 = -73
                    _ = v39
                }
                _ = v17
            }
            fmt.Println(p0, v12)
            var v40 int64 = 0
            _ = v40
        } else {
            fmt.Println(p1)
            for fi41 := 0; fi41 < 16; fi41++ {
                v12 = v12
                _ = fi41
            }
            fmt.Println(p1)
        }
    }
    v42 := []int{(^(p1)), 128, (p1 >> ((p0) & 63))}
    _ = p0
    _ = p1
    _ = v12
    _ = v42
    return (("foo bar" + (v12 + (("" + v12) + v12))) >= (v12 + (v12 + (v12 + v12))))
}

func main() {
    v1 := []uint32{65535, 691, 2147483647}
    var v2 int = 4294967295
    var v3 int32 = 1
    var v4 string = ("hello" + ("a" + ((("hello" + "hello") + ("\n\t" + "hello")) + "a")))
    v5 := []string{v4, (v4 + v4), v4, v4}
    v4 = "foo bar"
    var v43 int8 = 1
    if (896 < 1000) {
        v2 += v2
    }
    var v44 int32 = -2147483648
    var v45 int32 = 365
    v45 += (^(v3))
    v3 = (v44 % ((v44) | 1))
    v3 += (v44 / ((-1) | 1))
    v43 -= (v43 >> ((1) & 63))
    for fi46 := 0; fi46 < 6; fi46++ {
        v3 -= v3
        _ = fi46
    }
    switch v4 {
    case "\n\t":
        v3 = (v44 ^ v3)
    case "":
        for fi47 := 0; fi47 < 13; fi47++ {
            var v48 float64 = 1e10
            _ = fi47
            _ = v48
        }
    case "foo bar":
        v49 := GenS5{F0: 2, F1: 32767}
        _ = v49
    default:
        for fi50 := 0; fi50 < 4; fi50++ {
            v51 := GenS6{F0: 0, F1: (v3 >> ((18446744073709551615) & 63))}
            _ = fi50
            _ = v51
        }
    }
    switch v2 {
    case 9223372036854775807:
        fmt.Println(v45)
    case -2147483648:
        v3 = v45
    case 127:
        v52 := GenS7{F0: (v2 & (-(v2))), F1: (v2 + (v2 ^ (v2 + (-(v2))))), F2: 9223372036854775807}
        _ = v52
    default:
        _ = len(v5)
    }
    v53 := GenS4{F0: 30}
    v3 += 994
    for fi54 := 0; fi54 < 5; fi54++ {
        for fi55 := 0; fi55 < 15; fi55++ {
            if (-2147483648 >= v45) {
                v1 = append(v1, 104)
            }
            _ = fi55
        }
        if false {
            if false {
                fmt.Println(v44, v44, v44)
            }
        }
        _ = fi54
    }
    var v56 float64 = 1.5
    for fi57 := 0; fi57 < 18; fi57++ {
        fmt.Println(v4, v2)
        if false {
            var v58 float32 = 0.0
            var v59 int16 = -948
            _ = v58
            _ = v59
        } else {
            if true {
                if (!true) {
                    if false {
                        switch v4 {
                        case "foo bar":
                            v2 -= v2
                        default:
                            v5 = append(v5, (v4 + (v4 + "\n\t")))
                        }
                        for fi60 := 0; fi60 < 8; fi60++ {
                            var v61 int = v2
                            _ = fi60
                            _ = v61
                        }
                    }
                } else {
                    var v62 int16 = -1
                    _ = v62
                }
            }
        }
        _ = fi57
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v53
    _ = v56
}
