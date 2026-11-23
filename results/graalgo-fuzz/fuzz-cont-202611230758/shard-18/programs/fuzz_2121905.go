package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int16; F1 uint32; F2 int8 }
type GenS6 struct { F0 bool; F1 int64 }
func genFn0(p0 int) uint16 {
    for fi4 := 0; fi4 < 2; fi4++ {
        for fi5 := 0; fi5 < 19; fi5++ {
            p0 = 92
            _ = fi5
        }
        _ = fi4
    }
    p0 -= p0
    p0 -= (p0 + -2147483648)
    v6 := []int{(p0 << ((338) & 63))}
    _ = p0
    _ = v6
    return 8
}
func genFn1(p0 bool, p1 uint, p2 bool) int64 {
    var v7 float64 = (1.5 - float64(float32(-2.25)))
    p2 = (p0 || (p2 && p2))
    var v8 int = -9223372036854775808
    if p0 {
        for fi9 := 0; fi9 < 6; fi9++ {
            v7 = v7
            for fi10 := 0; fi10 < 14; fi10++ {
                if p0 {
                    v11 := GenS1{F0: 42, F1: 169}
                    var v12 int = fi10
                    _ = v11
                    _ = v12
                }
                fmt.Println(fi10)
                _ = fi10
            }
            _ = fi9
        }
    } else {
        for fi13 := 0; fi13 < 13; fi13++ {
            for fi14 := 0; fi14 < 11; fi14++ {
                fmt.Println(fi14)
                switch v8 {
                case 2147483647:
                    v15 := GenS1{F0: 42, F1: 0}
                    _ = v15
                case 602:
                    if (p2 || p0) {
                        if (false && (0 <= 4294967295)) {
                            var v16 float64 = v7
                            var v17 float32 = -2.25
                            p1 = p1
                            _ = v16
                            _ = v17
                        } else {
                            v18 := []int{fi13, (v8 << ((p1) & 63)), fi13, (fi13 & 32767)}
                            if len(v18) > 0 {
                                v18[0] = -1
                            }
                            _ = v18
                        }
                        var v19 uint8 = 42
                        switch "foo bar" {
                        case "a":
                            fmt.Println(v8, fi14)
                        case "\n\t":
                            fmt.Println(v19, fi14, p2)
                        case "":
                            if true {
                                var v20 uint8 = (v19 / (((^(v19))) | 1))
                                var v21 int = 42
                                _ = v20
                                _ = v21
                            }
                        default:
                            var v22 float32 = ((0.1 - -2.25) - 0.1)
                            _ = v22
                        }
                        _ = v19
                    }
                case -2147483648:
                    var v23 uint = p1
                    _ = v23
                default:
                    fmt.Println(v8, fi13)
                    fmt.Println(fi13, v7, p0)
                    var v24 string = ("" + "a")
                    _ = v24
                }
                _ = fi14
            }
            fmt.Println(p1)
            v7 = v7
            _ = fi13
        }
        for fi25 := 0; fi25 < 5; fi25++ {
            var v26 int16 = 0
            var v27 int64 = 42
            _ = fi25
            _ = v26
            _ = v27
        }
        var v28 float64 = v7
        _ = v28
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]int)
    var v3 int = 32767
    v29 := GenS1{F0: 128, F1: 127}
    for fi30 := 0; fi30 < 9; fi30++ {
        if (false && (v3 == (^(v3)))) {
            fmt.Println(fi30, v1)
            v31 := []int{fi30, (v3 >> ((18446744073709551615) & 63))}
            var v32 int16 = -32768
            _ = v31
            _ = v32
        } else {
            switch 0 {
            case -662:
                fmt.Println(v1, fi30)
            default:
                var v33 int32 = 894
                _ = v33
            }
            v34 := []int{}
            _ = v34
        }
        v35 := GenS1{F0: 1, F1: 0}
        for fi36 := 0; fi36 < 13; fi36++ {
            switch "foo bar" {
            case "":
                switch (fi36 - v3) {
                case -9223372036854775808:
                    fmt.Println(v1)
                case 9223372036854775807:
                    v1 = v1
                default:
                    fmt.Println(v3, v3)
                }
                switch v3 {
                case -2147483648:
                    v3 -= fi30
                }
                var v37 float64 = 3.14159
                _ = v37
            case "\n\t":
                v3 = fi30
                if (int8(fi36) != -1) {
                    var v38 uint32 = 42
                    v1 = v1
                    v39 := GenS2{F0: 127}
                    v38 = (^(v38))
                    var v40 int = fi30
                    _ = v38
                    _ = v39
                    _ = v40
                }
            }
            _ = fi36
        }
        _ = fi30
        _ = v35
    }
    v41 := []int{-9223372036854775808, 4294967295, -2147483648}
    switch "" {
    case "hello":
        var v42 float64 = 3.14159
        _ = v42
    }
    fmt.Println(v1)
    for fi43 := 0; fi43 < 16; fi43++ {
        v3 -= 32768
        if (true && ((1e10 > float32(float64(float32(((float64(3.14159) * v1) - (v1 / v1)))))) && true)) {
            var v44 uint64 = 3
            if (255 > 42) {
                v41 = append(v41, fi43)
                if len(v41) >= 2 {
                    v41 = v41[1:len(v41)]
                }
                v44 += (^(v44))
            }
            _ = v44
        } else {
            v3 -= 9223372036854775807
        }
        _ = fi43
    }
    if true {
        for fi45 := 0; fi45 < 20; fi45++ {
            var v46 int16 = 651
            _ = fi45
            _ = v46
        }
        if (false || (false || (2147483647 > -519))) {
            v47 := GenS3{F0: 127}
            switch v3 {
            case 0:
                var v48 int16 = 32767
                _ = v48
            }
            switch ((("\n\t" + "hello") + "") + ("hello" + ("" + "hello"))) {
            case "hello":
                v2["\n\t"] = (v3 ^ (v3 >> ((180) & 63)))
            case "a":
                var v49 float64 = v1
                _ = v49
            default:
                if (761 == 0) {
                    for fi50 := 0; fi50 < 4; fi50++ {
                        _ = len(v41)
                        _ = fi50
                    }
                    v3 = (^(v3))
                    v3 += v3
                } else {
                    v51 := GenS4{F0: 407, F1: 197, F2: 42}
                    _ = v51
                }
            }
            _ = v47
        }
        for fi52 := 0; fi52 < 9; fi52++ {
            for fi53 := 0; fi53 < 16; fi53++ {
                var v54 uint64 = 42
                v55 := GenS1{F0: 42, F1: 2}
                v3 -= 0
                v1 += v1
                if (((false || (768 <= -32768)) || true) || (42 < 65535)) {
                    var v56 bool = (!((!true) && true))
                    if v56 {
                        if (genFn1(v56, 4294967295, (((!((!true) || (444 == 65535))) || v56) || (fi52 == fi52))) < 42) {
                            v3 += (fi52 + v3)
                        } else {
                            _ = len(v41)
                            v1 += (0.0 / v1)
                        }
                        if (v56 || true) {
                            var v57 int64 = 1
                            _ = v57
                        }
                        var v58 uint16 = 799
                        _ = v58
                    } else {
                        v59 := GenS3{F0: 42}
                        if v56 {
                            v60 := GenS5{F0: 306, F1: 499, F2: 42}
                            var v61 uint32 = 592
                            _ = v60
                            _ = v61
                        }
                        _ = v59
                    }
                    _ = v56
                }
                _ = fi53
                _ = v54
                _ = v55
            }
            var v62 uint64 = 42
            _ = fi52
            _ = v62
        }
    }
    switch v3 {
    case 0:
        v63 := GenS5{F0: 1, F1: 126, F2: -90}
        _ = v63
    case 4294967295:
        for fi64 := 0; fi64 < 19; fi64++ {
            switch 2147483647 {
            case 4294967295:
                v1 -= v1
            }
            _ = fi64
        }
    case -9223372036854775808:
        if true {
            v41 = append(v41, v3)
            switch "" {
            case "":
                for fi65 := 0; fi65 < 6; fi65++ {
                    v66 := v2["a"]
                    if len(v41) > 0 {
                        v41[0] = fi65
                    }
                    if (66 != 42) {
                        v67 := GenS4{F0: 1000, F1: genFn0((^(v3))), F2: 0}
                        _ = v67
                    }
                    _ = fi65
                    _ = v66
                }
            }
            v68 := GenS6{F0: ("\n\t" != "\n\t"), F1: 14}
            _ = v68
        }
    default:
        fmt.Println(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v41
}
