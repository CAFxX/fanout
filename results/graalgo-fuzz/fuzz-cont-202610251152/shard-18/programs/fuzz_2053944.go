package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 bool; F1 int16 }
type GenS8 struct { F0 int16; F1 int; F2 bool }
func genFn0(p0 int16) string {
    v5 := []int{2147483647, -457, 2147483647, 0}
    var v6 int32 = 260
    var v7 float64 = (0.0 - 3.14159)
    if (-614 < 128) {
        var v8 uint32 = 991
        _ = v8
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        if len(v5) > 0 {
            v5[0] = fi9
        }
        _ = fi9
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return "foo bar"
}
func genFn1(p0 uint64, p1 uint64) uint64 {
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS1{F0: 4294967295, F1: 0}
        var v12 int32 = 42
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: -1, F1: -1}
    for fi14 := 0; fi14 < 8; fi14++ {
        fmt.Println(fi14, fi14)
        var v15 string = ""
        fmt.Println(fi14, fi14)
        _ = fi14
        _ = v15
    }
    p1 -= 1
    if false {
        var v16 float64 = 0.0
        var v17 int = 444
        _ = v16
        _ = v17
    } else {
        if (!((1e10 + 1e10) < -2.25)) {
            for fi18 := 0; fi18 < 19; fi18++ {
                if (!true) {
                    var v19 int64 = -1
                    _ = v19
                }
                var v20 int = fi18
                fmt.Println(fi18, fi18)
                _ = fi18
                _ = v20
            }
        }
        switch 42 {
        case 127:
            v21 := GenS3{F0: -364, F1: true}
            _ = v21
        case 4294967295:
            fmt.Println(p1, p1)
        case 1:
            v22 := GenS3{F0: 0, F1: (51 < 250)}
            _ = v22
        }
        p1 = (p1 | p0)
    }
    _ = p0
    _ = p1
    _ = v13
    return 1
}

