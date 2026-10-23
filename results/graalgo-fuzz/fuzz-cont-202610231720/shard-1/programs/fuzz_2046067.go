package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint64; F2 int32 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 string; F2 int8 }
type GenS6 struct { F0 uint64; F1 uint64; F2 uint64 }
func genFn0(p0 int) bool {
    if (239 != 1) {
        p0 += (p0 & (p0 * p0))
        p0 = p0
        for fi7 := 0; fi7 < 13; fi7++ {
            p0 += 127
            var v8 int16 = 0
            _ = fi7
            _ = v8
        }
    }
    switch (("" + (("hello" + "a") + ("foo bar" + "\n\t"))) + (("foo bar" + "") + ("\n\t" + "hello"))) {
    case "hello":
        v9 := GenS1{F0: 581}
        _ = v9
    case "foo bar":
        v10 := GenS2{F0: 164, F1: 370, F2: 2147483647}
        _ = v10
    default:
        v11 := []int{0}
        _ = v11
    }
    p0 = p0
    p0 -= 942
    p0 = -700
    _ = p0
    return (42 < 127)
}

func main() {
    var v1 bool = ((!(0.1 >= 0.1)) || false)
    var v2 uint8 = 255
    var v3 uint16 = 639
    v4 := make(map[int]int)
    var v5 int8 = -1
    var v6 uint64 = 274
    if (v1 || (!(42 <= 36))) {
        v12 := []int{2147483647, 0, 127, int(v2)}
        _ = v12
    }
    if (749 > 547) {
        var v13 int32 = 12
        _ = v13
    } else {
        var v14 uint64 = (^(v6))
        fmt.Println(v2, v5)
        v3 = (v3 * v3)
        _ = v14
    }
    fmt.Println(v2)
    v15 := v4[66]
    var v16 string = "foo bar"
    v17 := GenS2{F0: 9223372036854775807, F1: v6, F2: 480}
    fmt.Println(v3, v2, v5)
    v3 = v3
    var v18 int8 = (-(v5))
    switch int(v2) {
    case 127:
        fmt.Println(v3, v1, v6)
    case 32768:
        for fi19 := 0; fi19 < 8; fi19++ {
            v6 = (v6 ^ v6)
            if v1 {
                if ((-9223372036854775808 <= -1) && ((0 >= 680) && true)) {
                    v20 := GenS2{F0: 406, F1: 147, F2: 1}
                    if (12 != 2147483647) {
                        v21 := GenS2{F0: int64(v6), F1: 0, F2: -204}
                        _ = v21
                    }
                    _ = v20
                } else {
                    v3 += 741
                    if v1 {
                        var v22 int16 = -368
                        v6 -= (^(v6))
                        v23 := v4[fi19]
                        _ = v22
                        _ = v23
                    }
                }
                v24 := []int{v15}
                _ = len(v24)
                _ = v24
            }
            _ = fi19
        }
    }
    v5 += (v18 * 1)
    switch ("\n\t" + ("hello" + v16)) {
    case "\n\t":
        fmt.Println(v18)
    case "":
        switch (v15 + v15) {
        case 32767:
            v1 = (((0 > 961) && (false || ("" <= v16))) || (!true))
        case -985:
            fmt.Println(v3)
        case -516:
            fmt.Println(v3, v16)
        default:
            var v25 int8 = v18
            _ = v25
        }
    case "a":
        v18 += (v18 / ((v18) | 1))
    default:
        switch v16 {
        case "foo bar":
            v2 = (^(v2))
        case "hello":
            var v26 int32 = 902
            _ = v26
        case "\n\t":
            v3 += uint16(v6)
        }
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        if true {
            var v28 uint16 = v3
            if (-920 <= 0) {
                v29 := v4[(v15 & v15)]
                if (0 != 42) {
                    switch fi27 {
                    case 1:
                        var v30 float64 = (float64(1.5) * ((3.14159 + (((float64(-2.25) / 1.0) - float64(1.5)) * 0.0)) * 1.5))
                        _ = v30
                    case 0:
                        for fi31 := 0; fi31 < 17; fi31++ {
                            if v1 {
                                if ((((v28 << ((1) & 63)) <= (^(v3))) && (881 <= 2147483647)) || (("\n\t" == ("" + v16)) && (0.1 > float32(1.5)))) {
                                    switch fi27 {
                                    case 4294967295:
                                        if ((((("hello" + v16) + "a") + v16) + "hello") <= v16) {
                                            if (((v16 + "foo bar") + v16) <= v16) {
                                                v32 := GenS2{F0: 0, F1: 443, F2: -1}
                                                _ = v32
                                            } else {
                                                var v33 uint = 255
                                                _ = v33
                                            }
                                            switch ("\n\t" + (v16 + v16)) {
                                            case "hello":
                                                fmt.Println(v15, v16)
                                            case "a":
                                                v16 = (v16 + "foo bar")
                                            default:
                                                var v34 int = fi31
                                                _ = v34
                                            }
                                        } else {
                                            v3 -= v3
                                            v15 = v15
                                            v16 = v16
                                            var v35 string = "a"
                                            _ = v35
                                        }
                                    default:
                                        if v1 {
                                            v29 -= fi31
                                            switch v16 {
                                            case "\n\t":
                                                fmt.Println(v6, v16)
                                            case "hello":
                                                switch v16 {
                                                case "hello":
                                                    v36 := GenS1{F0: 391}
                                                    _ = v36
                                                case "":
                                                    var v37 int = 563
                                                    _ = v37
                                                }
                                            }
                                            v38 := GenS1{F0: -601}
                                            var v39 float32 = 1.5
                                            switch (v16 + ((v16 + ((("\n\t" + v16) + "foo bar") + "foo bar")) + ((v16 + (("\n\t" + (v16 + "")) + v16)) + v16))) {
                                            case "a":
                                                var v40 int = fi27
                                                _ = v40
                                            case "":
                                                var v41 int16 = -586
                                                _ = v41
                                            case "\n\t":
                                                if (v6 != (v6 ^ v6)) {
                                                    v42 := []int{(fi31 - v15)}
                                                    v16 = "\n\t"
                                                    if false {
                                                        if false {
                                                            switch v29 {
                                                            case -425:
                                                                var v43 uint32 = 999
                                                                _ = v43
                                                            default:
                                                                v44 := GenS1{F0: -18}
                                                                _ = v44
                                                            }
                                                            fmt.Println(v1)
                                                            var v45 int8 = -96
                                                            if v1 {
                                                                v15 -= v15
                                                            } else {
                                                                v2 -= v2
                                                            }
                                                            v2 = v2
                                                            v46 := GenS2{F0: -9223372036854775808, F1: 18446744073709551615, F2: -2147483648}
                                                            v47 := GenS2{F0: 0, F1: (^(v6)), F2: 1}
                                                            _ = v45
                                                            _ = v46
                                                            _ = v47
                                                        } else {
                                                            fmt.Println(v39, v18)
                                                        }
                                                        fmt.Println(v39)
                                                        v5 -= 100
                                                    }
                                                    _ = v42
                                                } else {
                                                    var v48 uint = 255
                                                    v49 := GenS3{F0: 0, F1: 1}
                                                    v50 := GenS4{F0: (v48 - v48)}
                                                    if true {
                                                        fmt.Println(v1, v6)
                                                        v51 := GenS3{F0: (v18 + (^(v18))), F1: 204}
                                                        _ = v51
                                                    } else {
                                                        if (!(!(!v1))) {
                                                            _ = len(v4)
                                                        } else {
                                                            if (1e10 > (3.14159 * v39)) {
                                                                var v52 int = v29
                                                                _ = v52
                                                            } else {
                                                                v4[fi27] = (fi31 / ((int(v48)) | 1))
                                                                v53 := GenS1{F0: 0}
                                                                _ = v53
                                                            }
                                                        }
                                                        v54 := []int{(v15 >> (((v48 ^ (v48 & 4294967295))) & 63)), v29}
                                                        _ = v54
                                                    }
                                                    var v55 float32 = 1.5
                                                    _ = v48
                                                    _ = v49
                                                    _ = v50
                                                    _ = v55
                                                }
                                            default:
                                                if (!true) {
                                                    var v56 int8 = (v5 - (v5 % (((^(v5))) | 1)))
                                                    _ = v56
                                                }
                                            }
                                            v57 := []int{v15, v15, (fi27 | fi27), 9223372036854775807}
                                            _ = v38
                                            _ = v39
                                            _ = v57
                                        } else {
                                            var v58 int = fi31
                                            _ = v58
                                        }
                                    }
                                    var v59 uint32 = 2147483648
                                    var v60 float64 = -2.25
                                    _ = v59
                                    _ = v60
                                }
                            }
                            v61 := []int{(fi31 / (((fi31 | 42)) | 1)), (v15 / (((v29 ^ fi31)) | 1))}
                            switch -197 {
                            case 9223372036854775807:
                                fmt.Println(v28, v2, v1)
                            case -9223372036854775808:
                                fmt.Println(v29, v15)
                            default:
                                v6 -= (v6 << ((169) & 63))
                            }
                            _ = fi31
                            _ = v61
                        }
                    case -990:
                        v62 := GenS5{F0: v16, F1: v16, F2: (v5 + (v18 - 127))}
                        _ = v62
                    }
                    fmt.Println(v16)
                    if v1 {
                        v1 = (!true)
                    } else {
                        fmt.Println(v15, v28)
                        switch 32767 {
                        case 42:
                            if v1 {
                                var v63 int8 = v5
                                _ = v63
                            }
                        case 0:
                            for fi64 := 0; fi64 < 3; fi64++ {
                                v65 := GenS6{F0: v6, F1: v6, F2: 979}
                                var v66 uint8 = (v2 % ((42) | 1))
                                v16 = "hello"
                                _ = fi64
                                _ = v65
                                _ = v66
                            }
                        case -9223372036854775808:
                            v67 := GenS3{F0: 100, F1: 2147483647}
                            _ = v67
                        }
                    }
                }
                var v68 uint32 = 0
                _ = v29
                _ = v68
            } else {
                if (!(-656 <= 837)) {
                    var v69 string = ("foo bar" + (v16 + ((v16 + "") + v16)))
                    _ = v69
                }
            }
            _ = v28
        } else {
            v1 = ((v2 - v2) >= v2)
            v70 := GenS3{F0: v18, F1: 65535}
            var v71 int = 128
            _ = v70
            _ = v71
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
