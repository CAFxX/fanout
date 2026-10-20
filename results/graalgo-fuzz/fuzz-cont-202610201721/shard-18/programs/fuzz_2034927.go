package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 string; F1 int16; F2 uint }
func genFn0() int64 {
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 string = "hello"
        fmt.Println(v5, fi4)
        _ = fi4
        _ = v5
    }
    switch "" {
    case "":
        var v6 float64 = float64(float32(3.14159))
        _ = v6
    case "\n\t":
        v7 := []int{854, 736, 128, 2147483647}
        _ = v7
    }
    var v8 uint16 = 2
    if false {
        var v9 float32 = float32(1e10)
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 uint64 = 4294967295
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        var v12 uint = 419
        _ = v12
    }
    _ = v8
    return -9223372036854775808
}
func genFn1(p0 int, p1 uint8, p2 uint8) uint {
    var v13 uint32 = 65535
    v14 := GenS1{F0: (9223372036854775807 >= p0)}
    if (-204 == 0) {
        var v15 string = (("\n\t" + (("a" + "foo bar") + ("a" + "a"))) + ("a" + "hello"))
        fmt.Println(v13, p0, p0)
        _ = v15
    }
    if (false && (!true)) {
        var v16 int64 = 1
        fmt.Println(p1, v13, p0)
        switch (-(p0)) {
        case -141:
            for fi17 := 0; fi17 < 11; fi17++ {
                v18 := GenS2{F0: -1, F1: (p0 << ((4294967295) & 63))}
                _ = fi17
                _ = v18
            }
        }
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 1
    var v3 float32 = (float32((0.1 / 1e10)) - (1.5 + 1e10))
    switch 4294967295 {
    case 2147483647:
        var v19 int8 = 42
        _ = v19
    }
    if (-1 >= 42) {
        _ = len(v1)
        v2 = 65535
        var v20 float64 = (1e10 * 0.0)
        _ = v20
    }
    v21 := GenS1{F0: true}
    v2 = (v2 - v2)
    if (((float32(0.0) / (0.1 + 0.0)) >= v3) || (true && false)) {
        for fi22 := 0; fi22 < 17; fi22++ {
            v2 = 567
            v3 -= ((((float32(-2.25) * 0.0) - 0.1) - 0.0) - 0.0)
            _ = fi22
        }
        if false {
            var v23 int32 = 0
            v3 = 0.1
            _ = v23
        }
        for fi24 := 0; fi24 < 2; fi24++ {
            if false {
                for fi25 := 0; fi25 < 7; fi25++ {
                    v3 += (float32(3.14159) - (3.14159 - v3))
                    _ = fi25
                }
            }
            if (!false) {
                var v26 string = (("" + "hello") + (("" + "\n\t") + ""))
                v27 := GenS1{F0: ((978 < 1) || true)}
                switch "" {
                case "hello":
                    for fi28 := 0; fi28 < 10; fi28++ {
                        switch fi24 {
                        case 1:
                            var v29 float32 = -2.25
                            _ = v29
                        }
                        var v30 int64 = -292
                        fmt.Println(v3, v3)
                        _ = fi28
                        _ = v30
                    }
                case "foo bar":
                    var v31 uint = 0
                    _ = v31
                case "\n\t":
                    switch ("" + ("\n\t" + "")) {
                    case "\n\t":
                        if (v2 < (v2 / (((v2 >> ((382) & 63))) | 1))) {
                            fmt.Println(fi24, v26, v26)
                            v26 = ("a" + v26)
                            v32 := []int{}
                            _ = v32
                        } else {
                            v2 += (v2 / (((v2 - (^(v2)))) | 1))
                        }
                    default:
                        v26 = ""
                    }
                }
                _ = v26
                _ = v27
            } else {
                if (true || (-9223372036854775808 > 257)) {
                    for fi33 := 0; fi33 < 18; fi33++ {
                        v34 := []int{607}
                        switch 469 {
                        case -2147483648:
                            v2 -= (v2 - v2)
                        case 128:
                            switch 1 {
                            case 32767:
                                v3 = v3
                            case 9223372036854775807:
                                switch "" {
                                case "a":
                                    v2 -= (v2 << ((517) & 63))
                                case "hello":
                                    v1[fi24] = fi33
                                default:
                                    switch (("foo bar" + "\n\t") + "a") {
                                    case "a":
                                        v35 := GenS2{F0: 381, F1: fi33}
                                        _ = v35
                                    case "\n\t":
                                        v34 = append(v34, (fi24 | fi24))
                                    default:
                                        var v36 int8 = -2
                                        _ = v36
                                    }
                                }
                            default:
                                v3 += (float32((-2.25 - ((0.0 - 1.5) + (((((0.0 / (-2.25 / 3.14159)) * 1e10) - 0.0) * 1e10) / 3.14159)))) * v3)
                            }
                        }
                        delete(v1, (fi24 * 2))
                        _ = fi33
                        _ = v34
                    }
                }
            }
            for fi37 := 0; fi37 < 20; fi37++ {
                var v38 string = "a"
                v39 := []int{150, (fi37 | (fi24 % (((^(fi37))) | 1))), (fi37 << ((genFn1((fi37 & fi37), uint8(v2), 1)) & 63))}
                var v40 int32 = 459
                _ = fi37
                _ = v38
                _ = v39
                _ = v40
            }
            _ = fi24
        }
    } else {
        if false {
            var v41 int64 = genFn0()
            v41 += (v41 & 0)
            _ = v41
        } else {
            v42 := GenS1{F0: false}
            _ = v42
        }
    }
    for fi43 := 0; fi43 < 12; fi43++ {
        var v44 float32 = v3
        if false {
            var v45 int = fi43
            v46 := []int{(v45 % (((-(fi43))) | 1))}
            fmt.Println(fi43, v2, v3)
            _ = v45
            _ = v46
        }
        _ = fi43
        _ = v44
    }
    v47 := GenS3{F0: "", F1: 1000, F2: 42}
    v3 = v3
    var v48 uint64 = 1
    for fi49 := 0; fi49 < 6; fi49++ {
        switch -1 {
        case 32767:
            var v50 float64 = (-2.25 - float64((v3 / 0.1)))
            _ = v50
        case 4294967295:
            var v51 int = fi49
            _ = v51
        case 127:
            v2 = 42
        }
        v2 -= (v2 | 42)
        _ = fi49
    }
    if (float64(0.0) <= 0.1) {
        switch 32767 {
        case -2147483648:
            for fi52 := 0; fi52 < 7; fi52++ {
                var v53 int32 = 42
                _ = fi52
                _ = v53
            }
        case 9223372036854775807:
            for fi54 := 0; fi54 < 17; fi54++ {
                v55 := GenS2{F0: 653, F1: (fi54 & fi54)}
                var v56 int64 = 502
                _ = fi54
                _ = v55
                _ = v56
            }
        }
        switch "" {
        case "\n\t":
            v57 := []int{127, 0, 4294967295, -453}
            _ = v57
        case "":
            v48 += 142
        case "a":
            fmt.Println(v2, v2)
        default:
            v58 := []int{}
            _ = v58
        }
        switch 697 {
        case 32768:
            v59 := []int{32767, 2147483647, 32767}
            _ = v59
        case 4294967295:
            v48 += (v48 & 42)
        default:
            v48 = v48
        }
    } else {
        for fi60 := 0; fi60 < 3; fi60++ {
            if true {
                v48 = (v48 & (v48 - v48))
                switch int(v48) {
                case 4294967295:
                    fmt.Println(v2, v2)
                case 42:
                    v2 = v2
                default:
                    fmt.Println(fi60, fi60)
                }
                v61 := v1[(fi60 + (fi60 * (fi60 % ((32767) | 1))))]
                _ = v61
            } else {
                for fi62 := 0; fi62 < 2; fi62++ {
                    var v63 uint16 = (v2 | v2)
                    _ = len(v1)
                    _ = fi62
                    _ = v63
                }
                v64 := []int{(fi60 >> ((1) & 63)), (fi60 % ((32768) | 1)), 240, (^(fi60))}
                _ = v64
            }
            _ = fi60
        }
        v3 += float32(1.5)
    }
    var v65 int = 9223372036854775807
    var v66 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v47
    _ = v48
    _ = v65
    _ = v66
}
