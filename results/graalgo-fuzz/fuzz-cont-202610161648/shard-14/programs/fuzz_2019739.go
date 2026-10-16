package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int8; F1 int64; F2 int }
type GenS5 struct { F0 bool; F1 uint8 }
func genFn0(p0 string, p1 int32) uint64 {
    var v7 float32 = -2.25
    if false {
        var v8 uint64 = 2
        v9 := []int{}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 0
}
func genFn1(p0 int8, p1 int) uint {
    p1 = (^(p1))
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := GenS1{F0: 1, F1: "a"}
        fmt.Println(p0, fi10)
        _ = fi10
        _ = v11
    }
    fmt.Println(p1, p0)
    var v12 uint = 0
    _ = p0
    _ = p1
    _ = v12
    return (^(v12))
}

func main() {
    var v1 uint = 4294967295
    var v2 uint = (^(v1))
    var v3 uint16 = 401
    v4 := make(map[int]bool)
    var v5 int = -189
    var v6 uint8 = 2
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 int8 = -109
        v3 = v3
        switch (fi13 + v5) {
        case 32768:
            v15 := v4[v5]
            _ = v15
        case -347:
            if false {
                v2 = (v2 % ((v1) | 1))
                var v16 int = (v5 + -2147483648)
                v17 := GenS2{F0: (1000 > 42), F1: "\n\t"}
                _ = v16
                _ = v17
            } else {
                delete(v4, (v5 / (((v5 << (((v2 * v2)) & 63))) | 1)))
            }
        default:
            v6 -= (v6 ^ v6)
        }
        _ = fi13
        _ = v14
    }
    v3 = v3
    if (-1 != 18) {
        v1 -= v1
        for fi18 := 0; fi18 < 18; fi18++ {
            var v19 float32 = 1.5
            var v20 uint64 = 1
            if (-1 > -158) {
                for fi21 := 0; fi21 < 19; fi21++ {
                    v22 := []int{(v5 % (((^(fi18))) | 1)), (fi21 | v5), 1, (fi18 << (((v2 << (((^(v1))) & 63))) & 63))}
                    v6 -= v6
                    _ = fi21
                    _ = v22
                }
            }
            _ = fi18
            _ = v19
            _ = v20
        }
        if true {
            var v23 bool = (true && (25 <= 100))
            v2 -= v1
            v24 := []int{int(v6), v5, (v5 << ((18446744073709551615) & 63)), v5}
            _ = v23
            _ = v24
        }
    }
    if false {
        v25 := GenS3{F0: -974, F1: v6}
        v26 := v4[v5]
        fmt.Println(v26)
        _ = v25
        _ = v26
    } else {
        fmt.Println(v3)
        if false {
            var v27 float32 = -2.25
            if ((v3 * (^(v3))) >= v3) {
                v2 -= v2
                if true {
                    if true {
                        var v28 int32 = -972
                        if false {
                            var v29 float32 = v27
                            _ = v29
                        }
                        _ = v28
                    }
                    v30 := GenS3{F0: 1, F1: (^(v6))}
                    v27 -= ((((v27 * (v27 - float32(1e10))) - v27) + -2.25) * v27)
                    _ = v30
                } else {
                    v31 := []int{int(v6), (v5 / ((v5) | 1)), (v5 & (v5 - v5)), v5}
                    _ = v31
                }
                var v32 uint8 = 0
                _ = v32
            }
            _ = v27
        }
        switch ("a" + "a") {
        case "a":
            v33 := []int{}
            _ = v33
        case "hello":
            if (((!(1 < 9223372036854775807)) && ((0.1 - -2.25) != float64(3.14159))) && false) {
                var v34 uint16 = v3
                v35 := []int{v5, 1}
                if false {
                    var v36 int32 = 15
                    v3 += (v3 | 1)
                    _ = v36
                } else {
                    switch (v5 << ((v2) & 63)) {
                    case -9223372036854775808:
                        if true {
                            for fi37 := 0; fi37 < 4; fi37++ {
                                var v38 int64 = -9223372036854775808
                                v2 = genFn1(-128, (fi37 - (v5 - fi37)))
                                var v39 uint64 = 42
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                            switch ("foo bar" + "a") {
                            case "\n\t":
                                switch v5 {
                                case -336:
                                    for fi40 := 0; fi40 < 5; fi40++ {
                                        var v41 int8 = -128
                                        var v42 int16 = 0
                                        _ = fi40
                                        _ = v41
                                        _ = v42
                                    }
                                }
                            default:
                                var v43 int = int(v34)
                                _ = v43
                            }
                        }
                    case 431:
                        for fi44 := 0; fi44 < 3; fi44++ {
                            for fi45 := 0; fi45 < 5; fi45++ {
                                var v46 bool = (0 <= int16(v3))
                                v2 += (v1 & genFn1(100, int(v34)))
                                if (!((true && v46) && ((("\n\t" + "\n\t") + "") == ("a" + ("\n\t" + "hello"))))) {
                                    var v47 float64 = 1.5
                                    _ = v47
                                } else {
                                    v5 = (v5 ^ fi45)
                                    v3 += v34
                                    var v48 uint = genFn1(int8(v3), (fi45 & int(v6)))
                                    _ = v48
                                }
                                _ = fi45
                                _ = v46
                            }
                            _ = fi44
                        }
                    default:
                        for fi49 := 0; fi49 < 12; fi49++ {
                            v50 := GenS4{F0: 39, F1: 0, F2: v5}
                            if (-1 >= -128) {
                                v3 -= 1
                            }
                            fmt.Println(v2, v34, v34)
                            _ = fi49
                            _ = v50
                        }
                    }
                    var v51 int8 = 127
                    _ = v51
                }
                _ = v34
                _ = v35
            }
        case "":
            switch "hello" {
            case "\n\t":
                if false {
                    delete(v4, (v5 - 2147483647))
                } else {
                    var v52 bool = (!(v3 <= (v3 % ((v3) | 1))))
                    _ = v52
                }
            }
        }
    }
    var v53 int = v5
    for fi54 := 0; fi54 < 13; fi54++ {
        for fi55 := 0; fi55 < 19; fi55++ {
            v2 = 2
            _ = fi55
        }
        v56 := GenS2{F0: true, F1: ("" + "hello")}
        var v57 int32 = 1
        _ = fi54
        _ = v56
        _ = v57
    }
    if true {
        var v58 int64 = -1
        var v59 uint16 = (v3 << ((v1) & 63))
        _ = v58
        _ = v59
    } else {
        for fi60 := 0; fi60 < 5; fi60++ {
            var v61 int32 = -706
            v53 = fi60
            _ = fi60
            _ = v61
        }
        for fi62 := 0; fi62 < 6; fi62++ {
            var v63 uint64 = uint64(v2)
            _ = fi62
            _ = v63
        }
        var v64 uint16 = (v3 / (((v3 / ((v3) | 1))) | 1))
        _ = v64
    }
    var v65 string = ("hello" + "a")
    v66 := []int{(v53 >> (((v1 % (((^(v2))) | 1))) & 63))}
    if (-92 != 488) {
        v67 := GenS5{F0: true, F1: v6}
        v3 = (v3 >> (((v2 << (((v1 - genFn1(-41, (v5 / (((v53 * (v53 & v53))) | 1))))) & 63))) & 63))
        if ((true && false) && false) {
            v68 := GenS3{F0: -9223372036854775808, F1: (^(v6))}
            v5 = (^(v53))
            _ = v68
        }
        _ = v67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("slicelen v66:", len(v66), v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v65
    _ = v66
}
