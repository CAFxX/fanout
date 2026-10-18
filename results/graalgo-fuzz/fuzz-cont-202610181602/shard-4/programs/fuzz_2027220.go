package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint; F1 bool; F2 uint }
func genFn0(p0 uint32, p1 uint64, p2 string) int8 {
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            fmt.Println(fi8, p0, fi8)
            if ("foo bar" > p2) {
                v9 := GenS1{F0: 127, F1: 398}
                v10 := []int{fi8}
                _ = len(v10)
                var v11 int32 = 1
                _ = v10
                _ = v11
                _ = v9
            }
            var v12 int32 = 1
            _ = fi8
            _ = v12
        }
        for fi13 := 0; fi13 < 17; fi13++ {
            if ((p1 % ((p1) | 1)) != p1) {
                switch fi7 {
                case 22:
                    if false {
                        p2 = ""
                    }
                case -2147483648:
                    if true {
                        var v14 int8 = int8(p0)
                        var v15 float32 = (3.14159 - (0.0 * 1.5))
                        _ = v14
                        _ = v15
                    } else {
                        fmt.Println(fi13, fi7)
                        p0 = 42
                        var v16 float64 = 0.1
                        _ = v16
                    }
                }
            }
            _ = fi13
        }
        p0 -= (p0 % (((p0 & (p0 + p0))) | 1))
        _ = fi7
    }
    var v17 int32 = 0
    p1 = (p1 % ((p1) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 1
}
func genFn1(p0 uint32, p1 int64, p2 string) string {
    v18 := GenS1{F0: -498, F1: 65535}
    p0 = p0
    p2 = (p2 + p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return "\n\t"
}

func main() {
    var v1 int = 32767
    var v2 int32 = -2147483648
    var v3 float64 = 3.14159
    var v4 uint64 = 18446744073709551615
    var v5 int64 = 9223372036854775807
    var v6 string = ((("\n\t" + ("a" + ("foo bar" + "a"))) + "hello") + ("foo bar" + "hello"))
    switch (v1 - (v1 - 9223372036854775807)) {
    case -1:
        for fi19 := 0; fi19 < 1; fi19++ {
            if true {
                var v20 uint = 18
                var v21 uint32 = uint32(v4)
                _ = v20
                _ = v21
            } else {
                if ((v5 % ((v5) | 1)) < (v5 / (((v5 / (((v5 >> ((0) & 63))) | 1))) | 1))) {
                    var v22 uint16 = 32767
                    fmt.Println(v22)
                    _ = v22
                }
            }
            var v23 bool = (-32768 > 830)
            v24 := GenS1{F0: (v1 + v1), F1: 0}
            _ = fi19
            _ = v23
            _ = v24
        }
    case 0:
        for fi25 := 0; fi25 < 1; fi25++ {
            for fi26 := 0; fi26 < 20; fi26++ {
                var v27 int16 = -1
                v1 += v1
                v2 -= (v2 ^ v2)
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
    default:
        switch 355 {
        case -60:
            v1 = v1
        case 42:
            var v28 float64 = (v3 / (((v3 * ((float64(0.0) - (0.0 - -2.25)) + v3)) / v3) / (1.5 / -2.25)))
            _ = v28
        case 4294967295:
            for fi29 := 0; fi29 < 11; fi29++ {
                for fi30 := 0; fi30 < 11; fi30++ {
                    switch 32767 {
                    case 4294967295:
                        switch v6 {
                        case "hello":
                            v31 := GenS1{F0: (^(fi29)), F1: 42}
                            _ = v31
                        case "\n\t":
                            v32 := []int{4294967295}
                            _ = v32
                        case "a":
                            v33 := []int{-547, fi30, fi29}
                            _ = v33
                        default:
                            switch fi29 {
                            case -338:
                                switch fi29 {
                                case 829:
                                    var v34 int8 = -128
                                    _ = v34
                                case 2147483647:
                                    switch (("\n\t" + (v6 + v6)) + "\n\t") {
                                    case "foo bar":
                                        var v35 int8 = genFn0(65535, 393, "\n\t")
                                        _ = v35
                                    }
                                }
                            case 2147483647:
                                if false {
                                    fmt.Println(v6, v2)
                                }
                                fmt.Println(v6, v1, v3)
                            case -1:
                                if true {
                                    v3 -= (v3 / (v3 / ((v3 / 1e10) - 0.1)))
                                    v36 := []int{(fi30 | v1)}
                                    var v37 int16 = 32767
                                    _ = v36
                                    _ = v37
                                }
                            }
                        }
                    case 551:
                        v4 -= v4
                    }
                    v38 := []int{-9223372036854775808, (-(fi29))}
                    var v39 string = v6
                    _ = fi30
                    _ = v38
                    _ = v39
                }
                fmt.Println(v4, fi29, v5)
                _ = fi29
            }
        }
    }
    var v40 int16 = -446
    fmt.Println(v5, v1)
    var v41 int8 = 127
    if (true && (((((v6 + "hello") + v6) >= v6) || (!(!(v5 <= (v5 % (((v5 << ((2) & 63))) | 1)))))) || (!(v40 <= v40)))) {
        for fi42 := 0; fi42 < 15; fi42++ {
            var v43 int8 = 100
            var v44 uint = 42
            _ = fi42
            _ = v43
            _ = v44
        }
        v45 := []int{128, 1}
        _ = v45
    } else {
        v2 -= v2
        switch (v1 ^ (v1 * v1)) {
        case -871:
            v46 := GenS2{F0: 741, F1: (v4 << ((686) & 63))}
            _ = v46
        }
        var v47 int32 = v2
        _ = v47
    }
    fmt.Println(v6, v5)
    for fi48 := 0; fi48 < 9; fi48++ {
        v49 := GenS2{F0: 255, F1: v4}
        _ = fi48
        _ = v49
    }
    if false {
        v50 := GenS1{F0: int(v41), F1: 955}
        var v51 int16 = (v40 << ((42) & 63))
        _ = v50
        _ = v51
    }
    if (!false) {
        var v52 int8 = v41
        v53 := GenS2{F0: 650, F1: (v4 ^ 0)}
        var v54 int = (^(v1))
        _ = v52
        _ = v53
        _ = v54
    } else {
        var v55 float64 = (v3 * 3.14159)
        var v56 int16 = v40
        _ = v55
        _ = v56
    }
    v57 := []int{-304, v1}
    if true {
        for fi58 := 0; fi58 < 18; fi58++ {
            if len(v57) >= 2 {
                v57 = v57[1:len(v57)]
            }
            v5 = v5
            _ = fi58
        }
    } else {
        for fi59 := 0; fi59 < 14; fi59++ {
            for fi60 := 0; fi60 < 7; fi60++ {
                v61 := GenS3{F0: 0, F1: (true || false), F2: 65535}
                _ = fi60
                _ = v61
            }
            _ = fi59
        }
        var v62 uint = 0
        _ = v62
    }
    fmt.Println(v2)
    var v63 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v57
    _ = v6
    _ = v63
}
