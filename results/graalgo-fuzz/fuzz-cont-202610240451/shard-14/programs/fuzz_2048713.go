package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS5 struct { F0 uint16; F1 uint32; F2 bool }
type GenS6 struct { F0 uint; F1 uint8 }
type GenS7 struct { F0 int8; F1 uint8 }
func genFn0() uint32 {
    var v4 int16 = 943
    var v5 uint16 = 32767
    if true {
        v4 -= (^(v4))
        switch ("\n\t" + "") {
        case "hello":
            v6 := GenS1{F0: (("a" + "") + "hello")}
            _ = v6
        case "\n\t":
            var v7 string = ""
            _ = v7
        default:
            for fi8 := 0; fi8 < 20; fi8++ {
                fmt.Println(v5)
                var v9 bool = true
                var v10 int8 = 100
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        if (42 >= 346) {
            fmt.Println(v4, v5)
            v5 -= v5
        } else {
            v11 := []int{-2147483648, 180}
            if true {
                v4 = v4
                var v12 uint = 1
                _ = v12
            }
            _ = v11
        }
    }
    _ = v4
    _ = v5
    return 2
}
func genFn1(p0 int) int16 {
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := []int{p0, fi13, (fi13 % ((p0) | 1)), -489}
        for fi15 := 0; fi15 < 5; fi15++ {
            if false {
                var v16 float64 = -2.25
                fmt.Println(fi15, fi13)
                p0 = (-(fi13))
                _ = v16
            } else {
                p0 -= (p0 - 127)
            }
            v14 = append(v14, (fi13 * (^(fi15))))
            v14 = append(v14, (-(fi13)))
            fmt.Println(fi13, fi15, fi13)
            var v17 float64 = -2.25
            _ = fi15
            _ = v17
        }
        var v18 float64 = 0.1
        _ = fi13
        _ = v14
        _ = v18
    }
    v19 := GenS1{F0: "foo bar"}
    p0 += 127
    v20 := []int{}
    _ = p0
    _ = v19
    _ = v20
    return 32767
}

func main() {
    var v1 uint32 = 876
    var v2 uint32 = v1
    var v3 uint32 = (v1 & v2)
    var v21 uint = 999
    if true {
        var v22 uint = v21
        for fi23 := 0; fi23 < 5; fi23++ {
            v21 += (v21 - v22)
            v24 := GenS1{F0: (((("foo bar" + "hello") + "") + "\n\t") + "foo bar")}
            v25 := []int{fi23, (fi23 * fi23)}
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = v22
    } else {
        for fi26 := 0; fi26 < 14; fi26++ {
            switch "hello" {
            case "hello":
                if true {
                    v1 += (v2 & v3)
                    if (1 == 102) {
                        v27 := GenS2{F0: -1, F1: (fi26 * -1)}
                        _ = v27
                    } else {
                        var v28 int64 = -9223372036854775808
                        var v29 uint8 = 0
                        var v30 uint8 = (v29 % (((v29 * (v29 + v29))) | 1))
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    if (v21 != v21) {
                        v31 := GenS3{F0: v1, F1: 42}
                        v32 := GenS2{F0: -406, F1: fi26}
                        if false {
                            v33 := GenS1{F0: ""}
                            _ = v33
                        }
                        _ = v31
                        _ = v32
                    } else {
                        v34 := []int{(fi26 & (fi26 & fi26)), fi26}
                        switch fi26 {
                        case 7:
                            if true {
                                var v35 float32 = float32((float64(0.0) - 3.14159))
                                _ = v35
                            } else {
                                switch "a" {
                                case "hello":
                                    switch (fi26 >> (((v21 | v21)) & 63)) {
                                    case -752:
                                        if (!(false || false)) {
                                            for fi36 := 0; fi36 < 4; fi36++ {
                                                fmt.Println(v3)
                                                v37 := GenS4{F0: -686, F1: 18446744073709551615, F2: 127}
                                                if ((true && ((((v21 / (((v21 ^ (v21 | (v21 - (v21 >> ((v21) & 63)))))) | 1)) <= 50) || false) && (false && (false && (-9223372036854775808 < 482))))) && ((1.5 * -2.25) < 1.5)) {
                                                    var v38 uint32 = (v1 & (v1 % ((296) | 1)))
                                                    v39 := GenS5{F0: 2, F1: 65535, F2: (true || (42 <= -2147483648))}
                                                    v34 = append(v34, 242)
                                                    v40 := GenS4{F0: -999, F1: 42, F2: 42}
                                                    _ = v38
                                                    _ = v39
                                                    _ = v40
                                                } else {
                                                    v41 := GenS4{F0: 382, F1: 671, F2: 127}
                                                    _ = len(v34)
                                                    _ = v41
                                                }
                                                _ = fi36
                                                _ = v37
                                            }
                                            var v42 int16 = genFn1((fi26 - fi26))
                                            _ = v42
                                        } else {
                                            fmt.Println(fi26, fi26)
                                            var v43 uint = v21
                                            if false {
                                                v44 := GenS1{F0: "foo bar"}
                                                _ = v44
                                            } else {
                                                v2 -= v2
                                                fmt.Println(v21, v2, fi26)
                                                fmt.Println(v2, fi26)
                                            }
                                            _ = v43
                                        }
                                    case 543:
                                        v21 = v21
                                    case 557:
                                        if false {
                                            var v45 string = "\n\t"
                                            for fi46 := 0; fi46 < 14; fi46++ {
                                                v45 = v45
                                                if false {
                                                    if (!false) {
                                                        v2 -= 0
                                                    }
                                                    var v47 uint8 = 128
                                                    v48 := GenS2{F0: -667, F1: (fi46 - (fi26 << (((v21 & (v21 / ((42) | 1)))) & 63)))}
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    v49 := GenS4{F0: 1, F1: 18446744073709551615, F2: 2}
                                                    _ = v49
                                                }
                                                _ = fi46
                                            }
                                            var v50 int64 = 42
                                            _ = v45
                                            _ = v50
                                        } else {
                                            v3 = 2147483647
                                        }
                                    }
                                default:
                                    for fi51 := 0; fi51 < 8; fi51++ {
                                        v2 -= genFn0()
                                        _ = fi51
                                    }
                                }
                                var v52 uint8 = 2
                                v2 += (v2 & 782)
                                _ = v52
                            }
                        case 127:
                            if len(v34) > 0 {
                                v34[0] = (fi26 ^ fi26)
                            }
                        }
                        _ = v34
                    }
                }
            case "":
                var v53 uint16 = 42
                _ = v53
            case "a":
                v54 := GenS6{F0: v21, F1: 1}
                _ = v54
            }
            _ = fi26
        }
    }
    var v55 uint32 = (^(v1))
    var v56 uint64 = 18446744073709551615
    var v57 int32 = -2147483648
    for fi58 := 0; fi58 < 3; fi58++ {
        v59 := GenS1{F0: "foo bar"}
        v60 := GenS7{F0: -1, F1: 255}
        _ = fi58
        _ = v59
        _ = v60
    }
    var v61 uint32 = (^(v55))
    v55 = v2
    for fi62 := 0; fi62 < 19; fi62++ {
        v56 -= 1
        v61 = 848
        var v63 int64 = 9223372036854775807
        _ = fi62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v55
    _ = v56
    _ = v57
    _ = v61
}
