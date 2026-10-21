package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int }
type GenS2 struct { F0 int8; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32; F1 bool }
func genFn0(p0 uint32, p1 string) bool {
    if false {
        p0 -= (p0 / ((p0) | 1))
        for fi4 := 0; fi4 < 5; fi4++ {
            v5 := []int{(fi4 & (-(fi4))), fi4, (fi4 | 4294967295)}
            var v6 bool = (179 == 32767)
            _ = fi4
            _ = v5
            _ = v6
        }
    } else {
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := []int{0, fi7, fi7, fi7}
            _ = fi7
            _ = v8
        }
    }
    p0 += p0
    if ((-432 >= 0) || (-1 > -9223372036854775808)) {
        switch "\n\t" {
        case "hello":
            for fi9 := 0; fi9 < 15; fi9++ {
                if true {
                    v10 := GenS1{F0: 0, F1: 0, F2: fi9}
                    var v11 uint16 = 0
                    _ = v10
                    _ = v11
                }
                _ = fi9
            }
        case "\n\t":
            for fi12 := 0; fi12 < 9; fi12++ {
                v13 := GenS2{F0: 1, F1: false, F2: p0}
                v14 := GenS2{F0: 127, F1: ((p0 / (((p0 << ((65535) & 63))) | 1)) <= (p0 | p0)), F2: (p0 / ((4294967295) | 1))}
                v15 := []int{}
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
        default:
            p1 = (p1 + ((p1 + ((p1 + ("hello" + (p1 + ""))) + p1)) + (p1 + (p1 + ("hello" + p1)))))
        }
        v16 := GenS1{F0: 886, F1: -1, F2: -220}
        var v17 float32 = float32(((-2.25 * float64((((1.5 / 1.0) + float32(3.14159)) * 1.5))) + 0.1))
        _ = v16
        _ = v17
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return true
}
func genFn1() uint32 {
    v18 := GenS3{F0: 42}
    v19 := []int{2147483647, 127, 32767, 9223372036854775807}
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 bool = ("a" < ("" + "a"))
        v21 = v21
        v22 := GenS3{F0: 90}
        _ = fi20
        _ = v21
        _ = v22
    }
    switch "foo bar" {
    case "hello":
        for fi23 := 0; fi23 < 4; fi23++ {
            if true {
                var v24 uint = 413
                _ = v24
            }
            _ = fi23
        }
    case "\n\t":
        for fi25 := 0; fi25 < 16; fi25++ {
            var v26 uint = 0
            _ = fi25
            _ = v26
        }
    case "":
        if (false && (255 != 0)) {
            for fi27 := 0; fi27 < 12; fi27++ {
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                for fi28 := 0; fi28 < 14; fi28++ {
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                    v19 = append(v19, (fi28 - -9223372036854775808))
                    var v29 uint32 = 4294967295
                    _ = fi28
                    _ = v29
                }
                v19 = append(v19, fi27)
                _ = fi27
            }
        } else {
            var v30 float64 = 0.0
            _ = len(v19)
            _ = v30
        }
    default:
        var v31 int8 = -55
        _ = v31
    }
    _ = v18
    _ = v19
    return 42
}

