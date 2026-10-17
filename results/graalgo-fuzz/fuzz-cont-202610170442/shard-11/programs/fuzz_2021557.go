package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 uint64; F1 uint; F2 int64 }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint; F1 int32; F2 uint32 }
func genFn0(p0 int32, p1 uint8, p2 int) uint8 {
    p2 = int(p1)
    fmt.Println(p1, p1)
    switch "foo bar" {
    case "":
        var v6 int8 = -1
        _ = v6
    case "hello":
        p0 = p0
    case "a":
        for fi7 := 0; fi7 < 4; fi7++ {
            p0 = p0
            p2 += p2
            _ = fi7
        }
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        switch -9223372036854775808 {
        case 1:
            p2 = (p2 >> ((886) & 63))
        case -513:
            for fi9 := 0; fi9 < 10; fi9++ {
                p2 -= 127
                var v10 uint16 = 32767
                _ = fi9
                _ = v10
            }
        case -792:
            p1 = p1
        }
        v11 := []int{fi8, -1}
        _ = len(v11)
        _ = fi8
        _ = v11
    }
    var v12 uint32 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return (^(p1))
}
func genFn1(p0 uint, p1 bool, p2 uint64) uint64 {
    p2 -= 4294967295
    if p1 {
        v13 := GenS1{F0: true, F1: 32767, F2: (p2 << ((p0) & 63))}
        var v14 uint16 = 32767
        _ = v13
        _ = v14
    } else {
        if p1 {
            for fi15 := 0; fi15 < 15; fi15++ {
                v16 := []int{}
                _ = fi15
                _ = v16
            }
        } else {
            for fi17 := 0; fi17 < 14; fi17++ {
                p0 = p0
                fmt.Println(p1, p1)
                v18 := []int{fi17, fi17, fi17, (fi17 - int(p2))}
                _ = fi17
                _ = v18
            }
            if false {
                v19 := GenS2{F0: 42}
                _ = v19
            }
            if p1 {
                if false {
                    if p1 {
                        var v20 bool = ((-1 != 1) && ((((1.5 / 1.0) - 1.5) / float64(1.5)) >= ((-2.25 + 1e10) / float64(-2.25))))
                        p1 = ((0.0 >= 1.5) || v20)
                        _ = v20
                    }
                    p2 = (p2 / (((p2 << (((p0 - uint(p2))) & 63))) | 1))
                }
            }
        }
        fmt.Println(p1, p0, p0)
    }
    if p1 {
        var v21 float64 = -2.25
        v21 += v21
        _ = v21
    }
    v22 := GenS3{F0: -520, F1: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return (p2 / ((2) | 1))
}

func main() {
    var v1 bool = (!(0 != 255))
    v2 := make(map[uint32]int)
    var v3 string = "foo bar"
    v4 := []int{}
    v5 := []string{("foo bar" + v3), v3}
    var v23 uint64 = 4294967295
    delete(v2, 2147483647)
    v1 = (false && (333 == 65535))
    if len(v4) > 0 {
        v4[0] = -1
    }
    if v1 {
        v24 := GenS4{F0: (v23 + v23), F1: 799, F2: 388}
        _ = v24
    }
    switch v3 {
    case "":
        v1 = v1
    default:
        v25 := GenS1{F0: (2 >= 192), F1: 922, F2: (v23 | v23)}
        _ = v25
    }
    v1 = v1
    var v26 int8 = 36
    var v27 bool = false
    switch 42 {
    case -694:
        delete(v2, 158)
    case 2147483647:
        v23 = (v23 % (((v23 | (v23 << ((1) & 63)))) | 1))
    }
    fmt.Println(v23, v3, v3)
    var v28 int16 = 654
    for fi29 := 0; fi29 < 1; fi29++ {
        var v30 int64 = 928
        for fi31 := 0; fi31 < 2; fi31++ {
            v1 = (!v27)
            var v32 uint64 = 4294967295
            fmt.Println(fi29, v26)
            _ = fi31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    switch -2147483648 {
    case 10:
        if (true || v27) {
            _ = len(v2)
            var v33 int16 = (v28 ^ 1000)
            switch ("a" + v3) {
            case "\n\t":
                v34 := GenS3{F0: 127, F1: 9223372036854775807}
                _ = v34
            default:
                if v1 {
                    var v35 string = "\n\t"
                    _ = v35
                }
            }
            _ = v33
        } else {
            v36 := v2[42]
            for fi37 := 0; fi37 < 7; fi37++ {
                switch ((v3 + "") + v3) {
                case "hello":
                    _ = len(v2)
                case "\n\t":
                    for fi38 := 0; fi38 < 20; fi38++ {
                        var v39 uint64 = genFn1(18446744073709551615, ((!(!v1)) || (!(!true))), genFn1(4294967295, (2 <= 128), genFn1(255, (!false), (v23 + 18446744073709551615))))
                        switch v3 {
                        case "\n\t":
                            var v40 bool = (v1 || ((^(fi37)) > (-(fi37))))
                            _ = v40
                        }
                        _ = fi38
                        _ = v39
                    }
                case "foo bar":
                    v4 = append(v4, (v36 / ((fi37) | 1)))
                }
                _ = fi37
            }
            var v41 int = (v36 | 4294967295)
            _ = v36
            _ = v41
        }
    case -986:
        if (2147483647 > 168) {
            v42 := v2[65535]
            if ((!(!v1)) || false) {
                switch v42 {
                case 32767:
                    delete(v2, 2)
                case 128:
                    var v43 int = v42
                    _ = v43
                default:
                    fmt.Println(v27, v23)
                }
                if v27 {
                    _ = len(v5)
                    var v44 uint8 = 128
                    for fi45 := 0; fi45 < 15; fi45++ {
                        if false {
                            switch (v3 + "a") {
                            case "hello":
                                v42 += (fi45 * fi45)
                            case "":
                                v44 = genFn0(-47, (^(v44)), v42)
                            }
                        }
                        _ = fi45
                    }
                    _ = v44
                }
                delete(v2, 4294967295)
            } else {
                var v46 string = ("" + v3)
                fmt.Println(v26, v46)
                v27 = true
                _ = v46
            }
            _ = v42
        } else {
            v47 := GenS1{F0: ((0 < (v26 ^ (-(v26)))) || v27), F1: -430, F2: (v23 | v23)}
            _ = v47
        }
    case 391:
        switch 32767 {
        case 0:
            for fi48 := 0; fi48 < 14; fi48++ {
                v27 = (588 > 4294967295)
                _ = fi48
            }
        }
    default:
        switch ("\n\t" + "") {
        case "a":
            var v49 int64 = 1
            _ = v49
        default:
            if v27 {
                fmt.Println(v3, v23)
                v26 += v26
            } else {
                switch 128 {
                case 127:
                    v50 := GenS5{F0: v28, F1: genFn1(0, v1, (v23 - (^(v23))))}
                    _ = v50
                case -1:
                    if false {
                        var v51 int16 = v28
                        v52 := GenS6{F0: 452}
                        for fi53 := 0; fi53 < 4; fi53++ {
                            for fi54 := 0; fi54 < 16; fi54++ {
                                v2[911] = fi53
                                _ = fi54
                            }
                            v55 := GenS2{F0: (v28 + (v51 % ((v28) | 1)))}
                            if v27 {
                                for fi56 := 0; fi56 < 10; fi56++ {
                                    var v57 int8 = v26
                                    var v58 uint64 = v23
                                    v59 := GenS4{F0: (^(v23)), F1: 2, F2: 0}
                                    _ = fi56
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                }
                                v28 += (v51 >> ((255) & 63))
                            } else {
                                delete(v2, 4294967295)
                                for fi60 := 0; fi60 < 1; fi60++ {
                                    v28 -= (v51 << ((42) & 63))
                                    if v1 {
                                        var v61 uint64 = v23
                                        _ = v61
                                    } else {
                                        v62 := GenS7{F0: 0}
                                        _ = v62
                                    }
                                    var v63 float32 = 0.1
                                    _ = fi60
                                    _ = v63
                                }
                            }
                            _ = fi53
                            _ = v55
                        }
                        _ = v51
                        _ = v52
                    }
                default:
                    v1 = (((3.14159 * 1.5) + ((float32(-2.25) + 0.1) / 1.0)) == 0.0)
                }
                v64 := GenS8{F0: 65535, F1: 817, F2: 945}
                _ = v64
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
