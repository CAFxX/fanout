package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 int32; F1 int16; F2 int }
func genFn0(p0 int16, p1 uint8, p2 uint) uint8 {
    var v6 int64 = -559
    for fi7 := 0; fi7 < 6; fi7++ {
        p0 += p0
        for fi8 := 0; fi8 < 20; fi8++ {
            if (int32(fi7) < 2147483647) {
                v9 := []int{}
                _ = v9
            } else {
                var v10 uint64 = 42
                p1 = (p1 / ((p1) | 1))
                p1 = p1
                v11 := GenS1{F0: p2, F1: 65535}
                if true {
                    v12 := GenS1{F0: 42, F1: 0}
                    var v13 int32 = 42
                    _ = v12
                    _ = v13
                } else {
                    v14 := GenS2{F0: 18446744073709551615, F1: uint8(p2)}
                    _ = v14
                }
                _ = v10
                _ = v11
            }
            p2 += 4294967295
            v15 := GenS2{F0: 655, F1: (^(p1))}
            fmt.Println(p2)
            _ = fi8
            _ = v15
        }
        p0 += p0
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 255
}
func genFn1() uint {
    var v16 int16 = 42
    var v17 string = ("foo bar" + "\n\t")
    v18 := []int{-564, -2147483648}
    switch v17 {
    case "foo bar":
        for fi19 := 0; fi19 < 10; fi19++ {
            if true {
                v20 := GenS1{F0: 217, F1: 1}
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                if true {
                    if false {
                        if true {
                            _ = len(v18)
                            switch 32767 {
                            case -315:
                                var v21 int32 = 1
                                _ = v21
                            case 32768:
                                v22 := GenS3{F0: 61, F1: 135}
                                _ = v22
                            case -2147483648:
                                if false {
                                    v16 -= v16
                                    fmt.Println(v17, v17, v16)
                                    v16 = (v16 ^ v16)
                                } else {
                                    switch v17 {
                                    case "foo bar":
                                        var v23 int64 = -74
                                        _ = v23
                                    case "\n\t":
                                        var v24 uint16 = 65535
                                        _ = v24
                                    }
                                    for fi25 := 0; fi25 < 10; fi25++ {
                                        v16 = v16
                                        fmt.Println(fi25, v16, fi19)
                                        var v26 uint32 = 0
                                        if len(v18) > 0 {
                                            v18[0] = fi19
                                        }
                                        _ = fi25
                                        _ = v26
                                    }
                                }
                            default:
                                if true {
                                    if (360 != 536) {
                                        v17 = v17
                                        var v27 int8 = -128
                                        _ = v27
                                    } else {
                                        if ((!(2 <= 0)) && (!((v16 & v16) > 1))) {
                                            if len(v18) >= 2 {
                                                v18 = v18[1:len(v18)]
                                            }
                                        }
                                    }
                                } else {
                                    v17 = v17
                                    var v28 float64 = -2.25
                                    _ = v28
                                }
                            }
                            v16 = 118
                        } else {
                            var v29 float64 = (((0.1 - 1e10) * -2.25) / 1.0)
                            switch (^(fi19)) {
                            case 32768:
                                var v30 string = v17
                                _ = v30
                            case 2147483647:
                                v16 -= v16
                            case -312:
                                v31 := GenS4{F0: 0, F1: 113}
                                _ = v31
                            default:
                                var v32 uint = 984
                                _ = v32
                            }
                            _ = v29
                        }
                    }
                    if true {
                        _ = len(v18)
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                    }
                    var v33 int = (fi19 ^ (fi19 << ((481) & 63)))
                    _ = v33
                }
                _ = v20
            } else {
                fmt.Println(v16, v17, v16)
            }
            v17 = "a"
            if (!((v17 >= ("\n\t" + "\n\t")) || true)) {
                v34 := GenS3{F0: 42, F1: 1}
                _ = v34
            } else {
                v17 = "foo bar"
                if (!(128 > 0)) {
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    v17 = (("" + "foo bar") + ("foo bar" + ((v17 + "\n\t") + v17)))
                    v16 += -1
                } else {
                    var v35 float64 = (0.1 - float64(0.0))
                    _ = v35
                }
            }
            _ = fi19
        }
    case "a":
        var v36 uint8 = 255
        _ = v36
    default:
        fmt.Println(v16)
    }
    switch v17 {
    case "":
        v37 := GenS5{F0: 2147483647, F1: v16, F2: -9223372036854775808}
        _ = v37
    case "hello":
        if (!false) {
            fmt.Println(v17, v17)
        }
    default:
        for fi38 := 0; fi38 < 20; fi38++ {
            if len(v18) > 0 {
                v18[0] = fi38
            }
            v39 := GenS1{F0: 419, F1: 920}
            for fi40 := 0; fi40 < 6; fi40++ {
                var v41 int = (^(fi38))
                _ = fi40
                _ = v41
            }
            _ = fi38
            _ = v39
        }
    }
    _ = v16
    _ = v17
    _ = v18
    return 1
}

func main() {
    v1 := []string{((("" + "a") + "a") + ("a" + "")), (("foo bar" + ("hello" + "foo bar")) + (("foo bar" + ("hello" + "foo bar")) + "foo bar")), "", ("" + "a")}
    v2 := []string{"a"}
    v3 := []uint32{65535, 4294967295}
    var v4 uint16 = 1
    var v5 bool = (true && false)
    var v42 int32 = -2147483648
    fmt.Println(v4, v4)
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    var v43 int16 = -483
    var v44 int32 = v42
    if (127 == int8(v43)) {
        if len(v2) > 0 {
            v2[0] = ("a" + (("foo bar" + "\n\t") + "a"))
        }
    } else {
        for fi45 := 0; fi45 < 14; fi45++ {
            if false {
                var v46 string = "\n\t"
                if v5 {
                    v47 := GenS2{F0: 828, F1: 238}
                    var v48 uint64 = 1
                    _ = v47
                    _ = v48
                }
                _ = v46
            } else {
                fmt.Println(fi45, v44, fi45)
                var v49 uint32 = 2147483647
                var v50 uint16 = 1
                _ = v49
                _ = v50
            }
            v42 = v42
            _ = fi45
        }
    }
    if ((v5 || v5) && true) {
        v42 += (-(v44))
    } else {
        var v51 string = "hello"
        var v52 uint16 = (v4 >> ((948) & 63))
        _ = v51
        _ = v52
    }
    var v53 uint32 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v53
}
