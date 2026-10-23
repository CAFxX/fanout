package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 bool; F2 int32 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int8; F1 bool; F2 int32 }
type GenS6 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS7 struct { F0 int; F1 uint32 }
type GenS8 struct { F0 int16; F1 int }
type GenS9 struct { F0 uint; F1 bool; F2 int32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint8; F1 int32 }
func genFn0() uint64 {
    if true {
        for fi7 := 0; fi7 < 4; fi7++ {
            if false {
                v8 := GenS1{F0: 42, F1: "hello", F2: 42}
                _ = v8
            } else {
                var v9 int = fi7
                if (((0 < 42) || (((0.1 - (float32(-2.25) - 3.14159)) + (float32(((0.1 + 1e10) + 1e10)) + 0.0)) > float32(1.5))) && true) {
                    var v10 uint16 = 2
                    fmt.Println(fi7, v10)
                    _ = v10
                }
                _ = v9
            }
            _ = fi7
        }
    }
    switch 32768 {
    case 2147483647:
        if false {
            v11 := []int{9223372036854775807, 4294967295}
            for fi12 := 0; fi12 < 15; fi12++ {
                if true {
                    fmt.Println(fi12)
                    switch "" {
                    case "hello":
                        var v13 int16 = 1
                        _ = v13
                    case "foo bar":
                        var v14 float32 = 3.14159
                        _ = v14
                    }
                } else {
                    var v15 float32 = 0.0
                    for fi16 := 0; fi16 < 19; fi16++ {
                        var v17 uint16 = 862
                        v15 = float32(0.1)
                        var v18 uint32 = 0
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    _ = v15
                }
                if (true || true) {
                    fmt.Println(fi12)
                    if true {
                        v11 = append(v11, (fi12 << ((65535) & 63)))
                        for fi19 := 0; fi19 < 1; fi19++ {
                            var v20 string = "hello"
                            _ = len(v11)
                            switch 128 {
                            case 32768:
                                var v21 int64 = -1
                                _ = v21
                            case 128:
                                var v22 uint32 = 399
                                _ = v22
                            }
                            _ = fi19
                            _ = v20
                        }
                        var v23 uint = 255
                        _ = v23
                    }
                }
                _ = fi12
            }
            var v24 int32 = 0
            _ = v11
            _ = v24
        }
    }
    var v25 uint = 18446744073709551615
    _ = v25
    return 42
}
func genFn1(p0 uint64) uint64 {
    var v26 uint = 1
    fmt.Println(v26, p0)
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 uint32 = 2
        _ = fi27
        _ = v28
    }
    p0 = (p0 >> ((0) & 63))
    _ = p0
    _ = v26
    return (p0 - p0)
}

