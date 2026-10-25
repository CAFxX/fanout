package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 int32; F1 uint8; F2 uint16 }
func genFn0() int64 {
    var v7 int32 = -250
    var v8 int16 = -982
    fmt.Println(v7, v8)
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint32, p1 uint16) uint8 {
    var v9 float32 = ((float32((0.0 - (3.14159 * 0.1))) * (1e10 * 1e10)) * 1e10)
    var v10 int = 0
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            p0 = p0
            _ = fi12
        }
        p1 += (p1 & p1)
        _ = fi11
    }
    var v13 string = "\n\t"
    _ = p0
    _ = p1
    _ = v10
    _ = v13
    _ = v9
    return 142
}

func main() {
    var v1 uint16 = 255
    var v2 uint16 = 42
    var v3 bool = (0 <= -20)
    var v4 int64 = 768
    v5 := make(map[uint32]string)
    var v6 string = "foo bar"
    v14 := []int{-502, 566, 531}
    v15 := v5[261]
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := GenS1{F0: 128}
        v15 = "a"
        v15 = (v6 + (v6 + v6))
        _ = fi16
        _ = v17
    }
    var v18 int16 = 42
    for fi19 := 0; fi19 < 19; fi19++ {
        switch fi19 {
        case 42:
            if true {
                if v3 {
                    v1 = v1
                    if v3 {
                        _ = len(v14)
                        var v20 uint64 = 486
                        _ = v20
                    } else {
                        switch 526 {
                        case 127:
                            var v21 int32 = 1
                            _ = v21
                        case 807:
                            if true {
                                if (false && v3) {
                                    fmt.Println(v3, v3)
                                } else {
                                    var v22 uint8 = 232
                                    v23 := GenS2{F0: -1}
                                    switch v15 {
                                    case "\n\t":
                                        var v24 float64 = 1.5
                                        _ = v24
                                    case "foo bar":
                                        for fi25 := 0; fi25 < 2; fi25++ {
                                            var v26 uint16 = v2
                                            var v27 float32 = 0.0
                                            v5[365] = v15
                                            _ = fi25
                                            _ = v26
                                            _ = v27
                                        }
                                    case "a":
                                        if v3 {
                                            if true {
                                                delete(v5, uint32(v22))
                                                switch v6 {
                                                case "":
                                                    v2 += (^(v2))
                                                case "foo bar":
                                                    for fi28 := 0; fi28 < 13; fi28++ {
                                                        fmt.Println(v4, v22, fi28)
                                                        var v29 uint64 = 2
                                                        v5[538] = v6
                                                        fmt.Println(v18)
                                                        v2 = v2
                                                        var v30 uint16 = (^(v2))
                                                        if len(v14) > 0 {
                                                            v14[0] = fi19
                                                        }
                                                        v31 := GenS3{F0: uint32(v22), F1: 42, F2: 255}
                                                        var v32 int64 = int64(fi19)
                                                        _ = fi28
                                                        _ = v29
                                                        _ = v30
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                case "a":
                                                    v1 = v1
                                                default:
                                                    v1 += v2
                                                }
                                            }
                                            v33 := GenS3{F0: 632, F1: (v2 ^ (^(v2))), F2: 2}
                                            _ = v33
                                        } else {
                                            _ = len(v14)
                                        }
                                    default:
                                        v3 = false
                                    }
                                    _ = v22
                                    _ = v23
                                }
                                var v34 uint16 = v2
                                _ = v34
                            } else {
                                switch fi19 {
                                case -485:
                                    var v35 uint64 = 552
                                    _ = v35
                                case 4294967295:
                                    fmt.Println(v1, v4)
                                default:
                                    if ((^(v4)) != v4) {
                                        if ("\n\t" <= "") {
                                            var v36 bool = v3
                                            var v37 uint32 = 42
                                            var v38 int = (fi19 % ((fi19) | 1))
                                            _ = v36
                                            _ = v37
                                            _ = v38
                                        } else {
                                            switch (v6 + v15) {
                                            case "a":
                                                switch fi19 {
                                                case 32768:
                                                    var v39 bool = (((42 <= 759) && v3) && (1000 <= -648))
                                                    _ = v39
                                                case 4294967295:
                                                    v18 = 32767
                                                default:
                                                    if (65535 > 255) {
                                                        var v40 int32 = 703
                                                        _ = v40
                                                    } else {
                                                        for fi41 := 0; fi41 < 19; fi41++ {
                                                            fmt.Println(fi41, v4, v1)
                                                            v42 := GenS3{F0: 654, F1: v2, F2: 18446744073709551615}
                                                            _ = fi41
                                                            _ = v42
                                                        }
                                                        v43 := GenS3{F0: 2, F1: (v1 * (^(v1))), F2: 2}
                                                        _ = v43
                                                    }
                                                }
                                            case "\n\t":
                                                for fi44 := 0; fi44 < 18; fi44++ {
                                                    var v45 float64 = (1.5 - (0.1 + 3.14159))
                                                    v46 := GenS4{F0: v15, F1: -12}
                                                    _ = fi44
                                                    _ = v45
                                                    _ = v46
                                                }
                                            default:
                                                if false {
                                                    var v47 int32 = -2147483648
                                                    if (false && (v1 != 2)) {
                                                        _ = len(v14)
                                                    }
                                                    var v48 uint = 0
                                                    _ = v47
                                                    _ = v48
                                                }
                                            }
                                        }
                                    }
                                }
                                fmt.Println(v2, v3, v1)
                            }
                        }
                        v2 += 255
                        if v3 {
                            var v49 string = v15
                            _ = v49
                        } else {
                            v3 = (1 <= -549)
                            var v50 float32 = 1e10
                            _ = v50
                        }
                    }
                    v2 = (^(v1))
                } else {
                    for fi51 := 0; fi51 < 18; fi51++ {
                        v14 = append(v14, int(v4))
                        _ = fi51
                    }
                    var v52 int = fi19
                    if v3 {
                        if false {
                            switch (v15 + ("hello" + v6)) {
                            case "\n\t":
                                v15 = v6
                            case "a":
                                var v53 uint16 = 312
                                _ = v53
                            case "hello":
                                if true {
                                    v3 = false
                                }
                            default:
                                var v54 int8 = 0
                                _ = v54
                            }
                        }
                        var v55 int64 = (-(v4))
                        v1 = (^(v1))
                        _ = v55
                    }
                    _ = v52
                }
                var v56 uint16 = v2
                _ = v56
            } else {
                v15 = v15
                for fi57 := 0; fi57 < 15; fi57++ {
                    var v58 bool = (0 != 898)
                    _ = fi57
                    _ = v58
                }
            }
        case -1:
            var v59 float32 = 1.5
            _ = v59
        case 590:
            v4 = v4
        }
        v60 := GenS5{F0: 1, F1: (v4 ^ 1)}
        if ((v1 + (v2 << ((246) & 63))) < (v2 & v2)) {
            if (!false) {
                v1 = (^(v1))
                for fi61 := 0; fi61 < 20; fi61++ {
                    delete(v5, 305)
                    fmt.Println(fi61, fi61)
                    _ = fi61
                }
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            }
            _ = len(v5)
        }
        _ = fi19
        _ = v60
    }
    v6 = v6
    v62 := GenS2{F0: -1}
    if false {
        var v63 float32 = (-2.25 / 1.0)
        _ = v63
    } else {
        if false {
            v4 += (v4 - 0)
        }
        v64 := GenS6{F0: 2147483647, F1: genFn1(4294967295, v2), F2: (v1 / ((v1) | 1))}
        v65 := GenS5{F0: 978, F1: 1}
        _ = v64
        _ = v65
    }
    _ = len(v5)
    fmt.Println(v18)
    switch "a" {
    case "a":
        v6 = ("" + (v6 + "\n\t"))
    case "hello":
        fmt.Println(v2)
    }
    var v66 float32 = (-2.25 / -2.25)
    var v67 int = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v66
    _ = v67
}
