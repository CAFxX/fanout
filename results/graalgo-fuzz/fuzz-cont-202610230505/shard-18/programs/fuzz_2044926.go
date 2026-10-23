package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 string }
type GenS2 struct { F0 uint; F1 uint64; F2 uint }
type GenS3 struct { F0 uint16; F1 string }
func genFn0(p0 uint64, p1 uint16) uint {
    p0 += p0
    fmt.Println(p0)
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := []int{}
        fmt.Println(p1)
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := []int{fi10, fi9, fi10, (fi9 % ((fi9) | 1))}
            _ = fi10
            _ = v11
        }
        p1 = (^(p1))
        _ = fi9
    }
    _ = p0
    _ = p1
    return 4294967295
}
func genFn1() int32 {
    var v12 uint8 = 127
    fmt.Println(v12)
    _ = v12
    return -2147483648
}

func main() {
    var v1 int32 = -2147483648
    var v2 int = 0
    var v3 int64 = -33
    v4 := make(map[uint32]string)
    var v5 int32 = -270
    v6 := []int{312, 411}
    var v13 float64 = 3.14159
    switch "foo bar" {
    case "":
        _ = len(v6)
    case "a":
        if false {
            if true {
                if (0 != 976) {
                    if true {
                        var v14 float32 = 1.5
                        delete(v4, 4294967295)
                        _ = v14
                    }
                    if (v5 <= 2147483647) {
                        var v15 bool = true
                        v4[2147483647] = "\n\t"
                        _ = v15
                    } else {
                        for fi16 := 0; fi16 < 20; fi16++ {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            var v17 uint32 = 128
                            v18 := GenS1{F0: (-(v2)), F1: -1, F2: ("" + "")}
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        v19 := GenS1{F0: v2, F1: 32767, F2: ("foo bar" + "foo bar")}
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 1; fi20++ {
                        fmt.Println(fi20, v1, v2)
                        _ = fi20
                    }
                } else {
                    for fi21 := 0; fi21 < 6; fi21++ {
                        var v22 int8 = 9
                        for fi23 := 0; fi23 < 5; fi23++ {
                            _ = len(v4)
                            _ = fi23
                        }
                        _ = fi21
                        _ = v22
                    }
                    v6 = append(v6, -9223372036854775808)
                    if (!false) {
                        v4[131] = ""
                    }
                }
                v24 := v4[65535]
                if (2147483647 > 298) {
                    switch "" {
                    case "a":
                        v13 -= ((float64(-2.25) + float64(0.1)) + ((((v13 + (float64(0.1) / ((v13 / (1e10 * ((v13 * 0.1) / (v13 / 1e10)))) / 1.5))) + (((v13 * v13) * v13) * (v13 - v13))) * ((v13 + ((1e10 / v13) * (float64(1.5) - v13))) + (v13 / ((v13 + v13) + float64((0.1 / -2.25)))))) / float64(((0.1 + 0.1) + 3.14159))))
                    case "foo bar":
                        var v25 uint64 = 42
                        _ = v25
                    case "hello":
                        v4[0] = ("\n\t" + "a")
                    default:
                        if true {
                            switch (v2 / ((v2) | 1)) {
                            case 136:
                                v24 = v24
                            case -2147483648:
                                if (v2 != v2) {
                                    v26 := GenS1{F0: (v2 ^ v2), F1: -228, F2: (v24 + (v24 + (((v24 + (v24 + v24)) + v24) + v24)))}
                                    var v27 float64 = v13
                                    _ = v26
                                    _ = v27
                                }
                            default:
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                        } else {
                            var v28 uint = genFn0(18446744073709551615, 32767)
                            if false {
                                v29 := GenS2{F0: (v28 + (^(v28))), F1: 118, F2: v28}
                                _ = v29
                            } else {
                                var v30 float32 = -2.25
                                fmt.Println(v30)
                                _ = v30
                            }
                            if ((0 != v2) || false) {
                                v1 += v1
                                v24 = "hello"
                            } else {
                                var v31 bool = (0 >= 42)
                                _ = v31
                            }
                            _ = v28
                        }
                    }
                    v1 = v5
                    var v32 int8 = 42
                    _ = v32
                } else {
                    v2 = v2
                    v33 := GenS3{F0: 255, F1: v24}
                    var v34 float32 = (-2.25 + ((3.14159 - float32((v13 / float64(1.5)))) + 1.5))
                    _ = v33
                    _ = v34
                }
                _ = v24
            }
            v35 := GenS2{F0: 18446744073709551615, F1: 1, F2: 2}
            _ = v35
        }
    case "\n\t":
        v2 = (-(v2))
    }
    var v36 uint64 = 1
    v5 -= genFn1()
    for fi37 := 0; fi37 < 12; fi37++ {
        v3 += (v3 & 42)
        v38 := GenS2{F0: 4294967295, F1: 302, F2: 4294967295}
        _ = fi37
        _ = v38
    }
    fmt.Println(v36)
    fmt.Println(v3)
    for fi39 := 0; fi39 < 14; fi39++ {
        v40 := GenS2{F0: genFn0((^(v36)), 2), F1: (v36 / ((42) | 1)), F2: 255}
        v41 := GenS1{F0: v2, F1: 319, F2: ""}
        _ = fi39
        _ = v40
        _ = v41
    }
    v2 = (v2 >> ((4294967295) & 63))
    v42 := GenS1{F0: v2, F1: 32767, F2: (("foo bar" + ("foo bar" + "a")) + (("\n\t" + "\n\t") + "hello"))}
    switch "hello" {
    case "hello":
        v43 := GenS1{F0: v2, F1: 195, F2: ("hello" + "\n\t")}
        _ = v43
    case "":
        v36 = v36
    case "\n\t":
        v44 := GenS1{F0: v2, F1: 1, F2: "\n\t"}
        _ = v44
    default:
        switch "foo bar" {
        case "foo bar":
            switch v2 {
            case -2147483648:
                v45 := GenS3{F0: 897, F1: ((("foo bar" + (("" + "hello") + "foo bar")) + "a") + ((("foo bar" + "\n\t") + (("" + (((((("foo bar" + "a") + ("a" + "foo bar")) + "") + "hello") + (("" + "foo bar") + ("\n\t" + ("foo bar" + "hello")))) + "hello")) + "hello")) + ("foo bar" + "a")))}
                _ = v45
            case 32768:
                switch "\n\t" {
                case "":
                    v46 := GenS1{F0: v2, F1: -445, F2: ("\n\t" + ("foo bar" + "a"))}
                    _ = v46
                default:
                    v5 = v5
                }
            case 2147483647:
                if (34 <= 2) {
                    var v47 int16 = 1
                    _ = v47
                } else {
                    var v48 int8 = 127
                    _ = v48
                }
            default:
                var v49 uint16 = 916
                _ = v49
            }
        case "":
            var v50 uint = genFn0(v36, 32767)
            _ = v50
        default:
            var v51 uint32 = 1
            _ = v51
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
