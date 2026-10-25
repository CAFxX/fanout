package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 bool; F1 bool; F2 string }
type GenS4 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint; F1 uint; F2 uint }
type GenS8 struct { F0 int32 }
func genFn0(p0 uint16, p1 string, p2 uint8) uint32 {
    for fi7 := 0; fi7 < 8; fi7++ {
        switch (fi7 - 32767) {
        case -655:
            var v8 int32 = 42
            _ = v8
        case -893:
            var v9 int32 = 751
            _ = v9
        }
        _ = fi7
    }
    v10 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1
}
func genFn1(p0 uint8, p1 int16) string {
    v11 := []int{}
    v12 := GenS1{F0: false}
    fmt.Println(p1)
    p0 += (p0 >> ((65535) & 63))
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return "foo bar"
}

func main() {
    var v1 int16 = 1000
    v2 := []string{((("" + ("hello" + "hello")) + "foo bar") + "\n\t")}
    v3 := []int{650, 32768, 128}
    v4 := make(map[uint32]bool)
    var v5 float32 = -2.25
    v6 := make(map[int]int)
    v13 := GenS1{F0: (0 >= -274)}
    if false {
        for fi14 := 0; fi14 < 19; fi14++ {
            delete(v4, 722)
            v15 := GenS2{F0: 260, F1: 32767}
            v16 := v4[0]
            _ = fi14
            _ = v15
            _ = v16
        }
    } else {
        switch 668 {
        case 789:
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 uint64 = 0
                v5 = (((3.14159 / 0.1) * v5) + 1e10)
                _ = fi17
                _ = v18
            }
        case 668:
            v5 += v5
        }
        v5 += v5
    }
    var v19 string = "foo bar"
    var v20 int16 = v1
    for fi21 := 0; fi21 < 20; fi21++ {
        for fi22 := 0; fi22 < 12; fi22++ {
            v23 := GenS2{F0: 65535, F1: 2}
            v20 = v1
            _ = fi22
            _ = v23
        }
        v24 := GenS3{F0: false, F1: true, F2: "hello"}
        _ = fi21
        _ = v24
    }
    v25 := GenS4{F0: 2147483647, F1: -26, F2: 42}
    if (2 > 517) {
        _ = len(v3)
    } else {
        for fi26 := 0; fi26 < 11; fi26++ {
            if false {
                v27 := GenS1{F0: (!(!(!false)))}
                _ = v27
            } else {
                v3 = append(v3, (fi26 & (^(fi26))))
            }
            _ = fi26
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if false {
        fmt.Println(v5)
        switch 32767 {
        case 0:
            var v28 int8 = -111
            _ = v28
        case 32767:
            v20 = (-(v20))
        case 868:
            fmt.Println(v1, v20, v19)
        default:
            fmt.Println(v5, v5)
        }
        var v29 uint64 = 0
        _ = v29
    } else {
        if (!false) {
            fmt.Println(v20, v1)
            if true {
                fmt.Println(v19)
            } else {
                fmt.Println(v19)
                for fi30 := 0; fi30 < 8; fi30++ {
                    v1 -= v20
                    v1 += -588
                    _ = fi30
                }
                for fi31 := 0; fi31 < 3; fi31++ {
                    var v32 uint8 = 127
                    v32 += (v32 / ((v32) | 1))
                    if (1 > -2147483648) {
                        var v33 string = (v19 + ("a" + "a"))
                        var v34 int16 = (v20 + v1)
                        v34 = -1
                        _ = v33
                        _ = v34
                    }
                    _ = fi31
                    _ = v32
                }
            }
            if ((v1 / (((^(v20))) | 1)) > (v1 << ((1) & 63))) {
                v5 -= v5
                if false {
                    v20 = (v20 & v20)
                }
            } else {
                v20 = (v20 % (((^(v20))) | 1))
                if ((!true) && false) {
                    v1 = -32768
                    for fi35 := 0; fi35 < 20; fi35++ {
                        v5 = v5
                        _ = fi35
                    }
                }
                fmt.Println(v19, v1, v1)
            }
        } else {
            var v36 int16 = 966
            switch 29 {
            case -9223372036854775808:
                switch v19 {
                case "\n\t":
                    v37 := GenS1{F0: (true || (false || (1.5 == 3.14159)))}
                    _ = v37
                case "hello":
                    var v38 uint16 = 2
                    _ = v38
                case "foo bar":
                    for fi39 := 0; fi39 < 12; fi39++ {
                        for fi40 := 0; fi40 < 8; fi40++ {
                            v2 = append(v2, ((v19 + "\n\t") + (v19 + ("a" + "foo bar"))))
                            var v41 uint8 = 128
                            _ = fi40
                            _ = v41
                        }
                        fmt.Println(v36)
                        if true {
                            if false {
                                switch v19 {
                                case "foo bar":
                                    v2 = append(v2, v19)
                                }
                                var v42 int64 = 9223372036854775807
                                _ = v42
                            }
                            v43 := GenS1{F0: (789 == 0)}
                            if false {
                                if false {
                                    v19 = ("" + "a")
                                }
                                if true {
                                    var v44 uint64 = 2
                                    v36 -= (v36 & (v36 << ((255) & 63)))
                                    _ = v44
                                } else {
                                    var v45 int32 = 42
                                    _ = v45
                                }
                                v5 = v5
                            } else {
                                v46 := GenS1{F0: true}
                                if len(v2) > 0 {
                                    v2[0] = v19
                                }
                                _ = v46
                            }
                            _ = v43
                        } else {
                            for fi47 := 0; fi47 < 4; fi47++ {
                                _ = len(v2)
                                v5 -= v5
                                _ = fi47
                            }
                        }
                        _ = fi39
                    }
                default:
                    var v48 int8 = 1
                    _ = v48
                }
            case 1:
                var v49 int32 = -298
                _ = v49
            case 42:
                fmt.Println(v5, v1)
            }
            _ = v36
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v5 += (v5 + v5)
    }
    switch -541 {
    case 917:
        fmt.Println(v1, v5)
    case 32768:
        if (!(!(2147483647 > 0))) {
            delete(v4, 2147483647)
            var v50 int64 = -148
            _ = v50
        }
    case -1:
        var v51 int8 = 127
        _ = v51
    default:
        fmt.Println(v5, v5, v20)
    }
    for fi52 := 0; fi52 < 16; fi52++ {
        switch v19 {
        case "a":
            v53 := GenS1{F0: true}
            _ = v53
        case "hello":
            fmt.Println(v1)
        case "\n\t":
            switch (fi52 ^ fi52) {
            case 32767:
                var v54 float32 = ((float32(-2.25) * v5) * v5)
                _ = v54
            case 42:
                delete(v4, 985)
            case 2147483647:
                for fi55 := 0; fi55 < 8; fi55++ {
                    switch (fi52 % ((fi55) | 1)) {
                    case 32767:
                        v5 = (-2.25 / ((1e10 + 3.14159) + (((v5 + (v5 * ((float32(1e10) / 0.1) / v5))) + 3.14159) - float32((1e10 + 0.1)))))
                    case 42:
                        var v56 uint = 1
                        _ = v56
                    default:
                        if false {
                            v57 := GenS5{F0: 0}
                            _ = v57
                        } else {
                            switch fi55 {
                            case 32767:
                                switch -971 {
                                case 9223372036854775807:
                                    var v58 int16 = v1
                                    _ = v58
                                default:
                                    var v59 uint8 = 2
                                    _ = v59
                                }
                            case 127:
                                if len(v3) > 0 {
                                    v3[0] = fi52
                                }
                            default:
                                v60 := GenS6{F0: -712}
                                _ = v60
                            }
                        }
                    }
                    var v61 float64 = 3.14159
                    if (-1 == 42) {
                        v62 := GenS7{F0: 0, F1: 123, F2: 540}
                        _ = v62
                    } else {
                        var v63 int64 = 947
                        _ = v63
                    }
                    _ = fi55
                    _ = v61
                }
            default:
                fmt.Println(v5, v19)
            }
        }
        if false {
            v64 := GenS8{F0: -1}
            _ = v64
        } else {
            var v65 string = v19
            _ = v65
        }
        v1 -= v1
        _ = fi52
    }
    v1 -= v1
    _ = len(v3)
    var v66 int8 = 42
    v20 += v1
    v5 = 0.0
    if (1.5 >= (3.14159 / (0.1 * (v5 / v5)))) {
        var v67 bool = (v19 == ("a" + v19))
        v68 := GenS2{F0: 1, F1: 255}
        _ = v67
        _ = v68
    } else {
        v19 = (v19 + (((v19 + "") + v19) + "\n\t"))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v66:", v66)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
}
