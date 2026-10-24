package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint8) uint32 {
    for fi4 := 0; fi4 < 17; fi4++ {
        p0 -= p0
        p0 -= (p0 * (p0 % (((^(p0))) | 1)))
        for fi5 := 0; fi5 < 3; fi5++ {
            p0 = p0
            var v6 int64 = 144
            v7 := []int{fi4, (fi4 % ((-57) | 1)), (fi5 % ((fi4) | 1)), (^(fi5))}
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    v8 := GenS1{F0: 9223372036854775807, F1: 100}
    _ = p0
    _ = v8
    return 2147483647
}
func genFn1() uint32 {
    var v9 float64 = 0.0
    v10 := []int{0, 42, 387, 0}
    if false {
        var v11 float64 = (v9 * (v9 / v9))
        _ = v11
    } else {
        if true {
            for fi12 := 0; fi12 < 10; fi12++ {
                fmt.Println(fi12)
                fmt.Println(fi12, fi12)
                _ = fi12
            }
            v13 := GenS2{F0: -1, F1: 0, F2: 248}
            if true {
                var v14 string = "\n\t"
                switch 2147483647 {
                case 276:
                    v15 := GenS2{F0: 32767, F1: 0, F2: 984}
                    _ = v15
                case 907:
                    fmt.Println(v14, v14)
                case -9223372036854775808:
                    if true {
                        v9 += 0.1
                        v16 := GenS2{F0: -32768, F1: 4294967295, F2: 845}
                        _ = v16
                    } else {
                        if false {
                            switch 0 {
                            case 4294967295:
                                var v17 uint32 = 1
                                _ = v17
                            case -9223372036854775808:
                                if false {
                                    fmt.Println(v9, v14)
                                    var v18 int8 = 1
                                    _ = v18
                                } else {
                                    var v19 uint = 675
                                    switch (v14 + ((v14 + ((((v14 + v14) + v14) + ((v14 + (v14 + v14)) + v14)) + v14)) + "a")) {
                                    case "":
                                        v14 = v14
                                    case "a":
                                        switch v14 {
                                        case "":
                                            var v20 int = -226
                                            _ = v20
                                        case "a":
                                            if (true || (762 != 788)) {
                                                var v21 int16 = -1
                                                _ = v21
                                            }
                                        case "foo bar":
                                            fmt.Println(v19, v14)
                                        }
                                    }
                                    v22 := GenS2{F0: 32767, F1: 4294967295, F2: 2}
                                    _ = v19
                                    _ = v22
                                }
                            default:
                                var v23 uint8 = 2
                                _ = v23
                            }
                            v14 = (v14 + v14)
                        } else {
                            v24 := GenS2{F0: 729, F1: 585, F2: 0}
                            _ = len(v10)
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 12; fi25++ {
                            v10 = append(v10, fi25)
                            if false {
                                v26 := GenS1{F0: 1, F1: -128}
                                _ = v26
                            }
                            _ = fi25
                        }
                        switch (v14 + ("hello" + "\n\t")) {
                        case "\n\t":
                            fmt.Println(v9)
                        case "":
                            v14 = "\n\t"
                        case "a":
                            v14 = ("\n\t" + v14)
                        }
                    }
                }
                if false {
                    var v27 string = v14
                    fmt.Println(v27, v9, v27)
                    _ = v27
                } else {
                    fmt.Println(v14, v14)
                    switch 222 {
                    case 42:
                        v14 = (v14 + v14)
                    case -526:
                        v10 = append(v10, -2147483648)
                    case 2147483647:
                        v28 := GenS2{F0: 0, F1: 2, F2: 876}
                        _ = v28
                    default:
                        v9 = (v9 * 1.5)
                    }
                }
                _ = v14
            }
            _ = v13
        }
        if false {
            var v29 uint32 = 2
            switch -2147483648 {
            case 128:
                v9 += v9
            case -733:
                if len(v10) > 0 {
                    v10[0] = -108
                }
            case 32768:
                for fi30 := 0; fi30 < 1; fi30++ {
                    var v31 uint32 = (v29 << ((4294967295) & 63))
                    if ((!false) && false) {
                        var v32 float64 = ((0.0 / (v9 + 1e10)) - v9)
                        var v33 string = "foo bar"
                        _ = v32
                        _ = v33
                    }
                    var v34 uint = 2
                    _ = fi30
                    _ = v31
                    _ = v34
                }
            default:
                _ = len(v10)
            }
            v9 += (v9 + ((v9 / float64(3.14159)) / (float64((3.14159 + 0.0)) * 1e10)))
            _ = v29
        }
    }
    v35 := GenS3{F0: "hello"}
    _ = v10
    _ = v35
    _ = v9
    return 2
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint = 0
    v3 := []string{}
    v36 := GenS1{F0: 1, F1: -128}
    var v37 bool = (-721 <= 4294967295)
    for fi38 := 0; fi38 < 14; fi38++ {
        v1 = v1
        v3 = append(v3, "foo bar")
        _ = fi38
    }
    switch 9223372036854775807 {
    case 636:
        switch ("a" + "hello") {
        case "\n\t":
            fmt.Println(v37, v37, v1)
        case "a":
            v39 := GenS4{F0: 2147483647, F1: 102}
            _ = v39
        case "":
            for fi40 := 0; fi40 < 4; fi40++ {
                fmt.Println(v1, v1, v2)
                _ = fi40
            }
        }
    case -959:
        v1 -= v1
    default:
        var v41 float64 = (0.1 * 1.5)
        _ = v41
    }
    var v42 bool = (65535 > 32767)
    switch -488 {
    case 2147483647:
        fmt.Println(v42, v1, v42)
    case 32767:
        _ = len(v3)
    default:
        var v43 bool = ((("hello" + "") + ("" + "")) >= "foo bar")
        _ = v43
    }
    v44 := GenS3{F0: "\n\t"}
    v1 += ((v1 - (((1e10 + -2.25) * v1) / float32(1e10))) * (1e10 - (float32((float64((v1 / v1)) + -2.25)) + v1)))
    v2 = (v2 + v2)
    if (!false) {
        fmt.Println(v1, v1)
    }
    v2 += (v2 << (((v2 & (^(v2)))) & 63))
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v42
    _ = v44
}
