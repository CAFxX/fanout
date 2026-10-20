package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint; F1 int16; F2 uint64 }
type GenS4 struct { F0 string; F1 string; F2 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 uint64 }
type GenS7 struct { F0 uint8 }
func genFn0(p0 int16, p1 string) int64 {
    for fi6 := 0; fi6 < 13; fi6++ {
        switch fi6 {
        case -856:
            v7 := []int{127, (fi6 / ((fi6) | 1)), fi6, fi6}
            _ = v7
        case 128:
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := GenS1{F0: 32767, F1: 1}
                p0 += p0
                _ = fi8
                _ = v9
            }
        case 32767:
            fmt.Println(fi6)
        default:
            if true {
                p1 = p1
            }
        }
        v10 := GenS2{F0: p0, F1: 2147483647}
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS3{F0: 2, F1: (p0 | (p0 % (((p0 & (p0 + p0))) | 1))), F2: 438}
            switch (fi6 - (fi11 / (((^(fi11))) | 1))) {
            case -422:
                if (423 == 2147483647) {
                    var v13 int16 = p0
                    var v14 bool = (42 < -1)
                    p0 = v13
                    _ = v13
                    _ = v14
                }
            case 128:
                if true {
                    p0 = p0
                }
            }
            _ = fi11
            _ = v12
        }
        _ = fi6
        _ = v10
    }
    p1 = "a"
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 string, p1 bool, p2 int16) int32 {
    switch p0 {
    case "a":
        p0 = p0
    case "\n\t":
        switch -9223372036854775808 {
        case -812:
            for fi15 := 0; fi15 < 7; fi15++ {
                v16 := GenS4{F0: ((p0 + p0) + p0), F1: p0, F2: 45}
                _ = fi15
                _ = v16
            }
        case -9223372036854775808:
            if true {
                p1 = p1
            }
        default:
            p1 = p1
        }
    case "hello":
        var v17 uint8 = 255
        _ = v17
    default:
        v18 := GenS5{F0: p2}
        _ = v18
    }
    v19 := []int{9223372036854775807, 0, -9223372036854775808, 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 0
}

func main() {
    var v1 bool = true
    var v2 float32 = float32(-2.25)
    var v3 int64 = 42
    var v4 uint16 = 2
    v5 := make(map[int]string)
    v20 := GenS3{F0: 1, F1: 1, F2: 18446744073709551615}
    var v21 int8 = -1
    _ = len(v5)
    switch "hello" {
    case "\n\t":
        v22 := GenS6{F0: -1, F1: 1}
        _ = v22
    default:
        if (!((^(v4)) >= (v4 << ((18446744073709551615) & 63)))) {
            switch "foo bar" {
            case "\n\t":
                var v23 int32 = 2147483647
                _ = v23
            case "hello":
                for fi24 := 0; fi24 < 15; fi24++ {
                    v25 := []int{int(v3), (fi24 / ((fi24) | 1)), (fi24 | (-(fi24))), fi24}
                    _ = fi24
                    _ = v25
                }
            case "":
                v26 := v5[-40]
                _ = v26
            }
        }
    }
    if (2 > 924) {
        v27 := []int{}
        _ = v27
    }
    v3 = (^(v3))
    v4 -= (^(v4))
    if (2147483647 > 42) {
        for fi28 := 0; fi28 < 15; fi28++ {
            var v29 int32 = 42
            _ = fi28
            _ = v29
        }
    } else {
        if false {
            var v30 uint16 = v4
            if v1 {
                switch ("\n\t" + "foo bar") {
                case "a":
                    v31 := []int{2147483647, 128}
                    _ = v31
                case "hello":
                    v32 := []int{32768, 127, int(v21)}
                    _ = v32
                }
                var v33 bool = false
                _ = v33
            } else {
                var v34 int8 = 127
                v21 = (v34 << ((18446744073709551615) & 63))
                _ = v34
            }
            var v35 int16 = 1
            _ = v30
            _ = v35
        }
    }
    if (2147483647 > 2147483647) {
        v36 := GenS6{F0: 1, F1: 42}
        v37 := []int{1}
        var v38 uint8 = 0
        _ = v36
        _ = v37
        _ = v38
    }
    v39 := v5[-798]
    if v1 {
        v2 = 0.1
        switch -1 {
        case 415:
            v3 = v3
        default:
            v3 = (v3 ^ 42)
        }
        var v40 uint64 = 2
        _ = v40
    } else {
        for fi41 := 0; fi41 < 20; fi41++ {
            fmt.Println(v1, v3)
            _ = fi41
        }
        if v1 {
            var v42 int64 = v3
            v43 := []int{128}
            _ = v42
            _ = v43
        } else {
            v2 -= v2
            v44 := []int{32768, 4294967295, 42}
            _ = v44
        }
        var v45 uint8 = 127
        _ = v45
    }
    var v46 uint64 = 2
    var v47 float64 = 0.0
    v5[-2147483648] = ((((v39 + "hello") + (v39 + v39)) + v39) + v39)
    var v48 uint16 = (v4 / (((v4 % ((517) | 1))) | 1))
    v2 = (float32(((((v47 + 1.5) + (v47 / v47)) / v47) / v47)) - v2)
    switch ("foo bar" + "foo bar") {
    case "hello":
        var v49 uint = 0
        _ = v49
    case "a":
        if v1 {
            for fi50 := 0; fi50 < 4; fi50++ {
                if v1 {
                    for fi51 := 0; fi51 < 4; fi51++ {
                        v52 := GenS7{F0: uint8(v46)}
                        var v53 float32 = (v2 * v2)
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                    for fi54 := 0; fi54 < 15; fi54++ {
                        fmt.Println(v1, v46)
                        _ = fi54
                    }
                    v55 := GenS7{F0: 128}
                    _ = v55
                } else {
                    for fi56 := 0; fi56 < 16; fi56++ {
                        v2 = float32((v47 - v47))
                        fmt.Println(v46, v39, v21)
                        _ = fi56
                    }
                    v57 := GenS3{F0: 684, F1: -1, F2: v46}
                    _ = v57
                }
                if v1 {
                    if (true || (v39 <= v39)) {
                        var v58 int16 = 1000
                        for fi59 := 0; fi59 < 15; fi59++ {
                            delete(v5, (fi50 - 113))
                            _ = fi59
                        }
                        v39 = v39
                        _ = v58
                    }
                }
                var v60 uint8 = 245
                _ = fi50
                _ = v60
            }
        } else {
            v2 += (0.1 * -2.25)
        }
    case "":
        v61 := []int{-149}
        _ = v61
    default:
        for fi62 := 0; fi62 < 19; fi62++ {
            if ((float64((v2 + v2)) + 3.14159) < (v47 + float64(v2))) {
                var v63 int32 = -2147483648
                _ = v63
            } else {
                var v64 uint64 = v46
                v4 -= 42
                if (!v1) {
                    var v65 int64 = 9223372036854775807
                    _ = v65
                } else {
                    fmt.Println(v47)
                    var v66 uint32 = 4294967295
                    fmt.Println(v4, v66)
                    _ = v66
                }
                _ = v64
            }
            fmt.Println(v4, fi62)
            v67 := GenS5{F0: 42}
            _ = fi62
            _ = v67
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v39
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
}
