package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 bool; F1 uint8; F2 int16 }
type GenS7 struct { F0 string; F1 int32; F2 int64 }
func genFn0(p0 int8, p1 uint64, p2 uint64) uint {
    v6 := GenS1{F0: -949, F1: -2147483648}
    v7 := GenS1{F0: -9223372036854775808, F1: 564}
    v8 := GenS2{F0: uint(p2), F1: 18446744073709551615, F2: 228}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 uint64 = 974
    var v2 uint32 = 1
    v3 := make(map[int]bool)
    var v4 uint8 = 174
    var v5 bool = true
    if false {
        var v9 int16 = 1
        switch 32767 {
        case -721:
            var v10 uint32 = v2
            _ = v10
        default:
            if (!(-830 > 1)) {
                var v11 uint16 = 255
                _ = v11
            }
        }
        _ = v9
    }
    v12 := GenS2{F0: 65535, F1: v1, F2: 255}
    if ((1 >= 100) && (((1e10 <= 0.0) && false) || (!v5))) {
        v5 = ("\n\t" < ("foo bar" + "a"))
    } else {
        var v13 int = 42
        _ = v13
    }
    _ = len(v3)
    var v14 uint = 2
    for fi15 := 0; fi15 < 1; fi15++ {
        v16 := GenS3{F0: -9223372036854775808, F1: (v2 + 2), F2: -542}
        switch (fi15 / ((-175) | 1)) {
        case -735:
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 string = ("" + "\n\t")
                v19 := GenS2{F0: genFn0(-53, v1, v1), F1: v1, F2: 65535}
                _ = fi17
                _ = v18
                _ = v19
            }
        case 32767:
            if ((0.0 + (-2.25 - 0.1)) > (3.14159 + (0.1 + (-2.25 * (0.1 * 1.5))))) {
                var v20 uint8 = v4
                for fi21 := 0; fi21 < 19; fi21++ {
                    fmt.Println(v2, v1)
                    var v22 int16 = 128
                    _ = fi21
                    _ = v22
                }
                switch fi15 {
                case 971:
                    for fi23 := 0; fi23 < 18; fi23++ {
                        v20 -= v20
                        if (!(!("a" > ((("a" + ("" + "foo bar")) + "hello") + ("\n\t" + "foo bar"))))) {
                            if ((0 == 1) && v5) {
                                v24 := GenS1{F0: -144, F1: (fi15 * 9223372036854775807)}
                                _ = v24
                            } else {
                                if (2147483647 >= 636) {
                                    _ = len(v3)
                                    v25 := GenS3{F0: -9223372036854775808, F1: (v2 % ((uint32(v14)) | 1)), F2: 376}
                                    var v26 int16 = 42
                                    _ = v25
                                    _ = v26
                                } else {
                                    fmt.Println(fi15, fi15, v4)
                                }
                                var v27 bool = ("foo bar" <= (("a" + "a") + ("hello" + "a")))
                                v2 += (v2 * v2)
                                fmt.Println(fi23, v5)
                                _ = v27
                            }
                        } else {
                            v4 -= v4
                            var v28 int8 = 20
                            switch (^(fi23)) {
                            case 127:
                                if true {
                                    if v5 {
                                        if ((((v1 | v1) != (v1 & (v1 - uint64(v2)))) || v5) || ((!v5) || ((784 == 65535) || v5))) {
                                            if v5 {
                                                switch "hello" {
                                                case "hello":
                                                    switch 42 {
                                                    case -783:
                                                        switch ("foo bar" + "\n\t") {
                                                        case "foo bar":
                                                            var v29 string = ("foo bar" + "a")
                                                            _ = v29
                                                        case "hello":
                                                            var v30 int8 = v28
                                                            _ = v30
                                                        case "\n\t":
                                                            v31 := []int{fi23, 127}
                                                            _ = v31
                                                        }
                                                    case 4294967295:
                                                        fmt.Println(fi23, v4, v2)
                                                    }
                                                case "foo bar":
                                                    var v32 int32 = 2147483647
                                                    _ = v32
                                                default:
                                                    v33 := GenS4{F0: 479, F1: 2, F2: (v28 - 100)}
                                                    _ = v33
                                                }
                                                var v34 int64 = 42
                                                _ = v34
                                            } else {
                                                v5 = v5
                                            }
                                            v35 := GenS1{F0: 42, F1: (fi15 / ((fi15) | 1))}
                                            _ = v35
                                        }
                                        v36 := []int{(-(fi15)), fi23, 2147483647, (fi15 & (^(fi23)))}
                                        var v37 uint8 = v20
                                        _ = v36
                                        _ = v37
                                    }
                                    var v38 int = fi23
                                    _ = v38
                                }
                            case 938:
                                v39 := GenS4{F0: 437, F1: (v2 & (v2 % ((743) | 1))), F2: (v28 & (^(v28)))}
                                var v40 bool = (!v5)
                                _ = v39
                                _ = v40
                            case 585:
                                if v5 {
                                    v20 += 1
                                    switch ("hello" + ("hello" + "foo bar")) {
                                    case "\n\t":
                                        var v41 uint8 = uint8(v1)
                                        _ = v41
                                    case "a":
                                        v42 := []int{(fi23 % ((fi23) | 1)), -965, (fi15 % ((fi23) | 1)), fi23}
                                        _ = v42
                                    }
                                }
                            default:
                                if ((false || v5) || (v14 != 4294967295)) {
                                    if (774 == 0) {
                                        v43 := GenS1{F0: 1, F1: fi15}
                                        var v44 int16 = 32767
                                        v14 = (v14 * v14)
                                        _ = v43
                                        _ = v44
                                    } else {
                                        if ((true && (!v5)) && v5) {
                                            v14 -= v14
                                            var v45 int32 = -1
                                            _ = v45
                                        } else {
                                            var v46 int16 = 42
                                            v28 = v28
                                            v46 = (v46 / ((42) | 1))
                                            var v47 uint32 = (^(v2))
                                            _ = v46
                                            _ = v47
                                        }
                                        var v48 int16 = -1
                                        _ = v48
                                    }
                                } else {
                                    v49 := GenS3{F0: 666, F1: (^(v2)), F2: -2147483648}
                                    v5 = (v5 && v5)
                                    var v50 uint = 0
                                    _ = v49
                                    _ = v50
                                }
                            }
                            _ = v28
                        }
                        switch "a" {
                        case "":
                            v4 = uint8(v2)
                        default:
                            v1 -= (v1 * v1)
                        }
                        _ = fi23
                    }
                case 2147483647:
                    v51 := GenS5{F0: v2, F1: (false && (v5 || (uint16(v20) == 42)))}
                    _ = v51
                }
                _ = v20
            }
        case 934:
            v52 := GenS6{F0: v5, F1: (v4 | 128), F2: 1000}
            _ = v52
        default:
            _ = len(v3)
        }
        v5 = (42 < -128)
        _ = fi15
        _ = v16
    }
    v53 := GenS7{F0: ("a" + ""), F1: int32(v4), F2: -9223372036854775808}
    v54 := v3[32767]
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
}
