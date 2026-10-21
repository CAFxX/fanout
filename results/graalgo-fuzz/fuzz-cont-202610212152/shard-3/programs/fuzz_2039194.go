package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint; F1 int; F2 int }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 uint8; F1 bool; F2 int }
type GenS8 struct { F0 uint8; F1 uint32; F2 bool }
type GenS9 struct { F0 uint64; F1 uint; F2 uint64 }
func genFn0() uint64 {
    var v5 uint = 2
    v5 = (v5 | v5)
    var v6 float64 = (0.0 - 3.14159)
    v7 := GenS1{F0: 63, F1: 1, F2: 2}
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1() int {
    v8 := GenS2{F0: 1, F1: 42, F2: 399}
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 10; fi10++ {
            switch "hello" {
            case "\n\t":
                v11 := GenS3{F0: 230, F1: -128}
                _ = v11
            case "hello":
                fmt.Println(fi9, fi10)
            default:
                v12 := []int{(fi10 ^ fi9), (fi10 + -417), 0, 32768}
                _ = v12
            }
            var v13 int = (^(fi9))
            _ = fi10
            _ = v13
        }
        switch (fi9 >> ((2) & 63)) {
        case -644:
            v14 := GenS4{F0: -26}
            _ = v14
        case 127:
            if (!true) {
                fmt.Println(fi9)
                var v15 int64 = 1
                _ = v15
            } else {
                if (0 < 0) {
                    var v16 uint = 838
                    _ = v16
                } else {
                    v17 := []int{(-(fi9)), 32768, (fi9 / (((fi9 / ((-1) | 1))) | 1))}
                    if (65535 <= 1) {
                        var v18 uint32 = 770
                        var v19 int = fi9
                        if ((v18 << ((4294967295) & 63)) > (^(v18))) {
                            if len(v17) > 0 {
                                v17[0] = v19
                            }
                            fmt.Println(v18, fi9, v18)
                        } else {
                            v17 = append(v17, fi9)
                        }
                        _ = v18
                        _ = v19
                    }
                    if ((26 < -77) && false) {
                        var v20 int16 = 270
                        var v21 int32 = 42
                        _ = v20
                        _ = v21
                    }
                    _ = v17
                }
            }
        default:
            var v22 uint64 = 4294967295
            _ = v22
        }
        var v23 int16 = 176
        _ = fi9
        _ = v23
    }
    if true {
        var v24 uint64 = 4294967295
        _ = v24
    }
    if true {
        var v25 int32 = -1
        var v26 float32 = ((1e10 * 3.14159) + 3.14159)
        _ = v25
        _ = v26
    } else {
        v27 := []int{1, -836, 0}
        _ = v27
    }
    var v28 string = ""
    _ = v28
    _ = v8
    return 32767
}

