package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int; F1 string; F2 string }
type GenS3 struct { F0 uint16; F1 bool; F2 uint }
type GenS4 struct { F0 int8 }

func main() {
    var v1 int16 = 32767
    var v2 int32 = 0
    v3 := make(map[int]int)
    var v4 string = "hello"
    delete(v3, 1)
    var v5 uint32 = 849
    if (-32 >= 42) {
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := GenS1{F0: 697, F1: 0}
            _ = fi6
            _ = v7
        }
        v5 += 209
    }
    v8 := GenS1{F0: 0, F1: 0}
    var v9 float64 = 0.1
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 int16 = 32767
        var v12 uint32 = 935
        if true {
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 string = v4
                _ = fi13
                _ = v14
            }
            if true {
                var v15 uint8 = 30
                var v16 int8 = -1
                _ = v15
                _ = v16
            } else {
                if (v1 > -672) {
                    for fi17 := 0; fi17 < 5; fi17++ {
                        v5 -= 65535
                        _ = fi17
                    }
                    v11 = (v1 + (-(v11)))
                } else {
                    var v18 uint8 = uint8(v5)
                    if true {
                        v19 := v3[(fi10 / (((fi10 * (fi10 + 1))) | 1))]
                        fmt.Println(v11, v11, v1)
                        if true {
                            fmt.Println(v2, v5, v4)
                        } else {
                            var v20 bool = true
                            var v21 int = (v19 * 1)
                            switch "a" {
                            case "":
                                v22 := GenS2{F0: 32767, F1: v4, F2: v4}
                                _ = v22
                            case "a":
                                v21 = (fi10 * v19)
                            }
                            _ = v20
                            _ = v21
                        }
                        _ = v19
                    } else {
                        for fi23 := 0; fi23 < 15; fi23++ {
                            _ = len(v3)
                            var v24 float64 = 1e10
                            fmt.Println(fi10)
                            var v25 int = 9223372036854775807
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        var v26 string = ((((v4 + (v4 + v4)) + v4) + (v4 + "hello")) + "")
                        _ = v26
                    }
                    var v27 int8 = 42
                    _ = v18
                    _ = v27
                }
            }
            if ((!false) || (42 > 630)) {
                for fi28 := 0; fi28 < 7; fi28++ {
                    v4 = v4
                    _ = fi28
                }
                v29 := []int{-777, -75, fi10}
                _ = v29
            } else {
                if ((false || true) && (true && true)) {
                    for fi30 := 0; fi30 < 16; fi30++ {
                        if (false && false) {
                            if (!true) {
                                v31 := []int{fi30, 32767, (-(fi10))}
                                var v32 string = "hello"
                                _ = v31
                                _ = v32
                            } else {
                                v4 = v4
                                v33 := []int{fi30, (fi30 - 9223372036854775807), fi30, fi10}
                                _ = v33
                            }
                        } else {
                            v34 := GenS2{F0: 0, F1: v4, F2: (v4 + v4)}
                            var v35 uint32 = v12
                            _ = v34
                            _ = v35
                        }
                        v5 += 2147483648
                        v5 = 416
                        _ = fi30
                    }
                    var v36 uint8 = 0
                    var v37 float64 = v9
                    _ = v36
                    _ = v37
                }
                if (2 < 2) {
                    if (int8(v5) != -48) {
                        if (795 <= 1) {
                            _ = len(v3)
                            v38 := GenS2{F0: (fi10 - fi10), F1: v4, F2: "hello"}
                            v12 += v5
                            _ = v38
                        } else {
                            v9 += -2.25
                            v1 -= 1000
                        }
                    } else {
                        if true {
                            var v39 int = (fi10 & fi10)
                            var v40 int = fi10
                            for fi41 := 0; fi41 < 2; fi41++ {
                                if (!true) {
                                    if ((!true) || false) {
                                        if (float64(0.1) != float64((-2.25 - (3.14159 - float32(0.1))))) {
                                            var v42 uint64 = 2
                                            _ = v42
                                        }
                                    }
                                }
                                v2 -= (v2 ^ (v2 * 2147483647))
                                v12 -= v5
                                _ = fi41
                            }
                            _ = v39
                            _ = v40
                        }
                        var v43 bool = false
                        _ = v43
                    }
                    for fi44 := 0; fi44 < 14; fi44++ {
                        var v45 string = (((("a" + "") + "foo bar") + "\n\t") + ("\n\t" + v4))
                        _ = fi44
                        _ = v45
                    }
                    var v46 int32 = (v2 * v2)
                    _ = v46
                }
            }
        } else {
            v47 := GenS3{F0: 797, F1: false, F2: 2}
            fmt.Println(v9, v12)
            var v48 uint8 = 138
            _ = v47
            _ = v48
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    if false {
        if (!(!(!(!(1 >= 42))))) {
            v5 += (^(v5))
            for fi49 := 0; fi49 < 20; fi49++ {
                v9 = 1e10
                _ = fi49
            }
            var v50 uint16 = 1
            _ = v50
        } else {
            var v51 float64 = v9
            if false {
                v5 += v5
                v2 += (v2 | v2)
            }
            _ = v51
        }
    }
    v52 := []int{}
    fmt.Println(v4, v9, v4)
    switch (v4 + v4) {
    case "foo bar":
        v5 -= v5
    case "a":
        v53 := GenS1{F0: 18446744073709551615, F1: 2}
        _ = v53
    case "\n\t":
        if false {
            if ((v1 % ((v1) | 1)) == (v1 + (v1 ^ -1))) {
                if (4294967295 < 246) {
                    if len(v52) >= 2 {
                        v52 = v52[1:len(v52)]
                    }
                    v1 = v1
                } else {
                    v54 := GenS4{F0: -36}
                    for fi55 := 0; fi55 < 16; fi55++ {
                        var v56 uint16 = 65535
                        var v57 int16 = -603
                        _ = fi55
                        _ = v56
                        _ = v57
                    }
                    _ = v54
                }
                switch 127 {
                case -2147483648:
                    v4 = v4
                case 32767:
                    v5 = v5
                }
            } else {
                for fi58 := 0; fi58 < 20; fi58++ {
                    var v59 int8 = 127
                    for fi60 := 0; fi60 < 8; fi60++ {
                        fmt.Println(v4)
                        if (v1 >= v1) {
                            var v61 int8 = (^(v59))
                            var v62 uint16 = 32767
                            _ = v61
                            _ = v62
                        } else {
                            var v63 string = v4
                            _ = v63
                        }
                        var v64 int32 = v2
                        _ = fi60
                        _ = v64
                    }
                    _ = fi58
                    _ = v59
                }
            }
            if (uint(v5) == 255) {
                switch 793 {
                case -385:
                    var v65 uint8 = 2
                    _ = v65
                case -1:
                    v5 = v5
                case -269:
                    v9 = (v9 * v9)
                default:
                    v66 := GenS4{F0: 93}
                    _ = v66
                }
            } else {
                switch "foo bar" {
                case "":
                    fmt.Println(v1, v1)
                case "foo bar":
                    v67 := v3[2147483647]
                    _ = v67
                case "a":
                    if len(v52) > 0 {
                        v52[0] = -884
                    }
                default:
                    v5 = (^(v5))
                }
            }
            for fi68 := 0; fi68 < 17; fi68++ {
                for fi69 := 0; fi69 < 13; fi69++ {
                    v70 := GenS1{F0: 0, F1: 4294967295}
                    _ = fi69
                    _ = v70
                }
                for fi71 := 0; fi71 < 2; fi71++ {
                    var v72 int = fi68
                    _ = fi71
                    _ = v72
                }
                var v73 int64 = int64(fi68)
                _ = fi68
                _ = v73
            }
        }
    }
    if (!(true && true)) {
        switch (("foo bar" + v4) + "a") {
        case "a":
            switch -1 {
            case 0:
                v5 += v5
            case 9223372036854775807:
                v3[32768] = 2147483647
            case -450:
                if false {
                    var v74 uint8 = 1
                    _ = v74
                } else {
                    if len(v52) >= 2 {
                        v52 = v52[1:len(v52)]
                    }
                }
            }
        case "":
            fmt.Println(v9, v2)
        default:
            switch "" {
            case "foo bar":
                if false {
                    var v75 uint64 = 0
                    v9 = 0.0
                    _ = v75
                } else {
                    var v76 string = v4
                    _ = v76
                }
            case "hello":
                v5 -= v5
            case "a":
                var v77 int8 = 127
                _ = v77
            default:
                if false {
                    v5 -= v5
                } else {
                    v78 := GenS3{F0: 125, F1: false, F2: 2}
                    _ = v78
                }
            }
        }
        switch 128 {
        case 32767:
            v4 = (v4 + "\n\t")
        }
    }
    v9 = (0.0 - v9)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v8
    _ = v9
}
