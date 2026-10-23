package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int16; F1 uint64; F2 int8 }
func genFn0(p0 uint8, p1 int16, p2 int8) int16 {
    var v7 int16 = (p1 & (p1 * 0))
    if false {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 bool = ((-2147483648 > -2147483648) && false)
            var v10 float32 = (-2.25 * 0.1)
            _ = fi8
            _ = v10
            _ = v9
        }
    } else {
        v11 := GenS1{F0: 900, F1: 2147483647}
        var v12 uint = 42
        p2 = (^(p2))
        _ = v11
        _ = v12
    }
    var v13 uint8 = (p0 >> ((4294967295) & 63))
    switch (("" + "") + ("hello" + ("\n\t" + ""))) {
    case "foo bar":
        v14 := GenS1{F0: 466, F1: 691}
        _ = v14
    case "":
        p0 = p0
    }
    switch -1 {
    case 32767:
        v7 = (-(p1))
    default:
        v15 := GenS2{F0: "", F1: 2}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v7
    return v7
}
func genFn1(p0 uint, p1 uint64, p2 uint32) uint16 {
    fmt.Println(p2, p0)
    var v16 uint8 = 2
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 float32 = float32(((0.0 / 0.1) - 1e10))
        if false {
            p1 = (p1 ^ p1)
            v19 := []int{fi17, fi17, fi17}
            p0 -= p0
            _ = v19
        }
        var v20 uint8 = v16
        _ = fi17
        _ = v18
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 633
}

func main() {
    var v1 float64 = (0.1 - 3.14159)
    var v2 uint8 = 121
    v3 := []uint32{65535, 530, 2, 248}
    var v4 uint8 = (v2 - v2)
    var v5 uint16 = 32767
    v6 := make(map[int]uint32)
    v21 := GenS2{F0: (((("hello" + ("\n\t" + "\n\t")) + "foo bar") + "foo bar") + "\n\t"), F1: 42}
    var v22 string = "a"
    _ = len(v6)
    _ = len(v6)
    var v23 float32 = ((0.1 * -2.25) + 1e10)
    v2 -= 42
    var v24 uint = 4294967295
    v4 = v2
    if false {
        for fi25 := 0; fi25 < 19; fi25++ {
            for fi26 := 0; fi26 < 8; fi26++ {
                switch v22 {
                case "hello":
                    var v27 int8 = -31
                    _ = v27
                }
                if (0 > (fi25 | 4294967295)) {
                    switch v22 {
                    case "":
                        v5 = 185
                    case "foo bar":
                        v5 -= (v5 ^ (^(v5)))
                        var v28 string = v22
                        v29 := GenS2{F0: (v22 + "\n\t"), F1: 4294967295}
                        _ = v28
                        _ = v29
                    }
                    if (!(9223372036854775807 == 0)) {
                        v30 := GenS2{F0: "", F1: 894}
                        switch v22 {
                        case "\n\t":
                            fmt.Println(v4, v1, v22)
                        default:
                            switch fi26 {
                            case -998:
                                var v31 int8 = -56
                                _ = v31
                            case 9223372036854775807:
                                var v32 int32 = -591
                                _ = v32
                            }
                        }
                        _ = v30
                    } else {
                        v33 := GenS1{F0: (v5 * (v5 / (((v5 - uint16(v24))) | 1))), F1: 42}
                        v22 = "a"
                        if (!false) {
                            if (int32(v24) > -1) {
                                fmt.Println(v22, v24, v22)
                            }
                        }
                        _ = v33
                    }
                } else {
                    var v34 uint = v24
                    v35 := GenS2{F0: v22, F1: 572}
                    _ = v34
                    _ = v35
                }
                _ = fi26
            }
            _ = fi25
        }
    }
    v2 += v4
    if true {
        v3 = append(v3, 952)
    } else {
        switch 32768 {
        case 155:
            switch "hello" {
            case "":
                if (true || (15 != 0)) {
                    v22 = "foo bar"
                    var v36 uint64 = 42
                    var v37 float32 = -2.25
                    _ = v36
                    _ = v37
                }
            case "\n\t":
                if len(v3) > 0 {
                    v3[0] = 386
                }
            case "hello":
                fmt.Println(v4, v1, v22)
            default:
                var v38 uint16 = (v5 << ((v24) & 63))
                _ = v38
            }
        case -9223372036854775808:
            var v39 float32 = (0.0 / float32(v1))
            _ = v39
        case -1:
            if (!(118 >= -32768)) {
                for fi40 := 0; fi40 < 19; fi40++ {
                    var v41 int16 = 1
                    var v42 string = (((((((v22 + v22) + "\n\t") + (v22 + ("\n\t" + ("\n\t" + ((("" + v22) + (("\n\t" + (v22 + v22)) + v22)) + (v22 + v22)))))) + "hello") + v22) + v22) + "\n\t")
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            } else {
                fmt.Println(v2, v22, v4)
                for fi43 := 0; fi43 < 10; fi43++ {
                    v44 := v6[4294967295]
                    v45 := GenS1{F0: genFn1(v24, 284, (^(v44))), F1: v44}
                    v1 = (1e10 - ((v1 / v1) + 3.14159))
                    _ = fi43
                    _ = v44
                    _ = v45
                }
            }
        default:
            for fi46 := 0; fi46 < 2; fi46++ {
                v47 := GenS3{F0: 32767, F1: 2, F2: 100}
                fmt.Println(v5, v5)
                var v48 int8 = 52
                _ = fi46
                _ = v47
                _ = v48
            }
        }
        v22 = (("a" + v22) + v22)
    }
    v4 = (v4 / ((v4) | 1))
    if (!((^(v4)) <= v2)) {
        switch -2147483648 {
        case -867:
            var v49 uint32 = 0
            _ = v49
        case -120:
            v4 -= (^(v4))
        case -23:
            v50 := GenS2{F0: v22, F1: 2}
            _ = v50
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        var v51 int64 = 9223372036854775807
        if len(v3) > 0 {
            v3[0] = 819
        }
        _ = v51
    }
    for fi52 := 0; fi52 < 14; fi52++ {
        fmt.Println(v23, v24, fi52)
        v5 -= (v5 + uint16(v4))
        v4 = 42
        _ = fi52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
