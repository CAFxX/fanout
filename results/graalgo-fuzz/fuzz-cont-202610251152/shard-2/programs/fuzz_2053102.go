package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int64, p1 bool, p2 int64) uint64 {
    fmt.Println(p1)
    if p1 {
        p0 = p2
        for fi5 := 0; fi5 < 4; fi5++ {
            if false {
                p1 = ((!true) || true)
                switch fi5 {
                case 775:
                    var v6 int = fi5
                    _ = v6
                case 32767:
                    v7 := GenS1{F0: (^(fi5)), F1: 0}
                    _ = v7
                case 9223372036854775807:
                    fmt.Println(p2, p2)
                }
            }
            v8 := []int{(-(fi5)), -906, fi5}
            p1 = (-800 != 1000)
            _ = fi5
            _ = v8
        }
    }
    var v9 float64 = 3.14159
    p1 = ((!p1) && false)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 0
}

func main() {
    var v1 int8 = -128
    v2 := make(map[int]int)
    var v3 uint64 = 18446744073709551615
    v4 := []int{852}
    v1 += (v1 / (((v1 + (v1 + (v1 & -125)))) | 1))
    if (42 > 127) {
        if false {
            v10 := GenS1{F0: 466, F1: 65}
            v3 -= v3
            fmt.Println(v1, v1)
            _ = v10
        }
        var v11 uint16 = 42
        _ = v11
    } else {
        var v12 uint16 = 255
        var v13 float64 = 1e10
        _ = v12
        _ = v13
    }
    fmt.Println(v3, v1)
    var v14 uint16 = 159
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 int = fi15
        if true {
            for fi17 := 0; fi17 < 19; fi17++ {
                var v18 int8 = (v1 | -128)
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 10; fi19++ {
                v1 = v1
                v20 := GenS1{F0: (fi19 * fi19), F1: 42}
                _ = fi19
                _ = v20
            }
        } else {
            if true {
                v21 := GenS1{F0: 127, F1: 0}
                switch fi15 {
                case 127:
                    switch (fi15 / ((v16) | 1)) {
                    case -1:
                        for fi22 := 0; fi22 < 3; fi22++ {
                            v23 := GenS1{F0: 4294967295, F1: 141}
                            if true {
                                var v24 uint = 140
                                _ = v24
                            } else {
                                var v25 int32 = -43
                                switch "\n\t" {
                                case "foo bar":
                                    v26 := GenS1{F0: v16, F1: 255}
                                    _ = v26
                                case "":
                                    var v27 uint32 = 4294967295
                                    _ = v27
                                case "a":
                                    v3 = 4294967295
                                default:
                                    var v28 string = "foo bar"
                                    _ = v28
                                }
                                _ = v25
                            }
                            v1 = (v1 / ((v1) | 1))
                            _ = fi22
                            _ = v23
                        }
                    case -2147483648:
                        v29 := GenS2{F0: fi15, F1: -1, F2: v14}
                        _ = v29
                    case 9223372036854775807:
                        v30 := GenS2{F0: (v16 << ((105) & 63)), F1: -32768, F2: 1}
                        _ = v30
                    }
                default:
                    if (0.1 <= (1e10 / -2.25)) {
                        v2[(^(fi15))] = v16
                        for fi31 := 0; fi31 < 2; fi31++ {
                            v32 := GenS1{F0: (v16 + fi31), F1: 4294967295}
                            v33 := GenS2{F0: (fi31 * int(v1)), F1: -32768, F2: (^(v14))}
                            var v34 int32 = 597
                            _ = fi31
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                    }
                }
                for fi35 := 0; fi35 < 15; fi35++ {
                    v16 -= (fi35 / ((v16) | 1))
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v36 := GenS3{F0: uint8(v14)}
                    var v37 float64 = (((((3.14159 - 0.1) / 1.0) - 0.1) * (0.1 * -2.25)) / float64((1e10 - -2.25)))
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                _ = v21
            }
        }
        _ = fi15
        _ = v16
    }
    if true {
        v38 := GenS3{F0: uint8(v3)}
        _ = v38
    } else {
        v39 := v2[-485]
        v40 := GenS3{F0: 2}
        var v41 int16 = 1
        _ = v39
        _ = v40
        _ = v41
    }
    v3 -= v3
    var v42 uint16 = (^(v14))
    var v43 string = "foo bar"
    for fi44 := 0; fi44 < 6; fi44++ {
        v45 := v2[fi44]
        v46 := GenS1{F0: 32768, F1: 707}
        _ = fi44
        _ = v45
        _ = v46
    }
    if true {
        for fi47 := 0; fi47 < 10; fi47++ {
            v1 = v1
            _ = fi47
        }
        var v48 int = 740
        var v49 uint = 618
        _ = v48
        _ = v49
    } else {
        if ((v43 + "foo bar") > "a") {
            var v50 uint = 584
            v51 := GenS2{F0: 499, F1: -1, F2: (v14 + (v14 % ((v14) | 1)))}
            _ = v50
            _ = v51
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
}
