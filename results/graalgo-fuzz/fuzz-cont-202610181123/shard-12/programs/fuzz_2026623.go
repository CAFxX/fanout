package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32 }
func genFn0(p0 int, p1 int16, p2 int64) uint {
    v5 := GenS1{F0: 4294967295, F1: (p2 * p2)}
    v6 := []int{p0, (p0 + (p0 - 9223372036854775807)), -2147483648, (-(p0))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 6
}
func genFn1(p0 string, p1 uint64) uint32 {
    if false {
        v7 := []int{-1, 0}
        fmt.Println(p0, p0)
        p0 = p0
        _ = v7
    } else {
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: 411, F1: 515}
            v10 := []int{(fi8 >> ((964) & 63))}
            _ = fi8
            _ = v10
            _ = v9
        }
        fmt.Println(p1, p0)
        v11 := GenS1{F0: 2, F1: -1}
        _ = v11
    }
    if (!true) {
        var v12 int8 = 127
        v13 := []int{}
        var v14 int = 1
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = v15
    return v15
}

func main() {
    v1 := []uint32{1, 717, 2147483647}
    v2 := []string{"\n\t"}
    v3 := []string{("foo bar" + "\n\t"), "foo bar", ""}
    var v4 float64 = 1.5
    v4 += (3.14159 + 1.5)
    fmt.Println(v4)
    v3 = append(v3, "hello")
    var v16 int = 127
    var v17 uint8 = 136
    fmt.Println(v17, v17, v17)
    v18 := GenS1{F0: 255, F1: -1}
    var v19 float64 = (v4 * 0.0)
    v17 = 42
    if false {
        if true {
            var v20 int = v16
            _ = v20
        } else {
            v3 = append(v3, (("hello" + ("\n\t" + ("foo bar" + ("foo bar" + ("a" + ("" + ((("" + ("\n\t" + (((("\n\t" + "\n\t") + "foo bar") + "") + "hello"))) + "hello") + "foo bar"))))))) + (((("hello" + ("a" + "foo bar")) + "\n\t") + ("a" + "")) + "\n\t")))
            v21 := GenS2{F0: 0}
            v22 := GenS3{F0: 20, F1: 42, F2: -32768}
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            v4 = v4
            _ = fi23
        }
        switch -2147483648 {
        case 4294967295:
            if true {
                if (495 > 1) {
                    switch "" {
                    case "a":
                        if true {
                            v24 := GenS2{F0: 210}
                            v19 = float64((1.5 + 3.14159))
                            v2 = append(v2, "foo bar")
                            _ = v24
                        }
                    case "":
                        v19 -= v19
                    case "foo bar":
                        switch -9223372036854775808 {
                        case 0:
                            if false {
                                if (((v4 - (((v19 - 0.0) - v19) + 3.14159)) + v4) >= -2.25) {
                                    for fi25 := 0; fi25 < 15; fi25++ {
                                        v4 += float64(((((3.14159 - ((0.1 + 3.14159) / float32(-2.25))) * (-2.25 + 1.5)) * ((1e10 * ((1.5 - 3.14159) - (3.14159 - 0.0))) + ((((-2.25 / 1.0) * -2.25) * 1e10) / ((0.1 / 3.14159) * float32(0.1))))) / 0.1))
                                        v26 := GenS2{F0: 327}
                                        _ = fi25
                                        _ = v26
                                    }
                                    for fi27 := 0; fi27 < 10; fi27++ {
                                        var v28 uint = genFn0((v16 - (-(fi27))), 42, 1)
                                        var v29 int16 = -396
                                        _ = fi27
                                        _ = v28
                                        _ = v29
                                    }
                                    if true {
                                        if len(v2) > 0 {
                                            v2[0] = (("" + (((("hello" + "") + "foo bar") + (("a" + "\n\t") + "foo bar")) + ("\n\t" + "foo bar"))) + "foo bar")
                                        }
                                    }
                                } else {
                                    var v30 bool = true
                                    v30 = (((float32(((v4 * float64(0.0)) + (v19 - (float64(1e10) / ((v4 - 0.0) / (0.1 * ((v19 + v4) * 1e10))))))) + (float32(float64((0.1 - (3.14159 + (1.5 * 0.1))))) / float32((v19 - float64(0.1))))) / (float32(v4) + 1e10)) != (0.1 - 1.5))
                                    _ = v30
                                }
                                fmt.Println(v17, v4)
                            }
                        case 272:
                            if true {
                                for fi31 := 0; fi31 < 19; fi31++ {
                                    v17 = (v17 ^ (v17 / ((v17) | 1)))
                                    if (42 >= 42) {
                                        v19 -= -2.25
                                    }
                                    _ = fi31
                                }
                                for fi32 := 0; fi32 < 6; fi32++ {
                                    v33 := GenS3{F0: 1, F1: 0, F2: 1000}
                                    v34 := GenS4{F0: 1}
                                    _ = fi32
                                    _ = v33
                                    _ = v34
                                }
                            } else {
                                fmt.Println(v17)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v35 := GenS5{F0: -891}
                                _ = v35
                            }
                        case 32768:
                            switch (v16 % (((v16 - v16)) | 1)) {
                            case 128:
                                v36 := GenS4{F0: 514}
                                _ = v36
                            case 127:
                                if true {
                                    var v37 int32 = 1
                                    _ = v37
                                }
                            case 49:
                                var v38 int32 = 0
                                _ = v38
                            }
                        }
                    default:
                        v39 := GenS4{F0: 4294967295}
                        _ = v39
                    }
                    v40 := GenS2{F0: 2147483647}
                    _ = v40
                } else {
                    for fi41 := 0; fi41 < 6; fi41++ {
                        v4 = float64((1.5 + (0.1 - (1e10 + -2.25))))
                        for fi42 := 0; fi42 < 20; fi42++ {
                            var v43 bool = (!(!("" < "\n\t")))
                            v3 = append(v3, "hello")
                            _ = fi42
                            _ = v43
                        }
                        _ = fi41
                    }
                    for fi44 := 0; fi44 < 19; fi44++ {
                        for fi45 := 0; fi45 < 19; fi45++ {
                            v4 -= -2.25
                            _ = fi45
                        }
                        _ = fi44
                    }
                }
                v46 := GenS3{F0: 240, F1: 2, F2: 0}
                _ = v46
            }
        default:
            for fi47 := 0; fi47 < 11; fi47++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v16 += 1
                var v48 string = ("\n\t" + "foo bar")
                _ = fi47
                _ = v48
            }
        }
    }
    fmt.Println(v17, v16)
    if (1 <= 0) {
        if len(v3) > 0 {
            v3[0] = (("foo bar" + "foo bar") + "foo bar")
        }
    } else {
        v17 = v17
        for fi49 := 0; fi49 < 15; fi49++ {
            v16 += fi49
            _ = fi49
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
