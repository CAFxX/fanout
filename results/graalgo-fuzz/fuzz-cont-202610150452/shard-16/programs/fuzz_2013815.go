package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 int16 }
func genFn0(p0 int16, p1 int32, p2 uint8) uint64 {
    var v5 uint = 255
    if false {
        v6 := GenS1{F0: (p0 << ((599) & 63)), F1: 136}
        if (!(((2 != 715) && ((true && false) || false)) || (!true))) {
            var v7 int64 = 317
            if true {
                v8 := []int{}
                _ = v8
            }
            _ = v7
        } else {
            switch -1 {
            case 1:
                v9 := []int{}
                _ = v9
            case 4294967295:
                p2 += 0
            case 127:
                p1 = p1
            default:
                v10 := []int{}
                _ = v10
            }
            p2 = (p2 / ((1) | 1))
        }
        var v11 int32 = p1
        _ = v11
        _ = v6
    }
    v12 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v5
    return 464
}
func genFn1(p0 uint16, p1 string, p2 uint32) int {
    v13 := []int{4294967295, -619, 0}
    for fi14 := 0; fi14 < 3; fi14++ {
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 uint8 = 42
            p0 = (^(p0))
            var v17 uint8 = 166
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi14
    }
    v18 := GenS1{F0: -32768, F1: 128}
    var v19 int32 = -650
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v18
    _ = v19
    return 128
}

