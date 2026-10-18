package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 int16; F2 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 int16; F2 uint8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8; F1 uint8 }
func genFn0(p0 uint8) uint64 {
    p0 -= (p0 << ((0) & 63))
    fmt.Println(p0)
    v4 := []int{1, 42, 128}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    p0 = p0
    _ = p0
    _ = v4
    return 2
}

func main() {
    var v1 uint = 1
    var v2 float64 = 1.5
    var v3 uint32 = 2147483648
    v2 = 0.0
    if false {
        v2 -= (v2 / v2)
    } else {
        var v5 uint64 = 18446744073709551615
        v6 := []int{2147483647, 128, 699}
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: -2147483648, F1: -876}
    var v8 uint64 = 42
    if false {
        v9 := GenS1{F0: -51, F1: 32767}
        if false {
            v10 := []int{}
            if true {
                if (!true) {
                    if (0 < 411) {
                        var v11 int32 = -1
                        _ = v11
                    }
                    _ = len(v10)
                }
            }
            _ = v10
        } else {
            var v12 float64 = v2
            _ = v12
        }
        _ = v9
    } else {
        for fi13 := 0; fi13 < 15; fi13++ {
            fmt.Println(v1, fi13)
            var v14 bool = true
            if (907 <= 32767) {
                for fi15 := 0; fi15 < 3; fi15++ {
                    switch 32767 {
                    case 942:
                        v16 := GenS1{F0: -2147483648, F1: 42}
                        _ = v16
                    case 2147483647:
                        switch -1 {
                        case -1:
                            var v17 uint32 = (v3 ^ (v3 * v3))
                            _ = v17
                        default:
                            v18 := []int{(fi15 << ((65535) & 63)), fi15, (fi13 + fi13), fi15}
                            _ = v18
                        }
                    default:
                        v3 = (v3 ^ v3)
                    }
                    v19 := GenS1{F0: 647, F1: 32767}
                    _ = fi15
                    _ = v19
                }
                if ((3.14159 + 1.5) <= v2) {
                    var v20 float32 = 0.1
                    v2 -= (float64((v20 + (v20 + v20))) * v2)
                    for fi21 := 0; fi21 < 5; fi21++ {
                        fmt.Println(fi13)
                        var v22 string = "a"
                        _ = fi21
                        _ = v22
                    }
                    _ = v20
                } else {
                    var v23 int8 = -1
                    if false {
                        if (false && true) {
                            if (false || (!v14)) {
                                v24 := []int{(^(fi13)), fi13, (fi13 | fi13), (^(fi13))}
                                _ = v24
                            }
                            v25 := []int{fi13, (fi13 ^ fi13), (fi13 + -1), (fi13 >> ((741) & 63))}
                            _ = v25
                        }
                    } else {
                        if false {
                            var v26 int8 = (v23 / ((-68) | 1))
                            _ = v26
                        } else {
                            var v27 string = ("a" + "foo bar")
                            _ = v27
                        }
                    }
                    for fi28 := 0; fi28 < 18; fi28++ {
                        switch ((("\n\t" + "foo bar") + "\n\t") + "foo bar") {
                        case "hello":
                            v2 = -2.25
                        case "foo bar":
                            v2 -= 3.14159
                        default:
                            switch ("hello" + ("hello" + "\n\t")) {
                            case "foo bar":
                                if (89 == v23) {
                                    var v29 bool = v14
                                    if ((fi13 << (((v1 - v1)) & 63)) <= fi28) {
                                        v8 = 301
                                        var v30 uint = uint(v3)
                                        _ = v30
                                    } else {
                                        v31 := []int{}
                                        fmt.Println(v23, v8, v14)
                                        _ = v31
                                    }
                                    v32 := GenS1{F0: 42, F1: 244}
                                    v33 := GenS2{F0: -1, F1: (v23 * v23)}
                                    if (2 > 2147483648) {
                                        v23 -= v23
                                        if true {
                                            var v34 uint64 = v8
                                            if false {
                                                if v29 {
                                                    v35 := []int{fi13, fi13}
                                                    v34 -= 1
                                                    var v36 bool = v29
                                                    v37 := GenS1{F0: -1, F1: 186}
                                                    v38 := GenS1{F0: -95, F1: 191}
                                                    _ = v35
                                                    _ = v36
                                                    _ = v37
                                                    _ = v38
                                                } else {
                                                    v39 := GenS3{F0: (127 <= uint8(v8))}
                                                    _ = v39
                                                }
                                                v2 = v2
                                            } else {
                                                if (fi28 >= fi28) {
                                                    v40 := GenS2{F0: 709, F1: v23}
                                                    _ = v40
                                                }
                                            }
                                            v41 := []int{int(v3)}
                                            _ = v34
                                            _ = v41
                                        } else {
                                            if true {
                                                if v14 {
                                                    var v42 int64 = 638
                                                    fmt.Println(v8, v3, v8)
                                                    v3 = (v3 >> ((255) & 63))
                                                    _ = v42
                                                } else {
                                                    if ((v14 && (609 <= 1)) || (v29 || true)) {
                                                        v8 += 0
                                                        v43 := []int{}
                                                        _ = v43
                                                    } else {
                                                        if ((false && false) && v14) {
                                                            v44 := []int{}
                                                            fmt.Println(v2, v8, v23)
                                                            switch (fi13 | (^(fi13))) {
                                                            case 32768:
                                                                var v45 string = "foo bar"
                                                                _ = v45
                                                            case 865:
                                                                if true {
                                                                    if len(v44) >= 2 {
                                                                        v44 = v44[1:len(v44)]
                                                                    }
                                                                    fmt.Println(fi13)
                                                                } else {
                                                                    v23 -= v23
                                                                    v46 := GenS4{F0: -128, F1: 0, F2: 843}
                                                                    var v47 int32 = 727
                                                                    _ = v46
                                                                    _ = v47
                                                                }
                                                            case -424:
                                                                fmt.Println(v3)
                                                            default:
                                                                var v48 int32 = 42
                                                                _ = v48
                                                            }
                                                            _ = v44
                                                        }
                                                        fmt.Println(v23, v14)
                                                        v49 := GenS5{F0: (v8 > v8)}
                                                        v50 := GenS4{F0: -128, F1: 32767, F2: 0}
                                                        _ = v49
                                                        _ = v50
                                                    }
                                                    var v51 uint32 = v3
                                                    _ = v51
                                                }
                                                fmt.Println(v23, v1)
                                                v3 -= v3
                                                if (0 > 9223372036854775807) {
                                                    v2 -= v2
                                                    switch "" {
                                                    case "\n\t":
                                                        v3 += 391
                                                    case "foo bar":
                                                        fmt.Println(v29, v8, v3)
                                                    }
                                                    v52 := GenS1{F0: -2147483648, F1: 1}
                                                    var v53 int32 = 2147483647
                                                    v3 = (v3 * (v3 ^ v3))
                                                    if false {
                                                        var v54 uint16 = 42
                                                        _ = v54
                                                    }
                                                    _ = v52
                                                    _ = v53
                                                }
                                                v55 := GenS2{F0: -521, F1: v23}
                                                _ = v55
                                            }
                                            fmt.Println(v3, v14, v3)
                                            var v56 uint32 = (v3 - (v3 % ((4294967295) | 1)))
                                            _ = v56
                                        }
                                    } else {
                                        v57 := []int{128, -9223372036854775808}
                                        fmt.Println(v8, fi13)
                                        v58 := GenS5{F0: true}
                                        _ = v57
                                        _ = v58
                                    }
                                    v1 -= (v1 >> ((4294967295) & 63))
                                    v59 := GenS6{F0: (v1 % ((v1) | 1)), F1: 42, F2: 71}
                                    _ = v29
                                    _ = v32
                                    _ = v33
                                    _ = v59
                                } else {
                                    if v14 {
                                        v60 := GenS2{F0: 0, F1: (v23 >> (((v1 | (^(v1)))) & 63))}
                                        _ = v60
                                    }
                                    var v61 int8 = 127
                                    var v62 int8 = (v23 | (v23 - (v61 % ((-128) | 1))))
                                    v14 = (1 >= 1)
                                    _ = v61
                                    _ = v62
                                }
                            case "hello":
                                switch -544 {
                                case -1:
                                    var v63 uint16 = 709
                                    _ = v63
                                }
                            default:
                                if (!v14) {
                                    var v64 int32 = -1
                                    v8 = (v8 * v8)
                                    _ = v64
                                } else {
                                    fmt.Println(v1, v1)
                                    if false {
                                        v14 = v14
                                    } else {
                                        v65 := GenS6{F0: 2, F1: 32767, F2: 2}
                                        _ = v65
                                    }
                                    var v66 int64 = -1
                                    _ = v66
                                }
                            }
                        }
                        v8 += 0
                        v67 := GenS7{F0: -1}
                        _ = fi28
                        _ = v67
                    }
                    _ = v23
                }
                v2 += float64((0.1 / 1e10))
            }
            _ = fi13
            _ = v14
        }
        v8 += v8
    }
    v2 -= (-2.25 + 1.5)
    for fi68 := 0; fi68 < 20; fi68++ {
        var v69 bool = ((128 >= (fi68 & fi68)) || (0 < int32(v1)))
        _ = fi68
        _ = v69
    }
    var v70 uint32 = (v3 << (((^(v1))) & 63))
    switch "\n\t" {
    case "foo bar":
        for fi71 := 0; fi71 < 14; fi71++ {
            if true {
                fmt.Println(v1)
                v8 = genFn0(72)
            }
            v72 := GenS1{F0: -997, F1: 42}
            _ = fi71
            _ = v72
        }
    case "hello":
        var v73 float64 = v2
        _ = v73
    case "a":
        var v74 bool = false
        _ = v74
    }
    fmt.Println(v3, v8)
    var v75 uint = 65535
    v76 := GenS8{F0: 100, F1: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v70:", v70)
    fmt.Println("v75:", v75)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v70
    _ = v75
    _ = v76
    _ = v8
}
