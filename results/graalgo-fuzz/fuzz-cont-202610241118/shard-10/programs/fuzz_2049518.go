package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 int16; F2 uint }
type GenS4 struct { F0 int16 }
func genFn0(p0 int32) uint8 {
    switch ("a" + "foo bar") {
    case "":
        for fi7 := 0; fi7 < 7; fi7++ {
            for fi8 := 0; fi8 < 20; fi8++ {
                fmt.Println(p0, fi7, p0)
                p0 -= (p0 / ((p0) | 1))
                p0 = p0
                _ = fi8
            }
            _ = fi7
        }
    case "a":
        switch -808 {
        case -2147483648:
            if (!false) {
                for fi9 := 0; fi9 < 3; fi9++ {
                    v10 := GenS1{F0: 17, F1: 712}
                    _ = fi9
                    _ = v10
                }
                v11 := GenS1{F0: 192, F1: 42}
                _ = v11
            } else {
                v12 := GenS2{F0: p0}
                _ = v12
            }
        case 538:
            switch -88 {
            case 465:
                for fi13 := 0; fi13 < 6; fi13++ {
                    var v14 int = fi13
                    v15 := []int{v14}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            }
        case 296:
            if true {
                var v16 int8 = 127
                fmt.Println(v16, p0)
                _ = v16
            } else {
                switch "" {
                case "":
                    for fi17 := 0; fi17 < 17; fi17++ {
                        switch fi17 {
                        case 128:
                            p0 -= p0
                        }
                        if (((0.1 - 1.5) - -2.25) > (0.1 + ((0.1 * 3.14159) * 0.1))) {
                            var v18 int8 = 109
                            _ = v18
                        } else {
                            if true {
                                if (2 <= 2147483647) {
                                    p0 += (p0 & (p0 | p0))
                                    for fi19 := 0; fi19 < 14; fi19++ {
                                        p0 = int32(fi17)
                                        p0 += p0
                                        var v20 uint16 = 1
                                        _ = fi19
                                        _ = v20
                                    }
                                }
                                for fi21 := 0; fi21 < 17; fi21++ {
                                    p0 = p0
                                    _ = fi21
                                }
                            } else {
                                if (662 != 1) {
                                    v22 := GenS3{F0: 2, F1: -459, F2: 4294967295}
                                    _ = v22
                                } else {
                                    switch fi17 {
                                    case 9223372036854775807:
                                        switch (fi17 | 32768) {
                                        case -1:
                                            v23 := []int{fi17, (fi17 - fi17)}
                                            _ = v23
                                        case 2147483647:
                                            var v24 bool = (int16(p0) == 42)
                                            _ = v24
                                        case 971:
                                            if (false || (p0 != p0)) {
                                                if (893 != 705) {
                                                    var v25 int = -926
                                                    var v26 int8 = -128
                                                    _ = v25
                                                    _ = v26
                                                } else {
                                                    v27 := []int{fi17, 127, fi17}
                                                    v28 := GenS3{F0: 255, F1: 669, F2: 860}
                                                    p0 = -574
                                                    _ = v27
                                                    _ = v28
                                                }
                                            } else {
                                                for fi29 := 0; fi29 < 14; fi29++ {
                                                    var v30 int = fi17
                                                    _ = fi29
                                                    _ = v30
                                                }
                                            }
                                        }
                                    }
                                    p0 -= (^(p0))
                                }
                                p0 = (p0 % (((p0 - (p0 & (-(p0))))) | 1))
                                for fi31 := 0; fi31 < 12; fi31++ {
                                    p0 += -436
                                    var v32 int32 = p0
                                    _ = fi31
                                    _ = v32
                                }
                            }
                            if true {
                                var v33 string = "hello"
                                var v34 string = (v33 + "hello")
                                v35 := GenS1{F0: 127, F1: 328}
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                        }
                        fmt.Println(fi17, p0)
                        _ = fi17
                    }
                }
            }
        default:
            v36 := GenS4{F0: 1}
            _ = v36
        }
    case "foo bar":
        var v37 int64 = -9223372036854775808
        _ = v37
    default:
        var v38 float64 = -2.25
        _ = v38
    }
    p0 += -1
    p0 += p0
    _ = p0
    return 1
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 42
    var v3 int8 = 42
    var v4 int = 0
    var v5 float64 = 1e10
    var v6 int64 = -386
    v39 := []int{0, (v4 + v4), 42}
    v2 = v2
    fmt.Println(v6, v4)
    v6 -= (^(v6))
    v40 := GenS3{F0: 4294967295, F1: 32767, F2: 255}
    for fi41 := 0; fi41 < 10; fi41++ {
        fmt.Println(v6, v4, v6)
        v1[("hello" + "foo bar")] = (fi41 * (fi41 - v4))
        v5 = 3.14159
        _ = fi41
    }
    fmt.Println(v5, v2)
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