func main() {
    var v1 float64 = -2.25
    var v2 string = "foo bar"
    var v3 string = ((v2 + "foo bar") + (v2 + v2))
    v4 := []string{v2}
    for fi20 := 0; fi20 < 17; fi20++ {
        v21 := GenS2{F0: 1, F1: -1}
        _ = fi20
        _ = v21
    }
    if (16 > 1) {
        fmt.Println(v3, v1, v3)
    } else {
        v3 = v2
        v3 = v3
    }
    v22 := GenS2{F0: 0, F1: 1}
    _ = len(v4)
    for fi23 := 0; fi23 < 5; fi23++ {
        switch fi23 {
        case 32767:
            v3 = v2
        case 2147483647:
            switch v3 {
            case "hello":
                v3 = "foo bar"
            case "a":
                v24 := GenS3{F0: 0, F1: -2147483648}
                _ = v24
            default:
                var v25 uint = 1
                _ = v25
            }
        default:
            if len(v4) > 0 {
                v4[0] = v3
            }
        }
        switch (fi23 | (fi23 >> ((495) & 63))) {
        case 698:
            v26 := GenS3{F0: 1, F1: 2147483647}
            _ = v26
        case -1:
            _ = len(v4)
        case -9223372036854775808:
            fmt.Println(fi23)
        }
        _ = fi23
    }
    v27 := GenS4{F0: 2}
    v4 = append(v4, ((((v3 + (v3 + v2)) + v2) + "a") + ("hello" + v2)))
    if true {
        switch "hello" {
        case "":
            if false {
                for fi28 := 0; fi28 < 1; fi28++ {
                    v3 = ((v2 + v3) + v3)
                    fmt.Println(v1, v2)
                    _ = fi28
                }
                var v29 int64 = 9223372036854775807
                _ = len(v4)
                _ = v29
            } else {
                v30 := GenS2{F0: 1000, F1: -475}
                _ = v30
            }
        case "foo bar":
            if true {
                var v31 int64 = 0
                _ = v31
            }
        case "a":
            if false {
                _ = len(v4)
                var v32 uint8 = 128
                if (true && (-1 <= 280)) {
                    v2 = (v2 + ("" + "foo bar"))
                    for fi33 := 0; fi33 < 3; fi33++ {
                        switch (fi33 / (((fi33 << ((4294967295) & 63))) | 1)) {
                        case 710:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        _ = fi33
                    }
                    for fi34 := 0; fi34 < 4; fi34++ {
                        if ((false || (654 <= 42)) && ((-309 != 32767) && false)) {
                            for fi35 := 0; fi35 < 10; fi35++ {
                                if true {
                                    v1 = v1
                                    if ((v1 - 0.1) < (((0.1 + 0.0) + ((1e10 / float64((1.5 * 1.5))) / ((v1 + (0.0 + (float64(0.1) * 1.5))) - float64((3.14159 + -2.25))))) / float64((-2.25 * float32((v1 * (v1 * (float64((3.14159 + 1.5)) + v1)))))))) {
                                        v36 := GenS5{F0: 255, F1: v32}
                                        _ = v36
                                    }
                                } else {
                                    switch (("foo bar" + v3) + (v3 + v3)) {
                                    case "hello":
                                        var v37 bool = true
                                        _ = v37
                                    case "\n\t":
                                        fmt.Println(fi34, v1, fi34)
                                    case "":
                                        v38 := GenS6{F0: -560, F1: 100}
                                        _ = v38
                                    }
                                }
                                if true {
                                    var v39 int = fi34
                                    v1 = float64(1.5)
                                    v32 -= v32
                                    _ = v39
                                } else {
                                    var v40 uint = 65535
                                    fmt.Println(v1)
                                    _ = v40
                                }
                                fmt.Println(fi35)
                                _ = fi35
                            }
                            fmt.Println(v1)
                            v2 = ((v2 + (v3 + "foo bar")) + ("\n\t" + ((v3 + (v2 + "hello")) + (v2 + ("\n\t" + ("a" + "hello"))))))
                        } else {
                            switch (fi34 + -779) {
                            case 0:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            default:
                                var v41 float32 = ((3.14159 * (3.14159 - 1e10)) - float32(v1))
                                _ = v41
                            }
                            var v42 string = (v2 + v3)
                            v32 = (v32 | v32)
                            _ = v42
                        }
                        var v43 int16 = 42
                        if true {
                            fmt.Println(v3, v2, v3)
                        } else {
                            v2 = "foo bar"
                            if (((float64((1e10 + (((1.5 / 0.1) * (1.5 / 0.1)) * 1e10))) / (v1 * v1)) <= (v1 * v1)) || false) {
                                var v44 bool = false
                                _ = v44
                            } else {
                                var v45 float32 = (((0.1 - ((((3.14159 * ((0.0 * 0.1) / -2.25)) + 1e10) * 1.5) * (3.14159 - (((float32((v1 / float64(float32(((((((v1 - (v1 * -2.25)) / 1.0) - v1) * (0.0 / 1.5)) * float64(-2.25)) * (v1 + 1e10)))))) + -2.25) / 0.1) + (0.0 - 1e10))))) + (float32((v1 / float64(3.14159))) - 0.1)) / 1.0)
                                _ = v45
                            }
                            for fi46 := 0; fi46 < 1; fi46++ {
                                if ((1 < 955) && (v1 < v1)) {
                                    switch v2 {
                                    case "":
                                        v47 := GenS6{F0: 2147483647, F1: -107}
                                        _ = v47
                                    case "hello":
                                        v48 := GenS2{F0: v43, F1: 525}
                                        _ = v48
                                    case "a":
                                        var v49 float32 = 0.0
                                        _ = v49
                                    default:
                                        if true {
                                            var v50 int = -2147483648
                                            _ = v50
                                        }
                                    }
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v4 = append(v4, v3)
                                }
                                switch v2 {
                                case "a":
                                    if false {
                                        v4 = append(v4, "hello")
                                        v51 := GenS5{F0: 1, F1: (v32 % (((^(v32))) | 1))}
                                        if (2147483647 == 2147483647) {
                                            var v52 float64 = v1
                                            v53 := GenS7{F0: v43}
                                            _ = len(v4)
                                            _ = v52
                                            _ = v53
                                        } else {
                                            fmt.Println(v3, v2, v2)
                                            v4 = append(v4, v2)
                                            if (-197 != -1) {
                                                if false {
                                                    switch "hello" {
                                                    case "hello":
                                                        if (42 == 232) {
                                                            switch "" {
                                                            case "\n\t":
                                                                v32 = 0
                                                            default:
                                                                var v54 uint64 = 395
                                                                v32 = 42
                                                                _ = v54
                                                            }
                                                        }
                                                    case "":
                                                        var v55 bool = (564 > 2)
                                                        _ = v55
                                                    }
                                                } else {
                                                    if (!true) {
                                                        var v56 uint = 42
                                                        _ = v56
                                                    }
                                                }
                                            } else {
                                                if len(v4) > 0 {
                                                    v4[0] = "\n\t"
                                                }
                                            }
                                            v3 = (v2 + (v2 + v3))
                                            var v57 int = 4294967295
                                            _ = v57
                                        }
                                        _ = v51
                                    } else {
                                        v43 = (v43 << ((336) & 63))
                                    }
                                case "\n\t":
                                    v43 += (v43 << ((2) & 63))
                                case "":
                                    var v58 uint64 = 42
                                    _ = v58
                                }
                                _ = fi46
                            }
                        }
                        _ = fi34
                        _ = v43
                    }
                } else {
                    _ = len(v4)
                }
                _ = v32
            } else {
                for fi59 := 0; fi59 < 18; fi59++ {
                    v1 += (v1 * ((-2.25 + -2.25) / ((1.5 - float64(0.0)) + v1)))
                    _ = fi59
                }
                v2 = (v3 + (("\n\t" + ("\n\t" + v3)) + ("a" + "\n\t")))
                for fi60 := 0; fi60 < 2; fi60++ {
                    v61 := GenS4{F0: 65535}
                    switch (fi60 + 9223372036854775807) {
                    case 724:
                        switch v2 {
                        case "a":
                            var v62 int = fi60
                            _ = v62
                        }
                    case 159:
                        var v63 uint = 255
                        _ = v63
                    }
                    _ = fi60
                    _ = v61
                }
            }
        }
    }
    if len(v4) > 0 {
        v4[0] = (v3 + "\n\t")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v4
}
