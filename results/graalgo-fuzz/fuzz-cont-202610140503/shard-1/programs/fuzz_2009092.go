package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 bool }
type GenS2 struct { F0 int64; F1 uint16; F2 int }
type GenS3 struct { F0 uint8 }
func genFn0() uint32 {
    var v7 float32 = 0.1
    v7 += v7
    var v8 float64 = -2.25
    _ = v7
    _ = v8
    return 2147483648
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    var v4 float64 = 0.1
    var v5 string = ("hello" + ("\n\t" + "foo bar"))
    var v6 float32 = 0.1
    switch v5 {
    case "hello":
        v9 := GenS1{F0: 127, F1: 1, F2: (97 < 24)}
        _ = v9
    default:
        fmt.Println(v6, v4)
    }
    v6 = 1e10
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 uint32 = 0
        _ = fi10
        _ = v11
    }
    v12 := []int{1, 32768, -265, 51}
    if false {
        v5 = "a"
        if false {
            var v13 uint16 = 2
            switch 1 {
            case -534:
                for fi14 := 0; fi14 < 3; fi14++ {
                    v6 += (v6 * float32(((v4 - v4) / (v4 + v4))))
                    switch int(v13) {
                    case 0:
                        if true {
                            for fi15 := 0; fi15 < 7; fi15++ {
                                v16 := GenS2{F0: 0, F1: 941, F2: fi14}
                                var v17 int16 = 1
                                var v18 uint16 = v13
                                v6 += (v6 - (0.1 + v6))
                                switch fi15 {
                                case 4294967295:
                                    if false {
                                        var v19 int64 = 93
                                        _ = v19
                                    }
                                case 683:
                                    fmt.Println(v4)
                                case 635:
                                    v3[""] = (fi15 | 1)
                                }
                                v20 := GenS1{F0: 13, F1: 128, F2: true}
                                var v21 int = (fi14 * -1)
                                fmt.Println(v18, v21, v18)
                                _ = fi15
                                _ = v16
                                _ = v17
                                _ = v18
                                _ = v20
                                _ = v21
                            }
                            v22 := GenS1{F0: 969, F1: uint8(v13), F2: true}
                            _ = v22
                        } else {
                            v6 += -2.25
                        }
                    default:
                        for fi23 := 0; fi23 < 17; fi23++ {
                            fmt.Println(v13, v4)
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            var v24 bool = (true || true)
                            var v25 bool = (3.14159 != v6)
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                    }
                    _ = fi14
                }
            default:
                if false {
                    var v26 uint16 = (v13 | (v13 << ((623) & 63)))
                    var v27 int8 = 0
                    v28 := GenS2{F0: -1, F1: v26, F2: 128}
                    _ = v26
                    _ = v27
                    _ = v28
                } else {
                    var v29 uint8 = 0
                    _ = v29
                }
            }
            v6 = 3.14159
            _ = v13
        } else {
            var v30 uint = 1
            _ = v30
        }
        v5 = v5
    }
    v6 -= (((((1e10 - 1.5) / (-2.25 + 3.14159)) - v6) + 1.5) / (v6 + float32(float64(v6))))
    v31 := GenS1{F0: 0, F1: 227, F2: (2 < 255)}
    if (((-2.25 * v6) == ((v6 + v6) + v6)) && false) {
        var v32 int16 = 32767
        if false {
            fmt.Println(v1, v1)
            v33 := v3[("" + (((("hello" + v5) + v5) + "foo bar") + v5))]
            _ = v33
        }
        _ = v32
    }
    if ((float64(1.5) >= v4) && (!(false && false))) {
        var v34 int64 = 0
        for fi35 := 0; fi35 < 2; fi35++ {
            switch fi35 {
            case -2147483648:
                var v36 int8 = 51
                _ = v36
            case 0:
                v37 := v3[(("foo bar" + v5) + v5)]
                _ = v37
            case 2147483647:
                v6 += 3.14159
            }
            var v38 int64 = v34
            _ = fi35
            _ = v38
        }
        v34 = (^(v34))
        _ = v34
    }
    if (false || (1 == 42)) {
        var v39 float32 = 0.0
        switch "foo bar" {
        case "hello":
            fmt.Println(v4, v1, v5)
        case "a":
            v4 = v4
        }
        for fi40 := 0; fi40 < 14; fi40++ {
            switch "" {
            case "":
                v39 = (v39 + v6)
            case "hello":
                v41 := GenS1{F0: -9223372036854775808, F1: 42, F2: (false || ((fi40 + -143) > 32768))}
                _ = v41
            }
            fmt.Println(v6, v5)
            for fi42 := 0; fi42 < 14; fi42++ {
                if false {
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                } else {
                    var v43 uint32 = 2147483648
                    _ = v43
                }
                var v44 uint32 = 1
                v45 := GenS1{F0: fi40, F1: 127, F2: ((255 != 0) && true)}
                v1 = (v1 | (v1 * 1))
                var v46 int8 = 2
                _ = fi42
                _ = v44
                _ = v45
                _ = v46
            }
            _ = fi40
        }
        _ = v39
    } else {
        v47 := GenS3{F0: 128}
        if (true || (((0 <= 127) && false) || (412 < 42))) {
            if (535 > -1) {
                var v48 bool = false
                var v49 float32 = 0.0
                _ = v48
                _ = v49
            } else {
                var v50 uint8 = 250
                v1 = v1
                _ = v50
            }
            _ = len(v12)
        }
        v51 := GenS1{F0: -86, F1: 246, F2: false}
        _ = v47
        _ = v51
    }
    switch 32768 {
    case 814:
        v1 = (v1 + v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
