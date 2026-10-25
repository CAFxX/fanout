package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 string }
func genFn0(p0 int64, p1 uint) int32 {
    p0 = (p0 * p0)
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint = p1
        v8 := []int{fi6, 127}
        fmt.Println(p1, fi6)
        _ = fi6
        _ = v7
        _ = v8
    }
    switch "hello" {
    case "a":
        v9 := []int{9223372036854775807, -407}
        _ = v9
    }
    _ = p0
    _ = p1
    return 791
}
func genFn1(p0 string, p1 string) uint16 {
    p1 = (p0 + ("\n\t" + ("foo bar" + p0)))
    var v10 uint64 = 2
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 bool = (!(0 > 621))
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := GenS1{F0: -1}
        _ = fi13
        _ = v14
    }
    var v15 int8 = 100
    _ = p0
    _ = p1
    _ = v10
    _ = v15
    return 1
}

func main() {
    var v1 uint8 = 128
    var v2 float64 = (-2.25 / 0.1)
    var v3 uint16 = 1
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]bool)
    v16 := GenS1{F0: 42}
    v17 := []int{138, 2147483647, 2147483647, -2147483648}
    var v18 uint8 = (v1 + v1)
    if (-458 >= 42) {
        v1 += (v18 | v1)
        v18 = (v18 - v1)
        for fi19 := 0; fi19 < 20; fi19++ {
            v20 := GenS1{F0: -91}
            var v21 int8 = 127
            _ = fi19
            _ = v20
            _ = v21
        }
    } else {
        fmt.Println(v2, v3, v3)
    }
    if (-2.25 <= 1.5) {
        v22 := GenS2{F0: (v3 - (^(v3)))}
        _ = v22
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 float64 = v2
        if true {
            var v25 int64 = -1
            _ = v25
        }
        _ = fi23
        _ = v24
    }
    if (true || true) {
        if (65535 >= 2147483647) {
            v26 := GenS3{F0: 76, F1: 1000}
            _ = v26
        }
    } else {
        v3 += 42
        if false {
            switch "foo bar" {
            case "":
                if true {
                    fmt.Println(v2, v3)
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                }
            case "hello":
                var v27 uint32 = 1
                _ = v27
            }
        } else {
            switch "a" {
            case "foo bar":
                if (139 > 18446744073709551615) {
                    v28 := v5[2]
                    _ = v28
                } else {
                    v29 := GenS1{F0: 0}
                    var v30 uint32 = 465
                    for fi31 := 0; fi31 < 8; fi31++ {
                        for fi32 := 0; fi32 < 6; fi32++ {
                            if false {
                                var v33 int16 = 869
                                _ = v33
                            } else {
                                v3 += (v3 << ((18446744073709551615) & 63))
                                if ((v3 / ((uint16(v18)) | 1)) >= v3) {
                                    if (false && true) {
                                        var v34 uint32 = (v30 + v30)
                                        v35 := GenS1{F0: 0}
                                        var v36 uint64 = 0
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    }
                                    v37 := GenS4{F0: 59, F1: 1000}
                                    v2 -= v2
                                    _ = v37
                                } else {
                                    v1 -= (v1 ^ (v18 >> ((460) & 63)))
                                    var v38 string = (("\n\t" + "") + "a")
                                    _ = v38
                                }
                                var v39 bool = false
                                _ = v39
                            }
                            switch ("hello" + "hello") {
                            case "a":
                                if (!(v30 != (^(v30)))) {
                                    _ = len(v17)
                                } else {
                                    v2 += float64(1e10)
                                    fmt.Println(v30, fi32, v3)
                                    if (459 <= 704) {
                                        if (!(v3 <= (v3 / ((uint16(v18)) | 1)))) {
                                            var v40 int64 = 42
                                            if (0.1 < (0.0 * (0.1 + ((0.0 * 1.5) * 0.0)))) {
                                                if (v40 != v40) {
                                                    var v41 uint64 = 433
                                                    v2 += v2
                                                    _ = v41
                                                } else {
                                                    fmt.Println(v18, v3)
                                                    fmt.Println(v3)
                                                }
                                                v42 := v5[v30]
                                                _ = v42
                                            }
                                            _ = len(v17)
                                            if len(v17) >= 2 {
                                                v17 = v17[1:len(v17)]
                                            }
                                            _ = v40
                                        } else {
                                            v3 = (v3 + uint16(v30))
                                            v5[v30] = true
                                        }
                                    } else {
                                        fmt.Println(v1, v1)
                                        var v43 int32 = -1
                                        var v44 uint32 = (^(v30))
                                        _ = v43
                                        _ = v44
                                    }
                                }
                            case "hello":
                                v45 := GenS5{F0: "foo bar"}
                                _ = v45
                            case "\n\t":
                                v1 = 204
                            }
                            fmt.Println(v1, fi32)
                            v3 = (v3 - uint16(v30))
                            if ((!(-9 != -93)) || true) {
                                v17 = append(v17, int(v3))
                                v18 = (v18 | (v18 / (((v1 | v1)) | 1)))
                                if len(v17) > 0 {
                                    v17[0] = -9223372036854775808
                                }
                            }
                            _ = fi32
                        }
                        _ = fi31
                    }
                    _ = v29
                    _ = v30
                }
            case "":
                v1 = (v1 ^ 128)
            case "a":
                if (false && (-2147483648 < 2147483647)) {
                    var v46 int = 128
                    _ = v46
                }
            }
            var v47 int32 = -1
            v48 := v5[4294967295]
            _ = v47
            _ = v48
        }
        switch 32768 {
        case 32767:
            v1 -= 255
        default:
            for fi49 := 0; fi49 < 4; fi49++ {
                v2 += (v2 - v2)
                switch "hello" {
                case "foo bar":
                    for fi50 := 0; fi50 < 12; fi50++ {
                        if false {
                            v51 := GenS4{F0: 1, F1: -1}
                            _ = v51
                        } else {
                            v18 = v1
                            v52 := v4[540]
                            _ = v52
                        }
                        fmt.Println(v1, v1, fi49)
                        _ = fi50
                    }
                case "a":
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                case "hello":
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                default:
                    v18 = uint8(v3)
                }
                _ = fi49
            }
        }
    }
    v53 := v5[910]
    v3 += 32767
    for fi54 := 0; fi54 < 5; fi54++ {
        v55 := GenS3{F0: 695, F1: 32767}
        _ = fi54
        _ = v55
    }
    v53 = false
    for fi56 := 0; fi56 < 4; fi56++ {
        for fi57 := 0; fi57 < 7; fi57++ {
            v53 = (521 == 830)
            var v58 string = ("\n\t" + "foo bar")
            switch fi57 {
            case -1:
                v1 = (v1 / ((v1) | 1))
            case 32768:
                v59 := GenS2{F0: v3}
                _ = v59
            case 128:
                v60 := GenS2{F0: (v3 & genFn1((("a" + ((v58 + (v58 + "a")) + v58)) + v58), (v58 + (v58 + v58))))}
                _ = v60
            }
            _ = fi57
            _ = v58
        }
        _ = fi56
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
}
