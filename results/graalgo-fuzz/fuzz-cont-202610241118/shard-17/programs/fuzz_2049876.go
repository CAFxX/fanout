package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 string; F2 int }
type GenS4 struct { F0 int32; F1 bool; F2 bool }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 int8; F1 int64; F2 int64 }
type GenS8 struct { F0 uint16; F1 int32 }
func genFn0(p0 uint16, p1 uint32, p2 int32) uint8 {
    v7 := []int{9223372036854775807, 32768}
    switch "hello" {
    case "hello":
        var v8 string = ""
        _ = v8
    case "foo bar":
        if true {
            var v9 uint16 = 255
            _ = v9
        } else {
            p2 -= 1
            v10 := GenS1{F0: p0, F1: uint(p0)}
            for fi11 := 0; fi11 < 18; fi11++ {
                v12 := GenS1{F0: (p0 % (((p0 >> ((820) & 63))) | 1)), F1: uint(p1)}
                _ = fi11
                _ = v12
            }
            _ = v10
        }
    case "":
        if (false && false) {
            p2 = 0
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            for fi13 := 0; fi13 < 2; fi13++ {
                p0 = 281
                _ = fi13
            }
        }
    default:
        v14 := GenS2{F0: -2147483648}
        _ = v14
    }
    p1 -= (p1 % ((p1) | 1))
    if (false || (((0.1 * 1.5) / 3.14159) == 1e10)) {
        var v15 uint64 = 439
        v16 := GenS3{F0: 781, F1: (((("foo bar" + "foo bar") + "") + "a") + ("foo bar" + "hello")), F2: 9223372036854775807}
        _ = v15
        _ = v16
    }
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}
func genFn1(p0 int, p1 uint32, p2 int8) uint8 {
    var v17 uint = 65535
    v18 := []int{(p0 << ((v17) & 63)), p0, (p0 >> (((v17 & (v17 >> ((v17) & 63)))) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return 60
}

func main() {
    var v1 uint16 = 5
    var v2 uint32 = 2
    var v3 int8 = 0
    v4 := []uint32{v2, (v2 + (v2 / ((42) | 1))), v2, 42}
    v5 := make(map[uint32]string)
    v6 := []string{}
    var v19 int8 = v3
    fmt.Println(v19)
    if (571 <= -1) {
        var v20 int16 = 1
        v2 = v2
        _ = v20
    }
    if (!(("hello" + "hello") > (("\n\t" + "foo bar") + "foo bar"))) {
        var v21 int = 2147483647
        var v22 int16 = -788
        for fi23 := 0; fi23 < 7; fi23++ {
            v5[(v2 << ((577) & 63))] = ("hello" + (("\n\t" + "a") + "\n\t"))
            v5[v2] = "\n\t"
            switch "" {
            case "\n\t":
                var v24 string = "foo bar"
                _ = v24
            case "foo bar":
                for fi25 := 0; fi25 < 13; fi25++ {
                    var v26 float32 = -2.25
                    v6 = append(v6, "")
                    fmt.Println(fi25, v1, v19)
                    _ = fi25
                    _ = v26
                }
            case "hello":
                if true {
                    if ((-(v19)) <= 42) {
                        v27 := GenS4{F0: -1, F1: ((!false) || true), F2: false}
                        delete(v5, v2)
                        for fi28 := 0; fi28 < 7; fi28++ {
                            switch fi23 {
                            case 660:
                                var v29 float32 = 0.0
                                _ = v29
                            default:
                                v1 -= v1
                            }
                            _ = fi28
                        }
                        _ = v27
                    }
                }
            default:
                if false {
                    v22 = (v22 * v22)
                    fmt.Println(fi23, v22)
                }
            }
            _ = fi23
        }
        _ = v21
        _ = v22
    }
    switch (("\n\t" + ("" + "hello")) + "hello") {
    case "foo bar":
        for fi30 := 0; fi30 < 10; fi30++ {
            v31 := GenS5{F0: ("\n\t" + ""), F1: 737}
            v6 = append(v6, "hello")
            for fi32 := 0; fi32 < 2; fi32++ {
                if true {
                    var v33 uint32 = 701
                    var v34 uint = 748
                    _ = v33
                    _ = v34
                }
                _ = fi32
            }
            _ = fi30
            _ = v31
        }
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        v3 = v3
        v2 = (^(v2))
        fmt.Println(v3, v3, v1)
        _ = fi35
    }
    var v36 int64 = 0
    if (((float32(3.14159) + 1.5) - -2.25) == 0.1) {
        if false {
            v1 = 688
            for fi37 := 0; fi37 < 10; fi37++ {
                fmt.Println(v2)
                _ = fi37
            }
            var v38 uint64 = 866
            _ = v38
        }
        v36 = (^(v36))
    } else {
        v39 := GenS6{F0: 2147483647, F1: 42}
        switch -661 {
        case 890:
            switch "" {
            case "foo bar":
                if false {
                    for fi40 := 0; fi40 < 1; fi40++ {
                        var v41 uint64 = 42
                        _ = fi40
                        _ = v41
                    }
                    var v42 uint32 = v2
                    v19 = v3
                    _ = v42
                } else {
                    for fi43 := 0; fi43 < 5; fi43++ {
                        var v44 uint16 = 1
                        _ = fi43
                        _ = v44
                    }
                    var v45 int8 = v3
                    var v46 int64 = v36
                    _ = v45
                    _ = v46
                }
            case "\n\t":
                delete(v5, 38)
            case "hello":
                v1 -= 32767
            }
        case 128:
            for fi47 := 0; fi47 < 14; fi47++ {
                fmt.Println(v3, v2)
                v1 = v1
                _ = fi47
            }
        case 554:
            v2 -= uint32(v1)
        default:
            var v48 int64 = int64(v3)
            _ = v48
        }
        fmt.Println(v3, v36, v3)
        _ = v39
    }
    fmt.Println(v19, v3)
    v36 = (v36 << ((1) & 63))
    if len(v4) > 0 {
        v4[0] = v2
    }
    for fi49 := 0; fi49 < 8; fi49++ {
        for fi50 := 0; fi50 < 14; fi50++ {
            var v51 float64 = 0.0
            v52 := GenS7{F0: (v19 | v19), F1: (^(v36)), F2: -391}
            var v53 int64 = 42
            _ = fi50
            _ = v51
            _ = v52
            _ = v53
        }
        v2 = 42
        _ = fi49
    }
    var v54 uint8 = genFn0((v1 | v1), 65535, -677)
    var v55 float32 = ((3.14159 / 1e10) - 1e10)
    delete(v5, v2)
    v56 := GenS8{F0: v1, F1: 1}
    switch -2147483648 {
    case 32768:
        for fi57 := 0; fi57 < 18; fi57++ {
            var v58 uint16 = uint16(v54)
            _ = fi57
            _ = v58
        }
    }
    var v59 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v59
    _ = v6
}
