package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 uint64 }
func genFn0(p0 int32) string {
    var v4 uint64 = 819
    p0 += (p0 | (p0 >> ((65535) & 63)))
    _ = p0
    _ = v4
    return ("\n\t" + ((("hello" + (("hello" + "foo bar") + (("" + ("\n\t" + "foo bar")) + (("foo bar" + (("foo bar" + "foo bar") + ("hello" + "foo bar"))) + "hello")))) + "a") + "a"))
}
func genFn1() uint32 {
    var v5 uint8 = 2
    if (42 >= 42) {
        v5 = v5
        var v6 int16 = -268
        _ = v6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        v5 -= (^(v5))
        for fi8 := 0; fi8 < 18; fi8++ {
            fmt.Println(v5, v5, fi8)
            v9 := []int{(fi7 ^ (^(fi8))), 32768, 1}
            v5 = 128
            if (true && true) {
                var v10 uint16 = 42
                _ = v10
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    switch "" {
    case "a":
        v11 := []int{2147483647, 127, -154}
        _ = v11
    case "":
        v12 := GenS1{F0: ("foo bar" + "hello")}
        _ = v12
    }
    v5 -= (v5 << ((4294967295) & 63))
    _ = v5
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 255
    v3 := []int{407, 0, 871, -152}
    switch 42 {
    case 128:
        if (-993 > 128) {
            var v13 int16 = 42
            _ = v13
        }
    }
    if true {
        delete(v1, genFn1())
        v14 := GenS1{F0: "foo bar"}
        _ = v14
    }
    var v15 uint32 = 249
    fmt.Println(v15)
    v15 = (v15 / ((2147483647) | 1))
    if false {
        var v16 uint64 = 2
        for fi17 := 0; fi17 < 14; fi17++ {
            v2 -= v2
            _ = fi17
        }
        _ = v16
    }
    _ = len(v3)
    var v18 uint16 = (v2 % (((^(v2))) | 1))
    var v19 int16 = 951
    v15 -= (v15 ^ (v15 % ((65535) | 1)))
    fmt.Println(v15, v2)
    v20 := GenS1{F0: "hello"}
    fmt.Println(v15, v15)
    var v21 int = 178
    var v22 int32 = 1
    var v23 float64 = 1.5
    if true {
        v23 -= (v23 + 1e10)
        v22 = (v22 | (-(v22)))
    } else {
        v24 := GenS2{F0: v15}
        if true {
            var v25 int8 = 100
            switch (v21 & v21) {
            case 130:
                for fi26 := 0; fi26 < 14; fi26++ {
                    for fi27 := 0; fi27 < 20; fi27++ {
                        v18 += v2
                        switch 307 {
                        case -9223372036854775808:
                            var v28 uint = 1
                            _ = v28
                        case -595:
                            switch fi27 {
                            case -9223372036854775808:
                                v29 := GenS3{F0: -2147483648}
                                _ = v29
                            default:
                                v1[(v15 | (v15 % (((v15 << ((808) & 63))) | 1)))] = "hello"
                            }
                        case 9223372036854775807:
                            v3 = append(v3, 1)
                        }
                        fmt.Println(fi27, fi26)
                        _ = fi27
                    }
                    if len(v3) > 0 {
                        v3[0] = fi26
                    }
                    var v30 int = v21
                    _ = fi26
                    _ = v30
                }
            case 127:
                v25 -= v25
            default:
                switch v21 {
                case -9223372036854775808:
                    if true {
                        v18 -= (^(v2))
                        var v31 uint32 = v15
                        if false {
                            v32 := GenS1{F0: ("a" + "")}
                            for fi33 := 0; fi33 < 20; fi33++ {
                                if (!true) {
                                    if false {
                                        var v34 float64 = v23
                                        v35 := GenS3{F0: v22}
                                        _ = v34
                                        _ = v35
                                    }
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    fmt.Println(v31, v2, v31)
                                    v36 := GenS4{F0: (v25 / ((1) | 1))}
                                    _ = v36
                                }
                                _ = fi33
                            }
                            switch (v21 & 127) {
                            case 2147483647:
                                var v37 int = (v21 * 850)
                                _ = v37
                            }
                            _ = v32
                        } else {
                            v19 = (v19 ^ v19)
                            delete(v1, (v31 ^ 0))
                        }
                        _ = v31
                    } else {
                        v2 = v18
                        if true {
                            for fi38 := 0; fi38 < 18; fi38++ {
                                for fi39 := 0; fi39 < 12; fi39++ {
                                    v25 -= (v25 | v25)
                                    v19 = int16(v15)
                                    _ = fi39
                                }
                                v25 = 104
                                switch "" {
                                case "hello":
                                    if false {
                                        var v40 int8 = -81
                                        v41 := GenS5{F0: 1, F1: 103}
                                        _ = v40
                                        _ = v41
                                    } else {
                                        v15 += (v15 % (((v15 / ((genFn1()) | 1))) | 1))
                                    }
                                }
                                _ = fi38
                            }
                            fmt.Println(v22)
                            var v42 int16 = v19
                            _ = v42
                        } else {
                            var v43 string = "hello"
                            _ = v43
                        }
                        v2 = 255
                    }
                case -431:
                    var v44 int = 820
                    _ = v44
                case -936:
                    for fi45 := 0; fi45 < 18; fi45++ {
                        if len(v3) > 0 {
                            v3[0] = v21
                        }
                        v46 := v1[v15]
                        _ = fi45
                        _ = v46
                    }
                }
            }
            var v47 uint16 = (v18 >> ((234) & 63))
            _ = v25
            _ = v47
        }
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
