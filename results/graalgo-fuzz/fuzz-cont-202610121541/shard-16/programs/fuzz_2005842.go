package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 bool }

func main() {
    var v1 int64 = -1
    var v2 float32 = ((1.5 - (-2.25 * 3.14159)) + 1e10)
    var v3 uint = 1
    var v4 int16 = -32768
    if (!false) {
        v5 := []int{127, 379, 32768, 2147483647}
        if len(v5) > 0 {
            v5[0] = int(v1)
        }
        var v6 uint8 = 2
        _ = v5
        _ = v6
    }
    if (65535 <= 2147483647) {
        v7 := GenS1{F0: -825, F1: (^(v4)), F2: (false && (v2 > v2))}
        if false {
            v8 := GenS1{F0: v4, F1: (v4 ^ v4), F2: false}
            v9 := GenS1{F0: (v4 % ((42) | 1)), F1: (v4 | v4), F2: true}
            _ = v8
            _ = v9
        } else {
            var v10 int16 = v4
            v11 := []int{192}
            if (-1 <= 0) {
                v12 := GenS2{F0: (v1 - 0)}
                fmt.Println(v1)
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = v12
            } else {
                for fi13 := 0; fi13 < 7; fi13++ {
                    v14 := GenS1{F0: (v10 + 1000), F1: v10, F2: (-1 > 42)}
                    v11 = append(v11, 128)
                    if (!(3.14159 == (3.14159 * 0.0))) {
                        switch "\n\t" {
                        case "hello":
                            var v15 int = int(v10)
                            _ = v15
                        }
                    }
                    _ = fi13
                    _ = v14
                }
                v16 := GenS2{F0: v1}
                _ = v16
            }
            _ = v10
            _ = v11
        }
        _ = v7
    }
    v17 := []int{}
    fmt.Println(v4, v1)
    fmt.Println(v4, v2, v4)
    fmt.Println(v4, v4)
    if (v2 >= (0.0 + (v2 / v2))) {
        fmt.Println(v1)
        fmt.Println(v4, v2)
        switch ("\n\t" + "hello") {
        case "\n\t":
            v4 = 0
        case "hello":
            if true {
                var v18 uint16 = 65535
                v19 := GenS1{F0: -239, F1: (v4 - (v4 ^ 1)), F2: false}
                if true {
                    var v20 float32 = (v2 * 0.1)
                    v4 = (v4 + v4)
                    if true {
                        if false {
                            v3 -= 310
                            for fi21 := 0; fi21 < 19; fi21++ {
                                v2 += v20
                                _ = fi21
                            }
                            v17 = append(v17, 128)
                        } else {
                            if true {
                                var v22 string = (("" + "hello") + ("\n\t" + ("a" + "")))
                                for fi23 := 0; fi23 < 1; fi23++ {
                                    var v24 float32 = 0.1
                                    if true {
                                        for fi25 := 0; fi25 < 14; fi25++ {
                                            var v26 float32 = (v24 - v20)
                                            var v27 bool = true
                                            _ = fi25
                                            _ = v26
                                            _ = v27
                                        }
                                        if (true || ((((!(-943 <= 1)) || ((v22 + "\n\t") <= "a")) && (v2 >= v24)) || (true && true))) {
                                            if len(v17) > 0 {
                                                v17[0] = 32768
                                            }
                                            if (2 != 1) {
                                                if false {
                                                    for fi28 := 0; fi28 < 3; fi28++ {
                                                        switch v22 {
                                                        case "a":
                                                            v2 = (1.5 - v2)
                                                        case "foo bar":
                                                            v4 = v4
                                                        case "hello":
                                                            fmt.Println(v18)
                                                        }
                                                        switch v22 {
                                                        case "hello":
                                                            if (v1 >= v1) {
                                                                if true {
                                                                    v29 := GenS2{F0: (v1 << ((255) & 63))}
                                                                    v4 += (v4 % (((v4 - (v4 % ((v4) | 1)))) | 1))
                                                                    _ = v29
                                                                } else {
                                                                    var v30 bool = true
                                                                    switch v22 {
                                                                    case "":
                                                                        var v31 uint64 = 613
                                                                        _ = v31
                                                                    case "foo bar":
                                                                        var v32 string = (("foo bar" + ("\n\t" + v22)) + v22)
                                                                        _ = v32
                                                                    default:
                                                                        var v33 string = "\n\t"
                                                                        _ = v33
                                                                    }
                                                                    v34 := GenS2{F0: v1}
                                                                    _ = v30
                                                                    _ = v34
                                                                }
                                                                v17 = append(v17, (fi28 % ((411) | 1)))
                                                                if true {
                                                                    v17 = append(v17, int(v18))
                                                                    var v35 float32 = (v24 * ((v24 + (float32(3.14159) / v2)) + v24))
                                                                    var v36 float32 = (0.0 * v2)
                                                                    _ = v35
                                                                    _ = v36
                                                                }
                                                                var v37 float32 = (v20 * ((v24 - (0.1 + -2.25)) * 0.0))
                                                                _ = v37
                                                            }
                                                        }
                                                        v38 := GenS2{F0: -1}
                                                        switch (v22 + v22) {
                                                        case "a":
                                                            var v39 uint32 = 4294967295
                                                            _ = v39
                                                        }
                                                        v1 = v1
                                                        _ = fi28
                                                        _ = v38
                                                    }
                                                } else {
                                                    for fi40 := 0; fi40 < 16; fi40++ {
                                                        v3 = (v3 - 560)
                                                        fmt.Println(v1, fi40, v1)
                                                        _ = len(v17)
                                                        _ = fi40
                                                    }
                                                    v18 = 745
                                                    if len(v17) >= 2 {
                                                        v17 = v17[1:len(v17)]
                                                    }
                                                }
                                            }
                                            fmt.Println(v4, v3)
                                        }
                                    } else {
                                        v41 := GenS3{F0: 1, F1: 1, F2: (v18 - v18)}
                                        if false {
                                            v20 += v2
                                            v2 = v20
                                            v22 = v22
                                        } else {
                                            for fi42 := 0; fi42 < 3; fi42++ {
                                                switch v22 {
                                                case "a":
                                                    var v43 uint64 = 4294967295
                                                    _ = v43
                                                case "hello":
                                                    switch -100 {
                                                    case 127:
                                                        switch (fi23 >> ((255) & 63)) {
                                                        case 127:
                                                            v1 -= v1
                                                        default:
                                                            fmt.Println(v24, v20)
                                                        }
                                                    }
                                                }
                                                _ = fi42
                                            }
                                            var v44 int16 = 1
                                            fmt.Println(v1)
                                            _ = v44
                                        }
                                        v45 := GenS4{F0: 65535, F1: 495}
                                        _ = v41
                                        _ = v45
                                    }
                                    if len(v17) > 0 {
                                        v17[0] = int(v18)
                                    }
                                    _ = fi23
                                    _ = v24
                                }
                                _ = v22
                            }
                            v46 := GenS1{F0: v4, F1: 1, F2: (!false)}
                            _ = v46
                        }
                        v20 += (1.5 - (v20 * 1.5))
                    } else {
                        switch 32768 {
                        case 339:
                            var v47 int16 = (v4 >> ((v3) & 63))
                            _ = v47
                        case -1:
                            fmt.Println(v18, v3)
                        case 2147483647:
                            fmt.Println(v20, v3, v1)
                        }
                    }
                    _ = v20
                }
                _ = v18
                _ = v19
            }
        case "foo bar":
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
        default:
            for fi48 := 0; fi48 < 12; fi48++ {
                var v49 int = -62
                _ = fi48
                _ = v49
            }
        }
    }
    var v50 uint = v3
    v4 -= 1
    v3 -= 255
    v51 := GenS2{F0: v1}
    for fi52 := 0; fi52 < 14; fi52++ {
        var v53 int64 = (-(v1))
        v54 := GenS5{F0: true}
        fmt.Println(v1, v3)
        _ = fi52
        _ = v53
        _ = v54
    }
    fmt.Println(v50)
    for fi55 := 0; fi55 < 14; fi55++ {
        v50 -= (v50 | 2)
        _ = fi55
    }
    if false {
        if len(v17) > 0 {
            v17[0] = -1
        }
    } else {
        v2 -= -2.25
        fmt.Println(v1, v2, v3)
        v50 = (v50 % ((v50) | 1))
    }
    v17 = append(v17, int(v4))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
}
