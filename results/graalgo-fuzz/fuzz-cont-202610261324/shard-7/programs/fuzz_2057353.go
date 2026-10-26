package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint64; F1 bool; F2 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint32; F1 uint }
type GenS9 struct { F0 string; F1 int32; F2 int32 }
type GenS10 struct { F0 uint; F1 int32 }
func genFn0() int16 {
    v4 := []int{42, 0, 275}
    if true {
        v5 := GenS1{F0: 32767, F1: -30}
        _ = len(v4)
        var v6 float64 = 3.14159
        _ = v5
        _ = v6
    } else {
        var v7 string = ("\n\t" + ("\n\t" + "foo bar"))
        _ = v7
    }
    v8 := GenS1{F0: -418, F1: 1000}
    switch "\n\t" {
    case "a":
        var v9 uint8 = 0
        _ = v9
    case "\n\t":
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 uint32 = 42
            _ = fi10
            _ = v11
        }
    case "foo bar":
        var v12 int16 = 1
        _ = v12
    }
    var v13 bool = false
    _ = v13
    _ = v4
    _ = v8
    return 692
}
func genFn1(p0 int8) uint64 {
    fmt.Println(p0)
    v14 := GenS1{F0: -174, F1: genFn0()}
    fmt.Println(p0)
    _ = p0
    _ = v14
    return 2
}

func main() {
    var v1 int = 4294967295
    var v2 int = v1
    v3 := make(map[int]int)
    var v15 uint64 = 274
    delete(v3, (^(v2)))
    v16 := GenS2{F0: (v15 & 258), F1: ((1 > 0) || (-399 != -1)), F2: 0}
    v17 := GenS3{F0: 0}
    var v18 int64 = -486
    if (-912 != int16(v18)) {
        v19 := GenS3{F0: 985}
        _ = v19
    } else {
        fmt.Println(v1, v15)
        fmt.Println(v15, v1, v18)
        var v20 int16 = 1
        _ = v20
    }
    v21 := []int{32768}
    if (v18 == v18) {
        fmt.Println(v18, v1)
        switch (-(v2)) {
        case -9223372036854775808:
            v22 := GenS3{F0: (v18 << ((65535) & 63))}
            _ = v22
        }
    } else {
        fmt.Println(v2, v2)
    }
    if (true && (-2147483648 == 0)) {
        var v23 float64 = 0.1
        if true {
            v3[(v1 / ((v2) | 1))] = 1
            if (v18 > (v18 ^ -800)) {
                if true {
                    v23 -= v23
                    _ = len(v3)
                } else {
                    var v24 uint32 = 4294967295
                    for fi25 := 0; fi25 < 2; fi25++ {
                        if false {
                            v26 := GenS2{F0: 334, F1: (true && false), F2: -739}
                            v27 := v3[int(v18)]
                            fmt.Println(v24)
                            _ = v26
                            _ = v27
                        } else {
                            var v28 bool = (!(19 == 127))
                            v29 := GenS4{F0: v24}
                            var v30 int8 = -1
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                        for fi31 := 0; fi31 < 5; fi31++ {
                            var v32 float32 = 3.14159
                            _ = len(v3)
                            _ = fi31
                            _ = v32
                        }
                        _ = fi25
                    }
                    _ = v24
                }
                v33 := GenS5{F0: (686 >= 65535)}
                v21 = append(v21, v2)
                _ = v33
            } else {
                var v34 float64 = v23
                for fi35 := 0; fi35 < 4; fi35++ {
                    v21 = append(v21, v1)
                    _ = fi35
                }
                switch -9223372036854775808 {
                case 32768:
                    if true {
                        v18 = (v18 & (v18 >> ((42) & 63)))
                        for fi36 := 0; fi36 < 4; fi36++ {
                            var v37 int = v2
                            v34 -= v23
                            switch "foo bar" {
                            case "a":
                                v38 := GenS6{F0: 430}
                                _ = v38
                            }
                            _ = fi36
                            _ = v37
                        }
                    } else {
                        v15 = v15
                        for fi39 := 0; fi39 < 16; fi39++ {
                            var v40 int16 = genFn0()
                            var v41 int16 = (^(v40))
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                        if (true || false) {
                            var v42 uint16 = 2
                            if len(v21) > 0 {
                                v21[0] = 127
                            }
                            _ = v42
                        } else {
                            for fi43 := 0; fi43 < 3; fi43++ {
                                switch (v2 + (v2 * (v2 >> ((1) & 63)))) {
                                case -1:
                                    for fi44 := 0; fi44 < 20; fi44++ {
                                        v45 := GenS1{F0: 127, F1: -333}
                                        v46 := GenS7{F0: false}
                                        v47 := GenS8{F0: 1, F1: 65535}
                                        _ = fi44
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                    }
                                }
                                for fi48 := 0; fi48 < 8; fi48++ {
                                    v18 = (v18 << ((2) & 63))
                                    v1 -= v1
                                    var v49 uint32 = 673
                                    _ = fi48
                                    _ = v49
                                }
                                v50 := GenS1{F0: v2, F1: 668}
                                _ = fi43
                                _ = v50
                            }
                        }
                    }
                case 32767:
                    var v51 uint = 0
                    _ = v51
                case -977:
                    var v52 uint = 65535
                    _ = v52
                }
                _ = v34
            }
            v53 := v3[-9223372036854775808]
            _ = v53
        } else {
            var v54 bool = false
            _ = v54
        }
        _ = v23
    } else {
        var v55 uint16 = 1
        _ = v55
    }
    for fi56 := 0; fi56 < 1; fi56++ {
        var v57 uint8 = 127
        _ = fi56
        _ = v57
    }
    _ = len(v21)
    v58 := GenS7{F0: ((v1 - v2) == v2)}
    if true {
        if false {
            for fi59 := 0; fi59 < 12; fi59++ {
                for fi60 := 0; fi60 < 14; fi60++ {
                    var v61 uint32 = 0
                    _ = fi60
                    _ = v61
                }
                _ = fi59
            }
            v62 := GenS5{F0: (true || ((((((0.0 * 1.5) * (3.14159 + (1.5 - (((0.1 + -2.25) - 0.1) * float32(0.0))))) / 1e10) * (3.14159 / float32((1e10 - 0.1)))) / 0.1) < (((-2.25 + (-2.25 * (float32(1e10) + float32(3.14159)))) - 1.5) * 0.0)))}
            switch "hello" {
            case "\n\t":
                for fi63 := 0; fi63 < 7; fi63++ {
                    v64 := GenS7{F0: (v18 <= v18)}
                    _ = fi63
                    _ = v64
                }
            case "foo bar":
                var v65 float64 = float64(0.0)
                _ = v65
            case "a":
                v66 := GenS9{F0: "foo bar", F1: 42, F2: -2147483648}
                _ = v66
            default:
                v1 += int(v15)
            }
            _ = v62
        } else {
            _ = len(v21)
        }
        v15 = v15
        v67 := GenS4{F0: uint32(v15)}
        _ = v67
    } else {
        v15 = v15
        var v68 float64 = (3.14159 * 1e10)
        _ = v68
    }
    for fi69 := 0; fi69 < 5; fi69++ {
        switch ("\n\t" + ("hello" + "a")) {
        case "":
            v70 := GenS6{F0: 183}
            _ = v70
        default:
            v71 := GenS10{F0: 925, F1: -68}
            _ = v71
        }
        for fi72 := 0; fi72 < 9; fi72++ {
            var v73 int8 = 38
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
            _ = fi72
            _ = v73
        }
        var v74 uint16 = 575
        _ = fi69
        _ = v74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v58
}