func main() {
    var v1 uint16 = 0
    var v2 int16 = 32767
    var v3 uint8 = 127
    var v4 uint16 = (v1 >> ((701) & 63))
    var v5 uint32 = 42
    var v6 uint32 = 1
    for fi29 := 0; fi29 < 1; fi29++ {
        v4 -= (^(v4))
        _ = fi29
    }
    var v30 string = ""
    switch 127 {
    case -9223372036854775808:
        if false {
            v4 += (v4 & v4)
            v31 := []int{-9223372036854775808, 32768, 9223372036854775807}
            fmt.Println(v2)
            _ = v31
        } else {
            v32 := []int{}
            for fi33 := 0; fi33 < 3; fi33++ {
                var v34 int32 = -65
                _ = fi33
                _ = v34
            }
            switch v30 {
            case "a":
                v6 = uint32(v1)
            case "":
                var v35 int32 = -1
                _ = v35
            case "foo bar":
                v2 = 1000
            default:
                for fi36 := 0; fi36 < 1; fi36++ {
                    var v37 uint16 = (^(v1))
                    switch ("\n\t" + (v30 + v30)) {
                    case "\n\t":
                        fmt.Println(v3)
                    case "hello":
                        v4 = (v4 | 255)
                    case "":
                        v38 := GenS2{F0: 106}
                        _ = v38
                    }
                    _ = fi36
                    _ = v37
                }
            }
            _ = v32
        }
    case 128:
        var v39 uint8 = 0
        _ = v39
    case 695:
        v40 := []int{123, 2147483647, 4294967295, 9223372036854775807}
        _ = v40
    default:
        v41 := GenS3{F0: 2147483647, F1: true, F2: -1}
        _ = v41
    }
    v42 := GenS4{F0: int(v2), F1: int16(v1)}
    if false {
        for fi43 := 0; fi43 < 17; fi43++ {
            v44 := []int{-874}
            _ = fi43
            _ = v44
        }
        v45 := GenS4{F0: -748, F1: (-(v2))}
        v46 := GenS2{F0: uint16(v6)}
        _ = v45
        _ = v46
    }
    v4 = v1
    v2 = v2
    var v47 bool = true
    for fi48 := 0; fi48 < 6; fi48++ {
        v49 := GenS2{F0: (v1 % ((uint16(v3)) | 1))}
        _ = fi48
        _ = v49
    }
    for fi50 := 0; fi50 < 10; fi50++ {
        for fi51 := 0; fi51 < 12; fi51++ {
            v52 := []int{fi51, (fi51 + int(v1)), (fi50 >> ((255) & 63))}
            v3 -= 128
            var v53 int32 = -2147483648
            _ = fi51
            _ = v52
            _ = v53
        }
        v54 := []int{}
        v55 := GenS2{F0: v4}
        _ = fi50
        _ = v54
        _ = v55
    }
    var v56 string = v30
    v2 = (v2 << ((18446744073709551615) & 63))
    v57 := GenS5{F0: 56, F1: v47, F2: 0}
    switch 9223372036854775807 {
    case 0:
        switch -2147483648 {
        case -1:
            for fi58 := 0; fi58 < 5; fi58++ {
                if v47 {
                    var v59 uint8 = (v3 * (v3 & v3))
                    var v60 uint = 65535
                    _ = v59
                    _ = v60
                } else {
                    v61 := GenS5{F0: int8(v6), F1: v47, F2: 958}
                    _ = v61
                }
                _ = fi58
            }
        default:
            v62 := []int{}
            _ = v62
        }
    case -74:
        var v63 string = v56
        _ = v63
    }
    for fi64 := 0; fi64 < 6; fi64++ {
        v65 := []int{fi64, (fi64 ^ 708), fi64}
        _ = fi64
        _ = v65
    }
    if (!((v47 || v47) || v47)) {
        if v47 {
            v66 := GenS6{F0: 42, F1: v3, F2: (v6 >> ((396) & 63))}
            _ = v66
        }
        var v67 uint = 65535
        v68 := []int{-9223372036854775808, -2147483648, 127, 128}
        _ = v67
        _ = v68
    } else {
        v69 := GenS6{F0: 529, F1: v3, F2: v6}
        _ = v69
    }
    for fi70 := 0; fi70 < 8; fi70++ {
        if ((v47 && false) || ((0.1 + 0.1) < 3.14159)) {
            v2 += 1
            if (-2.25 < 0.0) {
                switch (fi70 + (fi70 << ((2) & 63))) {
                case 909:
                    if (42 > genFn1(4294967295)) {
                        var v71 int32 = -296
                        v72 := GenS7{F0: (^(fi70)), F1: v6}
                        _ = v71
                        _ = v72
                    } else {
                        if (0.1 > -2.25) {
                            fmt.Println(v47)
                        } else {
                            var v73 int8 = 127
                            fmt.Println(v1, v30)
                            _ = v73
                        }
                        fmt.Println(v3, v1)
                    }
                case -9223372036854775808:
                    for fi74 := 0; fi74 < 13; fi74++ {
                        v5 -= 184
                        v4 = 0
                        v2 -= (v2 % ((-32768) | 1))
                        _ = fi74
                    }
                case -2147483648:
                    if v47 {
                        v75 := []int{int(v6), (fi70 ^ fi70), fi70, -9223372036854775808}
                        var v76 int8 = -128
                        if (((v30 + "\n\t") + v56) <= "hello") {
                            for fi77 := 0; fi77 < 17; fi77++ {
                                v75 = append(v75, fi70)
                                v78 := GenS8{F0: -1, F1: 2147483647}
                                _ = fi77
                                _ = v78
                            }
                            for fi79 := 0; fi79 < 4; fi79++ {
                                v2 = v2
                                _ = fi79
                            }
                        } else {
                            v5 = v5
                            v75 = append(v75, 2147483647)
                        }
                        _ = v75
                        _ = v76
                    } else {
                        v2 = v2
                        switch v30 {
                        case "hello":
                            if v47 {
                                var v80 int8 = 28
                                v81 := []int{(fi70 + (fi70 % (((-(fi70))) | 1))), fi70, (^(fi70)), 32768}
                                v3 += 2
                                _ = v80
                                _ = v81
                            } else {
                                switch (fi70 % ((1) | 1)) {
                                case 1:
                                    var v82 float64 = 0.1
                                    _ = v82
                                case 521:
                                    v83 := GenS9{F0: 2, F1: ((v2 / ((v2) | 1)) <= v2), F2: -1}
                                    _ = v83
                                case 379:
                                    v3 += (^(v3))
                                }
                            }
                        default:
                            switch fi70 {
                            case -2147483648:
                                v84 := GenS10{F0: fi70}
                                _ = v84
                            case 127:
                                v85 := []int{-9223372036854775808, -1}
                                _ = v85
                            }
                        }
                        v47 = (!((1.5 * ((1e10 + 1.5) + 0.1)) >= 1.5))
                    }
                }
                v86 := []int{}
                for fi87 := 0; fi87 < 9; fi87++ {
                    v88 := GenS11{F0: (v3 << ((1) & 63)), F1: -638}
                    _ = fi87
                    _ = v88
                }
                _ = v86
            } else {
                v89 := GenS4{F0: 1, F1: 1000}
                _ = v89
            }
            if (!(v30 < v30)) {
                switch ("\n\t" + (v30 + v56)) {
                case "foo bar":
                    v90 := GenS8{F0: (v2 >> ((963) & 63)), F1: 42}
                    _ = v90
                }
            }
        } else {
            var v91 uint32 = (v6 * 954)
            _ = v91
        }
        _ = fi70
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v42
    _ = v47
    _ = v5
    _ = v56
    _ = v57
    _ = v6
}
