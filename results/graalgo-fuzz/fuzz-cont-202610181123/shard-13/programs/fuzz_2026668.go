package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint32; F1 int64 }
type GenS9 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS10 struct { F0 int32 }

func main() {
    var v1 uint64 = 1
    var v2 uint32 = 2
    var v3 float32 = -2.25
    var v4 uint32 = 2147483648
    var v5 uint32 = (v2 >> ((136) & 63))
    var v6 int64 = -1
    var v7 string = "foo bar"
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := GenS1{F0: fi8, F1: 42, F2: (v1 % ((v1) | 1))}
        _ = fi8
        _ = v9
    }
    var v10 string = (v7 + v7)
    fmt.Println(v1, v10, v7)
    fmt.Println(v6, v1, v3)
    if ((-1 == 1) && false) {
        for fi11 := 0; fi11 < 12; fi11++ {
            for fi12 := 0; fi12 < 9; fi12++ {
                v13 := GenS1{F0: fi12, F1: 127, F2: v1}
                _ = fi12
                _ = v13
            }
            v14 := []int{(fi11 * (-(fi11))), fi11, int(v1), 369}
            fmt.Println(v10, v4)
            _ = fi11
            _ = v14
        }
        v15 := []int{-316}
        _ = v15
    } else {
        v16 := GenS2{F0: 32767, F1: 0}
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        if false {
            if (-917 >= 269) {
                switch "\n\t" {
                case "\n\t":
                    v18 := []int{fi17, (fi17 / ((fi17) | 1)), (-(fi17)), 957}
                    _ = v18
                }
                v19 := []int{fi17, (fi17 & (fi17 / (((fi17 - fi17)) | 1))), (fi17 / ((fi17) | 1)), (fi17 >> ((2) & 63))}
                _ = v19
            } else {
                if false {
                    v20 := []int{(fi17 / ((fi17) | 1)), int(v6), 842}
                    switch (fi17 % ((fi17) | 1)) {
                    case 206:
                        var v21 int16 = -355
                        _ = v21
                    case 448:
                        for fi22 := 0; fi22 < 16; fi22++ {
                            var v23 bool = false
                            v24 := GenS3{F0: -1, F1: 42}
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                    case 2147483647:
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                    }
                    for fi25 := 0; fi25 < 3; fi25++ {
                        var v26 bool = false
                        var v27 int16 = 1
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    _ = v20
                }
                if true {
                    switch (fi17 ^ (-(fi17))) {
                    case -258:
                        for fi28 := 0; fi28 < 20; fi28++ {
                            v1 = 2
                            var v29 string = v7
                            v30 := GenS4{F0: 1, F1: int8(v2)}
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    default:
                        var v31 float32 = v3
                        _ = v31
                    }
                    fmt.Println(v1)
                    v32 := []int{fi17, (fi17 ^ fi17)}
                    _ = v32
                }
            }
        }
        _ = fi17
    }
    for fi33 := 0; fi33 < 10; fi33++ {
        for fi34 := 0; fi34 < 12; fi34++ {
            var v35 int8 = 0
            var v36 int16 = 42
            v37 := GenS5{F0: 127, F1: v5}
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
        }
        if (1 <= 2147483647) {
            switch fi33 {
            case -991:
                fmt.Println(v4)
            default:
                var v38 int32 = 2147483647
                _ = v38
            }
        } else {
            if (v6 < v6) {
                var v39 float32 = v3
                v7 = "\n\t"
                if true {
                    v40 := []int{(-(fi33)), -2147483648}
                    if (!true) {
                        if len(v40) >= 2 {
                            v40 = v40[1:len(v40)]
                        }
                        var v41 uint16 = 902
                        _ = v41
                    } else {
                        var v42 int8 = -1
                        v40 = append(v40, fi33)
                        for fi43 := 0; fi43 < 17; fi43++ {
                            v2 -= 42
                            if ((v1 + v1) > (v1 % (((v1 % ((v1) | 1))) | 1))) {
                                if (42 < (-(fi43))) {
                                    var v44 int16 = 0
                                    _ = v44
                                } else {
                                    var v45 int16 = 1
                                    fmt.Println(v7, v42)
                                    v46 := GenS6{F0: v6, F1: -61}
                                    var v47 int32 = -2147483648
                                    _ = v45
                                    _ = v46
                                    _ = v47
                                }
                            } else {
                                v48 := GenS2{F0: 65535, F1: (v42 + v42)}
                                _ = v48
                            }
                            v49 := GenS2{F0: 1, F1: (v42 ^ (-(v42)))}
                            _ = fi43
                            _ = v49
                        }
                        _ = v42
                    }
                    if false {
                        v40 = append(v40, 127)
                        if len(v40) > 0 {
                            v40[0] = fi33
                        }
                    } else {
                        var v50 int8 = 1
                        var v51 int8 = v50
                        if len(v40) >= 2 {
                            v40 = v40[1:len(v40)]
                        }
                        _ = v50
                        _ = v51
                    }
                    _ = v40
                }
                _ = v39
            } else {
                v52 := GenS4{F0: 114, F1: 87}
                for fi53 := 0; fi53 < 12; fi53++ {
                    v54 := GenS7{F0: v6}
                    var v55 uint8 = 127
                    switch ("a" + (v10 + (v7 + v7))) {
                    case "":
                        if false {
                            switch 32768 {
                            case 169:
                                fmt.Println(fi33)
                            case -1:
                                if ((-2.25 * -2.25) == 0.0) {
                                    fmt.Println(v1, fi53, v7)
                                }
                            case -2147483648:
                                if (!false) {
                                    v56 := GenS4{F0: 42, F1: -128}
                                    v57 := []int{(-(fi53)), fi53, (fi33 >> ((uint(v1)) & 63)), (^(fi33))}
                                    if false {
                                        if (!false) {
                                            if true {
                                                v5 -= (^(v4))
                                                if (!true) {
                                                    v58 := GenS8{F0: 968, F1: (v6 * (v6 + (v6 * v6)))}
                                                    var v59 float32 = 0.1
                                                    fmt.Println(v55)
                                                    _ = v58
                                                    _ = v59
                                                } else {
                                                    v1 -= (v1 << ((65535) & 63))
                                                    v2 -= (v2 - (v5 | (v5 - v4)))
                                                    fmt.Println(v6, v6)
                                                }
                                                var v60 uint8 = (^(v55))
                                                _ = v60
                                            }
                                            v57 = append(v57, (^(fi33)))
                                            if (v1 <= v1) {
                                                v61 := GenS7{F0: (v6 >> ((14) & 63))}
                                                var v62 uint = 0
                                                if len(v57) >= 2 {
                                                    v57 = v57[1:len(v57)]
                                                }
                                                _ = v61
                                                _ = v62
                                            }
                                        }
                                        if (1e10 == 1e10) {
                                            var v63 float32 = 1.5
                                            var v64 int16 = 361
                                            var v65 uint32 = (^(v2))
                                            _ = v63
                                            _ = v64
                                            _ = v65
                                        }
                                        v6 = (v6 * v6)
                                    }
                                    _ = v56
                                    _ = v57
                                } else {
                                    v55 += 233
                                    var v66 int16 = 42
                                    var v67 bool = (42 != 596)
                                    _ = v66
                                    _ = v67
                                }
                            default:
                                var v68 int64 = v6
                                if (!(-1 >= int32(fi53))) {
                                    v69 := GenS9{F0: 321, F1: (^(v55)), F2: v6}
                                    v70 := []int{fi33}
                                    fmt.Println(fi33, v2, v1)
                                    _ = v69
                                    _ = v70
                                } else {
                                    var v71 uint64 = 2
                                    if false {
                                        v72 := GenS10{F0: -2147483648}
                                        var v73 int8 = -128
                                        switch fi33 {
                                        case -364:
                                            var v74 int32 = 0
                                            _ = v74
                                        }
                                        _ = v72
                                        _ = v73
                                    } else {
                                        fmt.Println(v2)
                                    }
                                    v75 := []int{fi33, (fi53 - 32767), fi33}
                                    _ = v71
                                    _ = v75
                                }
                                _ = v68
                            }
                            var v76 uint = 945
                            var v77 int8 = 127
                            _ = v76
                            _ = v77
                        }
                    case "\n\t":
                        v3 -= v3
                    case "hello":
                        v7 = "\n\t"
                    }
                    _ = fi53
                    _ = v54
                    _ = v55
                }
                v3 = 3.14159
                _ = v52
            }
            var v78 int32 = -78
            _ = v78
        }
        _ = fi33
    }
    v79 := []int{}
    for fi80 := 0; fi80 < 11; fi80++ {
        v1 += v1
        if len(v79) >= 2 {
            v79 = v79[1:len(v79)]
        }
        _ = fi80
    }
    var v81 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v79:", len(v79), v79)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v79
    _ = v81
}
