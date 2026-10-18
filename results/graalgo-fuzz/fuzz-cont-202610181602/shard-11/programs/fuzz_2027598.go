package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 bool; F1 int16; F2 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32; F1 int16; F2 int }
func genFn0(p0 bool) int16 {
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 int32 = 77
        p0 = (!(true || ((fi6 << ((4294967295) & 63)) < fi6)))
        _ = fi6
        _ = v7
    }
    v8 := []int{127}
    var v9 uint32 = 0
    v9 += v9
    _ = p0
    _ = v8
    _ = v9
    return -1
}
func genFn1() int16 {
    v10 := GenS1{F0: ((1 < -803) || false)}
    var v11 uint32 = 2147483647
    _ = v10
    _ = v11
    return -1
}

func main() {
    var v1 int = -1
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    var v4 bool = (474 >= -2147483648)
    var v5 bool = v4
    var v12 int16 = 1
    fmt.Println(v5, v4, v5)
    v13 := GenS2{F0: 698, F1: (v12 * (v12 << ((255) & 63))), F2: 42}
    var v14 uint32 = 2
    for fi15 := 0; fi15 < 20; fi15++ {
        if true {
            v14 = (v14 - (v14 ^ (v14 - v14)))
            var v16 float32 = 0.1
            _ = v16
        }
        var v17 string = "hello"
        if (-1 != 42) {
            for fi18 := 0; fi18 < 13; fi18++ {
                _ = len(v3)
                v4 = (true || v5)
                var v19 int64 = 9223372036854775807
                _ = fi18
                _ = v19
            }
            for fi20 := 0; fi20 < 8; fi20++ {
                v14 += v14
                var v21 string = (v17 + "\n\t")
                _ = fi20
                _ = v21
            }
            var v22 uint16 = 513
            _ = v22
        } else {
            v17 = (((v17 + "foo bar") + v17) + (v17 + (v17 + (v17 + (v17 + v17)))))
            if false {
                var v23 uint64 = 4294967295
                var v24 bool = v4
                v25 := GenS1{F0: v4}
                _ = v23
                _ = v24
                _ = v25
            } else {
                v17 = (v17 + (v17 + "a"))
                for fi26 := 0; fi26 < 1; fi26++ {
                    var v27 bool = v5
                    switch (v17 + v17) {
                    case "\n\t":
                        v17 = v17
                        var v28 int16 = v12
                        if true {
                            var v29 bool = false
                            v30 := GenS3{F0: 1, F1: -1}
                            _ = v29
                            _ = v30
                        } else {
                            v31 := GenS4{F0: (v27 && v4), F1: v12, F2: (v17 + ((v17 + v17) + (v17 + v17)))}
                            var v32 uint = 1
                            _ = v31
                            _ = v32
                        }
                        _ = v28
                    case "hello":
                        var v33 float64 = (float64(-2.25) * -2.25)
                        _ = v33
                    default:
                        switch 0 {
                        case 938:
                            v3[(fi15 | (v1 ^ (-(fi26))))] = 1
                        case 797:
                            v34 := GenS2{F0: 254, F1: (v12 - 219), F2: v14}
                            _ = v34
                        }
                    }
                    if v5 {
                        if false {
                            v17 = ("hello" + "a")
                            v3[42] = (v14 | 42)
                        } else {
                            if false {
                                v35 := v3[fi26]
                                var v36 uint32 = v35
                                if (0.0 <= (1.5 * (float64((1.5 / (-2.25 / 1.0))) / 1.0))) {
                                    v37 := GenS5{F0: (fi15 * 127)}
                                    _ = v37
                                }
                                _ = v35
                                _ = v36
                            }
                            if v5 {
                                v4 = v27
                                switch (fi15 + int(v12)) {
                                case 42:
                                    switch (fi26 ^ (-(v1))) {
                                    case 128:
                                        fmt.Println(fi26, v5)
                                    case 9223372036854775807:
                                        fmt.Println(v17, v14, v27)
                                    case -994:
                                        var v38 int = 0
                                        _ = v38
                                    default:
                                        if v5 {
                                            switch (v1 / ((-464) | 1)) {
                                            case -657:
                                                v39 := []int{(fi15 % ((v1) | 1))}
                                                _ = v39
                                            }
                                            v40 := []int{}
                                            _ = v40
                                        } else {
                                            _ = len(v3)
                                        }
                                    }
                                case 128:
                                    v5 = v4
                                case 4294967295:
                                    var v41 uint32 = 955
                                    _ = v41
                                default:
                                    var v42 uint = 255
                                    _ = v42
                                }
                                if (!(false && v5)) {
                                    if (false && v27) {
                                        v3[v1] = (v14 + (v14 & v14))
                                        var v43 uint8 = 1
                                        _ = v43
                                    }
                                }
                            }
                        }
                    } else {
                        switch fi15 {
                        case 546:
                            v14 += (v14 >> ((894) & 63))
                            delete(v3, 128)
                            var v44 uint8 = 2
                            _ = v44
                        case 320:
                            v45 := GenS2{F0: 1, F1: v12, F2: 140}
                            _ = v45
                        }
                    }
                    _ = fi26
                    _ = v27
                }
            }
            if v5 {
                _ = len(v2)
            }
        }
        _ = fi15
        _ = v17
    }
    var v46 bool = true
    for fi47 := 0; fi47 < 9; fi47++ {
        fmt.Println(v4, v14, v5)
        for fi48 := 0; fi48 < 16; fi48++ {
            if (true || v4) {
                v1 = fi47
                v49 := GenS6{F0: 2147483647, F1: 1000, F2: v1}
                if (664 == 899) {
                    delete(v2, fi47)
                }
                switch ("hello" + "") {
                case "a":
                    var v50 string = ("hello" + "a")
                    _ = v50
                case "":
                    var v51 int16 = 916
                    _ = v51
                }
                delete(v3, -797)
                _ = v49
            } else {
                v14 = v14
                if (249 <= 148) {
                    v12 -= (^(v12))
                }
            }
            var v52 uint32 = 4294967295
            _ = fi48
            _ = v52
        }
        _ = fi47
    }
    for fi53 := 0; fi53 < 12; fi53++ {
        v5 = ((255 >= 240) && true)
        _ = fi53
    }
    v54 := GenS3{F0: 0, F1: 1}
    switch "hello" {
    case "foo bar":
        var v55 uint16 = 189
        _ = v55
    default:
        for fi56 := 0; fi56 < 3; fi56++ {
            for fi57 := 0; fi57 < 20; fi57++ {
                var v58 int32 = 0
                var v59 uint64 = 18446744073709551615
                _ = fi57
                _ = v58
                _ = v59
            }
            _ = fi56
        }
    }
    fmt.Println(v5)
    var v60 uint16 = 265
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v54
    _ = v60
}
