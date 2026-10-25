package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint32; F2 uint64 }
type GenS4 struct { F0 bool; F1 int8; F2 uint64 }
type GenS5 struct { F0 int8; F1 bool; F2 int64 }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 int16; F1 uint }
type GenS8 struct { F0 int; F1 int; F2 int16 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int; F1 int }
type GenS12 struct { F0 int32; F1 uint }
type GenS13 struct { F0 int8; F1 bool; F2 int64 }
type GenS14 struct { F0 uint64; F1 int64; F2 int8 }
func genFn0(p0 uint64, p1 uint64, p2 int32) string {
    p0 = (^(p1))
    v7 := GenS1{F0: (p2 << ((2) & 63))}
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := []int{fi8, fi9, fi9}
            if true {
                if true {
                    var v11 int64 = 9223372036854775807
                    switch (fi9 * (fi8 >> ((255) & 63))) {
                    case 4294967295:
                        if false {
                            var v12 uint8 = 108
                            _ = v12
                        } else {
                            switch "hello" {
                            case "a":
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(fi8, p2)
                                case "hello":
                                    p1 = (^(p0))
                                case "":
                                    var v13 uint32 = 2147483647
                                    _ = v13
                                }
                            }
                        }
                    }
                    _ = v11
                }
                var v14 int32 = (p2 >> ((65535) & 63))
                p1 -= (^(p0))
                _ = v14
            } else {
                v15 := GenS2{F0: (fi8 | fi9)}
                _ = v15
            }
            _ = fi9
            _ = v10
        }
        p1 -= 722
        _ = fi8
    }
    p2 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (("foo bar" + ((("foo bar" + "") + "\n\t") + ("" + "\n\t"))) + ("hello" + "foo bar"))
}
func genFn1() string {
    if ((128 <= 185) && true) {
        var v16 uint32 = 1
        _ = v16
    } else {
        var v17 uint8 = 0
        v17 -= 0
        v18 := GenS1{F0: 0}
        _ = v17
        _ = v18
    }
    var v19 bool = ((true && (!(((1.5 / 1e10) + 0.0) > -2.25))) && (true && true))
    var v20 float64 = 1.5
    _ = v19
    _ = v20
    return (("a" + "") + "a")
}

