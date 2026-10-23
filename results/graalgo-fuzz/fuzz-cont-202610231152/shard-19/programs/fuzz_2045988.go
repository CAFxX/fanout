package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int8 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS4 struct { F0 int16; F1 string; F2 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint16, p1 uint8, p2 int8) int {
    v5 := GenS1{F0: 470, F1: 42, F2: p2}
    var v6 int64 = -9223372036854775808
    fmt.Println(p0, p0)
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int8 = p2
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{65535, 42, 966, 364}
    var v3 int16 = 32767
    var v4 int = -328
    var v9 float32 = 1.5
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := GenS1{F0: 1, F1: (fi11 + (fi11 + fi11)), F2: 127}
            v4 += (fi11 >> ((42) & 63))
            _ = len(v2)
            var v13 bool = (!false)
            v14 := v1[("a" + "foo bar")]
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        var v15 bool = (1 > -1)
        var v16 int64 = -1
        _ = fi10
        _ = v15
        _ = v16
    }
    var v17 int = int(v3)
    for fi18 := 0; fi18 < 16; fi18++ {
        v17 = 748
        switch ("\n\t" + "hello") {
        case "foo bar":
            if ((float64(v9) + 1e10) != (1.5 / 1e10)) {
                fmt.Println(v3)
                for fi19 := 0; fi19 < 5; fi19++ {
                    v1[((("a" + ("hello" + ("hello" + ("hello" + "a")))) + "hello") + (("" + "hello") + "a"))] = ("hello" + "\n\t")
                    v3 += 1
                    v20 := GenS2{F0: 1, F1: 4294967295}
                    _ = fi19
                    _ = v20
                }
                var v21 float64 = 0.1
                _ = v21
            } else {
                switch "foo bar" {
                case "\n\t":
                    v9 += v9
                }
                for fi22 := 0; fi22 < 13; fi22++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi22
                }
                var v23 uint8 = 146
                _ = v23
            }
        case "a":
            if ((false || true) || (true || (-422 >= -900))) {
                for fi24 := 0; fi24 < 2; fi24++ {
                    if (!((fi24 < (v4 ^ fi18)) && false)) {
                        v3 += v3
                        v9 += float32(-2.25)
                        v25 := GenS2{F0: 18446744073709551615, F1: 2}
                        _ = v25
                    } else {
                        if ((^(v3)) != (v3 | v3)) {
                            v3 = v3
                            _ = len(v2)
                            var v26 uint8 = 110
                            _ = v26
                        }
                    }
                    _ = fi24
                }
            } else {
                switch (fi18 + v17) {
                case 1:
                    for fi27 := 0; fi27 < 2; fi27++ {
                        var v28 string = (((("hello" + ("a" + (("foo bar" + ("foo bar" + "foo bar")) + ((("a" + "foo bar") + "a") + (("a" + "foo bar") + "foo bar"))))) + ("foo bar" + ("a" + "\n\t"))) + "foo bar") + "a")
                        if false {
                            fmt.Println(v28, v28)
                        }
                        v29 := GenS3{F0: 11, F1: 0, F2: 56}
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                default:
                    v30 := GenS1{F0: 4294967295, F1: fi18, F2: -1}
                    _ = v30
                }
                switch "" {
                case "":
                    delete(v1, (("foo bar" + "\n\t") + ("\n\t" + ((("hello" + "foo bar") + "foo bar") + "a"))))
                case "foo bar":
                    if false {
                        fmt.Println(v4)
                        fmt.Println(v4)
                        for fi31 := 0; fi31 < 8; fi31++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v32 float32 = float32(float64(v9))
                            _ = fi31
                            _ = v32
                        }
                    }
                case "a":
                    switch (v17 | genFn0(65535, 2, -128)) {
                    case -1:
                        var v33 int64 = 0
                        _ = v33
                    case -9223372036854775808:
                        delete(v1, "\n\t")
                    case 2147483647:
                        var v34 float64 = (((1e10 - (3.14159 + (0.0 + (float64((0.1 / v9)) - (float64(-2.25) * 1e10))))) - (((1.5 / 1.0) + 1.5) / (0.1 - (1e10 * 1.5)))) - (float64(0.1) - 0.1))
                        _ = v34
                    default:
                        for fi35 := 0; fi35 < 3; fi35++ {
                            if ((!false) && false) {
                                _ = len(v2)
                                var v36 int = 128
                                switch (-(v4)) {
                                case 42:
                                    v37 := GenS4{F0: v3, F1: (("foo bar" + "\n\t") + ("foo bar" + "hello")), F2: -757}
                                    _ = v37
                                case 1:
                                    _ = len(v1)
                                case 2147483647:
                                    if (1 == 0) {
                                        v9 = v9
                                    }
                                }
                                _ = v36
                            }
                            _ = fi35
                        }
                    }
                default:
                    var v38 uint = 783
                    _ = v38
                }
                v4 = 2147483647
            }
        default:
            var v39 uint32 = 599
            _ = v39
        }
        if ((false || false) || false) {
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
        }
        _ = fi18
    }
    v9 = (v9 * (v9 + (((v9 / 1.5) + float32(0.0)) + v9)))
    v40 := GenS5{F0: 1}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v41 float32 = v9
    switch "foo bar" {
    case "foo bar":
        v41 -= -2.25
    default:
        v42 := GenS4{F0: 711, F1: "", F2: -1}
        _ = v42
    }
    var v43 string = "hello"
    _ = len(v1)
    v44 := GenS6{F0: 0, F1: 65535}
    v1[v43] = (v43 + "hello")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v43
    _ = v44
    _ = v9
}
