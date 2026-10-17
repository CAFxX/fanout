package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 int64 }
type GenS8 struct { F0 uint64; F1 int8 }
type GenS9 struct { F0 uint8; F1 uint32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int16 }
func genFn0(p0 int32) int32 {
    v6 := []int{1, -600, 32767, 260}
    for fi7 := 0; fi7 < 9; fi7++ {
        p0 = 0
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS1{F0: 42, F1: 819, F2: -9223372036854775808}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    fmt.Println(p0)
    if false {
        for fi10 := 0; fi10 < 13; fi10++ {
            p0 = p0
            switch fi10 {
            case -1:
                var v11 bool = false
                _ = v11
            case 2147483647:
                for fi12 := 0; fi12 < 4; fi12++ {
                    var v13 float64 = (0.0 - 0.0)
                    if ((fi12 >> ((255) & 63)) > 4294967295) {
                        var v14 int16 = 0
                        if len(v6) > 0 {
                            v6[0] = fi10
                        }
                        var v15 int = (fi12 % (((fi12 % (((fi10 / ((32768) | 1))) | 1))) | 1))
                        _ = v14
                        _ = v15
                    } else {
                        var v16 bool = (42 < 127)
                        v17 := GenS1{F0: 255, F1: 1, F2: 163}
                        _ = v16
                        _ = v17
                    }
                    if true {
                        v6 = append(v6, fi10)
                        p0 = p0
                        var v18 float32 = 1e10
                        var v19 uint = 42
                        _ = v18
                        _ = v19
                    }
                    _ = fi12
                    _ = v13
                }
            }
            if (42 < 163) {
                p0 += (p0 - p0)
            } else {
                v20 := GenS1{F0: 904, F1: 296, F2: 42}
                var v21 int32 = 654
                _ = v20
                _ = v21
            }
            _ = fi10
        }
        for fi22 := 0; fi22 < 1; fi22++ {
            for fi23 := 0; fi23 < 3; fi23++ {
                var v24 int16 = -89
                v6 = append(v6, fi22)
                _ = fi23
                _ = v24
            }
            var v25 uint64 = 0
            v25 -= 2
            _ = fi22
            _ = v25
        }
    } else {
        var v26 uint32 = 2147483648
        _ = v26
    }
    _ = p0
    _ = v6
    return (-(p0))
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := []int{560}
    var v3 float64 = 0.1
    var v4 uint8 = 128
    var v5 int32 = 1
    if (false && true) {
        v27 := GenS2{F0: 4294967295}
        _ = v27
    }
    var v28 uint64 = uint64(v4)
    v29 := GenS2{F0: 345}
    fmt.Println(v3)
    v30 := GenS1{F0: v1, F1: (v1 << ((v1) & 63)), F2: 869}
    var v31 int16 = -820
    var v32 int16 = (^(v31))
    v33 := GenS3{F0: (v28 << (((v1 - 814)) & 63)), F1: 65535}
    var v34 int64 = 42
    for fi35 := 0; fi35 < 5; fi35++ {
        v36 := GenS3{F0: 0, F1: 65535}
        v28 -= v28
        for fi37 := 0; fi37 < 15; fi37++ {
            v38 := GenS4{F0: 42, F1: (fi37 | 2147483647)}
            v39 := GenS5{F0: false, F1: "a"}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v32 += (v32 % ((32767) | 1))
            v32 = v32
            if true {
                v34 += (v34 << (((v1 << ((v1) & 63))) & 63))
                fmt.Println(v4, v5, v5)
                fmt.Println(v5, fi35, v5)
            }
            _ = fi37
            _ = v38
            _ = v39
        }
        _ = fi35
        _ = v36
    }
    v5 = (v5 % ((-2147483648) | 1))
    var v40 int = int(v4)
    var v41 uint16 = 304
    switch (v40 - v40) {
    case 42:
        v2 = append(v2, v40)
    case -425:
        var v42 float64 = 1.5
        _ = v42
    case -2147483648:
        var v43 int32 = genFn0((v5 >> (((v1 << (((v1 >> (((v1 / (((v1 % ((v1) | 1))) | 1))) & 63))) & 63))) & 63)))
        _ = v43
    default:
        v28 = v28
    }
    for fi44 := 0; fi44 < 3; fi44++ {
        v45 := GenS2{F0: 776}
        _ = fi44
        _ = v45
    }
    if true {
        v46 := GenS4{F0: 1, F1: (v40 << ((255) & 63))}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = len(v2)
        _ = v46
    }
    for fi47 := 0; fi47 < 5; fi47++ {
        v48 := GenS6{F0: (v5 << ((0) & 63))}
        if (!false) {
            var v49 bool = true
            if (v49 && (v34 >= v34)) {
                fmt.Println(v1)
                v34 = int64(v31)
            } else {
                for fi50 := 0; fi50 < 5; fi50++ {
                    v51 := GenS7{F0: (v4 % (((v4 >> ((v1) & 63))) | 1)), F1: (^(v34))}
                    v5 = genFn0(v5)
                    v2 = append(v2, 128)
                    v52 := GenS8{F0: 18446744073709551615, F1: 100}
                    v31 -= -32768
                    _ = fi50
                    _ = v51
                    _ = v52
                }
                v34 = (v34 | v34)
            }
            _ = len(v2)
            _ = v49
        } else {
            if (!((-128 > 1) || ((!(!((1e10 <= float32(v3)) && (float64((0.1 / (0.1 - 3.14159))) != (v3 - -2.25))))) || true))) {
                var v53 int = fi47
                for fi54 := 0; fi54 < 1; fi54++ {
                    v55 := GenS9{F0: 255, F1: 832}
                    var v56 int32 = v5
                    if false {
                        var v57 int64 = (v34 + (v34 >> (((v1 + (^(v1)))) & 63)))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        switch v53 {
                        case 128:
                            fmt.Println(v57, v40)
                        case -1:
                            v4 -= (v4 / ((uint8(v28)) | 1))
                            v41 = (v41 & (v41 + (v41 * (^(v41)))))
                            v2 = append(v2, (fi54 + fi54))
                        }
                        if (v1 == 65535) {
                            _ = len(v2)
                            if (!(-693 > (v57 / (((v57 << ((v1) & 63))) | 1)))) {
                                if false {
                                    if (0 > 8) {
                                        v2 = append(v2, 127)
                                        v57 += (v34 % ((42) | 1))
                                        v2 = append(v2, (-(v53)))
                                    } else {
                                        v58 := GenS10{F0: v32}
                                        _ = v58
                                    }
                                }
                                var v59 uint32 = 2147483648
                                _ = v59
                            } else {
                                v60 := GenS4{F0: 2, F1: v53}
                                if (!true) {
                                    switch "" {
                                    case "a":
                                        fmt.Println(v57)
                                    }
                                    v41 = (v41 | (v41 + v41))
                                    v61 := GenS7{F0: 1, F1: (v34 | v57)}
                                    _ = v61
                                }
                                _ = v60
                            }
                            var v62 float32 = 0.0
                            _ = v62
                        }
                        _ = v57
                    } else {
                        fmt.Println(v56, v5, v40)
                        var v63 int16 = v32
                        _ = v63
                    }
                    _ = fi54
                    _ = v55
                    _ = v56
                }
                _ = v53
            } else {
                fmt.Println(v41)
                var v64 int64 = -9223372036854775808
                v65 := GenS11{F0: -636}
                _ = v64
                _ = v65
            }
            if (v32 > -32768) {
                fmt.Println(v28, v40)
                for fi66 := 0; fi66 < 12; fi66++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi66
                }
            }
            var v67 int32 = 1
            _ = v67
        }
        var v68 int64 = (v34 >> ((v1) & 63))
        _ = fi47
        _ = v48
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
