package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int; F1 bool; F2 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int, p1 uint) int16 {
    fmt.Println(p1, p0)
    var v7 int8 = -1
    _ = p0
    _ = p1
    _ = v7
    return -262
}
func genFn1() uint32 {
    var v8 int64 = 675
    v8 += -1
    v8 += v8
    var v9 string = "hello"
    v9 = (v9 + ("foo bar" + ((((v9 + v9) + v9) + "hello") + v9)))
    _ = v8
    _ = v9
    return 157
}

func main() {
    v1 := []int{32768, -87, 4294967295}
    var v2 uint = 579
    v3 := make(map[uint32]bool)
    var v4 int8 = -1
    var v5 int16 = 1
    var v6 string = "hello"
    v2 = 65535
    delete(v3, 36)
    if false {
        if ((1e10 * 3.14159) >= (1.5 / 3.14159)) {
            v6 = "hello"
            for fi10 := 0; fi10 < 18; fi10++ {
                v11 := GenS1{F0: (0.1 < (0.1 * (-2.25 - 0.1))), F1: (!(!true))}
                v6 = v6
                _ = fi10
                _ = v11
            }
            v6 = v6
        } else {
            v12 := GenS2{F0: -9223372036854775808, F1: false, F2: 65535}
            switch v6 {
            case "foo bar":
                v13 := v3[584]
                _ = v13
            case "a":
                switch v6 {
                case "\n\t":
                    if false {
                        v14 := GenS2{F0: -1, F1: false, F2: 2147483647}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v4 = v4
                        _ = v14
                    }
                case "foo bar":
                    if (true && true) {
                        switch 739 {
                        case 738:
                            var v15 uint64 = 0
                            _ = v15
                        }
                        var v16 bool = true
                        _ = v16
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                case "hello":
                    switch (("foo bar" + v6) + ((v6 + (v6 + "\n\t")) + "\n\t")) {
                    case "a":
                        for fi17 := 0; fi17 < 3; fi17++ {
                            for fi18 := 0; fi18 < 13; fi18++ {
                                v3[0] = (false && false)
                                _ = fi18
                            }
                            _ = fi17
                        }
                    case "\n\t":
                        switch v6 {
                        case "":
                            v4 += v4
                        }
                    case "foo bar":
                        var v19 int16 = 1000
                        _ = v19
                    }
                default:
                    v20 := GenS3{F0: ""}
                    _ = v20
                }
            }
            v21 := GenS1{F0: false, F1: true}
            _ = v12
            _ = v21
        }
        v22 := GenS1{F0: true, F1: (!(((float32(-2.25) - 0.0) - 0.1) != (0.0 / (0.1 + ((1.5 / float32((1e10 / 1.0))) + (3.14159 - 0.1))))))}
        if (0 >= 0) {
            _ = len(v3)
        } else {
            var v23 uint16 = 255
            v3[213] = (false || false)
            if ((true || ((0.0 / 0.1) > 1.5)) && false) {
                var v24 bool = true
                v25 := GenS2{F0: 127, F1: ((v2 / ((42) | 1)) <= (v2 - v2)), F2: 2}
                fmt.Println(v4)
                _ = v24
                _ = v25
            } else {
                v1 = append(v1, int(v2))
                if (184 == 2) {
                    switch -1 {
                    case 2147483647:
                        v26 := GenS4{F0: 0, F1: 42}
                        _ = v26
                    }
                    fmt.Println(v2)
                    for fi27 := 0; fi27 < 9; fi27++ {
                        var v28 int64 = 0
                        var v29 uint = (v2 / ((v2) | 1))
                        fmt.Println(v6, v28, v29)
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                }
            }
            _ = v23
        }
        _ = v22
    } else {
        for fi30 := 0; fi30 < 10; fi30++ {
            v2 = v2
            var v31 uint16 = uint16(v2)
            var v32 int8 = (^(v4))
            _ = fi30
            _ = v31
            _ = v32
        }
        if true {
            delete(v3, genFn1())
            fmt.Println(v2, v4)
        }
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        for fi34 := 0; fi34 < 14; fi34++ {
            v6 = v6
            v6 = v6
            _ = fi34
        }
        fmt.Println(v4, v2)
        v6 = ((v6 + ("\n\t" + ("foo bar" + ("a" + (v6 + (("hello" + v6) + v6)))))) + v6)
        _ = fi33
    }
    switch 4294967295 {
    case -973:
        v3[2147483648] = false
    case 2147483647:
        if false {
            for fi35 := 0; fi35 < 18; fi35++ {
                switch (-(fi35)) {
                case -186:
                    fmt.Println(v2)
                case 891:
                    if false {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v2, v2)
                        if false {
                            var v36 int64 = 0
                            switch v6 {
                            case "foo bar":
                                switch fi35 {
                                case 32768:
                                    v37 := v3[42]
                                    _ = v37
                                case 0:
                                    if false {
                                        v6 = "hello"
                                    }
                                case -784:
                                    v38 := GenS3{F0: v6}
                                    _ = v38
                                default:
                                    var v39 uint16 = 65535
                                    _ = v39
                                }
                            case "\n\t":
                                _ = len(v1)
                            case "a":
                                fmt.Println(v2, v4)
                            }
                            _ = v36
                        }
                    }
                default:
                    if true {
                        var v40 uint = (v2 & v2)
                        if true {
                            v40 -= v40
                        } else {
                            var v41 uint16 = 0
                            _ = v41
                        }
                        v40 -= 255
                        _ = v40
                    } else {
                        var v42 uint16 = 42
                        _ = v42
                    }
                }
                var v43 uint32 = 65535
                _ = fi35
                _ = v43
            }
            for fi44 := 0; fi44 < 18; fi44++ {
                var v45 float32 = 1e10
                _ = fi44
                _ = v45
            }
        } else {
            v6 = v6
            v6 = v6
        }
    default:
        for fi46 := 0; fi46 < 13; fi46++ {
            var v47 float64 = 1.5
            if (!((!true) && false)) {
                switch (fi46 | (fi46 + (-(fi46)))) {
                case 488:
                    for fi48 := 0; fi48 < 6; fi48++ {
                        var v49 uint32 = genFn1()
                        _ = fi48
                        _ = v49
                    }
                case -783:
                    v50 := GenS1{F0: (false || false), F1: (4294967295 >= 42)}
                    _ = v50
                default:
                    var v51 uint = 65535
                    _ = v51
                }
                fmt.Println(v2, v47, v2)
                if (255 < 255) {
                    if true {
                        for fi52 := 0; fi52 < 9; fi52++ {
                            var v53 uint8 = 7
                            var v54 float32 = ((0.1 - 0.1) / 0.1)
                            v55 := GenS4{F0: 0, F1: v5}
                            _ = fi52
                            _ = v53
                            _ = v54
                            _ = v55
                        }
                        var v56 string = v6
                        _ = v56
                    } else {
                        fmt.Println(v6, fi46)
                        var v57 float64 = 3.14159
                        v58 := GenS5{F0: v4}
                        _ = v57
                        _ = v58
                    }
                    v6 = "a"
                    _ = len(v1)
                } else {
                    v5 += (v5 | v5)
                    v6 = v6
                    fmt.Println(v5, v4)
                }
            } else {
                for fi59 := 0; fi59 < 3; fi59++ {
                    v4 = (v4 % ((v4) | 1))
                    v6 = v6
                    fmt.Println(fi46)
                    _ = fi59
                }
                var v60 int64 = 824
                v47 -= v47
                _ = v60
            }
            _ = fi46
            _ = v47
        }
    }
    v2 += (v2 * 519)
    v61 := GenS1{F0: ((-32768 < v5) && true), F1: (0 >= 0)}
    switch v6 {
    case "hello":
        fmt.Println(v2, v6, v6)
    case "":
        v1 = append(v1, 2147483647)
    case "a":
        v6 = v6
    default:
        switch "\n\t" {
        case "":
            v62 := GenS6{F0: 4294967295, F1: 255}
            _ = v62
        case "foo bar":
            for fi63 := 0; fi63 < 4; fi63++ {
                v4 = int8(fi63)
                _ = fi63
            }
        case "hello":
            var v64 int8 = 84
            _ = v64
        default:
            _ = len(v3)
        }
    }
    fmt.Println(v6, v4)
    var v65 float64 = ((float64(3.14159) * 1e10) - (0.1 - (1e10 * 0.1)))
    var v66 int32 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v65
    _ = v66
}
