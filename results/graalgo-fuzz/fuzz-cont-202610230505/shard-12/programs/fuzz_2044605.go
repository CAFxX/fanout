package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 int16 }
func genFn0(p0 int32, p1 int) int16 {
    if (!(2 <= 376)) {
        for fi6 := 0; fi6 < 14; fi6++ {
            p1 = (^(p1))
            p0 = (p0 << ((370) & 63))
            fmt.Println(fi6, p0)
            _ = fi6
        }
        p1 -= 42
        v7 := GenS1{F0: p0}
        _ = v7
    } else {
        var v8 int16 = 42
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 uint64 = 213
            v11 := GenS2{F0: 0, F1: 0}
            v8 += (v8 >> ((289) & 63))
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := GenS3{F0: 1}
        _ = v12
        _ = v8
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return -856
}
func genFn1(p0 bool, p1 uint64, p2 string) uint8 {
    if (127 != 1) {
        v13 := []int{}
        switch p2 {
        case "a":
            if (!((true && p0) && false)) {
                p1 -= 4294967295
            } else {
                fmt.Println(p2)
                if p0 {
                    for fi14 := 0; fi14 < 1; fi14++ {
                        var v15 int8 = int8(fi14)
                        for fi16 := 0; fi16 < 9; fi16++ {
                            fmt.Println(p2)
                            v17 := GenS2{F0: 2, F1: 893}
                            _ = fi16
                            _ = v17
                        }
                        switch p2 {
                        case "hello":
                            for fi18 := 0; fi18 < 2; fi18++ {
                                fmt.Println(fi18, v15, fi14)
                                _ = fi18
                            }
                        case "":
                            p1 += 4294967295
                        case "a":
                            var v19 uint = 811
                            _ = v19
                        }
                        _ = fi14
                        _ = v15
                    }
                    if ((!(false && ("\n\t" >= p2))) || p0) {
                        if (42 < 128) {
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            p1 -= p1
                        } else {
                            var v20 float32 = 1e10
                            _ = v20
                        }
                        v21 := GenS1{F0: 2147483647}
                        _ = v21
                    } else {
                        v22 := GenS4{F0: 42, F1: 1, F2: 32767}
                        _ = v22
                    }
                }
            }
        default:
            p0 = ((1.5 - 1.5) > 3.14159)
        }
        _ = v13
    }
    if true {
        switch p2 {
        case "":
            switch (p2 + (p2 + ("a" + "foo bar"))) {
            case "a":
                var v23 bool = p0
                _ = v23
            case "":
                p2 = p2
            case "foo bar":
                v24 := GenS5{F0: (p1 >> ((255) & 63)), F1: -272}
                _ = v24
            default:
                v25 := GenS1{F0: 1}
                _ = v25
            }
        }
        switch -9223372036854775808 {
        case 49:
            switch ((p2 + (p2 + p2)) + p2) {
            case "":
                p2 = "hello"
            case "\n\t":
                var v26 float32 = (float32((1.5 + (float64((-2.25 / 1.5)) + (-2.25 / 3.14159)))) * 1e10)
                _ = v26
            default:
                var v27 uint8 = 42
                _ = v27
            }
        }
        for fi28 := 0; fi28 < 8; fi28++ {
            for fi29 := 0; fi29 < 16; fi29++ {
                v30 := []int{(fi29 | -1), fi29, fi28}
                p1 = p1
                _ = fi29
                _ = v30
            }
            _ = fi28
        }
    }
    if (p0 && (-32768 > 667)) {
        var v31 uint64 = p1
        p2 = "hello"
        v32 := GenS6{F0: (v31 ^ 266), F1: 1, F2: 127}
        _ = v31
        _ = v32
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int = -661
    v2 := make(map[uint32]bool)
    var v3 int16 = 0
    v4 := make(map[string]bool)
    var v5 uint32 = 2
    switch ("\n\t" + "a") {
    case "foo bar":
        fmt.Println(v1, v5)
    default:
        v1 = v1
    }
    var v33 float64 = 0.1
    for fi34 := 0; fi34 < 12; fi34++ {
        for fi35 := 0; fi35 < 1; fi35++ {
            var v36 uint16 = 104
            v37 := GenS6{F0: 42, F1: 100, F2: 166}
            _ = fi35
            _ = v36
            _ = v37
        }
        fmt.Println(v3, v33, v3)
        for fi38 := 0; fi38 < 9; fi38++ {
            var v39 float64 = 0.0
            v40 := GenS4{F0: 32767, F1: 32767, F2: 874}
            v41 := GenS2{F0: 500, F1: (v3 + -162)}
            _ = fi38
            _ = v39
            _ = v40
            _ = v41
        }
        _ = fi34
    }
    v3 = v3
    fmt.Println(v3)
    if true {
        v1 -= (v1 >> ((255) & 63))
        var v42 uint64 = 148
        _ = v42
    } else {
        switch (v1 >> ((255) & 63)) {
        case 128:
            v43 := v4[("a" + "hello")]
            _ = v43
        case 2147483647:
            v1 -= (v1 * (v1 << ((255) & 63)))
        case -553:
            for fi44 := 0; fi44 < 1; fi44++ {
                var v45 uint = 18446744073709551615
                var v46 float64 = (v33 + 3.14159)
                _ = fi44
                _ = v45
                _ = v46
            }
        }
        var v47 int = v1
        for fi48 := 0; fi48 < 6; fi48++ {
            var v49 bool = false
            if (2 > 202) {
                v50 := GenS4{F0: 68, F1: 1000, F2: 699}
                v51 := GenS7{F0: 255, F1: uint64(v5)}
                var v52 int = -182
                _ = v50
                _ = v51
                _ = v52
            } else {
                if (!(false || v49)) {
                    var v53 string = "a"
                    v1 += v1
                    _ = v53
                } else {
                    _ = len(v2)
                    switch "\n\t" {
                    case "\n\t":
                        v54 := []int{}
                        _ = v54
                    case "":
                        var v55 uint32 = (v5 / ((995) | 1))
                        _ = v55
                    case "a":
                        var v56 int32 = int32(v3)
                        _ = v56
                    }
                }
                var v57 float64 = (-2.25 - (0.1 - (v33 + float64(((0.0 * 3.14159) - 1.5)))))
                v58 := []int{fi48}
                _ = v57
                _ = v58
            }
            v3 = (v3 / ((-32768) | 1))
            _ = fi48
            _ = v49
        }
        _ = v47
    }
    for fi59 := 0; fi59 < 10; fi59++ {
        if false {
            v60 := []int{}
            v5 = 65535
            _ = v60
        }
        _ = fi59
    }
    var v61 int64 = -9223372036854775808
    for fi62 := 0; fi62 < 17; fi62++ {
        _ = len(v4)
        if true {
            v5 -= (^(v5))
            if false {
                v3 -= v3
                v63 := GenS6{F0: 18446744073709551615, F1: 20, F2: 128}
                v61 = v61
                _ = v63
            }
        } else {
            v61 -= (v61 - 42)
            v1 += -2147483648
        }
        _ = fi62
    }
    for fi64 := 0; fi64 < 7; fi64++ {
        v65 := GenS8{F0: v3}
        v66 := GenS6{F0: 0, F1: int8(v3), F2: 68}
        v5 = v5
        _ = fi64
        _ = v65
        _ = v66
    }
    var v67 int16 = 1000
    _ = len(v4)
    if false {
        v5 = v5
    }
    var v68 string = ((("\n\t" + "a") + "") + ("hello" + (("hello" + "") + ("\n\t" + "foo bar"))))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v61
    _ = v67
    _ = v68
}
