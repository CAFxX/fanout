package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 int16; F2 int64 }
type GenS5 struct { F0 uint64; F1 bool; F2 int64 }
type GenS6 struct { F0 uint32 }

func main() {
    var v1 string = (("\n\t" + "foo bar") + "\n\t")
    var v2 bool = false
    var v3 bool = v2
    var v4 uint = 1
    var v5 uint16 = 1
    var v6 uint32 = 2147483647
    fmt.Println(v5, v3)
    v1 = v1
    switch -1 {
    case 32768:
        var v7 bool = v3
        _ = v7
    case 4294967295:
        v4 = v4
    default:
        if false {
            for fi8 := 0; fi8 < 5; fi8++ {
                if (!v2) {
                    v9 := GenS1{F0: 0, F1: -2147483648, F2: 0}
                    for fi10 := 0; fi10 < 19; fi10++ {
                        fmt.Println(v4)
                        _ = fi10
                    }
                    _ = v9
                } else {
                    v3 = (9223372036854775807 < 0)
                }
                v11 := []int{}
                v12 := GenS1{F0: (v5 ^ v5), F1: -391, F2: -1}
                _ = fi8
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 uint64 = 1
                var v15 int16 = 32767
                _ = fi13
                _ = v14
                _ = v15
            }
        }
    }
    var v16 int = -20
    fmt.Println(v6, v2)
    if (32767 < 1) {
        for fi17 := 0; fi17 < 10; fi17++ {
            v6 += v6
            _ = fi17
        }
        var v18 uint32 = v6
        var v19 bool = v3
        _ = v18
        _ = v19
    }
    v20 := []int{(v16 + v16), v16, (v16 / (((v16 + 32768)) | 1))}
    v21 := GenS2{F0: v6}
    switch v16 {
    case 599:
        if (int64(v5) <= -9223372036854775808) {
            v16 = 743
            if ((v1 != (("foo bar" + v1) + v1)) || (v2 && v2)) {
                if ("\n\t" < "\n\t") {
                    for fi22 := 0; fi22 < 2; fi22++ {
                        if v2 {
                            for fi23 := 0; fi23 < 1; fi23++ {
                                v3 = true
                                var v24 uint64 = 818
                                _ = fi23
                                _ = v24
                            }
                            v20 = append(v20, v16)
                            switch v1 {
                            case "":
                                v25 := GenS2{F0: v6}
                                _ = v25
                            default:
                                v2 = ((-(v16)) < fi22)
                            }
                        }
                        for fi26 := 0; fi26 < 12; fi26++ {
                            _ = len(v20)
                            _ = fi26
                        }
                        _ = fi22
                    }
                    v3 = (2 <= 255)
                }
                var v27 float32 = 1e10
                v28 := GenS2{F0: 65535}
                _ = v27
                _ = v28
            } else {
                var v29 float32 = (0.0 + -2.25)
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                _ = v29
            }
        } else {
            v30 := GenS3{F0: 32767}
            if (v2 || (((42 <= 1000) && ((0.1 / 0.1) >= float32((0.1 - (((0.0 + (0.0 - 0.1)) - (3.14159 + (0.0 * (0.1 / 0.1)))) - (1e10 - 1.5)))))) && false)) {
                var v31 uint = 1
                fmt.Println(v3, v5)
                _ = v31
            } else {
                _ = len(v20)
                if v2 {
                    for fi32 := 0; fi32 < 2; fi32++ {
                        for fi33 := 0; fi33 < 17; fi33++ {
                            v6 -= v6
                            if (1 == 2147483647) {
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                            } else {
                                var v34 bool = v2
                                fmt.Println(v2, v4, fi32)
                                v35 := GenS1{F0: (v5 % ((v5) | 1)), F1: 679, F2: -840}
                                v36 := GenS1{F0: 65535, F1: 1, F2: -1}
                                var v37 int64 = 42
                                if true {
                                    var v38 uint64 = 479
                                    var v39 float64 = (0.1 / -2.25)
                                    _ = v38
                                    _ = v39
                                } else {
                                    var v40 int32 = -2147483648
                                    _ = v40
                                }
                                _ = v34
                                _ = v35
                                _ = v36
                                _ = v37
                            }
                            v41 := GenS4{F0: int8(v6), F1: 1000, F2: 0}
                            _ = fi33
                            _ = v41
                        }
                        var v42 bool = (v1 < v1)
                        _ = fi32
                        _ = v42
                    }
                } else {
                    fmt.Println(v4, v6, v2)
                    var v43 int16 = 795
                    for fi44 := 0; fi44 < 11; fi44++ {
                        v43 += int16(v6)
                        for fi45 := 0; fi45 < 9; fi45++ {
                            fmt.Println(v5)
                            _ = len(v20)
                            _ = fi45
                        }
                        _ = fi44
                    }
                    _ = v43
                }
            }
            if v2 {
                v46 := GenS5{F0: 2, F1: v3, F2: 9223372036854775807}
                _ = v46
            }
            _ = v30
        }
    case -249:
        v47 := GenS6{F0: v6}
        _ = v47
    case 42:
        if len(v20) > 0 {
            v20[0] = (v16 >> (((v4 & v4)) & 63))
        }
    }
    if true {
        var v48 uint = 1
        _ = v48
    } else {
        v4 = (v4 | v4)
        fmt.Println(v1, v1, v3)
        fmt.Println(v3, v3, v5)
    }
    var v49 int16 = 0
    var v50 uint8 = 249
    if ((v16 * (v16 >> ((36) & 63))) < (v16 / (((v16 | v16)) | 1))) {
        var v51 int64 = 1
        _ = v51
    }
    var v52 int32 = 103
    if ((v4 | 18446744073709551615) <= (v4 % (((v4 + (v4 & v4))) | 1))) {
        if len(v20) > 0 {
            v20[0] = 9223372036854775807
        }
    } else {
        v1 = v1
        v20 = append(v20, (v16 << ((872) & 63)))
    }
    v49 = (v49 - int16(v16))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v52
    _ = v6
}