func main() {
    var v1 string = ("foo bar" + "foo bar")
    var v2 int8 = 48
    var v3 uint = 18446744073709551615
    var v4 int8 = (v2 * v2)
    if false {
        for fi29 := 0; fi29 < 9; fi29++ {
            if true {
                if true {
                    switch ("\n\t" + ((v1 + ("\n\t" + v1)) + "\n\t")) {
                    case "foo bar":
                        v30 := []int{(fi29 - (fi29 - (fi29 + (^(fi29))))), (fi29 << (((v3 >> ((v3) & 63))) & 63)), fi29}
                        _ = v30
                    }
                }
            }
            fmt.Println(v1, v2)
            _ = fi29
        }
        var v31 int8 = (v4 - -52)
        v32 := []int{1}
        _ = v31
        _ = v32
    }
    fmt.Println(v3, v4)
    for fi33 := 0; fi33 < 7; fi33++ {
        var v34 uint64 = 255
        if (v3 != v3) {
            for fi35 := 0; fi35 < 15; fi35++ {
                v36 := []int{549}
                if (0.1 > (1.5 - (float64(1.5) / (1.5 + 0.1)))) {
                    var v37 uint64 = (v34 * 1)
                    _ = v37
                }
                _ = fi35
                _ = v36
            }
            var v38 float64 = -2.25
            _ = v38
        }
        switch "foo bar" {
        case "":
            v39 := GenS1{F0: -1, F1: 0, F2: 0}
            _ = v39
        }
        _ = fi33
        _ = v34
    }
    var v40 uint64 = 42
    if (!((v2 % (((v4 & v4)) | 1)) >= -27)) {
        fmt.Println(v2, v2)
        for fi41 := 0; fi41 < 9; fi41++ {
            var v42 int8 = int8(v40)
            for fi43 := 0; fi43 < 18; fi43++ {
                if (114 <= (v42 << ((v3) & 63))) {
                    v44 := []int{127, fi41}
                    _ = v44
                }
                var v45 string = ("a" + v1)
                _ = fi43
                _ = v45
            }
            _ = fi41
            _ = v42
        }
        if (true || true) {
            switch 978 {
            case 1:
                var v46 uint = (v3 - (v3 >> ((v3) & 63)))
                _ = v46
            case -2147483648:
                var v47 uint16 = 673
                _ = v47
            case -512:
                for fi48 := 0; fi48 < 20; fi48++ {
                    var v49 string = v1
                    var v50 float64 = float64(1.5)
                    switch (fi48 ^ (fi48 + (fi48 % (((fi48 - (fi48 >> (((v3 % (((v3 / ((v3) | 1))) | 1))) & 63)))) | 1)))) {
                    case 1:
                        var v51 uint = v3
                        _ = v51
                    case 128:
                        if true {
                            if false {
                                var v52 uint32 = 2147483647
                                _ = v52
                            } else {
                                v49 = (v1 + "foo bar")
                                v53 := []int{42}
                                _ = v53
                            }
                            for fi54 := 0; fi54 < 11; fi54++ {
                                if (!(((-128 != (v2 >> (((^(v3))) & 63))) || (((v40 % ((uint64(v3)) | 1)) >= (v40 + v40)) && (42 >= 840))) || (false || false))) {
                                    var v55 uint8 = 127
                                    if true {
                                        var v56 uint32 = 4294967295
                                        v50 += v50
                                        _ = v56
                                    } else {
                                        v57 := GenS5{F0: (0 <= 42), F1: -1}
                                        if false {
                                            v4 = v4
                                            var v58 bool = ((v3 | 2) < (v3 - (v3 & v3)))
                                            _ = v58
                                        } else {
                                            v40 = 0
                                        }
                                        fmt.Println(v1, v55, v40)
                                        if false {
                                            if false {
                                                v59 := []int{fi54}
                                                if (!true) {
                                                    v59 = append(v59, 32768)
                                                }
                                                _ = v59
                                            }
                                            v1 = v49
                                            v60 := GenS6{F0: 9223372036854775807, F1: 9223372036854775807}
                                            _ = v60
                                        }
                                        _ = v57
                                    }
                                    v49 = v49
                                    _ = v55
                                }
                                v61 := []int{(fi54 + (fi54 << (((^(v3))) & 63)))}
                                _ = fi54
                                _ = v61
                            }
                            var v62 uint8 = 127
                            _ = v62
                        } else {
                            if (1 != 842) {
                                if (!(0 <= -2147483648)) {
                                    v50 += 1e10
                                }
                                for fi63 := 0; fi63 < 15; fi63++ {
                                    var v64 uint = (^(v3))
                                    v65 := GenS6{F0: 42, F1: 0}
                                    v66 := GenS7{F0: 42, F1: false, F2: fi48}
                                    _ = fi63
                                    _ = v64
                                    _ = v65
                                    _ = v66
                                }
                            }
                            fmt.Println(v1)
                        }
                    case 62:
                        if ((255 > 127) && (2147483648 <= 512)) {
                            if true {
                                v50 = ((float64(1.5) - 0.1) + float64(1.5))
                                if (0 >= 0) {
                                    fmt.Println(v2, v40, v50)
                                    var v67 uint8 = 54
                                    _ = v67
                                } else {
                                    if (true && (!(v3 <= (^(v3))))) {
                                        v40 = genFn0()
                                        for fi68 := 0; fi68 < 9; fi68++ {
                                            v50 += (v50 / v50)
                                            _ = fi68
                                        }
                                    }
                                    if (!true) {
                                        var v69 uint = (v3 | 762)
                                        _ = v69
                                    } else {
                                        var v70 float64 = v50
                                        var v71 int16 = -32768
                                        _ = v70
                                        _ = v71
                                    }
                                }
                            } else {
                                var v72 int16 = -1
                                if (true && false) {
                                    v1 = v1
                                    v73 := GenS8{F0: 86, F1: 65535, F2: true}
                                    _ = v73
                                }
                                switch -2147483648 {
                                case 127:
                                    v74 := []int{(fi48 + fi48), genFn1(), (fi48 << (((^(v3))) & 63)), 32768}
                                    _ = v74
                                }
                                _ = v72
                            }
                            v4 -= (v2 ^ v4)
                        }
                    }
                    _ = fi48
                    _ = v49
                    _ = v50
                }
            }
        } else {
            v3 -= (v3 % ((v3) | 1))
        }
    }
    v75 := GenS9{F0: (v40 + 42), F1: 209, F2: (v40 - 42)}
    var v76 uint16 = 590
    fmt.Println(v76, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v75
    _ = v76
}
