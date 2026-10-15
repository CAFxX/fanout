package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 string }
type GenS2 struct { F0 bool; F1 uint; F2 uint }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint }
func genFn0(p0 bool, p1 int) int32 {
    var v6 int16 = -628
    var v7 int8 = 42
    switch "\n\t" {
    case "a":
        if (true || p0) {
            if (v7 >= v7) {
                fmt.Println(p1)
            }
            v7 = (v7 << ((255) & 63))
            for fi8 := 0; fi8 < 8; fi8++ {
                p0 = true
                v6 = (v6 / (((v6 + 0)) | 1))
                _ = fi8
            }
        }
    default:
        if (p1 <= int(v7)) {
            if (v7 > v7) {
                for fi9 := 0; fi9 < 3; fi9++ {
                    v10 := GenS1{F0: 81, F1: (-(v7)), F2: ("a" + ("\n\t" + "\n\t"))}
                    var v11 uint32 = 1
                    v11 += v11
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
            v12 := []int{}
            var v13 int8 = (v7 * -1)
            _ = v12
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -287
}
func genFn1(p0 uint32) int8 {
    v14 := GenS2{F0: false, F1: 1, F2: 0}
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 int16 = 421
        switch ((((("" + "\n\t") + ("foo bar" + "a")) + "a") + "a") + "") {
        case "\n\t":
            if false {
                p0 = 109
                v17 := []int{-448, fi15}
                _ = v17
            }
        }
        _ = fi15
        _ = v16
    }
    switch 127 {
    case 435:
        if (!false) {
            if false {
                p0 = p0
                switch 1 {
                case -1:
                    if true {
                        if true {
                            for fi18 := 0; fi18 < 9; fi18++ {
                                var v19 int8 = 1
                                v20 := []int{(fi18 >> ((255) & 63)), fi18, (fi18 / ((int(p0)) | 1))}
                                var v21 bool = false
                                _ = fi18
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                        }
                        p0 -= 4294967295
                        for fi22 := 0; fi22 < 13; fi22++ {
                            v23 := []int{(fi22 % (((-(fi22))) | 1))}
                            var v24 uint64 = 378
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                    }
                default:
                    fmt.Println(p0)
                }
                fmt.Println(p0)
            }
        }
    }
    p0 = p0
    for fi25 := 0; fi25 < 10; fi25++ {
        p0 += 65535
        _ = fi25
    }
    _ = p0
    _ = v14
    return -128
}

func main() {
    v1 := []string{}
    var v2 uint = 2
    var v3 int16 = 42
    var v4 int16 = v3
    var v5 int64 = 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 -= v4
    var v26 int16 = 881
    var v27 bool = false
    if len(v1) > 0 {
        v1[0] = ""
    }
    v28 := GenS3{F0: 0, F1: v3}
    fmt.Println(v27, v5)
    v2 -= (v2 % (((v2 % (((v2 % ((151) | 1))) | 1))) | 1))
    var v29 float64 = 1e10
    v29 = (1.5 * (v29 * (0.0 / 1.5)))
    v30 := GenS2{F0: (2147483648 != 0), F1: v2, F2: 1}
    _ = len(v1)
    for fi31 := 0; fi31 < 4; fi31++ {
        for fi32 := 0; fi32 < 11; fi32++ {
            var v33 uint = v2
            if v27 {
                switch fi31 {
                case 128:
                    switch fi31 {
                    case -533:
                        fmt.Println(v3, v2)
                    case 546:
                        var v34 int16 = (v26 >> ((v33) & 63))
                        _ = v34
                    default:
                        var v35 uint8 = 2
                        _ = v35
                    }
                case 2147483647:
                    if false {
                        var v36 uint = v2
                        v29 = (v29 / (v29 / v29))
                        _ = v36
                    } else {
                        if true {
                            v4 = (v4 ^ (v3 * 1000))
                            v37 := GenS4{F0: v2}
                            if (!(!(!(0 <= genFn0((!((961 < 42) && v27)), fi31))))) {
                                v38 := GenS4{F0: 290}
                                _ = v38
                            }
                            _ = v37
                        }
                    }
                default:
                    var v39 uint64 = 648
                    _ = v39
                }
            }
            v1 = append(v1, "a")
            _ = fi32
            _ = v33
        }
        fmt.Println(v3, v26)
        _ = fi31
    }
    var v40 string = ("foo bar" + (("hello" + ("foo bar" + "")) + "foo bar"))
    for fi41 := 0; fi41 < 6; fi41++ {
        fmt.Println(v29, v2, v3)
        _ = fi41
    }
    for fi42 := 0; fi42 < 19; fi42++ {
        v1 = append(v1, ("a" + v40))
        _ = fi42
    }
    for fi43 := 0; fi43 < 15; fi43++ {
        v1 = append(v1, "foo bar")
        v3 = 1
        _ = fi43
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v5
}
