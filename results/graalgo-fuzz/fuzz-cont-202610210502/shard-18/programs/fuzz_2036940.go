package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 int; F1 int32; F2 int32 }
type GenS7 struct { F0 int32 }
func genFn0(p0 int8) uint32 {
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{fi7, fi7, fi7}
        var v9 uint32 = 296
        var v10 uint32 = 38
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS1{F0: 379, F1: ("\n\t" + ("\n\t" + ("hello" + ("" + "a"))))}
        _ = fi11
        _ = v12
    }
    var v13 bool = (0 < 32767)
    switch 128 {
    case -276:
        p0 = (p0 % ((-128) | 1))
    }
    _ = p0
    _ = v13
    return 199
}

func main() {
    var v1 string = "a"
    v2 := make(map[string]uint32)
    v3 := []string{("hello" + v1), "\n\t", v1}
    var v4 uint32 = 65535
    var v5 int64 = -421
    v6 := make(map[string]bool)
    var v14 int8 = 100
    switch v1 {
    case "foo bar":
        v15 := v6[v1]
        _ = v15
    case "hello":
        switch 127 {
        case 42:
            var v16 int32 = -425
            _ = v16
        case -1:
            v17 := GenS2{F0: 9223372036854775807, F1: 0}
            _ = v17
        default:
            v4 += (v4 * (v4 - 42))
        }
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        fmt.Println(v14, v1, v5)
        _ = fi18
    }
    v14 += v14
    fmt.Println(v1, v4, v1)
    if true {
        var v19 uint32 = v4
        switch 1 {
        case -1:
            if (-2.25 >= 0.0) {
                v19 += v4
                v3 = append(v3, v1)
                if true {
                    fmt.Println(v19, v1, v19)
                    fmt.Println(v1)
                }
            }
        case 579:
            if (false && false) {
                v6["hello"] = (1 >= -466)
            } else {
                for fi20 := 0; fi20 < 16; fi20++ {
                    v4 -= 42
                    var v21 uint32 = (v19 % ((v4) | 1))
                    _ = fi20
                    _ = v21
                }
            }
        case 439:
            switch 32768 {
            case 128:
                var v22 uint8 = 0
                _ = v22
            case 42:
                v19 = (v19 << ((1) & 63))
            case 1:
                var v23 string = ""
                _ = v23
            }
        default:
            v19 += v19
        }
        switch ("" + ("a" + ((v1 + (v1 + v1)) + "\n\t"))) {
        case "foo bar":
            var v24 int = 354
            _ = v24
        }
        _ = v19
    } else {
        var v25 uint32 = 1
        _ = v25
    }
    switch v1 {
    case "foo bar":
        v26 := GenS3{F0: 42, F1: (^(v14))}
        _ = v26
    case "hello":
        fmt.Println(v14, v1, v4)
    case "":
        fmt.Println(v1, v4, v5)
    }
    if true {
        fmt.Println(v5, v14)
    } else {
        var v27 uint8 = 42
        for fi28 := 0; fi28 < 15; fi28++ {
            if (!true) {
                for fi29 := 0; fi29 < 15; fi29++ {
                    v4 += (v4 & 791)
                    _ = fi29
                }
                v30 := GenS2{F0: (^(fi28)), F1: 673}
                _ = v30
            } else {
                for fi31 := 0; fi31 < 17; fi31++ {
                    switch (fi31 >> ((4294967295) & 63)) {
                    case 32768:
                        v27 = v27
                    case 814:
                        v27 = uint8(v4)
                    case 128:
                        var v32 int64 = v5
                        _ = v32
                    default:
                        var v33 bool = true
                        _ = v33
                    }
                    _ = fi31
                }
                v34 := GenS4{F0: 0, F1: 0, F2: 126}
                _ = v34
            }
            v35 := GenS5{F0: 1, F1: 127}
            _ = fi28
            _ = v35
        }
        _ = v27
    }
    var v36 int16 = 1
    switch 32767 {
    case -108:
        if (32767 < int(v5)) {
            var v37 bool = false
            _ = v37
        } else {
            for fi38 := 0; fi38 < 7; fi38++ {
                for fi39 := 0; fi39 < 1; fi39++ {
                    var v40 uint32 = v4
                    if false {
                        v14 = v14
                        v5 -= (v5 - 0)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v41 := GenS1{F0: genFn0((v14 + (^(v14)))), F1: ("hello" + v1)}
                        _ = v41
                    }
                    _ = fi39
                    _ = v40
                }
                fmt.Println(v4, v14, v4)
                for fi42 := 0; fi42 < 6; fi42++ {
                    v43 := GenS5{F0: (^(v5)), F1: 1}
                    switch (-(fi38)) {
                    case 652:
                        var v44 float64 = ((1.5 / float64(1e10)) * float64((1e10 * 1.5)))
                        _ = v44
                    case 128:
                        v45 := GenS2{F0: -2147483648, F1: 329}
                        _ = v45
                    default:
                        var v46 int = (fi42 - (fi38 ^ 9223372036854775807))
                        _ = v46
                    }
                    v47 := GenS6{F0: 0, F1: 0, F2: -46}
                    _ = fi42
                    _ = v43
                    _ = v47
                }
                _ = fi38
            }
        }
    case -1:
        if false {
            if (-1 != 1) {
                if false {
                    switch 42 {
                    case 128:
                        var v48 uint = 120
                        _ = v48
                    }
                    fmt.Println(v14, v4, v5)
                    if true {
                        if false {
                            var v49 bool = (false || false)
                            _ = v49
                        }
                        var v50 int8 = (v14 / ((1) | 1))
                        var v51 int8 = (v14 % (((v14 * v50)) | 1))
                        _ = v50
                        _ = v51
                    } else {
                        var v52 uint64 = 0
                        v14 -= v14
                        var v53 uint32 = 65535
                        _ = v52
                        _ = v53
                    }
                } else {
                    v5 = 427
                }
            }
            var v54 int32 = -705
            for fi55 := 0; fi55 < 20; fi55++ {
                if ((v54 | (v54 >> ((364) & 63))) > (-(v54))) {
                    var v56 uint64 = 0
                    _ = v56
                }
                v57 := GenS2{F0: fi55, F1: 18446744073709551615}
                for fi58 := 0; fi58 < 16; fi58++ {
                    fmt.Println(fi55, v14, v14)
                    _ = fi58
                }
                _ = fi55
                _ = v57
            }
            _ = v54
        }
    case 128:
        var v59 uint64 = 42
        _ = v59
    }
    v14 = 100
    switch -873 {
    case -94:
        var v60 int16 = v36
        _ = v60
    case 128:
        if ((true && false) || true) {
            v61 := GenS6{F0: 127, F1: 168, F2: 42}
            if ((!(!true)) || false) {
                var v62 float64 = 0.1
                switch 537 {
                case -603:
                    switch 127 {
                    case -42:
                        v36 = -32768
                    case 127:
                        var v63 int64 = 9223372036854775807
                        _ = v63
                    }
                case 18:
                    v4 += genFn0(v14)
                default:
                    var v64 uint = 4294967295
                    _ = v64
                }
                _ = v62
            } else {
                v4 = (v4 ^ (^(v4)))
                for fi65 := 0; fi65 < 15; fi65++ {
                    v66 := GenS5{F0: -460, F1: 127}
                    v67 := GenS4{F0: 65535, F1: 2147483648, F2: 522}
                    v2[(v1 + "")] = genFn0((v14 * v14))
                    _ = fi65
                    _ = v66
                    _ = v67
                }
            }
            _ = v61
        } else {
            _ = len(v6)
        }
    case 986:
        for fi68 := 0; fi68 < 13; fi68++ {
            v36 += (^(v36))
            switch -127 {
            case 2147483647:
                v69 := v2[((v1 + ("a" + ((((v1 + (v1 + (v1 + v1))) + v1) + "a") + v1))) + (((v1 + "\n\t") + "\n\t") + ((("foo bar" + v1) + ("" + v1)) + "")))]
                _ = v69
            case 4294967295:
                v70 := GenS7{F0: -957}
                _ = v70
            default:
                v3 = append(v3, "")
            }
            _ = fi68
        }
    default:
        var v71 uint64 = 608
        _ = v71
    }
    for fi72 := 0; fi72 < 5; fi72++ {
        v1 = ""
        _ = fi72
    }
    v14 += (v14 % ((v14) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
