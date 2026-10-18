package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 uint16; F2 uint }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 int32; F1 int32 }
type GenS9 struct { F0 int16; F1 int }
func genFn0() string {
    var v4 int16 = -786
    var v5 int16 = (v4 >> ((1) & 63))
    v6 := []int{2147483647, -555}
    _ = v4
    _ = v5
    _ = v6
    return ("foo bar" + ("a" + "\n\t"))
}
func genFn1(p0 int32) uint32 {
    v7 := GenS1{F0: "a"}
    for fi8 := 0; fi8 < 19; fi8++ {
        switch fi8 {
        case -825:
            p0 += p0
        default:
            v9 := GenS2{F0: "a"}
            _ = v9
        }
        _ = fi8
    }
    var v10 int8 = -128
    _ = p0
    _ = v10
    _ = v7
    return 0
}

func main() {
    var v1 uint64 = 1
    var v2 float64 = 1e10
    var v3 uint16 = 42
    v11 := []int{-2147483648, 1, -116}
    var v12 int64 = 9223372036854775807
    var v13 uint64 = 4294967295
    var v14 uint64 = v13
    var v15 int8 = 42
    for fi16 := 0; fi16 < 15; fi16++ {
        switch fi16 {
        case -324:
            for fi17 := 0; fi17 < 1; fi17++ {
                switch ("\n\t" + ("foo bar" + "hello")) {
                case "":
                    switch (("hello" + "hello") + ("\n\t" + "\n\t")) {
                    case "\n\t":
                        v18 := GenS3{F0: -513, F1: v3, F2: 124}
                        _ = v18
                    default:
                        fmt.Println(v13, fi17, v15)
                    }
                }
                if false {
                    switch fi16 {
                    case 0:
                        v1 = (^(v14))
                    default:
                        var v19 int8 = v15
                        _ = v19
                    }
                }
                if ((v1 % ((474) | 1)) != v14) {
                    switch fi16 {
                    case 9223372036854775807:
                        v15 += 1
                    case -434:
                        v12 = (v12 >> ((217) & 63))
                    case 2147483647:
                        v20 := GenS4{F0: 42, F1: 2, F2: uint8(v3)}
                        _ = v20
                    }
                    v11 = append(v11, (-(fi16)))
                    var v21 int32 = 991
                    _ = v21
                }
                _ = fi17
            }
        case 32767:
            v14 = 376
        case 1:
            for fi22 := 0; fi22 < 5; fi22++ {
                v14 = 0
                _ = fi22
            }
        }
        if ((^(v15)) > v15) {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            v13 = (v13 * v13)
            for fi23 := 0; fi23 < 2; fi23++ {
                if (4294967295 >= 42) {
                    var v24 int = (fi23 & (fi16 | 868))
                    _ = v24
                }
                var v25 bool = true
                switch fi16 {
                case 291:
                    v13 += (v13 >> ((4294967295) & 63))
                default:
                    if len(v11) > 0 {
                        v11[0] = fi16
                    }
                }
                if false {
                    var v26 int64 = v12
                    if true {
                        v27 := GenS5{F0: v14, F1: -1}
                        switch "foo bar" {
                        case "hello":
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                        }
                        if v25 {
                            var v28 uint = 42
                            v11 = append(v11, fi23)
                            _ = v28
                        }
                        _ = v27
                    }
                    _ = v26
                } else {
                    v29 := GenS6{F0: v25}
                    fmt.Println(v2)
                    _ = v29
                }
                var v30 uint = 1
                _ = fi23
                _ = v25
                _ = v30
            }
        } else {
            v12 = int64(v3)
            v14 += (v1 | (^(v14)))
        }
        _ = fi16
    }
    if (0 == -2147483648) {
        for fi31 := 0; fi31 < 14; fi31++ {
            v32 := GenS5{F0: 42, F1: (v15 + (v15 + (v15 % (((v15 >> ((42) & 63))) | 1))))}
            var v33 uint16 = (v3 ^ (v3 - 1))
            fmt.Println(v15, v2, v15)
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = len(v11)
    }
    if (737 > (v3 - v3)) {
        v34 := GenS2{F0: ("\n\t" + "a")}
        v35 := GenS7{F0: -9223372036854775808, F1: ""}
        _ = v34
        _ = v35
    }
    var v36 uint16 = v3
    for fi37 := 0; fi37 < 5; fi37++ {
        if len(v11) > 0 {
            v11[0] = (-(fi37))
        }
        _ = fi37
    }
    var v38 uint16 = (v36 & v36)
    for fi39 := 0; fi39 < 5; fi39++ {
        if (-1 != -439) {
            switch fi39 {
            case -9223372036854775808:
                for fi40 := 0; fi40 < 8; fi40++ {
                    v11 = append(v11, fi40)
                    _ = fi40
                }
            case 2147483647:
                if len(v11) > 0 {
                    v11[0] = -728
                }
            case 32:
                v2 = 0.1
            }
            for fi41 := 0; fi41 < 9; fi41++ {
                v11 = append(v11, (fi41 * (fi41 << ((1) & 63))))
                v3 -= (v3 ^ v3)
                _ = fi41
            }
            var v42 uint32 = 4294967295
            _ = v42
        } else {
            if false {
                if false {
                    for fi43 := 0; fi43 < 10; fi43++ {
                        fmt.Println(v3, fi43, v15)
                        if (true && false) {
                            v44 := GenS7{F0: -1, F1: "a"}
                            var v45 float64 = (v2 - v2)
                            _ = v44
                            _ = v45
                        } else {
                            switch fi43 {
                            case 9223372036854775807:
                                if len(v11) > 0 {
                                    v11[0] = (fi39 - fi39)
                                }
                            case -519:
                                var v46 float64 = v2
                                _ = v46
                            }
                            if true {
                                var v47 int8 = (v15 >> ((4294967295) & 63))
                                v13 = (v13 * v14)
                                _ = v47
                            } else {
                                var v48 uint64 = (v1 | 4294967295)
                                fmt.Println(v2)
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                _ = v48
                            }
                        }
                        if false {
                            v15 += v15
                            if false {
                                v11 = append(v11, fi39)
                                var v49 uint64 = 42
                                _ = len(v11)
                                _ = v49
                            } else {
                                v50 := GenS3{F0: v12, F1: 32767, F2: 65535}
                                if (!true) {
                                    v36 = 620
                                    v14 = (v13 ^ v1)
                                    v51 := GenS8{F0: 42, F1: 1}
                                    _ = v51
                                } else {
                                    fmt.Println(v14)
                                    v52 := GenS9{F0: 531, F1: -9223372036854775808}
                                    _ = v52
                                }
                                _ = v50
                            }
                        } else {
                            if true {
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                var v53 int = (-(fi39))
                                _ = v53
                            } else {
                                v38 = v38
                                switch 2147483647 {
                                case 522:
                                    var v54 bool = ((v38 / ((v36) | 1)) > 32767)
                                    _ = v54
                                }
                                v38 += (v3 | v38)
                            }
                        }
                        _ = len(v11)
                        fmt.Println(v3, v36, v15)
                        var v55 float64 = 1e10
                        if len(v11) > 0 {
                            v11[0] = int(v15)
                        }
                        _ = fi43
                        _ = v55
                    }
                    var v56 int64 = v12
                    var v57 int64 = v56
                    _ = v56
                    _ = v57
                }
            }
            var v58 uint64 = 2
            for fi59 := 0; fi59 < 20; fi59++ {
                v13 = v58
                switch fi39 {
                case -193:
                    var v60 int32 = 42
                    _ = v60
                }
                _ = fi59
            }
            _ = v58
        }
        v1 = (v14 ^ (v14 | v1))
        if false {
            fmt.Println(v13)
            v14 += 401
            var v61 int = fi39
            _ = v61
        }
        _ = fi39
    }
    _ = len(v11)
    for fi62 := 0; fi62 < 1; fi62++ {
        switch (fi62 / (((fi62 | (fi62 >> ((uint(v13)) & 63)))) | 1)) {
        case 4294967295:
            var v63 int32 = -994
            _ = v63
        case 0:
            switch (^(fi62)) {
            case -1:
                for fi64 := 0; fi64 < 7; fi64++ {
                    fmt.Println(fi62, v38, fi64)
                    switch fi64 {
                    case -267:
                        var v65 uint16 = v36
                        _ = v65
                    case -420:
                        fmt.Println(v14, v1, v12)
                    case 128:
                        v3 = v36
                    }
                    _ = fi64
                }
            case 4294967295:
                var v66 uint8 = 42
                _ = v66
            case 127:
                var v67 uint64 = 42
                _ = v67
            default:
                v3 -= (^(v38))
            }
        case 2147483647:
            _ = len(v11)
        }
        _ = fi62
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v38
}
