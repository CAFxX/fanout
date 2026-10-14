package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint; F1 int16; F2 int16 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int64, p1 int32, p2 uint) int16 {
    v4 := []int{128}
    var v5 int64 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 995
}
func genFn1() int32 {
    var v6 uint32 = 2147483648
    var v7 uint32 = v6
    var v8 uint32 = 65535
    v6 = 4294967295
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float64 = -2.25
    v3 := make(map[string]uint32)
    if (1.5 <= float32(1.5)) {
        var v9 string = "a"
        if (!(!true)) {
            for fi10 := 0; fi10 < 13; fi10++ {
                switch -9223372036854775808 {
                case -9223372036854775808:
                    v1 = 70
                default:
                    var v11 int8 = 0
                    _ = v11
                }
                _ = fi10
            }
            v2 = 3.14159
        } else {
            var v12 bool = true
            _ = v12
        }
        _ = v9
    } else {
        v2 = v2
        v2 = ((((0.0 / (v2 * (v2 * ((((3.14159 / 0.1) / (v2 - v2)) / v2) / v2)))) * 0.1) / 0.1) * float64((((((float32(v2) / 3.14159) / 0.1) - 0.0) + (1.5 / (0.1 * 0.1))) + 0.1)))
    }
    switch "\n\t" {
    case "foo bar":
        for fi13 := 0; fi13 < 3; fi13++ {
            v14 := []int{}
            v2 -= float64(-2.25)
            _ = fi13
            _ = v14
        }
    case "\n\t":
        var v15 int32 = -635
        _ = v15
    case "":
        switch ((("foo bar" + "foo bar") + "") + "a") {
        case "\n\t":
            var v16 string = (((("\n\t" + "") + ("" + "hello")) + "") + ("\n\t" + (("foo bar" + ("" + "foo bar")) + "\n\t")))
            _ = v16
        case "":
            var v17 uint = 0
            _ = v17
        case "foo bar":
            var v18 int16 = 1000
            _ = v18
        }
    }
    v19 := []int{32768, 9223372036854775807}
    _ = len(v3)
    v2 -= 1.5
    fmt.Println(v1)
    _ = len(v3)
    var v20 int64 = -1
    fmt.Println(v2, v20, v2)
    var v21 uint32 = 65535
    fmt.Println(v20, v1)
    for fi22 := 0; fi22 < 8; fi22++ {
        v2 -= v2
        v1 = 18446744073709551615
        delete(v3, "foo bar")
        _ = fi22
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        if (((false || (("\n\t" + "foo bar") <= "a")) && true) || (1 <= -1)) {
            v24 := GenS1{F0: v20, F1: -1}
            var v25 float64 = v2
            for fi26 := 0; fi26 < 3; fi26++ {
                v27 := GenS2{F0: v21, F1: (v21 / ((42) | 1)), F2: (^(v21))}
                var v28 int8 = -55
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = v24
            _ = v25
        } else {
            if true {
                fmt.Println(v21, fi23, v21)
                if (844 >= 867) {
                    v20 -= v20
                }
                var v29 uint16 = 0
                _ = v29
            } else {
                for fi30 := 0; fi30 < 8; fi30++ {
                    v21 += (v21 + 42)
                    _ = fi30
                }
                v21 -= 2147483647
                for fi31 := 0; fi31 < 16; fi31++ {
                    if (!true) {
                        v3[("" + "hello")] = v21
                        v21 = v21
                    } else {
                        v2 -= 0.1
                        var v32 bool = true
                        _ = v32
                    }
                    var v33 int = (^(fi31))
                    if true {
                        var v34 bool = (65535 >= 2)
                        v35 := GenS1{F0: (v20 & (v20 % (((v20 / ((v20) | 1))) | 1))), F1: 1}
                        _ = v34
                        _ = v35
                    }
                    v36 := v3["\n\t"]
                    switch "foo bar" {
                    case "foo bar":
                        delete(v3, ("" + ""))
                    case "":
                        v20 -= (v20 << ((602) & 63))
                    case "\n\t":
                        if (false || true) {
                            v37 := GenS3{F0: 1, F1: 32767, F2: 42}
                            if (2 > 1) {
                                var v38 uint64 = v1
                                if true {
                                    fmt.Println(v1, fi31, v33)
                                    v38 += v38
                                    v36 = (^(v21))
                                    var v39 uint64 = (v1 | v1)
                                    if len(v19) > 0 {
                                        v19[0] = (v33 & fi23)
                                    }
                                    _ = v39
                                }
                                _ = v38
                            } else {
                                if false {
                                    v40 := v3["foo bar"]
                                    v19 = append(v19, 32768)
                                    _ = v40
                                }
                                v21 = v36
                                switch 128 {
                                case 32768:
                                    switch v33 {
                                    case -763:
                                        v21 -= (v36 & 546)
                                    case 17:
                                        var v41 int = -9223372036854775808
                                        _ = v41
                                    }
                                default:
                                    if ((fi23 | (fi31 & 9223372036854775807)) < 32768) {
                                        v42 := GenS4{F0: -1}
                                        var v43 uint8 = 2
                                        var v44 int16 = -245
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    } else {
                                        var v45 float64 = ((v2 * -2.25) - float64(float32(v2)))
                                        _ = v45
                                    }
                                }
                                var v46 uint = 255
                                if false {
                                    if (1.5 <= 0.0) {
                                        var v47 float32 = 0.1
                                        v1 += v1
                                        v46 = (v46 - (v46 % (((^(v46))) | 1)))
                                        _ = v47
                                    }
                                } else {
                                    var v48 uint8 = 255
                                    var v49 float64 = v2
                                    _ = v48
                                    _ = v49
                                }
                                _ = v46
                            }
                            v20 = -9223372036854775808
                            _ = v37
                        } else {
                            var v50 float32 = 1.5
                            _ = v50
                        }
                    }
                    var v51 int32 = -136
                    v19 = append(v19, (fi31 * (^(fi23))))
                    _ = fi31
                    _ = v33
                    _ = v36
                    _ = v51
                }
            }
            delete(v3, ("foo bar" + "hello"))
            for fi52 := 0; fi52 < 6; fi52++ {
                var v53 string = ((("foo bar" + "") + ((("" + "a") + ("" + "foo bar")) + (("a" + "foo bar") + ("\n\t" + (("hello" + "") + "hello"))))) + (((("foo bar" + "hello") + "\n\t") + "") + ""))
                _ = fi52
                _ = v53
            }
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