func main() {
    var v1 bool = true
    var v2 uint64 = 623
    var v3 int8 = 0
    var v32 int32 = -1
    if v1 {
        v33 := GenS4{F0: (v2 <= (v2 * v2))}
        _ = v33
    }
    v34 := GenS1{F0: v2, F1: -117, F2: -2147483648}
    if (false || ((v32 / (((v32 & (v32 * v32))) | 1)) != (v32 << ((52) & 63)))) {
        v2 += 4294967295
    } else {
        v35 := []int{183, 32767, 666, 42}
        for fi36 := 0; fi36 < 20; fi36++ {
            for fi37 := 0; fi37 < 3; fi37++ {
                v1 = (-32768 >= 0)
                switch fi37 {
                case 857:
                    var v38 uint = 2
                    _ = v38
                default:
                    if (!(984 != v2)) {
                        v1 = false
                        v35 = append(v35, fi37)
                        if (v1 || v1) {
                            v3 -= (v3 - v3)
                            v39 := GenS2{F0: (^(v3)), F1: ((v32 ^ (v32 | (v32 << ((1) & 63)))) <= v32), F2: 2147483647}
                            _ = v39
                        }
                    } else {
                        v32 += 0
                        if len(v35) > 0 {
                            v35[0] = fi36
                        }
                        v40 := GenS5{F0: 637, F1: true}
                        _ = v40
                    }
                }
                _ = fi37
            }
            _ = fi36
        }
        switch -905 {
        case 984:
            v2 = v2
        }
        _ = v35
    }
    v1 = (!true)
    var v41 bool = true
    v42 := GenS5{F0: 0, F1: v1}
    var v43 uint = 65535
    fmt.Println(v41)
    for fi44 := 0; fi44 < 8; fi44++ {
        v32 = (-(v32))
        if (!v1) {
            v45 := []int{fi44, fi44, 825, (fi44 % ((128) | 1))}
            switch 0 {
            case -2147483648:
                for fi46 := 0; fi46 < 14; fi46++ {
                    var v47 uint8 = 127
                    if len(v45) > 0 {
                        v45[0] = fi46
                    }
                    _ = fi46
                    _ = v47
                }
            case 941:
                v32 = -1
            case 42:
                v48 := GenS1{F0: v2, F1: 42, F2: (fi44 % ((fi44) | 1))}
                _ = v48
            }
            for fi49 := 0; fi49 < 7; fi49++ {
                fmt.Println(v3)
                _ = fi49
            }
            _ = v45
        }
        if false {
            for fi50 := 0; fi50 < 17; fi50++ {
                var v51 uint16 = 42
                v43 += v43
                _ = fi50
                _ = v51
            }
            switch "" {
            case "hello":
                if false {
                    v52 := GenS6{F0: v43}
                    if false {
                        if v1 {
                            switch "hello" {
                            case "foo bar":
                                var v53 int32 = int32(v2)
                                _ = v53
                            case "hello":
                                v54 := []int{(fi44 * fi44), (fi44 * (-(fi44))), (fi44 + (fi44 % ((fi44) | 1))), 4294967295}
                                _ = v54
                            }
                            if true {
                                v43 += v43
                                v55 := []int{(-(fi44)), (fi44 | (fi44 * (fi44 | fi44))), fi44}
                                var v56 int64 = 0
                                _ = v55
                                _ = v56
                            }
                            if (false && true) {
                                v57 := []int{-9223372036854775808, fi44}
                                _ = v57
                            } else {
                                fmt.Println(v2, v43, fi44)
                            }
                        } else {
                            var v58 string = ("foo bar" + (("" + "foo bar") + "a"))
                            _ = v58
                        }
                        v59 := []int{fi44, 4294967295, 4294967295, fi44}
                        _ = v59
                    }
                    _ = v52
                } else {
                    fmt.Println(v41)
                }
            default:
                v60 := GenS7{F0: fi44}
                _ = v60
            }
        } else {
            if v41 {
                v1 = (1 > (fi44 % (((fi44 / ((9223372036854775807) | 1))) | 1)))
            }
            fmt.Println(v32, v43, v43)
            var v61 int = (^(fi44))
            _ = v61
        }
        _ = fi44
    }
    if (((v1 && true) || v1) || (v32 != 0)) {
        v62 := []int{925, 216}
        _ = v62
    } else {
        for fi63 := 0; fi63 < 12; fi63++ {
            for fi64 := 0; fi64 < 6; fi64++ {
                if v1 {
                    v65 := GenS5{F0: 1, F1: (1000 < 1)}
                    if false {
                        var v66 float64 = ((((((-2.25 * 0.1) + 0.1) * -2.25) * 0.1) / 1.5) * 3.14159)
                        _ = v66
                    }
                    if (-1 <= (-(v32))) {
                        fmt.Println(v32, v43)
                    }
                    _ = v65
                }
                switch "" {
                case "hello":
                    v67 := []int{fi64, fi64, fi63, (fi63 - 32767)}
                    _ = v67
                default:
                    v68 := GenS8{F0: 2147483648, F1: (!(v3 == -1))}
                    _ = v68
                }
                v43 -= v43
                _ = fi64
            }
            switch (fi63 ^ (fi63 ^ (fi63 / (((fi63 + 9223372036854775807)) | 1)))) {
            case 9223372036854775807:
                for fi69 := 0; fi69 < 11; fi69++ {
                    if v41 {
                        v32 = (^(v32))
                        v3 += (^(v3))
                        v70 := []int{(fi69 - fi69), int(v32)}
                        _ = v70
                    }
                    _ = fi69
                }
            }
            if v1 {
                fmt.Println(v2, v43)
            }
            _ = fi63
        }
        for fi71 := 0; fi71 < 15; fi71++ {
            v72 := []int{(fi71 - (fi71 * fi71))}
            _ = fi71
            _ = v72
        }
        var v73 uint64 = v2
        _ = v73
    }
    v74 := GenS4{F0: (v32 <= 201)}
    var v75 uint8 = 2
    fmt.Println(v41, v41, v43)
    var v76 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v41
    _ = v42
    _ = v43
    _ = v74
    _ = v75
    _ = v76
}