func main() {
    var v1 bool = (3.14159 == ((-2.25 + -2.25) + (3.14159 * ((1e10 + float32(-2.25)) + 0.1))))
    var v2 uint32 = 2147483647
    v3 := make(map[int]int)
    var v4 uint16 = 65535
    var v23 int8 = -50
    if (v1 && v1) {
        v24 := GenS4{F0: 18446744073709551615}
        v25 := v3[128]
        v26 := GenS3{F0: -31, F1: false}
        _ = v24
        _ = v25
        _ = v26
    } else {
        var v27 string = "hello"
        _ = v27
    }
    _ = len(v3)
    if v1 {
        if (v23 > (v23 + v23)) {
            if (0.1 > (float64((((0.0 + -2.25) / 1.5) * (3.14159 + (-2.25 - 0.0)))) * (float64((0.1 * -2.25)) * 0.1))) {
                switch -511 {
                case 0:
                    if v1 {
                        if v1 {
                            v28 := []int{}
                            _ = v28
                        }
                        fmt.Println(v2, v23, v23)
                        v29 := []int{9223372036854775807, -9223372036854775808}
                        _ = v29
                    } else {
                        v4 = (^(v4))
                    }
                case -1:
                    var v30 string = ""
                    _ = v30
                case 9223372036854775807:
                    v1 = false
                }
                v23 = v23
            } else {
                v23 += (v23 + (v23 - -1))
                fmt.Println(v1, v1)
                if ((!(9223372036854775807 != 9223372036854775807)) || true) {
                    for fi31 := 0; fi31 < 13; fi31++ {
                        for fi32 := 0; fi32 < 20; fi32++ {
                            if (false || true) {
                                fmt.Println(v2)
                                v33 := []int{(fi31 % ((1) | 1))}
                                var v34 int64 = 1
                                var v35 int8 = v23
                                _ = v33
                                _ = v34
                                _ = v35
                            } else {
                                var v36 int8 = v23
                                _ = v36
                            }
                            _ = fi32
                        }
                        _ = fi31
                    }
                    v1 = v1
                    v37 := GenS5{F0: 1, F1: v2}
                    _ = v37
                } else {
                    v1 = (v2 != (v2 & (v2 ^ 2)))
                    for fi38 := 0; fi38 < 18; fi38++ {
                        switch 9223372036854775807 {
                        case 42:
                            switch ((("a" + "\n\t") + "\n\t") + ("foo bar" + ("\n\t" + ""))) {
                            case "":
                                delete(v3, fi38)
                            case "hello":
                                var v39 int16 = 42
                                _ = v39
                            }
                        }
                        _ = fi38
                    }
                }
            }
            switch (("\n\t" + "\n\t") + "\n\t") {
            case "foo bar":
                if false {
                    if v1 {
                        var v40 int64 = 1
                        var v41 bool = (1 <= 2)
                        _ = v40
                        _ = v41
                    } else {
                        var v42 int8 = v23
                        _ = v42
                    }
                } else {
                    if (!(!v1)) {
                        v23 = (-(v23))
                    } else {
                        for fi43 := 0; fi43 < 6; fi43++ {
                            for fi44 := 0; fi44 < 19; fi44++ {
                                v2 += (^(v2))
                                _ = fi44
                            }
                            _ = fi43
                        }
                        for fi45 := 0; fi45 < 11; fi45++ {
                            for fi46 := 0; fi46 < 11; fi46++ {
                                v23 = v23
                                if v1 {
                                    var v47 uint = 575
                                    fmt.Println(v47, v1, v1)
                                    _ = v47
                                } else {
                                    var v48 int64 = -9223372036854775808
                                    var v49 int64 = v48
                                    _ = v48
                                    _ = v49
                                }
                                var v50 int64 = 9223372036854775807
                                _ = fi46
                                _ = v50
                            }
                            switch 32768 {
                            case -9223372036854775808:
                                v23 = v23
                            default:
                                var v51 uint32 = (v2 & (^(v2)))
                                _ = v51
                            }
                            if false {
                                var v52 int8 = int8(v2)
                                switch "\n\t" {
                                case "a":
                                    fmt.Println(v1, fi45)
                                case "\n\t":
                                    var v53 bool = (0 != 42)
                                    _ = v53
                                case "hello":
                                    v54 := GenS4{F0: 42}
                                    _ = v54
                                }
                                _ = v52
                            }
                            _ = fi45
                        }
                    }
                    switch ("\n\t" + "foo bar") {
                    case "a":
                        for fi55 := 0; fi55 < 2; fi55++ {
                            v1 = (v1 && v1)
                            v2 = (^(v2))
                            _ = fi55
                        }
                    case "foo bar":
                        if v1 {
                            for fi56 := 0; fi56 < 19; fi56++ {
                                v23 += v23
                                v57 := v3[2147483647]
                                _ = fi56
                                _ = v57
                            }
                        }
                    case "":
                        var v58 bool = true
                        _ = v58
                    }
                    v1 = (!v1)
                }
            }
        } else {
            for fi59 := 0; fi59 < 19; fi59++ {
                switch "foo bar" {
                case "foo bar":
                    var v60 int8 = -128
                    _ = v60
                default:
                    var v61 uint16 = (v4 >> ((458) & 63))
                    _ = v61
                }
                _ = fi59
            }
            switch 1 {
            case -548:
                if v1 {
                    switch 128 {
                    case -410:
                        var v62 uint32 = (v2 | v2)
                        _ = v62
                    case 861:
                        delete(v3, -9223372036854775808)
                    case -748:
                        v63 := []int{-2147483648}
                        _ = v63
                    default:
                        fmt.Println(v4, v4, v4)
                    }
                } else {
                    fmt.Println(v1)
                    v64 := GenS6{F0: "", F1: v23}
                    _ = v64
                }
            case 826:
                var v65 int64 = -255
                _ = v65
            case 42:
                fmt.Println(v23, v1)
            }
        }
        v4 -= (v4 * v4)
        v23 = (v23 >> ((2) & 63))
    }
    if (-2.25 <= 0.1) {
        var v66 float64 = 1.5
        _ = v66
    }
    v67 := GenS7{F0: false, F1: 1}
    if v1 {
        var v68 int64 = 182
        if v1 {
            v68 = -1
            v68 += 42
            for fi69 := 0; fi69 < 1; fi69++ {
                switch fi69 {
                case 1:
                    var v70 string = ("\n\t" + "")
                    _ = v70
                case 0:
                    for fi71 := 0; fi71 < 15; fi71++ {
                        var v72 uint16 = 737
                        var v73 string = (("\n\t" + "foo bar") + "\n\t")
                        _ = fi71
                        _ = v72
                        _ = v73
                    }
                }
                _ = fi69
            }
        }
        var v74 int64 = v68
        _ = v68
        _ = v74
    }
    v1 = ("\n\t" < "foo bar")
    v75 := GenS8{F0: -768, F1: 32768, F2: v1}
    var v76 int64 = 0
    v77 := v3[-9223372036854775808]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v67
    _ = v75
    _ = v76
    _ = v77
}
