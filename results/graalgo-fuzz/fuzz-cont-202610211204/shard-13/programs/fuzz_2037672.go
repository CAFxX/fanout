package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int8 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 int8 }
func genFn0(p0 bool, p1 uint16, p2 int32) uint8 {
    p0 = (-1 >= 32768)
    fmt.Println(p1, p0)
    var v6 string = ("\n\t" + "")
    v7 := GenS1{F0: 0, F1: 128, F2: 20}
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint8 = 255
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int32, p1 int64) uint32 {
    p0 += 0
    v10 := GenS1{F0: 475, F1: 42, F2: -128}
    _ = p0
    _ = p1
    _ = v10
    return 527
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    var v3 int32 = 31
    v4 := make(map[uint32]uint32)
    var v5 int32 = v3
    var v11 uint8 = 1
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 string = ""
        fmt.Println(v13, v11, v11)
        _ = fi12
        _ = v13
    }
    var v14 bool = false
    for fi15 := 0; fi15 < 4; fi15++ {
        if v14 {
            fmt.Println(v5, v11)
            var v16 uint8 = v11
            _ = v16
        } else {
            if (-32768 <= 32767) {
                switch fi15 {
                case 467:
                    switch -9223372036854775808 {
                    case -2147483648:
                        v11 = genFn0(false, 65535, v3)
                    case 2147483647:
                        v17 := v2[fi15]
                        _ = v17
                    case 548:
                        v18 := v4[genFn1(v3, 0)]
                        _ = v18
                    default:
                        v19 := GenS1{F0: 4294967295, F1: 1, F2: -93}
                        _ = v19
                    }
                case 594:
                    v20 := GenS2{F0: 99, F1: 670}
                    _ = v20
                }
            }
            v21 := []int{1, (fi15 % (((fi15 + fi15)) | 1)), fi15, fi15}
            if (102 < -128) {
                if v14 {
                    for fi22 := 0; fi22 < 8; fi22++ {
                        v11 += v11
                        _ = fi22
                    }
                }
                v23 := GenS1{F0: 1, F1: fi15, F2: 32}
                fmt.Println(v14, v3, v3)
                _ = v23
            } else {
                if (!((!(42 > 2)) || true)) {
                    if (255 > 1) {
                        v11 = v11
                    } else {
                        v21 = append(v21, (-(fi15)))
                    }
                    if (v14 || ((!true) && (((-2.25 * 3.14159) <= 0.0) && false))) {
                        var v24 float32 = -2.25
                        v25 := v2[(fi15 * 2147483647)]
                        var v26 float32 = (v24 * v24)
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                } else {
                    for fi27 := 0; fi27 < 7; fi27++ {
                        v5 += v3
                        v28 := GenS3{F0: 0}
                        _ = fi27
                        _ = v28
                    }
                }
                fmt.Println(v3, v3)
            }
            _ = v21
        }
        var v29 int64 = 28
        var v30 uint16 = 582
        _ = fi15
        _ = v29
        _ = v30
    }
    v31 := GenS1{F0: 450, F1: -1, F2: 1}
    var v32 int64 = 42
    var v33 string = "foo bar"
    v34 := []int{128, 9223372036854775807, -686, int(v11)}
    var v35 string = (("" + (("hello" + v33) + v33)) + (((v33 + "foo bar") + v33) + (v33 + v33)))
    if true {
        if false {
            if (0 <= 236) {
                v35 = (v35 + ("" + v33))
                fmt.Println(v33, v11)
            }
            v5 -= v3
        } else {
            v35 = v33
        }
        if v14 {
            v14 = true
        } else {
            var v36 bool = false
            _ = len(v34)
            _ = v36
        }
    } else {
        v37 := GenS2{F0: 0, F1: v32}
        var v38 bool = v14
        _ = v37
        _ = v38
    }
    fmt.Println(v3, v14, v33)
    if (!v14) {
        for fi39 := 0; fi39 < 20; fi39++ {
            if v14 {
                for fi40 := 0; fi40 < 2; fi40++ {
                    v14 = ((v33 + v35) != ("" + (v33 + v35)))
                    if ((v14 && false) || v14) {
                        var v41 uint64 = 352
                        _ = v41
                    } else {
                        var v42 uint64 = 18446744073709551615
                        _ = v42
                    }
                    var v43 float32 = 1e10
                    _ = fi40
                    _ = v43
                }
            }
            for fi44 := 0; fi44 < 8; fi44++ {
                if ((v3 | 835) >= (^(v3))) {
                    v45 := GenS4{F0: 0}
                    if ((v5 >> ((42) & 63)) == v3) {
                        var v46 uint = 42
                        _ = v46
                    } else {
                        if len(v34) >= 2 {
                            v34 = v34[1:len(v34)]
                        }
                        if (!v14) {
                            switch v35 {
                            case "\n\t":
                                v3 += v3
                            case "foo bar":
                                v33 = v33
                            }
                        } else {
                            _ = len(v2)
                        }
                        if v14 {
                            v11 -= genFn0(false, 0, v3)
                            var v47 bool = v14
                            var v48 int = fi44
                            _ = v47
                            _ = v48
                        }
                        var v49 string = v35
                        v50 := GenS4{F0: (v11 / ((genFn0((!v14), 255, v3)) | 1))}
                        _ = v49
                        _ = v50
                    }
                    _ = v45
                }
                _ = fi44
            }
            v51 := GenS5{F0: 65535, F1: int8(v5)}
            _ = fi39
            _ = v51
        }
        if len(v34) >= 2 {
            v34 = v34[1:len(v34)]
        }
        for fi52 := 0; fi52 < 14; fi52++ {
            var v53 uint8 = v11
            _ = fi52
            _ = v53
        }
    }
    v11 += 2
    v3 += v3
    for fi54 := 0; fi54 < 13; fi54++ {
        if (2 > 0) {
            var v55 float32 = -2.25
            fmt.Println(v11)
            _ = v55
        } else {
            if v14 {
                for fi56 := 0; fi56 < 9; fi56++ {
                    var v57 float32 = 0.1
                    _ = fi56
                    _ = v57
                }
            } else {
                if (!(v32 <= (v32 ^ v32))) {
                    var v58 uint16 = 32767
                    v11 = (v11 << ((1) & 63))
                    var v59 uint = 256
                    _ = v58
                    _ = v59
                } else {
                    v5 -= -2147483648
                    var v60 float64 = -2.25
                    _ = v60
                }
                v61 := v2[fi54]
                switch -2147483648 {
                case 4294967295:
                    v3 = v5
                case 128:
                    v32 = v32
                case 32767:
                    var v62 int8 = -89
                    _ = v62
                }
                _ = v61
            }
        }
        delete(v1, 4294967295)
        var v63 float64 = 3.14159
        _ = fi54
        _ = v63
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
