package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 int64; F2 int }
type GenS5 struct { F0 int }
func genFn0(p0 int16) uint16 {
    v5 := GenS1{F0: -117, F1: 194}
    if (true && false) {
        fmt.Println(p0)
        p0 -= 487
        fmt.Println(p0)
    } else {
        v6 := GenS2{F0: true}
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 uint16 = 255
                v10 := []int{fi7, (fi8 << ((42) & 63)), fi7}
                _ = fi8
                _ = v10
                _ = v9
            }
            var v11 float64 = (((float64(-2.25) * -2.25) * 0.1) + 1e10)
            fmt.Println(p0, fi7)
            _ = fi7
            _ = v11
        }
        var v12 string = (("a" + "") + "a")
        _ = v12
        _ = v6
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        if ((0 > -230) || false) {
            for fi14 := 0; fi14 < 6; fi14++ {
                v15 := []int{}
                var v16 float32 = 0.1
                _ = fi14
                _ = v15
                _ = v16
            }
            v17 := GenS3{F0: ("a" + ("a" + ("\n\t" + ("foo bar" + ("hello" + "")))))}
            _ = v17
        }
        switch "\n\t" {
        case "":
            var v18 float64 = 0.1
            _ = v18
        case "a":
            switch 9223372036854775807 {
            case 1:
                if false {
                    p0 -= 0
                    var v19 int32 = 42
                    v19 = (v19 | v19)
                    _ = v19
                }
            case -1:
                switch fi13 {
                case 339:
                    if false {
                        v20 := GenS1{F0: -1, F1: 19}
                        p0 += p0
                        _ = v20
                    } else {
                        v21 := []int{1}
                        if false {
                            if (2147483647 >= 0) {
                                switch fi13 {
                                case 190:
                                    var v22 int64 = -9223372036854775808
                                    _ = v22
                                default:
                                    var v23 int16 = (p0 / (((p0 * (p0 | p0))) | 1))
                                    _ = v23
                                }
                                if false {
                                    p0 -= p0
                                } else {
                                    fmt.Println(p0, p0)
                                    p0 -= (p0 - -140)
                                }
                            }
                        }
                        _ = v21
                    }
                case 128:
                    v24 := GenS3{F0: "hello"}
                    _ = v24
                default:
                    v25 := []int{fi13, -1, (-(fi13))}
                    _ = v25
                }
            default:
                fmt.Println(p0, fi13)
            }
        default:
            fmt.Println(fi13)
        }
        _ = fi13
    }
    _ = p0
    _ = v5
    return 0
}
func genFn1(p0 uint32, p1 uint8) int64 {
    fmt.Println(p1, p1)
    var v26 int64 = -90
    v27 := []int{32767}
    var v28 uint64 = 42
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    _ = v28
    return 9223372036854775807
}

func main() {
    var v1 uint64 = 576
    var v2 int16 = 32767
    var v3 string = "a"
    var v4 uint32 = 42
    for fi29 := 0; fi29 < 2; fi29++ {
        if (2147483647 >= (v4 | (v4 >> ((255) & 63)))) {
            var v30 int8 = 100
            v31 := []int{(fi29 % (((fi29 * fi29)) | 1)), fi29, 32768}
            _ = len(v31)
            _ = v30
            _ = v31
        } else {
            v32 := GenS4{F0: 1, F1: -622, F2: fi29}
            v3 = v3
            _ = v32
        }
        var v33 bool = true
        _ = fi29
        _ = v33
    }
    v34 := GenS5{F0: 4294967295}
    v35 := GenS1{F0: -2147483648, F1: 1}
    var v36 uint16 = 42
    for fi37 := 0; fi37 < 9; fi37++ {
        v38 := []int{fi37, (fi37 / (((fi37 & (^(fi37)))) | 1))}
        var v39 string = ("" + v3)
        v4 = v4
        _ = fi37
        _ = v38
        _ = v39
    }
    v2 = (v2 << ((18446744073709551615) & 63))
    v40 := []int{}
    _ = len(v40)
    v1 = v1
    v41 := GenS1{F0: -2147483648, F1: uint8(v36)}
    for fi42 := 0; fi42 < 17; fi42++ {
        var v43 int8 = 29
        v43 = (v43 * -128)
        if true {
            for fi44 := 0; fi44 < 7; fi44++ {
                var v45 uint = 2
                if len(v40) > 0 {
                    v40[0] = (fi42 >> (((v45 * v45)) & 63))
                }
                _ = fi44
                _ = v45
            }
            fmt.Println(fi42)
            _ = len(v40)
        } else {
            v46 := GenS5{F0: -102}
            _ = v46
        }
        _ = fi42
        _ = v43
    }
    for fi47 := 0; fi47 < 20; fi47++ {
        var v48 int32 = 0
        _ = len(v40)
        switch -371 {
        case 651:
            var v49 string = v3
            _ = v49
        default:
            if false {
                if len(v40) >= 2 {
                    v40 = v40[1:len(v40)]
                }
            }
        }
        _ = fi47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v41
}
