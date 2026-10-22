package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 bool; F1 bool; F2 string }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 uint; F1 uint16; F2 uint }
type GenS7 struct { F0 int16 }
func genFn0(p0 uint, p1 string, p2 uint64) int {
    switch p1 {
    case "\n\t":
        v5 := []int{}
        _ = v5
    case "a":
        if (float64(1.5) > 0.1) {
            var v6 float64 = 1e10
            fmt.Println(p1)
            _ = v6
        }
    case "hello":
        p1 = p1
    default:
        for fi7 := 0; fi7 < 17; fi7++ {
            switch "foo bar" {
            case "\n\t":
                for fi8 := 0; fi8 < 16; fi8++ {
                    fmt.Println(p1, p0, fi8)
                    switch p1 {
                    case "":
                        fmt.Println(p1, fi8, fi7)
                    }
                    p2 = p2
                    _ = fi8
                }
            }
            _ = fi7
        }
    }
    v9 := []int{32768, -1, 32767}
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 float64 = 0.1
            v12 += -2.25
            switch p1 {
            case "":
                var v13 int = (fi10 << ((p0) & 63))
                _ = v13
            case "hello":
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            }
            _ = fi11
            _ = v12
        }
        p1 = "foo bar"
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 1
}
func genFn1() bool {
    for fi14 := 0; fi14 < 2; fi14++ {
        switch ("\n\t" + "foo bar") {
        case "\n\t":
            for fi15 := 0; fi15 < 10; fi15++ {
                v16 := GenS1{F0: 931, F1: 988}
                var v17 string = "\n\t"
                if false {
                    v18 := GenS1{F0: 2147483647, F1: 2}
                    if (!true) {
                        v17 = "a"
                        var v19 int16 = 157
                        v20 := []int{fi14}
                        v19 = (^(v19))
                        switch fi14 {
                        case 92:
                            var v21 int32 = 119
                            v17 = "a"
                            _ = v21
                        case -845:
                            var v22 bool = true
                            _ = v22
                        case 128:
                            v19 -= (v19 << ((18446744073709551615) & 63))
                        }
                        fmt.Println(fi14)
                        if false {
                            switch (((("a" + "\n\t") + "\n\t") + (((v17 + v17) + "hello") + v17)) + v17) {
                            case "":
                                if true {
                                    fmt.Println(fi15, fi15)
                                }
                            case "foo bar":
                                v23 := GenS2{F0: (!false), F1: false, F2: (v17 + (v17 + ("foo bar" + v17)))}
                                _ = v23
                            default:
                                fmt.Println(v19)
                            }
                            var v24 bool = false
                            _ = v24
                        } else {
                            switch 833 {
                            case 742:
                                var v25 uint16 = 338
                                v17 = ((v17 + ((("foo bar" + ("\n\t" + v17)) + v17) + "")) + v17)
                                _ = v25
                            case 218:
                                v26 := GenS3{F0: 128, F1: int32(v19)}
                                _ = v26
                            }
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                        }
                        v17 = "hello"
                        v27 := GenS3{F0: fi15, F1: 42}
                        v20 = append(v20, genFn0(4294967295, "foo bar", 1))
                        _ = v19
                        _ = v20
                        _ = v27
                    }
                    _ = v18
                } else {
                    v17 = v17
                    fmt.Println(v17)
                }
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        _ = fi14
    }
    var v28 int = 9223372036854775807
    _ = v28
    return (42 == 0)
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 42
    var v3 int32 = 0
    v4 := make(map[string]int)
    v3 -= (v3 * v3)
    _ = len(v1)
    v4[("hello" + "hello")] = -867
    var v29 uint8 = 1
    switch "" {
    case "hello":
        for fi30 := 0; fi30 < 6; fi30++ {
            delete(v4, "foo bar")
            _ = fi30
        }
    default:
        v31 := []int{}
        _ = v31
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        var v33 bool = true
        if false {
            var v34 uint64 = 620
            v35 := GenS1{F0: (v3 / ((v3) | 1)), F1: 2}
            if (false && v33) {
                delete(v4, "hello")
            }
            _ = v34
            _ = v35
        } else {
            var v36 int16 = 42
            _ = v36
        }
        for fi37 := 0; fi37 < 14; fi37++ {
            var v38 uint16 = 1
            _ = fi37
            _ = v38
        }
        _ = fi32
        _ = v33
    }
    v39 := GenS3{F0: 9223372036854775807, F1: (-(v3))}
    _ = len(v4)
    var v40 uint = 42
    v41 := GenS4{F0: 1000, F1: 32767}
    if false {
        fmt.Println(v40)
    } else {
        v40 -= v40
        for fi42 := 0; fi42 < 10; fi42++ {
            v29 -= (^(v29))
            fmt.Println(v29)
            _ = fi42
        }
        if (((0.1 == ((0.1 - 0.0) / (0.0 + (0.1 * -2.25)))) || true) && true) {
            if ((0.1 - float32(0.0)) >= (1e10 * 1.5)) {
                v3 = 20
                switch "hello" {
                case "hello":
                    for fi43 := 0; fi43 < 11; fi43++ {
                        for fi44 := 0; fi44 < 6; fi44++ {
                            var v45 uint64 = 4294967295
                            v46 := []int{genFn0((v40 ^ 0), ("hello" + ("" + "hello")), 926)}
                            if true {
                                if false {
                                    v29 = v29
                                    var v47 int64 = 9223372036854775807
                                    _ = v47
                                } else {
                                    if false {
                                        v45 = 313
                                    } else {
                                        switch fi44 {
                                        case 32768:
                                            v48 := GenS5{F0: true, F1: (v29 / ((v29) | 1))}
                                            _ = v48
                                        case -1:
                                            switch (fi44 << (((v40 << (((v40 | 65535)) & 63))) & 63)) {
                                            case -512:
                                                var v49 int32 = 42
                                                _ = v49
                                            }
                                        case -366:
                                            v50 := GenS6{F0: v40, F1: 1, F2: 612}
                                            _ = v50
                                        default:
                                            _ = len(v46)
                                        }
                                        v51 := v4[(((("foo bar" + "a") + ("hello" + "foo bar")) + "foo bar") + ("a" + "foo bar"))]
                                        _ = v51
                                    }
                                    switch "\n\t" {
                                    case "a":
                                        var v52 int16 = 42
                                        _ = v52
                                    default:
                                        var v53 float32 = 0.1
                                        _ = v53
                                    }
                                    var v54 float32 = (3.14159 - ((1e10 + ((0.1 * 0.0) * float32(0.0))) / 0.1))
                                    _ = v54
                                }
                            }
                            if (255 > v40) {
                                v40 += v40
                                v3 = (v3 - (v3 / (((v3 >> (((v40 / ((65535) | 1))) & 63))) | 1)))
                                switch "" {
                                case "a":
                                    v55 := GenS5{F0: true, F1: uint8(v40)}
                                    _ = v55
                                }
                            }
                            v29 -= v29
                            v29 += 128
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                        v56 := GenS1{F0: (v3 ^ (v3 + v3)), F1: 1}
                        _ = fi43
                        _ = v56
                    }
                }
            }
            v57 := GenS2{F0: false, F1: (int64(v29) != 42), F2: ("hello" + "foo bar")}
            v3 -= 0
            _ = v57
        } else {
            v29 += v29
            v29 -= v29
            v58 := v1[303]
            _ = v58
        }
    }
    switch genFn0(v40, ("" + "hello"), 2) {
    case 9223372036854775807:
        if ((1.5 - ((((((float64(-2.25) - float64(1.5)) / (((float64(float32((((float64((((((1e10 - ((0.0 / 1.5) + 0.1)) - (-2.25 / 1.0)) / (float32(1e10) - 0.1)) + float32((-2.25 / (((0.1 * (1.5 + 0.0)) - 0.1) / (float64(0.1) - (0.1 - (1e10 + 1.5))))))) + 0.1)) / float64((float32((3.14159 - 3.14159)) - 1.5))) * 0.1) * (float64((1e10 * (1.5 / float32(3.14159)))) / 1e10)))) * 3.14159) / 1.5) * (0.1 / 3.14159))) * 0.0) + float64(1.5)) + -2.25) + 0.1)) < (0.0 - 1.5)) {
            v59 := GenS6{F0: (v40 | 1), F1: 283, F2: (v40 << ((255) & 63))}
            v3 += (v3 / (((v3 * (^(v3)))) | 1))
            v29 -= v29
            _ = v59
        }
    }
    fmt.Println(v2)
    switch "a" {
    case "":
        fmt.Println(v3, v3, v2)
    case "hello":
        switch ("" + ((("\n\t" + "\n\t") + "\n\t") + "\n\t")) {
        case "foo bar":
            v60 := GenS7{F0: -867}
            _ = v60
        case "":
            v40 = (v40 % ((299) | 1))
        case "\n\t":
            v2 = 449
        }
    case "\n\t":
        for fi61 := 0; fi61 < 20; fi61++ {
            var v62 uint8 = v29
            v3 = v3
            fmt.Println(v40, v62, v2)
            _ = fi61
            _ = v62
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
