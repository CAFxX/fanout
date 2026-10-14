package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 string; F1 uint8; F2 int32 }
type GenS7 struct { F0 int16 }
func genFn0() uint8 {
    var v6 int8 = 1
    v7 := GenS1{F0: 128, F1: 1, F2: 1}
    _ = v6
    _ = v7
    return 0
}
func genFn1() int {
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS2{F0: ("" + (("\n\t" + "a") + "hello"))}
        _ = fi8
        _ = v9
    }
    if (!(!false)) {
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            switch "" {
            case "hello":
                if false {
                    v10 := []int{}
                    var v11 int16 = 1000
                    switch "a" {
                    case "":
                        if false {
                            if ((false || (true || (2 == 64))) || true) {
                                for fi12 := 0; fi12 < 13; fi12++ {
                                    var v13 int16 = (v11 >> ((65535) & 63))
                                    fmt.Println(v13, v11)
                                    v11 -= v13
                                    _ = fi12
                                    _ = v13
                                }
                            }
                        } else {
                            var v14 int64 = 640
                            v15 := GenS3{F0: (v11 / (((v11 ^ (v11 % (((v11 & -97)) | 1)))) | 1)), F1: 65535}
                            v11 = 42
                            _ = v14
                            _ = v15
                        }
                    case "hello":
                        fmt.Println(v11)
                    default:
                        v11 += v11
                    }
                    _ = v10
                    _ = v11
                } else {
                    var v16 int16 = -32768
                    _ = v16
                }
            default:
                var v17 int32 = 497
                _ = v17
            }
        case "a":
            var v18 uint64 = 568
            _ = v18
        case "hello":
            if ((false || true) && false) {
                switch 842 {
                case -2147483648:
                    var v19 int8 = 27
                    _ = v19
                case 4294967295:
                    if (42 == 65535) {
                        for fi20 := 0; fi20 < 1; fi20++ {
                            if (1 <= 0) {
                                if (0 <= -128) {
                                    if ((65535 <= 1) || false) {
                                        var v21 uint16 = 65535
                                        _ = v21
                                    } else {
                                        v22 := GenS4{F0: 9223372036854775807}
                                        _ = v22
                                    }
                                }
                            }
                            v23 := GenS1{F0: 127, F1: fi20, F2: 121}
                            _ = fi20
                            _ = v23
                        }
                        v24 := GenS1{F0: 255, F1: 63, F2: 489}
                        _ = v24
                    }
                default:
                    if true {
                        var v25 string = "\n\t"
                        fmt.Println(v25)
                        _ = v25
                    } else {
                        var v26 uint8 = genFn0()
                        v27 := GenS5{F0: -1, F1: 42}
                        if (true && true) {
                            fmt.Println(v26)
                        } else {
                            v28 := GenS6{F0: "foo bar", F1: (v26 / ((genFn0()) | 1)), F2: 521}
                            _ = v28
                        }
                        _ = v26
                        _ = v27
                    }
                }
                switch 128 {
                case 32767:
                    var v29 int64 = 623
                    _ = v29
                case 9223372036854775807:
                    var v30 uint8 = 125
                    _ = v30
                case -2147483648:
                    var v31 uint8 = 0
                    _ = v31
                }
                v32 := GenS7{F0: 32767}
                _ = v32
            } else {
                for fi33 := 0; fi33 < 14; fi33++ {
                    var v34 uint = 535
                    v35 := []int{fi33, -2147483648}
                    fmt.Println(fi33, v34)
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                for fi36 := 0; fi36 < 8; fi36++ {
                    var v37 float32 = 1.5
                    fmt.Println(fi36)
                    _ = fi36
                    _ = v37
                }
            }
        }
    }
    var v38 float64 = float64(1.5)
    _ = v38
    return -950
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int32 = -2147483648
    var v3 int64 = v1
    v4 := make(map[string]bool)
    var v5 string = "foo bar"
    fmt.Println(v3, v3, v5)
    var v39 int64 = -1
    var v40 int8 = 100
    for fi41 := 0; fi41 < 17; fi41++ {
        v42 := v4["hello"]
        _ = fi41
        _ = v42
    }
    v2 = (v2 >> ((0) & 63))
    for fi43 := 0; fi43 < 17; fi43++ {
        v40 = v40
        if false {
            for fi44 := 0; fi44 < 10; fi44++ {
                if ((v2 != v2) && (!((v40 % (((v40 | (v40 << ((1) & 63)))) | 1)) >= (v40 | (v40 ^ (v40 * 0)))))) {
                    switch genFn1() {
                    case 32767:
                        var v45 uint8 = 153
                        v46 := []int{(-(fi44)), (fi44 ^ (fi43 ^ genFn1())), fi44, (fi43 | 995)}
                        _ = v45
                        _ = v46
                    }
                    var v47 uint8 = 185
                    v3 = v39
                    _ = v47
                } else {
                    v39 -= (^(v39))
                    var v48 uint16 = 2
                    var v49 uint32 = 42
                    var v50 int32 = (v2 << ((606) & 63))
                    _ = v48
                    _ = v49
                    _ = v50
                }
                switch v5 {
                case "\n\t":
                    if (!true) {
                        if false {
                            v3 += v3
                            switch "" {
                            case "\n\t":
                                if (65535 != 185) {
                                    var v51 uint32 = 42
                                    var v52 int32 = (^(v2))
                                    if (18446744073709551615 > 255) {
                                        v53 := GenS2{F0: (v5 + ("foo bar" + (v5 + ((v5 + v5) + v5))))}
                                        _ = v53
                                    }
                                    if true {
                                        if false {
                                            _ = len(v4)
                                            _ = len(v4)
                                        }
                                        v54 := GenS5{F0: 0, F1: 84}
                                        _ = v54
                                    } else {
                                        fmt.Println(fi44, v52)
                                    }
                                    _ = v51
                                    _ = v52
                                }
                            }
                        }
                        var v55 uint16 = 2
                        _ = v55
                    } else {
                        fmt.Println(v2)
                        v56 := GenS6{F0: v5, F1: 128, F2: (v2 ^ -2147483648)}
                        switch "" {
                        case "a":
                            if false {
                                v39 = -1
                                v39 = (v3 ^ 42)
                            }
                        case "hello":
                            v40 = (v40 - (v40 / (((v40 << ((448) & 63))) | 1)))
                        case "":
                            v40 = v40
                        default:
                            fmt.Println(v39, v40)
                        }
                        _ = v56
                    }
                default:
                    var v57 uint8 = 125
                    var v58 uint32 = 284
                    _ = v57
                    _ = v58
                }
                _ = fi44
            }
            fmt.Println(v3, v2, fi43)
        } else {
            v40 -= (-(v40))
            v59 := GenS7{F0: 272}
            _ = v59
        }
        _ = fi43
    }
    v40 += v40
    if false {
        if (!(false && (65535 == 798))) {
            v4[v5] = (105 > 150)
            var v60 int8 = (v40 % (((v40 % ((-128) | 1))) | 1))
            v39 = 660
            _ = v60
        }
        var v61 int16 = 1000
        _ = v61
    } else {
        v4[v5] = false
    }
    v62 := GenS3{F0: 1, F1: 974}
    switch v5 {
    case "\n\t":
        v63 := GenS6{F0: v5, F1: 0, F2: (-(v2))}
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v62
}
