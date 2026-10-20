package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int32; F2 int32 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 bool }
func genFn0(p0 int16, p1 int) string {
    p1 = (p1 ^ p1)
    v7 := []int{p1, (p1 << ((204) & 63)), (p1 + 128), p1}
    _ = p0
    _ = p1
    _ = v7
    return "\n\t"
}
func genFn1(p0 int64, p1 uint) uint32 {
    if (true && false) {
        v8 := []int{987, -9223372036854775808, 0}
        _ = len(v8)
        var v9 int = -406
        _ = v8
        _ = v9
    }
    var v10 int32 = 42
    if false {
        var v11 uint32 = 2147483647
        _ = v11
    } else {
        var v12 int = int(v10)
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    return 0
}

func main() {
    var v1 float32 = float32(0.0)
    v2 := []int{}
    v3 := []string{((("a" + "hello") + "\n\t") + "foo bar"), "hello", "foo bar"}
    var v4 int8 = -1
    var v5 float32 = v1
    var v6 bool = (52 > 0)
    switch "" {
    case "foo bar":
        v13 := GenS1{F0: -889}
        _ = v13
    case "hello":
        var v14 int64 = -295
        _ = v14
    default:
        v15 := GenS1{F0: -270}
        _ = v15
    }
    v1 = v1
    v5 -= float32(float64(v1))
    var v16 uint32 = 2147483647
    var v17 float32 = (1.5 + v5)
    v6 = false
    var v18 float64 = (float64(1e10) / (((1.5 + 1e10) - -2.25) * 1.5))
    v6 = (209 == 2)
    switch "hello" {
    case "a":
        if (-32768 != -344) {
            fmt.Println(v6, v4)
            var v19 string = "foo bar"
            _ = v19
        }
    default:
        switch 32767 {
        case 32768:
            var v20 string = "foo bar"
            _ = v20
        case 1:
            switch "" {
            case "":
                var v21 uint32 = (^(v16))
                _ = v21
            default:
                v16 += 2
            }
        default:
            _ = len(v2)
        }
    }
    v22 := GenS2{F0: 2, F1: -740, F2: 1}
    if true {
        for fi23 := 0; fi23 < 6; fi23++ {
            if v6 {
                for fi24 := 0; fi24 < 7; fi24++ {
                    switch "" {
                    case "hello":
                        v6 = (-219 == -232)
                    case "foo bar":
                        if v6 {
                            var v25 uint = 874
                            v26 := GenS3{F0: ("a" + "\n\t"), F1: (!v6)}
                            var v27 float64 = v18
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                    default:
                        v28 := GenS1{F0: 1}
                        _ = v28
                    }
                    if v6 {
                        var v29 uint32 = (^(v16))
                        if (!v6) {
                            var v30 int32 = 1
                            v31 := GenS3{F0: "\n\t", F1: (false && (v6 && (!(!v6))))}
                            _ = v30
                            _ = v31
                        }
                        _ = v29
                    }
                    _ = fi24
                }
                if (-1 > -76) {
                    v2 = append(v2, fi23)
                }
                var v32 uint = 1
                _ = v32
            }
            _ = fi23
        }
        v3 = append(v3, "a")
        if v6 {
            var v33 uint = 65535
            var v34 int32 = 2147483647
            v35 := GenS3{F0: (("" + "foo bar") + "a"), F1: false}
            _ = v33
            _ = v34
            _ = v35
        } else {
            v5 = float32((v18 - (3.14159 + float64(v1))))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v36 bool = v6
            _ = v36
        }
    } else {
        v5 = 1.5
        var v37 bool = (2147483647 == 0)
        _ = v37
    }
    if (("" + "foo bar") == "") {
        v38 := GenS1{F0: 932}
        _ = v38
    } else {
        var v39 string = "hello"
        for fi40 := 0; fi40 < 11; fi40++ {
            fmt.Println(fi40, v4, v1)
            var v41 int8 = v4
            v39 = "a"
            _ = fi40
            _ = v41
        }
        var v42 int16 = -32768
        _ = v39
        _ = v42
    }
    for fi43 := 0; fi43 < 16; fi43++ {
        var v44 uint16 = 42
        v45 := GenS2{F0: 2, F1: 949, F2: 0}
        v46 := GenS1{F0: 205}
        _ = fi43
        _ = v44
        _ = v45
        _ = v46
    }
    v16 += 922
    for fi47 := 0; fi47 < 18; fi47++ {
        v4 -= (v4 & -55)
        _ = fi47
    }
    switch 1 {
    case -454:
        if true {
            if v6 {
                fmt.Println(v18, v5)
            } else {
                var v48 int32 = -1
                if v6 {
                    var v49 uint8 = 171
                    v17 -= 1.5
                    var v50 bool = (18446744073709551615 > 42)
                    _ = v49
                    _ = v50
                } else {
                    if false {
                        for fi51 := 0; fi51 < 11; fi51++ {
                            v2 = append(v2, fi51)
                            _ = fi51
                        }
                    } else {
                        for fi52 := 0; fi52 < 7; fi52++ {
                            for fi53 := 0; fi53 < 12; fi53++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v54 uint8 = 224
                                v17 += (v17 * (((v1 / (-2.25 * (0.0 - float32(((1.5 / (-2.25 / 0.1)) - v18))))) / v17) * float32(0.0)))
                                v55 := GenS2{F0: 255, F1: v48, F2: 42}
                                _ = fi53
                                _ = v54
                                _ = v55
                            }
                            _ = fi52
                        }
                        v48 += v48
                    }
                    for fi56 := 0; fi56 < 11; fi56++ {
                        var v57 string = "a"
                        if v6 {
                            for fi58 := 0; fi58 < 15; fi58++ {
                                v59 := GenS2{F0: 392, F1: -293, F2: v48}
                                switch 793 {
                                case 2147483647:
                                    var v60 float64 = ((float64(v5) - ((((v18 - v18) - (0.1 * v18)) / 1.0) / v18)) - float64(v1))
                                    _ = v60
                                case -2147483648:
                                    v57 = (v57 + v57)
                                case -491:
                                    fmt.Println(v48)
                                }
                                v61 := GenS4{F0: 255}
                                if v6 {
                                    if (!(v48 > (v48 + v48))) {
                                        v18 += 1e10
                                        switch (fi58 + -32) {
                                        case -2147483648:
                                            v62 := GenS5{F0: "\n\t", F1: (int8(v48) > -1)}
                                            _ = v62
                                        case 1:
                                            v16 = (v16 >> ((2) & 63))
                                        case 32768:
                                            var v63 float64 = ((1e10 - v18) - 0.1)
                                            _ = v63
                                        }
                                    }
                                    if false {
                                        v17 = float32(v18)
                                        var v64 uint8 = 223
                                        _ = v64
                                    }
                                    v16 = (v16 + v16)
                                }
                                _ = fi58
                                _ = v59
                                _ = v61
                            }
                        } else {
                            v1 = v1
                        }
                        _ = fi56
                        _ = v57
                    }
                    switch 2147483647 {
                    case 32767:
                        var v65 int = 290
                        _ = v65
                    default:
                        v66 := GenS4{F0: 251}
                        _ = v66
                    }
                }
                _ = v48
            }
        } else {
            v18 = v18
            var v67 int16 = 1
            _ = v67
        }
    default:
        v68 := GenS4{F0: 1}
        _ = v68
    }
    fmt.Println(v5, v17)
    v6 = (1 == 87)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
