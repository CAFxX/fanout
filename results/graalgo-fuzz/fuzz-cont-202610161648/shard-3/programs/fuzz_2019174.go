package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 uint; F2 bool }
func genFn0(p0 int32, p1 bool, p2 string) int16 {
    p2 = p2
    var v5 int = 9223372036854775807
    v6 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint8 = 174
    v2 := []int{-865}
    var v3 uint = 4294967295
    var v4 string = ("hello" + (("a" + "\n\t") + "foo bar"))
    v4 = v4
    v7 := GenS1{F0: 2147483647}
    v8 := GenS1{F0: 2}
    v4 = "\n\t"
    v9 := GenS1{F0: 1}
    var v10 int16 = genFn0(0, false, (v4 + v4))
    var v11 string = "\n\t"
    var v12 uint64 = 739
    v11 = ("\n\t" + v4)
    for fi13 := 0; fi13 < 18; fi13++ {
        v11 = v11
        var v14 int = (fi13 | (^(fi13)))
        if len(v2) > 0 {
            v2[0] = fi13
        }
        _ = fi13
        _ = v14
    }
    if (595 != 101) {
        v3 = (v3 - v3)
        var v15 int16 = (v10 - (v10 % (((v10 ^ v10)) | 1)))
        _ = v15
    }
    var v16 int8 = 127
    var v17 float32 = 0.1
    if len(v2) > 0 {
        v2[0] = int(v12)
    }
    fmt.Println(v4, v17)
    switch "\n\t" {
    case "foo bar":
        switch 32767 {
        case 254:
            if ((false && (!false)) || false) {
                if len(v2) > 0 {
                    v2[0] = -703
                }
                v18 := GenS2{F0: 1}
                var v19 bool = false
                _ = v18
                _ = v19
            }
        case 1:
            for fi20 := 0; fi20 < 10; fi20++ {
                v1 = v1
                if true {
                    var v21 uint16 = 402
                    v22 := GenS3{F0: 1}
                    fmt.Println(v12)
                    _ = v21
                    _ = v22
                } else {
                    var v23 float64 = 0.1
                    for fi24 := 0; fi24 < 4; fi24++ {
                        v2 = append(v2, fi20)
                        _ = fi24
                    }
                    _ = v23
                }
                var v25 float64 = float64(v17)
                _ = fi20
                _ = v25
            }
        default:
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v11)
            case "\n\t":
                var v26 uint = 267
                _ = v26
            case "a":
                if false {
                    if (1 > 1) {
                        for fi27 := 0; fi27 < 6; fi27++ {
                            if len(v2) > 0 {
                                v2[0] = 32768
                            }
                            switch int(v16) {
                            case 0:
                                _ = len(v2)
                            case 745:
                                fmt.Println(v11)
                            }
                            _ = fi27
                        }
                    }
                    v11 = "\n\t"
                    if ((^(v1)) >= v1) {
                        for fi28 := 0; fi28 < 7; fi28++ {
                            for fi29 := 0; fi29 < 13; fi29++ {
                                v10 = v10
                                var v30 float32 = 1e10
                                v30 = ((float32(0.1) + 1.5) - (3.14159 / v17))
                                _ = fi29
                                _ = v30
                            }
                            v10 -= 32767
                            _ = fi28
                        }
                    } else {
                        if (-1 < 1) {
                            var v31 uint16 = 1
                            switch 2147483647 {
                            case -9223372036854775808:
                                var v32 int = -609
                                _ = v32
                            case 42:
                                for fi33 := 0; fi33 < 1; fi33++ {
                                    if (false || (!(((v31 - uint16(v3)) == (v31 + (^(v31)))) && false))) {
                                        v10 = v10
                                        if true {
                                            v34 := GenS3{F0: 1}
                                            _ = v34
                                        }
                                        v1 -= v1
                                    }
                                    fmt.Println(v17, v11)
                                    var v35 uint = (v3 * v3)
                                    _ = fi33
                                    _ = v35
                                }
                            }
                            if false {
                                fmt.Println(v4, v12, v3)
                                if true {
                                    v10 -= (-(v10))
                                    fmt.Println(v12, v31, v11)
                                    if false {
                                        for fi36 := 0; fi36 < 6; fi36++ {
                                            switch fi36 {
                                            case 128:
                                                var v37 int16 = (v10 % (((^(v10))) | 1))
                                                _ = v37
                                            case 42:
                                                v12 = 594
                                            case -232:
                                                fmt.Println(v17, v12)
                                            default:
                                                v38 := GenS4{F0: genFn0(-2147483648, false, "a"), F1: 620, F2: (!(true && false))}
                                                _ = v38
                                            }
                                            v2 = append(v2, fi36)
                                            _ = len(v2)
                                            _ = fi36
                                        }
                                        v17 = v17
                                        var v39 uint16 = v31
                                        _ = v39
                                    }
                                } else {
                                    v40 := GenS2{F0: 0}
                                    _ = v40
                                }
                            }
                            _ = v31
                        }
                        _ = len(v2)
                        v17 = v17
                    }
                }
            }
        }
    case "a":
        switch 32767 {
        case 9223372036854775807:
            if false {
                var v41 float32 = v17
                for fi42 := 0; fi42 < 6; fi42++ {
                    v3 += (v3 | v3)
                    fmt.Println(v11, v11)
                    if ((int(v3) != (fi42 << ((v3) & 63))) && true) {
                        if len(v2) > 0 {
                            v2[0] = (fi42 / ((fi42) | 1))
                        }
                        var v43 uint16 = 32
                        _ = v43
                    }
                    _ = fi42
                }
                _ = v41
            } else {
                if ((v12 + (v12 << ((v3) & 63))) <= 4294967295) {
                    fmt.Println(v3, v1, v4)
                } else {
                    if (0 == uint64(v3)) {
                        fmt.Println(v17, v17)
                    } else {
                        var v44 int32 = -1
                        for fi45 := 0; fi45 < 15; fi45++ {
                            fmt.Println(v44, v44)
                            _ = fi45
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v44
                    }
                }
                if (!(!(!(-1 == -9223372036854775808)))) {
                    for fi46 := 0; fi46 < 17; fi46++ {
                        v1 = 42
                        _ = fi46
                    }
                } else {
                    for fi47 := 0; fi47 < 19; fi47++ {
                        v12 = v12
                        for fi48 := 0; fi48 < 12; fi48++ {
                            var v49 int32 = 1
                            _ = fi48
                            _ = v49
                        }
                        _ = fi47
                    }
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
    case "\n\t":
        if false {
            for fi50 := 0; fi50 < 20; fi50++ {
                v51 := GenS4{F0: 42, F1: 2, F2: (0.0 >= (1.5 * 0.1))}
                _ = fi50
                _ = v51
            }
            v3 = 376
            v17 += (0.1 * 3.14159)
        } else {
            if true {
                var v52 string = v4
                _ = v52
            } else {
                _ = len(v2)
                v53 := GenS4{F0: v10, F1: v3, F2: true}
                var v54 float32 = ((v17 + v17) * 0.1)
                _ = v53
                _ = v54
            }
            v12 = 695
            fmt.Println(v17, v1)
        }
    default:
        for fi55 := 0; fi55 < 2; fi55++ {
            if (!false) {
                v56 := GenS3{F0: 625}
                var v57 uint8 = (v1 & 0)
                _ = v56
                _ = v57
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if true {
                var v58 string = "a"
                if true {
                    v4 = ((v4 + v4) + "hello")
                }
                fmt.Println(fi55, v12)
                _ = v58
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if (!false) {
                    switch fi55 {
                    case 743:
                        fmt.Println(v1, v17)
                    case -246:
                        for fi59 := 0; fi59 < 14; fi59++ {
                            v60 := GenS1{F0: 42}
                            _ = fi59
                            _ = v60
                        }
                    }
                }
                var v61 uint32 = 494
                _ = v61
            }
            _ = fi55
        }
    }
    fmt.Println(v3, v16, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
