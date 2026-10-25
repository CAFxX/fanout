package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 uint8; F1 int16; F2 int16 }
type GenS6 struct { F0 int64; F1 string }
func genFn0(p0 uint16, p1 int64) int8 {
    if false {
        v7 := GenS1{F0: p1, F1: -1, F2: p0}
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := []int{}
            var v10 float64 = -2.25
            _ = fi8
            _ = v10
            _ = v9
        }
        if true {
            if (1 > 1) {
                v11 := []int{4294967295}
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = v11
            } else {
                var v12 float32 = 1.5
                _ = v12
            }
        }
        _ = v7
    }
    switch 127 {
    case -2147483648:
        switch 42 {
        case 2147483647:
            var v13 int = 1
            _ = v13
        case 32768:
            v14 := GenS1{F0: 42, F1: -1, F2: (p0 - p0)}
            _ = v14
        }
    case 0:
        v15 := []int{-616}
        _ = v15
    default:
        v16 := GenS2{F0: p0, F1: 1}
        _ = v16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        fmt.Println(p0, p1, fi17)
        _ = fi17
    }
    p1 += (p1 ^ p1)
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = float64(float32(0.1))
    var v3 uint32 = 539
    var v4 uint16 = 755
    var v5 float32 = (0.0 + 3.14159)
    var v6 uint32 = v3
    v18 := []int{-9223372036854775808, 32767}
    v6 = (v3 << ((1) & 63))
    v6 -= (v3 * v3)
    v19 := GenS3{F0: (76 >= 0), F1: "a"}
    if (3.14159 > v2) {
        var v20 uint64 = 630
        var v21 int = -976
        var v22 uint8 = 21
        _ = v20
        _ = v21
        _ = v22
    }
    v23 := GenS2{F0: v4, F1: 810}
    if true {
        v24 := GenS2{F0: (v4 * (v4 % ((v4) | 1))), F1: 2147483647}
        for fi25 := 0; fi25 < 6; fi25++ {
            var v26 float32 = (v5 - float32(((3.14159 * v2) + v2)))
            v2 = float64((v5 - v5))
            _ = fi25
            _ = v26
        }
        switch -1 {
        case 2147483647:
            var v27 string = ("foo bar" + "")
            _ = v27
        case -9223372036854775808:
            for fi28 := 0; fi28 < 10; fi28++ {
                v2 = v2
                _ = fi28
            }
        }
        _ = v24
    } else {
        var v29 int32 = 42
        for fi30 := 0; fi30 < 6; fi30++ {
            if false {
                var v31 uint32 = (v3 / ((v3) | 1))
                var v32 int16 = 42
                v6 = v3
                _ = v31
                _ = v32
            } else {
                for fi33 := 0; fi33 < 4; fi33++ {
                    fmt.Println(v3, v4)
                    if (42 > 46) {
                        if ((42 == 42) || (v2 <= v2)) {
                            v3 = (v6 % (((v6 ^ (v3 | v6))) | 1))
                        } else {
                            var v34 bool = true
                            _ = v34
                        }
                        var v35 uint16 = 315
                        _ = v35
                    }
                    v29 = v29
                    _ = fi33
                }
            }
            var v36 uint32 = (^(v6))
            v37 := GenS3{F0: true, F1: ""}
            _ = fi30
            _ = v36
            _ = v37
        }
        _ = v29
    }
    if (!(127 > genFn0(1, -9223372036854775808))) {
        v38 := GenS3{F0: ((2 == 4294967295) || true), F1: "\n\t"}
        switch 1 {
        case 4294967295:
            for fi39 := 0; fi39 < 13; fi39++ {
                if (0.0 <= v2) {
                    v40 := GenS3{F0: ((v4 & (v4 ^ (^(v4)))) < v4), F1: ("\n\t" + ("hello" + "foo bar"))}
                    var v41 string = ("\n\t" + (("a" + ("foo bar" + "\n\t")) + "hello"))
                    fmt.Println(v3, v3)
                    _ = v40
                    _ = v41
                } else {
                    for fi42 := 0; fi42 < 19; fi42++ {
                        var v43 uint32 = 782
                        v44 := GenS4{F0: 127, F1: 42}
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                }
                var v45 int64 = 397
                switch -730 {
                case 4294967295:
                    for fi46 := 0; fi46 < 12; fi46++ {
                        v47 := GenS1{F0: v45, F1: 1, F2: 0}
                        var v48 float32 = (v5 * v5)
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                case 32767:
                    var v49 bool = false
                    _ = v49
                case 0:
                    v45 = (^(v45))
                }
                _ = fi39
                _ = v45
            }
        default:
            if (1 > 42) {
                var v50 int8 = genFn0((v4 / (((v4 / (((v4 * v4)) | 1))) | 1)), -9223372036854775808)
                for fi51 := 0; fi51 < 8; fi51++ {
                    fmt.Println(v3, v2)
                    _ = fi51
                }
                _ = v50
            }
        }
        v1["hello"] = 4294967295
        _ = v38
    }
    if (false && false) {
        for fi52 := 0; fi52 < 8; fi52++ {
            if true {
                for fi53 := 0; fi53 < 16; fi53++ {
                    v6 = v3
                    _ = fi53
                }
                for fi54 := 0; fi54 < 11; fi54++ {
                    fmt.Println(fi52, fi54, v2)
                    _ = fi54
                }
            }
            v5 = v5
            var v55 int8 = 127
            _ = fi52
            _ = v55
        }
        for fi56 := 0; fi56 < 10; fi56++ {
            var v57 bool = (!(42 == 65535))
            switch (fi56 >> ((18446744073709551615) & 63)) {
            case 2147483647:
                for fi58 := 0; fi58 < 3; fi58++ {
                    _ = len(v18)
                    v6 -= (v6 >> ((2) & 63))
                    v59 := GenS5{F0: 169, F1: -32768, F2: -1}
                    _ = fi58
                    _ = v59
                }
            case -703:
                v60 := GenS1{F0: 1, F1: 0, F2: v4}
                _ = v60
            default:
                if (v6 <= v6) {
                    _ = len(v18)
                    for fi61 := 0; fi61 < 11; fi61++ {
                        v57 = true
                        _ = fi61
                    }
                }
            }
            _ = fi56
            _ = v57
        }
    }
    for fi62 := 0; fi62 < 18; fi62++ {
        v63 := GenS3{F0: true, F1: ("\n\t" + "\n\t")}
        _ = fi62
        _ = v63
    }
    if false {
        v5 = v5
        _ = len(v1)
        for fi64 := 0; fi64 < 7; fi64++ {
            v4 = v4
            _ = fi64
        }
    }
    switch "\n\t" {
    case "":
        for fi65 := 0; fi65 < 14; fi65++ {
            for fi66 := 0; fi66 < 2; fi66++ {
                var v67 uint16 = (v4 << ((18446744073709551615) & 63))
                fmt.Println(v5, v6, v3)
                v68 := GenS1{F0: -1, F1: 81, F2: 255}
                _ = fi66
                _ = v67
                _ = v68
            }
            fmt.Println(fi65, v5)
            _ = fi65
        }
    default:
        v18 = append(v18, -904)
    }
    for fi69 := 0; fi69 < 18; fi69++ {
        if (!false) {
            for fi70 := 0; fi70 < 1; fi70++ {
                fmt.Println(v3, v3)
                _ = fi70
            }
            v71 := GenS6{F0: -858, F1: ("a" + "foo bar")}
            _ = v71
        }
        _ = fi69
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
