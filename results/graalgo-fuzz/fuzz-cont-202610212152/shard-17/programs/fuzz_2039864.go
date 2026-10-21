package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS6 struct { F0 uint16; F1 uint8 }
func genFn0(p0 bool, p1 int16) int8 {
    switch ("a" + ("foo bar" + "foo bar")) {
    case "a":
        var v6 uint = 42
        _ = v6
    case "hello":
        var v7 bool = p0
        _ = v7
    case "":
        v8 := []int{1, 128}
        _ = v8
    default:
        v9 := []int{205, 9223372036854775807, 42}
        _ = v9
    }
    if false {
        var v10 uint = 65535
        var v11 int32 = -737
        _ = v10
        _ = v11
    } else {
        v12 := GenS1{F0: 0}
        v13 := []int{-2, 4294967295}
        p0 = false
        _ = v12
        _ = v13
    }
    var v14 uint = 65535
    _ = p0
    _ = p1
    _ = v14
    return 0
}
func genFn1(p0 int16) uint16 {
    p0 = p0
    var v15 float32 = 0.1
    var v16 uint64 = 18446744073709551615
    _ = p0
    _ = v15
    _ = v16
    return 32767
}

func main() {
    var v1 bool = (!(!((-1 >= 2147483647) && (2147483647 >= 42))))
    var v2 int = 32767
    v3 := make(map[int]uint32)
    var v4 uint64 = 42
    var v5 float64 = 1e10
    var v17 string = ("hello" + "a")
    if ((!false) || v1) {
        fmt.Println(v17, v4)
        if v1 {
            v17 = "hello"
        } else {
            if (genFn0(((908 == 1) || (false && ("foo bar" == ""))), int16(v4)) == 127) {
                if (float64(((3.14159 - float32((v5 * ((((v5 * 1e10) - 0.0) / (v5 + float64(((0.1 / -2.25) + ((float32((v5 - 0.0)) - 0.0) * 0.1))))) / 0.1)))) * (0.1 * 1.5))) == ((v5 / (((v5 + v5) - v5) / (v5 + (((v5 - 1.5) / 1.0) + v5)))) / (v5 / v5))) {
                    for fi18 := 0; fi18 < 8; fi18++ {
                        for fi19 := 0; fi19 < 11; fi19++ {
                            var v20 uint = 42
                            v21 := GenS2{F0: (fi19 + 0), F1: uint(v4)}
                            v1 = (!(v17 != v17))
                            v5 -= 0.1
                            var v22 int16 = -343
                            switch "hello" {
                            case "a":
                                fmt.Println(v2, v4, fi19)
                            case "foo bar":
                                v22 = (v22 & v22)
                            default:
                                v5 -= (v5 - v5)
                            }
                            v23 := GenS2{F0: (v2 / ((fi18) | 1)), F1: v20}
                            switch fi19 {
                            case 32768:
                                v1 = (!false)
                            case -2147483648:
                                v1 = (995 != (fi19 + (fi19 * fi18)))
                            }
                            v24 := []int{(v2 | (fi19 << (((v20 | 0)) & 63))), (fi18 >> ((v20) & 63)), fi18}
                            _ = fi19
                            _ = v20
                            _ = v21
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        _ = fi18
                    }
                    v4 += v4
                } else {
                    for fi25 := 0; fi25 < 2; fi25++ {
                        for fi26 := 0; fi26 < 16; fi26++ {
                            v2 -= (-(fi25))
                            if (79 < 1) {
                                switch -732 {
                                case -9223372036854775808:
                                    switch "hello" {
                                    case "a":
                                        v27 := []int{-97}
                                        _ = v27
                                    default:
                                        if false {
                                            if (!v1) {
                                                var v28 int32 = -399
                                                _ = v28
                                            }
                                            var v29 int16 = 1000
                                            if false {
                                                v29 += (v29 | (v29 / (((v29 % ((-1) | 1))) | 1)))
                                                var v30 uint32 = 42
                                                var v31 float64 = ((float64((0.0 + -2.25)) + v5) - (1.5 - 0.1))
                                                _ = v30
                                                _ = v31
                                            } else {
                                                if (-2147483648 > 1) {
                                                    v3[(v2 / ((v2) | 1))] = 2147483648
                                                }
                                                v32 := []int{v2, fi26}
                                                _ = v32
                                            }
                                            _ = v29
                                        } else {
                                            v33 := GenS1{F0: -128}
                                            switch 638 {
                                            case 4294967295:
                                                var v34 bool = v1
                                                _ = v34
                                            case -9223372036854775808:
                                                var v35 uint16 = 2
                                                _ = v35
                                            default:
                                                if v1 {
                                                    var v36 uint32 = 42
                                                    v37 := []int{fi25, v2, (fi25 >> ((1) & 63))}
                                                    if false {
                                                        if false {
                                                            var v38 uint32 = (v36 << ((408) & 63))
                                                            fmt.Println(v4)
                                                            v39 := GenS1{F0: 100}
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                        v2 = fi25
                                                        fmt.Println(v17, fi25, fi25)
                                                    } else {
                                                        switch v2 {
                                                        case -2147483648:
                                                            v2 += -91
                                                        default:
                                                            v37 = append(v37, (fi25 >> ((42) & 63)))
                                                        }
                                                        if false {
                                                            v17 = ((("hello" + ((v17 + ((v17 + v17) + (v17 + ""))) + "\n\t")) + "foo bar") + "")
                                                            if false {
                                                                var v40 string = "a"
                                                                _ = v40
                                                            }
                                                        }
                                                        switch -1 {
                                                        case 32768:
                                                            var v41 float64 = (-2.25 - ((0.0 * 0.1) + v5))
                                                            _ = v41
                                                        case 576:
                                                            v42 := GenS2{F0: (fi26 / (((fi26 / (((v2 | fi26)) | 1))) | 1)), F1: 255}
                                                            _ = v42
                                                        default:
                                                            v43 := GenS1{F0: -128}
                                                            _ = v43
                                                        }
                                                        if (127 != 83) {
                                                            v44 := v3[v2]
                                                            _ = v44
                                                        }
                                                        if v1 {
                                                            v4 = (v4 >> ((4294967295) & 63))
                                                        }
                                                        fmt.Println(fi25, v17, fi26)
                                                        switch "foo bar" {
                                                        case "a":
                                                            if false {
                                                                var v45 int32 = 2147483647
                                                                switch v17 {
                                                                case "foo bar":
                                                                    var v46 float32 = 3.14159
                                                                    _ = v46
                                                                case "a":
                                                                    v2 = 712
                                                                case "hello":
                                                                    var v47 uint64 = (v4 + (v4 - 118))
                                                                    _ = v47
                                                                }
                                                                _ = v45
                                                            }
                                                        }
                                                        var v48 int16 = 290
                                                        v49 := GenS3{F0: 42}
                                                        if (562 <= 42) {
                                                            if len(v37) > 0 {
                                                                v37[0] = v2
                                                            }
                                                            v2 = v2
                                                        }
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                    _ = len(v37)
                                                    var v50 uint64 = v4
                                                    _ = v36
                                                    _ = v37
                                                    _ = v50
                                                }
                                            }
                                            var v51 float32 = 0.1
                                            _ = v33
                                            _ = v51
                                        }
                                    }
                                default:
                                    v17 = (v17 + v17)
                                }
                                if v1 {
                                    switch (fi26 & (v2 / ((fi25) | 1))) {
                                    case 4294967295:
                                        if ((v1 && (v2 == (-(fi26)))) || true) {
                                            v52 := GenS4{F0: 127}
                                            v53 := GenS5{F0: -128, F1: v4, F2: 1}
                                            _ = v52
                                            _ = v53
                                        } else {
                                            var v54 uint32 = 1
                                            v55 := GenS6{F0: 32767, F1: 42}
                                            _ = v54
                                            _ = v55
                                        }
                                    case -725:
                                        var v56 int8 = 0
                                        _ = v56
                                    default:
                                        var v57 uint32 = 0
                                        _ = v57
                                    }
                                    v4 += (^(v4))
                                    v58 := GenS4{F0: 127}
                                    v59 := []int{}
                                    if (((v4 | 86) >= v4) && false) {
                                        if len(v59) >= 2 {
                                            v59 = v59[1:len(v59)]
                                        }
                                    }
                                    if (float32(v5) == 3.14159) {
                                        v60 := v3[fi25]
                                        _ = v60
                                    }
                                    _ = v58
                                    _ = v59
                                }
                                delete(v3, fi26)
                            } else {
                                var v61 uint32 = 0
                                v5 = -2.25
                                _ = v61
                            }
                            var v62 uint64 = v4
                            fmt.Println(v2)
                            _ = fi26
                            _ = v62
                        }
                        delete(v3, fi25)
                        var v63 int64 = -175
                        _ = fi25
                        _ = v63
                    }
                    var v64 uint16 = 42
                    _ = v64
                }
            }
        }
    }
    fmt.Println(v17)
    fmt.Println(v17, v17)
    for fi65 := 0; fi65 < 9; fi65++ {
        var v66 float64 = v5
        _ = fi65
        _ = v66
    }
    var v67 float64 = ((1.5 - 1.5) + (float64(1.5) / -2.25))
    for fi68 := 0; fi68 < 9; fi68++ {
        switch v2 {
        case -9223372036854775808:
            v67 += v67
        }
        _ = fi68
    }
    switch v17 {
    case "a":
        fmt.Println(v5)
    case "\n\t":
        v17 = (((("\n\t" + (v17 + ((((((v17 + v17) + v17) + ("" + v17)) + v17) + (v17 + "a")) + v17))) + v17) + v17) + ("a" + "\n\t"))
    case "":
        delete(v3, 9223372036854775807)
    }
    fmt.Println(v1, v5)
    var v69 uint8 = 50
    var v70 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v67
    _ = v69
    _ = v70
}
