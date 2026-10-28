package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint32; F1 uint8; F2 string }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 uint64; F1 int64 }
type GenS7 struct { F0 int64; F1 uint16; F2 int32 }
func genFn0(p0 uint, p1 string, p2 bool) int8 {
    var v6 uint8 = 42
    var v7 int32 = 42
    for fi8 := 0; fi8 < 16; fi8++ {
        if (true || (((0.0 + (1.5 + 0.1)) == (3.14159 - (0.1 - ((((1e10 * (-2.25 * (1.5 - 1e10))) * 1.5) + 0.1) * float32((1.5 * 3.14159)))))) && true)) {
            switch p1 {
            case "":
                if p2 {
                    var v9 int32 = v7
                    v10 := GenS1{F0: (126 != 137), F1: -234}
                    _ = v10
                    _ = v9
                }
            default:
                var v11 int16 = 1
                _ = v11
            }
            switch p1 {
            case "foo bar":
                if (!(p1 == "a")) {
                    var v12 int16 = 1
                    _ = v12
                } else {
                    var v13 uint8 = v6
                    fmt.Println(p0, v7)
                    v14 := GenS2{F0: (^(v6))}
                    _ = v13
                    _ = v14
                }
            default:
                v15 := GenS1{F0: (p2 || (!(1000 != 1))), F1: fi8}
                _ = v15
            }
        } else {
            if (!p2) {
                for fi16 := 0; fi16 < 2; fi16++ {
                    if p2 {
                        var v17 float32 = -2.25
                        var v18 uint8 = 7
                        _ = v17
                        _ = v18
                    }
                    if p2 {
                        var v19 int32 = 0
                        v20 := []int{fi16, (fi8 + (fi16 / ((128) | 1))), fi16, 32768}
                        var v21 int32 = (v7 % (((v7 >> (((p0 - p0)) & 63))) | 1))
                        _ = v19
                        _ = v20
                        _ = v21
                    } else {
                        v22 := GenS1{F0: p2, F1: -9223372036854775808}
                        _ = v22
                    }
                    _ = fi16
                }
                p0 = p0
                for fi23 := 0; fi23 < 6; fi23++ {
                    var v24 uint64 = 0
                    _ = fi23
                    _ = v24
                }
            } else {
                for fi25 := 0; fi25 < 19; fi25++ {
                    v26 := []int{(fi8 | (fi25 | fi25)), fi25, fi25, fi25}
                    p1 = p1
                    var v27 uint = (p0 + 42)
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                fmt.Println(v6, fi8)
            }
            v28 := []int{fi8, fi8, fi8, fi8}
            _ = v28
        }
        if ("a" >= (p1 + p1)) {
            v29 := GenS1{F0: (p2 && p2), F1: (fi8 - 127)}
            p2 = (p2 || (-72 == 1))
            v30 := GenS3{F0: 43, F1: 4294967295}
            _ = v29
            _ = v30
        }
        _ = fi8
    }
    for fi31 := 0; fi31 < 1; fi31++ {
        v32 := GenS4{F0: 42, F1: 27, F2: ((p1 + (p1 + (p1 + ""))) + "foo bar")}
        switch (fi31 >> (((p0 * (^(p0)))) & 63)) {
        case -1:
            v33 := GenS5{F0: 4294967295, F1: 382}
            _ = v33
        case 32767:
            var v34 uint8 = (v6 - v6)
            _ = v34
        case -9223372036854775808:
            var v35 uint64 = 18446744073709551615
            _ = v35
        }
        _ = fi31
        _ = v32
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 72
}

func main() {
    v1 := []string{("foo bar" + ""), ("a" + "a"), ("foo bar" + "hello")}
    v2 := make(map[string]int)
    var v3 int64 = 9223372036854775807
    v4 := []int{32768, 127, -2147483648, -398}
    var v5 uint = 255
    if (400 == 2) {
        v36 := v2[""]
        for fi37 := 0; fi37 < 1; fi37++ {
            v38 := GenS6{F0: 1, F1: (v3 & (v3 % (((v3 % ((v3) | 1))) | 1)))}
            switch (fi37 / ((2147483647) | 1)) {
            case 32768:
                switch v36 {
                case 9223372036854775807:
                    v2["hello"] = (-(v36))
                case -792:
                    v5 = (^(v5))
                }
            case -2147483648:
                var v39 uint8 = 1
                _ = v39
            }
            v40 := v2[(("hello" + ("" + "a")) + "")]
            _ = fi37
            _ = v38
            _ = v40
        }
        _ = v36
    } else {
        var v41 uint32 = 4294967295
        _ = v41
    }
    var v42 uint8 = 33
    _ = len(v1)
    fmt.Println(v5)
    switch 0 {
    case -9223372036854775808:
        fmt.Println(v42, v5, v42)
    case 0:
        if ((290 < 0) && false) {
            for fi43 := 0; fi43 < 10; fi43++ {
                for fi44 := 0; fi44 < 14; fi44++ {
                    switch fi43 {
                    case -9223372036854775808:
                        v3 -= 9223372036854775807
                    }
                    v42 = v42
                    _ = fi44
                }
                _ = fi43
            }
            if (1 <= -1) {
                if (((^(v3)) == v3) && ((v3 >= (^(v3))) || true)) {
                    for fi45 := 0; fi45 < 10; fi45++ {
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                        v2["hello"] = fi45
                        _ = fi45
                    }
                    v5 = 450
                } else {
                    fmt.Println(v5, v3, v42)
                    var v46 int64 = v3
                    v3 = -9223372036854775808
                    _ = v46
                }
            }
            if false {
                v42 += (^(v42))
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                switch 486 {
                case 9223372036854775807:
                    v5 = 0
                case -1:
                    v47 := GenS1{F0: (!true), F1: -213}
                    _ = v47
                default:
                    v5 += v5
                }
            }
        }
    default:
        fmt.Println(v42, v42, v3)
    }
    v48 := GenS6{F0: 18446744073709551615, F1: (v3 << (((v5 ^ (v5 << (((^(v5))) & 63)))) & 63))}
    switch 9223372036854775807 {
    case -1:
        for fi49 := 0; fi49 < 20; fi49++ {
            _ = len(v2)
            _ = fi49
        }
    case 4294967295:
        v2[("" + "")] = 190
    case -2147483648:
        v5 += (v5 / ((v5) | 1))
    }
    v2["hello"] = 32767
    var v50 int16 = 0
    v42 = (v42 * 172)
    var v51 float32 = (3.14159 * 3.14159)
    v5 -= (v5 << (((v5 / (((^(v5))) | 1))) & 63))
    v3 = (v3 << ((v5) & 63))
    var v52 int32 = -286
    if true {
        var v53 bool = false
        v54 := GenS7{F0: int64(v5), F1: 42, F2: 0}
        v55 := GenS1{F0: ((!(0.0 == 1.5)) && (v52 <= (v52 >> ((v5) & 63)))), F1: 99}
        _ = v53
        _ = v54
        _ = v55
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v48
    _ = v5
    _ = v50
    _ = v51
    _ = v52
}
