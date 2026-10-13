package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int8 }
type GenS2 struct { F0 uint32; F1 int; F2 string }
type GenS3 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int; F1 int; F2 bool }
type GenS8 struct { F0 uint }
func genFn0(p0 uint16) int32 {
    var v6 uint64 = 730
    v6 += 2
    _ = p0
    _ = v6
    return 2147483647
}
func genFn1() uint16 {
    switch ("hello" + "") {
    case "":
        if true {
            if (false || true) {
                switch 502 {
                case 42:
                    for fi7 := 0; fi7 < 9; fi7++ {
                        var v8 uint16 = 0
                        _ = fi7
                        _ = v8
                    }
                case -65:
                    var v9 int32 = -1
                    _ = v9
                }
                var v10 bool = (3.14159 < (1.5 - (1.5 + 0.1)))
                for fi11 := 0; fi11 < 14; fi11++ {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        switch fi12 {
                        case 9223372036854775807:
                            var v13 string = ""
                            _ = v13
                        default:
                            var v14 int8 = 0
                            _ = v14
                        }
                        if (!(false || (1 > -324))) {
                            v10 = ((!(18446744073709551615 < 675)) || v10)
                            v15 := []int{}
                            switch "\n\t" {
                            case "":
                                if ((32767 == 1) && false) {
                                    var v16 int16 = 32767
                                    v17 := GenS1{F0: 831, F1: -32768, F2: 127}
                                    var v18 int32 = 791
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                } else {
                                    var v19 uint16 = 65535
                                    var v20 int32 = 24
                                    v21 := GenS1{F0: 0, F1: 1, F2: 70}
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                }
                            case "foo bar":
                                switch ("\n\t" + "a") {
                                case "hello":
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                default:
                                    v10 = v10
                                }
                            case "\n\t":
                                if v10 {
                                    fmt.Println(fi12, v10, fi12)
                                    v10 = (!(v10 || false))
                                    var v22 uint8 = 165
                                    fmt.Println(fi11, fi11, fi12)
                                    if ((v22 >> ((4294967295) & 63)) == 1) {
                                        if v10 {
                                            fmt.Println(fi11, fi12)
                                            var v23 int32 = -147
                                            v24 := GenS2{F0: 2147483648, F1: fi11, F2: "\n\t"}
                                            _ = v23
                                            _ = v24
                                        }
                                        var v25 bool = (32767 < 575)
                                        _ = v25
                                    } else {
                                        if (2 > 332) {
                                            switch fi11 {
                                            case 9223372036854775807:
                                                if len(v15) >= 2 {
                                                    v15 = v15[1:len(v15)]
                                                }
                                            case 4294967295:
                                                if v10 {
                                                    if v10 {
                                                        fmt.Println(fi11, fi12)
                                                    }
                                                    v26 := GenS3{F0: 65535, F1: 582, F2: 65535}
                                                    v27 := GenS3{F0: 2147483647, F1: 42, F2: 2}
                                                    var v28 uint32 = 0
                                                    _ = v26
                                                    _ = v27
                                                    _ = v28
                                                } else {
                                                    switch int(v22) {
                                                    case -123:
                                                        var v29 int16 = 315
                                                        _ = v29
                                                    case 0:
                                                        v10 = v10
                                                    default:
                                                        if ((!(false && (9223372036854775807 == 9223372036854775807))) || ((1 < 32767) && v10)) {
                                                            var v30 int8 = 109
                                                            if len(v15) > 0 {
                                                                v15[0] = fi12
                                                            }
                                                            v10 = false
                                                            if true {
                                                                var v31 int16 = -607
                                                                var v32 uint = 114
                                                                v32 = 65535
                                                                _ = v31
                                                                _ = v32
                                                            }
                                                            v22 -= (v22 % ((v22) | 1))
                                                            _ = v30
                                                        }
                                                    }
                                                    var v33 uint8 = v22
                                                    _ = v33
                                                }
                                            }
                                        }
                                    }
                                    _ = v22
                                }
                            }
                            _ = v15
                        }
                        _ = fi12
                    }
                    _ = fi11
                }
                _ = v10
            }
            var v34 uint = 1
            _ = v34
        }
    default:
        var v35 float64 = 0.1
        _ = v35
    }
    var v36 float32 = 0.1
    var v37 uint64 = 0
    v37 = 42
    _ = v36
    _ = v37
    return 65535
}

