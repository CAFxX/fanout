package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 int32; F1 string; F2 uint16 }
type GenS7 struct { F0 uint16; F1 int32; F2 int16 }
type GenS8 struct { F0 int16; F1 uint64 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint64 }
func genFn0(p0 int, p1 uint) int8 {
    fmt.Println(p1, p1)
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := []int{(fi7 * -556)}
        switch fi7 {
        case 979:
            var v9 int = fi7
            _ = v9
        case 2147483647:
            var v10 bool = (-735 >= 227)
            _ = v10
        default:
            var v11 string = "foo bar"
            _ = v11
        }
        fmt.Println(p1, fi7, fi7)
        _ = fi7
        _ = v8
    }
    if true {
        v12 := GenS1{F0: 0, F1: p1, F2: 2}
        p1 -= p1
        _ = v12
    }
    v13 := []int{17, (p0 + p0)}
    _ = p0
    _ = p1
    _ = v13
    return 0
}
func genFn1(p0 int64, p1 uint8) uint {
    var v14 int32 = 1
    if false {
        if (false || true) {
            v14 += (v14 % ((v14) | 1))
            var v15 uint16 = 837
            _ = v15
        } else {
            if (true || true) {
                switch "\n\t" {
                case "foo bar":
                    v16 := GenS1{F0: (p0 + (p0 ^ (p0 >> ((4294967295) & 63)))), F1: 42, F2: p1}
                    _ = v16
                case "a":
                    fmt.Println(v14, v14, v14)
                case "hello":
                    var v17 int8 = genFn0(32767, 18446744073709551615)
                    _ = v17
                default:
                    if true {
                        v14 = (v14 / ((v14) | 1))
                        for fi18 := 0; fi18 < 15; fi18++ {
                            var v19 uint8 = p1
                            _ = fi18
                            _ = v19
                        }
                        v20 := []int{-1, 2147483647, 308}
                        _ = v20
                    } else {
                        p1 = 120
                        fmt.Println(p0, v14)
                        v21 := GenS2{F0: 359}
                        _ = v21
                    }
                }
            } else {
                fmt.Println(p1, p1, v14)
                var v22 int = 1
                p0 = (p0 * (-(p0)))
                _ = v22
            }
        }
        p0 -= p0
        switch "hello" {
        case "":
            fmt.Println(v14, v14, p0)
        case "a":
            for fi23 := 0; fi23 < 16; fi23++ {
                v14 -= (v14 >> ((18446744073709551615) & 63))
                var v24 uint64 = 4294967295
                if false {
                    v25 := []int{fi23, 72, -9223372036854775808, 32767}
                    p0 = p0
                    _ = v25
                } else {
                    p1 -= 127
                    var v26 int64 = (p0 ^ (p0 & p0))
                    v27 := []int{fi23}
                    _ = v26
                    _ = v27
                }
                _ = fi23
                _ = v24
            }
        default:
            var v28 uint32 = 529
            _ = v28
        }
    }
    if (!true) {
        switch "a" {
        case "foo bar":
            for fi29 := 0; fi29 < 9; fi29++ {
                var v30 uint8 = 2
                v31 := GenS1{F0: p0, F1: 0, F2: (v30 ^ p1)}
                p1 -= 0
                _ = fi29
                _ = v30
                _ = v31
            }
        default:
            v32 := []int{-1, 42, -2147483648}
            _ = v32
        }
        if (554 > 625) {
            if false {
                switch (("foo bar" + "") + (("hello" + "a") + ("\n\t" + "a"))) {
                case "hello":
                    fmt.Println(p1, p0)
                default:
                    v14 = (v14 / ((v14) | 1))
                }
                if true {
                    fmt.Println(v14, p1)
                    v33 := []int{}
                    _ = v33
                }
                fmt.Println(p0, p1)
            }
            var v34 uint32 = 2147483647
            _ = v34
        } else {
            var v35 int32 = v14
            if false {
                p0 += (p0 ^ 830)
                v35 = 763
            } else {
                switch 9223372036854775807 {
                case 0:
                    v36 := []int{885, 23, 1}
                    _ = v36
                case 128:
                    p0 = (p0 * 9223372036854775807)
                case -46:
                    var v37 float32 = ((1.5 / ((1.5 - 3.14159) + float32((float64((1.5 - (-2.25 * ((((1e10 / 1.5) - float32(0.1)) * float32((-2.25 / 1.5))) * -2.25)))) / 1.0)))) - (float32(0.1) * ((float32(0.0) + (3.14159 - 0.1)) - -2.25)))
                    _ = v37
                default:
                    if (false || (!false)) {
                        var v38 int = 32768
                        var v39 uint8 = 1
                        _ = v38
                        _ = v39
                    }
                }
                var v40 int8 = 127
                _ = v40
            }
            v14 += v14
            _ = v35
        }
    } else {
        var v41 uint64 = 584
        _ = v41
    }
    var v42 int = int(p1)
    _ = p0
    _ = p1
    _ = v14
    _ = v42
    return uint(p1)
}

