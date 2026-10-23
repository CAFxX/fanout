package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int }
func genFn0(p0 int) uint64 {
    var v5 uint64 = 42
    var v6 int16 = 0
    v6 = v6
    p0 = p0
    fmt.Println(v5, v5, p0)
    _ = p0
    _ = v5
    _ = v6
    return 1
}
func genFn1() string {
    v7 := GenS1{F0: 822}
    var v8 int16 = -292
    v8 += v8
    _ = v7
    _ = v8
    return "hello"
}

func main() {
    v1 := []string{"hello", "", "a"}
    var v2 uint8 = 255
    var v3 string = "a"
    var v4 int8 = -113
    fmt.Println(v4, v3, v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v9 int8 = 0
    var v10 uint = 42
    var v11 uint16 = 42
    fmt.Println(v10, v9, v9)
    for fi12 := 0; fi12 < 2; fi12++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v13 uint64 = 0
        var v14 string = v3
        _ = fi12
        _ = v13
        _ = v14
    }
    if false {
        if false {
            v11 = v11
        }
        v1 = append(v1, v3)
        if (!(!((int32(v11) > 2147483647) || false))) {
            v15 := GenS2{F0: (^(v10)), F1: v10}
            if false {
                if true {
                    if false {
                        v16 := GenS1{F0: 65535}
                        _ = v16
                    } else {
                        if (!(3.14159 != 3.14159)) {
                            switch -190 {
                            case 42:
                                v1 = append(v1, (((v3 + "hello") + v3) + v3))
                            case 128:
                                switch ("" + "foo bar") {
                                case "a":
                                    for fi17 := 0; fi17 < 3; fi17++ {
                                        var v18 uint8 = 127
                                        fmt.Println(v9, v9)
                                        _ = fi17
                                        _ = v18
                                    }
                                case "":
                                    fmt.Println(v3)
                                default:
                                    v4 = (-(v4))
                                }
                            default:
                                if false {
                                    fmt.Println(v9, v11)
                                }
                            }
                            if len(v1) > 0 {
                                v1[0] = v3
                            }
                        }
                    }
                    var v19 int16 = int16(v9)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v19
                }
                var v20 int8 = 42
                _ = v20
            }
            if true {
                if (!false) {
                    if (false && true) {
                        var v21 uint8 = (v2 % (((v2 + v2)) | 1))
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v22 float64 = (-2.25 - ((1.5 * (1.5 / 1.0)) / 3.14159))
                        _ = v21
                        _ = v22
                    } else {
                        v1 = append(v1, (("hello" + "a") + v3))
                        v11 += 42
                    }
                    var v23 uint16 = (v11 / ((v11) | 1))
                    var v24 uint16 = v11
                    _ = v23
                    _ = v24
                }
                v3 = (v3 + v3)
            }
            _ = v15
        } else {
            if (65535 != 2) {
                if true {
                    var v25 float32 = 3.14159
                    _ = v25
                } else {
                    switch 42 {
                    case -30:
                        v26 := GenS1{F0: 326}
                        _ = v26
                    default:
                        for fi27 := 0; fi27 < 10; fi27++ {
                            var v28 string = "hello"
                            _ = fi27
                            _ = v28
                        }
                    }
                    for fi29 := 0; fi29 < 8; fi29++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch fi29 {
                        case 364:
                            var v30 uint = (v10 ^ (v10 << ((v10) & 63)))
                            _ = v30
                        case 453:
                            for fi31 := 0; fi31 < 15; fi31++ {
                                fmt.Println(v3, fi31)
                                _ = fi31
                            }
                        default:
                            if len(v1) > 0 {
                                v1[0] = v3
                            }
                        }
                        _ = fi29
                    }
                    var v32 uint32 = 2
                    _ = v32
                }
            } else {
                v9 = v4
                if len(v1) > 0 {
                    v1[0] = (("hello" + v3) + "\n\t")
                }
            }
            v4 = int8(v2)
            for fi33 := 0; fi33 < 10; fi33++ {
                v34 := GenS2{F0: 65535, F1: v10}
                v35 := GenS1{F0: uint32(v11)}
                _ = fi33
                _ = v34
                _ = v35
            }
        }
    }
    v36 := GenS1{F0: 65535}
    for fi37 := 0; fi37 < 12; fi37++ {
        v38 := GenS3{F0: v11}
        _ = fi37
        _ = v38
    }
    var v39 int8 = v9
    v40 := GenS4{F0: 127}
    var v41 uint64 = genFn0(61)
    var v42 int8 = 0
    v1 = append(v1, v3)
    var v43 uint16 = v11
    v43 += 580
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v9
}
