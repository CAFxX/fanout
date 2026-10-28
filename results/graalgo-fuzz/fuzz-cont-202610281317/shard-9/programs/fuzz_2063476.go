package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int; F2 int8 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 int; F1 int64 }
func genFn0() uint32 {
    var v6 uint8 = 2
    if true {
        if true {
            var v7 uint32 = 375
            switch ("foo bar" + ("hello" + "")) {
            case "hello":
                if false {
                    if (65535 <= 42) {
                        var v8 int32 = 159
                        for fi9 := 0; fi9 < 18; fi9++ {
                            v7 = v7
                            fmt.Println(v7, v7)
                            _ = fi9
                        }
                        _ = v8
                    } else {
                        if false {
                            for fi10 := 0; fi10 < 1; fi10++ {
                                v7 += v7
                                var v11 bool = ((v7 >> ((46) & 63)) > v7)
                                v11 = (1 <= fi10)
                                _ = fi10
                                _ = v11
                            }
                            var v12 int64 = -953
                            _ = v12
                        } else {
                            v13 := GenS1{F0: -9223372036854775808}
                            _ = v13
                        }
                        v14 := GenS2{F0: 1, F1: -1, F2: 42}
                        for fi15 := 0; fi15 < 12; fi15++ {
                            var v16 uint = 2
                            _ = fi15
                            _ = v16
                        }
                        _ = v14
                    }
                    var v17 bool = (18446744073709551615 == 0)
                    for fi18 := 0; fi18 < 9; fi18++ {
                        v6 = 255
                        var v19 uint8 = v6
                        switch "hello" {
                        case "hello":
                            var v20 float32 = 0.0
                            _ = v20
                        case "foo bar":
                            v21 := []int{fi18, -1, (fi18 % (((fi18 * fi18)) | 1)), 92}
                            _ = v21
                        default:
                            if v17 {
                                v17 = v17
                                fmt.Println(v7, v6)
                            }
                        }
                        _ = fi18
                        _ = v19
                    }
                    _ = v17
                } else {
                    v22 := GenS3{F0: 42, F1: 1}
                    v7 -= (v7 * v7)
                    _ = v22
                }
            default:
                var v23 int = 4294967295
                _ = v23
            }
            _ = v7
        } else {
            for fi24 := 0; fi24 < 2; fi24++ {
                for fi25 := 0; fi25 < 18; fi25++ {
                    switch fi24 {
                    case 0:
                        switch "" {
                        case "hello":
                            var v26 int16 = 0
                            _ = v26
                        }
                    }
                    var v27 string = "hello"
                    _ = fi25
                    _ = v27
                }
                _ = fi24
            }
            v28 := []int{9, 42, 40, 203}
            _ = v28
        }
    }
    _ = v6
    return 388
}
func genFn1(p0 bool, p1 uint64) int16 {
    p0 = p0
    p1 = (^(p1))
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{}
    v3 := make(map[int]bool)
    var v4 uint8 = 42
    var v5 uint16 = 850
    _ = len(v3)
    if false {
        fmt.Println(v4, v4)
    } else {
        v5 = v5
        v29 := v3[128]
        var v30 int = -82
        _ = v29
        _ = v30
    }
    v31 := GenS2{F0: 125, F1: -9223372036854775808, F2: 100}
    switch "a" {
    case "foo bar":
        var v32 uint = 4294967295
        _ = v32
    case "\n\t":
        v4 = v4
    default:
        v5 += (^(v5))
    }
    v5 += (^(v5))
    switch "hello" {
    case "foo bar":
        switch -777 {
        case -9223372036854775808:
            if len(v2) > 0 {
                v2[0] = -697
            }
        case 42:
            var v33 int8 = -30
            _ = v33
        }
    case "a":
        v4 = v4
    case "hello":
        v34 := GenS4{F0: 127, F1: -9223372036854775808}
        _ = v34
    }
    if false {
        fmt.Println(v4)
        var v35 uint32 = 65535
        var v36 uint32 = genFn0()
        _ = v35
        _ = v36
    }
    var v37 int8 = -128
    if (true && (false && ((1e10 / 1.0) <= 0.1))) {
        v4 = v4
    } else {
        fmt.Println(v37, v4)
        var v38 uint = 730
        _ = v38
    }
    for fi39 := 0; fi39 < 12; fi39++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v40 int32 = 1
        v2 = append(v2, fi39)
        _ = fi39
        _ = v40
    }
    var v41 int16 = 42
    if (v41 >= genFn1(false, 18446744073709551615)) {
        var v42 uint64 = 425
        for fi43 := 0; fi43 < 15; fi43++ {
            if false {
                if true {
                    for fi44 := 0; fi44 < 8; fi44++ {
                        _ = len(v2)
                        if false {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if false {
                                fmt.Println(v42, v5)
                                v45 := GenS2{F0: genFn1(((-1 >= (fi44 % ((fi43) | 1))) || (1 != 1)), v42), F1: (fi44 ^ fi44), F2: (v37 * (-(v37)))}
                                fmt.Println(fi43)
                                var v46 string = "hello"
                                v47 := GenS4{F0: (v37 << ((18446744073709551615) & 63)), F1: 42}
                                v3[fi44] = ((!(false || true)) || ((true || ((fi43 / ((862) | 1)) != int(v4))) || true))
                                var v48 uint64 = v42
                                _ = v45
                                _ = v46
                                _ = v47
                                _ = v48
                            } else {
                                fmt.Println(v41, v37)
                            }
                            switch (fi44 & (^(fi43))) {
                            case 87:
                                v49 := GenS3{F0: (v37 * (-(v37))), F1: 255}
                                _ = v49
                            default:
                                v37 = int8(v41)
                            }
                        }
                        switch 426 {
                        case -343:
                            var v50 bool = false
                            _ = v50
                        default:
                            var v51 int16 = v41
                            _ = v51
                        }
                        _ = fi44
                    }
                    switch "hello" {
                    case "foo bar":
                        _ = len(v2)
                    case "a":
                        var v52 int64 = 0
                        _ = v52
                    case "":
                        for fi53 := 0; fi53 < 12; fi53++ {
                            v54 := GenS5{F0: 841, F1: 42}
                            switch (fi53 - (^(fi53))) {
                            case -9223372036854775808:
                                var v55 uint64 = 1
                                _ = v55
                            case 1:
                                switch fi43 {
                                case 32767:
                                    if false {
                                        v37 = v37
                                    } else {
                                        v56 := GenS1{F0: 9223372036854775807}
                                        switch fi43 {
                                        case -325:
                                            var v57 int = fi43
                                            _ = v57
                                        default:
                                            v58 := GenS1{F0: int64(v41)}
                                            _ = v58
                                        }
                                        v5 -= (v5 ^ 2)
                                        var v59 uint8 = (v4 / (((v4 >> ((4294967295) & 63))) | 1))
                                        _ = v56
                                        _ = v59
                                    }
                                case 338:
                                    var v60 string = "hello"
                                    _ = v60
                                }
                            case 177:
                                _ = len(v3)
                            default:
                                var v61 int8 = (v37 << ((65535) & 63))
                                _ = v61
                            }
                            _ = fi53
                            _ = v54
                        }
                    }
                } else {
                    var v62 uint64 = 1
                    _ = v62
                }
            } else {
                fmt.Println(v37, v37)
                var v63 uint16 = (v5 / (((v5 % ((v5) | 1))) | 1))
                _ = v63
            }
            var v64 int = fi43
            v42 = v42
            _ = fi43
            _ = v64
        }
        _ = v42
    }
    v65 := v3[32767]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v37
    _ = v4
    _ = v41
    _ = v5
    _ = v65
}