func main() {
    var v1 uint32 = 851
    var v2 string = "\n\t"
    v3 := []int{}
    v4 := []int{}
    var v5 string = ("" + "")
    v6 := []uint32{v1}
    v1 += v1
    for fi43 := 0; fi43 < 7; fi43++ {
        if true {
            if (!(65535 < (v1 / ((v1) | 1)))) {
                v2 = ("a" + ("foo bar" + v5))
            } else {
                v5 = "\n\t"
                var v44 float64 = 1e10
                var v45 int8 = genFn0(fi43, 279)
                _ = v44
                _ = v45
            }
            v46 := GenS3{F0: 527}
            _ = v46
        }
        var v47 uint8 = 42
        _ = fi43
        _ = v47
    }
    var v48 uint64 = 1
    if len(v4) > 0 {
        v4[0] = -2147483648
    }
    v49 := GenS3{F0: -32768}
    if true {
        switch -408 {
        case -483:
            v50 := GenS3{F0: -1}
            _ = v50
        case 2147483647:
            if (478 != 740) {
                fmt.Println(v5, v2, v2)
                for fi51 := 0; fi51 < 14; fi51++ {
                    var v52 uint64 = (v48 ^ (v48 | v48))
                    _ = fi51
                    _ = v52
                }
                for fi53 := 0; fi53 < 18; fi53++ {
                    for fi54 := 0; fi54 < 2; fi54++ {
                        v1 += 2147483647
                        _ = fi54
                    }
                    _ = fi53
                }
            } else {
                var v55 int16 = 1000
                var v56 string = v2
                _ = v55
                _ = v56
            }
        }
    } else {
        v57 := GenS4{F0: 42, F1: 1}
        var v58 uint8 = 208
        _ = v57
        _ = v58
    }
    v3 = append(v3, 324)
    var v59 bool = (1 < 60)
    var v60 uint64 = (^(v48))
    if (v59 || (float32(3.14159) == 0.0)) {
        v6 = append(v6, v1)
    } else {
        for fi61 := 0; fi61 < 1; fi61++ {
            for fi62 := 0; fi62 < 7; fi62++ {
                v1 = (v1 & (v1 << ((65535) & 63)))
                v63 := GenS5{F0: v59, F1: 0}
                _ = fi62
                _ = v63
            }
            v64 := GenS6{F0: 222, F1: (((v2 + (v2 + ((v5 + v5) + ("\n\t" + v5)))) + ("foo bar" + (v2 + v2))) + v2), F2: 0}
            _ = fi61
            _ = v64
        }
        v65 := GenS7{F0: 0, F1: 80, F2: 32767}
        if len(v6) > 0 {
            v6[0] = v1
        }
        _ = v65
    }
    for fi66 := 0; fi66 < 20; fi66++ {
        var v67 uint64 = v60
        _ = fi66
        _ = v67
    }
    v68 := GenS8{F0: -1, F1: v60}
    v48 += v60
    for fi69 := 0; fi69 < 19; fi69++ {
        v4 = append(v4, fi69)
        if (0 >= 2) {
            fmt.Println(fi69, v2, v59)
            fmt.Println(v5, v1)
            v70 := GenS9{F0: 65535}
            _ = v70
        } else {
            v71 := GenS10{F0: v60}
            _ = v71
        }
        _ = fi69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v68
}
