package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 uint8; F1 bool }
func genFn0(p0 bool) uint {
    var v5 int32 = 0
    var v6 uint = 65535
    _ = p0
    _ = v5
    _ = v6
    return 2
}
func genFn1() uint64 {
    if (true && true) {
        v7 := GenS1{F0: 802, F1: genFn0((false && false))}
        var v8 uint16 = 888
        _ = v7
        _ = v8
    } else {
        var v9 float32 = 0.1
        v10 := []int{}
        _ = v10
        _ = v9
    }
    if (!false) {
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := GenS1{F0: 264, F1: 2}
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := []int{(fi13 & -68), (fi13 | fi13), 127}
            _ = fi13
            _ = v14
        }
    } else {
        v15 := GenS1{F0: 255, F1: 255}
        var v16 int = 128
        switch v16 {
        case 32768:
            for fi17 := 0; fi17 < 13; fi17++ {
                if true {
                    fmt.Println(v16, v16)
                    v16 = (fi17 ^ fi17)
                    if (false || true) {
                        switch (-(fi17)) {
                        case 42:
                            for fi18 := 0; fi18 < 6; fi18++ {
                                switch "" {
                                case "a":
                                    if true {
                                        fmt.Println(fi18, fi17, fi18)
                                        var v19 float64 = ((3.14159 - ((1.5 * float64(1e10)) + float64(0.0))) * 0.0)
                                        v20 := GenS2{F0: 103}
                                        _ = v19
                                        _ = v20
                                    }
                                default:
                                    switch fi18 {
                                    case 759:
                                        var v21 uint = 65535
                                        _ = v21
                                    case 407:
                                        v22 := GenS2{F0: 127}
                                        _ = v22
                                    default:
                                        var v23 int8 = 42
                                        _ = v23
                                    }
                                }
                                _ = fi18
                            }
                        case -1:
                            for fi24 := 0; fi24 < 14; fi24++ {
                                switch fi24 {
                                case -1:
                                    if (167 >= 0) {
                                        v25 := GenS2{F0: 127}
                                        _ = v25
                                    } else {
                                        if true {
                                            switch fi17 {
                                            case -757:
                                                v26 := []int{(fi17 >> ((42) & 63)), -2147483648, (fi17 >> ((18446744073709551615) & 63)), (fi24 + v16)}
                                                _ = v26
                                            }
                                            var v27 int16 = 32767
                                            v16 -= 32767
                                            _ = v27
                                        } else {
                                            switch "foo bar" {
                                            case "":
                                                var v28 int8 = 127
                                                _ = v28
                                            case "hello":
                                                if (197 != 2) {
                                                    var v29 int = fi24
                                                    if (42 >= -25) {
                                                        var v30 float32 = ((3.14159 * -2.25) + 0.0)
                                                        v31 := []int{(fi17 | (v16 >> ((138) & 63))), fi24, -1}
                                                        _ = v30
                                                        _ = v31
                                                    } else {
                                                        v29 = (^(v16))
                                                    }
                                                    var v32 float64 = 0.1
                                                    v32 += 0.1
                                                    switch "\n\t" {
                                                    case "":
                                                        v33 := GenS3{F0: (-830 < -413), F1: 1, F2: 42}
                                                        var v34 uint64 = 42
                                                        v35 := GenS1{F0: 705, F1: genFn0(((0 >= 1000) || ((167 >= uint8(v34)) || (("\n\t" + "a") < "foo bar"))))}
                                                        _ = v33
                                                        _ = v34
                                                        _ = v35
                                                    default:
                                                        var v36 string = ("a" + "hello")
                                                        _ = v36
                                                    }
                                                    var v37 float64 = v32
                                                    var v38 float32 = 1e10
                                                    var v39 int8 = -128
                                                    _ = v29
                                                    _ = v32
                                                    _ = v37
                                                    _ = v38
                                                    _ = v39
                                                }
                                            default:
                                                fmt.Println(fi24, fi24)
                                            }
                                            v40 := GenS1{F0: 42, F1: 255}
                                            var v41 int = v16
                                            if (!(("" + "foo bar") == ("a" + ""))) {
                                                v42 := GenS4{F0: 0, F1: 978}
                                                _ = v42
                                            } else {
                                                var v43 uint32 = 2147483647
                                                switch fi17 {
                                                case -463:
                                                    v41 = (^(v16))
                                                default:
                                                    v44 := []int{(fi24 + fi17), (fi17 & (fi24 | 127))}
                                                    _ = v44
                                                }
                                                v16 = (-(fi17))
                                                fmt.Println(v43, v41)
                                                _ = v43
                                            }
                                            _ = v40
                                            _ = v41
                                        }
                                        v16 -= 4294967295
                                    }
                                case 32768:
                                    var v45 int64 = 1
                                    _ = v45
                                case 893:
                                    if (42 != 0) {
                                        if (83 <= 65535) {
                                            switch "foo bar" {
                                            case "\n\t":
                                                v46 := []int{fi17, 9223372036854775807}
                                                _ = v46
                                            }
                                            var v47 bool = false
                                            _ = v47
                                        }
                                        v16 -= fi24
                                    } else {
                                        v16 = (fi17 ^ fi17)
                                        var v48 uint64 = 18446744073709551615
                                        _ = v48
                                    }
                                default:
                                    var v49 bool = true
                                    _ = v49
                                }
                                _ = fi24
                            }
                        case -403:
                            for fi50 := 0; fi50 < 12; fi50++ {
                                v16 -= fi50
                                fmt.Println(v16, fi50)
                                switch (v16 | fi50) {
                                case 2147483647:
                                    var v51 int16 = 32767
                                    _ = v51
                                case -1:
                                    if (1e10 > (float32(-2.25) - 1.5)) {
                                        var v52 uint16 = 1
                                        v53 := GenS5{F0: 0, F1: (477 <= 1)}
                                        _ = v52
                                        _ = v53
                                    } else {
                                        v16 = (^(v16))
                                    }
                                }
                                v16 = 1
                                v54 := GenS3{F0: true, F1: 500, F2: 65535}
                                var v55 uint = 2
                                _ = fi50
                                _ = v54
                                _ = v55
                            }
                        }
                        for fi56 := 0; fi56 < 20; fi56++ {
                            if true {
                                var v57 int32 = -1
                                v58 := []int{(fi17 | 65), (fi56 * (fi56 << ((18446744073709551615) & 63))), 2147483647, fi17}
                                var v59 int64 = 42
                                _ = v57
                                _ = v58
                                _ = v59
                            } else {
                                var v60 uint32 = 65535
                                _ = v60
                            }
                            _ = fi56
                        }
                    }
                } else {
                    for fi61 := 0; fi61 < 10; fi61++ {
                        v62 := GenS5{F0: 0, F1: false}
                        _ = fi61
                        _ = v62
                    }
                    if false {
                        var v63 uint = 4294967295
                        v64 := []int{(fi17 + (fi17 + -194)), 42}
                        _ = v63
                        _ = v64
                    }
                }
                _ = fi17
            }
        }
        _ = v15
        _ = v16
    }
    v65 := GenS5{F0: 42, F1: false}
    _ = v65
    return 2
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = 109
    var v3 float32 = -2.25
    v4 := make(map[int]string)
    v66 := GenS1{F0: 2, F1: 2}
    if false {
        v3 = v3
    }
    var v67 int8 = (v2 << ((844) & 63))
    for fi68 := 0; fi68 < 11; fi68++ {
        v69 := []int{(fi68 << ((genFn0(false)) & 63))}
        _ = len(v4)
        var v70 string = ("\n\t" + ("foo bar" + "\n\t"))
        _ = fi68
        _ = v69
        _ = v70
    }
    var v71 string = ("\n\t" + "")
    var v72 uint8 = 60
    var v73 int = -9223372036854775808
    var v74 int = (v73 * v73)
    fmt.Println(v3)
    var v75 int64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v67:", v67)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v66
    _ = v67
    _ = v71
    _ = v72
    _ = v73
    _ = v74
    _ = v75
}