func main() {
    var v1 int = 842
    var v2 uint64 = 42
    var v3 uint8 = 127
    var v4 int32 = 42
    var v5 uint16 = 65535
    v6 := make(map[uint32]int)
    v21 := []int{v1}
    delete(v6, uint32(v2))
    if (-960 < -1) {
        fmt.Println(v4, v4, v5)
        var v22 uint16 = v5
        if (true && true) {
            var v23 uint16 = (v5 + (v5 % ((v22) | 1)))
            _ = v23
        } else {
            if (!(1 < (v5 ^ (v22 ^ (v22 / (((v5 >> ((2) & 63))) | 1)))))) {
                v24 := GenS3{F0: v1, F1: 734, F2: v2}
                v25 := GenS3{F0: (v1 | (v1 | (v1 >> ((2) & 63)))), F1: 725, F2: 749}
                v26 := GenS4{F0: true, F1: -66, F2: (v2 % (((v2 << ((uint(v22)) & 63))) | 1))}
                _ = v24
                _ = v25
                _ = v26
            }
            var v27 float64 = 3.14159
            v28 := GenS5{F0: 42, F1: (!(!(!(v27 <= v27)))), F2: -576}
            _ = v27
            _ = v28
        }
        _ = v22
    } else {
        if (1e10 <= 0.1) {
            v29 := GenS1{F0: -2147483648}
            for fi30 := 0; fi30 < 2; fi30++ {
                for fi31 := 0; fi31 < 13; fi31++ {
                    var v32 int32 = int32(v2)
                    v33 := GenS6{F0: v32, F1: (v5 >> ((0) & 63))}
                    v1 = fi30
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                v4 = (v4 >> ((805) & 63))
                var v34 int = -2147483648
                _ = fi30
                _ = v34
            }
            _ = v29
        } else {
            v21 = append(v21, v1)
            if (int8(v1) != 1) {
                var v35 string = (("" + "foo bar") + "\n\t")
                _ = len(v21)
                _ = v35
            }
        }
        fmt.Println(v3)
        var v36 string = "\n\t"
        _ = v36
    }
    v2 = v2
    v37 := GenS1{F0: v4}
    v38 := GenS7{F0: 1000, F1: 894}
    v39 := GenS1{F0: v4}
    if ((-2.25 / (1e10 * ((1.5 * -2.25) * -2.25))) == ((float32(3.14159) / (3.14159 + 1.5)) * (1e10 * 1.5))) {
        v2 = v2
        if false {
            if (!(!(v1 >= 9223372036854775807))) {
                for fi40 := 0; fi40 < 17; fi40++ {
                    var v41 uint32 = 2
                    var v42 bool = (!false)
                    for fi43 := 0; fi43 < 4; fi43++ {
                        v44 := GenS8{F0: fi40, F1: fi40, F2: 0}
                        if ((3.14159 + 0.1) == 1e10) {
                            switch v1 {
                            case 0:
                                var v45 uint8 = (v3 + 42)
                                var v46 int = (-(fi40))
                                _ = v45
                                _ = v46
                            default:
                                v2 += 507
                            }
                            if (42 <= 0) {
                                v47 := GenS4{F0: false, F1: 0, F2: v2}
                                fmt.Println(v2, fi43, v1)
                                _ = v47
                            } else {
                                v3 = (^(v3))
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                            }
                        } else {
                            var v48 uint = 607
                            _ = v48
                        }
                        _ = fi43
                        _ = v44
                    }
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            } else {
                v3 = (v3 | (v3 ^ v3))
                delete(v6, 352)
                if (0 == 1) {
                    for fi49 := 0; fi49 < 9; fi49++ {
                        v50 := GenS9{F0: 65535}
                        _ = fi49
                        _ = v50
                    }
                } else {
                    if ((^(v2)) <= v2) {
                        v2 = 42
                        v2 += (v2 * v2)
                        if (1 == 2147483648) {
                            for fi51 := 0; fi51 < 10; fi51++ {
                                var v52 uint8 = v3
                                v4 = -32
                                _ = fi51
                                _ = v52
                            }
                            var v53 uint8 = (^(v3))
                            _ = v53
                        } else {
                            fmt.Println(v1, v1)
                            var v54 int64 = 870
                            _ = v54
                        }
                    } else {
                        v55 := v6[0]
                        _ = v55
                    }
                }
            }
            var v56 int8 = 12
            _ = v56
        } else {
            v57 := GenS5{F0: 40, F1: (2147483648 != 4294967295), F2: -1}
            _ = v57
        }
    } else {
        var v58 int = 4294967295
        _ = v58
    }
    var v59 uint16 = 1
    var v60 int8 = 79
    if (v5 <= v59) {
        if (!(!true)) {
            switch "a" {
            case "\n\t":
                var v61 int64 = -9223372036854775808
                _ = v61
            case "hello":
                v60 = (v60 * -110)
            case "":
                v62 := GenS8{F0: (v1 % ((v1) | 1)), F1: (^(v1)), F2: 42}
                _ = v62
            default:
                switch "hello" {
                case "":
                    if (-1 > 768) {
                        v59 += (v59 / ((v59) | 1))
                        var v63 uint = 18446744073709551615
                        switch v1 {
                        case -2147483648:
                            v64 := GenS10{F0: v59}
                            _ = v64
                        case -9223372036854775808:
                            if true {
                                var v65 int16 = -1
                                _ = v65
                            } else {
                                var v66 int8 = (-(v60))
                                fmt.Println(v3)
                                v67 := GenS11{F0: (v1 | 42), F1: -1}
                                _ = v66
                                _ = v67
                            }
                        case 128:
                            v68 := GenS12{F0: (v4 + v4), F1: v63}
                            _ = v68
                        }
                        _ = v63
                    }
                case "hello":
                    if (int16(v4) >= 42) {
                        if (((v3 << ((251) & 63)) <= (^(v3))) && ((v3 * 42) < 2)) {
                            var v69 uint = 255
                            _ = v69
                        } else {
                            for fi70 := 0; fi70 < 18; fi70++ {
                                v71 := GenS13{F0: v60, F1: (v3 <= (v3 & (v3 + (v3 / ((v3) | 1))))), F2: 1000}
                                _ = fi70
                                _ = v71
                            }
                            if false {
                                if (881 > -9223372036854775808) {
                                    fmt.Println(v5, v5)
                                } else {
                                    switch -776 {
                                    case 128:
                                        if ((v60 / ((v60) | 1)) != v60) {
                                            v60 = (v60 ^ -1)
                                            _ = len(v6)
                                            fmt.Println(v4)
                                        }
                                    case -2147483648:
                                        var v72 float64 = 1e10
                                        _ = v72
                                    case 32768:
                                        v59 -= 2
                                    }
                                    switch (v1 ^ v1) {
                                    case 127:
                                        v73 := GenS14{F0: v2, F1: 1, F2: v60}
                                        _ = v73
                                    }
                                    _ = len(v21)
                                }
                                v59 = 385
                                var v74 int32 = v4
                                _ = v74
                            }
                            if false {
                                _ = len(v6)
                            } else {
                                v4 = 145
                                var v75 float32 = 1.5
                                _ = v75
                            }
                        }
                        v2 -= (v2 / ((v2) | 1))
                    }
                }
            }
        } else {
            fmt.Println(v60)
        }
    } else {
        var v76 int64 = -69
        v1 = v1
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
}