func main() {
    var v1 string = (("\n\t" + "\n\t") + "a")
    var v2 uint64 = 1
    var v3 int8 = 1
    var v4 uint32 = 2147483647
    var v5 uint16 = 2
    fmt.Println(v2, v5, v3)
    var v38 int8 = (v3 | -1)
    fmt.Println(v2, v5, v38)
    var v39 uint = 797
    var v40 uint = (v39 << ((2) & 63))
    fmt.Println(v39, v39, v1)
    v41 := []int{32768, int(v2), -620, -1}
    v41 = append(v41, int(v5))
    if (4294967295 != -2147483648) {
        if true {
            if (true || (-2147483648 > 0)) {
                v42 := GenS4{F0: -9223372036854775808}
                var v43 uint32 = uint32(v2)
                v41 = append(v41, 9223372036854775807)
                _ = v42
                _ = v43
            }
            fmt.Println(v5, v2, v38)
        }
        var v44 float32 = (1.5 * 1e10)
        _ = v44
    }
    if false {
        for fi45 := 0; fi45 < 17; fi45++ {
            v46 := GenS5{F0: false}
            _ = fi45
            _ = v46
        }
        if false {
            v3 -= v38
            v3 += -41
        }
    } else {
        var v47 float64 = 3.14159
        switch 127 {
        case 2147483647:
            for fi48 := 0; fi48 < 20; fi48++ {
                if true {
                    v49 := GenS6{F0: v5}
                    _ = v49
                } else {
                    v50 := GenS6{F0: (v5 >> ((2) & 63))}
                    v5 = v5
                    switch fi48 {
                    case -585:
                        var v51 uint8 = 42
                        _ = v51
                    case 32767:
                        var v52 float64 = (v47 * v47)
                        _ = v52
                    default:
                        switch -312 {
                        case 257:
                            v53 := GenS7{F0: (fi48 / (((fi48 >> ((v40) & 63))) | 1)), F1: fi48, F2: ((v3 % (((v38 - v38)) | 1)) <= v38)}
                            _ = v53
                        case 827:
                            v4 = 2147483648
                        default:
                            v54 := GenS8{F0: (^(v40))}
                            _ = v54
                        }
                    }
                    _ = v50
                }
                v40 += uint(v4)
                _ = fi48
            }
        case -2147483648:
            if ((v1 > v1) && false) {
                var v55 uint64 = (^(v2))
                switch -149 {
                case 32768:
                    v56 := GenS6{F0: v5}
                    _ = v56
                case 0:
                    var v57 int8 = v3
                    _ = v57
                case 42:
                    if ((v3 % (((v3 + 48)) | 1)) == -1) {
                        var v58 uint = 4294967295
                        var v59 float32 = (-2.25 - (0.0 + (3.14159 + 1.5)))
                        _ = len(v41)
                        _ = v58
                        _ = v59
                    }
                }
                for fi60 := 0; fi60 < 19; fi60++ {
                    fmt.Println(v2)
                    _ = fi60
                }
                _ = v55
            } else {
                v40 -= 4
                v38 = 100
            }
        default:
            switch (v1 + v1) {
            case "a":
                for fi61 := 0; fi61 < 8; fi61++ {
                    v4 += (^(v4))
                    v47 = 1e10
                    _ = fi61
                }
            default:
                v62 := GenS6{F0: (v5 - 42)}
                _ = v62
            }
        }
        if (true || false) {
            if (true && (true && false)) {
                if true {
                    v1 = "\n\t"
                } else {
                    if len(v41) >= 2 {
                        v41 = v41[1:len(v41)]
                    }
                    if len(v41) >= 2 {
                        v41 = v41[1:len(v41)]
                    }
                    fmt.Println(v1, v4, v2)
                }
                v63 := GenS8{F0: (^(v39))}
                switch v1 {
                case "":
                    for fi64 := 0; fi64 < 4; fi64++ {
                        var v65 uint8 = 48
                        _ = fi64
                        _ = v65
                    }
                default:
                    v47 -= v47
                }
                _ = v63
            } else {
                v4 = v4
                if true {
                    if len(v41) >= 2 {
                        v41 = v41[1:len(v41)]
                    }
                } else {
                    switch (v1 + v1) {
                    case "":
                        var v66 uint16 = 65535
                        _ = v66
                    }
                }
            }
        } else {
            if (!(v2 != (^(v2)))) {
                var v67 uint64 = (v2 * (v2 & (^(v2))))
                _ = v67
            }
        }
        _ = v47
    }
    v68 := GenS6{F0: (^(v5))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v68
}
