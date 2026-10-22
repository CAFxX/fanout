package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 bool }
func genFn0(p0 int) string {
    var v4 float32 = 1.5
    switch (("" + "hello") + "foo bar") {
    case "hello":
        switch "a" {
        case "":
            p0 = (-(p0))
        }
    case "\n\t":
        v5 := GenS1{F0: -101, F1: 153}
        _ = v5
    default:
        v6 := GenS2{F0: 0}
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint32 = 511
        var v9 string = "foo bar"
        if false {
            var v10 int = p0
            for fi11 := 0; fi11 < 1; fi11++ {
                v4 += float32((-2.25 / float64((v4 - 3.14159))))
                _ = fi11
            }
            var v12 bool = true
            _ = v10
            _ = v12
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v4
    return ("" + ("hello" + "\n\t"))
}
func genFn1() int {
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 uint8 = 0
        v15 := GenS3{F0: 291, F1: "hello"}
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 bool = (65535 <= 65535)
    switch 128 {
    case 32767:
        for fi17 := 0; fi17 < 4; fi17++ {
            var v18 uint32 = 42
            if ("foo bar" > "a") {
                var v19 uint16 = 480
                var v20 float64 = 0.1
                if (!v16) {
                    v20 = -2.25
                }
                _ = v19
                _ = v20
            }
            _ = fi17
            _ = v18
        }
    case -971:
        v16 = (1e10 >= ((-2.25 * 0.0) - 1.5))
    case -510:
        var v21 float64 = (1e10 * (((float64(-2.25) - 0.1) / 1e10) / ((1e10 / (0.0 - (0.0 - (3.14159 - ((3.14159 / 1e10) - 0.0))))) - -2.25)))
        _ = v21
    default:
        var v22 uint32 = 2147483648
        _ = v22
    }
    _ = v16
    return 9223372036854775807
}

func main() {
    var v1 uint8 = 135
    var v2 uint64 = 42
    var v3 int = 32767
    var v23 int64 = int64(v2)
    for fi24 := 0; fi24 < 11; fi24++ {
        if (((1 <= 24) || true) || (18446744073709551615 >= 138)) {
            v23 = (v23 / ((v23) | 1))
            v25 := GenS2{F0: v23}
            _ = v25
        } else {
            v2 += (^(v2))
            v26 := []int{genFn1(), int(v1), fi24, v3}
            var v27 int = 975
            _ = v26
            _ = v27
        }
        var v28 uint16 = 0
        switch "hello" {
        case "foo bar":
            for fi29 := 0; fi29 < 11; fi29++ {
                v30 := GenS3{F0: 2, F1: ("foo bar" + "foo bar")}
                if false {
                    var v31 string = ("foo bar" + "a")
                    if true {
                        fmt.Println(v1, v3, v31)
                    }
                    _ = v31
                } else {
                    var v32 bool = ((int(v1) <= (fi29 - (v3 / ((v3) | 1)))) && ((2 == v28) || true))
                    v32 = v32
                    _ = v32
                }
                _ = fi29
                _ = v30
            }
        case "hello":
            v28 -= 870
        case "":
            v33 := []int{fi24, v3}
            _ = v33
        }
        _ = fi24
        _ = v28
    }
    var v34 bool = true
    var v35 int = -2147483648
    v2 = v2
    v34 = false
    for fi36 := 0; fi36 < 12; fi36++ {
        for fi37 := 0; fi37 < 2; fi37++ {
            v38 := []int{genFn1(), (-(fi36))}
            switch 462 {
            case 2147483647:
                v39 := GenS4{F0: 398}
                _ = v39
            default:
                if len(v38) >= 2 {
                    v38 = v38[1:len(v38)]
                }
            }
            v40 := GenS1{F0: 0, F1: 1}
            _ = fi37
            _ = v38
            _ = v40
        }
        _ = fi36
    }
    var v41 uint8 = 1
    if true {
        for fi42 := 0; fi42 < 7; fi42++ {
            if (0.1 == 1.5) {
                if true {
                    var v43 uint = 18446744073709551615
                    for fi44 := 0; fi44 < 2; fi44++ {
                        if true {
                            fmt.Println(v3, fi44, v35)
                        } else {
                            v41 = (v1 >> (((v43 ^ v43)) & 63))
                            v2 -= (v2 >> (((v43 ^ (v43 + 563))) & 63))
                            v45 := GenS2{F0: v23}
                            _ = v45
                        }
                        fmt.Println(v35, v34, v34)
                        v46 := GenS5{F0: 4294967295, F1: 2, F2: (v23 / ((724) | 1))}
                        _ = fi44
                        _ = v46
                    }
                    _ = v43
                } else {
                    v47 := []int{(v35 << ((0) & 63)), genFn1()}
                    fmt.Println(v23)
                    var v48 string = (("foo bar" + "foo bar") + "hello")
                    _ = v47
                    _ = v48
                }
                v2 -= v2
            }
            fmt.Println(fi42, v35)
            _ = fi42
        }
    }
    for fi49 := 0; fi49 < 13; fi49++ {
        if v34 {
            var v50 uint64 = v2
            v41 = (v41 / ((v41) | 1))
            _ = v50
        }
        for fi51 := 0; fi51 < 3; fi51++ {
            if v34 {
                v41 = (^(v41))
                var v52 uint8 = 214
                fmt.Println(fi51, fi49)
                var v53 int32 = 1
                v1 -= 255
                v23 += v23
                v54 := []int{fi51, (fi49 >> ((42) & 63))}
                _ = v52
                _ = v53
                _ = v54
            } else {
                if true {
                    v2 += v2
                } else {
                    var v55 uint64 = 42
                    v35 = fi51
                    var v56 float32 = 1e10
                    if (!(1 < 979)) {
                        v57 := GenS1{F0: 0, F1: 42}
                        switch fi51 {
                        case 32767:
                            fmt.Println(v3, v1)
                        case -877:
                            if v34 {
                                v58 := GenS6{F0: uint16(v41), F1: v35}
                                _ = v58
                            }
                        case -9223372036854775808:
                            fmt.Println(v34, fi51)
                        default:
                            v41 = 127
                        }
                        switch "hello" {
                        case "a":
                            var v59 int32 = 509
                            _ = v59
                        }
                        _ = v57
                    } else {
                        if true {
                            if ((1.5 / (3.14159 * 1e10)) != ((((float64(v56) / 1.5) + (float64((v56 + 0.0)) + (-2.25 - (float64(v56) - float64(float32((0.1 + (float64(v56) * 0.1)))))))) - 1.5) - 1.5)) {
                                var v60 float32 = (0.1 + (float32(0.0) * v56))
                                v61 := GenS2{F0: 1}
                                if v34 {
                                    fmt.Println(v41)
                                } else {
                                    var v62 int = (v3 | v3)
                                    var v63 int32 = 42
                                    var v64 uint = 90
                                    _ = v62
                                    _ = v63
                                    _ = v64
                                }
                                _ = v60
                                _ = v61
                            }
                            if v34 {
                                switch "foo bar" {
                                case "a":
                                    v65 := []int{(fi51 - genFn1()), genFn1()}
                                    _ = v65
                                default:
                                    var v66 uint = 0
                                    _ = v66
                                }
                            }
                            var v67 uint8 = v41
                            _ = v67
                        }
                        var v68 uint16 = 2
                        v35 -= fi49
                        _ = v68
                    }
                    _ = v55
                    _ = v56
                }
                fmt.Println(v2, v41, v23)
            }
            var v69 int64 = (v23 ^ v23)
            _ = fi51
            _ = v69
        }
        _ = fi49
    }
    v34 = v34
    fmt.Println(v23, v2, v35)
    if false {
        v3 = (v3 - (v3 >> ((4294967295) & 63)))
        switch (v3 * (v35 & int(v41))) {
        case -9223372036854775808:
            for fi70 := 0; fi70 < 18; fi70++ {
                for fi71 := 0; fi71 < 1; fi71++ {
                    if (0 >= 65535) {
                        v72 := GenS6{F0: 32767, F1: v35}
                        var v73 int32 = -1
                        v34 = v34
                        _ = v72
                        _ = v73
                    }
                    var v74 float32 = float32(-2.25)
                    _ = fi71
                    _ = v74
                }
                v35 = genFn1()
                v1 = (^(v1))
                _ = fi70
            }
        case 0:
            for fi75 := 0; fi75 < 2; fi75++ {
                fmt.Println(v3)
                var v76 int64 = v23
                var v77 uint32 = 42
                _ = fi75
                _ = v76
                _ = v77
            }
        case 127:
            fmt.Println(v41, v35)
        default:
            v78 := GenS3{F0: 974, F1: ("a" + ((("" + "hello") + ("foo bar" + "foo bar")) + "a"))}
            _ = v78
        }
    }
    v79 := GenS7{F0: v34}
    fmt.Println(v41, v41)
    fmt.Println(v41, v23)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v34
    _ = v35
    _ = v41
    _ = v79
}
